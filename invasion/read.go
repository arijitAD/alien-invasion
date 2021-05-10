package invasion

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

// asyncLineReader reads each line from a file and write to a channel.
func asyncLineReader(path string, lineCh chan string) error {
	curr, err := os.Getwd()
	if err != nil {
		return err
	}

	filePath := filepath.Join(curr, path)
	log.Println("filepath: ", filePath)

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	r := bufio.NewReader(file)

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lineCh <- scanner.Text()
	}

	return nil
}
