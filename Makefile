include ./scripts/system/Makefile

dep:
	@GO111MODULE=on go mod vendor -v

test:
	@go test ./... -cover -coverprofile c.out

semver:
	@cd ./cmd/semver && go generate
	@CGO_ENABLED=0 \
		GO111MODULE=on \
		GOARCH=$(SYS_ARCH) \
		GOOS=$(SYS_OS) \
		go build -mod vendor -v -a \
			-ldflags "-extldflags '-static'" \
			-o ./bin/semver \
			./cmd/semver

semver_run:
	@go run ./cmd/semver ${ARGS}

image: semver
	@docker build \
		--build-arg BIN_NAME=semver \
		--file ./build/Dockerfile \
		--tag usvc/semver \
		.
