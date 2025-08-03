package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // comnverts name to coursename while converting to json
	Price    int
	Platform string
	Password string   `json:"-"`              // won't display password
	Tags     []string `json:"tags,omitempty"` // if empty then omitted
}

func main() {

	fmt.Println("Go Json")
	// encodeJson()
	decodeJson()
}

func encodeJson() {
	courses := []course{
		{"React", 288, "youtube", "123456", []string{"web dev", "js"}},
		{"angular", 288, "youtube", "123456", []string{"web dev", "js"}},
		{"redux", 288, "youtube", "123456", nil},
	}
	finalJson, err := json.Marshal(courses)
	if err != nil {
		panic(err)
	}
	fmt.Println("json value", string(finalJson))
	// alternative way
	finalJson2, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println("json value prettified using MarshallIndent", string(finalJson2))

}

func decodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "React",
		"Price": 288,
		"Platform": "youtube",
		"tags": ["web dev", "js"]
	}
`)

	var courseParsed course

	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &courseParsed)
	}
	fmt.Println("parsed courses ", courseParsed)

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Println("parsed using map:\n", myOnlineData)
}
