package main

import (
	"fmt"
	"io/ioutil"
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
func getDlDirs() []string {
	return []string{"dl_images", "dl_docs"}
}

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

//Need to add integrations with directory func
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

func listDirs() []string {
	var lsd []string
	home, _ := os.UserHomeDir()
	downloads := "/Downloads"
	dirs, err := ioutil.ReadDir(home + downloads)
	if err != nil {
		log.Fatal(err)
	}

	for _, d := range dirs {
		if d.IsDir() {
			//fmt.Println(d.Name())
			lsd = append(lsd, d.Name())
		}
	}
	return lsd
}
func dirChecker(dir string) {
	dirs := getDlDirs()
	for _, d := range dirs {
		if d == dir {
			err := os.Mkdir(d, 0750)
			if err != nil && !os.IsExist(err) {
				log.Fatal(err)
			}
			fmt.Println("Directory created ", d)
		}
	}
}

func dlPath() string {
	homedir, _ := os.UserHomeDir()
	dlDir := "/Downloads"
	return homedir + dlDir
}

//func moveFiles(ext string) {
//	dlDir := dlPath()
//
//}

func main() {
	dlDir := dlPath()

	fmt.Println("listing files: ", listFiles(dlDir))
	files, err := os.ReadDir(dlDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		ext := filepath.Ext(file.Name())
		dir := checkExt(ext)
		fmt.Println(os.Rename(dlDir+file.Name(), dir+file.Name()))
	}
	fmt.Println(os.UserHomeDir())
	home, _ := os.UserHomeDir()
	downloads := "/Downloads"
	fmt.Println(home + downloads)
	dlPath := home + downloads
	fmt.Println(listFiles(dlPath))
	listDirs()

}
