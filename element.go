package bootstrapper

type Element interface {
	HTML() (string, error)
}

type ElementBase struct {
	AdditionalClasses    []string
	Style                string
	InnerHTML            []Element
	AdditionalProperties map[string]string
}
