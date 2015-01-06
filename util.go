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

	if p, ok := params["page"]; ok {
		pint64, err := strconv.ParseInt(p, 0, 0)
		if err != nil {
			return 1
		}
		page = int(pint64)
	}
	return
}

func paginate(articles ArticleSlice, numPerPage, page int) (rArticles ArticleSlice, prev, next int) {
	//pages are one indexed whereas slices
	//are zero indexed.
	start := (page - 1) * numPerPage
	end := start + numPerPage
	prev, next = -1, page+1
	//If we got handed an empty list of articles
	//just return that list and signal there is
	//no next page
	if len(articles) == 0 {
		rArticles = articles
		next = -1
		return
	}

	if start > 0 && start < len(articles) {
		prev = page - 1
	}

	//bounds checks to make sure we
	//are within the slice and don't
	//indicate pages that aren't there
	if start < 0 || start >= len(articles) {
		//This is bad, somehow we got to
		//a page that doesn't exist. This
		//was probably someone messing with
		//the url parameter. Oh Well.
		rArticles = make(ArticleSlice, 0)
		next = -1
	} else if end >= len(articles) {
		//Handles the end of the list
		//we have no more pages after this one!
		rArticles = articles[start:]
		next = -1
	} else {
		//The normal case, everything within
		//the proper limits
		rArticles = articles[start:end]
		//if this isn't the first page we'll return
		//a valid previous page number
	}
	return
}
