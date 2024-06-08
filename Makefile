BIN=river-buddy

export CGO_ENABED=0
GOSRC=$(wildcard *.go) $(wildcard **/*.go)
GOMOD=$(wildcard go.mod) $(wildcard go.sum)

GOBUILD_TAGS=-tags netgo

$(BIN): $(GOSRC) $(GOMOD) $(BIN_STATIC_RESOURCES)
	go build $(GOBUILD_TAGS) -o $(BIN)
