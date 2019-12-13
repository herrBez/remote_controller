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
  "os/exec"
  "os"
  "bytes"
)

// Code for the welcome page, simply load the static web page
func sayHello(w http.ResponseWriter, r *http.Request) {
  message, err := ioutil.ReadFile("./html/index.html")
  if err != nil {
    panic(err)
  }
  w.Write([]byte(message))
}


// Simple wrapper that execute a command and prints the output result in the
// command line
// https://stackoverflow.com/questions/18159704/how-to-debug-exit-status-1-error-when-running-exec-command-in-golang
func executeCommand(cmd *exec.Cmd) { 
  var out bytes.Buffer
  var stderr bytes.Buffer
  cmd.Stdout = &out
  cmd.Stderr = &stderr
  err := cmd.Run()
  if err != nil {
    fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    return
  }
  fmt.Println("Result: " + out.String())   
}

// Backend function that increases the volume of the computer of the specified
// percent value. 65535 is the maximal value.
func _increase(percent int) {
  delta := 65535*percent/100
  parameter := fmt.Sprintf("%d", delta)
  cmd := exec.Command("cmd", "/C", "nircmd.exe", "changesysvolume", parameter)
  working_directory, _ := os.Getwd()
  cmd.Dir = working_directory
  executeCommand(cmd) 
}

// Function that mute and unmute the volume of the computer
func _mute() {
  parameter := fmt.Sprintf("%d", 2)
  cmd := exec.Command("cmd", "/C", "nircmd.exe", "mutesysvolume", parameter)
  executeCommand(cmd)
}

// Send a space to the system https://docs.microsoft.com/en-us/windows/win32/inputdev/virtual-key-codes?redirectedfrom=MSDN
func _pause() {
  parameter := fmt.Sprintf("0x20")
  cmd := exec.Command("cmd", "/C", "nircmd.exe", "sendkey", parameter, "press")
  executeCommand(cmd)
}

func _send(key string) {
  cmd := exec.Command("cmd", "/C", "nircmd.exe", "sendkey", key, "press")
  executeCommand(cmd)
}

func increase (w http.ResponseWriter, r *http.Request) {
  fmt.Println("Increase Volume")
  go _increase(2)
  sayHello(w, r)
}

func decrease (w http.ResponseWriter, r *http.Request) {
  fmt.Println("Decrease Volume")
  go _increase(-2)
  sayHello(w, r)
}

func mute (w http.ResponseWriter, r *http.Request) {
  fmt.Println("Mute/Unmute Volume")
  go _mute()
  sayHello(w, r)
}

func pause (w http.ResponseWriter, r *http.Request) {
  fmt.Println("Mute/Unmute Volume")
  go _pause()
  sayHello(w, r)
}

func graceful_shutdown(w http.ResponseWriter, r *http.Request, shutdown_channel chan bool) {
  fmt.Println("Going to shutdown the server gracefully")
  shutdown_channel <- true
}

func send (w http.ResponseWriter, r *http.Request) {
  body, err := ioutil.ReadAll(r.Body)
  fmt.Println(string(body))
  if err != nil {
    http.Error(w, "can't read body", http.StatusBadRequest)
    return
  }
  _send(string(body))
}

func main() {
  // Channel used to shutdown gracefully the server
  shutdown_channel := make(chan bool)

  server := &http.Server{Addr: ":8080"}
  http.HandleFunc("/", sayHello)
  http.HandleFunc("/increase", increase)
  http.HandleFunc("/decrease", decrease)
  http.HandleFunc("/mute", mute)
  http.HandleFunc("/pause", pause)
  http.HandleFunc("/send", send)
  http.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) { 
    graceful_shutdown(w, r, shutdown_channel)
  })
  
  go func() {
    if err := server.ListenAndServe(); err != nil {
      panic(err)
    }
  }()
  
  <- shutdown_channel
  server.Close()
}