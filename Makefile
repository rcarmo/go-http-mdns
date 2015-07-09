export GOPATH:=$(shell pwd)/vendor
export PATH:=$(PATH):$(GOPATH)/bin
BINARY=web

$(BINARY): *.go
	go build -o $(BINARY)

deps:
	mkdir -p vendor
	go get github.com/zenazn/goji
	go get github.com/hashicorp/mdns
	# save dependencies for buildpack
	#godep save github.com/zenazn/goji

clean:
	rm -f $(BINARY) $(BINARY)-linux $(BINARY)-rpi $(BINARY)-syno
	go fmt *.go

bootstrap:
	go get github.com/tools/godep

linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY)-linux

container: linux
	docker build -t rcarmo/go-http-mdns .

pi:
	GOARCH=arm GOARM=6 go build -o $(BINARY)-rpi

synology:
	GOARCH=arm GOARM=5 go build -o $(BINARY)-syno

print-%: ; @echo $*=$($*)
