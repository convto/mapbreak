package main

func main() {
	ma := map[string]string{"a": "item a", "b": "item b"}
	mb := map[string]string{}
	for k, v := range ma {
		mb[k] = "reassigned: " + v
	}
}
