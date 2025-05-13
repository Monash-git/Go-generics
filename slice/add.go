package slice

import "go-generics/internal/slice"

//对底层的封装，提供给用户调用
//对任意类型的切片中的index位置进行插入
//特别的，若index == len(src)则
func Add[Src any](src []Src, element Src, index int)([]Src, error){
	res, err := slice.Add[Src](src,element,index)
	return res,err
}