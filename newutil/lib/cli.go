package newutil

import (
  "fmt"
  "flag"
  "regexp"
)

var Verbose bool
var Template string
var Output string
var Macros map[string]string
var macDefinition *regexp.Regexp

func parseMacro(arg string) (string,string) {

  defn := macDefinition.FindAllStringSubmatch(arg,-1)
  fmt.Printf("%v macro=%s value=%s\n",defn,defn[0][1],defn[0][2])
  return defn[0][1] , defn[0][2]
}

func init() {

  flag.BoolVar( &Verbose , "verbose" , true , "verbose")
  flag.StringVar( &Template , "template" , "./template" , "template dir")
  flag.StringVar( &Output , "output" , "./output" , "output dir") 
  flag.Parse()
  Macros = make(map[string]string)
  macDefinition = regexp.MustCompile("([a-zA-Z][a-zA-Z0-9]*)=([a-zA-Z0-9]+)")
  for arg:=0; arg < flag.NArg(); arg++ {
    mac,val := parseMacro(flag.Arg(arg))
    Define(mac,val)
  }
}

func showMacros() {
  fmt.Printf("Macros defined\n")
  for k,v := range( Macros ) {
    fmt.Printf("Macro = %s Value = %s\n",k,v)
  }
}

func Macro(key string) string {
   return Macros[key]
}

func Define(key string , val string) {
  Macros[key] = val
}

func AnalyzeCommandLine() {
     if Verbose {
       showMacros()
     }
}
