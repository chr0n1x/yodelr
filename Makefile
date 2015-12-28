default: compile

install-deps:
	go get github.com/tools/godep
	go get -u github.com/jteeuwen/go-bindata/...
	godep restore
	go-bindata -pkg templates -o templates/bindata.go templates/...

compile: install-deps
	go build -o ./yodelr
