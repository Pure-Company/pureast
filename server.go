package cgi

const (
	StateNew ConnState = iota
	StateActive
	StateIdle
	StateHijacked
	StateClosed
) // StateNew represents a new connection that is expected to
// send a request immediately. Connections begin at this
// state and then transition to either StateActive or
// StateClosed.
// StateClosed represents a closed connection.
// This is a terminal state. Hijacked connections do not
// transition to StateClosed.

func (s *Server) shouldConfigureHTTP2ForServe() bool {
	if s.TLSConfig == nil {
		return true
	}
	if s.protocols().UnencryptedHTTP2() {
		return true
	}
	return slices.Contains(s.TLSConfig.NextProtos, http2NextProtoTLS)
} // shouldConfigureHTTP2ForServe reports whether Server.Serve should configure
// automatic HTTP/2. (which sets up the s.TLSNextProto map)

func (s *Server) closeConnChan(c net.Conn, done chan<- struct{}) {
	c.Close()
	if done != nil {
		done <- struct{}{}
	}
} // closeConnChan is like closeConn, but takes an optional channel to receive a value
// when the goroutine closing c is done.

func NewServer(handler http.Handler) *Server {
	ts := NewUnstartedServer(handler)
	ts.Start()
	return ts
} // NewServer starts and returns a new [Server].
// The caller should call Close when finished, to shut it down.

type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
} // A Handler responds to an HTTP request.
//
// [Handler.ServeHTTP] should write reply headers and data to the [ResponseWriter]
// and then return. Returning signals that the request is finished; it
// is not valid to use the [ResponseWriter] or read from the
// [Request.Body] after or concurrently with the completion of the
// ServeHTTP call.
//
// Depending on the HTTP client software, HTTP protocol version, and
// any intermediaries between the client and the Go server, it may not
// be possible to read from the [Request.Body] after writing to the
// [ResponseWriter]. Cautious handlers should read the [Request.Body]
// first, and then reply.
//
// Except for reading the body, handlers should not modify the
// provided Request.
//
// If ServeHTTP panics, the server (the caller of ServeHTTP) assumes
// that the effect of the panic was isolated to the active request.
// It recovers the panic, logs a stack trace to the server error log,
// and either closes the network connection or sends an HTTP/2
// RST_STREAM, depending on the HTTP protocol. To abort a handler so
// the client sees an interrupted response but the server doesn't log
// an error, panic with the value [ErrAbortHandler].

type connectMethodKey struct {
	proxy, scheme, addr string
	onlyH1              bool
} // connectMethodKey is the map key version of connectMethod, with a
// stringified proxy URL (or the empty string) instead of a pointer to
// a URL.

type incomparable [0]func() // incomparable is a zero-width, non-comparable type. Adding it to a struct
// makes that struct also non-comparable, and generally doesn't add
// any size (as long as it's first).

func (s *Server) ListenAndServe() error {
	if s.shuttingDown() {
		return ErrServerClosed
	}
	addr := s.Addr
	if addr == "" {
		addr = ":http"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	return s.Serve(ln)
} // ListenAndServe listens on the TCP network address s.Addr and then
// calls [Serve] to handle requests on incoming connections.
// Accepted connections are configured to enable TCP keep-alives.
//
// If s.Addr is blank, ":http" is used.
//
// ListenAndServe always returns a non-nil error. After [Server.Shutdown] or [Server.Close],
// the returned error is [ErrServerClosed].

func (s *Server) closeIdleConns() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	quiescent := true
	for c := range // closeIdleConns closes all idle connections and reports whether the
	// server is quiescent.
	s.activeConn {
		st, unixSec := c.getState()
		if st == StateNew && unixSec < time.Now().Unix()-5 {
			st = StateIdle
		}
		if st != StateIdle || unixSec == 0 {
			quiescent = false
			continue
		}
		c.rwc.Close()
		delete(s.activeConn, c)
	}
	return quiescent
}

const http2NextProtoTLS = "h2"

func (s *Server) newConn(rwc net.Conn) *conn {
	c := &conn{server: s, rwc: rwc}
	if debugServerConnections {
		c.rwc = newLoggingConn("server", c.rwc)
	}
	return c
} // Create new connection from rwc.

func http2ConfigureServer(s *Server, conf *http2Server) error {
	panic(noHTTP2)
}

func newLoggingConn(baseName string, c net.Conn) net.Conn {
	uniqNameMu.Lock()
	defer uniqNameMu.Unlock()
	uniqNameNext[baseName]++
	return &loggingConn{name: fmt.Sprintf("%s-%d", baseName, uniqNameNext[baseName]), Conn: c}
}

type responseAndError struct {
	_   incomparable
	res *Response
	err error
} // responseAndError is how the goroutine reading from an HTTP/1 server
// communicates with the goroutine doing the RoundTrip.
// else use this response (see res method)

var testHookServerServe func(*Server, net.Listener) // used if non-nil

type Server struct {
	Addr                         string
	Handler                      Handler
	DisableGeneralOptionsHandler bool
	TLSConfig                    *tls.Config
	ReadTimeout                  time.Duration
	ReadHeaderTimeout            time.Duration
	WriteTimeout                 time.Duration
	IdleTimeout                  time.Duration
	MaxHeaderBytes               int
	TLSNextProto                 map[ // A Server defines parameters for running an HTTP server.
	// The zero value for Server is a valid configuration.
	// TLSNextProto optionally specifies a function to take over
	// ownership of the provided TLS connection when an ALPN
	// protocol upgrade has occurred. The map key is the protocol
	// name negotiated. The Handler argument should be used to
	// handle HTTP requests and will initialize the Request's TLS
	// and RemoteAddr if not already set. The connection is
	// automatically closed when the function returns.
	// If TLSNextProto is not nil, HTTP/2 support is not enabled
	// automatically.
	string]func(*Server, *tls.Conn, Handler)
	ConnState         func(net.Conn, ConnState)
	ErrorLog          *log.Logger
	BaseContext       func(net.Listener) context.Context
	ConnContext       func(ctx context.Context, c net.Conn) context.Context
	HTTP2             *HTTP2Config
	Protocols         *Protocols
	inShutdown        atomic.Bool
	disableKeepAlives atomic.Bool
	nextProtoOnce     sync.Once
	nextProtoErr      error
	mu                sync.Mutex
	listeners         map[ // ConnState specifies an optional callback function that is
	// called when a client connection changes state. See the
	// ConnState type and associated constants for details.
	// result of http2.ConfigureServer if used
	*net.Listener]struct{}
	activeConn    map[*conn]struct{}
	onShutdown    []func()
	listenerGroup sync.WaitGroup
}

