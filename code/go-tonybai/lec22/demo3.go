//package bufio
//
//import (
//	"bufio"
//	"errors"
//)
//
///**
// * @Author: Alan Yin
// * @Date: 2022/5/5
// * @Desc: 错误处理策略2
// */
//
//func main() {
//	data, err := test2()
//	if err != nil {
//		switch err.Error() {
//		case bufio.ErrBadReadCount:
//			// ... ...
//			return
//		case bufio.ErrBufferFull:
//			// ... ...
//			return
//		case bufio.ErrInvalidUnreadByte:
//			// ... ...
//			return
//		default:
//			// ... ...
//			return
//		}
//	}
//	println(data)
//}
//
//func test2() (string, error) {
//	return "err msg", errors.New("bufio: invalid use of UnreadByte")
//}
