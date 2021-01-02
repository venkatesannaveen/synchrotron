package convert

import (
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"log"
)

// ParseBinary returns
func ParseBinary(filename string) []float32 {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var output []float32

	for i := 0; i <= len(file)-4; i += 4 {
		var intensity float32
		intensityBytes := file[i : i+4]
		buf := bytes.NewReader(intensityBytes)
		err = binary.Read(buf, binary.LittleEndian, &intensity)
		if err != nil {
			log.Fatal(err)
		}
		output = append(output, intensity)
	}

	return output
}
