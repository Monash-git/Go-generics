package slice

//FilterMap对切片中的值进行过滤，将符合条件的值进行类型转换
//若m的返回值为false则忽略当前的值
func FilterMap[Src any, Dst any](src []Src, m func(idx int, src Src)(Dst, bool)) []Dst {
	res := make([]Dst, 0 , len(src))
	for i, s := range src {
		//Src转换为Dst
		dst, ok := m(i,s)
		if ok {
			res = append(res, dst)
		}
	}
	return res
}

//不进行过滤，对切片中的全部元素进行转换
func Map[Src any, Dst any](src []Src, m func(idx int, src Src)Dst)[]Dst{
	dst := make([]Dst, len(src))
	for i,s := range src {
		dst[i] = m(i,s)
	}
	return dst
}

//将切片转换为Map类型
func ToMap[Ele any, Key comparable](elements []Ele, fn func(element Ele) Key) map[Key]Ele {
	return ToMapV(elements, func(element Ele) (Key, Ele) {
		return fn(element), element
	})
}

func ToMapV[Ele any, Key comparable, Val any](elements []Ele, fn func(element Ele) (Key, Val))(resultMap map[Key]Val) {
	resultMap = make(map[Key]Val, len(elements))
	for _, element := range elements {
		k,v := fn(element)
		resultMap[k] = v
	}
	return
}

//构造map，将切片转换为集合类型
func toMap[T comparable](src []T)map[T]struct{}{
	var dataMap = make(map[T]struct{},len(src))
	for _, v := range src {
		//使用空结构体，减少内存消耗
		dataMap[v] = struct{}{}
	}
	return dataMap
}

//检查并去除切片中的重复元素
func deduplicateFunc[T any](data []T, equal equalFunc[T]) []T {
	var newData = make ([]T, 0 , len(data))
	for k, v := range data {
		//判断当前元素是否在后续的子切片中重复出现
		//若没有再出现则进行保留
		if !ContainsFunc[T](data[k+1:],func(src T) bool {
			return equal(src,v)
		}){
			newData = append(newData, v)
		}
	}
	return newData
}

//利用map类型key的唯一性进行去重
//只能用于comparable类型
func deduplicate[T comparable](data []T) []T {
	//将切片转换为map类型
	dataMap := toMap[T](data)
	var newData = make([]T, 0, len(dataMap))
	//重新收集Key
	for key := range dataMap {
		newData = append(newData, key)
	}
	return newData
}