package find

import (
	"regexp"
)

func FindingsUrlInFile(read_file_bytes []byte) ([][][]byte, error) {
	var regxp = regexp.MustCompile(`([\w+]+\:\/\/)?([\w\d-]+\.)*[\w-]+[\.\:]\w+([\/\?\=\&\#\.]?[\w-]+)*\/?/gm`)

	return regxp.FindAllSubmatch(read_file_bytes, -1), nil
}
