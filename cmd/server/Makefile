VERSION=latest
TAG=mboye/primes-server:$(VERSION)

cmd:
	GOOS=linux go build -o primes-server .

image: cmd
	docker build -t $(TAG) .

run: image
	docker run -it $(TAG)

push: image
	docker push $(TAG)

.PHONY: cmd

