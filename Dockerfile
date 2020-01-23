FROM golang:1.13 as builder

WORKDIR /app
COPY . .
COPY . "/go/src/github.com/tpetrychyn/nes-emu"
RUN go build -o main .
RUN chmod 755 .
RUN ls -alh

FROM lucastetreault/ssl-scratch
WORKDIR /app
COPY --from=builder /app/nestest.nes .
COPY --from=builder /app/main .
CMD ["/app/main"]