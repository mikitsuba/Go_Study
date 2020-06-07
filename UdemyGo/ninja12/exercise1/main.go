package main

import (
	"fmt"

	"./dog" // packageのimportについて：https://qiita.com/ogady/items/0cedd3599c4dc13e9a95
)

func main() {
	dy := dog.Years(5)
	fmt.Println(dy)
}
