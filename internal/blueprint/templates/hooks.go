package templates

const PreHookTmpl = `#!/bin/bash
echo "================================"
echo "Carriage {{ .Type }} Hook"
echo "================================"
{{ .Path }}/conductor {{ .Type }} $1`
