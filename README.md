# Semver

[![release github](https://badge.fury.io/gh/usvc%2Fsemver.svg)](https://github.com/usvc/semver/releases)
[![pipeline status](https://gitlab.com/usvc/utils/semver/badges/master/pipeline.svg)](https://gitlab.com/usvc/utils/semver/commits/master)
[![build btatus](https://travis-ci.org/usvc/semver.svg?branch=master)](https://travis-ci.org/usvc/semver)

An easy-peasy CLI tool to bump semver versions.

| | |
| --- | --- |
| Github | [https://github.com/usvc/semver](https://github.com/usvc/semver) |
| Gitlab | [https://gitlab.com/usvc/utils/semver](https://gitlab.com/usvc/utils/semver) |

- - -

- [Semver](#semver)
- [Installation](#installation)
  - [Via Go Get](#via-go-get)
  - [Binary Download](#binary-download)
  - [Via cURL](#via-curl)
  - [Via `/bin/sh`](#via-binsh)
- [Usage](#usage)
  - [CLI Help](#cli-help)
  - [Bump a provided version](#bump-a-provided-version)
  - [Bump a version using Git tags](#bump-a-version-using-git-tags)
  - [Bump Git tag version](#bump-git-tag-version)
  - [Usage via Dockerfile](#usage-via-dockerfile)
  - [Usage in a CI pipeline](#usage-in-a-ci-pipeline)
    - [GitLab CI (Automated)](#gitlab-ci-automated)
    - [GitLab CI (Manual)](#gitlab-ci-manual)
  - [Usage notes](#usage-notes)
- [Development](#development)
  - [Install dependencies](#install-dependencies)
  - [Run tests](#run-tests)
  - [Run the Go code](#run-the-go-code)
    - [Without arguments](#without-arguments)
    - [With arguments](#with-arguments)
  - [Create semver binary](#create-semver-binary)
  - [Development Runbook](#development-runbook)
    - [Getting Started](#getting-started)
    - [Continuous Integration (CI) Pipeline](#continuous-integration-ci-pipeline)
      - [On Github](#on-github)
        - [Releasing](#releasing)
      - [On Gitlab](#on-gitlab)
        - [Version Bumping](#version-bumping)
        - [DockerHub Publishing](#dockerhub-publishing)
- [License](#license)

- - -

# Installation

## Via Go Get

```sh
go get -v -u github.com/usvc/semver/cmd/semver
```

## Binary Download

Get the latest binary for your operating system from [https://github.com/usvc/semver/releases](https://github.com/usvc/semver/releases).

## Via cURL

```sh
# linux
curl -oL https://github.com/usvc/semver/releases/download/v0.3.17/semver_linux_amd64

# macos
curl -oL https://github.com/usvc/semver/releases/download/v0.3.17/semver_darwin_amd64

# windows
curl -oL https://github.com/usvc/semver/releases/download/v0.3.17/semver_windows_386.exe
```

## Via `/bin/sh`

```sh
curl -L https://raw.githubusercontent.com/usvc/semver/master/init/install.sh | sh
```

> This script downloads the latest binary to your working directory.

- - -

# Usage

## CLI Help

**Use Case** - As a developer, I'd like to find out how to use `semver` from the CLI

```sh
semver -h
```

## Bump a provided version

**Use Case** - As a developer, I'd like to do a bump of a semver version string and receive the next version as the output.

```sh
# major version bump:
semver bump major 1.2.3-label.1
# > 2.0.0

# minor version bump:
semver bump minor 1.2.3-label.1
# > 1.3.0

# patch version bump
semver bump patch 1.2.3-label.1
# OR simply...
semver bump 1.2.3-label.1
# > 1.2.4

# prerelease iteration bump
semver bump 1.2.3-prerelease.1 -l
# > 1.2.3-prerelease.2
```

## Bump a version using Git tags

**Use Case** - As a developer, I'd like to receive an output of the bumped version, using the latest semver version in a Git repository's tags as input.

```sh
# bump the major version (assuming git tag 'v1.2.3' exists):
semver bump major --git
# > v2.0.0

# bump the minor version (assuming git tag 'v1.2.3' exists):
semver bump minor --git
# > v1.3.0

# bump the patch version (assuming git tag 'v1.2.3' exists):
semver bump patch --git
# OR simply...
semver bump --git
# > v1.2.4

# bump the prerelease version (assuming git tag 'v1.2.3-prerelease.0' exists):
semver bump --git
# > v1.2.3-prerelease.1
```

## Bump Git tag version

**Use Case** - As a developer, I'd like to run a single command that gets the latest semver version in a Git repository's tags, bumps it, and adds it to the repository's tags.

```sh
# bump the major version:
semver bump --git --apply -M
# > added git tag 'vX.0.0'

# bump the minor version:
semver bump --git --apply -m
# > added git tag 'vX.Y.0'

# bump the patch version:
semver bump --git --apply -p
# > added git tag 'vX.Y.Z'

# bump the label version:
semver bump --git --apply -l
# > added git tag 'vX.Y.Z-label.A'
```

## Usage via Dockerfile

```sh
docker run -it -v $(pwd):/repo usvc/semver:latest ${SUBCOMMANDS_AND_FLAGS}
```

> Replace `${SUBCOMMANDS_AND_FLAGS}` with whatever you would run behind the main `semver` command.

## Usage in a CI pipeline

### GitLab CI (Automated)

To use the inbuilt scripts, you need to define the following variables in your CI pipeline:

| Key | Description |
| ---: | :--- |
| `BASE64_DEPLOY_KEY` | Base64-encoded deploy key |
| `GIT_EMAIL` | Email of the Git user to use to push |
| `GIT_NAME` | Name of the Git user to use to push |
| `REPO_HOSTNAME` | Hostname of the repository |
| `REPO_URL` | Repository clone URL (the SSH one) |

Then use the following job specification:

```yaml
bump:
  only: ["master"]
  stage: versioning
  image: usvc/semver:gitlab-latest
  before_script:
    - bump-before
  script:
    - bump-script
  after_script:
    - bump-after
```

### GitLab CI (Manual)

An example version bump job together with pushing back to the repository can be as such:

```yaml
bump:
  only: ["master"]
  stage: versioning
  before_script:
    - mkdir -p ~/.ssh
    - 'printf -- "${DEPLOY_KEY}" | base64 -d > ~/.ssh/id_rsa'
    - chmod 600 -R ~/.ssh/id_rsa
    - ssh-keyscan -t rsa gitlab.com >> ~/.ssh/known_hosts
  script:
    - git remote set-url origin "${DEPLOY_URL}"
    - git checkout master
    - docker run -v $(pwd):/repo usvc/semver:latest + --git --apply
    - git push origin master --verbose --tags
  after_script:
    - rm -rf ~/.ssh/*
```

Set the `DEPLOY_KEY` environment variable from your CI/CD settings to a base64 encoded version of your private key. To generate a private/public key pair, use `ssh-keygen -t rsa -b 4096`. To encode it into base64 without line breaks, `cat` it and pipe it to `base64 -w 0` (eg. `cat ./path/to/id_rsa | base64 -w 0 > ./path/to/id_rsa.b64`).

Set the `DEPLOY_URL` environment variable from your CI/CD settings to the SSH clone URL of the repository you'd like to push to.

## Usage notes

- If the major (`-M`), minor (`-m`), patch (`-p`), or label (`-l`) flag is not specified, the patch will be bumped by default.
- If more than one instance of a version indicator flag is specified, the lowest priority will be executed. For example, if both `-M` and `-m` is specified, `-m` will be applied.

- - -

# Development

Canonical repository URL: https://gitlab.com/usvc/utils/semver
Public URL: https://github.com/usvc/semver

## Install dependencies

```sh
make dep
```

## Run tests

```sh
make test
```

## Run the Go code

### Without arguments

```sh
make semver_run
```

### With arguments

```sh
make semver_run ARGS="bump 1.2.3 -m"
```

## Create semver binary

```sh
make semver
```

- - -

## Development Runbook

### Getting Started

1. Clone this repository
2. Run `make deps` to pull in external dependencies
3. Write some awesome stuff
4. Run `make test` to ensure unit tests are passing
5. Push

### Continuous Integration (CI) Pipeline

#### On Github

Github is used to deploy binaries/libraries because of it's ease of access by other developers.

##### Releasing

Releasing of the binaries can be done via Travis CI.

1. On Github, navigate to the [tokens settings page](https://github.com/settings/tokens) (by clicking on your profile picture, selecting **Settings**, selecting **Developer settings** on the left navigation menu, then **Personal Access Tokens** again on the left navigation menu)
2. Click on **Generate new token**, give the token an appropriate name and check the checkbox on **`public_repo`** within the **repo** header
3. Copy the generated token
4. Navigate to [travis-ci.org](https://travis-ci.org) and access the cooresponding repository there. Click on the **More options** button on the top right of the repository page and select **Settings**
5. Scroll down to the section on **Environment Variables** and enter in a new **NAME** with `RELEASE_TOKEN` and the **VALUE** field cooresponding to the generated personal access token, and hit **Add**

#### On Gitlab

Gitlab is used to run tests and ensure that builds run correctly.

##### Version Bumping

1. Run `make .ssh`
2. Copy the contents of the file generated at `./.ssh/id_rsa.base64` into an environment variable named **`DEPLOY_KEY`** in **Settings > CI/CD > Variables**
3. Navigate to the **Deploy Keys** section of the **Settings > Repository > Deploy Keys** and paste in the contents of the file generated at `./.ssh/id_rsa.pub` with the **Write access allowed** checkbox enabled

- **`DEPLOY_KEY`**: generate this by running `make .ssh` and copying the contents of the file generated at `./.ssh/id_rsa.base64`

##### DockerHub Publishing

1. Login to [https://hub.docker.com](https://hub.docker.com), or if you're using your own private one, log into yours
2. Navigate to [your security settings at the `/settings/security` endpoint](https://hub.docker.com/settings/security)
3. Click on **Create Access Token**, type in a name for the new token, and click on **Create**
4. Copy the generated token that will be displayed on the screen
5. Enter the following varialbes into the CI/CD Variables page at **Settings > CI/CD > Variables** in your Gitlab repository:

- **`DOCKER_REGISTRY_URL`**: The hostname of the Docker registry (defaults to `docker.io` if not specified)
- **`DOCKER_REGISTRY_USERNAME`**: The username you used to login to the Docker registry
- **`DOCKER_REGISTRY_PASSWORD`**: The generated access token

# License

This project is licensed under the MIT license. [See the full text](./LICENSE).
