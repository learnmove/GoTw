package stubs

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func InitStub(name string) {

}
func CreateModel(name string) {
	pwd, _ := os.Getwd()
	stub, err := ioutil.ReadFile(pwd + "\\..\\stubs\\template\\Model.txt")
	if err != nil {
		fmt.Print(err)

	}
	content := string(stub)
	content = strings.Replace(content, "$model_name", name, -1)
	va := []byte(content)
	err = ioutil.WriteFile(pwd+"\\..\\app\\models\\"+name+".go", va, os.ModeAppend)
	if err != nil {
		fmt.Print(err)

	}
}
func CreateController(name string) {
	pwd, _ := os.Getwd()
	stub, err := ioutil.ReadFile(pwd + "\\..\\stubs\\template\\Controller.txt")
	if err != nil {
		fmt.Print(err)

	}
	content := string(stub)
	content = strings.Replace(content, "$controller_name", name, -1)
	va := []byte(content)
	err = ioutil.WriteFile(pwd+"\\..\\app\\controllers\\"+name+"Controller"+".go", va, os.ModeAppend)
	if err != nil {
		fmt.Print(err)

	}
}
func CreateRepository(name string) {
	pwd, _ := os.Getwd()
	stub, err := ioutil.ReadFile(pwd + "\\..\\stubs\\template\\Repository.txt")
	if err != nil {
		fmt.Print(err)

	}
	content := string(stub)
	content = strings.Replace(content, "$repository_name", name, -1)
	va := []byte(content)
	err = ioutil.WriteFile(pwd+"\\..\\app\\repositories\\"+name+"Repository"+".go", va, os.ModeAppend)
	if err != nil {
		fmt.Print(err)

	}
}
