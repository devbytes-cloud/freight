package templates

const PreHookTmpl = `#!/bin/bash
echo "================================"
echo "Carriage {{ .Type }} Hook"
echo "================================"
{{ .Path }}/railcar {{ .Type }} $1`
