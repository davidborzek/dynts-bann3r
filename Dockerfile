FROM golang:1.17.0-alpine3.14

WORKDIR /home/go/app

COPY . .

RUN go mod tidy

RUN go build -o ./dynts-bann3r src/main.go

RUN chmod +x ./dynts-bann3r

CMD ["./dynts-bann3r"]