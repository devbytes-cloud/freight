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
- `-a, --allow`: Specific Git hooks to install (default: all). Valid options: `pre-commit`, `prepare-commit-msg`, `commit-msg`, `post-commit`, `post-checkout`.
    
    When this flag is used, Freight will **only** rewire the hooks you explicitly specify. Any existing hooks in your `.git/hooks` directory that are NOT in the allow list will remain untouched. This is useful if you want to use Freight alongside other hook managers or if you only want to manage a subset of hooks with Freight.
- `-q, --quiet`: Quiet mode, will only display errors. Suppresses informational and success output during initialization.

## Examples

Basic initialization:
```bash
freight init
```

Initialize with specific hooks (comma-separated):
```bash
freight init --allow pre-commit,commit-msg
```

Initialize with specific hooks (multiple flags):
```bash
freight init -a pre-commit -a post-checkout
```

Initialize in quiet mode (only errors are printed):
```bash
freight init --quiet
```

**Output:**
```text
✔ Extracting conductor binary...
✔ Generating railcar.json...
✔ Rewiring git hooks...
Freight initialized successfully!
```
