package utils

import (
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"path"
	"runtime"
)

func GetMyPath() string {
	if _, filename, _, ok := runtime.Caller(1); ok {
		return path.Dir(filename)
	}

	return ""
}

func GetTempPath() string {
	var temp_path = "./tmp"

	if _, err := os.Stat(temp_path); err != nil {
		if err := os.Mkdir(temp_path, os.ModeAppend); err != nil {
			log.Println(err)
			return ""
		}
	}

	return temp_path
}

func RunCmd(cmdline string, combine bool) ([]byte, error) {
	var command = exec.Command("/bin/sh", "-c", cmdline)

	if combine {
		return command.CombinedOutput()
	} else {
		return command.Output()
	}

}

const (
	kilo float64 = 1024
)

func FormatBytes(bytes float64, decimals int) string {
	var sizes = []string{"Bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	if decimals < 0 {
		decimals = 0
	}

	var i = math.Floor(math.Log(bytes) / math.Log(kilo))
	var tmpl = fmt.Sprintf("%s0.%df %s", "%", decimals, "%s")
	return fmt.Sprintf(tmpl, bytes/math.Pow(kilo, i), sizes[int(i)])
}
