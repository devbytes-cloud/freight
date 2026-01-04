# Best Practices

Follow these guidelines to ensure a smooth and consistent experience for your entire team using Freight.

### Version Control
It is crucial to commit both the `railcar.json` configuration file AND the `conductor` binary to your repository.

*   **railcar.json**: Defines what hooks run and what they do.
*   **conductor**: Ensures everyone is using the exact same orchestrator version without needing to install external dependencies like Node.js or Python.

:::info
By committing the binary, you guarantee that Git hooks will work immediately upon cloning the repository, regardless of the developer's local environment.
:::

### Team Onboarding
When a new team member clones the project, they need to initialize the hooks. You can provide a simple one-liner in your `README.md` or onboarding documentation:

```bash
./freight init
```

:::tip
Adding this to your project's `bootstrap` or `setup` script ensures it's never forgotten.
:::
