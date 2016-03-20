package main

import (
  "html/template"
  "path/filepath"
  "io/ioutil"
  "strings"
  "time"
  "log"
  "os"

  blfr "github.com/russross/blackfriday"
)

const (
  flags = blfr.HTML_USE_XHTML             |
          blfr.HTML_USE_SMARTYPANTS       |
          blfr.HTML_SMARTYPANTS_FRACTIONS |
          blfr.HTML_SMARTYPANTS_DASHES    |
          blfr.HTML_SMARTYPANTS_LATEX_DASHES

  extensions =  blfr.EXTENSION_NO_INTRA_EMPHASIS |
		            blfr.EXTENSION_TABLES            |
		            blfr.EXTENSION_FENCED_CODE       |
		            blfr.EXTENSION_AUTOLINK          |
		            blfr.EXTENSION_STRIKETHROUGH     |
		            blfr.EXTENSION_SPACE_HEADERS     |
		            blfr.EXTENSION_HEADER_IDS        |
		            blfr.EXTENSION_AUTO_HEADER_IDS   |
		            blfr.EXTENSION_DEFINITION_LISTS  |
                blfr.EXTENSION_BACKSLASH_LINE_BREAK
)

type page struct { Name, File string }

var funcs = template.FuncMap{
  "readdir":  ioutil.ReadDir,
  "readfile": ioutil.ReadFile,
  "noescape": func(s string)template.HTML { return template.HTML(s) },
  "slug": func(s string)string {
    return strings.ToLower(strings.Replace(s, " ", "-", -1))
  },
  "markdown": func(b []byte)string {
    r := blfr.HtmlRenderer(flags, "", "")
    return string(blfr.MarkdownOptions(b, r, blfr.Options{ Extensions: extensions }))
  },
}

var templ = template.Must(template.New("index").Funcs(funcs).ParseGlob("../views/*.html"))

var pages = []page{
    {"Teams", "team.md"},
    //{"Project", "project.md"},
}

func main() {
    start := time.Now()
    data := struct{
      Docs string
      Pages []page
    }{
      Docs:  filepath.Join("../pages"),
      Pages: pages,
    }

    log.Printf("Building %s", data.Docs)

    err := templ.Execute(os.Stdout, data)

    if err != nil {
      log.Fatalf("Error: %s", err)
    }

    log.Printf("Build completed in %s", time.Since(start))
}
