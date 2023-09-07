package fnEnv

import (
	"bufio"
	"os"
	"path"
	"strings"
)

func ReadFromFile(fp string) (err error) {
	var wd string
	if wd, err = os.Getwd(); err != nil {
		return
	}

	if !strings.HasPrefix(fp, "/") {
		fp = path.Join(wd, fp)
	}

	var file *os.File
	if file, err = os.Open(fp); err != nil {
		return
	}

	defer file.Close()
	var reader = bufio.NewReader(file)

	for {
		var body []byte
		if body, _, err = reader.ReadLine(); err != nil {
			if err.Error() == "EOF" {
				err = nil
			}
			break
		}

		var strBody = string(body)
		var strLs = strings.Split(strBody, "=")
		if len(strLs) != 2 {
			continue
		}

		if err = os.Setenv(strLs[0], strLs[1]); err != nil {
			return
		}
	}

	return
}
