package main

import (
	"fmt"
	database "hue/database/add_group"
	"net/http"
	"strconv"
)

func HandleChangeColor(w http.ResponseWriter, r *http.Request, client *http.Client){
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

func HandleChangeBrightness(w http.ResponseWriter, r *http.Request, client *http.Client){
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

  changeBrigtness(id, brightnesInt, client)
}

func HandlePower(w http.ResponseWriter, r *http.Request, client *http.Client){
    id := r.PathValue("id")

    lamp := getLampData(id, client)
    isOn := lamp.On.On

    if !isOn == false {
      w.Write([]byte("Powered off"))
    } else {
      w.Write([]byte("Powered on"))
    }

    power(id, !isOn, client)
}

func HandleAddGroup(w http.ResponseWriter, r *http.Request, client *http.Client){
  database.Add_group()
}

