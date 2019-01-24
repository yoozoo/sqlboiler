{{- $alias := .Aliases.Table .Table.Name -}}
{{- $g := . -}}
{{- $colDefs := sqlColDefinitions .Table.Columns .Table.PKey.Columns -}}
{{- $pkNames := $colDefs.Names | stringMap (aliasCols $alias) | stringMap .StringFuncs.camelCase | stringMap .StringFuncs.replaceReserved -}}
{{- $pkTitles := $colDefs.Names | stringMap (aliasCols $alias) | stringMap .StringFuncs.titleCase | stringMap .StringFuncs.replaceReserved -}}
{{- $pkArgs := joinSlices " " $pkNames $colDefs.Types | join ", "}}

// Manager Handy methods
type {{$alias.DownPlural}}Mgr int

var {{$alias.UpPlural}}Mgr {{$alias.DownPlural}}Mgr

func ({{$alias.DownPlural}}Mgr) GetBy{{$pkTitles | join ""}}({{$pkArgs}}, selectCols ...string) (*{{$alias.UpSingular}}, error) {
	{{- if .NoContext}}
	return Find{{$alias.UpSingular}}(boil.GetDB(), {{$pkNames | join ", "}}, selectCols...)
	{{else}}
	return Find{{$alias.UpSingular}}(nil, boil.GetContextDB(), {{$pkNames | join ", "}}, selectCols...)
	{{end -}}
}

{{- range $idx := .Table.Indexes}}
{{- $titles := $idx.Columns | stringMap $g.StringFuncs.titleCase | stringMap $g.StringFuncs.replaceReserved -}}
{{- $args := $idx.Columns | stringMap $g.StringFuncs.camelCase | stringMap $g.StringFuncs.replaceReserved -}}
{{- $colDefs := sqlColDefinitions $g.Table.Columns $idx.Columns -}}
{{- $funcArgs := joinSlices " " $args $colDefs.Types | join ", "}}
{{if $idx.Unique}}

func ({{$alias.DownPlural}}Mgr) GetBy{{ $titles | join ""}}({{ $funcArgs }}, selectCols ...string) (*{{$alias.UpSingular}}, error) {
	{{$alias.DownSingular}}Obj := &{{$alias.UpSingular}}{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from {{$g.Table.Name | $g.SchemaTable}} where {{if $g.Dialect.UseIndexPlaceholders}}{{whereClause $g.LQ $g.RQ 1 $idx.Columns}}{{else}}{{whereClause $g.LQ $g.RQ 0 $idx.Columns}}{{end}}", sel,
	)

	q := queries.Raw(query, {{$args | join ", "}})

	err := q.Bind(nil, {{- if $g.NoContext}}boil.GetDB(){{else}}boil.GetContextDB(){{end}}, {{$alias.DownSingular}}Obj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "{{$g.PkgName}}: unable to select from {{$g.Table.Name}}")
	}

	return {{$alias.DownSingular}}Obj, nil
}
{{else}}

func ({{$alias.DownPlural}}Mgr) FindAllBy{{ $titles | join ""}}({{ $funcArgs }}, selectCols ...string) ([]*{{$alias.UpSingular}}, error) {
	var o []*{{$alias.UpSingular}}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from {{$g.Table.Name | $g.SchemaTable}} where {{if $g.Dialect.UseIndexPlaceholders}}{{whereClause $g.LQ $g.RQ 1 $idx.Columns}}{{else}}{{whereClause $g.LQ $g.RQ 0 $idx.Columns}}{{end}}", sel,
	)

	q := queries.Raw(query, {{$args | join ", "}})

	err := q.Bind(nil, {{- if $g.NoContext}}boil.GetDB(){{else}}boil.GetContextDB(){{end}}, &o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "{{$g.PkgName}}: unable to select from {{$g.Table.Name}}")
	}

	return o, nil
}

func ({{$alias.DownPlural}}Mgr) FindBy{{ $titles | join ""}}({{ $funcArgs }}, limit, offset int, selectCols ...string) ([]*{{$alias.UpSingular}}, error) {
	var o []*{{$alias.UpSingular}}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from {{$g.Table.Name | $g.SchemaTable}} where {{if $g.Dialect.UseIndexPlaceholders}}{{whereClause $g.LQ $g.RQ 1 $idx.Columns}}{{else}}{{whereClause $g.LQ $g.RQ 0 $idx.Columns}}{{end}} limit ?, ?", sel,
	)

	q := queries.Raw(query, {{$args | join ", "}}, offset, limit)

	err := q.Bind(nil, {{- if $g.NoContext}}boil.GetDB(){{else}}boil.GetContextDB(){{end}}, &o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "{{$g.PkgName}}: unable to select from {{$g.Table.Name}}")
	}

	return o, nil
}
{{ end -}}
{{ end }}
