package config

import (
	"fmt"
	"strconv"
	"strings"
)

// IntSliceValue represents a slice of integers that satisfies the flag.Value interface.
type IntSliceValue []int

// String returns the string representation of the slice.
func (i *IntSliceValue) String() string {
	return fmt.Sprintf("%v", *i)
}

// Set parses the input string and appends the integer value to the slice.
func (i *IntSliceValue) Set(value string) error {
	vals := strings.Split(value, ",")
	for _, v := range vals {
		num, err := strconv.Atoi(v)
		if err != nil {
			return err
		}

		*i = append(*i, num)
	}
	
	return nil
}