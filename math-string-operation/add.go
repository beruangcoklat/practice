package main

import "strconv"

func add(a string, b string) string {
	i := 1
	res := []int{}
	carry := 0

	for len(a)-i >= 0 && len(b)-i >= 0 {
		aa := a[len(a)-i]
		bb := b[len(b)-i]

		x := byteToInt(aa) + byteToInt(bb) + carry
		carry = 0
		if x > 9 {
			x -= 10
			carry += 1
		}
		res = append([]int{x}, res...)
		i++
	}

	var s string
	if len(a)-i >= 0 {
		s = a
	} else if len(b)-i >= 0 {
		s = b
	}

	for len(s)-i >= 0 {
		x := byteToInt(s[len(s)-i]) + carry
		carry = 0
		if x > 9 {
			x -= 10
			carry += 1
		}
		res = append([]int{x}, res...)
		i++
	}

	if carry > 0 {
		res = append([]int{carry}, res...)
	}

	strRes := ""
	for _, a := range res {
		strRes += strconv.Itoa(a)
	}
	return strRes
}
