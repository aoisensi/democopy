package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

var (
	fSrcDir, fDstDir string
	fHelp            bool
	yesterday        time.Time
)

func init() {
	flag.StringVar(&fSrcDir, "s", "", "Source Demo Directory Name")
	flag.StringVar(&fDstDir, "d", "", "Output Web Server Directory Name")
	yesterday = time.Now().AddDate(0, 0, -1)
}

func main() {
	if err := checkFlags(); err != nil {
		panic(err)
	}
	fis, err := ioutil.ReadDir(fSrcDir)
	if err != nil {
		panic(err)
	}
	var newf os.FileInfo
	//remove newest file
	for _, fi := range fis {
		if newf == nil || newf.ModTime().Before(fi.ModTime()) {
			newf = fi
		}
	}
	log.Printf("Ignore newest file is \"%v\"", newf.Name())
	for _, fi := range fis {
		if fi.IsDir() || fi == newf {
			continue
		}
		fn := fSrcDir + "/" + fi.Name()
		{
			f, err := os.Open(fn)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			t := fi.ModTime().Format("2006.01.02")
			ofn := fmt.Sprintf("%v/%v/%v.zip", fDstDir, t, fi.Name())
			if err := os.MkdirAll(filepath.Dir(ofn), 0777); err != nil {
				log.Fatal(err)
			}
			cmp, err := os.Create(ofn)
			defer cmp.Close()
			if err != nil {
				log.Fatal(err)
				continue
			}
			cmpz := zip.NewWriter(cmp)
			defer cmpz.Close()
			cp, _ := cmpz.Create(fi.Name())
			io.Copy(cp, f)
			log.Printf("Success \"%v\".", ofn)
		}
		os.Remove(fn)
	}
}

func checkFlags() error {
	flag.Parse()
	files := []string{fSrcDir, fDstDir}
	for _, fn := range files {
		fi, err := os.Stat(fn)
		if err != nil {
			return err
		}
		if !fi.IsDir() {
			return fmt.Errorf("\"%v\" is not Directory", fn)
		}
	}
	return nil
}
