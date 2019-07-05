include ./scripts/system/Makefile

dep:
	@GO111MODULE=on go mod vendor -v

semver:
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