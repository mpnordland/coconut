package main

import (
    "os"
    "fmt"
    "testing"
)


func createTestFile(fileName, fileCont string) {
    file, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed setting up for tests", err)
		os.Exit(1)
	}
	file.WriteString(fileCont)
	file.Close()
}

func TestMain(m *testing.M) {
	//file names for the tests
	articleFile := "./articles/test-article.md"
	pageFile := "./static/pages/test-page.md"

	//create the files
    createTestFile(articleFile, ArticleContent)
    createTestFile(pageFile, PageContent)
    createTestFile(ConfigFile, ConfigContent)

	//run the tests
	exitCode := m.Run()

	//clean up after the tests
	os.Remove(articleFile)
	os.Remove(pageFile)
    os.Remove(ConfigFile)
	//finally, exit
	os.Exit(exitCode)

}
