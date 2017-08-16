package main

import "fmt"

func main() {

fmt.Printf("%d\n",*fun1())


}


func fun1() *int {

return  new(int)
}


