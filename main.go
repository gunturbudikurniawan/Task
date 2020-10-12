package main

import "fmt"

func asteric(num int) {
	if num%2 == 0 {
		fmt.Print("Harus Bilangan Ganjil")
	} else {
		for i := 0; i < num; i++ {
			star := ""
			for j := 0; j < num; j++ {
				if j < 1 || j == num-1 || i == num/2 {
					star += "* "
				} else {
					star += "= "
				}
			}
			fmt.Println(star)
		}
	}
	fmt.Println("")
}
func main() {
	asteric(5)
	asteric(6)

}
