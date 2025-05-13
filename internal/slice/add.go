package slice

import "go-generics/errs"

func Add[T any](src []T, element T, index int)([]T, error){
	length := len(src)
	if index < 0 || index > length {
		return nil, errs.NewErrIndexOutOfRange(length, index)
	}
	//先空出一个位置，默认值为零
	var empty T
	src = append(src, empty)
	//将index与之后的元素全部后移一位
	for i := length-1 ; i > index ; i-- {
		if i-1 >= 0 {
			src[i] = src[i-1]
		}
	}
	src[index] = element
	return src, nil
}