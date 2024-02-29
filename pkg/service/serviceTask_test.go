package service

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArraySorted(t *testing.T) {
	var data []int
	data = append(data, 48, 34, 23, 6, 5, 4, 2, 1)
	output := ArraySorted(data)
	require.Equal(t, false, output)
}

func TestArraySortedTrue(t *testing.T) {
	var data []int
	data = append(data, 48, 4, 23, 6, 5, 4, 2, 1)
	output := ArraySorted(data)
	require.Equal(t, true, output)
}

func TestService_AnswerTask1(t *testing.T) {
	var data []int
	data = append(data, 3, 48, 34, 5, 64, 23, 56, 1, 5, 9)
	var Serv Service

	output, err := Serv.AnswerTask1(data)
	require.NoError(t, err)
	require.Equal(t, 63, output)
}

func TestService_AnswerTask1SortArray(t *testing.T) {
	var data []int
	data = append(data, 48, 34, 23, 6, 5, 4, 2, 1)
	var Serv Service
	output, err := Serv.AnswerTask1(data)
	require.NoError(t, err)
	require.Equal(t, 0, output)
}

func TestService_AnswerTask1ErrorLen(t *testing.T) {
	var data []int
	data = append(data, 5)
	var Serv Service
	output, err := Serv.AnswerTask1(data)
	require.EqualError(t, err, "expect more one number")
	require.Equal(t, 0, output)
}
func TestService_AnswerTask1ErrorEmptyArray(t *testing.T) {
	var data []int
	var Serv Service
	output, err := Serv.AnswerTask1(data)
	require.EqualError(t, err, "empty array")
	require.Equal(t, 0, output)
}
func TestService_AnswerTask2(t *testing.T) {
	var data1 string = "abet"
	var data2 string = "beat"
	var Serv Service

	output, err := Serv.AnswerTask2(data1, data2)
	require.NoError(t, err)
	require.Equal(t, true, output)
}

func TestService_AnswerTask2False(t *testing.T) {
	var data1 string = "aber"
	var data2 string = "beat"
	var Serv Service

	output, err := Serv.AnswerTask2(data1, data2)
	require.NoError(t, err)
	require.Equal(t, false, output)
}

func TestService_AnswerTask2ErrorDifferentLengths(t *testing.T) {
	var data1 string = "abet"
	var data2 string = "beati"
	var Serv Service

	output, err := Serv.AnswerTask2(data1, data2)
	require.EqualError(t, err, "different lengths")
	require.Equal(t, false, output)
}

func TestService_AnswerTask2ErrorEmptyString(t *testing.T) {
	var data1 string = "tdds"
	var data2 string = ""
	var Serv Service

	output, err := Serv.AnswerTask2(data1, data2)
	require.EqualError(t, err, "empty string")
	require.Equal(t, false, output)
}
func TestService_AnswerTask2ErrorOneLetter(t *testing.T) {
	var data1 string = "tdds"
	var data2 string = "e"
	var Serv Service

	output, err := Serv.AnswerTask2(data1, data2)
	require.EqualError(t, err, "expect more one letter")
	require.Equal(t, false, output)
}

func TestService_AnswerTask3(t *testing.T) {
	var data string = "incomprehensibilities"
	var key int = 4
	var Serv Service

	output, err := Serv.AnswerTask3(data, key)
	require.NoError(t, err)
	require.Equal(t, "iriinpesbtecmhniisoel", output)
}

func TestService_AnswerTask3ErrorEmptyString(t *testing.T) {
	var data string = ""
	var key int = 4
	var Serv Service
	output, err := Serv.AnswerTask3(data, key)
	require.EqualError(t, err, "empty string")
	require.Equal(t, "", output)
}

func TestService_AnswerTask3ErrorOneLetter(t *testing.T) {
	var data string = "a"
	var key int = 4
	var Serv Service
	output, err := Serv.AnswerTask3(data, key)
	require.EqualError(t, err, "expect more one letter")
	require.Equal(t, "", output)
}
