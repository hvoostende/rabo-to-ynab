package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"time"
)

//GetYNABDate converts Rabo date format to YNAB date format
func GetYNABDate(r1 string) string {
	// example 20161028 -> 28/10/2016
	return r1[6:8] + "/" + r1[4:6] + "/" + r1[:4]
}

//GetYNABPayee converts Rabo info fields to a YNAB Payee field
func GetYNABPayee(r1, r2 string) string {
	if r1 == "" {
		return r2
	}
	return r1
}

//GetYNABInflow converts Rabo Debet/Credit field and value to YNAB inflow
func GetYNABInflow(r1, r2 string) string {
	if r1 == "C" {
		return r2
	}
	return ""
}

//GetYNABOutflow converts Rabo Debet/Credit field and value to YNAB outflow
func GetYNABOutflow(r1, r2 string) string {
	if r1 == "D" {
		return r2
	}
	return ""
}

func readCSV(filepath string) ([][]string, error) {
	csvfile, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}
	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	fields, err := reader.ReadAll()
	if err != nil {
		log.Fatal("cannot read csv file:", err)
	}

	return fields, nil
}

var input string
var output string
var test bool

func init() {
	flag.StringVar(&input, "input", "transactions.txt", "csv file")
	flag.StringVar(&output, "output", "ynabTransactions", "csv file")
	flag.BoolVar(&test, "test", false, "switch to true for main testing")
	flag.Parse()
}

func main() {

	// load data csv
	records, err := readCSV(input)
	if err != nil {
		log.Fatal("could not find your bank transactions file", err)
	}

	// write results to a new csv
	if !test {
		output = output + time.Now().String()
	}

	outfile, err := os.Create(output + ".csv")
	if err != nil {
		log.Fatal("Unable to open output")
	}
	defer outfile.Close()
	writer := csv.NewWriter(outfile)

	//write header
	writer.Write([]string{"Date", "Payee", "Category", "Memo", "Outflow", "Inflow"})

	for _, record := range records {
		date := GetYNABDate(record[2])
		payee := GetYNABPayee(record[6], record[10])
		category := "" //maybe create something to add category's for now let's do it in YNAB
		memo := record[10]
		outflow := GetYNABOutflow(record[3], record[4])
		inflow := GetYNABInflow(record[3], record[4])

		writer.Write([]string{date, payee, category, memo, outflow, inflow})
	}

	writer.Flush()

}
