package main

import (
	"fmt"

	"github.com/grokify/go-bootstrap-components"
	"github.com/grokify/go-bootstrap-components/examples"
)

func main() {

	fmt.Println(examples.ExampleWebpageHtmlString())

	if 1 == 0 {
		nav := bootstrap.Navbar{
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
		//navHtml := bootstrap.NavbarHtml(nav)
		//fmt.Println(navHtml)

		pg := bootstrap.Webpage{
			Title:  "FooBar Home",
			Navbar: nav,
		}
		pgHtml := bootstrap.WebpageHtml(pg)
		fmt.Println(pgHtml)
	}
	//fmt.Println("DONE")
}
