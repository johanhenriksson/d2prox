package d2prox

import (
	"encoding/csv"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type TsvUnmarshaller func(record []string) error

func ParseTsv(path string, unmarshal TsvUnmarshaller) error {
	csvFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = '\t'

	for {
		record, err := reader.Read()
		if record == nil {
			break
		}
		if err != nil {
			return err
		}

		if err := unmarshal(record); err != nil {
			return err
		}
	}

	return nil
}

func UnmarshalCsv(record []string, v interface{}) error {
	s := reflect.ValueOf(v).Elem()
	if s.NumField() != len(record) {
		return fmt.Errorf("record length doesn't match number of struct fields")
	}
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		switch f.Type().String() {
		case "string":
			f.SetString(record[i])
		case "int":
			ival, err := strconv.ParseInt(record[i], 10, 0)
			if err != nil {
				return err
			}
			f.SetInt(ival)
		default:
			return fmt.Errorf("unsupported type %s", f.Type().String())
		}
	}
	return nil
}
