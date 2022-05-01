FROM golang:1.14.15 AS builder
ENV GO111MODULE on
WORKDIR /workspace
ADD . /workspace
RUN apt-get update
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build

FROM alpine
RUN apk upgrade && apk add --no-cache ca-certificates
COPY --from=builder /workspace/app ./
COPY --from=builder /workspace/resources ./resources
ENTRYPOINT ["./app"]