package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func getImgExt() []string {
	return []string{".jpg", ".png", ".jpeg", ".gif", ".svg"}
}
func getDocExt() []string {
	return []string{".pdf", ".txt", ".csv"}
}

// list files
func listFiles(dir string) []string {
	var ls []string
	files, err := os.ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {

		ls = append(ls, file.Name())
	}
	return ls
}
func imgChecker(ext string) bool {
	imgs := getImgExt()
	for _, i := range imgs {
		if i == ext {
			return true
		}
	}
	return false
}
func docsChecker(ext string) bool {
	docs := getDocExt()
	for _, i := range docs {
		if i == ext {
			return true
		}
	}
	return false
}
func checkExt(ext string) string {
	imgDir := "/home/johannes/github.com/ekn-j/filer/dl_images/"
	docDir := "/home/johannes/github.com/ekn-j/filer/dl_docs/"

	if docsChecker(ext) {
		return imgDir
	} else if imgChecker(ext) {
		return docDir
	}
	return ""
}

func main() {
	downloadsDir := "/home/johannes/github.com/ekn-j/filer/"

	fmt.Println("listing files: ", listFiles(downloadsDir))
	files, err := os.ReadDir(downloadsDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		ext := filepath.Ext(file.Name())
		dir := checkExt(ext)
		fmt.Println(os.Rename(downloadsDir+file.Name(), dir+file.Name()))
	}
	fmt.Println(os.UserHomeDir())
	home, _ := os.UserHomeDir()
	downloads := "/Downloads"
	fmt.Println(home + downloads)
	dlPath := home + downloads
	fmt.Println(listFiles(dlPath))
}
