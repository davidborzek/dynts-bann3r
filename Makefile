deps:
	go mod tidy
	go mod download

run: deps
	go run src/main.go

build: deps
	go build -o out/dynts-bann3r src/main.go

test_unit: deps
	go test ./...

docker_test:
	docker build -t dynts-bann3r-unit-test --target test .

docker_build:
	docker build -t dynts-bann3r .

docker_run: docker_build
	docker run --rm -it -p 9000:9000 -v $(PWD)/config.json:/config.json -v $(PWD)/template.png:/template.png -v /etc/localtime:/etc/localtime:ro  dynts-bann3r