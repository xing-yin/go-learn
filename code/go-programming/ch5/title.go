package ch5

import (
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s,not text/html", url, ct)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as html:%v", url, err)
	}
	fmt.Println(doc)
	return nil
}

func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()

	return m[key]
}
