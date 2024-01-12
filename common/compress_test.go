package common

import (
	"log"
	"testing"
)

func TestCheckApp(t *testing.T) {
	tcases := map[string]bool{
		"zip":   true,
		"unzip": true,
		"abc":   false,
	}
	for app, has := range tcases {
		err := CheckFor(app)
		if has && err == nil {
			continue
		}
		if has && err != nil {
			t.Fail()
		}
	}
}
func TestCreateZipWithPasswordFromFile(t *testing.T) {

	err := CreateZipWithPasswordFromFile("zXqw2", "http.zip", "http.go")
	if err != nil {
		log.Print(err)
	}
}

func TestCreateZipWithPasswordFromFolder(t *testing.T) {
	err := CreateZipWithPasswordFromFolder("UrBox@2023", "common.zip", "folder")
	if err != nil {
		log.Print(err)
	}
}
func TestZipExtract(t *testing.T) {
	err := ZipExtract("123456", "common.zip")
	if err != nil {
		log.Print(err)
	}
}
