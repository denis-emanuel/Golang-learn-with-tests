package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Post represents a post on a blog
type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func newPost(postBody io.Reader) (Post, error) {
	//to scan through data line by line
	scanner := bufio.NewScanner(postBody)

	readMetaLine := func(tagName string) string {
		//Scan() to read a line
		scanner.Scan()
		//extract the data using Text()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title:       readMetaLine(titleSeparator),
		Description: readMetaLine(descriptionSeparator),
		Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // ignore a line
	buf := bytes.Buffer{}
	for scanner.Scan() { //scanner.Scan() returns a bool which indicates wheter there's more data to scan 
		//after every Scan() we write the data into the buffer using Fprintln()
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}