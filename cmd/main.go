package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const server = "http://vietnamtudien.org/hanviet/ajax.php?"

func main() {
	args := os.Args[1:]

	lines := lookup(args)

	for index := range lines {
		line := strings.Split(string(lines[index]), ":")
		if len(line) > 1 {
			fmt.Println("- " + line[2])
		}
	}
}

func lookup(args []string) []string {
	resp, err := http.Get(server + buildQuery(args[0]).Encode())
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

func buildQuery(keyword string) url.Values {
	query := url.Values{}
	query.Set("query", keyword)
	query.Set("methode", "normal")

	return query
}
