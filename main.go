package main

import(
	"fmt"
	"os"
	"log"
	"./graphparser"
)

func main(){
	file, err := os.Open("./testdata/test01.graphml")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	g, err := graphparser.New(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", g)
}
