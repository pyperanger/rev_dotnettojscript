package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"errors"
)

var (
	fname = flag.String("f", "", "Filename")
	out	= flag.String("o", "", "Output")
)

func LocateByte(f []byte, HEX string) int {
	for i, _ := range f {
		if i+2 == len(f) {
			break
		}
		HEX_ADDR := fmt.Sprintf("%x%x%x", f[i], f[i+1], f[i+2])
		if HEX_ADDR == HEX {
			return i
		}
	}
	return 0
}

func ExtractBin(f []byte) ([]byte, error) {
	INIT_HEX := "4d5a90"
	END_HEX := "275379"
	var INIT_ADDR int
	var END_ADDR int

	INIT_ADDR = LocateByte(f, INIT_HEX)
	END_ADDR = LocateByte(f, END_HEX)

	if INIT_ADDR == 0 || END_ADDR == 0 {
		return []byte{0}, errors.New("Cannot find the andrress")
	}

	return f[INIT_ADDR:END_ADDR], nil
}

func main() {
	flag.Parse()
	if *fname == "" || *out == "" {
		flag.Usage()
		return
	}

	// Open Serialized data
	f, err := ioutil.ReadFile(*fname)
	if err != nil {
		fmt.Println("File not found")
		os.Exit(1)
	}

	fmt.Println("Extracting .. ")

	bin, err := ExtractBin(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	} 

	fmt.Println("Saving PE/DLL in "+*out)
	err = ioutil.WriteFile(*out, bin, 0444)
	if err != nil {
		fmt.Println("Cannot save file, check directory permissions")
		os.Exit(1)
	}

	fmt.Println("Finish.")
	return
}
