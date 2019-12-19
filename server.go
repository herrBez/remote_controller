/*
 * author: Mirko Bez (c) 2019
 *
 * description: simple web server, that allows to send commands to the os such
 * as increase and decreasing the volume, of your laptop
 *
 */
package main

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "github.com/herrBez/remote_controller/platform_specific"
)

// Code for the welcome page, simply load the static web page
func sayHello(w http.ResponseWriter, r *http.Request) {
  message, err := ioutil.ReadFile("./html/index.html")
  if err != nil {
    panic(err)
  }
  w.Write([]byte(message))
}

func increase (w http.ResponseWriter, r *http.Request) {
  fmt.Println("Increase Volume")
  go platform_specific.IncreaseVolume(2)
}

func decrease (w http.ResponseWriter, r *http.Request) {
  fmt.Println("Decrease Volume")
  go platform_specific.IncreaseVolume(-2)
}

func mute (w http.ResponseWriter, r *http.Request) {
  fmt.Println("Mute/Unmute Volume")
  go platform_specific.MuteVolume()
}

func pause (w http.ResponseWriter, r *http.Request) {
  fmt.Println("Mute/Unmute Volume")
  go platform_specific.Pause()
}

func switch_app(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Switch App")
  go platform_specific.SwitchApp()
}

func send (w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)
  fmt.Println(string(body))
  if err != nil {
    http.Error(w, "can't read body", http.StatusBadRequest)
    return
  }
  go platform_specific.SendKeyBoardInput(string(body))
}

func graceful_shutdown(w http.ResponseWriter, r *http.Request, shutdown_channel chan bool) {
  fmt.Println("Going to shutdown the server gracefully")
  shutdown_channel <- true
}

func main() {
  // Channel used to shutdown gracefully the server
  shutdown_channel := make(chan bool)

  server := &http.Server{Addr: ":8080"}
  defer server.Close()
  http.HandleFunc("/", sayHello)
  http.HandleFunc("/increase", increase)
  http.HandleFunc("/decrease", decrease)
  http.HandleFunc("/mute", mute)
  http.HandleFunc("/pause", pause)
  http.HandleFunc("/send", send)
  http.HandleFunc("/switch", switch_app)
  http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) { 
    graceful_shutdown(w, r, shutdown_channel)
  })
  
  go func() {
    if err := server.ListenAndServe(); err != nil {
      panic(err)
    }
  }()
  
  // Waiting for a message in the shutdown channel
  <- shutdown_channel
}