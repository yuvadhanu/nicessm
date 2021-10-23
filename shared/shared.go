package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"nicessm-api-service/config"
	"strings"
	"time"
)

//Shared : ""
type Shared struct {
	commandArgs   map[string]string
	Config        *config.ViperConfigReader
	PatymCheckSum *PatymCheckSum
}

//NewShared : Shared Factory
func NewShared(commandArgs map[string]string, configuration *config.ViperConfigReader) *Shared {
	return &Shared{commandArgs: commandArgs, Config: configuration, PatymCheckSum: new(PatymCheckSum)}
}

//GetCmdArg : ""
func (sh *Shared) GetCmdArg(key string) string {
	return sh.commandArgs[key]
}

//SplitCmdArguments : ""
func SplitCmdArguments(args []string) map[string]string {
	// fmt.Println(args)
	m := make(map[string]string)
	for _, v := range args {
		strs := strings.Split(v, "=")
		if len(strs) == 2 {
			m[strs[0]] = strs[1]
		} else {
			log.Println("not proper arguments", strs)
		}
	}
	// fmt.Print(args, m)
	return m
}

//GetTransactionID : ""
func (sh *Shared) GetTransactionID(charset string, length int) string {
	charset = "ABCDEFFHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

//Get : ""
func (sh *Shared) Get(url string, h map[string]string) (resp *http.Response, e error) {
	client := &http.Client{}
	req, err1 := http.NewRequest("GET", url, nil)
	if err1 != nil {
		return nil, err1
	}
	for key, val := range h {
		req.Header.Add(key, val)
	}
	return client.Do(req)
}

//Post : ""
func (sh *Shared) Post(url string, h map[string]string, postBody interface{}) (resp *http.Response, e error) {
	fmt.Println("url - ", url)
	sh.BsonToJSONPrintTag("body - ", postBody)
	client := &http.Client{}
	data, err := json.Marshal(postBody)
	if err != nil {
		return nil, errors.New("Errro in marshaling post data - " + err.Error())
	}
	fmt.Println("prepare buffer")
	body := bytes.NewBuffer(data)
	fmt.Println("posting")
	req, err1 := http.NewRequest("POST", url, body)
	if err1 != nil {
		return nil, err1
	}
	fmt.Println("headering")
	for key, val := range h {
		req.Header.Add(key, val)
	}
	fmt.Println("post doing")
	return client.Do(req)
}

// GetRandomOTP : returns random numeric string
// @param n length
func (sh *Shared) GetRandomOTP(n int) string {
	var x = []byte("0123456789")
	return genRandomStr(n, x)
}

func random(min, max int) int {
	// Sourse of new Random
	return rand.Intn(max-min) + min
}

func genRandomStr(n int, x []byte) string {
	var t = ""
	for i := 0; i < n; i++ {
		rr := random(0, len(x)-1)
		t += string(x[rr])
	}
	return t
}
