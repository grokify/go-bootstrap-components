package bootstrapper

import (
	"strings"

	"github.com/grokify/mogo/html/htmlutil"
	"golang.org/x/net/html/atom"
)

// Button return ssomething like `<button type="button" class="btn btn-success" style="width:100%" onclick="this.blur()">MyText</button>`
type Button struct {
	htmlutil.Element
}

func NewButton() *Button {
	but := &Button{}
	but.Element.TagName = atom.Button.String()
	but.Element.Attrs = map[string][]string{
		"class": {ClassButton},
		"type":  {"button"}}
	return but
}

func (but *Button) SetColor(c string) error {
	c = strings.ToLower(strings.TrimSpace(c))
	if IsColor(c) {
		return but.Element.AddAttribute(htmlutil.AttributeClass, ClassButton+"-"+c)
	}
	return ErrColorUnknown
}
