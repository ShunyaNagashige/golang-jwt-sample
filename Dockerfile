# FROM golang:1.16.7-alpine
FROM golang:1.16

# RUN apk update && apk --no-cache add git gcc

RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest \
  && go install github.com/ramya-rao-a/go-outline@latest \
  && go install github.com/nsf/gocode@latest \
  && go install github.com/acroca/go-symbols@latest \
  && go install github.com/fatih/gomodifytags@latest \
  && go install github.com/go-delve/delve/cmd/dlv@latest \
  && go install golang.org/x/lint/golint@latest \
  && go install golang.org/x/tools/gopls@latest \
  && go install golang.org/x/tools/cmd/goimports@latest \
  && go install github.com/cosmtrek/air@latest

WORKDIR /go/src/github.com/ShunyaNagashige/golang-jwt-sample

# RUN go mod init github.com/ShunyaNagashige/go-mysql-de
