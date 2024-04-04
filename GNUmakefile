.POSIX:
.SUFFIXES:
.PHONY: all clean install check
all:
PROJECT=coingecko
VERSION=1.0.0
PREFIX=/usr/local

## -- BLOCK:go --
build/coingecko$(EXE):
	mkdir -p build
	go build -o $@ $(GO_CONF) ./cmd/coingecko
all: build/coingecko$(EXE)
install: all
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp build/coingecko$(EXE) $(DESTDIR)$(PREFIX)/bin
clean:
	rm -f build/coingecko$(EXE)
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: 
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp LICENSE $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
## -- BLOCK:license --
