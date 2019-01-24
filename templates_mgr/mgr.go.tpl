{{- $alias := .Aliases.Table .Table.Name -}}
{{- $colDefs := sqlColDefinitions .Table.Columns .Table.PKey.Columns -}}
{{- $pkNames := $colDefs.Names | stringMap (aliasCols $alias) | stringMap .StringFuncs.camelCase | stringMap .StringFuncs.replaceReserved -}}
{{- $pkTitles := $colDefs.Names | stringMap (aliasCols $alias) | stringMap .StringFuncs.titleCase | stringMap .StringFuncs.replaceReserved -}}
{{- $pkArgs := joinSlices " " $pkNames $colDefs.Types | join ", "}}

// Manager Handy methods
type {{$alias.DownPlural}}Mgr int

var {{$alias.UpPlural}}Mgr {{$alias.DownPlural}}Mgr

func ({{$alias.DownPlural}}Mgr) GetBy{{$pkTitles | join ""}}({{$pkArgs}}) (*{{$alias.UpSingular}}, error) {
	{{- if .NoContext}}
	return Find{{$alias.UpSingular}}(boil.GetDB(), {{$pkNames | join ", "}})
	{{else}}
	return Find{{$alias.UpSingular}}(nil, boil.GetContextDB(), {{$pkNames | join ", "}})
	{{end -}}
}

