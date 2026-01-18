# freight init

Bootstrap Freight in the current repository.

## Description

The `init` command sets up Freight in your local repository. It ensures all necessary components are in place for managing your Git hooks.

**What it does:**
- Extracts the **Conductor** binary to your repository root.
- Generates a starter **Railcar** manifest (`railcar.json`).
- Rewires your `.git/hooks` to point to the Conductor.

## Flags

- `-c, --config-force`: Overwrite an existing `railcar.json` file if it already exists.

## Examples

```bash
freight init
```

**Output:**
```text
✔ Extracting conductor binary...
✔ Generating railcar.json...
✔ Rewiring git hooks...
Freight initialized successfully!
```
