package config

import (
	"os"
	"bufio"
	"io"
	"strings"
	"strconv"
)

type Property struct {
	//键值对
	Store map[string]string
	//配置文件路径
	Path  string
}

func NewProperty(path string) *Property {
	p:=&Property{}
	p.Load(path)
	return p
}

/**
 *初始化配置文件
 */
func (p *Property) Load(path string) (*Property, error) {
	p.Store = make(map[string]string)
	p.Path = path
	f, err := os.Open(p.Path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	p.readFile(f)

	return p, nil
}

func (p *Property) readFile(f *os.File) error {
	r := bufio.NewReader(f)

	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		s := strings.TrimSpace(string(b))

		if strings.Index(s, "#") > 0 {
			continue
		}

		index := strings.Index(s, "=")
		if index < 0 {
			continue
		}

		key := strings.TrimSpace(s[:index])
		if len(key) == 0 {
			continue
		}
		value := strings.TrimSpace(s[index + 1:])

		p.Store[key] = value

	}
	return nil

}

func (p *Property) reload() {
	p.Load(p.Path)
}

func (p *Property) Get(key string, def ...interface{}) interface{} {
	value := p.Store[key]
	if len(value) == 0 {
		if (len(def) == 0) {
			return nil
		}
		return def[0]
	}
	return value
}

func (p *Property) GetString(key string, def ...interface{}) string {
	v := p.Get(key,def...)
	b:=v.(string)
	return b
}

func (p *Property) GetBool(key string, def ...interface{}) bool {
	v := p.Get(key,def...)

	b, _ := strconv.ParseBool(v.(string))
	return b
}

func (p *Property) GetInt(key string, def ...interface{}) int {
	v := p.Get(key,def...)
	i, _ := strconv.Atoi(v.(string))
	return i
}