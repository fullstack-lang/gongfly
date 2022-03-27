package models

type MarkdownContent struct {
	Name    string
	Content string

	Root *Element
}

// UpdateContent updates Content from the root
func (markdownContent *MarkdownContent) UpdateContent() {

	markdownContent.Content = ""

	// verification that there are no cycles in the graph

	if markdownContent.Root != nil {
		markdownContent.RecursiveUpdateContent(markdownContent.Root, 1)
	}

}

func (markdownContent *MarkdownContent) RecursiveUpdateContent(element *Element, depth int) {

	switch element.GetType() {
	case TITLE:
		for level := 0; level < depth; level++ {
			markdownContent.Content = markdownContent.Content + "#"
		}
		markdownContent.Content = markdownContent.Content + " "
	case TABLE:
		if len(element.Rows) == 0 {
			break
		}

		// create title row
		markdownContent.Content = markdownContent.Content + "|"
		for _, cell := range element.Rows[0].Cells {
			markdownContent.Content = markdownContent.Content + cell.Name + "|"
		}
		markdownContent.Content = markdownContent.Content + "\n"

		// creates line under title row
		markdownContent.Content = markdownContent.Content + "|"
		for idx := 0; idx < len(element.Rows[0].Cells); idx++ {
			markdownContent.Content = markdownContent.Content + "-" + "|"
		}
		markdownContent.Content = markdownContent.Content + "\n"

		// create one line per remaining row
		for rowId := 1; rowId < len(element.Rows); rowId++ {
			for _, cell := range element.Rows[rowId].Cells {
				markdownContent.Content = markdownContent.Content + cell.Name + "|"
			}
			markdownContent.Content = markdownContent.Content + "\n"
		}
		markdownContent.Content = markdownContent.Content + "\n"

	}

	markdownContent.Content = markdownContent.Content + element.GetContent()
	markdownContent.Content = markdownContent.Content + "\n\n"

	for _, subElement := range element.GetSubElements() {
		markdownContent.RecursiveUpdateContent(subElement, depth+1)
	}

}
