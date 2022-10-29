package bootstrapper

import (
	"strings"

	"github.com/grokify/mogo/html/htmlutil"
	"golang.org/x/net/html/atom"
)

type Panel struct {
	htmlutil.Element
}

func NewPanel() *Panel {
	p := &Panel{}
	p.Element.TagName = atom.Div.String()
	err := p.Element.AddAttribute(htmlutil.AttributeClass, "panel")
	if err != nil {
		panic(err)
	}
	return p
}

func (p *Panel) Color(c string) error {
	if IsColor(c) {
		return p.Element.AddAttribute(htmlutil.AttributeClass, "panel-"+c)
	}
	return ErrColorUnknown
}

func (p *Panel) AddTitle(text string, escaped bool, tagName string) {
	elText := &htmlutil.Text{
		Text:    text,
		Escaped: escaped}

	tagName = strings.TrimSpace(tagName)
	if len(tagName) == 0 {
		tagName = atom.H1.String()
	}

	elTitle := &htmlutil.Element{}
	elTitle.TagName = tagName
	err := elTitle.AddAttribute(htmlutil.AttributeClass, ClassPanelTitle)
	if err != nil {
		panic(err)
	}
	elTitle.AddInnerHTML(elText)

	elHeading := &htmlutil.Element{}
	elHeading.TagName = atom.Div.String()
	err = elHeading.AddAttribute(htmlutil.AttributeClass, ClassPanelHeading)
	if err != nil {
		panic(err)
	}
	elHeading.AddInnerHTML(elTitle)
	p.Element.AddInnerHTML(elHeading)
}

func (p *Panel) AddBody(text string, escaped bool) {
	elBody := &htmlutil.Element{}
	elBody.TagName = atom.Div.String()
	err := elBody.AddAttribute(htmlutil.AttributeClass, ClassPanelBody)
	if err != nil {
		panic(err)
	}
	elBody.AddInnerHTML(&htmlutil.Text{
		Text:    text,
		Escaped: escaped})
	p.Element.AddInnerHTML(elBody)
}
