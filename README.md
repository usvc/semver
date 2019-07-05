# Semver

`WORK-IN-PROGRESS`

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
