package main

import "fmt"

type Locker interface {
	Lock() uint16
	Unlock(id uint16)
}

// 以下两个结构体均实现了Locker接口
type interLocker struct {
	lockID uint16 
}

func (l *interLocker) Lock() uint16 {
	l.lockID++
	fmt.Println("系统已锁定")
	return l.lockID
}

func (l *interLocker) Unlock(id uint16) {
	if id != l.lockID {
		fmt.Println("锁定标识不匹配")
		return
	}
	fmt.Println("系统已解锁")
}


type custLocker struct {
	locked bool
}

func (l *custLocker) Lock() uint16 {
	fmt.Println("线程已锁定")
	return 0
}

func (l *custLocker) Unlock(id uint16) {
	fmt.Println("线程已解锁")
}


// 与Locker接口不兼容
type repreLocker struct { }

func (l *repreLocker) Lock() uint16 {
	fmt.Println("资源被上锁")
	return 1
}

func (l *repreLocker) Unload(id uint16) {
	fmt.Println("资源被卸载")
}

// 未实现Locker接口
type codeLocker struct { }

func (l *codeLocker) Lock() string {
	return "xx-xxx"
}

func (l *codeLocker) Unlock(id string) {
	fmt.Printf("id: %s\n", id)
}

func main() {
	var lk Locker

	// 引用interLocker实例
	lk = &interLocker{ }
	id := lk.Lock()
	lk.Unlock(id)

	// 引用custLocker实例
	lk = &custLocker{ }
	id = lk.Lock()
	lk.Unlock(id)

	// 无法兼容接口
	//lk = &repreLocker{ }
	//id = lk.Lock()

	// 无法兼容接口
	//lk = &codeLocker{ }
	//var sd = lk.Lock()
	//lk.Unlock(sd)
}