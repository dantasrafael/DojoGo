package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type animal struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	showMarshal()
	showUnmarshal()
}

func showMarshal() {
	bruce := animal{"Bruce", "Yorkshire", 1}
	bruceJson, err := json.Marshal(bruce)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bruceJson)
	fmt.Println(bytes.NewBuffer(bruceJson))

	rex := map[string]string{
		"nome": "rex",
		"raca": "vira-lata",
	}
	rexJson, err := json.Marshal(rex)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rexJson)
	fmt.Println(bytes.NewBuffer(rexJson))
}

func showUnmarshal() {
	ninaJson := `{"nome":"Nina","raca":"Golden","idade":5}`
	var nina animal

	err := json.Unmarshal([]byte(ninaJson), &nina)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(ninaJson)
	fmt.Println(nina)

	tobyJson := `{"nome":"Toby","raca":"Pitbull","idade":2}`
	toby := make(map[string]interface{})

	if err := json.Unmarshal([]byte(tobyJson), &toby); err != nil {
		log.Fatal(err)
	}

	fmt.Println(tobyJson)
	fmt.Println(toby)
}
