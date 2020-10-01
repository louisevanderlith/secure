FROM golang:1.13 as build_base

WORKDIR /box

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base as builder

COPY main.go .
COPY handles ./handles
COPY core ./core

RUN CGO_ENABLED="0" go build

FROM alpine:3.12.0

COPY --from=builder /box/secure .

EXPOSE 8086

ENTRYPOINT [ "./secure" ]
