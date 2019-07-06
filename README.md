# Semver

`WORK-IN-PROGRESS`

# Installation

```sh
go get -u gitlab.com/usvc/utils/semver
```

# Usage

## Bump a patch version

```sh
semver bump 1.2.3-label.1
# > 1.2.4
```

## Bump a minor version

```sh
semver bump 1.2.3-label.1 -m
# > 1.3.0
```

## Bump a major version

```sh
semver bump 1.2.3-label.1 -M
# > 2.0.0
```

## Bump the label iteration

```sh
semver bump 1.2.3-label.1 -l
# > 1.2.3-label.2
```

## Bump a Git tag version

```sh
semver bump --git
```

> The `-M`, `-m`, `-p` and `-l` flags can be used in conjunction with `--git`.

## End-to-end with Git tags

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

# Development

## Install dependencies

```sh
make dep
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
