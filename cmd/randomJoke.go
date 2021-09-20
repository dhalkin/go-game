package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

// randomJokeCmd represents the randomJoke command
var randomJokeCmd = &cobra.Command{
	Use:   "randomJoke",
	Short: "Get a random joke",
	Long: `Get a random joke from :

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("randomJoke called")
		getRandomJoke()
	},
}

func init() {
	rootCmd.AddCommand(randomJokeCmd)
}

func getRandomJoke() {

	url1 := viper.Get("api.dadjoke.baseurl")
	// convert to string ?
	url2 := fmt.Sprintf("%v", url1)

	responseBytes := getJokeData(url2)
	joke := Joke{}
	if err := json.Unmarshal(responseBytes, &joke); err != nil {
		fmt.Println("Could not unmarshall response - %v", err)
	}

	fmt.Println(string(joke.Joke))
}

func getJokeData(baseApi string) []byte {
	request, err := http.NewRequest(
		http.MethodGet,
		baseApi,
		nil,
	)

	if err != nil {
		fmt.Println("Error while creating a request - %v", err)
	}

	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Go Game CLI (github.com/dhalkin/go-game)")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("Could not make a request - %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Could not read a response - %v", err)
	}

	return responseBytes
}
