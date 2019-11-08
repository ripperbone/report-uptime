package main

import (
   "fmt"
   "flag"
   "time"
   "net/http"
   "os"
   "encoding/json"
   "github.com/hako/durafmt"
   "golang.org/x/sys/unix"
)

type Response struct {
   Hostname string `json:"hostname"`
   Uptime string `json:"uptime"`
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
   var sysinfo unix.Sysinfo_t

   var err error = unix.Sysinfo(&sysinfo)

   if err != nil {
      fmt.Errorf("Error reading sysinfo: %v", err)
   }

   // sysinfo.Uptime is seconds since boot. Convert to nanoseconds
   var uptime time.Duration = time.Duration(sysinfo.Uptime * 1e9)
   hostname, _ := os.Hostname()

   response := Response{ Hostname: hostname, Uptime: durafmt.Parse(uptime).String() }
   data, err := json.Marshal(response)

   if err != nil {
      fmt.Errorf("Error forming response: %v", err)
   }
   writer.Header().Set("Content-Type", "application/json")
   writer.Write(data)
}
