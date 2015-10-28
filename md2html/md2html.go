package md2html

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/russross/blackfriday"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

//mardown is getting from github, but add more
type MarkDown struct {
	Title   string
	Author  string
	Content string
	Date    string
}

var DirTimestamp = make(map[string]int64)

func IsDirUpdated(dir string) bool {
	fileinfo, err := os.Stat(dir)
	if err != nil {
		fmt.Println(err)
		return false
	}
	timestamp := fileinfo.ModTime().Unix()
	if val, ok := DirTimestamp[dir]; ok {
		if val >= timestamp {
			return false
		}
	}
	DirTimestamp[dir] = timestamp
	return true
}

//return "" means donnot regenerate
func NewsGenerator(tplFile string, mdDir string) string {
	if !IsDirUpdated(mdDir) {
		fmt.Println("no need to updated.")
		return ""
	}
	t, err := template.ParseFiles(tplFile)
	if err != nil {
		return ""
	}

	type Input struct {
		News []MarkDown
	}
	var input Input

	files_info, _ := ioutil.ReadDir(mdDir)
	for _, file := range files_info {
		if !file.IsDir() {
			mdFile := path.Join(mdDir, file.Name())
			md, err := readMarkDown(mdFile)
			if err == nil {
				input.News = append(input.News, md)
			}
		}
	}

	var buf bytes.Buffer

	err = t.Execute(&buf, input)
	if err != nil {
		fmt.Println("error happen")
		return ""
	}

	return buf.String()
}

func readMarkDown(mdFile string) (md MarkDown, err error) {
	speratorLine := "---"
	f, err := os.OpenFile(mdFile, os.O_RDONLY, 0660)
	if err != nil {
		return md, err
	}
	r := bufio.NewReader(f)
	var mdValue string
	begin := 0
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		if begin == 2 {
			mdValue = mdValue + line
			continue
		}

		if strings.HasPrefix(line, "date: ") {
			md.Date = strings.TrimPrefix(line, "date: ")
		} else if strings.HasPrefix(line, "author: ") {
			md.Author = strings.TrimPrefix(line, "author: ")
		} else if strings.HasPrefix(line, "title: ") {
			md.Title = strings.TrimPrefix(line, "title: ")
		} else if strings.HasPrefix(line, speratorLine) {
			begin++
		}
	}
	md.Content = string(blackfriday.MarkdownCommon([]byte(mdValue)))
	return md, nil
}
