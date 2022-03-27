package models

type Element struct {
	Name    string
	Content string

	// kinda polymorphism
	Type ElementType

	SubElements []*Element

	// Rows is set up when element is a Table
	Rows []*Row
}

func (element *Element) GetType() ElementType {
	return element.Type
}

func (element *Element) GetContent() string {
	return element.Content
}

func (element *Element) SetContent(content string) {
	element.Content = content
}

func (element *Element) AppendContent(content string) {
	element.Content = element.Content + content
}

func (element *Element) GetSubElements() []*Element {
	return element.SubElements
}
