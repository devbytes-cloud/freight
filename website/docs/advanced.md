# Advanced Configuration

Take full control of your Git hooks with advanced configuration options and deep dives into Freight's execution model.

:::info
This page covers advanced topics. If you're just getting started, check out the [Configuration](./configuration.md) guide first.
:::

### Execution Order
Actions defined for a hook in `railcar.json` are executed **sequentially** in the order they appear in the list.

*   Each command must finish before the next one starts.
*   **The first non-zero exit code stops the entire chain.** If a command fails, subsequent commands in that hook will not be executed, and the Git operation will be aborted.

```json
{
  "config": {
    "commit-operations": {
      "pre-commit": [
        {
          "name": "lint",
          "command": "npm run lint"
        },
        {
          "name": "test",
          "command": "npm run test:unit"
        }
      ]
    }
  }
}
```

### Shell Context
The Conductor executes commands via the system's default shell:
*   **Unix-like (Linux/macOS)**: Uses `sh`.
*   **Windows**: Uses `cmd.exe`.

To ensure **cross-platform compatibility**, avoid shell-specific syntax when possible, or use tools that work across platforms (like `node` or `python` scripts).

:::caution
Be mindful of path separators and environment variable syntax differences between shells if your team uses multiple operating systems.
:::

### Hook Arguments
Some Git hooks receive input via stdin or arguments. Freight provides access to these through the `${HOOK_INPUT}` environment variable.

#### Deep Dive into `${HOOK_INPUT}`
Freight captures hook arguments and makes them available to your commands.

*   **commit-msg**: Receives the path to the commit message file.
*   **post-checkout**: Receives information about the previous and new HEAD.

You can use `${HOOK_INPUT}` in your `railcar.json` commands to pass this data to your scripts:

```json
{
  "config": {
    "commit-operations": {
      "commit-msg": [
        {
          "name": "validate",
          "command": "./scripts/validate-msg.sh \"${HOOK_INPUT}\""
        }
      ]
    }
  }
}
```
In your script, you can then parse this input to perform advanced validations based on the specific commits or branches being pushed.
