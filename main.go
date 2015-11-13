package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/stacktic/dropbox"
)

type config struct {
	Clientid     string `json:"clientid"`
	Clientsecret string `json:"clientsecret"`
	Token        string `json:"token"`
	Destination  string `json:"destination"`
	DropboxPath  string `json:"dropboxpath"`
}

func main() {

	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("Error opening config file: " + err.Error())
	}

	var c config
	err = json.Unmarshal(configFile, &c)
	if err != nil {
		panic(err)
	}

	var db *dropbox.Dropbox
	// 1. Create a new dropbox object.
	db = dropbox.NewDropbox()

	// 2. Provide your clientid and clientsecret (see prerequisite).
	db.SetAppInfo(c.Clientid, c.Clientsecret)

	// 3. Provide the user token.
	db.SetAccessToken(c.Token)
	fmt.Println(c.DropboxPath)
	err = db.DownloadToFile(c.DropboxPath, c.Destination, "")
	if err != nil {
		panic(err)
	}
}
