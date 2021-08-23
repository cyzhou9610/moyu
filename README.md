rsrc -manifest moyu.manifest -o moyu.syso

go build -ldflags="-H windowsgui -w -s"
