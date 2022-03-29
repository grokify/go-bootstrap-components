// Code generated by qtc from "webpage.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line webpage.qtpl:1
package bootstrap

//line webpage.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line webpage.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line webpage.qtpl:1
func StreamWebpageHTML(qw422016 *qt422016.Writer, pageData Webpage) {
//line webpage.qtpl:1
	qw422016.N().S(`!<DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>`)
//line webpage.qtpl:5
	qw422016.E().S(pageData.Title)
//line webpage.qtpl:5
	qw422016.N().S(`</title>
    <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" rel="stylesheet" crossorigin="anonymous">
  </head>
  <body>

  `)
//line webpage.qtpl:10
	qw422016.N().S(pageData.NavbarString())
//line webpage.qtpl:10
	qw422016.N().S(`

  `)
//line webpage.qtpl:12
	qw422016.N().S(pageData.MainString())
//line webpage.qtpl:12
	qw422016.N().S(`

  </body>
</html>  
`)
//line webpage.qtpl:16
}

//line webpage.qtpl:16
func WriteWebpageHTML(qq422016 qtio422016.Writer, pageData Webpage) {
//line webpage.qtpl:16
	qw422016 := qt422016.AcquireWriter(qq422016)
//line webpage.qtpl:16
	StreamWebpageHTML(qw422016, pageData)
//line webpage.qtpl:16
	qt422016.ReleaseWriter(qw422016)
//line webpage.qtpl:16
}

//line webpage.qtpl:16
func WebpageHTML(pageData Webpage) string {
//line webpage.qtpl:16
	qb422016 := qt422016.AcquireByteBuffer()
//line webpage.qtpl:16
	WriteWebpageHTML(qb422016, pageData)
//line webpage.qtpl:16
	qs422016 := string(qb422016.B)
//line webpage.qtpl:16
	qt422016.ReleaseByteBuffer(qb422016)
//line webpage.qtpl:16
	return qs422016
//line webpage.qtpl:16
}
