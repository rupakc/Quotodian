package commonutils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	s "strings"
)

func GetHttpJsonResponseFromURL(urlString string, queryParamMap map[string]string) interface{} {

	httpClient := http.Client{
		Timeout: time.Second * 10, // Timeout after 2 seconds
	}

	request, error := http.NewRequest(http.MethodGet, urlString, nil)
	if error != nil {
		log.Fatal(error)
	}

  query := request.URL.Query()

  if queryParamMap != nil {
    for key, value := range queryParamMap {
      query.Add(key,value)
    }
  }

  request.URL.RawQuery = query.Encode()
	request.Header.Set("User-Agent", "quotes-aggregator")
  request.Header.Set("Accept", "application/json")

	response, getError := httpClient.Do(request)

  if getError != nil {
		log.Fatal(getError)
	}

	if response.Body != nil {
		defer response.Body.Close()
	}

	body, readError := ioutil.ReadAll(response.Body)

	if readError != nil {
		log.Fatal(readError)
	}
  var resultJson map[string]interface{}
  var alternateResultJson []map[string]interface{}

	if !s.Contains(urlString,"ondesign") && !s.Contains(urlString,"namo-memes") {
    _ = json.Unmarshal([]byte(body), &resultJson)
    return resultJson
	} else {
    _ = json.Unmarshal([]byte(body), &alternateResultJson)
    return alternateResultJson
  }
}
