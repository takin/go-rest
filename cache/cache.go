package cache

import (
	"fmt"
	"io/ioutil"
	"os"

	jsoniter "github.com/json-iterator/go"
)

type Channel struct {
	Data   chan interface{}
	Error  chan error
	IsOpen chan bool
}

type ChannelHandler interface {
	Close()
}

type FileCache struct {
	Path      string
	Overwrite bool
}

type CacheHandler interface {
	Write(interface{}) (bool, error)
	Read() (interface{}, error)
}

func (c *Channel) Close() {
	close(c.Data)
	close(c.Error)
	close(c.IsOpen)
}

func (c *FileCache) Read() (content interface{}, err error) {
	var cacheFile *os.File
	var cacheContent []byte
	cacheFile, err = os.Open(c.Path)
	defer cacheFile.Close()
	if err == nil {
		if cacheContent, err = ioutil.ReadAll(cacheFile); err == nil {
			err = jsoniter.Unmarshal([]byte(cacheContent), &content)
		}
	}
	return
}

func (c *FileCache) Write(content interface{}) (ok bool, err error) {
	var cache []byte
	var file *os.File
	ok = false
	if cache, err = jsoniter.Marshal(content); err != nil {
		return
	}
	file, err = os.Open(c.Path)
	defer file.Close()
	if err == nil && c.Overwrite {
		if _, err = file.Write(cache); err == nil {
			ok = true
		}
		return
	}
	if file, err = os.Create(c.Path); err == nil {
		if _, err = file.Write(cache); err == nil {
			ok = true
		}
		return
	}
	err = fmt.Errorf("File %s already exist in the path", c.Path)
	return
}
