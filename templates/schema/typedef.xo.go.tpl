{{- $t := .Data -}}
{{- if $t.Comment -}}
// {{ $t.Comment | eval $t.GoName }}
{{- else -}}
// {{ $t.GoName }} represents a row from '{{ $t.SQLName }}'.
{{- end }}
type {{ $t.GoName }} struct {
{{ range $t.Fields -}}
	{{ field . }}
{{ end }}
}

// {{ $t.GoName }}Storage is {{ $t.GoName }} service implementation
type {{ $t.GoName }}Storage struct {
	db *tx.Manager
}

// New{{ $t.GoName }}Storage creates new instance of {{ $t.GoName }}Storage
func New{{ $t.GoName }}Storage(db *sql.DB) *{{ $t.GoName }}Storage {
	return &{{ $t.GoName }}Storage{
		db: tx.NewManager(db),
	}
}

// GetAll returns all records
func (s *{{ $t.GoName }}Storage) GetAll(ctx context.Context) ([]{{ $t.GoName }}, error) {
    return s.Find(ctx, nil, "")
}

// Find perform find request by params
func (s *{{ $t.GoName }}Storage) Find(ctx context.Context, q builder.Cond, sort string) ([]{{ $t.GoName }}, error) {
	query := `{{ sqlstr_select $t }}`

	var args  []interface{}

	if q != nil {
		var where string
		var err error
		where, args, err = builder.ToSQL(q)
		if err != nil {
			return nil, err
		}
		query += ` WHERE ` + where
	}

	if sort != "" {
		query += ` ` + getOrderStmt(sort)
	}

	rows, err := s.db.Sync(ctx).QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []{{ $t.GoName }}

	for rows.Next() {
		var item {{ $t.GoName }}
		if err := rows.Scan(
			{{ range $t.Fields -}}
				&item.{{ .GoName }},
			{{ end }}
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// FindWithLimit perform find request by params, offset and limit
func (s *{{ $t.GoName }}Storage) FindWithLimit(ctx context.Context, q builder.Cond, sort string, offset, limit uint64) ([]{{ $t.GoName }}, error) {
	query := `{{ sqlstr_select $t }}`

	var args []interface{}

	if q != nil {
		var where string
		var err error
		where, args, err = builder.ToSQL(q)
		if err != nil {
			return nil, err
		}
		query += ` WHERE ` + where
	}

	if sort != "" {
		query += ` ` + getOrderStmt(sort)
	}

	if limit != 0 {
		query += ` OFFSET ? LIMIT ?`
		args = append(args, offset)
		args = append(args, limit)
	}

	rows, err := s.db.Sync(ctx).QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []{{ $t.GoName }}

	for rows.Next() {
		var item {{ $t.GoName }}
		if err := rows.Scan(
			{{ range $t.Fields -}}
				&item.{{ .GoName }},
			{{ end }}
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	return items, nil
}

// Total return count(*) by params
func (s *{{ $t.GoName }}Storage) Total(ctx context.Context, q builder.Cond) (uint64, error) {
	query := `SELECT count(*) FROM {{ $t.SQLName }}`

	var args  []interface{}

	if q != nil {
		var where string
		var err   error
		where, args, err = builder.ToSQL(q)
		if err != nil {
			return 0, err
		}
		query += ` WHERE ` + where
	}

	rows, err := s.db.Sync(ctx).QueryContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	var count uint64

	for rows.Next() {
		if err := rows.Scan(
			&count,
		); err != nil {
			return 0, err
		}
	}

	return count, nil
}

{{ if $t.PrimaryKeys -}}
// {{ func_name_context "Insert" }} inserts the {{ $t.GoName }} to the database.
func (s *{{ $t.GoName }}Storage) Insert(ctx context.Context, item {{ $t.GoName }}) (int, error) {
	// insert (primary key generated and returned by database)
	{{ sqlstr "insert" $t }}
	// run
	{{ logf $t $t.PrimaryKeys "CreatedAt" "UpdatedAt" "DeletedAt" }}

	res, err := s.{{ db_prefix "Exec" true $t }}
	if err != nil {
		return 0, err
	}
	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// {{ func_name_context "Update" }} updates a {{ $t.GoName }} in the database.
func (s *{{ $t.GoName }}Storage) Update(ctx context.Context, item {{ $t.GoName }}) error {
	// update with {{ if driver "postgres" }}composite {{ end }}primary key
	{{ sqlstr "update" $t }}
	// run
	{{ logf_update $t }}
	if _, err := s.{{ db_update "Exec" $t }}; err != nil {
		return err
	}

	return nil
}


// {{ func_name_context "Upsert" }} performs an upsert for {{ $t.GoName }}.
func (s *{{ $t.GoName }}Storage) Upsert(ctx context.Context, item {{ $t.GoName }}) error {
	// upsert
	{{ sqlstr "upsert" $t }}
	// run
	{{ logf $t }}
	if _, err := s.{{ db_prefix "Exec" false $t }}; err != nil {
		return err
	}

	return nil
}

// {{ func_name_context "Delete" }} deletes the {{ $t.GoName }} from the database.
func (s *{{ $t.GoName }}Storage) Delete(ctx context.Context, id int) error {
	{{ if hasfield "DeletedAt" $t -}}
	// update with {{ if driver "postgres" }}composite {{ end }}primary key
	{{ if hasfield "UpdatedAt" $t -}}
	const sqlstr = `{{ sqlstr_update_deleted_at $t }}`
	{{- else -}}
	const sqlstr = `{{ sqlstr_update_deleted_at_without_updated_at $t }}`
	{{- end }}
	{{- else -}}
	// delete with single primary key
	{{ sqlstr "delete" $t }}
	{{- end }}
	// run
	logger.Debugf(ctx, sqlstr, id)

	if _, err := s.db.Master(ctx).ExecContext(ctx, sqlstr, id); err != nil {
		return err
	}

	return nil
}
{{- end }}
