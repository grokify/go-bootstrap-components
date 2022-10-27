package bootstrapper

import (
	"errors"
	"fmt"
	"strings"

	"github.com/grokify/mogo/type/stringsutil"
)

const (
	ColSizeXs   = "xs" // for phones - screens less than 768px wide
	ColSizeSm   = "sm" // for tablets - screens equal to or greater than 768px wide
	ColSizeMd   = "md" // for small laptops - screens equal to or greater than 992px wide
	ColSizeLg   = "lg" // for laptops and desktops - screens equal to or greater than 1200px wide
	ColMaxWidth = 12
)

type Column struct {
	ElementBase
	SizeDefaultWidth int
	SizeAllHidden    bool
	XsWidth          int
	XsHidden         bool
	SmWidth          int
	SmHidden         bool
	MdWidth          int
	MdHidden         bool
	LgWidth          int
	LgHidden         bool
}

var (
	ErrWidthTooLarge   = errors.New("max width of 12 exceeded")
	ErrWidthIsNegative = errors.New("width cannot be negative")
)

func (col *Column) Class() (string, error) {
	classes := []string{}
	var err error
	classes, err = colClassAppendNotErr(classes, ColSizeXs, col.XsWidth, col.SizeDefaultWidth)
	if err != nil {
		return "", err
	}
	classes, err = colClassAppendNotErr(classes, ColSizeSm, col.SmWidth, col.SizeDefaultWidth)
	if err != nil {
		return "", err
	}
	classes, err = colClassAppendNotErr(classes, ColSizeMd, col.MdWidth, col.SizeDefaultWidth)
	if err != nil {
		return "", err
	}
	classes, err = colClassAppendNotErr(classes, ColSizeLg, col.LgWidth, col.SizeDefaultWidth)
	if err != nil {
		return "", err
	}
	if col.XsHidden || col.SizeAllHidden {
		classes = append(classes, "hidden-"+ColSizeXs)
	}
	if col.SmHidden || col.SizeAllHidden {
		classes = append(classes, "hidden-"+ColSizeSm)
	}
	if col.MdHidden || col.SizeAllHidden {
		classes = append(classes, "hidden-"+ColSizeMd)
	}
	if col.LgHidden || col.SizeAllHidden {
		classes = append(classes, "hidden-"+ColSizeLg)
	}
	classes = append(classes, col.AdditionalClasses...)
	classes = stringsutil.SliceCondenseSpace(classes, true, true)
	return strings.Join(classes, " "), nil
}

func colClassAppendNotErr(s []string, size string, width, defaultWidth int) ([]string, error) {
	computedWidth := width
	if computedWidth <= 0 && defaultWidth > 0 {
		computedWidth = defaultWidth
	}
	cls, err := ColWidthToClass(size, computedWidth)
	if err != nil {
		return s, err
	} else if len(cls) > 0 {
		s = append(s, cls)
	}
	return s, nil
}

func (col *Column) HTML() (string, error) {
	cls, err := col.Class()
	if err != nil {
		return "", err
	}
	classHTML := ""
	if len(cls) > 0 {
		classHTML = fmt.Sprintf(` class="%s"`, cls)
	}
	styleHTML := ""
	if len(col.Style) > 0 {
		styleHTML = fmt.Sprintf(` style='%s"`, col.Style)
	}
	addlHTML := ""
	for k, v := range col.AdditionalProperties {
		addlHTML += fmt.Sprintf(` %s="%s"`, k, v)
	}
	innerHTML := ""
	for _, el := range col.InnerHTML {
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

func ColWidthToClass(size string, width int) (string, error) {
	if width > ColMaxWidth {
		return "", ErrWidthTooLarge
	} else if width < 0 {
		return "", ErrWidthIsNegative
	} else if width == 0 {
		return "", nil
	}
	return fmt.Sprintf("col-%x-%d", size, width), nil
}
