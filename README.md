# Freight 🚂

[![Go Version](https://img.shields.io/github/go-mod/go-version/devbytes-cloud/freight)](https://go.dev/)
[![License](https://img.shields.io/github/license/devbytes-cloud/freight)](https://opensource.org/licenses/BSD-3-Clause)

**The professional, zero-dependency Git hook manager for modern teams.**

Freight streamlines Git workflows by rewiring every Git hook in your repository to a single **Conductor** binary. All logic is defined in a declarative **Railcar** manifest (`railcar.json`), ensuring your hooks are portable, fast, and easy to manage.

## Why Freight?

### 🚀 Zero Runtime Dependencies
Unlike Husky (which requires Node.js) or pre-commit (which requires Python), Freight is a single, static Go binary. Your developers don't need to install a specific runtime just to run Git hooks.

### 📦 Unified Configuration
Manage every hook—from `pre-commit` to `post-merge`—in one `railcar.json` manifest. No more messy `.git/hooks` directory filled with ad-hoc scripts.

### 🛠️ Built for Portability
Freight's 'Conductor/Railcar' architecture ensures that your hooks work identically across Windows, macOS, and Linux.

### 🥊 Freight vs. Husky
| Feature | Freight | Husky |
|---------|---------|-------|
| **Runtime** | None (Static Binary) | Node.js |
| **Setup** | `freight init` | `npm install` |
| **Config** | Single JSON file | Multiple files/package.json |
| **Portability** | High (Binary included) | Moderate (Requires Node) |

---

## Quick Start

### 1. Install
- **Homebrew (macOS):** `brew install --cask devbytes-cloud/tap/freight`
- **Precompiled Binaries:** `[GitHub Releases](https://github.com/devbytes-cloud/freight/releases)`

### 2. Setup
Run the following command in your Git repository:
```bash
freight init
```
This installs the **Conductor** binary and creates a starter **Railcar** manifest (`railcar.json`).

### 3. Verify
Add a script to your `railcar.json` and watch it run on your next commit!

---

## Architecture: Conductor & Railcar

Freight operates on a simple, powerful metaphor:
- **The Conductor:** A tiny, high-performance binary placed at your repo root. It is the single entry point for all Git hooks.
- **The Railcar:** A `railcar.json` manifest that defines exactly what the Conductor should execute for each hook.

When a Git hook fires, the Conductor extracts the logic from the Railcar and executes it with precision.

---

## Documentation

For full documentation, recipes, and advanced configuration, visit [freight.devbytes.cloud](https://freight.devbytes.cloud).

## License

BSD-style. See `LICENSE` for full text.
