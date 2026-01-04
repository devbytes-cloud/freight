---
sidebar_position: 3
---

# Configuration (Railcar Manifest)

The `railcar.json` file is the heart of Freight. It is a declarative manifest that defines exactly what the **Conductor** should execute for each Git hook.

## Syntax

* **config** – Root object.
* **\<operation-family\>** – Logical grouping of hooks (e.g., `commit-operations`, `checkout-operations`).
* **\<git-hook\>** – The specific Git hook name (e.g., `pre-commit`, `commit-msg`).
* **actions array** – A list of commands to run for the hook:
    * `name` – A human-readable label for the action.
    * `command` – The shell command or script to execute.

## Handling Arguments with `${HOOK_INPUT}`

Many Git hooks pass arguments to their scripts (e.g., `commit-msg` passes the path to the commit message file). Freight captures these and makes them available via the `${HOOK_INPUT}` variable.

**Example: Validating a commit message**
```json
{
  "commit-msg": [
    { 
      "name": "lint-message", 
      "command": "grep -E '^(feat|fix|docs|style|refactor|test|chore): ' ${HOOK_INPUT}" 
    }
  ]
}
```

---

## Configuration Recipes

Use these real-world examples to supercharge your `railcar.json`.

### Recipe: Golang Quality Gates
Run `golangci-lint` and tests before every commit.

```json
{
  "pre-commit": [
    { 
      "name": "lint", 
      "command": "golangci-lint run ./..." 
    },
    { 
      "name": "unit-tests", 
      "command": "go test -v ./..." 
    }
  ]
}
```

### Recipe: Conventional Commits (via grep)
Ensure every commit message follows the Conventional Commits specification without needing heavy dependencies.

```json
{
  "commit-msg": [
    { 
      "name": "conventional-check", 
      "command": "grep -E '^(feat|fix|chore|docs|refactor|test): .+' ${HOOK_INPUT}" 
    }
  ]
}
```

### Recipe: Security Scanning
Scan for secrets before pushing code to the remote.

```json
{
  "pre-push": [
    { 
      "name": "gitleaks", 
      "command": "gitleaks detect --source . -v" 
    }
  ]
}
```

---

## Supported Git Hooks

Freight supports the standard execution order for various Git operations:

* **Commit**: `pre-commit` → `prepare-commit-msg` → `commit-msg` → `post-commit`
* **Merge**: `pre-merge-commit` → `post-merge`
* **Rebase**: `pre-rebase` → `post-rewrite`
* **Push**: `pre-push` → `update` → `post-update` → `post-receive`
* **Checkout**: `pre-checkout` → `post-checkout`
* **ApplyPatch**: `applypatch-msg` → `pre-applypatch` → `post-applypatch`
