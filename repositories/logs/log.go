package logs

import (
	"os"
	"log"
)

type FileLog interface {
	Count() os.FileInfo
}

type Logs struct {
	Name string
}

func NewCount(path string) FileLog {
	return &Logs{path}
}

func (l *Logs) Count() os.FileInfo {
	f, e := os.Stat(l.Name)
	if e != nil {
		log.Printf("%v", e)
	}
	return f
}