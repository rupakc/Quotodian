package handlers

import (
	"net/http"
  "encoding/json"
  "errors"
	"github.com/gorilla/mux"
	"strings"
  commonutils "backend/commonutils"
  commonconstants "backend/constants"
	quotesParser "backend/parse"
)

func GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
  vars := mux.Vars(r)
  sourceName := vars["sourceName"]
  urlForSource,keyPresent := commonconstants.UrlMap[sourceName]
  if !keyPresent {
    panic(errors.New("source name not valid or supported"))
  }
  jsonResponse := commonutils.GetHttpJsonResponseFromURL(urlForSource, nil)
	parsedJsonResponse := make(map[string]interface{})
	if !strings.Contains(sourceName, "design") && !strings.Contains(sourceName, "namo") {
		parsedJsonResponse = quotesParser.QuoteParserFactory(sourceName, jsonResponse.(map[string]interface{}), nil)
	} else {
		parsedJsonResponse = quotesParser.QuoteParserFactory(sourceName, nil, jsonResponse.([]map[string]interface{}))
	}
  if err := json.NewEncoder(w).Encode(parsedJsonResponse); err != nil {
		panic(err)
	}
}
