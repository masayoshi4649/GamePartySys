package common

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// ReadSQLFile ... common
func ReadSQLFile(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	b := bytes.NewBuffer(content)
	return b.String(), nil
}
