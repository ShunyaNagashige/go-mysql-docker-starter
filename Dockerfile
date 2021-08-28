FROM golang:1.16

RUN go install github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest \
  && go install github.com/ramya-rao-a/go-outline@latest \
  && go install github.com/nsf/gocode@latest \
  && go install github.com/acroca/go-symbols@latest \
  && go install github.com/fatih/gomodifytags@latest \
  && go install github.com/go-delve/delve/cmd/dlv@latest \
  && go install golang.org/x/lint/golint@latest \
  && go install golang.org/x/tools/gopls@latest \
  && go install golang.org/x/tools/cmd/goimports@latest

WORKDIR /go/src/github.com/ShunyaNagashige/cohduc-backend

# RUN go mod init github.com/ShunyaNagashige/go-mysql-de
