package particle

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func Trigger() {
	deviceID := os.Getenv("DEVICE_ID")
	accessToken := os.Getenv("ACCESS_TOKEN")
	functionName := os.Getenv("PARTICLE_FUNCTION_NAME")
	endpoint := "https://api.particle.io/v1/devices/" + deviceID + "/" + functionName
	resp, err := http.PostForm(endpoint, url.Values{"access_token": {accessToken}})
	if err != nil {
		fmt.Println("ERRORED")
		panic(err)
	}
	fmt.Println(resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("ERRORED ON BODY")
		panic(err)
	}
	fmt.Println(string(body))
}
