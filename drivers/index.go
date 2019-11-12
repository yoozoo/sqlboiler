package drivers

type Index struct {
	Name    string   `json:"name" toml:"name"`
	Columns []string `json:"columns" toml:"columns"`
	Unique  bool     `json:"unique" toml:"unique"`
}
