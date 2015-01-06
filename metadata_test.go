package main

import "testing"

//ArticleContent is defined in content_test.go
//This tests the GetMetadata function and Metadata
//struct. If this test passes and the content_test.go
//tests don't the error is probably in content.go or
//there is a broken test

func TestGetMetadataArticle(t *testing.T) {
	m, _ := GetMetadata([]byte(ArticleContent))
	if m.Title != "Test Article" {
		t.Fail()
	}
	if m.Author != "Coconut Test" {
		t.Fail()
	}
	if m.Date != "Jan  5 2015 13:50" {
		t.Fail()
	}
	if m.Image != "" {
		t.Fail()
	}

	if m.Tags[0] != "article" || m.Tags[1] != "test" || m.Tags[2] != "post" {
		t.Fail()
	}
}

func TestGetMetadataPage(t *testing.T) {
	m, _ := GetMetadata([]byte(PageContent))
	if m.Title != "Test Page" {
		t.Fail()
	}

	if m.Author != "" {
		t.Fail()
	}
	if m.Date != "" {
		t.Fail()
	}
	if m.Image != "" {
		t.Fail()
	}

	if len(m.Tags) > 0 {
		t.Fail()
	}
}
