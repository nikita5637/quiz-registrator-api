{{- $i := .Data -}}
// Get{{ func_name_context $i }} retrieves a row from '{{ $i.Table.SQLName }}' as a {{ $i.Table.GoName }}.
//
// Generated from index '{{ $i.SQLName }}'.
{{ func_context $i }} {
	// query
	{{ sqlstr "index" $i }}
	// run
	logger.Debugf(ctx, sqlstr, {{ params $i.Fields false }})
{{- if $i.IsUnique }}
	{{ short $i.Table }} := {{ $i.Table.GoName }}{}
	if err := s.{{ db "QueryRow"  $i }}.Scan({{ names (print "&" (short $i.Table) ".") $i.Table }}); err != nil {
		return nil, err
	}
	return &{{ short $i.Table }}, nil
{{- else }}
	rows, err := s.{{ db "Query" $i }}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// process
	var res []*{{ $i.Table.GoName }}
	for rows.Next() {
		{{ short $i.Table }} := {{ $i.Table.GoName }}{}
		// scan
		if err := rows.Scan({{ names_ignore (print "&" (short $i.Table) ".")  $i.Table }}); err != nil {
			return nil, err
		}
		res = append(res, &{{ short $i.Table }})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
{{- end }}
}
