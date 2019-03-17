package cache

import (
	"fmt"
	"io/ioutil"
	"os"

	jsoniter "github.com/json-iterator/go"
)

type FileCache struct {
	Path      string
	Overwrite bool
}

type CacheHandler interface {
	Write(interface{}) (bool, error)
	Read() (interface{}, error)
}

func (c *FileCache) Read() (content interface{}, err error) {
	cache, err := os.Open(c.Path)
	defer cache.Close()
	if err == nil {
		if byteContent, err := ioutil.ReadAll(cache); err == nil {
			err = jsoniter.Unmarshal([]byte(byteContent), &content)
		}
	}
	return
}

func (c *FileCache) Write(content interface{}) (ok bool, err error) {
	file, err := os.Open(c.Path)
	defer file.Close()
	if (err != nil) || (err == nil && c.Overwrite) {
		if cache, err := jsoniter.Marshal(content); err != nil {
			file.Write(cache)
			return true, nil
		}
		return false, err
	}
	return false, fmt.Errorf("File Already Exixted")
}