func (s *Server) SetKeepAlivesEnabled(v bool) {
	if v {
		s.disableKeepAlives.Store(false)
		return
	}
	s.disableKeepAlives.Store(true)
	s.closeIdleConns()
} // SetKeepAlivesEnabled controls whether HTTP keep-alives are enabled.
// By default, keep-alives are always enabled. Only very
// resource-constrained environments or servers in the process of
// shutting down should disable them.

type loggingConn struct {
	name string
	net.Conn
} // loggingConn is used for debugging.

func NewUnstartedServer(handler http.Handler) *Server {
	return &Server{Listener: newLocalListener(), Config: &http.Server{Handler: handler}}
} // NewUnstartedServer returns a new [Server] but doesn't start it.
//
// After changing its configuration, the caller should call Start or
// StartTLS.
//
// The caller should call Close when finished, to shut it down.

type ResponseWriter interface {
	Header() Header
	Write([]byte) (// A ResponseWriter interface is used by an HTTP handler to
	// construct an HTTP response.
	//
	// A ResponseWriter may not be used after [Handler.ServeHTTP] has returned.
	// Write writes the data to the connection as part of an HTTP reply.
	//
	// If [ResponseWriter.WriteHeader] has not yet been called, Write calls
	// WriteHeader(http.StatusOK) before writing the data. If the Header
	// does not contain a Content-Type line, Write adds a Content-Type set
	// to the result of passing the initial 512 bytes of written data to
	// [DetectContentType]. Additionally, if the total size of all written
	// data is under a few KB and there are no Flush calls, the
	// Content-Length header is added automatically.
	//
	// Depending on the HTTP protocol version and the client, calling
	// Write or WriteHeader may prevent future reads on the
	// Request.Body. For HTTP/1.x requests, handlers should read any
	// needed request body data before writing the response. Once the
	// headers have been flushed (due to either an explicit Flusher.Flush
	// call or writing enough data to trigger a flush), the request body
	// may be unavailable. For HTTP/2 requests, the Go HTTP server permits
	// handlers to continue to read the request body while concurrently
	// writing the response. However, such behavior may not be supported
	// by all HTTP/2 clients. Handlers should read before writing if
	// possible to maximize compatibility.
	int, error)
	WriteHeader(statusCode int)
} // WriteHeader sends an HTTP response header with the provided
// status code.
//
// If WriteHeader is not called explicitly, the first call to Write
// will trigger an implicit WriteHeader(http.StatusOK).
// Thus explicit calls to WriteHeader are mainly used to
// send error codes or 1xx informational responses.
//
// The provided code must be a valid HTTP 1xx-5xx status code.
// Any number of 1xx headers may be written, followed by at most
// one 2xx-5xx header. 1xx headers are sent immediately, but 2xx-5xx
// headers may be buffered. Use the Flusher interface to send
// buffered data. The header map is cleared when 2xx-5xx headers are
// sent, but not with 1xx headers.
//
// The server will automatically send a 100 (Continue) header
// on the first read from the request body if the request has
// an "Expect: 100-continue" header.

func (s *Server) Certificate() *x509.Certificate {
	return s.certificate
} // Certificate returns the certificate used by the server, or nil if
// the server doesn't use TLS.

type handler string

type connOrError struct {
	pc     *persistConn
	err    error
	idleAt time.Time
}

