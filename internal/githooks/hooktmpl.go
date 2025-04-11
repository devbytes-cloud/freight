package githooks

// gitHookTemplate defines the template for a Git hook script.
// The template includes placeholders for the hook type (`{{ .Type }}`) and path (`{{ .Path }}`),
// which are replaced with actual values when the template is rendered.
const gitHookTemplate = `#!/bin/bash
echo "================================"
echo "Railcar {{ .Type }} Hook"
echo "================================"
{{ .Path }}/conductor {{ .Type }} $1`
