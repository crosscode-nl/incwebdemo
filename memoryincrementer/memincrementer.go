package memoryincrementer

type MemIncrementer struct {
	count int
}

func (m *MemIncrementer) Inc() int {
	m.count = m.count + 1
	return m.count
}

func New() (*MemIncrementer) {
	return &MemIncrementer{count: 0}
}