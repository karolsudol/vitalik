package travelsaver

import (
	"encoding/json"
	"fmt"
	"log"
)

func prettyPrint(d ...interface{}) {
	b, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
