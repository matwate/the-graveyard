package main

func isSquare(n int32) bool {
	var i int32 = 1
	for i*i < n {
		i++
	}
	return i*i == n
}

func square(a int32, b int32) int32 {

	var count int32 = 0
	for i := a; i <= b; i++ {
		if isSquare(i) {
			count++
		}
	}
	return count
}
