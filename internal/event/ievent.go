package event

type IEvent interface {
	GetDesc() string
	GetFromModuleName() string
	GetToModuleName() string
	GetCategory() int
	SetName(name string)
	SetCategory(category int)
}

type Base struct {
	Name     string
	Category int
}

func (b *Base) GetDesc() string {
	//TODO implement me
	panic("implement me")
}

func (b *Base) GetFromModuleName() string {
	return b.Name
}
func (b *Base) GetToModuleName() string {
	return b.Name
}

func (b *Base) GetCategory() int {
	return b.Category
}

func (b *Base) SetName(name string) {
	b.Name = name
}

func (b *Base) SetCategory(category int) {
	b.Category = category
}
