package main

import "fmt"

func main() {

	app()
	confDB()
	parser()
	auth()

}

func app() {
	fmt.Println("Main app")
}

func confDB() {
	fmt.Println("Conf Database")
}

func parser() {
	fmt.Println("Menu parser")
}

func auth() {
	fmt.Println("Auth")
}
