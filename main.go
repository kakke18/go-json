package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	inputFileName  = "input.json"
	outputFileName = "output.json"
)

func main() {
	// step1: jsonファイルを読み込み
	inputBytes, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		fmt.Printf("[Error] %s\n", err.Error())
		return
	}

	// step2: ファイルデータから構造体を生成
	var users []User
	if err := json.Unmarshal(inputBytes, &users); err != nil {
		fmt.Printf("[Error] %s\n", err.Error())
		return
	}

	// step3: 構造体をjsonファイルに書き込みできるデータに変換
	outputBytes, err := json.MarshalIndent(&users, "", "  ") // 改行して出力するため
	if err != nil {
		fmt.Printf("[Error] %s\n", err.Error())
		return
	}

	// step4: 出力ファイルを生成
	file, err := os.OpenFile(outputFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Printf("[Error] %s\n", err.Error())
		return
	}
	defer file.Close()

	// step5: 出力ファイルに書き込み
	if _, err := file.Write(outputBytes); err != nil {
		fmt.Printf("[Error] %s\n", err.Error())
		return
	}

	return
}

type User struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Hobbies []Hobby `json:"hobbies"`
}

type Hobby struct {
	Category string `json:"category"`
	Name     string `json:"name"`
}
