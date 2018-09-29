package lib

import (
// "fmt"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
