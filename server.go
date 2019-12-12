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

// Backend function that increases the volume of the computer of the specified
// percent value. 65535 is the maximal value.
func _increase(percent int) {
  delta := 65535*percent/100
  parameter := fmt.Sprintf("%d", delta)
  cmd := exec.Command("cmd", "/C", "nircmd.exe", "changesysvolume", parameter)
  working_directory, err := os.Getwd()
  cmd.Dir = working_directory

  // https://stackoverflow.com/questions/18159704/how-to-debug-exit-status-1-error-when-running-exec-command-in-golang
  var out bytes.Buffer
  var stderr bytes.Buffer
  cmd.Stdout = &out
  cmd.Stderr = &stderr
  err = cmd.Run()
  if err != nil {
    fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    return
  }
  fmt.Println("Result: " + out.String())   
}

func increase (w http.ResponseWriter, r *http.Request) {
  fmt.Println("Increase Volume")
  _increase(5)
  sayHello(w, r)
}

func decrease (w http.ResponseWriter, r *http.Request) {
  fmt.Println("Decrease Volume")
  _increase(-5)
  sayHello(w, r)
}

func main() {
  http.HandleFunc("/", sayHello)
  http.HandleFunc("/increase", increase)
  http.HandleFunc("/decrease", decrease)

  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}