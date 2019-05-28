package libs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Actions interface {
	Add(wrapPoint string, path string) error
	Rm(wrapPoint string) error
	Show(wrapPoint string) (path string, err error)
	List() (records map[string]string, err error)
	Ls(wrapPoint string)
	Clean()
	Cd(wrapPoint string) error
}

type Sekiro struct {
	path    string
	records map[string]string
}

func (s *Sekiro) load() error {
	f, err := os.Open(s.path)
	if err != nil {
		return err
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		data := strings.Split(string(line), ":")
		s.records[data[0]] = data[1]
	}
	return nil
}

func (s *Sekiro) dump() error {
	f, err := os.Create(s.path)
	if err != nil {
		return err
	}
	defer f.Close()
	for wrapPoint, path := range s.records {
		line := fmt.Sprintf("%s:%s\n", wrapPoint, path)
		f.Write([]byte(line))
	}
	return nil
}

func (s *Sekiro) Add(wrapPoint string, path string) error {
	f, err := os.OpenFile(s.path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	line := fmt.Sprintf("%s:%s\n", wrapPoint, path)
	_, err = f.Write([]byte(line))
	return err
}

func (s *Sekiro) Rm(wrapPoint string) error {
	if err := s.load(); err != nil {
		return err
	}
	delete(s.records, wrapPoint)
	if err := s.dump(); err != nil {
		return err
	}

}

func (s *Sekiro) Show(wrapPoint string) (path string, err error) {
	if err := s.load(); err != nil {
		return
	}
	path, ok := s.records[wrapPoint]
	if !ok {
		return path, fmt.Errorf("wrapPoint %s not found!", wrapPoint)
	}
	return path, nil
}

func (s *Sekiro) List() (records map[string]string, err error) {
	if err := s.load(); err != nil {
		return
	}
	return s.records, nil
}

func Ls(wrapPoint string) {
	panic("implement me")
}

func (*Sekiro) Clean() {
	panic("implement me")
}

func (*Sekiro) Cd() error {
}
