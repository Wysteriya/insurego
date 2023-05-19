package tools

import "baby-chain/tools/data"

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
