package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

const server = "http://vietnamtudien.org"

var (
	choice int
)

func main() {
	args := os.Args[1:]

	words := getListOfWords(args)

	for index := range words {
		line := strings.Split(string(words[index]), ":")
		if len(line) > 1 {
			fmt.Printf("%d. %s\n", index, line[2])
		}
	}
	fmt.Print("- Tra nghĩa của từ bằng cách nhập vào số thứ tự:")

	fmt.Scanln(&choice)
	word := strings.Split(words[choice], ":")[1]

	definitions := lookup(word)
	list := definitions[1 : len(definitions)-1]

	var re = regexp.MustCompile(`<[^>]+>`)

	fmt.Println(words[choice])

	for index := range list {
		s := re.ReplaceAllString(html.UnescapeString(list[index]), "")
		fmt.Printf("- %s\n", s)
	}
}

func lookup(word string) []string {
	resp, err := http.Get(queryForLookUp(word))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return strings.Split(string(body), "<BR>&FilledSmallSquare;&nbsp;")
}

func getListOfWords(args []string) []string {
	resp, err := http.Get(queryForListOfWords(args[0]))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return parseBody(resp)
}

func parseBody(resp *http.Response) []string {
	body, _ := ioutil.ReadAll(resp.Body)
	lines := strings.Split(string(body), "|")

	return lines
}

func queryForLookUp(keyword string) string {
	query := url.Values{}
	query.Set("unichar", keyword)

	return fmt.Sprintf("%s/hanviet/hv_timchu_ndv.php?%s", server, query.Encode())
}

func queryForListOfWords(keyword string) string {
	query := url.Values{}
	query.Set("query", keyword)
	query.Set("methode", "normal")

	return fmt.Sprintf("%s/hanviet/ajax.php?%s", server, query.Encode())
}
