include ./scripts/system/Makefile

dep:
	# install dependencies
	@GO111MODULE=on go mod vendor -v

test:
	# runs tests
	@go test ./... -cover -coverprofile c.out

build:
	# builds a binary at ./bin/semver
	@cd ./cmd/semver && go generate
	@CGO_ENABLED=0 \
		GO111MODULE=on \
		GOARCH=$(SYS_ARCH) \
		GOOS=$(SYS_OS) \
		go build -mod vendor -v -a \
			-ldflags "-extldflags '-static'" \
			-o ./bin/semver \
			./cmd/semver

run:
	# runs the semver application in development
	@go run ./cmd/semver ${ARGS}

image: semver
	# builds the docker image
	@docker build \
		--build-arg BIN_NAME=semver \
		--file ./build/Dockerfile \
		--tag usvc/semver:latest \
		.

publish_image: image
	# publishes the docker image
	@docker push usvc/semver:latest
