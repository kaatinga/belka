package belka

import (
	"errors"
)

const digit = 0b110000

type outputValue interface {
	~int32
}

// Sm converts input string which implies to have a value in centimeters to meters in int32 value.
func Sm[V ~string, O outputValue](value V) (O, error) {
	fractionalPart := false
	fractionalPartLength := O(2)
	var output O
	for i, char := range value {
		if i > 255 {
			return 0, errors.New("the input string is too long")
		}

		if fractionalPart == false {
			if char == '.' {
				fractionalPart = true
				continue
			}
		} else {
			fractionalPartLength--
		}

		if char > 47 && char < 58 {
			output = O(char)&^digit + output*10
		} else {
			return 0, errors.New("the input is not a number")
		}

		if fractionalPartLength == 0 {
			break
		}
	}
	return output * multiplier[O](fractionalPartLength), nil
}

// multiplier creates a value using input exponent.
func multiplier[O outputValue](fractionalPartLength O) O {
	output := O(1)
	for ; fractionalPartLength > 0; fractionalPartLength-- {
		output *= 10
	}
	return output
}
