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
	grp.Element.TagName = atom.Div.String()
	err := grp.Element.AddAttribute(htmlutil.AttributeClass, ClassListGroup)
	if err != nil {
		panic(err)
	}
	return grp
}

func (grp *ListGroup) AddItemText(text string, escaped, active bool) {
	elText := &htmlutil.Element{}
	elText.TagName = atom.P.String()
	err := elText.AddAttribute(htmlutil.AttributeClass, ClassListGroupItemText)
	if err != nil {
		panic(err)
	}
	elText.AddInnerHTMLText(text, escaped)

	wrapper := &htmlutil.Element{}
	wrapper.TagName = atom.A.String()
	err = wrapper.AddAttribute(htmlutil.AttributeHref, "#")
	if err != nil {
		panic(err)
	}
	err = wrapper.AddAttribute(htmlutil.AttributeClass, ClassListGroupItem)
	if err != nil {
		panic(err)
	}
	if active {
		err = wrapper.AddAttribute(htmlutil.AttributeClass, ClassActive)
		if err != nil {
			panic(err)
		}
	}
	wrapper.AddInnerHTML(elText)

	grp.Element.AddInnerHTML(wrapper)
}
