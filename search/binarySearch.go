package search

import (
	"github.com/juju/errors"
)

type Number interface {
	int | int32 | int64 | float32 | float64
}

func BinarySearch[T Number](array []T, target T, lowIndex, highIndex int) (int, error) {
	if len(array) == 0 {
		return -1, errors.NotFoundf("length of input array is zero")
	}

	if lowIndex > highIndex || lowIndex < 0 || highIndex > len(array)-1 {
		return -1, errors.NotValidf("invalid input of lowIndex and/or highIndex")
	}

	for lowIndex <= highIndex {
		midIndex := lowIndex + (highIndex-lowIndex)/2
		switch {
		case array[midIndex] == target:
			return midIndex, nil
		case array[midIndex] < target:
			lowIndex = midIndex + 1
		default:
			highIndex = midIndex - 1
		}
	}

	return -1, errors.NotFoundf("%v", target)
}
