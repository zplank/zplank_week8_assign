package main

import (
	"fmt"
	"log"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	// Open the Excel file
	f, err := excelize.OpenFile("ames_housing_data.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	// Get the sheet name
	sheetName := f.GetSheetName(1)

	// Get all the rows in the sheet
	rows := f.GetRows(sheetName)

	// Loop through each row
	for _, row := range rows {
		// Assuming the columns are in order as specified
		salePrice := row[0]
		totalBsmtSF := row[1]
		firstFlrSF := row[2]
		secondFlrSF := row[3]
		lotArea := row[4]
		overallQual := row[5]
		bsmtFullBath := row[6]
		bsmtHalfBath := row[7]
		fullBath := row[8]
		halfBath := row[9]
		bedroomAbvGr := row[10]
		kitchenAbvGr := row[11]
		totRmsAbvGrd := row[12]
		fireplaces := row[13]

		// Do something with the data, for example, print it
		fmt.Printf("SalePrice: %s, TotalBsmtSF: %s, FirstFlrSF: %s, SecondFlrSF: %s, LotArea: %s, OverallQual: %s, BsmtFullBath: %s, BsmtHalfBath: %s, FullBath: %s, HalfBath: %s, BedroomAbvGr: %s, KitchenAbvGr: %s, TotRmsAbvGrd: %s, Fireplaces: %s\n", salePrice, totalBsmtSF, firstFlrSF, secondFlrSF, lotArea, overallQual, bsmtFullBath, bsmtHalfBath, fullBath, halfBath, bedroomAbvGr, kitchenAbvGr, totRmsAbvGrd, fireplaces)
	}
}
