package main

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
