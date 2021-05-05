NAME=sysproxy
BINDIR=bin
GOBUILD=CGO_ENABLED=0 go build -trimpath -ldflags '-w -s'

PLATFORM_LIST = \
	darwin-amd64 \
	darwin-arm64 \

all: darwin-amd64 darwin-arm64

darwin-amd64:
	GOARCH=amd64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$@/$(NAME)

darwin-arm64:
	GOARCH=arm64 GOOS=darwin $(GOBUILD) -o $(BINDIR)/$@/$(NAME)

clean:
	rm -rf $(BINDIR)