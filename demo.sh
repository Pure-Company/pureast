./pureast \
  -file $GOROOT/src/net/http \
  -symbol Accept,AddrFromSlice,AllowQuerySemicolons,ArgServer,CIDRMask,CRAMMD5Auth,CachedCanonicalHeader,CanonicalHeader,CanonicalHeaderKey,CanonicalMIMEHeaderKey,ChanCreate \
  -structs \
  -output extracted.go
#  -batch 10 \

#  -minimal \
#  -deps \
#  -methods \
#  -report \
#  -dot \
#  -all-types \
#  -structs \
#  -interfaces \
#  -list-symbols \
#  -group \
#  -search \
#  -pattern "Handler" \
#  -index \
#  -index-path ".pureast-index.json" \
#  -types-summary

#
#./pureast \
#  -file ./pkg/mcp \
#  -symbol "Server" \
#  -output extracted.go \
#  -workers 12 \
#  -batch 10 \
#  -recursive \
#  -minimal \
#  -deps \
#  -methods \
#  -report \
#  -dot \
#  -all-types \
#  -structs \
#  -interfaces \
#  -list-symbols \
#  -group \
#  -search \
#  -pattern "Handler" \
#  -index \
#  -index-path ".pureast-index.json" \
#  -types-summary
#
#
#
## All types for LLM
#./pureast \
##  -file ./pkg/mcp \
#  -file ./$GOROOT/src/net/http \
#  -file ./$GOROOT/src/net/ \
#  -all-types \
#  -output types.go
#
#
## Extract with deps
#./pureast \
#  -file ./pkg/mcp \
#  -symbol "Server" \
#  -output extracted.go \
#  -deps
#
## Search
#./pureast \
#  -file ./pkg/mcp \
#  -search \
#  -pattern "Handler"
#
## List symbols
#./pureast \
#  -file ./pkg/mcp \
#  -list-symbols \
#  -group
#
## Dependency graph
#./pureast \
#  -file ./pkg/mcp \
#  -symbol "Server" \
#  -dot \
#  -output server.dot
#
## Full analysis
#./pureast \
#  -file ./pkg/mcp \
#  -symbol "Server" \
#  -deps \
#  -methods \
#  -report \
#  -output report.txt
#
#
