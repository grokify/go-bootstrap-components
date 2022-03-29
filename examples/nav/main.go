package main

import (
	"fmt"

	"github.com/grokify/go-bootstrap-components"
	"github.com/grokify/go-bootstrap-components/examples"
)

func main() {

	fmt.Println(examples.ExampleWebpageHTMLString())

	if 1 == 0 {
		nav := bootstrap.Navbar{
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
		//navHtml := bootstrap.NavbarHtml(nav)
		//fmt.Println(navHtml)

		pg := bootstrap.Webpage{
			Title:  "FooBar Home",
			Navbar: nav,
		}
		pgHtml := bootstrap.WebpageHTML(pg)
		fmt.Println(pgHtml)
	}
	//fmt.Println("DONE")
}
