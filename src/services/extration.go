package services

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"
)

//Extract - Extração dos dados do arquivo csv/txt
func Extract(eChan chan []string) {
	f, _ := os.Open("./base_teste.txt")
	defer f.Close()

	r, err := readSample(f)
	if err != nil {
		fmt.Println("ERROU")
	}

	for record, err := r.Read(); err == nil; record, err = r.Read() {
		eChan <- record
	}

	close(eChan)
}

func readSample(rs io.ReadSeeker) (*csv.Reader, error) {
	// Skip first row (line)
	row1, err := bufio.NewReader(rs).ReadSlice('\n')
	if err != nil {
		return nil, err
	}
	_, err = rs.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		return nil, err
	}
	fileBytes, _ := ioutil.ReadAll(rs)

	m := regexp.MustCompile(`(?sm)[ ]{1,}`)
	comma := regexp.MustCompile(`,`)
	fileSanitizada := comma.ReplaceAll(fileBytes, []byte("."))
	fileSanitizada = m.ReplaceAll(fileSanitizada, []byte(","))

	// Read remaining rows
	r := csv.NewReader(bytes.NewReader(fileSanitizada))
	return r, nil
}
