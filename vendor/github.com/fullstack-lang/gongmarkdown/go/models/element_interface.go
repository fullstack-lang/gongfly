package models

type ElementInterface interface {
	GetType() ElementType
	GetContent() string
	SetContent(content string)
	AppendContent(content string)
	GetSubElements() []ElementInterface
}
