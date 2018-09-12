FROM golang:1.11.0-stretch as build

WORKDIR $GOPATH/src/github.com/quinlanmorake/verisart-go
COPY . .

RUN go version && go get -u -v golang.org/x/vgo
RUN CC=gcc vgo install ./...

FROM gcr.io/distroless/base
COPY --from=build /go/bin/verisart-go /
COPY app.config /
  
CMD ["/verisart-go"]
