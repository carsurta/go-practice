package cases

import "fmt"

var fruits []string

func ExampleArraySlice() {
	fruits2 := make([]string, 3)
	fruits2[0] = "사과"
	fruits2[1] = "바나나"
	fruits2[2] = "토마토"
	fmt.Println(fruits2[0], fruits2[1], fruits2[2])
	// Output:
	// 사과 바나나 토마토
}
