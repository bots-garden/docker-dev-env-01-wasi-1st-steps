package main

import (
	"github.com/extism/go-pdk"
)

//export hello
func hello()  {
	// read function argument from the memory
	input := pdk.Input()

	// create the output
	output := "Hello " + string(input)
		
	// copy output to host memory
	mem := pdk.AllocateString(output)
	
	pdk.OutputMemory(mem)
}

func main() {}
