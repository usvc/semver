# Semver

`WORK-IN-PROGRESS`

# Installation

```sh
go get -u gitlab.com/usvc/utils/semver
```

# Usage

## Bump a provided version

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

```sh
# bump the major version:
git tag "$(semver bump --git -M)"

# bump the minor version:
git tag "$(semver bump --git -m)"

# bump the patch version:
git tag "$(semver bump --git -p)"

# bump the label version:
git tag "$(semver bump --git -l)"
```

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
