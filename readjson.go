package stratifier

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ReadJson(filename string) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file %s", filename, err)
	}
	model := new(map[string][]string)
	err = json.Unmarshal(contents, model)
	if err != nil {
		log.Fatalf("failed to encode json", err)
	}
	for k, v := range *model {
		log.Print(k, v)
	}
}
