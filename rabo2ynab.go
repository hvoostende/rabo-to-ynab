package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
)

//GetYNABDate converts Rabo date format to YNAB date format
func GetYNABDate(r1 string) string {
	return "date"
}

//GetYNABPayee converts Rabo info fields to a YNAB Payee field
func GetYNABPayee(r1, r2, r3, r4 string) string {
	return "payee"
}

//GetYNABMemo converts Rabo info field to YNAB Memo field
func GetYNABMemo(r1 string) string {
	return "memo"
}

//GetYNABCategory for now returns nothing can be further implemented
func GetYNABCategory() string {
	return ""
}

//GetYNABInflow converts Rabo Debet/Credit field and value to YNAB inflow
func GetYNABInflow(r1, r2 string) string {
	return "inflow"
}

//GetYNABOutflow converts Rabo Debet/Credit field and value to YNAB outflow
func GetYNABOutflow(r1, r2 string) string {
	return "outflow"
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

func main() {
	//mandatory csv file
	var input string

	flag.StringVar(&input, "input", "transactions.txt", "csv file")

	// load data csv
	records, err := readCSV(input)
	if err != nil {
		log.Fatal(err)
	}

	// write results to a new csv
	outfile, err := os.Create("resultsfile.csv")
	if err != nil {
		log.Fatal("Unable to open output")
	}
	defer outfile.Close()
	writer := csv.NewWriter(outfile)

	for _, record := range records {
		date := GetYNABDate(record[2])
		payee := GetYNABPayee(record[10], record[11], record[12], record[13])
		category := GetYNABCategory()
		memo := GetYNABMemo(record[10])
		outflow := GetYNABOutflow(record[3], record[4])
		inflow := GetYNABOutflow(record[3], record[4])

		writer.Write([]string{date, payee, category, memo, outflow, inflow})
	}

	writer.Flush()

}
