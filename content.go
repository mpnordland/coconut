package main

import (
        "io/ioutil"
	    "github.com/russross/blackfriday"
        "os"
        "strings"
        "fmt"
        "time"
        "sort"
)

const DateFormat = "Jan _2 2006"


type Article struct {
    Title string
    Author string
    Tags []string
    Time time.Time
    Path string
    Body string
}

func GetArticle(fileName string) (*Article, error) {
    data, err := ioutil.ReadFile("./articles/"+fileName)
    if err != nil {
        return nil, err
    }
    md, cont := GetMetadata(data)

    d, err := time.Parse(DateFormat, md.Date)
    if err != nil {
        fmt.Println("error getting pubdate:", err)
    }

    return &Article{md.Title, md.Author, md.Tags, d, strings.TrimSuffix(fileName, ".md"), string(blackfriday.MarkdownCommon(cont))}, nil
}

func (a *Article) Date() string {
    return a.Time.Format(DateFormat)
}

func (a *Article) HasTag(tag string) bool {
    for _, t := range a.Tags{
        if tag == t {
            return true
        }
    }
    return false
}

func GetArticles() ArticleSlice {
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
        fmt.Println(f)
        if strings.HasSuffix(f, ".md") {
            a, err := GetArticle(f)
            if err != nil {
                continue
            }
            articles = append(articles, a)
        }
    }
    sort.Sort(sort.Reverse(articles))
    return articles
}


type Page struct {
    Title string
    Body string
}

func GetPage(path string) (*Page, error) {
    data, err := ioutil.ReadFile("./static/"+path)
    if err != nil {
        return nil, err
    }
    md, cont := GetMetadata(data)
    return &Page{md.Title, string(blackfriday.MarkdownCommon(cont))}, nil
}
