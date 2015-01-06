package main

import (
	"github.com/hoisie/web"
	"io"
	"log"
	"os"
)

type Controller struct {
	themeEngine     *ThemeEngine
	sessionManager  *SessionManager
	articlesPerPage int
	log             *log.Logger
}

func (c *Controller) Init(conf *Config, s *web.Server) {
	logFile, err := os.Create("coconut.log")
	if err != nil {
		log.Fatalln("Error creating log file:", err)
	}
	c.log = log.New(logFile, "[coconut]", log.LstdFlags|log.Lshortfile)
	s.Get("/", c.Front)
	s.Get("/login", c.Login)
	s.Get("/publish", c.Publish)
	s.Post("/login", c.LoginPost)
	s.Post("/publish", c.PublishPost)
	for url, filename := range conf.Pages {
		s.Get(url, c.makePageFunc(filename))
	}
	s.Get("/tag/(.*)", c.Tag)
	s.Get("/(.*)", c.Article)
}

func (c *Controller) makePageFunc(filename string) func() string {
	return func() string {
		return c.Page(filename)
	}
}

func (c *Controller) Front(ctx *web.Context) string {
	var content string
	articles, prev, next := paginate(GetArticles(func(a *Article) bool { return true }), c.articlesPerPage, getPageNum(ctx.Params))
	for _, a := range articles {
		a.FullView = false
		content += c.themeEngine.ThemeArticle(a)
	}
	if content == "" {
		content = "No articles found"
	}

	return c.themeEngine.Theme(c.themeEngine.ThemeList(content, prev, next))
}

func (c *Controller) Article(name string) string {
	a, err := GetArticle(name + ".md")
	if err != nil {
		return c.themeEngine.Theme("Couldn't find article \"" + name + "\"")
	}
	return c.themeEngine.Theme(c.themeEngine.ThemeArticle(a))
}

func (c *Controller) Page(path string) string {
	p, err := GetPage(path)
	if err != nil {
		return c.themeEngine.Theme("No page at \"" + path + "\"")
	}
	return c.themeEngine.Theme(c.themeEngine.ThemePage(p))
}

func (c *Controller) Tag(ctx *web.Context, tag string) string {
	var content string
	articles, prev, next := paginate(GetArticles(func(a *Article) bool { return a.HasTag(tag) }), c.articlesPerPage, getPageNum(ctx.Params))
	for _, a := range articles {
		a.FullView = false
		content += c.themeEngine.ThemeArticle(a)
	}
	if content == "" {
		content = "No articles found"
	}
	return c.themeEngine.Theme(c.themeEngine.ThemeList(content, prev, next))
}

func (c *Controller) Login(ctx *web.Context) string {
	c.log.Println(ctx.Request.RemoteAddr)
	return c.Page("login.html")
}

func (c *Controller) LoginPost(ctx *web.Context) {
	user := ctx.Params["user"]
	pass := ctx.Params["pass"]
	dest := "/login"
	if c.sessionManager.Login(ctx, user, pass) {
		ctx.WriteString(c.themeEngine.Theme("You are logged in"))
		dest = "/publish"
	} else {
		ctx.WriteString(c.themeEngine.Theme("Incorrect Username/Password"))
	}
	ctx.WriteString("<meta HTTP-EQUIV=\"REFRESH\" content=\"2; url=" + dest + "\">")
}

func (c *Controller) Publish(ctx *web.Context) {
	if !c.sessionManager.LoggedIn(ctx) {
		ctx.Redirect(303, "/login")
		return
	}
	ctx.WriteString(c.Page("publish.html"))
}

func (c *Controller) PublishPost(ctx *web.Context) {
	if !c.sessionManager.LoggedIn(ctx) {
		ctx.Redirect(303, "/login")
		return
	}

	file, head, err := ctx.Request.FormFile("publishFile")
	if err != nil {
		ctx.Abort(405, "error, post without a file")
		return
	}
	saveFile, err := os.Create("articles/" + head.Filename)
	defer saveFile.Close()
	io.Copy(saveFile, file)
	file.Close()
	ctx.Redirect(303, "/publish")
}
