package csv

import (
	"bufio"
	"encoding/csv"
	"os"
	"strings"
	"ww/stocks/template"
	"ww/stocks/xerr"
)

type FileService struct {
	headerMap map[int]string           // [列位置][中文列名]
	dataMap   []map[string]interface{} // [英文数据库字段名][值]
}

// data [][英文数据库字段名][值]
func (c *FileService)ImportCSV(filePath string) (data []map[string]interface{}) {
	// get all rows from csv
	records := c.readData(filePath)
	// set head position/name to map
	c.setHeaderMap(records)
	// convert csv rows to []map[col]value ==> records [][]string
	// caller should convert them to database models from outside
	c.convertToMap(records)

	return c.dataMap
}

func (c *FileService)readData(filePath string)(records [][]string){
	file, err := os.Open(filePath)
	defer file.Close()
	xerr.ThrowError(err)

	rr := bufio.NewReader(file)
	tsv := csv.NewReader(rr)
	//tsv.Comma = '\t'
	tsv.Comment = '#'
	tsv.LazyQuotes = true
	//tsv.TrailingComma = true     // retain rather than remove empty slots
	tsv.TrimLeadingSpace = false // retain rather than remove empty slots

	records, err = tsv.ReadAll()
	xerr.ThrowError(err)

	return records
}

func (c *FileService)setHeaderMap(records [][]string){
	c.headerMap = make(map[int]string) // [列位置][中文列名]

	for _, rec := range records {
		for idx, name := range rec {
			n := strings.ReplaceAll(strings.TrimSpace(name), "\"", "")
			n = strings.TrimSpace(n)
			c.headerMap[idx] = n
		}
		break // just read first header line
	}
}

func (c *FileService)convertToMap(records [][]string){
	for idx, record := range records {
		if idx == 0 { // ignore header
			continue
		}
		mm := make(map[string]interface{})
		for idx, value := range record {
			if columnName, ok := c.headerMap[idx]; ok {
				// get english name from chinese
				key := c.getEnglishColumnName(columnName)
				if key != "" {
					c.setColumnValue(mm,key, strings.TrimSpace(value))
				}
			}
		}

		c.dataMap = append(c.dataMap,mm)
	}
}

func (c *FileService)getEnglishColumnName(key string) string {
	if v, ok := template.FutuColumnMap[key]; ok {
		return v
	}

	return ""
}

// key = engColumnName,datetime
func (c *FileService)setColumnValue(mm map[string]interface{},key string, value string) {
	arr := strings.Split(key,",")
	col := arr[0]
	columnType := arr[1] // string,stringToInt,datetime

	var v interface{}
	if  columnType == "string" {
		v = toString(value)
	}
	if  columnType == "float" {
		v = toFloat(value)
	}
	if  columnType == "int" {
		v = toInt(value)
	}
	if  columnType == "datetime" {
		v = toDatetime(value)
	}
	if  columnType == "percent" {
		v = toPercent(value)
	}
	if  columnType == "money" {
		v = toMoney(value)
	}
	if  columnType == "money2" {
		v = toMoney2(value)
	}
	if columnType == "stringToFloat" {
		v = stringToFloat(value)
	}

	mm[col] = v
}