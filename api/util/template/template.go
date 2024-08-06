package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var path string
var appPath string

func main() {
	var err error
	path, err = os.Getwd()
	appPath = path + "/../../app"
	if err != nil {
		return
	}
	temp := flag.String("f", "", "")
	name := flag.String("n", "", "")
	flag.Parse()
	file, err := os.ReadFile(path + "/examples/" + *temp)
	if err != nil {
		return
	}
	switch *temp {
	case "getXById":
		getXById(string(file), *name)
	}

}

func getXById(content string, name string) {
	str := content
	str = strings.Replace(str, "#Name$", name, -1)
	str = strings.Replace(str, "#name$", strings.ToLower(name), -1)
	str = strings.Replace(str, "#cacheKey$", strings.ToUpper(name), -1)
	serviceFileName := appPath + "/service/" + strings.ToLower(name) + ".go"
	modelFileName := appPath + "/model/" + strings.ToLower(name) + ".go"
	serviceFile, err := os.OpenFile(serviceFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err.Error())
	}
	modelFile, err := os.OpenFile(modelFileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer func(serviceFile *os.File) {
		err = serviceFile.Close()
		if err != nil {
			return
		}
	}(serviceFile)
	defer func(serviceFile *os.File) {
		err = serviceFile.Close()
		if err != nil {
			return
		}
	}(modelFile)
	strArr := strings.Split(str, "====")
	serviceWriter := bufio.NewWriter(serviceFile)
	modelWriter := bufio.NewWriter(modelFile)
	_, _ = serviceWriter.WriteString(strArr[0])
	_, _ = modelWriter.WriteString(strArr[1])
	_ = serviceWriter.Flush()
	_ = modelWriter.Flush()
}
