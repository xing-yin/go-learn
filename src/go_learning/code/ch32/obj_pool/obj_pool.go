package object_pool

import (
	"errors"
	"time"
)

type ReusableObj struct {
	// 可以放置连接、超时等属性
}

type ObjPool struct {
	bufChan chan *ReusableObj //用于缓冲可重用对象
}

// 创建对象
func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOfObj)
	for i := 0; i < numOfObj; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

// 从池中获取: 不建议 *ReusableObj 换成 interface{}  可以存放任意对象，但是建议使用具体类型的对象池
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	// 超时控制：慢请求比快速失败更糟糕
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}

}

// 释放归还
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		// 放不进去，满了
		return errors.New("overflow")
	}
}
