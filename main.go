package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sato48/mimo-counter/lib"
	"github.com/spf13/viper"
)

func main() {
	m := lib.NewMimo(
		viper.GetString("MIMO_HOST"),
		viper.GetInt("MIMO_PORT"),
		viper.GetInt("MIMO_DOCUMENT_ID"),
	)
	ts := lib.NewTCPServer(
		viper.GetString("TCP_HOST"),
		viper.GetInt("TCP_PORT"),
	)

	ts.RegisterHandler("inc", func(string) string {
		pnum, _ := m.IncrementNumericText(
			viper.GetString("MIMO_LAYER_ID"),
		)

		return strconv.Itoa(pnum)
	})

	ts.RegisterHandler("dec", func(string) string {
		pnum, _ := m.DecrementNumericText(
			viper.GetString("MIMO_LAYER_ID"),
		)

		return strconv.Itoa(pnum)
	})

	ts.RegisterHandler("get", func(string) string {
		pnum, _ := m.GetNumericText(
			viper.GetString("MIMO_LAYER_ID"),
		)

		return strconv.Itoa(pnum)
	})

	ts.RegisterHandler("set", func(req string) string {
		args := strings.Split(req, " ")
		if len(args) != 2 {
			return string("err")
		}
		arg := args[1]
		argI, err := strconv.Atoi(arg)
		if err != nil {
			return string("err")
		}

		m.SetNumericText(
			viper.GetString("MIMO_LAYER_ID"),
			argI,
		)

		return arg
	})

	fmt.Println(ts.ListenAndServe())
}

func init() {
	viper.AutomaticEnv()

	viper.SetDefault("MIMO_HOST", "localhost")
	viper.SetDefault("MIMO_PORT", 8989)

	viper.SetDefault("TCP_HOST", "localhost")
	viper.SetDefault("TCP_PORT", 6000)
}
