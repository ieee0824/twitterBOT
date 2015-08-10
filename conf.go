package bot

import (
	"encoding/json"
	"io/ioutil"

	"github.com/shirase-aoi/log"
)

type Conf struct {
	AccessToken       string
	AccessTokenSelect string
	ConsumerKey       string
	ConsumerSelect    string
}

func CreateConfigureFile(accessToken string, accessTokenSelect string, consumerKey string, consumerSelect string, fileName string) {
	if fileName == "" {
		fileName = "../conf/configure.json"
	}
	c := Conf{accessToken, accessTokenSelect, consumerKey, consumerSelect}
	js, err := json.Marshal(c)
	log.Terminate(err, "../log/error.json")
	err = ioutil.WriteFile(fileName, js, 0644)
	log.Terminate(err, "../log/error.json")
}

func CreateConfigureTemplate() {
	CreateConfigureFile("", "", "", "", "../conf/configureTemp.json")
}

func LoadConfigure(fileName string) *Conf {
	if fileName == "" {
		fileName = "../conf/configure.json"
	}
	js, err := ioutil.ReadFile(fileName)
	log.Terminate(err, "../log/error.json")
	c := Conf{}
	err = json.Unmarshal(js, &c)
	log.Terminate(err, "../log/error.json")
	return &c
}
