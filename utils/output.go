package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Output(data []byte, download bool, filename, filepath string) error {
	if !download {
		fmt.Println("++++++++++++++" + filename + "++++++++++++++++++")
		fmt.Println(string(data))
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++")
		return nil
	}
	if filepath == "" {
		err := os.MkdirAll("out", 0777)
		if err != nil {
			return err
		}
		filepath = "out"
	}
	file, err := CreateFile(filepath + "/" + filename + ".json")
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func CreateFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.OpenFile(p, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
}
