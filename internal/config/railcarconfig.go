package config

// RailcarJson is a JSON string that defines the default configuration for RailCar.
// It includes configurations for commit-related operations (e.g., pre-commit, post-commit)
// and checkout-related operations (e.g., post-checkout). Each operation can specify
// a sequence of hook steps, where each step has a name and a command to execute.
const RailcarJson = `{
  "config": {
    "commit-operations": {
      "pre-commit": [
        {
          "name": "echo",
          "command": "echo conductor is running!"
        }
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
`
