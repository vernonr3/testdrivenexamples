package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Body        string
	Tags        []string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, _ := fs.ReadDir(fileSystem, ".")
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, name string) (Post, error) {
	postFile, err := fileSystem.Open(name)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
	bodyseparator        = "---"
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)
	readMetaLine := func(tagName string) string {
		scanner.Scan()
		mscan := strings.TrimPrefix(scanner.Text(), tagName)
		return mscan
	}
	title := readMetaLine(titleSeparator)

	description := readMetaLine(descriptionSeparator)

	tags := readMetaLine(tagsSeparator)
	mtags := strings.Split(tags, ", ")
	body := readBody(scanner)

	return Post{Title: title, Description: description, Tags: mtags, Body: body}, nil

}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // discard this line
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	body := strings.TrimSuffix(buf.String(), "\n")
	return body
}
