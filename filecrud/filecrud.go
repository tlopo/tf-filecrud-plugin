package filecrud

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type FileCrud struct {
	Folder string
}

func (f FileCrud) Create(content string) string {
	id := fmt.Sprintf("%d", time.Now().UnixNano())
	data := []byte(content)
	err := ioutil.WriteFile(filepath.FromSlash(f.Folder+"/"+id), data, 0644)
	if err != nil {
		panic(err)
	}
	return id
}

func (f FileCrud) Update(id string, content string) {
	data := []byte(content)
	err := ioutil.WriteFile(filepath.FromSlash(f.Folder+"/"+id), data, 0644)
	if err != nil {
		panic(err)
	}
}

func (f FileCrud) Read(id string) string {
	result, err := ioutil.ReadFile(filepath.FromSlash(f.Folder + "/" + id))
	if err != nil {
		panic(err)
	}
	return string(result)
}

func (f FileCrud) Delete(id string) {
	path := filepath.FromSlash(f.Folder + "/" + id)
	fmt.Println("Removing " + path)
	err := os.Remove(path)
	if err != nil {
		panic(err)
	}
}

func (f FileCrud) Exist(id string) bool {
	path := filepath.FromSlash(f.Folder + "/" + id)
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
