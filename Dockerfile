FROM golang:alpine
RUN mkdir -p /go/src \
 && mkdir -p /go/bin \
 && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH
RUN mkdir -p $GOPATH/src/watcher
ADD . $GOPATH/src/watcher
WORKDIR $GOPATH/src/watcher
RUN go build -o watch .
CMD ["/go/src/watcher/watch", "-dir", "/tmp"]
