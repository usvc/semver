# Semver

[![pipeline status](https://gitlab.com/usvc/utils/semver/badges/master/pipeline.svg)](https://gitlab.com/usvc/utils/semver/commits/master)

An easy-peasy CLI tool to bump semver versions.

# Installation

```sh
go get -v -u gitlab.com/usvc/utils/semver/cmd/semver;
```

# Usage

## Bump a provided version

**Use Case** - As a developer, I'd like to do a bump of a semver version string and receive the next version as the output.

```sh
# major version bump:
semver bump 1.2.3-label.1 -M
# > 1.2.4

# minor version bump:
semver bump 1.2.3-label.1 -m
# > 1.3.0

# patch version bump
semver bump 1.2.3-label.1 -p
# > 2.0.0

# label iteration bump
semver bump 1.2.3-label.1 -l
# > 1.2.3-label.2
```

## Bump a version using Git tags

**Use Case** - As a developer, I'd like to receive an output of the bumped version, using the latest semver version in a Git repository's tags as input.

```sh
# bump the major version (assuming git tag 'v1.2.3' exists):
semver bump --git -M
# > v2.0.0

# bump the minor version (assuming git tag 'v1.2.3' exists):
semver bump --git -m
# > v1.3.0

# bump the patch version (assuming git tag 'v1.2.3' exists):
semver bump --git -p
# > v1.2.4

# bump the label version (assuming git tag 'v1.2.3-label.0' exists):
semver bump --git -l
# > v1.2.3-label.1
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

## Usage notes

- If the major (`-M`), minor (`-m`), patch (`-p`), or label (`-l`) flag is not specified, the patch will be bumped by default.
- If more than one instance of a version indicator flag is specified, the lowest priority will be executed. For example, if both `-M` and `-m` is specified, `-m` will be applied.

# Development

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

## Configuring the CI

### Gitlab

Go to the settings page under CI/CD > Variables and input the following CI variables:

| Key | Description |
| ---: | :--- |
| `DOCKER_REGISTRY_URL` | URL to the Docker registry to push to |
| `DOCKER_REGISTRY_USERNAME` | Username of the Docker registry to push to |
| `DOCKER_REGISTRY_PASSWORD` | Password of the Docker registry user |

# License

This project is licensed under the MIT license. [See the full text](./LICENSE).
