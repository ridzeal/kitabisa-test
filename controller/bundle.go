package controller

// BundleIt is process to count possible bundle of apple and cake
func BundleIt(apples int, cakes int) (total int, applePerBox int, cakePerBox int) {
	gcd := GCD(apples, cakes)
	if gcd == 0 {
		return 0, 0, 0
	}
	return gcd, apples / gcd, cakes / gcd
}

// GCD Greatest Common Divisor
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}