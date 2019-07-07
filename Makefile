include ./scripts/system/Makefile

binaries: dep
	# builds a binaries
	@$(MAKE) _binary BIN_NAME=semver
	@$(MAKE) _binary BIN_NAME=bump
	@$(MAKE) _binary BIN_NAME=get
_binary:
	@cd ./cmd/${BIN_NAME} && go generate
	@CGO_ENABLED=0 \
		GO111MODULE=on \
		GOARCH=$(SYS_ARCH) \
		GOOS=$(SYS_OS) \
		go build -mod vendor -v -a \
			-ldflags "-extldflags '-static'" \
			-o ./bin/${BIN_NAME} \
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
	@git push origin master
	@git remote set-url origin $$(cat ./.publish_github)
	@rm -rf ./.publish_github

publish_images: images
	# publishes the docker image
	@docker push usvc/semver:latest
	@docker tag usvc/semver:latest usvc/semver:$$(docker run -it -v $$(pwd):/repo usvc/semver-get:latest)
	@docker push usvc/semver-bump:latest
	@docker tag usvc/semver-bump:latest usvc/semver-bump:$$(docker run -it -v $$(pwd):/repo usvc/semver-get:latest)
	@docker push usvc/semver-bump:$$(docker run -it -v $$(pwd):/repo usvc/semver-get:latest)
	@docker push usvc/semver-get:latest
	@docker tag usvc/semver-get:latest usvc/semver-get:$$(docker run -it -v $$(pwd):/repo usvc/semver-get:latest)
	@docker push usvc/semver-get:$$(docker run -it -v $$(pwd):/repo usvc/semver-get:latest)
