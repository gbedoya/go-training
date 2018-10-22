package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page int
	Fruits []string
}

type response2 struct {
	Page	int		`json:"page"`
	Fruits	[]string `json:"fruits"`
}

// https://blog.goland.org/json-and-go
// https://goland.org/pkg/encoding/json
func JSON() {

	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.15)
	fmt.Println(string(fltB))

	StrB, _ := json.Marshal("GOLANG")
	fmt.Println(string(StrB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// The JSON package can automatically encode your custom data types,
	// It will only include exported fields in the encoded output
	// and will by default use those names as the JSON keys.
	res1D := &response1 {
		Page: 	1,
		Fruits:	[]string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// You can use tags on struct field declarations
	// to customize	the encoded JSON key names.
	// Check the definition of response2 above to see an example of such tags.
	res2D := &response2{
		Page: 	1,
		Fruits:	[]string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Decoding JSON data into Go values
	byt := []byte(`{"num": 6.13, "strs": ["a", "b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(float64)
	//num := dat["num"]			// returns 6.13
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)
	//fmt.Println(strs[0])		// returns a


	// We can also decode JSON into custom data types.
	// This has the advantages of adding additional type-safety to
	// our programs and eliminating the need for type assertions when
	// accessing the decoded data
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "pear": 3}
	enc.Encode(d)

}
