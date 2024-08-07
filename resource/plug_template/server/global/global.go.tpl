package global

{{- if .HasGlobal }}

import "leiserv/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}