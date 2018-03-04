package action

import (
	"io/ioutil"
	"bufio"
	"os"
	"strconv"
	"errors"
)

var listPath string =os.Getenv("GOPATH")+"/src/github.com/ryomak/musicbox/list.txt"

func WriteList(str string) error {
	keyAndValue := []byte(str + "\n")
	return ioutil.WriteFile(listPath, keyAndValue, 0666)
}

func ReadList(musicNumber string)(string,error){
	num ,err := strconv.Atoi(musicNumber)
	if err != nil{
		return "",err
	}
	fp, err := os.Open(listPath)
	if err != nil {
		return "",err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	sum := 1
	for scanner.Scan() {
		if num == sum{
			return scanner.Text(),nil
		}
		sum ++
	}
	return "",errors.New("not match number")
}