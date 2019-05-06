package file

import (
	"bufio"
	"crypto/sha1"
	"errors"
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

func IsFileExist(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		fmt.Println(info)
		return false
	}
	return true
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

func CopyFile(source, dest string) (ok bool, err error) {
	if source == "" || dest == "" {
		err = errors.New("source or dest is null")
		return
	}
	//打开文件资源
	sourceOpen, err := os.Open(source)
	//养成好习惯。操作文件时候记得添加 defer 关闭文件资源代码
	if err != nil {
		return
	}
	defer sourceOpen.Close()
	//只写模式打开文件 如果文件不存在进行创建 并赋予 644的权限。详情查看linux 权限解释
	destOpen, err := os.OpenFile(dest, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	//养成好习惯。操作文件时候记得添加 defer 关闭文件资源代码
	defer destOpen.Close()
	//进行数据拷贝
	_, err = io.Copy(destOpen, sourceOpen)
	if err != nil {
		return
	}
	ok = true
	return
}