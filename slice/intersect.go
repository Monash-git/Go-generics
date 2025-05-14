package slice

//取交集，只支持comparable类型
func IntersectSet[T comparable](src []T,dst []T) []T {
	srcMap := toMap[T](src)
	var res = make ([]T, 0, len(src))
	for _, val := range dst {
		if _, exist := srcMap[val]; exist{
			res = append(res, val)
		}
	}
	//去重
	return deduplicate[T](res)
}

//取交集，支持任意类型
//优先使用IntersectSet
func IntersectSetFunc[T any](src []T, dst []T, equal equalFunc[T]) []T {
	var res = make([]T, 0 , len(src))
	for _, v := range dst {
		if ContainsFunc[T](src, func(t T) bool {
			return equal(t, v)
		}){
			res = append(res, v)
		}
	}
	return deduplicateFunc[T](res, equal)
}