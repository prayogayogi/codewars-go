package data

func Grow(arr []int)int{
	number := 1
	for i := 0; i < len(arr); i++ {
		number = number * arr[i]
	}
	return number
}