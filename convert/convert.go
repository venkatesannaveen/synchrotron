package convert

import (
	"bufio"
	"log"
	"os"
)

// ParseBinary returns
func ParseBinary(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		log.Fatal(err)
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	return bytes, err
}
