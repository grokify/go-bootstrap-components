package examples

import (
	"github.com/grokify/go-bootstrap-quicktemplate"
)

func ExampleNavbar() bootstrap.Navbar {
	return bootstrap.Navbar{
		Title: bootstrap.Link{
			Href:      "/",
			InnerHtml: "FooBar"},
		MenuLinks: []bootstrap.Link{
			{
				Href:      "/solutions",
				InnerHtml: "Solutions"},
			{
				Href:      "/products",
				InnerHtml: "Products"},
			{
				Href:      "/resources",
				InnerHtml: "Resources",
				SubLinks: []bootstrap.Link{
					{
						Href:      "/resources/foo",
						InnerHtml: "Foo"},
					{
						Href:      "/resources/bar",
						InnerHtml: "Bar"},
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

func ExampleWebpageHtmlString() string {
	pg := ExampleWebpage()
	pgHtmlString := bootstrap.WebpageHtml(pg)
	return pgHtmlString
}
