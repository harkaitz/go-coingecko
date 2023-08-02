VERSION=1.0.2b
PROJECT=go-coingecko
PREFIX=/usr/local

## -- BLOCK:go --
all: all-go
install: install-go
clean: clean-go
deps: deps-go

build/coingecko$(EXE): deps
	go build -o $@ $(GO_CONF) ./cmd/coingecko

all-go:  build/coingecko$(EXE)
deps-go:
	mkdir -p build
install-go:
	install -d $(DESTDIR)$(PREFIX)/bin
	cp  build/coingecko$(EXE) $(DESTDIR)$(PREFIX)/bin
clean-go:
	rm -f  build/coingecko$(EXE)
## -- BLOCK:go --
## -- BLOCK:license --
install: install-license
install-license: 
	mkdir -p $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
	cp LICENSE README.md $(DESTDIR)$(PREFIX)/share/doc/$(PROJECT)
update: update-license
update-license:
	ssnip README.md
## -- BLOCK:license --
