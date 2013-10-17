package main

import (
	"fmt"
	"github.com/hoisie/web"
	"github.com/russross/blackfriday"
	"html/template"
	"io"
	"io/ioutil"
	"os"
)

type PageHandler struct {
	sm   *SessionManager
	tmpl *template.Template
}

func NewPageHandler(sm *SessionManager) *PageHandler {
    return &PageHandler{sm, template.Must(template.ParseFiles("./static/theme.html"))}
}

func (ph *PageHandler) setupHandlers(s *web.Server, c *Config) {
    //handle any pages setup in the config file
    for url, fileName := range c.Pages{
        s.Get(url, makeHandler(ph, "./static/"+fileName))
    }

    //add handlers for default pages
	s.Get("/", ph.front)
	s.Get("/publish", ph.publish)
	s.Post("/publish", ph.publishPost)
	s.Get("/login", ph.login)
	s.Post("/login", ph.loginPost)
	s.Get("/(.*)", ph.article)
}

func (ph *PageHandler) front(ctx *web.Context) {
	getPage(ctx, ph.tmpl, "static/front.html")
}

func (ph *PageHandler) article(ctx *web.Context, val string) {
	cont, err := ioutil.ReadFile("./articles/" + val + ".md")
	if err != nil {
		ctx.NotFound("Sorry, the article " + val + " just isn't here!")
		return
	}
	theme(ctx, ph.tmpl, template.HTML(string(blackfriday.MarkdownCommon(cont))))
}

func (ph *PageHandler) login(ctx *web.Context) {
	getPage(ctx, ph.tmpl, "static/login.html")
}

func (ph *PageHandler) loginPost(ctx *web.Context) {
	user := ctx.Params["user"]
	pass := ctx.Params["pass"]
	if ph.sm.CreateSession(ctx, user, pass) {
		theme(ctx, ph.tmpl, template.HTML("Success!\n <META HTTP-EQUIV=\"refresh\" CONTENT=\"5;URL=/publish\">"))
	} else {
		ctx.Redirect(303, "/login")
	}
}

func (ph *PageHandler) publish(ctx *web.Context) {
	if ph.sm.LoggedIn(ctx) {
		getPage(ctx, ph.tmpl, "static/publish.html")
	} else {
		ctx.Redirect(303, "/login")
	}
}

func (ph *PageHandler) publishPost(ctx *web.Context) {
	if !ph.sm.LoggedIn(ctx) {
		ctx.Redirect(303, "/login")
		return
	}
	file, head, err := ctx.Request.FormFile("publishFile")
	if err != nil {
		fmt.Println(err)
		ctx.Abort(405, "error, post without a file")
		return
	}
	saveFile, err := os.Create("articles/" + head.Filename)
	defer saveFile.Close()
	io.Copy(saveFile, file)
	file.Close()
	ctx.Redirect(303, "/publish")
}
