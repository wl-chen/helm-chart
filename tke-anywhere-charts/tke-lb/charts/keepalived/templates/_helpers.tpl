{{/*
Create the vrid of the keepalived config
*/}}
{{- define "generate.vrid" -}}
  {{- if .Values.vip -}}
    {{- $split := splitList "." .Values.vip -}}
    {{- $num := index $split 3 -}}
    {{- $vrid := add $num 1 -}}
    {{- if and (lt 0 $vrid) (lt $vrid 256) -}}
      {{- print $vrid -}}
    {{- end -}}
  {{- end -}}
{{- end -}}

{{/*
Create the random num range [1, 255]
*/}}
{{- define "generate.random" -}}
  {{- $base := randNumeric 8 -}}
  {{- $num := add (mod $base 254) 1 -}}
  {{- print $num -}}
{{- end -}}