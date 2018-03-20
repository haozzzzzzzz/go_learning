package main

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"fmt"
)

func main() {
	xlsx, err := excelize.OpenFile("F:\\github\\go_learning\\src\\excel\\new\\simple_create.xlsx")
	if err != nil {
		fmt.Println(err)
	    return
	}

	cell := xlsx.GetCellValue("Sheet1", "B2")
	fmt.Println(cell)

	xlsx.SetCellValue("Sheet1", "C1", 666)

	rows := xlsx.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}

		fmt.Println()
	}
	xlsx.Save()
}