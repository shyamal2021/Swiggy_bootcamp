package main

import "fmt"

func main(){
	states := make(map[string] string)
	states["KA"] = "BLR"
	states["TN"] = "Chennai"
	states["Maharastra"] = "Mumbai"

	fmt.Println(states)

	states["MP"] = "Bhopal"
	fmt.Println(states)

	city :=  states["TN"]
	fmt.Println("Capital of TN is ", city)

	for k,v := range states{
		fmt.Println(k,v)
	}

	delete(states,"Maharastra")

	for _ ,v := range states{
		fmt.Println(v)
	}

	keys := make([] string,len(states))
	i:=0
	for k := range states{
		keys[i]=k
		i++
	}

	fmt.Println(keys)

	fmt.Println("Keys are : ")
	for i:= range keys{
		fmt.Println(keys[i])
	}


}
