//栈的链式存储实现
package stack

import (
	"errors"
	"fmt"
)

type LineStack struct {
	data int
	next *LineStack
}

func initLinkStack() (linkStack *LineStack) {
	linkStack = &LineStack{
		data: -1,
		next: nil,
	}
	return linkStack
}

func (s *LineStack) push(v int) {
	//链栈不需要考虑栈满问题

	pushNode := &LineStack{
		data: v,
		next: nil,
	}
	pushNode.next = s.next
	s.next = pushNode
}

func (s *LineStack) pop() (int, error) {
	var v int
	if s.next == nil {
		return 0, errors.New("error:栈为空")
	}
	tmpTop := s.next
	v = tmpTop.data
	s.next = tmpTop.next //头节点指向原栈顶节点的下一个节点
	tmpTop.next = nil    //原栈顶节点指向nil
	return v, nil
}

func Init3() {
	linkStack := initLinkStack()

	fmt.Println("---入栈---")
	for i := 1; i <= 5; i++ {
		linkStack.push(i)
		fmt.Printf("第%v次入栈，值为：%v\n", i, i)
	}

	fmt.Println("---出栈---")
	for i := 1; i <= 6; i++ {
		v, err := linkStack.pop()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("第%v次出栈，值为：%v\n", i, v)
	}

}
