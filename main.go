package main

import (
  "fmt"
  "log"
  "time"
  "os"
  "os/exec"
  "strings"
  "bufio"
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
  ID          int       `json:"id"`
  Name        string    `json:"name"`
  Description string    `json:"description"`
  FileName    string    `json:"fileName"`
  PublishTime time.Time `json:"publishTime"`
}

type LastUpdate struct {
	Value  int       `json:"update"`
	Time   time.Time `json:"time"`
}

func main() {
  var jsonData JsonData
  var lastUpdate LastUpdate
  readJsonData(&jsonData)
  readLastUpdate(&lastUpdate)
  update(jsonData, lastUpdate)
}

func update(jsonData JsonData, lastUpdate LastUpdate) {
  clear()
  jsonLastUpdate := -1
  if len(jsonData.Updates) > 0 {
    jsonLastUpdate = jsonData.Updates[len(jsonData.Updates) - 1].ID
  }

  lastUpdateTime := "N/A"
  if lastUpdate.Value >= 0 {
    lastUpdateTime = lastUpdate.Time.Format("January 2, 2006, 15:04")
  }

  if lastUpdate.Value < jsonLastUpdate {
    clear()
    out:
    for {
      banner()
      fmt.Printf("Last update time: \"%v\" \n", lastUpdateTime)
      fmt.Printf("You have %v new updates.\n", jsonLastUpdate - lastUpdate.Value)
      fmt.Printf("\nDo you want to upgrade your system? [Y/N] ")
      var answer string
      fmt.Scan(&answer)
      switch answer {
      case "y", "Y":
        break out
      case "n", "N":
        fmt.Println("\nOkay. Update service is cancelled.")
        os.Exit(1)
      default:
        clear()
        fmt.Println("(!) Please use only Y or N")
      }
    }
    upgrade(jsonData.Updates[lastUpdate.Value+1:])
    clear()
    banner()
    fmt.Println("All updates are done.\nYour system is up to date.")
  } else {
    banner()
    fmt.Printf("Last update time: \"%v\" \n", lastUpdate.Time.Format("January 2, 2006, 15:04"))
    fmt.Printf("You have %v new updates.\n", jsonLastUpdate - lastUpdate.Value)
    fmt.Println("System is up to date.")
  }
}

func upgrade(updates []Update) {
  for _, update := range updates {
    clear()
    out:
    for {
      banner()
      fmt.Printf("Name: %v\nDescription: %v\nPublish Time: %v\n\n", update.Name,
                                                                  update.Description,
                                                                  update.PublishTime.Format("January 2, 2006, 15:04"))

      showMeCode("updates/" + update.FileName)
      fmt.Printf("\nDo you want to run this upgration? [Y/N] ")
      var answer string
      fmt.Scan(&answer)
      switch answer {
      case "y", "Y":
        break out
      case "n", "N":
        fmt.Println("\nOkay. Update service is cancelled.")
        os.Exit(1)
      default:
        clear()
        fmt.Println("(!) Please use only Y or N")
      }
    }

    fmt.Printf("\nUpgration is started.\n")
    fmt.Printf("--------------------------------------\n")
    cmd := exec.Command("sudo", "/bin/bash", "updates/" + update.FileName )
    cmd.Stderr = os.Stderr
    cmd.Stdout = os.Stdout
    cmd.Stdin = os.Stdin
    cmd.Run()
    fmt.Printf("\n--------------------------------------\n")
    fmt.Println("Running upgration is completed.")
    writeLastUpdate(update.ID)
    fmt.Printf("\n[ Push enter to continue ] ")
    fmt.Scanln()
  }
}

func showMeCode(path string) {
  file, err := os.Open(path)
  if err != nil {
    log.Fatal("Error occur")
  }
  defer file.Close()

  fmt.Println("\t -----------")
  fmt.Println("\t|  Code:")
  fmt.Println("\t|")
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Println("\t|\t" + scanner.Text())
  }
  fmt.Println("\t|")
  fmt.Println("\t -----------")
}

func writeLastUpdate(id int) {
  lastUpdate := &LastUpdate{
    Value: id,
    Time: time.Now(),
  }
  lastUpdateJson, _ := json.MarshalIndent(lastUpdate, "", "  ")
  ioutil.WriteFile("./updates/last_update.json", lastUpdateJson, 0644)
}

func readLastUpdate(lastUpdate *LastUpdate) {
  data, err := ioutil.ReadFile("./updates/last_update.json")
  if err != nil {
    // Creates json file if it do not exist.
    // lastUpdate.Value = -1
    lastUpdate.Time = time.Now()
    lastUpdateJson, _ := json.MarshalIndent(lastUpdate, "", "  ")
    ioutil.WriteFile("./updates/last_update.json", lastUpdateJson, 0644)
  } else {
    json.Unmarshal(data, lastUpdate)
  }
}

func readJsonData(jsonData *JsonData) {
  data, err := ioutil.ReadFile("./updates/updates.json")
  if err != nil {
    log.Fatal("Error occur while reading \"updates.json\" file.")
  }

  json.Unmarshal(data, jsonData)
}

func banner() {
  banner := `  ____   _   _   _ ____ ___ ____  _____ ____       ____  _______     __` + "\n" +
            ` |  _ \ / \ | | | / ___|_ _| __ )| ____|  _ \     |  _ \| ____\ \   / /` + "\n" +
            ` | |_) / _ \| | | \___ \| ||  _ \|  _| | |_) |    | | | |  _|  \ \ / /` + "\n" +
            ` |  __/ ___ \ |_| |___) | || |_) | |___|  _ <     | |_| | |___  \ V /` + "\n" +
            ` |_| /_/   \_\___/|____/___|____/|_____|_| \_\    |____/|_____|  \_/` + "\n"
  fmt.Println(banner)
  fmt.Printf("\t\t\t- Update Service %v - \n\n", version())
}

func version() string {
  data, err := ioutil.ReadFile("version.txt")
  if err != nil {
    return ""
  } else {
    return strings.TrimSpace(string(data))
  }
}

func clear() {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()
}
