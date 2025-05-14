package slice

//Contains判断切片中是否包含dst
func Contains[T comparable](src []T, dst T) bool {
	return ContainsFunc[T](src, func(src T) bool {
		return src == dst
	})
}

//ContainsFunc判断切片中是否包含dst
//优先使用Contains
func ContainsFunc[T any](src []T, equal func(src T) bool) bool {
	//遍历调用equal函数进行判断
	for _, v := range src {
		if equal(v) {
			return true
		}
	}
	return false
}