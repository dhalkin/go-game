/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
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
	url := "https://icanhazdadjoke.com/"
	responseBytes := getJokeData(url)
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
