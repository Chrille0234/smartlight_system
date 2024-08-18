package presets

import (
	"encoding/json"
	"fmt"
	"hue/types"
	"os"
)

func Get_presets() (types.Presets, error) {
  filePath := "./database/presets/presets.json" 
  file, err := os.OpenFile(filePath, os.O_RDWR, 0644)

  defer file.Close()

  if err != nil {
    err = fmt.Errorf("Error opening presets.json: %v", err)
    return types.Presets{}, err
  }

  fileContent, err := os.ReadFile(filePath)

  if err != nil {
    err = fmt.Errorf("Error reading presets.json: %v", err)
    return types.Presets{}, err
  }

  var fileJSON types.Presets
  err = json.Unmarshal(fileContent, &fileJSON)

  if err != nil {
    err = fmt.Errorf("Error unmarshalling json to fileJSON: %v", err)
    return types.Presets{}, err
  }

  return fileJSON, nil
}

func Get_preset(preset_name string) (types.Preset, error) {
  if preset_name == "" {
    return types.Preset{}, fmt.Errorf("The preset name cannot be an empty string")
  }

  filePath := "./database/add_preset/presets.json" 
  file, err := os.OpenFile(filePath, os.O_RDWR, 0644)

  defer file.Close()

  if err != nil {
    err := fmt.Errorf("Error opening file: %v", err)
    return types.Preset{}, err
  }

  fileContent, err := os.ReadFile(filePath)

  if err != nil {
    err := fmt.Errorf("Error reading file: %v", err)
    return types.Preset{}, err
  }

  var fileJSON types.Presets
  err = json.Unmarshal(fileContent, &fileJSON)

  if err != nil {
    err := fmt.Errorf("Error converting: %v", err)
    return types.Preset{}, err
  }

  return fileJSON.Presets[preset_name], nil
  
}
