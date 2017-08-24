package main

//для сбори проекта ввести в консоль
//go build maine.go

//пример запроса
// ./maine --currency rub --value 123
// ./maine --currency RUB --value 512
// ./maine --currency usd --value 228
// ./maine --currency USD --value 522

//для проекта использовалось апи http://fixer.io/

//поддерживаемые валюты: "AUD", "BGN","BRL","CAD","CHF","CNY","CZK","DKK", "GBP","HKD" ,"HRK","HUF","IDR",
// "ILS", "INR", "JPY","KRW","MXN","MYR","NOK","NZD","PHP","PLN","RON","SEK","SGD","THB","TRY","USD","ZAR"

import (
	"os"
	"encoding/json"
	//для работы с консолью
	"github.com/urfave/cli"
	"net/http"
	"io/ioutil"
	"strconv"
	"log"
	"strings"
)

const (
	fixerPath = "http://api.fixer.io/latest?base="
)

type FixerResponse struct {
	_ string
	_ string
	Rates map[string]float64 `json:"rates"`
}
pus
func main() {

	app := cli.NewApp()

	app.Flags = []cli.Flag {
		cli.StringFlag{

			Name: "currency",
			Value: "def_cur",
			Usage: "2",
		},

		cli.StringFlag{
			Name: "value",
			Value: "def_sum",
			Usage: "2",
		},
	}

	app.Action = func(c *cli.Context) error {

		curr := strings.ToUpper(c.String("currency"))
		val := c.Float64("value")

		dotRequest(val, curr)

		return nil
	}

	app.Run(os.Args)
}

//сделать запрос и распечатать ответ
func dotRequest(value float64, currency string)  {

	client := http.Client{}

	resp, err := client.Get (fixerPath + currency)
	checkError(err)

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	fixer := FixerResponse{}
	err = json.Unmarshal(body, &fixer)
	checkError(err)

	log.Println(fixer.Rates)

	for k, v := range fixer.Rates {
		log.Print(floatToString(value) + " " + currency  + " = " +
			floatToString(v) + " " +  k)
	}
}

//преобразовать флот 64 в строку
func floatToString(input float64) string {
	return strconv.FormatFloat(input, 'f', 2, 64)
}

//проверка ошбки
func checkError(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}
