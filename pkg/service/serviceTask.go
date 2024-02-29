package service

import (
	"errors"
	"slices"
	"sort"
)

func ArraySorted(array []int) bool {
	for i := 0; i < len(array)-1; i++ {
		if array[i] < array[i+1] {
			return true
		}
	}
	return false
}

func (s *Service) AnswerTask1(data []int) (int, error) {
	err := errors.New("expect more one number")
	if len(data) == 1 {
		return 0, err
	}
	err = errors.New("empty array")
	if len(data) == 0 {
		return 0, err
	}
	if flag := ArraySorted(data); flag == false {
		return 0, nil
	}
	return slices.Max(data) - slices.Min(data), nil
}

func (s *Service) AnswerTask2(input1, input2 string) (bool, error) {
	err := errors.New("empty string")
	if len(input1) == 0 || len(input2) == 0 {
		return false, err
	}
	err = errors.New("expect more one letter")
	if len(input1) == 1 || len(input2) == 1 {
		return false, err
	}
	err = errors.New("different lengths")
	if len(input1) != len(input2) {
		return false, err
	}
	input1Byte := []byte(input1)
	input2Byte := []byte(input2)
	sort.Slice(input1Byte, func(i, j int) bool { return input1Byte[i] < input1Byte[j] })
	sort.Slice(input2Byte, func(i, j int) bool { return input2Byte[i] < input2Byte[j] })
	if string(input1Byte) == string(input2Byte) {
		return true, nil
	} else {
		return false, nil
	}

}

func (s *Service) AnswerTask3(data string, key int) (string, error) {
	err := errors.New("expect more one letter")
	if len(data) == 1 {
		return "", err
	}
	err = errors.New("empty string")
	if len(data) == 0 {
		return "", err
	}

	dataRune := make([]rune, len(data))
	dataRune = []rune(data)
	var newData string
	for i := 0; i < key; i++ {
		for j := 0; j < len(dataRune)+key*2-2; j = j + key*2 - 2 {
			if j+i < len(dataRune) {
				if i == 0 || i == key-1 {
					newData = newData + string(dataRune[i+j])
				} else {
					if j-i > 0 {
						newData = newData + string(dataRune[j-i])
					}
					newData = newData + string(dataRune[i+j])
				}

			}
		}
	}
	return newData, nil
}
