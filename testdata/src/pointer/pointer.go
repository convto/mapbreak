package main

func main() {
	m := &map[string]string{"a": "item a", "b": "item b"}
	for k, v := range *m {
		(*m)[k] = "reassigned: " + v // want "detected range access to map and reassigning"
	}
}
