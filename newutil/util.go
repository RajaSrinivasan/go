package main

import (
  "os"
  "path/filepath"
  "fmt"
  "./lib"
)

func processFile (path string, info os.FileInfo, err error) error {
  fmt.Printf("%s\n",path)
  return nil
}
func main() {
  newutil.AnalyzeCommandLine()
  if _, err := os.Stat(newutil.Template); os.IsNotExist(err) {
     fmt.Printf("Template Directory %s does not exist\n",newutil.Template)
     return
  }
  if newutil.Verbose {
     filepath.Walk(newutil.Template,processFile)
  }
  os.Mkdir(newutil.Output,os.ModeDir+os.ModePerm)
  newutil.Create(newutil.Template,newutil.Output)
}
