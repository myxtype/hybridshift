package sqltypes

import (
	"database/sql/driver"
	"frame/pkg/sql/format"
)

type Attribute struct {
	TraitType string      `json:"trait_type"`
	Value     interface{} `json:"value"`
}

type Attributes []Attribute

func (t Attributes) GormDataType() string {
	return "jsonb"
}

func (t *Attributes) Scan(src interface{}) error {
	return format.Scan(t, src)
}

func (t Attributes) Value() (driver.Value, error) {
	return format.Value(t)
}

func (t Attributes) GetValue(traitType string) interface{} {
	for _, n := range t {
		if n.TraitType == traitType {
			return n.Value
		}
	}
	return nil
}
