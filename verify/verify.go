package verify

import "fmt"

func isPosInt( num int ) (int, error) {
	if num < 1 {
		err := fmt.Errorf("%v is not a positive integer", num)
		return -1, err
	}
	return num, nil
}