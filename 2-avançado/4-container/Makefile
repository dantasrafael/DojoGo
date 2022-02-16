build1:
	docker build -t demo:1 -f Dockerfile.1 .

build2:
	docker build -t demo:2 -f Dockerfile.2 .

build: build1 build2

start:
	docker run --rm -p 8080:8080 demo:2
