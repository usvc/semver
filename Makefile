include ./scripts/system/Makefile

binary: dep
	# builds main binary
	@$(MAKE) _binary GOARCH=$(SYS_ARCH) GOOS=$(SYS_OS) BIN_NAME=semver
	cp ./bin/semver-$(SYS_OS)-$(SYS_ARCH) ./bin/semver
binaries.all:
	@$(MAKE) _binary.all.supported BIN_NAME=semver
	@$(MAKE) _binary.all.supported BIN_NAME=bump
	@$(MAKE) _binary.all.supported BIN_NAME=get
_binary.all.supported: dep
	# generic method to build binaries for all oses/architectures
	@$(MAKE) _binary GOARCH=amd64 GOOS=linux BIN_NAME=${BIN_NAME}
	@$(MAKE) _binary GOARCH=amd64 GOOS=darwin BIN_NAME=${BIN_NAME}
	@$(MAKE) _binary GOARCH=386 GOOS=windows BIN_NAME=${BIN_NAME}
_binary:
	# generic method to build a binary
	@cd ./cmd/${BIN_NAME} && go generate
	@CGO_ENABLED=0 \
		GO111MODULE=on \
		GOARCH=${GOARCH} \
		GOOS=${GOOS} \
		go build -mod vendor -v -a \
			-ldflags "-extldflags '-static'" \
			-o ./bin/${BIN_NAME}-${GOOS}-${GOARCH}${BIN_EXT} \
			./cmd/${BIN_NAME}

dep:
	# install dependencies
	@GO111MODULE=on go mod vendor -v
	@GO111MODULE=on go mod download

test:
	# running tests with output coverage at ./c.out
	@go test ./... -cover -coverprofile c.out

run_bump:
	# runs the bump application in development
	@go run ./cmd/bump ${ARGS}

run_semver:
	# runs the semver application in development
	@go run ./cmd/semver ${ARGS}

images:
	# builds the docker image for `semver`
	@docker build \
		--file ./build/Dockerfile \
		--tag usvc/semver:latest \
		--target semver \
		.
	# builds the docker image for `semver-bump`
	@docker build \
		--file ./build/Dockerfile \
		--tag usvc/semver-bump:latest \
		--target bump \
		.
	# builds the docker image for `semver-get`
	@docker build \
		--file ./build/Dockerfile \
		--tag usvc/semver-get:latest \
		--target get \
		.

publish_github:
	# publish repository to github
	@git remote get-url origin > ./.publish_github
	@git remote set-url origin git@github.com:usvc/semver.git
	@git push origin master --tags
	@git remote set-url origin $$(cat ./.publish_github)
	@rm -rf ./.publish_github

publish_images: images
	# publishes the docker image
	@docker push usvc/semver:latest
	# @docker tag usvc/semver:latest usvc/semver:$$(docker run -v $$(pwd):/repo usvc/semver-get:latest)
	# @docker push usvc/semver:$$(docker run -v $$(pwd):/repo usvc/semver-get:latest)
	@docker push usvc/semver-bump:latest
	# @docker tag usvc/semver-bump:latest usvc/semver-bump:$$(docker run -v $$(pwd):/repo usvc/semver-get:latest)
	# @docker push usvc/semver-bump:$$(docker run -v $$(pwd):/repo usvc/semver-get:latest)
	@docker push usvc/semver-get:latest
	# @docker tag usvc/semver-get:latest usvc/semver-get:$$(docker run -v $$(pwd):/repo usvc/semver-get:latest)
	# @docker push usvc/semver-get:$$(docker run -v $$(pwd):/repo usvc/semver-get:latest)
