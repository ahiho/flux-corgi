package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	swaggerFile := args[0]
	b, e := os.ReadFile(swaggerFile)
	if e != nil {
		log.Fatalln(e)
	}
	content := string(b)
	swaggerMap := map[string]interface{}{}
	e = json.Unmarshal(b, &swaggerMap)
	if e != nil {
		log.Fatalln(e)
	}
	tags := swaggerMap["tags"].([]interface{})

	newTags := []interface{}{}

	for _, tag := range tags {
		tagMap := tag.(map[string]interface{})
		tagName := tagMap["name"].(string)
		if strings.Contains(content, fmt.Sprintf(`"operationId": "%v_`, tagName)) {
			newTags = append(newTags, tag)
		}
	}
	swaggerMap["tags"] = newTags

	b, e = json.MarshalIndent(swaggerMap, "", "  ")
	if e != nil {
		log.Fatalln(e)
	}
	os.WriteFile(swaggerFile, b, 0644)
}
