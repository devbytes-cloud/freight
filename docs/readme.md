## README

### Project Overview

Freight is a Go-based CLI application designed to streamline your Git workflow by modifying all Git hooks to point to a binary called `railcar` located in your project’s root directory. The `railcar` binary reads from a `railcar.json` file, also installed by Freight, where you can define actions for specific Git hooks. Each action is specified with a key (the action's name) and a corresponding value (the terminal command to be executed).

### Features

- **Git Hooks Setup**: Automatically configures Git hooks to invoke the `railcar` binary.
- **Configuration Management**: Creates and manages the `railcar.json` file for hook actions.
- **Binary Installation**: Installs the `railcar` binary to your project’s root directory.

### Installation 

#### Local

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

### Usage

After initializing a repository with `freight init`, two new files will be added to your repo:

- **railcar**: This binary is invoked by each Git hook, executing the actions defined for that specific hook in `railcar.json`.
- **railcar.json**: This file specifies the actions to be taken for each Git hook.

To initialize the application:

```sh
./freight init
```

#### Flags

- `--config-force` or `-c`: Force overwrite the `railcar.json` file if it already exists.

### Understanding and Using `railcar.json`

The `railcar.json` file is the central configuration for the `railcar` binary. It allows you to define specific actions to be executed during various Git hook events. The structure is hierarchical, with hooks and their associated actions organized under the `commit-operations` key.

#### Structure Breakdown

- **Top-Level Key**: `"carriage"` – contains all Git hook configurations.
- **Nested Key**: `"commit-operations"` – defines actions for different Git hooks.
- **Git Hook Keys**: `"pre-commit"`, `"prepare-commit-msg"`, `"commit-msg"`, `"post-commit"` – these represent specific Git hooks that can be configured.
- **Actions**: Each Git hook key can contain multiple actions, where the key is the action’s name, and the value is the command to be executed.

#### Example Configuration

Here’s how the provided `railcar.json` works:

```json
{
  "carriage": {
    "commit-operations": {
      "pre-commit": {
        "echo": "echo skiff is running!"
      },
      "prepare-commit-msg": {},
      "commit-msg": {},
      "post-commit": {
        "weather": "curl wttr.in"
      }
    }
  }
}
```

- **`pre-commit` Hook**: Prints "skiff is running!" to the terminal before committing.
- **`post-commit` Hook**: Fetches and displays the weather information from `wttr.in` after committing.
- **Empty Hooks**: `prepare-commit-msg` and `commit-msg` are defined but currently have no actions. These can be populated as needed.

#### Adding Actions

To add more actions, simply expand the `commit-operations` object. For example, to add a linting step to the `commit-msg` hook:

```json
"commit-msg": {
  "lint": "your-linting-command-here"
}
```

#### Usage Notes

- **Order of Execution**: Actions within a hook may be executed in any order. It is not guaranteed to run in the order they are defined.
- **Custom Commands**: You can define any terminal command as an action, providing flexibility for various project needs.
- **railcar.json**: If a `railcar.json` exists in your directory railcar will not overwrite it unless `--config-force` is set.
- 
### Example

To initialize the application with forced configuration writing:

```sh
./freight init 
```

### Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Open a pull request.

### License

This project is licensed under the BSD-style license. See the `LICENSE` file for more details.
