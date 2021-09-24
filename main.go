package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

type Format struct {
	Data [][]string
}

func main() {
	colorCyan := "\033[36m"
	fmt.Println(string(colorCyan), "Inicio =>", now())
	sourcePathV1 := "/Users/lraga/Downloads/Base V1.csv"
	sourcePathV2 := "/Users/lraga/Downloads/Base V2.csv"
	linesV1 := readCsv(sourcePathV1)
	linesV2 := readCsv(sourcePathV2)

	fmt.Println(len(linesV1))
	fmt.Println(len(linesV2))
	createCsv(linesV2, linesV1)
	fmt.Println(string(colorCyan), "Fin =>", now())

}

func createCsv(data1 [][]string, data2 [][]string) {
	csvFile, err := os.Create("compare.csv")
	if err != nil {
		fmt.Println("failed creating file:", err)
	}
	csvwriter := csv.NewWriter(csvFile)
	// lineAll := append(data1,  data2)

	for i, line1 := range data1 {
		if i == 0 {
			newlene2 := append(line1, "compare1", "version")

			fmt.Println(newlene2)

			_ = csvwriter.Write(newlene2)
			continue
		}
		if i > 0 {
			for y, line2 := range data2 {
				if y == 0 {
					continue
				}
				if i > 0 {
					// todo v2 y v1 con una columna adicional
					if line1[2] == line2[2] {
						line1 := append(line1, line1[3], "V2")
						fmt.Println("if => ", line1, line2)
						_ = csvwriter.Write(line1)
						break
					}
				}
			}
		}
	}

	for i, line2 := range data2 {
		if i == 0 {
			newlene2 := append(line2, "compare1", "version")

			fmt.Println(newlene2)

			_ = csvwriter.Write(newlene2)
			continue
		}
		if i > 0 {
			for y, line1 := range data1 {
				if y == 0 {
					continue
				}
				if i > 0 {
					// todo v2 y v1 con una columna adicional

					if line1[2] == line2[2] {
						// line2[11] = line1[3]
						line2 := append(line2, line1[3], "V1")
						fmt.Println("if => line1 ", line1, line2)
						_ = csvwriter.Write(line2)
						break
					}
				}
			}
		}
	}
	// _ = csvwriter.WriteAll(data1)
	// _ = csvwriter.WriteAll(data2)
	csvFile.Close()
}

func readCsv(sourcePath string) [][]string {
	file, _ := os.Open(sourcePath)
	lines, _ := csv.NewReader(file).ReadAll()
	return lines
}

func now() string {
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	return fecha
}
