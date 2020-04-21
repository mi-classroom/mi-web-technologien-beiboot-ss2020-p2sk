FROM golang:latest

WORKDIR /backend

COPY ./ /backend

RUN go env -w GOBIN=$GOPATH/bin
RUN go get -v
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build server.go bild.go" --command=./server