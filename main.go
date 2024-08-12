package main

import (
	"crypto/tls"
	"fmt"
	Index "hue/views"
	"net/http"
	"github.com/a-h/templ"
)

var key string = "lwXMceW4jRVPJ5ti5o34CwYvpTgrP9QS39KAx3vq"
var productId = "2b2805f8-fee5-4c60-bcc7-f8706ed1968f"

func main() {
  tr := &http.Transport{
    TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
  }

  client := &http.Client{Transport: tr}

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    handler := templ.Handler(Index.Index())
    handler.ServeHTTP(w, r)
  })

  http.HandleFunc("POST /power/{id}", func(w http.ResponseWriter, r *http.Request) {
    HandlePower(w, r, client)
  })

  http.HandleFunc("POST /change_color/{id}", func(w http.ResponseWriter, r *http.Request) {
    HandleChangeColor(w, r, client)
  })

  http.HandleFunc("POST /change_brightness/{id}", func(w http.ResponseWriter, r *http.Request) {
    HandleChangeBrightness(w, r, client)
  })

  http.HandleFunc("/lamp_name/{id}", func(w http.ResponseWriter, r *http.Request) {
    lamp := getLampData(r.PathValue("id"), client)
    message := fmt.Sprintf("Edit lamp %s", lamp.Metadata.Name)
    w.Write([]byte(message))
  })

  http.HandleFunc("GET /add_group/{name}", func(w http.ResponseWriter, r *http.Request) {
    HandleAddGroup(w, r, client)
  })


  http.ListenAndServe(":3000", nil)
}


func power(id string, powerMode bool, client *http.Client){
  data := map[string]interface{}{
    "on": map[string]interface{}{
      "on": powerMode,
    },
  }

  changeBulbSetting(id, client, data)
}



