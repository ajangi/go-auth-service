FROM golang:1.14
RUN mkdir -p $GOPATH/src/github.com/ajangi/gAuthService
COPY ./app $GOPATH/src/github.com/ajangi/gAuthService
WORKDIR $GOPATH/src/github.com/ajangi/gAuthService
RUN go get ./
RUN go build ./app.go
ENTRYPOINT [ "./app" ]
