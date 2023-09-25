package sqltypes

type Address string

func (t *Address) GormDataType() string {
	return "varchar(42)"
}

func (t Address) String() string {
	return string(t)
}

func (t Address) Equal(a Address) bool {
	return t.String() == a.String()
}
