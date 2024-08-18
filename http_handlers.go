package main

import (
	"crypto/tls"
	"fmt"
	"hue/database/groups"
	presets "hue/database/presets"
	Index "hue/views"
	"net/http"
	"strconv"
	"strings"

	"github.com/a-h/templ"
)

var tr *http.Transport = &http.Transport{
  TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
}
var client *http.Client = &http.Client{Transport: tr}

func HandleAddPreset(w http.ResponseWriter, r *http.Request){
  if err := r.ParseForm(); err != nil {
    fmt.Printf("Error parsing form: %v", err)
  }

  preset_name := r.FormValue("preset_name")
 
  preset,_ := presets.Get_preset(preset_name)

  

  presets.Add_preset(preset_name, client, key)

  if len(preset.Colors) == 0 {
    templhandler := templ.Handler(Index.Preset(preset_name))
    templhandler.ServeHTTP(w, r)
  }
}

func HandleGetPresets(w http.ResponseWriter, r *http.Request){
  presets, err := presets.Get_presets()

  if err != nil {
    fmt.Printf("Error getting all presets: %v", err)
    return
  }

  for key := range presets.Presets {
    templhandler := templ.Handler(Index.Preset(key))
    templhandler.ServeHTTP(w, r)
  }
}

func HandleChangePreset(w http.ResponseWriter, r *http.Request){
  preset,_ := presets.Get_preset(r.PathValue("preset_name"))

  for id, color := range preset.Colors {
    color = strings.Trim(color, "{}")
    colorSplit := strings.Split(color, " ")

    x, err := strconv.ParseFloat(colorSplit[0], 64)

    if err != nil {
      fmt.Printf("Could not convert x to float64: %v", err)
    }

    y, err := strconv.ParseFloat(colorSplit[1], 64)

    if err != nil {
      fmt.Printf("Could not convert y to float64: %v", err)
    }

    XY := XY{
      X: x,
      Y: y,
    }

    changeColorWithIdArg(id, XY, client)
  }

}

func HandleChangeColor(w http.ResponseWriter, r *http.Request){
    id := r.PathValue("id")

    err := r.ParseForm();

    if err != nil {
      fmt.Printf("Error parsing form: %v", err)
    }

    hex := r.FormValue("color")

    if hex == ""{
      fmt.Println("Cannot find parameter 'color'")
    }

    changeColor(id, hex, client)
}

func HandleChangeBrightness(w http.ResponseWriter, r *http.Request){
  id := r.PathValue("id")
  err := r.ParseForm()

  if err != nil {
    fmt.Println("Error parsing form:", err)
  }

  brightness := r.FormValue("brightness")

  if brightness == "" {
    fmt.Println("A brightness is required")
  }

  brightnesInt, err := strconv.Atoi(brightness)

  if err != nil {
    fmt.Printf("Couldnt convert brightness to an int: %v", err)
  }
  changeBrightness(id, brightnesInt, client)
}

func HandleDeletePreset(w http.ResponseWriter, r *http.Request){
  preset_name := r.PathValue("preset_name")
  presets.Delete_preset(preset_name, client)
}

func HandleGetLampNames(w http.ResponseWriter, r *http.Request){
  allLamps := getAllLampsData(client)


  for _, lamp := range allLamps {
    inputName := fmt.Sprintf("%s:%s", lamp.ID, lamp.Metadata.Name)
    templhandler := templ.Handler(Index.GroupLabel(lamp.Metadata.Name, inputName))
    templhandler.ServeHTTP(w, r)
  }
}

func HandleAddGroup(w http.ResponseWriter, r *http.Request){
  err := r.ParseForm()

  if err != nil {
    fmt.Printf("http_handler.go: Error parsing form: %v", err)
  }

  lampsToAdd := make(map[string]string)

  for name, status := range r.Form {
    if name == "group_name" { continue }
    nameSplit := strings.Split(name, ":")
    id := nameSplit[0]
    lampName := nameSplit[1]

    if status[0] != "on" { continue }

    lampsToAdd[id] = lampName
  }

  groups.Add_group(r.FormValue("group_name"), lampsToAdd)

}

func HandlePower(w http.ResponseWriter, r *http.Request){
    id := r.PathValue("id")

    lamp := getLampData(id, client)
    isOn := lamp.On.On

    if !isOn {
      w.Write([]byte("Powered off"))
    } else {
      w.Write([]byte("Powered on"))
    }

   data := map[string]interface{}{
    "on": map[string]interface{}{
      "on": !isOn,
    },
  }

  changeBulbSetting(id, client, data)
}


func HandleGetLamps(w http.ResponseWriter, r *http.Request){
  allLamps := getAllLampsData(client)

  for _, lamp := range allLamps {
    XY := XY {
      X: lamp.Color.XY.X,
      Y: lamp.Color.XY.Y,
    }
    hex := XY.ToHex()
    brightness := fmt.Sprint(lamp.Dimming.Brightness)

    templHandler := templ.Handler(Index.Lamp(lamp.ID, lamp.Metadata.Name, hex, brightness))
    templHandler.ServeHTTP(w, r)
  }
}
