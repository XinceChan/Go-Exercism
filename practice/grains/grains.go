package grains

import "fmt"

// Square returns the number of grains on a square on a chess board where the
// first square has 1 and every subsequent square doubles the number.
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, fmt.Errorf("input[%d] invalid. Input should be between 1 and 64 (inclusive)", number)
	}
	// a left shift is equivalent to multiplying by 2.  So we need to left
	// shift by num-1 times to get the number of grains on that square
	return 1 << (uint16(number) - 1), nil
}

func Total() uint64 {
	return (1 << 64) - 1
}
