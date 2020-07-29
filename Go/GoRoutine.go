package main

import "net/http"

func main(){
	urls := [] string {
		"https://golang.org",
		"https://www.tutorialspoint.com/spring_boot/spring_boot_introduction.htm",
		"https://www.geeksforgeeks.org/what-is-the-python-global-interpreter-lock-gil/",
	}
}

func returnType(url string){
	resp, err := http.Get(url)

}
