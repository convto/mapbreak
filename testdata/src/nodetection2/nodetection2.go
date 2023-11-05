package array

func main() {
	arr := [2]string{"item a", "item b"}
	for i, v := range arr {
		arr[i] = "reassigned: " + v
	}
}
