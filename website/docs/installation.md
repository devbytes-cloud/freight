---
sidebar_position: 2
---

# Installation & Setup

Getting started with Freight is a three-step process designed for maximum speed and zero friction.

### 1. Install

#### Homebrew (macOS)
Install Freight as a Cask:
```bash
brew install --cask devbytes-cloud/tap/freight
```

#### GitHub Releases
You can also download the pre-compiled binaries directly from the [GitHub Releases](https://github.com/devbytes-cloud/freight/releases) page.

#### Build from source

If you prefer to build Freight from source:

```bash
git clone https://github.com/devbytes-cloud/freight.git
cd freight
go mod tidy
make build-all
```

### 2. Setup
Navigate to any existing Git repository and initialize Freight:

```bash
freight init
```

This command performs the following actions:
- Extracts the **Conductor** binary to your repository root.
- Generates a starter **Railcar** manifest (`railcar.json`).
- Rewires your `.git/hooks` to point to the Conductor.

:::tip Pro-Tip
For total team portability, **commit the `conductor` binary** directly to your repository. This ensures that every team member (and your CI/CD pipeline) can execute hooks immediately without needing to install the `freight` CLI tool themselves.
:::

### 3. Verify
Add a simple command to your `railcar.json` and trigger a hook:

```bash
git add .
git commit -m "Testing Freight"
```

---

## Command-line Reference

| Command        | Description                                 |
|----------------|---------------------------------------------|
| `freight init` | Bootstrap Freight in the current repository |
| `freight help` | Show global or command-specific help        |

### Global flags:

* `-c, --config-force` – overwrite an existing `railcar.json`
* `-h, --help` – display help
