/* RZFeeser | Alta3 Research
   Generating YAML files      */

package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

func main() {

	trees := []string{"elm", "oak", "maple", "sycamore", "chestnut"}

	data, err := yaml.Marshal(&trees)

	if err != nil {
		log.Fatal(err)
	}

	err2 := ioutil.WriteFile("trees2.yaml", data, 0x777)

	if err2 != nil {

		log.Fatal(err2)
	}

	fmt.Println("data written")
}
