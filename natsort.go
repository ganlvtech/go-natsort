package natsort

// Compare compares two strings a and b using natural order sorting. returns -1 if a < b, 1 if a > b, and 0 if the strings are equal.
func Compare(a string, b string) int {
	i, j := 0, 0
	lenA, lenB := len(a), len(b)

	for i < lenA && j < lenB {
		isNumA := isNumber(a[i])
		isNumB := isNumber(b[j])
		// If both current characters are digits, perform numeric comparison
		if isNumA && isNumB {
			// Count and skip leading zeros
			zeroCountA, zeroCountB := 0, 0
			for i < lenA && a[i] == '0' {
				zeroCountA++
				i++
			}
			for j < lenB && b[j] == '0' {
				zeroCountB++
				j++
			}

			// Record the start position of numbers
			startA, startB := i, j

			// Find the end position of numbers
			for i < lenA && isNumber(a[i]) {
				i++
			}
			for j < lenB && isNumber(b[j]) {
				j++
			}

			// Compare number length first (longer number is greater)
			numLenA := i - startA
			numLenB := j - startB
			if numLenA < numLenB {
				return -1
			}
			if numLenA > numLenB {
				return 1
			}

			// If length is same, compare digit by digit
			for k := 0; k < numLenA; k++ {
				if a[startA+k] < b[startB+k] {
					return -1
				}
				if a[startA+k] > b[startB+k] {
					return 1
				}
			}

			// If both numbers are equal (or both are all zeros), compare zero count
			// More leading zeros means smaller number
			if zeroCountA < zeroCountB {
				return 1
			}
			if zeroCountA > zeroCountB {
				return -1
			}
		} else {
			// One or both are non-numeric characters
			// Digits come before non-digits
			if isNumA {
				return -1
			}
			if isNumB {
				return 1
			}

			// Both are non-numeric, compare directly
			if a[i] < b[j] {
				return -1
			}
			if a[i] > b[j] {
				return 1
			}
			i++
			j++
		}
	}

	// The shorter string is considered smaller
	if i < lenA {
		return 1
	}
	if j < lenB {
		return -1
	}
	return 0
}

func isNumber(input uint8) bool {
	return input >= '0' && input <= '9'
}
