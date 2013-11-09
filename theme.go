package main

import (
	"github.com/hoisie/mustache"
	"github.com/howeyc/fsnotify"
	"log"
	"strings"
)

type ThemeRequest struct {
	template string
	data     interface{}
}

type ThemeEngine struct {
	watcher   *fsnotify.Watcher
	templates map[string]*mustache.Template
	done      chan bool
	request   chan ThemeRequest
	response  chan string
}

func makeThemeFileName(name string) string {
	return "./static/theme/" + name + ".tmpl.html"
}

func NewThemeEngine() (*ThemeEngine, error) {
	w, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	t := ThemeEngine{w, make(map[string]*mustache.Template), make(chan bool, 1), make(chan ThemeRequest, 1), make(chan string, 1)}
	files := []string{"layout", "article", "page"}
	for _, f := range files {
		path := makeThemeFileName(f)
		t.loadTemplate(path)

	}
	//fsnotify can't handle delete/rename
	//so we work around that by watching
	//the directory

	t.watcher.Watch("./static/theme")
	return &t, nil
}

func (t *ThemeEngine) ThemeArticle(a *Article) string {
	return t.sendThemeRequest(makeThemeFileName("article"), a)
}

func (t *ThemeEngine) ThemePage(p *Page) string {
	return t.sendThemeRequest(makeThemeFileName("page"), p)
}

func (t *ThemeEngine) Theme(content string) string {
	return t.sendThemeRequest(makeThemeFileName("layout"), map[string]interface{}{"content": content})
}

func (t *ThemeEngine) loadTemplate(name string) {
	tmpl, err := mustache.ParseFile(name)
	if err != nil {
		return
	}
	t.templates[name] = tmpl
}

func (t *ThemeEngine) sendThemeRequest(tmpl string, data interface{}) string {
	t.request <- ThemeRequest{tmpl, data}
	return <-t.response
}

func (t *ThemeEngine) Run() {
	go func() {
		for {
			select {

			case tr := <-t.request:
				if tmpl, ok := t.templates[tr.template]; ok {
					t.response <- tmpl.Render(tr.data)
				} else {
					t.response <- ""
				}

			case ev := <-t.watcher.Event:
				//filter for just template files
				//see comment on line 39
				log.Print(ev)
				if (ev.IsModify() || ev.IsRename()) && strings.HasSuffix(ev.Name, ".tmpl.html") {
					t.loadTemplate(ev.Name)
				}
			case err := <-t.watcher.Error:
				log.Print("Error in file watcher:", err)
			}
		}
	}()
}
