package main

// map word count练习 (https://tour.go-zh.org/moretypes/20)
import (
	"strings"
	"golang.org/x/tour/wc"
)
func WordCount(s string) map[string]int{
	strs := strings.Fields(s)
	records := make(map[string]int)
	for _, s := range strs {
		if _, ok := records[s]; !ok {
			records[s] =0
		} 
		records[s]++
	}
	return records	
}

func main() {
	wc.Test(WordCount)
}