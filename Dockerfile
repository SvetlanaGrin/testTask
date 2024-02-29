FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o testTask cmd/main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/testTask /build/testTask

CMD ["/testTask"]