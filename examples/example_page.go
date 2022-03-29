package examples

import (
	"github.com/grokify/go-bootstrap-components"
)

func ExampleNavbar() bootstrap.Navbar {
	return bootstrap.Navbar{
		Title: bootstrap.Link{
			Href:      "/",
			InnerHTML: "FooBar"},
		MenuLinks: []bootstrap.Link{
			{
				Href:      "/solutions",
				InnerHTML: "Solutions"},
			{
				Href:      "/products",
				InnerHTML: "Products"},
			{
				Href:      "/resources",
				InnerHTML: "Resources",
				SubLinks: []bootstrap.Link{
					{
						Href:      "/resources/foo",
						InnerHTML: "Foo"},
					{
						Href:      "/resources/bar",
						InnerHTML: "Bar"},
				}},
		},
		Search: bootstrap.Search{
			Action: "/search",
			Method: "GET",
		},
	}
}

func ExampleWebpage() bootstrap.Webpage {
	return bootstrap.Webpage{
		Title:  "FooBar Home",
		Navbar: ExampleNavbar(),
	}
}

func ExampleWebpageHTMLString() string {
	pg := ExampleWebpage()
	pgHtmlString := bootstrap.WebpageHTML(pg)
	return pgHtmlString
}
