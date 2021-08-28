FROM golang:1.17.0-alpine3.14 AS base

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY src src

FROM base as test

CMD [ "go", "test", "./...", "-v" ]

FROM base as build

RUN [ "go", "test", "./..." ]
RUN go build -o ./dynts-bann3r src/main.go

WORKDIR /dist

RUN cp /build/dynts-bann3r .

FROM scratch as prod

COPY --from=build /dist/dynts-bann3r /
COPY fonts /fonts

CMD ["./dynts-bann3r"]