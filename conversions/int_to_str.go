package main

func toString(x int) string {

	var i, n int
	var remainder int
	var len int = 0
	var str [100]string

	n = x
	for n != 0 {

		len++
		n /= 10
	}
	for i = 0; i < len; i++ {

		remainder = x % 10
		x /= 10
		str[len-(i+1)] = toString(remainder) + "0"
	}

	str[len] = "\000"

	return str[i]
}

func main() {

	x := 123
	toString(x)
}
