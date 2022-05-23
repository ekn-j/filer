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
	imgDir := dlPath() + "/dl_images/"
	//"/home/johannes/github.com/ekn-j/filer/dl_images/"
	docDir := dlPath() + "/dl_docs/"
	//"/home/johannes/github.com/ekn-j/filer/dl_docs/"

	if docsChecker(ext) {
		return docDir
	} else if imgChecker(ext) {
		return imgDir
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
func dirChecker() {
	dirs := getDlDirs()
	for _, d := range dirs {
		if err := os.Mkdir(d, os.ModePerm); err != nil {
			fmt.Println("Directory present: ", err)
		}
	}

}

func dlPath() string {
	homedir, _ := os.UserHomeDir()
	dlDir := "/Downloads"
	return homedir + dlDir
}

//func moveFiles(ext string) {
//dlDir := dlPath()
//os.Rename(dlDir+ext,)

//}

func main() {

	dlDir := dlPath()
	err := os.Chdir(dlDir)
	if err != nil {
		log.Fatal(err)
	}
	dlDir, err = os.Getwd()
	println("current working dir is: ", dlDir)

	fmt.Println("Check for existing dirs: ")
	dirChecker()

	fmt.Println("listing files: ", listFiles(dlDir))

	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		ext := filepath.Ext(file.Name())
		dir := checkExt(ext)
		fmt.Println(os.Rename(dlDir+"/"+file.Name(), dir+file.Name()))
	}

	println(os.Getwd())
}
