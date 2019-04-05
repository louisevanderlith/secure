FROM golang:1.11 as builder

WORKDIR /box
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY main.go .
COPY controllers ./controllers
COPY core ./core
COPY logic ./logic
COPY routers ./routers

RUN CGO_ENABLED="0" go build

FROM scratch

COPY --from=builder /box/secure .
COPY conf conf

EXPOSE 8086

ENTRYPOINT [ "./secure" ]
