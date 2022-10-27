package bootstrapper

import (
	"fmt"
	"strings"

	"github.com/grokify/mogo/type/stringsutil"
)

type Generic struct {
	ElementBase
}

func (gen *Generic) Class() (string, error) {
	classes := []string{}
	classes = append(classes, gen.AdditionalClasses...)
	classes = stringsutil.SliceCondenseSpace(classes, true, true)
	return strings.Join(classes, " "), nil
}

func (gen *Generic) HTML() (string, error) {
	cls, err := gen.Class()
	if err != nil {
		return "", err
	}
	classHTML := ""
	if len(cls) > 0 {
		classHTML = fmt.Sprintf(` class="%s"`, cls)
	}
	styleHTML := ""
	if len(gen.Style) > 0 {
		styleHTML = fmt.Sprintf(` style='%s"`, gen.Style)
	}
	addlHTML := ""
	for k, v := range gen.AdditionalProperties {
		addlHTML += fmt.Sprintf(` %s="%s"`, k, v)
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
	return fmt.Sprintf(`<div%s%s%s>%s</div>`, classHTML, styleHTML, addlHTML, innerHTML), nil
}
