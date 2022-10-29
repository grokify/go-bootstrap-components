package bootstrapper

import (
	"strings"

	"github.com/grokify/mogo/html/htmlutil"
)

// Button return ssomething like `<button type="button" class="btn btn-success" style="width:100%" onclick="this.blur()">MyText</button>``
type Button struct {
	htmlutil.Element
}

func NewButton() *Button {
	but := &Button{}
	but.TagName = "button"
	but.Attrs = map[string][]string{
		"class": {ClassButton},
		"type":  {"button"}}
	return but
}

func (but *Button) SetColor(c string) error {
	c = strings.ToLower(strings.TrimSpace(c))
	if IsColor(c) {
		but.AddAttribute(htmlutil.AttributeClass, ClassButton+"-"+c)
		return nil
	}
	return ErrColorUnknown
}
