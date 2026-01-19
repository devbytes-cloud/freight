# freight status

Report the current state of Freight in the repository.

## Description

The `status` command provides a comprehensive overview of the current Freight setup in your repository. It checks for the existence of required files, displays version information, and analyzes the status of Git hooks.

**What it checks:**
- **Valid Git repository:** Verifies that the current directory is within a Git repository.
- **Railcar Manifest (`railcar.json`):** Checks if the manifest file exists.
- **Fingerprint (`.git/hooks/.fingerprint.yaml`):** Checks if the Freight fingerprint file exists in `.git/hooks/`.
- **Conductor Binary:** Checks if the `conductor` binary is present in the repository root.
- **Version Information:** Displays both the Freight version used during initialization and the current binary version.
- **Git Hooks Status:** Provides a detailed table showing:
    - Which hooks are present.
    - Which hooks are managed by Freight.
    - Detection of "drift" (if the hook content differs from what Freight expects).

## Status Fields

### Core Status

- **Valid Git repository**: Confirms if the current directory is a valid Git repository where Freight can operate.
- **Railcar Manifest (`railcar.json`)**: The configuration file that defines your hook steps (commands to run for each hook).
- **Fingerprint (`.git/hooks/.fingerprint.yaml`)**: A file stored in `.git/hooks/` that tracks the state of hooks managed by Freight and the version used for initialization.
- **Conductor Binary**: The small executable that Freight installs into your repository to orchestrate the execution of hooks defined in `railcar.json`.
- **Freight Version**: The version recorded in the fingerprint when the repository was last initialized or migrated.
- **Current Freight Binary Version**: The version of the `freight` CLI tool you are currently running.

### Git Hooks Table

The Git Hooks Status table provides a granular look at each supported hook:

- **Hook**: The name of the Git hook (e.g., `pre-commit`).
- **Exists**: Indicates whether a file for this hook exists in the `.git/hooks/` directory.
- **Managed by Freight**: Shows if Freight is configured to manage this hook. This is determined by the "allow list" created during `freight init`. If a hook is not managed, Freight will not attempt to rewire it or check for drift.
- **Drift**: Measures whether the current hook file in `.git/hooks/` matches the template Freight expects.
    - `none`: The hook is correctly managed and matches the expected state.
    - `drifted`: The hook file exists and is managed by Freight, but its content has been modified or overwritten by something else.
    - `missing`: The hook is supposed to be managed by Freight, but the file is missing from `.git/hooks/`.
    - `error`: An error occurred while trying to read the hook file or calculate the expected state.
    - `N/A`: Drift detection is not applicable (usually because the hook is not managed by Freight).

## Usage

```bash
freight status
```

## Examples

Check the status of Freight in your repository:
```bash
freight status
```

**Output:**
```text
                               
     Freight Status Report     
                               
 SUCCESS  Valid Git repository
 INFO  Railcar Manifest (railcar.json): Found
 INFO  Fingerprint (.git/hooks/.fingerprint.yaml): Found
 INFO  Conductor Binary: Found
 INFO  Last Applied Freight Version: dev
 INFO  Current Freight Version: dev

# Git Hooks Status

Hook               | Exists | Managed by Freight | Drift
commit-msg         | ✔      | ✔                  | none 
post-checkout      | ✔      | ✔                  | none 
post-commit        | ✔      | ✔                  | none 
pre-commit         | ✔      | ✔                  | none 
prepare-commit-msg | ✔      | ✔                  | none 
```
