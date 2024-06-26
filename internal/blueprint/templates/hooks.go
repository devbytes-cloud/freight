package templates

const PreHookTmpl = `#!/bin/bash
echo "================================"
echo "Skiff {{ .Type }} Hook"
echo "================================"
{{ .Path }}/parser {{ .Type }} $1`
