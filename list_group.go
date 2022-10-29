package bootstrapper

import (
	"github.com/grokify/mogo/html/htmlutil"
	"golang.org/x/net/html/atom"
)

type ListGroup struct {
	htmlutil.Element
}

func NewListGroup() *ListGroup {
	grp := &ListGroup{}
	grp.TagName = atom.Div.String()
	grp.AddAttribute(htmlutil.AttributeClass, ClassListGroup)
	return grp
}

func (lg *ListGroup) AddItemText(text string, escaped, active bool) {
	elText := &htmlutil.Element{}
	elText.TagName = atom.P.String()
	elText.AddAttribute(htmlutil.AttributeClass, ClassListGroupItemText)
	elText.AddInnerHTMLText(text, escaped)

	wrapper := &htmlutil.Element{}
	wrapper.TagName = atom.A.String()
	wrapper.AddAttribute(htmlutil.AttributeHref, "#")
	wrapper.AddAttribute(htmlutil.AttributeClass, ClassListGroupItem)
	if active {
		wrapper.AddAttribute(htmlutil.AttributeClass, ClassActive)
	}
	wrapper.AddInnerHTML(elText)

	lg.AddInnerHTML(wrapper)
}
