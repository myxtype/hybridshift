package sqltypes

type Hash string

func (t *Hash) GormDataType() string {
	return "varchar(66)"
}

func (t Hash) String() string {
	return string(t)
}
