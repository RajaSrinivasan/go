package newutil

import (
  "os"
  "io"
  "io/ioutil"
  "fmt"
  "path"
  "strings"
)

func copyFile(from string, to string){
  outfilename := to
  if Filenames {
     for k,v := range(Macros) {
        outfilename = strings.Replace( outfilename , k , v , -1 )
     }
  }
  fmt.Printf("Output file is %s\n",outfilename)
  inpfile,_ := os.Open(from)
  outfile,_ := os.Create(outfilename)
  io.Copy(outfile,inpfile)
  outfile.Sync()
  inpfile.Close()
  outfile.Close()
}

func Create(template string, output string) {
  if Verbose {
    fmt.Printf("Create template=%s output=%s\n",template,output)
  }
    files, err := ioutil.ReadDir(template)
    if err != nil {
      panic(err)
    }
    for _,f := range(files) {
      if Verbose {
        fmt.Printf("%s\n",f.Name())
      }
      if f.IsDir() {
        if Verbose {
           fmt.Printf("Copy Dir %s\n",f.Name())
        }
        os.Mkdir(path.Join(output,f.Name()),os.ModeDir+os.ModePerm)
        Create( path.Join(template,f.Name()) , path.Join(output,f.Name()))
      } else {
        copyFile( path.Join(template,f.Name()) , path.Join(output,f.Name()))
      }
    }
}

func init() {

}
