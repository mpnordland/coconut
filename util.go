package main

import (
    "strconv"
)

type ArticleSlice []*Article

func (as ArticleSlice) Len() int {
	return len(as)
}

func (as ArticleSlice) Less(i, j int) bool {
	return as[i].Time.Before(as[j].Time)
}

func (as ArticleSlice) Swap(i, j int) {
	as[i], as[j] = as[j], as[i]
}


func getPageNum(params map[string]string) (page int) {
    //the default page is the first one
    page = 1

    if p, ok := params["page"]; ok{
        pint64, err := strconv.ParseInt(p, 0, 0)
        if err != nil {
            return 1;
        }
        page = int(pint64)
    }
    return
}

func paginate(articles ArticleSlice, numPerPage, page int) (rArticles ArticleSlice, prev, next int) {
    start := (page-1)*numPerPage
    end := start + numPerPage
    prev, next = -1, page+1
    if len(articles) == 0 {
        rArticles = articles
        next = -1
        return
    }

    if start != 0 {
        prev = page - 1
    }

    if end >= len(articles){
        rArticles = articles[start:]
        next = -1
    } else if start > len(articles) {
        rArticles = make(ArticleSlice, 0)
        next = -1
    } else {
        rArticles = articles[start:end]
    }
    return
}
