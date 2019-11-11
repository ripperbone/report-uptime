package main

import (
   "fmt"
   "flag"
   "net/http"
   "os"
   "encoding/json"
   "github.com/shirou/gopsutil/host"
)

type Response struct {
   Hostname string `json:"hostname"`
   Uptime uint64 `json:"uptime"`
   UptimeString string `json:"uptime_string"`
}

func main() {

   var port int
   flag.IntVar(&port, "port", 9095, "the port to listen on")
   flag.Parse()

   http.HandleFunc("/", getUptime)
   fmt.Printf("Listening on port: %d...\n", port)
   http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func getUptime(writer http.ResponseWriter, request *http.Request) {

   hostname, err := os.Hostname()

   if err != nil {
      fmt.Errorf("Error getting hostname: %v", err)
   }

   uptime, err := host.Uptime()

   if err != nil {
      fmt.Errorf("Error getting uptime: %v", err)
   }

   response := Response{ Hostname: hostname, Uptime: uptime, UptimeString: formatUptime(uptime) }
   data, err := json.Marshal(response)

   if err != nil {
      fmt.Errorf("Error forming response: %v", err)
   }

   writer.Header().Set("Content-Type", "application/json")
   writer.Write(data)
}

func formatUptime(uptime uint64) string {
   days := uptime / 86400
   hours := (uptime % 86400) / 3600
   minutes := (uptime % 3600) / 60
   return fmt.Sprintf("%d days %d hours %d minutes", days, hours, minutes)
}
