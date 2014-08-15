package main

import (
	"github.com/robertkrimen/otto"
	"log"
	"runtime/pprof"
	"os"
)

/*
http://blog.golang.org/profiling-go-programs
 */

func main(){
	vm := otto.New()
	vm.Run(`
	abc = 2 + 2
	console.log("The value of abs is " + abc);
	`)
	value, err := vm.Get("abc")
	if err != nil {
		log.Panic(err)
	}
	valueI, _ := value.ToInteger()
	log.Println(valueI)
	f, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.WriteHeapProfile(f)
	f.Close()
	return
}
