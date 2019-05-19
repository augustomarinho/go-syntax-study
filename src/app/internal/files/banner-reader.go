package files

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Banner struct {
	content string
}

func (banner *Banner) NewBanner(pathFile string) {

	file, err := os.Open(pathFile)

	if err != nil {
		fmt.Println("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	var str strings.Builder

	for scanner.Scan() {
		str.Write(scanner.Bytes())
	}

	file.Close()

	banner.content = str.String()
}

func (b Banner) GetContent() string {
	return b.content
}
