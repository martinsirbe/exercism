/*
Package hamming implements logic for calculating the Hamming
difference between two DNA strands.
*/
package hamming

import "github.com/pkg/errors"

// Distance calculates the minimum number of point mutations
// between the two DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.Errorf("dna strands aren't the same length: strand a = %s, strand b = %s", a, b)
	}

	var count int
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
