package bootstrapper

import (
	"fmt"
	"strings"

	"github.com/grokify/mogo/type/maputil"
	"github.com/grokify/mogo/type/stringsutil"
)

type Generic struct {
	TagName string
	ElementBase
}

func (gen *Generic) Class() (string, error) {
	classes := []string{}
	classes = append(classes, gen.AdditionalClasses...)
	classes = stringsutil.SliceCondenseSpace(classes, true, true)
	return strings.Join(classes, " "), nil
}

func (gen *Generic) HTML() (string, error) {
	attrs := map[string]string{}
	for k, v := range gen.AdditionalProperties {
		attrs[k] = v
	}
	cls, err := gen.Class()
	if err != nil {
		return "", err
	} else if len(cls) > 0 {
		attrs["class"] = cls
	}
	innerHTML := ""
	for _, el := range gen.InnerHTML {
		if el == nil {
			continue
		}
		elHTML, err := el.HTML()
		if err != nil {
			return "", err
		}
		if len(innerHTML) > 0 {
			innerHTML += elHTML
		}
	}
	tag := strings.TrimSpace(gen.TagName)
	if len(tag) == 0 {
		tag = "div"
	}
	return fmt.Sprintf(`%s%s</%s>`, TagOpening(tag, attrs, false), innerHTML, tag), nil
}

// TagOpening creates HTML for an opening tag.
func TagOpening(tagName string, attrs map[string]string, close bool) string {
	parts := []string{tagName}
	if len(attrs) > 0 {
		keys := maputil.StringKeys(attrs, nil, true)
		for _, key := range keys {
			val, ok := attrs[key]
			if !ok {
				panic("extracted key not found")
			}
			parts = append(parts, fmt.Sprintf(`%s="%s"`, key, val))
		}
	}
	if close {
		parts = append(parts, "/")
	}
	if len(parts) == 0 {
		return ""
	}
	return "<" + strings.Join(parts, " ") + ">"
}
