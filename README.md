# retlen

Checks the number of return values from functions.

## Installation

```shell
go install github.com/Namchee/retlen@latest
```

## Usage

```
retlen [-flag] [package]
```

## Flags

All options are provided through CLI flags.

Name | Type | Default | Description
---- | ---- | ------- | -----------
`maxReturn` | `int` | `2` | Maximum number of variables allowed to be returned by a function
`includeTest` | `bool` | `false` | Include test functions (starts with `Test`)