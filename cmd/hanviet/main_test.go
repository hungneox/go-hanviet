package main

import (
	"testing"
)

func TestQueryForLookUp(t *testing.T) {
	url := "http://vietnamtudien.org"

	urlForChar := QueryForLookUp("36234", "Char")

	timchuURL := url + "/hanviet/hv_timchu_ndv.php?unichar=36234"

	if urlForChar != timchuURL {
		t.Fatalf("Url for Char was wrong")
	}

	urlForWord := QueryForLookUp("175813", "Word")

	timtukepURL := url + "/hanviet/hv_timtukep_ndv.php?wordid=175813"

	if urlForWord != timtukepURL {
		t.Fatalf("Url for Char was wrong")
	}
}
