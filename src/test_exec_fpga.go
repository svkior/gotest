package main

import (
    //mmap "github.com/riobard/go-mmap"
	"log"
	"os"
	//"encoding/hex"
	"syscall"
	"fmt"
	"unsafe"
	"time"
)

const (
	MAPSIZE=4096
	MAPMASK=(MAPSIZE - 1)
	target = uint64(0xFFFFEC00)
)


func getRealOffset(offset uint64) (uint64){
	realOffset := uint64( offset & ^uint64(MAPMASK))
	return realOffset
}


func getArrOffset(offset uint64, addr uint64) (uint64){
	offsetAddr := (addr & MAPMASK)
	return offsetAddr
}

func dumpUint32(mapp (*[1024]uint32), targ uint64, offset uint) {
	log.Printf("Four Horstman: map_array[0x%08x] = 0x%08x",getRealOffset(uint64(targ)) + uint64(offset) , mapp[uint(offset) >> 2])
}

func mflush(addr, len uintptr) error {
	_, _, errno := syscall.Syscall(uintptr(syscall.SYS_MSYNC), addr, len, syscall.MS_SYNC)
	if errno != 0 {
		return syscall.Errno(errno)
	}
	return nil
}

func main(){

	//sfcSetupBase := uint64(0xFFFFEC00)

	fpgaBaseMemory := uint64(0x10000000)

	f, err := os.OpenFile("/dev/mem", os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	realOffset := int64( getRealOffset(target) )
	log.Printf("Real offset: 0x%x\n", realOffset)

	mmap, err := syscall.Mmap(int(f.Fd()), realOffset, 4096, syscall.PROT_READ|syscall.PROT_WRITE| syscall.SYS_SYNC, syscall.MAP_SHARED)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	/*
	myMap, err := mmap.Map(f, realOffset, 4096, mmap.PROT_WRITE + mmap.PROT_READ, mmap.MAP_SHARED)
	if err != nil {
		log.Panic(err)
	}*/

	map_array := (*[1024]uint32)(unsafe.Pointer(&mmap[0]))

//	dumpUint32(map_array, target,  uint(0xc00))
//	dumpUint32(map_array, target, uint(0xc04))
//	dumpUint32(map_array, target, uint(0xc08))
//	dumpUint32(map_array, target, uint(0xc0C))

	log.Println("Program fpga mapping...")

	map_array[0xc00 >> 2] = 0x00000000
	map_array[0xc04 >> 2] = 0x0b080b08
	map_array[0xc08 >> 2] = 0x000b000b
	map_array[0xc0C >> 2] = 0x00000000
	log.Println("Done.")

//	dumpUint32(map_array, uint64(target), uint(0xc00))
//	dumpUint32(map_array,uint64(target), uint(0xc04))
//	dumpUint32(map_array,uint64(target), uint(0xc08))
//	dumpUint32(map_array,target, uint(0xc0C))


	err = syscall.Munmap(mmap)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f2, err := os.OpenFile("/dev/mem", os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}


	realOffset2 := int64( getRealOffset(fpgaBaseMemory) )
	log.Printf("Real offset: 0x%x\n", realOffset)

	myMap2, err := syscall.Mmap(int(f2.Fd()), realOffset2, 4096, syscall.PROT_READ|syscall.PROT_WRITE| syscall.SYS_SYNC, syscall.MAP_SHARED)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	red := byte(0)

	valInc := myMap2[2000]

	myMap2[2001] = 1
	mflush(uintptr(fpgaBaseMemory), uintptr(4096))
	myMap2[2000] = valInc
	mflush(uintptr(fpgaBaseMemory), uintptr(4096))
	myMap2[1] = 50	// 1 - Dimmer
	mflush(uintptr(fpgaBaseMemory), uintptr(4096))
	myMap2[2] = 0	// 2 - Flash
	mflush(uintptr(fpgaBaseMemory), uintptr(4096))
	myMap2[3] = 000	// 3 - Mode
	mflush(uintptr(fpgaBaseMemory), uintptr(4096))
	myMap2[4] = 000	// 4 - ?????????? SPEED
	mflush(uintptr(fpgaBaseMemory), uintptr(4096))
	myMap2[5] = 0  // 5 - Red
	mflush(uintptr(fpgaBaseMemory), uintptr(4096))
	myMap2[6] = 0  // 6 - Green
	mflush(uintptr(fpgaBaseMemory), uintptr(4096))
	myMap2[7] = 0  // 7 - Blue
	mflush(uintptr(fpgaBaseMemory), uintptr(4096))
	myMap2[8] = 0  // 8 - White
	mflush(uintptr(fpgaBaseMemory), uintptr(4096))
	myMap2[2000] = (valInc + 1) & 0xff
	mflush(uintptr(fpgaBaseMemory), uintptr(4096))
	myMap2[2001] = 0
	mflush(uintptr(fpgaBaseMemory), uintptr(4096))

	for {

		red += 50



		valInc = myMap2[2000]

		myMap2[2001] = 1
		mflush(uintptr(fpgaBaseMemory), uintptr(4096))
		myMap2[2000] = valInc
		myMap2[5] = red  // 5 - Red
		mflush(uintptr(fpgaBaseMemory), uintptr(4096))
		myMap2[2000] = (valInc + 1) & 0xff
		mflush(uintptr(fpgaBaseMemory), uintptr(4096))
		myMap2[2001] = 0
		mflush(uintptr(fpgaBaseMemory), uintptr(4096))
		//log.Print(myMap2)


		time.Sleep(1 * time.Second)
		valInc = myMap2[2000]

		myMap2[2001] = 1
		mflush(uintptr(fpgaBaseMemory), uintptr(4096))
		myMap2[2000] = valInc
		myMap2[6] = red  // 5 - Red
		mflush(uintptr(fpgaBaseMemory), uintptr(4096))
		myMap2[2000] = (valInc + 1) & 0xff
		mflush(uintptr(fpgaBaseMemory), uintptr(4096))
		myMap2[2001] = 0
		mflush(uintptr(fpgaBaseMemory), uintptr(4096))
		//log.Print(myMap2)

		time.Sleep(1 * time.Second)
		valInc = myMap2[2000]

		myMap2[2001] = 1
		mflush(uintptr(fpgaBaseMemory), uintptr(4096))
		myMap2[2000] = valInc
		myMap2[7] = red  // 5 - Red
		mflush(uintptr(fpgaBaseMemory), uintptr(4096))
		myMap2[2000] = (valInc + 1) & 0xff
		mflush(uintptr(fpgaBaseMemory), uintptr(4096))
		myMap2[2001] = 0
		mflush(uintptr(fpgaBaseMemory), uintptr(4096))
		//log.Print(myMap2)

		time.Sleep(1 * time.Second)
		valInc = myMap2[2000]

		myMap2[2001] = 1
		mflush(uintptr(fpgaBaseMemory), uintptr(4096))
		myMap2[2000] = valInc
		myMap2[8] = red  // 5 - Red
		mflush(uintptr(fpgaBaseMemory), uintptr(4096))
		myMap2[2000] = (valInc + 1) & 0xff
		mflush(uintptr(fpgaBaseMemory), uintptr(4096))
		myMap2[2001] = 0
		mflush(uintptr(fpgaBaseMemory), uintptr(4096))
		//log.Print(myMap2)

		time.Sleep(1 * time.Second)
	}




}
