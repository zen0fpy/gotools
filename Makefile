WORKDIR=`pwd`

default: build

vet:
	go vet ./...

tools:
	go get github.com/golangci/golangci-lint/cmd/golangci-lint
	go get github.com/golang/lint/golint
	go get github.com/axw/gocov/gocov
	go get github.com/matm/gocov-html
	go get github.com/fzipp/gocyclo  # 检查函数的复杂度
	go get github.com/mvdan/interfacer/cmd/interfacer # 接口类型
	go get github.com/tsenart/deadcode # 告诉你哪些代码片段根本没用
	go get golang.org/x/tools/cmd/gotype #语法分析
	go get github.com/client9/misspell # 拼写
	go get github.com/jgautheron/goconst/cmd/goconst # 重复字符串，提取为常量
	go get honnef.co/go/tools/cmd/staticcheck # 静态检查

golangci-lint:
	golangci-lint run -D errcheck --build-tags 'quic kcp'

lint:
	golint ./...

doc:
	godoc -http=:6060

deps:
	go list -f '{{ join .Deps  "\n"}}' ./... |grep "/" | grep -v "github.com/smallnest/rpcx"| grep "\." | sort |uniq

fmt:
	go fmt ./...

build:
	go build ./...

build-all:
	go build -tags "kcp quic" ./...

test:
	go test -race -tags "hello" --cover -v -coverprofile cover.out ./...

cover:
	gocov test -tags "kcp quic" ./... | gocov-html > cover.html
	open cover.html

check-libs:
	GIT_TERMINAL_PROMPT=1 GO111MODULE=on go list -m -u all | column -t

update-libs:
	GIT_TERMINAL_PROMPT=1 GO111MODULE=on go get -u -v ./...

mod-tidy:
	GIT_TERMINAL_PROMPT=1 GO111MODULE=on go mod tidy
