package main

import (
	Index "hue/views"
	"net/http"
	"github.com/a-h/templ"
)

var key string = "lwXMceW4jRVPJ5ti5o34CwYvpTgrP9QS39KAx3vq"
var productId = "2b2805f8-fee5-4c60-bcc7-f8706ed1968f"

func main() {

  // serve tailwindcss
  fs := http.FileServer(http.Dir("./static"))
  http.Handle("GET /static/", http.StripPrefix("/static/", fs))

  http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request){
    handler := templ.Handler(Index.Index())
    handler.ServeHTTP(w, r)
  })

  http.HandleFunc("GET /get_lamps", HandleGetLamps)
  http.HandleFunc("GET /get_lamp_names", HandleGetLampNames)
  http.HandleFunc("POST /power/{id}", HandlePower)
  http.HandleFunc("POST /change_color/{id}", HandleChangeColor)
  http.HandleFunc("POST /change_brightness/{id}", HandleChangeBrightness)
  http.HandleFunc("POST /change_preset/{preset_name}", HandleChangePreset)
  http.HandleFunc("GET /add_preset", HandleAddPreset)
  http.HandleFunc("GET /get_presets", HandleGetPresets)
  http.HandleFunc("GET /delete_preset/{preset_name}", HandleDeletePreset)
  http.HandleFunc("POST /add_group", HandleAddGroup)

  // I dont know if i even need this endpoint anymore....
//  http.HandleFunc("POST /lamp_name/{id}", func(w http.ResponseWriter, r *http.Request) {
//    lamp := getLampData(r.PathValue("id"), client)
//    message := fmt.Sprintf("Edit lamp %s", lamp.Metadata.Name)
//    w.Write([]byte(message))
//  })

  http.ListenAndServe(":3000", nil)
}
