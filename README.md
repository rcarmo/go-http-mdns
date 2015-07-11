# go-http-mdns

This is a very simple HTTP daemon for testing mDNS service discovery.

It was initially developed for use as a test Docker container (since 
Docker [doesn't support multicast by default](https://github.com/docker/docker/issues/3043), 
I wanted a quick and lightweight way to test custom network setups).

(Testing is easy enough - the container will fail to start if you don't have any real multicast support.)

## Requirements

You need to have Go 1.4+ installed. 

## Quick Start

```bash
make bootstrap; make deps; make
./web
```

## Using the Container

```bash
make container
docker run -ti rcarmo/go-http-mdns
```

## Building under Mac OS X

If you're on a Mac and have Go from [Homebrew](http://brew.sh), you need to enable cross-compiling for Linux like so:

```bash
cd /usr/local/opt/go/libexec/src/
GOOS=linux GOARCH=amd64 ./make.bash --no-clean
```
