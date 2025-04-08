package templates

const Config = `{
  "railcar": {
    "commit-operations": {
      "pre-commit": [
        {
          "name": "echo",
          "command": "echo conductor is running!"
        }
      ]
    },
    "prepare-commit-msg": {
    },
    "commit-msg": {
    },
    "post-commit": {
    }
  }
}
`
