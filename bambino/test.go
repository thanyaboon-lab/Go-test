package bambino

import (
	"Bambino/bambino/internal/bam"
	"fmt"
)

func generateTest() {
	fmt.Println("Say Hello Test")
}

func SayTest() {
	generateTest()
	bam.SayInternalBam()
}
