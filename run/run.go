package main

import (
	"fmt"
	"os"

	"github.com/second-state/WasmEdge-go/wasmedge"
)

func main() {
	wasmedge.SetLogErrorLevel()
	var conf = wasmedge.NewConfigure(wasmedge.WASI)
	var store = wasmedge.NewStore()
	var vm = wasmedge.NewVMWithConfig(conf)
	var wasi = vm.GetImportObject(wasmedge.WASI)
	wasi.InitWasi([]string{}, os.Environ(), []string{".:."})

	fmt.Println(" ### Loading wasm file: ", os.Args[1])
	vm.LoadWasmFile(os.Args[1])
	vm.Validate()
	vm.Instantiate()

	var a, b int32 = 23, 16

	fmt.Println(" ### Running sum: ", a, b)
	var ret, err = vm.ExecuteBindgen("sum", wasmedge.Bindgen_return_i32, a, b)
	if err == nil {
		fmt.Println(" Return value: ", ret.(int32))
	} else {
		fmt.Println(" !!! Error: ", err.Error())
	}

	vm.Release()
	conf.Release()
	store.Release()
}
