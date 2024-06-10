package main

func main() {
	s := []int{}
	for i := 0; i <= 10; i++ {
		s = append(s, i)
	}

	for _, num := range s {
		if num%2 == 0 {
			println(num, "even")
		} else {
			println(num, "odd")
		}
	}
}
