package connection

import (
	"os"
	"path"
)

type JSONString string

func GetConnectionList() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	filename := path.Join(cwd, "mylist.json")
	result, err := os.ReadFile(filename)
	return string(result), err

}
