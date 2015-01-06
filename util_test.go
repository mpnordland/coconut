package main

import (
	"encoding/base64"
	"math/rand"
	"testing"
	"time"
)

//tests for getPageNum()
func TestGetPageNumWithArg(t *testing.T) {
	params := make(map[string]string)
	params["page"] = "4"
	pageNum := getPageNum(params)
	if pageNum != 4 {
		t.Fail()
	}
}

func TestGetPageNumWithoutArg(t *testing.T) {
	params := make(map[string]string)
	params["some_other_arg"] = "HelloWorld!"
	pageNum := getPageNum(params)
	if pageNum != 1 {
		t.Fail()
	}
}

func TestGetPageNumWithInvalidArg(t *testing.T) {
	params := make(map[string]string)
	params["page"] = "bloop bloop! Error!"
	pageNum := getPageNum(params)
	if pageNum != 1 {
		t.Fail()
	}
}

//tests for paginate()
func makeRandomArticle() *Article {
	randomBytes := make([]byte, 0)
	for i := 0; i <= 10; i++ {
		randomBytes = append(randomBytes, byte(rand.Uint32()))
	}
	randomString := base64.StdEncoding.EncodeToString(randomBytes)
	return &Article{Title: randomString, Author: randomString}
}

func makeRandomArticleSlice(length int) (articles ArticleSlice) {
	for i := 0; i < length; i++ {
		articles = append(articles, makeRandomArticle())
	}
	return
}

func TestPaginateWithOneArticleAndMutiplePerPage(t *testing.T) {
	articles := makeRandomArticleSlice(1)
	pArticles, prev, next := paginate(articles, 5, 1)
	if pArticles[0] != articles[0] {
		t.Fail()
	}
	if prev != -1 && next != -1 {
		t.Fail()
	}
}

func TestPaginateWithMultipleArticlesOnFirstPage(t *testing.T) {
	articles := makeRandomArticleSlice(10)
	pArticles, prev, next := paginate(articles, 5, 1)
	if pArticles[0] != articles[0] || pArticles[4] != articles[4] {
		t.Fail()
	}
	if prev != -1 || next != 2 {
		t.Fail()
	}
}

func TestPaginateWithMultipleArticlesOnLastPage(t *testing.T) {
	articles := makeRandomArticleSlice(10)
	pArticles, prev, next := paginate(articles, 5, 2)
	if pArticles[0] != articles[5] || pArticles[4] != articles[9] {
		t.Log("The articles weren't equal")
		t.Log("pArticles[0] is", pArticles[0], "and articles[4] is", articles[4])
		t.Fail()
	}
	if prev != 1 || next != -1 {
		t.Log("The page Numbers weren't correct")
		t.Fail()
	}
}

func TestPaginateWithMultipleArticlesOnLastPageFewer(t *testing.T) {
	articles := makeRandomArticleSlice(9)
	pArticles, prev, next := paginate(articles, 5, 2)
	if pArticles[0] != articles[5] || pArticles[3] != articles[8] {
		t.Log("The articles weren't equal")
		t.Fail()
	}
	if len(pArticles) != 4 {
		t.Fail()
	}
	if prev != 1 || next != -1 {
		t.Fail()
	}
}

func TestPaginateWithEmptyPage(t *testing.T) {
	articles := makeRandomArticleSlice(5)
	pArticles, prev, next := paginate(articles, 5, 2)
	if len(pArticles) > 0 {
		t.Log("The article list wasn't empty")
		t.Fail()
	}
	if prev != -1 || next != -1 {
		t.Log("prev:", prev, ", next:", next)
		t.Fail()
	}
}

func TestPaginateWithNegativePageNum(t *testing.T) {
	articles := makeRandomArticleSlice(5)
	pArticles, prev, next := paginate(articles, 5, -1)
	if len(pArticles) > 0 {
		t.Log("The article list wasn't empty")
		t.Fail()
	}
	if prev != -1 || next != -1 {
		t.Fail()
	}
}

func TestPaginageWithEmptyList(t *testing.T) {
	articles := makeRandomArticleSlice(0)
	pArticles, prev, next := paginate(articles, 5, 0)
	if len(pArticles) > 0 {
		t.Log("The article list wasn't empty")
		t.Fail()
	}
	if prev != -1 || next != -1 {
		t.Fail()
	}
}

//Test ArticleSlice
func TestArticleSliceLen(t *testing.T) {
	articles := makeRandomArticleSlice(5)
	if articles.Len() != 5 {
		t.Fail()
	}
}

func TestArticleSliceLess(t *testing.T) {
	articles := makeRandomArticleSlice(5)
	articles[1].Time = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	if !articles.Less(0, 1) {
		t.Fail()
	}
}

func TestArticleSliceSwap(t *testing.T) {
	articles := makeRandomArticleSlice(5)
	firstArticle := articles[0]
	articles.Swap(0, 1)
	if firstArticle != articles[1] {
		t.Fail()
	}
}
