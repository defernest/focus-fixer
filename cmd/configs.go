package cmd

import (
	"encoding/csv"
	"io"
	"os"

	api "focus-fixer/axis"

	"github.com/gocarina/gocsv"
)

func GetCameras(filepath string) ([]api.Camera, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return []api.Camera{}, err
	}
	defer file.Close()
	cameras := []api.Camera{}
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comment = '#'
		r.Comma = ','
		return r
	})
	gocsv.FailIfUnmatchedStructTags = true
	if err := gocsv.UnmarshalFile(file, &cameras); err != nil {
		return []api.Camera{}, err
	}
	return cameras, nil
}
