FROM golang:1.17.0-alpine3.14 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./dynts-bann3r src/main.go

WORKDIR /dist

RUN cp /build/dynts-bann3r .

FROM scratch

COPY --from=builder /dist/dynts-bann3r /
COPY --from=builder /build/fonts /fonts

CMD ["./dynts-bann3r"]