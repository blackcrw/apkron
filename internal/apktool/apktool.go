package apktool

import (
	"fmt"
	"os/exec"
)

func Exists() error {
	var _, err = exec.LookPath("apktool")

	if err != nil {
		return fmt.Errorf("%w - (%s)", err, "Didn't find 'apktool' executable. Install instructions is here: https://ibotpeaches.github.io/Apktool/install/")
	}
	
	return nil
}

func Decompile(apk_path, temp_path string) (string, error) {
	var apktool_exec = exec.Command("apktool", "d", apk_path, "-o", temp_path, "-fq")

	var apktool_output, err = apktool_exec.CombinedOutput()

	if err != nil {
		return "", fmt.Errorf("%w - (%s)", err, "When running apktool")
	}

	return string(apktool_output), nil
}