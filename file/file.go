package file

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

func GetAllFile(path string, patterns []string) (files []string) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("ReadDir error %s\n", err)
		return
	}

	for _, fi := range entries {
		filename := fi.Name()
		if fi.IsDir() {
			fmt.Printf("dir [%s]\n", path+"/"+filename)
			files = append(files, GetAllFile(path + fi.Name() + "/", patterns)...)
		} else {
			for _, pattern := range patterns {
				if ok,_ := regexp.MatchString(pattern, filename); ok {
					reg := regexp.MustCompile(pattern)
					files = append(files, reg.FindAllString(filename, -1)...)
					//fmt.Printf("dev file %q, %d\n", files, len(files))
				}
			}
		}
	}
	return
}

func GetFileHash(filename string) (hash string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
		return
	}
	defer f.Close()

	br := bufio.NewReader(f)

	h := sha1.New()
	_, err = io.Copy(h, br)

	if err != nil {
		panic(err)
		return
	}
	hash = fmt.Sprintf("%x", h.Sum(nil))
	return
}