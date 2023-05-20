package tools

import (
	"baby-chain/tools/data"
	"fmt"
)

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}

func Reverse[T any](arr []T) chan T {
	ret := make(chan T)
	go func() {
		for i := range arr {
			ret <- arr[len(arr)-1-i]
		}
		close(ret)
	}()
	return ret
}

func Contains[T comparable](arr data.Array, element T) bool {
	for _, a := range arr {
		if a == element {
			return true
		}
	}
	return false
}
func Remove[T comparable](arr data.Array, element T) bool {
	var idx int
	found := false
	fmt.Println(arr)
	for index, a := range arr {
		if a == element {
			idx = index
			found = true
			break
		}
	}
	if found {
		arr = append(arr[:idx], arr[idx+1:]...)
		fmt.Println("found : ", arr)
		return true
	}

	return false
}
