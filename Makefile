build:
	@go build -o /bin/server .

run: build
	./bin/server

test: 
	@go test ./... --race -v


docker_build:
	docker build -t tcp-server .

docker_run: docker_build
	docker run --name tcp-server -p 6379:6379 tcp-server

docker_stop:
	docker stop tcp-server && docker remove tcp-server
