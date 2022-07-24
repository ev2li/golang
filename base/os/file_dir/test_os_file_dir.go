package main

import (
	"fmt"
	"os"
)

func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

func createDir() {
	/* 	err := os.Mkdir("test", os.ModePerm)
	   	if err != nil {
	   		fmt.Printf("err: %v\n", err)
	   	} */
	err := os.MkdirAll("a/b/c", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
func remove() {
	err := os.Remove("d.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func getwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("dir: %v\n", dir)
	err2 := os.Chdir("/var/www/source")
	dir, _ = os.Getwd()
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	}
	fmt.Printf("dir: %v\n", dir)

	s := os.TempDir()
	fmt.Printf("s: %v\n", s)
}

func rename() {
	err := os.Rename("d.txt", "a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func readFile() {
	b, _ := os.ReadFile("1.txt")
	fmt.Printf("b: %v\n", string(b[:]))
}

func writeFile() {
	err := os.WriteFile("test.txt", []byte("hello world"), os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func main() {
	// createFile()
	// createDir()
	// remove()
	// getwd()
	// rename()
	// readFile()
	writeFile()
}
