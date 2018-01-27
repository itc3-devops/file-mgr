package fileMgr

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

var Debug bool

// PrintError : Error handling
func PrintError(err error, label string, message string) {
	// If Debug is set, print messages to screen
	if Debug == true {
		fmt.Println("Output:\n", label, message, err)
	}
	log.WithFields(log.Fields{"vrctl": label}).Error(message, err)
}

// Before : Trim a string based on a delimiter
func Before(value string, a string) string {
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

// Between : Trim a string based on a delimiter
func Between(value string, a string, b string) string {
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

// After : Trim a string based on a delimiter
func After(value string, a string) string {
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}

// CreateFile : Creates a new file, and folder path if not already there
func CreateFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		dir, _ := filepath.Split(path)
		MkDir(dir)
	}
	// create file if not exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		var file, err = os.Create(path)
		
		defer file.Close()
	}
}

// WriteFile : Writes contents of a string to a file
func WriteFile(path string, contents string) {
	err := ioutil.WriteFile(path, []byte(contents), 0666)
	if err != nil {
		log.Fatal(err)
	}
}

// ReadFile : Reads contents of a file into byte array
func ReadFile(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return content
}

// MkDir : Create a new directory
func MkDir(dir string) {
	if _, mkDirErr := os.Stat(dir); os.IsNotExist(mkDirErr) {
		mkDirErr = os.MkdirAll(dir, 0755)
		
	}

}

// Rmdir : Delete directory from local storage
func RmDir(dir string) {
	err := os.RemoveAll(dir)
	

}
