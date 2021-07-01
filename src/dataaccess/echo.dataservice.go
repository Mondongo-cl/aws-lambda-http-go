package dataaccess

import (
	"errors"
)

var (
	data       = map[int32]*MessageRow{1: {1, "Hard Coded"}}
	next int32 = 1
)

func GetAll() ([]MessageRow, error) {
	var slice []MessageRow
	for _, v := range data {
		slice = append(slice, *v)
	}
	return slice, nil
}

func Add(message string) (*MessageRow, error) {
	next++
	data[next] = &MessageRow{Id: int32(next), Message: message}
	return data[next], nil

}
func Remove(id int32) error {
	if _, ok := data[id]; ok {
		delete(data, id)
		return nil
	}
	return errors.New("not found")
}

func Find(id int32) (*string, error) {
	r, ok := data[id]
	if !ok {
		return nil, errors.New("not found")
	}
	return &r.Message, nil
}
