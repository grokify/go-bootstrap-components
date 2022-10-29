package bootstrapper

import (
	"errors"
	"fmt"

	"github.com/grokify/mogo/html/htmlutil"
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
	htmlutil.Element
	//ElementBase
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

// Classes returns a list of column-related Bootstrap classes.
func (col *Column) Classes() ([]string, error) {
	classes := []string{}
	var err error
	classes, err = colClassAppendNotErr(classes, ColSizeXs, col.XsWidth, col.SizeDefaultWidth)
	if err != nil {
		return classes, err
	}
	classes, err = colClassAppendNotErr(classes, ColSizeSm, col.SmWidth, col.SizeDefaultWidth)
	if err != nil {
		return classes, err
	}
	classes, err = colClassAppendNotErr(classes, ColSizeMd, col.MdWidth, col.SizeDefaultWidth)
	if err != nil {
		return classes, err
	}
	classes, err = colClassAppendNotErr(classes, ColSizeLg, col.LgWidth, col.SizeDefaultWidth)
	if err != nil {
		return classes, err
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
	// classes = append(classes, col.AdditionalClasses...)
	classes = stringsutil.SliceCondenseSpace(classes, true, true)
	return classes, nil
	// return strings.Join(classes, " "), nil
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

func (col *Column) String() (string, error) {
	classes, err := col.Classes()
	if err != nil {
		return "", err
	}
	err = col.AddAttribute(htmlutil.AttributeClass, classes...)
	if err != nil {
		return "", err
	}
	return col.Element.String() // from `html.Element`
	/*
		attrs := map[string]string{}
		for k, v := range col.AdditionalProperties {
			attrs[k] = v
		}
		cls, err := col.Class()
		if err != nil {
			return "", err
		} else if len(cls) > 0 {
			attrs["class"] = cls
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
		return fmt.Sprintf(`%s%s</div>`, TagOpening("div", attrs, false), innerHTML), nil
	*/
}

// ColWidthToClass creates a class string where `size` must be in `[xs,sm,md,lg]` and size must be `[1,12]`.
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
