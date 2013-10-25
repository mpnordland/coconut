package main

import (
        "github.com/hoisie/web"
)

type Controller struct {
    themeEngine *ThemeEngine
    sessionManager *SessionManager
}

func (c *Controller) Init(conf *Config, s *web.Server) {
    s.Get("/", c.Front)
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

func (c *Controller) Front() string {
    var content string
    articles := GetArticles()
    for _, a := range articles {
        content += c.themeEngine.ThemeArticle(a)
    }
    if content == "" {
        content = "No articles found"
    }
    return c.themeEngine.Theme(content)
}

func (c *Controller) Article(name string) string {
    a, err := GetArticle(name+".md")
    if err != nil {
        return c.themeEngine.Theme("Couldn't find article \""+name+"\"")
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


func (c *Controller) Tag(tag string) string{
     var content string
    articles := GetArticles()
    for _, a := range articles {
        if a.HasTag(tag) {
            content += c.themeEngine.ThemeArticle(a)
        }
    }
    if content == "" {
        content = "No articles found"
    }
    return c.themeEngine.Theme(content)
}
