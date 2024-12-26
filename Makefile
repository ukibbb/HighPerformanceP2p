# GOROOT directory of go instalation - standard library and compiler - default /usr/local/go
# GOPATH workspace dir for go projects and deps 
# - contains src (projects source codes) bin (executable binaries created bo `go install` pkg (compiled package objects. ) )
# Defaults to $HOME/go if not set. 

# GOBIN specifies where compiled binaries should be installed overrides $GOPATH/bin
# GOOS - specifies target operating system for build (linux, darwin, windows)
# GOARCH - specifies target CPU architecture (amd64, arm, amr64, 386)

# CGO_ENABLED - determines whether Cgo (C language interoperation) is enabled. controles whether go can use C libraries. 
# Disable for pure go build (usefull for cross-comiling)

# GOMOD - path to the current go.mod file or off if the projet isn't module mode. Indicates if project is using Go modules.
# GOPROXY - url of go module proxy, specifies how go retrivies module from remote sources.
# GOSUMDB - checksum database fo go modules, ensures integrity and authenticity
# GODEBUG - enables debugging options for the go runtime, used for performance tuning, diagnostic and experimentation
#	# gctrace=1: Enables garbage collection tracing.
#	# schedtrace=1000: Enables scheduler tracing every 1 ms.
#	# cgocheck=2: Enables stricter Cgo pointer checks.

# GOTMPDIR - directory for temporary files create by the go toolchain, overrides the default system directory
# GOVERSION - go version in use. indicates the version of go being used for the build


build:
	@go build -o $(HOME)/go/bin/p2p 

run: build
	$(HOME)/go/bin/p2p

test: 
	@go test ./... --race -v

bench_test:
	@go test ./... -bench=.

docker_build:
	docker build -t p2pr .

docker_run: docker_build
	docker run --name p2p -p 6379:6379 p2p

docker_stop:
	docker stop p2p && docker remove p2p 
