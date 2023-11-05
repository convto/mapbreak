package array

func main() {
	arr := []string{"item a", "item b"}
	for i, v := range arr {
		arr[i] = "reassigned: " + v
	}
}
