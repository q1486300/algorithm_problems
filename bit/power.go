package bit

func Is2Power(n int) bool {
	return n&(^n+1) == 1
}

func Is4Power(n int) bool {
	//                       010101...010101
	return n&(^n+1) == 1 && (n&0x55555555) != 0
}