func (s *Server) Serve(l net.Listener) error {
	if fn := testHookServerServe; fn != nil {
		fn(s, l)
	}
	origListener := l
	l = &onceCloseListener{Listener: l}
	defer l.Close()
	if err := s.setupHTTP2_Serve(); err != nil {
		return err
	}
	if !s.trackListener(&l, true) {
		return ErrServerClosed
	}
	defer s.trackListener(&l, false)
	baseCtx := context.Background()
	if s.BaseContext != nil {
		baseCtx = s.BaseContext(origListener)
		if baseCtx == nil {
			panic("BaseContext returned a nil context")
		}
	}
	var tempDelay time.Duration
	ctx := context.WithValue(baseCtx, ServerContextKey, s)
	for {
		rw, err := l.Accept()
		if err != nil {
			if s.shuttingDown() {
				return ErrServerClosed
			}
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				s.logf("http: Accept error: %v; retrying in %v", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return err
		}
		connCtx := ctx
		if cc := s.ConnContext; cc != nil {
			connCtx = cc(connCtx, rw)
			if connCtx == nil {
				panic("ConnContext returned nil")
			}
		}
		tempDelay = 0
		c := s.newConn(rw)
		c.setState(c.rwc, StateNew, runHooks)
		go c.serve(connCtx)
	}
} // Serve accepts incoming connections on the Listener l, creating a
// new service goroutine for each. The service goroutines read requests and
// then call s.Handler to reply to them.
//
// HTTP/2 support is only enabled if the Listener returns [*tls.Conn]
// connections and they were configured with "h2" in the TLS
// Config.NextProtos.
//
// Serve always returns a non-nil error and closes l.
// After [Server.Shutdown] or [Server.Close], the returned error is [ErrServerClosed].
// how long to sleep on accept failure

type http2Server struct{ NewWriteScheduler func() http2WriteScheduler }

type h2Transport interface{ CloseIdleConnections() } // h2Transport is the interface we expect to be able to call from
// net/http against an *http2.Transport that's either bundled into
// h2_bundle.go or supplied by the user via x/net/http2.
//
// We name it with the "h2" prefix to stay out of the "http2" prefix
// namespace used by x/tools/cmd/bundle for h2_bundle.go.

type onceCloseListener struct {
	net.Listener
	once     sync.Once
	closeErr error
} // onceCloseListener wraps a net.Listener, protecting it from
// multiple Close calls.

func (s *Server) idleTimeout() time.Duration {
	if s.IdleTimeout != 0 {
		return s.IdleTimeout
	}
	return s.ReadTimeout
}

const DefaultMaxHeaderBytes = 1 << 20 // DefaultMaxHeaderBytes is the maximum permitted size of the headers
// in an HTTP request.
// This can be overridden by setting [Server.MaxHeaderBytes].
// 1 MB

func (s *Server) onceSetNextProtoDefaults_Serve() {
	if s.shouldConfigureHTTP2ForServe() {
		s.onceSetNextProtoDefaults()
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.inShutdown.Store(true)
	s.mu.Lock()
	lnerr := s.closeListenersLocked()
	for _, f := range // Shutdown gracefully shuts down the server without interrupting any
	// active connections. Shutdown works by first closing all open
	// listeners, then closing all idle connections, and then waiting
	// indefinitely for connections to return to idle and then shut down.
	// If the provided context expires before the shutdown is complete,
	// Shutdown returns the context's error, otherwise it returns any
	// error returned from closing the [Server]'s underlying Listener(s).
	//
	// When Shutdown is called, [Serve], [ServeTLS], [ListenAndServe], and
	// [ListenAndServeTLS] immediately return [ErrServerClosed]. Make sure the
	// program doesn't exit and waits instead for Shutdown to return.
	//
	// Shutdown does not attempt to close nor wait for hijacked
	// connections such as WebSockets. The caller of Shutdown should
	// separately notify such long-lived connections of shutdown and wait
	// for them to close, if desired. See [Server.RegisterOnShutdown] for a way to
	// register shutdown notification functions.
	//
	// Once Shutdown has been called on a server, it may not be reused;
	// future calls to methods such as Serve will return ErrServerClosed.
	s.onShutdown {
		go f()
	}
	s.mu.Unlock()
	s.listenerGroup.Wait()
	pollIntervalBase := time.Millisecond
	nextPollInterval := func() time.Duration {
		interval := pollIntervalBase + time.Duration(rand.Intn(int(pollIntervalBase/10)))
		pollIntervalBase *= 2
		if pollIntervalBase > shutdownPollIntervalMax {
			pollIntervalBase = shutdownPollIntervalMax
		}
		return interval
	}
	timer := time.NewTimer(nextPollInterval())
	defer timer.Stop()
	for {
		if s.closeIdleConns() {
			return lnerr
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-timer.C:
			timer.Reset(nextPollInterval())
		}
	}
}

var (
	ServerContextKey    = &contextKey{"http-server"}
	LocalAddrContextKey = &contextKey{"local-addr"}
) // ServerContextKey is a context key. It can be used in HTTP
// handlers with Context.Value to access the server that
// started the handler. The associated value will be of
// type *Server.
// LocalAddrContextKey is a context key. It can be used in
// HTTP handlers with Context.Value to access the local
// address the connection arrived on.
// The associated value will be of type net.Addr.

type http2WriteScheduler any

type requestAndChan struct {
	_          incomparable
	treq       *transportRequest
	ch         chan responseAndError
	addedGzip  bool
	continueCh chan<- struct{}
	callerGone <-chan // unbuffered; always send in select on callerGone
	// Optional blocking chan for Expect: 100-continue (for send).
	// If the request has an "Expect: 100-continue" header and
	// the server responds 100 Continue, readLoop send a value
	// to writeLoop via this chan.
	struct{}
} // closed when roundTrip caller has returned

type HTTP2Config struct {
	MaxConcurrentStreams          int
	MaxDecoderHeaderTableSize     int
	MaxEncoderHeaderTableSize     int
	MaxReadFrameSize              int
	MaxReceiveBufferPerConnection int
	MaxReceiveBufferPerStream     int
	SendPingTimeout               time.Duration
	PingTimeout                   time.Duration
	WriteByteTimeout              time.Duration
	PermitProhibitedCipherSuites  bool
	CountError                    func(errType string)
} // HTTP2Config defines HTTP/2 configuration parameters common to
// both [Transport] and [Server].
// CountError, if non-nil, is called on HTTP/2 errors.
// It is intended to increment a metric for monitoring.
// The errType contains only lowercase letters, digits, and underscores
// (a-z, 0-9, _).

func (s *Server) trackConn(c *conn, add bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.activeConn == nil {
		s.activeConn = make(map[*conn]struct{})
	}
	if add {
		s.activeConn[c] = struct{}{}
	} else {
		delete(s.activeConn, c)
	}
}

type Protocols struct{ bits uint8 } // Protocols is a set of HTTP protocols.
// The zero value is an empty set of protocols.
//
// The supported protocols are:
//
//   - HTTP1 is the HTTP/1.0 and HTTP/1.1 protocols.
//     HTTP1 is supported on both unsecured TCP and secured TLS connections.
//
//   - HTTP2 is the HTTP/2 protcol over a TLS connection.
//
//   - UnencryptedHTTP2 is the HTTP/2 protocol over an unsecured TCP connection.

func (s *Server) setupHTTP2_Serve() error {
	s.nextProtoOnce.Do(s.onceSetNextProtoDefaults_Serve)
	return s.nextProtoErr
} // setupHTTP2_Serve is called from (*Server).Serve and conditionally
// configures HTTP/2 on s using a more conservative policy than
// setupHTTP2_ServeTLS because Serve is called after tls.Listen,
// and may be called concurrently. See shouldConfigureHTTP2ForServe.
//
// The tests named TestTransportAutomaticHTTP2* and
// TestConcurrentServerServe in server_test.go demonstrate some
// of the supported use cases and motivations.

var ErrServerClosed = errors.New("http: Server closed") // ErrServerClosed is returned by the [Server.Serve], [ServeTLS], [ListenAndServe],
// and [ListenAndServeTLS] methods after a call to [Server.Shutdown] or [Server.Close].

func (s *Server) tlsHandshakeTimeout() time.Duration {
	var ret time.Duration
	for _, v := range // tlsHandshakeTimeout returns the time limit permitted for the TLS
	// handshake, or zero for unlimited.
	//
	// It returns the minimum of any positive ReadHeaderTimeout,
	// ReadTimeout, or WriteTimeout.
	[...]time.Duration{s.ReadHeaderTimeout, s.ReadTimeout, s.WriteTimeout} {
		if v <= 0 {
			continue
		}
		if ret == 0 || v < ret {
			ret = v
		}
	}
	return ret
}

func (s *Server) setupHTTP2_ServeTLS() error {
	s.nextProtoOnce.Do(s.onceSetNextProtoDefaults)
	return s.nextProtoErr
} // setupHTTP2_ServeTLS conditionally configures HTTP/2 on
// s and reports whether there was an error setting it up. If it is
// not configured for policy reasons, nil is returned.

type RoundTripper interface {
	RoundTrip(*Request) (*Response, error)
} // RoundTripper is an interface representing the ability to execute a
// single HTTP transaction, obtaining the [Response] for a given [Request].
//
// A RoundTripper must be safe for concurrent use by multiple
// goroutines.
// RoundTrip executes a single HTTP transaction, returning
// a Response for the provided Request.
//
// RoundTrip should not attempt to interpret the response. In
// particular, RoundTrip must return err == nil if it obtained
// a response, regardless of the response's HTTP status code.
// A non-nil err should be reserved for failure to obtain a
// response. Similarly, RoundTrip should not attempt to
// handle higher-level protocol details such as redirects,
// authentication, or cookies.
//
// RoundTrip should not modify the request, except for
// consuming and closing the Request's Body. RoundTrip may
// read fields of the request in a separate goroutine. Callers
// should not mutate or reuse the request until the Response's
// Body has been closed.
//
// RoundTrip must always close the body, including on errors,
// but depending on the implementation may do so in a separate
// goroutine even after RoundTrip returns. This means that
// callers wanting to reuse the body for subsequent requests
// must arrange to wait for the Close call before doing so.
//
// The Request's URL and Header fields must be initialized.

type transportRequest struct {
	*Request
	extra  Header
	trace  *httptrace.ClientTrace
	ctx    context.Context
	cancel context.CancelCauseFunc
	mu     sync.Mutex
	err    error
} // transportRequest is a wrapper around a *Request that adds
// optional extra headers to write and stores any error to return
// from roundTrip.
// first setError value for mapRoundTripError to consider

func (s *Server) CloseClientConnections() {
	s.mu.Lock()
	nconn := len(s.conns)
	ch := make(chan struct{}, nconn)
	for c := range // CloseClientConnections closes any open HTTP connections to the test Server.
	s.conns {
		go s.closeConnChan(c, ch)
	}
	s.mu.Unlock()
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	for i := 0; i < nconn; i++ {
		select {
		case <-ch:
		case <-timer.C:
			return
		}
	}
}

type ConnState int // A ConnState represents the state of a client connection to a server.
// It's used by the optional [Server.ConnState] hook.

type pattern struct {
	str      string
	method   string
	host     string
	segments []segment// A pattern is something that can be matched against an HTTP request.
	// It has an optional method, an optional host, and a path.
	// The representation of a path differs from the surface syntax, which
	// simplifies most algorithms.
	//
	// Paths ending in '/' are represented with an anonymous "..." wildcard.
	// For example, the path "a/" is represented as a literal segment "a" followed
	// by a segment with multi==true.
	//
	// Paths ending in "{$}" are represented with the literal segment "/".
	// For example, the path "a/{$}" is represented as a literal segment "a" followed
	// by a literal segment "/".

	loc string
} // source location of registering call, for helpful messages

func (s *Server) onceSetNextProtoDefaults() {
	if omitBundledHTTP2 {
		return
	}
	p := s.protocols()
	if !p.HTTP2() && !p.UnencryptedHTTP2() {
		return
	}
	if http2server.Value() == "0" {
		http2server.IncNonDefault()
		return
	}
	if _, ok := s.TLSNextProto["h2"]; ok {
		return
	}
	conf := &http2Server{}
	s.nextProtoErr = http2ConfigureServer(s, conf)
} // onceSetNextProtoDefaults configures HTTP/2, if the user hasn't
// configured otherwise. (by setting s.TLSNextProto non-nil)
// It must only be called via s.nextProtoOnce (use s.setupHTTP2_*).

type contextKey struct{ name string } // contextKey is a value for use with context.WithValue. It's used as
// a pointer so it fits in an interface{} without allocation.

type wantConnQueue struct {
	head []*// A wantConnQueue is a queue of wantConns.
	// This is a queue, not a deque.
	// It is split into two stages - head[headPos:] and tail.
	// popFront is trivial (headPos++) on the first stage, and
	// pushBack is trivial (append) on the second stage.
	// If the first stage is empty, popFront can swap the
	// first and second stages to remedy the situation.
	//
	// This two-stage split is analogous to the use of two lists
	// in Okasaki's purely functional queue but without the
	// overhead of reversing the list when swapping stages.
	wantConn
	headPos int
	tail    []*wantConn
}

var http2server = godebug.New("http2server")

type Transport struct {
	idleMu    sync.Mutex
	closeIdle bool
	idleConn  map[ // Transport is an implementation of [RoundTripper] that supports HTTP,
	// HTTPS, and HTTP proxies (for either HTTP or HTTPS with CONNECT).
	//
	// By default, Transport caches connections for future re-use.
	// This may leave many open connections when accessing many hosts.
	// This behavior can be managed using [Transport.CloseIdleConnections] method
	// and the [Transport.MaxIdleConnsPerHost] and [Transport.DisableKeepAlives] fields.
	//
	// Transports should be reused instead of created as needed.
	// Transports are safe for concurrent use by multiple goroutines.
	//
	// A Transport is a low-level primitive for making HTTP and HTTPS requests.
	// For high-level functionality, such as cookies and redirects, see [Client].
	//
	// Transport uses HTTP/1.1 for HTTP URLs and either HTTP/1.1 or HTTP/2
	// for HTTPS URLs, depending on whether the server supports HTTP/2,
	// and how the Transport is configured. The [DefaultTransport] supports HTTP/2.
	// To explicitly enable HTTP/2 on a transport, set [Transport.Protocols].
	//
	// Responses with status codes in the 1xx range are either handled
	// automatically (100 expect-continue) or ignored. The one
	// exception is HTTP status code 101 (Switching Protocols), which is
	// considered a terminal status and returned by [Transport.RoundTrip]. To see the
	// ignored 1xx responses, use the httptrace trace package's
	// ClientTrace.Got1xxResponse.
	//
	// Transport only retries a request upon encountering a network error
	// if the connection has been already been used successfully and if the
	// request is idempotent and either has no body or has its [Request.GetBody]
	// defined. HTTP requests are considered idempotent if they have HTTP methods
	// GET, HEAD, OPTIONS, or TRACE; or if their [Header] map contains an
	// "Idempotency-Key" or "X-Idempotency-Key" entry. If the idempotency key
	// value is a zero-length slice, the request is treated as idempotent but the
	// header is not sent on the wire.
	// user has requested to close all idle conns
	connectMethodKey][]*persistConn
	idleConnWait map[ // most recently used at end
	connectMethodKey]wantConnQueue
	idleLRU     connLRU
	reqMu       sync.Mutex
	reqCanceler map[ // waiting getConns
	*Request]context.CancelCauseFunc
	altMu          sync.Mutex
	altProto       atomic.Value
	connsPerHostMu sync.Mutex
	connsPerHost   map[ // guards changing altProto only
	// of nil or map[string]RoundTripper, key is URI scheme
	connectMethodKey]int
	connsPerHostWait       map[connectMethodKey]wantConnQueue
	dialsInProgress        wantConnQueue
	Proxy                  func(*Request) (*url.URL, error)
	OnProxyConnectResponse func(ctx context.Context, proxyURL *url.URL, connectReq *Request, connectRes *Response) error
	DialContext            func(ctx context.Context, network, addr string) (net.Conn, error)
	Dial                   func(network, addr string) (net.Conn, error)
	DialTLSContext         func(ctx context.Context, network, addr string) (net.Conn, error)
	DialTLS                func(network, addr string) (net.Conn, error)
	TLSClientConfig        *tls.Config
	TLSHandshakeTimeout    time.Duration
	DisableKeepAlives      bool
	DisableCompression     bool
	MaxIdleConns           int
	MaxIdleConnsPerHost    int
	MaxConnsPerHost        int
	IdleConnTimeout        time.Duration
	ResponseHeaderTimeout  time.Duration
	ExpectContinueTimeout  time.Duration
	TLSNextProto           map[ // waiting getConns
	// TLSNextProto specifies how the Transport switches to an
	// alternate protocol (such as HTTP/2) after a TLS ALPN
	// protocol negotiation. If Transport dials a TLS connection
	// with a non-empty protocol name and TLSNextProto contains a
	// map entry for that key (such as "h2"), then the func is
	// called with the request's authority (such as "example.com"
	// or "example.com:1234") and the TLS connection. The function
	// must return a RoundTripper that then handles the request.
	// If TLSNextProto is not nil, HTTP/2 support is not enabled
	// automatically.
	string]func(authority string, c *tls.Conn) RoundTripper
	ProxyConnectHeader     Header
	GetProxyConnectHeader  func(ctx context.Context, proxyURL *url.URL, target string) (Header, error)
	MaxResponseHeaderBytes int64
	WriteBufferSize        int
	ReadBufferSize         int
	nextProtoOnce          sync.Once
	h2transport            h2Transport
	tlsNextProtoWasNil     bool
	ForceAttemptHTTP2      bool
	HTTP2                  *HTTP2Config
	Protocols              *Protocols
} // ProxyConnectHeader optionally specifies headers to send to
// proxies during CONNECT requests.
// To set the header dynamically, see GetProxyConnectHeader.
// Protocols is the set of protocols supported by the transport.
//
// If Protocols includes UnencryptedHTTP2 and does not include HTTP1,
// the transport will use unencrypted HTTP/2 for requests for http:// URLs.
//
// If Protocols is nil, the default is usually HTTP/1 only.
// If ForceAttemptHTTP2 is true, or if TLSNextProto contains an "h2" entry,
// the default is HTTP/1 and HTTP/2.

func (s *Server) shuttingDown() bool {
	return s.inShutdown.Load()
}

const (
	runHooks  = true
	skipHooks = false
)

func adjustNextProtos(nextProtos []string,// adjustNextProtos adds or removes "http/1.1" and "h2" entries from
// a tls.Config.NextProtos list, according to the set of protocols in protos.
protos Protocols) []string {
	nextProtos = slices.Clone(nextProtos)
	var have Protocols
	nextProtos = slices.DeleteFunc(nextProtos, func(s string) bool {
		switch s {
		case "http/1.1":
			if !protos.HTTP1() {
				return true
			}
			have.SetHTTP1(true)
		case "h2":
			if !protos.HTTP2() {
				return true
			}
			have.SetHTTP2(true)
		}
		return false
	})
	if protos.HTTP2() && !have.HTTP2() {
		nextProtos = append(nextProtos, "h2")
	}
	if protos.HTTP1() && !have.HTTP1() {
		nextProtos = append(nextProtos, "http/1.1")
	}
	return nextProtos
}

func (s *Server) ListenAndServeTLS(certFile, keyFile string) error {
	if s.shuttingDown() {
		return ErrServerClosed
	}
	addr := s.Addr
	if addr == "" {
		addr = ":https"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()
	return s.ServeTLS(ln, certFile, keyFile)
} // ListenAndServeTLS listens on the TCP network address s.Addr and
// then calls [ServeTLS] to handle requests on incoming TLS connections.
// Accepted connections are configured to enable TCP keep-alives.
//
// Filenames containing a certificate and matching private key for the
// server must be provided if neither the [Server]'s TLSConfig.Certificates
// nor TLSConfig.GetCertificate are populated. If the certificate is
// signed by a certificate authority, the certFile should be the
// concatenation of the server's certificate, any intermediates, and
// the CA's certificate.
//
// If s.Addr is blank, ":https" is used.
//
// ListenAndServeTLS always returns a non-nil error. After [Server.Shutdown] or
// [Server.Close], the returned error is [ErrServerClosed].

type wantConn struct {
	cm         connectMethod
	key        connectMethodKey
	beforeDial func()
	afterDial  func()
	mu         sync.Mutex
	ctx        context.Context
	cancelCtx  context.CancelFunc
	done       bool
	result     chan connOrError
} // A wantConn records state about a wanted connection
// (that is, an active call to getConn).
// The conn may be gotten by dialing or by finding an idle connection,
// or a cancellation may make the conn no longer wanted.
// These three options are racing against each other and use
// wantConn to coordinate and agree about the winning outcome.
// channel to deliver connection or error

var (
	uniqNameMu   sync.Mutex
	uniqNameNext = make(map[string]int)
)

const shutdownPollIntervalMax = 500 * time.Millisecond // shutdownPollIntervalMax is the max polling interval when checking
// quiescence during Server.Shutdown. Polling starts with a small
// interval and backs off to the max.
// Ideally we could find a solution that doesn't involve polling,
// but which also doesn't have a high runtime cost (and doesn't
// involve any contentious mutexes), but that is left as an
// exercise for the reader.

func (s *Server) closeConn(c net.Conn) {
	s.closeConnChan(c, nil)
} // closeConn closes c.
// s.mu must be held.

type Header map[ // A Header represents the key-value pairs in an HTTP header.
//
// The keys should be in canonical form, as returned by
// [CanonicalHeaderKey].
string][]string

const noHTTP2 = "no bundled HTTP/2" // should never see this

func (s *Server) wrap() {
	oldHook := s.Config.ConnState
	s.Config.ConnState = func(c net.Conn, cs http.ConnState) {
		s.mu.Lock()
		defer s.mu.Unlock()
		switch cs {
		case http.StateNew:
			if _, exists := s.conns[c]; exists {
				panic("invalid state transition")
			}
			if s.conns == nil {
				s.conns = make(map[ // wrap installs the connection state-tracking hook to know which
				// connections are idle.
				net.Conn]http.ConnState)
			}
			s.wg.Add(1)
			s.conns[c] = cs
			if s.closed {
				s.closeConn(c)
			}
		case http.StateActive:
			if oldState, ok := s.conns[c]; ok {
				if oldState != http.StateNew && oldState != http.StateIdle {
					panic("invalid state transition")
				}
				s.conns[c] = cs
			}
		case http.StateIdle:
			if oldState, ok := s.conns[c]; ok {
				if oldState != http.StateActive {
					panic("invalid state transition")
				}
				s.conns[c] = cs
			}
			if s.closed {
				s.closeConn(c)
			}
		case http.StateHijacked, http.StateClosed:
			if _, ok := s.conns[c]; ok {
				delete(s.conns, c)
				defer s.wg.Done()
			}
		}
		if oldHook != nil {
			oldHook(c, cs)
		}
	}
}

func (s *Server) logf(format string, args ...any) {
	if s.ErrorLog != nil {
		s.ErrorLog.Printf(format, args...)
	} else {
		log.Printf(format, args...)
	}
}

type Request struct {
	Method           string
	URL              *url.URL
	Proto            string
	ProtoMajor       int
	ProtoMinor       int
	Header           Header
	Body             io.ReadCloser
	GetBody          func() (io.ReadCloser, error)
	ContentLength    int64
	TransferEncoding []string// A Request represents an HTTP request received by a server
	// or to be sent by a client.
	//
	// The field semantics differ slightly between client and server
	// usage. In addition to the notes on the fields below, see the
	// documentation for [Request.Write] and [RoundTripper].
	// TransferEncoding lists the transfer encodings from outermost to
	// innermost. An empty list denotes the "identity" encoding.
	// TransferEncoding can usually be ignored; chunked encoding is
	// automatically added and removed as necessary when sending and
	// receiving requests.

	Close         bool
	Host          string
	Form          url.Values
	PostForm      url.Values
	MultipartForm *multipart.Form
	Trailer       Header
	RemoteAddr    string
	RequestURI    string
	TLS           *tls.ConnectionState
	Cancel        <-chan // Close indicates whether to close the connection after
	// replying to this request (for servers) or after sending this
	// request and reading its response (for clients).
	//
	// For server requests, the HTTP server handles this automatically
	// and this field is not needed by Handlers.
	//
	// For client requests, setting this field prevents re-use of
	// TCP connections between requests to the same hosts, as if
	// Transport.DisableKeepAlives were set.
	// Cancel is an optional channel whose closure indicates that the client
	// request should be regarded as canceled. Not all implementations of
	// RoundTripper may support Cancel.
	//
	// For server requests, this field is not applicable.
	//
	// Deprecated: Set the Request's context with NewRequestWithContext
	// instead. If a Request's Cancel field and context are both
	// set, it is undefined whether Cancel is respected.
	struct{}
	Response *Response
	Pattern  string
	ctx      context.Context
	pat      *pattern
	matches  []string// Response is the redirect response which caused this request
	// to be created. This field is only populated during client
	// redirects.
	// the pattern that matched

	otherValues map[ // values for the matching wildcards in pat
	string]string
} // for calls to SetPathValue that don't match a wildcard

var serveFlag string // When debugging a particular http server-based test,
// this flag lets you run
//
//	go test -run='^BrokenTest$' -httptest.serve=127.0.0.1:8000
//
// to start the broken server so you can interact with it manually.
// We only register this flag if it looks like the caller knows about it
// and is trying to use it as we don't want to pollute flags and this
// isn't really part of our API. Don't depend on this.

type connectMethod struct {
	_            incomparable
	proxyURL     *url.URL
	targetScheme string
	targetAddr   string
	onlyH1       bool
} // connectMethod is the map key (in its String form) for keeping persistent
// TCP connections alive for subsequent HTTP requests.
//
// A connect method may be of the following types:
//
//	connectMethod.key().String()      Description
//	------------------------------    -------------------------
//	|http|foo.com                     http directly to server, no proxy
//	|https|foo.com                    https directly to server, no proxy
//	|https,h1|foo.com                 https directly to server w/o HTTP/2, no proxy
//	http://proxy.com|https|foo.com    http to proxy, then CONNECT to foo.com
//	http://proxy.com|http             http to proxy, http to anywhere after that
//	socks5://proxy.com|http|foo.com   socks5 to proxy, then http to foo.com
//	socks5://proxy.com|https|foo.com  socks5 to proxy, then https to foo.com
//	https://proxy.com|https|foo.com   https to proxy, then CONNECT to foo.com
//	https://proxy.com|http            https to proxy, http to anywhere after that
// whether to disable HTTP/2 and force HTTP/1

func (s *Server) maxHeaderBytes() int {
	if s.MaxHeaderBytes > 0 {
		return s.MaxHeaderBytes
	}
	return DefaultMaxHeaderBytes
}

func (s *Server) readHeaderTimeout() time.Duration {
	if s.ReadHeaderTimeout != 0 {
		return s.ReadHeaderTimeout
	}
	return s.ReadTimeout
}

func (s *Server) trackListener(ln *net.Listener, add bool) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.listeners == nil {
		s.listeners = make(map[ // trackListener adds or removes a net.Listener to the set of tracked
		// listeners.
		//
		// We store a pointer to interface in the map set, in case the
		// net.Listener is not comparable. This is safe because we only call
		// trackListener via Serve and can track+defer untrack the same
		// pointer to local variable there. We never need to compare a
		// Listener from another caller.
		//
		// It reports whether the server is still up (not Shutdown or Closed).
		*net.Listener]struct{})
	}
	if add {
		if s.shuttingDown() {
			return false
		}
		s.listeners[ln] = struct{}{}
		s.listenerGroup.Add(1)
	} else {
		delete(s.listeners, ln)
		s.listenerGroup.Done()
	}
	return true
}

func (s *Server) initialReadLimitSize() int64 {
	return int64(s.maxHeaderBytes()) + 4096
}

var _ io.ReaderFrom = (*persistConnWriter)(nil)

type connReader struct {
	rwc     net.Conn
	mu      sync.Mutex
	conn    *conn
	hasByte bool
	byteBuf [1]byte
	cond    *sync.Cond
	inRead  bool
	aborted bool
	remain  int64
} // connReader is the io.Reader wrapper used by *conn. It combines a
// selectively-activated io.LimitedReader (to bound request header
// read sizes) with support for selectively keeping an io.Reader.Read
// call blocked in a background goroutine to wait for activity and
// trigger a CloseNotifier channel.
// After a Handler has hijacked the conn and exited, connReader behaves like a
// proxy for the net.Conn and the aforementioned behavior is bypassed.
// bytes remaining

func (s *Server) logCloseHangDebugInfo() {
	s.mu.Lock()
	defer s.mu.Unlock()
	var buf strings.Builder
	buf.WriteString("httptest.Server blocked in Close after 5 seconds, waiting for connections:\n")
	for c, st := range s.conns {
		fmt.Fprintf(&buf, "  %T %p %v in state %v\n", c, c, c.RemoteAddr(), st)
	}
	log.Print(buf.String())
}

func cloneTLSConfig(cfg *tls.Config) *tls.Config {
	if cfg == nil {
		return &tls.Config{}
	}
	return cfg.Clone()
} // cloneTLSConfig returns a shallow clone of cfg, or a new zero tls.Config if
// cfg is nil. This is safe to call even if cfg is in active use by a TLS
// client or server.
//
// cloneTLSConfig should be an internal detail,
// but widely used packages access it using linkname.
// Notable members of the hall of shame include:
//   - github.com/searKing/golang
//
// Do not remove or change the type signature.
// See go.dev/issue/67401.
//
//go:linkname cloneTLSConfig

type writeRequest struct {
	req        *transportRequest
	ch         chan<- error
	continueCh <-chan // A writeRequest is sent by the caller's goroutine to the
	// writeLoop's goroutine to write a request while the read loop
	// concurrently waits on both the write response and the server's
	// reply.
	// Optional blocking chan for Expect: 100-continue (for receive).
	// If not nil, writeLoop blocks sending request body until
	// it receives from this chan.
	struct{}
}

type Response struct {
	Status           string
	StatusCode       int
	Proto            string
	ProtoMajor       int
	ProtoMinor       int
	Header           Header
	Body             io.ReadCloser
	ContentLength    int64
	TransferEncoding []string// Response represents the response from an HTTP request.
	//
	// The [Client] and [Transport] return Responses from servers once
	// the response headers have been received. The response body
	// is streamed on demand as the Body field is read.
	// Contains transfer encodings from outer-most to inner-most. Value is
	// nil, means that "identity" encoding is used.

	Close        bool
	Uncompressed bool
	Trailer      Header
	Request      *Request
	TLS          *tls.ConnectionState
} // Close records whether the header directed that the connection be
// closed after reading Body. The value is advice for clients: neither
// ReadResponse nor Response.Write ever closes a connection.
// TLS contains information about the TLS connection on which the
// response was received. It is nil for unencrypted responses.
// The pointer is shared between responses and should not be
// modified.

type segment struct {
	s     string
	wild  bool
	multi bool
} // A segment is a pattern piece that matches one or more path segments, or
// a trailing slash.
//
// If wild is false, it matches a literal segment, or, if s == "/", a trailing slash.
// Examples:
//
//	"a" => segment{s: "a"}
//	"/{$}" => segment{s: "/"}
//
// If wild is true and multi is false, it matches a single path segment.
// Example:
//
//	"{x}" => segment{s: "x", wild: true}
//
// If both wild and multi are true, it matches all remaining path segments.
// Example:
//
//	"{rest...}" => segment{s: "rest", wild: true, multi: true}
// "..." wildcard

func (s *Server) ServeTLS(l net.Listener, certFile, keyFile string) error {
	if err := s.setupHTTP2_ServeTLS(); err != nil {
		return err
	}
	config := cloneTLSConfig(s.TLSConfig)
	config.NextProtos = adjustNextProtos(config.NextProtos, s.protocols())
	configHasCert := len(config.Certificates) > 0 || config.GetCertificate != nil || config.GetConfigForClient != nil
	if !configHasCert || certFile != "" || keyFile != "" {
		var err error
		config.Certificates = make([]tls.// ServeTLS accepts incoming connections on the Listener l, creating a
		// new service goroutine for each. The service goroutines perform TLS
		// setup and then read requests, calling s.Handler to reply to them.
		//
		// Files containing a certificate and matching private key for the
		// server must be provided if neither the [Server]'s
		// TLSConfig.Certificates, TLSConfig.GetCertificate nor
		// config.GetConfigForClient are populated.
		// If the certificate is signed by a certificate authority, the
		// certFile should be the concatenation of the server's certificate,
		// any intermediates, and the CA's certificate.
		//
		// ServeTLS always returns a non-nil error. After [Server.Shutdown] or [Server.Close], the
		// returned error is [ErrServerClosed].
		Certificate, 1)
		config.Certificates[0], err = tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			return err
		}
	}
	tlsListener := tls.NewListener(l, config)
	return s.Serve(tlsListener)
}

func (s *Server) doKeepAlives() bool {
	return !s.disableKeepAlives.Load() && !s.shuttingDown()
}

type connLRU struct {
	ll *list.List
	m  map[ // list.Element.Value type is of *persistConn
	*persistConn]*list.Element
}

func (s *Server) StartTLS() {
	if s.URL != "" {
		panic("Server already started")
	}
	if s.client == nil {
		s.client = &http.Client{}
	}
	cert, err := tls.X509KeyPair(testcert.LocalhostCert, testcert.LocalhostKey)
	if err != nil {
		panic(fmt.Sprintf("httptest: NewTLSServer: %v", err))
	}
	existingConfig := s.TLS
	if existingConfig != nil {
		s.TLS = existingConfig.Clone()
	} else {
		s.TLS = new(tls.Config)
	}
	if s.TLS.NextProtos == nil {
		nextProtos := []string{// StartTLS starts TLS on a server from NewUnstartedServer.
		"http/1.1"}
		if s.EnableHTTP2 {
			nextProtos = []string{"h2"}
		}
		s.TLS.NextProtos = nextProtos
	}
	if len(s.TLS.Certificates) == 0 {
		s.TLS.Certificates = []tls.Certificate{cert}
	}
	s.certificate, err = x509.ParseCertificate(s.TLS.Certificates[0].Certificate[0])
	if err != nil {
		panic(fmt.Sprintf("httptest: NewTLSServer: %v", err))
	}
	certpool := x509.NewCertPool()
	certpool.AddCert(s.certificate)
	s.client.Transport = &http.Transport{TLSClientConfig: &tls.Config{RootCAs: certpool}, ForceAttemptHTTP2: s.EnableHTTP2}
	s.Listener = tls.NewListener(s.Listener, s.TLS)
	s.URL = "https://" + s.Listener.Addr().String()
	s.wrap()
	s.goServe()
}

type conn struct {
	server     *Server
	cancelCtx  context.CancelFunc
	rwc        net.Conn
	remoteAddr string
	tlsState   *tls.ConnectionState
	werr       error
	r          *connReader
	bufr       *bufio.Reader
	bufw       *bufio.Writer
	lastMethod string
	curReq     atomic.Pointer[response]
	curState   atomic.Uint64
	mu         sync.Mutex
	hijackedv  bool
} // A conn represents the server side of an HTTP connection.
// hijackedv is whether this connection has been hijacked
// by a Handler with the Hijacker interface.
// It is guarded by mu.

func (s *Server) RegisterOnShutdown(f func()) {
	s.mu.Lock()
	s.onShutdown = append(s.onShutdown, f)
	s.mu.Unlock()
} // RegisterOnShutdown registers a function to call on [Server.Shutdown].
// This can be used to gracefully shutdown connections that have
// undergone ALPN protocol upgrade or that have been hijacked.
// This function should start protocol-specific graceful shutdown,
// but should not wait for shutdown to complete.

func (s *Server) protocols() Protocols {
	if s.Protocols != nil {
		return *s.Protocols
	}
	_, hasH2 := s.TLSNextProto["h2"]
	http2Disabled := s.TLSNextProto != nil && !hasH2
	if http2server.Value() == "0" && !hasH2 {
		http2Disabled = true
	}
	var p Protocols
	p.SetHTTP1(true)
	if !http2Disabled {
		p.SetHTTP2(true)
	}
	return p
}

func (s *Server) Start() {
	if s.URL != "" {
		panic("Server already started")
	}
	if s.client == nil {
		s.client = &http.Client{Transport: &http.Transport{}}
	}
	s.URL = "http://" + s.Listener.Addr().String()
	s.wrap()
	s.goServe()
	if serveFlag != "" {
		fmt.Fprintln(os.Stderr, "httptest: serving on", s.URL)
		select {}
	}
} // Start starts a server from NewUnstartedServer.

type persistConnWriter struct{ pc *persistConn } // persistConnWriter is the io.Writer written to by pc.bw.
// It accumulates the number of bytes written to the underlying conn,
// so the retry logic can determine whether any bytes made it across
// the wire.
// This is exactly 1 pointer field wide so it can go into an interface
// without allocation.

const debugServerConnections = false // debugServerConnections controls whether all server connections are wrapped
// with a verbose logging wrapper.

var omitBundledHTTP2 bool // omitBundledHTTP2 is set by omithttp2.go when the nethttpomithttp2
// build tag is set. That means h2_bundle.go isn't compiled in and we
// shouldn't try to use it.

type persistConn struct {
	alt                  RoundTripper
	t                    *Transport
	cacheKey             connectMethodKey
	conn                 net.Conn
	tlsState             *tls.ConnectionState
	br                   *bufio.Reader
	bw                   *bufio.Writer
	nwrite               int64
	reqch                chan requestAndChan
	writech              chan writeRequest
	closech              chan struct{}
	isProxy              bool
	sawEOF               bool
	readLimit            int64
	writeErrCh           chan error
	writeLoopDone        chan struct{}
	idleAt               time.Time
	idleTimer            *time.Timer
	mu                   sync.Mutex
	numExpectedResponses int
	closed               error
	canceledErr          error
	broken               bool
	reused               bool
	mutateHeaderFunc     func(Header)
} // persistConn wraps a connection, usually a persistent one
// (but may be used for non-keep-alive requests as well)
// mutateHeaderFunc is an optional func to modify extra
// headers on each outbound request before it's written. (the
// original Request given to RoundTrip is not modified)

func (s *Server) goServe() {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.Config.Serve(s.Listener)
	}()
}

