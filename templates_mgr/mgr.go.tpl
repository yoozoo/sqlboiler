{{- $alias := .Aliases.Table .Table.Name -}}
{{- $g := . -}}
{{- $nouniqueFlag := false -}}
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

{{if len $pkTitles | ge 1}}
{{- $pkTitlePlural := $pkTitles | join "" | plural -}}
{{- $pkParamName := index .Table.PKey.Columns 0 -}}
{{- $pkParamType := index $colDefs.Types 0 -}}
{{- $pkArgsPlural := plural $pkParamName -}}
func ({{$alias.DownPlural}}Mgr) GetBy{{$pkTitlePlural}}({{$pkArgsPlural}} []{{$pkParamType}}, selectCols ...string) (map[{{$pkParamType}}]*{{$alias.UpSingular}}, error) {
	var o []*{{$alias.UpSingular}}

	if len({{$pkArgsPlural}}) == 0 {
		return nil, errors.New("models: no {{$g.Table.Name}} {{$pkParamName}} provided for GetBy{{$pkTitlePlural}}")
	}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from {{$g.Table.Name | $g.SchemaTable}} where `{{$pkParamName}}` in (?"+strings.Repeat(", ?", len({{$pkArgsPlural}})-1)+")", sel,
	)

	i := make([]interface{}, len({{$pkArgsPlural}}))
	for k, v := range {{$pkArgsPlural}} {
		i[k] = v
	}
	q := queries.Raw(query, i...)

	err := q.Bind(nil, {{- if $g.NoContext}}boil.GetDB(){{else}}boil.GetContextDB(){{end}}, &o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "{{$g.PkgName}}: unable to select from {{$g.Table.Name}}")
	}

	m := make(map[{{$pkParamType}}]*{{$alias.UpSingular}}, len(o))
	for _, v := range o {
		m[v.{{$pkTitles | join ""}}] = v
	}

	return m, nil
}
{{end}}

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

{{if len $titles | ge 1}}
{{- $paramName := index $args 0 -}}
{{- $paramType := index $colDefs.Types 0 -}}
{{- $titlePlural := $titles | join "" | plural -}}
{{- $argsPlural := plural $paramName -}}
func ({{$alias.DownPlural}}Mgr) GetBy{{$titlePlural}}({{$argsPlural}} []{{$paramType}}, selectCols ...string) (map[{{$paramType}}]*{{$alias.UpSingular}}, error) {
	var o []*{{$alias.UpSingular}}

	if len({{$argsPlural}}) == 0 {
		return nil, errors.New("models: no {{$g.Table.Name}} {{$paramName}} provided for GetBy{{$titlePlural}}")
	}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from {{$g.Table.Name | $g.SchemaTable}} where `{{$idx.Columns | join ""}}` in (?"+strings.Repeat(", ?", len({{$argsPlural}})-1)+")", sel,
	)

	i := make([]interface{}, len({{$argsPlural}}))
	for k, v := range {{$argsPlural}} {
		i[k] = v
	}
	q := queries.Raw(query, i...)

	err := q.Bind(nil, {{- if $g.NoContext}}boil.GetDB(){{else}}boil.GetContextDB(){{end}}, &o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "{{$g.PkgName}}: unable to select from {{$g.Table.Name}}")
	}

	m := make(map[{{$paramType}}]*{{$alias.UpSingular}}, len(o))
	for _, v := range o {
		m[v.{{$titles | join ""}}] = v
	}

	return m, nil
}
{{end}}
{{else}}
{{- $nouniqueFlag = true}}
func ({{$alias.DownPlural}}Mgr) FindBy{{ $titles | join ""}}({{ $funcArgs }}) (*{{$alias.DownPlural}}MgrQueries) {
	queries := []qm.QueryMod{
		{{- range $col := $idx.Columns}}
		qm.Where("{{$col}} = ?", authorID),
		{{- end }}
	}

	return &{{$alias.DownPlural}}MgrQueries{queries: queries}
}
{{ end -}}
{{ end -}}
{{ if $nouniqueFlag -}}
type {{$alias.DownPlural}}MgrQueries struct {
	queries []qm.QueryMod
}

// PaginateWithToal returns certain page of {{$alias.DownSingular}} record and total number of all record
func (q *{{$alias.DownPlural}}MgrQueries) PaginateWithToal(limit int, offset int) ({{$alias.UpSingular}}Slice, int64, error) {
	total, err := {{$alias.UpPlural}}(q.queries...).Count(context.TODO(), boil.GetContextDB())
	if err != nil {
		return nil, 0, err
	}

	q.queries = append(q.queries, qm.Limit(limit), qm.Offset(offset))
	{{$alias.DownPlural}}, err := {{$alias.UpPlural}}(q.queries...).All(nil, boil.GetContextDB())
	if err != nil {
		return nil, 0, err
	}

	return {{$alias.DownPlural}}, total, nil
}

// Paginate returns certain page of {{$alias.DownSingular}} record
func (q *{{$alias.DownPlural}}MgrQueries) Paginate(limit int, offset int) ({{$alias.UpSingular}}Slice, error) {
	q.queries = append(q.queries, qm.Limit(limit), qm.Offset(offset))
	{{$alias.DownPlural}}, err := {{$alias.UpPlural}}(q.queries...).All(nil, boil.GetContextDB())
	if err != nil {
		return nil, 0, err
	}

	return {{$alias.DownPlural}}, nil
}

// All returns all {{$alias.DownSingular}} record
func (q *{{$alias.DownPlural}}MgrQueries) All() ({{$alias.UpSingular}}Slice, error) {
	return {{$alias.UpPlural}}(q.queries...).All(nil, boil.GetContextDB())
}

func (q *{{$alias.DownPlural}}MgrQueries) OrderBy(orderClause string) *{{$alias.DownPlural}}MgrQueries {
	q.queries = append(q.queries, qm.OrderBy(orderClause))
	return q
}

func (q *{{$alias.DownPlural}}MgrQueries) Select(cols ...string) *{{$alias.DownPlural}}MgrQueries {
	q.queries = append(q.queries, qm.Select(cols...))
	return q
}

func (q *{{$alias.DownPlural}}MgrQueries) Where(queries ...qm.QueryMod) *{{$alias.DownPlural}}MgrQueries {
	q.queries = append(q.queries, queries...)
	return q
}
{{ end }}
