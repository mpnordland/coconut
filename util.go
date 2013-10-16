package main

import (
	"fmt"
	"github.com/hoisie/web"
	"html/template"
	"io"
	"io/ioutil"
    "log"
)

func getPage(ctx *web.Context, tmpl *template.Template, name string) {
	cont, err := ioutil.ReadFile(name)
	if err != nil {
		ctx.NotFound("Woops, something's messed up, your " + name + " file is gone!")
		return
	}
	theme(ctx, tmpl, template.HTML(string(cont)))
}

func theme(w io.Writer, temp *template.Template, cont interface{}) {
	err := temp.ExecuteTemplate(w, "theme.html", cont)
	if err != nil {
		fmt.Println(err)
	}
}

func handleFatalError(err error) {
    if err != nil {
        log.Fatalln("Fatal: coconut has exerienced an error:", err)
    }
}
