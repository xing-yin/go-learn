package bufio

import (
	"errors"
)

/**
 * @Author: Alan Yin
 * @Date: 2022/5/5
 * @Desc:
 */

func main() {
	data, err := test()
	if err != nil {
		switch err.Error() {
		case "bufio: negative count":
			// ... ...
			return
		case "bufio: buffer full":
			// ... ...
			return
		case "bufio: invalid use of UnreadByte":
			// ... ...
			return
		default:
			// ... ...
			return
		}
	}
	println(data)
}

func test() (string, error) {
	return "err msg", errors.New("bufio: invalid use of UnreadByte")
}
