package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/xor-gate/goexif2/exif"
	"github.com/xor-gate/goexif2/mknote"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if strings.ToLower(filepath.Ext(f.Name())) != ".jpg" {
			continue
		}

		fmt.Println(">>>> Reading", f.Name())
		f, err := os.Open(f.Name())
		if err != nil {
			log.Println(err)
			continue
		}

		// Optionally register camera makenote data parsing - currently Nikon and
		// Canon are supported.
		exif.RegisterParsers(mknote.All...)

		x, err := exif.Decode(f)
		if err != nil {
			log.Println(err)
			continue
		}

		// camModel, _ := x.Get(exif.Model) // normally, don't ignore errors!
		// fmt.Println(camModel.StringVal())

		// focal, _ := x.Get(exif.FocalLength)
		// numer, denom, _ := focal.Rat2(0) // retrieve first (only) rat. value
		// fmt.Printf("%v/%v", numer, denom)

		// Two convenience functions exist for date/time taken and GPS coords:
		tm, _ := x.DateTime()
		fmt.Println("Taken: ", tm)

		os.Chtimes(f.Name(), tm, tm)

		fmt.Println("Fixed atime and mtime")

		// lat, long, _ := x.LatLong()
		// fmt.Println("lat, long: ", lat, ", ", long)
	}
}
