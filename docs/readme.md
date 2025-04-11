# Freight

## Project Overview

Freight is a Go-based CLI application designed to streamline your Git workflow by modifying all Git hooks to point to a
binary called `conductor` located in your project's root directory. The `conductor` binary reads from a `railcar.json` file,
also installed by Freight, where you can define actions for specific Git hooks. Each action is specified with a key (the
action's name) and a corresponding value (the terminal command to be executed).

## Features

- **Git Hooks Setup**: Automatically configures Git hooks to invoke the `railcar` binary.
- **Configuration Management**: Creates and manages the `railcar.json` file for hook actions.
- **Binary Installation**: Installs the `railcar` binary to your project's root directory.

## Installation

### Local

1. **Clone the repository**:
   ```sh
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Install dependencies**:
   ```sh
   go mod tidy
   ```

3. **Build the application**:
   ```sh
   make build-all
   ```

## Usage

After initializing a repository with `freight init`, two new files will be added to your repo:

- **conductor**: This binary is invoked by each Git hook, executing the actions defined for that specific hook in
  `railcar.json`.
- **railcar.json**: This file specifies the actions to be taken for each Git hook.

To initialize the application, run:

```sh
./freight init
```

### Flags

- `--config-force` or `-c`: Force overwrite the railcar.json file if it already exists.

## Understanding and Using railcar.json

The `railcar.json` file is the central configuration for the conductor binary. It allows you to define specific actions to
be executed during various Git hook events. The structure is hierarchical, with hooks and their associated actions
organized under the commit-operations key.

### Structure Breakdown

- **Top-Level Key**: "config" – contains all Git hook configurations.
- **Nested Key**: "commit-operations" – defines actions for different Git hooks.
- **Git Hook Keys**: "pre-commit", "prepare-commit-msg", "commit-msg", "post-commit" – these represent specific Git
  hooks that can be configured.
- **Actions**: Each Git hook key can contain multiple actions, where the key is the action's name, and the value is the
  command to be executed.

### Example Configuration

Here's how the provided railcar.json works:

```json
{
  "config": {
    "commit-operations": {
      "pre-commit": [],
      "prepare-commit-msg": [],
      "commit-msg": [],
      "post-commit": []
    },
    "checkout-operations": {
      "post-checkout": [
        {
          "name": "hello",
          "command": "go env"
        }
      ]
    }
  }
}
```

- **Commit Hooks**: In this example, all commit-related hooks (such as pre-commit, commit-msg, etc.) are defined as
  empty arrays. You can add actions to these arrays as needed.
- **Checkout Hooks**: The post-checkout hook is configured with one action that runs the go env command, labeled with
  the name "hello".

### Adding Actions

To add an action to a hook, simply insert an object into the corresponding array. For example, to add a linting step to
the pre-commit hook, update the configuration as follows:

```json
"pre-commit": [
  {
    "name": "lint",
    "command": "golangci-lint run"
  }
]
```

Similarly, to add an action to the commit-msg hook:

```json
"commit-msg": [
  {
    "name": "format",
    "command": "gofmt -l ."
  }
]
```

### Usage Notes

- **Order of Execution**: Actions within a hook may be executed in any order.
- **Custom Commands**: You can define any terminal command as an action, providing flexibility for various project
  needs.
- **railcar.json Behavior**: If a railcar.json file already exists in your directory, the railcar binary will not
  overwrite it unless you specify `--config-force`.



### Example

To initialize the application with forced configuration writing, run:

```sh
./freight init --config-force
```

## Contributing

1. Fork the repository.
2. Create a new branch:
   ```sh
   git checkout -b feature-branch
   ```
3. Make your changes.
4. Commit your changes:
   ```sh
   git commit -m 'Add some feature'
   ```
5. Push to the branch:
   ```sh
   git push origin feature-branch
   ```
6. Open a pull request.

## License

This project is licensed under the BSD-style license. See the LICENSE file for more details.