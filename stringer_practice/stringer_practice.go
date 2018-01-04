package main

import "fmt"
type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	ip := IPAddr{1, 2, 3, 4}
	fmt.Println(ip)
}