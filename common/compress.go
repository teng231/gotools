package common

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// zip -P password F.zip F
func CreateZipWithPasswordFromFile(password, zipPath, filename string) error {
	if strings.Contains(password, " ") {
		return errors.New("password require not contains space")
	}
	log.Print("Creating", zipPath)
	cmd := fmt.Sprintf(`zip -P %v %v %v`, password, zipPath, filename)
	cmdBash := strings.Fields(cmd)
	log.Println(cmd)
	c := exec.Command(cmdBash[0], cmdBash[1:]...)
	return c.Run()
}

// zip -P password -r F.zip F
func CreateZipWithPasswordFromFolder(password, zipPath, folder string) error {
	if strings.Contains(password, " ") {
		return errors.New("password require not contains space")
	}
	log.Print("Creating", zipPath)
	cmd := fmt.Sprintf(`zip -P %v -r %v %v`, password, zipPath, folder)
	cmdBash := strings.Fields(cmd)
	log.Println(cmd)
	c := exec.Command(cmdBash[0], cmdBash[1:]...)
	return c.Run()
}

// unzip -P your-password zipfile.zip
func ZipExtract(password, zipPath string) error {
	if strings.Contains(password, " ") {
		return errors.New("password require not contains space")
	}
	log.Printf("Unzipping `%s` to directory `%s`\n", zipPath, password)
	cmd := fmt.Sprintf(`unzip -P %v %v`, password, zipPath)
	cmdBash := strings.Fields(cmd)
	log.Println(cmd)
	c := exec.Command(cmdBash[0], cmdBash[1:]...)
	return c.Run()
}

func CheckFor(app string) error {
	_, e := exec.LookPath(app)
	if e != nil {
		log.Println("Make sure " + app + " is install and include your path.")
	}
	return e
}
