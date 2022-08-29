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

func ParseBinaryFloat(filename string, width int16, height int16) ([][]float32, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return [][]float32{}, err
	}

	buf := bytes.NewReader(file)
	var output [][]float32

	for i := 0; i < int(height); i++ {
		row := make([]float32, width)
		for j := 0; j < int(width); j++ {
			var intensity float32
			err = binary.Read(buf, binary.LittleEndian, &intensity)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, intensity)
		}
		output = append(output, row)
	}

	return output, nil
}
