package main

import (
   "fmt"
   "flag"
   "time"
   "github.com/hako/durafmt"
   "golang.org/x/sys/unix"
)

func main() {

   var url string

   flag.StringVar(&url, "url", "", "the url where to send the report")

   flag.Parse()

   if len(url) > 0 {
      fmt.Println("url:", url)
   }

   var sysinfo unix.Sysinfo_t

   var err error = unix.Sysinfo(&sysinfo)

   if err != nil {
      fmt.Errorf("Error reading sysinfo: %v", err)
   }

   // sysinfo.Uptime is seconds since boot. Convert to nanoseconds
   var uptime time.Duration = time.Duration(sysinfo.Uptime * 1e9)

   fmt.Println(durafmt.Parse(uptime))
}
