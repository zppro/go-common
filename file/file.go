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

type ExtInfo struct {
	FI os.FileInfo
	Path string
	Reg *regexp.Regexp
}
func (fei *ExtInfo) Name () string {
	return fei.FI.Name()
}
func (fei *ExtInfo) FullName () string {
	return fmt.Sprintf("%s/%s", fei.Path, fei.FI.Name())
}

func (fei *ExtInfo) Hash () string {
	return GetFileHash(fei.FullName())
}

func GetAllFileExtInfo(path string, patterns []string) (extInfos []ExtInfo) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Printf("ReadDir error %s\n", err)
		return
	}

	for _, fi := range entries {
		filename := fi.Name()
		if fi.IsDir() {
			fmt.Printf("dir [%s]\n", path+"/"+filename)
			extInfos = append(extInfos, GetAllFileExtInfo(path+fi.Name()+"/", patterns)...)
		} else {
			if len(patterns) > 0 {
				for _, pattern := range patterns {
					if ok, _ := regexp.MatchString(pattern, filename); ok {
						reg := regexp.MustCompile(pattern)
						extInfos = append(extInfos, ExtInfo{FI: fi, Path: path, Reg: reg})
					}
				}
			} else {
				extInfos = append(extInfos, ExtInfo{FI: fi, Path: path})
			}
		}
	}
	return
}

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
					files = append(files, filename)
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