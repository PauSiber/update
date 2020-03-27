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
  "github.com/fatih/color"
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
    upgrade(jsonData.Updates[lastUpdate.Value+1:], false)
    clear()
    banner()
    fmt.Println("All updates are done.\nYour system is up to date.")
  } else {
    banner()
    fmt.Printf("Last update time: \"%v\" \n", lastUpdateTime)
    fmt.Printf("You have %v new updates.\n", jsonLastUpdate - lastUpdate.Value)
    fmt.Println("System is up to date.")
  }
}

func upgrade(updates []Update, codeFlag bool) {
  cancel:
  for i, update := range updates {
    clear()
    out:
    for {
      banner()
      r := color.New(color.FgRed).SprintFunc()
      fmt.Printf("Name: %v\nDescription: %v\nPublish Time: %v\n", r(update.Name),
                                                                  r(update.Description),
                                                                  r(update.PublishTime.Format("January 2, 2006, 15:04")))

      if codeFlag {
        showMeCode("updates/" + update.FileName)
      }
      if codeFlag {
        fmt.Printf("\nDo you want to run this upgration? [Y/N/S(close code)] ")
      } else {
        fmt.Printf("\nDo you want to run this upgration? [Y/N/S(show code)] ")
      }
      var answer string
      fmt.Scan(&answer)
      switch answer {
      case "s", "S":
        if codeFlag {
          codeFlag = false
        } else {
          codeFlag = true
        }
        upgrade(updates[i:], codeFlag)
        break cancel
      case "y", "Y":
        break out
      case "n", "N":
        fmt.Println("\nOkay. Update service is cancelled.")
        os.Exit(1)
      default:
        clear()
        fmt.Println("(!) Please use only Y, N or S")
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

  r := color.New(color.FgRed)

  r.Println("\n\t -----------")
  r.Println("\t|  Code:")
  r.Println("\t|")
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Println("\t"+ r.Sprintf("|") +"\t" + scanner.Text())
  }
  r.Println("\t|")
  r.Println("\t -----------")
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
  c := color.New(color.FgGreen)
  y := color.New(color.FgYellow).SprintFunc()
  fmt.Println("\n --------------------------------------------------------------------------")
  banner := `|    ____   _   _   _ ____ ___ ____  _____ ____   ` + y(`  ____  _______     __`) + "   |\n" +
            `|   |  _ \ / \ | | | / ___|_ _| __ )| ____|  _ \  ` + y(` |  _ \| ____\ \   / /`) + "   |\n" +
            `|   | |_) / _ \| | | \___ \| ||  _ \|  _| | |_) | ` + y(` | | | |  _|  \ \ / /`) + "    |\n" +
            `|   |  __/ ___ \ |_| |___) | || |_) | |___|  _ <  ` + y(` | |_| | |___  \ V /`) + "     |\n" +
            `|   |_| /_/   \_\___/|____/___|____/|_____|_| \_\ ` + y(` |____/|_____|  \_/`) + "      |\n|\t\t\t\t\t\t\t\t\t   |"
  fmt.Println(banner)
  fmt.Printf("|\t\t\t\t\t\t\t\t\t   |\n|")
  c.Printf("\t\t\t- Update Service %v - \t\t", version())
  fmt.Printf("\t   |\n")
  fmt.Println(" --------------------------------------------------------------------------\n")
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
  cmd = exec.Command("printf", "\\e[3J")
  cmd.Stdout = os.Stdout
  cmd.Run()
}
