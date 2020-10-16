package main

func Exchange(a, b *int) {
	*a, *b = *b, *a
}
