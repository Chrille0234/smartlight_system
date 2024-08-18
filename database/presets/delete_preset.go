package presets

import (
	"encoding/json"
	"fmt"
	"hue/types"
	"net/http"
	"os"
)

func Delete_preset(preset_name string, client *http.Client) {
  filePath := "./database/presets/presets.json"
  file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
  if err != nil {
    fmt.Printf("delete_preset.go: There was an error opening file: %v \n", err)
    return
  }
  defer file.Close()

  fileContent, err := os.ReadFile(filePath)
  if err != nil {
    fmt.Printf("delete_preset.go: There was an error reading file: %v", err)
    return
  }

  var fileContentUnmarshalled types.Presets
  err = json.Unmarshal(fileContent, &fileContentUnmarshalled)
  if err != nil {
    fmt.Printf("delete_preset.go: Error unmarshalling file: %v", err)
    return
  }

  delete(fileContentUnmarshalled.Presets, preset_name)

  fileMarshaled, err := json.MarshalIndent(fileContentUnmarshalled, "", "  ")
  if err != nil {
    fmt.Printf("delete_preset.go: Error MarshalIndenting new content: %v", err)
    return
  }

  if err := file.Truncate(0); err != nil {
    fmt.Printf("delete_preset.go: Error truncating file: %v", err)
    return
  }
  if _, err := file.Seek(0, 0); err != nil {
    fmt.Printf("delete_preset.go: Error seeking file: %v", err)
    return
  }
  if _, err := file.Write(fileMarshaled); err != nil {
    fmt.Printf("delete_preset.go: Error writing to file: %v", err)
    return
  }
}
