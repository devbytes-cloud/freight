# Freight

## Project Overview

Freight is a Go-based CLI tool that streamlines Git workflows by rewiring every Git hook in your repository to call a
single binary named `conductor` (placed at the repo root).`conductor` reads a JSON configuration file (`railcar.json`)
that lists the shell commands you want to run for each
hook. All logic therefore lives in one easy-to-version file instead of a dozen ad-hoc hook scripts.

---

## Features

* **One-step hook bootstrap** – `freight init` installs `conductor`, generates `railcar.json`, and rewrites every Git
  hook in one go.
* **Declarative configuration** – add, remove, or reorder hook commands by editing JSON.
* **Cross-platform binaries** – pre-built for Linux, macOS, and Windows.
* **Positional-argument support** – hooks such as `commit-msg` or `pre-push` automatically pass their arguments to your
  commands (via `${HOOK_INPUT}`).
* **Zero vendor lock-in** – all generated files live inside your repo; deleting them restores the default Git behaviour.

---

## Installation

### From Source

```
git clone https://github.com/devbytes-cloud/freight.git
cd freight
go mod tidy
make build-all
```

---

## Quick Start

```
# Inside any existing Git repository
freight init                   # installs conductor + railcar.json + rewired hooks

# Optionally install only specific hooks
freight init --allow pre-commit,post-checkout

git add . && git commit -m "test"  # your new hooks will now fire
```

Need to overwrite an existing `railcar.json`?  
`freight init --config-force`

---

## Command-line Reference

| Command        | Description                                 |
|----------------|---------------------------------------------|
| `freight init` | Bootstrap Freight in the current repository |
| `freight help` | Show global or command-specific help        |

Global flags:

* `-c, --config-force` – overwrite an existing `railcar.json`
* `-a, --allow` – specific Git hooks to install (default: all). Valid options: `pre-commit`, `prepare-commit-msg`, `commit-msg`, `post-commit`, `post-checkout`
* `-h, --help` – display help

---

## How It Works (under the hood)

1. **Bootstrap (`freight init`)**
    * Places a self-contained `conductor` binary at your repo root.
    * Generates a starter `railcar.json`.
    * Replaces every file in `.git/hooks/` with a tiny wrapper script that simply executes `conductor` with the hook
      name and original arguments.

2. **Hook trigger** – Git fires `pre-commit`, `commit-msg`, etc.
    * The wrapper calls `conductor`.
    * `conductor` loads `railcar.json`, finds the matching section, and runs each configured action.
    * Any non-zero exit in a pre-hook aborts the Git operation.

---

## `railcar.json` Syntax

Hierarchical structure

* **config** – top level
* **\<operation-family\>** – e.g. `commit-operations`, `checkout-operations`
* **\<git-hook\>** – e.g. `pre-commit`, `commit-msg`, `post-checkout`
* **actions array** – each item needs:
    * `name` – label for readability
    * `command` – shell snippet to run

Example starter file:

```
{
  "config": {
    "commit-operations": {
      "pre-commit": [
        { "name": "echo", "command": "echo conductor is running!" }
      ],
      "prepare-commit-msg": [],
      "commit-msg": [],
      "post-commit": []
    },
    "checkout-operations": {
      "post-checkout": []
    }
  }
}
```

### Referencing Hook Arguments

Hooks that receive parameters expose them in two interchangeable ways:

| Placeholder     | Meaning (example: `commit-msg`) |
|-----------------|---------------------------------|
| `${HOOK_INPUT}` | Alias for the  parameter (`$1`) |

Use whichever style you prefer:

```
{
  "commit-msg": [
    { "name": "validate", "command": "grep -E '^(feat|fix): ' ${HOOK_INPUT}" },
    { "name": "print",    "command": "echo 'MSG file → ${HOOK_INPUT}'" }
  ]
}
```

### Real-world Examples

* Run tests before committing

```
{
    "pre-commit": [
      { "name": "tests", "command": "go test ./..." }
    ]
  }
```

* Enforce Conventional Commits format

```
{
    "commit-msg": [
      { "name": "conventional", "command": "npx commitlint --edit $1" }
    ]
  }
```

* Verify tags before pushing

```
{
    "pre-push": [
      { "name": "verify-tags", "command": "./scripts/check_tags.sh $@" }
    ]
  }
```

Notes

* **Order** – actions execute sequentially in array order.
* **Shell chaining** – combine commands (`go vet ./... && go test ./...`).
* **Environment vars** – standard shell expansion works (`FOO=bar ./script.sh`).
* **Idempotency** – `freight init` never overwrites `railcar.json` unless `--config-force` is supplied.

---

## Supported Git Hooks & Execution Order

```
Commit     : pre-commit → prepare-commit-msg → commit-msg → post-commit
Merge      : pre-merge-commit → post-merge
Rebase     : pre-rebase → post-rewrite
Push       : pre-push → update → post-update → post-receive
Checkout   : pre-checkout → post-checkout
ApplyPatch : applypatch-msg → pre-applypatch → post-applypatch
```

---

## Troubleshooting

| Issue                        | Fix                                                                               |
|------------------------------|-----------------------------------------------------------------------------------|
| Permission denied on hooks   | `chmod +x ./conductor` (and ensure hooks are executable)                          |
| Hook seems to do nothing     | Check `.git/hooks/<hook>` – it should contain the wrapper that calls `conductor`. |
| Command not found            | Ensure the command exists in `$PATH` or use an absolute path in `railcar.json`.   |
| Need to debug a failing hook | Run the failing hook script manually or add `set -x` inside your action command.  |

---

## Contributing

```
git clone https://github.com/yourusername/freight.git
git checkout -b my-feature
# make changes
git commit -s -m "feat: awesome contribution"
git push origin my-feature
```

Open a Pull Request—thank you!

### Build & Release

### 1. Release & Distribution (GoReleaser + ) `go:embed`

- The repository contains a that builds **platform-specific `conductor` binaries** for Linux, macOS and Windows (amd64,
  arm, arm64). `.goreleaser.yaml`
- These binaries are dropped into `assets/dist/` and then **embedded directly into the main `freight` executable** via
  Go’s mechanism (). `//go:embed``assets/embed.go`
- At runtime, `freight init` extracts the correct pre-built `conductor` for the user’s OS/CPU.
- CGO is disabled () so binaries are fully statically linked and portable. `CGO_ENABLED=0`

### 2. Hook-Generation Template

- Git hooks are produced from a **single script template** (`internal/githooks/gitHookTemplate`) so every generated hook
  is tiny, consistent, and easy to audit.
- Unit tests in verify that every hook file is generated with the expected path and template. `githooks_test.go`

A short note in the README can highlight the attention to reliability and test coverage.

### 3. Testing

- The project ships with a Go test suite () that covers hook generation, path handling, and validation helpers.
  `go test ./...`
- Mentioning this encourages contributors to run tests before submitting pull requests.

### 4. Make Targets

- If you already have a , consider listing other useful targets (`make test`, `make lint`, etc.) so newcomers can find
  them quickly. `make build-all`

---

## License

BSD-style. See `LICENSE` for full text.