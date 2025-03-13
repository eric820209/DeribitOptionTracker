package main

import (
  "fmt"
  "strings"
  "net/http"
  "io"
)

func main() {

  url := "https://derivatives-graphql.amberdata.com/graphql"
  method := "POST"

  payload := strings.NewReader(`{"operationName":"Level1Quotes","variables":{"symbol":"ETH","exchange":"deribit"},"query":"query Level1Quotes($symbol: SymbolEnumType, $exchange: ExchangeEnumType) {\n  Level1Quotes(currency: $symbol, exchange: $exchange) {\n    strike\n    expirationTimestamp\n    openInterest\n    openInterestUSD\n    putCall\n    instrument\n    __typename\n  }\n}"}`)

  client := &http.Client {
  }
  req, err := http.NewRequest(method, url, payload)

  if err != nil {
    fmt.Println(err)
    return
  }
  req.Header.Add("Content-Type", "application/json")

  res, err := client.Do(req)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer res.Body.Close()

  body, err := io.ReadAll(res.Body)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(string(body))
}