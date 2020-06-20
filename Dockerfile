FROM golang:1.14
EXPOSE 8080
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN go get golang.org/x/tools/cmd/godoc
CMD ["/go/bin/godoc","-http=:8080"]