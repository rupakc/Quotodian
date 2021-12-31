package parse

import (
        "math/rand"
        "time"
        "strings"
      )

func QuoteParserFactory(quoteChannelName string, QuotesMap map[string]interface{}, QuotesMapList []map[string]interface{}) map[string]interface{} {
  if strings.ToLower(quoteChannelName) == "programming" {
      return parseProgrammingQuotes(QuotesMap)
  } else if strings.ToLower(quoteChannelName) == "garden" {
      return parseGardenQuotes(QuotesMap)
  } else if strings.ToLower(quoteChannelName) == "advice" {
      return parseAdviceQuotes(QuotesMap)
  } else if strings.ToLower(quoteChannelName) == "design" {
      return parseDesignQuotes(QuotesMapList)
  } else if strings.ToLower(quoteChannelName) == "favorite" {
      return parseFavoriteQuotes(QuotesMap)
  } else if strings.ToLower(quoteChannelName) == "taylor" {
      return parseTaylorSwiftQuotes(QuotesMap)
  } else if strings.ToLower(quoteChannelName) == "trump" {
      return parseDonaldTrumpQuotes(QuotesMap)
  } else if strings.ToLower(quoteChannelName) == "kanye" {
      return parseKanyeWestQuotes(QuotesMap)
  } else {
      return parseForismaticQuotes(QuotesMap)
  }
}

func parseProgrammingQuotes(QuotesMap map[string]interface{}) map[string]interface{} {
    parsedMap := make(map[string]interface{})
    parsedMap["author"] = QuotesMap["author"]
    parsedMap["quote"] = QuotesMap["en"]
    return parsedMap
}

func parseGardenQuotes(QuotesMap map[string]interface{}) map[string]interface{} {
    parsedMap := make(map[string]interface{})
    parsedMap["author"] = QuotesMap["quote"].(map[string]interface{})["quoteAuthor"]
    parsedMap["quote"] = QuotesMap["quote"].(map[string]interface{})["quoteText"]
    parsedMap["genre"] = QuotesMap["quote"].(map[string]interface{})["quoteGenre"]
    return parsedMap
}

func parseAdviceQuotes(QuotesMap map[string]interface{}) map[string]interface{} {
    parsedMap := make(map[string]interface{})
    parsedMap["quote"] = QuotesMap["slip"].(map[string]interface{})["advice"]
    return parsedMap
}

func parseDesignQuotes(QuotesMapList []map[string]interface{}) map[string]interface{} {
    source := rand.NewSource(time.Now().UnixNano())
    randomNumGenerator := rand.New(source)
    randomIndex := randomNumGenerator.Intn(len(QuotesMapList))
    randomQuotesMap := QuotesMapList[randomIndex]
    parsedMap := make(map[string]interface{})
    parsedMap["quote"] = randomQuotesMap["content"].(map[string]interface{})["rendered"]
    parsedMap["quote"] = strings.Replace(parsedMap["quote"].(string),"<p>","",-1)
    parsedMap["quote"] = strings.Replace(parsedMap["quote"].(string),"</p>","",-1)
    parsedMap["quote"] = strings.Replace(parsedMap["quote"].(string),"&#8217","",-1)
    parsedMap["quote"] = strings.Replace(parsedMap["quote"].(string),"&#8211","",-1)
    parsedMap["quote"] = strings.Replace(parsedMap["quote"].(string),"&#8243;","",-1)
    parsedMap["quote"] = strings.Replace(parsedMap["quote"].(string),"&#8220;","",-1)
    parsedMap["quote"] = strings.Replace(parsedMap["quote"].(string),"&#8221;","",-1)
    parsedMap["quote"] = strings.Replace(parsedMap["quote"].(string),"<strong>","",-1)
    parsedMap["quote"] = strings.Replace(parsedMap["quote"].(string),"</strong>","",-1)
    parsedMap["author"] = randomQuotesMap["slug"]
    return parsedMap
}

func parseFavoriteQuotes(QuotesMap map[string]interface{}) map[string]interface{} {
    parsedMap := make(map[string]interface{})
    parsedMap["quote"] = QuotesMap["quote"].(map[string]interface{})["body"]
    parsedMap["author"] = QuotesMap["quote"].(map[string]interface{})["author"]
    return parsedMap
}

func parseForismaticQuotes(QuotesMap map[string]interface{}) map[string]interface{} {
    parsedMap := make(map[string]interface{})
    parsedMap["quote"] = QuotesMap["quoteText"]
    parsedMap["author"] = QuotesMap["quoteAuthor"]
    return parsedMap
}

func parseTaylorSwiftQuotes(QuotesMap map[string]interface{}) map[string]interface{} {
    parsedMap := make(map[string]interface{})
    parsedMap["quote"] = QuotesMap["quote"]
    return parsedMap
}

func parseDonaldTrumpQuotes(QuotesMap map[string]interface{}) map[string]interface{} {
    parsedMap := make(map[string]interface{})
    parsedMap["quote"] = QuotesMap["value"]
    return parsedMap
}

func parseKanyeWestQuotes(QuotesMap map[string]interface{}) map[string]interface{} {
    parsedMap := make(map[string]interface{})
    parsedMap["quote"] = QuotesMap["quote"]
    return parsedMap
}
