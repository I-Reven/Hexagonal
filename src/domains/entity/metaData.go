package entity

type MetaData struct {
	Key   string `cql:"key" json:"key"`
	Kind  int32  `cql:"kind" json:"kind"`
	Value string `cql:"value" json:"value"`
}

func (m *MetaData) GetKey() string                  { return m.Key }
func (m *MetaData) SetKey(key string) *MetaData     { m.Key = key; return m }
func (m *MetaData) GetKind() int32                  { return m.Kind }
func (m *MetaData) SetKind(kind int32) *MetaData    { m.Kind = kind; return m }
func (m *MetaData) GetValue() string                { return m.Value }
func (m *MetaData) SetValue(value string) *MetaData { m.Value = value; return m }

func (m *MetaData) Make(key string, value string, kind int32) *MetaData {
	return m.SetKey(key).SetValue(value).SetKind(kind)
}
