package ygo

import (
	"errors"
)

var CardBag_VOL01 = NewCardVersion(1024)

type CardVersion struct {
	List map[uint]*CardOriginal
}

func NewCardVersion(max int) *CardVersion {
	return &CardVersion{
		List: make(map[uint]*CardOriginal, max),
	}
}

func (cv *CardVersion) Get(id uint) *CardOriginal {
	return cv.List[id]
}

func (cv *CardVersion) Register(co *CardOriginal) error {
	if co == nil {
		return errors.New("RegisterCard: Nil")
	}
	if co.Name == "" {
		return errors.New("RegisterCard: Name is empty")
	}
	if co.Id == 0 || cv.List[co.Id] != nil {
		return errors.New("RegisterCard: Duplicate ID")
	}
	cv.List[co.Id] = co
	return nil
}
