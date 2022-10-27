package examples

import (
	"github.com/grokify/bootstrapper"
)

func ExampleNavbar() bootstrapper.Navbar {
	return bootstrapper.Navbar{
		Title: bootstrapper.Link{
			Href:      "/",
			InnerHTML: "FooBar"},
		MenuLinks: []bootstrapper.Link{
			{
				Href:      "/solutions",
				InnerHTML: "Solutions"},
			{
				Href:      "/products",
				InnerHTML: "Products"},
			{
				Href:      "/resources",
				InnerHTML: "Resources",
				SubLinks: []bootstrapper.Link{
					{
						Href:      "/resources/foo",
						InnerHTML: "Foo"},
					{
						Href:      "/resources/bar",
						InnerHTML: "Bar"},
				}},
		},
		Search: bootstrapper.Search{
			Action: "/search",
			Method: "GET",
		},
	}
}

func ExampleWebpage() bootstrapper.Webpage {
	return bootstrapper.Webpage{
		Title:  "FooBar Home",
		Navbar: ExampleNavbar(),
	}
}

func ExampleWebpageHTMLString() string {
	pg := ExampleWebpage()
	pgHtmlString := bootstrapper.WebpageHTML(pg)
	return pgHtmlString
}
