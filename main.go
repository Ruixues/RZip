package main

import "fmt"

func main() {
	dealer := GetDealer("glib_2.26.0-1_win64.zip")
	fmt.Println(dealer.GetNowFiles())
}
