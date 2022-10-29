package bootstrapper

import (
	"strings"

	"github.com/grokify/mogo/html/htmlutil"
)

type Panel struct {
	htmlutil.Element
}

func NewPanel() *Panel {
	p := &Panel{}
	p.TagName = htmlutil.TagDiv
	p.AddAttribute(htmlutil.AttributeClass, "panel")
	return p
}

func (p *Panel) Color(c string) error {
	if IsColor(c) {
		p.AddAttribute(htmlutil.AttributeClass, "panel-"+c)
		return nil
	}
	return ErrColorUnknown
}

func (p *Panel) AddTitle(text string, escaped bool, tagName string) {
	elText := &htmlutil.Text{
		Text:    text,
		Escaped: escaped}

	tagName = strings.TrimSpace(tagName)
	if len(tagName) == 0 {
		tagName = "h1"
	}

	elTitle := &htmlutil.Element{}
	elTitle.TagName = tagName
	elTitle.AddAttribute(htmlutil.AttributeClass, ClassPanelTitle)
	elTitle.AddInnerHTML(elText)

	elHeading := &htmlutil.Element{}
	elHeading.TagName = htmlutil.TagDiv
	elHeading.AddAttribute(htmlutil.AttributeClass, ClassPanelHeading)
	elHeading.AddInnerHTML(elTitle)
	p.AddInnerHTML(elHeading)
}

func (p *Panel) AddBody(text string, escaped bool) {
	elBody := &htmlutil.Element{}
	elBody.TagName = htmlutil.TagDiv
	elBody.AddAttribute(htmlutil.AttributeClass, ClassPanelBody)
	elBody.AddInnerHTML(&htmlutil.Text{
		Text:    text,
		Escaped: escaped})
	p.AddInnerHTML(elBody)
}
