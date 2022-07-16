package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"

	"github.com/denisbrodbeck/machineid"
	"github.com/hpcloud/tail"
	//"github.com/mdp/qrterminal/v3"
)

var EventNumber int

func PostResults(MapName string, QueueTime int) {
	EventNumber++
	fmt.Println("[#] Event number:", EventNumber)
	fmt.Println("  [i] Map:", MapName)
	fmt.Println("  [i] Sec:", QueueTime)
	fmt.Println()
}

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned iff there is an error with the
// buffered reader.
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func Map(line string) string {
	var map_name string

	// Format the map name
	map_name = strings.TrimSpace(strings.Split(strings.Split(line, "Location:")[1], ",")[0])
	if map_name == "factory4_day" {
		map_name = "factory"
	}
	if map_name == "RezervBase" {
		map_name = "reserve"
	}
	if map_name == "bigmap" {
		map_name = "customs"
	}

	map_name = strings.ToLower(map_name)

	return map_name
}

func Queue(line string) int {
	var queue_time_raw string

	// Format the queue time
	queue_time_raw = strings.TrimSpace(strings.Split(strings.Split(line, "real:")[1], " ")[0])

	queue_time_float, _ := strconv.ParseFloat(queue_time_raw, 64)
	queue_time := int(queue_time_float)

	return queue_time
}

func main() {
	id, err := machineid.ProtectedID("tarkov")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)

	//qrterminal.Generate(id, qrterminal.L, os.Stdout)

	// Get the current user's home directory
	currentUser, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	homedir := currentUser.HomeDir

	// Format the player log file
	player_log_file := homedir + "\\AppData\\LocalLow\\Battlestate Games\\EscapeFromTarkov\\Player.log"

	// Read the Player.log config file to get the game install path
	install_search_prefix := "Mono path[0] = "
	var install_path string
	f, err := os.Open(player_log_file)
	if err != nil {
		fmt.Printf("error opening file Player.log: %v\n", err)
		//os.Exit(1)
	}
	r := bufio.NewReader(f)
	s, e := Readln(r)
	for e == nil {
		// Search each line for the install_search_prefix
		if strings.Contains(s, install_search_prefix) {
			install_path = strings.Replace(strings.Replace(strings.TrimSpace(strings.Replace(strings.Split(s, install_search_prefix)[1], "'", "", -1)), "EscapeFromTarkov_Data/Managed", "", -1), "/", "\\", -1)
			break
		}
		s, e = Readln(r)
	}
	fmt.Println("[#] EFT Install Path: " + install_path)

	// Log Directory
	log_dir := install_path + "\\Logs"

	// Scan the log directory for all log folders
	files, err := ioutil.ReadDir(log_dir)
	if err != nil {
		//log.Fatal(err)
	}

	// Grab the last log folder in the list which will be the latest
	var latest_log_dir string
	for _, f := range files {
		latest_log_dir = log_dir + "\\" + f.Name()
	}
	fmt.Println("[#] Latest Log Directory: " + latest_log_dir)
	// Loop through all files in the log directory
	log_files, err := ioutil.ReadDir(latest_log_dir)
	if err != nil {
		//log.Fatal(err)
	}

	// Grab the last log folder in the list which will be the latest
	var app_log_file string
	for _, f := range log_files {
		if strings.Contains(f.Name(), "application") {
			app_log_file = latest_log_dir + "\\" + f.Name()
		}
	}
	fmt.Println("[#] Application Log File: " + app_log_file)

	fmt.Println("[#] Starting queue-scanner\n")
	time.Sleep(2 * time.Second)

	var QueueTime int
	// Tail the application log file
	t, _ := tail.TailFile(app_log_file, tail.Config{Follow: true, Poll: true})
	for line := range t.Lines {

		if strings.Contains(line.Text, "GamePrepared") {
			QueueTime = Queue(line.Text)
		}
		if strings.Contains(line.Text, "RaidMode: Online") {
			//MapName = Map(line.Text)
			PostResults(Map(line.Text), QueueTime)
		}
	}

	fmt.Println("[#] Exiting queue-scanner")
}
