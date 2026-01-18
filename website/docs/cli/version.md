# freight version

Print the version information for the Freight CLI.

## Description

The `version` command displays the current version of the Freight CLI tool. This is useful for troubleshooting and ensuring you are running the expected version.

## Flags

- `-v, --verbose`: Show additional build details (commit hash and build date).

## Examples

### Normal mode
```bash
freight version
```

**Output:**
```text
Freight version: 0.1.0
```

### Verbose mode
```bash
freight version --verbose
```

**Output:**
```text
Freight version: 0.1.0
Commit: abc123
Date: 2026-01-18T12:00:00Z
```

:::info Note about development builds
When running Freight from source or a non-release build, the version will show as `dev`, the commit will be `none`, and the date will be `unknown`.
:::
