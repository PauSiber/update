package main

import (
  "fmt"
  "log"
  "time"
  "os/exec"
  "io/ioutil"
  "encoding/json"
)

type JsonData struct {
  Authority string `json:"authority"`
  Name      string `json:"name"`
  Version   string `json:"version"`
  Updates   []Update `json:"updates"`
}

type Update struct {
  ID          string    `json:"id"`
  Description string    `json:"description"`
  FileName    string    `json:"fileName"`
  PublishTime time.Time `json:"publishTime"`
}

func main() {
  data, err := ioutil.ReadFile("./updates/updates.json")
  if err != nil {
    log.Fatal("Error occur while opening updates.json file.")
  }

  jsonData := JsonData{}
  json.Unmarshal(data, &jsonData)

  for _, update := range jsonData.Updates {
    out, err := exec.Command("/bin/bash", "./updates/" + update.FileName).Output()
    if err != nil {
        fmt.Printf("%s", err)
    }
    output := string(out[:])
    fmt.Print(output)
  }
}
