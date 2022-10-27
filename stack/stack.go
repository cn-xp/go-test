package stack

import (
	"errors"
	"fmt"
)

type Stack struct {
	MaxSize int
	Top     int
	arr     [10]int
}

//初始化栈
func initStack() (stack *Stack) {
	stack = &Stack{
		MaxSize: 10,
		Top:     -1,
		arr:     [10]int{},
	}
	return stack
}

//入栈
func (s *Stack) push(v int) error {
	//判断栈是否满
	if s.MaxSize-1 == s.Top {
		return errors.New("栈已满，无法插入数据")
	}
	s.Top++          //栈顶+1
	s.arr[s.Top] = v //入栈
	return nil
}

//出栈
func (s *Stack) pop() (int, error) {
	//判断栈是否为空
	if s.Top == -1 {
		return 0, errors.New("error:栈为空")
	}
	v := s.arr[s.Top] //出栈
	s.Top--           //栈顶-1

	return v, nil
}

func Init() {
	//初始化一个栈
	stack1 := initStack()

	//验证空栈报错
	fmt.Println("---空栈pop报错---")
	_, err := stack1.pop()
	if err != nil {
		fmt.Println(err)
	}

	//入栈
	fmt.Println("---栈满push报错---")
	for i := 1; i <= 20; i++ {
		err := stack1.push(i)
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	//出栈
	v, _ := stack1.pop()
	fmt.Println("---pop出栈---:", v)
}
