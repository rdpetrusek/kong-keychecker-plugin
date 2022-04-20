FROM golang:alpine as builder
# delete with 'docker image prune --filter label=stage=temporary'
LABEL stage=temporary
RUN apk add --no-cache git gcc libc-dev curl
RUN mkdir /go-plugins 
WORKDIR /go-plugins
RUN cd /go-plugins

COPY key-checker.go .
COPY go.mod .
COPY go.sum .
COPY copy-plugin.sh .
RUN go mod download
RUN go mod tidy
RUN go build key-checker.go

FROM busybox
COPY --from=builder /go-plugins/key-checker .
COPY --from=builder /go-plugins/copy-plugin.sh .
CMD ["sh", "copy-plugin.sh"]