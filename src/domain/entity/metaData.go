package entity

type MetaData struct {
	Key   string `cql:"key" json:"key"`
	Kind  int32  `cql:"kind" json:"kind"`
	Value string `cql:"value" json:"value"`
}

func (e *MetaData) GetKey() string                  { return e.Key }
func (e *MetaData) SetKey(key string) *MetaData     { e.Key = key; return e }
func (e *MetaData) GetKind() int32                  { return e.Kind }
func (e *MetaData) SetKind(kind int32) *MetaData    { e.Kind = kind; return e }
func (e *MetaData) GetValue() string                { return e.Value }
func (e *MetaData) SetValue(value string) *MetaData { e.Value = value; return e }

func (e *MetaData) Make(key string, value string, kind int32) *MetaData {
	return e.SetKey(key).SetValue(value).SetKind(kind)
}
