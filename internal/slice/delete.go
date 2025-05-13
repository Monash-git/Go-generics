package slice

import "go-generics/errs"

func Delete[T any](src []T, index int)([]T,T,error){
	length := len(src)
	if index < 0 || index > length {
		var empty T
		return nil, empty, errs.NewErrIndexOutOfRange(length, index)
	}
	//index后的元素逐一向前覆盖
	res := src[index]
	for i := index ; i < length - 1 ; i++ {
		src[i] = src[i+1]
	}
	//去掉最后一个重复的元素
	src = src[:length-1]
	return src, res, nil
}

