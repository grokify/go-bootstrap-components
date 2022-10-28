package bootstrapper

type Element interface {
	HTML() (string, error)
}

type ElementBase struct {
	AdditionalClasses    []string
	AdditionalProperties map[string]string
	InnerHTML            []Element
}
