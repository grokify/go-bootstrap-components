package bootstrapper

import (
	"errors"

	"github.com/grokify/mogo/type/stringsutil"
)

const (
	ClassActive = "active"

	ClassButton         = "btn"
	ClassButtonGroup    = "btn-group"
	ClassButtonSuccess  = "btn-success"
	ClassDropdown       = "dropdown"
	ClassDropdownMenu   = "dropdown-menu"
	ClassDropdownToggle = "dropdown-toggle"

	ClassListGroup         = "list-group"
	ClassListGroupItem     = "list-group-item"
	ClassListGroupItemText = "list-group-item-text"

	ClassPanel        = "panel"
	ClassPanelHeading = "panel-heading"
	ClassPanelTitle   = "panel-title"
	ClassPanelBody    = "panel-body"

	ColorBlack  = "inverse"
	ColorBlue   = "primary"
	ColorCyan   = "info"
	ColorGreen  = "success"
	ColorOrange = "warning"
	ColorRed    = "danger"
)

var ErrColorUnknown = errors.New("color unknown")

func Colors() []string {
	return []string{ColorBlack, ColorBlue, ColorCyan, ColorGreen, ColorOrange, ColorRed}
}

func IsColor(c string) bool {
	colors := Colors()
	return stringsutil.SliceIndex(colors, c) != -1
}
