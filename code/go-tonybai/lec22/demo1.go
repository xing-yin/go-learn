package bufio

import "errors"

/**
 * @Author: Alan Yin
 * @Date: 2022/5/5
 * @Desc:
 */

func main() {

	//err := errors.New("demo error")
	//errWithCtx := fmt.Errorf("index %d is out of bounds", 1)

	err := dosomething()
	if err != nil {
		println(err.Error())
	}

}

func dosomething() error {
	return errors.New("error msg")
}
