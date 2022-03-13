package temp_path

import (
	"fmt"
	"io/ioutil"
	"os"
)

func CreateTempDirectory() (string, error) {
	var temp_dir_path, err = ioutil.TempDir("", "")

	if err != nil { return "", fmt.Errorf("%w", err) }
	
	return temp_dir_path, nil
}

func DeleteTempDirectory(temp_dir_path string) error { return os.RemoveAll(temp_dir_path) }