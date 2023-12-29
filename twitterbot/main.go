package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/ChimeraCoder/anaconda"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	api := anaconda.NewTwitterApi(os.Getenv("TWITTER_ACCESS_TOKEN"), os.Getenv("TWITTER_ACCESS_SECRET"))

	v := url.Values{}
	v.Set("status", "Hello World!")

	// Example tweet
	resp, err := api.Client.PostForm("https://api.twitter.com/1.1/statuses/update.json", v)
	if err != nil {
		fmt.Println("Failed to send tweet:", err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to send tweet with status code", resp.StatusCode)
		return
	}

	fmt.Println("Successfully sent tweet.")
}
