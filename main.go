package main

import(
	"fmt"
	"./graphparser"
)

func main(){
	g, err := graphparser.New()
	fmt.Println(err)
	fmt.Printf("%+v\n", g)
}
