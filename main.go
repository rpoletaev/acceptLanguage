package main

import (
  "fmt"
  //"context"
  "sort"
  "strings"
  "strconv"
  // "net/http"
  "golang.org/x/text/language"

)


type acceptLanguageKey struct{}

type AcceptLanguage struct {
  Lang string
  QFactor float32
}

type AcceptLanguages []AcceptLanguage

func(list AcceptLanguages) Len() int {
  return len(list)
}

func (list AcceptLanguages) Less(i,j int) bool {
  return list[i].QFactor < list[j].QFactor
}

func (list AcceptLanguages) Swap(i,j int) {
  list[i], list[j] = list[j], list[i]
}

func ParseLanguagesList(raw string) AcceptLanguages {
  rawList := strings.Split(raw, ",")
  list := AcceptLanguages(make([]AcceptLanguage, 0, len(rawList)))
  
  for _, rawLang := range rawList{
    list = append(list, ParseLanguage(rawLang)) 
  }

  sort.Sort(sort.Reverse(list))
  return list
}

func ParseLanguage(rawLang string) AcceptLanguage {
  parsed := strings.Split(rawLang, ";")
  l := AcceptLanguage{
    Lang: strings.TrimSpace(parsed[0]),
  }

  if len(parsed) > 1 {
    l.QFactor = parseQFactorString(parsed[1])
  } else {
    // returns as preferred in case: ru-RU,
    l.QFactor = 1.0
  }
  return l
}

func parseQFactorString(raw string) float32 {
  splited := strings.Split(raw, "=")
  
  if len(splited) == 1 {
    return 0.0
  }

  f, err := strconv.ParseFloat(strings.TrimSpace(splited[1]), 32)
  if err != nil {
    return 0.0
  }

  return float32(f)
}

// func LangToContext(req http.Request) context.Context {
  
// }


func main() {
  languages := ParseLanguagesList("ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
  fmt.Printf("%+v\n", languages)
  for _, l := range languages {
    fmt.Println(language.Parse(l.Lang))
  }
}