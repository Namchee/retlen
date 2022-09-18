# retlen

Checks the number of return values from functions.

## Usage

```
retlen [-flag] [package]
```

## Flags

Name | Type | Default | Description
---- | ---- | ------- | -----------
`maxReturn` | `uint` | `2` | Maximum number of variables allowed to be returned by a function
`includeTest` | `bool` | `false` | Include test functions (starts with `Test`)