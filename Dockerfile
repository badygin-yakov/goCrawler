FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/yakovbadygin/parses/

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o /go/bin/crawler

# Final docker image
FROM alpine:3.7

WORKDIR /

COPY --from=builder /go/bin/crawler .

# List of CMD. The last one will be executed.
ENTRYPOINT ["./crawler"]