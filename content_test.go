package main

import (
	"testing"
	"time"
)

const ArticleContent string = `---
title: Test Article
author: Coconut Test
image: 
tags: 
 - article 
 - test 
 - post 
date: Jan  5 2015 13:50
---

This is an article test
'""<>[]{}&'`

const PageContent string = `---
title: Test Page
---

This is a test Page!`

//The article we get is written out in TestMain
//The tests for Article related functions all
//get that article and test based upon it

func getArticle(t *testing.T, articleName string) *Article {
	article, err := GetArticle(articleName)
	if err != nil {
		t.Fatal("Existing article not found", err)
	}
	return article
}

func TestGetArticleWithExistingArticle(t *testing.T) {
	getArticle(t, "test-article.md")
}

func TestGetArticleWithMissingArticle(t *testing.T) {
	article, err := GetArticle("missing-article.md")
	if err == nil && article != nil {
		t.Fail()
	}
}

func TestGetArticleMetadata(t *testing.T) {
	article := getArticle(t, "test-article.md")
	if article.Title != "Test Article" {
		t.Fail()
	}
	if article.Author != "Coconut Test" {
		t.Fail()
	}
	if article.Tags[0] != "article" || article.Tags[1] != "test" || article.Tags[2] != "post" {
		t.Fail()
	}

	if article.Image != "" && article.HaveImage {
		t.Fail()
	}

	time, err := time.Parse(DateFormat, "Jan  5 2015 13:50")
	if err != nil {
		t.Fatal("Parsing Date failed, this is a broken test.", err)
	}

	if !article.Time.Equal(time) {
		t.Fail()
	}
}

func TestGetArticleBody(t *testing.T) {
	article := getArticle(t, "test-article.md")
	if article.Body == "" {
		t.Fail()
	}
}

func TestArticleHasTag(t *testing.T) {
	article := getArticle(t, "test-article.md")
	if !article.HasTag("article") {
		t.Fail()
	}
	if !article.HasTag("test") {
		t.Fail()
	}
	if !article.HasTag("post") {
		t.Fail()
	}
}

func TestArticleDate(t *testing.T) {
	article := getArticle(t, "test-article.md")
	if article.Date() != "Jan  5 2015 13:50" {
		t.Fail()
	}
}

func TestGetArticles(t *testing.T) {
	articles := GetArticles(func(a *Article) bool { return true })
	if len(articles) != 2 {
		t.Fail()
	}
}

//Test pages

func getPage(t *testing.T, pageName string) *Page {
	page, err := GetPage(pageName)
	if err != nil || page == nil {
		t.Fatal("Error getting existing page:", err)
	}
	return page
}

func TestGetPageWithExistingPage(t *testing.T) {
	page, err := GetPage("pages/test-page.md")
	if err != nil || page == nil {
		t.Fail()
	}
}

func TestGetPageWithMissingPage(t *testing.T) {
	page, err := GetPage("pages/missing-page.md")
	if err == nil || page != nil {
		t.Fail()
	}
}

func TestGetPageMetadata(t *testing.T) {
	page := getPage(t, "pages/test-page.md")
	if page.Title != "Test Page" {
		t.Fail()
	}
}

func TestGetPageBody(t *testing.T) {
	page := getPage(t, "pages/test-page.md")
	if page.Body == "" {
		t.Fail()
	}
}

