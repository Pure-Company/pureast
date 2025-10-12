// Type definitions extracted by pureast

// Suitable for LLM context - contains only type structures

package main

import (
	"./a"
	"C"
	"a"
	"archive/tar"
	"archive/zip"
	"arena"
	"b"
	"bufio"
	"bytes"
	"c"
	"cgolife"
	"cgosotest"
	"cgostdio/stdio"
	"chans"
	"cmd/asm/internal/arch"
	"cmd/asm/internal/asm"
	"cmd/asm/internal/flags"
	"cmd/asm/internal/lex"
	"cmd/cgo/internal/test/gcc68255"
	"cmd/cgo/internal/test/issue23555a"
	"cmd/cgo/internal/test/issue23555b"
	"cmd/cgo/internal/test/issue26213"
	"cmd/cgo/internal/test/issue26430"
	"cmd/cgo/internal/test/issue26743"
	"cmd/cgo/internal/test/issue27340"
	"cmd/cgo/internal/test/issue29563"
	"cmd/cgo/internal/test/issue30527"
	"cmd/cgo/internal/test/issue41761a"
	"cmd/cgo/internal/test/issue43639"
	"cmd/cgo/internal/test/issue52611a"
	"cmd/cgo/internal/test/issue52611b"
	"cmd/cgo/internal/test/issue8756"
	"cmd/cgo/internal/test/issue8828"
	"cmd/cgo/internal/test/issue9026"
	"cmd/cgo/internal/test/issue9400"
	"cmd/cgo/internal/test/issue9510a"
	"cmd/cgo/internal/test/issue9510b"
	"cmd/cgo/internal/testsanitizers/testdata/asan_linkerx/p"
	"cmd/compile/internal/abi"
	"cmd/compile/internal/abt"
	"cmd/compile/internal/amd64"
	"cmd/compile/internal/arm"
	"cmd/compile/internal/arm64"
	"cmd/compile/internal/base"
	"cmd/compile/internal/bitvec"
	"cmd/compile/internal/compare"
	"cmd/compile/internal/coverage"
	"cmd/compile/internal/deadlocals"
	"cmd/compile/internal/devirtualize"
	"cmd/compile/internal/dwarfgen"
	"cmd/compile/internal/escape"
	"cmd/compile/internal/gc"
	"cmd/compile/internal/importer"
	"cmd/compile/internal/inline"
	"cmd/compile/internal/inline/inlheur"
	"cmd/compile/internal/inline/interleaved"
	"cmd/compile/internal/ir"
	"cmd/compile/internal/liveness"
	"cmd/compile/internal/logopt"
	"cmd/compile/internal/loong64"
	"cmd/compile/internal/loopvar"
	"cmd/compile/internal/loopvar/testdata/inlines/a"
	"cmd/compile/internal/loopvar/testdata/inlines/b"
	"cmd/compile/internal/loopvar/testdata/inlines/c"
	"cmd/compile/internal/mips"
	"cmd/compile/internal/mips64"
	"cmd/compile/internal/noder"
	"cmd/compile/internal/objw"
	"cmd/compile/internal/pgoir"
	"cmd/compile/internal/pkginit"
	"cmd/compile/internal/ppc64"
	"cmd/compile/internal/rangefunc"
	"cmd/compile/internal/reflectdata"
	"cmd/compile/internal/riscv64"
	"cmd/compile/internal/rttype"
	"cmd/compile/internal/s390x"
	"cmd/compile/internal/ssa"
	"cmd/compile/internal/ssagen"
	"cmd/compile/internal/staticdata"
	"cmd/compile/internal/staticinit"
	"cmd/compile/internal/syntax"
	"cmd/compile/internal/test/testdata/mysort"
	"cmd/compile/internal/typebits"
	"cmd/compile/internal/typecheck"
	"cmd/compile/internal/types"
	"cmd/compile/internal/types2"
	"cmd/compile/internal/walk"
	"cmd/compile/internal/wasm"
	"cmd/compile/internal/x86"
	"cmd/go/internal/auth"
	"cmd/go/internal/base"
	"cmd/go/internal/bug"
	"cmd/go/internal/cache"
	"cmd/go/internal/cacheprog"
	"cmd/go/internal/cfg"
	"cmd/go/internal/clean"
	"cmd/go/internal/cmdflag"
	"cmd/go/internal/doc"
	"cmd/go/internal/envcmd"
	"cmd/go/internal/fips140"
	"cmd/go/internal/fix"
	"cmd/go/internal/fmtcmd"
	"cmd/go/internal/fsys"
	"cmd/go/internal/generate"
	"cmd/go/internal/gover"
	"cmd/go/internal/help"
	"cmd/go/internal/imports"
	"cmd/go/internal/imports/testdata/test/child"
	"cmd/go/internal/list"
	"cmd/go/internal/load"
	"cmd/go/internal/lockedfile"
	"cmd/go/internal/lockedfile/internal/filelock"
	"cmd/go/internal/mmap"
	"cmd/go/internal/modcmd"
	"cmd/go/internal/modfetch"
	"cmd/go/internal/modfetch/codehost"
	"cmd/go/internal/modget"
	"cmd/go/internal/modindex"
	"cmd/go/internal/modinfo"
	"cmd/go/internal/modload"
	"cmd/go/internal/mvs"
	"cmd/go/internal/run"
	"cmd/go/internal/search"
	"cmd/go/internal/str"
	"cmd/go/internal/telemetrycmd"
	"cmd/go/internal/telemetrystats"
	"cmd/go/internal/test"
	"cmd/go/internal/test/internal/genflags"
	"cmd/go/internal/tool"
	"cmd/go/internal/toolchain"
	"cmd/go/internal/trace"
	"cmd/go/internal/vcs"
	"cmd/go/internal/vcweb"
	"cmd/go/internal/version"
	"cmd/go/internal/vet"
	"cmd/go/internal/web"
	"cmd/go/internal/web/intercept"
	"cmd/go/internal/work"
	"cmd/go/internal/workcmd"
	"cmd/internal/archive"
	"cmd/internal/bio"
	"cmd/internal/browser"
	"cmd/internal/buildid"
	"cmd/internal/codesign"
	"cmd/internal/cov"
	"cmd/internal/cov/covcmd"
	"cmd/internal/disasm"
	"cmd/internal/doc"
	"cmd/internal/dwarf"
	"cmd/internal/edit"
	"cmd/internal/gcprog"
	"cmd/internal/goobj"
	"cmd/internal/hash"
	"cmd/internal/macho"
	"cmd/internal/obj"
	"cmd/internal/obj/arm"
	"cmd/internal/obj/arm64"
	"cmd/internal/obj/loong64"
	"cmd/internal/obj/mips"
	"cmd/internal/obj/ppc64"
	"cmd/internal/obj/riscv"
	"cmd/internal/obj/s390x"
	"cmd/internal/obj/wasm"
	"cmd/internal/obj/x86"
	"cmd/internal/objabi"
	"cmd/internal/objfile"
	"cmd/internal/osinfo"
	"cmd/internal/par"
	"cmd/internal/pathcache"
	"cmd/internal/pgo"
	"cmd/internal/pkgpath"
	"cmd/internal/pkgpattern"
	"cmd/internal/quoted"
	"cmd/internal/robustio"
	"cmd/internal/script"
	"cmd/internal/src"
	"cmd/internal/sys"
	"cmd/internal/telemetry"
	"cmd/internal/telemetry/counter"
	"cmd/internal/test2json"
	"cmd/link/internal/amd64"
	"cmd/link/internal/arm"
	"cmd/link/internal/arm64"
	"cmd/link/internal/benchmark"
	"cmd/link/internal/ld"
	"cmd/link/internal/ld/testdata/issue25459/a"
	"cmd/link/internal/ld/testdata/issue26237/b.dir"
	"cmd/link/internal/ld/testdata/issue32233/lib"
	"cmd/link/internal/loadelf"
	"cmd/link/internal/loader"
	"cmd/link/internal/loadmacho"
	"cmd/link/internal/loadpe"
	"cmd/link/internal/loadxcoff"
	"cmd/link/internal/loong64"
	"cmd/link/internal/mips"
	"cmd/link/internal/mips64"
	"cmd/link/internal/ppc64"
	"cmd/link/internal/riscv64"
	"cmd/link/internal/s390x"
	"cmd/link/internal/sym"
	"cmd/link/internal/wasm"
	"cmd/link/internal/x86"
	"cmd/link/testdata/dynimportvar/asm"
	"cmd/link/testdata/linkname/p"
	"cmp"
	"compress/flate"
	"compress/gzip"
	"compress/zlib"
	"container/heap"
	"container/list"
	"context"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/tls"
	"crypto/x509"
	"d"
	"debug/buildinfo"
	"debug/dwarf"
	"debug/elf"
	"debug/gosym"
	"debug/macho"
	"debug/pe"
	"debug/plan9obj"
	"dep"
	"e"
	"embed"
	"encoding/base64"
	"encoding/binary"
	"encoding/gob"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"encoding/xml"
	"errors"
	"example.com/pgo/devirtualize/mult.pkg"
	"f"
	"flag"
	"fmt"
	"g"
	"github.com/google/pprof/driver"
	"github.com/google/pprof/internal/binutils"
	"github.com/google/pprof/internal/driver"
	"github.com/google/pprof/internal/elfexec"
	"github.com/google/pprof/internal/graph"
	"github.com/google/pprof/internal/measurement"
	"github.com/google/pprof/internal/plugin"
	"github.com/google/pprof/internal/report"
	"github.com/google/pprof/internal/symbolizer"
	"github.com/google/pprof/internal/symbolz"
	"github.com/google/pprof/internal/transport"
	"github.com/google/pprof/profile"
	"github.com/google/pprof/third_party/svgpan"
	"github.com/ianlancetaylor/demangle"
	"go/ast"
	"go/build"
	"go/build/constraint"
	"go/constant"
	"go/doc"
	"go/format"
	"go/importer"
	"go/parser"
	"go/printer"
	"go/scanner"
	"go/token"
	"go/types"
	"go/version"
	"golang.org/x/arch/arm/armasm"
	"golang.org/x/arch/arm64/arm64asm"
	"golang.org/x/arch/loong64/loong64asm"
	"golang.org/x/arch/ppc64/ppc64asm"
	"golang.org/x/arch/riscv64/riscv64asm"
	"golang.org/x/arch/s390x/s390xasm"
	"golang.org/x/arch/x86/x86asm"
	"golang.org/x/mod/internal/lazyregexp"
	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
	"golang.org/x/mod/semver"
	"golang.org/x/mod/sumdb"
	"golang.org/x/mod/sumdb/dirhash"
	"golang.org/x/mod/sumdb/note"
	"golang.org/x/mod/sumdb/tlog"
	"golang.org/x/mod/zip"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
	"golang.org/x/sys/plan9"
	"golang.org/x/sys/unix"
	"golang.org/x/sys/windows"
	"golang.org/x/telemetry"
	"golang.org/x/telemetry/counter"
	"golang.org/x/telemetry/internal/config"
	"golang.org/x/telemetry/internal/configstore"
	"golang.org/x/telemetry/internal/counter"
	"golang.org/x/telemetry/internal/crashmonitor"
	"golang.org/x/telemetry/internal/mmap"
	"golang.org/x/telemetry/internal/telemetry"
	"golang.org/x/telemetry/internal/upload"
	"golang.org/x/term"
	"golang.org/x/text/cases"
	"golang.org/x/text/internal"
	"golang.org/x/text/internal/language"
	"golang.org/x/text/internal/language/compact"
	"golang.org/x/text/internal/tag"
	"golang.org/x/text/language"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"golang.org/x/tools/cmd/bisect"
	"golang.org/x/tools/cover"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/internal/analysisflags"
	"golang.org/x/tools/go/analysis/passes/appends"
	"golang.org/x/tools/go/analysis/passes/asmdecl"
	"golang.org/x/tools/go/analysis/passes/assign"
	"golang.org/x/tools/go/analysis/passes/atomic"
	"golang.org/x/tools/go/analysis/passes/bools"
	"golang.org/x/tools/go/analysis/passes/buildtag"
	"golang.org/x/tools/go/analysis/passes/cgocall"
	"golang.org/x/tools/go/analysis/passes/composite"
	"golang.org/x/tools/go/analysis/passes/copylock"
	"golang.org/x/tools/go/analysis/passes/ctrlflow"
	"golang.org/x/tools/go/analysis/passes/defers"
	"golang.org/x/tools/go/analysis/passes/directive"
	"golang.org/x/tools/go/analysis/passes/errorsas"
	"golang.org/x/tools/go/analysis/passes/framepointer"
	"golang.org/x/tools/go/analysis/passes/hostport"
	"golang.org/x/tools/go/analysis/passes/httpresponse"
	"golang.org/x/tools/go/analysis/passes/ifaceassert"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/passes/internal/analysisutil"
	"golang.org/x/tools/go/analysis/passes/loopclosure"
	"golang.org/x/tools/go/analysis/passes/lostcancel"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shift"
	"golang.org/x/tools/go/analysis/passes/sigchanyzer"
	"golang.org/x/tools/go/analysis/passes/slog"
	"golang.org/x/tools/go/analysis/passes/stdmethods"
	"golang.org/x/tools/go/analysis/passes/stdversion"
	"golang.org/x/tools/go/analysis/passes/stringintconv"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"golang.org/x/tools/go/analysis/passes/testinggoroutine"
	"golang.org/x/tools/go/analysis/passes/tests"
	"golang.org/x/tools/go/analysis/passes/timeformat"
	"golang.org/x/tools/go/analysis/passes/unmarshal"
	"golang.org/x/tools/go/analysis/passes/unreachable"
	"golang.org/x/tools/go/analysis/passes/unsafeptr"
	"golang.org/x/tools/go/analysis/passes/unusedresult"
	"golang.org/x/tools/go/analysis/passes/waitgroup"
	"golang.org/x/tools/go/analysis/unitchecker"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/ast/edge"
	"golang.org/x/tools/go/ast/inspector"
	"golang.org/x/tools/go/cfg"
	"golang.org/x/tools/go/types/objectpath"
	"golang.org/x/tools/go/types/typeutil"
	"golang.org/x/tools/internal/aliases"
	"golang.org/x/tools/internal/analysisinternal"
	"golang.org/x/tools/internal/analysisinternal/typeindex"
	"golang.org/x/tools/internal/astutil"
	"golang.org/x/tools/internal/bisect"
	"golang.org/x/tools/internal/facts"
	"golang.org/x/tools/internal/fmtstr"
	"golang.org/x/tools/internal/stdlib"
	"golang.org/x/tools/internal/typeparams"
	"golang.org/x/tools/internal/typesinternal"
	"golang.org/x/tools/internal/typesinternal/typeindex"
	"golang.org/x/tools/internal/versions"
	"h"
	"hash"
	"hash/crc32"
	"hash/fnv"
	"hash/maphash"
	"html"
	"html/template"
	"import1"
	"import2"
	"import3"
	"import4"
	"indirect"
	"internal/abi"
	"internal/asan"
	"internal/bisect"
	"internal/buildcfg"
	"internal/cfg"
	"internal/coverage"
	"internal/coverage/calloc"
	"internal/coverage/cformat"
	"internal/coverage/cmerge"
	"internal/coverage/decodecounter"
	"internal/coverage/decodemeta"
	"internal/coverage/encodecounter"
	"internal/coverage/encodemeta"
	"internal/coverage/pods"
	"internal/coverage/slicewriter"
	"internal/diff"
	"internal/exportdata"
	"internal/goarch"
	"internal/godebug"
	"internal/godebugs"
	"internal/goroot"
	"internal/gover"
	"internal/goversion"
	"internal/lazyregexp"
	"internal/lazytemplate"
	"internal/pkgbits"
	"internal/platform"
	"internal/profile"
	"internal/race"
	"internal/singleflight"
	"internal/syscall/unix"
	"internal/syscall/windows"
	"internal/sysinfo"
	"internal/syslist"
	"internal/testenv"
	"internal/trace"
	"internal/trace/raw"
	"internal/trace/tracev2"
	"internal/trace/traceviewer"
	"internal/trace/traceviewer/format"
	"internal/trace/version"
	"internal/txtar"
	"internal/types/errors"
	"internal/unsafeheader"
	"internal/xcoff"
	"io"
	"io/fs"
	"iter"
	"log"
	"log/slog"
	"maps"
	"math"
	"math/big"
	"math/bits"
	"math/rand"
	"mime"
	"net"
	"net/http"
	"net/http/cgi"
	"net/http/httptest"
	"net/http/httputil"
	"net/http/pprof"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"os/user"
	"p2"
	"path"
	"path/filepath"
	"plugin"
	"prog/dep"
	"reflect"
	"regexp"
	"rsc.io/markdown"
	"runtime"
	"runtime/cgo"
	"runtime/debug"
	"runtime/metrics"
	"runtime/pprof"
	"runtime/trace"
	"slices"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"syscall/js"
	"testcarchive/p"
	"testcshared/p"
	"testing"
	"testplugin/common"
	"testplugin/iface_i"
	"testplugin/issue18676/dynamodbstreamsevt"
	"testplugin/issue44956/base"
	"testplugin/issue53989/p"
	"testplugin/method2/p"
	"testplugin/method3/p"
	"testshared/dep2"
	"testshared/dep3"
	"testshared/depBase"
	"testshared/depBaseInternal"
	"testshared/explicit"
	"testshared/gcdata/p"
	"testshared/globallib"
	"testshared/iface_a"
	"testshared/iface_b"
	"testshared/iface_i"
	"testshared/implicit"
	"testshared/issue39777/b"
	"testshared/issue44031/a"
	"testshared/issue44031/b"
	"testshared/issue47837/a"
	"text/scanner"
	"text/tabwriter"
	"text/template"
	"time"
	"unicode"
	"unicode/utf16"
	"unicode/utf8"
	"unique"
	"unsafe"
)

type Interface interface{ N([]byte) }

type algo struct{ indrt func(dep.Interface) }

type Time struct{}

type S struct {
	Public     *int
	private    *int
	PublicTime Time
} // Deprecated: use PublicTime.

type URL struct{} // Deprecated: use URI.

type EmbedURLPtr struct{ *URL }

type S2 struct {
	S
	Extra bool
} // Deprecated: use T.

type Namer interface{ Name() string }

type I interface {
	Namer
	ptwo.Twoer
	Set(name string, balance int64)
	Get(string) int64
	GetNamed(string) (balance int64)
	private()
} // Deprecated: use GetNamed.

type Public interface {
	X()
	Y()
}

type Private interface {
	X()
	y()
} // Deprecated: Use Unexported.

type Error interface {
	error
	Temporary() bool
}

type s struct{}

type Codec struct{ Func func(x int, y int) (z int) }

type SI struct{ I int }

type T struct{ common }

type B struct{ common }

type common struct{ i int }

type TPtrUnexported struct{ *common }

type TPtrExported struct{ *Embedded }

type Embedded struct{}

type EmbedSelector struct{ Time }

type ByteStruct struct {
	B byte
	R rune
}

type Twoer interface{ PackageTwoMeth() }

type ThirdBase struct{}

type Pair[T1 interface{ M() }, T2 ~int] struct {
	f1 T1
	f2 T2
}

type Arch struct {
	*obj.LinkArch
	Instructions map[ // Arch wraps the link architecture object with more architecture-specific information.
	// Map of instruction names to enumeration.
	string]obj.As
	Register map[ // Map of register names to enumeration.
	string]int16
	RegisterPrefix map[ // Table of register prefix names. These are things like R for R(0) and SPR for SPR(268).
	string]bool
	RegisterNumber func(string, int16) (int16, bool)
	IsJump         func(word string) bool
} // RegisterNumber converts R(10) into arm.REG_R10.
// Instruction is a jump.

type Parser struct {
	lex        lex.TokenReader
	lineNum    int
	errorLine  int
	errorCount int
	sawCode    bool
	pc         int64
	input      []lex.// Line number in source file.
	// virtual PC; count of Progs; doesn't advance for GLOBL or DATA.
	Token
	inputPos      int
	pendingLabels []string
	labels        map[ // Labels to attach to next instruction.
	string]*obj.Prog
	toPatch     []Patch
	addr        []obj.Addr
	arch        *arch.Arch
	ctxt        *obj.Link
	firstProg   *obj.Prog
	lastProg    *obj.Prog
	dataAddr    map[string]int64
	isJump      bool
	allowABI    bool
	pkgPrefix   string
	errorWriter io.Writer
} // Most recent address for DATA for this symbol.
// Prefix to add to local symbols.

type Patch struct {
	addr  *obj.Addr
	label string
}

type Input struct {
	Stack
	includes []string// Input is the main input: a stack of readers and some macro definitions.
	// It also handles #include processing (by pushing onto the input stack)
	// and parses and instantiates macro definitions.

	beginningOfLine bool
	ifdefStack      []bool
	macros          map[string]*Macro
	text            string
	peek            bool
	peekToken       ScanToken
	peekText        string
} // Text of last token returned by Next.

type TokenReader interface {
	Next() ScanToken
	Text() string
	File() string
	Base() *src.PosBase
	SetBase(*src.PosBase)
	Line() int
	Col() int
	Close()
} // A TokenReader is like a reader, but returns lex tokens of type Token. It also can tell you what
// the text of the most recently returned token is, and where it was found.
// The underlying scanner elides all spaces except newline, so the input looks like a stream of
// Tokens; original spacing is lost but we don't need it.
// Close does any teardown required.

type Token struct {
	ScanToken
	text string
} // A Token is a scan token plus its string value.
// A macro is stored as a sequence of Tokens with spaces stripped.

type Macro struct {
	name string
	args []string// A Macro represents the definition of a #defined macro.
	// The #define name.

	tokens []Token// Formal arguments.

} // Body of macro.

type Slice struct {
	tokens []Token// A Slice reads from a slice of Tokens.

	base *src.PosBase
	line int
	pos  int
}

type Stack struct {
	tr []TokenReader// A Stack is a stack of TokenReaders. As the top TokenReader hits EOF,
	// it resumes reading the next one down.
}

type Tokenizer struct {
	tok  ScanToken
	s    *scanner.Scanner
	base *src.PosBase
	line int
	file *os.File
} // A Tokenizer is a simple wrapping of text/scanner.Scanner, configured
// for our purposes and made a TokenReader. It forms the lowest level,
// turning text from readers into tokens.
// If non-nil, file descriptor to close.

type typeConv struct {
	m map[ // A typeConv is a translator from dwarf types to Go types
	// with equivalent memory layout.
	// Cache of already-translated or in-progress types.
	string]*Type
	ptrs map[ // Map from types to incomplete pointers to those types.
	string][]*Type
	ptrKeys []dwarf.// Keys of ptrs in insertion order (deterministic worklist)
	// ptrKeys contains exactly the keys in ptrs.
	Type
	getTypeIDs map[ // Type names X for which there exists an XGetTypeID function with type func() CFTypeID.
	string]bool
	incompleteStructs map[ // incompleteStructs contains C structs that should be marked Incomplete.
	string]bool
	bool                                   ast.Expr
	byte                                   ast.Expr
	int8, int16, int32, int64              ast.Expr
	uint8, uint16, uint32, uint64, uintptr ast.Expr
	float32, float64                       ast.Expr
	complex64, complex128                  ast.Expr
	void                                   ast.Expr
	string                                 ast.Expr
	goVoid                                 ast.Expr
	goVoidPtr                              ast.Expr
	ptrSize                                int64
	intSize                                int64
} // Predeclared types.
// unsafe.Pointer or *byte

type GoCallback struct{}

type testPair struct {
	Name      string
	Got, Want interface{}
}

type Context struct{ ctx *C.struct_ibv_context }

type AsyncEvent struct{ event C.struct_ibv_async_event } // issue 1222

type type52542[T ~*C.float] struct{}

type data49633 struct{ msg string }

type ts struct {
	tv *C.SV
} // ERROR HERE

type A struct {
	a *_Ctype_struct_a
} // ERROR HERE

type dwarfer interface{ DWARF() (*dwarf.Data, error) }

type C struct{}

type Event struct{}

type Foo struct {
	Bar string `json:"Bar@baz,omitempty"`
}

type sameNameReusedInPlugins struct{ X string }

type sameNameHolder struct{ F *sameNameReusedInPlugins }

type Y struct{ X *X }

type X struct{ Y Y }

type Any struct {
	s string
	b int64
}

type Dep2 struct{ depBase.Dep }

type Dep3 struct {
	dep  depBase.Dep
	dep2 dep2.Dep2
}

type HasProg struct{ array [1024]*byte }

type Dep struct{ X int }

type ATypeWithALoooooongName interface{ M() }

type i interface{ m() } // test that unexported method is correctly marked

type ImplA struct{}

type Package struct {
	PackageName string
	PackagePath string
	PtrSize     int64
	IntSize     int64
	GccOptions  []string// A Package collects information about the package we're going to write.
	// name of package

	GccIsClang bool
	LdFlags    []string
	Written    map[ // #cgo LDFLAGS
	string]bool
	Name    map[string]*Name
	ExpFunc []*// accumulated Name from Files
	ExpFunc
	Decl []ast.// accumulated ExpFunc from Files
	Decl
	GoFiles  []string
	GccFiles []string// list of Go files

	Preamble string
	typedefs map[ // list of gcc output files
	// collected preamble for _cgo_export.h
	string]bool
	typedefList []typedefInfo// type names that appear in the types of the objects we're interested in

	noCallbacks map[string]bool
	noEscapes   map[ // C function names with #cgo nocallback directive
	string]bool
} // C function names with #cgo noescape directive

type typedefInfo struct {
	typedef string
	pos     token.Pos
} // A typedefInfo is an element on Package.typedefList: a typedef name
// and the position where it was required.

type Call struct {
	Call     *ast.CallExpr
	Deferred bool
	Done     bool
} // A Call refers to a call of a C.xxx function in the AST.

type Ref struct {
	Name    *Name
	Expr    *ast.Expr
	Context astContext
	Done    bool
} // A Ref refers to an expression of the form C.xxx in the AST.

type Name struct {
	Go       string
	Mangle   string
	C        string
	Define   string
	Kind     string
	Type     *Type
	FuncType *FuncType
	AddError bool
	Const    string
} // A Name collects information about C.xxx.
// constant definition

type ExpFunc struct {
	Func    *ast.FuncDecl
	ExpName string
	Doc     string
} // An ExpFunc is an exported function, callable from C.
// Such functions are identified in the Go input file
// by doc comments containing the line //export ExpName
// name to use from C

type TypeRepr struct {
	Repr       string
	FormatArgs []interface// A TypeRepr contains the string representation of a type.
	{

	}
}

type Type struct {
	Size       int64
	Align      int64
	C          *TypeRepr
	Go         ast.Expr
	EnumValues map[ // A Type collects information about a type in both the C and Go worlds.
	string]int64
	Typedef    string
	BadPointer bool
} // this pointer type should be represented as a uintptr (deprecated)

type ABIParamResultInfo struct {
	inparams []ABIParamAssignment// ABIParamResultInfo stores the results of processing a given
	// function type to compute stack layout and register assignments. For
	// each input and output parameter we capture whether the param was
	// register-assigned (and to which register(s)) or the stack offset
	// for the param if is not going to be passed in registers according
	// to the rules in the Go internal ABI specification (1.17).

	outparams []ABIParamAssignment// Includes receiver for method calls.  Does NOT include hidden closure pointer.

	offsetToSpillArea int64
	spillAreaSize     int64
	inRegistersUsed   int
	outRegistersUsed  int
	config            *ABIConfig
} // to enable String() method

type ABIParamAssignment struct {
	Type      *types.Type
	Name      *ir.Name
	Registers []RegIndex// ABIParamAssignment holds information about how a specific param or
	// result will be passed: in registers (in which case 'Registers' is
	// populated) or on the stack (in which case 'Offset' is set to a
	// non-negative stack offset). The values in 'Registers' are indices
	// (as described above), not architected registers.

	offset int32
}

type RegAmounts struct {
	intRegs   int
	floatRegs int
} // RegAmounts holds a specified number of integer/float registers.

type ABIConfig struct {
	offsetForLocals int64
	regAmounts      RegAmounts
	which           obj.ABI
} // ABIConfig captures the number of registers made available
// by the ABI rules for parameter passing and result returning.
// e.g., obj.(*Link).Arch.FixedFrameSize -- extra linkage information on some architectures.

type assignState struct {
	rTotal      RegAmounts
	rUsed       RegAmounts
	stackOffset int64
	spillOffset int64
} // assignState holds intermediate state during the register assigning process
// for a given function signature.
// current spill offset

type node32 struct {
	left, right *node32
	data        interface{}
	key         int32
	height_     int8
} // node32 is the internal tree node data type
// Standard conventions hold for left = smaller, right = larger

type iterator struct{ parents []*node32 }

type Iterator struct{ it iterator }

type DebugFlags struct {
	AlignHot              int    `help:"enable hot block alignment (currently requires -pgo)" concurrent:"ok"`
	Append                int    `help:"print information about append compilation"`
	Checkptr              int    `help:"instrument unsafe pointer conversions\n0: instrumentation disabled\n1: conversions involving unsafe.Pointer are instrumented\n2: conversions to unsafe.Pointer force heap allocation" concurrent:"ok"`
	Closure               int    `help:"print information about closure compilation"`
	Defer                 int    `help:"print information about defer compilation"`
	DisableNil            int    `help:"disable nil checks" concurrent:"ok"`
	DumpInlFuncProps      string `help:"dump function properties from inl heuristics to specified file"`
	DumpInlCallSiteScores int    `help:"dump scored callsites during inlining"`
	InlScoreAdj           string `help:"set inliner score adjustments (ex: -d=inlscoreadj=panicPathAdj:10/passConstToNestedIfAdj:-90)"`
	InlBudgetSlack        int    `help:"amount to expand the initial inline budget when new inliner enabled. Defaults to 80 if option not set." concurrent:"ok"`
	DumpPtrs              int    `help:"show Node pointers values in dump output"`
	DwarfInl              int    `help:"print information about DWARF inlined function creation"`
	EscapeMutationsCalls  int    `help:"print extra escape analysis diagnostics about mutations and calls" concurrent:"ok"`
	EscapeDebug           int    `help:"print information about escape analysis and resulting optimizations" concurrent:"ok"`
	Export                int    `help:"print export data"`
	FIPSHash              string `help:"hash value for FIPS debugging" concurrent:"ok"`
	Fmahash               string `help:"hash value for use in debugging platform-dependent multiply-add use" concurrent:"ok"`
	GCAdjust              int    `help:"log adjustments to GOGC" concurrent:"ok"`
	GCCheck               int    `help:"check heap/gc use by compiler" concurrent:"ok"`
	GCProg                int    `help:"print dump of GC programs"`
	Gossahash             string `help:"hash value for use in debugging the compiler"`
	InlFuncsWithClosures  int    `help:"allow functions with closures to be inlined" concurrent:"ok"`
	InlStaticInit         int    `help:"allow static initialization of inlined calls" concurrent:"ok"`
	Libfuzzer             int    `help:"enable coverage instrumentation for libfuzzer"`
	LiteralAllocHash      string `help:"hash value for use in debugging literal allocation optimizations" concurrent:"ok"`
	LoopVar               int    `help:"shared (0, default), 1 (private loop variables), 2, private + log"`
	LoopVarHash           string `help:"for debugging changes in loop behavior. Overrides experiment and loopvar flag."`
	LocationLists         int    `help:"print information about DWARF location list creation"`
	MaxShapeLen           int    `help:"hash shape names longer than this threshold (default 500)" concurrent:"ok"`
	MergeLocals           int    `help:"merge together non-interfering local stack slots" concurrent:"ok"`
	MergeLocalsDumpFunc   string `help:"dump specified func in merge locals"`
	MergeLocalsHash       string `help:"hash value for debugging stack slot merging of local variables" concurrent:"ok"`
	MergeLocalsTrace      int    `help:"trace debug output for locals merging"`
	MergeLocalsHTrace     int    `help:"hash-selected trace debug output for locals merging"`
	Nil                   int    `help:"print information about nil checks"`
	NoDeadLocals          int    `help:"disable deadlocals pass" concurrent:"ok"`
	NoOpenDefer           int    `help:"disable open-coded defers" concurrent:"ok"`
	NoRefName             int    `help:"do not include referenced symbol names in object file" concurrent:"ok"`
	PCTab                 string `help:"print named pc-value table\nOne of: pctospadj, pctofile, pctoline, pctoinline, pctopcdata"`
	Panic                 int    `help:"show all compiler panics"`
	Reshape               int    `help:"print information about expression reshaping"`
	Shapify               int    `help:"print information about shaping recursive types"`
	Slice                 int    `help:"print information about slice compilation"`
	SoftFloat             int    `help:"force compiler to emit soft-float code" concurrent:"ok"`
	StaticCopy            int    `help:"print information about missed static copies" concurrent:"ok"`
	SyncFrames            int    `help:"how many writer stack frames to include at sync points in unified export data"`
	TailCall              int    `help:"print information about tail calls"`
	TypeAssert            int    `help:"print information about type assertion inlining"`
	WB                    int    `help:"print information about write barriers"`
	ABIWrap               int    `help:"print information about ABI wrapper generation"`
	MayMoreStack          string `help:"call named function before all stack growth checks" concurrent:"ok"`
	PGODebug              int    `help:"debug profile-guided optimizations"`
	PGOHash               string `help:"hash value for debugging profile-guided optimizations" concurrent:"ok"`
	PGOInline             int    `help:"enable profile-guided inlining" concurrent:"ok"`
	PGOInlineCDFThreshold string `help:"cumulative threshold percentage for determining call sites as hot candidates for inlining" concurrent:"ok"`
	PGOInlineBudget       int    `help:"inline budget for hot functions" concurrent:"ok"`
	PGODevirtualize       int    `help:"enable profile-guided devirtualization; 0 to disable, 1 to enable interface devirtualization, 2 to enable function devirtualization" concurrent:"ok"`
	RangeFuncCheck        int    `help:"insert code to check behavior of range iterator functions" concurrent:"ok"`
	VariableMakeHash      string `help:"hash value for debugging stack allocation of variable-sized make results" concurrent:"ok"`
	VariableMakeThreshold int    `help:"threshold in bytes for possible stack allocation of variable-sized make results" concurrent:"ok"`
	WrapGlobalMapDbg      int    `help:"debug trace output for global map init wrapping"`
	WrapGlobalMapCtl      int    `help:"global map init wrap control (0 => default, 1 => off, 2 => stress mode, no size cutoff)"`
	ZeroCopy              int    `help:"enable zero-copy string->[]byte conversions" concurrent:"ok"`
	ConcurrentOk          bool
} // DebugFlags defines the debugging configuration values (see var Debug).
// Each struct field is a different value, named for the lower-case of the field name.
// Each field must be an int or string and must have a `help` struct tag.
//
// The -d option takes a comma-separated list of settings.
// Each setting is name=value; for ints, name is short for name=1.
// true if only concurrentOk flags seen

type CmdFlags struct {
	B                  CountFlag    "help:\"disable bounds checking\""
	C                  CountFlag    "help:\"disable printing of columns in error messages\""
	D                  string       "help:\"set relative `path` for local imports\""
	E                  CountFlag    "help:\"debug symbol export\""
	I                  func(string) "help:\"add `directory` to import search path\""
	K                  CountFlag    "help:\"debug missing line numbers\""
	L                  CountFlag    "help:\"also show actual source file names in error messages for positions affected by //line directives\""
	N                  CountFlag    "help:\"disable optimizations\""
	S                  CountFlag    "help:\"print assembly listing\""
	W                  CountFlag    "help:\"debug parse tree after type checking\""
	LowerC             int          "help:\"concurrency during compilation (1 means no concurrency)\""
	LowerD             flag.Value   "help:\"enable debugging settings; try -d help\""
	LowerE             CountFlag    "help:\"no limit on number of errors reported\""
	LowerH             CountFlag    "help:\"halt on error\""
	LowerJ             CountFlag    "help:\"debug runtime-initialized variables\""
	LowerL             CountFlag    "help:\"disable inlining\""
	LowerM             CountFlag    "help:\"print optimization decisions\""
	LowerO             string       "help:\"write output to `file`\""
	LowerP             *string      "help:\"set expected package import `path`\""
	LowerR             CountFlag    "help:\"debug generated wrappers\""
	LowerT             bool         "help:\"enable tracing for debugging the compiler\""
	LowerW             CountFlag    "help:\"debug type checking\""
	LowerV             *bool        "help:\"increase debug verbosity\""
	Percent            CountFlag    "flag:\"%\" help:\"debug non-static initializers\""
	CompilingRuntime   bool         "flag:\"+\" help:\"compiling runtime\""
	AsmHdr             string       "help:\"write assembly header to `file`\""
	ASan               bool         "help:\"build code compatible with C/C++ address sanitizer\""
	Bench              string       "help:\"append benchmark times to `file`\""
	BlockProfile       string       "help:\"write block profile to `file`\""
	BuildID            string       "help:\"record `id` as the build id in the export metadata\""
	CPUProfile         string       "help:\"write cpu profile to `file`\""
	Complete           bool         "help:\"compiling complete package (no C or assembly)\""
	ClobberDead        bool         "help:\"clobber dead stack slots (for debugging)\""
	ClobberDeadReg     bool         "help:\"clobber dead registers (for debugging)\""
	Dwarf              bool         "help:\"generate DWARF symbols\""
	DwarfBASEntries    *bool        "help:\"use base address selection entries in DWARF\""
	DwarfLocationLists *bool        "help:\"add location lists to DWARF in optimized mode\""
	Dynlink            *bool        "help:\"support references to Go symbols defined in other shared libraries\""
	EmbedCfg           func(string) "help:\"read go:embed configuration from `file`\""
	Env                func(string) "help:\"add `definition` of the form key=value to environment\""
	GenDwarfInl        int          "help:\"generate DWARF inline info records\""
	GoVersion          string       "help:\"required version of the runtime\""
	ImportCfg          func(string) "help:\"read import configuration from `file`\""
	InstallSuffix      string       "help:\"set pkg directory `suffix`\""
	JSON               string       "help:\"version,file for JSON compiler/optimizer detail output\""
	Lang               string       "help:\"Go language version source code expects\""
	LinkObj            string       "help:\"write linker-specific object to `file`\""
	LinkShared         *bool        "help:\"generate code that will be linked against Go shared libraries\""
	Live               CountFlag    "help:\"debug liveness analysis\""
	MSan               bool         "help:\"build code compatible with C/C++ memory sanitizer\""
	MemProfile         string       "help:\"write memory profile to `file`\""
	MemProfileRate     int          "help:\"set runtime.MemProfileRate to `rate`\""
	MutexProfile       string       "help:\"write mutex profile to `file`\""
	NoLocalImports     bool         "help:\"reject local (relative) imports\""
	CoverageCfg        func(string) "help:\"read coverage configuration from `file`\""
	Pack               bool         "help:\"write to file.a instead of file.o\""
	Race               bool         "help:\"enable race detector\""
	Shared             *bool        "help:\"generate code that can be linked into a shared library\""
	SmallFrames        bool         "help:\"reduce the size limit for stack allocated objects\""
	Spectre            string       "help:\"enable spectre mitigations in `list` (all, index, ret)\""
	Std                bool         "help:\"compiling standard library\""
	SymABIs            string       "help:\"read symbol ABIs from `file`\""
	TraceProfile       string       "help:\"write an execution trace to `file`\""
	TrimPath           string       "help:\"remove `prefix` from recorded source file paths\""
	WB                 bool         "help:\"enable write barrier\""
	PgoProfile         string       "help:\"read profile or pre-process profile from `file`\""
	ErrorURL           bool         "help:\"print explanatory URL with error message if applicable\""
	Cfg                struct {
		Embed struct {
			Patterns map[ // CmdFlags defines the command-line flags (see var Flag).
			// Each struct field is a different flag, by default named for the lower-case of the field name.
			// If the flag name is a single letter, the default flag name is left upper-case.
			// If the flag name is "Lower" followed by a single letter, the default flag name is the lower-case of the last letter.
			//
			// If this default flag name can't be made right, the `flag` struct tag can be used to replace it,
			// but this should be done only in exceptional circumstances: it helps everyone if the flag name
			// is obvious from the field name when the flag is used elsewhere in the compiler sources.
			// The `flag:"-"` struct tag makes a field invisible to the flag logic and should also be used sparingly.
			//
			// Each field must have a `help` struct tag giving the flag help message.
			//
			// The allowed field types are bool, int, string, pointers to those (for values stored elsewhere),
			// CountFlag (for a counting flag), and func(string) (for a flag that uses special code for parsing).
			// Configuration derived from flags; not a flag itself.
			string][]string
			Files map[string]string
		}
		ImportDirs []string
		ImportMap  map[ // appended to by -I
		string]string
		PackageFile map[ // set by -importcfg
		string]string
		CoverageInfo  *covcmd.CoverFixupConfig
		SpectreIndex  bool
		Instrumenting bool
	}
} // set by -importcfg; nil means not in use
// Whether we are adding any sort of code instrumentation, such as
// when the race detector is enabled.

type hashAndMask struct {
	hash uint64
	mask uint64
	name string
} // a hash h matches if (h^hash)&mask == 0
// base name, or base name + "0", "1", etc.

type HashDebug struct {
	mu      sync.Mutex
	name    string
	logfile io.Writer
	posTmp  []src.// for logfile, posTmp, bytesTmp
	// what file (if any) receives the yes/no logging?
	// default is os.Stdout
	Pos
	bytesTmp bytes.Buffer
	matches  []hashAndMask
	excludes []hashAndMask// A hash matches if one of these matches.

	bisect           *bisect.Matcher
	fileSuffixOnly   bool
	inlineSuffixOnly bool
} // explicitly excluded hash suffixes
// for Pos hashes, remove all but the most inline position.

type errorMsg struct {
	pos  src.XPos
	msg  string
	code errors.Code
} // An errorMsg is a queued error message, waiting to be printed.

type Timings struct {
	list []timestamp// Timings collects the execution times of labeled phases
	// which are added through a sequence of Start/Stop calls.
	// Events may be associated with each phase via AddEvent.

	events map[int][]*event
} // lazily allocated

type timestamp struct {
	time  time.Time
	label string
	start bool
}

type event struct {
	size int64
	unit string
} // count or amount of data processed (allocations, data size, lines, funcs, ...)
// unit of size measure (count, MB, lines, funcs, ...)

type BitVec struct {
	N int32
	B []uint32// A BitVec is a bit vector.
	// number of bits in vector

} // words holding bits

type Bulk struct {
	words []uint32
	nbit  int32
	nword int32
}

type names struct {
	MetaVar     *ir.Name
	PkgIdVar    *ir.Name
	InitFn      *ir.Func
	CounterMode coverage.CounterMode
	CounterGran coverage.CounterGranularity
} // names records state information collected in the first fixup
// phase so that it can be passed to the second fixup phase.

type visitor struct {
	curfn *ir.Func
	defs  map[ // defs[name] contains assignments that can be discarded if name can be discarded.
	// if defs[name] is defined nil, then name is actually used.
	*ir.Name][]assign
	defsKeys []*ir.Name
	doNode   func(ir.Node) bool
} // insertion order of keys, for reproducible iteration (and builds)

type assign struct {
	pos      src.XPos
	lhs, rhs *ir.Node
}

type CallStat struct {
	Pkg                 string
	Pos                 string
	Caller              string
	Direct              bool
	Interface           bool
	Weight              int64
	Hottest             string
	HottestWeight       int64
	Devirtualized       string
	DevirtualizedWeight int64
} // CallStat summarizes a single call site.
//
// This is used only for debug logging.
// Devirtualized callee if != "".
//
// Note that this may be different than Hottest because we apply
// type-check restrictions, which helps distinguish multiple calls on
// the same line.

type varsAndDecls struct {
	decls      []*ir.Name
	vars       []*dwarf.Var
	paramOrder map[*ir.Name]int
}

type varPos struct {
	DeclName string
	DeclFile string
	DeclLine uint
	DeclCol  uint
} // To identify variables by original source position.

type ScopeMarker struct {
	parents []ir.// A ScopeMarker tracks scope nesting and boundaries for later use
	// during DWARF generation.
	ScopeID
	marks []ir.Mark
}

type varsByScopeAndOffset struct {
	vars   []*dwarf.Var
	scopes []ir.ScopeID
}

type varsByScope struct {
	vars   []*dwarf.Var
	scopes []ir.ScopeID
}

type batch struct {
	allLocs []*// A batch holds escape analysis state that's shared across an entire
	// batch of functions being analyzed at once.
	location
	closures        []closure
	reassignOracles map[*ir.Func]*ir.ReassignOracle
	heapLoc         location
	mutatorLoc      location
	calleeLoc       location
	blankLoc        location
}

type closure struct {
	k   hole
	clo *ir.ClosureExpr
} // A closure holds a closure expression and its spill hole (i.e.,
// where the hole representing storing into its closure record).

type escape struct {
	*batch
	curfn  *ir.Func
	labels map[ // An escape holds state specific to a single function being analyzed
	// within a batch.
	// function being analyzed
	*types.Sym]labelState
	loopDepth int
} // known labels
// loopDepth counts the current loop nesting depth within
// curfn. It increments within each "for" loop and at each
// label with a corresponding backwards "goto" (i.e.,
// unstructured loop).

type location struct {
	n     ir.Node
	curfn *ir.Func
	edges []edge// A location represents an abstract location that stores a Go
	// variable.
	// enclosing function

	loopDepth     int
	resultIndex   int
	derefs        int
	walkgen       uint32
	dst           *location
	dstEdgeIdx    int
	queuedWalkAll bool
	queuedWalkOne uint32
	attrs         locAttr
	paramEsc      leaks
	captured      bool
	reassigned    bool
	addrtaken     bool
	param         bool
	paramOut      bool
} // incoming edges
// is this variable an out parameter (ONAME of class ir.PPARAMOUT)?

type edge struct {
	src    *location
	derefs int
	notes  *note
} // An edge represents an assignment edge between two Go variables.
// >= -1

type hole struct {
	dst       *location
	derefs    int
	notes     *note
	addrtaken bool
} // A hole represents a context for evaluation of a Go
// expression. E.g., when evaluating p in "x = **p", we'd have a hole
// with dst==x and derefs==2.
// addrtaken indicates whether this context is taking the address of
// the expression, independent of whether the address will actually
// be stored into a variable.

type note struct {
	next  *note
	where ir.Node
	why   string
}

type queue struct {
	locs []*// queue implements a queue of locations for use in WalkAll and WalkOne.
	// It supports pushing to front & back, and popping from front.
	// TODO(thepudds): does cmd/compile have a deque or similar somewhere?
	location
	head  int
	tail  int
	elems int
} // index of front element
// next back element

type fakeFileSet struct {
	fset  *token.FileSet
	files map[ // Synthesize a token.Pos
	string]*token.File
}

type anyType struct{}

type derivedInfo struct {
	idx    pkgbits.Index
	needed bool
} // See cmd/compile/internal/noder.derivedInfo.

type typeInfo struct {
	idx     pkgbits.Index
	derived bool
} // See cmd/compile/internal/noder.typeInfo.

type E interface{ M() T }

type BlankField struct{ _ int } // Any release before and including Go 1.7 didn't encode
// the package for a blank struct field.

type pkgReader struct {
	pkgbits.PkgDecoder
	ctxt        *types2.Context
	imports     map[string]*types2.Package
	enableAlias bool
	posBases    []*// whether to use aliases
	syntax.PosBase
	pkgs []*types2.Package
	typs []types2.Type
}

type reader struct {
	pkgbits.Decoder
	p    *pkgReader
	dict *readerDict
}

type readerDict struct {
	bounds       []typeInfo
	tparams      []*types2.TypeParam
	derived      []derivedInfo
	derivedTypes []types2.Type
}

type readerTypeBound struct {
	derived  bool
	boundIdx int
}

type hairyVisitor struct {
	curFunc       *ir.Func
	isBigFunc     bool
	budget        int32
	maxBudget     int32
	reason        string
	extraCallCost int32
	usedLocals    ir.NameSet
	do            func(ir.Node) bool
	profile       *pgoir.Profile
} // hairyVisitor visits a function body to determine its inlining
// hairiness and whether or not it can be inlined.
// This is needed to access the current caller in the doNode function.

type propAnalyzer interface {
	nodeVisitPre(n ir.Node)
	nodeVisitPost(n ir.Node)
	setResults(funcProps *FuncProps)
} // propAnalyzer interface is used for defining one or more analyzer
// helper objects, each tasked with computing some specific subset of
// the properties we're interested in. The assumption is that
// properties are independent, so each new analyzer that implements
// this interface can operate entirely on its own. For a given analyzer
// there will be a sequence of calls to nodeVisitPre and nodeVisitPost
// as the nodes within a function are visited, then a followup call to
// setResults so that the analyzer can transfer its results into the
// final properties object.

type fnInlHeur struct {
	props *FuncProps
	cstab CallSiteTab
	fname string
	file  string
	line  uint
} // fnInlHeur contains inline heuristics state information about a
// specific Go function being analyzed/considered by the inliner. Note
// that in addition to constructing a fnInlHeur object by analyzing a
// specific *ir.Func, there is also code in the test harness
// (funcprops_test.go) that builds up fnInlHeur's by reading in and
// parsing a dump. This is the reason why we have file/fname/line
// fields below instead of just an *ir.Func field.

type callSiteAnalyzer struct {
	fn *ir.Func
	*nameFinder
}

type callSiteTableBuilder struct {
	fn *ir.Func
	*nameFinder
	cstab    CallSiteTab
	ptab     map[ir.Node]pstate
	nstack   []ir.Node
	loopNest int
	isInit   bool
}

type funcFlagsAnalyzer struct {
	fn     *ir.Func
	nstate map[ // funcFlagsAnalyzer computes the "Flags" value for the FuncProps
	// object we're computing. The main item of interest here is "nstate",
	// which stores the disposition of a given ir Node with respect to the
	// flags/properties we're trying to compute.
	ir.Node]pstate
	noInfo bool
} // set if we see something inscrutable/un-analyzable

type paramsAnalyzer struct {
	fname  string
	values []ParamPropBits// paramsAnalyzer holds state information for the phase that computes
	// flags for a Go functions parameters, for use in inline heuristics.
	// Note that the params slice below includes entries for blanks.

	params []*ir.Name
	top    []bool
	*condLevelTracker
	*nameFinder
}

type condLevelTracker struct{ condLevel int } // condLevelTracker helps keeps track very roughly of "level of conditional
// nesting", e.g. how many "if" statements you have to go through to
// get to the point where a given stmt executes. Example:
//
//	                      cond nesting level
//	func foo() {
//	 G = 1                   0
//	 if x < 10 {             0
//	  if y < 10 {            1
//	   G = 0                 2
//	  }
//	 }
//	}
//
// The intent here is to provide some sort of very abstract relative
// hotness metric, e.g. "G = 1" above is expected to be executed more
// often than "G = 0" (in the aggregate, across large numbers of
// functions).

type resultsAnalyzer struct {
	fname string
	props []ResultPropBits// resultsAnalyzer stores state information for the process of
	// computing flags/properties for the return values of a specific Go
	// function, as part of inline heuristics synthesis.

	values          []resultVal
	inlineMaxBudget int
	*nameFinder
}

type resultVal struct {
	cval    constant.Value
	fn      *ir.Name
	fnClo   bool
	top     bool
	derived bool
} // resultVal captures information about a specific result returned from
// the function we're analyzing; we are interested in cases where
// the func always returns the same constant, or always returns
// the same function, etc. This container stores info on a the specific
// scenarios we're looking for.
// see deriveReturnFlagsFromCallee below

type CallSite struct {
	Callee   *ir.Func
	Call     *ir.CallExpr
	parent   *CallSite
	Assign   ir.Node
	Flags    CSPropBits
	ArgProps []ActualExprPropBits// CallSite records useful information about a potentially inlinable
	// (direct) function call. "Callee" is the target of the call, "Call"
	// is the ir node corresponding to the call itself, "Assign" is
	// the top-level assignment statement containing the call (if the call
	// appears in the form of a top-level statement, e.g. "x := foo()"),
	// "Flags" contains properties of the call that might be useful for
	// making inlining decisions, "Score" is the final score assigned to
	// the site, and "ID" is a numeric ID for the site within its
	// containing function.

	Score     int
	ScoreMask scoreAdjustTyp
	ID        uint
	aux       uint8
}

type propsAndScore struct {
	props CSPropBits
	score int
	mask  scoreAdjustTyp
}

type exprClassifier struct {
	names map[ // exprClassifier holds intermediate state about nodes within an
	// expression tree being analyzed by ShouldFoldIfNameConstant. Here
	// "name" is the name node passed in, and "disposition" stores the
	// result of classifying a given IR node.
	*ir.Name]bool
	disposition map[ir.Node]disp
}

type FuncProps struct {
	Flags      FuncPropBits
	ParamFlags []ParamPropBits// FuncProps describes a set of function or method properties that may
	// be useful for inlining heuristics. Here 'Flags' are properties that
	// we think apply to the entire function; 'RecvrParamFlags' are
	// properties of specific function params (or the receiver), and
	// 'ResultFlags' are things properties we think will apply to values
	// of specific results. Note that 'ParamFlags' includes and entry for
	// the receiver if applicable, and does include etries for blank
	// params; for a function such as "func foo(_ int, b byte, _ float32)"
	// the length of ParamFlags will be 3.

	ResultFlags []ResultPropBits// slot 0 receiver if applicable

}

type nameFinder struct{ ro *ir.ReassignOracle } // nameFinder provides a set of "isXXX" query methods for clients to
// ask whether a given AST node corresponds to a function, a constant
// value, and so on. These methods use an underlying ir.ReassignOracle
// to return more precise results in cases where an "interesting"
// value is assigned to a singly-defined local temp. Example:
//
//	const q = 101
//	fq := func() int { return q }
//	copyOfConstant := q
//	copyOfFunc := f
//	interestingCall(copyOfConstant, copyOfFunc)
//
// A name finder query method invoked on the arguments being passed to
// "interestingCall" will be able detect that 'copyOfConstant' always
// evaluates to a constant (even though it is in fact a PAUTO local
// variable). A given nameFinder can also operate without using
// ir.ReassignOracle (in cases where it is not practical to look
// at the entire function); in such cases queries will still work
// for explicit constant values and functions.

type resultPropAndCS struct {
	defcs *CallSite
	props ResultPropBits
}

type resultUseAnalyzer struct {
	resultNameTab map[*ir.Name]resultPropAndCS
	fn            *ir.Func
	cstab         CallSiteTab
	*condLevelTracker
}

type scoreCallsCacheType struct {
	tab CallSiteTab
	csl []*CallSite
}

type Bar struct {
	x int
	y string
}

type Itf interface{ Plark() }

type callSite struct {
	fn         *ir.Func
	whichParen int
}

type inlClosureState struct {
	fn        *ir.Func
	profile   *pgoir.Profile
	callSites map[*ir.ParenExpr]bool
	resolved  []*// callSites[p] == "p appears in parens" (do not append again)
	ir.Func
	useCounts map[ // for each call in parens, the resolved target of the call
	*ir.Func]int
	parens []*// shared among all InlClosureStates
	ir.ParenExpr
	bigCaller bool
}

type dumper struct {
	output  io.Writer
	fieldrx *regexp.Regexp
	ptrmap  map[ // field name filter
	uintptr]int
	lastadr string
	indent  int
	last    byte
	line    int
} // ptr -> dump line number
// current line number

type Expr interface {
	Node
	isExpr()
} // An Expr is a Node that can appear as an expression.

type miniExpr struct {
	miniNode
	flags bitset8
	typ   *types.Type
	init  Nodes
} // A miniExpr is a miniNode with extra fields common to expressions.
// TODO(rsc): Once we are sure about the contents, compact the bools
// into a bit field and leave extra bits available for implementations
// embedding miniExpr. Right now there are ~24 unused bits sitting here.
// TODO(rsc): Don't require every Node to have an init

type AddStringExpr struct {
	miniExpr
	List     Nodes
	Prealloc *Name
} // An AddStringExpr is a string concatenation List[0] + List[1] + ... + List[len(List)-1].

type AddrExpr struct {
	miniExpr
	X        Node
	Prealloc *Name
} // An AddrExpr is an address-of expression &X.
// It may end up being a normal address-of or an allocation of a composite literal.
// preallocated storage if any

type BasicLit struct {
	miniExpr
	val constant.Value
} // A BasicLit is a literal of basic type.

type BinaryExpr struct {
	miniExpr
	X     Node
	Y     Node
	RType Node `mknode:"-"`
} // A BinaryExpr is a binary expression X Op Y,
// or Op(X, Y) for builtin functions that do not become calls.
// see reflectdata/helpers.go

type CallExpr struct {
	miniExpr
	Fun       Node
	Args      Nodes
	DeferAt   Node
	RType     Node `mknode:"-"`
	KeepAlive []*// A CallExpr is a function call Fun(Args).
	// see reflectdata/helpers.go
	Name
	IsDDD    bool
	GoDefer  bool
	NoInline bool
} // vars to be kept alive until call returns
// whether this call must not be inlined

type ClosureExpr struct {
	miniExpr
	Func     *Func `mknode:"-"`
	Prealloc *Name
	IsGoWrap bool
} // A ClosureExpr is a function literal expression.
// whether this is wrapper closure of a go statement

type CompLitExpr struct {
	miniExpr
	List     Nodes
	RType    Node `mknode:"-"`
	Prealloc *Name
	Len      int64
} // A CompLitExpr is a composite literal Type{Vals}.
// Before type-checking, the type is Ntype.
// For OSLICELIT, Len is the backing array length.
// For OMAPLIT, Len is the number of entries that we've removed from List and
// generated explicit mapassign calls for. This is used to inform the map alloc hint.

type ConvExpr struct {
	miniExpr
	X             Node
	TypeWord      Node `mknode:"-"`
	SrcRType      Node `mknode:"-"`
	ElemRType     Node `mknode:"-"`
	ElemElemRType Node `mknode:"-"`
} // A ConvExpr is a conversion Type(X).
// It may end up being a value or a type.
// For -d=checkptr instrumentation of conversions from
// unsafe.Pointer to *Elem or *[Len]Elem.
//
// TODO(mdempsky): We only ever need one of these, but currently we
// don't decide which one until walk. Longer term, it probably makes
// sense to have a dedicated IR op for `(*[Len]Elem)(ptr)[:n:m]`
// expressions.

type IndexExpr struct {
	miniExpr
	X        Node
	Index    Node
	RType    Node `mknode:"-"`
	Assigned bool
} // An IndexExpr is an index expression X[Index].
// see reflectdata/helpers.go

type KeyExpr struct {
	miniExpr
	Key   Node
	Value Node
} // A KeyExpr is a Key: Value composite literal key.

type StructKeyExpr struct {
	miniExpr
	Field *types.Field
	Value Node
} // A StructKeyExpr is a Field: Value composite literal key.

type InlinedCallExpr struct {
	miniExpr
	Body       Nodes
	ReturnVars Nodes
} // An InlinedCallExpr is an inlined function call.
// must be side-effect free

type LogicalExpr struct {
	miniExpr
	X Node
	Y Node
} // A LogicalExpr is an expression X Op Y where Op is && or ||.
// It is separate from BinaryExpr to make room for statements
// that must be executed before Y but after X.

type MakeExpr struct {
	miniExpr
	RType Node `mknode:"-"`
	Len   Node
	Cap   Node
} // A MakeExpr is a make expression: make(Type[, Len[, Cap]]).
// Op is OMAKECHAN, OMAKEMAP, OMAKESLICE, or OMAKESLICECOPY,
// but *not* OMAKE (that's a pre-typechecking CallExpr).
// see reflectdata/helpers.go

type NilExpr struct{ miniExpr } // A NilExpr represents the predefined untyped constant nil.

type ParenExpr struct {
	miniExpr
	X Node
} // A ParenExpr is a parenthesized expression (X).
// It may end up being a value or a type.

type ResultExpr struct {
	miniExpr
	Index int64
} // A ResultExpr represents a direct access to a result.
// index of the result expr.

type LinksymOffsetExpr struct {
	miniExpr
	Linksym *obj.LSym
	Offset_ int64
} // A LinksymOffsetExpr refers to an offset within a global variable.
// It is like a SelectorExpr but without the field name.

type SelectorExpr struct {
	miniExpr
	X         Node
	Sel       *types.Sym
	Selection *types.Field
	Prealloc  *Name
} // A SelectorExpr is a selector expression X.Sel.
// preallocated storage for OMETHVALUE, if any

type SliceExpr struct {
	miniExpr
	X    Node
	Low  Node
	High Node
	Max  Node
} // A SliceExpr is a slice expression X[Low:High] or X[Low:High:Max].

type SliceHeaderExpr struct {
	miniExpr
	Ptr Node
	Len Node
	Cap Node
} // A SliceHeader expression constructs a slice header from its parts.

type StringHeaderExpr struct {
	miniExpr
	Ptr Node
	Len Node
} // A StringHeaderExpr expression constructs a string header from its parts.

type StarExpr struct {
	miniExpr
	X Node
} // A StarExpr is a dereference expression *X.
// It may end up being a value or a type.

type TypeAssertExpr struct {
	miniExpr
	X          Node
	ITab       Node `mknode:"-"`
	Descriptor *obj.LSym
} // A TypeAssertionExpr is a selector expression X.(Type).
// Before type-checking, the type is Ntype.
// An internal/abi.TypeAssert descriptor to pass to the runtime.

type DynamicTypeAssertExpr struct {
	miniExpr
	X        Node
	SrcRType Node
	RType    Node
	ITab     Node
} // A DynamicTypeAssertExpr asserts that X is of dynamic type RType.
// ITab is an expression that yields a *runtime.itab value
// representing the asserted type within the assertee expression's
// original interface type.
//
// ITab is only used for assertions from non-empty interface type to
// a concrete (i.e., non-interface) type. For all other assertions,
// ITab is nil.

type UnaryExpr struct {
	miniExpr
	X Node
} // A UnaryExpr is a unary expression Op X,
// or Op(X) for a builtin function that does not end up being a call.

type Func struct {
	miniNode
	Body     Nodes
	Nname    *Name
	OClosure *ClosureExpr
	Dcl      []*// A Func corresponds to a single function in a Go program
	// (and vice versa: each function is denoted by exactly one *Func).
	//
	// There are multiple nodes that represent a Func in the IR.
	//
	// The ONAME node (Func.Nname) is used for plain references to it.
	// The ODCLFUNC node (the Func itself) is used for its declaration code.
	// The OCLOSURE node (Func.OClosure) is used for a reference to a
	// function literal.
	//
	// An imported function will have an ONAME node which points to a Func
	// with an empty body.
	// A declared function or method has an ODCLFUNC (the Func itself) and an ONAME.
	// A function literal is represented directly by an OCLOSURE, but it also
	// has an ODCLFUNC (and a matching ONAME) representing the compiled
	// underlying form of the closure, which accesses the captured variables
	// using a special data structure passed in a register.
	//
	// A method declaration is represented like functions, except f.Sym
	// will be the qualified method name (e.g., "T.m").
	//
	// A method expression (T.M) is represented as an OMETHEXPR node,
	// in which n.Left and n.Right point to the type and method, respectively.
	// Each distinct mention of a method expression in the source code
	// constructs a fresh node.
	//
	// A method value (t.M) is represented by ODOTMETH/ODOTINTER
	// when it is called directly and by OMETHVALUE otherwise.
	// These are like method expressions, except that for ODOTMETH/ODOTINTER,
	// the method name is stored in Sym instead of Right.
	// Each OMETHVALUE ends up being implemented as a new
	// function, a bit like a closure, with its own ODCLFUNC.
	// The OMETHVALUE uses n.Func to record the linkage to
	// the generated ODCLFUNC, but there is no
	// pointer from the Func back to the OMETHVALUE.
	// ONAME nodes for all params/locals for this func/closure, does NOT
	// include closurevars until transforming closures during walk.
	// Names must be listed PPARAMs, PPARAMOUTs, then PAUTOs,
	// with PPARAMs and PPARAMOUTs in order corresponding to the function signature.
	// Anonymous and blank params are declared as ~pNN (for PPARAMs) and ~rNN (for PPARAMOUTs).
	Name
	ClosureVars []*// ClosureVars lists the free variables that are used within a
	// function literal, but formally declared in an enclosing
	// function. The variables in this slice are the closure function's
	// own copy of the variables, which are used within its function
	// body. They will also each have IsClosureVar set, and will have
	// Byval set if they're captured by value.
	Name
	Closures []*// Enclosed functions that need to be compiled.
	// Populated during walk.
	Func
	ClosureParent *Func
	Parents       []ScopeID// Parent of a closure
	// Parents records the parent scope of each scope within a
	// function. The root scope (0) has no parent, so the i'th
	// scope's parent is stored at Parents[i-1].

	Marks []Mark// Marks records scope boundary changes.

	FieldTrack  map[*obj.LSym]struct{}
	DebugInfo   interface{}
	LSym        *obj.LSym
	Inl         *Inline
	RangeParent *Func
	funcLitGen  int32
	rangeLitGen int32
	goDeferGen  int32
	Label       int32
	Endlineno   src.XPos
	WBPos       src.XPos
	Pragma      PragmaFlag
	flags       bitset16
	ABI         obj.ABI
	ABIRefs     obj.ABISet
	NumDefers   int32
	NumReturns  int32
	NWBRCalls   *[]SymAndPos// Linker object in this function's native ABI (Func.ABI)
	// NWBRCalls records the LSyms of functions called by this
	// function for go:nowritebarrierrec analysis. Only filled in
	// if nowritebarrierrecCheck != nil.

	WrappedFunc *Func
	WasmImport  *WasmImport
	WasmExport  *WasmExport
} // For wrapper functions, WrappedFunc point to the original Func.
// Currently only used for go/defer wrappers.
// WasmExport is used by the //go:wasmexport directive to store info about
// a WebAssembly function import.

type WasmImport struct {
	Module string
	Name   string
} // WasmImport stores metadata associated with the //go:wasmimport pragma.

type WasmExport struct{ Name string } // WasmExport stores metadata associated with the //go:wasmexport pragma.

type Inline struct {
	Cost int32
	Dcl  []*// An Inline holds fields used for function bodies that can be inlined.
	// Copy of Func.Dcl for use during inlining. This copy is needed
	// because the function's Dcl may change from later compiler
	// transformations. This field is also populated when a function
	// from another package is imported and inlined.
	Name
	HaveDcl         bool
	Properties      string
	CanDelayResults bool
} // whether we've loaded Dcl
// CanDelayResults reports whether it's safe for the inliner to delay
// initializing the result parameters until immediately before the
// "return" statement.

type Mark struct {
	Pos   src.XPos
	Scope ScopeID
} // A Mark represents a scope boundary.
// Scope identifies the innermost scope to the right of Pos.

type SymAndPos struct {
	Sym *obj.LSym
	Pos src.XPos
} // LSym of callee
// line of call

type miniNode struct {
	pos  src.XPos
	op   Op
	bits bitset8
	esc  uint16
} // A miniNode is a minimal node implementation,
// meant to be embedded as the first field in a larger node implementation,
// at a cost of 12 bytes.
//
// A miniNode is NOT a valid Node by itself: the embedding struct
// must at the least provide:
//
//	func (n *MyNode) String() string { return fmt.Sprint(n) }
//	func (n *MyNode) rawCopy() Node { c := *n; return &c }
//	func (n *MyNode) Format(s fmt.State, verb rune) { FmtNode(n, s, verb) }
//
// The embedding struct should also fill in n.op in its constructor,
// for more useful panic messages when invalid methods are called,
// instead of implementing Op itself.

type Ident struct {
	miniExpr
	sym *types.Sym
} // An Ident is an identifier, possibly qualified.

type Embed struct {
	Pos      src.XPos
	Patterns []string
}

type Node interface {
	Format(s fmt.State, verb rune)
	Pos() src.XPos
	SetPos(x src.XPos)
	copy() Node
	doChildren(func(Node) bool) bool
	doChildrenWithHidden(func(Node) bool) bool
	editChildren(func(Node) Node)
	editChildrenWithHidden(func(Node) Node)
	Op() Op
	Init() Nodes
	Type() *types.Type
	SetType(t *types.Type)
	Name() *Name
	Sym() *types.Sym
	Val() constant.Value
	SetVal(v constant.Value)
	Esc() uint16
	SetEsc(x uint16)
	Typecheck() uint8
	SetTypecheck(x uint8)
	NonNil() bool
	MarkNonNil()
} // A Node is the abstract interface to an IR node.
// Typecheck values:
//  0 means the node is not typechecked
//  1 means the node is completely typechecked
//  2 means typechecking of the node is in progress

type InitNode interface {
	Node
	PtrInit() *Nodes
	SetInit(x Nodes)
}

type NameQueue struct {
	ring []*// NameQueue is a FIFO queue of *Name. The zero value of NameQueue is
	// a ready-to-use empty queue.
	Name
	head, tail int
}

type ReassignOracle struct {
	fn        *Func
	singleDef map[ // A ReassignOracle efficiently answers queries about whether local
	// variables are reassigned. This helper works by looking for function
	// params and short variable declarations (e.g.
	// https://go.dev/ref/spec#Short_variable_declarations) that are
	// neither address taken nor subsequently re-assigned. It is intended
	// to operate much like "ir.StaticValue" and "ir.Reassigned", but in a
	// way that does just a single walk of the containing function (as
	// opposed to a new walk on every call).
	// maps candidate name to its defining assignment (or
	// for params, defining func).
	*Name]Node
}

type bottomUpVisitor struct {
	analyze  func([]*Func, bool)
	visitgen uint32
	nodeID   map[*Func]uint32
	stack    []*Func
}

type Decl struct {
	miniNode
	X *Name
} // A Decl is a declaration of a const, type, or var. (A declared func is a Func.)
// the thing being declared

type Stmt interface {
	Node
	isStmt()
} // A Stmt is a Node that can appear as a statement.
// This includes statement-like expressions such as f().
//
// (It's possible it should include <-c, but that would require
// splitting ORECV out of UnaryExpr, which hasn't yet been
// necessary. Maybe instead we will introduce ExprStmt at
// some point.)

type miniStmt struct {
	miniNode
	init Nodes
} // A miniStmt is a miniNode with extra fields common to statements.

type AssignListStmt struct {
	miniStmt
	Lhs Nodes
	Def bool
	Rhs Nodes
} // An AssignListStmt is an assignment statement with
// more than one item on at least one side: Lhs = Rhs.
// If Def is true, the assignment is a :=.

type AssignStmt struct {
	miniStmt
	X   Node
	Def bool
	Y   Node
} // An AssignStmt is a simple assignment statement: X = Y.
// If Def is true, the assignment is a :=.

type AssignOpStmt struct {
	miniStmt
	X      Node
	AsOp   Op
	Y      Node
	IncDec bool
} // An AssignOpStmt is an AsOp= assignment statement: X AsOp= Y.
// actually ++ or --

type BlockStmt struct {
	miniStmt
	List Nodes
} // A BlockStmt is a block: { List }.

type BranchStmt struct {
	miniStmt
	Label *types.Sym
} // A BranchStmt is a break, continue, fallthrough, or goto statement.
// label if present

type CaseClause struct {
	miniStmt
	Var    *Name
	List   Nodes
	RTypes Nodes
	Body   Nodes
} // A CaseClause is a case statement in a switch or select: case List: Body.
// RTypes is a list of RType expressions, which are copied to the
// corresponding OEQ nodes that are emitted when switch statements
// are desugared. RTypes[i] must be non-nil if the emitted
// comparison for List[i] will be a mixed interface/concrete
// comparison; see reflectdata.CompareRType for details.
//
// Because mixed interface/concrete switch cases are rare, we allow
// len(RTypes) < len(List). Missing entries are implicitly nil.

type CommClause struct {
	miniStmt
	Comm Node
	Body Nodes
} // communication case

type ForStmt struct {
	miniStmt
	Label        *types.Sym
	Cond         Node
	Post         Node
	Body         Nodes
	DistinctVars bool
} // A ForStmt is a non-range for loop: for Init; Cond; Post { Body }

type GoDeferStmt struct {
	miniStmt
	Call    Node
	DeferAt Expr
} // A GoDeferStmt is a go or defer statement: go Call / defer Call.
//
// The two opcodes use a single syntax because the implementations
// are very similar: both are concerned with saving Call and running it
// in a different context (a separate goroutine or a later time).

type IfStmt struct {
	miniStmt
	Cond   Node
	Body   Nodes
	Else   Nodes
	Likely bool
} // An IfStmt is a return statement: if Init; Cond { Body } else { Else }.
// code layout hint

type JumpTableStmt struct {
	miniStmt
	Idx   Node
	Cases []constant.// A JumpTableStmt is used to implement switches. Its semantics are:
	//
	//	tmp := jt.Idx
	//	if tmp == Cases[0] goto Targets[0]
	//	if tmp == Cases[1] goto Targets[1]
	//	...
	//	if tmp == Cases[n] goto Targets[n]
	//
	// Note that a JumpTableStmt is more like a multiway-goto than
	// a multiway-if. In particular, the case bodies are just
	// labels to jump to, not full Nodes lists.
	// If Idx is equal to Cases[i], jump to Targets[i].
	// Cases entries must be distinct and in increasing order.
	// The length of Cases and Targets must be equal.
	Value
	Targets []*types.Sym
}

type InterfaceSwitchStmt struct {
	miniStmt
	Case        Node
	Itab        Node
	RuntimeType Node
	Hash        Node
	Descriptor  *obj.LSym
} // An InterfaceSwitchStmt is used to implement type switches.
// Its semantics are:
//
//	if RuntimeType implements Descriptor.Cases[0] {
//	    Case, Itab = 0, itab<RuntimeType, Descriptor.Cases[0]>
//	} else if RuntimeType implements Descriptor.Cases[1] {
//	    Case, Itab = 1, itab<RuntimeType, Descriptor.Cases[1]>
//	...
//	} else if RuntimeType implements Descriptor.Cases[N-1] {
//	    Case, Itab = N-1, itab<RuntimeType, Descriptor.Cases[N-1]>
//	} else {
//	    Case, Itab = len(cases), nil
//	}
//
// RuntimeType must be a non-nil *runtime._type.
// Hash must be the hash field of RuntimeType (or its copy loaded from an itab).
// Descriptor must represent an abi.InterfaceSwitch global variable.

type InlineMarkStmt struct {
	miniStmt
	Index int64
} // An InlineMarkStmt is a marker placed just before an inlined body.

type LabelStmt struct {
	miniStmt
	Label *types.Sym
} // A LabelStmt is a label statement (just the label, not including the statement it labels).
// "Label:"

type RangeStmt struct {
	miniStmt
	Label         *types.Sym
	Def           bool
	X             Node
	RType         Node `mknode:"-"`
	Key           Node
	Value         Node
	Body          Nodes
	DistinctVars  bool
	Prealloc      *Name
	KeyTypeWord   Node `mknode:"-"`
	KeySrcRType   Node `mknode:"-"`
	ValueTypeWord Node `mknode:"-"`
	ValueSrcRType Node `mknode:"-"`
} // A RangeStmt is a range loop: for Key, Value = range X { Body }
// When desugaring the RangeStmt during walk, the assignments to Key
// and Value may require OCONVIFACE operations. If so, these fields
// will be copied to their respective ConvExpr fields.

type ReturnStmt struct {
	miniStmt
	Results Nodes
} // A ReturnStmt is a return statement.
// return list

type SelectStmt struct {
	miniStmt
	Label *types.Sym
	Cases []*// A SelectStmt is a block: { Cases }.
	CommClause
	Compiled Nodes
} // TODO(rsc): Instead of recording here, replace with a block?
// compiled form, after walkSelect

type SendStmt struct {
	miniStmt
	Chan  Node
	Value Node
} // A SendStmt is a send statement: X <- Y.

type SwitchStmt struct {
	miniStmt
	Tag   Node
	Cases []*// A SwitchStmt is a switch statement: switch Init; Tag { Cases }.
	CaseClause
	Label    *types.Sym
	Compiled Nodes
} // TODO(rsc): Instead of recording here, replace with a block?
// compiled form, after walkSwitch

type TailCallStmt struct {
	miniStmt
	Call *CallExpr
} // A TailCallStmt is a tail call statement, which is used for back-end
// code generation to jump directly to another function entirely.
// the underlying call

type TypeSwitchGuard struct {
	miniNode
	Tag  *Ident
	X    Node
	Used bool
} // A TypeSwitchGuard is the [Name :=] X.(type) in a type switch.

type symsStruct struct {
	AssertE2I         *obj.LSym
	AssertE2I2        *obj.LSym
	Asanread          *obj.LSym
	Asanwrite         *obj.LSym
	CgoCheckMemmove   *obj.LSym
	CgoCheckPtrWrite  *obj.LSym
	CheckPtrAlignment *obj.LSym
	Deferproc         *obj.LSym
	Deferprocat       *obj.LSym
	DeferprocStack    *obj.LSym
	Deferreturn       *obj.LSym
	Duffcopy          *obj.LSym
	Duffzero          *obj.LSym
	GCWriteBarrier    [8]*obj.LSym
	Goschedguarded    *obj.LSym
	Growslice         *obj.LSym
	InterfaceSwitch   *obj.LSym
	MallocGC          *obj.LSym
	Memmove           *obj.LSym
	Msanread          *obj.LSym
	Msanwrite         *obj.LSym
	Msanmove          *obj.LSym
	Newobject         *obj.LSym
	Newproc           *obj.LSym
	Panicdivide       *obj.LSym
	Panicshift        *obj.LSym
	PanicdottypeE     *obj.LSym
	PanicdottypeI     *obj.LSym
	Panicnildottype   *obj.LSym
	Panicoverflow     *obj.LSym
	Racefuncenter     *obj.LSym
	Racefuncexit      *obj.LSym
	Raceread          *obj.LSym
	Racereadrange     *obj.LSym
	Racewrite         *obj.LSym
	Racewriterange    *obj.LSym
	TypeAssert        *obj.LSym
	WBZero            *obj.LSym
	WBMove            *obj.LSym
	SigPanic          *obj.LSym
	Staticuint64s     *obj.LSym
	Typedmemmove      *obj.LSym
	Udiv              *obj.LSym
	WriteBarrier      *obj.LSym
	Zerobase          *obj.LSym
	ZeroVal           *obj.LSym
	ARM64HasATOMICS   *obj.LSym
	ARMHasVFPv4       *obj.LSym
	Loong64HasLAMCAS  *obj.LSym
	Loong64HasLAM_BH  *obj.LSym
	Loong64HasLSX     *obj.LSym
	RISCV64HasZbb     *obj.LSym
	X86HasFMA         *obj.LSym
	X86HasPOPCNT      *obj.LSym
	X86HasSSE41       *obj.LSym
	WasmDiv           *obj.LSym
	WasmTruncS        *obj.LSym
	WasmTruncU        *obj.LSym
} // Wasm
// Wasm

type typeNode struct {
	miniNode
	typ *types.Type
} // A typeNode is a Node wrapper for type t.

type DynamicType struct {
	miniExpr
	RType Node
	ITab  Node
} // A DynamicType represents a type expression whose exact type must be
// computed dynamically.
// ITab is an expression that yields a *runtime.itab value
// representing the asserted type within the assertee expression's
// original interface type.
//
// ITab is only used for assertions (including type switches) from
// non-empty interface type to a concrete (i.e., non-interface)
// type. For all other assertions, ITab is nil.

type nameOff struct {
	n   *ir.Name
	off int64
} // name and offset

type blockArgEffects struct {
	livein  bitvec.BitVec
	liveout bitvec.BitVec
} // variables live at block entry
// variables live at block exit

type argLiveness struct {
	fn   *ir.Func
	f    *ssa.Func
	args []nameOff
	idx  map[ // name and offset of spill slots
	nameOff]int32
	be []blockArgEffects// index in args

	bvset    bvecSet
	blockIdx map[ // indexed by block ID
	// Liveness map indices at each Value (where it changes) and Block entry.
	// During the computation the indices are temporarily index to bvset.
	// At the end they will be index (offset) to the output funcdata (changed
	// in (*argLiveness).emit).
	ssa.ID]int
	valueIdx map[ssa.ID]int
}

type bvecSet struct {
	index []int// bvecSet is a set of bvecs, in initial insertion order.

	uniq []bitvec.// hash -> uniq index. -1 indicates empty slot.
	BitVec
} // unique bvecs, in insertion order

type Interval struct{ st, en int } // Interval hols the range [st,en).

type IntervalsBuilder struct {
	s    Intervals
	lidx int
} // IntervalsBuilder is a helper for constructing intervals based on
// live dataflow sets for a series of BBs where we're making a
// backwards pass over each BB looking for uses and kills. The
// expected use case is:
//
//   - invoke MakeIntervalsBuilder to create a new object "b"
//   - series of calls to b.Live/b.Kill based on a backwards reverse layout
//     order scan over instructions
//   - invoke b.Finish() to produce final set
//
// See the Live method comment for an IR example.
// index of last instruction visited plus 1

type intWithIdx struct {
	i         Interval
	pairIndex int
} // intWithIdx holds an interval i and an index pairIndex storing i's
// position (either 0 or 1) within some previously specified interval
// pair <I1,I2>; a pairIndex of -1 is used to signal "end of
// iteration". Used for Intervals operations, not expected to be
// exported.

type pairVisitor struct {
	cur    intWithIdx
	i1pos  int
	i2pos  int
	i1, i2 Intervals
} // pairVisitor provides a way to visit (iterate through) each interval
// within a pair of Intervals in order of increasing start time. Expected
// usage model:
//
//	func example(i1, i2 Intervals) {
//	  var pairVisitor pv
//	  cur := pv.init(i1, i2);
//	  for !cur.done() {
//	     fmt.Printf("interval %s from i%d", cur.i.String(), cur.pairIndex+1)
//	     cur = pv.nxt()
//	  }
//	}
//
// Used internally for Intervals operations, not expected to be exported.

type MergeLocalsState struct {
	vars []*// MergeLocalsState encapsulates information about which AUTO
	// (stack-allocated) variables within a function can be safely
	// merged/overlapped, e.g. share a stack slot with some other auto).
	// An instance of MergeLocalsState is produced by MergeLocals() below
	// and then consumed in ssagen.AllocFrame. The map 'partition'
	// contains entries of the form <N,SL> where N is an *ir.Name and SL
	// is a slice holding the indices (within 'vars') of other variables
	// that share the same slot, specifically the slot of the first
	// element in the partition, which we'll call the "leader". For
	// example, if a function contains five variables where v1/v2/v3 are
	// safe to overlap and v4/v5 are safe to overlap, the MergeLocalsState
	// content might look like
	//
	//	vars: [v1, v2, v3, v4, v5]
	//	partition: v1 -> [1, 0, 2], v2 -> [1, 0, 2], v3 -> [1, 0, 2]
	//	           v4 -> [3, 4], v5 -> [3, 4]
	//
	// A nil MergeLocalsState indicates that no local variables meet the
	// necessary criteria for overlap.
	// contains auto vars that participate in overlapping
	ir.Name
	partition map[ // maps auto variable to overlap partition
	*ir.Name][]int
}

type candRegion struct{ st, en int } // candRegion is a sub-range (start, end) corresponding to an interval
// [st,en] within the list of candidate variables.

type cstate struct {
	fn    *ir.Func
	f     *ssa.Func
	lv    *Liveness
	cands []*// cstate holds state information we'll need during the analysis
	// phase of stack slot merging but can be discarded when the analysis
	// is done.
	ir.Name
	nameToSlot     map[*ir.Name]int32
	regions        []candRegion
	indirectUE     map[ssa.ID][]*ir.Name
	ivs            []Intervals
	hashDeselected map[*ir.Name]bool
	trace          int
} // debug trace level

type nameCount struct {
	n     *ir.Name
	count int32
}

type blockEffects struct {
	uevar   bitvec.BitVec
	varkill bitvec.BitVec
	livein  bitvec.BitVec
	liveout bitvec.BitVec
} // blockEffects summarizes the liveness effects on an SSA block.
// Computed during Liveness.solve using control flow information:
//
//	livein: variables live at block entry
//	liveout: variables live at block exit

type Liveness struct {
	fn   *ir.Func
	f    *ssa.Func
	vars []*// A collection of global state used by Liveness analysis.
	ir.Name
	idx          map[*ir.Name]int32
	stkptrsize   int64
	be           []blockEffects
	allUnsafe    bool
	unsafePoints bitvec.BitVec
	unsafeBlocks bitvec.BitVec
	livevars     []bitvec.// allUnsafe indicates that all points in this function are
	// unsafe-points.
	// An array with a bit vector for each safe point in the
	// current Block during liveness.epilogue. Indexed in Value
	// order for that block. Additionally, for the entry block
	// livevars[0] is the entry bitmap. liveness.compact moves
	// these to stackMaps.
	BitVec
	livenessMap Map
	stackMapSet bvecSet
	stackMaps   []bitvec.// livenessMap maps from safe points (i.e., CALLs) to their
	// liveness map indexes.
	BitVec
	cache        progeffectscache
	partLiveArgs map[ // partLiveArgs includes input arguments (PPARAM) that may
	// be partially live. That is, it is considered live because
	// a part of it is used, but we may not initialize all parts.
	*ir.Name]bool
	doClobber          bool
	noClobberArgs      bool
	conservativeWrites bool
} // Whether to clobber dead stack slots in this function.
// treat "dead" writes as equivalent to reads during the analysis;
// used only during liveness analysis for stack slot merging (doesn't
// make sense for stackmap analysis).

type Map struct {
	Vals map[ // Map maps from *ssa.Value to StackMapIndex.
	// Also keeps track of unsafe ssa.Values and ssa.Blocks.
	// (unsafe = can't be interrupted during GC.)
	ssa.ID]objw.StackMapIndex
	UnsafeVals   map[ssa.ID]bool
	UnsafeBlocks map[ssa.ID]bool
	DeferReturn  objw.StackMapIndex
} // The set of live, pointer-containing variables at the DeferReturn
// call (only set when open-coded defers are used).

type progeffectscache struct {
	retuevar    []int32
	tailuevar   []int32
	initialized bool
}

type livenessFuncCache struct {
	be          []blockEffects
	livenessMap Map
}

type VersionHeader struct {
	Version   int    `json:"version"`
	Package   string `json:"package"`
	Goos      string `json:"goos"`
	Goarch    string `json:"goarch"`
	GcVersion string `json:"gc_version"`
	File      string `json:"file,omitempty"`
} // LSP requires an enclosing resource, i.e., a file

type Position struct {
	Line      uint `json:"line"`
	Character uint `json:"character"`
} // gopls uses float64, but json output is the same for integers
// gopls uses float64, but json output is the same for integers

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
} // A Range in a text document expressed as (zero-based) start and end positions.
// A range is comparable to a selection in an editor. Therefore the end position is exclusive.
// If you want to specify a range that contains a line including the line ending character(s)
// then use an end position denoting the start of the next line.
// exclusive

type Location struct {
	URI   DocumentURI `json:"uri"`
	Range Range       `json:"range"`
} // A Location represents a location inside a resource, such as a line inside a text file.
// Range is

type DiagnosticRelatedInformation struct {
	Location Location `json:"location"`
	Message  string   `json:"message"`
} /* DiagnosticRelatedInformation defined:
 * Represents a related message and source code location for a diagnostic. This should be
 * used to point to code locations that cause or related to a diagnostics, e.g when duplicating
 * a symbol in a scope.
 */ /*Message defined:
 * The message of this related diagnostic information.
 */

type Diagnostic struct {
	Range    Range              `json:"range"`
	Severity DiagnosticSeverity `json:"severity,omitempty"`
	Code     string             `json:"code,omitempty"`
	Source   string             `json:"source,omitempty"`
	Message  string             `json:"message"`
	Tags     []DiagnosticTag    `json:"tags,omitempty"`/*Diagnostic defined:
	 * Represents a diagnostic, such as a compiler error or warning. Diagnostic objects
	 * are only valid in the scope of a resource.
	 */ /*Tags defined:
	 * Additional metadata about the diagnostic.
	 */
	RelatedInformation []DiagnosticRelatedInformation `json:"relatedInformation,omitempty"`// always empty for logging optimizations.
	/*RelatedInformation defined:
	 * An array of related diagnostic information, e.g. when symbol-names within
	 * a scope collide all definitions can be marked via this property.
	 */
}

type LoggedOpt struct {
	pos          src.XPos
	lastPos      src.XPos
	compilerPass string
	functionName string
	what         string
	target       []interface// A LoggedOpt is what the compiler produces and accumulates,
	// to be converted to JSON for human or IDE consumption.
	// The (non) optimization; "nilcheck", "boundsCheck", "inline", "noInline"
	{

	}
} // Optional target(s) or parameter(s) of "what" -- what was inlined, why it was not, size of copy, etc. 1st is most important/relevant.

type VarAndLoop struct {
	Name    *ir.Name
	Loop    ir.Node
	LastPos src.XPos
} // the *ir.RangeStmt or *ir.ForStmt. Used for identity and position
// the last position observed within Loop

type ImplicitNode interface {
	ir.Node
	SetImplicit(x bool)
}

type gcimports struct {
	ctxt     *types2.Context
	packages map[string]*types2.Package
}

type cycleFinder struct {
	cyclic map[ // A cycleFinder detects anonymous interface cycles (go.dev/issue/56103).
	*types2.Interface]bool
}

type linker struct {
	pw   pkgbits.PkgEncoder
	pkgs map[ // A linker combines a package's stub export data with any referenced
	// elements from imported packages into a single, self-contained
	// export data file.
	string]index
	decls  map[*types.Sym]index
	bodies map[*types.Sym]index
}

type noder struct {
	file      *syntax.File
	linknames []linkname// noder transforms package syntax's AST into a Node tree.

	pragcgobuf [][]string
	err        chan syntax.Error
}

type linkname struct {
	pos    syntax.Pos
	local  string
	remote string
} // linkname records a //go:linkname directive.

type pragmas struct {
	Flag ir.PragmaFlag
	Pos  []pragmaPos// *pragmas is the value stored in a syntax.pragmas during parsing.
	// collected bits

	Embeds []pragmaEmbed// position of each individual flag

	WasmImport *WasmImport
	WasmExport *WasmExport
}

type pragmaPos struct {
	Flag ir.PragmaFlag
	Pos  syntax.Pos
}

type pragmaEmbed struct {
	Pos      syntax.Pos
	Patterns []string
}

type posMap struct {
	bases map[ // A posMap handles mapping from syntax.Pos to src.XPos.
	*syntax.PosBase]*src.PosBase
	cache struct {
		last *syntax.PosBase
		base *src.PosBase
	}
}

type poser interface{ Pos() syntax.Pos }

type ender interface{ End() syntax.Pos }

type pkgReaderIndex struct {
	pr        *pkgReader
	idx       index
	dict      *readerDict
	methodSym *types.Sym
	synthetic func(pos src.XPos, r *reader)
} // A pkgReaderIndex compactly identifies an index (and its
// corresponding dictionary) within a package's export data.

type readerMethodExprInfo struct {
	typeParamIdx int
	method       *types.Sym
}

type methodValueWrapper struct {
	rcvr   *types.Type
	method *types.Field
}

type pkgWriter struct {
	pkgbits.PkgEncoder
	m                     posMap
	curpkg                *types2.Package
	info                  *types2.Info
	rangeFuncBodyClosures map[ // A pkgWriter constructs Unified IR export data from the results of
	// running the types2 type checker on a Go compilation unit.
	*syntax.FuncLit]bool
	posBasesIdx map[ // non-public information, e.g., which functions are closures range function bodies?
	*syntax.PosBase]index
	pkgsIdx   map[*types2.Package]index
	typsIdx   map[types2.Type]index
	objsIdx   map[types2.Object]index
	funDecls  map[*types2.Func]*syntax.FuncDecl
	typDecls  map[*types2.TypeName]typeDeclGen
	linknames map[ // linknames maps package-scope objects to their linker symbol name,
	// if specified by a //go:linkname directive.
	types2.Object]string
	cgoPragmas [][// cgoPragmas accumulates any //go:cgo_* pragmas that need to be
	// passed through to cmd/link.
	]string
}

type writer struct {
	p *pkgWriter
	*pkgbits.Encoder
	sig       *types2.Signature
	localsIdx map[ // A writer provides APIs for writing out an individual element.
	// localsIdx tracks any local variables declared within this
	// function body. It's unused for writing out non-body things.
	*types2.Var]int
	closureVars []posVar// closureVars tracks any free variables that are referenced by this
	// function body. It's unused for writing out non-body things.

	closureVarsIdx map[*types2.Var]int
	dict           *writerDict
	derived        bool
} // index of previously seen free variables
// derived tracks whether the type being written out references any
// type parameters. It's unused for writing non-type things.

type writerDict struct {
	implicits []*// A writerDict tracks types and objects that are used by a declaration.
	// implicits is a slice of type parameters from the enclosing
	// declarations.
	types2.TypeParam
	derived []derivedInfo// derived is a slice of type indices for computing derived types
	// (i.e., types that depend on the declaration's type parameters).

	derivedIdx map[ // derivedIdx maps a Type to its corresponding index within the
	// derived slice, if present.
	types2.Type]index
	typeParamMethodExprs []writerMethodExprInfo// These slices correspond to entries in the runtime dictionary.

	subdicts []objInfo
	rtypes   []typeInfo
	itabs    []itabInfo
}

type itabInfo struct {
	typ   typeInfo
	iface typeInfo
}

type objInfo struct {
	idx       index
	explicits []typeInfo// An objInfo represents a reference to an encoded, instantiated (if
	// applicable) Go object.
	// index for the generic function declaration

} // info for the type arguments

type selectorInfo struct {
	pkgIdx  index
	nameIdx index
} // A selectorInfo represents a reference to an encoded field or method
// name (i.e., objects that can only be accessed using selector
// expressions).

type writerMethodExprInfo struct {
	typeParamIdx int
	methodInfo   selectorInfo
}

type posVar struct {
	pos  syntax.Pos
	var_ *types2.Var
}

type typeDeclGen struct {
	*syntax.TypeDecl
	gen       int
	implicits []*// Implicit type parameters in scope at this type declaration.
	types2.TypeParam
}

type fileImports struct{ importedEmbed, importedUnsafe bool }

type declCollector struct {
	pw         *pkgWriter
	typegen    *int
	file       *fileImports
	withinFunc bool
	implicits  []*// declCollector is a visitor type that collects compiler-needed
	// information about declarations that types2 doesn't track.
	//
	// Notably, it maps declared types and functions back to their
	// declaration statement, keeps track of implicit type parameters, and
	// assigns unique type "generation" numbers to local defined types.
	types2.TypeParam
}

type Progs struct {
	Text    *obj.Prog
	Next    *obj.Prog
	PC      int64
	Pos     src.XPos
	CurFunc *ir.Func
	Cache   []obj.// Progs accumulates Progs for a function and converts them into machine code.
	// fn these Progs are for
	Prog
	CacheIndex int
	NextLive   StackMapIndex
	PrevLive   StackMapIndex
	NextUnsafe bool
	PrevUnsafe bool
} // local progcache
// last emitted unsafe mark

type IRGraph struct {
	IRNodes map[ // IRGraph is a call graph with nodes pointing to IRs of functions and edges
	// carrying weights and callsite information.
	//
	// Nodes for indirect calls may have missing IR (IRNode.AST == nil) if the node
	// is not visible from this package (e.g., not in the transitive deps). Keeping
	// these nodes allows determining the hottest edge from a call even if that
	// callee is not available.
	//
	// TODO(prattmic): Consider merging this data structure with Graph. This is
	// effectively a copy of Graph aggregated to line number and pointing to IR.
	string]*IRNode
}

type IRNode struct {
	AST              *ir.Func
	LinkerSymbolName string
	OutEdges         map[ // IRNode represents a node (function) in the IRGraph.
	// Set of out-edges in the callgraph. The map uniquely identifies each
	// edge based on the callsite and callee, for fast lookup.
	pgo.NamedCallEdge]*IREdge
}

type IREdge struct {
	Src, Dst       *IRNode
	Weight         int64
	CallSiteOffset int
} // IREdge represents a call edge in the IRGraph with source, destination,
// weight, callsite, and line number information.
// Line offset from function start line.

type CallSiteInfo struct {
	LineOffset int
	Caller     *ir.Func
	Callee     *ir.Func
} // CallSiteInfo captures call-site information and its caller/callee.
// Line offset from function start line.

type Profile struct {
	*pgo.Profile
	WeightedCG *IRGraph
} // Profile contains the processed PGO profile and weighted call graph used for
// PGO optimizations.
// WeightedCG represents the IRGraph built from profile, which we will
// update as part of inlining.

type rewriter struct {
	pkg        *types2.Package
	info       *types2.Info
	sig        *types2.Signature
	outer      *syntax.FuncType
	body       *syntax.BlockStmt
	any        types2.Object
	bool       types2.Object
	int        types2.Object
	true       types2.Object
	false      types2.Object
	branchNext map[ // A rewriter implements rewriting the range-over-funcs in a given function.
	// Branch numbering, computed as needed.
	branch]int
	labelLoop map[ // branch -> #next value
	string]*syntax.ForStmt
	stack []syntax.// label -> innermost rangefunc loop it is declared inside (nil for no loop)
	// Stack of nodes being visited.
	Node
	forStack []*// all nodes
	forLoop
	rewritten map[ // range-over-func loops
	*syntax.ForStmt]syntax.Stmt
	declStmt              *syntax.DeclStmt
	nextVar               types2.Object
	defers                types2.Object
	stateVarCount         int
	bodyClosureCount      int
	rangefuncBodyClosures map[ // Declared variables in generated code for outermost loop.
	// to help the debugger, the closures generated for loop bodies get names
	*syntax.FuncLit]bool
}

type branch struct {
	tok   syntax.Token
	label string
} // A branch is a single labeled branch.

type forLoop struct {
	nfor          *syntax.ForStmt
	stateVar      *types2.Var
	stateVarDecl  *syntax.VarDecl
	depth         int
	checkRet      bool
	checkBreak    bool
	checkContinue bool
	checkBranch   []branch// A forLoop describes a single range-over-func loop being processed.
	// add check for "continue" after loop

} // add check for labeled branch after loop

type ptabEntry struct {
	s *types.Sym
	t *types.Type
}

type typeSig struct {
	name  *types.Sym
	isym  *obj.LSym
	tsym  *obj.LSym
	type_ *types.Type
	mtype *types.Type
}

type typeAndStr struct {
	t       *types.Type
	short   string
	regular string
} // "short" here means TypeSymName

type Cursor struct {
	lsym   *obj.LSym
	offset int64
	typ    *types.Type
} // A Cursor represents a typed location inside a static variable where we
// are going to write.

type ArrayCursor struct {
	c Cursor
	n int
} // cursor pointing at first element
// number of elements

type allocator struct {
	name     string
	typ      string
	mak      string
	capacity string
	resize   string
	clear    string
	minLog   int
	maxLog   int
} // name for alloc/free functions
// log_2 of maximum allocation size

type derived struct {
	name string
	typ  string
	base string
} // name for alloc/free functions
// underlying allocator

type arch struct {
	name    string
	pkg     string
	genfile string
	ops     []opData// obj package to import for this arch.
	// source file containing opcode code generation.

	blocks             []blockData
	regnames           []string
	ParamIntRegNames   string
	ParamFloatRegNames string
	gpregmask          regMask
	fpregmask          regMask
	fp32regmask        regMask
	fp64regmask        regMask
	specialregmask     regMask
	framepointerreg    int8
	linkreg            int8
	generic            bool
	imports            []string
}

type opData struct {
	name              string
	reg               regInfo
	asm               string
	typ               string
	aux               string
	rematerializeable bool
	argLength         int32
	commutative       bool
	resultInArg0      bool
	resultNotInArgs   bool
	clobberFlags      bool
	needIntTemp       bool
	call              bool
	tailCall          bool
	nilCheck          bool
	faultOnNilArg0    bool
	faultOnNilArg1    bool
	hasSideEffects    bool
	zeroWidth         bool
	unsafePoint       bool
	fixedReg          bool
	symEffect         string
	scale             uint8
} // default result type
// amd64/386 indexed load scale

type blockData struct {
	name     string
	controls int
	aux      string
} // the suffix for this block ("EQ", "LT", etc.)
// the type of the Aux/AuxInt value, if any

type regInfo struct {
	inputs []regMask// inputs[i] encodes the set of registers allowed for the i'th input.
	// Inputs that don't use registers (flags, memory, etc.) should be 0.

	clobbers regMask
	outputs  []regMask// clobbers encodes the set of registers that are overwritten by
	// the instruction (other than the output registers).
	// outputs[i] encodes the set of registers allowed for the i'th output.

}

type intPair struct{ key, val int } // for sorting a pair of integers by key

type Rule struct {
	Rule string
	Loc  string
} // file name & line number

type unusedInspector struct {
	scope  *scope
	unused map[ // unusedInspector can be used to detect unused variables and imports in an
	// ast.Node via its node method. The result is available in the "unused" map.
	//
	// note that unusedInspector is lazy and best-effort; it only supports the node
	// types and patterns used by the rulegen program.
	// unused is the resulting set of unused declared names, indexed by the
	// starting position of the node that declared the name.
	token.Pos]bool
	defining *object
} // defining is the object currently being defined; this is useful so
// that if "foo := bar" is unused and removed, we can then detect if
// "bar" becomes unused as well.

type scope struct {
	outer   *scope
	objects map[ // scope keeps track of a certain scope and its declared names, as well as the
	// outer (parent) scope.
	// can be nil, if this is the top-level scope
	string]*object
} // indexed by each declared name

type object struct {
	name    string
	pos     token.Pos
	numUses int
	used    []*// object keeps track of a declared name, such as a variable or import.
	// number of times this object is used
	object
} // objects that its declaration makes use of

type Statement interface{} // Statement can be one of our high-level statement struct types, or an
// ast.Stmt under some limited circumstances.

type BodyBase struct {
	List []Statement// BodyBase is shared by all of our statement pseudo-node types which can
	// contain other statements.

	CanFail bool
}

type tokenNode struct {
	pos token.Pos
	end token.Pos
} // tokenNode is a dummy implementation of ast.Node for a single token.
// They are used transiently by PathEnclosingInterval but never escape
// this package.

type application struct {
	pre, post ApplyFunc
	cursor    Cursor
	iter      iterator
} // application carries all the shared data so we can pass it around cheaply.

type biasedSparseMap struct {
	s     *sparseMap
	first int
} // A biasedSparseMap is a sparseMap for integers between J and K inclusive,
// where J might be somewhat larger than zero (and K-J is probably much smaller than J).
// (The motivating use case is the line numbers of statements for a single function.)
// Not all features of a SparseMap are exported, and it is also easy to treat a
// biasedSparseMap like a SparseSet.

type Block struct {
	ID             ID
	Pos            src.XPos
	Kind           BlockKind
	Likely         BranchPrediction
	FlagsLiveAtEnd bool
	Hotness        Hotness
	Succs          []Edge// Block represents a basic block in the control flow graph of a function.
	// Subsequent blocks, if any. The number and order depend on the block kind.

	Preds []Edge// Inverse of successors.
	// The order is significant to Phi nodes in the block.
	// TODO: predecessors is a pain to maintain. Can we somehow order phi
	// arguments by block id and have this field computed explicitly when needed?

	Controls [2]*Value
	Aux      Aux
	AuxInt   int64
	Values   []*// A list of values that determine how the block is exited. The number
	// and type of control values depends on the Kind of the block. For
	// instance, a BlockIf has a single boolean control value and BlockExit
	// has a single memory control value.
	//
	// The ControlValues() method may be used to get a slice with the non-nil
	// control values that can be ranged over.
	//
	// Controls[1] must be nil if Controls[0] is nil.
	// The unordered set of Values that define the operation of this block.
	// After the scheduling pass, this list is ordered.
	Value
	Func        *Func
	succstorage [2]Edge
	predstorage [4]Edge
	valstorage  [9]*Value
} // The containing function
// Storage for Succs, Preds and Values.

type Edge struct {
	b *Block
	i int
} // Edge represents a CFG edge.
// Example edges for b branching to either c or d.
// (c and d have other predecessors.)
//
//	b.Succs = [{c,3}, {d,1}]
//	c.Preds = [?, ?, ?, {b,0}]
//	d.Preds = [?, {b,1}, ?]
//
// These indexes allow us to edit the CFG in constant time.
// In addition, it informs phi ops in degenerate cases like:
//
//	b:
//	   if k then c else c
//	c:
//	   v = Phi(x, y)
//
// Then the indexes tell you whether x is chosen from
// the if or else branch from b.
//
//	b.Succs = [{c,0},{c,1}]
//	c.Preds = [{b,0},{b,1}]
//
// means x is chosen if k is true.
// index of reverse edge.  Invariant:
//   e := x.Succs[idx]
//   e.b.Preds[e.i] = Edge{x,idx}
// and similarly for predecessors.

type Cache struct {
	values          [2000]Value
	blocks          [200]Block
	locs            [2000]Location
	stackAllocState *stackAllocState
	scrPoset        []*// A Cache holds reusable compiler state.
	// It is intended to be re-used for multiple Func compilations.
	// Reusable stackAllocState.
	// See stackalloc.go's {new,put}StackAllocState.
	poset
	regallocValues []valState// scratch poset to be reused
	// Reusable regalloc state.

	ValueToProgAfter []*obj.Prog
	debugState       debugState
	Liveness         interface{}
	hdrValueSlice    []*// *gc.livenessFuncCache
	// Free "headers" for use by the allocators in allocators.go.
	// Used to put slices in sync.Pools without allocation.
	[]*Value
	hdrLimitSlice []*[]limit
}

type pass struct {
	name     string
	fn       func(*Func)
	required bool
	disabled bool
	time     bool
	mem      bool
	stats    int
	debug    int
	test     int
	dump     map[ // report time to run pass
	// pass-specific ad-hoc option, perhaps useful in development
	string]bool
} // dump if function name matches

type constraint struct {
	a, b string
} // Double-check phase ordering constraints.
// This code is intended to document the ordering requirements
// between different phases. It does not override the passes
// list above.
// a must come before b

type Config struct {
	arch           string
	PtrSize        int64
	RegSize        int64
	Types          Types
	lowerBlock     blockRewriter
	lowerValue     valueRewriter
	lateLowerBlock blockRewriter
	lateLowerValue valueRewriter
	splitLoad      valueRewriter
	registers      []Register// A Config holds readonly compilation information.
	// It is created once, early during compilation,
	// and shared across all compilations.
	// function for splitting merged load ops; only used on some architectures

	gpRegMask      regMask
	fpRegMask      regMask
	fp32RegMask    regMask
	fp64RegMask    regMask
	specialRegMask regMask
	intParamRegs   []int8// machine registers
	// special register mask

	floatParamRegs []int8// register numbers of integer param (in/out) registers

	ABI1        *abi.ABIConfig
	ABI0        *abi.ABIConfig
	FPReg       int8
	LinkReg     int8
	hasGReg     bool
	ctxt        *obj.Link
	optimize    bool
	useAvg      bool
	useHmul     bool
	SoftFloat   bool
	Race        bool
	BigEndian   bool
	unalignedOK bool
	haveBswap64 bool
	haveBswap32 bool
	haveBswap16 bool
	mulRecipes  map[ // register numbers of floating param (in/out) registers
	// mulRecipes[x] = function to build v * x from v.
	int64]mulRecipe
}

type mulRecipe struct {
	cost  int
	build func(*Value, *Value) *Value
} // build(m, v) returns v * x built at m.

type Types struct {
	Bool       *types.Type
	Int8       *types.Type
	Int16      *types.Type
	Int32      *types.Type
	Int64      *types.Type
	UInt8      *types.Type
	UInt16     *types.Type
	UInt32     *types.Type
	UInt64     *types.Type
	Int        *types.Type
	Float32    *types.Type
	Float64    *types.Type
	UInt       *types.Type
	Uintptr    *types.Type
	String     *types.Type
	BytePtr    *types.Type
	Int32Ptr   *types.Type
	UInt32Ptr  *types.Type
	IntPtr     *types.Type
	UintptrPtr *types.Type
	Float32Ptr *types.Type
	Float64Ptr *types.Type
	BytePtrPtr *types.Type
} // TODO: use unsafe.Pointer instead?

type Logger interface {
	Logf(string, ...interface{})
	Log() bool
	Fatalf(pos src.XPos, msg string, args ...interface{})
	Warnl(pos src.XPos, fmt_ string, args ...interface{})
	Debug_checknil() bool
} // Logf logs a message from the compiler.
// Forwards the Debug flags from gc

type Frontend interface {
	Logger
	StringData(string) *obj.LSym
	SplitSlot(parent *LocalSlot, suffix string, offset int64, t *types.Type) LocalSlot
	Syslook(string) *obj.LSym
	UseWriteBarrier() bool
	Func() *ir.Func
} // StringData returns a symbol pointing to the given string's contents.
// Func returns the ir.Func of the function being compiled.

type FuncDebug struct {
	Slots []LocalSlot// A FuncDebug contains all the debug information for the variables in a
	// function. Variables are identified by their LocalSlot, which may be
	// the result of decomposing a larger variable.
	// Slots is all the slots used in the debug info, indexed by their SlotID.

	Vars []*// The user variables, indexed by VarID.
	ir.Name
	VarSlots [][// The slots that make up each variable, indexed by VarID.
	]SlotID
	LocationLists [][// The location list data, indexed by VarID. Must be processed by PutLocationList.
	]byte
	RegOutputParams []*// Register-resident output parameters for the function. This is filled in at
	// SSA generation time.
	ir.Name
	OptDcl []*// Variable declarations that were removed during optimization
	ir.Name
	EntryID ID
	GetPC   func(block, value ID) int64
} // The ssa.Func.EntryID value, used to build location lists for
// return values promoted to heap in later DWARF generation.
// Filled in by the user. Translates Block and Value ID to PC.
//
// NOTE: block is only used if value is BlockStart.ID or BlockEnd.ID.
// Otherwise, it is ignored.

type BlockDebug struct {
	startState, endState             abt.T
	lastCheckedTime, lastChangedTime int32
	relevant                         bool
	everProcessed                    bool
} // State at the start and end of the block. These are initialized,
// and updated from new information that flows on back edges.
// false until the block has been processed at least once. This
// affects how the merge is done; the goal is to maximize sharing
// and avoid allocation.

type liveSlot struct{ VarLoc } // A liveSlot is a slot that's live in loc at entry/exit of a block.

type stateAtPC struct {
	slots []VarLoc// stateAtPC is the current state of all variables at some point.
	// The location of each known slot, indexed by SlotID.

	registers [][// The slots present in each register, indexed by register number.
	]SlotID
}

type VarLoc struct {
	Registers RegisterSet
	StackOffset
} // A VarLoc describes the storage for part of a user variable.
// The registers this variable is available in. There can be more than
// one in various situations, e.g. it's being moved between registers.

type debugState struct {
	slots []LocalSlot// See FuncDebug.

	vars     []*ir.Name
	varSlots [][]SlotID
	lists    [][]byte
	slotVars []VarID// The user variable that each slot rolls up to, indexed by SlotID.

	f             *Func
	loggingLevel  int
	convergeCount int
	registers     []Register// testing; iterate over block debug state this many times

	stackOffset func(LocalSlot) int32
	ctxt        *obj.Link
	valueNames  [][// The names (slots) associated with each value, indexed by Value ID.
	]SlotID
	currentState   stateAtPC
	changedVars    *sparseSet
	changedSlots   *sparseSet
	pendingEntries []pendingEntry// The current state of whatever analysis is running.
	// The pending location list entry for each user variable, indexed by VarID.

	varParts        map[*ir.Name][]SlotID
	blockDebug      []BlockDebug
	pendingSlotLocs []VarLoc
}

type slotCanonicalizer struct {
	slmap map[ // slotCanonicalizer is a table used to lookup and canonicalize
	// LocalSlot's in a type insensitive way (e.g. taking into account the
	// base name, offset, and width of the slot, but ignoring the slot
	// type).
	slotKey]SlKeyIdx
	slkeys []LocalSlot
}

type slotKey struct {
	name        *ir.Name
	offset      int64
	width       int64
	splitOf     SlKeyIdx
	splitOffset int64
} // slotKey is a type-insensitive encapsulation of a LocalSlot; it
// is used to key a map within slotCanonicalizer.
// idx in slkeys slice in slotCanonicalizer

type pendingEntry struct {
	present                bool
	startBlock, startValue ID
	pieces                 []VarLoc// A pendingEntry represents the beginning of a location list entry, missing
	// only its end coordinate.
	// The location of each piece of the variable, in the same order as the
	// SlotIDs in varParts.

}

type namedVal struct {
	locIndex, valIndex int
} // f.NamedValues[f.Names[locIndex]][valIndex] = key

type blockAndIndex struct {
	b     *Block
	index int
} // index is the number of successor edges of b that have already been explored.

type registerCursor struct {
	storeDest   *Value
	storeOffset int64
	regs        []abi.// A registerCursor tracks which register is used for an Arg or regValues, or a piece of such.
	// if there are no register targets, then this is the base of the store.
	RegIndex
	nextSlice Abi1RO
	config    *abi.ABIConfig
	regValues *[]*// the registers available for this Arg/result (which is all in registers or not at all)
	// the next register/register-slice offset
	Value
} // values assigned to registers accumulate here

type selKey struct {
	from          *Value
	offsetOrIndex int64
	size          int64
	typ           *types.Type
} // what is selected from
// whatever is appropriate for the selector

type expandState struct {
	f           *Func
	debug       int
	regSize     int64
	sp          *Value
	typs        *Types
	firstOp     Op
	secondOp    Op
	firstType   *types.Type
	secondType  *types.Type
	wideSelects map[ // odd values log lost statement markers, so likely settings are 1 (stmts), 2 (expansion), and 3 (both)
	// second half type, for Int64
	*Value]*Value
	commonSelectors map[ // Selects that are not SSA-able, mapped to consuming stores.
	selKey]*Value
	commonArgs map[ // used to de-dupe selectors
	selKey]*Value
	memForCall map[ // used to de-dupe OpArg/OpArgIntReg/OpArgFloatReg
	ID]*Value
	indentLevel int
} // For a call, need to know the unique selector that gets the mem.
// Indentation for debugging recursion

type LocalSlotSplitKey struct {
	parent *LocalSlot
	Off    int64
	Type   *types.Type
} // offset of slot in N
// type of slot

type HTMLWriter struct {
	w             io.WriteCloser
	Func          *Func
	path          string
	dot           *dotWriter
	prevHash      []byte
	pendingPhases []string
	pendingTitles []string
}

type FuncLines struct {
	Filename    string
	StartLineno uint
	Lines       []string// FuncLines contains source code for a function to be displayed
	// in sources column.

}

type htmlFuncPrinter struct{ w io.Writer }

type dotWriter struct {
	path   string
	broken bool
	phases map[string]bool
} // keys specify phases with CFGs

type idAlloc struct{ last ID } // idAlloc provides an allocator for unique integers.

type lcaRange struct {
	blocks []lcaRangeBlock// lcaRange is a data structure that can compute lowest common ancestor queries
	// in O(n lg n) precomputed space and O(1) time per query.
	// Additional information about each block (indexed by block ID).

	rangeMin [][// Data structure for range minimum queries.
	// rangeMin[k][i] contains the ID of the minimum depth block
	// in the Euler tour from positions i to i+1<<k-1, inclusive.
	]ID
}

type lcaRangeBlock struct {
	b          *Block
	parent     ID
	firstChild ID
	sibling    ID
	pos        int32
	depth      int32
} // parent in dominator tree.  0 = no parent (entry or unreachable)
// depth in dominator tree (root=0, its children=1, etc.)

type loop struct {
	header   *Block
	outer    *loop
	children []*// The header node of this (reducible) loop
	// By default, children, exits, and depth are not initialized.
	loop
	exits []*// loops nested directly within this loop. Initialized by assembleChildren().
	Block
	nBlocks                 int32
	depth                   int16
	isInner                 bool
	containsUnavoidableCall bool
} // exits records blocks reached by exits from this loop. Initialized by findExits().
// True if all paths through the loop have a call

type loopnest struct {
	f                                                       *Func
	b2l                                                     []*loop
	po                                                      []*Block
	sdom                                                    SparseTree
	loops                                                   []*loop
	hasIrreducible                                          bool
	initializedChildren, initializedDepth, initializedExits bool
} // TODO current treatment of irreducible loops is very flaky, if accurate loops are needed, must punt at function level.
// Record which of the lazily initialized fields have actually been initialized.

type Register struct {
	num    int32
	objNum int16
	name   string
} // A Register is a machine register, like AX.
// They are numbered densely from 0 (for each architecture).
// register number from cmd/internal/obj/$ARCH

type LocalSlot struct {
	N           *ir.Name
	Type        *types.Type
	Off         int64
	SplitOf     *LocalSlot
	SplitOffset int64
} // A LocalSlot is a location in the stack frame, which identifies and stores
// part or all of a PPARAM, PPARAMOUT, or PAUTO ONAME node.
// It can represent a whole variable, part of a larger stack slot, or part of a
// variable that has been decomposed into multiple stack slots.
// As an example, a string could have the following configurations:
//
//	          stack layout              LocalSlots
//
//	Optimizations are disabled. s is on the stack and represented in its entirety.
//	[ ------- s string ---- ] { N: s, Type: string, Off: 0 }
//
//	s was not decomposed, but the SSA operates on its parts individually, so
//	there is a LocalSlot for each of its fields that points into the single stack slot.
//	[ ------- s string ---- ] { N: s, Type: *uint8, Off: 0 }, {N: s, Type: int, Off: 8}
//
//	s was decomposed. Each of its fields is in its own stack slot and has its own LocalSLot.
//	[ ptr *uint8 ] [ len int] { N: ptr, Type: *uint8, Off: 0, SplitOf: parent, SplitOffset: 0},
//	                          { N: len, Type: int, Off: 0, SplitOf: parent, SplitOffset: 8}
//	                          parent = &{N: s, Type: string}
// .. at this offset.

type Spill struct {
	Type   *types.Type
	Offset int64
	Reg    int16
}

type indVar struct {
	ind   *Value
	nxt   *Value
	min   *Value
	max   *Value
	entry *Block
	flags indVarFlags
} // induction variable
// entry block in the loop.

type edgeMem struct {
	e Edge
	m *Value
} // an edgeMem records a backedge, together with the memory
// phi functions at the target of the backedge that must
// be updated when a rescheduling check replaces the backedge.
// phi for memory at dest of e

type rewriteTarget struct {
	v *Value
	i int
} // a rewriteTarget is a value-argindex pair indicating
// where a rewrite is applied.  Note that this is for values,
// not for block controls, because block controls are not targets
// for the rewrites performed in inserting rescheduling checks.

type rewrite struct {
	before, after *Value
	rewrites      []rewriteTarget// before is the expected value before rewrite, after is the new value installed.

} // all the targets for this rewrite.

type backedgesState struct {
	b *Block
	i int
}

type umagicData struct {
	s int64
	m uint64
} // ⎡log2(c)⎤
// ⎡2^(n+s)/c⎤ - 2^n

type smagicData struct {
	s int64
	m uint64
} // ⎡log2(c)⎤-1
// ⎡2^(n+s)/c⎤

type udivisibleData struct {
	k   int64
	m   uint64
	max uint64
} // trailingZeros(c)
// ⎣(2^n - 1)/ c⎦ max value to for divisibility

type sdivisibleData struct {
	k   int64
	m   uint64
	a   uint64
	max uint64
} // trailingZeros(c)
// ⎣(2 a) / (1<<k)⎦ max value to for divisibility

type BaseAddress struct {
	ptr *Value
	idx *Value
} // A BaseAddress represents the address ptr+idx, where
// ptr is a pointer type and idx is an integer type.
// idx may be nil, in which case it is treated as 0.

type fileAndPair struct {
	f  int32
	lp lineRange
}

type opInfo struct {
	name              string
	reg               regInfo
	auxType           auxType
	argLen            int32
	asm               obj.As
	generic           bool
	rematerializeable bool
	commutative       bool
	resultInArg0      bool
	resultNotInArgs   bool
	clobberFlags      bool
	needIntTemp       bool
	call              bool
	tailCall          bool
	nilCheck          bool
	faultOnNilArg0    bool
	faultOnNilArg1    bool
	usesScratch       bool
	hasSideEffects    bool
	zeroWidth         bool
	unsafePoint       bool
	fixedReg          bool
	symEffect         SymEffect
	scale             uint8
} // the number of arguments, -1 if variable length
// amd64/386 indexed load scale

type inputInfo struct {
	idx  int
	regs regMask
} // index in Args array
// allowed input registers

type outputInfo struct {
	idx  int
	regs regMask
} // index in output tuple
// allowed output registers

type AuxNameOffset struct {
	Name   *ir.Name
	Offset int64
}

type AuxCall struct {
	Fn      *obj.LSym
	reg     *regInfo
	abiInfo *abi.ABIParamResultInfo
} // regInfo for this call

type Sym interface {
	Aux
	CanBeAnSSASym()
} // A Sym represents a symbolic offset from a base register.
// Currently a Sym can be one of 3 things:
//   - a *ir.Name, for an offset from SP (the stack pointer)
//   - a *obj.LSym, for an offset from SB (the global pointer)
//   - nil, for no offset

type pairableLoadInfo struct {
	width int64
	pair  Op
} // width of one element in the pair, in bytes

type pairableStoreInfo struct {
	width int64
	pair  Op
} // width of one element in the pair, in bytes

type posetUndo struct {
	typ  undoType
	idx  uint32
	ID   ID
	edge posetEdge
} // posetUndo represents an undo pass to be performed.
// It's a union of fields that can be used to store information,
// and typ is the discriminant, that specifies which kind
// of operation must be performed. Not all fields are always used.

type posetNode struct{ l, r posetEdge } // posetNode is a node of a DAG within the poset.

type poset struct {
	lastidx uint32
	flags   uint8
	values  map[ // poset is a union-find data structure that can represent a partially ordered set
	// of SSA values. Given a binary relation that creates a partial order (eg: '<'),
	// clients can record relations between SSA values using SetOrder, and later
	// check relations (in the transitive closure) with Ordered. For instance,
	// if SetOrder is called to record that A<B and B<C, Ordered will later confirm
	// that A<C.
	//
	// It is possible to record equality relations between SSA values with SetEqual and check
	// equality with Equal. Equality propagates into the transitive closure for the partial
	// order so that if we know that A<B<C and later learn that A==D, Ordered will return
	// true for D<C.
	//
	// It is also possible to record inequality relations between nodes with SetNonEqual;
	// non-equality relations are not transitive, but they can still be useful: for instance
	// if we know that A<=B and later we learn that A!=B, we can deduce that A<B.
	// NonEqual can be used to check whether it is known that the nodes are different, either
	// because SetNonEqual was called before, or because we know that they are strictly ordered.
	//
	// poset will refuse to record new relations that contradict existing relations:
	// for instance if A<B<C, calling SetOrder for C<A will fail returning false; also
	// calling SetEqual for C==A will fail.
	//
	// poset is implemented as a forest of DAGs; in each DAG, if there is a path (directed)
	// from node A to B, it means that A<B (or A<=B). Equality is represented by mapping
	// two SSA values to the same DAG node; when a new equality relation is recorded
	// between two existing nodes, the nodes are merged, adjusting incoming and outgoing edges.
	//
	// poset is designed to be memory efficient and do little allocations during normal usage.
	// Most internal data structures are pre-allocated and flat, so for instance adding a
	// new relation does not cause any allocation. For performance reasons,
	// each node has only up to two outgoing edges (like a binary tree), so intermediate
	// "extra" nodes are required to represent more than two relations. For instance,
	// to record that A<I, A<J, A<K (with no known relation between I,J,K), we create the
	// following DAG:
	//
	//	  A
	//	 / \
	//	I  extra
	//	    /  \
	//	   J    K
	// internal flags
	ID]uint32
	nodes []posetNode// map SSA values to dense indexes

	roots []uint32// nodes (in all DAGs)

	noneq map[ // list of root nodes (forest)
	uint32]bitset
	undo []posetUndo// non-equal relations

} // undo chain

type funcPrinter interface {
	header(f *Func)
	startBlock(b *Block, reachable bool)
	endBlock(b *Block, reachable bool)
	value(v *Value, live bool)
	startDepCycle()
	endDepCycle()
	named(n LocalSlot, vals []*Value)
}

type stringFuncPrinter struct {
	w         io.Writer
	printDead bool
}

type limit struct {
	min, max   int64
	umin, umax uint64
} // a limit records known upper and lower bounds for a value.
//
// If we have min>max or umin>umax, then this limit is
// called "unsatisfiable". When we encounter such a limit, we
// know that any code for which that limit applies is unreachable.
// We don't particularly care how unsatisfiable limits propagate,
// including becoming satisfiable, because any optimization
// decisions based on those limits only apply to unreachable code.
// umin <= value <= umax, unsigned

type limitFact struct {
	vid   ID
	limit limit
} // a limitFact is a limit known for a particular value.

type ordering struct {
	next *ordering
	w    *Value
	d    domain
	r    relation
} // An ordering encodes facts like v < w.
// one of ==,!=,<,<=,>,>=

type factsTable struct {
	unsat      bool
	unsatDepth int
	orderS     *poset
	orderU     *poset
	orderings  map[ // factsTable keeps track of relations between pairs of values.
	//
	// The fact table logic is sound, but incomplete. Outside of a few
	// special cases, it performs no deduction or arithmetic. While there
	// are known decision procedures for this, the ad hoc approach taken
	// by the facts table is effective for real code while remaining very
	// efficient.
	// orderings contains a list of known orderings between values.
	// These lists are indexed by v.ID.
	// We do not record transitive orderings. Only explicitly learned
	// orderings are recorded. Transitive orderings can be obtained
	// by walking along the individual orderings.
	ID]*ordering
	orderingsStack []ID// stack of IDs which have had an entry added in orderings.
	// In addition, ID==0 are checkpoint markers.

	orderingCache *ordering
	limits        []limit// unused ordering records
	// known lower and upper constant bounds on individual values.

	limitStack []limitFact// indexed by value ID

	recurseCheck []bool// previous entries

	lens map[ // recursion detector for limit propagation
	// For each slice s, a map from s to a len(s)/cap(s) value (if any)
	// TODO: check if there are cases that matter where we have
	// more than one len(s) for a slice. We could keep a list if necessary.
	ID]*Value
	caps map[ID]*Value
}

type use struct {
	dist int32
	pos  src.XPos
	next *use
} // distance from start of the block to a use of a value
//   dist == 0                 used by first instruction in block
//   dist == len(b.Values)-1   used by last instruction in block
//   dist == len(b.Values)     used by block's control value
//   dist  > len(b.Values)     used by a subsequent block
// linked list of uses of a value in nondecreasing dist order

type valState struct {
	regs              regMask
	uses              *use
	spill             *Value
	restoreMin        int32
	restoreMax        int32
	needReg           bool
	rematerializeable bool
} // A valState records the register allocation state for a (pre-regalloc) value.
// cached value of v.rematerializeable()

type regState struct {
	v *Value
	c *Value
} // Original (preregalloc) Value stored in this register.
// A Value equal to v which is currently in a register.  Might be v or a copy of it.

type regAllocState struct {
	f           *Func
	sdom        SparseTree
	registers   []Register
	numRegs     register
	SPReg       register
	SBReg       register
	GReg        register
	ZeroIntReg  register
	allocatable regMask
	live        [][// live values at the end of each block.  live[b.ID] is a list of value IDs
	// which are live at the end of b, together with a count of how many instructions
	// forward to the next use.
	]liveInfo
	desired []desiredState// desired register assignments at the end of each block.
	// Note that this is a static map computed before allocation occurs. Dynamic
	// register desires (from partially completed allocations) will trump
	// this information.

	values []valState// current state of each (preregalloc) Value

	sp, sb ID
	orig   []*// ID of SP, SB values
	// For each Value, map from its value ID back to the
	// preregalloc Value it was derived from.
	Value
	regs []regState// current state of each register.
	// Includes only registers in allocatable.

	nospill             regMask
	used                regMask
	usedSinceBlockStart regMask
	tmpused             regMask
	curBlock            *Block
	freeUseRecords      *use
	endRegs             [][// registers that contain values which can't be kicked out
	// endRegs[blockid] is the register state at the end of each block.
	// encoded as a set of endReg records.
	]endReg
	startRegs [][// startRegs[blockid] is the register state at the start of merge blocks.
	// saved state does not include the state of phi ops in the block.
	]startReg
	startRegsMask regMask
	spillLive     [][// startRegsMask is a mask of the registers in startRegs[curBlock.ID].
	// Registers dropped from startRegsMask are later synchronoized back to
	// startRegs by dropping from there as well.
	// spillLive[blockid] is the set of live spills at the end of each block
	]ID
	copies map[ // a set of copies we generated to move things around, and
	// whether it is used in shuffle. Unused copies will be deleted.
	*Value]bool
	loopnest   *loopnest
	visitOrder []*// choose a good order in which to visit blocks for allocation purposes.
	Block
	blockOrder []int32// blockOrder[b.ID] corresponds to the index of block b in visitOrder.

	doClobber bool
	nextCall  []int32// whether to insert instructions that clobber dead registers at call sites
	// For each instruction index in a basic block, the index of the next call
	// at or after that instruction index.
	// If there is no next call, returns maxInt32.
	// nextCall for a call instruction points to itself.
	// (Indexes and results are pre-regalloc.)

	curIdx int
} // Index of the instruction we're currently working on.
// Index is expressed in terms of the pre-regalloc b.Values list.

type endReg struct {
	r register
	v *Value
	c *Value
} // pre-regalloc value held in this register (TODO: can we use ID here?)
// cached version of the value

type startReg struct {
	r   register
	v   *Value
	c   *Value
	pos src.XPos
} // pre-regalloc value needed in this register
// source position of use of this register

type edgeState struct {
	s     *regAllocState
	p, b  *Block
	cache map[ // edge goes from p->b.
	// for each pre-regalloc value, a list of equivalent cached values
	ID][]*Value
	cachedVals []ID
	contents   map[ // (superset of) keys of the above map, for deterministic iteration
	// map from location to the value it contains
	Location]contentRecord
	destinations []dstRecord// desired destination locations

	extra                 []dstRecord
	usedRegs              regMask
	uniqueRegs            regMask
	finalRegs             regMask
	rematerializeableRegs regMask
} // registers currently holding something
// registers that hold rematerializeable values

type contentRecord struct {
	vid   ID
	c     *Value
	final bool
	pos   src.XPos
} // pre-regalloc value
// source position of use of the value

type dstRecord struct {
	loc    Location
	vid    ID
	splice **Value
	pos    src.XPos
} // register or stack slot
// source position of use of this location

type liveInfo struct {
	ID   ID
	dist int32
	pos  src.XPos
} // ID of value
// source position of next use

type desiredState struct {
	entries []desiredStateEntry// A desiredState represents desired register assignments.
	// Desired assignments will be small, so we just use a list
	// of valueID+registers entries.

	avoid regMask
} // Registers that other values want to be in.  This value will
// contain at least the union of the regs fields of entries, but
// may contain additional entries for values that were once in
// this data structure but are no longer.

type desiredStateEntry struct {
	ID   ID
	regs [4]register
} // (pre-regalloc) value
// Registers it would like to be in, in priority order.
// Unused slots are filled with noRegister.
// For opcodes that return tuples, we track desired registers only
// for the first element of the tuple (see desiredSecondReg for
// tracking the desired register for second part of a tuple).

type Aux interface{ CanBeAnSSAAux() } // Aux is an interface to hold miscellaneous data in Blocks and Values.

type flagConstantBuilder struct {
	N bool
	Z bool
	C bool
	V bool
}

type lattice struct {
	tag int8
	val *Value
} // lattice type
// constant value

type worklist struct {
	f     *Func
	edges []Edge// the target function to be optimized out

	uses []*// propagate constant facts through edges
	Value
	visited map[ // re-visiting set
	Edge]bool
	latticeCells map[ // visited edges
	*Value]lattice
	defUse map[ // constant lattices
	*Value][]*Value
	defBlock map[ // def-use chains for some values
	*Value][]*Block
	visitedBlock []bool// use blocks of def

} // visited block

type ValHeap struct {
	a           []*Value
	score       []int8
	inBlockUses []bool
}

type sparseEntry struct {
	key ID
	val int32
}

type sparseMap struct {
	dense  []sparseEntry
	sparse []int32
}

type sparseEntryPos struct {
	key ID
	val int32
	pos src.XPos
}

type sparseMapPos struct {
	dense  []sparseEntryPos
	sparse []int32
}

type sparseSet struct {
	dense  []ID
	sparse []int32
}

type SparseTreeNode struct {
	child       *Block
	sibling     *Block
	parent      *Block
	entry, exit int32
} // Every block has 6 numbers associated with it:
// entry-1, entry, entry+1, exit-1, and exit, exit+1.
// entry and exit are conceptually the top of the block (phi functions)
// entry+1 and exit-1 are conceptually the bottom of the block (ordinary defs)
// entry-1 and exit+1 are conceptually "just before" the block (conditions flowing in)
//
// This simplifies life if we wish to query information about x
// when x is both an input to and output of a block.

type stackAllocState struct {
	f    *Func
	live [][// live is the output of stackalloc.
	// live[b.id] = live values at the end of block b.
	]ID
	values []stackValState// The following slices are reused across multiple users
	// of stackAllocState.

	interfere [][]ID
	names     []LocalSlot// interfere[v.id] = values that interfere with v.

	nArgSlot, nNotNeed, nNamedSlot, nReuse, nAuto, nSelfInterfere int32
} // Number of self-interferences

type stackValState struct {
	typ      *types.Type
	spill    *Value
	needSlot bool
	isArg    bool
}

type point struct{ x, y int }

type line struct{ begin, end point }

type big struct{ pile [768]int8 }

type thing struct {
	name  string
	next  *thing
	self  *thing
	stuff []big
}

type ExtNode[V any] struct {
	v V
	Node
}

type List[V any] struct {
	root *ExtNode[V]
	len  int
}

type Value struct {
	ID     ID
	Op     Op
	Type   *types.Type
	AuxInt int64
	Aux    Aux
	Args   []*// A Value represents a value in the SSA representation of the program.
	// The ID and Type fields must not be modified. The remainder may be modified
	// if they preserve the value of the Value (e.g. changing a (mul 2 x) to an (add x x)).
	// Arguments of this value
	Value
	Block       *Block
	Pos         src.XPos
	Uses        int32
	OnWasmStack bool
	InCache     bool
	argstorage  [3]*Value
} // Containing basic block
// Storage for the first three args

type ZeroRegion struct {
	base *Value
	mask uint64
} // A ZeroRegion records parts of an object which are known to be zero.
// A ZeroRegion only applies to a single memory state.
// Each bit in mask is set if the corresponding pointer-sized word of
// the base object is known to be zero.
// In other words, if mask & (1<<i) != 0, then [base+i*ptrSize, base+(i+1)*ptrSize)
// is known to be zero.

type lineRange struct{ first, last uint32 }

type xposmap struct {
	maps map[ // An xposmap is a map from fileindex and line of src.XPos to int32,
	// implemented sparsely to save space (column and statement status are ignored).
	// The sparse skeleton is constructed once, and then reused by ssa phases
	// that (re)move values with statements attached.
	// A map from file index to maps from line range to integers (block numbers)
	int32]*biasedSparseMap
	lastIndex int32
	lastMap   *biasedSparseMap
} // The next two fields provide a single-item cache for common case of repeated lines from same file.
// map found at maps[lastIndex]

type vkey struct {
	op Op
	ai int64
	ax Aux
	t  *types.Type
} // vkey is a type used to uniquely identify a zero arg value.
// type

type SymABIs struct {
	defs map[ // SymABIs records information provided by the assembler about symbol
	// definition ABIs and reference ABIs.
	string]obj.ABI
	refs map[string]obj.ABISet
}

type ArchInfo struct {
	LinkArch      *obj.LinkArch
	REGSP         int
	MAXWIDTH      int64
	SoftFloat     bool
	PadFrame      func(int64) int64
	ZeroRange     func(*objw.Progs, *obj.Prog, int64, int64, *uint32) *obj.Prog
	Ginsnop       func(*objw.Progs) *obj.Prog
	SSAMarkMoves  func(*State, *ssa.Block)
	SSAGenValue   func(*State, *ssa.Value)
	SSAGenBlock   func(s *State, b, next *ssa.Block)
	LoadRegResult func(s *State, f *ssa.Func, t *types.Type, reg int16, n *ir.Name, off int64) *obj.Prog
	SpillArgReg   func(pp *objw.Progs, p *obj.Prog, f *ssa.Func, t *types.Type, reg int16, n *ir.Name, off int64) *obj.Prog
} // ZeroRange zeroes a range of memory on stack. It is only inserted
// at function entry, and it is ok to clobber registers.
// SpillArgReg emits instructions that spill reg to n+off.

type intrinsicKey struct {
	arch *sys.Arch
	pkg  string
	fn   string
}

type intrinsicBuildConfig struct {
	instrumenting bool
	go386         string
	goamd64       int
	goarm         buildcfg.GoarmFeatures
	goarm64       buildcfg.Goarm64Features
	gomips        string
	gomips64      string
	goppc64       int
	goriscv64     int
} // intrinsicBuildConfig specifies the config to use for intrinsic building.

type nowritebarrierrecChecker struct {
	extraCalls map[ // extraCalls contains extra function calls that may not be
	// visible during later analysis. It maps from the ODCLFUNC of
	// the caller to a list of callees.
	*ir.Func][]nowritebarrierrecCall
	curfn *ir.Func
} // curfn is the current function during AST walks.

type nowritebarrierrecCall struct {
	target *ir.Func
	lineno src.XPos
} // caller or callee
// line of call

type largeStack struct {
	locals int64
	args   int64
	callee int64
	pos    src.XPos
} // largeStack is info about a function whose stack frame is too large (rare).

type fwdRefAux struct {
	_ [0]func()
	N ir.Node
} // fwdRefAux wraps an arbitrary ir.Node as an ssa.Aux for use with OpFwdref.
// ensure ir.Node isn't compared for equality

type phiState struct {
	s       *state
	f       *ssa.Func
	defvars []map// SSA state
	// function to work on
	[ir.Node]*ssa.Value
	varnum map[ // defined variables at end of each block
	ir.Node]int32
	idom []*// variable numbering
	// properties of the dominator tree
	ssa.Block
	tree []domBlock// dominator parents

	level []int32// dominator child+sibling

	priq blockHeap
	q    []*// level in dominator tree (0 = root or unreachable, 1 = children of root, ...)
	// priority queue of blocks, higher level (toward leaves) = higher priority
	ssa.Block
	queued      *sparseSet
	hasPhi      *sparseSet
	hasDef      *sparseSet
	placeholder *ssa.Value
} // inner loop queue
// value to use as a "not set yet" placeholder.

type domBlock struct {
	firstChild *ssa.Block
	sibling    *ssa.Block
} // domBlock contains extra per-block information to record the dominator tree.
// next child of parent in dominator tree

type blockHeap struct {
	a []*// A block heap is used as a priority queue to implement the PiggyBank
	// from Sreedhar and Gao.  That paper uses an array which is better
	// asymptotically but worse in the common case when the PiggyBank
	// holds a sparse set of blocks.
	ssa.Block
	level []int32// block IDs in heap

} // depth in dominator tree (static, used for determining priority)

type simplePhiState struct {
	s       *state
	f       *ssa.Func
	fwdrefs []*// Variant to use for small functions.
	// function to work on
	ssa.Value
	defvars []map// list of FwdRefs to be processed
	[ir.Node]*ssa.Value
	reachable []bool// defined variables at end of each block

} // which blocks are reachable

type openDeferInfo struct {
	n           *ir.CallExpr
	closure     *ssa.Value
	closureNode *ir.Name
} // Information about each open-coded defer.
// The node representing the argtmp where the closure is stored - used for
// function, method, or interface call, to store a closure that panic
// processing can use for this defer.

type state struct {
	config *ssa.Config
	f      *ssa.Func
	curfn  *ir.Func
	labels map[ // configuration (arch) information
	// labels in f
	string]*ssaLabel
	breakTo    *ssa.Block
	continueTo *ssa.Block
	curBlock   *ssa.Block
	vars       map[ // unlabeled break and continue statement tracking
	// variable assignments in the current block (map from variable symbol to ssa value)
	// *Node is the unique identifier (an ONAME Node) for the variable.
	// TODO: keep a single varnum map, then make all of these maps slices instead?
	ir.Node]*ssa.Value
	fwdVars map[ // fwdVars are variables that are used before they are defined in the current block.
	// This map exists just to coalesce multiple references into a single FwdRef op.
	// *Node is the unique identifier (an ONAME Node) for the variable.
	ir.Node]*ssa.Value
	defvars []map// all defined variables at the end of each block. Indexed by block ID.
	[ir.Node]*ssa.Value
	decladdrs map[ // addresses of PPARAM and PPARAMOUT variables on the stack.
	*ir.Name]*ssa.Value
	startmem      *ssa.Value
	sp            *ssa.Value
	sb            *ssa.Value
	deferBitsAddr *ssa.Value
	deferBitsTemp *ir.Name
	line          []src.// starting values. Memory, stack pointer, and globals pointer
	// line number stack. The current line number is top of stack
	XPos
	lastPos src.XPos
	panics  map[ // the last line number processed; it may have been popped
	// list of panic calls by function name and line number.
	// Used to deduplicate panic calls.
	funcLine]*ssa.Block
	cgoUnsafeArgs       bool
	hasdefer            bool
	softFloat           bool
	hasOpenDefers       bool
	checkPtrEnabled     bool
	instrumentEnterExit bool
	instrumentMemory    bool
	openDefers          []*// whether the function contains a defer statement
	// If doing open-coded defers, list of info about the defer calls in
	// scanning order. Hence, at exit we should run these defers in reverse
	// order of this list
	openDeferInfo
	lastDeferExit          *ssa.Block
	lastDeferFinalBlock    *ssa.Block
	lastDeferCount         int
	prevCall               *ssa.Value
	pendingHeapAllocations []*// For open-coded defers, this is the beginning and end blocks of the last
	// defer exit code that we have generated so far. We use these to share
	// code between exits if the shareDeferExits option (disabled by default)
	// is on.
	// List of allocations in the current block that are still pending.
	// They are all (OffPtr (Select0 (runtime call))) and have the correct types,
	// but the offsets are not set yet, and the type of the runtime call is also not final.
	ssa.Value
	appendTargets map[ // First argument of append calls that could be stack allocated.
	ir.Node]bool
}

type funcLine struct {
	f    *obj.LSym
	base *src.PosBase
	line uint
}

type ssaLabel struct {
	target         *ssa.Block
	breakTarget    *ssa.Block
	continueTarget *ssa.Block
} // block identified by this label
// block to continue to in control flow node identified by this label

type opAndType struct {
	op    ir.Op
	etype types.Kind
}

type opAndTwoTypes struct {
	op     ir.Op
	etype1 types.Kind
	etype2 types.Kind
}

type twoTypes struct {
	etype1 types.Kind
	etype2 types.Kind
}

type twoOpsAndType struct {
	op1              ssa.Op
	op2              ssa.Op
	intermediateType types.Kind
}

type sfRtCallDef struct {
	rtfn  *obj.LSym
	rtype types.Kind
}

type u642fcvtTab struct {
	leq, cvt2F, and, rsh, or, add ssa.Op
	one                           func(*state, *types.Type, int64) *ssa.Value
}

type u322fcvtTab struct{ cvtI2F, cvtF2F ssa.Op }

type f2uCvtTab struct {
	ltf, cvt2U, subf, or ssa.Op
	floatValue           func(*state, *types.Type, float64) *ssa.Value
	intValue             func(*state, *types.Type, int64) *ssa.Value
	cutoff               uint64
}

type Branch struct {
	P *obj.Prog
	B *ssa.Block
} // Branch is an unresolved branch.
// target

type IndexJump struct {
	Jump  obj.As
	Index int
} // For generating consecutive jump instructions to model a specific branching

type ssafn struct {
	curfn   *ir.Func
	strings map[ // ssafn holds frontend information about a function that the backend is processing.
	// It also exports a bunch of compiler services for the ssa backend.
	string]*obj.LSym
	stksize    int64
	stkptrsize int64
	stkalign   int64
	log        bool
} // map from constant string to data symbols
// print ssa debug to the stdout

type Entry struct {
	Xoffset int64
	Expr    ir.Node
} // struct, array only
// bytes of run-time computed expressions

type Plan struct{ E []Entry }

type Schedule struct {
	Out []ir.// An Schedule is used to decompose assignment statements into
	// static and dynamic initialization parts. Static initializations are
	// handled by populating variables' linker symbol data, while dynamic
	// initializations are accumulated to be executed in order.
	// Out is the ordered list of dynamic initialization
	// statements.
	Node
	Plans        map[ir.Node]*Plan
	Temps        map[ir.Node]*ir.Name
	seenMutation bool
} // seenMutation tracks whether we've seen an initialization
// expression that may have modified other package-scope variables
// within this package.

type labelScope struct {
	errh   ErrorHandler
	labels map[string]*label
} // all label declarations inside the function; allocated lazily

type label struct {
	parent *block
	lstmt  *LabeledStmt
	used   bool
} // block containing this label declaration
// whether the label is used or not

type block struct {
	parent *block
	start  Pos
	lstmt  *LabeledStmt
} // immediately enclosing block, or nil
// labeled statement associated with this block, or nil

type targets struct {
	breaks    Stmt
	continues *ForStmt
	caseIndex int
} // targets describes the target statements within which break
// or continue statements are valid.
// case index of immediately enclosing switch statement, or < 0

type writeError struct{ err error } // writeError wraps locally caught write errors so we can distinguish
// them from genuine panics which we don't want to return as errors.

type node struct{ pos Pos }

type decl struct{ node }

type Group struct {
	_ int
} // All declarations belonging to the same group point to the same Group node.
// not empty so we are guaranteed different Group instances

type expr struct {
	node
	typeAndValue
} // After typechecking, contains the results of typechecking this expression.

type (
	RangeClause struct {
		Lhs Expr
		Def bool
		X   Expr
		simpleStmt
	}
	CaseClause struct {
		Cases Expr
		Body  []Stmt// nil means no Lhs = or Lhs :=
		// nil means default clause

		Colon Pos
		node
	}
	CommClause struct {
		Comm SimpleStmt
		Body []Stmt// send or receive stmt; nil means default clause

		Colon Pos
		node
	}
)

type stmt struct{ node }

type simpleStmt struct{ stmt }

type Comment struct {
	Kind CommentKind
	Text string
	Next *Comment
}

type parser struct {
	file  *PosBase
	errh  ErrorHandler
	mode  Mode
	pragh PragmaHandler
	scanner
	base      *PosBase
	first     error
	errcnt    int
	pragma    Pragma
	goVersion string
	top       bool
	fnest     int
	xnest     int
	indent    []byte// current position base
	// expression nesting level (for complit ambiguity resolution)

} // tracing support

type Pos struct {
	base      *PosBase
	line, col uint32
} // A Pos represents an absolute (line, col) source position
// with a reference to position base for computing relative
// (to a file, or line directive) position information.
// Pos values are intentionally light-weight so that they
// can be created without too much concern about space use.

type position_ struct {
	filename  string
	line, col uint
} // TODO(gri) cleanup: find better name, avoid conflict with position in error_test.go

type PosBase struct {
	pos       Pos
	filename  string
	line, col uint32
	trimmed   bool
} // A PosBase represents the base for relative position information:
// At position pos, the relative position is filename:line:col.
// whether -trimpath has been applied

type whitespace struct {
	last token
	kind ctrlSymbol
}

type printer struct {
	output     io.Writer
	written    int
	form       Form
	linebreaks bool
	indent     int
	nlcount    int
	pending    []whitespace// number of bytes written
	// number of consecutive newlines

	lastTok token
} // pending whitespace
// last token (after any pending semi) processed by print

type printGroup struct {
	node
	Tok   token
	Decls []Decl
}

type scanner struct {
	source
	mode      uint
	nlsemi    bool
	line, col uint
	blank     bool
	tok       token
	lit       string
	bad       bool
	kind      LitKind
	op        Operator
	prec      int
} // if set '\n' and EOF translate to ';'
// valid if tok is _Operator, _Star, _AssignOp, or _IncOp

type source struct {
	in        io.Reader
	errh      func(line, col uint, msg string)
	buf       []byte
	ioerr     error
	b, r, e   int
	line, col uint
	ch        rune
	chw       int
} // source buffer
// width of ch

type Pragma interface{} // A Pragma value augments a package, import, const, func, type, or var declaration.
// Its meaning is entirely up to the PragmaHandler,
// except that nil is used to mean “no pragma seen.”

type Sender[T any] struct {
	values chan<- T
	done   <-chan // A sender is used to send values to a Receiver.
	bool
}

type Receiver[T any] struct {
	values <-chan // A Receiver receives values from a Sender.
	T
	done chan<- bool
}

type _ interface {
	m()
	E
}

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | float32 | ~float64 | complex64 | ~complex128
} // Numeric is type bound that matches any numeric type.
// It would likely be in a constraints package in the standard library.

type NumericAbs[T any] interface {
	Numeric
	Abs() T
} // NumericAbs matches numeric types with an Abs method.

type OrderedNumeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | float32 | ~float64
} // OrderedNumeric is a type bound that matches numeric types that support the < operator.

type Complex interface{ ~complex64 | ~complex128 } // Complex is a type bound that matches the two complex types, which do not have a < operator.

type keyValue[K, V any] struct {
	key K
	val V
} // keyValue is a pair of key and value used when iterating.

type chans_Sender[T any] struct {
	values chan<- T
	done   <-chan // A sender is used to send values to a Receiver.
	bool
}

type chans_Receiver[T any] struct {
	values <-chan T
	done   chan<- bool
}

type TypeAndValue struct {
	Type  Type
	Value constant.Value
	exprFlags
} // A TypeAndValue records the type information, constant
// value if known, and various other flags associated with
// an expression.
// This type is similar to types2.TypeAndValue, but exposes
// none of types2's internals.

type typeAndValue struct{ tv TypeAndValue } // a typeAndValue contains the results of typechecking an expression.
// It is embedded in expression nodes.

type Visitor interface{ Visit(node Node) (w Visitor) } // A Visitor's Visit method is invoked for each node encountered by Walk.
// If the result visitor w is not nil, Walk visits each of the children
// of node with the visitor w, followed by a call of w.Visit(nil).

type walker struct{ v Visitor }

type T55357[T any] struct{}

type blo struct {
	inc   int64
	cond  bool
	succs [2]int64
} // blo describes a block in the generated/interpreted code
// block ends in conditional

type tmplData struct{ Name, Stype, Symbol string } // used for interpolation in a text template

type sizedTestData struct {
	name string
	sn   string
	u    []uint64
	i    []int64
}

type op struct{ name, symbol string }

type szD struct {
	name   string
	sn     string
	u      []uint64
	i      []int64
	oponly string
}

type Pointery struct {
	p *Pointery
	x [1024]int
} // This type and the following one will share the same GC shape and size.

type Pointery2 struct {
	p *Pointery2
	x [1024]int
}

type Vanilla struct {
	np uintptr
	x  [1024]int
} // This type and the following one will have the same size.

type Vanilla2 struct {
	np uintptr
	x  [1023]int
	y  int
}

type Single struct {
	np uintptr
	x  [1023]int
}

type LessConstraint[T any] interface{ Less(T) bool }

type Adder interface{ Add(a, b int) int }

type Add struct{}

type Sub struct{}

type Multiplier interface{ Multiply(a, b int) int }

type Mult struct{}

type NegMult struct{}

type BS struct {
	length uint
	s      []uint64
}

type MyString struct{ string }

type methods struct {
	m1 func(a *A, x string, y int) string
	m2 func(a *A, x string, y int) string
}

type ClosureStructIter struct {
	closureVars []*// ClosureStructIter iterates through a slice of closure variables returning
	// their type and offset in the closure struct.
	ir.Name
	offset int64
	next   int
}

type typeInterner struct {
	typs []string// typeInterner maps Go type expressions to compiler code that
	// constructs the denoted type. It recognizes and reuses common
	// subtype expressions.

	hash map[string]int
}

type typeSet struct{ m map[string]src.XPos }

type dlist struct{ field *types.Field } // A dlist stores a pointer to a TFIELD Type embedded within
// a TSTRUCT or TINTER Type.

type symlink struct{ field *types.Field }

type lang struct{ major, minor int } // A lang is a language version broken into major and minor numbers.

type typePair struct {
	t1 *Type
	t2 *Type
}

type Pkg struct {
	Path   string
	Name   string
	Prefix string
	Syms   map[ // string literal used in import statement, e.g. "internal/runtime/sys"
	// escaped path for use in symbol table
	string]*Sym
	Pathsym *obj.LSym
	Direct  bool
} // imported directly

type Object interface {
	Pos() src.XPos
	Sym() *Sym
	Type() *Type
} // Object represents an ir.Node, but without needing to import cmd/compile/internal/ir,
// which would cause an import cycle. The uses in other packages must type assert
// values of type Object to ir.Node or a more specific type.

type Forward struct {
	Copyto []*// Forward contains Type fields specific to forward types.
	Type
	Embedlineno src.XPos
} // where to copy the eventual value to
// first use of this type as an embedded type

type Struct struct {
	fields     fields
	Map        *Type
	ParamTuple bool
} // StructType contains Type fields specific to struct types.
// whether this struct is actually a tuple of signature parameters

type Ptr struct {
	Elem *Type
} // Ptr contains Type fields specific to pointer types.
// element type

type ChanArgs struct {
	T *Type
} // ChanArgs contains Type fields specific to TCHANARGS types.
// reference to a chan type whose elements need a width check

type FuncArgs struct {
	T *Type
} // FuncArgs contains Type fields specific to TFUNCARGS types.
// reference to a func type whose elements need a width check

type Chan struct {
	Elem *Type
	Dir  ChanDir
} // Chan contains Type fields specific to channel types.
// channel direction

type Tuple struct {
	first  *Type
	second *Type
}

type Results struct {
	Types []*// Results are the output from calls that will be late-expanded.
	Type
} // Last element is memory output from call.

type Array struct {
	Elem  *Type
	Bound int64
} // Array contains Type fields specific to array types.
// number of elements; <0 if unknown yet

type Field struct {
	flags    bitset8
	Embedded uint8
	Pos      src.XPos
	Sym      *Sym
	Type     *Type
	Note     string
	Nname    Object
	Offset   int64
} // A Field is a (Sym, Type) pairing along with some other information, and,
// depending on the context, is used to represent:
//   - a field in a struct
//   - a method in an interface or associated with a named type
//   - a function parameter
// Offset in bytes of this field or method within its enclosing struct
// or interface Type. For parameters, this is BADWIDTH.

type fields struct {
	s *[]*// fields is a pointer to a slice of *Field.
	// This saves space in Types that do not have fields or methods
	// compared to a simple slice of *Field.
	Field
}

type Alias struct {
	obj     *TypeName
	orig    *Alias
	tparams *TypeParamList
	targs   *TypeList
	fromRHS Type
	actual  Type
} // An Alias represents an alias type.
//
// Alias types are created by alias declarations such as:
//
//	type A = int
//
// The type on the right-hand side of the declaration can be accessed
// using [Alias.Rhs]. This type may itself be an alias.
// Call [Unalias] to obtain the first non-alias type in a chain of
// alias type declarations.
//
// Like a defined ([Named]) type, an alias type has a name.
// Use the [Alias.Obj] method to access its [TypeName] object.
//
// Historically, Alias types were not materialized so that, in the example
// above, A's type was represented by a Basic (int), not an Alias
// whose [Alias.Rhs] is int. But Go 1.24 allows you to declare an
// alias type with type parameters or arguments:
//
//	type Set[K comparable] = map[K]bool
//	s := make(Set[String])
//
// and this requires that Alias types be materialized. Use the
// [Alias.TypeParams] and [Alias.TypeArgs] methods to access them.
//
// To ease the transition, the Alias type was introduced in go1.22,
// but the type-checker would not construct values of this type unless
// the GODEBUG=gotypesalias=1 environment variable was provided.
// Starting in go1.23, this variable is enabled by default.
// This setting also causes the predeclared type "any" to be
// represented as an Alias, not a bare [Interface].
// actual (aliased) type; never an alias

type ArgumentError struct {
	Index int
	Err   error
} // An ArgumentError holds an error associated with an argument index.

type Importer interface {
	Import(path string) (*Package, error)
} // An Importer resolves import paths to Packages.
//
// CAUTION: This interface does not support the import of locally
// vendored packages. See https://golang.org/s/go15vendor.
// If possible, external implementations should implement ImporterFrom.
// Import returns the imported package for the given import path.
// The semantics is like for ImporterFrom.ImportFrom except that
// dir and mode are ignored (since they are not present).

type ImporterFrom interface {
	Importer
	ImportFrom(path, dir string, mode ImportMode) (*Package, error)
} // An ImporterFrom resolves import paths to packages; it
// supports vendoring per https://golang.org/s/go15vendor.
// Use go/importer to obtain an ImporterFrom implementation.
// ImportFrom returns the imported package for the given import
// path when imported by a package file located in dir.
// If the import failed, besides returning an error, ImportFrom
// is encouraged to cache and return a package anyway, if one
// was created. This will reduce package inconsistencies and
// follow-on type checker errors due to the missing package.
// The mode value must be 0; it is reserved for future use.
// Two calls to ImportFrom with the same path and dir must
// return the same package.

type Info struct {
	Types map[ // Info holds result type information for a type-checked package.
	// Only the information for which a map is provided is collected.
	// If the package has type errors, the collected information may
	// be incomplete.
	// Types maps expressions to their types, and for constant
	// expressions, also their values. Invalid expressions are
	// omitted.
	//
	// For (possibly parenthesized) identifiers denoting built-in
	// functions, the recorded signatures are call-site specific:
	// if the call result is not a constant, the recorded type is
	// an argument-specific signature. Otherwise, the recorded type
	// is invalid.
	//
	// The Types map does not record the type of every identifier,
	// only those that appear where an arbitrary expression is
	// permitted. For instance:
	// - an identifier f in a selector expression x.f is found
	//   only in the Selections map;
	// - an identifier z in a variable declaration 'var z int'
	//   is found only in the Defs map;
	// - an identifier p denoting a package in a qualified
	//   identifier p.X is found only in the Uses map.
	//
	// Similarly, no type is recorded for the (synthetic) FuncType
	// node in a FuncDecl.Type field, since there is no corresponding
	// syntactic function type expression in the source in this case
	// Instead, the function type is found in the Defs.map entry for
	// the corresponding function declaration.
	syntax.Expr]TypeAndValue
	StoreTypesInSyntax bool
	Instances          map[ // If StoreTypesInSyntax is set, type information identical to
	// that which would be put in the Types map, will be set in
	// syntax.Expr.TypeAndValue (independently of whether Types
	// is nil or not).
	// Instances maps identifiers denoting generic types or functions to their
	// type arguments and instantiated type.
	//
	// For example, Instances will map the identifier for 'T' in the type
	// instantiation T[int, string] to the type arguments [int, string] and
	// resulting instantiated *Named type. Given a generic function
	// func F[A any](A), Instances will map the identifier for 'F' in the call
	// expression F(int(1)) to the inferred type arguments [int], and resulting
	// instantiated *Signature.
	//
	// Invariant: Instantiating Uses[id].Type() with Instances[id].TypeArgs
	// results in an equivalent of Instances[id].Type.
	*syntax.Name]Instance
	Defs map[ // Defs maps identifiers to the objects they define (including
	// package names, dots "." of dot-imports, and blank "_" identifiers).
	// For identifiers that do not denote objects (e.g., the package name
	// in package clauses, or symbolic variables t in t := x.(type) of
	// type switch headers), the corresponding objects are nil.
	//
	// For an embedded field, Defs returns the field *Var it defines.
	//
	// Invariant: Defs[id] == nil || Defs[id].Pos() == id.Pos()
	*syntax.Name]Object
	Uses map[ // Uses maps identifiers to the objects they denote.
	//
	// For an embedded field, Uses returns the *TypeName it denotes.
	//
	// Invariant: Uses[id].Pos() != id.Pos()
	*syntax.Name]Object
	Implicits map[ // Implicits maps nodes to their implicitly declared objects, if any.
	// The following node and object types may appear:
	//
	//     node               declared object
	//
	//     *syntax.ImportDecl    *PkgName for imports without renames
	//     *syntax.CaseClause    type-specific *Var for each type switch case clause (incl. default)
	//     *syntax.Field         anonymous parameter *Var (incl. unnamed results)
	//
	syntax.Node]Object
	Selections map[ // Selections maps selector expressions (excluding qualified identifiers)
	// to their corresponding selections.
	*syntax.SelectorExpr]*Selection
	Scopes map[ // Scopes maps syntax.Nodes to the scopes they define. Package scopes are not
	// associated with a specific node but with all files belonging to a package.
	// Thus, the package scope can be found in the type-checked Package object.
	// Scopes nest, with the Universe scope being the outermost scope, enclosing
	// the package scope, which contains (one or more) files scopes, which enclose
	// function scopes which in turn enclose statement and function literal scopes.
	// Note that even though package-level functions are declared in the package
	// scope, the function scopes are embedded in the file scope of the file
	// containing the function declaration.
	//
	// The Scope of a function contains the declarations of any
	// type parameters, parameters, and named results, plus any
	// local declarations in the body block.
	// It is coextensive with the complete extent of the
	// function's syntax ([*ast.FuncDecl] or [*ast.FuncLit]).
	// The Scopes mapping does not contain an entry for the
	// function body ([*ast.BlockStmt]); the function's scope is
	// associated with the [*ast.FuncType].
	//
	// The following node types may appear in Scopes:
	//
	//     *syntax.File
	//     *syntax.FuncType
	//     *syntax.TypeDecl
	//     *syntax.BlockStmt
	//     *syntax.IfStmt
	//     *syntax.SwitchStmt
	//     *syntax.CaseClause
	//     *syntax.CommClause
	//     *syntax.ForStmt
	//
	syntax.Node]*Scope
	InitOrder []*// InitOrder is the list of package-level initializers in the order in which
	// they must be executed. Initializers referring to variables related by an
	// initialization dependency appear in topological order, the others appear
	// in source order. Variables without an initialization expression do not
	// appear in this list.
	Initializer
	FileVersions map[ // FileVersions maps a file to its Go version string.
	// If the file doesn't specify a version, the reported
	// string is Config.GoVersion.
	// Version strings begin with “go”, like “go1.21”, and
	// are suitable for use with the [go/version] package.
	*syntax.PosBase]string
}

type Instance struct {
	TypeArgs *TypeList
	Type     Type
} // Instance reports the type arguments and instantiated type for type and
// function instantiations. For type instantiations, Type will be of dynamic
// type *Named. For function instantiations, Type will be of dynamic type
// *Signature.

type Initializer struct {
	Lhs []*// An Initializer describes a package-level variable, or a list of variables in case
	// of a multi-valued initialization expression, and the corresponding initialization
	// expression.
	Var
	Rhs syntax.Expr
} // var Lhs = Rhs

type Basic struct {
	kind BasicKind
	info BasicInfo
	name string
} // A Basic represents a basic type.

type exprInfo struct {
	isLhs bool
	mode  operandMode
	typ   *Basic
	val   constant.Value
} // exprInfo stores information about an untyped expression.
// constant value; or nil (if not a constant)

type environment struct {
	decl         *declInfo
	scope        *Scope
	version      goVersion
	iota         constant.Value
	errpos       syntax.Pos
	inTParamList bool
	sig          *Signature
	isPanic      map[ // An environment represents the environment within which an object is
	// type-checked.
	// function signature if inside a function; nil otherwise
	*syntax.CallExpr]bool
	hasLabel      bool
	hasCallOrRecv bool
} // set of panic call expressions (used for termination check)
// set if an expression contains a function call or channel receive operation

type importKey struct{ path, dir string } // An importKey identifies an imported package by import path and source directory
// (directory containing the file containing the import). In practice, the directory
// may always be the same, or may not matter. Given an (import path, directory), an
// importer must always return the same package (but given two different import paths,
// an importer may still return the same package by mapping them to the same package
// paths).

type dotImportKey struct {
	scope *Scope
	name  string
} // A dotImportKey describes a dot-imported object in the given scope.

type action struct {
	version goVersion
	f       func()
	desc    *actionDesc
} // An action describes a (delayed) action.
// action description; may be nil, requires debug to be set

type actionDesc struct {
	pos    poser
	format string
	args   []interface// An actionDesc provides information on an action.
	// For debugging only.
	{

	}
}

type Checker struct {
	conf *Config
	ctxt *Context
	pkg  *Package
	*Info
	nextID uint64
	objMap map[ // A Checker maintains the state of the type checker.
	// It must be created with NewChecker.
	// unique Id for type parameters (first valid Id is 1)
	Object]*declInfo
	impMap map[ // maps package-level objects and (non-interface) methods to declaration info
	importKey]*Package
	pkgPathMap map[ // maps (import path, source directory) to (complete or fake) package
	// pkgPathMap maps package names to the set of distinct import paths we've
	// seen for that name, anywhere in the import graph. It is used for
	// disambiguating package names in error messages.
	//
	// pkgPathMap is allocated lazily, so that we don't pay the price of building
	// it on the happy path. seenPkgMap tracks the packages that we've already
	// walked.
	string]map[string]bool
	seenPkgMap map[*Package]bool
	files      []*// information collected during type-checking of a set of package files
	// (initialized by Files, valid only for the duration of check.Files;
	// maps and lists are allocated on demand)
	syntax.File
	versions map[ // list of package files
	*syntax.PosBase]string
	imports []*// maps files to version strings (each file has an entry); shared with Info.FileVersions if present; may be unaltered Config.GoVersion
	PkgName
	dotImportMap map[ // list of imported packages
	dotImportKey]*PkgName
	brokenAliases map[ // maps dot-imported objects to the package they were dot-imported through
	*TypeName]bool
	unionTypeSets map[ // set of aliases with broken (not yet determined) types
	*Union]*_TypeSet
	usedVars map[ // computed type sets for union types
	*Var]bool
	usedPkgNames map[ // set of used variables
	*PkgName]bool
	mono     monoGraph
	firstErr error
	methods  map[ // set of used package names
	// first error encountered
	*TypeName][]*Func
	untyped map[ // maps package scope type names to associated non-blank (non-interface) methods
	syntax.Expr]exprInfo
	delayed []action// map of expressions without final type

	objPath []Object// stack of delayed action segments; segments are processed in FIFO order

	cleaners []cleaner// path of object dependencies during type inference (for cycle reporting)

	environment
	posStack []syntax.// list of types that may need a final cleanup at the end of type-checking
	// debugging
	Pos
	indent int
} // stack of source positions seen; used for panic tracing
// indentation for tracing

type cleaner interface{ cleanup() }

type bailout struct{} // A bailout panic is used for early termination.

type ctxtEntry struct {
	orig     Type
	targs    []Type
	instance Type
} // = orig[targs]

type errorDesc struct {
	pos syntax.Pos
	msg string
} // An errorDesc describes part of a type-checking error.

type error_ struct {
	check *Checker
	desc  []errorDesc// An error_ represents a type-checking error.
	// A new error_ is created with Checker.newError.
	// To report an error_, call error_.report.

	code Code
	soft bool
} // TODO(gri) eventually determine this from an error code

type target struct {
	sig  *Signature
	desc string
} // target represent the (signature) type and description of the LHS
// variable of an assignment, or of a function result variable.

type gcSizes struct {
	WordSize int64
	MaxAlign int64
} // word size in bytes - must be >= 4 (32bits)
// maximum alignment in bytes - must be >= 1

type tpWalker struct {
	tparams []*TypeParam
	seen    map[Type]bool
}

type dependency interface {
	Object
	isDependency()
} // A dependency is an object that may be a dependency in an initialization
// expression. Only constants, variables, and functions can be dependencies.
// Constants are here because constant expression cycles are reported during
// initialization order computation.

type graphNode struct {
	obj        dependency
	pred, succ nodeSet
	index      int
	ndeps      int
} // A graphNode represents a node in the object dependency graph.
// Each node p in n.pred represents an edge p->n, and each node
// s in n.succ represents an edge n->s; with a->b indicating that
// a depends on b.
// number of outstanding dependencies before this object can be initialized

type genericType interface {
	Type
	TypeParams() *TypeParamList
} // A genericType implements access to its type parameters.

type embeddedType struct {
	typ   Type
	index []int// embeddedType represents an embedded type

	indirect  bool
	multiples bool
} // embedded field indices, starting with index at depth 0
// if set, typ appears multiple times at this depth

type instanceLookup struct {
	buf [3]*Named
	m   map[ // buf is used to avoid allocating the map m in the common case of a small
	// number of instances.
	*Named][]*Named
}

type monoGraph struct {
	vertices []monoVertex
	edges    []monoEdge
	canon    map[ // canon maps method receiver type parameters to their respective
	// receiver type's type parameters.
	*TypeParam]*TypeParam
	nameIdx map[ // nameIdx maps a defined type or (canonical) type parameter to its
	// vertex index.
	*TypeName]int
}

type monoVertex struct {
	weight int
	pre    int
	len    int
	obj    *TypeName
} // weight of heaviest known path to this vertex
// obj is the defined type or type parameter represented by this
// vertex.

type monoEdge struct {
	dst, src int
	weight   int
	pos      syntax.Pos
	typ      Type
}

type Named struct {
	check      *Checker
	obj        *TypeName
	fromRHS    Type
	inst       *instance
	mu         sync.Mutex
	state_     uint32
	underlying Type
	tparams    *TypeParamList
	methods    []*// A Named represents a named (defined) type.
	//
	// A declaration such as:
	//
	//	type S struct { ... }
	//
	// creates a defined type whose underlying type is a struct,
	// and binds this type to the object S, a [TypeName].
	// Use [Named.Underlying] to access the underlying type.
	// Use [Named.Obj] to obtain the object S.
	//
	// Before type aliases (Go 1.9), the spec called defined types "named types".
	// methods declared for this type (not the method set of this type)
	// Signatures are type-checked lazily.
	// For non-instantiated types, this is a fully populated list of methods. For
	// instantiated types, methods are individually expanded when they are first
	// accessed.
	Func
	loader func(*Named) (tparams []*// loader may be provided to lazily load type parameters, underlying type, and methods.
	TypeParam, underlying Type, methods []*Func)
}

type instance struct {
	orig            *Named
	targs           *TypeList
	expandedMethods int
	ctxt            *Context
} // instance holds information that is only necessary for instantiated named
// types.
// local Context; set to nil after full expansion

type PkgName struct {
	object
	imported *Package
} // A PkgName represents an imported Go package.
// PkgNames don't have a type.

type Const struct {
	object
	val constant.Value
} // A Const represents a declared constant.

type TypeName struct{ object } // A TypeName is an [Object] that represents a type with a name:
// a defined type ([Named]),
// an alias type ([Alias]),
// a type parameter ([TypeParam]),
// or a predeclared type such as int or error.

type Var struct {
	object
	origin   *Var
	kind     VarKind
	embedded bool
} // A Variable represents a declared variable (including function parameters and results, and struct fields).
// if set, the variable is an embedded struct field, and name is the type name

type Label struct {
	object
	used bool
} // A Label represents a declared label.
// Labels don't have a type.
// set if the label was used

type Builtin struct {
	object
	id builtinId
} // A Builtin represents a built-in function.
// Builtins don't have a valid type.

type Nil struct{ object } // Nil represents the predeclared value nil.

type operand struct {
	mode operandMode
	expr syntax.Expr
	typ  Type
	val  constant.Value
	id   builtinId
} // An operand represents an intermediate value during type checking.
// Operands have an (addressing) mode, the expression evaluating to
// the operand, the operand's type, a value for constants, and an id
// for built-in functions.
// The zero value of operand is a ready to use invalid operand.

type Pointer struct {
	base Type
} // A Pointer represents a pointer type.
// element type

type ifacePair struct {
	x, y *Interface
	prev *ifacePair
} // An ifacePair is a node in a stack of interface type pairs compared for identity.

type comparer struct {
	ignoreTags     bool
	ignoreInvalids bool
} // A comparer is used to compare types.
// if set, identical treats an invalid type as identical to any type

type declInfo struct {
	file    *Scope
	version goVersion
	lhs     []*// A declInfo describes a package-level const, type, var, or func declaration.
	// Go version of file containing this declaration
	Var
	vtyp      syntax.Expr
	init      syntax.Expr
	inherited bool
	tdecl     *syntax.TypeDecl
	fdecl     *syntax.FuncDecl
	deps      map[ // lhs of n:1 variable declarations, or nil
	// The deps field tracks initialization expression dependencies.
	Object]bool
} // lazily initialized

type Scope struct {
	parent   *Scope
	children []*// A Scope maintains a set of objects and links to its containing
	// (parent) and contained (children) scopes. Objects may be inserted
	// and looked up by name. The zero value for Scope is a ready-to-use
	// empty scope.
	Scope
	number int
	elems  map[ // parent.children[number-1] is this scope; 0 if there is no parent
	string]Object
	pos, end syntax.Pos
	comment  string
	isFunc   bool
} // lazily allocated
// set if this is a function scope (internal use only)

type lazyObject struct {
	parent  *Scope
	resolve func() Object
	obj     Object
	once    sync.Once
} // A lazyObject represents an imported Object that has not been fully
// resolved yet by its importer.

type Selection struct {
	kind  SelectionKind
	recv  Type
	obj   Object
	index []int// A Selection describes a selector expression x.f.
	// For the declarations:
	//
	//	type T struct{ x int; E }
	//	type E struct{}
	//	func (e E) m() {}
	//	var p *T
	//
	// the following relations exist:
	//
	//	Selector    Kind          Recv    Obj    Type       Index     Indirect
	//
	//	p.x         FieldVal      T       x      int        {0}       true
	//	p.m         MethodVal     *T      m      func()     {1, 0}    true
	//	T.m         MethodExpr    T       m      func(T)    {1, 0}    false
	// object denoted by x.f

	indirect bool
} // path from x to x.f
// set if there was any pointer indirection on the path

type Signature struct {
	rparams  *TypeParamList
	tparams  *TypeParamList
	scope    *Scope
	recv     *Var
	params   *Tuple
	results  *Tuple
	variadic bool
} // A Signature represents a (non-builtin) function or method type.
// The receiver is ignored when comparing signatures for identity.
// true if the last parameter's type is of the form ...T (or string, for append built-in only)

type Sizes interface {
	Alignof(T Type) int64
	Offsetsof(fields []*// Sizes defines the sizing functions for package unsafe.
	// Offsetsof returns the offsets of the given struct fields, in bytes.
	// Offsetsof must implement the offset guarantees required by the spec.
	// A negative entry in the result indicates that the struct is too large.
	Var) []int64
	Sizeof(T Type) int64
} // Sizeof returns the size of a variable of type T.
// Sizeof must implement the size guarantees required by the spec.
// A negative result indicates that T is too large.

type StdSizes struct {
	WordSize int64
	MaxAlign int64
} // StdSizes is a convenience type for creating commonly used Sizes.
// It makes the following simplifying assumptions:
//
//   - The size of explicitly sized basic types (int16, etc.) is the
//     specified size.
//   - The size of strings and interfaces is 2*WordSize.
//   - The size of slices is 3*WordSize.
//   - The size of an array of n elements corresponds to the size of
//     a struct of n consecutive fields of the array's element type.
//   - The size of a struct is the offset of the last field plus that
//     field's size. As with all element types, if the struct is used
//     in an array its size must first be aligned to a multiple of the
//     struct's alignment.
//   - All other types have size WordSize.
//   - Arrays and structs are aligned per spec definition; all other
//     types are naturally aligned with a maximum alignment MaxAlign.
//
// *StdSizes implements Sizes.
// maximum alignment in bytes - must be >= 1

type subster struct {
	pos       syntax.Pos
	smap      substMap
	check     *Checker
	expanding *Named
	ctxt      *Context
} // nil if called via Instantiate
// if non-nil, the instance that is being expanded

type TypeParamList struct {
	tparams []*// TypeParamList holds a list of type parameters.
	TypeParam
}

type TypeList struct {
	types []Type// TypeList holds a list of types.
}

type TypeParam struct {
	check *Checker
	id    uint64
	obj   *TypeName
	index int
	bound Type
} // A TypeParam represents the type of a type parameter in a generic declaration.
//
// A TypeParam has a name; use the [TypeParam.Obj] method to access
// its [TypeName] object.
// any type, but underlying is eventually *Interface for correct programs (see TypeParam.iface)

type _TypeSet struct {
	methods []*// A _TypeSet represents the type set of an interface.
	// Because of existing language restrictions, methods can be "factored out"
	// from the terms. The actual type set is the intersection of the type set
	// implied by the methods and the type set described by the terms and the
	// comparable bit. To test whether a type is included in a type set
	// ("implements" relation), the type must implement all methods _and_ be
	// an element of the type set described by the terms and the comparable bit.
	// If the term list describes the set of all types and comparable is true,
	// only comparable types are meant; in all other cases comparable is false.
	Func
	terms      termlist
	comparable bool
} // all methods of the interface; sorted by unique ID
// invariant: !comparable || terms.isAll()

type typeWriter struct {
	buf          *bytes.Buffer
	seen         map[Type]bool
	qf           Qualifier
	ctxt         *Context
	tparams      *TypeParamList
	paramNames   bool
	tpSubscripts bool
	pkgInfo      bool
} // if non-nil, we are type hashing
// package-annotate first unexported-type field to avoid confusing type description

type term struct {
	tilde bool
	typ   Type
} // A term describes elementary type sets:
//
//	 ∅:  (*term)(nil)     == ∅                      // set of no types (empty set)
//	 𝓤:  &term{}          == 𝓤                      // set of all types (𝓤niverse)
//	 T:  &term{false, T}  == {T}                    // set of type T
//	~t:  &term{true, t}   == {t' | under(t') == t}  // set of types with underlying type t
// valid if typ != nil

type typeError struct {
	format_ string
	args    []any// A typeError describes a type error.

}

type unifier struct {
	handles map[ // A unifier maintains a list of type parameters and
	// corresponding types inferred for each type parameter.
	// A unifier is created by calling newUnifier.
	// handles maps each type parameter to its inferred type through
	// an indirection *Type called (inferred type) "handle".
	// Initially, each type parameter has its own, separate handle,
	// with a nil (i.e., not yet inferred) type.
	// After a type parameter P is unified with a type parameter Q,
	// P and Q share the same handle (and thus type). This ensures
	// that inferring the type for a given type parameter P will
	// automatically infer the same type for all other parameters
	// unified (joined) with P.
	*TypeParam]*Type
	depth                    int
	enableInterfaceInference bool
} // recursion depth during unification
// use shared methods for better inference

type Union struct {
	terms []*// A Union represents a union of terms embedded in an interface.
	Term
} // list of syntactical terms (not a canonicalized termlist)

type hookInfo struct {
	paramType   types.Kind
	argsNum     int
	runtimeFunc string
}

type orderState struct {
	out []ir.// orderState holds state during the ordering process.
	Node
	temp []*// list of generated statements
	ir.Name
	free map[ // stack of temporary variables
	string][]*ir.Name
	edit func(ir.Node) ir.Node
} // free list of unused temporaries, by type.LinkString().
// cached closure of o.exprNoLHS

type exprSwitch struct {
	pos      src.XPos
	exprname ir.Node
	done     ir.Nodes
	clauses  []exprClause// An exprSwitch walks an expression switch.
	// value being switched on

}

type exprClause struct {
	pos    src.XPos
	lo, hi ir.Node
	rtype  ir.Node
	jmp    ir.Node
} // *runtime._type for OEQ node

type typeSwitch struct {
	srcName  ir.Node
	hashName ir.Node
	okName   ir.Node
	itabName ir.Node
} // A typeSwitch walks a type switch.
// itab value to use for first word of non-empty interface

type typeClause struct {
	hash uint32
	body ir.Nodes
}

type argvalues struct {
	osargs []string
	goos   string
	goarch string
}

type argstate struct {
	state       argvalues
	initialized bool
}

type covOperation interface {
	cov.CovDataVisitor
	Setup()
	Usage(string)
}

type dstate struct {
	calloc.BatchCounterAlloc
	cm     *cmerge.Merger
	format *cformat.Formatter
	mm     map[ // dstate encapsulates state and provides methods for implementing
	// various dump operations. Specifically, dstate implements the
	// CovDataVisitor interface, and is designed to be used in
	// concert with the CovDataReader utility, which abstracts away most
	// of the grubby details of reading coverage data files.
	// 'mm' stores values read from a counter data file; the pkfunc key
	// is a pkgid/funcid pair that uniquely identifies a function in
	// instrumented application.
	pkfunc]decodecounter.FuncPayload
	pkm map[ // pkm maps package ID to the number of functions in the package
	// with that ID. It is used to report inconsistencies in counter
	// data (for example, a counter data entry with pkgid=N funcid=10
	// where package N only has 3 functions).
	uint32]uint32
	pkgpaths map[ // pkgpaths records all package import paths encountered while
	// visiting coverage data files (used to implement the "pkglist"
	// subcommand).
	string]struct{}
	pkgName                  string
	pkgImportPath            string
	modulePath               string
	cmd                      string
	textfmtoutf              *os.File
	totalStmts, coveredStmts int
	preambleEmitted          bool
} // Current package name and import path.
// Records whether preamble has been emitted for current pkg
// (used when in "debugdump" mode)

type mstate struct{ mm *metaMerge } // mstate encapsulates state and provides methods for implementing the
// merge operation. This type implements the CovDataVisitor interface,
// and is designed to be used in concert with the CovDataReader
// utility, which abstracts away most of the grubby details of reading
// coverage data files. Most of the heavy lifting for merging is done
// using apis from 'metaMerge' (this is mainly a wrapper around that
// functionality).

type metaMerge struct {
	calloc.BatchCounterAlloc
	cmerge.Merger
	pkm map[ // metaMerge provides state and methods to help manage the process
	// of selecting or merging meta data files. There are three cases
	// of interest here: the "-pcombine" flag provided by merge, the
	// "-pkg" option provided by all merge/subtract/intersect, and
	// a regular vanilla merge with no package selection
	//
	// In the -pcombine case, we're essentially glomming together all the
	// meta-data for all packages and all functions, meaning that
	// everything we see in a given package needs to be added into the
	// meta-data file builder; we emit a single meta-data file at the end
	// of the run.
	//
	// In the -pkg case, we will typically emit a single meta-data file
	// per input pod, where that new meta-data file contains entries for
	// just the selected packages.
	//
	// In the third case (vanilla merge with no combining or package
	// selection) we can carry over meta-data files without touching them
	// at all (only counter data files will be merged).
	// maps package import path to package state
	string]*pkstate
	pkgs []*// list of packages
	pkstate
	p      *pkstate
	pod    *podstate
	astate *argstate
} // current package state
// counter data file osargs/goos/goarch state

type pkstate struct {
	pkgIdx uint32
	ctab   map[ // pkstate
	// this maps function index within the package to counter data payload
	uint32]decodecounter.FuncPayload
	mdblob []byte// pointer to meta-data blob for package

	*pcombinestate
} // filled in only for -pcombine merges

type podstate struct {
	pmm      map[pkfunc]decodecounter.FuncPayload
	mdf      string
	mfr      *decodemeta.CoverageMetaFileReader
	fileHash [16]byte
}

type pkfunc struct{ pk, fcn uint32 }

type pcombinestate struct {
	cmdb *encodemeta.CoverageMetaDataBuilder
	ftab map[ // pcombinestate
	// Maps function meta-data hash to new function index in the
	// new version of the package we're building.
	[16]byte]uint32
}

type sstate struct {
	mm    *metaMerge
	inidx int
	mode  string
	imm   map[ // sstate holds state needed to implement subtraction and intersection
	// operations on code coverage data files. This type provides methods
	// to implement the CovDataVisitor interface, and is designed to be
	// used in concert with the CovDataReader utility, which abstracts
	// away most of the grubby details of reading coverage data files.
	// Used only for intersection; keyed by pkg/fn ID, it keeps track of
	// just the set of functions for which we have data in the current
	// input directory.
	pkfunc]struct{}
}

type block1 struct {
	Block
	index int
}

type pos2 struct{ p1, p2 token.Position } // pos2 is a pair of token.Position values, used as a map key type.

type FuncExtent struct {
	name      string
	startLine int
	startCol  int
	endLine   int
	endCol    int
} // FuncExtent describes a function's extent in the source by file and position.

type FuncVisitor struct {
	fset    *token.FileSet
	name    string
	astFile *ast.File
	funcs   []*// FuncVisitor implements the visitor that builds the function position list for a file.
	// Name of file.
	FuncExtent
}

type templateData struct {
	Files []*templateFile
	Set   bool
}

type templateFile struct {
	Name     string
	Body     template.HTML
	Coverage float64
}

type exprParser struct {
	x string
	t exprToken
} // exprParser is a //go:build expression parser and evaluator.
// The parser is a trivial precedence-based parser which is still
// almost overkill for these very simple expressions.
// upcoming token

type exprToken struct {
	tok    string
	prec   int
	prefix func(*exprParser) val
	infix  func(val, val) val
} // exprToken describes a single token in the input.
// Prefix operators define a prefix func that parses the
// upcoming value. Binary operators define an infix func
// that combines two values according to the operator.
// In that case, the parsing loop parses the two values.

type importReader struct {
	b    *bufio.Reader
	buf  []byte
	peek byte
	err  error
	eof  bool
	nerr int
}

type systeminfo struct {
	wProcessorArchitecture      uint16
	wReserved                   uint16
	dwPageSize                  uint32
	lpMinimumApplicationAddress uintptr
	lpMaximumApplicationAddress uintptr
	dwActiveProcessorMask       uintptr
	dwNumberOfProcessors        uint32
	dwProcessorType             uint32
	dwAllocationGranularity     uint32
	wProcessorLevel             uint16
	wProcessorRevision          uint16
} // see https://learn.microsoft.com/en-us/windows/win32/api/sysinfoapi/ns-sysinfoapi-system_info

type tester struct {
	race        bool
	msan        bool
	asan        bool
	listMode    bool
	rebuild     bool
	failed      bool
	keepGoing   bool
	compileOnly bool
	runRxStr    string
	runRx       *regexp.Regexp
	runRxWant   bool
	runNames    []string// tester executes cmdtest.
	// want runRx to match (true) or not match (false)

	banner      string
	lastHeading string
	short       bool
	cgoEnabled  bool
	json        bool
	tests       []distTest// tests to run, exclusive with runRx; empty means all
	// last dir heading printed

	testNames map[ // use addTest to extend
	string]bool
	timeoutScale int
	worklist     []*work
}

type work struct {
	dt    *distTest
	cmd   *exec.Cmd
	flush func()
	start chan bool
	out   bytes.Buffer
	err   error
	end   chan struct{}
} // work tracks command execution for a test.
// a value means cmd ended (or was skipped)

type distTest struct {
	name    string
	heading string
	fn      func(*distTest) error
} // A distTest is a test run by dist test.
// Each test has a unique name and belongs to a group (heading)
// group section; this header is printed before the test is run.

type goTest struct {
	timeout time.Duration
	short   bool
	tags    []string// goTest represents all options to a "go test" command. The final command will
	// combine configuration from goTest and tester flags.
	// If true, force -short

	race      bool
	bench     bool
	runTests  string
	cpu       string
	skip      string
	gcflags   string
	ldflags   string
	buildmode string
	env       []string// Build tags
	// If non-empty, -buildmode flag

	runOnHost   bool
	variant     string
	omitVariant bool
	pkgs        []string// Environment variables to add, as KEY=VAL. KEY= unsets a variable
	// We have both pkg and pkgs as a convenience. Both may be set, in which
	// case they will be combined. At least one must be set.

	pkg       string
	testFlags []string// Multiple packages to test
	// A single package to test

} // Additional flags accepted by this test

type registerTestOpt interface{ isRegisterTestOpt() }

type rtSkipFunc struct {
	skip func(*distTest) (string, bool)
} // rtSkipFunc is a registerTest option that runs a skip check function before
// running the test.
// Return message, true to skip the test

type lockedWriter struct {
	lock sync.Mutex
	w    io.Writer
} // lockedWriter serializes Write calls to an underlying Writer.

type testJSONFilter struct {
	w       io.Writer
	variant string
	lineBuf bytes.Buffer
} // testJSONFilter is an io.Writer filter that replaces the Package field in
// test2json output.
// Buffer for incomplete lines

type jsonValue struct {
	atom json.Token
	seq  []jsonValue// If json.Delim, then seq will be filled

} // If atom == json.Delim('{'), alternating pairs

type Archive struct {
	Files []File// An Archive describes an archive to write: a collection of files.
	// Directories are implied by the files and not explicitly listed.
}

type fileInfo struct{ f *File } // A fileInfo is an implementation of fs.FileInfo describing a File.

type testRule struct {
	name    string
	goos    string
	exclude bool
}

type fix struct {
	name     string
	date     string
	f        func(*ast.File) bool
	desc     string
	disabled bool
} // date that fix was introduced, in YYYY-MM-DD format
// whether this fix should be disabled by default

type TypeConfig struct {
	Type     map[string]*Type
	Var      map[string]string
	Func     map[string]string
	External map[ // External maps from a name to its type.
	// It provides additional typings not present in the Go source itself.
	// For now, the only additional typings are those generated by cgo.
	string]string
}

type netrcLine struct {
	machine  string
	login    string
	password string
}

type Command struct {
	Run func(ctx context.Context, cmd *Command, args []string)// A Command is an implementation of a go command
	// like go build or go fix.
	// Run runs the command.
	// The args are the arguments after the command name.

	UsageLine   string
	Short       string
	Long        string
	Flag        flag.FlagSet
	CustomFlags bool
	Commands    []*// UsageLine is the one-line usage message.
	// The words between "go" and the first flag or argument in the line are taken to be the command name.
	// Commands lists the available commands and help topics.
	// The order here is the order in which they are printed by 'go help'.
	// Note that subcommands are in general best avoided.
	Command
}

type explicitStringFlag struct {
	value    *string
	explicit *bool
} // explicitStringFlag is like a regular string flag, but it also tracks whether
// the string was set explicitly to a non-empty value.

type boolFlag interface {
	flag.Value
	IsBoolFlag() bool
} // boolFlag is the optional interface for flag.Value known to the flag package.
// (It is not clear why package flag does not export this interface.)

type netTokenChecker struct {
	released                 bool
	unusedAvoidTinyAllocator string
} // We want to use a finalizer to check that all acquired tokens are returned,
// so we arbitrarily pad the tokens with a string to defeat the runtime's
// “tiny allocator”.

type DiskCache struct {
	dir string
	now func() time.Time
} // A Cache is a package cache, backed by a file system directory tree.

type entryNotFoundError struct{ Err error } // An entryNotFoundError indicates that a cache entry was not found, with an
// optional underlying reason.

type noVerifyReadSeeker struct{ io.ReadSeeker } // noVerifyReadSeeker is an io.ReadSeeker wrapper sentinel type
// that says that Cache.Put should skip the verify check
// (from GODEBUG=goverifycache=1).

type Hash struct {
	h    hash.Hash
	name string
	buf  *bytes.Buffer
} // A Hash provides access to the canonical hash function used to index the cache.
// The current implementation uses salted SHA256, but clients must not assume this.
// for verify

type ProgCache struct {
	cmd    *exec.Cmd
	stdout io.ReadCloser
	stdin  io.WriteCloser
	bw     *bufio.Writer
	jenc   *json.Encoder
	can    map[ // ProgCache implements Cache via JSON messages over stdin/stdout to a child
	// helper process which can then implement whatever caching policy/mechanism it
	// wants.
	//
	// See https://github.com/golang/go/issues/59719
	// can are the commands that the child process declared that it supports.
	// This is effectively the versioning mechanism.
	cacheprog.Cmd]bool
	fuzzDirCache Cache
	closing      atomic.Bool
	ctx          context.Context
	ctxCancel    context.CancelFunc
	readLoopDone chan struct{}
	mu           sync.Mutex
	nextID       int64
	inFlight     map[ // fuzzDirCache is another Cache implementation to use for the FuzzDir
	// method. In practice this is the default GOCACHE disk-based
	// implementation.
	//
	// TODO(bradfitz): maybe this isn't ideal. But we'd need to extend the Cache
	// interface and the fuzzing callers to be less disk-y to do more here.
	// guards following fields
	int64]chan<- *cacheprog.Response
	outputFile map[OutputID]string
	writeMu    sync.Mutex
} // object => abs path on disk
// writeMu serializes writing to the child process.
// It must never be held at the same time as mu.

type Request struct {
	ID       int64
	Command  Cmd
	ActionID []byte `json:",omitempty"`// Request is the JSON-encoded message that's sent from the go command to
	// the GOCACHEPROG child process over stdin. Each JSON object is on its own
	// line. A ProgRequest of Type "put" with BodySize > 0 will be followed by a
	// line containing a base64-encoded JSON string literal of the body.
	// ActionID is the cache key for "put" and "get" requests.

	OutputID []byte `json:",omitempty"`// or nil if not used
	// OutputID is stored with the body for "put" requests.

	Body     io.Reader `json:"-"`
	BodySize int64     `json:",omitempty"`
} // or nil if not used
// BodySize is the number of bytes of Body. If zero, the body isn't written.

type Response struct {
	ID            int64
	Err           string `json:",omitempty"`
	KnownCommands []Cmd  `json:",omitempty"`// Response is the JSON response from the child process to the go command.
	//
	// With the exception of the first protocol message that the child writes to its
	// stdout with ID==0 and KnownCommands populated, these are only sent in
	// response to a Request from the go command.
	//
	// Responses can be sent in any order. The ID must match the request they're
	// replying to.
	// KnownCommands is included in the first message that cache helper program
	// writes to stdout on startup (with ID==0). It includes the
	// Request.Command types that are supported by the program.
	//
	// This lets the go command extend the protocol gracefully over time (adding
	// "get2", etc), or fail gracefully when needed. It also lets the go command
	// verify the program wants to be a cache helper.

	Miss     bool   `json:",omitempty"`
	OutputID []byte `json:",omitempty"`// cache miss

	Size     int64      `json:",omitempty"`
	Time     *time.Time `json:",omitempty"`
	DiskPath string     `json:",omitempty"`
} // the OutputID stored with the body
// DiskPath is the absolute path on disk of the body corresponding to a
// "get" (on cache hit) or "put" request's ActionID.

type EnvVar struct {
	Name    string
	Value   string
	Changed bool
} // An EnvVar is an environment variable Name=Value.
// effective Value differs from default

type buildXContextKey struct{}

type dirInfo struct{ dir fs.DirEntry } // A dirInfo implements fs.FileInfo from fs.DirEntry.
// We know that go/build doesn't use the non-DirEntry parts,
// so we can panic instead of doing difficult work.

type FlagNotDefinedError struct {
	RawArg   string
	Name     string
	HasValue bool
	Value    string
} // A FlagNotDefinedError indicates a flag-like argument that does not correspond
// to any registered flag in a FlagSet.
// only provided if HasValue is true

type NonFlagError struct{ RawArg string } // A NonFlagError indicates an argument that is not a syntactically-valid flag.

type overlayJSON struct {
	Replace map[ // overlayJSON is the format for the -overlay file.
	string]string
}

type replace struct {
	from string
	to   string
} // A replace represents a single replaced path.
// to is the replacement for the old path.
// It is an absolute path returned by abs.
// If it is the empty string, the old path appears deleted.
// Otherwise the old path appears to be the file named by to.
// If to ends in a trailing slash, the overlay code below treats
// it as a directory replacement, akin to a bind mount.
// However, our processing of external overlay maps removes
// such paths by calling abs, except for / or C:\.

type info struct {
	abs      string
	deleted  bool
	replaced bool
	dir      bool
	file     bool
	actual   string
} // info is a summary of the known information about a path
// being looked up in the virtual file system.
// must be file

type fakeFile struct {
	name string
	real fs.FileInfo
} // fakeFile provides an fs.FileInfo implementation for an overlaid file,
// so that the file has the name of the overlaid file, but takes all
// other characteristics of the replacement file.

type Generator struct {
	r        io.Reader
	path     string
	dir      string
	file     string
	pkg      string
	commands map[ // A Generator represents the state of a single Go source file
	// being scanned for generator commands.
	// base name of file.
	string][]string
	lineNum int
	env     []string// current line number.

}

type TooNewError struct {
	What      string
	GoVersion string
	Toolchain string
} // A TooNewError explains that a module is too new for this version of Go.
// for callers if they want to use it, but not printed

type Switcher interface {
	Error(err error)
	Switch(ctx context.Context)
} // A Switcher provides the ability to switch to a new toolchain in response to TooNewErrors.
// See [cmd/go/internal/toolchain.Switcher] for documentation.

type commentWriter struct {
	W            io.Writer
	wroteSlashes bool
} // commentWriter writes a Go comment to the underlying io.Writer,
// using line comment form (//).
// Wrote "//" at the beginning of the current line.

type errWriter struct {
	w   io.Writer
	err error
} // An errWriter wraps a writer, recording whether a write error occurred.

type TrackingWriter struct {
	w    *bufio.Writer
	last byte
} // TrackingWriter tracks the last byte written on every write so
// we can avoid printing a newline if one was already written or
// if there is no output at all.

type PerPackageFlag struct {
	raw     string
	present bool
	values  []ppfValue// A PerPackageFlag is a command-line flag implementation (a flag.Value)
	// that allows specifying different effective flags for different packages.
	// See 'go help build' for more details about per-package flags.

}

type ppfValue struct {
	match func(*Package) bool
	flags []string// A ppfValue is a single <pattern>=<flags> per-package flag value.
	// compiled pattern

}

type PackagePublic struct {
	Dir           string                `json:",omitempty"`
	ImportPath    string                `json:",omitempty"`
	ImportComment string                `json:",omitempty"`
	Name          string                `json:",omitempty"`
	Doc           string                `json:",omitempty"`
	Target        string                `json:",omitempty"`
	Shlib         string                `json:",omitempty"`
	Root          string                `json:",omitempty"`
	ConflictDir   string                `json:",omitempty"`
	ForTest       string                `json:",omitempty"`
	Export        string                `json:",omitempty"`
	BuildID       string                `json:",omitempty"`
	Module        *modinfo.ModulePublic `json:",omitempty"`
	Match         []string              `json:",omitempty"`// Note: These fields are part of the go command's public API.
	// See list.go. It is okay to add fields, but not to change or
	// remove existing ones. Keep in sync with ../list/list.go
	// info about package's module, if any

	Goroot         bool     `json:",omitempty"`
	Standard       bool     `json:",omitempty"`
	DepOnly        bool     `json:",omitempty"`
	BinaryOnly     bool     `json:",omitempty"`
	Incomplete     bool     `json:",omitempty"`
	DefaultGODEBUG string   `json:",omitempty"`
	Stale          bool     `json:",omitempty"`
	StaleReason    string   `json:",omitempty"`
	GoFiles        []string `json:",omitempty"`// command-line patterns matching this package
	// Source files
	// If you add to this list you MUST add to p.AllFiles (below) too.
	// Otherwise file name security lists will not apply to any new additions.

	CgoFiles []string `json:",omitempty"`// .go source files (excluding CgoFiles, TestGoFiles, XTestGoFiles)

	CompiledGoFiles []string `json:",omitempty"`// .go source files that import "C"

	IgnoredGoFiles []string `json:",omitempty"`// .go output from running cgo on CgoFiles

	InvalidGoFiles []string `json:",omitempty"`// .go source files ignored due to build constraints

	IgnoredOtherFiles []string `json:",omitempty"`// .go source files with detected problems (parse error, wrong package name, and so on)

	CFiles []string `json:",omitempty"`// non-.go source files ignored due to build constraints

	CXXFiles []string `json:",omitempty"`// .c source files

	MFiles []string `json:",omitempty"`// .cc, .cpp and .cxx source files

	HFiles []string `json:",omitempty"`// .m source files

	FFiles []string `json:",omitempty"`// .h, .hh, .hpp and .hxx source files

	SFiles []string `json:",omitempty"`// .f, .F, .for and .f90 Fortran source files

	SwigFiles []string `json:",omitempty"`// .s source files

	SwigCXXFiles []string `json:",omitempty"`// .swig files

	SysoFiles []string `json:",omitempty"`// .swigcxx files

	EmbedPatterns []string `json:",omitempty"`// .syso system object files added to package
	// Embedded files

	EmbedFiles []string `json:",omitempty"`// //go:embed patterns

	CgoCFLAGS []string `json:",omitempty"`// files matched by EmbedPatterns
	// Cgo directives

	CgoCPPFLAGS []string `json:",omitempty"`// cgo: flags for C compiler

	CgoCXXFLAGS []string `json:",omitempty"`// cgo: flags for C preprocessor

	CgoFFLAGS []string `json:",omitempty"`// cgo: flags for C++ compiler

	CgoLDFLAGS []string `json:",omitempty"`// cgo: flags for Fortran compiler

	CgoPkgConfig []string `json:",omitempty"`// cgo: flags for linker

	Imports []string `json:",omitempty"`// cgo: pkg-config names
	// Dependency information

	ImportMap     map[ // import paths used by this package
	string]string      `json:",omitempty"`
	Deps []string `json:",omitempty"`// map from source import to ImportPath (identity entries omitted)

	Error      *PackageError `json:",omitempty"`
	DepsErrors []*// all (recursively) imported dependencies
	// error loading this package (not dependencies)
	PackageError `json:",omitempty"`
	TestGoFiles []string `json:",omitempty"`// errors loading dependencies, collected by go list before output
	// Test information
	// If you add to this list you MUST add to p.AllFiles (below) too.
	// Otherwise file name security lists will not apply to any new additions.

	TestImports []string `json:",omitempty"`// _test.go files in package

	TestEmbedPatterns []string `json:",omitempty"`// imports from TestGoFiles

	TestEmbedFiles []string `json:",omitempty"`// //go:embed patterns

	XTestGoFiles []string `json:",omitempty"`// files matched by TestEmbedPatterns

	XTestImports []string `json:",omitempty"`// _test.go files outside package

	XTestEmbedPatterns []string `json:",omitempty"`// imports from XTestGoFiles

	XTestEmbedFiles []string `json:",omitempty"`// //go:embed patterns

} // files matched by XTestEmbedPatterns

type PackageInternal struct {
	Build   *build.Package
	Imports []*// Unexported fields are not part of the public API.
	Package
	CompiledImports []string// this package's direct imports

	RawImports []string// additional Imports necessary when using CompiledGoFiles (all from standard library); 1:1 with the end of PackagePublic.Imports

	ForceLibrary      bool
	CmdlineFiles      bool
	CmdlinePkg        bool
	CmdlinePkgLiteral bool
	Local             bool
	LocalPrefix       string
	ExeName           string
	FuzzInstrument    bool
	Cover             CoverSetup
	OmitDebug         bool
	GobinSubdir       bool
	BuildInfo         *debug.BuildInfo
	TestmainGo        *[]byte// this package's original imports as they appear in the text of the program; 1:1 with the end of PackagePublic.Imports
	// add this info to package main

	Embed map[ // content for _testmain.go
	string][]string
	OrigImportPath string
	PGOProfile     string
	ForMain        string
	Asmflags       []string// //go:embed comment mapping
	// the main package if this package is built specifically for it

	Gcflags []string// -asmflags for this package

	Ldflags []string// -gcflags for this package

	Gccgoflags []string// -ldflags for this package

} // -gccgoflags for this package

type NoGoError struct{ Package *Package } // A NoGoError indicates that no Go files for the package were applicable to the
// build for that package.
//
// That may be because there were no files whatsoever, or because all files were
// excluded, or because all non-excluded files were test sources.

type CoverSetup struct {
	Mode    string
	Cfg     string
	GenMeta bool
} // CoverSetup holds parameters related to coverage setup for a given package (covermode, etc).
// ask cover tool to emit a static meta data if set

type PackageError struct {
	ImportStack      ImportStack
	Pos              string
	Err              error
	IsImportCycle    bool
	alwaysPrintStack bool
} // A PackageError describes an error loading information about a package.
// whether to always print the ImportStack

type ImportPathError interface {
	error
	ImportPath() string
} // ImportPathError is a type of error that prevents a package from being loaded
// for a given import path. When such a package is loaded, a *Package is
// returned with Err wrapping an ImportPathError: the error is attached to
// the imported package, not the importing package.
//
// The string returned by ImportPath must appear in the string returned by
// Error. Errors that wrap ImportPathError (such as PackageError) may omit
// the import path.

type importError struct {
	importPath string
	err        error
} // created with fmt.Errorf

type ImportInfo struct {
	Pkg string
	Pos *token.Position
}

type importSpec struct {
	path                              string
	parentPath, parentDir, parentRoot string
	parentIsStd                       bool
	mode                              int
} // importSpec describes an import declaration in source code. It is used as a
// cache key for resolvedImportCache.

type resolvedImport struct {
	path, dir string
	err       error
} // resolvedImport holds a canonical identifier for a package. It may also contain
// a path to the package's directory and an error if one occurred. resolvedImport
// is the value type in resolvedImportCache.

type preload struct {
	cancel chan struct{}
	sema   chan struct{}
} // preload holds state for managing concurrent preloading of package data.
//
// A preload should be created with newPreload before loading a large
// package graph. flush must be called when package loading is complete
// to ensure preload goroutines are no longer active. This is necessary
// because of global mutable state that cannot safely be read and written
// concurrently. In particular, packageDataCache may be cleared by "go get"
// in GOPATH mode, and modload.loaded (accessed via modload.Lookup) may be
// modified by modload.LoadPackages.

type EmbedError struct {
	Pattern string
	Err     error
} // An EmbedError indicates a problem with a go:embed directive.

type PackageOpts struct {
	IgnoreImports      bool
	ModResolveTests    bool
	MainOnly           bool
	AutoVCS            bool
	SuppressBuildInfo  bool
	SuppressEmbedFiles bool
} // PackageOpts control the behavior of PackagesAndErrors and other package
// loading functions.
// SuppressEmbedFiles is true if the caller does not need any embed files to be populated on the
// package.

type mainPackageError struct{ importPath string }

type Printer interface {
	Printf(pkg *Package, format string, args ...any)
	Errorf(pkg *Package, format string, args ...any)
} // A Printer reports output about a Package.
// Errorf prints output in the form of `log.Errorf` and reports that
// building pkg failed.
//
// This ensures the output is terminated with a new line if there's any
// output, but does not do any other formatting. Callers should generally
// use a higher-level output abstraction, such as (*Shell).reportCmd.
//
// pkg may be nil if this output is not associated with the build of a
// particular package.
//
// This sets the process exit status to 1.

type TextPrinter struct{ Writer io.Writer } // A TextPrinter emits text format output to Writer.

type JSONPrinter struct{ enc *json.Encoder } // A JSONPrinter emits output about a build in JSON format.

type jsonBuildEvent struct {
	ImportPath string
	Action     string
	Output     string `json:",omitempty"`
} // Non-empty if Action == “build-output”

type TestCover struct {
	Mode  string
	Local bool
	Pkgs  []*Package
	Paths []string
}

type testFuncs struct {
	Tests       []testFunc
	Benchmarks  []testFunc
	FuzzTargets []testFunc
	Examples    []testFunc
	TestMain    *testFunc
	Package     *Package
	ImportTest  bool
	NeedTest    bool
	ImportXtest bool
	NeedXtest   bool
	Cover       *TestCover
}

type testFunc struct {
	Package   string
	Name      string
	Output    string
	Unordered bool
} // imported package name (_test or _xtest)
// output is allowed to be unordered.

type inodeLock struct {
	owner File
	queue []<-chan File
}

type osFile struct{ *os.File } // osFile embeds a *os.File while keeping the pointer itself unexported.
// (When we close a File, it must be the same file descriptor that we opened!)

type Mutex struct {
	Path string
	mu   sync.Mutex
} // A Mutex provides mutual exclusion within and across processes by locking a
// well-known file. Such a file generally guards some other part of the
// filesystem: for example, a Mutex file in a directory might guard access to
// the entire tree rooted in that directory.
//
// Mutex does not implement sync.Locker: unlike a sync.Mutex, a lockedfile.Mutex
// can fail to lock (e.g. if there is a permission error in the filesystem).
//
// Like a sync.Mutex, a Mutex may be included as a field of a larger struct but
// must not be copied after first use. The Path field must be set before first
// use and must not be change thereafter.
// A redundant mutex. The race detector doesn't know about file locking, so in tests we may need to lock something that it understands.

type Data struct {
	f    *os.File
	Data []byte// Data is mmap'ed read-only data from a file.
	// The backing file is never closed, so Data
	// remains valid for the lifetime of the process.

}

type ModuleJSON struct {
	Path     string           `json:",omitempty"`
	Version  string           `json:",omitempty"`
	Query    string           `json:",omitempty"`
	Error    string           `json:",omitempty"`
	Info     string           `json:",omitempty"`
	GoMod    string           `json:",omitempty"`
	Zip      string           `json:",omitempty"`
	Dir      string           `json:",omitempty"`
	Sum      string           `json:",omitempty"`
	GoModSum string           `json:",omitempty"`
	Origin   *codehost.Origin `json:",omitempty"`
	Reuse    bool             `json:",omitempty"`
} // A ModuleJSON describes the result of go mod download.

type fileJSON struct {
	Module    editModuleJSON
	Go        string `json:",omitempty"`
	Toolchain string `json:",omitempty"`
	Require   []requireJSON// fileJSON is the -json output data structure.

	Exclude []module.Version
	Replace []replaceJSON
	Retract []retractJSON
	Tool    []toolJSON
	Ignore  []ignoreJSON
}

type editModuleJSON struct {
	Path       string
	Deprecated string `json:",omitempty"`
}

type requireJSON struct {
	Path     string
	Version  string `json:",omitempty"`
	Indirect bool   `json:",omitempty"`
}

type replaceJSON struct {
	Old module.Version
	New module.Version
}

type retractJSON struct {
	Low       string `json:",omitempty"`
	High      string `json:",omitempty"`
	Rationale string `json:",omitempty"`
}

type toolJSON struct{ Path string }

type ignoreJSON struct{ Path string }

type goVersionFlag struct{ v string } // A goVersionFlag is a flag.Value representing a supported Go version.
//
// (Note that the -go argument to 'go mod edit' is *not* a goVersionFlag.
// It intentionally allows newer-than-supported versions as arguments.)

type metakey struct {
	modPath string
	dst     string
}

type DownloadDirPartialError struct {
	Dir string
	Err error
} // DownloadDirPartialError is returned by DownloadDir if a module directory
// exists but was not completely populated.
//
// DownloadDirPartialError is equivalent to fs.ErrNotExist.

type cachingRepo struct {
	path          string
	versionsCache par.ErrCache[string, *Versions]
	statCache     par.ErrCache[string, *RevInfo]
	latestCache   par.ErrCache[struct{}, *RevInfo]
	gomodCache    par.ErrCache[string, []byte]// A cachingRepo is a cache around an underlying Repo,
	// avoiding redundant calls to ModulePath, Versions, Stat, Latest, and GoMod (but not CheckReuse or Zip).
	// It is also safe for simultaneous use by multiple goroutines
	// (so that it can be returned from Lookup multiple times).
	// It serializes calls to the underlying Repo.

	once     sync.Once
	initRepo func(context.Context) (Repo, error)
	r        Repo
}

type cachedInfo struct {
	info *RevInfo
	err  error
}

type Repo interface {
	CheckReuse(ctx context.Context, old *Origin, subdir string) error
	Tags(ctx context.Context, prefix string) (*Tags, error)
	Stat(ctx context.Context, rev string) (*RevInfo, error)
	Latest(ctx context.Context) (*RevInfo, error)
	ReadFile(ctx context.Context, rev, file string, maxSize int64) (data []byte,// A Repo represents a code hosting source.
	// Typical implementations include local version control repositories,
	// remote version control servers, and code hosting sites.
	//
	// A Repo must be safe for simultaneous use by multiple goroutines,
	// and callers must not modify returned values, which may be cached and shared.
	// ReadFile reads the given file in the file tree corresponding to revision rev.
	// It should refuse to read more than maxSize bytes.
	//
	// If the requested file does not exist it should return an error for which
	// os.IsNotExist(err) returns true.
	err error)
	ReadZip(ctx context.Context, rev, subdir string, maxSize int64) (zip io.ReadCloser, err error)
	RecentTag(ctx context.Context, rev, prefix string, allowed func(tag string) bool) (tag string, err error)
	DescendsFrom(ctx context.Context, rev, tag string) (bool, error)
} // ReadZip downloads a zip file for the subdir subdirectory
// of the given revision to a new file in a given temporary directory.
// It should refuse to read more than maxSize bytes.
// It returns a ReadCloser for a streamed copy of the zip file.
// All files in the zip file are expected to be
// nested in a single top-level directory, whose name is not specified.
// DescendsFrom reports whether rev or any of its ancestors has the given tag.
//
// DescendsFrom must return true for any tag returned by RecentTag for the
// same revision.

type Origin struct {
	VCS       string `json:",omitempty"`
	URL       string `json:",omitempty"`
	Subdir    string `json:",omitempty"`
	Hash      string `json:",omitempty"`
	TagPrefix string `json:",omitempty"`
	TagSum    string `json:",omitempty"`
	Ref       string `json:",omitempty"`
	RepoSum   string `json:",omitempty"`
} // An Origin describes the provenance of a given repo method result.
// It can be passed to CheckReuse (usually in a different go command invocation)
// to see whether the result remains up-to-date.
// If RepoSum is non-empty, then the resolution of this module version
// failed due to the repo being available but the version not being present.
// This depends on the entire state of the repo, which RepoSum summarizes.
// For Git, this is a hash of all the refs and their hashes.

type Tags struct {
	Origin *Origin
	List   []Tag// A Tags describes the available tags in a code repository.

}

type Tag struct {
	Name string
	Hash string
} // A Tag describes a single tag in a code repository.
// content hash identifying tag's content, if available

type RevInfo struct {
	Origin  *Origin
	Name    string
	Short   string
	Version string
	Time    time.Time
	Tags    []string// A RevInfo describes a single revision in a source code repository.
	// commit time

} // known tags for commit

type UnknownRevisionError struct{ Rev string } // UnknownRevisionError is an error equivalent to fs.ErrNotExist, but for a
// revision rather than a file.

type noCommitsError struct{}

type RunError struct {
	Cmd      string
	Err      error
	Stderr   []byte
	HelpText string
}

type RunArgs struct {
	cmdline []any
	dir     string
	local   bool
	env     []string// the command to run
	// true if the VCS information is local

	stdin io.Reader
} // environment variables for the command

type notExistError struct{ err error } // A notExistError wraps another error to retain its original text
// but makes it opaquely equivalent to fs.ErrNotExist.

type gitRepo struct {
	ctx               context.Context
	remote, remoteURL string
	local             bool
	dir               string
	sha256Hashes      bool
	mu                lockedfile.Mutex
	fetchLevel        int
	statCache         par.ErrCache[string, *RevInfo]
	refsOnce          sync.Once
	refs              map[ // local only lookups; no remote fetches
	// refs maps branch and tag refs (e.g., "HEAD", "refs/heads/master")
	// to commits (e.g., "37ffd2e798afde829a34e8955b716ab730b2a6d6")
	string]string
	refsErr       error
	localTagsOnce sync.Once
	localTags     sync.Map
} // map[string]bool

type VCSError struct{ Err error } // A VCSError indicates an error using a version control system.
// The implication of a VCSError is that we know definitively where
// to get the code, but we can't access it due to the error.
// The caller should report this error instead of continuing to probe
// other possible module paths.
//
// TODO(golang.org/issue/31730): See if we can invert this. (Return a
// distinguished error for “repo not found” and treat everything else
// as terminal.)

type vcsCacheKey struct {
	vcs    string
	remote string
	local  bool
}

type vcsRepo struct {
	mu       lockedfile.Mutex
	remote   string
	cmd      *vcsCmd
	dir      string
	local    bool
	tagsOnce sync.Once
	tags     map[ // protects all commands, so we don't have to decide which are safe on a per-VCS basis
	string]bool
	branchesOnce sync.Once
	branches     map[string]bool
	fetchOnce    sync.Once
	fetchErr     error
}

type vcsCmd struct {
	vcs  string
	init func(remote string) []string// vcs name "hg"

	tags func(remote string) []string// cmd to init repo to track remote

	tagRE    *lazyregexp.Regexp
	branches func(remote string) []string// cmd to list local tags
	// regexp to extract tag names from output of tags cmd

	branchRE      *lazyregexp.Regexp
	badLocalRevRE *lazyregexp.Regexp
	statLocal     func(rev, remote string) []string// cmd to list local branches
	// regexp of names that must not be served out of local cache without doing fetch first

	parseStat func(rev, out string) (*RevInfo, error)
	fetch     []string// cmd to stat local rev
	// cmd to parse output of statLocal

	latest   string
	readFile func(rev, file, remote string) []string// cmd to fetch everything from remote
	// name of latest commit on remote (tip, HEAD, etc)

	readZip func(rev, subdir, remote, target string) []string// cmd to read rev's file

	doReadZip func(ctx context.Context, dst io.Writer, workDir, rev, subdir, remote string) error
} // cmd to read rev's subdir as zip file
// arbitrary function to read rev's subdir as zip file

type deleteCloser struct{ *os.File } // deleteCloser is a file that gets deleted on Close.

type limitedWriter struct {
	W               io.Writer
	N               int64
	ErrLimitReached error
}

type codeRepo struct {
	modPath     string
	code        codehost.Repo
	codeRoot    string
	codeDir     string
	pathMajor   string
	pathPrefix  string
	pseudoMajor string
} // A codeRepo implements modfetch.Repo using an underlying codehost.Repo.
// pseudoMajor is the major version prefix to require when generating
// pseudo-versions for this module, derived from the module path. pseudoMajor
// is empty if the module path does not include a version suffix (that is,
// accepts either v0 or v1).

type zipFile struct {
	name string
	f    *zip.File
}

type dataFile struct {
	name string
	data []byte
}

type dataFileInfo struct{ f dataFile }

type modSum struct {
	mod module.Version
	sum string
}

type sumState struct {
	m map[module.Version][]string
	w map[ // content of go.sum file
	string]map[module.Version][]string
	status map[ // sum file in workspace -> content of that sum file
	modSum]modSumStatus
	overwrite bool
	enabled   bool
} // state of sums in m
// whether to use go.sum at all

type modSumStatus struct{ used, dirty bool }

type proxySpec struct {
	url             string
	fallBackOnError bool
} // url is the proxy URL or one of "off", "direct", "noproxy".
// fallBackOnError is true if a request should be attempted on the next proxy
// in the list after any error from this proxy. If fallBackOnError is false,
// the request will only be attempted on the next proxy if the error is
// equivalent to os.ErrNotFound, which is true for 404 and 410 responses.

type proxyRepo struct {
	url            *url.URL
	path           string
	redactedBase   string
	listLatestOnce sync.Once
	listLatest     *RevInfo
	listLatestErr  error
} // The combined module proxy URL joined with the module path.
// The base module proxy URL in [url.URL.Redacted] form.

type Versions struct {
	Origin *codehost.Origin `json:",omitempty"`
	List   []string// A Versions describes the available versions in a module repository.
	// origin information for reuse

} // semver versions

type lookupCacheKey struct{ proxy, path string }

type lookupDisabledError struct{}

type loggingRepo struct{ r Repo } // A loggingRepo is a wrapper around an underlying Repo
// that prints a log message at the start and end of each call.
// It can be inserted when debugging.

type errRepo struct {
	modulePath string
	err        error
} // errRepo is a Repo that returns the same error for all operations.
//
// It is useful in conjunction with caching, since cache hits will not attempt
// the prohibited operations.

type dbClient struct {
	key     string
	name    string
	direct  *url.URL
	once    sync.Once
	base    *url.URL
	baseErr error
}

type toolchainRepo struct {
	path string
	repo Repo
} // A toolchainRepo is a synthesized repository reporting Go toolchain versions.
// It has path "go" or "toolchain". The "go" repo reports versions like "1.2".
// The "toolchain" repo reports versions like "go1.2".
//
// Note that the repo ONLY reports versions. It does not actually support
// downloading of the actual toolchains. Instead, that is done using
// the regular repo code with "golang.org/toolchain".
// The naming conflict is unfortunate: "golang.org/toolchain"
// should perhaps have been "go.dev/dl", but it's too late.
//
// For clarity, this file refers to golang.org/toolchain as the "DL" repo,
// the one you can actually download.
// underlying DL repo

type upgradeFlag struct {
	rawVersion string
	version    string
} // upgradeFlag is a custom flag.Value for -u.

type dFlag struct {
	value bool
	set   bool
} // dFlag is a custom flag.Value for the deprecated -d flag
// which will be used to provide warnings or errors if -d
// is provided.

type resolver struct {
	localQueries []*query
	pathQueries  []*// queries for absolute or relative paths
	query
	wildcardQueries []*// package path literal queries in original order
	query
	patternAllQueries []*// path wildcard queries in original order
	query
	workQueries []*// queries with the pattern "all"
	query
	toolQueries []*// queries with the pattern "work"
	query
	nonesByPath map[ // queries with the pattern "tool"
	// Indexed "none" queries. These are also included in the slices above;
	// they are indexed here to speed up noneForPath.
	string]*query
	wildcardNones []*// path-literal "@none" queries indexed by path
	query
	resolvedVersion map[ // wildcard "@none" queries
	// resolvedVersion maps each module path to the version of that module that
	// must be selected in the final build list, along with the first query
	// that resolved the module to that version (the “reason”).
	string]versionReason
	buildList        []module.Version
	buildListVersion map[string]string
	initialVersion   map[ // index of buildList (module path → version)
	string]string
	missing []pathSet// index of the initial build list at the start of 'go get'

	work               *par.Queue
	matchInModuleCache par.ErrCache[matchInModuleKey, []string]// candidates for missing transitive dependencies

	workspace *workspace
} // workspace is used to check whether, in workspace mode, any of the workspace
// modules would contain a package.

type versionReason struct {
	version string
	reason  *query
}

type matchInModuleKey struct {
	pattern string
	m       module.Version
}

type workspace struct {
	modules map[ // workspace represents the set of modules in a workspace.
	// It can be used
	string]string
} // path -> modroot

type query struct {
	raw                      string
	rawVersion               string
	pattern                  string
	patternIsLocal           bool
	version                  string
	matchWildcard            func(path string) bool
	canMatchWildcardInModule func(mPath string) bool
	conflict                 *query
	candidates               []pathSet// A query describes a command-line argument and the modules and/or packages
	// to which that argument may resolve..
	// candidates is a list of sets of alternatives for a path that matches (or
	// contains packages that match) the pattern. The query can be resolved by
	// choosing exactly one alternative from each set in the list.
	//
	// A path-literal query results in only one set: the path itself, which
	// may resolve to either a package path or a module path.
	//
	// A wildcard query results in one set for each matching module path, each
	// module for which the matching version contains at least one matching
	// package, and (if no other modules match) one candidate set for the pattern
	// overall if no existing match is identified in the build list.
	//
	// A query for pattern "all" results in one set for each package transitively
	// imported by the main module.
	//
	// The special query for the "-u" flag results in one set for each
	// otherwise-unconstrained package that has available upgrades.

	candidatesMu sync.Mutex
	pathSeen     sync.Map
	resolved     []module.// pathSeen ensures that only one pathSet is added to the query per
	// unique path.
	// resolved contains the set of modules whose versions have been determined by
	// this query, in the order in which they were determined.
	//
	// The resolver examines the candidate sets for each query, resolving one
	// module per candidate set in a way that attempts to avoid obvious conflicts
	// between the versions resolved by different queries.
	Version
	matchesPackages bool
} // matchesPackages is true if the resolved modules provide at least one
// package matching q.pattern.

type pathSet struct {
	path    string
	pkgMods []module.// A pathSet describes the possible options for resolving a specific path
	// to a package and/or module.
	// pkgMods is a set of zero or more modules, each of which contains the
	// package with the indicated path. Due to the requirement that imports be
	// unambiguous, only one such module can be in the build list, and all others
	// must be excluded.
	Version
	mod module.Version
	err error
} // mod is either the zero Version, or a module that does not contain any
// packages matching the query but for which the module path itself
// matches the query pattern.
//
// We track this module separately from pkgMods because, all else equal, we
// prefer to match a query to a package rather than just a module. Also,
// unlike the modules in pkgMods, this module does not inherently exclude
// any other module in pkgMods.

type conflictError struct {
	mPath    string
	proposed versionReason
	conflict versionReason
}

type MultiplePackageError struct {
	Dir      string
	Packages []string// MultiplePackageError describes a directory containing
	// multiple buildable Go source files for multiple packages.
	// directory containing files

	Files []string// package names found

} // corresponding files: Files[i] declares package Packages[i]

type fileImport struct {
	path string
	pos  token.Pos
	doc  *ast.CommentGroup
}

type fileEmbed struct {
	pattern string
	pos     token.Position
}

type Module struct {
	modroot string
	d       *decoder
	n       int
} // Module represents and encoded module index file. It is used to
// do the equivalent of build.Import of packages in the module and answer other
// questions based on the index file's data.
// number of packages

type IndexPackage struct {
	error       error
	dir         string
	modroot     string
	sourceFiles []*// IndexPackage holds the information in the index
	// needed to load a package in a specific directory.
	// Source files
	sourceFile
}

type sourceFile struct {
	d               *decoder
	pos             int
	onceReadImports sync.Once
	savedImports    []rawImport// sourceFile represents the information of a given source file in the module index.
	// start of sourceFile encoding in d

} // saved imports so that they're only read once

type decoder struct {
	data []byte// A decoder helps decode the index format.

	str []byte// data after header

} // string table

type rawPackage struct {
	error       string
	dir         string
	sourceFiles []*// rawPackage holds the information from each package that's needed to
	// fill a build.Package once the context is available.
	// Source files
	rawFile
}

type parseError struct {
	ErrorList   *scanner.ErrorList
	ErrorString string
}

type rawFile struct {
	error                string
	parseError           string
	name                 string
	synopsis             string
	pkgName              string
	ignoreFile           bool
	binaryOnly           bool
	cgoDirectives        string
	goBuildConstraint    string
	plusBuildConstraints []string// rawFile is the struct representation of the file holding all
	// information in its fields.
	// the #cgo directive lines in the comment on import "C"

	imports    []rawImport
	embeds     []embed
	directives []build.Directive
}

type rawImport struct {
	path     string
	position token.Position
}

type embed struct {
	pattern  string
	position token.Position
}

type encoder struct {
	b           []byte
	stringTable []byte
	strings     map[string]int
}

type ModulePublic struct {
	Path     string   `json:",omitempty"`
	Version  string   `json:",omitempty"`
	Query    string   `json:",omitempty"`
	Versions []string `json:",omitempty"`// module path
	// version query corresponding to this version

	Replace   *ModulePublic `json:",omitempty"`
	Time      *time.Time    `json:",omitempty"`
	Update    *ModulePublic `json:",omitempty"`
	Main      bool          `json:",omitempty"`
	Indirect  bool          `json:",omitempty"`
	Dir       string        `json:",omitempty"`
	GoMod     string        `json:",omitempty"`
	GoVersion string        `json:",omitempty"`
	Retracted []string      `json:",omitempty"`// available module versions
	// go version used in module

	Deprecated string           `json:",omitempty"`
	Error      *ModuleError     `json:",omitempty"`
	Sum        string           `json:",omitempty"`
	GoModSum   string           `json:",omitempty"`
	Origin     *codehost.Origin `json:",omitempty"`
	Reuse      bool             `json:",omitempty"`
} // retraction information, if any (with -retracted or -u)
// reuse of old module info is safe

type ModuleError struct {
	Err string
} // error text

type Requirements struct {
	pruning     modPruning
	rootModules []module.// A Requirements represents a logically-immutable set of root module requirements.
	// rootModules is the set of root modules of the graph, sorted and capped to
	// length. It may contain duplicates, and may contain multiple versions for a
	// given module path. The root modules of the graph are the set of main
	// modules in workspace mode, and the main module's direct requirements
	// outside workspace mode.
	//
	// The roots are always expected to contain an entry for the "go" module,
	// indicating the Go language version in use.
	Version
	maxRootVersion map[string]string
	direct         map[ // direct is the set of module paths for which we believe the module provides
	// a package directly imported by a package or test in the main module.
	//
	// The "direct" map controls which modules are annotated with "// indirect"
	// comments in the go.mod file, and may impact which modules are listed as
	// explicit roots (vs. indirect-only dependencies). However, it should not
	// have a semantic effect on the build list overall.
	//
	// The initial direct map is populated from the existing "// indirect"
	// comments (or lack thereof) in the go.mod file. It is updated by the
	// package loader: dependencies may be promoted to direct if new
	// direct imports are observed, and may be demoted to indirect during
	// 'go mod tidy' or 'go mod vendor'.
	//
	// The direct map is keyed by module paths, not module versions. When a
	// module's selected version changes, we assume that it remains direct if the
	// previous version was a direct dependency. That assumption might not hold in
	// rare cases (such as if a dependency splits out a nested module, or merges a
	// nested module back into a parent module).
	string]bool
	graphOnce sync.Once
	graph     atomic.Pointer[cachedGraph]
} // guards writes to (but not reads from) graph

type cachedGraph struct {
	mg  *ModuleGraph
	err error
} // A cachedGraph is a non-nil *ModuleGraph, together with any error discovered
// while loading that graph.
// If err is non-nil, mg may be incomplete (but must still be non-nil).

type ModuleGraph struct {
	g             *mvs.Graph
	loadCache     par.ErrCache[module.Version, *modFileSummary]
	buildListOnce sync.Once
	buildList     []module.// A ModuleGraph represents the complete graph of module dependencies
	// of a main module.
	//
	// If the main module supports module graph pruning, the graph does not include
	// transitive dependencies of non-root (implicit) dependencies.
	Version
}

type ConstraintError struct {
	Conflicts []Conflict// A ConstraintError describes inconsistent constraints in EditBuildList
}

type Conflict struct {
	Path []module.// A Conflict is a path of requirements starting at a root or proposed root in
	// the requirement graph, explaining why that root either causes a module passed
	// in the mustSelect list to EditBuildList to be unattainable, or introduces an
	// unresolvable error in loading the requirement graph.
	// Path is a path of requirements starting at some module version passed in
	// the mustSelect argument and ending at a module whose requirements make that
	// version unacceptable. (Path always has len ≥ 1.)
	Version
	Constraint module.Version
	Err        error
} // If Err is nil, Constraint is a module version passed in the mustSelect
// argument that has the same module path as, and a lower version than,
// the last element of the Path slice.
// If Constraint is unset, Err is an error encountered when loading the
// requirements of the last element in Path.

type perPruning[T any] struct {
	pruned   T
	unpruned T
}

type dqTracker struct {
	extendedRootPruning map[ // A dqTracker tracks and propagates the reason that each module version
	// cannot be included in the module graph.
	// extendedRootPruning is the modPruning given the go.mod file for each root
	// in the extended module graph.
	module.Version]modPruning
	dqReason map[ // dqReason records whether and why each encountered version is
	// disqualified in a pruned or unpruned context.
	module.Version]perPruning[dqState]
	requiring map[ // requiring maps each not-yet-disqualified module version to the versions
	// that would cause that module's requirements to be included in a pruned or
	// unpruned context. If that version becomes disqualified, the
	// disqualification will be propagated to all of the versions in the
	// corresponding list.
	//
	// This map is similar to the module requirement graph, but includes more
	// detail about whether a given dependency edge appears in a pruned or
	// unpruned context. (Other commands do not need this level of detail.)
	module.Version][]module.Version
}

type dqState struct {
	err error
	dep module.Version
} // A dqState indicates whether and why a module version is “disqualified” from
// being used in a way that would incorporate its requirements.
//
// The zero dqState indicates that the module version is not known to be
// disqualified, either because it is ok or because we are currently traversing
// a cycle that includes it.
// disqualified because the module is or requires dep

type ImportMissingError struct {
	Path                string
	Module              module.Version
	QueryErr            error
	ImportingMainModule module.Version
	isStd               bool
	importerGoVersion   string
	replaced            module.Version
	newMissingVersion   string
} // isStd indicates whether we would expect to find the package in the standard
// library. This is normally true for all dotless import paths, but replace
// directives can cause us to treat the replaced paths as also being in
// modules.
// newMissingVersion is set to a newer version of Module if one is present
// in the build list. When set, we can't automatically upgrade.

type AmbiguousImportError struct {
	importPath string
	Dirs       []string// An AmbiguousImportError indicates an import of a package found in multiple
	// modules in the build list, or found in both the main module and its vendor
	// directory.

	Modules []module.Version
} // Either empty or 1:1 with Dirs.

type DirectImportFromImplicitDependencyError struct {
	ImporterPath string
	ImportedPath string
	Module       module.Version
} // A DirectImportFromImplicitDependencyError indicates a package directly
// imported by a package or test in the main module that is satisfied by a
// dependency that is not explicit in the main module's go.mod file.

type ImportMissingSumError struct {
	importPath string
	found      bool
	mods       []module.// ImportMissingSumError is reported in readonly mode when we need to check
	// if a module contains a package, but we don't have a sum for its .zip file.
	// We might need sums for multiple modules to verify the package is unique.
	//
	// TODO(#43653): consolidate multiple errors of this type into a single error
	// that suggests a 'go get' command for root packages that transitively import
	// packages from modules with missing sums. load.CheckPackageErrors would be
	// a good place to consolidate errors, but we'll need to attach the import
	// stack here.
	Version
	importer, importerVersion string
	importerIsTest            bool
} // optional, but used for additional context

type invalidImportError struct {
	importPath string
	err        error
}

type sumMissingError struct{ suggestion string }

type MainModuleSet struct {
	versions []module.// versions are the module.Version values of each of the main modules.
	// For each of them, the Path fields are ordinary module paths and the Version
	// fields are empty strings.
	// versions is clipped (len=cap).
	Version
	modRoot map[ // modRoot maps each module in versions to its absolute filesystem path.
	module.Version]string
	pathPrefix map[ // pathPrefix is the path prefix for packages in the module, without a trailing
	// slash. For most modules, pathPrefix is just version.Path, but the
	// standard-library module "std" has an empty prefix.
	module.Version]string
	inGorootSrc map[ // inGorootSrc caches whether modRoot is within GOROOT/src.
	// The "std" module is special within GOROOT/src, but not otherwise.
	module.Version]bool
	modFiles           map[module.Version]*modfile.File
	tools              map[string]bool
	modContainingCWD   module.Version
	workFile           *modfile.WorkFile
	workFileReplaceMap map[module.Version]module.Version
	highestReplaced    map[ // highest replaced version of each module path; empty string for wildcard-only replacements
	string]string
	indexMu sync.Mutex
	indices map[module.Version]*modFileIndex
}

type noMainModulesError struct{} // noMainModulesError returns the appropriate error if there is no main module or
// main modules depending on whether the go command is in workspace mode.

type goModDirtyError struct{}

type WriteOpts struct {
	DropToolchain     bool
	ExplicitToolchain bool
	AddTools          []string// WriteOpts control the behavior of WriteGoMod.
	// go get has set explicit toolchain version

	DropTools []string// go get -tool example.com/m1

	TidyWroteGo bool
} // go get -tool example.com/m1@none
// Go.Version field already updated by 'go mod tidy'

type loader struct {
	loaderParams
	allClosesOverTests bool
	skipImportModFiles bool
	work               *par.Queue
	roots              []*// A loader manages the process of loading information about
	// the required packages for a particular build,
	// checking that the packages are available in the module set,
	// and updating the module set if needed.
	// reset on each iteration
	loadPkg
	pkgCache *par.Cache[string, *loadPkg]
	pkgs     []*loadPkg
} // transitive closure of loaded packages and tests; populated in buildStacks

type loaderParams struct {
	PackageOpts
	requirements     *Requirements
	allPatternIsRoot bool
	listRoots        func(rs *Requirements) []string// loaderParams configure the packages loaded by, and the properties reported
	// by, a loader instance.
	// Is the "all" pattern an additional root?

}

type loadPkg struct {
	path    string
	testOf  *loadPkg
	flags   atomicLoadPkgFlags
	mod     module.Version
	dir     string
	err     error
	imports []*// A loadPkg records information about a single loaded package.
	// error loading package
	loadPkg
	testImports []string// packages imported by this one

	inStd   bool
	altMods []module.// test-only imports, saved for use by pkg.test.
	Version
	testOnce sync.Once
	test     *loadPkg
	stack    *loadPkg
} // modules that could have contained the package but did not
// package importing this one in minimal import stack for this pkg

type atomicLoadPkgFlags struct{ bits atomic.Int32 } // An atomicLoadPkgFlags stores a loadPkgFlags for which individual flags can be
// added atomically.

type modFileIndex struct {
	data []byte// A modFileIndex is an index of data corresponding to a modFile
	// at a specific point in time.

	dataNeedsFix bool
	module       module.Version
	goVersion    string
	toolchain    string
	require      map[ // true if fixVersion applied a change while parsing data
	// Go version (no "v" or "go" prefix)
	module.Version]requireMeta
	replace map[module.Version]module.Version
	exclude map[module.Version]bool
	ignore  []string
}

type requireMeta struct{ indirect bool }

type excludedError struct{}

type ModuleRetractedError struct{ Rationale []string }

type retractionLoadingError struct {
	m   module.Version
	err error
}

type modFileSummary struct {
	module    module.Version
	goVersion string
	toolchain string
	ignore    []string// A modFileSummary is a summary of a go.mod file for which we do not need to
	// retain complete information — for example, the go.mod file of a dependency
	// module.

	pruning    modPruning
	require    []module.Version
	retract    []retraction
	deprecated string
}

type retraction struct {
	modfile.VersionInterval
	Rationale string
} // A retraction consists of a retracted version interval and rationale.
// retraction is like modfile.Retract, but it doesn't point to the syntax tree.

type mvsReqs struct {
	roots []module.// mvsReqs implements mvs.Reqs for module semantic versions,
	// with any exclusions or replacements applied internally.
	Version
}

type queryDisabledError struct{}

type queryMatcher struct {
	path               string
	prefix             string
	filter             func(version string) bool
	allowed            AllowedFunc
	canStat            bool
	preferLower        bool
	mayUseLatest       bool
	preferIncompatible bool
} // if true, the query can be resolved by repo.Stat
// if true, choose the lowest matching version

type QueryResult struct {
	Mod      module.Version
	Rev      *modfetch.RevInfo
	Packages []string
}

type NoMatchingVersionError struct{ query, current string } // A NoMatchingVersionError indicates that Query found a module at the requested
// path, but not at any versions satisfying the query string and allow-function.
//
// NOTE: NoMatchingVersionError MUST NOT implement Is(fs.ErrNotExist).
//
// If the module came from a proxy, that proxy had to return a successful status
// code for the versions it knows about, and thus did not have the opportunity
// to return a non-400 status code to suppress fallback.

type NoPatchBaseError struct{ path string } // A NoPatchBaseError indicates that Query was called with the query "patch"
// but with a current version of "" or "none".

type WildcardInFirstElementError struct {
	Pattern string
	Query   string
} // A WildcardInFirstElementError indicates that a pattern passed to QueryPattern
// had a wildcard in its first path element, and therefore had no pattern-prefix
// modules to search in.

type PackageNotInModuleError struct {
	MainModules []module.// A PackageNotInModuleError indicates that QueryPattern found a candidate
	// module at the requested version, but that module did not contain any packages
	// matching the requested pattern.
	//
	// NOTE: PackageNotInModuleError MUST NOT implement Is(fs.ErrNotExist).
	//
	// If the module came from a proxy, that proxy had to return a successful status
	// code for the versions it knows about, and thus did not have the opportunity
	// to return a non-400 status code to suppress fallback.
	Version
	Mod         module.Version
	Replacement module.Version
	Query       string
	Pattern     string
}

type versionRepo interface {
	ModulePath() string
	CheckReuse(context.Context, *codehost.Origin) error
	Versions(ctx context.Context, prefix string) (*modfetch.Versions, error)
	Stat(ctx context.Context, rev string) (*modfetch.RevInfo, error)
	Latest(context.Context) (*modfetch.RevInfo, error)
} // A versionRepo is a subset of modfetch.Repo that can report information about
// available versions, but cannot fetch specific source files.

type emptyRepo struct {
	path string
	err  error
} // An emptyRepo is a versionRepo that contains no versions.

type replacementRepo struct{ repo versionRepo } // A replacementRepo augments a versionRepo to include the replacement versions
// (if any) found in the main module's go.mod file.
//
// A replacementRepo suppresses "not found" errors for otherwise-nonexistent
// modules, so a replacementRepo should only be constructed for a module that
// actually has one or more valid replacements.

type QueryMatchesMainModulesError struct {
	MainModules []module.// A QueryMatchesMainModulesError indicates that a query requests
	// a version of the main module that cannot be satisfied.
	// (The main module's version cannot be changed.)
	Version
	Pattern string
	Query   string
}

type QueryUpgradesAllError struct {
	MainModules []module.// A QueryUpgradesAllError indicates that a query requests
	// an upgrade on the all pattern.
	// (The main module's version cannot be changed.)
	Version
	Query string
}

type QueryMatchesPackagesInMainModuleError struct {
	Pattern  string
	Query    string
	Packages []string// A QueryMatchesPackagesInMainModuleError indicates that a query cannot be
	// satisfied because it matches one or more packages found in the main module.

}

type vendorMetadata struct {
	Explicit    bool
	Replacement module.Version
	GoVersion   string
}

type BuildListError struct {
	Err   error
	stack []buildListErrorElem// BuildListError decorates an error that occurred gathering requirements
	// while constructing a build list. BuildListError prints the chain
	// of requirements to the module where the error occurred.

}

type buildListErrorElem struct {
	m          module.Version
	nextReason string
} // nextReason is the reason this module depends on the next module in the
// stack. Typically either "requires", or "updating to".

type Graph struct {
	cmp   func(p, v1, v2 string) int
	roots []module.// Graph implements an incremental version of the MVS algorithm, with the
	// requirements pushed by the caller instead of pulled by the MVS traversal.
	Version
	required map[module.Version][]module.Version
	isRoot   map[module.Version]bool
	selected map[ // contains true for roots and false for reachable non-roots
	string]string
} // path → version

type Reqs interface {
	Required(m module.Version) ([]module.// A Reqs is the requirement graph on which Minimal Version Selection (MVS) operates.
	//
	// The version strings are opaque except for the special version "none"
	// (see the documentation for module.Version). In particular, MVS does not
	// assume that the version strings are semantic versions; instead, the Max method
	// gives access to the comparison operation.
	//
	// It must be safe to call methods on a Reqs from multiple goroutines simultaneously.
	// Because a Reqs may read the underlying graph from the network on demand,
	// the MVS algorithms parallelize the traversal to overlap network delays.
	// Required returns the module versions explicitly required by m itself.
	// The caller must not modify the returned list.
	Version, error)
	Max(p, v1, v2 string) string
} // Max returns the maximum of v1 and v2 (it returns either v1 or v2)
// in the module with path p.
//
// For all versions v, Max(v, "none") must be v,
// and for the target passed as the first argument to MVS functions,
// Max(target, v) must be target.
//
// Note that v1 < v2 can be written Max(v1, v2) != v1
// and similarly v1 <= v2 can be written Max(v1, v2) == v2.

type UpgradeReqs interface {
	Reqs
	Upgrade(m module.Version) (module.Version, error)
} // An UpgradeReqs is a Reqs that can also identify available upgrades.
// Upgrade returns the upgraded version of m,
// for use during an UpgradeAll operation.
// If m should be kept as is, Upgrade returns m.
// If m is not yet used in the build, then m.Version will be "none".
// More typically, m.Version will be the version required
// by some other module in the build.
//
// If no module version is available for the given path,
// Upgrade returns a non-nil error.
// TODO(rsc): Upgrade must be able to return errors,
// but should "no latest version" just return m instead?

type DowngradeReqs interface {
	Reqs
	Previous(m module.Version) (module.Version, error)
} // A DowngradeReqs is a Reqs that can also identify available downgrades.
// Previous returns the version of m.Path immediately prior to m.Version,
// or "none" if no such version is known.

type override struct {
	target module.Version
	list   []module.Version
	Reqs
}

type Match struct {
	pattern string
	Dirs    []string// A Match represents the result of matching a single package pattern.
	// the pattern itself

	Pkgs []string// if the pattern is local, directories that potentially contain matching packages

	Errs []error// matching packages (import paths)

} // errors matching the patterns to packages, NOT errors loading those packages

type MatchError struct {
	Match *Match
	Err   error
} // A MatchError indicates an error that occurred while attempting to match a
// pattern.

type IgnorePatterns struct {
	relativePatterns []string// IgnorePatterns is normalized with normalizePath.

	anyPatterns []string
}

type testVFlag struct {
	on   bool
	json bool
} // -v is set in some form
// -v=test2json is set, to make output better for test2json

type runTestActor struct {
	c                 runCache
	writeCoverMetaAct *work.Action
	prev              <-chan // runTestActor is the actor for running a test.
	// sequencing of json start messages, to preserve test order
	struct{}
	next chan<- struct{}
} // wait to start until prev is closed
// close next once the next test can start.

type runCache struct {
	disableCache bool
	buf          *bytes.Buffer
	id1          cache.ActionID
	id2          cache.ActionID
} // runCache is the cache for running a single test.
// cache should be disabled for this run

type lockedStdout struct{}

type outputdirFlag struct{ abs string } // outputdirFlag implements the -outputdir flag.
// It interprets an empty value as the working directory of the 'go' command.

type vetFlag struct {
	explicit bool
	off      bool
	flags    []string// vetFlag implements the special parsing logic for the -vet flag:
	// a comma-separated list, with distinguished values "all" and
	// "off", plus a boolean tracking whether it was set explicitly.
	//
	// "all" is encoded as vetFlag{true, false, nil}, since it will
	// pass no flags to the vet binary, and by default, it runs all
	// analyzers.

} // passed to vet when invoked automatically during 'go test'

type shuffleFlag struct {
	on   bool
	seed *int64
}

type Span struct {
	t     *tracer
	name  string
	tid   uint64
	start time.Time
	end   time.Time
}

type tracer struct {
	file       chan traceFile
	nextTID    atomic.Uint64
	nextFlowID atomic.Uint64
} // 1-buffered

type traceKey struct{} // traceKey is the context key for tracing information. It is unexported to prevent collisions with context keys defined in
// other packages.

type traceContext struct {
	t   *tracer
	tid uint64
}

type traceFile struct {
	f       *os.File
	sb      *strings.Builder
	enc     *json.Encoder
	entries int64
}

type Status struct {
	Revision    string
	CommitTime  time.Time
	Uncommitted bool
} // Status is the current state of a local repository.
// Required.

type tagCmd struct {
	cmd     string
	pattern string
} // A tagCmd describes a command to list available tags
// that can be passed to tagSyncCmd.
// regexp to extract tags from list

type vcsPath struct {
	pathPrefix string
	regexp     *lazyregexp.Regexp
	repo       string
	vcs        string
	check      func(match map[ // A vcsPath describes how to convert an import path into a
	// version control system and repository name.
	// version control system to use (expand with match of re)
	string]string) error
	schemelessRepo bool
} // additional checks
// if true, the repo pattern lacks a scheme

type rootName struct {
	filename string
	isDir    bool
}

type vcsNotFoundError struct{ dir string }

type govcsRule struct {
	pattern string
	allowed []string// A govcsRule is a single GOVCS rule like private:hg|svn.

}

type RepoRoot struct {
	Repo     string
	Root     string
	SubDir   string
	IsCustom bool
	VCS      *Cmd
} // RepoRoot describes the repository root for a tree of source code.
// defined by served <meta> tags (as opposed to hard-coded pattern)

type fetchResult struct {
	url     *urlpkg.URL
	imports []metaImport
	err     error
}

type metaImport struct{ Prefix, VCS, RepoRoot, SubDir string } // metaImport represents the parsed <meta name="go-import"
// content="prefix vcs reporoot subdir" /> tags from HTML files.

type ImportMismatchError struct {
	importPath string
	mismatches []string// An ImportMismatchError is returned where metaImport/s are present
	// but none match our import path.

} // the meta imports that were discarded for not matching our importPath

type authHandler struct{} // authHandler serves requests only if the Basic Auth data sent with the request
// matches the contents of a ".access" file in the requested directory.
//
// For each request, the handler looks for a file named ".access" and parses it
// as a JSON-serialized accessToken. If the credentials from the request match
// the accessToken, the file is served normally; otherwise, it is rejected with
// the StatusCode and Message provided by the token.

type accessToken struct {
	Username, Password string
	StatusCode         int
	Message            string
} // defaults to 401.

type bzrHandler struct{}

type dirHandler struct{} // dirHandler is a vcsHandler that serves the raw contents of a directory.

type fossilHandler struct {
	once          sync.Once
	fossilPath    string
	fossilPathErr error
}

type gitHandler struct {
	once       sync.Once
	gitPath    string
	gitPathErr error
}

type hgHandler struct {
	once      sync.Once
	hgPath    string
	hgPathErr error
}

type insecureHandler struct{} // insecureHandler redirects requests to the same host and path but using the
// "http" scheme instead of "https".

type scriptCtx struct {
	context.Context
	server      *Server
	commitTime  time.Time
	handlerName string
	handler     http.Handler
} // A scriptCtx is a context.Context that stores additional state for script
// commands.

type scriptCtxKey struct{} // scriptCtxKey is the key associating the *scriptCtx in a script's Context..

type SkipError struct{ Msg string }

type svnHandler struct {
	svnRoot      string
	logger       *log.Logger
	pathOnce     sync.Once
	svnservePath string
	svnserveErr  error
	listenOnce   sync.Once
	s            chan *svnState
} // An svnHandler serves requests for Subversion repos.
//
// Unlike the other vcweb handlers, svnHandler does not serve the Subversion
// protocol directly over the HTTP connection. Instead, it opens a separate port
// that serves the (non-HTTP) 'svn' protocol. The test binary can retrieve the
// URL for that port by sending an HTTP request with the query parameter
// "vcwebsvn=1".
//
// We take this approach because the 'svn' protocol is implemented by a
// lightweight 'svnserve' binary that is usually packaged along with the 'svn'
// client binary, whereas only known implementation of the Subversion HTTP
// protocol is the mod_dav_svn apache2 module. Apache2 has a lot of dependencies
// and also seems to rely on global configuration via well-known file paths, so
// implementing a hermetic test using apache2 would require the test to run in a
// complicated container environment, which wouldn't be nearly as
// straightforward for Go contributors to set up and test against on their local
// machine.
// 1-buffered

type svnState struct {
	listener  net.Listener
	listenErr error
	conns     map[ // An svnState describes the state of a port serving the 'svn://' protocol.
	net.Conn]struct{}
	closing bool
	done    chan struct{}
}

type Server struct {
	vcweb   *vcweb.Server
	workDir string
	HTTP    *httptest.Server
	HTTPS   *httptest.Server
}

type vcsHandler interface {
	Available() bool
	Handler(dir string, env []string,// A vcsHandler serves repositories over HTTP for a known version-control tool.
	logger *log.Logger) (http.Handler, error)
}

type scriptResult struct {
	mu       sync.RWMutex
	hash     [sha256.Size]byte
	hashTime time.Time
	handler  http.Handler
	err      error
} // A scriptResult describes the cached result of executing a vcweb script.
// error from executing the script, if any

type ScriptNotFoundError struct{ err error } // A ScriptNotFoundError indicates that the requested script file does not exist.
// (It typically wraps a "stat" error for the script file.)

type ServerNotInstalledError struct{ name string } // A ServerNotInstalledError indicates that the server binary required for the
// indicated VCS does not exist.

type HTTPError struct {
	URL        string
	Status     string
	StatusCode int
	Err        error
	Detail     string
} // An HTTPError describes an HTTP error response (non-200 result).
// limited to maxErrorDetailLines and maxErrorDetailBytes

type errorDetailBuffer struct {
	r        io.ReadCloser
	buf      strings.Builder
	bufLines int
} // An errorDetailBuffer is an io.ReadCloser that copies up to
// maxErrorDetailLines into a buffer for later inspection.

type hookCloser struct {
	io.ReadCloser
	afterClose func()
}

type Interceptor struct {
	Scheme   string
	FromHost string
	ToHost   string
	Client   *http.Client
} // Interceptor is used to change the host, and maybe the client,
// for a request to point to a test host.

type Builder struct {
	WorkDir     string
	actionCache map[ // A Builder holds global state about a build.
	// It does not hold per-package state, because we
	// build packages in parallel, and the builder is shared.
	// the temporary work directory (ends in filepath.Separator)
	cacheKey]*Action
	flagCache map[ // a cache of already-constructed actions
	[2]string]bool
	gccCompilerIDCache map[ // a cache of supported compiler flags
	string]cache.ActionID
	IsCmdList           bool
	NeedError           bool
	NeedExport          bool
	NeedCompiledGoFiles bool
	AllowErrors         bool
	objdirSeq           int
	pkgSeq              int
	backgroundSh        *Shell
	exec                sync.Mutex
	readySema           chan bool
	ready               actionQueue
	id                  sync.Mutex
	toolIDCache         par.Cache[string, string]
	gccToolIDCache      map[ // cache for gccCompilerID
	// tool name -> tool ID
	string]string
	buildIDCache map[ // tool name -> tool ID
	string]string
} // file name -> build ID

type Actor interface {
	Act(*Builder, context.Context, *Action) error
} // An Actor runs an action.

type Action struct {
	Mode    string
	Package *load.Package
	Deps    []*// An Action represents a single action in the action graph.
	// the package this action works on
	Action
	Actor      Actor
	IgnoreFail bool
	TestOutput *bytes.Buffer
	Args       []string// actions that must happen before this one
	// test output buffer

	triggers []*// additional args for runProgram
	Action
	buggyInstall     bool
	TryCache         func(*Builder, *Action) bool
	CacheExecutable  bool
	Objdir           string
	Target           string
	built            string
	cachedExecutable string
	actionID         cache.ActionID
	buildID          string
	VetxOnly         bool
	needVet          bool
	needBuild        bool
	vetCfg           *vetConfig
	output           []byte// inverse of deps
	// vet config

	sh           *Shell
	pending      int
	priority     int
	Failed       *Action
	json         *actionJSON
	nonGoOverlay map[ // output redirect buffer (nil means use b.Print)
	// action graph information
	string]string
	traceSpan *trace.Span
} // map from non-.go source files to copied files in objdir. Nil if no overlay is used.

type actionJSON struct {
	ID         int
	Mode       string
	Package    string
	Deps       []int     `json:",omitempty"`
	IgnoreFail bool      `json:",omitempty"`
	Args       []string  `json:",omitempty"`
	Link       bool      `json:",omitempty"`
	Objdir     string    `json:",omitempty"`
	Target     string    `json:",omitempty"`
	Priority   int       `json:",omitempty"`
	Failed     bool      `json:",omitempty"`
	Built      string    `json:",omitempty"`
	VetxOnly   bool      `json:",omitempty"`
	NeedVet    bool      `json:",omitempty"`
	NeedBuild  bool      `json:",omitempty"`
	ActionID   string    `json:",omitempty"`
	BuildID    string    `json:",omitempty"`
	TimeReady  time.Time `json:",omitempty"`
	TimeStart  time.Time `json:",omitempty"`
	TimeDone   time.Time `json:",omitempty"`
	Cmd        []string
	CmdReal    time.Duration `json:",omitempty"`
	CmdUser    time.Duration `json:",omitempty"`
	CmdSys     time.Duration `json:",omitempty"`
} // `json:",omitempty"`

type cacheKey struct {
	mode string
	p    *load.Package
} // cacheKey is the key for the action cache.

type buildActor struct{ covMetaFileName string } // buildActor implements the Actor interface for package build
// actions. For most package builds this simply means invoking th
// *Builder.build method; in the case of "go test -cover" for
// a package with no test files, we stores some additional state
// information in the build actor to help with reporting.

type pgoActor struct{ input string } // pgoActor implements the Actor interface for preprocessing PGO profiles.

type buildCompiler struct{} // buildCompiler implements flag.Var.
// It implements Set by updating both
// BuildToolchain and buildContext.Compiler.

type coverFlag struct{ V flag.Value } // A coverFlag is a flag.Value that also implies -cover.

type commaListFlag struct {
	Vals *[]string// A commaListFlag is a flag.Value representing a comma-separated list.
}

type stringFlag struct{ val *string } // A stringFlag is a flag.Value representing a single string.

type vetConfig struct {
	ID         string
	Compiler   string
	Dir        string
	ImportPath string
	GoFiles    []string// vetConfig is the configuration passed to vet describing a single package.
	// canonical import path ("package path")

	NonGoFiles []string// absolute paths to package source files

	IgnoredFiles []string// absolute paths to package non-Go files

	ModulePath    string
	ModuleVersion string
	ImportMap     map[ // absolute paths to ignored source files
	// module version (may be "" on main module or module error)
	string]string
	PackageFile map[ // map import path in source code to package path
	string]string
	Standard map[ // map package path to .a file with export data
	string]bool
	PackageVetx map[ // map package path to whether it's in the standard library
	string]string
	VetxOnly                  bool
	VetxOutput                string
	GoVersion                 string
	SucceedOnTypecheckFailure bool
} // map package path to vetx data from earlier vet run
// awful hack; see #18395 and below

type toolchain interface {
	gc(b *Builder, a *Action, archive string, importcfg, embedcfg []byte,// gc runs the compiler in a specific directory on a set of files
	// and returns the name of the generated output file.
	symabis string, asmhdr bool, pgoProfile string, gofiles []string) (ofile string, out []byte, err error)
	cc(b *Builder, a *Action, ofile, cfile string) error
	asm(b *Builder, a *Action, sfiles []string) (// cc runs the toolchain's C compiler in a directory on a C file
	// to produce an output file.
	// asm runs the assembler in a specific directory on specific files
	// and returns a list of named output files.
	[]string, error)
	symabis(b *Builder, a *Action, sfiles []string) (// symabis scans the symbol ABIs from sfiles and returns the
	// path to the output symbol ABIs file, or "" if none.
	string, error)
	pack(b *Builder, a *Action, afile string, ofiles []string) error// pack runs the archive packer in a specific directory to create
	// an archive from a set of object files.
	// typically it is run in the object directory.

	ld(b *Builder, root *Action, targetPath, importcfg, mainpkg string) error
	ldShared(b *Builder, root *Action, toplevelactions []*// ld runs the linker to create an executable starting at mainpkg.
	// ldShared runs the linker to create a shared library containing the pkgs built by toplevelactions
	Action, targetPath, importcfg string, allactions []*Action) error
	compiler() string
	linker() string
}

type noToolchain struct{}

type gcToolchain struct{}

type gccgoToolchain struct{}

type version struct {
	name         string
	major, minor int
}

type Shell struct {
	action *Action
	*shellShared
} // A Shell runs shell commands and performs shell-like file system operations.
//
// Shell tracks context related to running commands, and form a tree much like
// context.Context.
// per-Builder state shared across Shells

type shellShared struct {
	workDir    string
	printLock  sync.Mutex
	printer    load.Printer
	scriptDir  string
	mkdirCache par.Cache[string, error]
} // shellShared is Shell state shared across all Shells derived from a single
// root shell (generally a single Builder).
// a cache of created directories

type cmdError struct {
	desc       string
	text       string
	importPath string
	needsPath  bool
} // Set if desc does not already include the import path

type workfileJSON struct {
	Go  string `json:",omitempty"`
	Use []useJSON// workfileJSON is the -json output data structure.

	Replace []replaceJSON
}

type useJSON struct {
	DiskPath string
	ModPath  string `json:",omitempty"`
}

type sequencer struct {
	maxWeight int64
	sem       *semaphore.Weighted
	prev      <-chan // A sequencer performs concurrent tasks that may write output, but emits that
	// output in a deterministic order.
	// weighted by input bytes (an approximate proxy for memory overhead)
	*reporterState
} // 1-buffered

type reporter struct {
	prev <-chan // A reporter reports output, warnings, and errors.
	*reporterState
	state *reporterState
}

type reporterState struct {
	out, err io.Writer
	exitCode int
} // reporterState carries the state of a reporter instance.
//
// Only one reporter at a time may have access to a reporterState.

type simplifier struct{}

type GoObj struct {
	TextHeader []byte
	Arch       string
	Data
}

type ErrGoObjOtherVersion struct{ magic []byte }

type objReader struct {
	a      *Archive
	b      *bio.Reader
	err    error
	offset int64
	limit  int64
	tmp    [256]byte
} // An objReader is an object file reader.

type Reader struct {
	f *os.File
	*bufio.Reader
} // Reader implements a seekable buffered io.Reader.

type Writer struct {
	f *os.File
	*bufio.Writer
} // Writer implements a seekable buffered io.Writer.

type excludedReader struct {
	r          io.Reader
	off        int64
	start, end int64
} // excludedReader wraps an io.Reader. Reading from it returns the bytes from
// the underlying reader, except that when the byte offset is within the
// range between start and end, it returns zero bytes.
// the range to be excluded (read as zero)

type Blob struct {
	typ    uint32
	offset uint32
} // type of entry
// offset of entry

type SuperBlob struct {
	magic  uint32
	length uint32
	count  uint32
} // magic number
// number of index entries following

type CodeDirectory struct {
	magic         uint32
	length        uint32
	version       uint32
	flags         uint32
	hashOffset    uint32
	identOffset   uint32
	nSpecialSlots uint32
	nCodeSlots    uint32
	codeLimit     uint32
	hashSize      uint8
	hashType      uint8
	_pad1         uint8
	pageSize      uint8
	_pad2         uint32
	scatterOffset uint32
	teamOffset    uint32
	_pad3         uint32
	codeLimit64   uint64
	execSegBase   uint64
	execSegLimit  uint64
	execSegFlags  uint64
} // magic number (CSMAGIC_CODEDIRECTORY)
// unused (must be zero)

type CodeSigCmd struct {
	Cmd      uint32
	Cmdsize  uint32
	Dataoff  uint32
	Datasize uint32
} // CodeSigCmd is Mach-O LC_CODE_SIGNATURE load command.
// file size of data in __LINKEDIT segment

type CoverPkgConfig struct {
	OutConfig    string
	PkgPath      string
	PkgName      string
	Granularity  string
	ModulePath   string
	Local        bool
	EmitMetaFile string
} // CoverPkgConfig is a bundle of information passed from the Go
// command to the cover command during "go build -cover" runs. The
// Go command creates and fills in a struct as below, then passes
// file containing the encoded JSON for the struct to the "cover"
// tool when instrumenting the source files in a Go package.
// EmitMetaFile if non-empty is the path to which the cover tool should
// directly emit a coverage meta-data file for the package, if the
// package has any functions in it. The go command will pass in a value
// here if we've been asked to run "go test -cover" on a package that
// doesn't have any *_test.go files.

type CoverFixupConfig struct {
	MetaVar            string
	MetaLen            int
	MetaHash           string
	Strategy           string
	CounterPrefix      string
	PkgIdVar           string
	CounterMode        string
	CounterGranularity string
} // CoverFixupConfig contains annotations/notes generated by the
// cmd/cover tool (during instrumentation) to be passed on to the
// compiler when the instrumented code is compiled. The cmd/cover tool
// creates a struct of this type, JSON-encodes it, and emits the
// result to a file, which the Go command then passes to the compiler
// when the instrumented package is built.
// Counter granularity (perblock or perfunc).

type MReader struct {
	f        *os.File
	rdr      *bio.Reader
	fileView []byte
	off      int64
}

type CovDataReader struct {
	vis    CovDataVisitor
	indirs []string// CovDataReader is a general-purpose helper/visitor object for
	// reading coverage data files in a structured way. Clients create a
	// CovDataReader to process a given collection of coverage data file
	// directories, then pass in a visitor object with methods that get
	// invoked at various important points. CovDataReader is intended
	// to facilitate common coverage data file operations such as
	// merging or intersecting data files, analyzing data files, or
	// dumping data files.

	matchpkg       func(name string) bool
	flags          CovDataReaderFlags
	err            error
	verbosityLevel int
}

type CovDataVisitor interface {
	BeginPod(p pods.Pod)
	EndPod(p pods.Pod)
	VisitMetaDataFile(mdf string, mfr *decodemeta.CoverageMetaFileReader)
	BeginCounterDataFile(cdf string, cdr *decodecounter.CounterDataReader, dirIdx int)
	EndCounterDataFile(cdf string, cdr *decodecounter.CounterDataReader, dirIdx int)
	VisitFuncCounterData(payload decodecounter.FuncPayload)
	EndCounters()
	BeginPackage(pd *decodemeta.CoverageMetaDataDecoder, pkgIdx uint32)
	EndPackage(pd *decodemeta.CoverageMetaDataDecoder, pkgIdx uint32)
	VisitFunc(pkgIdx uint32, fnIdx uint32, fd *coverage.FuncDesc)
	Finish()
} // Invoked at the start and end of a given pod (a pod here is a
// specific coverage meta-data files with the counter data files
// that correspond to it).
// Invoked when all counter + meta-data file processing is complete.

type Disasm struct {
	syms []objfile.// Disasm is a disassembler for a given File.
	Sym
	pcln objfile.Liner
	text []byte// symbols in file, sorted by address
	// pcln table

	textStart uint64
	textEnd   uint64
	goarch    string
	disasm    disasmFunc
	byteOrder binary.ByteOrder
} // bytes of text segment (actual instructions)
// byte order for goarch

type CachedFile struct {
	FileName string
	Lines    [][// CachedFile contains the content of a file split into lines.
	]byte
}

type FileCache struct {
	files  *list.List
	maxLen int
} // FileCache is a simple LRU cache of file contents.

type textReader struct {
	code []byte
	pc   uint64
}

type Dir struct {
	importPath string
	dir        string
	inModule   bool
} // A Dir describes a directory holding code by specifying
// the expected import path and the file system directory.
// file system directory

type Dirs struct {
	scan chan Dir
	hist []Dir// Dirs is a structure for scanning the directory tree.
	// Its Next method returns the next Go source directory it finds.
	// Although it can be used to scan the tree multiple times, it
	// only walks the tree once, caching the data it finds.
	// Directories generated by walk.

	offset int
} // History of reported Dirs.
// Counter for Next.

type moduleJSON struct{ Path, Dir, GoVersion string }

type pkgBuffer struct {
	pkg     *Package
	printed bool
	bytes.Buffer
} // pkgBuffer is a wrapper for bytes.Buffer that prints a package clause the
// first time Write is called.
// Prevent repeated package clauses.

type ExportedType struct {
	ExportedField   int
	unexportedField int
	ExportedEmbeddedType
	*ExportedEmbeddedType
	*qualified.ExportedEmbeddedType
	unexportedType
	*unexportedType
	io.Reader
	error
} // Comment about exported type.
// Comment on line with embedded error.

type ExportedStructOneField struct {
	OnlyField int
} // the only field

type ExportedInterface interface {
	ExportedMethod()
	unexportedMethod()
	io.Reader
	error
} // Comment about exported interface.
// Comment on line with embedded error.

type ExportedFormattedType struct{ ExportedField int }

type SimpleConstraint interface{ ~int | ~float64 }

type TildeConstraint interface{ ~int }

type StructConstraint interface{ struct{ F int } }

type FnState struct {
	Name     string
	Info     Sym
	Loc      Sym
	Ranges   Sym
	Absfn    Sym
	StartPC  Sym
	StartPos src.Pos
	Size     int64
	External bool
	Scopes   []Scope// This container is used by the PutFunc* variants below when
	// creating the DWARF subprogram DIE(s) for a function.

	InlCalls          InlCalls
	UseBASEntries     bool
	dictIndexToOffset []int64
}

type InlCalls struct{ Calls []InlCall }

type InlCall struct {
	InlIndex  int
	CallPos   src.Pos
	AbsFunSym Sym
	Children  []int// index into ctx.InlTree describing the call inlined here
	// Indices of child inlines within Calls array above.

	InlVars []*// entries in this list are PAUTO's created by the inliner to
	// capture the promoted formals and locals of the inlined callee.
	Var
	Ranges []Range// PC ranges for this inlined call.

	Root bool
} // Root call (not a child of some other call).

type dwAttrForm struct {
	attr uint16
	form uint8
} /*
 * Defining Abbrevs. This is hardcoded on a per-platform basis (that is,
 * each platform will see a fixed abbrev table for all objects); the number
 * of abbrev entries is fairly small (compared to C++ objects).  The DWARF
 * spec places no restriction on the ordering of attributes in the
 * Abbrevs and DIEs, and we will always write them out in the order
 * of declaration in the abbrev.
 */

type dwAbbrev struct {
	tag      uint8
	children uint8
	attr     []dwAttrForm
}

type DWAttr struct {
	Link  *DWAttr
	Atr   uint16
	Cls   uint8
	Value int64
	Data  interface{}
} // DWAttr represents an attribute of a DWDie.
//
// For DW_CLS_string and _block, value should contain the length, and
// data the data, for _reference, value is 0 and data is a DWDie* to
// the referenced instance, for all others, value is the whole thing
// and data is null.
// DW_CLS_

type DWDie struct {
	Abbrev int
	Link   *DWDie
	Child  *DWDie
	Attr   *DWAttr
	Sym    Sym
} // DWDie represents a DWARF debug info entry.

type Buffer struct {
	old []byte// A Buffer is a queue of edits to apply to a given byte slice.

	q edits
}

type edit struct {
	start int
	end   int
	new   string
} // An edit records a single text modification: change the bytes in [start,end) to new.

type FuncInfo struct {
	Args      uint32
	Locals    uint32
	FuncID    abi.FuncID
	FuncFlag  abi.FuncFlag
	StartLine int32
	File      []CUFileIndex// FuncInfo is serialized as a symbol (aux symbol). The symbol data is
	// the binary encoding of the struct below.

	InlTree []InlTreeNode
}

type FuncInfoLengths struct {
	NumFile     uint32
	FileOff     uint32
	NumInlTree  uint32
	InlTreeOff  uint32
	Initialized bool
} // FuncInfoLengths is a cache containing a roadmap of offsets and
// lengths for things within a serialized FuncInfo. Each length field
// stores the number of items (e.g. files, inltree nodes, etc), and the
// corresponding "off" field stores the byte offset of the start of
// the items in question.

type InlTreeNode struct {
	Parent   int32
	File     CUFileIndex
	Line     int32
	Func     SymRef
	ParentPC int32
} // InlTreeNode is the serialized form of FileInfo.InlTree.

type extra struct {
	name string
	abi  int
}

type Header struct {
	Magic       string
	Fingerprint FingerprintType
	Flags       uint32
	Offsets     [NBlk]uint32
} // File header.
// TODO: probably no need to export this.

type ImportedPkg struct {
	Pkg         string
	Fingerprint FingerprintType
} // Autolib

type SymRef struct {
	PkgIdx uint32
	SymIdx uint32
} // Symbol reference.

type LoadCmd struct {
	Cmd macho.LoadCmd
	Len uint32
} // LoadCmd is macho.LoadCmd with its length, which is also
// the load command header in the Mach-O file.

type LoadCmdReader struct {
	offset, next int64
	f            io.ReadSeeker
	order        binary.ByteOrder
}

type LoadCmdUpdater struct{ LoadCmdReader }

type ctxt5 struct {
	ctxt       *obj.Link
	newprog    obj.ProgAlloc
	cursym     *obj.LSym
	printp     *obj.Prog
	blitrl     *obj.Prog
	elitrl     *obj.Prog
	autosize   int64
	instoffset int64
	pc         int64
	pool       struct {
		start uint32
		size  uint32
		extra uint32
	}
} // ctxt5 holds state while assembling a single function.
// Each function gets a fresh ctxt5.
// This allows for multiple functions to be safely concurrently assembled.

type Optab struct {
	as       obj.As
	a1       uint8
	a2       int8
	a3       uint8
	type_    uint8
	size     int8
	param    int16
	flag     int8
	pcrelsiz uint8
	scond    uint8
} // optional flags accepted by the instruction

type ctxt7 struct {
	ctxt       *obj.Link
	newprog    obj.ProgAlloc
	cursym     *obj.LSym
	blitrl     *obj.Prog
	elitrl     *obj.Prog
	autosize   int32
	extrasize  int32
	instoffset int64
	pc         int64
	pool       struct {
		start uint32
		size  uint32
	}
} // ctxt7 holds state while assembling a single function.
// Each function gets a fresh ctxt7.
// This allows for multiple functions to be safely concurrently assembled.

type codeBuffer struct{ data *[]byte }

type dwCtxt struct{ *Link } // implement dwarf.Context

type DwarfFixupTable struct {
	ctxt   *Link
	mu     sync.Mutex
	symtab map[ // This table is designed to aid in the creation of references between
	// DWARF subprogram DIEs.
	//
	// In most cases when one DWARF DIE has to refer to another DWARF DIE,
	// the target of the reference has an LSym, which makes it easy to use
	// the existing relocation mechanism. For DWARF inlined routine DIEs,
	// however, the subprogram DIE has to refer to a child
	// parameter/variable DIE of the abstract subprogram. This child DIE
	// doesn't have an LSym, and also of interest is the fact that when
	// DWARF generation is happening for inlined function F within caller
	// G, it's possible that DWARF generation hasn't happened yet for F,
	// so there is no way to know the offset of a child DIE within F's
	// abstract function. Making matters more complex, each inlined
	// instance of F may refer to a subset of the original F's variables
	// (depending on what happens with optimization, some vars may be
	// eliminated).
	//
	// The fixup table below helps overcome this hurdle. At the point
	// where a parameter/variable reference is made (via a call to
	// "ReferenceChildDIE"), a fixup record is generate that records
	// the relocation that is targeting that child variable. At a later
	// point when the abstract function DIE is emitted, there will be
	// a call to "RegisterChildDIEOffsets", at which point the offsets
	// needed to apply fixups are captured. Finally, once the parallel
	// portion of the compilation is done, fixups can actually be applied
	// during the "Finalize" method (this can't be done during the
	// parallel portion of the compile due to the possibility of data
	// races).
	//
	// This table is also used to record the "precursor" function node for
	// each function that is the target of an inline -- child DIE references
	// have to be made with respect to the original pre-optimization
	// version of the function (to allow for the fact that each inlined
	// body may be optimized differently).
	*LSym]int
	svec []symFixups// maps abstract fn LSYM to index in svec

	precursor map[*LSym]fnState
} // maps fn Lsym to precursor Node, absfn sym

type symFixups struct {
	fixups   []relFixup
	doffsets []declOffset
	inlIndex int32
	defseen  bool
}

type declOffset struct {
	dclIdx int32
	offset int32
} // Index of variable within DCL list of pre-optimization function
// Offset of var's child DIE with respect to containing subprogram DIE

type relFixup struct {
	refsym *LSym
	relidx int32
	dclidx int32
}

type fnState struct {
	precursor Func
	absfn     *LSym
} // precursor function
// abstract function symbol

type InlTree struct {
	nodes []InlinedCall// InlTree is a collection of inlined calls. The Parent field of an
	// InlinedCall is the index of another InlinedCall in InlTree.
	//
	// The compiler maintains a global inlining tree and adds a node to it
	// every time a function is inlined. For example, suppose f() calls g()
	// and g has two calls to h(), and that f, g, and h are inlineable:
	//
	//	 1 func main() {
	//	 2     f()
	//	 3 }
	//	 4 func f() {
	//	 5     g()
	//	 6 }
	//	 7 func g() {
	//	 8     h()
	//	 9     h()
	//	10 }
	//	11 func h() {
	//	12     println("H")
	//	13 }
	//
	// Assuming the global tree starts empty, inlining will produce the
	// following tree:
	//
	//	[]InlinedCall{
	//	  {Parent: -1, Func: "f", Pos: <line 2>},
	//	  {Parent:  0, Func: "g", Pos: <line 5>},
	//	  {Parent:  1, Func: "h", Pos: <line 8>},
	//	  {Parent:  1, Func: "h", Pos: <line 9>},
	//	}
	//
	// The nodes of h inlined into main will have inlining indexes 2 and 3.
	//
	// Eventually, the compiler extracts a per-function inlining tree from
	// the global inlining tree (see pcln.go).
}

type InlinedCall struct {
	Parent   int
	Pos      src.XPos
	Func     *LSym
	Name     string
	ParentPC int32
} // InlinedCall is a node in an InlTree.
// PC of instruction just before inlined body. Only valid in local trees.

type Addr struct {
	Reg    int16
	Index  int16
	Scale  int16
	Type   AddrType
	Name   AddrName
	Class  int8
	Offset int64
	Sym    *LSym
	Val    interface{}
} // Sometimes holds a register.
// argument value:
//	for TYPE_SCONST, a string
//	for TYPE_FCONST, a float64
//	for TYPE_BRANCH, a *Prog (optional)
//	for TYPE_TEXTSIZE, an int32 (optional)

type Prog struct {
	Ctxt     *Link
	Link     *Prog
	From     Addr
	RestArgs []AddrPos// Prog describes a single machine instruction.
	//
	// The general instruction form is:
	//
	//	(1) As.Scond From [, ...RestArgs], To
	//	(2) As.Scond From, Reg [, ...RestArgs], To, RegTo2
	//
	// where As is an opcode and the others are arguments:
	// From, Reg are sources, and To, RegTo2 are destinations.
	// RestArgs can hold additional sources and destinations.
	// Usually, not all arguments are present.
	// For example, MOVL R1, R2 encodes using only As=MOVL, From=R1, To=R2.
	// The Scond field holds additional condition bits for systems (like arm)
	// that have generalized conditional execution.
	// (2) form is present for compatibility with older code,
	// to avoid too much changes in a single swing.
	// (1) scheme is enough to express any kind of operand combination.
	//
	// Jump instructions use the To.Val field to point to the target *Prog,
	// which must be in the same linked list as the jump instruction.
	//
	// The Progs for a given function are arranged in a list linked through the Link field.
	//
	// Each Prog is charged to a specific source line in the debug information,
	// specified by Pos.Line().
	// Every Prog has a Ctxt field that defines its context.
	// For performance reasons, Progs are usually bulk allocated, cached, and reused;
	// those bulk allocators should always be used, rather than new(Prog).
	//
	// The other fields not yet mentioned are for use by the back ends and should
	// be left zeroed by creators of Prog lists.
	// first source operand

	To     Addr
	Pool   *Prog
	Forwd  *Prog
	Rel    *Prog
	Pc     int64
	Pos    src.XPos
	Spadj  int32
	As     As
	Reg    int16
	RegTo2 int16
	Mark   uint16
	Optab  uint16
	Scond  uint8
	Back   uint8
	Ft     uint8
	Tt     uint8
	Isize  uint8
} // can pack any operands that not fit into {Prog.From, Prog.To}, same kinds of operands are saved in order
// for x86 back end: size of the instruction in bytes

type AddrPos struct {
	Addr
	Pos OperandPos
} // AddrPos indicates whether the operand is the source or the destination.

type LSym struct {
	Name string
	Type objabi.SymKind
	Attribute
	Size   int64
	Gotype *LSym
	P      []byte// An LSym is the sort of symbol that is written to an object file.
	// It represents Go symbols in a flat pkg+"."+name namespace.

	R      []Reloc
	Extra  *interface{}
	Pkg    string
	PkgIdx int32
	SymIdx int32
} // *FuncInfo, *VarInfo, *FileInfo, or *TypeInfo, if present

type JumpTable struct {
	Sym     *LSym
	Targets []*// JumpTable represents a table used for implementing multi-way
	// computed branching, used typically for implementing switches.
	// Sym is the table itself, and Targets is a list of target
	// instructions to go to for the computed branch index.
	Prog
}

type VarInfo struct{ dwarfInfoSym *LSym }

type FileInfo struct {
	Name string
	Size int64
} // A FileInfo contains extra fields for SDATA symbols backed by files.
// (If LSym.Extra is a *FileInfo, LSym.P == nil.)
// length of file

type TypeInfo struct {
	Type interface{}
} // A TypeInfo contains information for a symbol
// that contains a runtime._type.
// a *cmd/compile/internal/types.Type

type ItabInfo struct {
	Type interface{}
} // An ItabInfo contains information for a symbol
// that contains a runtime.itab.
// a *cmd/compile/internal/types.Type

type WasmFuncType struct {
	Params []WasmField// WasmFuncType represents a WebAssembly (WASM) function type with
	// parameters and results translated into WASM types based on the Go function
	// declaration.
	// Params holds the function parameter fields.

	Results []WasmField// Results holds the function result fields.

}

type WasmField struct {
	Type   WasmFieldType
	Offset int64
} // Offset holds the frame-pointer-relative locations for Go's stack-based
// ABI. This is used by the src/cmd/internal/wasm package to map WASM
// import parameters to the Go stack in a wrapper function.

type InlMark struct {
	p  *Prog
	id int32
} // When unwinding from an instruction in an inlined body, mark
// where we should unwind to.
// id records the global inlining id of the inlined body.
// p records the location of an instruction in the parent (inliner) frame.

type Pcln struct {
	Pcsp     *LSym
	Pcfile   *LSym
	Pcline   *LSym
	Pcinline *LSym
	Pcdata   []*// Aux symbols for pcln
	LSym
	Funcdata  []*LSym
	UsedFiles map[goobj.CUFileIndex]struct{}
	InlTree   InlTree
} // file indices used while generating pcfile
// per-function inlining tree extracted from the global tree

type Auto struct {
	Asym    *LSym
	Aoffset int32
	Name    AddrName
	Gotype  *LSym
}

type RegSpill struct {
	Addr           Addr
	Reg            int16
	Reg2           int16
	Spill, Unspill As
} // RegSpill provides spill/fill information for a register-resident argument
// to a function.  These need spilling/filling in the safepoint/stackgrowth case.
// At the time of fill/spill, the offset must be adjusted by the architecture-dependent
// adjustment to hardware SP that occurs in a call instruction.  E.g., for AMD64,
// at Offset+8 because the return address was pushed.
// If not 0, a second register to spill at Addr+regSize. Only for some archs.

type Link struct {
	Headtype           objabi.HeadType
	Arch               *LinkArch
	Debugasm           int
	Debugvlog          bool
	Debugpcln          string
	Flag_shared        bool
	Flag_dynlink       bool
	Flag_linkshared    bool
	Flag_optimize      bool
	Flag_locationlists bool
	Flag_noRefName     bool
	Retpoline          bool
	Flag_maymorestack  string
	Bso                *bufio.Writer
	Pathname           string
	Pkgpath            string
	hashmu             sync.Mutex
	hash               map[ // Link holds the context for writing object code from a compiler
	// to be linker input or for reading that input into the linker.
	// protects hash, funchash
	string]*LSym
	funchash map[ // name -> sym mapping
	string]*LSym
	statichash map[ // name -> sym mapping for ABIInternal syms
	string]*LSym
	PosTable    src.PosTable
	InlTree     InlTree
	DwFixups    *DwarfFixupTable
	DwTextCount int
	Imports     []goobj.// name -> sym mapping for static syms
	// global inlining tree used by gc/inl.go
	ImportedPkg
	DiagFunc        func(string, ...interface{})
	DiagFlush       func()
	DebugInfo       func(ctxt *Link, fn *LSym, info *LSym, curfn Func) ([]dwarf.Scope, dwarf.InlCalls)
	GenAbstractFunc func(fn *LSym)
	Errors          int
	InParallel      bool
	UseBASEntries   bool
	IsAsm           bool
	Std             bool
	Text            []*// parallel backend phase in effect
	// state for writing objects
	LSym
	Data      []*LSym
	constSyms []*// Constant symbols (e.g. $i64.*) are data symbols created late
	// in the concurrent phase. To ensure a deterministic order, we
	// add them to a separate list, sort at the end, and append it
	// to Data.
	LSym
	SEHSyms []*// Windows SEH symbols are also data symbols that can be created
	// concurrently.
	LSym
	pkgIdx map[ // pkgIdx maps package path to index. The index is used for
	// symbol reference in the object file.
	string]int32
	defs         []*LSym
	hashed64defs []*// list of defined symbols in the current package
	LSym
	hasheddefs []*// list of defined short (64-bit or less) hashed (content-addressable) symbols
	LSym
	nonpkgdefs []*// list of defined hashed (content-addressable) symbols
	LSym
	nonpkgrefs []*// list of defined non-package symbols
	LSym
	Fingerprint goobj.FingerprintType
} // list of referenced non-package symbols
// fingerprint of symbol indices, to catch index mismatch

type LinkArch struct {
	*sys.Arch
	Init       func(*Link)
	ErrorCheck func(*Link, *LSym)
	Preprocess func(*Link, *LSym, ProgAlloc)
	Assemble   func(*Link, *LSym, ProgAlloc)
	Progedit   func(*Link, *Prog, ProgAlloc)
	SEH        func(*Link, *LSym) *LSym
	UnaryDst   map[ // LinkArch is the definition of a single architecture.
	As]bool
	DWARFRegisters map[ // Instruction takes one operand, a destination.
	int16]int16
}

type ctxt0 struct {
	ctxt       *obj.Link
	newprog    obj.ProgAlloc
	cursym     *obj.LSym
	autosize   int32
	instoffset int64
	pc         int64
} // ctxt0 holds state while assembling a single function.
// Each function gets a fresh ctxt0.
// This allows for multiple functions to be safely concurrently assembled.

type Sch struct {
	p       obj.Prog
	set     Dep
	used    Dep
	soffset int32
	size    uint8
	nop     uint8
	comp    bool
}

type pcinlineState struct {
	globalToLocal map[ // pcinlineState holds the state used to create a function's inlining
	// tree and the PC-value table that maps PCs to nodes in that tree.
	int]int
	localTree InlTree
}

type PCIter struct {
	p []byte// PCIter iterates over encoded pcdata tables.

	PC      uint32
	NextPC  uint32
	PCScale uint32
	Value   int32
	start   bool
	Done    bool
}

type Plist struct {
	Firstpc *Prog
	Curfn   Func
}

type ctxt9 struct {
	ctxt       *obj.Link
	newprog    obj.ProgAlloc
	cursym     *obj.LSym
	autosize   int32
	instoffset int64
	pc         int64
} // ctxt9 holds state while assembling a single function.
// Each function gets a fresh ctxt9.
// This allows for multiple functions to be safely concurrently assembled.

type PrefixableOptab struct {
	Optab
	minGOPPC64 int
	pfxsize    int8
} // These are opcodes above which may generate different sequences depending on whether prefix opcode support
// is available
// Instruction sequence size when prefixed opcodes are used

type inst struct {
	opcode uint32
	funct3 uint32
	rs1    uint32
	rs2    uint32
	csr    int64
	funct7 uint32
}

type encoding struct {
	encode   func(*instruction) uint32
	validate func(*obj.Link, *instruction)
	length   int
} // encode returns the machine code for an instruction
// length of encoded instruction; 0 for pseudo-ops, 4 otherwise

type instructionData struct {
	enc     encoding
	immForm obj.As
	ternary bool
} // instructionData specifies details relating to a RISC-V instruction.
// immediate form of this instruction

type instruction struct {
	p      *obj.Prog
	as     obj.As
	rd     uint32
	rs1    uint32
	rs2    uint32
	rs3    uint32
	imm    int64
	funct3 uint32
	funct7 uint32
} // Prog that instruction is for
// Function 7 (or Function 2)

type ctxtz struct {
	ctxt       *obj.Link
	newprog    obj.ProgAlloc
	cursym     *obj.LSym
	autosize   int32
	instoffset int64
	pc         int64
} // ctxtz holds state while assembling a single function.
// Each function gets a fresh ctxtz.
// This allows for multiple functions to be safely concurrently assembled.

type RotateParams struct {
	Start  uint8
	End    uint8
	Amount uint8
} // RotateParams represents the immediates required for a "rotate
// then ... selected bits instruction".
//
// The Start and End values are the indexes that represent
// the masked region. They are inclusive and are in big-
// endian order (bit 0 is the MSB, bit 63 is the LSB). They
// may wrap around.
//
// Some examples:
//
// Masked region             | Start | End
// --------------------------+-------+----
// 0x00_00_00_00_00_00_00_0f | 60    | 63
// 0xf0_00_00_00_00_00_00_00 | 0     | 3
// 0xf0_00_00_00_00_00_00_0f | 60    | 3
//
// The Amount value represents the amount to rotate the
// input left by. Note that this rotation is performed
// before the masked region is used.
// amount to rotate left

type opSuffixSet struct {
	arch  string
	cconv func(suffix uint8) string
} // opSuffixSet is like regListSet, but for opcode suffixes.
//
// Unlike some other similar structures, uint8 space is not
// divided by its own values set (because there are only 256 of them).
// Instead, every arch may interpret/format all 8 bits as they like,
// as long as they register proper cconv function for it.

type regSet struct {
	lo    int
	hi    int
	Rconv func(int) string
}

type regListSet struct {
	lo     int64
	hi     int64
	RLconv func(int64) string
}

type spcSet struct {
	lo      int64
	hi      int64
	SPCconv func(int64) string
} // Special operands

type opSet struct {
	lo    As
	names []string
}

type movtab struct {
	as   obj.As
	ft   uint8
	f3t  uint8
	tt   uint8
	code uint8
	op   [4]uint8
}

type nopPad struct {
	p *obj.Prog
	n int32
} // Instruction before the pad
// Size of the pad

type AsmBuf struct {
	buf      [100]byte
	off      int
	rexflag  int
	vexflag  bool
	evexflag bool
	rep      bool
	repn     bool
	lock     bool
	evex     evexBits
} // AsmBuf is a simple buffer to assemble variable-length x86 instructions into
// and hold assembly state.
// Initialized when evexflag is true

type evexBits struct {
	b1     byte
	b2     byte
	opcode byte
} // evexBits stores EVEX prefix info that is used during instruction encoding.
// Associated instruction opcode.

type evexSuffix struct {
	rounding  byte
	sae       bool
	zeroing   bool
	broadcast bool
} // evexSuffixBits carries instruction EVEX suffix set flags.
//
// Examples:
//
//	"RU_SAE.Z" => {rounding: 3, zeroing: true}
//	"Z" => {zeroing: true}
//	"BCST" => {broadcast: true}
//	"SAE.Z" => {sae: true, zeroing: true}

type sehbuf struct {
	ctxt *obj.Link
	data []byte
	off  int
}

type ytab struct {
	zcase   uint8
	zoffset uint8
	args    argList
} // Last arg is usually destination.
// For unary instructions unaryDst is used to determine
// if single argument is a source or destination.

type versionFlag struct{}

type debugField struct {
	name         string
	help         string
	concurrentOk bool
	val          interface{}
} // true if this field/flag is compatible with concurrent compilation
// *int or *string

type DebugFlag struct {
	tab          map[string]debugField
	concurrentOk *bool
	debugSSA     DebugSSA
} // this is non-nil only for compiler's DebugFlags, but only compiler has concurrent:ok fields
// this is non-nil only for compiler's DebugFlags.

type PkgSpecial struct {
	Runtime      bool
	NoInstrument bool
	NoRaceFunc   bool
	AllowAsmABI  bool
} // PkgSpecial indicates special build properties of a given runtime-related
// package.
// AllowAsmABI indicates that assembly in this package is allowed to use ABI
// selectors in symbol names. Generally this is needed for packages that
// interact closely with the runtime package or have performance-critical
// assembly.

type elfFile struct{ elf *elf.File }

type goobjFile struct {
	goobj *archive.GoObj
	r     *goobj.Reader
	f     *os.File
	arch  *sys.Arch
}

type goobjReloc struct {
	Off  int32
	Size uint8
	Type objabi.RelocType
	Add  int64
	Sym  string
}

type machoFile struct{ macho *macho.File }

type RelocStringer interface {
	String(insnOffset uint64) string
} // insnOffset is the offset of the instruction containing the relocation
// from the start of the symbol containing the relocation.

type Liner interface {
	PCToLine(uint64) (string, int, *gosym.Func)
} // Given a pc, returns the corresponding file, line, and function data.
// If unknown, returns "",0,nil.

type peFile struct{ pe *pe.File }

type plan9File struct{ plan9 *plan9obj.File }

type xcoffFile struct{ xcoff *xcoff.File }

type utsname struct {
	Sysname  [257]byte
	Nodename [257]byte
	Release  [257]byte
	Version  [257]byte
	Machine  [257]byte
}

type Queue struct {
	maxActive int
	st        chan queueState
} // Queue manages a set of work items to be executed in parallel. The number of
// active work items is limited, and excess items are queued sequentially.

type queueState struct {
	active  int
	backlog []func// number of goroutines processing work; always nonzero when len(backlog) > 0
	()
	idle chan struct{}
} // if non-nil, closed when active becomes 0

type Work[T comparable] struct {
	f       func(T)
	running int
	mu      sync.Mutex
	added   map[ // Work manages a set of work items to be executed in parallel, at most once each.
	// The items in the set must all be valid map keys.
	// total number of runners
	T]bool
	todo []T// items added to set

	wait    sync.Cond
	waiting int
} // items yet to be run
// number of runners waiting for todo

type ErrCache[K comparable, V any] struct{ Cache[K, errValue[V]] } // ErrCache is like Cache except that it also stores
// an error value alongside the cached value V.

type errValue[V any] struct {
	v   V
	err error
}

type cacheEntry[V any] struct {
	done   atomic.Bool
	mu     sync.Mutex
	result V
}

type NamedCallEdge struct {
	CallerName     string
	CalleeName     string
	CallSiteOffset int
} // NamedCallEdge identifies a call edge by linker symbol names and call site
// offset.
// Line offset from function start line.

type NamedEdgeMap struct {
	Weight map[ // NamedEdgeMap contains all unique call edges in the profile and their
	// edge weight.
	NamedCallEdge]int64
	ByWeight []NamedCallEdge// ByWeight lists all keys in Weight, sorted by edge weight from
	// highest to lowest.

}

type funcCmd struct {
	usage CmdUsage
	run   func(*State, ...string) (WaitFunc, error)
} // A funcCmd implements Cmd using a function value.

type stopError struct{ msg string } // stopError is the sentinel error type returned by the Stop command.

type waitError struct {
	errs []*// A waitError wraps one or more errors returned by background commands.
	CommandError
}

type funcCond struct {
	eval  func(*State) (bool, error)
	usage CondUsage
}

type prefixCond struct {
	eval  func(*State, string) (bool, error)
	usage CondUsage
}

type boolCond struct {
	v     bool
	usage CondUsage
}

type onceCond struct {
	eval  func() (bool, error)
	usage CondUsage
}

type cachedCond struct {
	m     sync.Map
	eval  func(string) (bool, error)
	usage CondUsage
}

type Engine struct {
	Cmds map[ // An Engine stores the configuration for executing a set of scripts.
	//
	// The same Engine may execute multiple scripts concurrently.
	string]Cmd
	Conds map[string]Cond
	Quiet bool
} // If Quiet is true, Execute deletes log prints from the previous
// section when starting a new section.

type CmdUsage struct {
	Summary string
	Args    string
	Detail  []string// A CmdUsage describes the usage of a Cmd, independent of its name
	// (which can change based on its registration).
	// a brief synopsis of the command's arguments (only)

	Async      bool
	RegexpArgs func(rawArgs ...string) []int// zero or more sentences in the style of the Description section of a Unix 'man' page
	// RegexpArgs reports which arguments, if any, should be treated as regular
	// expressions. It takes as input the raw, unexpanded arguments and returns
	// the list of argument indices that will be interpreted as regular
	// expressions.
	//
	// If RegexpArgs is nil, all arguments are assumed not to be regular
	// expressions.

}

type Cond interface {
	Eval(s *State, suffix string) (bool, error)
	Usage() *CondUsage
} // A Cond is a condition deciding whether a command should be run.
// Usage returns the usage for the condition, which the caller must not modify.

type CondUsage struct {
	Summary string
	Prefix  bool
} // A CondUsage describes the usage of a Cond, independent of its name
// (which can change based on its registration).
// If Prefix is true, the condition is a prefix and requires a
// colon-separated suffix (like "[GOOS:linux]" for the "GOOS" condition).
// The suffix may be the empty string (like "[prefix:]").

type command struct {
	file  string
	line  int
	want  expectedStatus
	conds []condition// A command is a complete command parsed from a script.

	name    string
	rawArgs [][// all must be satisfied
	// the name of the command; must be non-empty
	]argFragment
	args       []string
	background bool
} // shell-expanded arguments following name
// command should run in background (ends with a trailing &)

type argFragment struct {
	s      string
	quoted bool
} // if true, disable variable expansion for this fragment

type condition struct {
	want bool
	tag  string
}

type CommandError struct {
	File string
	Line int
	Op   string
	Args []string// A CommandError describes an error resulting from attempting to execute a
	// specific command.

	Err error
}

type UsageError struct {
	Name    string
	Command Cmd
} // A UsageError reports the valid arguments for a command.
//
// It may be returned in response to invalid arguments.

type ToolReplacement struct {
	ToolName        string
	ReplacementPath string
	EnvVar          string
} // ToolReplacement records the name of a tool to replace
// within a given GOROOT for script testing purposes.
// env var setting (e.g. "FOO=BAR")

type skipError struct{ msg string }

type backgroundCmd struct {
	*command
	wait WaitFunc
}

type XPos struct {
	index int32
	lico
} // XPos is a more compact representation of Pos.

type PosTable struct {
	baseList []*// A PosTable tracks Pos -> XPos conversions and vice versa.
	// Its zero value is a ready-to-use PosTable.
	PosBase
	indexMap map[*PosBase]int
	nameMap  map[string]int
} // Maps file symbol name to index for debug information.

type dummyCounter struct{}

type Converter struct {
	w        io.Writer
	pkg      string
	mode     Mode
	start    time.Time
	testName string
	report   []*// A Converter holds the state of a test-to-JSON conversion.
	// It implements io.WriteCloser; the caller writes test output in,
	// and the converter writes JSON output to w.
	// name of current test, for output attribution
	event
	result      string
	input       lineBuffer
	output      lineBuffer
	needMarker  bool
	failedBuild string
} // pending test result reports (nested for subtests)
// failedBuild is set to the package ID of the cause of a build failure,
// if that's what caused this test to fail.

type lineBuffer struct {
	b []byte// A lineBuffer is an I/O buffer that reacts to writes by invoking
	// input-processing callbacks on whole lines or (for long lines that
	// have been split) line fragments.
	//
	// It should be initialized with b set to a buffer of length 0 but non-zero capacity,
	// and line and part set to the desired input processors.
	// The lineBuffer will call line(x) for any whole line x (including the final newline)
	// that fits entirely in cap(b). It will handle input lines longer than cap(b) by
	// calling part(x) for sections of the line. The line will be split at UTF8 boundaries,
	// and the final call to part for a long line includes the final newline.

	mid  bool
	line func([]byte)// buffer
	// whether we're in the middle of a long line

	part func([]byte)// line callback

} // partial line callback

type Metrics struct {
	gc        Flags
	marks     []*mark
	curMark   *mark
	filebase  string
	pprofFile *os.File
}

type mark struct {
	name              string
	startM, endM, gcM runtime.MemStats
	startT, endT      time.Time
}

type Examiner struct {
	dies        []*dwarf.Entry
	idxByOffset map[dwarf.Offset]int
	kids        map[int][]int
	parent      map[int]int
	byname      map[string][]int
}

type ArHdr struct {
	name string
	date string
	uid  string
	gid  string
	mode string
	size string
	fmag string
}

type relocSymState struct {
	target *Target
	ldr    *loader.Loader
	err    *ErrorReporter
	syms   *ArchSyms
} // relocSymState hold state information needed when making a series of
// successive calls to relocsym(). The items here are invariant
// (meaning that they are set up once initially and then don't change
// during the execution of relocsym), with the exception of a slice
// used to facilitate batch allocation of external relocations. Calls
// to relocsym happen in parallel; the assumption is that each
// parallel thread will have its own state object.

type GCProg struct {
	ctxt *Link
	sym  *loader.SymbolBuilder
	w    gcprog.Writer
}

type dodataState struct {
	ctxt *Link
	data [sym.SXREF][]loader.// dodataState holds bits of state information needed by dodata() and the
	// various helpers it calls. The lifetime of these items should not extend
	// past the end of dodata().
	// Data symbols bucketed by type.
	Sym
	dataMaxAlign [sym.SXREF]int32
	symGroupType []sym.// Max alignment for each flavor of data symbol.
	// Overridden sym type
	SymKind
	datsize int64
} // Current data size so far.

type symNameSize struct {
	name string
	sz   int64
	val  int64
	sym  loader.Sym
}

type deadcodePass struct {
	ctxt        *Link
	ldr         *loader.Loader
	wq          heap
	ifaceMethod map[ // work queue, using min-heap for better locality
	methodsig]bool
	genericIfaceMethod map[ // methods called from reached interface call sites
	string]bool
	markableMethods []methodref// names of methods called from reached generic interface call sites

	reflectSeen   bool
	dynlink       bool
	methodsigstmp []methodsig// methods of reached types
	// whether we have seen a reflect method call

	pkginits []loader.// scratch buffer for decoding method signatures
	Sym
	mapinitnoop loader.Sym
}

type methodsig struct {
	name string
	typ  loader.Sym
} // methodsig is a typed method signature (name + type).
// type descriptor symbol of the function

type methodref struct {
	m   methodsig
	src loader.Sym
	r   int
} // methodref holds the relocations from a receiver type symbol to its
// method. There are three relocations, one for each of the fields in
// the reflect.method struct: mtyp, ifn, and tfn.
// the index of R_METHODOFF relocations

type dwctxt struct {
	linkctxt *Link
	ldr      *loader.Loader
	arch     *sys.Arch
	tmap     map[ // dwctxt is a wrapper intended to satisfy the method set of
	// dwarf.Context, so that functions like dwarf.PutAttrs will work with
	// DIEs that use loader.Sym as opposed to *sym.Symbol. It is also
	// being used as a place to store tables/maps that are useful as part
	// of type conversion (this is just a convenience; it would be easy to
	// split these things out into another type if need be).
	// This maps type name string (e.g. "uintptr") to loader symbol for
	// the DWARF DIE for that type (e.g. "go:info.type.uintptr")
	string]loader.Sym
	rtmap map[ // This maps loader symbol for the DWARF DIE symbol generated for
	// a type (e.g. "go:info.uintptr") to the type symbol itself
	// ("type:uintptr").
	// FIXME: try converting this map (and the next one) to a single
	// array indexed by loader.Sym -- this may perform better.
	loader.Sym]loader.Sym
	tdmap map[ // This maps Go type symbol (e.g. "type:XXX") to loader symbol for
	// the typedef DIE for that type (e.g. "go:info.XXX..def")
	loader.Sym]loader.Sym
	typeRuntimeEface loader.Sym
	typeRuntimeIface loader.Sym
	uintptrInfoSym   loader.Sym
	dwmu             *sync.Mutex
} // Cache these type symbols, so as to avoid repeatedly looking them up
// Used at various points in that parallel portion of DWARF gen to
// protect against conflicting updates to globals (such as "gdbscript")

type dwarfSecInfo struct {
	syms []loader.// dwarfSecInfo holds information about a DWARF output section,
	// specifically a section symbol and a list of symbols contained in
	// that section. On the syms list, the first symbol will always be the
	// section symbol, then any remaining symbols (if any) will be
	// sub-symbols in that section. Note that for some sections (eg:
	// .debug_abbrev), the section symbol is all there is (all content is
	// contained in it). For other sections (eg: .debug_info), the section
	// symbol is empty and all the content is in the sub-symbols. Finally
	// there are some sections (eg: .debug_ranges) where it is a mix (both
	// the section symbol and the sub-symbols have content)
	Sym
}

type dwUnitSyms struct {
	lineProlog  loader.Sym
	rangeProlog loader.Sym
	infoEpilog  loader.Sym
	linesyms    []loader.// dwUnitSyms stores input and output symbols for DWARF generation
	// for a given compilation unit.
	// Outputs for a given unit.
	Sym
	infosyms   []loader.Sym
	locsyms    []loader.Sym
	rangessyms []loader.Sym
	addrsym    loader.Sym
}

type elfNote struct {
	nNamesz uint32
	nDescsz uint32
	nType   uint32
} /*
 * Note header.  The ".note" section contains an array of notes.  Each
 * begins with this header, aligned to a word boundary.  Immediately
 * following the note header is n_namesz bytes of name, padded to the
 * next word boundary.  Then comes n_descsz bytes of descriptor, again
 * padded to a word boundary.  The values of n_namesz and n_descsz do
 * not include the padding.
 */

type ElfShdr struct {
	elf.Section64
	shnum elf.SectionIndex
} /*
 * Section header.
 */

type ELFArch struct {
	Androiddynld    string
	Linuxdynld      string
	LinuxdynldMusl  string
	Freebsddynld    string
	Netbsddynld     string
	Openbsddynld    string
	Dragonflydynld  string
	Solarisdynld    string
	Reloc1          func(*Link, *OutBuf, *loader.Loader, loader.Sym, loader.ExtReloc, int, int64) bool
	RelocSize       uint32
	SetupPLT        func(ctxt *Link, ldr *loader.Loader, plt, gotplt *loader.SymbolBuilder, dynamic loader.Sym)
	DynamicReadOnly bool
} // ELFArch includes target-specific hooks for ELF targets.
// This is initialized by the target-specific Init function
// called by the linker's main function in cmd/link/main.go.
// DynamicReadOnly can be set to true to make the .dynamic
// section read-only. By default it is writable.
// This is used by MIPS targets.

type Elfstring struct {
	s   string
	off int
}

type Elfaux struct {
	next *Elfaux
	num  int
	vers string
}

type Elflib struct {
	next *Elflib
	aux  *Elfaux
	file string
}

type unresolvedSymKey struct {
	from loader.Sym
	to   loader.Sym
} // Symbol that referenced unresolved "to"
// Unresolved symbol referenced by "from"

type ErrorReporter struct {
	loader.ErrorReporter
	unresSyms map[ // ErrorReporter is used to make error reporting thread safe.
	unresolvedSymKey]bool
	unresMutex sync.Mutex
	SymName    symNameFn
}

type fipsObj struct {
	r   io.ReaderAt
	w   io.Writer
	wf  *os.File
	h   hash.Hash
	tmp [8]byte
} // fipsObj calculates the fips object hash and optionally writes
// the hashed content to a file for debugging.

type ArchSyms struct {
	Rel         loader.Sym
	Rela        loader.Sym
	RelPLT      loader.Sym
	RelaPLT     loader.Sym
	LinkEditGOT loader.Sym
	LinkEditPLT loader.Sym
	TOC         loader.Sym
	DotTOC      []loader.// ArchSyms holds a number of architecture specific symbols used during
	// relocation.  Rather than allowing them universal access to all symbols,
	// we keep a subset for relocation application.
	Sym
	GOT               loader.Sym
	PLT               loader.Sym
	GOTPLT            loader.Sym
	Tlsg              loader.Sym
	Tlsoffset         int
	Dynamic           loader.Sym
	DynSym            loader.Sym
	DynStr            loader.Sym
	unreachableMethod loader.Sym
	mainInittasks     loader.Sym
} // for each version
// Symbol containing a list of all the inittasks that need
// to be run at startup.

type Hostobj struct {
	ld     func(*Link, *bio.Reader, string, int64, string)
	pkg    string
	pn     string
	file   string
	off    int64
	length int64
}

type Shlib struct {
	Path    string
	Hash    []byte
	Deps    []string
	File    *elf.File
	symAddr map[ // For every symbol defined in the shared library, record its address
	// in the original shared library address space.
	string]uint64
	relocTarget map[ // For relocations in the shared library, map from the address
	// (in the shared library address space) at which that
	// relocation applies to the target symbol.  We only keep
	// track of a single kind of relocation: a standard absolute
	// address relocation with no addend. These were R_ADDR
	// relocations when the shared library was built.
	uint64]string
}

type shlibReloc struct {
	addr   uint64
	target string
} // A relocation that applies to part of the shared library.
// Target symbol name.

type cgodata struct {
	file       string
	pkg        string
	directives [][]string
}

type MachoHdr struct {
	cpu    uint32
	subcpu uint32
}

type MachoSect struct {
	name    string
	segname string
	addr    uint64
	size    uint64
	off     uint32
	align   uint32
	reloc   uint32
	nreloc  uint32
	flag    uint32
	res1    uint32
	res2    uint32
}

type MachoSeg struct {
	name       string
	vsize      uint64
	vaddr      uint64
	fileoffset uint64
	filesize   uint64
	prot1      uint32
	prot2      uint32
	nsect      uint32
	msect      uint32
	sect       []MachoSect
	flag       uint32
}

type MachoPlatformLoad struct {
	platform MachoPlatform
	cmd      MachoLoad
} // MachoPlatformLoad represents a LC_VERSION_MIN_* or
// LC_BUILD_VERSION load command.
// One of PLATFORM_* constants.

type MachoLoad struct {
	type_ uint32
	data  []uint32
}

type machoRebaseRecord struct {
	sym loader.Sym
	off int64
} // A rebase entry tells the dynamic linker the data at sym+off needs to be
// relocated when the in-memory image moves. (This is somewhat like, say,
// ELF R_X86_64_RELATIVE).
// For now, the only kind of entry we support is that the data is an absolute
// address. That seems all we need.
// In the binary it uses a compact stateful bytecode encoding. So we record
// entries as we go and build the table at the end.

type machoBindRecord struct {
	off  int64
	targ loader.Sym
} // A bind entry tells the dynamic linker the data at GOT+off should be bound
// to the address of the target symbol, which is a dynamic import.
// For now, the only kind of entry we support is that the data is an absolute
// address, and the source symbol is always the GOT. That seems all we need.
// In the binary it uses a compact stateful bytecode encoding. So we record
// entries as we go and build the table at the end.

type dyldInfoCmd struct {
	Cmd                      macho.LoadCmd
	Len                      uint32
	RebaseOff, RebaseLen     uint32
	BindOff, BindLen         uint32
	WeakBindOff, WeakBindLen uint32
	LazyBindOff, LazyBindLen uint32
	ExportOff, ExportLen     uint32
}

type linkEditDataCmd struct {
	Cmd              macho.LoadCmd
	Len              uint32
	DataOff, DataLen uint32
}

type encryptionInfoCmd struct {
	Cmd                macho.LoadCmd
	Len                uint32
	CryptOff, CryptLen uint32
	CryptId            uint32
}

type uuidCmd struct {
	Cmd  macho.LoadCmd
	Len  uint32
	Uuid [16]byte
}

type Rpath struct {
	set bool
	val string
}

type OutBuf struct {
	arch *sys.Arch
	off  int64
	buf  []byte// OutBuf is a buffered file writer.
	//
	// It is similar to the Writer in cmd/internal/bio with a few small differences.
	//
	// First, it tracks the output architecture and uses it to provide
	// endian helpers.
	//
	// Second, it provides a very cheap offset counter that doesn't require
	// any system calls to read the value.
	//
	// Third, it also mmaps the output file (if available). The intended usage is:
	//   - Mmap the output file
	//   - Write the content
	//   - possibly apply any edits in the output buffer
	//   - possibly write more content to the file. These writes take place in a heap
	//     backed buffer that will get synced to disk.
	//   - Munmap the output file
	//
	// And finally, it provides a mechanism by which you can multithread the
	// writing of output files. This mechanism is accomplished by copying a OutBuf,
	// and using it in the thread/goroutine.
	//
	// Parallel OutBuf is intended to be used like:
	//
	//	func write(out *OutBuf) {
	//	  var wg sync.WaitGroup
	//	  for i := 0; i < 10; i++ {
	//	    wg.Add(1)
	//	    view, err := out.View(start[i])
	//	    if err != nil {
	//	       // handle output
	//	       continue
	//	    }
	//	    go func(out *OutBuf, i int) {
	//	      // do output
	//	      wg.Done()
	//	    }(view, i)
	//	  }
	//	  wg.Wait()
	//	}

	heap []byte// backing store of mmap'd output file

	name   string
	f      *os.File
	encbuf [8]byte
	isView bool
} // backing store for non-mmapped data
// true if created from View()

type pclntab struct {
	firstFunc, lastFunc loader.Sym
	size                int64
	carrier             loader.Sym
	pclntab             loader.Sym
	pcheader            loader.Sym
	funcnametab         loader.Sym
	findfunctab         loader.Sym
	cutab               loader.Sym
	filetab             loader.Sym
	pctab               loader.Sym
	nfunc               int32
	nfiles              uint32
} // pclntab holds the state needed for pclntab generation.
// The number of filenames in runtime.filetab.

type IMAGE_IMPORT_DESCRIPTOR struct {
	OriginalFirstThunk uint32
	TimeDateStamp      uint32
	ForwarderChain     uint32
	Name               uint32
	FirstThunk         uint32
}

type IMAGE_EXPORT_DIRECTORY struct {
	Characteristics       uint32
	TimeDateStamp         uint32
	MajorVersion          uint16
	MinorVersion          uint16
	Name                  uint32
	Base                  uint32
	NumberOfFunctions     uint32
	NumberOfNames         uint32
	AddressOfFunctions    uint32
	AddressOfNames        uint32
	AddressOfNameOrdinals uint32
}

type Imp struct {
	s       loader.Sym
	off     uint64
	next    *Imp
	argsize int
}

type Dll struct {
	name     string
	nameoff  uint64
	thunkoff uint64
	ms       *Imp
	next     *Dll
}

type peStringTable struct {
	strings []string// peStringTable is a COFF string table.

	stringsLen int
}

type peSection struct {
	name                 string
	shortName            string
	index                int
	virtualSize          uint32
	virtualAddress       uint32
	sizeOfRawData        uint32
	pointerToRawData     uint32
	pointerToRelocations uint32
	numberOfRelocations  uint16
	characteristics      uint32
} // peSection represents section from COFF section table.
// one-based index into the Section Table

type peBaseRelocEntry struct{ typeOff uint16 } // peBaseRelocEntry represents a single relocation entry.

type peBaseRelocBlock struct {
	entries []peBaseRelocEntry// peBaseRelocBlock represents a Base Relocation Block. A block
	// is a collection of relocation entries in a page, where each
	// entry describes a single relocation.
	// The block page RVA (Relative Virtual Address) is the index
	// into peBaseRelocTable.blocks.
}

type peBaseRelocTable struct {
	blocks map[ // A PE base relocation table is a list of blocks, where each block
	// contains relocation information for a single page. The blocks
	// must be emitted in order of page virtual address.
	// See https://docs.microsoft.com/en-us/windows/desktop/debug/pe-format#the-reloc-section-image-only
	uint32]peBaseRelocBlock
	pages pePages
} // pePages is a list of keys into blocks map.
// It is stored separately for ease of sorting.

type stackCheck struct {
	ctxt      *Link
	ldr       *loader.Loader
	morestack loader.Sym
	callSize  int
	height    map[ // The number of bytes added by a CALL
	// height records the maximum number of bytes a function and
	// its callees can add to the stack without a split check.
	loader.Sym]int16
	graph map[ // graph records the out-edges from each symbol. This is only
	// populated on a second pass if the first pass reveals an
	// over-limit function.
	loader.Sym][]stackCheckEdge
}

type stackCheckEdge struct {
	growth int
	target loader.Sym
} // Stack growth in bytes at call to target
// 0 for stack growth without a call

type stackCheckChain struct {
	stackCheckEdge
	printed bool
}

type Target struct {
	Arch          *sys.Arch
	HeadType      objabi.HeadType
	LinkMode      LinkMode
	BuildMode     BuildMode
	linkShared    bool
	canUsePlugins bool
	IsELF         bool
} // Target holds the configuration we're building for.

type typelinkSortKey struct {
	TypeStr string
	Type    loader.Sym
}

type XcoffFileHdr64 struct {
	Fmagic   uint16
	Fnscns   uint16
	Ftimedat int32
	Fsymptr  uint64
	Fopthdr  uint16
	Fflags   uint16
	Fnsyms   int32
} // File Header
// Number of entries in symbol table

type XcoffAoutHdr64 struct {
	Omagic      int16
	Ovstamp     int16
	Odebugger   uint32
	Otextstart  uint64
	Odatastart  uint64
	Otoc        uint64
	Osnentry    int16
	Osntext     int16
	Osndata     int16
	Osntoc      int16
	Osnloader   int16
	Osnbss      int16
	Oalgntext   int16
	Oalgndata   int16
	Omodtype    [2]byte
	Ocpuflag    uint8
	Ocputype    uint8
	Otextpsize  uint8
	Odatapsize  uint8
	Ostackpsize uint8
	Oflags      uint8
	Otsize      uint64
	Odsize      uint64
	Obsize      uint64
	Oentry      uint64
	Omaxstack   uint64
	Omaxdata    uint64
	Osntdata    int16
	Osntbss     int16
	Ox64flags   uint16
	Oresv3a     int16
	Oresv3      [2]int32
} // Auxiliary Header
// Reserved

type XcoffScnHdr64 struct {
	Sname    [8]byte
	Spaddr   uint64
	Svaddr   uint64
	Ssize    uint64
	Sscnptr  uint64
	Srelptr  uint64
	Slnnoptr uint64
	Snreloc  uint32
	Snlnno   uint32
	Sflags   uint32
} // Section Header
// flags

type xcoffSym interface{} // Type representing all XCOFF symbols.

type XcoffSymEnt64 struct {
	Nvalue  uint64
	Noffset uint32
	Nscnum  int16
	Ntype   uint16
	Nsclass uint8
	Nnumaux int8
} // Symbol Table Entry
// Number of auxiliary entries

type XcoffAuxFile64 struct {
	Xzeroes  uint32
	Xoffset  uint32
	X_pad1   [6]byte
	Xftype   uint8
	X_pad2   [2]byte
	Xauxtype uint8
} // File Auxiliary Entry
// Type of auxiliary entry

type XcoffAuxFcn64 struct {
	Xlnnoptr uint64
	Xfsize   uint32
	Xendndx  uint32
	Xpad     uint8
	Xauxtype uint8
} // Function Auxiliary Entry
// Type of auxiliary entry

type XcoffAuxCSect64 struct {
	Xscnlenlo uint32
	Xparmhash uint32
	Xsnhash   uint16
	Xsmtyp    uint8
	Xsmclas   uint8
	Xscnlenhi uint32
	Xpad      uint8
	Xauxtype  uint8
} // csect Auxiliary Entry.
// Type of auxiliary entry

type XcoffAuxDWARF64 struct {
	Xscnlen  uint64
	X_pad    [9]byte
	Xauxtype uint8
} // DWARF Auxiliary Entry
// Type of auxiliary entry

type XcoffLdHdr64 struct {
	Lversion int32
	Lnsyms   int32
	Lnreloc  int32
	Listlen  uint32
	Lnimpid  int32
	Lstlen   uint32
	Limpoff  uint64
	Lstoff   uint64
	Lsymoff  uint64
	Lrldoff  uint64
} // Loader Header
// Offset to start of relocation entries

type XcoffLdSym64 struct {
	Lvalue  uint64
	Loffset uint32
	Lscnum  int16
	Lsmtype int8
	Lsmclas int8
	Lifile  int32
	Lparm   uint32
} // Loader Symbol
// Parameter type-check field

type xcoffLoaderSymbol struct {
	sym    loader.Sym
	smtype int8
	smclas int8
}

type XcoffLdImportFile64 struct {
	Limpidpath string
	Limpidbase string
	Limpidmem  string
}

type XcoffLdRel64 struct {
	Lvaddr  uint64
	Lrtype  uint16
	Lrsecnm int16
	Lsymndx int32
} // Address Field
// Loader-Section symbol table index

type xcoffLoaderReloc struct {
	sym    loader.Sym
	roff   int32
	rtype  uint16
	symndx int32
} // xcoffLoaderReloc holds information about a relocation made by the loader.

type XcoffLdStr64 struct {
	size uint16
	name string
}

type xcoffStringTable struct {
	strings []string// xcoffStringTable is a XCOFF string table.

	stringsLen int
}

type xcoffSymSrcFile struct {
	name         string
	file         *XcoffSymEnt64
	csectAux     *XcoffAuxCSect64
	csectSymNb   uint64
	csectVAStart int64
	csectVAEnd   int64
} // type records C_FILE information needed for genasmsym in XCOFF.
// Symbol number for the current .csect

type ElfSect struct {
	name        string
	nameoff     uint32
	type_       elf.SectionType
	flags       elf.SectionFlag
	addr        uint64
	off         uint64
	size        uint64
	link        uint32
	info        uint32
	align       uint64
	entsize     uint64
	base        []byte
	readOnlyMem bool
	sym         loader.Sym
} // Is this section in readonly memory?

type ElfObj struct {
	f      *bio.Reader
	base   int64
	length int64
	is64   int
	name   string
	e      binary.ByteOrder
	sect   []ElfSect// offset in f where ELF begins
	// length of ELF

	nsect     uint
	nsymtab   int
	symtab    *ElfSect
	symstr    *ElfSect
	type_     uint32
	machine   uint32
	version   uint32
	entry     uint64
	phoff     uint64
	shoff     uint64
	flags     uint32
	ehsize    uint32
	phentsize uint32
	phnum     uint32
	shentsize uint32
	shnum     uint32
	shstrndx  uint32
}

type ElfSym struct {
	name  string
	value uint64
	size  uint64
	bind  elf.SymBind
	type_ elf.SymType
	other uint8
	shndx elf.SectionIndex
	sym   loader.Sym
}

type elfAttribute struct {
	tag  uint64
	sval string
	ival uint64
}

type elfAttributeList struct {
	data []byte
	err  error
}

type Relocs struct {
	rs []goobj.// Relocs encapsulates the set of relocations on a given symbol; an
	// instance of this type is returned by the Loader Relocs() method.
	Reloc
	li uint32
	r  *oReader
	l  *Loader
} // local index of symbol whose relocs we're examining
// loader

type ExtReloc struct {
	Xsym Sym
	Xadd int64
	Type objabi.RelocType
	Size uint8
} // ExtReloc contains the payload for an external relocation.

type oReader struct {
	*goobj.Reader
	unit      *sym.CompilationUnit
	version   int
	pkgprefix string
	syms      []Sym// oReader is a wrapper type of obj.Reader, along with some
	// extra information.
	// version of static symbol

	pkg []uint32// Sym's global index, indexed by local index

	ndef         int
	nhashed64def int
	nhasheddef   int
	objidx       uint32
} // indices of referenced package by PkgIdx (index into loader.objs array)
// index of this reader in the objs slice

type objSym struct {
	objidx uint32
	s      uint32
} // objSym represents a symbol in an object file. It is a tuple of
// the object and the symbol's local index.
// For external symbols, objidx is the index of l.extReader (extObj),
// s is its index into the payload array.
// {0, 0} represents the nil symbol.
// local index

type nameVer struct {
	name string
	v    int
}

type symAndSize struct {
	sym  Sym
	size uint32
}

type Loader struct {
	objs []*// A Loader loads new object files and resolves indexed symbol references.
	//
	// Notes on the layout of global symbol index space:
	//
	//   - Go object files are read before host object files; each Go object
	//     read adds its defined package symbols to the global index space.
	//     Nonpackage symbols are not yet added.
	//
	//   - In loader.LoadNonpkgSyms, add non-package defined symbols and
	//     references in all object files to the global index space.
	//
	//   - Host object file loading happens; the host object loader does a
	//     name/version lookup for each symbol it finds; this can wind up
	//     extending the external symbol index space range. The host object
	//     loader stores symbol payloads in loader.payloads using SymbolBuilder.
	//
	//   - Each symbol gets a unique global index. For duplicated and
	//     overwriting/overwritten symbols, the second (or later) appearance
	//     of the symbol gets the same global index as the first appearance.
	oReader
	extStart    Sym
	builtinSyms []Sym// from this index on, the symbols are externally defined

	objSyms []objSym// global index of builtin symbols

	symsByName [2]map[ // global index mapping to local index
	string]Sym
	extStaticSyms map[ // map symbol name to index, two maps are for ABI0 and ABIInternal
	nameVer]Sym
	extReader    *oReader
	payloadBatch []extSymPayload// externally defined static symbols, keyed by name
	// a dummy oReader, for external symbols

	payloads []*extSymPayload
	values   []int64// contents of linker-materialized external syms

	sects []*// symbol values, indexed by global sym index
	sym.Section
	symSects []uint16// sections

	align []uint8// symbol's section, index to sects array

	deferReturnTramp map[ // symbol 2^N alignment, indexed by global index
	Sym]bool
	objByPkg map[ // whether the symbol is a trampoline of a deferreturn call
	string]uint32
	anonVersion          int
	attrReachable        Bitmap
	attrOnList           Bitmap
	attrLocal            Bitmap
	attrNotInSymbolTable Bitmap
	attrUsedInIface      Bitmap
	attrSpecial          Bitmap
	attrVisibilityHidden Bitmap
	attrDuplicateOK      Bitmap
	attrShared           Bitmap
	attrExternal         Bitmap
	generatedSyms        Bitmap
	attrReadOnly         map[ // map package path to the index of its Go object reader
	// symbols that generate their content, indexed by ext sym idx
	Sym]bool
	attrCgoExportDynamic map[ // readonly data for this sym
	Sym]struct{}
	attrCgoExportStatic map[ // "cgo_export_dynamic" symbols
	Sym]struct{}
	outer []Sym// "cgo_export_static" symbols
	// Outer and Sub relations for symbols.

	sub map[ // indexed by global index
	Sym]Sym
	dynimplib  map[Sym]string
	dynimpvers map[ // stores Dynimplib symbol attribute
	Sym]string
	localentry map[ // stores Dynimpvers symbol attribute
	Sym]uint8
	extname map[ // stores Localentry symbol attribute
	Sym]string
	elfType map[ // stores Extname symbol attribute
	Sym]elf.SymType
	elfSym map[ // stores elf type symbol property
	Sym]int32
	localElfSym map[ // stores elf sym symbol property
	Sym]int32
	symPkg map[ // stores "local" elf sym symbol property
	Sym]string
	plt map[ // stores package for symbol, or library for shlib-derived syms
	Sym]int32
	got map[ // stores dynimport for pe objects
	Sym]int32
	dynid map[ // stores got for pe objects
	Sym]int32
	relocVariant map[ // stores Dynid for symbol
	relocId]sym.RelocVariant
	Reachparent []Sym// stores variant relocs
	// Used to implement field tracking; created during deadcode if
	// field tracking is enabled. Reachparent[K] contains the index of
	// the symbol that triggered the marking of symbol K as live.

	CgoExports map[ // CgoExports records cgo-exported symbols by SymName.
	string]Sym
	WasmExports []Sym
	sizeFixups  []symAndSize// sizeFixups records symbols that we need to fix up the size
	// after loading. It is very rarely needed, only for a DATA symbol
	// and a BSS symbol with the same name, and the BSS symbol has
	// larger size.

	flags         uint32
	strictDupMsgs int
	errorReporter *ErrorReporter
	npkgsyms      int
	nhashedsyms   int
} // number of strict-dup warning/errors, when FlagStrictDups is enabled
// number of hashed symbols, for accounting

type extSymPayload struct {
	name   string
	size   int64
	ver    int
	kind   sym.SymKind
	objidx uint32
	relocs []goobj.// extSymPayload holds the payload (data + relocations) for linker-synthesized
	// external symbols (note that symbol value is stored in a separate slice).
	// index of original object if sym made by cloneToExternal
	Reloc
	data []byte
	auxs []goobj.Aux
}

type symWithVal struct {
	s Sym
	v int64
}

type loadState struct {
	l            *Loader
	hashed64Syms map[ // Holds the loader along with temporary states for loading symbols.
	uint64]symAndSize
	hashedSyms map[ // short hashed (content-addressable) symbols, keyed by content hash
	goobj.HashType]symAndSize
	linknameVarRefs []linknameVarRef// hashed (content-addressable) symbols, keyed by content hash

} // linknamed var refererces

type linknameVarRef struct {
	pkg  string
	name string
	sym  Sym
} // package of reference (not definition)

type relocId struct {
	sym  Sym
	ridx int
} // relocId is essentially a <S,R> tuple identifying the Rth
// relocation of symbol S.

type SymbolBuilder struct {
	*extSymPayload
	symIdx Sym
	l      *Loader
} // SymbolBuilder is a helper designed to help with the construction
// of new symbol contents.
// loader

type ldMachoObj struct {
	f          *bio.Reader
	base       int64
	length     int64
	is64       bool
	name       string
	e          binary.ByteOrder
	cputype    uint
	subcputype uint
	filetype   uint32
	flags      uint32
	cmd        []ldMachoCmd// off in f where Mach-O begins
	// length of Mach-O

	ncmd uint
}

type ldMachoCmd struct {
	type_ int
	off   uint32
	size  uint32
	seg   ldMachoSeg
	sym   ldMachoSymtab
	dsym  ldMachoDysymtab
}

type ldMachoSeg struct {
	name     string
	vmaddr   uint64
	vmsize   uint64
	fileoff  uint32
	filesz   uint32
	maxprot  uint32
	initprot uint32
	nsect    uint32
	flags    uint32
	sect     []ldMachoSect
}

type ldMachoSect struct {
	name    string
	segname string
	addr    uint64
	size    uint64
	off     uint32
	align   uint32
	reloff  uint32
	nreloc  uint32
	flags   uint32
	res1    uint32
	res2    uint32
	sym     loader.Sym
	rel     []ldMachoRel
}

type ldMachoRel struct {
	addr      uint32
	symnum    uint32
	pcrel     uint8
	length    uint8
	extrn     uint8
	type_     uint8
	scattered uint8
	value     uint32
}

type ldMachoSymtab struct {
	symoff  uint32
	nsym    uint32
	stroff  uint32
	strsize uint32
	str     []byte
	sym     []ldMachoSym
}

type ldMachoSym struct {
	name    string
	type_   uint8
	sectnum uint8
	desc    uint16
	kind    int8
	value   uint64
	sym     loader.Sym
}

type ldMachoDysymtab struct {
	ilocalsym      uint32
	nlocalsym      uint32
	iextdefsym     uint32
	nextdefsym     uint32
	iundefsym      uint32
	nundefsym      uint32
	tocoff         uint32
	ntoc           uint32
	modtaboff      uint32
	nmodtab        uint32
	extrefsymoff   uint32
	nextrefsyms    uint32
	indirectsymoff uint32
	nindirectsyms  uint32
	extreloff      uint32
	nextrel        uint32
	locreloff      uint32
	nlocrel        uint32
	indir          []uint32
}

type peImportSymsState struct {
	secSyms []loader.// peImportSymsState tracks the set of DLL import symbols we've seen
	// while reading host objects. We create a singleton instance of this
	// type, which will persist across multiple host objects.
	// Text and non-text sections read in by the host object loader.
	Sym
	l    *loader.Loader
	arch *sys.Arch
} // Loader and arch, for use in postprocessing.

type peLoaderState struct {
	l        *loader.Loader
	arch     *sys.Arch
	f        *pe.File
	pn       string
	sectsyms map[ // peLoaderState holds various bits of useful state information needed
	// while loading a single PE object file.
	*pe.Section]loader.Sym
	comdats  map[uint16]int64
	sectdata map[ // key is section index, val is size
	*pe.Section][]byte
	localSymVersion int
}

type Symbols struct {
	Textp []loader.// Symbols contains the symbols that can be loaded from a PE file.
	Sym
	Resources []loader.// text symbols
	Sym
	PData loader.Sym
	XData loader.Sym
} // .rsrc section or set of .rsrc$xx sections

type ldSection struct {
	xcoff.Section
	sym loader.Sym
} // ldSection is an XCOFF section with its symbols.

type CompilationUnit struct {
	Lib       *Library
	PclnIndex int
	PCs       []dwarf.// A CompilationUnit represents a set of source files that are compiled
	// together. Since all Go sources in a Go package are compiled together,
	// there's one CompilationUnit per package that represents all Go sources in
	// that package, plus one for each assembly file.
	//
	// Equivalently, there's one CompilationUnit per object file in each Library
	// loaded by the linker.
	//
	// These are used for both DWARF and pclntab generation.
	// Index of this CU in pclntab
	Range
	DWInfo    *dwarf.DWDie
	FileTable []string// PC ranges, relative to Textp[0]
	// CU root DIE

	Consts   LoaderSym
	FuncDIEs []LoaderSym// The file table used in this compilation unit.
	// Package constants DIEs

	VarDIEs []LoaderSym// Function DIE subtrees

	AbsFnDIEs []LoaderSym// Global variable DIEs

	RangeSyms []LoaderSym// Abstract function DIE subtrees

	Textp []LoaderSym// Symbols for debug_range

	Addrs map[ // Text symbols in this CU
	LoaderSym]uint32
} // slot in .debug_addr for fn sym (DWARF5)

type Library struct {
	Objref      string
	Srcref      string
	File        string
	Pkg         string
	Shlib       string
	Fingerprint goobj.FingerprintType
	Autolib     []goobj.ImportedPkg
	Imports     []*Library
	Main        bool
	Units       []*CompilationUnit
	Textp       []LoaderSym
	DupTextSyms []LoaderSym// text syms defined in this library

} // dupok text syms defined in this library

type Segment struct {
	Rwx      uint8
	Vaddr    uint64
	Length   uint64
	Fileoff  uint64
	Filelen  uint64
	Sections []*// permission as usual unix bits (5 = r-x etc)
	// length on disk
	Section
}

type Section struct {
	Rwx        uint8
	Extnum     int16
	Align      int32
	Name       string
	Vaddr      uint64
	Length     uint64
	Seg        *Segment
	Elfsect    interface{}
	Reloff     uint64
	Rellen     uint64
	Relcount   uint32
	Sym        LoaderSym
	Index      uint16
	Compressed bool
} // an *ld.ElfShdr
// each section has a unique index, used internally

type wasmFunc struct {
	Module string
	Name   string
	Type   uint32
	Code   []byte
}

type wasmFuncType struct {
	Params  []byte
	Results []byte
}

type wasmDataSect struct {
	sect *sym.Section
	data []byte
}

type nameWriter interface {
	io.ByteWriter
	io.Writer
}

type schedt struct{}

type FileLike interface {
	Name() string
	Stat() (fs.FileInfo, error)
	Read([]byte) (// FileLike abstracts the few methods we need, so we can test without needing real files.
	int, error)
	Close() error
}

type fetcher struct{}

type objTool struct {
	mu          sync.Mutex
	disasmCache map[ // objTool implements driver.ObjTool using Go libraries
	// (instead of invoking GNU binutils).
	string]*disasm.Disasm
}

type file struct {
	name   string
	offset uint64
	sym    []objfile.// file implements driver.ObjFile using Go libraries
	// (instead of invoking GNU binutils).
	// A file represents a single executable being analyzed.
	Sym
	file       *objfile.File
	pcln       objfile.Liner
	triedDwarf bool
	dwarf      *dwarf.Data
}

type readlineUI struct{ term *term.Terminal } // readlineUI implements driver.UI interface using the
// golang.org/x/term package.
// The upstream pprof command implements the same functionality
// using the github.com/chzyer/readline package.

type countWriter struct {
	n int64
	w io.Writer
}

type generator interface {
	Sync()
	StackSample(ctx *traceContext, ev *trace.Event)
	GlobalRange(ctx *traceContext, ev *trace.Event)
	GlobalMetric(ctx *traceContext, ev *trace.Event)
	GoroutineLabel(ctx *traceContext, ev *trace.Event)
	GoroutineRange(ctx *traceContext, ev *trace.Event)
	GoroutineTransition(ctx *traceContext, ev *trace.Event)
	ProcRange(ctx *traceContext, ev *trace.Event)
	ProcTransition(ctx *traceContext, ev *trace.Event)
	Log(ctx *traceContext, ev *trace.Event)
	Finish(ctx *traceContext)
} // generator is an interface for generating a JSON trace for the trace viewer
// from a trace. Each method in this interface is a handler for a kind of event
// that is interesting to render in the UI via the JSON trace.
// Finish indicates the end of the trace and finalizes generation.

type stackSampleGenerator[R resource] struct{ getResource func(*trace.Event) R } // stackSampleGenerator implements a generic handler for stack sample events.
// The provided resource is the resource the stack sample should count against.

type globalRangeGenerator struct {
	ranges map[ // globalRangeGenerator implements a generic handler for EventRange* events that pertain
	// to trace.ResourceNone (the global scope).
	string]activeRange
	seenSync int
}

type globalMetricGenerator struct{} // globalMetricGenerator implements a generic handler for Metric events.

type procRangeGenerator struct {
	ranges map[ // procRangeGenerator implements a generic handler for EventRange* events whose Scope.Kind is
	// ResourceProc.
	trace.Range]activeRange
	seenSync int
}

type activeRange struct {
	time  trace.Time
	stack trace.Stack
} // activeRange represents an active EventRange* range.

type completedRange struct {
	name       string
	startTime  trace.Time
	endTime    trace.Time
	startStack trace.Stack
	endStack   trace.Stack
	arg        any
} // completedRange represents a completed EventRange* range.

type logEventGenerator[R resource] struct{ getResource func(*trace.Event) R }

type goroutineGenerator struct {
	globalRangeGenerator
	globalMetricGenerator
	stackSampleGenerator[trace.GoID]
	logEventGenerator[trace.GoID]
	gStates map[trace.GoID]*gState[trace.GoID]
	focus   trace.GoID
	filter  map[trace.GoID]struct{}
}

type resource interface {
	trace.GoID | trace.ProcID | trace.ThreadID
} // resource is a generic constraint interface for resource IDs.

type gState[R resource] struct {
	baseName      string
	named         bool
	label         string
	isSystemG     bool
	executing     R
	lastStopStack trace.Stack
	activeRanges  map[ // gState represents the trace viewer state of a goroutine in a trace.
	//
	// The type parameter on this type is the resource which is used to construct
	// a timeline of events. e.g. R=ProcID for a proc-oriented view, R=GoID for
	// a goroutine-oriented view, etc.
	// activeRanges is the set of all active ranges on the goroutine.
	string]activeRange
	completedRanges []completedRange// completedRanges is a list of ranges that completed since before the
	// goroutine stopped executing. These are flushed on every stop or block.

	startRunningTime trace.Time
	syscall          struct {
		time   trace.Time
		stack  trace.Stack
		active bool
	}
	startBlockReason string
	startCause       struct {
		time     trace.Time
		name     string
		resource uint64
		stack    trace.Stack
	}
} // startRunningTime is the most recent event that caused a goroutine to
// transition to GoRunning.
// startCause is the event that allowed this goroutine to start running.
// It's used to generate flow events. This is typically something like
// an unblock event or a goroutine creation event.
//
// startCause.resource is the resource on which startCause happened, but is
// listed separately because the cause may have happened on a resource that
// isn't R (or perhaps on some abstract nebulous resource, like trace.NetpollP).

type genOpts struct {
	mode           traceviewer.Mode
	startTime      time.Duration
	endTime        time.Duration
	focusGoroutine trace.GoID
	goroutines     map[ // Used if mode != 0.
	trace.GoID]struct{}
	tasks []*// Goroutines to be displayed for goroutine-oriented or task-oriented view. goroutines[0] is the main goroutine.
	trace.UserTaskSummary
}

type parsedTrace struct {
	events      []trace.Event
	summary     *trace.Summary
	size, valid int64
	err         error
}

type countingReader struct {
	r         io.Reader
	bytesRead atomic.Int64
}

type interval struct{ start, end trace.Time } // interval represents a time interval in the trace.

type stackMap struct {
	stacks map[ // stackMap is a map of trace.Stack to some value V.
	// stacks contains the full list of stacks in the set, however
	// it is insufficient for deduplication because trace.Stack
	// equality is only optimistic. If two trace.Stacks are equal,
	// then they are guaranteed to be equal in content. If they are
	// not equal, then they might still be equal in content.
	trace.Stack]*traceviewer.ProfileRecord
	pcs map[ // pcs is the source-of-truth for deduplication. It is a map of
	// the actual PCs in the stack to a trace.Stack.
	[pprofMaxStack]uint64]trace.Stack
}

type procGenerator struct {
	globalRangeGenerator
	globalMetricGenerator
	procRangeGenerator
	stackSampleGenerator[trace.ProcID]
	logEventGenerator[trace.ProcID]
	gStates   map[trace.GoID]*gState[trace.ProcID]
	inSyscall map[trace.ProcID]*gState[trace.ProcID]
	maxProc   trace.ProcID
}

type regionFingerprint struct {
	Frame trace.StackFrame
	Type  string
} // regionFingerprint is a way to categorize regions that goes just one step beyond the region's Type
// by including the top stack frame.

type regionStats struct {
	regionFingerprint
	Histogram traceviewer.TimeHistogram
}

type regionFilter struct {
	name   string
	params url.Values
	cond   []func// regionFilter represents a region filter specified by a user of cmd/trace.
	(*parsedTrace, *trace.UserRegionSummary) bool
}

type taskStats struct {
	Type      string
	Count     int
	Histogram traceviewer.TimeHistogram
} // Complete + incomplete tasks
// Complete tasks only

type taskFilter struct {
	name string
	cond []func// taskFilter represents a task filter specified by a user of cmd/trace.
	(*parsedTrace, *trace.UserTaskSummary) bool
}

type threadGenerator struct {
	globalRangeGenerator
	globalMetricGenerator
	stackSampleGenerator[trace.ThreadID]
	logEventGenerator[trace.ThreadID]
	gStates map[trace.GoID]*gState[trace.ThreadID]
	threads map[trace.ThreadID]struct{}
}

type Options struct {
	Writer        Writer
	Flagset       FlagSet
	Fetch         Fetcher
	Sym           Symbolizer
	Obj           ObjTool
	UI            UI
	HTTPServer    func(*HTTPServerArgs) error
	HTTPTransport http.RoundTripper
} // Options groups all the optional plugins into pprof.

type FlagSet interface {
	Bool(name string, def bool, usage string) *bool
	Int(name string, def int, usage string) *int
	Float64(name string, def float64, usage string) *float64
	String(name string, def string, usage string) *string
	StringList(name string, def string, usage string) *[]*// A FlagSet creates and parses command-line flags.
	// It is similar to the standard flag.FlagSet.
	// StringList is similar to String but allows multiple values for a
	// single flag
	string
	ExtraUsage() string
	AddExtraUsage(eu string)
	Parse(usage func()) []string// ExtraUsage returns any additional text that should be printed after the
	// standard usage message. The extra usage message returned includes all text
	// added with AddExtraUsage().
	// The typical use of ExtraUsage is to show any custom flags defined by the
	// specific pprof plugins being used.
	// Parse initializes the flags with their values for this run
	// and returns the non-flag command line arguments.
	// If an unknown flag is encountered or there are no arguments,
	// Parse should call usage and return nil.

}

type Fetcher interface {
	Fetch(src string, duration, timeout time.Duration) (*profile.Profile, string, error)
} // A Fetcher reads and returns the profile named by src, using
// the specified duration and timeout. It returns the fetched
// profile and a string indicating a URL from where the profile
// was fetched, which may be different than src.

type Symbolizer interface {
	Symbolize(mode string, srcs MappingSources, prof *profile.Profile) error
} // A Symbolizer introduces symbol information into a profile.

type ObjTool interface {
	Open(file string, start, limit, offset uint64, relocationSymbol string) (ObjFile, error)
	Disasm(file string, start, end uint64, intelSyntax bool) ([]Inst,// An ObjTool inspects shared libraries and executable files.
	// Disasm disassembles the named object file, starting at
	// the start address and stopping at (before) the end address.
	error)
}

type Inst struct {
	Addr     uint64
	Text     string
	Function string
	File     string
	Line     int
} // An Inst is a single instruction in an assembly listing.
// source line

type ObjFile interface {
	Name() string
	ObjAddr(addr uint64) (uint64, error)
	BuildID() string
	SourceLine(addr uint64) ([]Frame,// An ObjFile is a single object file: a shared library or executable.
	// SourceLine reports the source line information for a given
	// address in the file. Due to inlining, the source line information
	// is in general a list of positions representing a call stack,
	// with the leaf function first.
	error)
	Symbols(r *regexp.Regexp, addr uint64) ([]*// Symbols returns a list of symbols in the object file.
	// If r is not nil, Symbols restricts the list to symbols
	// with names matching the regular expression.
	// If addr is not zero, Symbols restricts the list to symbols
	// containing that address.
	Sym, error)
	Close() error
} // Close closes the file, releasing associated resources.

type Frame struct {
	Func      string
	File      string
	Line      int
	Column    int
	StartLine int
} // A Frame describes a single line in a source file.
// start line of function (if available)

type UI interface {
	ReadLine(prompt string) (string, error)
	Print(...interface{})
	PrintErr(...interface{})
	IsTerminal() bool
	WantBrowser() bool
	SetAutoComplete(complete func(string) string)
} // A UI manages user interactions.
// SetAutoComplete instructs the UI to call complete(cmd) to obtain
// the auto-completion of cmd, if the UI supports auto-completion at all.

type internalObjTool struct{ ObjTool } // internalObjTool is a wrapper to map from the pprof external
// interface to the internal interface.

type internalObjFile struct{ ObjFile }

type internalSymbolizer struct{ Symbolizer } // internalSymbolizer is a wrapper to map from the pprof external
// interface to the internal interface.

type addr2Liner struct {
	mu   sync.Mutex
	rw   lineReaderWriter
	base uint64
	nm   *addr2LinerNM
} // addr2Liner is a connection to an addr2line command for obtaining
// address and line number information from a binary.
// nm holds an addr2Liner using nm tool. Certain versions of addr2line
// produce incomplete names due to
// https://sourceware.org/bugzilla/show_bug.cgi?id=17541. As a workaround,
// the names from nm are used when they look more complete. See addrInfo()
// code below for the exact heuristic.

type lineReaderWriter interface {
	write(string) error
	readLine() (string, error)
	close()
} // lineReaderWriter is an interface to abstract the I/O to an addr2line
// process. It writes a line of input to the job, and reads its output
// one line at a time.

type addr2LinerJob struct {
	cmd *exec.Cmd
	in  io.WriteCloser
	out *bufio.Reader
}

type llvmSymbolizer struct {
	sync.Mutex
	filename string
	rw       lineReaderWriter
	base     uint64
	isData   bool
} // llvmSymbolizer is a connection to an llvm-symbolizer command for
// obtaining address and line number information from a binary.

type llvmSymbolizerJob struct {
	cmd     *exec.Cmd
	in      io.WriteCloser
	out     *bufio.Reader
	symType string
} // llvm-symbolizer requires the symbol type, CODE or DATA, for symbolization.

type addr2LinerNM struct {
	m []symbolInfo// addr2LinerNM is a connection to an nm command for obtaining symbol
	// information from a binary.

} // Sorted list of symbol addresses from binary.

type symbolInfo struct {
	address uint64
	size    uint64
	name    string
	symType string
}

type Binutils struct {
	mu  sync.Mutex
	rep *binrep
} // A Binutils implements plugin.ObjTool by invoking the GNU binutils.

type binrep struct {
	llvmSymbolizer      string
	llvmSymbolizerFound bool
	addr2line           string
	addr2lineFound      bool
	nm                  string
	nmFound             bool
	objdump             string
	objdumpFound        bool
	isLLVMObjdump       bool
	fast                bool
} // binrep is an immutable representation for Binutils.  It is atomically
// replaced on every mutation to provide thread-safe access.
// if fast, perform symbolization using nm (symbol names only),
// instead of file-line detail from the slower addr2line.

type elfMapping struct {
	start, limit, offset uint64
	kernelOffset         *uint64
} // elfMapping stores the parameters of a runtime mapping that are needed to
// identify the ELF segment associated with a mapping.
// Offset of kernel relocation symbol. Only defined for kernel images, nil otherwise.

type fileNM struct {
	file
	addr2linernm *addr2LinerNM
} // fileNM implements the binutils.ObjFile interface, using 'nm' to map
// addresses to symbols (without file/line number information). It is
// faster than fileAddr2Line.

type fileAddr2Line struct {
	once sync.Once
	file
	addr2liner     *addr2Liner
	llvmSymbolizer *llvmSymbolizer
	isData         bool
} // fileAddr2Line implements the binutils.ObjFile interface, using
// llvm-symbolizer, if that's available, or addr2line to map addresses to
// symbols (with file/line number information). It can be slow for large
// binaries with debug information.

type config struct {
	Output              string  `json:"-"`
	CallTree            bool    `json:"call_tree,omitempty"`
	RelativePercentages bool    `json:"relative_percentages,omitempty"`
	Unit                string  `json:"unit,omitempty"`
	CompactLabels       bool    `json:"compact_labels,omitempty"`
	SourcePath          string  `json:"-"`
	TrimPath            string  `json:"-"`
	IntelSyntax         bool    `json:"intel_syntax,omitempty"`
	Mean                bool    `json:"mean,omitempty"`
	SampleIndex         string  `json:"-"`
	DivideBy            float64 `json:"-"`
	Normalize           bool    `json:"normalize,omitempty"`
	Sort                string  `json:"sort,omitempty"`
	TagRoot             string  `json:"tagroot,omitempty"`
	TagLeaf             string  `json:"tagleaf,omitempty"`
	DropNegative        bool    `json:"drop_negative,omitempty"`
	NodeCount           int     `json:"nodecount,omitempty"`
	NodeFraction        float64 `json:"nodefraction,omitempty"`
	EdgeFraction        float64 `json:"edgefraction,omitempty"`
	Trim                bool    `json:"trim,omitempty"`
	Focus               string  `json:"focus,omitempty"`
	Ignore              string  `json:"ignore,omitempty"`
	PruneFrom           string  `json:"prune_from,omitempty"`
	Hide                string  `json:"hide,omitempty"`
	Show                string  `json:"show,omitempty"`
	ShowFrom            string  `json:"show_from,omitempty"`
	TagFocus            string  `json:"tagfocus,omitempty"`
	TagIgnore           string  `json:"tagignore,omitempty"`
	TagShow             string  `json:"tagshow,omitempty"`
	TagHide             string  `json:"taghide,omitempty"`
	NoInlines           bool    `json:"noinlines,omitempty"`
	ShowColumns         bool    `json:"showcolumns,omitempty"`
	Granularity         string  `json:"granularity,omitempty"`
} // config holds settings for a single named config.
// The JSON tag name for a field is used both for JSON encoding and as
// a named variable.
// Output granularity

type configField struct {
	name     string
	urlparam string
	saved    bool
	field    reflect.StructField
	choices  []string// configField contains metadata for a single configuration field.
	// Field in config

	defaultValue string
} // Name Of variables in group
// Default value for this field.

type profileSource struct {
	addr   string
	source *source
	p      *profile.Profile
	msrc   plugin.MappingSources
	remote bool
	err    error
}

type GoFlags struct {
	UsageMsgs []string// GoFlags implements the plugin.FlagSet interface.
}

type stdUI struct{ r *bufio.Reader }

type oswriter struct{} // oswriter implements the Writer interface using a regular file.

type settings struct {
	Configs []namedConfig `json:"configs"`// settings holds pprof settings.
	// Configs holds a list of named UI configurations.

}

type namedConfig struct {
	Name string `json:"name"`
	config
} // namedConfig associates a name with a config.

type configMenuEntry struct {
	Name       string
	URL        string
	Current    bool
	UserConfig bool
} // configMenuEntry holds information for a single config menu entry.
// Is this a user-provided config?

type webInterface struct {
	prof    *profile.Profile
	copier  profileCopier
	options *plugin.Options
	help    map[ // webInterface holds the state needed for serving a browser based interface.
	string]string
	settingsFile string
}

type errorCatcher struct {
	plugin.UI
	errors []string// errorCatcher is a UI that captures errors for reporting to the browser.

}

type webArgs struct {
	Title  string
	Errors []string// webArgs contains arguments passed to templates in webhtml.go.

	Total       int64
	SampleTypes []string
	Legend      []string
	DocURL      string
	Standalone  bool
	Help        map[ // True for command-line generation of HTML
	string]string
	Nodes      []string
	HTMLBody   template.HTML
	TextBody   string
	Top        []report.TextItem
	Listing    report.WebListData
	FlameGraph template.JS
	Stacks     template.JS
	Configs    []configMenuEntry
	UnitDefs   []measurement.UnitType
}

type DotAttributes struct {
	Nodes map[ // DotAttributes contains details about the graph itself, giving
	// insight into how its elements should be rendered.
	*Node]*DotNodeAttributes
} // A map allowing each Node to have its own visualization option

type DotNodeAttributes struct {
	Shape       string
	Bold        bool
	Peripheries int
	URL         string
	Formatter   func(*NodeInfo) string
} // DotNodeAttributes contains Node specific visualization options.
// An optional formatter for the node's label

type DotConfig struct {
	Title     string
	LegendURL string
	Labels    []string// DotConfig contains attributes about how a graph should be
	// constructed and how it should look.
	// The URL to link to from the legend.

	FormatValue func(int64) string
	Total       int64
} // The labels for the DOT's legend
// The total weight of the graph, used to compute percentages

type builder struct {
	io.Writer
	attributes *DotAttributes
	config     *DotConfig
} // builder wraps an io.Writer and understands how to compose DOT formatted elements.

type NodeInfo struct {
	Name              string
	OrigName          string
	Address           uint64
	File              string
	StartLine, Lineno int
	Columnno          int
	Objfile           string
} // NodeInfo contains the attributes for a node.

type nodePair struct{ src, dest *Node }

type tags struct {
	t    []*Tag
	flat bool
}

type nodeSorter struct {
	rs   Nodes
	less func(l, r *Node) bool
} // nodeSorter is a mechanism used to allow a report to be sorted
// in different ways.

type Unit struct {
	CanonicalName string
	aliases       []string// Unit includes a list of aliases representing a specific unit and a factor
	// which one can multiple a value in the specified unit by to get the value
	// in terms of the base unit.

	Factor float64
}

type UnitType struct {
	DefaultUnit Unit
	Units       []Unit// UnitType includes a list of units that are within the same category (i.e.
	// memory or time units) and a default unit to use for this type of unit.

}

type objSymbol struct {
	sym  *plugin.Sym
	file plugin.ObjFile
} // objSym represents a symbol identified from a binary. It includes
// the SymbolInfo from the disasm package and the base that must be
// added to correspond to sample addresses

type orderSyms struct {
	v []*// orderSyms is a wrapper type to sort []*objSymbol by a supplied comparator.
	objSymbol
	less func(a, b *objSymbol) bool
}

type assemblyInstruction struct {
	address         uint64
	instruction     string
	function        string
	file            string
	line            int
	flat, cum       int64
	flatDiv, cumDiv int64
	startsBlock     bool
	inlineCalls     []callID
}

type callID struct {
	file string
	line int
}

type TextItem struct {
	Name                  string
	InlineLabel           string
	Flat, Cum             int64
	FlatFormat, CumFormat string
} // TextItem holds a single text report entry.
// Formatted values

type Report struct {
	prof        *profile.Profile
	total       int64
	options     *Options
	formatValue func(int64) string
} // Report contains the data and associated routines to extract a
// report from a profile.

type sourcePrinter struct {
	reader     *sourceReader
	synth      *synthCode
	objectTool plugin.ObjTool
	objects    map[ // sourcePrinter holds state needed for generating source+asm HTML listing.
	string]plugin.ObjFile
	sym   *regexp.Regexp
	files map[ // Opened object files
	// May be nil
	string]*sourceFile
	insts map[ // Set of files to print.
	uint64]instructionInfo
	interest map[ // Instructions of interest (keyed by address).
	// Set of function names that we are interested in (because they had
	// a sample and match sym).
	string]bool
	prettyNames map[ // Mapping from system function names to printable names.
	string]string
}

type addrInfo struct {
	loc *profile.Location
	obj plugin.ObjFile
} // addrInfo holds information for an address we are interested in.
// May be nil

type instructionInfo struct {
	objAddr   uint64
	length    int
	disasm    string
	file      string
	line      int
	flat, cum int64
} // instructionInfo holds collected information for an instruction.
// Samples to report (divisor already applied)

type sourceInst struct {
	addr  uint64
	stack []callID// sourceInst holds information for an instruction to be displayed.

} // Inlined call-stack

type sourceFunction struct {
	name       string
	begin, end int
	flat, cum  int64
} // sourceFunction contains information for a contiguous range of lines per function we
// will print.
// Line numbers (end is not included in the range)

type addressRange struct {
	begin, end uint64
	obj        plugin.ObjFile
	mapping    *profile.Mapping
	score      int64
} // addressRange is a range of addresses plus the object file that contains it.
// Used to order ranges for processing

type WebListData struct {
	Total string
	Files []WebListFile// WebListData holds the data needed to generate HTML source code listing.

}

type WebListFile struct {
	Funcs []WebListFunc// WebListFile holds the per-file information for HTML source code listing.
}

type WebListFunc struct {
	Name       string
	File       string
	Flat       string
	Cumulative string
	Percent    string
	Lines      []WebListLine// WebListFunc holds the per-function information for HTML source code listing.

}

type WebListLine struct {
	SrcLine      string
	HTMLClass    string
	Line         int
	Flat         string
	Cumulative   string
	Instructions []WebListInstruction// WebListLine holds the per-source-line information for HTML source code listing.

}

type WebListInstruction struct {
	NewBlock     bool
	Flat         string
	Cumulative   string
	Synthetic    bool
	Address      uint64
	Disasm       string
	FileLine     string
	InlinedCalls []WebListCall// WebListInstruction holds the per-instruction information for HTML source code listing.
	// Insert marker that indicates separation from previous block

}

type WebListCall struct {
	SrcLine  string
	FileBase string
	Line     int
} // WebListCall holds the per-inlined-call information for HTML source code listing.

type sourceReader struct {
	searchPath string
	trimPath   string
	files      map[ // sourceReader provides access to source code with caching of file contents.
	// files maps from path name to a list of lines.
	// files[*][0] is unused since line numbering starts at 1.
	string][]string
	errors map[ // errors collects errors encountered per file. These errors are
	// consulted before returning out of these module.
	string]error
}

type StackSet struct {
	Total  int64
	Scale  float64
	Type   string
	Unit   string
	Stacks []Stack// StackSet holds a set of stacks corresponding to a profile.
	//
	// Slices in StackSet and the types it contains are always non-nil,
	// which makes Javascript code that uses the JSON encoding less error-prone.
	// One of "B", "s", "GCU", or "" (if unknown)

	Sources []StackSource// List of stored stacks

	report *Report
} // Mapping from source index to info

type StackSource struct {
	FullName   string
	FileName   string
	UniqueName string
	Inlined    bool
	Display    []string// StackSource holds function/location info for a stack entry.
	// Alternative names to display (with decreasing lengths) to make text fit.
	// Guaranteed to be non-empty.

	Places []StackSlot// Places holds the list of stack slots where this source occurs.
	// In particular, if [a,b] is an element in Places,
	// StackSet.Stacks[a].Sources[b] points to this source.
	//
	// No stack will be referenced twice in the Places slice for a given
	// StackSource. In case of recursion, Places will contain the outer-most
	// entry in the recursive stack. E.g., if stack S has source X at positions
	// 4,6,9,10, the Places entry for X will contain [S,4].

	Self  int64
	Color int
} // Combined count of stacks where this source is the leaf.
// Color number to use for this source.
// Colors with high numbers than supported may be treated as zero.

type StackSlot struct {
	Stack int
	Pos   int
} // StackSlot identifies a particular StackSlot.
// Index in Stack.Sources

type synthCode struct {
	next uint64
	addr map[ // synthCode assigns addresses to locations without an address.
	*profile.Location]uint64
} // Synthesized address assigned to a location

type transport struct {
	cert       *string
	key        *string
	ca         *string
	caCertPool *x509.CertPool
	certs      []tls.Certificate
	initOnce   sync.Once
	initErr    error
}

type profileMerger struct {
	p             *Profile
	locationsByID locationIDMap
	functionsByID map[ // Memoization tables within a profile.
	uint64]*Function
	mappingsByID map[uint64]mapInfo
	samples      map[ // Memoization tables for profile entities.
	sampleKey]*Sample
	locations map[locationKey]*Location
	functions map[functionKey]*Function
	mappings  map[mappingKey]*Mapping
}

type mapInfo struct {
	m      *Mapping
	offset int64
}

type locationKey struct {
	addr, mappingID uint64
	lines           string
	isFolded        bool
}

type mappingKey struct {
	size, offset  uint64
	buildIDOrFile string
}

type functionKey struct {
	startLine                  int64
	name, systemName, fileName string
}

type locationIDMap struct {
	dense []*// locationIDMap is like a map[uint64]*Location, but provides efficiency for
	// ids that are densely numbered, which is often the case.
	Location
	sparse map[ // indexed by id for id < len(dense)
	uint64]*Location
} // indexed by id for id >= len(dense)

type ValueType struct {
	Type  string
	Unit  string
	typeX int64
	unitX int64
} // ValueType corresponds to Profile.ValueType
// seconds, nanoseconds, bytes, etc

type Sample struct {
	Location []*// Sample corresponds to Profile.Sample
	Location
	Value []int64
	Label map[ // Label is a per-label-key map to values for string labels.
	//
	// In general, having multiple values for the given label key is strongly
	// discouraged - see docs for the sample label field in profile.proto.  The
	// main reason this unlikely state is tracked here is to make the
	// decoding->encoding roundtrip not lossy. But we expect that the value
	// slices present in this map are always of length 1.
	string][]string
	NumLabel map[ // NumLabel is a per-label-key map to values for numeric labels. See a note
	// above on handling multiple values for a label.
	string][]int64
	NumUnit map[ // NumUnit is a per-label-key map to the unit names of corresponding numeric
	// label values. The unit info may be missing even if the label is in
	// NumLabel, see the docs in profile.proto for details. When the value is
	// slice is present and not nil, its length must be equal to the length of
	// the corresponding value slice in NumLabel.
	string][]string
	locationIDX []uint64
	labelX      []label
}

type Mapping struct {
	ID                     uint64
	Start                  uint64
	Limit                  uint64
	Offset                 uint64
	File                   string
	BuildID                string
	HasFunctions           bool
	HasFilenames           bool
	HasLineNumbers         bool
	HasInlineFrames        bool
	fileX                  int64
	buildIDX               int64
	KernelRelocationSymbol string
} // Mapping corresponds to Profile.Mapping
// Name of the kernel relocation symbol ("_text" or "_stext"), extracted from File.
// For linux kernel mappings generated by some tools, correct symbolization depends
// on knowing which of the two possible relocation symbols was used for `Start`.
// This is given to us as a suffix in `File` (e.g. "[kernel.kallsyms]_stext").
//
// Note, this public field is not persisted in the proto. For the purposes of
// copying / merging / hashing profiles, it is considered subsumed by `File`.

type Line struct {
	Function    *Function
	Line        int64
	Column      int64
	functionIDX uint64
} // Line corresponds to Profile.Line

type Function struct {
	ID          uint64
	Name        string
	SystemName  string
	Filename    string
	StartLine   int64
	nameX       int64
	systemNameX int64
	filenameX   int64
} // Function corresponds to Profile.Function

type buffer struct {
	field int
	typ   int
	u64   uint64
	data  []byte// field tag
	// proto wire type code for field

	tmp      [16]byte
	tmpLines []Line
} // temporary storage used while decoding "repeated Line".

type message interface {
	decoder() []decoder
	encode(*buffer)
}

type AST interface {
	print(*printState)
	Traverse(func(AST) bool)
	Copy(copy func(AST) AST, skip func(AST) bool) AST
	GoString() string
	goString(indent int, field string) string
} // AST is an abstract syntax tree representing a C++ declaration.
// This is sufficient for the demangler but is by no means a general C++ AST.
// This abstract syntax tree is only used for C++ symbols, not Rust symbols.
// Implement the fmt.GoStringer interface.

type printState struct {
	tparams         bool
	enclosingParams bool
	llvmStyle       bool
	max             int
	scopes          int
	buf             strings.Builder
	last            byte
	inner           []AST// The printState type holds information needed to print an AST.
	// The inner field is a list of items to print for a type
	// name.  This is used by types to implement the inside-out
	// C++ declaration syntax.

	printing []AST// The printing field is a list of items we are currently
	// printing.  This avoids endless recursion if a substitution
	// reference creates a cycle in the graph.

}

type hasPrec interface{ prec() precedence } // hasPrec matches the AST nodes that have a prec method that returns
// the node's precedence.

type Typed struct {
	Name AST
	Type AST
} // Typed is a typed name.

type Qualified struct {
	Scope     AST
	Name      AST
	LocalName bool
} // Qualified is a name in a scope.
// A full local name encoding

type Template struct {
	Name AST
	Args []AST// Template is a template with arguments.

}

type TemplateParam struct {
	Index    int
	Template *Template
} // TemplateParam is a template parameter.  The Template field is
// filled in while parsing the demangled string.  We don't normally
// see these while printing--they are replaced by the simplify
// function.

type LambdaAuto struct{ Index int } // LambdaAuto is a lambda auto parameter.

type TemplateParamQualifiedArg struct {
	Param AST
	Arg   AST
} // TemplateParamQualifiedArg is used when the mangled name includes
// both the template parameter declaration and the template argument.
// See https://github.com/itanium-cxx-abi/cxx-abi/issues/47.

type Qualifiers struct {
	Qualifiers []AST// Qualifiers is an ordered list of type qualifiers.
}

type TypeWithQualifiers struct {
	Base       AST
	Qualifiers AST
} // TypeWithQualifiers is a type with standard qualifiers.

type MethodWithQualifiers struct {
	Method       AST
	Qualifiers   AST
	RefQualifier string
} // MethodWithQualifiers is a method with qualifiers.
// "" or "&" or "&&"

type BuiltinType struct{ Name string } // BuiltinType is a builtin type, like "int".

type PointerType struct{ Base AST } // PointerType is a pointer type.

type ReferenceType struct{ Base AST } // ReferenceType is a reference type.

type RvalueReferenceType struct{ Base AST } // RvalueReferenceType is an rvalue reference type.

type ComplexType struct{ Base AST } // ComplexType is a complex type.

type ImaginaryType struct{ Base AST } // ImaginaryType is an imaginary type.

type SuffixType struct {
	Base   AST
	Suffix string
} // SuffixType is an type with an arbitrary suffix.

type TransformedType struct {
	Name string
	Base AST
} // TransformedType is a builtin type with a template argument.

type VendorQualifier struct {
	Qualifier AST
	Type      AST
} // VendorQualifier is a type qualified by a vendor-specific qualifier.

type ArrayType struct {
	Dimension AST
	Element   AST
} // ArrayType is an array type.

type FunctionType struct {
	Return AST
	Args   []AST// FunctionType is a function type.

	ForLocalName bool
} // The forLocalName field reports whether this FunctionType
// was created for a local name. With the default GNU demangling
// output we don't print the return type in that case.

type FunctionParam struct{ Index int } // FunctionParam is a parameter of a function, used for last-specified
// return type in a closure.

type PtrMem struct {
	Class  AST
	Member AST
} // PtrMem is a pointer-to-member expression.

type FixedType struct {
	Base  AST
	Accum bool
	Sat   bool
} // FixedType is a fixed numeric type of unknown size.

type BinaryFP struct{ Bits int } // BinaryFP is a binary floating-point type.

type BitIntType struct {
	Size   AST
	Signed bool
} // BitIntType is the C++23 _BitInt(N) type.

type VectorType struct {
	Dimension AST
	Base      AST
} // VectorType is a vector type.

type ElaboratedType struct {
	Kind string
	Type AST
} // ElaboratedType is an elaborated struct/union/enum type.

type Decltype struct{ Expr AST } // Decltype is the decltype operator.

type Constructor struct {
	Name AST
	Base AST
} // Constructor is a constructor.
// base class of inheriting constructor

type Destructor struct{ Name AST } // Destructor is a destructor.

type GlobalCDtor struct {
	Ctor bool
	Key  AST
} // GlobalCDtor is a global constructor or destructor.

type TaggedName struct {
	Name AST
	Tag  AST
} // TaggedName is a name with an ABI tag.

type PackExpansion struct {
	Base AST
	Pack *ArgumentPack
} // PackExpansion is a pack expansion.  The Pack field may be nil.

type ArgumentPack struct {
	Args []AST// ArgumentPack is an argument pack.
}

type SizeofPack struct{ Pack *ArgumentPack } // SizeofPack is the sizeof operator applied to an argument pack.

type SizeofArgs struct {
	Args []AST// SizeofArgs is the size of a captured template parameter pack from
	// an alias template.
}

type TemplateParamName struct {
	Prefix string
	Index  int
} // TemplateParamName is the name of a template parameter that the
// demangler introduced for a lambda that has explicit template
// parameters.  This is a prefix with an index.

type TypeTemplateParam struct{ Name AST } // TypeTemplateParam is a type template parameter that appears in a
// lambda with explicit template parameters.

type NonTypeTemplateParam struct {
	Name AST
	Type AST
} // NonTypeTemplateParam is a non-type template parameter that appears
// in a lambda with explicit template parameters.

type TemplateTemplateParam struct {
	Name   AST
	Params []AST// TemplateTemplateParam is a template template parameter that appears
	// in a lambda with explicit template parameters.

	Constraint AST
}

type ConstrainedTypeTemplateParam struct {
	Name       AST
	Constraint AST
} // ConstrainedTypeTemplateParam is a constrained template type
// parameter declaration.

type TemplateParamPack struct{ Param AST } // TemplateParamPack is a template parameter pack that appears in a
// lambda with explicit template parameters.

type Cast struct{ To AST } // Cast is a type cast.

type Nullary struct{ Op AST } // Nullary is an operator in an expression with no arguments, such as
// throw.

type Unary struct {
	Op         AST
	Expr       AST
	Suffix     bool
	SizeofType bool
} // Unary is a unary operation in an expression.
// true for sizeof (type)

type Binary struct {
	Op    AST
	Left  AST
	Right AST
} // Binary is a binary operation in an expression.

type Trinary struct {
	Op     AST
	First  AST
	Second AST
	Third  AST
} // Trinary is the ?: trinary operation in an expression.

type Fold struct {
	Left bool
	Op   AST
	Arg1 AST
	Arg2 AST
} // Fold is a C++17 fold-expression.  Arg2 is nil for a unary operator.

type Subobject struct {
	Type      AST
	SubExpr   AST
	Offset    int
	Selectors []int// Subobject is a a reference to an offset in an expression.  This is
	// used for C++20 manglings of class types used as the type of
	// non-type template arguments.
	//
	// See https://github.com/itanium-cxx-abi/cxx-abi/issues/47.

	PastEnd bool
}

type PtrMemCast struct {
	Type   AST
	Expr   AST
	Offset int
} // PtrMemCast is a conversion of an expression to a pointer-to-member
// type.  This is used for C++20 manglings of class types used as the
// type of non-type template arguments.
//
// See https://github.com/itanium-cxx-abi/cxx-abi/issues/47.

type New struct {
	Op    AST
	Place AST
	Type  AST
	Init  AST
} // New is a use of operator new in an expression.

type Literal struct {
	Type AST
	Val  string
	Neg  bool
} // Literal is a literal in an expression.

type StringLiteral struct{ Type AST } // StringLiteral is a string literal.

type LambdaExpr struct{ Type AST } // LambdaExpr is a literal that is a lambda expression.

type ExprList struct {
	Exprs []AST// ExprList is a list of expressions, typically arguments to a
	// function call in an expression.
}

type InitializerList struct {
	Type  AST
	Exprs AST
} // InitializerList is an initializer list: an optional type with a
// list of expressions.

type DefaultArg struct {
	Num int
	Arg AST
} // DefaultArg holds a default argument for a local name.

type Closure struct {
	TemplateArgs []AST// Closure is a closure, or lambda expression.

	TemplateArgsConstraint AST
	Types                  []AST
	Num                    int
	CallConstraint         AST
}

type StructuredBindings struct {
	Bindings []AST// StructuredBindings is a structured binding declaration.
}

type UnnamedType struct{ Num int } // UnnamedType is an unnamed type, that just has an index.

type Clone struct {
	Base   AST
	Suffix string
} // Clone is a clone of a function, with a distinguishing suffix.

type Special struct {
	Prefix string
	Val    AST
} // Special is a special symbol, printed as a prefix plus another
// value.

type Special2 struct {
	Prefix string
	Val1   AST
	Middle string
	Val2   AST
} // Special2 is like special, but uses two values.

type EnableIf struct {
	Type AST
	Args []AST// EnableIf is used by clang for an enable_if attribute.

}

type ModuleName struct {
	Parent      AST
	Name        AST
	IsPartition bool
} // ModuleName is a C++20 module.

type ModuleEntity struct {
	Module AST
	Name   AST
} // ModuleEntity is a name inside a module.

type Friend struct{ Name AST } // Friend is a member like friend name.

type Constraint struct {
	Name     AST
	Requires AST
} // Constraint represents an AST with a constraint.

type RequiresExpr struct {
	Params []AST// RequiresExpr is a C++20 requires expression.

	Requirements []AST
}

type ExprRequirement struct {
	Expr     AST
	Noexcept bool
	TypeReq  AST
} // ExprRequirement is a simple requirement in a requires expression.
// This is an arbitrary expression.

type TypeRequirement struct{ Type AST } // TypeRequirement is a type requirement in a requires expression.

type NestedRequirement struct{ Constraint AST } // NestedRequirement is a nested requirement in a requires expression.

type ExplicitObjectParameter struct{ Base AST } // ExplicitObjectParameter represents a C++23 explicit object parameter.

type innerPrinter interface{ printInner(*printState) } // innerPrinter is an interface for types that can print themselves as
// inner types.

type demangleErr struct {
	err string
	off int
} // A demangleErr is an error at a specific offset in the mangled
// string.

type operator struct {
	name string
	args int
	prec precedence
} // An operator is the demangled name, and the number of arguments it
// takes in an expression.

type rustState struct {
	orig          string
	str           string
	off           int
	buf           strings.Builder
	skip          bool
	lifetimes     int64
	last          byte
	noGenericArgs bool
	max           int
} // A rustState holds the current state of demangling a Rust string.
// maximum output length

type instFormat struct {
	mask     uint32
	value    uint32
	priority int8
	op       Op
	opBits   uint64
	args     instArgs
} // An instFormat describes the format of an instruction encoding.
// An instruction with 32-bit value x matches the format if x&mask == value
// and the condition matches.
// The condition matches if x>>28 == 0xF && value>>28==0xF
// or if x>>28 != 0xF and value>>28 == 0.
// If x matches the format, then the rest of the fields describe how to interpret x.
// The opBits describe bits that should be extracted from x and added to the opcode.
// For example opBits = 0x1234 means that the value
//
//	(2 bits at offset 1) followed by (4 bits at offset 3)
//
// should be added to op.
// Finally the args describe how to decode the instruction arguments.
// args is stored as a fixed-size array; if there are fewer than len(args) arguments,
// args[i] == 0 marks the end of the argument list.

type Arg interface {
	IsArg()
	String() string
} // An Arg is a single instruction argument, one of these types:
// Endian, Imm, Mem, PCRel, Reg, RegList, RegShift, RegShiftReg.

type ImmAlt struct {
	Val uint8
	Rot uint8
} // An ImmAlt is an alternate encoding of an integer constant.

type RegX struct {
	Reg   Reg
	Index int
} // A RegX represents a fraction of a multi-value register.
// The Index field specifies the index number,
// but the size of the fraction is not specified.
// It must be inferred from the instruction and the register type.
// For example, in a VMOV instruction, RegX{D5, 1} represents
// the top 32 bits of the 64-bit D5 register.

type RegShift struct {
	Reg   Reg
	Shift Shift
	Count uint8
} // A RegShift is a register shifted by a constant.

type RegShiftReg struct {
	Reg      Reg
	Shift    Shift
	RegCount Reg
} // A RegShiftReg is a register shifted by a register.

type Mem struct {
	Base   Reg
	Mode   AddrMode
	Sign   int8
	Index  Reg
	Shift  Shift
	Count  uint8
	Offset int16
} // A Mem is a memory reference made up of a base R and index expression X.
// The effective memory address is R or R+X depending on AddrMode.
// The index expression is X = Sign*(Index Shift Count) + Offset,
// but in any instruction either Sign = 0 or Offset = 0.

type goFPInfo struct {
	op        Op
	transArgs []int
	gnuName   string
	goName    string
} // indexes of arguments which need transformation
// instruction name in Plan 9 syntax

type ImmShift struct {
	imm   uint16
	shift uint8
}

type RegExtshiftAmount struct {
	reg       Reg
	extShift  ExtShift
	amount    uint8
	show_zero bool
}

type MemImmediate struct {
	Base RegSP
	Mode AddrMode
	imm  int32
} // A MemImmediate is a memory reference made up of a base R and immediate X.
// The effective memory address is R or R+X depending on AddrMode.

type MemExtend struct {
	Base            RegSP
	Index           Reg
	Extend          ExtShift
	Amount          uint8
	ShiftMustBeZero bool
} // A MemExtend is a memory reference made up of a base R and index expression X.
// The effective memory address is R or R+X depending on Index, Extend and Amount.
// Refer to ARM reference manual, for byte load/store(register), the index
// shift amount must be 0, encoded in "S" as 0 if omitted, or as 1 if present.
// a.ShiftMustBeZero is set true indicates the index shift amount must be 0.
// In GNU syntax, a #0 shift amount is printed if Amount is 1 but ShiftMustBeZero
// is true; #0 is not printed if Amount is 0 and ShiftMustBeZero is true.
// Both cases represent shift by 0 bit.

type Imm64 struct {
	Imm     uint64
	Decimal bool
}

type Systemreg struct {
	op0 uint8
	op1 uint8
	cn  uint8
	cm  uint8
	op2 uint8
}

type Imm_fp struct {
	s   uint8
	exp int8
	pre uint8
} // An Imm_fp is a signed floating-point constant.

type RegisterWithArrangement struct {
	r   Reg
	a   Arrangement
	cnt uint8
} // Register with arrangement: <Vd>.<T>, { <Vt>.8B, <Vt2>.8B},

type RegisterWithArrangementAndIndex struct {
	r     Reg
	a     Arrangement
	index uint8
	cnt   uint8
} // Register with arrangement and index:
//
//	<Vm>.<Ts>[<index>],
//	{ <Vt>.B, <Vt2>.B }[<index>].

type sysOp struct {
	op          sysInstFields
	r           Reg
	hasOperand2 bool
}

type sysInstFields struct {
	op1 uint8
	cn  uint8
	cm  uint8
	op2 uint8
}

type sysInstAttrs struct {
	typ         sys
	name        string
	hasOperand2 bool
}

type Uimm struct {
	Imm     uint32
	Decimal bool
} // An Imm is an integer constant.

type Simm16 struct {
	Imm   int16
	Width uint8
}

type Simm32 struct {
	Imm   int32
	Width uint8
}

type OffsetSimm struct {
	Imm   int32
	Width uint8
}

type argField struct {
	Type  ArgType
	Shift uint8
	BitFields
} // argField indicate how to decode an argument to an instruction.
// First parse the value from the BitFields, shift it left by Shift
// bits to get the actual numerical value.

type InstMaskMap struct {
	mask uint64
	insn map[uint64]*instFormat
}

type BitField struct {
	Offs uint8
	Bits uint8
	Word uint8
} // A BitField is a bit-field in a 32-bit word.
// Bits are counted from 0 from the MSB to 31 as the LSB.
// This instruction word holding this field.
// It is always 0 for ISA < 3.1 instructions. It is
// in decoding order. (0 == prefix, 1 == suffix on ISA 3.1)

type Simm struct {
	Imm     int32
	Decimal bool
	Width   uint8
} // A Simm is a signed immediate number
// Actual width of the Simm

type AmoReg struct {
	reg Reg
} // An AmoReg is an atomic address register used in AMO instructions
// Avoid promoted String method

type RegOffset struct {
	OfsReg Reg
	Ofs    Simm
} // A RegOffset is a register with offset value

type typ1ExtndMnics struct {
	BaseOpStr string
	Value     uint8
	Offset    uint8
	ExtnOpStr string
} // Typ1 - Instructions having different base and extended mnemonic strings.
//
//	These instructions have single M-field value and single offset.

type typ2ExtndMnics struct {
	Value     uint8
	Offset    uint8
	ExtnOpStr string
} // Typ2 - Instructions having couple of extra strings added to the base mnemonic string,
//
//	depending on the condition code evaluation.
//	These instructions have single M-field value and single offset.

type typ3ExtndMnics struct {
	Value1    uint8
	Value2    uint8
	Offset1   uint8
	Offset2   uint8
	ExtnOpStr string
} // Typ3 - Instructions having couple of extra strings added to the base mnemonic string,
//
//	depending on the condition code evaluation.
//	These instructions have two M-field values and two offsets.

type typ4ExtndMnics struct {
	BaseOpStr string
	Value1    uint8
	Value2    uint8
	Offset1   uint8
	Offset2   uint8
	ExtnOpStr string
} // Typ4 - Instructions having different base and extended mnemonic strings.
//
//	These instructions have two M-field values and two offsets.

type typ5ExtndMnics struct {
	BaseOpStr string
	Value1    uint8
	Value2    uint8
	Value3    uint8
	Offset1   uint8
	Offset2   uint8
	Offset3   uint8
	ExtnOpStr string
} // Typ5 - Instructions having different base and extended mnemonic strings.
//
//	These instructions have three M-field values and three offsets.

type APIFeature struct {
	Package string
	Build   string
	Feature string
	Issue   int
} // An APIFeature is a symbol mentioned in an API file,
// like the ones in the main go repo in the api directory.
// the issue that introduced the feature, or 0 if none

type Regexp struct {
	str  string
	once sync.Once
	rx   *regexp.Regexp
} // Regexp is a wrapper around [regexp.Regexp], where the underlying regexp will be
// compiled the first time it is needed.

type Comments struct {
	Before []Comment// Comments collects the comments associated with an expression.

	Suffix []Comment// whole-line comments before this expression

	After []Comment// end-of-line comments after this expression
	// For top-level expressions only, After lists whole-line
	// comments following the expression.

}

type FileSyntax struct {
	Name string
	Comments
	Stmt []Expr// A FileSyntax represents an entire go.mod file.
	// file path

}

type CommentBlock struct {
	Comments
	Start Position
} // A CommentBlock represents a top-level block of comments separate
// from any rule.

type LineBlock struct {
	Comments
	Start  Position
	LParen LParen
	Token  []string// A LineBlock is a factored block of lines, like
	//
	//	require (
	//		"x"
	//		"y"
	//	)

	Line   []*Line
	RParen RParen
}

type LParen struct {
	Comments
	Pos Position
} // An LParen represents the beginning of a parenthesized line block.
// It is a place to store suffix comments.

type RParen struct {
	Comments
	Pos Position
} // An RParen represents the end of a parenthesized line block.
// It is a place to store whole-line (before) comments.

type input struct {
	filename string
	complete []byte// An input represents a single input file being parsed.
	// name of input file, for errors

	remaining []byte// entire input

	tokenStart []byte// remaining input

	token    token
	pos      Position
	comments []Comment// token being scanned to end of input
	// current input position

	file        *FileSyntax
	parseErrors ErrorList
	pre         []Expr// accumulated comments
	// Comment assignment state.

	post []Expr// all expressions, in preorder traversal

} // all expressions, in postorder traversal

type Go struct {
	Version string
	Syntax  *Line
} // A Go is the go statement.
// "1.23"

type Toolchain struct {
	Name   string
	Syntax *Line
} // A Toolchain is the toolchain statement.
// "go1.21rc1"

type Godebug struct {
	Key    string
	Value  string
	Syntax *Line
} // A Godebug is a single godebug key=value statement.

type Exclude struct {
	Mod    module.Version
	Syntax *Line
} // An Exclude is a single exclude statement.

type Replace struct {
	Old    module.Version
	New    module.Version
	Syntax *Line
} // A Replace is a single replace statement.

type Retract struct {
	VersionInterval
	Rationale string
	Syntax    *Line
} // A Retract is a single retract statement.

type Tool struct {
	Path   string
	Syntax *Line
} // A Tool is a single tool statement.

type Ignore struct {
	Path   string
	Syntax *Line
} // An Ignore is a single ignore statement.

type VersionInterval struct{ Low, High string } // A VersionInterval represents a range of versions with upper and lower bounds.
// Intervals are closed: both bounds are included. When Low is equal to High,
// the interval may refer to a single version ('v1.2.3') or an interval
// ('[v1.2.3, v1.2.3]'); both have the same representation.

type Require struct {
	Mod      module.Version
	Indirect bool
	Syntax   *Line
} // A Require is a single require statement.
// has "// indirect" comment

type WorkFile struct {
	Go        *Go
	Toolchain *Toolchain
	Godebug   []*// A WorkFile is the parsed, interpreted form of a go.work file.
	Godebug
	Use     []*Use
	Replace []*Replace
	Syntax  *FileSyntax
}

type Use struct {
	Path       string
	ModulePath string
	Syntax     *Line
} // A Use is a single directory statement.
// Module path in the comment.

type Version struct {
	Path    string
	Version string `json:",omitempty"`
} // A Version (for clients, a module.Version) is defined by a module path and version pair.
// These are stored in their plain (unescaped) form.
// Version is usually a semantic version in canonical form.
// There are three exceptions to this general rule.
// First, the top-level target of a build has no specific version
// and uses Version = "".
// Second, during MVS calculations the version "none" is used
// to represent the decision to take no version of a given module.
// Third, filesystem paths found in "replace" directives are
// represented by a path with an empty version.

type InvalidVersionError struct {
	Version string
	Pseudo  bool
	Err     error
} // An InvalidVersionError indicates an error specific to a version, with the
// module path unknown or specified externally.
//
// A [ModuleError] may wrap an InvalidVersionError, but an InvalidVersionError
// must not wrap a ModuleError.

type InvalidPathError struct {
	Kind string
	Path string
	Err  error
} // An InvalidPathError indicates a module, import, or file path doesn't
// satisfy all naming constraints. See [CheckPath], [CheckImportPath],
// and [CheckFilePath] for specific restrictions.
// "module", "import", or "file"

type parsed struct {
	major      string
	minor      string
	patch      string
	short      string
	prerelease string
	build      string
} // parsed returns the parsed form of a semantic version string.

type parCache struct{ m sync.Map } // parCache runs an action once per key and caches the result.

type ClientOps interface {
	ReadRemote(path string) ([]byte,// A ClientOps provides the external operations
	// (file caching, HTTP fetches, and so on) needed by the [Client].
	// The methods must be safe for concurrent use by multiple goroutines.
	// ReadRemote reads and returns the content served at the given path
	// on the remote database server. The path begins with "/lookup" or "/tile/",
	// and there is no need to parse the path in any way.
	// It is the implementation's responsibility to turn that path into a full URL
	// and make the HTTP request. ReadRemote should return an error for
	// any non-200 HTTP response status.
	error)
	ReadConfig(file string) ([]byte,// ReadConfig reads and returns the content of the named configuration file.
	// There are only a fixed set of configuration files.
	//
	// "key" returns a file containing the verifier key for the server.
	//
	// serverName + "/latest" returns a file containing the latest known
	// signed tree from the server.
	// To signal that the client wishes to start with an "empty" signed tree,
	// ReadConfig can return a successful empty result (0 bytes of data).
	error)
	WriteConfig(file string, old, new []byte) error// WriteConfig updates the content of the named configuration file,
	// changing it from the old []byte to the new []byte.
	// If the old []byte does not match the stored configuration,
	// WriteConfig must return ErrWriteConflict.
	// Otherwise, WriteConfig should atomically replace old with new.
	// The "key" configuration file is never written using WriteConfig.

	ReadCache(file string) ([]byte,// ReadCache reads and returns the content of the named cache file.
	// Any returned error will be treated as equivalent to the file not existing.
	// There can be arbitrarily many cache files, such as:
	//	serverName/lookup/pkg@version
	//	serverName/tile/8/1/x123/456
	error)
	WriteCache(file string, data []byte)// WriteCache writes the named cache file.

	Log(msg string)
	SecurityError(msg string)
} // Log prints the given log message (such as with log.Print)
// SecurityError prints the given security error log message.
// The Client returns ErrSecurity from any operation that invokes SecurityError,
// but the return value is mainly for testing. In a real program,
// SecurityError should typically print the message and call log.Fatal or os.Exit.

type Client struct {
	ops        ClientOps
	didLookup  uint32
	initOnce   sync.Once
	initErr    error
	name       string
	verifiers  note.Verifiers
	tileReader tileReader
	tileHeight int
	nosumdb    string
	record     parCache
	tileCache  parCache
	latestMu   sync.Mutex
	latest     tlog.Tree
	latestMsg  []byte// A Client is a client connection to a checksum database.
	// All the methods are safe for simultaneous use by multiple goroutines.
	// latest known tree head

	tileSavedMu sync.Mutex
	tileSaved   map[ // encoded signed note for latest
	tlog.Tile]bool
} // which tiles have been saved using c.ops.WriteCache already

type tileReader struct{ c *Client } // tileReader is a *Client wrapper that implements tlog.TileReader.
// The separate type avoids exposing the ReadTiles and SaveTiles
// methods on Client itself.

type Verifier interface {
	Name() string
	KeyHash() uint32
	Verify(msg, sig []byte) bool// A Verifier verifies messages signed with a specific key.
	// Verify reports whether sig is a valid signature of msg.

}

type Signer interface {
	Name() string
	KeyHash() uint32
	Sign(msg []byte) (// A Signer signs messages using a specific key.
	// Sign returns a signature for the given message.
	[]byte, error)
}

type verifier struct {
	name   string
	hash   uint32
	verify func([]byte,// verifier is a trivial Verifier implementation.
	[]byte) bool
}

type signer struct {
	name string
	hash uint32
	sign func([]byte) (// signer is a trivial Signer implementation.
	[]byte, error)
}

type Verifiers interface {
	Verifier(name string, hash uint32) (Verifier, error)
} // A Verifiers is a collection of known verifier keys.
// Verifier returns the Verifier associated with the key
// identified by the name and hash.
// If the name, hash pair is unknown, Verifier should return
// an UnknownVerifierError.

type UnknownVerifierError struct {
	Name    string
	KeyHash uint32
} // An UnknownVerifierError indicates that the given key is not known.
// The Open function records signatures without associated verifiers as
// unverified signatures.

type ambiguousVerifierError struct {
	name string
	hash uint32
} // An ambiguousVerifierError indicates that the given name and hash
// match multiple keys passed to [VerifierList].
// (If this happens, some malicious actor has taken control of the
// verifier list, at which point we may as well give up entirely,
// but we diagnose the problem instead.)

type nameHash struct {
	name string
	hash uint32
}

type Note struct {
	Text string
	Sigs []Signature// A Note is a text and signatures.
	// text of note

	UnverifiedSigs []Signature// verified signatures

} // unverified signatures

type UnverifiedNoteError struct{ Note *Note } // An UnverifiedNoteError indicates that the note
// successfully parsed but had no verifiable signatures.

type InvalidSignatureError struct {
	Name string
	Hash uint32
} // An InvalidSignatureError indicates that the given key was known
// and the associated Verifier rejected the signature.

type ServerOps interface {
	Signed(ctx context.Context) ([]byte,// A ServerOps provides the external operations
	// (underlying database access and so on) needed by the [Server].
	// Signed returns the signed hash of the latest tree.
	error)
	ReadRecords(ctx context.Context, id, n int64) ([][// ReadRecords returns the content for the n records id through id+n-1.
	]byte, error)
	Lookup(ctx context.Context, m module.Version) (int64, error)
	ReadTileData(ctx context.Context, t tlog.Tile) ([]byte,// Lookup looks up a record for the given module,
	// returning the record ID.
	// ReadTileData reads the content of tile t.
	// It is only invoked for hash tiles (t.L ≥ 0).
	error)
}

type TestServer struct {
	signer string
	gosum  func(path, vers string) ([]byte,// A TestServer is an in-memory implementation of [ServerOps] for testing.
	error)
	mu      sync.Mutex
	hashes  testHashes
	records [][]byte
	lookup  map[string]int64
}

type Tree struct {
	N    int64
	Hash Hash
} // A Tree is a tree description, to be signed by a go.sum database server.

type Tile struct {
	H int
	L int
	N int64
	W int
} // A Tile is a description of a transparency log tile.
// A tile of height H at level L offset N lists W consecutive hashes
// at level H*L of the tree starting at offset N*(2**H).
// A complete tile lists 2**H hashes; a partial tile lists fewer.
// Note that a tile represents the entire subtree of height H
// with those hashes as the leaves. The levels above H*L
// can be reconstructed by hashing the leaves.
//
// Each Tile can be encoded as a “tile coordinate path”
// of the form tile/H/L/NNN[.p/W].
// The .p/W suffix is present only for partial tiles, meaning W < 2**H.
// The NNN element is an encoding of N into 3-digit path elements.
// All but the last path element begins with an "x".
// For example,
// Tile{H: 3, L: 4, N: 1234067, W: 1}'s path
// is tile/3/4/x001/x234/067.p/1, and
// Tile{H: 3, L: 4, N: 1234067, W: 8}'s path
// is tile/3/4/x001/x234/067.
// See the [Tile.Path] method and the [ParseTilePath] function.
//
// The special level L=-1 holds raw record data instead of hashes.
// In this case, the level encodes into a tile path as the path element
// "data" instead of "-1".
//
// See also https://golang.org/design/25530-sumdb#checksum-database
// and https://research.swtch.com/tlog#tiling_a_log.
// width of tile (1 ≤ W ≤ 2**H; 2**H is complete tile)

type badPathError struct{ path string }

type TileReader interface {
	Height() int
	ReadTiles(tiles []Tile) (// A TileReader reads tiles from a go.sum database log.
	// ReadTiles returns the data for each requested tile.
	// If ReadTiles returns err == nil, it must also return
	// a data record for each tile (len(data) == len(tiles))
	// and each data record must be the correct length
	// (len(data[i]) == tiles[i].W*HashSize).
	//
	// An implementation of ReadTiles typically reads
	// them from an on-disk cache or else from a remote
	// tile server. Tile data downloaded from a server should
	// be considered suspect and not saved into a persistent
	// on-disk cache before returning from ReadTiles.
	// When the client confirms the validity of the tile data,
	// it will call SaveTiles to signal that they can be safely
	// written to persistent storage.
	// See also https://research.swtch.com/tlog#authenticating_tiles.
	data [][]byte, err error)
	SaveTiles(tiles []Tile,// SaveTiles informs the TileReader that the tile data
	// returned by ReadTiles has been confirmed as valid
	// and can be saved in persistent storage (on disk).
	data [][]byte)
}

type tileHashReader struct {
	tree Tree
	tr   TileReader
}

type HashReader interface {
	ReadHashes(indexes []int64) (// A HashReader can read hashes for nodes in the log's tree structure.
	// ReadHashes returns the hashes with the given stored hash indexes
	// (see StoredHashIndex and SplitStoredHashIndex).
	// ReadHashes must return a slice of hashes the same length as indexes,
	// or else it must return a non-nil error.
	// ReadHashes may run faster if indexes is sorted in increasing order.
	[]Hash, error)
}

type CheckedFiles struct {
	Valid []string// CheckedFiles reports whether a set of files satisfy the name and size
	// constraints required by module zip files. The constraints are listed in the
	// package documentation.
	//
	// Functions that produce this report may include slightly different sets of
	// files. See documentation for CheckFiles, CheckDir, and CheckZip for details.
	// Valid is a list of file paths that should be included in a zip file.

	Omitted []FileError// Omitted is a list of files that are ignored when creating a module zip
	// file, along with the reason each file is ignored.

	Invalid []FileError// Invalid is a list of files that should not be included in a module zip
	// file, along with the reason each file is invalid.

	SizeError error
} // SizeError is non-nil if the total uncompressed size of the valid files
// exceeds the module zip size limit or if the zip file itself exceeds the
// limit.

type FileError struct {
	Path string
	Err  error
}

type UnrecognizedVCSError struct{ RepoRoot string } // UnrecognizedVCSError indicates that no recognized version control system was
// found in the given directory.

type dirFile struct {
	filePath, slashPath string
	info                os.FileInfo
}

type pathInfo struct {
	path  string
	isDir bool
}

type zipError struct {
	verb, path string
	err        error
}

type PanicError struct {
	Recovered error
	Stack     []byte// PanicError wraps an error recovered from an unhandled panic
	// when calling a function passed to Go or TryGo.

} // result of call to [debug.Stack]

type PanicValue struct {
	Recovered any
	Stack     []byte// PanicValue wraps a value that does not implement the error interface,
	// recovered from an unhandled panic when calling a function passed to Go or
	// TryGo.

} // result of call to [debug.Stack]

type waiter struct {
	n     int64
	ready chan<- struct{}
} // Closed when semaphore acquired.

type Weighted struct {
	size    int64
	cur     int64
	mu      sync.Mutex
	waiters list.List
} // Weighted provides a way to bound concurrent access to a resource.
// The callers can request access with a given weight.

type Qid struct {
	Path uint64
	Vers uint32
	Type uint8
} // A Qid represents a 9P server's unique identification for a file.
// the type of the file (plan9.QTDIR for example)

type Waitmsg struct {
	Pid  int
	Time [3]uint32
	Msg  string
}

type Timespec struct {
	Sec  int32
	Nsec int32
}

type Timeval struct {
	Sec  int32
	Usec int32
}

type Pgtha struct {
	Pid        uint32
	Tid0       uint32
	Tid1       uint32
	Accesspid  byte
	Accesstid  byte
	Accessasid uint16
	Loginname  [8]byte
	Flag1      byte
	Flag1b2    byte
} // 0
// 19

type Bpxystat_t struct {
	St_id           [4]uint8
	St_length       uint16
	St_version      uint16
	St_mode         uint32
	St_ino          uint32
	St_dev          uint32
	St_nlink        uint32
	St_uid          uint32
	St_gid          uint32
	St_size         uint64
	St_atime        uint32
	St_mtime        uint32
	St_ctime        uint32
	St_rdev         uint32
	St_auditoraudit uint32
	St_useraudit    uint32
	St_blksize      uint32
	St_createtime   uint32
	St_auditid      [4]uint32
	St_res01        uint32
	Ft_ccsid        uint16
	Ft_flags        uint16
	St_res01a       [2]uint32
	St_res02        uint32
	St_blocks       uint32
	St_opaque       [3]uint8
	St_visible      uint8
	St_reftime      uint32
	St_fid          uint64
	St_filefmt      uint8
	St_fspflag2     uint8
	St_res03        [2]uint8
	St_ctimemsec    uint32
	St_seclabel     [8]uint8
	St_res04        [4]uint8
	_               uint32
	St_atime64      uint64
	St_mtime64      uint64
	St_ctime64      uint64
	St_createtime64 uint64
	St_reftime64    uint64
	_               uint64
	St_res05        [16]uint8
} // 0
// 0xc8

type BpxFilestatus struct {
	Oflag1 byte
	Oflag2 byte
	Oflag3 byte
	Oflag4 byte
}

type BpxMode struct {
	Ftype byte
	Mode1 byte
	Mode2 byte
	Mode3 byte
}

type Bpxyatt_t struct {
	Att_id           [4]uint8
	Att_version      uint16
	Att_res01        [2]uint8
	Att_setflags1    uint8
	Att_setflags2    uint8
	Att_setflags3    uint8
	Att_setflags4    uint8
	Att_mode         uint32
	Att_uid          uint32
	Att_gid          uint32
	Att_opaquemask   [3]uint8
	Att_visblmaskres uint8
	Att_opaque       [3]uint8
	Att_visibleres   uint8
	Att_size_h       uint32
	Att_size_l       uint32
	Att_atime        uint32
	Att_mtime        uint32
	Att_auditoraudit uint32
	Att_useraudit    uint32
	Att_ctime        uint32
	Att_reftime      uint32
	Att_filefmt      uint8
	Att_res02        [3]uint8
	Att_filetag      uint32
	Att_res03        [8]uint8
	Att_atime64      uint64
	Att_mtime64      uint64
	Att_ctime64      uint64
	Att_reftime64    uint64
	Att_seclabel     [8]uint8
	Att_ver3res02    [8]uint8
} // Thr attribute structure for extended attributes
// end of version 2

type Ifreq struct{ raw ifreq } // An Ifreq is a type-safe wrapper around the raw ifreq struct. An Ifreq
// contains an interface name and a union of arbitrary data which can be
// accessed using the Ifreq's methods. To create an Ifreq, use the NewIfreq
// function.
//
// Use the Name method to access the stored interface name. The union data
// fields can be get and set using the following methods:
//   - Uint16/SetUint16: flags
//   - Uint32/SetUint32: ifindex, metric, mtu

type ifreqData struct {
	name [IFNAMSIZ]byte
	data unsafe.Pointer
	_    [len(ifreq{}.Ifru) - SizeofPtr]byte
} // An ifreqData is an Ifreq which carries pointer data. To produce an ifreqData,
// use the Ifreq.withData method.
// Pad to the same size as ifreq.

type FileDedupeRange struct {
	Src_offset uint64
	Src_length uint64
	Reserved1  uint16
	Reserved2  uint32
	Info       []FileDedupeRangeInfo
}

type FileDedupeRangeInfo struct {
	Dest_fd       int64
	Dest_offset   uint64
	Bytes_deduped uint64
	Status        int32
	Reserved      uint32
}

type mremapMmapper struct {
	mmapper
	mremap func(oldaddr uintptr, oldlength uintptr, newlength uintptr, flags int, newaddr uintptr) (xaddr uintptr, err error)
}

type SocketControlMessage struct {
	Header Cmsghdr
	Data   []byte// SocketControlMessage represents a socket control message.

}

type SockaddrDatalink struct {
	Len    uint8
	Family uint8
	Index  uint16
	Type   uint8
	Nlen   uint8
	Alen   uint8
	Slen   uint8
	Data   [12]int8
	raw    RawSockaddrDatalink
} // SockaddrDatalink implements the Sockaddr interface for AF_LINK type sockets.

type SockaddrCtl struct {
	ID   uint32
	Unit uint32
	raw  RawSockaddrCtl
} // SockaddrCtl implements the Sockaddr interface for AF_SYSTEM type sockets.

type SockaddrVM struct {
	CID  uint32
	Port uint32
	raw  RawSockaddrVM
} // SockaddrVM implements the Sockaddr interface for AF_VSOCK type sockets.
// SockaddrVM provides access to Darwin VM sockets: a mechanism that enables
// bidirectional communication between a hypervisor and its guest virtual
// machines.
// CID and Port specify a context ID and port address for a VM socket.
// Guests have a unique CID, and hosts may have a well-known CID of:
//  - VMADDR_CID_HYPERVISOR: refers to the hypervisor process.
//  - VMADDR_CID_LOCAL: refers to local communication (loopback).
//  - VMADDR_CID_HOST: refers to other processes on the host.

type IfreqMTU struct {
	Name [IFNAMSIZ]byte
	MTU  int32
} // IfreqMTU is struct ifreq used to get or set a network device's MTU.

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

type Termios struct {
	Iflag  uint32
	Oflag  uint32
	Cflag  uint32
	Lflag  uint32
	Cc     [20]uint8
	Ispeed int32
	Ospeed int32
}

type SockaddrLinklayer struct {
	Protocol uint16
	Ifindex  int
	Hatype   uint16
	Pkttype  uint8
	Halen    uint8
	Addr     [8]byte
	raw      RawSockaddrLinklayer
} // SockaddrLinklayer implements the Sockaddr interface for AF_PACKET type sockets.

type SockaddrNetlink struct {
	Family uint16
	Pad    uint16
	Pid    uint32
	Groups uint32
	raw    RawSockaddrNetlink
} // SockaddrNetlink implements the Sockaddr interface for AF_NETLINK type sockets.

type SockaddrHCI struct {
	Dev     uint16
	Channel uint16
	raw     RawSockaddrHCI
} // SockaddrHCI implements the Sockaddr interface for AF_BLUETOOTH type sockets
// using the HCI protocol.

type SockaddrL2 struct {
	PSM      uint16
	CID      uint16
	Addr     [6]uint8
	AddrType uint8
	raw      RawSockaddrL2
} // SockaddrL2 implements the Sockaddr interface for AF_BLUETOOTH type sockets
// using the L2CAP protocol.

type SockaddrRFCOMM struct {
	Addr    [6]uint8
	Channel uint8
	raw     RawSockaddrRFCOMM
} // SockaddrRFCOMM implements the Sockaddr interface for AF_BLUETOOTH type sockets
// using the RFCOMM protocol.
//
// Server example:
//
//	fd, _ := Socket(AF_BLUETOOTH, SOCK_STREAM, BTPROTO_RFCOMM)
//	_ = unix.Bind(fd, &unix.SockaddrRFCOMM{
//		Channel: 1,
//		Addr:    [6]uint8{0, 0, 0, 0, 0, 0}, // BDADDR_ANY or 00:00:00:00:00:00
//	})
//	_ = Listen(fd, 1)
//	nfd, sa, _ := Accept(fd)
//	fmt.Printf("conn addr=%v fd=%d", sa.(*unix.SockaddrRFCOMM).Addr, nfd)
//	Read(nfd, buf)
//
// Client example:
//
//	fd, _ := Socket(AF_BLUETOOTH, SOCK_STREAM, BTPROTO_RFCOMM)
//	_ = Connect(fd, &SockaddrRFCOMM{
//		Channel: 1,
//		Addr:    [6]byte{0x11, 0x22, 0x33, 0xaa, 0xbb, 0xcc}, // CC:BB:AA:33:22:11
//	})
//	Write(fd, []byte(`hello`))
// Channel is a designated bluetooth channel, only 1-30 are available for use.
// Since Linux 2.6.7 and further zero value is the first available channel.

type SockaddrCAN struct {
	Ifindex int
	RxID    uint32
	TxID    uint32
	raw     RawSockaddrCAN
} // SockaddrCAN implements the Sockaddr interface for AF_CAN type sockets.
// The RxID and TxID fields are used for transport protocol addressing in
// (CAN_TP16, CAN_TP20, CAN_MCNET, and CAN_ISOTP), they can be left with
// zero values for CAN_RAW and CAN_BCM sockets as they have no meaning.
//
// The SockaddrCAN struct must be bound to the socket file descriptor
// using Bind before the CAN socket can be used.
//
//	// Read one raw CAN frame
//	fd, _ := Socket(AF_CAN, SOCK_RAW, CAN_RAW)
//	addr := &SockaddrCAN{Ifindex: index}
//	Bind(fd, addr)
//	frame := make([]byte, 16)
//	Read(fd, frame)
//
// The full SocketCAN documentation can be found in the linux kernel
// archives at: https://www.kernel.org/doc/Documentation/networking/can.txt

type SockaddrCANJ1939 struct {
	Ifindex int
	Name    uint64
	PGN     uint32
	Addr    uint8
	raw     RawSockaddrCAN
} // SockaddrCANJ1939 implements the Sockaddr interface for AF_CAN using J1939
// protocol (https://en.wikipedia.org/wiki/SAE_J1939). For more information
// on the purposes of the fields, check the official linux kernel documentation
// available here: https://www.kernel.org/doc/Documentation/networking/j1939.rst

type SockaddrALG struct {
	Type    string
	Name    string
	Feature uint32
	Mask    uint32
	raw     RawSockaddrALG
} // SockaddrALG implements the Sockaddr interface for AF_ALG type sockets.
// SockaddrALG enables userspace access to the Linux kernel's cryptography
// subsystem. The Type and Name fields specify which type of hash or cipher
// should be used with a given socket.
//
// To create a file descriptor that provides access to a hash or cipher, both
// Bind and Accept must be used. Once the setup process is complete, input
// data can be written to the socket, processed by the kernel, and then read
// back as hash output or ciphertext.
//
// Here is an example of using an AF_ALG socket with SHA1 hashing.
// The initial socket setup process is as follows:
//
//	// Open a socket to perform SHA1 hashing.
//	fd, _ := unix.Socket(unix.AF_ALG, unix.SOCK_SEQPACKET, 0)
//	addr := &unix.SockaddrALG{Type: "hash", Name: "sha1"}
//	unix.Bind(fd, addr)
//	// Note: unix.Accept does not work at this time; must invoke accept()
//	// manually using unix.Syscall.
//	hashfd, _, _ := unix.Syscall(unix.SYS_ACCEPT, uintptr(fd), 0, 0)
//
// Once a file descriptor has been returned from Accept, it may be used to
// perform SHA1 hashing. The descriptor is not safe for concurrent use, but
// may be re-used repeatedly with subsequent Write and Read operations.
//
// When hashing a small byte slice or string, a single Write and Read may
// be used:
//
//	// Assume hashfd is already configured using the setup process.
//	hash := os.NewFile(hashfd, "sha1")
//	// Hash an input string and read the results. Each Write discards
//	// previous hash state. Read always reads the current state.
//	b := make([]byte, 20)
//	for i := 0; i < 2; i++ {
//	    io.WriteString(hash, "Hello, world.")
//	    hash.Read(b)
//	    fmt.Println(hex.EncodeToString(b))
//	}
//	// Output:
//	// 2ae01472317d1935a84797ec1983ae243fc6aa28
//	// 2ae01472317d1935a84797ec1983ae243fc6aa28
//
// For hashing larger byte slices, or byte streams such as those read from
// a file or socket, use Sendto with MSG_MORE to instruct the kernel to update
// the hash digest instead of creating a new one for a given chunk and finalizing it.
//
//	// Assume hashfd and addr are already configured using the setup process.
//	hash := os.NewFile(hashfd, "sha1")
//	// Hash the contents of a file.
//	f, _ := os.Open("/tmp/linux-4.10-rc7.tar.xz")
//	b := make([]byte, 4096)
//	for {
//	    n, err := f.Read(b)
//	    if err == io.EOF {
//	        break
//	    }
//	    unix.Sendto(hashfd, b[:n], unix.MSG_MORE, addr)
//	}
//	hash.Read(b)
//	fmt.Println(hex.EncodeToString(b))
//	// Output: 85cdcad0c06eef66f805ecce353bec9accbeecc5
//
// For more information, see: http://www.chronox.de/crypto-API/crypto/userspace-if.html.

type SockaddrXDP struct {
	Flags        uint16
	Ifindex      uint32
	QueueID      uint32
	SharedUmemFD uint32
	raw          RawSockaddrXDP
}

type SockaddrPPPoE struct {
	SID    uint16
	Remote []byte
	Dev    string
	raw    RawSockaddrPPPoX
}

type SockaddrTIPC struct {
	Scope int
	Addr  TIPCAddr
	raw   RawSockaddrTIPC
} // SockaddrTIPC implements the Sockaddr interface for AF_TIPC type sockets.
// For more information on TIPC, see: http://tipc.sourceforge.net/.
// Addr is the type of address used to manipulate a socket. Addr must be
// one of:
//  - *TIPCSocketAddr: "id" variant in the C addr union
//  - *TIPCServiceRange: "nameseq" variant in the C addr union
//  - *TIPCServiceName: "name" variant in the C addr union
//
// If nil, EINVAL will be returned when the structure is used.

type TIPCAddr interface {
	tipcAddrtype() uint8
	tipcAddr() [12]byte
} // TIPCAddr is implemented by types that can be used as an address for
// SockaddrTIPC. It is only implemented by *TIPCSocketAddr, *TIPCServiceRange,
// and *TIPCServiceName.

type SockaddrL2TPIP struct {
	Addr   [4]byte
	ConnId uint32
	raw    RawSockaddrL2TPIP
} // SockaddrL2TPIP implements the Sockaddr interface for IPPROTO_L2TP/AF_INET sockets.

type SockaddrL2TPIP6 struct {
	Addr   [16]byte
	ZoneId uint32
	ConnId uint32
	raw    RawSockaddrL2TPIP6
} // SockaddrL2TPIP6 implements the Sockaddr interface for IPPROTO_L2TP/AF_INET6 sockets.

type SockaddrIUCV struct {
	UserID string
	Name   string
	raw    RawSockaddrIUCV
} // SockaddrIUCV implements the Sockaddr interface for AF_IUCV sockets.

type SockaddrNFC struct {
	DeviceIdx   uint32
	TargetIdx   uint32
	NFCProtocol uint32
	raw         RawSockaddrNFC
}

type SockaddrNFCLLCP struct {
	DeviceIdx      uint32
	TargetIdx      uint32
	NFCProtocol    uint32
	DestinationSAP uint8
	SourceSAP      uint8
	ServiceName    string
	raw            RawSockaddrNFCLLCP
}

type fileHandle struct {
	Bytes uint32
	Type  int32
} // fileHandle is the argument to nameToHandleAt and openByHandleAt. We
// originally tried to generate it via unix/linux/types.go with "type
// fileHandle C.struct_file_handle" but that generated empty structs
// for mips64 and mips64le. Instead, hard code it for now (it's the
// same everywhere else) until the mips64 generator issue is fixed.

type FileHandle struct{ *fileHandle } // FileHandle represents the C struct file_handle used by
// name_to_handle_at (see NameToHandleAt) and open_by_handle_at (see
// OpenByHandleAt).

type RemoteIovec struct {
	Base uintptr
	Len  int
} // RemoteIovec is Iovec with the pointer replaced with an integer.
// It is used for ProcessVMReadv and ProcessVMWritev, where the pointer
// refers to a location in a different process' address space, which
// would confuse the Go garbage collector.

type rlimit32 struct {
	Cur uint32
	Max uint32
}

type stat_t struct {
	Dev        uint32
	Pad0       [3]int32
	Ino        uint64
	Mode       uint32
	Nlink      uint32
	Uid        uint32
	Gid        uint32
	Rdev       uint32
	Pad1       [3]uint32
	Size       int64
	Atime      uint32
	Atime_nsec uint32
	Mtime      uint32
	Mtime_nsec uint32
	Ctime      uint32
	Ctime_nsec uint32
	Blksize    uint32
	Pad2       uint32
	Blocks     int64
}

type fileObjCookie struct {
	fobj   *fileObj
	cookie interface{}
}

type EventPort struct {
	port int
	mu   sync.Mutex
	fds  map[ // EventPort provides a safe abstraction on top of Solaris/illumos Event Ports.
	uintptr]*fileObjCookie
	paths   map[string]*fileObjCookie
	cookies map[ // The user cookie presents an interesting challenge from a memory management perspective.
	// There are two paths by which we can discover that it is no longer in use:
	// 1. The user calls port_dissociate before any events fire
	// 2. An event fires and we return it to the user
	// The tricky situation is if the event has fired in the kernel but
	// the user hasn't requested/received it yet.
	// If the user wants to port_dissociate before the event has been processed,
	// we should handle things gracefully. To do so, we need to keep an extra
	// reference to the cookie around until the event is processed
	// thus the otherwise seemingly extraneous "cookies" map
	// The key of this map is a pointer to the corresponding fCookie
	*fileObjCookie]struct{}
}

type PortEvent struct {
	Cookie interface{}
	Events int32
	Fd     uintptr
	Path   string
	Source uint16
	fobj   *fileObj
} // PortEvent is an abstraction of the port_event C struct.
// Compare Source against PORT_SOURCE_FILE or PORT_SOURCE_FD
// to see if Path or Fd was the event source. The other will be
// uninitialized.

type Ucred struct{ ucred uintptr } // Ucred is an opaque struct that holds user credentials.

type mmapper struct {
	sync.Mutex
	active map[*byte][]byte
	mmap   func(addr, length uintptr, prot, flags, fd int, offset int64) (uintptr, error)
	munmap func(addr uintptr, length uintptr) error
} // active mappings; key is last byte in mapping

type Sockaddr interface {
	sockaddr() (ptr unsafe.Pointer, len _Socklen, err error)
} // Sockaddr represents a socket address.
// lowercase; only we can define Sockaddrs

type SockaddrInet4 struct {
	Port int
	Addr [4]byte
	raw  RawSockaddrInet4
} // SockaddrInet4 implements the Sockaddr interface for AF_INET type sockets.

type SockaddrInet6 struct {
	Port   int
	ZoneId uint32
	Addr   [16]byte
	raw    RawSockaddrInet6
} // SockaddrInet6 implements the Sockaddr interface for AF_INET6 type sockets.

type SockaddrUnix struct {
	Name string
	raw  RawSockaddrUnix
} // SockaddrUnix implements the Sockaddr interface for AF_UNIX type sockets.

type nwmTriplet struct {
	offset uint32
	length uint32
	number uint32
}

type nwmQuadruplet struct {
	offset uint32
	length uint32
	number uint32
	match  uint32
}

type nwmHeader struct {
	ident       uint32
	length      uint32
	version     uint16
	nwmType     uint16
	bytesNeeded uint32
	options     uint32
	_           [16]byte
	inputDesc   nwmTriplet
	outputDesc  nwmQuadruplet
}

type nwmFilter struct {
	ident         uint32
	flags         uint32
	resourceName  [8]byte
	resourceId    uint32
	listenerId    uint32
	local         [28]byte
	remote        [28]byte
	_             uint16
	_             uint16
	asid          uint16
	_             [2]byte
	tnLuName      [8]byte
	tnMonGrp      uint32
	tnAppl        [8]byte
	applData      [40]byte
	nInterface    [16]byte
	dVipa         [16]byte
	dVipaPfx      uint16
	dVipaPort     uint16
	dVipaFamily   byte
	_             [3]byte
	destXCF       [16]byte
	destXCFPfx    uint16
	destXCFFamily byte
	_             [1]byte
	targIP        [16]byte
	targIPPfx     uint16
	targIPFamily  byte
	_             [1]byte
	_             [20]byte
} // union of sockaddr4 and sockaddr6
// union of sockaddr4 and sockaddr6

type nwmRecHeader struct {
	ident  uint32
	length uint32
	number byte
	_      [3]byte
}

type nwmTCPStatsEntry struct {
	ident             uint64
	currEstab         uint32
	activeOpened      uint32
	passiveOpened     uint32
	connClosed        uint32
	estabResets       uint32
	attemptFails      uint32
	passiveDrops      uint32
	timeWaitReused    uint32
	inSegs            uint64
	predictAck        uint32
	predictData       uint32
	inDupAck          uint32
	inBadSum          uint32
	inBadLen          uint32
	inShort           uint32
	inDiscOldTime     uint32
	inAllBeforeWin    uint32
	inSomeBeforeWin   uint32
	inAllAfterWin     uint32
	inSomeAfterWin    uint32
	inOutOfOrder      uint32
	inAfterClose      uint32
	inWinProbes       uint32
	inWinUpdates      uint32
	outWinUpdates     uint32
	outSegs           uint64
	outDelayAcks      uint32
	outRsts           uint32
	retransSegs       uint32
	retransTimeouts   uint32
	retransDrops      uint32
	pmtuRetrans       uint32
	pmtuErrors        uint32
	outWinProbes      uint32
	probeDrops        uint32
	keepAliveProbes   uint32
	keepAliveDrops    uint32
	finwait2Drops     uint32
	acceptCount       uint64
	inBulkQSegs       uint64
	inDiscards        uint64
	connFloods        uint32
	connStalls        uint32
	cfgEphemDef       uint16
	ephemInUse        uint16
	ephemHiWater      uint16
	flags             byte
	_                 [1]byte
	ephemExhaust      uint32
	smcRCurrEstabLnks uint32
	smcRLnkActTimeOut uint32
	smcRActLnkOpened  uint32
	smcRPasLnkOpened  uint32
	smcRLnksClosed    uint32
	smcRCurrEstab     uint32
	smcRActiveOpened  uint32
	smcRPassiveOpened uint32
	smcRConnClosed    uint32
	smcRInSegs        uint64
	smcROutSegs       uint64
	smcRInRsts        uint32
	smcROutRsts       uint32
	smcDCurrEstabLnks uint32
	smcDActLnkOpened  uint32
	smcDPasLnkOpened  uint32
	smcDLnksClosed    uint32
	smcDCurrEstab     uint32
	smcDActiveOpened  uint32
	smcDPassiveOpened uint32
	smcDConnClosed    uint32
	smcDInSegs        uint64
	smcDOutSegs       uint64
	smcDInRsts        uint32
	smcDOutRsts       uint32
}

type nwmConnEntry struct {
	ident             uint32
	local             [28]byte
	remote            [28]byte
	startTime         [8]byte
	lastActivity      [8]byte
	bytesIn           [8]byte
	bytesOut          [8]byte
	inSegs            [8]byte
	outSegs           [8]byte
	state             uint16
	activeOpen        byte
	flag01            byte
	outBuffered       uint32
	inBuffered        uint32
	maxSndWnd         uint32
	reXmtCount        uint32
	congestionWnd     uint32
	ssThresh          uint32
	roundTripTime     uint32
	roundTripVar      uint32
	sendMSS           uint32
	sndWnd            uint32
	rcvBufSize        uint32
	sndBufSize        uint32
	outOfOrderCount   uint32
	lcl0WindowCount   uint32
	rmt0WindowCount   uint32
	dupacks           uint32
	flag02            byte
	sockOpt6Cont      byte
	asid              uint16
	resourceName      [8]byte
	resourceId        uint32
	subtask           uint32
	sockOpt           byte
	sockOpt6          byte
	clusterConnFlag   byte
	proto             byte
	targetAppl        [8]byte
	luName            [8]byte
	clientUserId      [8]byte
	logMode           [8]byte
	timeStamp         uint32
	timeStampAge      uint32
	serverResourceId  uint32
	intfName          [16]byte
	ttlsStatPol       byte
	ttlsStatConn      byte
	ttlsSSLProt       uint16
	ttlsNegCiph       [2]byte
	ttlsSecType       byte
	ttlsFIPS140Mode   byte
	ttlsUserID        [8]byte
	applData          [40]byte
	inOldestTime      [8]byte
	outOldestTime     [8]byte
	tcpTrustedPartner byte
	_                 [3]byte
	bulkDataIntfName  [16]byte
	ttlsNegCiph4      [4]byte
	smcReason         uint32
	lclSMCLinkId      uint32
	rmtSMCLinkId      uint32
	smcStatus         byte
	smcFlags          byte
	_                 [2]byte
	rcvWnd            uint32
	lclSMCBufSz       uint32
	rmtSMCBufSz       uint32
	ttlsSessID        [32]byte
	ttlsSessIDLen     int16
	_                 [1]byte
	smcDStatus        byte
	smcDReason        uint32
} // union of sockaddr4 and sockaddr6
// uint64

type PtraceRegsArm struct{ Uregs [18]uint32 } // PtraceRegsArm is the registers used by arm binaries.

type PtraceRegsArm64 struct {
	Regs   [31]uint64
	Sp     uint64
	Pc     uint64
	Pstate uint64
} // PtraceRegsArm64 is the registers used by arm64 binaries.

type PtraceRegsMips struct {
	Regs     [32]uint64
	Lo       uint64
	Hi       uint64
	Epc      uint64
	Badvaddr uint64
	Status   uint64
	Cause    uint64
} // PtraceRegsMips is the registers used by mips binaries.

type PtraceRegsMips64 struct {
	Regs     [32]uint64
	Lo       uint64
	Hi       uint64
	Epc      uint64
	Badvaddr uint64
	Status   uint64
	Cause    uint64
} // PtraceRegsMips64 is the registers used by mips64 binaries.

type PtraceRegsMipsle struct {
	Regs     [32]uint64
	Lo       uint64
	Hi       uint64
	Epc      uint64
	Badvaddr uint64
	Status   uint64
	Cause    uint64
} // PtraceRegsMipsle is the registers used by mipsle binaries.

type PtraceRegsMips64le struct {
	Regs     [32]uint64
	Lo       uint64
	Hi       uint64
	Epc      uint64
	Badvaddr uint64
	Status   uint64
	Cause    uint64
} // PtraceRegsMips64le is the registers used by mips64le binaries.

type PtraceRegs386 struct {
	Ebx      int32
	Ecx      int32
	Edx      int32
	Esi      int32
	Edi      int32
	Ebp      int32
	Eax      int32
	Xds      int32
	Xes      int32
	Xfs      int32
	Xgs      int32
	Orig_eax int32
	Eip      int32
	Xcs      int32
	Eflags   int32
	Esp      int32
	Xss      int32
} // PtraceRegs386 is the registers used by 386 binaries.

type PtraceRegsAmd64 struct {
	R15      uint64
	R14      uint64
	R13      uint64
	R12      uint64
	Rbp      uint64
	Rbx      uint64
	R11      uint64
	R10      uint64
	R9       uint64
	R8       uint64
	Rax      uint64
	Rcx      uint64
	Rdx      uint64
	Rsi      uint64
	Rdi      uint64
	Orig_rax uint64
	Rip      uint64
	Cs       uint64
	Eflags   uint64
	Rsp      uint64
	Ss       uint64
	Fs_base  uint64
	Gs_base  uint64
	Ds       uint64
	Es       uint64
	Fs       uint64
	Gs       uint64
} // PtraceRegsAmd64 is the registers used by amd64 binaries.

type mibentry struct {
	ctlname string
	ctloid  []_C_int
}

type Timeval32 struct {
	Sec  int32
	Usec int32
}

type Timex struct{}

type Tms struct{}

type Utimbuf struct {
	Actime  int32
	Modtime int32
}

type Timezone struct {
	Minuteswest int32
	Dsttime     int32
}

type Rusage struct {
	Utime    Timeval
	Stime    Timeval
	Maxrss   int32
	Ixrss    int32
	Idrss    int32
	Isrss    int32
	Minflt   int32
	Majflt   int32
	Nswap    int32
	Inblock  int32
	Oublock  int32
	Msgsnd   int32
	Msgrcv   int32
	Nsignals int32
	Nvcsw    int32
	Nivcsw   int32
}

type Rlimit struct {
	Cur uint64
	Max uint64
}

type Stat_t struct {
	Dev      uint32
	Ino      uint32
	Mode     uint32
	Nlink    int16
	Flag     uint16
	Uid      uint32
	Gid      uint32
	Rdev     uint32
	Size     int32
	Atim     Timespec
	Mtim     Timespec
	Ctim     Timespec
	Blksize  int32
	Blocks   int32
	Vfstype  int32
	Vfs      uint32
	Type     uint32
	Gen      uint32
	Reserved [9]uint32
}

type StatxTimestamp struct{}

type Statx_t struct{}

type Dirent struct {
	Offset uint32
	Ino    uint32
	Reclen uint16
	Namlen uint16
	Name   [256]uint8
}

type RawSockaddrInet4 struct {
	Len    uint8
	Family uint8
	Port   uint16
	Addr   [4]byte
	Zero   [8] /* in_addr */ uint8
}

type RawSockaddrInet6 struct {
	Len      uint8
	Family   uint8
	Port     uint16
	Flowinfo uint32
	Addr     [16]byte
	Scope_id uint32
} /* in6_addr */

type RawSockaddrUnix struct {
	Len    uint8
	Family uint8
	Path   [1023]uint8
}

type RawSockaddrDatalink struct {
	Len    uint8
	Family uint8
	Index  uint16
	Type   uint8
	Nlen   uint8
	Alen   uint8
	Slen   uint8
	Data   [120]uint8
}

type RawSockaddr struct {
	Len    uint8
	Family uint8
	Data   [14]uint8
}

type RawSockaddrAny struct {
	Addr RawSockaddr
	Pad  [1012]uint8
}

type Cmsghdr struct {
	Len   uint32
	Level int32
	Type  int32
}

type ICMPv6Filter struct{ Filt [8]uint32 }

type Iovec struct {
	Base *byte
	Len  uint32
}

type IPMreq struct {
	Multiaddr [4]byte
	Interface [4] /* in_addr */ byte
} /* in_addr */

type IPv6Mreq struct {
	Multiaddr [16]byte
	Interface uint32
} /* in6_addr */

type IPv6MTUInfo struct {
	Addr RawSockaddrInet6
	Mtu  uint32
}

type Linger struct {
	Onoff  int32
	Linger int32
}

type Msghdr struct {
	Name       *byte
	Namelen    uint32
	Iov        *Iovec
	Iovlen     int32
	Control    *byte
	Controllen uint32
	Flags      int32
}

type IfMsgHdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Addrs   int32
	Flags   int32
	Index   uint16
	Addrlen uint8
	_       [1]byte
}

type FdSet struct{ Bits [2048]int32 }

type Utsname struct {
	Sysname  [32]byte
	Nodename [32]byte
	Release  [32]byte
	Version  [32]byte
	Machine  [32]byte
}

type Ustat_t struct{}

type Sigset_t struct {
	Losigs uint32
	Hisigs uint32
}

type Termio struct {
	Iflag uint16
	Oflag uint16
	Cflag uint16
	Lflag uint16
	Line  uint8
	Cc    [8]uint8
	_     [1]byte
}

type PollFd struct {
	Fd      int32
	Events  uint16
	Revents uint16
}

type Flock_t struct {
	Type   int16
	Whence int16
	Sysid  uint32
	Pid    int32
	Vfs    int32
	Start  int64
	Len    int64
}

type Fsid_t struct{ Val [2]uint32 }

type Fsid64_t struct{ Val [2]uint64 }

type Statfs_t struct {
	Version   int32
	Type      int32
	Bsize     uint32
	Blocks    uint32
	Bfree     uint32
	Bavail    uint32
	Files     uint32
	Ffree     uint32
	Fsid      Fsid_t
	Vfstype   int32
	Fsize     uint32
	Vfsnumber int32
	Vfsoff    int32
	Vfslen    int32
	Vfsvers   int32
	Fname     [32]uint8
	Fpack     [32]uint8
	Name_max  int32
}

type Fstore_t struct {
	Flags      uint32
	Posmode    int32
	Offset     int64
	Length     int64
	Bytesalloc int64
}

type Radvisory_t struct {
	Offset int64
	Count  int32
	_      [4]byte
}

type Fbootstraptransfer_t struct {
	Offset int64
	Length uint64
	Buffer *byte
}

type Log2phys_t struct {
	Flags uint32
	_     [16]byte
}

type Fsid struct{ Val [2]int32 }

type Attrlist struct {
	Bitmapcount uint16
	Reserved    uint16
	Commonattr  uint32
	Volattr     uint32
	Dirattr     uint32
	Fileattr    uint32
	Forkattr    uint32
}

type RawSockaddrCtl struct {
	Sc_len      uint8
	Sc_family   uint8
	Ss_sysaddr  uint16
	Sc_id       uint32
	Sc_unit     uint32
	Sc_reserved [5]uint32
}

type RawSockaddrVM struct {
	Len       uint8
	Family    uint8
	Reserved1 uint16
	Port      uint32
	Cid       uint32
}

type XVSockPCB struct {
	Xv_len           uint32
	Xv_vsockpp       uint64
	Xvp_local_cid    uint32
	Xvp_local_port   uint32
	Xvp_remote_cid   uint32
	Xvp_remote_port  uint32
	Xvp_rxcnt        uint32
	Xvp_txcnt        uint32
	Xvp_peer_rxhiwat uint32
	Xvp_peer_rxcnt   uint32
	Xvp_last_pid     int32
	Xvp_gencnt       uint64
	Xv_socket        XSocket
	_                [4]byte
}

type XSocket struct {
	Xso_len      uint32
	Xso_so       uint32
	So_type      int16
	So_options   int16
	So_linger    int16
	So_state     int16
	So_pcb       uint32
	Xso_protocol int32
	Xso_family   int32
	So_qlen      int16
	So_incqlen   int16
	So_qlimit    int16
	So_timeo     int16
	So_error     uint16
	So_pgid      int32
	So_oobmark   uint32
	So_rcv       XSockbuf
	So_snd       XSockbuf
	So_uid       uint32
}

type XSocket64 struct {
	Xso_len      uint32
	_            [8]byte
	So_type      int16
	So_options   int16
	So_linger    int16
	So_state     int16
	_            [8]byte
	Xso_protocol int32
	Xso_family   int32
	So_qlen      int16
	So_incqlen   int16
	So_qlimit    int16
	So_timeo     int16
	So_error     uint16
	So_pgid      int32
	So_oobmark   uint32
	So_rcv       XSockbuf
	So_snd       XSockbuf
	So_uid       uint32
}

type XSockbuf struct {
	Cc    uint32
	Hiwat uint32
	Mbcnt uint32
	Mbmax uint32
	Lowat int32
	Flags int16
	Timeo int16
}

type XVSockPgen struct {
	Len   uint32
	Count uint64
	Gen   uint64
	Sogen uint64
}

type SaEndpoints struct {
	Srcif      uint32
	Srcaddr    *RawSockaddr
	Srcaddrlen uint32
	Dstaddr    *RawSockaddr
	Dstaddrlen uint32
	_          [4]byte
}

type Xucred struct {
	Version uint32
	Uid     uint32
	Ngroups int16
	Groups  [16]uint32
}

type IPMreqn struct {
	Multiaddr [4]byte
	Address   [4] /* in_addr */ byte
	Ifindex   int32
} /* in_addr */

type Inet4Pktinfo struct {
	Ifindex  uint32
	Spec_dst [4]byte
	Addr     [4] /* in_addr */ byte
} /* in_addr */

type Inet6Pktinfo struct {
	Addr    [16]byte
	Ifindex uint32
} /* in6_addr */

type TCPConnectionInfo struct {
	State               uint8
	Snd_wscale          uint8
	Rcv_wscale          uint8
	_                   uint8
	Options             uint32
	Flags               uint32
	Rto                 uint32
	Maxseg              uint32
	Snd_ssthresh        uint32
	Snd_cwnd            uint32
	Snd_wnd             uint32
	Snd_sbbytes         uint32
	Rcv_wnd             uint32
	Rttcur              uint32
	Srtt                uint32
	Rttvar              uint32
	Txpackets           uint64
	Txbytes             uint64
	Txretransmitbytes   uint64
	Rxpackets           uint64
	Rxbytes             uint64
	Rxoutoforderbytes   uint64
	Txretransmitpackets uint64
}

type Kevent_t struct {
	Ident  uint64
	Filter int16
	Flags  uint16
	Fflags uint32
	Data   int64
	Udata  *byte
}

type IfMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Addrs   int32
	Flags   int32
	Index   uint16
	Data    IfData
}

type IfMsghdr2 struct {
	Msglen     uint16
	Version    uint8
	Type       uint8
	Addrs      int32
	Flags      int32
	Index      uint16
	Snd_len    int32
	Snd_maxlen int32
	Snd_drops  int32
	Timer      int32
	Data       IfData64
}

type IfData struct {
	Type       uint8
	Typelen    uint8
	Physical   uint8
	Addrlen    uint8
	Hdrlen     uint8
	Recvquota  uint8
	Xmitquota  uint8
	Unused1    uint8
	Mtu        uint32
	Metric     uint32
	Baudrate   uint32
	Ipackets   uint32
	Ierrors    uint32
	Opackets   uint32
	Oerrors    uint32
	Collisions uint32
	Ibytes     uint32
	Obytes     uint32
	Imcasts    uint32
	Omcasts    uint32
	Iqdrops    uint32
	Noproto    uint32
	Recvtiming uint32
	Xmittiming uint32
	Lastchange Timeval32
	Unused2    uint32
	Hwassist   uint32
	Reserved1  uint32
	Reserved2  uint32
}

type IfData64 struct {
	Type       uint8
	Typelen    uint8
	Physical   uint8
	Addrlen    uint8
	Hdrlen     uint8
	Recvquota  uint8
	Xmitquota  uint8
	Unused1    uint8
	Mtu        uint32
	Metric     uint32
	Baudrate   uint64
	Ipackets   uint64
	Ierrors    uint64
	Opackets   uint64
	Oerrors    uint64
	Collisions uint64
	Ibytes     uint64
	Obytes     uint64
	Imcasts    uint64
	Omcasts    uint64
	Iqdrops    uint64
	Noproto    uint64
	Recvtiming uint32
	Xmittiming uint32
	Lastchange Timeval32
}

type IfaMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Addrs   int32
	Flags   int32
	Index   uint16
	Metric  int32
}

type IfmaMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Addrs   int32
	Flags   int32
	Index   uint16
	_       [2]byte
}

type IfmaMsghdr2 struct {
	Msglen   uint16
	Version  uint8
	Type     uint8
	Addrs    int32
	Flags    int32
	Index    uint16
	Refcount int32
}

type RtMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Index   uint16
	Flags   int32
	Addrs   int32
	Pid     int32
	Seq     int32
	Errno   int32
	Use     int32
	Inits   uint32
	Rmx     RtMetrics
}

type RtMsghdr2 struct {
	Msglen      uint16
	Version     uint8
	Type        uint8
	Index       uint16
	Flags       int32
	Addrs       int32
	Refcnt      int32
	Parentflags int32
	Reserved    int32
	Use         int32
	Inits       uint32
	Rmx         RtMetrics
}

type RtMetrics struct {
	Locks    uint32
	Mtu      uint32
	Hopcount uint32
	Expire   int32
	Recvpipe uint32
	Sendpipe uint32
	Ssthresh uint32
	Rtt      uint32
	Rttvar   uint32
	Pksent   uint32
	State    uint32
	Filler   [3]uint32
}

type BpfVersion struct {
	Major uint16
	Minor uint16
}

type BpfStat struct {
	Recv uint32
	Drop uint32
}

type BpfProgram struct {
	Len   uint32
	Insns *BpfInsn
}

type BpfInsn struct {
	Code uint16
	Jt   uint8
	Jf   uint8
	K    uint32
}

type BpfHdr struct {
	Tstamp  Timeval32
	Caplen  uint32
	Datalen uint32
	Hdrlen  uint16
	_       [2]byte
}

type Clockinfo struct {
	Hz      int32
	Tick    int32
	Tickadj int32
	Stathz  int32
	Profhz  int32
}

type CtlInfo struct {
	Id   uint32
	Name [96]byte
}

type Eproc struct {
	Paddr   uintptr
	Sess    uintptr
	Pcred   Pcred
	Ucred   Ucred
	Vm      Vmspace
	Ppid    int32
	Pgid    int32
	Jobc    int16
	Tdev    int32
	Tpgid   int32
	Tsess   uintptr
	Wmesg   [8]byte
	Xsize   int32
	Xrssize int16
	Xccount int16
	Xswrss  int16
	Flag    int32
	Login   [12]byte
	Spare   [4]int32
	_       [4]byte
}

type ExternProc struct {
	P_starttime Timeval
	P_vmspace   *Vmspace
	P_sigacts   uintptr
	P_flag      int32
	P_stat      int8
	P_pid       int32
	P_oppid     int32
	P_dupfd     int32
	User_stack  *int8
	Exit_thread *byte
	P_debugger  int32
	Sigwait     int32
	P_estcpu    uint32
	P_cpticks   int32
	P_pctcpu    uint32
	P_wchan     *byte
	P_wmesg     *int8
	P_swtime    uint32
	P_slptime   uint32
	P_realtimer Itimerval
	P_rtime     Timeval
	P_uticks    uint64
	P_sticks    uint64
	P_iticks    uint64
	P_traceflag int32
	P_tracep    uintptr
	P_siglist   int32
	P_textvp    uintptr
	P_holdcnt   int32
	P_sigmask   uint32
	P_sigignore uint32
	P_sigcatch  uint32
	P_priority  uint8
	P_usrpri    uint8
	P_nice      int8
	P_comm      [17]byte
	P_pgrp      uintptr
	P_addr      uintptr
	P_xstat     uint16
	P_acflag    uint16
	P_ru        *Rusage
}

type Itimerval struct {
	Interval Timeval
	Value    Timeval
}

type KinfoProc struct {
	Proc  ExternProc
	Eproc Eproc
}

type Vmspace struct {
	Dummy  int32
	Dummy2 *int8
	Dummy3 [5]int32
	Dummy4 [3]*int8
}

type Pcred struct {
	Pc_lock  [72]int8
	Pc_ucred uintptr
	P_ruid   uint32
	P_svuid  uint32
	P_rgid   uint32
	P_svgid  uint32
	P_refcnt int32
	_        [4]byte
}

type SysvIpcPerm struct {
	Uid  uint32
	Gid  uint32
	Cuid uint32
	Cgid uint32
	Mode uint16
	_    uint16
	_    int32
}

type SysvShmDesc struct {
	Perm   SysvIpcPerm
	Segsz  uint64
	Lpid   int32
	Cpid   int32
	Nattch uint16
	_      [34]byte
}

type IfAnnounceMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Index   uint16
	Name    [16]int8
	What    uint16
}

type PtraceLwpInfoStruct struct {
	Lwpid        int32
	Event        int32
	Flags        int32
	Sigmask      Sigset_t
	Siglist      Sigset_t
	Siginfo      __PtraceSiginfo
	Tdname       [20]int8
	Child_pid    int32
	Syscall_code uint32
	Syscall_narg uint32
}

type __Siginfo struct {
	Signo  int32
	Errno  int32
	Code   int32
	Pid    int32
	Uid    uint32
	Status int32
	Addr   *byte
	Value  [4]byte
	_      [32]byte
}

type __PtraceSiginfo struct {
	Signo  int32
	Errno  int32
	Code   int32
	Pid    int32
	Uid    uint32
	Status int32
	Addr   uintptr
	Value  [4]byte
	_      [32]byte
}

type FpReg struct {
	Env   [7]uint32
	Acc   [8][10]uint8
	Ex_sw uint32
	Pad   [64]uint8
}

type FpExtendedPrecision struct{}

type PtraceIoDesc struct {
	Op   int32
	Offs uintptr
	Addr *byte
	Len  uint32
}

type ifMsghdr struct {
	Msglen  uint16
	Version uint8
	Type    uint8
	Addrs   int32
	Flags   int32
	Index   uint16
	_       uint16
	Data    ifData
}

type ifData struct {
	Type       uint8
	Physical   uint8
	Addrlen    uint8
	Hdrlen     uint8
	Link_state uint8
	Vhid       uint8
	Datalen    uint16
	Mtu        uint32
	Metric     uint32
	Baudrate   uint64
	Ipackets   uint64
	Ierrors    uint64
	Opackets   uint64
	Oerrors    uint64
	Collisions uint64
	Ibytes     uint64
	Obytes     uint64
	Imcasts    uint64
	Omcasts    uint64
	Iqdrops    uint64
	Oqdrops    uint64
	Noproto    uint64
	Hwassist   uint64
	_          [8]byte
	_          [16]byte
}

type BpfZbuf struct {
	Bufa   *byte
	Bufb   *byte
	Buflen uint32
}

type BpfZbufHeader struct {
	Kernel_gen uint32
	Kernel_len uint32
	User_gen   uint32
	_          [5]uint32
}

type CapRights struct{ Rights [2]uint64 }

type ItimerSpec struct {
	Interval Timespec
	Value    Timespec
}

type FileCloneRange struct {
	Src_fd      int64
	Src_offset  uint64
	Src_length  uint64
	Dest_offset uint64
}

type RawFileDedupeRange struct {
	Src_offset uint64
	Src_length uint64
	Dest_count uint16
	Reserved1  uint16
	Reserved2  uint32
}

type RawFileDedupeRangeInfo struct {
	Dest_fd       int64
	Dest_offset   uint64
	Bytes_deduped uint64
	Status        int32
	Reserved      uint32
}

type FscryptPolicy struct {
	Version                   uint8
	Contents_encryption_mode  uint8
	Filenames_encryption_mode uint8
	Flags                     uint8
	Master_key_descriptor     [8]uint8
}

type FscryptKey struct {
	Mode uint32
	Raw  [64]uint8
	Size uint32
}

type FscryptPolicyV1 struct {
	Version                   uint8
	Contents_encryption_mode  uint8
	Filenames_encryption_mode uint8
	Flags                     uint8
	Master_key_descriptor     [8]uint8
}

type FscryptPolicyV2 struct {
	Version                   uint8
	Contents_encryption_mode  uint8
	Filenames_encryption_mode uint8
	Flags                     uint8
	Log2_data_unit_size       uint8
	_                         [3]uint8
	Master_key_identifier     [16]uint8
}

type FscryptGetPolicyExArg struct {
	Size   uint64
	Policy [24]byte
}

type FscryptKeySpecifier struct {
	Type uint32
	_    uint32
	U    [32]byte
}

type FscryptAddKeyArg struct {
	Key_spec FscryptKeySpecifier
	Raw_size uint32
	Key_id   uint32
	_        [8]uint32
}

type FscryptRemoveKeyArg struct {
	Key_spec             FscryptKeySpecifier
	Removal_status_flags uint32
	_                    [5]uint32
}

type FscryptGetKeyStatusArg struct {
	Key_spec     FscryptKeySpecifier
	_            [6]uint32
	Status       uint32
	Status_flags uint32
	User_count   uint32
	_            [13]uint32
}

type DmIoctl struct {
	Version      [3]uint32
	Data_size    uint32
	Data_start   uint32
	Target_count uint32
	Open_count   int32
	Flags        uint32
	Event_nr     uint32
	_            uint32
	Dev          uint64
	Name         [128]byte
	Uuid         [129]byte
	Data         [7]byte
}

type DmTargetSpec struct {
	Sector_start uint64
	Length       uint64
	Status       int32
	Next         uint32
	Target_type  [16]byte
}

type DmTargetDeps struct {
	Count uint32
	_     uint32
}

type DmTargetVersions struct {
	Next    uint32
	Version [3]uint32
}

type DmTargetMsg struct{ Sector uint64 }

type KeyctlDHParams struct {
	Private int32
	Prime   int32
	Base    int32
}

type RawSockaddrLinklayer struct {
	Family   uint16
	Protocol uint16
	Ifindex  int32
	Hatype   uint16
	Pkttype  uint8
	Halen    uint8
	Addr     [8]uint8
}

type RawSockaddrNetlink struct {
	Family uint16
	Pad    uint16
	Pid    uint32
	Groups uint32
}

type RawSockaddrHCI struct {
	Family  uint16
	Dev     uint16
	Channel uint16
}

type RawSockaddrL2 struct {
	Family      uint16
	Psm         uint16
	Bdaddr      [6]uint8
	Cid         uint16
	Bdaddr_type uint8
	_           [1]byte
}

type RawSockaddrRFCOMM struct {
	Family  uint16
	Bdaddr  [6]uint8
	Channel uint8
	_       [1]byte
}

type RawSockaddrCAN struct {
	Family  uint16
	Ifindex int32
	Addr    [16]byte
}

type RawSockaddrALG struct {
	Family uint16
	Type   [14]uint8
	Feat   uint32
	Mask   uint32
	Name   [64]uint8
}

type RawSockaddrXDP struct {
	Family         uint16
	Flags          uint16
	Ifindex        uint32
	Queue_id       uint32
	Shared_umem_fd uint32
}

type RawSockaddrTIPC struct {
	Family   uint16
	Addrtype uint8
	Scope    int8
	Addr     [12]byte
}

type RawSockaddrL2TPIP struct {
	Family  uint16
	Unused  uint16
	Addr    [4]byte
	Conn_id uint32
	_       [4] /* in_addr */ uint8
}

type RawSockaddrL2TPIP6 struct {
	Family   uint16
	Unused   uint16
	Flowinfo uint32
	Addr     [16]byte
	Scope_id uint32
	Conn_id  uint32
} /* in6_addr */

type RawSockaddrIUCV struct {
	Family  uint16
	Port    uint16
	Addr    uint32
	Nodeid  [8]int8
	User_id [8]int8
	Name    [8]int8
}

type RawSockaddrNFC struct {
	Sa_family    uint16
	Dev_idx      uint32
	Target_idx   uint32
	Nfc_protocol uint32
}

type PacketMreq struct {
	Ifindex int32
	Type    uint16
	Alen    uint16
	Address [8]uint8
}

type TCPInfo struct {
	State                uint8
	Ca_state             uint8
	Retransmits          uint8
	Probes               uint8
	Backoff              uint8
	Options              uint8
	Rto                  uint32
	Ato                  uint32
	Snd_mss              uint32
	Rcv_mss              uint32
	Unacked              uint32
	Sacked               uint32
	Lost                 uint32
	Retrans              uint32
	Fackets              uint32
	Last_data_sent       uint32
	Last_ack_sent        uint32
	Last_data_recv       uint32
	Last_ack_recv        uint32
	Pmtu                 uint32
	Rcv_ssthresh         uint32
	Rtt                  uint32
	Rttvar               uint32
	Snd_ssthresh         uint32
	Snd_cwnd             uint32
	Advmss               uint32
	Reordering           uint32
	Rcv_rtt              uint32
	Rcv_space            uint32
	Total_retrans        uint32
	Pacing_rate          uint64
	Max_pacing_rate      uint64
	Bytes_acked          uint64
	Bytes_received       uint64
	Segs_out             uint32
	Segs_in              uint32
	Notsent_bytes        uint32
	Min_rtt              uint32
	Data_segs_in         uint32
	Data_segs_out        uint32
	Delivery_rate        uint64
	Busy_time            uint64
	Rwnd_limited         uint64
	Sndbuf_limited       uint64
	Delivered            uint32
	Delivered_ce         uint32
	Bytes_sent           uint64
	Bytes_retrans        uint64
	Dsack_dups           uint32
	Reord_seen           uint32
	Rcv_ooopack          uint32
	Snd_wnd              uint32
	Rcv_wnd              uint32
	Rehash               uint32
	Total_rto            uint16
	Total_rto_recoveries uint16
	Total_rto_time       uint32
}

type TCPVegasInfo struct {
	Enabled uint32
	Rttcnt  uint32
	Rtt     uint32
	Minrtt  uint32
}

type TCPDCTCPInfo struct {
	Enabled  uint16
	Ce_state uint16
	Alpha    uint32
	Ab_ecn   uint32
	Ab_tot   uint32
}

type TCPBBRInfo struct {
	Bw_lo       uint32
	Bw_hi       uint32
	Min_rtt     uint32
	Pacing_gain uint32
	Cwnd_gain   uint32
}

type CanFilter struct {
	Id   uint32
	Mask uint32
}

type TCPRepairOpt struct {
	Code uint32
	Val  uint32
}

type NlMsghdr struct {
	Len   uint32
	Type  uint16
	Flags uint16
	Seq   uint32
	Pid   uint32
}

type NlMsgerr struct {
	Error int32
	Msg   NlMsghdr
}

type RtGenmsg struct{ Family uint8 }

type NlAttr struct {
	Len  uint16
	Type uint16
}

type RtAttr struct {
	Len  uint16
	Type uint16
}

type IfInfomsg struct {
	Family uint8
	_      uint8
	Type   uint16
	Index  int32
	Flags  uint32
	Change uint32
}

type IfAddrmsg struct {
	Family    uint8
	Prefixlen uint8
	Flags     uint8
	Scope     uint8
	Index     uint32
}

type IfaCacheinfo struct {
	Prefered uint32
	Valid    uint32
	Cstamp   uint32
	Tstamp   uint32
}

type RtMsg struct {
	Family   uint8
	Dst_len  uint8
	Src_len  uint8
	Tos      uint8
	Table    uint8
	Protocol uint8
	Scope    uint8
	Type     uint8
	Flags    uint32
}

type RtNexthop struct {
	Len     uint16
	Flags   uint8
	Hops    uint8
	Ifindex int32
}

type NdUseroptmsg struct {
	Family    uint8
	Pad1      uint8
	Opts_len  uint16
	Ifindex   int32
	Icmp_type uint8
	Icmp_code uint8
	Pad2      uint16
	Pad3      uint32
}

type NdMsg struct {
	Family  uint8
	Pad1    uint8
	Pad2    uint16
	Ifindex int32
	State   uint16
	Flags   uint8
	Type    uint8
}

type SockFilter struct {
	Code uint16
	Jt   uint8
	Jf   uint8
	K    uint32
}

type SockFprog struct {
	Len    uint16
	Filter *SockFilter
}

type InotifyEvent struct {
	Wd     int32
	Mask   uint32
	Cookie uint32
	Len    uint32
}

type OpenHow struct {
	Flags   uint64
	Mode    uint64
	Resolve uint64
}

type sigset_argpack struct {
	ss    *Sigset_t
	ssLen uintptr
}

type SignalfdSiginfo struct {
	Signo     uint32
	Errno     int32
	Code      int32
	Pid       uint32
	Uid       uint32
	Fd        int32
	Tid       uint32
	Band      uint32
	Overrun   uint32
	Trapno    uint32
	Status    int32
	Int       int32
	Ptr       uint64
	Utime     uint64
	Stime     uint64
	Addr      uint64
	Addr_lsb  uint16
	_         uint16
	Syscall   int32
	Call_addr uint64
	Arch      uint32
	_         [28]uint8
}

type CGroupStats struct {
	Sleeping        uint64
	Running         uint64
	Stopped         uint64
	Uninterruptible uint64
	Io_wait         uint64
}

type Genlmsghdr struct {
	Cmd      uint8
	Version  uint8
	Reserved uint16
}

type PerfEventAttr struct {
	Type               uint32
	Size               uint32
	Config             uint64
	Sample             uint64
	Sample_type        uint64
	Read_format        uint64
	Bits               uint64
	Wakeup             uint32
	Bp_type            uint32
	Ext1               uint64
	Ext2               uint64
	Branch_sample_type uint64
	Sample_regs_user   uint64
	Sample_stack_user  uint32
	Clockid            int32
	Sample_regs_intr   uint64
	Aux_watermark      uint32
	Sample_max_stack   uint16
	_                  uint16
	Aux_sample_size    uint32
	_                  uint32
	Sig_data           uint64
}

type PerfEventMmapPage struct {
	Version        uint32
	Compat_version uint32
	Lock           uint32
	Index          uint32
	Offset         int64
	Time_enabled   uint64
	Time_running   uint64
	Capabilities   uint64
	Pmc_width      uint16
	Time_shift     uint16
	Time_mult      uint32
	Time_offset    uint64
	Time_zero      uint64
	Size           uint32
	_              uint32
	Time_cycles    uint64
	Time_mask      uint64
	_              [928]uint8
	Data_head      uint64
	Data_tail      uint64
	Data_offset    uint64
	Data_size      uint64
	Aux_head       uint64
	Aux_tail       uint64
	Aux_offset     uint64
	Aux_size       uint64
}

type TCPMD5Sig struct {
	Addr      SockaddrStorage
	Flags     uint8
	Prefixlen uint8
	Keylen    uint16
	Ifindex   int32
	Key       [80]uint8
}

type HDDriveCmdHdr struct {
	Command uint8
	Number  uint8
	Feature uint8
	Count   uint8
}

type HDDriveID struct {
	Config         uint16
	Cyls           uint16
	Reserved2      uint16
	Heads          uint16
	Track_bytes    uint16
	Sector_bytes   uint16
	Sectors        uint16
	Vendor0        uint16
	Vendor1        uint16
	Vendor2        uint16
	Serial_no      [20]uint8
	Buf_type       uint16
	Buf_size       uint16
	Ecc_bytes      uint16
	Fw_rev         [8]uint8
	Model          [40]uint8
	Max_multsect   uint8
	Vendor3        uint8
	Dword_io       uint16
	Vendor4        uint8
	Capability     uint8
	Reserved50     uint16
	Vendor5        uint8
	TPIO           uint8
	Vendor6        uint8
	TDMA           uint8
	Field_valid    uint16
	Cur_cyls       uint16
	Cur_heads      uint16
	Cur_sectors    uint16
	Cur_capacity0  uint16
	Cur_capacity1  uint16
	Multsect       uint8
	Multsect_valid uint8
	Lba_capacity   uint32
	Dma_1word      uint16
	Dma_mword      uint16
	Eide_pio_modes uint16
	Eide_dma_min   uint16
	Eide_dma_time  uint16
	Eide_pio       uint16
	Eide_pio_iordy uint16
	Words69_70     [2]uint16
	Words71_74     [4]uint16
	Queue_depth    uint16
	Words76_79     [4]uint16
	Major_rev_num  uint16
	Minor_rev_num  uint16
	Command_set_1  uint16
	Command_set_2  uint16
	Cfsse          uint16
	Cfs_enable_1   uint16
	Cfs_enable_2   uint16
	Csf_default    uint16
	Dma_ultra      uint16
	Trseuc         uint16
	TrsEuc         uint16
	CurAPMvalues   uint16
	Mprc           uint16
	Hw_config      uint16
	Acoustic       uint16
	Msrqs          uint16
	Sxfert         uint16
	Sal            uint16
	Spg            uint32
	Lba_capacity_2 uint64
	Words104_125   [22]uint16
	Last_lun       uint16
	Word127        uint16
	Dlf            uint16
	Csfo           uint16
	Words130_155   [26]uint16
	Word156        uint16
	Words157_159   [3]uint16
	Cfa_power      uint16
	Words161_175   [15]uint16
	Words176_205   [30]uint16
	Words206_254   [49]uint16
	Integrity_word uint16
}

type Tpacket2Hdr struct {
	Status    uint32
	Len       uint32
	Snaplen   uint32
	Mac       uint16
	Net       uint16
	Sec       uint32
	Nsec      uint32
	Vlan_tci  uint16
	Vlan_tpid uint16
	_         [4]uint8
}

type Tpacket3Hdr struct {
	Next_offset uint32
	Sec         uint32
	Nsec        uint32
	Snaplen     uint32
	Len         uint32
	Status      uint32
	Mac         uint16
	Net         uint16
	Hv1         TpacketHdrVariant1
	_           [8]uint8
}

type TpacketHdrVariant1 struct {
	Rxhash    uint32
	Vlan_tci  uint32
	Vlan_tpid uint16
	_         uint16
}

type TpacketBlockDesc struct {
	Version uint32
	To_priv uint32
	Hdr     [40]byte
}

type TpacketBDTS struct {
	Sec  uint32
	Usec uint32
}

type TpacketHdrV1 struct {
	Block_status        uint32
	Num_pkts            uint32
	Offset_to_first_pkt uint32
	Blk_len             uint32
	Seq_num             uint64
	Ts_first_pkt        TpacketBDTS
	Ts_last_pkt         TpacketBDTS
}

type TpacketReq struct {
	Block_size uint32
	Block_nr   uint32
	Frame_size uint32
	Frame_nr   uint32
}

type TpacketReq3 struct {
	Block_size       uint32
	Block_nr         uint32
	Frame_size       uint32
	Frame_nr         uint32
	Retire_blk_tov   uint32
	Sizeof_priv      uint32
	Feature_req_word uint32
}

type TpacketStats struct {
	Packets uint32
	Drops   uint32
}

type TpacketStatsV3 struct {
	Packets      uint32
	Drops        uint32
	Freeze_q_cnt uint32
}

type TpacketAuxdata struct {
	Status    uint32
	Len       uint32
	Snaplen   uint32
	Mac       uint16
	Net       uint16
	Vlan_tci  uint16
	Vlan_tpid uint16
}

type Nfgenmsg struct {
	Nfgen_family uint8
	Version      uint8
	Res_id       uint16
}

type RTCTime struct {
	Sec   int32
	Min   int32
	Hour  int32
	Mday  int32
	Mon   int32
	Year  int32
	Wday  int32
	Yday  int32
	Isdst int32
}

type RTCWkAlrm struct {
	Enabled uint8
	Pending uint8
	Time    RTCTime
}

type BlkpgIoctlArg struct {
	Op      int32
	Flags   int32
	Datalen int32
	Data    *byte
}

type XDPRingOffset struct {
	Producer uint64
	Consumer uint64
	Desc     uint64
	Flags    uint64
}

type XDPMmapOffsets struct {
	Rx XDPRingOffset
	Tx XDPRingOffset
	Fr XDPRingOffset
	Cr XDPRingOffset
}

type XDPUmemReg struct {
	Addr            uint64
	Len             uint64
	Size            uint32
	Headroom        uint32
	Flags           uint32
	Tx_metadata_len uint32
}

type XDPStatistics struct {
	Rx_dropped               uint64
	Rx_invalid_descs         uint64
	Tx_invalid_descs         uint64
	Rx_ring_full             uint64
	Rx_fill_ring_empty_descs uint64
	Tx_ring_empty_descs      uint64
}

type XDPDesc struct {
	Addr    uint64
	Len     uint32
	Options uint32
}

type ScmTimestamping struct{ Ts [3]Timespec }

type SockExtendedErr struct {
	Errno  uint32
	Origin uint8
	Type   uint8
	Code   uint8
	Pad    uint8
	Info   uint32
	Data   uint32
}

type FanotifyEventMetadata struct {
	Event_len    uint32
	Vers         uint8
	Reserved     uint8
	Metadata_len uint16
	Mask         uint64
	Fd           int32
	Pid          int32
}

type FanotifyResponse struct {
	Fd       int32
	Response uint32
}

type CapUserHeader struct {
	Version uint32
	Pid     int32
}

type CapUserData struct {
	Effective   uint32
	Permitted   uint32
	Inheritable uint32
}

type LoopInfo64 struct {
	Device           uint64
	Inode            uint64
	Rdevice          uint64
	Offset           uint64
	Sizelimit        uint64
	Number           uint32
	Encrypt_type     uint32
	Encrypt_key_size uint32
	Flags            uint32
	File_name        [64]uint8
	Crypt_name       [64]uint8
	Encrypt_key      [32]uint8
	Init             [2]uint64
}

type LoopConfig struct {
	Fd   uint32
	Size uint32
	Info LoopInfo64
	_    [8]uint64
}

type TIPCSocketAddr struct {
	Ref  uint32
	Node uint32
}

type TIPCServiceRange struct {
	Type  uint32
	Lower uint32
	Upper uint32
}

type TIPCServiceName struct {
	Type     uint32
	Instance uint32
	Domain   uint32
}

type TIPCEvent struct {
	Event uint32
	Lower uint32
	Upper uint32
	Port  TIPCSocketAddr
	S     TIPCSubscr
}

type TIPCGroupReq struct {
	Type     uint32
	Instance uint32
	Scope    uint32
	Flags    uint32
}

type FsverityDigest struct {
	Algorithm uint16
	Size      uint16
}

type FsverityEnableArg struct {
	Version        uint32
	Hash_algorithm uint32
	Block_size     uint32
	Salt_size      uint32
	Salt_ptr       uint64
	Sig_size       uint32
	_              uint32
	Sig_ptr        uint64
	_              [11]uint64
}

type Nhmsg struct {
	Family   uint8
	Scope    uint8
	Protocol uint8
	Resvd    uint8
	Flags    uint32
}

type NexthopGrp struct {
	Id     uint32
	Weight uint8
	High   uint8
	Resvd2 uint16
}

type WatchdogInfo struct {
	Options  uint32
	Version  uint32
	Identity [32]uint8
}

type PPSFData struct {
	Info    PPSKInfo
	Timeout PPSKTime
}

type PPSKParams struct {
	Api_version   int32
	Mode          int32
	Assert_off_tu PPSKTime
	Clear_off_tu  PPSKTime
}

type PPSKTime struct {
	Sec   int64
	Nsec  int32
	Flags uint32
}

type EthtoolDrvinfo struct {
	Cmd          uint32
	Driver       [32]byte
	Version      [32]byte
	Fw_version   [32]byte
	Bus_info     [32]byte
	Erom_version [32]byte
	Reserved2    [12]byte
	N_priv_flags uint32
	N_stats      uint32
	Testinfo_len uint32
	Eedump_len   uint32
	Regdump_len  uint32
}

type EthtoolTsInfo struct {
	Cmd             uint32
	So_timestamping uint32
	Phc_index       int32
	Tx_types        uint32
	Tx_reserved     [3]uint32
	Rx_filters      uint32
	Rx_reserved     [3]uint32
}

type HwTstampConfig struct {
	Flags     int32
	Tx_type   int32
	Rx_filter int32
}

type (
	PtpClockCaps struct {
		Max_adj            int32
		N_alarm            int32
		N_ext_ts           int32
		N_per_out          int32
		Pps                int32
		N_pins             int32
		Cross_timestamping int32
		Adjust_phase       int32
		Max_phase_adj      int32
		Rsv                [11]int32
	}
	PtpClockTime struct {
		Sec      int64
		Nsec     uint32
		Reserved uint32
	}
	PtpExttsEvent struct {
		T     PtpClockTime
		Index uint32
		Flags uint32
		Rsv   [2]uint32
	}
	PtpExttsRequest struct {
		Index uint32
		Flags uint32
		Rsv   [2]uint32
	}
	PtpPeroutRequest struct {
		StartOrPhase PtpClockTime
		Period       PtpClockTime
		Index        uint32
		Flags        uint32
		On           PtpClockTime
	}
	PtpPinDesc struct {
		Name  [64]byte
		Index uint32
		Func  uint32
		Chan  uint32
		Rsv   [5]uint32
	}
	PtpSysOffset struct {
		Samples uint32
		Rsv     [3]uint32
		Ts      [51]PtpClockTime
	}
	PtpSysOffsetExtended struct {
		Samples uint32
		Clockid int32
		Rsv     [2]uint32
		Ts      [25][3]PtpClockTime
	}
	PtpSysOffsetPrecise struct {
		Device   PtpClockTime
		Realtime PtpClockTime
		Monoraw  PtpClockTime
		Rsv      [4]uint32
	}
)

type (
	HIDRawReportDescriptor struct {
		Size  uint32
		Value [4096]uint8
	}
	HIDRawDevInfo struct {
		Bustype uint32
		Vendor  int16
		Product int16
	}
)

type (
	EraseInfo struct {
		Start  uint32
		Length uint32
	}
	EraseInfo64 struct {
		Start  uint64
		Length uint64
	}
	MtdOobBuf struct {
		Start  uint32
		Length uint32
		Ptr    *uint8
	}
	MtdOobBuf64 struct {
		Start  uint64
		Pad    uint32
		Length uint32
		Ptr    uint64
	}
	MtdWriteReq struct {
		Start  uint64
		Len    uint64
		Ooblen uint64
		Data   uint64
		Oob    uint64
		Mode   uint8
		_      [7]uint8
	}
	MtdInfo struct {
		Type      uint8
		Flags     uint32
		Size      uint32
		Erasesize uint32
		Writesize uint32
		Oobsize   uint32
		_         uint64
	}
	RegionInfo struct {
		Offset      uint32
		Erasesize   uint32
		Numblocks   uint32
		Regionindex uint32
	}
	OtpInfo struct {
		Start  uint32
		Length uint32
		Locked uint32
	}
	NandOobinfo struct {
		Useecc   uint32
		Eccbytes uint32
		Oobfree  [8][2]uint32
		Eccpos   [32]uint32
	}
	NandOobfree struct {
		Offset uint32
		Length uint32
	}
	NandEcclayout struct {
		Eccbytes uint32
		Eccpos   [64]uint32
		Oobavail uint32
		Oobfree  [8]NandOobfree
	}
	MtdEccStats struct {
		Corrected uint32
		Failed    uint32
		Badblocks uint32
		Bbtblocks uint32
	}
)

type LandlockRulesetAttr struct {
	Access_fs  uint64
	Access_net uint64
	Scoped     uint64
}

type LandlockPathBeneathAttr struct {
	Allowed_access uint64
	Parent_fd      int32
}

type MountAttr struct {
	Attr_set    uint64
	Attr_clr    uint64
	Propagation uint64
	Userns_fd   uint64
}

type CANBitTiming struct {
	Bitrate      uint32
	Sample_point uint32
	Tq           uint32
	Prop_seg     uint32
	Phase_seg1   uint32
	Phase_seg2   uint32
	Sjw          uint32
	Brp          uint32
}

type CANBitTimingConst struct {
	Name      [16]uint8
	Tseg1_min uint32
	Tseg1_max uint32
	Tseg2_min uint32
	Tseg2_max uint32
	Sjw_max   uint32
	Brp_min   uint32
	Brp_max   uint32
	Brp_inc   uint32
}

type CANClock struct{ Freq uint32 }

type CANBusErrorCounters struct {
	Txerr uint16
	Rxerr uint16
}

type CANCtrlMode struct {
	Mask  uint32
	Flags uint32
}

type CANDeviceStats struct {
	Bus_error        uint32
	Error_warning    uint32
	Error_passive    uint32
	Bus_off          uint32
	Arbitration_lost uint32
	Restarts         uint32
}

type KCMAttach struct {
	Fd     int32
	Bpf_fd int32
}

type KCMUnattach struct{ Fd int32 }

type KCMClone struct{ Fd int32 }

type SchedAttr struct {
	Size     uint32
	Policy   uint32
	Flags    uint64
	Nice     int32
	Priority uint32
	Runtime  uint64
	Deadline uint64
	Period   uint64
	Util_min uint32
	Util_max uint32
}

type Cachestat_t struct {
	Cache            uint64
	Dirty            uint64
	Writeback        uint64
	Evicted          uint64
	Recently_evicted uint64
}

type CachestatRange struct {
	Off uint64
	Len uint64
}

type SockDiagReq struct {
	Family   uint8
	Protocol uint8
}

type DmNameList struct {
	Dev  uint64
	Next uint32
}

type RawSockaddrNFCLLCP struct {
	Sa_family        uint16
	Dev_idx          uint32
	Target_idx       uint32
	Nfc_protocol     uint32
	Dsap             uint8
	Ssap             uint8
	Service_name     [63]uint8
	Service_name_len uint32
}

type ifreq struct {
	Ifrn [16]byte
	Ifru [16]byte
}

type PtraceRegs struct {
	Ebx      int32
	Ecx      int32
	Edx      int32
	Esi      int32
	Edi      int32
	Ebp      int32
	Eax      int32
	Xds      int32
	Xes      int32
	Xfs      int32
	Xgs      int32
	Orig_eax int32
	Eip      int32
	Xcs      int32
	Eflags   int32
	Esp      int32
	Xss      int32
}

type Sysinfo_t struct {
	Uptime    int32
	Loads     [3]uint32
	Totalram  uint32
	Freeram   uint32
	Sharedram uint32
	Bufferram uint32
	Totalswap uint32
	Freeswap  uint32
	Procs     uint16
	Pad       uint16
	Totalhigh uint32
	Freehigh  uint32
	Unit      uint32
	_         [8]int8
}

type EpollEvent struct {
	Events uint32
	Fd     int32
	Pad    int32
}

type Siginfo struct {
	Signo int32
	Errno int32
	Code  int32
	_     [116]byte
}

type Taskstats struct {
	Version                   uint16
	Ac_exitcode               uint32
	Ac_flag                   uint8
	Ac_nice                   uint8
	_                         [4]byte
	Cpu_count                 uint64
	Cpu_delay_total           uint64
	Blkio_count               uint64
	Blkio_delay_total         uint64
	Swapin_count              uint64
	Swapin_delay_total        uint64
	Cpu_run_real_total        uint64
	Cpu_run_virtual_total     uint64
	Ac_comm                   [32]int8
	Ac_sched                  uint8
	Ac_pad                    [3]uint8
	_                         [4]byte
	Ac_uid                    uint32
	Ac_gid                    uint32
	Ac_pid                    uint32
	Ac_ppid                   uint32
	Ac_btime                  uint32
	_                         [4]byte
	Ac_etime                  uint64
	Ac_utime                  uint64
	Ac_stime                  uint64
	Ac_minflt                 uint64
	Ac_majflt                 uint64
	Coremem                   uint64
	Virtmem                   uint64
	Hiwater_rss               uint64
	Hiwater_vm                uint64
	Read_char                 uint64
	Write_char                uint64
	Read_syscalls             uint64
	Write_syscalls            uint64
	Read_bytes                uint64
	Write_bytes               uint64
	Cancelled_write_bytes     uint64
	Nvcsw                     uint64
	Nivcsw                    uint64
	Ac_utimescaled            uint64
	Ac_stimescaled            uint64
	Cpu_scaled_run_real_total uint64
	Freepages_count           uint64
	Freepages_delay_total     uint64
	Thrashing_count           uint64
	Thrashing_delay_total     uint64
	Ac_btime64                uint64
	Compact_count             uint64
	Compact_delay_total       uint64
	Ac_tgid                   uint32
	_                         [4]byte
	Ac_tgetime                uint64
	Ac_exe_dev                uint64
	Ac_exe_inode              uint64
	Wpcopy_count              uint64
	Wpcopy_delay_total        uint64
	Irq_count                 uint64
	Irq_delay_total           uint64
}

type SockaddrStorage struct {
	Family uint16
	Data   [122]byte
	_      uint32
}

type HDGeometry struct {
	Heads     uint8
	Sectors   uint8
	Cylinders uint16
	Start     uint32
}

type TpacketHdr struct {
	Status  uint32
	Len     uint32
	Snaplen uint32
	Mac     uint16
	Net     uint16
	Sec     uint32
	Usec    uint32
}

type RTCPLLInfo struct {
	Ctrl    int32
	Value   int32
	Max     int32
	Min     int32
	Posmult int32
	Negmult int32
	Clock   int32
}

type BlkpgPartition struct {
	Start   int64
	Length  int64
	Pno     int32
	Devname [64]uint8
	Volname [64]uint8
}

type CryptoUserAlg struct {
	Name        [64]int8
	Driver_name [64]int8
	Module_name [64]int8
	Type        uint32
	Mask        uint32
	Refcnt      uint32
	Flags       uint32
}

type CryptoStatAEAD struct {
	Type         [64]int8
	Encrypt_cnt  uint64
	Encrypt_tlen uint64
	Decrypt_cnt  uint64
	Decrypt_tlen uint64
	Err_cnt      uint64
}

type CryptoStatAKCipher struct {
	Type         [64]int8
	Encrypt_cnt  uint64
	Encrypt_tlen uint64
	Decrypt_cnt  uint64
	Decrypt_tlen uint64
	Verify_cnt   uint64
	Sign_cnt     uint64
	Err_cnt      uint64
}

type CryptoStatCipher struct {
	Type         [64]int8
	Encrypt_cnt  uint64
	Encrypt_tlen uint64
	Decrypt_cnt  uint64
	Decrypt_tlen uint64
	Err_cnt      uint64
}

type CryptoStatCompress struct {
	Type            [64]int8
	Compress_cnt    uint64
	Compress_tlen   uint64
	Decompress_cnt  uint64
	Decompress_tlen uint64
	Err_cnt         uint64
}

type CryptoStatHash struct {
	Type      [64]int8
	Hash_cnt  uint64
	Hash_tlen uint64
	Err_cnt   uint64
}

type CryptoStatKPP struct {
	Type                      [64]int8
	Setsecret_cnt             uint64
	Generate_public_key_cnt   uint64
	Compute_shared_secret_cnt uint64
	Err_cnt                   uint64
}

type CryptoStatRNG struct {
	Type          [64]int8
	Generate_cnt  uint64
	Generate_tlen uint64
	Seed_cnt      uint64
	Err_cnt       uint64
}

type CryptoStatLarval struct{ Type [64]int8 }

type CryptoReportLarval struct{ Type [64]int8 }

type CryptoReportHash struct {
	Type       [64]int8
	Blocksize  uint32
	Digestsize uint32
}

type CryptoReportCipher struct {
	Type        [64]int8
	Blocksize   uint32
	Min_keysize uint32
	Max_keysize uint32
}

type CryptoReportBlkCipher struct {
	Type        [64]int8
	Geniv       [64]int8
	Blocksize   uint32
	Min_keysize uint32
	Max_keysize uint32
	Ivsize      uint32
}

type CryptoReportAEAD struct {
	Type        [64]int8
	Geniv       [64]int8
	Blocksize   uint32
	Maxauthsize uint32
	Ivsize      uint32
}

type CryptoReportComp struct{ Type [64]int8 }

type CryptoReportRNG struct {
	Type     [64]int8
	Seedsize uint32
}

type CryptoReportAKCipher struct{ Type [64]int8 }

type CryptoReportKPP struct{ Type [64]int8 }

type CryptoReportAcomp struct{ Type [64]int8 }

type LoopInfo struct {
	Number           int32
	Device           uint16
	Inode            uint32
	Rdevice          uint16
	Offset           int32
	Encrypt_type     int32
	Encrypt_key_size int32
	Flags            int32
	Name             [64]int8
	Encrypt_key      [32]uint8
	Init             [2]uint32
	Reserved         [4]int8
}

type TIPCSubscr struct {
	Seq     TIPCServiceRange
	Timeout uint32
	Filter  uint32
	Handle  [8]int8
}

type TIPCSIOCLNReq struct {
	Peer     uint32
	Id       uint32
	Linkname [68]int8
}

type TIPCSIOCNodeIDReq struct {
	Peer uint32
	Id   [16]int8
}

type PPSKInfo struct {
	Assert_sequence uint32
	Clear_sequence  uint32
	Assert_tu       PPSKTime
	Clear_tu        PPSKTime
	Current_mode    int32
}

type RISCVHWProbePairs struct {
	Key   int64
	Value uint64
}

type PtracePsw struct {
	Mask uint64
	Addr uint64
}

type PtraceFpregs struct {
	Fpc  uint32
	Fprs [16]float64
}

type PtracePer struct {
	Control_regs  [3]uint64
	_             [8]byte
	Starting_addr uint64
	Ending_addr   uint64
	Perc_atmid    uint16
	Address       uint64
	Access_id     uint8
	_             [7]byte
}

type Statvfs_t struct {
	Flag        uint32
	Bsize       uint32
	Frsize      uint32
	Iosize      uint32
	Blocks      uint64
	Bfree       uint64
	Bavail      uint64
	Bresvd      uint64
	Files       uint64
	Ffree       uint64
	Favail      uint64
	Fresvd      uint64
	Syncreads   uint64
	Syncwrites  uint64
	Asyncreads  uint64
	Asyncwrites uint64
	Fsidx       Fsid
	Fsid        uint32
	Namemax     uint32
	Owner       uint32
	Spare       [4]uint32
	Fstypename  [32]byte
	Mntonname   [1024]byte
	Mntfromname [1024]byte
}

type BpfTimeval struct {
	Sec  int32
	Usec int32
}

type Ptmget struct {
	Cfd int32
	Sfd int32
	Cn  [1024]byte
	Sn  [1024]byte
}

type Sysctlnode struct {
	Flags           uint32
	Num             int32
	Name            [32]int8
	Ver             uint32
	X__rsvd         uint32
	Un              [16]byte
	X_sysctl_size   [8]byte
	X_sysctl_func   [8]byte
	X_sysctl_parent [8]byte
	X_sysctl_desc   [8]byte
}

type Uvmexp struct {
	Pagesize           int64
	Pagemask           int64
	Pageshift          int64
	Npages             int64
	Free               int64
	Active             int64
	Inactive           int64
	Paging             int64
	Wired              int64
	Zeropages          int64
	Reserve_pagedaemon int64
	Reserve_kernel     int64
	Freemin            int64
	Freetarg           int64
	Inactarg           int64
	Wiredmax           int64
	Nswapdev           int64
	Swpages            int64
	Swpginuse          int64
	Swpgonly           int64
	Nswget             int64
	Unused1            int64
	Cpuhit             int64
	Cpumiss            int64
	Faults             int64
	Traps              int64
	Intrs              int64
	Swtch              int64
	Softs              int64
	Syscalls           int64
	Pageins            int64
	Swapins            int64
	Swapouts           int64
	Pgswapin           int64
	Pgswapout          int64
	Forks              int64
	Forks_ppwait       int64
	Forks_sharevm      int64
	Pga_zerohit        int64
	Pga_zeromiss       int64
	Zeroaborts         int64
	Fltnoram           int64
	Fltnoanon          int64
	Fltpgwait          int64
	Fltpgrele          int64
	Fltrelck           int64
	Fltrelckok         int64
	Fltanget           int64
	Fltanretry         int64
	Fltamcopy          int64
	Fltnamap           int64
	Fltnomap           int64
	Fltlget            int64
	Fltget             int64
	Flt_anon           int64
	Flt_acow           int64
	Flt_obj            int64
	Flt_prcopy         int64
	Flt_przero         int64
	Pdwoke             int64
	Pdrevs             int64
	Unused4            int64
	Pdfreed            int64
	Pdscans            int64
	Pdanscan           int64
	Pdobscan           int64
	Pdreact            int64
	Pdbusy             int64
	Pdpageouts         int64
	Pdpending          int64
	Pddeact            int64
	Anonpages          int64
	Filepages          int64
	Execpages          int64
	Colorhit           int64
	Colormiss          int64
	Ncolors            int64
	Bootpages          int64
	Poolpages          int64
}

type fileObj struct {
	Atim Timespec
	Mtim Timespec
	Ctim Timespec
	Pad  [3]uint64
	Name *int8
}

type portEvent struct {
	Events int32
	Source uint16
	Pad    uint16
	Object uint64
	User   *byte
}

type strbuf struct {
	Maxlen int32
	Len    int32
	Buf    *int8
}

type Strioctl struct {
	Cmd    int32
	Timout int32
	Len    int32
	Dp     *int8
}

type Lifreq struct {
	Name   [32]int8
	Lifru1 [4]byte
	Type   uint32
	Lifru  [336]byte
}

type timeval_zos struct {
	Sec  int64
	_    [4]byte
	Usec int32
} // pad

type rusage_zos struct {
	Utime timeval_zos
	Stime timeval_zos
}

type Stat_LE_t struct {
	_            [4]byte
	Length       uint16
	Version      uint16
	Mode         int32
	Ino          uint32
	Dev          uint32
	Nlink        int32
	Uid          int32
	Gid          int32
	Size         int64
	Atim31       [4]byte
	Mtim31       [4]byte
	Ctim31       [4]byte
	Rdev         uint32
	Auditoraudit uint32
	Useraudit    uint32
	Blksize      int32
	Creatim31    [4]byte
	AuditID      [16]byte
	_            [4]byte
	File_tag     struct {
		Ccsid   uint16
		Txtflag uint16
	}
	CharsetID [8]byte
	Blocks    int64
	Genvalue  uint32
	Reftim31  [4]byte
	Fid       [8]byte
	Filefmt   byte
	Fspflag2  byte
	_         [2]byte
	Ctimemsec int32
	Seclabel  [8]byte
	_         [4]byte
	_         [4]byte
	Atim      Time_t
	Mtim      Time_t
	Ctim      Time_t
	Creatim   Time_t
	Reftim    Time_t
	_         [24]byte
} // eye catcher
// rsrvd5

type direntLE struct {
	Reclen uint16
	Namlen uint16
	Ino    uint32
	Extra  uintptr
	Name   [256]byte
}

type F_cnvrt struct {
	Cvtcmd int32
	Pccsid int16
	Fccsid int16
}

type W_Mnth struct {
	Hid   [4]byte
	Size  int32
	Cur1  int32
	Cur2  int32
	Devno uint32
	_     [4]byte
} //32bit pointer
//^

type W_Mntent struct {
	Fstype       uint32
	Mode         uint32
	Dev          uint32
	Parentdev    uint32
	Rootino      uint32
	Status       byte
	Ddname       [9]byte
	Fstname      [9]byte
	Fsname       [45]byte
	Pathlen      uint32
	Mountpoint   [1024]byte
	Jobname      [8]byte
	PID          int32
	Parmoffset   int32
	Parmlen      int16
	Owner        [8]byte
	Quiesceowner [8]byte
	_            [38]byte
}

type ConsMsg2 struct {
	Cm2Format       uint16
	Cm2R1           uint16
	Cm2Msglength    uint32
	Cm2Msg          *byte
	Cm2R2           [4]byte
	Cm2R3           [4]byte
	Cm2Routcde      *uint32
	Cm2Descr        *uint32
	Cm2Msgflag      uint32
	Cm2Token        uint32
	Cm2Msgid        *uint32
	Cm2R4           [4]byte
	Cm2DomToken     uint32
	Cm2DomMsgid     *uint32
	Cm2ModCartptr   *byte
	Cm2ModConsidptr *byte
	Cm2MsgCart      [8]byte
	Cm2MsgConsid    [4]byte
	Cm2R5           [12]byte
}

type SysvShmDesc64 struct {
	Perm   SysvIpcPerm
	_      [4]byte
	Lpid   int32
	Cpid   int32
	Nattch uint32
	_      [4]byte
	_      [4]byte
	_      [4]byte
	_      int32
	_      byte
	_      uint8
	_      uint16
	_      *byte
	Segsz  uint64
	Atime  int64
	Dtime  int64
	Ctime  int64
}

type DLLError struct {
	Err     error
	ObjName string
	Msg     string
} // DLLError describes reasons for DLL load failures.

type DLL struct {
	Name   string
	Handle Handle
} // A DLL implements access to a single DLL.

type Proc struct {
	Dll  *DLL
	Name string
	addr uintptr
} // A Proc implements access to a procedure inside a DLL.

type LazyDLL struct {
	Name   string
	System bool
	mu     sync.Mutex
	dll    *DLL
} // A LazyDLL implements access to a single DLL.
// It will delay the load of the DLL until the first
// call to its Handle method or to one of its
// LazyProc's Addr method.
// non nil once DLL is loaded

type LazyProc struct {
	Name string
	mu   sync.Mutex
	l    *LazyDLL
	proc *Proc
} // A LazyProc implements access to a procedure inside a LazyDLL.
// It delays the lookup until the Addr method is called.

type MemoryBasicInformation struct {
	BaseAddress       uintptr
	AllocationBase    uintptr
	AllocationProtect uint32
	PartitionId       uint16
	RegionSize        uintptr
	State             uint32
	Protect           uint32
	Type              uint32
}

type UserInfo10 struct {
	Name       *uint16
	Comment    *uint16
	UsrComment *uint16
	FullName   *uint16
}

type SidIdentifierAuthority struct{ Value [6]byte }

type SID struct{} // The security identifier (SID) structure is a variable-length
// structure used to uniquely identify users or groups.

type LUID struct {
	LowPart  uint32
	HighPart int32
}

type LUIDAndAttributes struct {
	Luid       LUID
	Attributes uint32
}

type SIDAndAttributes struct {
	Sid        *SID
	Attributes uint32
}

type Tokenuser struct{ User SIDAndAttributes }

type Tokenprimarygroup struct{ PrimaryGroup *SID }

type Tokengroups struct {
	GroupCount uint32
	Groups     [1]SIDAndAttributes
} // Use AllGroups() for iterating.

type Tokenprivileges struct {
	PrivilegeCount uint32
	Privileges     [1]LUIDAndAttributes
} // Use AllPrivileges() for iterating.

type Tokenmandatorylabel struct{ Label SIDAndAttributes }

type WTSSESSION_NOTIFICATION struct {
	Size      uint32
	SessionID uint32
}

type WTS_SESSION_INFO struct {
	SessionID         uint32
	WindowStationName *uint16
	State             uint32
}

type ACL struct {
	aclRevision byte
	sbz1        byte
	aclSize     uint16
	AceCount    uint16
	sbz2        uint16
}

type SECURITY_DESCRIPTOR struct {
	revision byte
	sbz1     byte
	control  SECURITY_DESCRIPTOR_CONTROL
	owner    *SID
	group    *SID
	sacl     *ACL
	dacl     *ACL
}

type SECURITY_QUALITY_OF_SERVICE struct {
	Length              uint32
	ImpersonationLevel  uint32
	ContextTrackingMode byte
	EffectiveOnly       byte
}

type SecurityAttributes struct {
	Length             uint32
	SecurityDescriptor *SECURITY_DESCRIPTOR
	InheritHandle      uint32
}

type EXPLICIT_ACCESS struct {
	AccessPermissions ACCESS_MASK
	AccessMode        ACCESS_MODE
	Inheritance       uint32
	Trustee           TRUSTEE
}

type ACE_HEADER struct {
	AceType  uint8
	AceFlags uint8
	AceSize  uint16
} // https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-ace_header

type ACCESS_ALLOWED_ACE struct {
	Header   ACE_HEADER
	Mask     ACCESS_MASK
	SidStart uint32
} // https://learn.microsoft.com/en-us/windows/win32/api/winnt/ns-winnt-access_allowed_ace

type TRUSTEE struct {
	MultipleTrustee          *TRUSTEE
	MultipleTrusteeOperation MULTIPLE_TRUSTEE_OPERATION
	TrusteeForm              TRUSTEE_FORM
	TrusteeType              TRUSTEE_TYPE
	TrusteeValue             TrusteeValue
}

type OBJECTS_AND_SID struct {
	ObjectsPresent          uint32
	ObjectTypeGuid          GUID
	InheritedObjectTypeGuid GUID
	Sid                     *SID
}

type OBJECTS_AND_NAME struct {
	ObjectsPresent          uint32
	ObjectType              SE_OBJECT_TYPE
	ObjectTypeName          *uint16
	InheritedObjectTypeName *uint16
	Name                    *uint16
}

type ENUM_SERVICE_STATUS struct {
	ServiceName   *uint16
	DisplayName   *uint16
	ServiceStatus SERVICE_STATUS
}

type SERVICE_STATUS struct {
	ServiceType             uint32
	CurrentState            uint32
	ControlsAccepted        uint32
	Win32ExitCode           uint32
	ServiceSpecificExitCode uint32
	CheckPoint              uint32
	WaitHint                uint32
}

type SERVICE_TABLE_ENTRY struct {
	ServiceName *uint16
	ServiceProc uintptr
}

type QUERY_SERVICE_CONFIG struct {
	ServiceType      uint32
	StartType        uint32
	ErrorControl     uint32
	BinaryPathName   *uint16
	LoadOrderGroup   *uint16
	TagId            uint32
	Dependencies     *uint16
	ServiceStartName *uint16
	DisplayName      *uint16
}

type SERVICE_DESCRIPTION struct{ Description *uint16 }

type SERVICE_DELAYED_AUTO_START_INFO struct{ IsDelayedAutoStartUp uint32 }

type SERVICE_STATUS_PROCESS struct {
	ServiceType             uint32
	CurrentState            uint32
	ControlsAccepted        uint32
	Win32ExitCode           uint32
	ServiceSpecificExitCode uint32
	CheckPoint              uint32
	WaitHint                uint32
	ProcessId               uint32
	ServiceFlags            uint32
}

type ENUM_SERVICE_STATUS_PROCESS struct {
	ServiceName          *uint16
	DisplayName          *uint16
	ServiceStatusProcess SERVICE_STATUS_PROCESS
}

type SERVICE_NOTIFY struct {
	Version               uint32
	NotifyCallback        uintptr
	Context               uintptr
	NotificationStatus    uint32
	ServiceStatus         SERVICE_STATUS_PROCESS
	NotificationTriggered uint32
	ServiceNames          *uint16
}

type SERVICE_FAILURE_ACTIONS struct {
	ResetPeriod  uint32
	RebootMsg    *uint16
	Command      *uint16
	ActionsCount uint32
	Actions      *SC_ACTION
}

type SERVICE_FAILURE_ACTIONS_FLAG struct{ FailureActionsOnNonCrashFailures int32 }

type SC_ACTION struct {
	Type  uint32
	Delay uint32
}

type QUERY_SERVICE_LOCK_STATUS struct {
	IsLocked     uint32
	LockOwner    *uint16
	LockDuration uint32
}

type DevInfoData struct {
	size      uint32
	ClassGUID GUID
	DevInst   DEVINST
	_         uintptr
} // DevInfoData is a device information structure (references a device instance that is a member of a device information set)

type DevInfoListDetailData struct {
	size                uint32
	ClassGUID           GUID
	RemoteMachineHandle Handle
	remoteMachineName   [SP_MAX_MACHINENAME_LENGTH]uint16
} // DevInfoListDetailData is a structure for detailed information on a device information set (used for SetupDiGetDeviceInfoListDetail which supersedes the functionality of SetupDiGetDeviceInfoListClass).
// Use unsafeSizeOf method

type DevInstallParams struct {
	size                     uint32
	Flags                    DI_FLAGS
	FlagsEx                  DI_FLAGSEX
	hwndParent               uintptr
	InstallMsgHandler        uintptr
	InstallMsgHandlerContext uintptr
	FileQueue                HSPFILEQ
	_                        uintptr
	_                        uint32
	driverPath               [MAX_PATH]uint16
} // DevInstallParams is device installation parameters structure (associated with a particular device information element, or globally with a device information set)

type ClassInstallHeader struct {
	size            uint32
	InstallFunction DI_FUNCTION
} // ClassInstallHeader is the first member of any class install parameters structure. It contains the device installation request code that defines the format of the rest of the install parameters structure.

type PropChangeParams struct {
	ClassInstallHeader ClassInstallHeader
	StateChange        DICS_STATE
	Scope              DICS_FLAG
	HwProfile          uint32
} // PropChangeParams is a structure corresponding to a DIF_PROPERTYCHANGE install function.

type RemoveDeviceParams struct {
	ClassInstallHeader ClassInstallHeader
	Scope              DI_REMOVEDEVICE
	HwProfile          uint32
} // RemoveDeviceParams is a structure corresponding to a DIF_REMOVE install function.

type DrvInfoData struct {
	size          uint32
	DriverType    uint32
	_             uintptr
	description   [LINE_LEN]uint16
	mfgName       [LINE_LEN]uint16
	providerName  [LINE_LEN]uint16
	DriverDate    Filetime
	DriverVersion uint64
} // DrvInfoData is driver information structure (member of a driver info list that may be associated with a particular device instance, or (globally) with a device information set)

type DrvInfoDetailData struct {
	size            uint32
	InfDate         Filetime
	compatIDsOffset uint32
	compatIDsLength uint32
	_               uintptr
	sectionName     [LINE_LEN]uint16
	infFileName     [MAX_PATH]uint16
	drvDescription  [LINE_LEN]uint16
	hardwareID      [1]uint16
} // DrvInfoDetailData is driver information details structure (provides detailed information about a particular driver information structure)
// Use unsafeSizeOf method

type DEVPROPKEY struct {
	FmtID DEVPROPGUID
	PID   DEVPROPID
} // DEVPROPKEY represents a device property key for a device property in the
// unified device property model.

type RawSockaddrBth struct {
	AddressFamily  [2]byte
	BtAddr         [8]byte
	ServiceClassId [16]byte
	Port           [4]byte
}

type SockaddrBth struct {
	BtAddr         uint64
	ServiceClassId GUID
	Port           uint32
	raw            RawSockaddrBth
}

type sysLinger struct {
	Onoff  uint16
	Linger uint16
}

type PSAPI_WORKING_SET_EX_INFORMATION struct {
	VirtualAddress    Pointer
	VirtualAttributes PSAPI_WORKING_SET_EX_BLOCK
} // PSAPI_WORKING_SET_EX_INFORMATION contains extended working set information for a process.
// A PSAPI_WORKING_SET_EX_BLOCK union that indicates the attributes of the page at VirtualAddress.

type Overlapped struct {
	Internal     uintptr
	InternalHigh uintptr
	Offset       uint32
	OffsetHigh   uint32
	HEvent       Handle
}

type FileNotifyInformation struct {
	NextEntryOffset uint32
	Action          uint32
	FileNameLength  uint32
	FileName        uint16
}

type Filetime struct {
	LowDateTime  uint32
	HighDateTime uint32
}

type Win32finddata struct {
	FileAttributes    uint32
	CreationTime      Filetime
	LastAccessTime    Filetime
	LastWriteTime     Filetime
	FileSizeHigh      uint32
	FileSizeLow       uint32
	Reserved0         uint32
	Reserved1         uint32
	FileName          [MAX_PATH - 1]uint16
	AlternateFileName [13]uint16
}

type win32finddata1 struct {
	FileAttributes    uint32
	CreationTime      Filetime
	LastAccessTime    Filetime
	LastWriteTime     Filetime
	FileSizeHigh      uint32
	FileSizeLow       uint32
	Reserved0         uint32
	Reserved1         uint32
	FileName          [MAX_PATH]uint16
	AlternateFileName [14]uint16
} // This is the actual system call structure.
// Win32finddata is what we committed to in Go 1.

type ByHandleFileInformation struct {
	FileAttributes     uint32
	CreationTime       Filetime
	LastAccessTime     Filetime
	LastWriteTime      Filetime
	VolumeSerialNumber uint32
	FileSizeHigh       uint32
	FileSizeLow        uint32
	NumberOfLinks      uint32
	FileIndexHigh      uint32
	FileIndexLow       uint32
}

type Win32FileAttributeData struct {
	FileAttributes uint32
	CreationTime   Filetime
	LastAccessTime Filetime
	LastWriteTime  Filetime
	FileSizeHigh   uint32
	FileSizeLow    uint32
}

type StartupInfo struct {
	Cb            uint32
	_             *uint16
	Desktop       *uint16
	Title         *uint16
	X             uint32
	Y             uint32
	XSize         uint32
	YSize         uint32
	XCountChars   uint32
	YCountChars   uint32
	FillAttribute uint32
	Flags         uint32
	ShowWindow    uint16
	_             uint16
	_             *byte
	StdInput      Handle
	StdOutput     Handle
	StdErr        Handle
}

type StartupInfoEx struct {
	StartupInfo
	ProcThreadAttributeList *ProcThreadAttributeList
}

type ProcThreadAttributeList struct{} // ProcThreadAttributeList is a placeholder type to represent a PROC_THREAD_ATTRIBUTE_LIST.
//
// To create a *ProcThreadAttributeList, use NewProcThreadAttributeList, update
// it with ProcThreadAttributeListContainer.Update, free its memory using
// ProcThreadAttributeListContainer.Delete, and access the list itself using
// ProcThreadAttributeListContainer.List.

type ProcThreadAttributeListContainer struct {
	data     *ProcThreadAttributeList
	pointers []unsafe.Pointer
}

type ProcessInformation struct {
	Process   Handle
	Thread    Handle
	ProcessId uint32
	ThreadId  uint32
}

type ProcessEntry32 struct {
	Size            uint32
	Usage           uint32
	ProcessID       uint32
	DefaultHeapID   uintptr
	ModuleID        uint32
	Threads         uint32
	ParentProcessID uint32
	PriClassBase    int32
	Flags           uint32
	ExeFile         [MAX_PATH]uint16
}

type ThreadEntry32 struct {
	Size           uint32
	Usage          uint32
	ThreadID       uint32
	OwnerProcessID uint32
	BasePri        int32
	DeltaPri       int32
	Flags          uint32
}

type ModuleEntry32 struct {
	Size         uint32
	ModuleID     uint32
	ProcessID    uint32
	GlblcntUsage uint32
	ProccntUsage uint32
	ModBaseAddr  uintptr
	ModBaseSize  uint32
	ModuleHandle Handle
	Module       [MAX_MODULE_NAME32 + 1]uint16
	ExePath      [MAX_PATH]uint16
}

type Systemtime struct {
	Year         uint16
	Month        uint16
	DayOfWeek    uint16
	Day          uint16
	Hour         uint16
	Minute       uint16
	Second       uint16
	Milliseconds uint16
}

type Timezoneinformation struct {
	Bias         int32
	StandardName [32]uint16
	StandardDate Systemtime
	StandardBias int32
	DaylightName [32]uint16
	DaylightDate Systemtime
	DaylightBias int32
}

type WSABuf struct {
	Len uint32
	Buf *byte
}

type WSAMsg struct {
	Name        *syscall.RawSockaddrAny
	Namelen     int32
	Buffers     *WSABuf
	BufferCount uint32
	Control     WSABuf
	Flags       uint32
}

type WSACMSGHDR struct {
	Len   uintptr
	Level int32
	Type  int32
}

type IN_PKTINFO struct {
	Addr    [4]byte
	Ifindex uint32
}

type IN6_PKTINFO struct {
	Addr    [16]byte
	Ifindex uint32
}

type Hostent struct {
	Name     *byte
	Aliases  **byte
	AddrType uint16
	Length   uint16
	AddrList **byte
}

type Protoent struct {
	Name    *byte
	Aliases **byte
	Proto   uint16
}

type DNSSRVData struct {
	Target   *uint16
	Priority uint16
	Weight   uint16
	Port     uint16
	Pad      uint16
}

type DNSPTRData struct{ Host *uint16 }

type DNSMXData struct {
	NameExchange *uint16
	Preference   uint16
	Pad          uint16
}

type DNSTXTData struct {
	StringCount uint16
	StringArray [1]*uint16
}

type DNSRecord struct {
	Next     *DNSRecord
	Name     *uint16
	Type     uint16
	Length   uint16
	Dw       uint32
	Ttl      uint32
	Reserved uint32
	Data     [40]byte
}

type TransmitFileBuffers struct {
	Head       uintptr
	HeadLength uint32
	Tail       uintptr
	TailLength uint32
}

type InterfaceInfo struct {
	Flags            uint32
	Address          SockaddrGen
	BroadcastAddress SockaddrGen
	Netmask          SockaddrGen
}

type IpAddressString struct{ String [16]byte }

type IpAddrString struct {
	Next      *IpAddrString
	IpAddress IpAddressString
	IpMask    IpMaskString
	Context   uint32
}

type IpAdapterInfo struct {
	Next                *IpAdapterInfo
	ComboIndex          uint32
	AdapterName         [MAX_ADAPTER_NAME_LENGTH + 4]byte
	Description         [MAX_ADAPTER_DESCRIPTION_LENGTH + 4]byte
	AddressLength       uint32
	Address             [MAX_ADAPTER_ADDRESS_LENGTH]byte
	Index               uint32
	Type                uint32
	DhcpEnabled         uint32
	CurrentIpAddress    *IpAddrString
	IpAddressList       IpAddrString
	GatewayList         IpAddrString
	DhcpServer          IpAddrString
	HaveWins            bool
	PrimaryWinsServer   IpAddrString
	SecondaryWinsServer IpAddrString
	LeaseObtained       int64
	LeaseExpires        int64
}

type MibIfRow struct {
	Name            [MAX_INTERFACE_NAME_LEN]uint16
	Index           uint32
	Type            uint32
	Mtu             uint32
	Speed           uint32
	PhysAddrLen     uint32
	PhysAddr        [MAXLEN_PHYSADDR]byte
	AdminStatus     uint32
	OperStatus      uint32
	LastChange      uint32
	InOctets        uint32
	InUcastPkts     uint32
	InNUcastPkts    uint32
	InDiscards      uint32
	InErrors        uint32
	InUnknownProtos uint32
	OutOctets       uint32
	OutUcastPkts    uint32
	OutNUcastPkts   uint32
	OutDiscards     uint32
	OutErrors       uint32
	OutQLen         uint32
	DescrLen        uint32
	Descr           [MAXLEN_IFDESCR]byte
}

type CertInfo struct {
	Version              uint32
	SerialNumber         CryptIntegerBlob
	SignatureAlgorithm   CryptAlgorithmIdentifier
	Issuer               CertNameBlob
	NotBefore            Filetime
	NotAfter             Filetime
	Subject              CertNameBlob
	SubjectPublicKeyInfo CertPublicKeyInfo
	IssuerUniqueId       CryptBitBlob
	SubjectUniqueId      CryptBitBlob
	CountExtensions      uint32
	Extensions           *CertExtension
}

type CertExtension struct {
	ObjId    *byte
	Critical int32
	Value    CryptObjidBlob
}

type CryptAlgorithmIdentifier struct {
	ObjId      *byte
	Parameters CryptObjidBlob
}

type CertPublicKeyInfo struct {
	Algorithm CryptAlgorithmIdentifier
	PublicKey CryptBitBlob
}

type DataBlob struct {
	Size uint32
	Data *byte
}

type CryptBitBlob struct {
	Size       uint32
	Data       *byte
	UnusedBits uint32
}

type CertContext struct {
	EncodingType uint32
	EncodedCert  *byte
	Length       uint32
	CertInfo     *CertInfo
	Store        Handle
}

type CertChainContext struct {
	Size                       uint32
	TrustStatus                CertTrustStatus
	ChainCount                 uint32
	Chains                     **CertSimpleChain
	LowerQualityChainCount     uint32
	LowerQualityChains         **CertChainContext
	HasRevocationFreshnessTime uint32
	RevocationFreshnessTime    uint32
}

type CertTrustListInfo struct{}

type CertSimpleChain struct {
	Size                       uint32
	TrustStatus                CertTrustStatus
	NumElements                uint32
	Elements                   **CertChainElement
	TrustListInfo              *CertTrustListInfo
	HasRevocationFreshnessTime uint32
	RevocationFreshnessTime    uint32
}

type CertChainElement struct {
	Size              uint32
	CertContext       *CertContext
	TrustStatus       CertTrustStatus
	RevocationInfo    *CertRevocationInfo
	IssuanceUsage     *CertEnhKeyUsage
	ApplicationUsage  *CertEnhKeyUsage
	ExtendedErrorInfo *uint16
}

type CertRevocationCrlInfo struct{}

type CertRevocationInfo struct {
	Size             uint32
	RevocationResult uint32
	RevocationOid    *byte
	OidSpecificInfo  Pointer
	HasFreshnessTime uint32
	FreshnessTime    uint32
	CrlInfo          *CertRevocationCrlInfo
}

type CertTrustStatus struct {
	ErrorStatus uint32
	InfoStatus  uint32
}

type CertUsageMatch struct {
	Type  uint32
	Usage CertEnhKeyUsage
}

type CertEnhKeyUsage struct {
	Length           uint32
	UsageIdentifiers **byte
}

type CertChainPara struct {
	Size                         uint32
	RequestedUsage               CertUsageMatch
	RequstedIssuancePolicy       CertUsageMatch
	URLRetrievalTimeout          uint32
	CheckRevocationFreshnessTime uint32
	RevocationFreshnessTime      uint32
	CacheResync                  *Filetime
}

type CertChainPolicyPara struct {
	Size            uint32
	Flags           uint32
	ExtraPolicyPara Pointer
}

type SSLExtraCertChainPolicyPara struct {
	Size       uint32
	AuthType   uint32
	Checks     uint32
	ServerName *uint16
}

type CertChainPolicyStatus struct {
	Size              uint32
	Error             uint32
	ChainIndex        uint32
	ElementIndex      uint32
	ExtraPolicyStatus Pointer
}

type CertPolicyInfo struct {
	Identifier      *byte
	CountQualifiers uint32
	Qualifiers      *CertPolicyQualifierInfo
}

type CertPoliciesInfo struct {
	Count       uint32
	PolicyInfos *CertPolicyInfo
}

type CertPolicyQualifierInfo struct{}

type CertStrongSignPara struct {
	Size                      uint32
	InfoChoice                uint32
	InfoOrSerializedInfoOrOID unsafe.Pointer
}

type CryptProtectPromptStruct struct {
	Size        uint32
	PromptFlags uint32
	App         HWND
	Prompt      *uint16
}

type CertChainFindByIssuerPara struct {
	Size                   uint32
	UsageIdentifier        *byte
	KeySpec                uint32
	AcquirePrivateKeyFlags uint32
	IssuerCount            uint32
	Issuer                 Pointer
	FindCallback           Pointer
	FindArg                Pointer
	IssuerChainIndex       *uint32
	IssuerElementIndex     *uint32
}

type WinTrustData struct {
	Size                            uint32
	PolicyCallbackData              uintptr
	SIPClientData                   uintptr
	UIChoice                        uint32
	RevocationChecks                uint32
	UnionChoice                     uint32
	FileOrCatalogOrBlobOrSgnrOrCert unsafe.Pointer
	StateAction                     uint32
	StateData                       Handle
	URLReference                    *uint16
	ProvFlags                       uint32
	UIContext                       uint32
	SignatureSettings               *WinTrustSignatureSettings
}

type WinTrustFileInfo struct {
	Size         uint32
	FilePath     *uint16
	File         Handle
	KnownSubject *GUID
}

type WinTrustSignatureSettings struct {
	Size             uint32
	Index            uint32
	Flags            uint32
	SecondarySigs    uint32
	VerifiedSigIndex uint32
	CryptoPolicy     *CertStrongSignPara
}

type AddrinfoW struct {
	Flags     int32
	Family    int32
	Socktype  int32
	Protocol  int32
	Addrlen   uintptr
	Canonname *uint16
	Addr      uintptr
	Next      *AddrinfoW
}

type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

type WSAProtocolInfo struct {
	ServiceFlags1     uint32
	ServiceFlags2     uint32
	ServiceFlags3     uint32
	ServiceFlags4     uint32
	ProviderFlags     uint32
	ProviderId        GUID
	CatalogEntryId    uint32
	ProtocolChain     WSAProtocolChain
	Version           int32
	AddressFamily     int32
	MaxSockAddr       int32
	MinSockAddr       int32
	SocketType        int32
	Protocol          int32
	ProtocolMaxOffset int32
	NetworkByteOrder  int32
	SecurityScheme    int32
	MessageSize       uint32
	ProviderReserved  uint32
	ProtocolName      [WSAPROTOCOL_LEN + 1]uint16
}

type WSAProtocolChain struct {
	ChainLen     int32
	ChainEntries [MAX_PROTOCOL_CHAIN]uint32
}

type TCPKeepalive struct {
	OnOff    uint32
	Time     uint32
	Interval uint32
}

type symbolicLinkReparseBuffer struct {
	SubstituteNameOffset uint16
	SubstituteNameLength uint16
	PrintNameOffset      uint16
	PrintNameLength      uint16
	Flags                uint32
	PathBuffer           [1]uint16
}

type mountPointReparseBuffer struct {
	SubstituteNameOffset uint16
	SubstituteNameLength uint16
	PrintNameOffset      uint16
	PrintNameLength      uint16
	PathBuffer           [1]uint16
}

type reparseDataBuffer struct {
	ReparseTag        uint32
	ReparseDataLength uint16
	Reserved          uint16
	reparseBuffer     byte
} // GenericReparseBuffer

type SocketAddress struct {
	Sockaddr       *syscall.RawSockaddrAny
	SockaddrLength int32
}

type IpAdapterUnicastAddress struct {
	Length             uint32
	Flags              uint32
	Next               *IpAdapterUnicastAddress
	Address            SocketAddress
	PrefixOrigin       int32
	SuffixOrigin       int32
	DadState           int32
	ValidLifetime      uint32
	PreferredLifetime  uint32
	LeaseLifetime      uint32
	OnLinkPrefixLength uint8
}

type IpAdapterAnycastAddress struct {
	Length  uint32
	Flags   uint32
	Next    *IpAdapterAnycastAddress
	Address SocketAddress
}

type IpAdapterMulticastAddress struct {
	Length  uint32
	Flags   uint32
	Next    *IpAdapterMulticastAddress
	Address SocketAddress
}

type IpAdapterDnsServerAdapter struct {
	Length   uint32
	Reserved uint32
	Next     *IpAdapterDnsServerAdapter
	Address  SocketAddress
}

type IpAdapterPrefix struct {
	Length       uint32
	Flags        uint32
	Next         *IpAdapterPrefix
	Address      SocketAddress
	PrefixLength uint32
}

type IpAdapterAddresses struct {
	Length                 uint32
	IfIndex                uint32
	Next                   *IpAdapterAddresses
	AdapterName            *byte
	FirstUnicastAddress    *IpAdapterUnicastAddress
	FirstAnycastAddress    *IpAdapterAnycastAddress
	FirstMulticastAddress  *IpAdapterMulticastAddress
	FirstDnsServerAddress  *IpAdapterDnsServerAdapter
	DnsSuffix              *uint16
	Description            *uint16
	FriendlyName           *uint16
	PhysicalAddress        [syscall.MAX_ADAPTER_ADDRESS_LENGTH]byte
	PhysicalAddressLength  uint32
	Flags                  uint32
	Mtu                    uint32
	IfType                 uint32
	OperStatus             uint32
	Ipv6IfIndex            uint32
	ZoneIndices            [16]uint32
	FirstPrefix            *IpAdapterPrefix
	TransmitLinkSpeed      uint64
	ReceiveLinkSpeed       uint64
	FirstWinsServerAddress *IpAdapterWinsServerAddress
	FirstGatewayAddress    *IpAdapterGatewayAddress
	Ipv4Metric             uint32
	Ipv6Metric             uint32
	Luid                   uint64
	Dhcpv4Server           SocketAddress
	CompartmentId          uint32
	NetworkGuid            GUID
	ConnectionType         uint32
	TunnelType             uint32
	Dhcpv6Server           SocketAddress
	Dhcpv6ClientDuid       [MAX_DHCPV6_DUID_LENGTH]byte
	Dhcpv6ClientDuidLength uint32
	Dhcpv6Iaid             uint32
	FirstDnsSuffix         *IpAdapterDNSSuffix
}

type IpAdapterWinsServerAddress struct {
	Length   uint32
	Reserved uint32
	Next     *IpAdapterWinsServerAddress
	Address  SocketAddress
}

type IpAdapterGatewayAddress struct {
	Length   uint32
	Reserved uint32
	Next     *IpAdapterGatewayAddress
	Address  SocketAddress
}

type IpAdapterDNSSuffix struct {
	Next   *IpAdapterDNSSuffix
	String [MAX_DNS_SUFFIX_STRING_LENGTH]uint16
}

type MibIfRow2 struct {
	InterfaceLuid               uint64
	InterfaceIndex              uint32
	InterfaceGuid               GUID
	Alias                       [IF_MAX_STRING_SIZE + 1]uint16
	Description                 [IF_MAX_STRING_SIZE + 1]uint16
	PhysicalAddressLength       uint32
	PhysicalAddress             [IF_MAX_PHYS_ADDRESS_LENGTH]uint8
	PermanentPhysicalAddress    [IF_MAX_PHYS_ADDRESS_LENGTH]uint8
	Mtu                         uint32
	Type                        uint32
	TunnelType                  uint32
	MediaType                   uint32
	PhysicalMediumType          uint32
	AccessType                  uint32
	DirectionType               uint32
	InterfaceAndOperStatusFlags uint8
	OperStatus                  uint32
	AdminStatus                 uint32
	MediaConnectState           uint32
	NetworkGuid                 GUID
	ConnectionType              uint32
	TransmitLinkSpeed           uint64
	ReceiveLinkSpeed            uint64
	InOctets                    uint64
	InUcastPkts                 uint64
	InNUcastPkts                uint64
	InDiscards                  uint64
	InErrors                    uint64
	InUnknownProtos             uint64
	InUcastOctets               uint64
	InMulticastOctets           uint64
	InBroadcastOctets           uint64
	OutOctets                   uint64
	OutUcastPkts                uint64
	OutNUcastPkts               uint64
	OutDiscards                 uint64
	OutErrors                   uint64
	OutUcastOctets              uint64
	OutMulticastOctets          uint64
	OutBroadcastOctets          uint64
	OutQLen                     uint64
} // MibIfRow2 stores information about a particular interface. See
// https://learn.microsoft.com/en-us/windows/win32/api/netioapi/ns-netioapi-mib_if_row2.

type MibUnicastIpAddressRow struct {
	Address            RawSockaddrInet6
	InterfaceLuid      uint64
	InterfaceIndex     uint32
	PrefixOrigin       uint32
	SuffixOrigin       uint32
	ValidLifetime      uint32
	PreferredLifetime  uint32
	OnLinkPrefixLength uint8
	SkipAsSource       uint8
	DadState           uint32
	ScopeId            uint32
	CreationTimeStamp  Filetime
} // MIB_UNICASTIPADDRESS_ROW stores information about a unicast IP address. See
// https://learn.microsoft.com/en-us/windows/win32/api/netioapi/ns-netioapi-mib_unicastipaddress_row.
// SOCKADDR_INET union

type MibIpInterfaceRow struct {
	Family                               uint16
	InterfaceLuid                        uint64
	InterfaceIndex                       uint32
	MaxReassemblySize                    uint32
	InterfaceIdentifier                  uint64
	MinRouterAdvertisementInterval       uint32
	MaxRouterAdvertisementInterval       uint32
	AdvertisingEnabled                   uint8
	ForwardingEnabled                    uint8
	WeakHostSend                         uint8
	WeakHostReceive                      uint8
	UseAutomaticMetric                   uint8
	UseNeighborUnreachabilityDetection   uint8
	ManagedAddressConfigurationSupported uint8
	OtherStatefulConfigurationSupported  uint8
	AdvertiseDefaultRoute                uint8
	RouterDiscoveryBehavior              uint32
	DadTransmits                         uint32
	BaseReachableTime                    uint32
	RetransmitTime                       uint32
	PathMtuDiscoveryTimeout              uint32
	LinkLocalAddressBehavior             uint32
	LinkLocalAddressTimeout              uint32
	ZoneIndices                          [ScopeLevelCount]uint32
	SitePrefixLength                     uint32
	Metric                               uint32
	NlMtu                                uint32
	Connected                            uint8
	SupportsWakeUpPatterns               uint8
	SupportsNeighborDiscovery            uint8
	SupportsRouterDiscovery              uint8
	ReachableTime                        uint32
	TransmitOffload                      uint32
	ReceiveOffload                       uint32
	DisableDefaultRoutes                 uint8
} // MIB_IPINTERFACE_ROW stores interface management information for a particular IP address family on a network interface.
// See https://learn.microsoft.com/en-us/windows/win32/api/netioapi/ns-netioapi-mib_ipinterface_row.

type Coord struct {
	X int16
	Y int16
}

type SmallRect struct {
	Left   int16
	Top    int16
	Right  int16
	Bottom int16
}

type ConsoleScreenBufferInfo struct {
	Size              Coord
	CursorPosition    Coord
	Attributes        uint16
	Window            SmallRect
	MaximumWindowSize Coord
}

type IO_COUNTERS struct {
	ReadOperationCount  uint64
	WriteOperationCount uint64
	OtherOperationCount uint64
	ReadTransferCount   uint64
	WriteTransferCount  uint64
	OtherTransferCount  uint64
}

type JOBOBJECT_EXTENDED_LIMIT_INFORMATION struct {
	BasicLimitInformation JOBOBJECT_BASIC_LIMIT_INFORMATION
	IoInfo                IO_COUNTERS
	ProcessMemoryLimit    uintptr
	JobMemoryLimit        uintptr
	PeakProcessMemoryUsed uintptr
	PeakJobMemoryUsed     uintptr
}

type JOBOBJECT_BASIC_UI_RESTRICTIONS struct{ UIRestrictionsClass uint32 }

type OsVersionInfoEx struct {
	osVersionInfoSize uint32
	MajorVersion      uint32
	MinorVersion      uint32
	BuildNumber       uint32
	PlatformId        uint32
	CsdVersion        [128]uint16
	ServicePackMajor  uint16
	ServicePackMinor  uint16
	SuiteMask         uint16
	ProductType       byte
	_                 byte
}

type CommTimeouts struct {
	ReadIntervalTimeout         uint32
	ReadTotalTimeoutMultiplier  uint32
	ReadTotalTimeoutConstant    uint32
	WriteTotalTimeoutMultiplier uint32
	WriteTotalTimeoutConstant   uint32
}

type NTUnicodeString struct {
	Length        uint16
	MaximumLength uint16
	Buffer        *uint16
} // NTUnicodeString is a UTF-16 string for NT native APIs, corresponding to UNICODE_STRING.
// Note: Length and MaximumLength are in *bytes*, not uint16s.
// They should always be even.

type NTString struct {
	Length        uint16
	MaximumLength uint16
	Buffer        *byte
} // NTString is an ANSI string for NT native APIs, corresponding to STRING.

type LIST_ENTRY struct {
	Flink *LIST_ENTRY
	Blink *LIST_ENTRY
}

type RUNTIME_FUNCTION struct {
	BeginAddress uint32
	EndAddress   uint32
	UnwindData   uint32
}

type LDR_DATA_TABLE_ENTRY struct {
	reserved1          [2]uintptr
	InMemoryOrderLinks LIST_ENTRY
	reserved2          [2]uintptr
	DllBase            uintptr
	reserved3          [2]uintptr
	FullDllName        NTUnicodeString
	reserved4          [8]byte
	reserved5          [3]uintptr
	reserved6          uintptr
	TimeDateStamp      uint32
}

type PEB_LDR_DATA struct {
	reserved1               [8]byte
	reserved2               [3]uintptr
	InMemoryOrderModuleList LIST_ENTRY
}

type CURDIR struct {
	DosPath NTUnicodeString
	Handle  Handle
}

type RTL_DRIVE_LETTER_CURDIR struct {
	Flags     uint16
	Length    uint16
	TimeStamp uint32
	DosPath   NTString
}

type RTL_USER_PROCESS_PARAMETERS struct {
	MaximumLength, Length                                                         uint32
	Flags, DebugFlags                                                             uint32
	ConsoleHandle                                                                 Handle
	ConsoleFlags                                                                  uint32
	StandardInput, StandardOutput, StandardError                                  Handle
	CurrentDirectory                                                              CURDIR
	DllPath                                                                       NTUnicodeString
	ImagePathName                                                                 NTUnicodeString
	CommandLine                                                                   NTUnicodeString
	Environment                                                                   unsafe.Pointer
	StartingX, StartingY, CountX, CountY, CountCharsX, CountCharsY, FillAttribute uint32
	WindowFlags, ShowWindowFlags                                                  uint32
	WindowTitle, DesktopInfo, ShellInfo, RuntimeData                              NTUnicodeString
	CurrentDirectories                                                            [32]RTL_DRIVE_LETTER_CURDIR
	EnvironmentSize, EnvironmentVersion                                           uintptr
	PackageDependencyData                                                         unsafe.Pointer
	ProcessGroupId                                                                uint32
	LoaderThreads                                                                 uint32
	RedirectionDllName                                                            NTUnicodeString
	HeapPartitionName                                                             NTUnicodeString
	DefaultThreadpoolCpuSetMasks                                                  uintptr
	DefaultThreadpoolCpuSetMaskCount                                              uint32
}

type PEB struct {
	reserved1              [2]byte
	BeingDebugged          byte
	BitField               byte
	reserved3              uintptr
	ImageBaseAddress       uintptr
	Ldr                    *PEB_LDR_DATA
	ProcessParameters      *RTL_USER_PROCESS_PARAMETERS
	reserved4              [3]uintptr
	AtlThunkSListPtr       uintptr
	reserved5              uintptr
	reserved6              uint32
	reserved7              uintptr
	reserved8              uint32
	AtlThunkSListPtr32     uint32
	reserved9              [45]uintptr
	reserved10             [96]byte
	PostProcessInitRoutine uintptr
	reserved11             [128]byte
	reserved12             [1]uintptr
	SessionId              uint32
}

type OBJECT_ATTRIBUTES struct {
	Length             uint32
	RootDirectory      Handle
	ObjectName         *NTUnicodeString
	Attributes         uint32
	SecurityDescriptor *SECURITY_DESCRIPTOR
	SecurityQoS        *SECURITY_QUALITY_OF_SERVICE
}

type IO_STATUS_BLOCK struct {
	Status      NTStatus
	Information uintptr
}

type RTLP_CURDIR_REF struct {
	RefCount int32
	Handle   Handle
}

type RTL_RELATIVE_NAME struct {
	RelativeName        NTUnicodeString
	ContainingDirectory Handle
	CurDirRef           *RTLP_CURDIR_REF
}

type PROCESS_BASIC_INFORMATION struct {
	ExitStatus                   NTStatus
	PebBaseAddress               *PEB
	AffinityMask                 uintptr
	BasePriority                 int32
	UniqueProcessId              uintptr
	InheritedFromUniqueProcessId uintptr
}

type SYSTEM_PROCESS_INFORMATION struct {
	NextEntryOffset              uint32
	NumberOfThreads              uint32
	WorkingSetPrivateSize        int64
	HardFaultCount               uint32
	NumberOfThreadsHighWatermark uint32
	CycleTime                    uint64
	CreateTime                   int64
	UserTime                     int64
	KernelTime                   int64
	ImageName                    NTUnicodeString
	BasePriority                 int32
	UniqueProcessID              uintptr
	InheritedFromUniqueProcessID uintptr
	HandleCount                  uint32
	SessionID                    uint32
	UniqueProcessKey             *uint32
	PeakVirtualSize              uintptr
	VirtualSize                  uintptr
	PageFaultCount               uint32
	PeakWorkingSetSize           uintptr
	WorkingSetSize               uintptr
	QuotaPeakPagedPoolUsage      uintptr
	QuotaPagedPoolUsage          uintptr
	QuotaPeakNonPagedPoolUsage   uintptr
	QuotaNonPagedPoolUsage       uintptr
	PagefileUsage                uintptr
	PeakPagefileUsage            uintptr
	PrivatePageCount             uintptr
	ReadOperationCount           int64
	WriteOperationCount          int64
	OtherOperationCount          int64
	ReadTransferCount            int64
	WriteTransferCount           int64
	OtherTransferCount           int64
}

type RTL_PROCESS_MODULE_INFORMATION struct {
	Section          Handle
	MappedBase       uintptr
	ImageBase        uintptr
	ImageSize        uint32
	Flags            uint32
	LoadOrderIndex   uint16
	InitOrderIndex   uint16
	LoadCount        uint16
	OffsetToFileName uint16
	FullPathName     [256]byte
}

type RTL_PROCESS_MODULES struct {
	NumberOfModules uint32
	Modules         [1]RTL_PROCESS_MODULE_INFORMATION
}

type ResourceIDOrString interface{} // ResourceIDOrString must be either a ResourceID, to specify a resource or resource type by ID,
// or a string, to specify a resource or resource type by name.

type VS_FIXEDFILEINFO struct {
	Signature        uint32
	StrucVersion     uint32
	FileVersionMS    uint32
	FileVersionLS    uint32
	ProductVersionMS uint32
	ProductVersionLS uint32
	FileFlagsMask    uint32
	FileFlags        uint32
	FileOS           uint32
	FileType         uint32
	FileSubtype      uint32
	FileDateMS       uint32
	FileDateLS       uint32
}

type COAUTHIDENTITY struct {
	User           *uint16
	UserLength     uint32
	Domain         *uint16
	DomainLength   uint32
	Password       *uint16
	PasswordLength uint32
	Flags          uint32
}

type COAUTHINFO struct {
	AuthnSvc           uint32
	AuthzSvc           uint32
	ServerPrincName    *uint16
	AuthnLevel         uint32
	ImpersonationLevel uint32
	AuthIdentityData   *COAUTHIDENTITY
	Capabilities       uint32
}

type COSERVERINFO struct {
	Reserved1 uint32
	Aame      *uint16
	AuthInfo  *COAUTHINFO
	Reserved2 uint32
}

type BIND_OPTS3 struct {
	CbStruct          uint32
	Flags             uint32
	Mode              uint32
	TickCountDeadline uint32
	TrackFlags        uint32
	ClassContext      uint32
	Locale            uint32
	ServerInfo        *COSERVERINFO
	Hwnd              HWND
}

type ModuleInfo struct {
	BaseOfDll   uintptr
	SizeOfImage uint32
	EntryPoint  uintptr
}

type Rect struct {
	Left   int32
	Top    int32
	Right  int32
	Bottom int32
}

type GUIThreadInfo struct {
	Size        uint32
	Flags       uint32
	Active      HWND
	Focus       HWND
	Capture     HWND
	MenuOwner   HWND
	MoveSize    HWND
	CaretHandle HWND
	CaretRect   Rect
}

type WSAQUERYSET struct {
	Size                uint32
	ServiceInstanceName *uint16
	ServiceClassId      *GUID
	Version             *WSAVersion
	Comment             *uint16
	NameSpace           uint32
	NSProviderId        *GUID
	Context             *uint16
	NumberOfProtocols   uint32
	AfpProtocols        *AFProtocols
	QueryString         *uint16
	NumberOfCsAddrs     uint32
	SaBuffer            *CSAddrInfo
	OutputFlags         uint32
	Blob                *BLOB
}

type WSAVersion struct {
	Version                 uint32
	EnumerationOfComparison int32
}

type AFProtocols struct {
	AddressFamily int32
	Protocol      int32
}

type CSAddrInfo struct {
	LocalAddr  SocketAddress
	RemoteAddr SocketAddress
	SocketType int32
	Protocol   int32
}

type BLOB struct {
	Size     uint32
	BlobData *byte
}

type ComStat struct {
	Flags    uint32
	CBInQue  uint32
	CBOutQue uint32
}

type DCB struct {
	DCBlength  uint32
	BaudRate   uint32
	Flags      uint32
	wReserved  uint16
	XonLim     uint16
	XoffLim    uint16
	ByteSize   uint8
	Parity     uint8
	StopBits   uint8
	XonChar    byte
	XoffChar   byte
	ErrorChar  byte
	EofChar    byte
	EvtChar    byte
	wReserved1 uint16
}

type WSAData struct {
	Version      uint16
	HighVersion  uint16
	Description  [WSADESCRIPTION_LEN + 1]byte
	SystemStatus [WSASYS_STATUS_LEN + 1]byte
	MaxSockets   uint16
	MaxUdpDg     uint16
	VendorInfo   *byte
}

type Servent struct {
	Name    *byte
	Aliases **byte
	Port    uint16
	Proto   *byte
}

type JOBOBJECT_BASIC_LIMIT_INFORMATION struct {
	PerProcessUserTimeLimit int64
	PerJobUserTimeLimit     int64
	LimitFlags              uint32
	MinimumWorkingSetSize   uintptr
	MaximumWorkingSetSize   uintptr
	ActiveProcessLimit      uint32
	Affinity                uintptr
	PriorityClass           uint32
	SchedulingClass         uint32
	_                       uint32
} // pad to 8 byte boundary

type pgkey struct{ program, key string }

type counterPtr struct {
	m     *mappedFile
	count *atomic.Uint64
}

type counterState struct{ bits atomic.Uint64 }

type mappedFile struct {
	meta      string
	hdrLen    uint32
	zero      [4]byte
	closeOnce sync.Once
	f         *os.File
	mapping   *mmap.Data
} // A mappedFile is a counter file mmapped into memory.
//
// The file layout for a mappedFile m is as follows:
//
//	offset, byte size:                 description
//	------------------                 -----------
//	0, hdrLen:                         header, containing metadata; see [mappedHeader]
//	hdrLen+limitOff, 4:                uint32 allocation limit (byte offset of the end of counter records)
//	hdrLen+hashOff, 4*numHash:         hash table, stores uint32 heads of a linked list of records, keyed by name hash
//	hdrLen+hashOff+4*numHash to limit: counter records: see record syntax below
//
// The record layout is as follows:
//
//	offset, byte size: description
//	------------------ -----------
//	0, 8:              uint64 counter value
//	8, 12:             uint32 name length
//	12, 16:            uint32 offset of next record in linked list
//	16, name length:   counter name

type stack struct {
	pcs     []uintptr
	counter *Counter
}

type UploadConfig struct {
	GOOS []string// An UploadConfig controls what data is uploaded.

	GOARCH     []string
	GoVersion  []string
	SampleRate float64
	Programs   []*ProgramConfig
}

type ProgramConfig struct {
	Name     string
	Versions []string// the counter names may have to be
	// repeated for each program. (e.g., if the counters are in a package
	// that is used in more than one program.)

	Counters []CounterConfig `json:",omitempty"`// versions present in a counterconfig

	Stacks []CounterConfig `json:",omitempty"`
}

type CounterConfig struct {
	Name  string
	Rate  float64
	Depth int `json:",omitempty"`
} // The "collapsed" counter: <chart>:{<bucket1>,<bucket2>,...}
// for stack counters

type ProgramReport struct {
	Program   string
	Version   string
	GoVersion string
	GOOS      string
	GOARCH    string
	Counters  map[ // Package path of the program.
	// Go version used to build the program.
	string]int64
	Stacks map[string]int64
}

type parsedCache struct {
	mu sync.Mutex
	m  map[ // avoid parsing count files multiple times
	string]*counter.File
}

type RunConfig struct {
	TelemetryDir string
	UploadURL    string
	LogWriter    io.Writer
	Env          []string// RunConfig configures non-default behavior of a call to Run.
	//
	// All fields are optional, for testing or observability.
	// if set, used for detailed logging of the upload process

	StartTime time.Time
} // if set, appended to the config download environment
// if set, overrides the upload start time

type uploader struct {
	config          *telemetry.UploadConfig
	configVersion   string
	dir             telemetry.Dir
	uploadServerURL string
	startTime       time.Time
	cache           parsedCache
	logFile         *os.File
	logger          *log.Logger
} // uploader encapsulates a single upload operation, carrying parameters and
// shared state.
// the telemetry dir to process

type StartResult struct{ wg sync.WaitGroup } // A StartResult is a handle to the result of a call to [Start]. Call
// [StartResult.Wait] to wait for the completion of all work done on behalf of
// Start.

type EscapeCodes struct {
	Black, Red, Green, Yellow, Blue, Magenta, Cyan, White []byte// EscapeCodes contains escape sequences that can be written to the terminal in
	// order to achieve different styles of text.
	// Foreground colors

	Reset []byte// Reset all attributes

}

type History interface {
	Add(entry string)
	Len() int
	At(idx int) string
} // A History provides a (possibly bounded) queue of input lines read by [Terminal.ReadLine].
// At returns an entry from the history.
// Index 0 is the most-recently added entry and
// index Len()-1 is the least-recently added entry.
// If index is < 0 or >= Len(), it panics.

type Terminal struct {
	AutoCompleteCallback func(line string, pos int, key rune) (newLine string, newPos int, ok bool)
	Escape               *EscapeCodes
	lock                 sync.Mutex
	c                    io.ReadWriter
	prompt               []rune// Terminal contains the state for running a VT100 terminal that is capable of
	// reading lines of input.
	// lock protects the terminal and the state in this object from
	// concurrent processing of a key press and a Write() call.

	line []rune// line is the current line being entered.

	pos                   int
	echo                  bool
	pasteActive           bool
	cursorX, cursorY      int
	maxLine               int
	termWidth, termHeight int
	outBuf                []byte// pos is the logical position of the cursor in line
	// outBuf contains the terminal data to be sent.

	remainder []byte// remainder contains the remainder of any partial key sequences after
	// a read. It aliases into inBuf.

	inBuf          [256]byte
	History        History
	historyIndex   int
	historyPending string
} // History records and retrieves lines of input read by [ReadLine] which
// a user can retrieve and navigate using the up and down arrow keys.
//
// It is not safe to call ReadLine concurrently with any methods on History.
//
// [NewTerminal] sets this to a default implementation that records the
// last 100 lines of input.
// When navigating up and down the history it's possible to return to
// the incomplete, initial line. That value is stored in
// historyPending.

type pasteIndicatorError struct{}

type stRingBuffer struct {
	entries []string// stRingBuffer is a ring buffer of strings.
	// entries contains max elements.

	max  int
	head int
	size int
} // head contains the index of the element most recently added to the ring.
// size contains the number of elements in the ring.

type Caser struct{ t transform.SpanningTransformer } // A Caser transforms given input to a certain case. It implements
// transform.Transformer.
//
// A Caser may be stateful and should therefore not be shared between
// goroutines.

type options struct {
	noLower          bool
	simple           bool
	ignoreFinalSigma bool
}

type context struct {
	dst, src []byte// A context is used for iterating over source bytes, fetching case info and
	// writing to a destination buffer.
	//
	// Casing operations may need more than one rune of context to decide how a rune
	// should be cased. Casing implementations should call checkpoint on context
	// whenever it is known to be safe to return the runes processed so far.
	//
	// It is recommended for implementations to not allow for more than 30 case
	// ignorables as lookahead (analogous to the limit in norm) and to use state if
	// unbounded lookahead is needed for cased runes.

	atEOF      bool
	pDst       int
	pSrc       int
	nDst, nSrc int
	err        error
	sz         int
	info       info
	isMidWord  bool
} // pDst points past the last written rune in dst.
// false if next cased letter needs to be title-cased.

type caseFolder struct{ transform.NopResetter }

type undUpperCaser struct{ transform.NopResetter }

type undLowerIgnoreSigmaCaser struct{ transform.NopResetter } // undLowerIgnoreSigmaCaser implements the Transformer interface for doing
// a lower case mapping for the root locale (und) ignoring final sigma
// handling. This casing algorithm is used in some performance-critical packages
// like secure/precis and x/net/http/idna, which warrants its special-casing.

type simpleCaser struct {
	context
	f    mapFunc
	span spanFunc
}

type undLowerCaser struct{ transform.NopResetter } // undLowerCaser implements the Transformer interface for doing a lower case
// mapping for the root locale (und) ignoring final sigma handling. This casing
// algorithm is used in some performance-critical packages like secure/precis
// and x/net/http/idna, which warrants its special-casing.

type lowerCaser struct {
	undLowerIgnoreSigmaCaser
	context
	first, midWord mapFunc
} // lowerCaser implements the Transformer interface. The default Unicode lower
// casing requires different treatment for the first and subsequent characters
// of a word, most notably to handle the Greek final Sigma.

type titleCaser struct {
	context
	title     mapFunc
	lower     mapFunc
	titleSpan spanFunc
	rewrite   func(*context)
} // titleCaser implements the Transformer interface. Title casing algorithms
// distinguish between the first letter of a word and subsequent letters of the
// same word. It uses state to avoid requiring a potentially infinite lookahead.
// rune mappings used by the actual casing algorithms.

type caseTrie struct{} // caseTrie. Total size: 11892 bytes (11.61 KiB). Checksum: c6f15484b7653775.

type valueRange struct {
	value  uint16
	lo, hi byte
} // valueRange is an entry in a sparse block.

type sparseBlocks struct {
	values  []valueRange
	offsets []uint16
}

type fullTag interface {
	IsRoot() bool
	Parent() language.Tag
}

type Variant struct {
	ID  uint8
	str string
} // Variant represents a registered variant of a language as defined by BCP 47.

type ValueError struct{ v [8]byte } // ValueError is returned by any of the parsing functions when the
// input is well-formed but the respective subtag is not recognized
// as a valid value.

type variantsSort struct {
	i []uint8
	v [][]byte
}

type bytesSort struct {
	b [][]byte
	n int
} // first n bytes to compare

type FromTo struct {
	From uint16
	To   uint16
}

type likelyLangRegion struct {
	lang   uint16
	region uint16
}

type likelyScriptRegion struct {
	region uint16
	script uint16
	flags  uint8
}

type likelyLangScript struct {
	lang   uint16
	script uint16
	flags  uint8
}

type likelyTag struct {
	lang   uint16
	region uint16
	script uint16
}

type parentRel struct {
	lang       uint16
	script     uint16
	maxScript  uint16
	toRegion   uint16
	fromRegion []uint16
}

type InheritanceMatcher struct{ index map[language.Tag]int }

type Coverage interface {
	Tags() []Tag// The Coverage interface is used to define the level of coverage of an
	// internationalization service. Note that not all types are supported by all
	// services. As lists may be generated on the fly, it is recommended that users
	// of a Coverage cache the results.
	// Tags returns the list of supported tags.

	BaseLanguages() []Base// BaseLanguages returns the list of supported base languages.

	Scripts() []Script// Scripts returns the list of supported scripts.

	Regions() []Region// Regions returns the list of supported regions.

}

type allSubtags struct{}

type coverage struct {
	tags func() []Tag// coverage is used by NewCoverage which is used as a convenient way for
	// creating Coverage implementations for partially defined data. Very often a
	// package will only need to define a subset of slices. coverage provides a
	// convenient way to do this. Moreover, packages using NewCoverage, instead of
	// their own implementation, will not break if later new slice types are added.

	bases   func() []Base
	scripts func() []Script
	regions func() []Region
}

type Extension struct{ s string } // Extension is a single BCP 47 extension.

type Matcher interface {
	Match(t ...Tag) (tag Tag, index int, c Confidence)
} // Matcher is the interface that wraps the Match method.
//
// Match returns the best match for any of the given tags, along with
// a unique index associated with the returned tag and a confidence
// score.

type matcher struct {
	default_  *haveTag
	supported []*// matcher keeps a set of supported language tags, indexed by language.
	haveTag
	index            map[language.Language]*matchHeader
	passSettings     bool
	preferSameScript bool
}

type matchHeader struct {
	haveTags []*// matchHeader has the lists of tags for exact matches and matches based on
	// maximized and canonicalized tags for a given language.
	haveTag
	original bool
}

type haveTag struct {
	tag       language.Tag
	index     int
	conf      Confidence
	maxRegion language.Region
	maxScript language.Script
	altScript language.Script
	nextMax   uint16
} // haveTag holds a supported Tag and its maximized script and region. The maximized
// or canonicalized language is not stored as it is not needed during matching.
// nextMax is the index of the next haveTag with the same maximized tags.

type bestMatch struct {
	have            *haveTag
	want            language.Tag
	conf            Confidence
	pinnedRegion    language.Region
	pinLanguage     bool
	sameRegionGroup bool
	origLang        bool
	origReg         bool
	paradigmReg     bool
	regGroupDist    uint8
	origScript      bool
} // bestMatch accumulates the best match so far.
// Cached results from applying tie-breaking rules.

type tagSort struct {
	tag []Tag
	q   []float32
}

type mutualIntelligibility struct {
	want     uint16
	have     uint16
	distance uint8
	oneway   bool
}

type scriptIntelligibility struct {
	wantLang   uint16
	haveLang   uint16
	wantScript uint8
	haveScript uint8
	distance   uint8
}

type regionIntelligibility struct {
	lang     uint16
	script   uint8
	group    uint8
	distance uint8
}

type Transformer interface {
	Transform(dst, src []byte,// Transformer transforms bytes.
	// Transform writes to dst the transformed bytes read from src, and
	// returns the number of dst bytes written and src bytes read. The
	// atEOF argument tells whether src represents the last bytes of the
	// input.
	//
	// Callers should always process the nDst bytes produced and account
	// for the nSrc bytes consumed before considering the error err.
	//
	// A nil error means that all of the transformed bytes (whether freshly
	// transformed from src or left over from previous Transform calls)
	// were written to dst. A nil error can be returned regardless of
	// whether atEOF is true. If err is nil then nSrc must equal len(src);
	// the converse is not necessarily true.
	//
	// ErrShortDst means that dst was too short to receive all of the
	// transformed bytes. ErrShortSrc means that src had insufficient data
	// to complete the transformation. If both conditions apply, then
	// either error may be returned. Other than the error conditions listed
	// here, implementations are free to report other errors that arise.
	atEOF bool) (nDst, nSrc int, err error)
	Reset()
} // Reset resets the state and allows a Transformer to be reused.

type SpanningTransformer interface {
	Transformer
	Span(src []byte,// SpanningTransformer extends the Transformer interface with a Span method
	// that determines how much of the input already conforms to the Transformer.
	// Span returns a position in src such that transforming src[:n] results in
	// identical output src[:n] for these bytes. It does not necessarily return
	// the largest such n. The atEOF argument tells whether src represents the
	// last bytes of the input.
	//
	// Callers should always account for the n bytes consumed before
	// considering the error err.
	//
	// A nil error means that all input bytes are known to be identical to the
	// output produced by the Transformer. A nil error can be returned
	// regardless of whether atEOF is true. If err is nil, then n must
	// equal len(src); the converse is not necessarily true.
	//
	// ErrEndOfSpan means that the Transformer output may differ from the
	// input after n bytes. Note that n may be len(src), meaning that the output
	// would contain additional bytes after otherwise identical output.
	// ErrShortSrc means that src had insufficient data to determine whether the
	// remaining bytes would change. Other than the error conditions listed
	// here, implementations are free to report other errors that arise.
	//
	// Calling Span can modify the Transformer state as a side effect. In
	// effect, it does the transformation just as calling Transform would, only
	// without copying to a destination buffer and only up to a point it can
	// determine the input and output bytes are the same. This is obviously more
	// limited than calling Transform, but can be more efficient in terms of
	// copying and allocating buffers. Calls to Span and Transform may be
	// interleaved.
	atEOF bool) (n int, err error)
}

type NopResetter struct{} // NopResetter can be embedded by implementations of Transformer to add a nop
// Reset method.

type nop struct{ NopResetter }

type discard struct{ NopResetter }

type chain struct {
	link []link// chain is a sequence of links. A chain with N Transformers has N+1 links and
	// N+1 buffers. Of those N+1 buffers, the first and last are the src and dst
	// buffers given to chain.Transform and the middle N-1 buffers are intermediate
	// buffers owned by the chain. The i'th link transforms bytes from the i'th
	// buffer chain.link[i].b at read offset chain.link[i].p to the i+1'th buffer
	// chain.link[i+1].b at write offset chain.link[i+1].n, for i in [0, N).

	err      error
	errStart int
} // errStart is the index at which the error occurred plus 1. Processing
// errStart at this level at the next call to Transform. As long as
// errStart > 0, chain will not consume any more source bytes.

type link struct {
	t Transformer
	b []byte// b[p:n] holds the bytes to be transformed by t.

	p int
	n int
}

type reorderBuffer struct {
	rune     [maxBufferSize]Properties
	byte     [maxByteBufferSize]byte
	nbyte    uint8
	ss       streamSafe
	nrune    int
	f        formInfo
	src      input
	nsrc     int
	tmpBytes input
	out      []byte// reorderBuffer is used to normalize a single segment.  Characters inserted with
	// insert are decomposed and reordered based on CCC. The compose method can
	// be used to recombine characters.  Note that the byte buffer does not hold
	// the UTF-8 characters in order.  Only the rune array is maintained in sorted
	// order. flush writes the resulting segment to a byte array.
	// Number of runeInfos.

	flushF func(*reorderBuffer) bool
}

type Properties struct {
	pos   uint8
	size  uint8
	ccc   uint8
	tccc  uint8
	nLead uint8
	flags qcInfo
	index uint16
} // Properties provides access to normalization properties of a rune.
// quick check flags

type formInfo struct {
	form                     Form
	composing, compatibility bool
	info                     lookupFunc
	nextMain                 iterFunc
} // formInfo holds Form-specific functions and tables.
// form type

type Iter struct {
	rb       reorderBuffer
	buf      [maxByteBufferSize]byte
	info     Properties
	next     iterFunc
	asciiF   iterFunc
	p        int
	multiSeg []byte// An Iter iterates over a string or byte slice, while normalizing it
	// to a given Form.
	// current position in input source

} // remainder of multi-segment decomposition

type normWriter struct {
	rb  reorderBuffer
	w   io.Writer
	buf []byte
}

type normReader struct {
	rb           reorderBuffer
	r            io.Reader
	inbuf        []byte
	outbuf       []byte
	bufStart     int
	lastBoundary int
	err          error
}

type nfcTrie struct{} // nfcTrie. Total size: 10442 bytes (10.20 KiB). Checksum: 4ba400a9d8208e03.

type nfkcTrie struct{} // nfkcTrie. Total size: 17104 bytes (16.70 KiB). Checksum: d985061cf5307b35.

type Bisect struct {
	Env []string// A Bisect holds the state for a bisect invocation.
	// Env is the additional environment variables for the command.
	// PATTERN and RANDOM are substituted in the values, but not the names.

	Cmd  string
	Args []string// Cmd is the command (program name) to run.
	// PATTERN and RANDOM are not substituted.
	// Args is the command arguments.
	// PATTERN and RANDOM are substituted anywhere they appear.

	Max     int
	MaxSet  int
	Timeout time.Duration
	Count   int
	Verbose bool
	Stdout  io.Writer
	Stderr  io.Writer
	TestRun func(env []string,// Command-line flags controlling bisect behavior.
	// where to write standard error (usually os.Stderr)
	cmd string, args []string) (out []byte, err error)
	Disable       bool
	SkipHexDigits int
	Add           []string// if non-nil, used instead of exec.Command
	// Add is a list of suffixes to add to every trial, because they
	// contain changes that are necessary for a group we are assembling.

	Skip []string// Skip is a list of suffixes that uniquely identify changes to exclude from every trial,
	// because they have already been used in failing change sets.
	// Suffixes later in the list may only be unique after removing
	// the ones earlier in the list.
	// Skip applies after Add.

}

type Result struct {
	Success  bool
	Cmd      string
	Out      string
	Suffix   string
	MatchIDs []uint64// A Result holds the result of a single target trial.
	// the suffix used for collecting MatchIDs, MatchText, and MatchFull

	MatchText []string// match IDs enabled during this trial

	MatchFull []string// match reports for the IDs, with match markers removed

} // full match lines for the IDs, with match markers kept

type ProfileBlock struct {
	StartLine, StartCol int
	EndLine, EndCol     int
	NumStmt, Count      int
} // ProfileBlock represents a single block of profiling data.

type Boundary struct {
	Offset int
	Start  bool
	Count  int
	Norm   float64
	Index  int
} // Boundary represents the position in a source file of the beginning or end of a
// block as reported by the coverage profile. In HTML mode, it will correspond to
// the opening or closing of a <span> tag and will be used to colorize the source
// Order in input file.

type Analyzer struct {
	Name             string
	Doc              string
	URL              string
	Flags            flag.FlagSet
	Run              func(*Pass) (any, error)
	RunDespiteErrors bool
	Requires         []*// An Analyzer describes an analysis function and its options.
	// Requires is a set of analyzers that must run successfully
	// before this one on a given package. This analyzer may inspect
	// the outputs produced by each analyzer in Requires.
	// The graph over analyzers implied by Requires edges must be acyclic.
	//
	// Requires establishes a "horizontal" dependency between
	// analysis passes (different analyzers, same package).
	Analyzer
	ResultType reflect.Type
	FactTypes  []Fact// ResultType is the type of the optional result of the Run function.
	// FactTypes indicates that this analyzer imports and exports
	// Facts of the specified concrete types.
	// An analyzer that uses facts may assume that its import
	// dependencies have been similarly analyzed before it runs.
	// Facts must be pointers.
	//
	// FactTypes establishes a "vertical" dependency between
	// analysis passes (same analyzer, different packages).

}

type Pass struct {
	Analyzer *Analyzer
	Fset     *token.FileSet
	Files    []*// A Pass provides information to the Run function that
	// applies a specific analyzer to a single Go package.
	//
	// It forms the interface between the analysis logic and the driver
	// program, and has both input and an output components.
	//
	// As in a compiler, one pass may depend on the result computed by another.
	//
	// The Run function should not call any of the Pass functions concurrently.
	// file position information; Run may add new files
	ast.File
	OtherFiles []string// the abstract syntax tree of each file

	IgnoredFiles []string// names of non-Go files of this package

	Pkg        *types.Package
	TypesInfo  *types.Info
	TypesSizes types.Sizes
	TypeErrors []types.// names of ignored source files in this package
	// function for computing sizes of types
	Error
	Module   *Module
	Report   func(Diagnostic)
	ResultOf map[ // type errors (only if Analyzer.RunDespiteErrors)
	// ResultOf provides the inputs to this analysis pass, which are
	// the corresponding results of its prerequisite analyzers.
	// The map keys are the elements of Analysis.Required,
	// and the type of each corresponding value is the required
	// analysis's ResultType.
	*Analyzer]any
	ReadFile func(filename string) ([]byte,// ReadFile returns the contents of the named file.
	//
	// The only valid file names are the elements of OtherFiles
	// and IgnoredFiles, and names returned by
	// Fset.File(f.FileStart).Name() for each f in Files.
	//
	// Analyzers must use this function (if provided) instead of
	// accessing the file system directly. This allows a driver to
	// provide a virtualized file tree (including, for example,
	// unsaved editor buffers) and to track dependencies precisely
	// to avoid unnecessary recomputation.
	error)
	ImportObjectFact  func(obj types.Object, fact Fact) bool
	ImportPackageFact func(pkg *types.Package, fact Fact) bool
	ExportObjectFact  func(obj types.Object, fact Fact)
	ExportPackageFact func(fact Fact)
	AllPackageFacts   func() []PackageFact// ImportObjectFact retrieves a fact associated with obj.
	// Given a value ptr of type *T, where *T satisfies Fact,
	// ImportObjectFact copies the value to *ptr.
	//
	// ImportObjectFact panics if called after the pass is complete.
	// ImportObjectFact is not concurrency-safe.
	// AllPackageFacts returns a new slice containing all package
	// facts of the analysis's FactTypes in unspecified order.
	// See comments for AllObjectFacts.

	AllObjectFacts func() []ObjectFact// AllObjectFacts returns a new slice containing all object
	// facts of the analysis's FactTypes in unspecified order.
	//
	// The result includes all facts exported by packages
	// whose symbols are referenced by the current package
	// (by qualified identifiers or field/method selections).
	// And it includes all facts exported from the current
	// package by the current analysis pass.

}

type PackageFact struct {
	Package *types.Package
	Fact    Fact
} // PackageFact is a package together with an associated fact.

type ObjectFact struct {
	Object types.Object
	Fact   Fact
} // ObjectFact is an object together with an associated fact.

type Fact interface {
	AFact()
} // A Fact is an intermediate fact produced during analysis.
//
// Each fact is associated with a named declaration (a types.Object) or
// with a package as a whole. A single object or package may have
// multiple associated facts, but only one of any particular fact type.
//
// A Fact represents a predicate such as "never returns", but does not
// represent the subject of the predicate such as "function F" or "package P".
//
// Facts may be produced in one analysis pass and consumed by another
// analysis pass even if these are in different address spaces.
// If package P imports Q, all facts about Q produced during
// analysis of that package will be available during later analysis of P.
// Facts are analogous to type export data in a build system:
// just as export data enables separate compilation of several passes,
// facts enable "separate analysis".
//
// Each pass (a, p) starts with the set of facts produced by the
// same analyzer a applied to the packages directly imported by p.
// The analysis may add facts to the set, and they may be exported in turn.
// An analysis's Run function may retrieve facts by calling
// Pass.Import{Object,Package}Fact and update them using
// Pass.Export{Object,Package}Fact.
//
// A fact is logically private to its Analysis. To pass values
// between different analyzers, use the results mechanism;
// see Analyzer.Requires, Analyzer.ResultType, and Pass.ResultOf.
//
// A Fact type must be a pointer.
// Facts are encoded and decoded using encoding/gob.
// A Fact may implement the GobEncoder/GobDecoder interfaces
// to customize its encoding. Fact encoding should not fail.
//
// A Fact should not be modified once exported.
// dummy method to avoid type errors

type RelatedInformation struct {
	Pos     token.Pos
	End     token.Pos
	Message string
} // RelatedInformation contains information related to a diagnostic.
// For example, a diagnostic that flags duplicated declarations of a
// variable may include one RelatedInformation per existing
// declaration.
// optional

type SuggestedFix struct {
	Message   string
	TextEdits []TextEdit// A SuggestedFix is a code change associated with a Diagnostic that a
	// user can choose to apply to their code. Usually the SuggestedFix is
	// meant to fix the issue flagged by the diagnostic.
	//
	// The TextEdits must not overlap, nor contain edits for other
	// packages. Edits need not be totally ordered, but the order
	// determines how insertions at the same point will be applied.
	// A verb phrase describing the fix, to be shown to
	// a user trying to decide whether to accept it.
	//
	// Example: "Remove the surplus argument"

}

type TextEdit struct {
	Pos     token.Pos
	End     token.Pos
	NewText []byte// A TextEdit represents the replacement of the code between Pos and End with the new text.
	// Each TextEdit should apply to a single file. End should not be earlier in the file than Pos.
	// For a pure insertion, End can either be set to Pos or token.NoPos.

}

type JSONTextEdit struct {
	Filename string `json:"filename"`
	Start    int    `json:"start"`
	End      int    `json:"end"`
	New      string `json:"new"`
} // A TextEdit describes the replacement of a portion of a file.
// Start and End are zero-based half-open indices into the original byte
// sequence of the file, and New is the new text.

type JSONSuggestedFix struct {
	Message string         `json:"message"`
	Edits   []JSONTextEdit `json:"edits"`// A JSONSuggestedFix describes an edit that should be applied as a whole or not
	// at all. It might contain multiple TextEdits/text_edits if the SuggestedFix
	// consists of multiple non-contiguous edits.

}

type JSONDiagnostic struct {
	Category       string             `json:"category,omitempty"`
	Posn           string             `json:"posn"`
	Message        string             `json:"message"`
	SuggestedFixes []JSONSuggestedFix `json:"suggested_fixes,omitempty"`// A JSONDiagnostic describes the JSON schema of an analysis.Diagnostic.
	//
	// TODO(matloob): include End position if present.
	// e.g. "file.go:line:column"

	Related []JSONRelatedInformation `json:"related,omitempty"`
}

type JSONRelatedInformation struct {
	Posn    string `json:"posn"`
	Message string `json:"message"`
} // A JSONRelated describes a secondary position and message related to
// a primary diagnostic.
//
// TODO(adonovan): include End position if present.
// e.g. "file.go:line:column"

type asmArch struct {
	name      string
	bigEndian bool
	stack     string
	lr        bool
	retRegs   []string// An asmArch describes assembly parameters for an architecture
	// retRegs is a list of registers for return value in register ABI (ABIInternal).
	// For now, as we only check whether we write to any result, here we only need to
	// include the first integer register and first floating-point register. Accessing
	// any of them counts as writing to result.

	writeResult []string// writeResult is a list of instructions that will change result register implicitly.

	sizes    types.Sizes
	intSize  int
	ptrSize  int
	maxAlign int
} // calculated during initialization

type asmFunc struct {
	arch *asmArch
	size int
	vars map[ // An asmFunc describes the expected variables for a function on a given architecture.
	// size of all arguments
	string]*asmVar
	varByOffset map[int]*asmVar
}

type asmVar struct {
	name  string
	kind  asmKind
	typ   string
	off   int
	size  int
	inner []*// An asmVar describes a single assembly variable.
	asmVar
}

type component struct {
	size   int
	offset int
	kind   asmKind
	typ    string
	suffix string
	outer  string
} // A component is an assembly-addressable component of a composite type,
// or a composite type itself.
// The suffix for immediately containing composite type.

type boolOp struct {
	name  string
	tok   token.Token
	badEq token.Token
} // token corresponding to this operator
// token corresponding to the equality test that should not be used with this operator

type checker struct {
	pass         *analysis.Pass
	plusBuildOK  bool
	goBuildOK    bool
	crossCheck   bool
	inStar       bool
	goBuildPos   token.Pos
	plusBuildPos token.Pos
	goBuild      constraint.Expr
	plusBuild    constraint.Expr
} // "+build" lines still OK
// AND of +build constraints found

type noReturn struct{} // noReturn is a fact indicating that a function does not return.

type CFGs struct {
	defs map[ // A CFGs holds the control-flow graphs
	// for all the functions of the current package.
	*ast.Ident]types.Object
	funcDecls map[ // from Pass.TypesInfo.Defs
	*types.Func]*declInfo
	funcLits map[*ast.FuncLit]*litInfo
	pass     *analysis.Pass
} // transient; nil after construction

type litInfo struct {
	cfg      *cfg.CFG
	noReturn bool
}

type isWrapper struct{ Kind Kind } // isWrapper is a fact indicating that a function is a print or printf wrapper.

type printfWrapper struct {
	obj     *types.Func
	fdecl   *ast.FuncDecl
	format  *types.Var
	args    *types.Var
	callers []printfCaller
	failed  bool
} // if true, not a printf wrapper

type printfCaller struct {
	w    *printfWrapper
	call *ast.CallExpr
}

type printVerb struct {
	verb  rune
	flags string
	typ   printfArgType
} // User may provide verb through Formatter; could be a rune.
// known flags are all ASCII

type argMatcher struct {
	t    printfArgType
	seen map[ // argMatcher recursively matches types against the printfArgType t.
	//
	// To short-circuit recursion, it keeps track of types that have already been
	// matched (or are in the process of being matched) via the seen map. Recursion
	// arises from the compound types {map,chan,slice} which may be printed with %d
	// etc. if that is appropriate for their element types, as well as from type
	// parameters, which are expanded to the constituents of their type set.
	//
	// The reason field may be set to report the cause of the mismatch.
	types.Type]bool
	reason string
}

type uniqueName struct {
	key   string
	name  string
	level int
} // "xml" or "json"
// anonymous struct nesting level

type asyncCall struct {
	region ast.Node
	async  ast.Node
	scope  ast.Node
	fun    ast.Expr
} // asyncCall describes a region of code that needs to be checked for
// t.Forbidden() calls as it is started asynchronously from an async
// node go fun() or t.Run(name, fun).
// fun in go fun() or t.Run(name, fun)

type commentMetadata struct {
	isOutput bool
	pos      token.Pos
}

type deadState struct {
	pass        *analysis.Pass
	hasBreak    map[ast.Stmt]bool
	hasGoto     map[string]bool
	labels      map[string]ast.Stmt
	breakTarget ast.Stmt
	reachable   bool
}

type result struct {
	a           *analysis.Analyzer
	diagnostics []analysis.Diagnostic
	err         error
}

type CycleInRequiresGraphError struct{ AnalyzerNames map[string]bool }

type fieldInfo struct {
	nodeType  reflect.Type
	name      string
	index     int
	fieldType reflect.Type
} // pointer-to-struct type of ast.Node implementation

type Inspector struct {
	events []event// An Inspector provides methods for inspecting
	// (traversing) the syntax trees of a package.
}

type item struct {
	index            int32
	parentIndex      int32
	typAccum         uint64
	edgeKindAndIndex int32
} // index of current node's push event
// edge.Kind and index, bit packed

type lblock struct {
	_goto     *Block
	_break    *Block
	_continue *Block
} // Destinations associated with a labeled block.
// We populate these as labels are encountered in forward gotos or
// labeled statements.

type CFG struct {
	fset   *token.FileSet
	Blocks []*// A CFG represents the control-flow graph of a single function.
	//
	// The entry point is Blocks[0]; there may be multiple return blocks.
	Block
} // block[0] is entry; order otherwise undefined

type Encoder struct {
	scopeMemo map[ // An Encoder amortizes the cost of encoding the paths of multiple objects.
	// The zero value of an Encoder is ready to use.
	*types.Scope][]types.Object
} // memoization of scopeObjects

type finder struct {
	obj             types.Object
	seenTParamNames map[ // finder closes over search state for a call to find.
	// the sought object
	*types.TypeName]bool
	seenMethods map[ // for cycle breaking through type parameters
	*types.Func]bool
} // for cycle breaking through recursive interfaces

type entry struct {
	key   types.Type
	value any
} // entry is an entry (key/value association) in a hash bucket.

type Hasher struct{} // A Hasher provides a [Hasher.Hash] method to map a type to its hash value.
// Hashers are stateless, and all are equivalent.

type hasher struct{ inGenericSig bool } // hasher holds the state of a single Hash traversal: whether we are
// inside the signature of a generic function; this is used to
// optimize [hasher.hashTypeParam].

type MethodSetCache struct {
	mu    sync.Mutex
	named map[ // A MethodSetCache records the method set of each type T for which
	// MethodSet(T) is called so that repeat queries are fast.
	// The zero value is a ready-to-use cache instance.
	*types.Named]struct{ value, pointer *types.MethodSet }
	others map[ // method sets for named N and *N
	types.Type]*types.MethodSet
} // all other types

type tokenRange struct{ StartPos, EndPos token.Pos } // tokenRange is an implementation of the [analysis.Range] interface.

type Directive struct {
	Pos  token.Pos
	Tool string
	Name string
	Args string
} // A directive is a comment line with special meaning to the Go
// toolchain or another tool. It has the form:
//
//	//tool:name args
//
// The "tool:" portion is missing for the three directives named
// line, extern, and export.
//
// See https://go.dev/doc/comment#Syntax for details of Go comment
// syntax and https://pkg.go.dev/cmd/compile#hdr-Compiler_Directives
// for details of directives used by the Go compiler.
// may contain internal spaces

type cond struct {
	mask   uint64
	bits   uint64
	result bool
} // A cond is a single condition in the matcher.
// Given an input id, if id&mask == bits, return the result.

type Set struct {
	pkg *types.Package
	mu  sync.Mutex
	m   map[ // A Set is a set of analysis.Facts.
	//
	// Decode creates a Set of facts by reading from the imports of a given
	// package, and Encode writes out the set. Between these operation,
	// the Import and Export methods will query and update the set.
	//
	// All of Set's methods except String are safe to call concurrently.
	key]analysis.Fact
}

type key struct {
	pkg *types.Package
	obj types.Object
	t   reflect.Type
} // (object facts only)

type gobFact struct {
	PkgPath string
	Object  objectpath.Path
	Fact    analysis.Fact
} // gobFact is the Gob declaration of a serialized fact.
// type and value of user-defined Fact

type Decoder struct {
	pkg        *types.Package
	getPackage GetPackageFunc
} // A Decoder decodes the facts from the direct imports of the package
// provided to NewEncoder. A single decoder may be used to decode
// multiple fact sets (e.g. each for a different set of fact types)
// for the same package. Each call to Decode returns an independent
// fact set.

type Operation struct {
	Text  string
	Verb  Verb
	Range Range
	Flags string
	Width Size
	Prec  Size
} // Operation holds the parsed representation of a printf operation such as "%3.*[4]d".
// It is constructed by [Parse].
// precision specifier, e.g., '.4' in '%.4f'

type Size struct {
	Fixed   int
	Dynamic int
	Index   int
	Range   Range
} // Size describes an optional width or precision in a format operation.
// It may represent no value, a literal number, an asterisk, or an indexed asterisk.
// position of the size specifier within the operation

type Verb struct {
	Verb     rune
	Range    Range
	Index    int
	ArgIndex int
} // Verb represents the verb character of a format operation (e.g., 'd', 's', 'f').
// It also includes positional information and any explicit argument indexing.
// argument index (0-based) associated with this verb, relative to CallExpr

type pkginfo struct {
	name string
	deps string
} // list of indices of dependencies, as varint-encoded deltas

type Symbol struct {
	Name      string
	Kind      Kind
	Version   Version
	Signature string
} // Go version that first included the symbol
// if Kind == stdlib.Func

type Free struct {
	seen map[ // Free is a memoization of the set of free type parameters within a
	// type. It makes a sequence of calls to [Free.Has] for overlapping
	// types more efficient. The zero value is ready for use.
	//
	// NOTE: Adapted from go/types/infer.go. If it is later exported, factor.
	types.Type]bool
}

type termSet struct {
	complete bool
	terms    termlist
} // A termSet holds the normalized set of terms for a given type.
//
// The name termSet is intentionally distinct from 'type set': a type set is
// all types that implement a type (and includes method restrictions), whereas
// a term set just represents the structural restrictions on a type.

type uses struct {
	code []byte// A uses holds the list of Cursors of Idents that use a given symbol.
	//
	// The Uses map of [types.Info] is substantial, so it pays to compress
	// its inverse mapping here, both in space and in CPU due to reduced
	// allocation. A Cursor is 2 words; a Cursor.Index is 4 bytes; but
	// since Cursors are naturally delivered in ascending order, we can
	// use varint-encoded deltas at a cost of only ~1.7-2.2 bytes per use.
	//
	// Many variables have only one or two uses, so their encoded uses may
	// fit in the 4 bytes of initial, saving further CPU and space
	// essentially for free since the struct's size class is 4 words.

	last    int32
	initial [4]byte
} // varint-encoded deltas of successive Cursor.Index values
// use slack in size class as initial space for code

type NamedOrAlias interface {
	types.Type
	Obj() *types.TypeName
	TypeArgs() *types.TypeList
	TypeParams() *types.TypeParamList
	SetTypeParams(tparams []*// A NamedOrAlias is a [types.Type] that is named (as
	// defined by the spec) and capable of bearing type parameters: it
	// abstracts aliases ([types.Alias]) and defined types
	// ([types.Named]).
	//
	// Every type declared by an explicit "type" declaration is a
	// NamedOrAlias. (Built-in type symbols may additionally
	// have type [types.Basic], which is not a NamedOrAlias,
	// though the spec regards them as "named"; see [TypeNameFor].)
	//
	// NamedOrAlias cannot expose the Origin method, because
	// [types.Alias.Origin] and [types.Named.Origin] have different
	// (covariant) result types; use [Origin] instead.
	types.TypeParam)
}

type gover struct {
	major string
	minor string
	patch string
	kind  string
	pre   string
} // A gover is a parsed Go gover: major[.Minor[.Patch]][kind[pre]]
// The numbers are the original decimal strings to avoid integer overflows
// and since there is very little actual math. (Probably overflow doesn't matter in practice,
// but at the time this code was written, there was an existing test that used
// go1.99999999999, which does not fit in an int on 32-bit platforms.
// The "big decimal" representation avoids the problem entirely.)
// decimal or ""

type ThematicBreak struct {
	Position
	raw string
}

type HardBreak struct{}

type SoftBreak struct{}

type CodeBlock struct {
	Position
	Fence string
	Info  string
	Text  []string
}

type preBuilder struct {
	indent string
	text   []string// For indented code blocks.

}

type fenceBuilder struct {
	fence string
	info  string
	n     int
	text  []string
}

type Heading struct {
	Position
	Level int
	Text  *Text
	ID    string
} // The HTML id attribute. The parser populates this field if
// [Parser.HeadingIDs] is true and the heading ends with text like "{#id}".

type HTMLBlock struct {
	Position
	Text []string
}

type htmlBuilder struct {
	endBlank bool
	text     []string
	endFunc  func(string) bool
}

type HTMLTag struct{ Text string }

type Plain struct{ Text string }

type openPlain struct {
	Plain
	i int
} // position in input where bracket is

type emphPlain struct {
	Plain
	canOpen  bool
	canClose bool
	i        int
	n        int
} // position in output where emph is
// length of original span

type Escaped struct{ Plain }

type Code struct{ Text string }

type Strong struct {
	Marker string
	Inner  []Inline
}

type Del struct {
	Marker string
	Inner  []Inline
}

type Emph struct {
	Marker string
	Inner  []Inline
}

type backtickParser struct {
	last    [maxBackticks]int
	scanned bool
}

type Emoji struct {
	Name string
	Text string
} // emoji :name:, including colons
// Unicode for emoji sequence

type AutoLink struct {
	Text string
	URL  string
}

type Image struct {
	Inner     []Inline
	URL       string
	Title     string
	TitleChar byte
	corner    bool
}

type validDomainChecker struct {
	s   string
	cut int
} // before this index, no valid domains

type Item struct {
	Position
	Blocks []Block
	width  int
}

type listBuilder struct {
	bullet rune
	num    int
	loose  bool
	item   *itemBuilder
	todo   func() line
}

type Task struct{ Checked bool }

type Empty struct{ Position }

type Paragraph struct {
	Position
	Text *Text
}

type paraBuilder struct {
	text  []string
	table *tableBuilder
}

type mdState struct {
	prefix  string
	prefix1 string
	bullet  rune
	num     int
} // for first line only
// for numbered list items

type buildState interface {
	blocks() []Block
	pos() Position
	last() Block
	deleteLast()
	link(label string) *Link
	defineLink(label string, link *Link)
	newText(pos Position, text string) *Text
}

type blockBuilder interface {
	extend(p *parseState, s line) (line, bool)
	build(buildState) Block
}

type openBlock struct {
	builder blockBuilder
	inner   []Block
	pos     Position
}

type itemBuilder struct {
	list        *listBuilder
	width       int
	haveContent bool
}

type Text struct {
	Position
	Inline []Inline
	raw    string
}

type rootBuilder struct{}

type Document struct {
	Position
	Blocks []Block
	Links  map[string]*Link
}

type parseState struct {
	*Parser
	root      *Document
	links     map[string]*Link
	lineno    int
	stack     []openBlock
	lineDepth int
	corner    bool
	s         string
	emitted   int
	list      []Inline// noticed corner case to ignore in cross-implementation testing
	// s[:emitted] has been emitted into list

	lists []*// for fixup at end
	List
	texts     []*Text
	backticks backtickParser
}

type Quote struct {
	Position
	Blocks []Block
}

type quoteBuilder struct{}

type tableBuilder struct {
	hdr   tableTrimmed
	delim tableTrimmed
	rows  []tableTrimmed
}

type Table struct {
	Position
	Header []*Text
	Align  []string
	Rows   [][// 'l', 'c', 'r' for left, center, right; 0 for unset
	]*Text
}

type ST struct {
	x int
	l []int
}

type someStruct struct{}

type embeddedStringer struct {
	foo string
	ptrStringer
	bar int
}

type notstringer struct{ f float64 }

type percentDStruct struct {
	a int
	b []byte// A data type we can print with "%d".

	c *float64
}

type notPercentDStruct struct {
	a int
	b []byte// A data type we cannot print correctly with "%d".

	c bool
}

type percentSStruct struct {
	a string
	b []byte// A data type we can print with "%s".

	C stringerarray
}

type RecursiveStruct struct{ next *RecursiveStruct }

type RecursiveStruct1 struct{ next *RecursiveStruct2 }

type RecursiveStruct2 struct{ next *RecursiveStruct1 }

type unexportedInterface struct{ f interface{} }

type unexportedStringer struct{ t ptrStringer } // Issue 17798: unexported ptrStringer cannot be formatted.

type unexportedStringerOtherFields struct {
	s string
	t ptrStringer
	S string
}

type unexportedError struct{ e error } // Issue 17798: unexported error cannot be formatted.

type unexportedErrorOtherFields struct {
	s string
	e error
	S string
}

type errorer struct{}

type unexportedCustomError struct{ e errorer }

type errorInterface interface {
	error
	ExtraMethod()
}

type unexportedErrorInterface struct{ e errorInterface }

type StructTagTest struct {
	A int "hello"
} // ERROR "`hello` not compatible with reflect.StructTag.Get: bad syntax for struct tag pair"
