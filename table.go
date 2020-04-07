package table

// Table interface
type Table interface {
	NewRecord() []string
	AddFieldHeader(name string)
	AddFieldValueByName(name, value string) error
	AddFieldValueByIndex(index int, value string) error
}

type table struct {
	header  []string
	records []string
}

func (t *table) AddFieldHeader(name string) {
	if t.header == nil {
		t.header = make([]string, 0)
	}
	t.header = append(t.header, name)
}

func (t *table) AddFieldValueByName(name, value string) error {
	return nil
}

func (t *table) AddFieldValueByIndex(index int, value string) error {
	return nil
}
