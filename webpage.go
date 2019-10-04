package bootstrap

// view-source:https://getbootstrap.com/docs/4.3/examples/jumbotron/

type Webpage struct {
	Title    string
	Navbar   Navbar
	MainHtml string
}

func (pg *Webpage) NavbarString() string {
	return NavbarHtml(pg.Navbar)
}

func (pg *Webpage) MainString() string {
	return pg.MainHtml
}

func (pg *Webpage) FooterString() string {
	return ""
}
