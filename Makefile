GOCMD=go
GOBUILD=$(GOCMD) build -race -ldflags="-w -s"
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINDIR=$(CURDIR)/bin
CMDDIR=$(CURDIR)/cmd

TARGETS = day1 day2 day3

.PHONY: $(TARGETS)

all: $(TARGETS)

$(TARGETS):
	$(GOBUILD) -o $(BINDIR)/$@ $(CMDDIR)/$@/main.go

test: 
	$(GOTEST) ./... | grep -v 'no test file'

clean:
	$(GOCLEAN) -testcache
	rm -f $(BINDIR)/*
