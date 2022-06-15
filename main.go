package main

import (
	"fmt"
	"net/http"

	"github.com/Jacksmall/go-basic/morestrings"
	"github.com/google/go-cmp/cmp"
)

type Counter int

func main() {
	s := morestrings.ReverseRunes("Hello World")
	fmt.Println(s)
	fmt.Println(cmp.Diff("Hello Wolrd", "Hello Go"))
	// defer
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	// slice
	slices := make([]int, 10)
	slices = append(slices, 1, 2, 3, 4)
	for _, item := range slices {
		fmt.Println(item)
	}
	// map
	attend := map[string]bool{
		"ck": true,
		"lc": false,
	}

	for key, item := range attend {
		fmt.Printf("%v => %v\n", key, item)
	}
	fmt.Println(attend["ck"])

}

func (ctr *Counter) ServeHttp(w http.ResponseWriter, r *http.Request) {
	*ctr++
	fmt.Println(ctr)
}
