package testTask

type InputTask1 struct {
	Array []int `json:"array"`
}
type InputTask2 struct {
	InputString string `json:"inputString"`
	Anagram     string `json:"anagram"`
}

type InputTask3 struct {
	InputString string `json:"inputString"`
	Key         int    `json:"key"`
}
