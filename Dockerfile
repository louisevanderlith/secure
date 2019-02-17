FROM alpine:latest as builder

COPY secure .
COPY conf conf

ENTRYPOINT [ "./secure" ]
