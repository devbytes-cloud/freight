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

#### One-liner Install (Linux/macOS/WSL)
Install Freight directly with a single command:
```bash
curl -fsSL https://raw.githubusercontent.com/devbytes-cloud/freight/main/curl.sh | bash
```

**Custom Install Directory:** By default, Freight installs to `/usr/local/bin`. You can customize this:
```bash
INSTALL_DIR=~/.local/bin curl -fsSL https://raw.githubusercontent.com/devbytes-cloud/freight/main/curl.sh | bash
```

The installer script:
- Auto-detects OS (Linux, macOS, Windows/MSYS)
- Auto-detects architecture (x86_64, arm64, armv6)
- Downloads the latest release from GitHub
- Installs to the specified directory (default: `/usr/local/bin`)
- May require `sudo` for system directories

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

By default, Freight installs all supported Git hooks. You can use the `--allow` flag to specify only the hooks you want:
```bash
freight init --allow pre-commit,commit-msg
```

This is particularly useful for **incremental adoption**. If you already have a complex set of hooks and only want to move `pre-commit` to Freight for now, you can do so without affecting your other hooks.

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

For the complete CLI reference, see [CLI Reference](./cli/index.md).
