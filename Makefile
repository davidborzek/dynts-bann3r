build:
	docker build -t dynts-bann3r .

run: build
	docker run --rm -it -p 9000:9000 -v $(HOME)/git/dynts-bann3r/config.json:/config.json -v $(HOME)/git/dynts-bann3r/template.png:/template.png -v /etc/localtime:/etc/localtime:ro  dynts-bann3r