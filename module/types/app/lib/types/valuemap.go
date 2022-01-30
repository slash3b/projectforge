package types

var valueMapInst = &Map{K: NewString(), V: NewAny()}

const KeyValueMap = "valuemap"

var _ Type = (*ValueMap)(nil)

type ValueMap struct {
}

func (x *ValueMap) Key() string {
	return KeyValueMap
}

func (x *ValueMap) String() string {
	return valueMapInst.String()
}

func (x *ValueMap) Sortable() bool {
	return valueMapInst.Sortable()
}

func (x *ValueMap) Scalar() bool {
	return valueMapInst.Scalar()
}

func (x *ValueMap) From(v interface{}) interface{} {
	return valueMapInst.From(v)
}

func (x *ValueMap) Default(s string) interface{} {
	return valueMapInst.Default(s)
}

func NewValueMap() *Wrapped {
	return Wrap(&ValueMap{})
}
