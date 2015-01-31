package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"
)

const DateFormat = "Jan _2 2006 15:04"

type Article struct {
	Title     string
	Author    string
	Tags      []string
	Image     string
	HaveImage bool
	Time      time.Time
	FullView  bool
	Path      string
	Body      string
}

func GetArticle(fileName string) (*Article, error) {
	haveImage := true
	data, err := ioutil.ReadFile("./articles/" + fileName)
	if err != nil {
		return nil, err
	}
	md, cont := GetMetadata(data)

	d, err := time.Parse(DateFormat, md.Date)
	if err != nil {
		fmt.Println("error getting pubdate:", err)
	}
	if md.Image == "" {
		haveImage = false
	}
	return &Article{md.Title, md.Author, md.Tags, md.Image, haveImage, d, true, fileName, string(blackfriday.MarkdownCommon(cont))}, nil
}

func (a *Article) Date() string {
	return a.Time.Format(DateFormat)
}

func (a *Article) HasTag(tag string) bool {
	for _, t := range a.Tags {
		if tag == t {
			return true
		}
	}
	return false
}

func GetArticles(include func(*Article) bool) ArticleSlice {
	articles := make(ArticleSlice, 0)
	file, err := os.Open("./articles")
	if err != nil {
		return nil
	}

	files, err := file.Readdirnames(0)
	if err != nil {
		return nil
	}
	for _, f := range files {
		if strings.HasSuffix(f, ".md") {
			a, err := GetArticle(f)
			if err != nil {
				continue
			}
			if include(a) {
				articles = append(articles, a)
			}
		}
	}
	sort.Sort(sort.Reverse(articles))
	return articles
}

type Page struct {
	Title string
	Body  string
}

func GetPage(path string) (*Page, error) {
	data, err := ioutil.ReadFile("./static/" + path)
	if err != nil {
		return nil, err
	}
	md, cont := GetMetadata(data)
	return &Page{md.Title, string(blackfriday.MarkdownCommon(cont))}, nil
}
