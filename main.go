package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Level1Quote 代表 API 回應中的單筆數據
type Level1Quote struct {
	Strike              float64 `json:"strike"`
	ExpirationTimestamp string  `json:"expirationTimestamp"`
	OpenInterest        int     `json:"openInterest"`
	OpenInterestUSD     float64 `json:"openInterestUSD"`
	PutCall             string  `json:"putCall"`
	Instrument          string  `json:"instrument"`
}

// APIResponse 代表完整的 API 回應
type APIResponse struct {
	Data struct {
		Level1Quotes []Level1Quote `json:"Level1Quotes"`
	} `json:"data"`
}

func main() {

	url := "https://derivatives-graphql.amberdata.com/graphql"
	method := "POST"

	payload := strings.NewReader(`{"operationName":"Level1Quotes","variables":{"symbol":"ETH","exchange":"deribit"},"query":"query Level1Quotes($symbol: SymbolEnumType, $exchange: ExchangeEnumType) {\n  Level1Quotes(currency: $symbol, exchange: $exchange) {\n    strike\n    expirationTimestamp\n    openInterest\n    openInterestUSD\n    putCall\n    instrument\n    __typename\n  }\n}"}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		LogError(err) // 記錄錯誤
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		LogError(err) // 記錄錯誤
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		LogError(err) // 記錄錯誤
		return
	}

	// 假設這裡要解析 JSON
	var response APIResponse
	if err := json.Unmarshal(body, &response); err != nil {
		LogError(err) // 記錄錯誤
		return
	}

	fmt.Println("✅ API 呼叫成功")
}

// LogError 將錯誤寫入 error.log
func LogError(err error) {
	if err == nil {
		return
	}

	// 開啟（或創建）error.log，並以追加模式寫入
	f, fileErr := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if fileErr != nil {
		fmt.Println("無法開啟錯誤日誌:", fileErr)
		return
	}
	defer f.Close()

	// 格式化錯誤訊息
	logMessage := fmt.Sprintf("[%s] %v\n", time.Now().Format("2006-01-02 15:04:05"), err)

	// 寫入日誌
	_, writeErr := f.WriteString(logMessage)
	if writeErr != nil {
		fmt.Println("無法寫入錯誤日誌:", writeErr)
	}
}
