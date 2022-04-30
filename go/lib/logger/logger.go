package logger

import (
	"log"
	"os"
	"path"
	"strings"
)

var (
	_, filename = path.Split(os.Args[0])
	prefix      = strings.ToUpper(filename)
	HTTP        = log.New(os.Stdout, prefix+" [HTTP] ", log.LstdFlags)
	Access      = log.New(os.Stdout, prefix+" [ACCESS] ", log.LstdFlags)
	Warn        = log.New(os.Stdout, prefix+" [WARN] ", log.LstdFlags)
	Info        = log.New(os.Stdout, prefix+" [INFO] ", log.LstdFlags)
	Debug       = log.New(os.Stdout, prefix+" [DEBUG] ", log.LstdFlags)
	Error       = log.New(os.Stdout, prefix+" [ERROR] ", log.LstdFlags)
)
