package bot

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/shirase-aoi/log"
)

type Data struct {
	Post string
	Week string
	Time int
}

type Random struct {
	PostData []Data
}

func RegistRandomPost(post string) {
	rp := Random{}
	rjs, err := ioutil.ReadFile("../data/random.json")

	if err == nil {
		err = json.Unmarshal(rjs, &rp)
		log.Terminate(err, "../log/error.json")
	}
	d := Data{Post: post}
	datas := rp.PostData
	datas = append(datas, d)
	rp.PostData = datas
	bin, _ := json.Marshal(rp)
	ioutil.WriteFile("../data/random.json", bin, 0644)
}

func RandomPost() {
	rand.Seed(time.Now().UnixNano())
	c := LoadConfigure("")
	rp := Random{}
	rjs, err := ioutil.ReadFile("../data/random.json")
	log.Terminate(err, "../log/posterr.json")
	err = json.Unmarshal(rjs, &rp)
	log.Terminate(err, "../log/posterr.json")
	p := rand.Intn(len(rp.PostData))

	anaconda.SetConsumerKey(c.ConsumerKey)
	anaconda.SetConsumerSecret(c.ConsumerSelect)
	api := anaconda.NewTwitterApi(c.AccessToken, c.AccessTokenSelect)
	api.PostTweet(rp.PostData[p].Post, nil)

}
