package search

import (
	"github.com/juju/errors"
	"github.com/ljg-cqu/gotypes"
)

func BinarySearch[T gotypes.Number](array []T, target T, lowIndex, highIndex int) (int, error) {
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

func BinarySearchRec[T gotypes.Number](array []T, target T, lowIndex, highIndex int) (int, error) {
	if len(array) == 0 {
		return -1, errors.NotFoundf("length of input array is zero")
	}

	if lowIndex > highIndex || lowIndex < 0 || highIndex > len(array)-1 {
		return -1, errors.NotValidf("invalid input of lowIndex and/or highIndex")
	}

	return binarySearchRec(array, target, lowIndex, highIndex)
}

func binarySearchRec[T gotypes.Number](array []T, target T, lowIndex, highIndex int) (int, error) {
	if lowIndex <= highIndex {
		midIndex := lowIndex + (highIndex-lowIndex)/2
		switch {
		case array[midIndex] == target:
			return midIndex, nil
		case array[midIndex] < target:
			lowIndex = midIndex + 1
			return binarySearchRec(array, target, lowIndex, highIndex)
		default:
			highIndex = midIndex - 1
			return binarySearchRec(array, target, lowIndex, highIndex)
		}
	}

	return -1, errors.NotFoundf("%v", target)
}
