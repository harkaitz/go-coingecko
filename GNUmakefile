.POSIX:
.SUFFIXES:
.PHONY: all clean install check
PROJECT    =coingecko
VERSION    =1.0.1
PREFIX     =/usr/local
BUILDDIR  ?=.build
UNAME_S   ?=$(shell uname -s)
EXE       ?=$(shell uname -s | awk '/Windows/ || /MSYS/ || /CYG/ { print ".exe" }')
TOOLCHAINS =x86_64-linux-gnu x86_64-w64-mingw32

all:
clean:
install:
check:

release:
	mkdir -p $(BUILDDIR)
	hrelease -w github -t "$(TOOLCHAINS)" -N $(PROJECT) -R $(VERSION) -o $(BUILDDIR)/Release
	gh release create v$(VERSION) $$(cat $(BUILDDIR)/Release)

## -- BLOCK:go --
.PHONY: all-go install-go clean-go $(BUILDDIR)/coingecko$(EXE)
all: all-go
install: install-go
clean: clean-go
all-go: $(BUILDDIR)/coingecko$(EXE)
install-go:
	mkdir -p $(DESTDIR)$(PREFIX)/bin
	cp  $(BUILDDIR)/coingecko$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f $(BUILDDIR)/coingecko$(EXE)
##
$(BUILDDIR)/coingecko$(EXE): $(GO_DEPS)
	mkdir -p $(BUILDDIR)
	go build -o $@ $(GO_CONF) ./cmd/coingecko
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: README.md LICENSE
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp README.md LICENSE $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
## -- BLOCK:license --
