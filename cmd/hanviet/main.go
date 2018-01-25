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

var re = regexp.MustCompile(`<[^>]+>`)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		log.Fatal("Kết thúc. Vui lòng thử lại với từ cần tra")
	}

	words := getListOfWords(args)

	for index := range words {
		line := strings.Split(string(words[index]), ":")
		if len(line) > 1 {
			fmt.Printf("%d. %s\n", index, line[2])
		}
	}

	for {
		fmt.Printf("-> Tra nghĩa của từ bằng cách nhập vào số thứ tự (0-%d):", len(words)-2)
		fmt.Scanln(&choice)
		if choice >= 0 && choice <= len(words)-2 {
			break
		}
	}

	raw := strings.Split(words[choice], ":")
	wordType := raw[0]
	word := raw[1]

	definitions := lookup(word, wordType)
	list := definitions

	if len(definitions) > 1 {
		list = definitions[1:len(definitions)]
	}

	fmt.Println(strings.Split(words[choice], ":")[2])
	for index := range list {
		fmt.Printf("- %s\n", clean(list[index], re))
	}
}

func clean(input string, re *regexp.Regexp) string {
	return re.ReplaceAllString(html.UnescapeString(strings.Split(input, "<hr>")[0]), "")
}

func lookup(word string, wordType string) []string {
	resp, err := http.Get(queryForLookUp(word, wordType))
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

func queryForLookUp(keyword string, wordType string) string {
	query := url.Values{}

	if wordType == "Word" {
		query.Set("wordid", keyword)
		return fmt.Sprintf("%s/hanviet/hv_timtukep_ndv.php?%s", server, query.Encode())
	}

	query.Set("unichar", keyword)
	return fmt.Sprintf("%s/hanviet/hv_timchu_ndv.php?%s", server, query.Encode())
}

func queryForListOfWords(keyword string) string {
	query := url.Values{}
	query.Set("query", keyword)
	query.Set("methode", "normal")

	return fmt.Sprintf("%s/hanviet/ajax.php?%s", server, query.Encode())
}
