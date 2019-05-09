package configuration

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
)

type Parser interface {
	Parse (file string, m map[string]string) (err error)
}

type confParser struct {}

func (cp confParser) Parse (file string, m map[string]string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	pattern := `^(?P<key>\S+)\s+(?P<value>\S+)`
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			continue
		}
		re := regexp.MustCompile(pattern)
		k, v := re.ReplaceAllString(line, "$1"), re.ReplaceAllString(line, "$2")
		m[k] = v
		//fmt.Printf("k:%s,v:%s\n", k, v)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}