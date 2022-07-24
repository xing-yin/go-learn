package ch5

import (
	"fmt"
	"golang.org/x/net/html"
	"math"
	"net/http"
	"testing"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func Test1(t *testing.T) {
	fmt.Println(hypot(3, 4))
}

// 不宜过度使用 bare return
// 如果一个函数将所有的返回值都显示的变量名，那么该函数的return语句可以省略操作数。这称之为 「bare return」。
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing html failed.err: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(doc *html.Node) (words, images int) {
	return 1, 1
}
