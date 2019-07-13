include ./scripts/system/Makefile

# ------------------------------------------------------------------------
# development recipes
# ------------------------------------------------------------------------

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
run_get:
	# runs the get application in development
	@go run ./cmd/get ${ARGS}
run_semver:
	# runs the semver application in development
	@go run ./cmd/semver ${ARGS}

# ------------------------------------------------------------------------
# ci pipeline recipes
# ------------------------------------------------------------------------

binary:
	# builds main binary
	@$(MAKE) _binary ARCH=$(SYS_ARCH) OS=$(SYS_OS) BIN_EXT=$(BIN_EXT) BIN_NAME=semver
	cp ./bin/semver-$(SYS_OS)-$(SYS_ARCH)$(BIN_EXT) ./bin/semver$(BIN_EXT)
binaries:
	# builds all binaries
	@$(MAKE) _binaries BIN_NAME=semver \
		& $(MAKE) _binaries BIN_NAME=bump \
		& $(MAKE) _binaries BIN_NAME=get \
		& wait
_binaries:
	# generic method to build binaries for all oses/architectures
	@$(MAKE) _binary ARCH=amd64 OS=linux BIN_NAME=${BIN_NAME}
	@$(MAKE) _binary ARCH=amd64 OS=darwin BIN_NAME=${BIN_NAME}
	@$(MAKE) _binary ARCH=386 OS=windows BIN_NAME=${BIN_NAME} BIN_EXT=$(BIN_EXT)
_binary:
	# generic method to build a binary
	@cd ./cmd/semver && go generate
	@CGO_ENABLED=0 \
		GO111MODULE=on \
		GOARCH=${ARCH} \
		GOOS=${OS} \
		go build -mod vendor -v -a \
			-ldflags "-extldflags '-static'" \
			-o ./bin/${BIN_NAME}-${OS}-${ARCH}${BIN_EXT} \
			./cmd/${BIN_NAME}

image:
	# builds the main docker image
	@$(MAKE) _image IMAGE_NAME=semver TARGET=semver
images:
	# builds all docker images for all binaries
	# do it sequentially so that we have caches
	@$(MAKE) _image IMAGE_NAME=semver-get TARGET=get
	@$(MAKE) _image IMAGE_NAME=semver-bump TARGET=bump
	@$(MAKE) _image IMAGE_NAME=semver TARGET=semver
	@$(MAKE) _image IMAGE_NAME=semver TARGET=gitlab PRE_TAG="gitlab-"
_image:
	# driver function for building images
	@docker build \
		--file ./build/Dockerfile \
		--tag usvc/${IMAGE_NAME}:${PRE_TAG}latest \
		--target ${TARGET} \
		.

publish_image: image
	# publishes main docker image
	@$(MAKE) _publish_image IMAGE_NAME=semver
publish_images: images
	# publishes the docker image
	@$(MAKE) _publish_image IMAGE_NAME=semver \
		& $(MAKE) _publish_image IMAGE_NAME=semver-bump \
		& $(MAKE) _publish_image IMAGE_NAME=semver-get \
		& $(MAKE) _publish_image IMAGE_NAME=semver PRE_TAG="gitlab-" \
		& wait
_publish_image:
	# driver function for publishing images
	@docker tag usvc/${IMAGE_NAME}:${PRE_TAG}latest usvc/${IMAGE_NAME}:${PRE_TAG}$$(docker run usvc/semver:latest -v | cut -f 3 -d ' ' | sed -e 's/v//g')
	@docker push usvc/${IMAGE_NAME}:${PRE_TAG}latest
	@docker push usvc/${IMAGE_NAME}:${PRE_TAG}$$(docker run usvc/semver:latest -v | cut -f 3 -d ' ' | sed -e 's/v//g')

publish_github:
	# publish repository to github
	@git remote add downstream git@github.com:usvc/semver.git
	@git fetch
	@git config --global user.email "usvc.semver@usvc.dev"
  @git config --global user.name "usvc_publisher"
	@git commit --allow-empty -m "from gitlab to github"
	@git push -u downstream master --tags --force

# ------------------------------------------------------------------------
# misc recipes
# ------------------------------------------------------------------------

deploy_keys:
	@mkdir -p ./.ssh
	# creating keypair...
	@ssh-keygen -t rsa -b 4096 -f ./.ssh/id_rsa -N ""
	# creating base64 encoded version for usage in ci variables...
	@cat ./.ssh/id_rsa | base64 -w 0 > ./.ssh/id_rsa.b64
