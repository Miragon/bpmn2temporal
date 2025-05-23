package generator

type Field struct {
	Name string
	Type string
}

type StructModel struct {
	Name      string
	Fields    []Field
	LastIndex int
}
