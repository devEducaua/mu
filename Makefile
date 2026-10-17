
TARGET = mu
TARGETDIR = build
MANPAGE = $(TARGET).1
SRCS = $(shell fd -e .go)

PREFIX ?= /usr/local
BINDIR = $(PREFIX)/bin
MANDIR = $(PREFIX)/share/man/man1

.PHONY: install clean

all: $(TARGETDIR)/$(TARGET)

$(TARGETDIR)/$(TARGET): $(SRCS)
	mkdir -p $(TARGETDIR)
	go build -o $@ cmd/main.go

install:
	cp $(TARGETDIR)/$(TARGET) $(BINDIR)/
	cp $(MANPAGE) $(MANDIR)/
	gzip -f $(MANDIR)/$(MANPAGE)

uninstall:
	rm -f $(BINDIR)/$(TARGET) 
	rm -f $(MANDIR)/$(MANPAGE) 

clean:
	rm -r $(TARGETDIR)

