FROM golang:1.14
RUN mkdir -p $GOPATH/src/github.com/ajangi/gAuthService
COPY . $GOPATH/src/github.com/ajangi/gAuthService
WORKDIR $GOPATH/src/github.com/ajangi/gAuthService
RUN go get ./app
RUN go build ./app/app.go
ENTRYPOINT [ "./app/app" ]
