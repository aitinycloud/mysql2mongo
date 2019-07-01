all: build

build: build-elasticsearch

build-elasticsearch:
	go build -o bin/go-mysql-mongodb ./cmd/go-mysql-mongodb

test:
	go test -timeout 1m --race ./...

clean:
	go clean -i ./...
	@rm -rf bin


update_vendor:
	which glide >/dev/null || curl https://glide.sh/get | sh
	which glide-vc || go get -v -u github.com/sgotti/glide-vc
	glide --verbose update --strip-vendor --skip-test
	@echo "removing test files"
	glide vc --only-code --no-tests
