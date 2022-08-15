package Filemgmnt

import (
	"log"
	"os"
	"path"
)

func Filelistmusic() {
	dir := "./ALARMS"
	oggfiles := [50]string
	f, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
	}
	files, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if path.Ext(file.Name()) == ".ogg" {

		}
	}
}
