package gostuff

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"

	"golang.org/x/text/language"
)

type IPLocation struct {
	IP           string
	Country_code string
	Country_name string
	Region_code  string
	Region_name  string
	City         string
	Zip_code     string
	Time_zone    string
	Latitude     float32
	Longitude    float32
	Metro_code   int
}

func GetCountry(w http.ResponseWriter, r *http.Request) {

	ipAddress, _, _ := net.SplitHostPort(r.RemoteAddr)
	//fmt.Println(ipAddress)
	response, err := http.Get("http://freegeoip.net/json/" + ipAddress)
	defer response.Body.Close()
	if err != nil {
		fmt.Println("error in get language", err)
		// default to globe
		w.Write([]byte("globe"))
		return
	}
	htmlData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		w.Write([]byte("globe"))
		return
	}

	//fmt.Println(string(htmlData))

	var ipLocation IPLocation

	if err := json.Unmarshal(htmlData, &ipLocation); err != nil {
		fmt.Println("Just receieved a message I couldn't decode:", string(htmlData), err)
	}
	w.Write([]byte(strings.ToLower(ipLocation.Country_code)))
}

func getLocale(w http.ResponseWriter, r *http.Request) {
	var matcher = language.NewMatcher([]language.Tag{
		language.BritishEnglish,
		language.Norwegian,
		language.German,
		language.English,
	})

	t, _, _ := language.ParseAcceptLanguage(r.Header.Get("Accept-Language"))
	// We ignore the error: the default language will be selected for t == nil.
	tag, _, _ := matcher.Match(t...)
	fmt.Println(tag)
}
