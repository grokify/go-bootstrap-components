package bootstrapper

import "github.com/grokify/mogo/html/htmlutil"

type ListGroup struct {
	htmlutil.Element
}

func NewListGroup() *ListGroup {
	grp := &ListGroup{}
	grp.TagName = htmlutil.TagDiv
	grp.AddAttribute(htmlutil.AttributeClass, ClassListGroup)
	return grp
}

func (lg *ListGroup) AddItemText(text string, escaped, active bool) {
	elText := &htmlutil.Element{}
	elText.TagName = "p"
	elText.AddAttribute(htmlutil.AttributeClass, ClassListGroupItemText)
	elText.AddInnerHTMLText(text, escaped)

	wrapper := &htmlutil.Element{}
	wrapper.TagName = "a"
	wrapper.AddAttribute(htmlutil.AttributeHref, "#")
	wrapper.AddAttribute(htmlutil.AttributeClass, ClassListGroupItem)
	if active {
		wrapper.AddAttribute(htmlutil.AttributeClass, ClassActive)
	}
	wrapper.AddInnerHTML(elText)

	lg.AddInnerHTML(wrapper)
}
