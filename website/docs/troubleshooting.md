# Troubleshooting & FAQ

Learn how to resolve common issues and find answers to frequently asked questions about Freight.

:::info
If you encounter an issue not covered here, please feel free to open an issue on our [GitHub repository](https://github.com/devbytes-cloud/freight).
:::

### Hook not firing
If your Git hooks are not executing as expected, check the following:

1.  **Check if `.git/hooks` exists**: Freight relies on the standard Git hooks directory. Ensure this directory exists in your repository root.
2.  **Conductor Permissions**: The `conductor` binary must have execute permissions. You can fix this by running:
    ```bash
    chmod +x conductor
    ```
3.  **Correct Installation**: Ensure you have run the installation command to set up the hooks.

### Bypassing hooks
Sometimes you might need to skip hook execution (e.g., during an emergency hotfix). Standard Git flags work as expected with Freight:

*   **Commit**: `git commit --no-verify`
*   **Push**: `git push --no-verify`

:::caution
Bypassing hooks should be done sparingly, as it avoids important checks like linting and testing.
:::

### Environment variables
You can pass secrets or custom paths into `railcar.json` commands using environment variables. These variables are accessible within the shell context where the command is executed.

Example `railcar.json`:
```json
{
  "pre-push": [
    "npm run deploy -- --token=$DEPLOY_TOKEN"
  ]
}
```

### Permissions
If you see "Permission Denied" errors when a hook tries to run, it's likely that the `conductor` binary doesn't have the necessary execution bit set.

**Fixing Permission Denied:**
Run the following command in your project root:
```bash
chmod +x conductor
```
If you are on Windows, ensure your shell has the appropriate permissions to execute the binary.
