package main

import (
	"github.com/extism/go-pdk"
)

//export say_hello
func say_hello()  {

	// read function argument from the memory
	input := pdk.Input()

	pdk.Log(pdk.LogInfo, "👋 Hello GoLab 2023 💜")

	firstName, _ := pdk.GetConfig("firstName")
	lastName, _ := pdk.GetConfig("lastName")

	pdk.Log(pdk.LogInfo, "📝" + firstName)
	pdk.Log(pdk.LogInfo, "📝" + lastName)

	output := "👋 (From Tinygo) " + string(input)

	// copy output to host memory
	mem := pdk.AllocateString(output)
	pdk.OutputMemory(mem)

}

func main() {}

