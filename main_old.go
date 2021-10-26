package main

import (
	"bufio"
	"encoding/csv"
	"fmt"

	"github.com/nahuelmarianolosada/test_project/conversions"
	dateUtils "github.com/nahuelmarianolosada/test_project/dateutils"

	"io"
	"log"
	"os"
	"time"

	stringUtils "github.com/nahuelmarianolosada/test_project/stringutils"
)

func mainOld() {

	fmt.Println(stringUtils.Reverse("hola"))

	fmt.Println("es Nuevo? " + conversions.BoolToString(true))

	current := time.Now()
	fmt.Println(dateUtils.ToString(current))
	fmt.Println("Hoy: ", dateUtils.FormatShort(current))
	fmt.Println("Dentro de un mes: ", dateUtils.FormatShort(dateUtils.AddMonths(current, 1)))
	fmt.Println("Mes pasado: ", dateUtils.FormatShort(dateUtils.AddMonths(current, -1)))
	fmt.Println("MM-DD-YYYY hh:mm:ss :", dateUtils.FormatLong(current))

	dateToValidated := current.Add(-24 * time.Hour)
	fmt.Println(dateToValidated)

	fmt.Println(current.Local())
	fmt.Println(current.Location() == time.UTC)

	dateToValidate2 := current.AddDate(0, 0, -1)
	fmt.Println(dateToValidate2)

	/*
		fmt.Println(stringUtils.IsEmpty(""))
		fmt.Println(stringUtils.GetOSEnvironment())

		helloWorld := " Hello, world "
		fmt.Printf("Largo: %d \t '%s'\n", len(helloWorld), helloWorld)
		fmt.Printf("Largo: %d \t '%s'\n",len(stringUtils.Trimmer(helloWorld)), stringUtils.Trimmer(helloWorld))*/

	/*list := readDataFromCSV("user_ids.csv")

	for _,value := range list {
		if value != "" {
			fmt.Println(`"` + value + `",` + getEmailByUserId(value))
		}
	}*/

}

func readDataFromCSV(fileName string) []string {
	var listResponse []string
	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		listResponse = append(listResponse, line[0])
	}

	return listResponse
}
