package fileini

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

type INIFile struct {
	content string
}

func (iniFile *INIFile) Init(path string) {
	cfg, err := ini.Load(path)

	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	fmt.Println(cfg)
}
