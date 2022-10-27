//拓展-一个数组实现两个栈
package stack

import (
	"errors"
	"fmt"
)

type StackArray struct {
	Top1 int
	Top2 int
	arr  [10]int
}

func initStackArray() (stack *StackArray) {
	stack = &StackArray{
		Top1: -1,
		Top2: 10,
		arr:  [10]int{},
	}
	return stack
}

func (s *StackArray) push(v int, tag int) error {
	if s.Top1+1 == s.Top2 {
		return errors.New("栈已满，无法插入数据")
	}
	if tag == 1 {
		s.Top1++
		s.arr[s.Top1] = v
	} else if tag == 2 {
		s.Top2--
		s.arr[s.Top2] = v
	} else {
		return errors.New("tag参数不合法，需要传入1或2")
	}

	return nil
}

func (s *StackArray) pop(tag int) (int, error) {
	var v int
	if tag == 1 {
		if s.Top1 == -1 {
			return 0, errors.New("error:栈A已空")
		}
		v = s.arr[s.Top1]
		s.Top1--
	} else if tag == 2 {
		if s.Top2 == 10 {
			return 0, errors.New("error:栈B已空")
		}
		v = s.arr[s.Top2]
		s.Top2++
	} else {
		return 0, errors.New("tag参数不合法，需要输入1或2")
	}
	return v, nil
}

func Init2() {
	stack1 := initStackArray()
	fmt.Println(stack1.arr)

	fmt.Println("---空栈pop报错---")
	_, err := stack1.pop(1)
	if err != nil {
		fmt.Println(err)
	}
	_, err = stack1.pop(2)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("---栈满push报错---")
	for i := 1; i <= 5; i++ {
		err := stack1.push(i, 1)
		if err != nil {
			fmt.Println("栈A：", err)
			break
		}
	}

	for i := 1; i <= 6; i++ {
		err := stack1.push(i, 2)
		if err != nil {
			fmt.Println("栈B：", err)
			break
		}
	}
	fmt.Println(stack1.arr)

	v, _ := stack1.pop(1)
	fmt.Println("---pop出栈---:", v)
	v, _ = stack1.pop(2)
	fmt.Println("---pop出栈---:", v)
	v, _ = stack1.pop(2)
	fmt.Println("---pop出栈---:", v)
}
