package main

import (
  "fmt"
  "io/ioutil"
  "log"

  "gopkg.in/yaml.v2"
  "github.com/cbroglie/mustache"
)

func generatePage(name string, context map[string]interface{}) {
  f, err := ioutil.ReadFile(fmt.Sprintf("mustache/%s.yaml", name))
  if err != nil {
      log.Printf("yamlFile.Get err   #%v ", err)
  }

  var yamldata map[string]interface{}
  err = yaml.Unmarshal(f, &yamldata)
  if err != nil {
    log.Fatalf("Unmarshal: %v", err)
  }

  for k, v := range yamldata {
      context[k] = v
  }
  data, err := mustache.RenderFileInLayout(fmt.Sprintf("mustache/%s.html.mustache", name), "mustache/layout.html.mustache", context)

  if err != nil {
    log.Fatalf(fmt.Sprintf("error generating %s page: %s.", name, err))
    return
  }

  if name == "video" {
    ioutil.WriteFile("index.html", []byte(data), 0644)
  } else {
    ioutil.WriteFile(fmt.Sprintf("%s.html", name), []byte(data), 0644)
  }
}

func main() {
  fmt.Println("generating site...")

  aboutcontext := map[string]interface{}{"about": true}
  generatePage("about", aboutcontext)

  videocontext := map[string]interface{}{
    "video": true,
  }
  generatePage("video", videocontext)

  photocontext := map[string]interface{}{"photo": true}
  generatePage("photo", photocontext)

  fmt.Println("done.")
}
