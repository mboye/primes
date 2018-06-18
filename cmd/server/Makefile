VERSION=latest
TAG=mboye/primes-backend:$(VERSION)

cmd:
	GOOS=linux go build -o primes-backend .

image: cmd
	docker build -t $(TAG) .

run: image
	docker run -it $(TAG)

push: image
	docker push $(TAG)

clean:
	rm -f primes-backend

.PHONY: cmd clean

