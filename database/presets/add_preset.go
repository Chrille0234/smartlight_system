package presets

import (
  "encoding/json"
  "fmt"
  "net/http"
  "os"
  "hue/types"
)

func Add_preset(preset_name string, client *http.Client, key string){
  fileName := "./database/presets/presets.json" 
  file, err := os.OpenFile(fileName, os.O_RDWR, 0644)

  defer file.Close()

  if err != nil {
    fmt.Printf("Error opening file: %v", err)
    return 
  }

  fileContent, err := os.ReadFile(fileName)

  if err != nil {
    fmt.Println("Error reading file:", err)
  }

  var fileJSON types.Presets
  err = json.Unmarshal(fileContent, &fileJSON)

  if err != nil {
    fmt.Printf("Error converting: %v", err)
  }

  var newColors map[string]string = make(map[string]string)

  allLamps := getAllLampsData(client, key)
  for _, val := range allLamps {
    newColors[val.ID] = fmt.Sprintf("%v", val.Color.XY)
  }


  fileJSON.Presets[preset_name] = types.Preset{
    Colors: newColors,
  }

  fileMarshaled, err := json.MarshalIndent(fileJSON, "", "  ")

  file.Truncate(0)
  file.Write(fileMarshaled)
}
