{{- define "query" -}}
{{- $table := . -}}
func Iterate{{ $table.GoName }}(sc interface{ Scan(...interface{}) error}) ({{ $table.GoName }}, error) {
    t := {{ $table.GoName }}{}
    if err := sc.Scan(t.Ptrs()...); err != nil {
        return {{ $table.GoName }}{}, dberror.MapError(err)
    }
    return t, nil
}

{{ range $index := $table.Indexes -}}
func Select{{ if $index.IsUnique }}One{{ end }}{{ $table.GoName }}By{{ range $i, $f := $index.Fields }} {{- $f.GoName -}} {{- if not ($index.Tail $i) -}} And {{- end }}{{ end }}(ctx context.Context, txn *sql.Tx, {{ range $i, $f := $index.Fields }} {{- $f.Name }} *{{ $f.GoType }} {{- if not ($index.Tail $i) -}} , {{- end }}{{ end }}) ({{ if not $index.IsUnique -}} []* {{- end }} {{- $table.GoName }}, error) {
    eq := squirrel.Eq{}
    {{- range $index.Fields }}
    if {{ .Name }} != nil {
        eq["{{ .Name }}"] = *{{ .Name }}
    }
    {{- end }}
    query, params, err := squirrel.
        Select({{ $table.GoName }}AllColumns...).
        From({{ $table.GoName}}TableName).
        Where(eq).
        ToSql()
    if err != nil {
        return {{ if not $index.IsUnique -}} nil {{- else -}} {{ $table.GoName }}{} {{- end -}}, dberror.MapError(err)
    }
    stmt, err := txn.PrepareContext(ctx, query)
    if err != nil {
        return {{ if not $index.IsUnique -}} nil {{- else -}} {{ $table.GoName }}{} {{- end -}}, dberror.MapError(err)
    }

    {{- if $index.IsUnique }}
    return Iterate{{ $table.GoName }}(stmt.QueryRowContext(ctx, params...))
    {{- else }}
    rows, err := stmt.QueryContext(ctx, params...)
    if err != nil {
        return nil, dberror.MapError(err)
    }
    res := make([]*{{ $table.GoName }}, 0)
    for rows.Next() {
        t, err := Iterate{{ $table.GoName }}(rows)
        if err != nil {
            return nil, dberror.MapError(err)
        }
        res = append(res, &t)
    }
    return res, nil
    {{- end }}
}

{{ end }}
{{ end }}