package log

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	pkgFile "log-generator/pkg/lib/file"
	tmpLog "log-generator/pkg/lib/template"
	pkgTime "log-generator/pkg/lib/time"
)

const (
	FORMAT_JSON = "json"
	FORMAT_TEXT = "text"
)

type FakeLog struct {
	config      *Config
	currentFile string
}

type Config struct {
	Name            string
	FilePath        string
	Source          string
	Separation      bool
	ErrCountLimit   int
	WriteCountLimit int
	Interval        int
}

func New(config *Config) *FakeLog {
	return &FakeLog{
		config: config,
	}
}

func (l *FakeLog) Mock() error {
	config := l.config
	logName := config.Name
	path := config.FilePath
	interval := time.Duration(config.Interval)
	fmt.Printf("[%s] mock strating...\n", logName)
	var (
		errBreak             bool
		errCount, writeCount int
	)
	t := tmpLog.New()
	l.refreshCurrentFile(path)
	for {
		if err := l.writeErrorLog(t, config); err != nil {
			log.Fatalln(err)
			errCount++
			continue
		} else {
			writeCount++
		}
		if errCount >= config.ErrCountLimit {
			errBreak = true
			break
		}
		if writeCount >= config.WriteCountLimit {
			break
		}
		time.Sleep(time.Second * interval)
	}
	if errBreak {
		return fmt.Errorf("failed to mock the fake logs")
	}
	fmt.Printf("[%s] mock done\n", logName)
	return nil
}

func (l *FakeLog) getCurrentFile() string {
	return l.currentFile
}

func (l *FakeLog) refreshCurrentFile(path string) {
	nowTime := pkgTime.GetLogTime()
	l.currentFile = fmt.Sprintf("%s/%s.log", path, nowTime)
}

func (l *FakeLog) writeErrorLog(t *tmpLog.Log, config *Config) error {
	source := config.Source
	path := config.FilePath
	errLog, err := genErrorLog(t, source)
	if err != nil {
		return fmt.Errorf("failed to generate error log: %v", err)
	}

	f, err := l.getWriteFile(path)
	if err != nil {
		return fmt.Errorf("failed to open file by path %s: %v", path, err)
	}
	defer f.Close()

	// write file
	log := fmt.Sprintf("%s\n", errLog)
	logText := []byte(log)
	if _, err := f.Write(logText); err != nil {
		return fmt.Errorf("failed to write the new context: %v", err)
	}
	return f.Sync()
}

func (l *FakeLog) getWriteFile(path string) (*os.File, error) {
	var (
		fileCount int
		done      bool
		f         *os.File
	)
	separation := l.config.Separation
	for {
		fileCount++
		currentFile := l.getCurrentFile()
		of, count, err := getFileLineCount(currentFile)
		if err != nil {
			return nil, fmt.Errorf("failed to get file's line count: %v", err)
		}
		if count > 5 && separation {
			defer f.Close()
			l.refreshCurrentFile(path)
			continue
		}
		f = of
		done = true
		break
	}
	if !done {
		return nil, fmt.Errorf("failed to found correct file")
	}
	return f, nil
}

func getFileLineCount(path string) (*os.File, int, error) {
	of, err := openFile(path)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to open file by path %s: %v", path, err)
	}
	count, err := pkgFile.CountLine(bufio.NewReader(of))
	if err != nil {
		return nil, 0, fmt.Errorf("failed to read file's content: %v", err)
	}
	return of, count, nil
}

func openFile(path string) (*os.File, error) {
	var f *os.File
	if _, err := os.Stat(path); err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("failed to check file exist: %v", err)
		}
		f, err = os.Create(path)
		if err != nil {
			return nil, fmt.Errorf("failed to create a new file: %v", err)
		}
	} else {
		f, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			return nil, fmt.Errorf("failed to open this file: %v", err)
		}
	}
	return f, nil
}

func genErrorLog(t *tmpLog.Log, source string) (string, error) {
	mackLevel := randLevels([]string{
		"debug",
		"info",
		"warn",
	})

	nowTime := pkgTime.GetNowTime()
	var errLog string
	switch source {
	case FORMAT_JSON:
		if s, err := t.GetJsonFormatTemplate(nowTime, mackLevel); err != nil {
			return "", fmt.Errorf("failed to get json template: %v", err)
		} else {
			errLog = s
		}
	case FORMAT_TEXT:
		if s, err := t.GetTextFormatTemplate(nowTime, mackLevel); err != nil {
			return "", fmt.Errorf("failed to get text template: %v", err)
		} else {
			errLog = s
		}
	default:
		return "", fmt.Errorf("failed to get mock error log")
	}
	return errLog, nil
}

func randLevels(levels []string) string {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	i := r.Intn(len(levels))
	return levels[i]
}
