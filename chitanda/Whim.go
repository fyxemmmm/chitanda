package chitanda

import (
	"fmt"
	"io/ioutil"
	"os"
)


func LoadConfigFile() []byte {
	dir, _ := os.Getwd()
	file := dir + "/application.yaml"
	fmt.Println(file)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}
	return b
}

