# go-natsort: Natural String Sorting in Go

![AI Assisted](https://img.shields.io/badge/AI%20Assisted-Yes-blue)
![Human Validated](https://img.shields.io/badge/Human%20Validated-100%25-brightgreen)

No regexp. No copy. No parse number. No dependency.

Sorting rules:

1. Numeric substrings are compared based on their integer values. `a1 < a2 < a10 < a11 < a20`.
2. Leading zeros in numeric substrings are ignored during comparison. `a001 < a2 < a010 < a11 < a00020`.
3. If the integer values are equal, more leading zeros means a smaller value. `a001 < a01 < a1`
4. Digits come before non-digits. `a1 < a_`, `a1 < aa`.
5. Non-numeric substrings are compared lexicographically. `aa < ab`.
6. The shorter string is considered smaller. `a < aa`.

Notices:

* `9.2 < 9.11`
* `a01b2 < a1b1`
* Two strings are considered equal only if they are exactly the same.
* Strings are compared byte by byte (not rune).

## Usage

```shell
go get github.com/ganlvtech/go-natsort
```

```go
package main

import (
	"fmt"
	"slices"
	"sort"

	"github.com/ganlvtech/go-natsort"
)

func main() {
	strs := []string{"a1", "a2", "a10", "a11", "a20", "a001", "a2", "a010", "a11", "a00020", "a001", "a01", "a1", "a1", "a_", "a1", "aa", "aa", "ab", "9.2", "a", "aa", "9.11", "a1b2", "a01b1"}
	fmt.Println(strs)
	strs1 := make([]string, len(strs))
	copy(strs1, strs)
	sort.Slice(strs1, func(i, j int) bool {
		return natsort.Compare(strs1[i], strs1[j]) < 0
	})
	fmt.Println(strs1)
	// go 1.21 slices
	strs2 := make([]string, len(strs))
	copy(strs2, strs)
	slices.SortFunc(strs2, natsort.Compare)
	fmt.Println(strs2)
	// go 1.21 slices with iterator
	strs3 := slices.SortedFunc(slices.Values(strs), natsort.Compare)
	fmt.Println(strs3)
}
```

## Similar Projects

* https://github.com/facette/natsort
* https://github.com/maruel/natural
* https://github.com/skarademir/naturalsort

## LICENSE

[MIT License](https://mit-license.org/)
