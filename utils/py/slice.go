package py

// Remove 传入索引，删除元素
func Remove[Type any](arr []Type, i int) []Type {
	if i < 0 || i > len(arr) {
		return arr
	}
	return append(arr[:i], arr[i+1:]...)
}

// Insert 传入索引，插入元素
func Insert[Type any](arr []Type, i int, new Type) []Type {
	if i < 0 || i > len(arr) {
		return arr
	}
	return append(arr[:i], append([]Type{new}, arr[i:]...)...)
}
