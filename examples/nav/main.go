package main

import (
	"fmt"

	"github.com/grokify/bootstrapper"
	"github.com/grokify/bootstrapper/examples"
)

func main() {
	fmt.Println(examples.ExampleWebpageHTMLString())

	if 1 == 0 {
		nav := bootstrapper.Navbar{
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
		//navHtml := bootstrap.NavbarHtml(nav)
		//fmt.Println(navHtml)

		pg := bootstrapper.Webpage{
			Title:  "FooBar Home",
			Navbar: nav,
		}
		pgHtml := bootstrapper.WebpageHTML(pg)
		fmt.Println(pgHtml)
	}
	//fmt.Println("DONE")
}
