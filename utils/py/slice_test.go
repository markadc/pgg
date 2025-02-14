package py

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("原始值：", arr)

	arr = Remove(arr, 2)
	fmt.Println("删除索引2后：", arr)

	arr = Insert(arr, 0, 200)
	fmt.Println("索引0插入200后：", arr)
}
