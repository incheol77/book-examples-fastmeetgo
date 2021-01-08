package main

/*
extern void CExample();
extern struct swap_return swap(int a, int b);

*/
import "C"

func swap(a, b C.int) (C.int, C.int) {
	return b, a
}

func main() {
	C.CExample()
}
