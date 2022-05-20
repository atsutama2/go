package main

import "fmt"

func main() {
	name := "太郎"
	fmt.Printf("name :%v\n", name)

	namePoint := &name
	fmt.Printf("namePoint :%v\n", namePoint)
	fmt.Printf("namePoint :%v\n", *namePoint)

	*namePoint = "木村"
	// *namePointに値を代入することもできる。
	fmt.Printf("*namePointに二郎を代入後の*namePoint :%v\n", *namePoint)
	// 再代入したところで、namePointに格納されている*string型の値(アドレス)自体は、変わらない
	fmt.Printf("*namePointに二郎を代入後の*namePoint :%v\n", namePoint)
	// stringへのポインタである*string型の値(nameに格納されている値)を書き換えたので、nameの値も変更される
	fmt.Printf("*namePointに二郎を代入後の*namePoint :%v\n", name)
}
