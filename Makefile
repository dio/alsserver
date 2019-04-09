.PHONY: build
build/log:
	CGO_ENABLED=0 GOOS=linux go build \
		-a --ldflags '-extldflags "-static"' -tags netgo -installsuffix netgo \
		-o build/log main.go

.PHONY: docker-build
docker-build: build
	docker build . -t gcr.io/tetratelabs/accesslogs-debugger:0.0.1

.PHONY: docker-push
docker-push:
	docker push gcr.io/tetratelabs/accesslogs-debugger:0.0.1

.PHONY: clean
clean:
	rm -fr build
