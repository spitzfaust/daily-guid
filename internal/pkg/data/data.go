package data

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/afero"
)

// FileSystem is the used file sytem.
var FileSystem = afero.NewOsFs()

const fileName = "./daily-guid.txt"

type saveFile struct {
	timestamp time.Time
	guid      uuid.UUID
}

const separator = " "

func unmarshalSaveFile(fileContents string) (*saveFile, error) {
	cleaned := strings.TrimSpace(fileContents)
	tokens := strings.Split(cleaned, separator)
	if len(tokens) != 2 {
		return nil, errors.New("save file does not contain exactly two tokens")
	}
	timestamp, err := time.Parse(time.RFC3339, tokens[0])
	if err != nil {
		return nil, err
	}
	guid, err := uuid.FromString(tokens[1])
	if err != nil {
		return nil, err
	}

	return &saveFile{timestamp: timestamp, guid: guid}, nil
}

func marshallSaveFile(saveFile saveFile) string {
	guid := saveFile.guid.String()
	timestamp := saveFile.timestamp.UTC().Format(time.RFC3339)
	return fmt.Sprintf("%s %s", timestamp, guid)
}

func datesEqualUTC(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.UTC().Date()
	y2, m2, d2 := date2.UTC().Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func generateNewSaveFile() *saveFile {
	guid := uuid.NewV4()
	s := &saveFile{timestamp: time.Now(), guid: guid}
	return s
}

func writeSaveFile(s saveFile) error {
	marshalled := marshallSaveFile(s)

	f, err := FileSystem.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(marshalled + "\n")
	if err != nil {
		return err
	}
	return nil
}

func readLineFromFile() (string, error) {
	f, err := FileSystem.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", errors.New("could not read file contents")
}

func readSaveFileFromExistingFile() (*saveFile, error) {
	fileContents, err := readLineFromFile()
	if err != nil {
		return nil, err
	}
	s, err := unmarshalSaveFile(fileContents)
	if err != nil {
		return nil, err
	}
	if !datesEqualUTC(s.timestamp, time.Now()) {
		s = generateNewSaveFile()
		err = writeSaveFile(*s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func createSaveFile() (*saveFile, error) {
	s := generateNewSaveFile()
	err := writeSaveFile(*s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// GetDailyGUID retrieves the daily GUID.
func GetDailyGUID() (*uuid.UUID, error) {
	fileExists, err := afero.Exists(FileSystem, fileName)
	if err != nil {
		return nil, err
	}
	var s *saveFile
	if fileExists {
		s, err = readSaveFileFromExistingFile()
	} else {
		s, err = createSaveFile()
	}
	if err != nil {
		return nil, err
	}
	return &s.guid, nil
}
