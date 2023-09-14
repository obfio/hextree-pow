package main

func trippleShift(num, t int) int {
	overflow := int32(num)
	return int(uint32(overflow) >> t)
}
