package slice

//UnionSet 求并集，只支持comparable类型
//返回值的顺序是不确定的
func UnionSet[T comparable](src, dst []T) []T {
	srcMap, dstMap := toMap[T](src), toMap[T](dst)
	//将srcMap中的Key全部插入到dstMap中，利用map key的唯一性实现去重
	for key := range srcMap {
		dstMap[key] = struct{}{}
	}
	//把去重后的并集结果重新收集为切片，但是顺序是不确定的
	var res = make([]T, 0, len(dstMap))
	for key := range dstMap {
		res = append(res, key)
	}
	return res
}

//支持任意类型求并集，但是请优先使用UionSet
func UnionSetV[T any](src, dst[]T, equal equalFunc[T]) []T {
	var res = make([]T, 0 , len(src)+len(dst))
	res = append(res, dst...)
	res = append(res, src...)
	return deduplicateFunc[T](res, equal)
}