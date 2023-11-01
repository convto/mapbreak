# mapbreak
Detects if there is map reassignment in the range access.

## Install
```bash
$go install github.com/convto/mapbreak/cmd/mapbreak@latest
```

## Usage

The following is an example of code that writes to map with range access

```main.go
package main

func main() {
	m := map[string]string{"a": "item a", "b": "item b"}
	for k, v := range m {
		m[k] = "reassigned: " + v
	}
}
```

Running go vet will result in the following

```bash
$ go vet -vettool=$(which mapbreak) ./...
# github.com/convto/mapbreak/testdata/src/singlefile
./singlefile.go:6:3: detected range access to map and reassigning
```
