package env

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var envContent string

func GetParamValue(value string) (paramVal string, err error) {
	content := readEnvFile()
	paramVal, err = getConfigParamValue(value, content)

	return paramVal, err
}

func readEnvFile() string {
	content := getEnvFileContent()

	return content
}

func getConfigParamValue(param string, data string) (paramVal string, err error) {
	pattern := "(?m)^" + param + ".+"
	re := regexp.MustCompile(pattern)
	match := re.FindString(data)

	if match == "" {
		log.Fatal("Param not found in env")
	}

	value := strings.Split(match, "=")
	val := value[1]

	for i := 2; i < len(value); i++ {
		val += "=" + value[i]
	}

	return strings.Trim(val, "\""), err
}

func getEnvFileContent() string {
	pwd, err := os.Getwd()

	err = filepath.Walk(pwd,
		func(path string, info os.FileInfo, err error) error {

			if info.Mode().String() != "drwxrwxr-x" {
				return nil
			}

			if err != nil {
				return err
			}

			ep, err :=ioutil.ReadFile(path + "/.env")

			if ep != nil{
				envContent = string(ep)
			}

			return nil
		})

	if err != nil {
		log.Println(err)
	}

	return envContent
}