package yelling

import "strings"

type Yeller interface {
	String() string
	Change(s string)
	Blank()
	Len() int
}

type LoudString struct {
	s string
}

func NewLoudStirng() *LoudString {
	return &LoudString{}
}

/*Pass by reference or Pointer Receiver*/
func (ls *LoudString) Change(s string) {
	/*This method receives a pointer to the loud string passed*/
	ls.s = strings.ToUpper(s)
}

func (ls *LoudString) String() string {
	return ls.s
}

func (ls *LoudString) Len() int {
	return len(ls.s)
}

func (ls *LoudString) Blank() {
	ls.s = ""
}
