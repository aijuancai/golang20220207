
package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	out, err := exec.Command("nslookup", "google.com").Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))

	out2, err2 := exec.Command("ping", "-c 4 google.com").Output()

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(string(out2))


}
