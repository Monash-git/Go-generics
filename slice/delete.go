package slice

import (
	"go-generics/internal/slice"
)

//Delete删除切片index处的元素
func Delete[Src any](src []Src, index int)([]Src, error){
	res, _ , err := slice.Delete[Src](src,index)
	return res, err
}

//删除符合一定条件的元素，即m()返回true的元素
//在原数组上进行操作
func DeleteIf[Src any](src []Src, m func(idx int, src Src) bool) []Src {
	//用于记录不进行删除的元素的副本的位置
	empty := 0
	for idx := range src {
		if m(idx,src[idx]) {
			//true不进行复制
			continue
		}
		//不用删除的元素复制到empty处
		src[empty] = src[idx]
		empty++
	}
	//去除剩下的需要删除的元素
	return src[:empty]
}