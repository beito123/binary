# Binary

[![GoDoc](https://godoc.org/github.com/beito123/binary?status.svg)](https://godoc.org/github.com/beito123/binary)

This is a simple binary library written in Go.

## Installation

You can get the package with go get command.

```shell
go get -u github.com/beito123/binary
```

*-u is a option updating the package*

## License
These codes are under the MIT License.

## Examples

### Read

```go
func main() {
	data := []byte{0xff, 0xff, 0xff, 0xff} // = -1 (int32)

	stream := binary.NewStreamBytes(data) // stream with bytes
	
	value, err := stream.Int() // int32
	if err != nil {
		panic(err)
	}

	fmt.Printf("Int32 value: %d", value) // Int value: -1
}
```

### Write

```go
func main() {
	stream := binary.NewStream() // empty bytes

	var value int32 = -1
	
	err := stream.WriteInt(value)
	if err != nil {
		panic(err)
	}

	// Bytes: []byte{0xff, 0xff, 0xff, 0xff}
	fmt.Printf("Bytes: %#v", stream.Bytes())
}
```
