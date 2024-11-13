package main

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

type versionStruct struct {
	year      int
	dayOfYear int
	micro     int
}

func CreateVersionStruct(micro int) *versionStruct {
	now := time.Now().UTC()
	version := &versionStruct{}
	version.year = now.Year() - 2000
	version.dayOfYear = int(now.YearDay())
	version.micro = 1
	return version
}

func ParseVersionStructFromString(s string) (*versionStruct, error) {
	data, err := os.ReadFile(s)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(string(data), ".")
	if len(parts) != 3 {
		return nil, errors.New("Invalid version file")
	}

	rawYear := parts[0]
	rawDayOfYear := parts[1]
	rawMicro := parts[2]

	year, err := strconv.Atoi(rawYear)
	if err != nil {
		return nil, errors.New("Invalid year in version file")
	}

	dayOfYear, err := strconv.Atoi(rawDayOfYear)
	if err != nil {
		return nil, errors.New("Invalid day of year in version file")
	}

	micro, err := strconv.Atoi(rawMicro)
	if err != nil {
		return nil, errors.New("Invalid micro in version file")
	}

	return &versionStruct{
		year:      year,
		dayOfYear: dayOfYear,
		micro:     micro,
	}, nil
}

func (self *versionStruct) String() string {
	return strconv.Itoa(self.year) + "." + strconv.Itoa(self.dayOfYear) + "." + strconv.Itoa(self.micro)
}

func (self *versionStruct) WriteToFile(name string) error {
	return os.WriteFile(name, []byte(self.String()), 0644)
}
