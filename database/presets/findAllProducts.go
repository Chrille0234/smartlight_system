package presets

import (
	"encoding/json"
	"fmt"
	"hue/types"
	"io"
	"net/http"
)


func getAllLampsData(client *http.Client, key string) []types.Light{
  url := fmt.Sprintf("https://192.168.8.100/clip/v2/resource/light")

  req, err := http.NewRequest(http.MethodGet, url, nil)
  if err != nil {
      fmt.Println("error making new request:", err)
  }

  req.Header.Add("hue-application-key", key)

  res, err := client.Do(req)
  if err != nil {
      fmt.Println("error doing request:", err)
  }

  body, err := io.ReadAll(res.Body)
  if err != nil {
      fmt.Println("Error reading body:", err)
  }

  var jsonData types.LightResponse
  err = json.Unmarshal(body, &jsonData)
  if err != nil {
      fmt.Println("error unmarshalling:", err)
  }

  if len(jsonData.Data) > 0 {
      return jsonData.Data
  }

  return []types.Light{}

}
