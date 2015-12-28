default: run

install-deps:
	go get github.com/tools/godep
	godep restore

run: install-deps
	go run yodelr.go ${ARGS}

install-dev: install-deps
	go build
