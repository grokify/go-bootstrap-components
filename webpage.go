package bootstrapper

// view-source:https://getbootstrap.com/docs/4.3/examples/jumbotron/

type Webpage struct {
	Title    string
	Navbar   Navbar
	MainHTML string
}

func (pg *Webpage) NavbarString() string {
	return NavbarHtml(pg.Navbar)
}

func (pg *Webpage) MainString() string {
	return pg.MainHTML
}

func (pg *Webpage) FooterString() string {
	return ""
}