func (s *Server) Close() error {
	s.inShutdown.Store(true)
	s.mu.Lock()
	defer s.mu.Unlock()
	err := s.closeListenersLocked()
	s.mu.Unlock()
	s.listenerGroup.Wait()
	s.mu.Lock()
	for c := range // Close immediately closes all active net.Listeners and any
	// connections in state [StateNew], [StateActive], or [StateIdle]. For a
	// graceful shutdown, use [Server.Shutdown].
	//
	// Close does not attempt to close (and does not even know about)
	// any hijacked connections, such as WebSockets.
	//
	// Close returns any error returned from closing the [Server]'s
	// underlying Listener(s).
	s.activeConn {
		c.rwc.Close()
		delete(s.activeConn, c)
	}
	return err
}

func NewTLSServer(handler http.Handler) *Server {
	ts := NewUnstartedServer(handler)
	ts.StartTLS()
	return ts
} // NewTLSServer starts and returns a new [Server] using TLS.
// The caller should call Close when finished, to shut it down.

func (s *Server) Client() *http.Client {
	return s.client
} // Client returns an HTTP client configured for making requests to the server.
// It is configured to trust the server's TLS test certificate and will
// close its idle connections on [Server.Close].
// Use Server.URL as the base URL to send requests to the server.

func (s *Server) closeListenersLocked() error {
	var err error
	for ln := range s.listeners {
		if cerr := (*ln).Close(); cerr != nil && err == nil {
			err = cerr
		}
	}
	return err
}
