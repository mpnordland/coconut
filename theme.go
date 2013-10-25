package main

import (
        "github.com/hoisie/mustache"
)

type ThemeEngine struct {
    layout *mustache.Template
    articleTmpl *mustache.Template
    pageTmpl *mustache.Template
}

func makeThemeFileName(name string) string {
    return "./static/theme/"+name+".tmpl.html"
}

func NewThemeEngine() (*ThemeEngine, error) {
    files := []string{"layout", "article", "page"}
    tmpls := make([]*mustache.Template, 0)
    for _, f := range files {
        tmpl, err := mustache.ParseFile(makeThemeFileName(f))
        if err != nil {
            return nil, err
        }
        tmpls = append(tmpls, tmpl)
    }
    return &ThemeEngine{tmpls[0], tmpls[1], tmpls[2]}, nil
}

func (t *ThemeEngine) ThemeArticle(a *Article) string {
    return t.articleTmpl.Render(a)
}

func (t *ThemeEngine) ThemePage(p *Page) string {
    return t.pageTmpl.Render(p)
}

func (t *ThemeEngine) Theme(content string) string {
    return t.layout.Render(map[string]interface{}{"content":content})
}
