package main

import (
	"encoding/json"
	"github.com/c-bata/go-prompt"
	"strconv"
	"strings"
	//"flag"
	"fmt"
)

type Data uint

const (
	MapRoundWins Data = iota
	Map
	PlayerID
	PlayerMatchStats
	PlayerState
	PlayerWeapons
	Provider
	Round

	// Below this line must be spectating or observing

	AllGrenades
	AllPlayersID
	AllPlayersMatchStats
	AllPlayersPosition
	AllPlayersState
	AllPlayersWeapons
	Bomb
	PhaseCountdowns
	PlayerPosition
)

// TODO : Support non-verbose mode

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

type Config struct {
	Name      string            `json:"-"`
	URI       string            `json:"uri"`
	Timeout   float64           `json:"timeout"`
	Buffer    float64           `json:"buffer"`
	Throttle  float64           `json:"throttle"`
	Heartbeat float64           `json:"heartbeat"`
	Auth      map[string]string `json:"auth"`
	Output    map[string]uint   `json:"output"`
	Data      map[string]uint   `json:"data"`
}

func GetStringFromInput(input, prefix, def string) string {
	fmt.Printf("Please Input your %s...(DEFAULT:\"%s\")\n", input, def)
	in := prompt.Input(prefix, completer)
	if in == "" {
		return def
	}
	return in
}

func GetIntFromInput(input string, prefix string, def int) int {
	fmt.Printf("Please Input your %s...(DEFAULT:\"%d\")\n", input, def)
	for {
		in := prompt.Input(prefix, completer)
		if in == "" {
			return def
		}
		res, err := strconv.Atoi(in)
		if err != nil {
			fmt.Println("Please Input numeric value!!")
			continue
		}
		return res
	}
}

func GetFloatFromInput(input string, prefix string, def float64) float64 {
	fmt.Printf("Please Input your %s...(DEFAULT:\"%f\")\n", input, def)
	for {
		in := prompt.Input(prefix, completer)
		if in == "" {
			return def
		}
		res, err := strconv.ParseFloat(in, 64)
		if err != nil {
			fmt.Println("Please Input numeric value!!")
			continue
		}
		return res
	}
}

func main() {
	prefix := "> "
	cfg := Config{}
	fmt.Println("Welcome to GSI cfg Generator!")
	fmt.Println("Developed by Shugo \"FlowingSPDG\" Kawamura")

	// Name
	cfg.Name = GetStringFromInput("Name", prefix, "GameStateIntegration Auto Generated")

	// URI
	cfg.URI = GetStringFromInput("URL Endpoint", prefix, "http://localhost:3090")

	// Timeout
	cfg.Timeout = GetFloatFromInput("Timeout", prefix, 5.0)

	// Buffer
	cfg.Buffer = GetFloatFromInput("Buffer", prefix, 0.1)

	// Throttle
	cfg.Throttle = GetFloatFromInput("Throttle", prefix, 0.1)

	// Heartbeat
	cfg.Heartbeat = GetFloatFromInput("Heartbeat", prefix, 30.0)

	// Auth
	cfg.Auth = make(map[string]string)

	// Output
	cfg.Output = make(map[string]uint)

	// Data
	cfg.Data = make(map[string]uint)

	jsonobj, err := json.MarshalIndent(&cfg, "", "    ")
	if err != nil {
		fmt.Printf("Something went wrong!\nERR : %v\n", err)
		return
	}
	jsonstr := strings.ReplaceAll(fmt.Sprintf("\"%s\"%v", cfg.Name, string(jsonobj)), ",", "")

	fmt.Printf("JSON OUTPUT...\n%v\n", jsonstr)
}
