package main

import (
   "fmt"
   "flag"
   "time"
   "net/http"
   "io"
   "github.com/hako/durafmt"
   "golang.org/x/sys/unix"
)

func main() {

   var port int
   flag.IntVar(&port, "port", 9090, "the port to listen on")
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

   io.WriteString(writer, durafmt.Parse(uptime).String())
}
