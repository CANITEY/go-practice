package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"errors"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		log.Fatal("A path/directory required")
	}
	path := os.Args[1]
	AbsPath, err := filepath.Abs(path)
	AbsPath = AbsPath + "/"
	if err != nil {
		log.Fatal(err)
	}
	
	fileInfo, err := os.Stat(AbsPath)
	if err != nil {
		log.Fatal(err)
	}

	if !fileInfo.IsDir() {
		log.Fatal("This is not a directory")
	}

	files, err := filepath.Glob(AbsPath + "/*")
	if err != nil {
		log.Fatal(err)
	}

	extensions := map[string]bool{}

	for _, file := range files {
		move(file)
		extension := getExtension(file)
		extensions[extension] = true
	}

	for key := range extensions {
		if os.Mkdir(AbsPath + key, os.ModePerm); errors.Is(err, fs.ErrExist) {
			fmt.Printf("[*] %v Exists\n", key)
		}else if err != nil {
			log.Fatal(err)
		} else {
			if key == "" {
				continue
			}
			fmt.Printf("[*] %v Created\n", key)
		}
	}

	for _, file := range files {
		filestate, err := os.Stat(file)
		if err != nil {
			log.Fatal(err)
		}
		if filestate.IsDir() || getExtension(file) == ""{
			continue
		}
		if err := move(file); err != nil {
			log.Fatal(err)
		}
		ext := getExtension(file)
		_, fi := filepath.Split(file)
		fmt.Printf("[*] %v is moved to %v directory\n", fi, ext)
		time.Sleep(time.Millisecond * 500)
	}

}


func getExtension(path string) (ext string) {
	ext = strings.Join(strings.Split(filepath.Ext(path), ".")[1:], "")
	return
}

func move(p string) (error) {
	fileInfo, err := os.Stat(p)
	if err != nil {
		return err
	}

	if fileInfo.IsDir() {
		return nil
	}

	path, file := filepath.Split(p)
	extDir := getExtension(p)

	if err := os.Rename(p, fmt.Sprintf("%v%v/%v", path, extDir, file)); err != nil {
		return err
	}

	return nil
}
