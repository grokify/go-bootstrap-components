package bootstrapper

// view-source:https://getbootstrap.com/docs/4.3/examples/jumbotron/

type Navbar struct {
	Title     Link
	MenuLinks []Link
	Search    Search
}

type Link struct {
	Href      string
	InnerHTML string
	Current   bool
	SubLinks  []Link
}

type Search struct {
	Action      string
	Method      string
	Onclick     string
	Placeholder string
	Text        string
}
