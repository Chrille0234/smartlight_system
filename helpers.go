package main

import (
	"bytes"
  "io"
	"encoding/json"
	"fmt"
	"net/http"
)

func changeBulbSetting(id string, client *http.Client, settings map[string]interface{}) {
  url := fmt.Sprintf("https://192.168.8.100/clip/v2/resource/light/%v", id)

  settingsJson, err := json.Marshal(settings);

  if err != nil {
    fmt.Printf("Error marshalling settings to json: %v", err)
  }

  req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(settingsJson))

  if err != nil {
    fmt.Println("Error creating request:", err)
  }

  req.Header.Set("Content-Type", "application/json")
  req.Header.Set("hue-application-key", key)


  resp, err := client.Do(req)
  if err != nil {
    fmt.Println("Error sending request:", err)
  }
  defer resp.Body.Close()
}

func changeColor(id string, hex string, client *http.Client){
  XY := getXYFromHex(hex)


  obj := map[string]interface{}{
    "color": map[string]interface{}{
      "xy": map[string]interface{}{
        "x": XY.X,
        "y": XY.Y,
      },
    },
  }

  changeBulbSetting(id, client, obj)
}

func getLampData(id string, client *http.Client) Light {
  url := fmt.Sprintf("https://192.168.8.100/clip/v2/resource/light/%v", id)

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

  var jsonData LightResponse
  err = json.Unmarshal(body, &jsonData)
  if err != nil {
      fmt.Println("error unmarshalling:", err)
  }

  if len(jsonData.Data) > 0 {
      return jsonData.Data[0]
  }

  return Light{}
}

func changeBrigtness(id string, brightness int, client *http.Client){
  obj := map[string]interface{}{
    "dimming": map[string]interface{}{
      "brightness": brightness,
    },
  }

  changeBulbSetting(id, client, obj)
}
