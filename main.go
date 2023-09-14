package main

import (
	"fmt"
	"strings"

	"github.com/obfio/hextree-pow-golang/hextree"
	"github.com/obfio/hextree-pow-golang/sha256"
)

func main() {
	client, err := hextree.MakeClient()
	if err != nil {
		panic(err)
	}
	config, err := client.GetConfig()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", config)
	difficultyPrefix := strings.Repeat("0", config.Difficulty)
	output := []string{}
	for i := 1; len(output) < 20; i++ {
		num, _ := sha256.Sum256([]byte(config.Prefix + fmt.Sprint(i)))
		num = trippleShift(num, 0)
		str := fmt.Sprintf("%b", num)
		str = strings.Repeat("0", 32-len(str)) + str
		if checkResult(str, difficultyPrefix) {
			output = append(output, fmt.Sprint(i))
		}
	}
	header := config.Prefix + ":" + strings.Join(output, ",")
	body, err := client.TestHeader(header)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Header sent:\n	Header: %s\n	Response Body: %s\n", header, body)
}

func checkResult(str, difficultyPrefix string) bool {
	return strings.HasPrefix(str, difficultyPrefix)
}
