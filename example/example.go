package main

import (
	"fmt"
	"time"

	"github.com/phf/go-openal"
)

func main2() {
	dev := openal.OpenDevice("")
	fmt.Printf("dev: %s\n", dev)

	err := dev.GetError()
	fmt.Printf("err: %s\n", err)

	con := dev.CreateContext()
	fmt.Printf("con: %s\n", con)

	err = dev.GetError()
	fmt.Printf("err: %s\n", err)

	ok := con.MakeContextCurrent()
	fmt.Printf("MakeContextCurrent ok: %s\n", ok)

	con.DestroyContext()
	fmt.Println("Context destroyed!")

	// according to the OpenAL 1.1 spec this should have
	// resulted in an error; hmmm...
	err = dev.GetError()
	fmt.Printf("err: %s\n", err)

	ok = dev.CloseDevice()
	fmt.Printf("CloseDevice ok: %s\n", ok)

	mic := openal.CaptureOpenDevice("", 8000, openal.AlFormatMono16, 16000)
	fmt.Printf("mic: %s\n", mic)

	err = mic.GetError()
	fmt.Printf("err: %s\n", err)

	mic.CaptureStart()
	fmt.Println("capture started!")

	err = mic.GetError()
	fmt.Printf("err: %s\n", err)

	time.Sleep(1 * 1000 * 1000 * 1000)

	smp := mic.GetInteger(openal.AlcCaptureSamples)
	fmt.Printf("smp: %s\n", smp)

	buf := mic.CaptureSamples(smp)
	fmt.Printf("buf: %v\n", buf)

	mic.CaptureStop()
	fmt.Println("capture stopped!")

	err = mic.GetError()
	fmt.Printf("err: %s\n", err)

	ok = mic.CaptureCloseDevice()
	fmt.Printf("ok: %s\n", ok)
}
