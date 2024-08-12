package database

import (
	"encoding/json"
	"fmt"
	"os"
)

type Group struct {
  Name string `json:"name"`
  Colors map[string]string `json:"colors"`
}

type Groups struct {
	Groups map[string]Group `json:"groups"`
}

func Add_group(){
  file, err := os.ReadFile("./database/add_group/groups.json")

  if err != nil {
    fmt.Printf("Error opening file: %v", err)
    return 
  }

  var fileJSON Groups
  err = json.Unmarshal(file, &fileJSON)

  if err != nil {
    fmt.Printf("Error converting: %v", err)
  }
  
}
