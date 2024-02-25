package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/sajari/regression"
)

func main() {
	//start time usage
	start := time.Now()

	//import data
	f, err := excelize.OpenFile("ames_housing_data.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	sheetName := f.GetSheetName(1)
	rows := f.GetRows(sheetName)

	//build regression model
	r := new(regression.Regression)

	//set Y variable
	r.SetObserved("SalePrice")

	for _, row := range rows {
		if row[0] == "SalePrice" {
			continue
		}

		var values []float64
		for _, v := range row {
			val, err := strconv.ParseFloat(v, 64)
			if err != nil {
				log.Printf("Failed to parse float: %v", err)
				continue
			}
			values = append(values, val)
		}

		r.Train(regression.DataPoint(values[0], values[1:]))
	}

	//run regression
	r.Run()

	//print summary stats for regression model
	fmt.Printf("Regression Formula: %s\n", r.Formula)
	fmt.Printf("R^2: %f\n", r.R2)
	fmt.Printf("Coefficients: %v\n", r.Coeff)

	//calculate processing time
	elapsed := time.Since(start)
	fmt.Printf("Processing Time: %s\n", elapsed)
}
