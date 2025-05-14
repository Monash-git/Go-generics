package slice

//Find查找符合传入函数的值，若未找到返回false
func Find[T any](src []T, match matchFunc[T])(T,bool){
	for _, val := range src{
		if match(val){
			return val, true
		}	
	}
	var t T
	return t, false
}

//
func FindAll[T any](src []T, match matchFunc[T])[]T{
	//将存储切片容量定为原始切片的八分之一，最多三次扩容即可达到原切片长度级别
	//避免频繁的扩容
	res := make([]T , 0 , len(src)>>3 + 1)
	for _, val := range src {
		if match(val) {
			res = append(res, val)
		}
	}
	return res
}