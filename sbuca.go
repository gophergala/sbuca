package main

import (
  "fmt"
  "github.com/codegangsta/cli"
  "github.com/gophergala/sbuca/server"
  "os"
)

func main() {

  app := cli.NewApp()
  app.Name = "sbuca"
  app.Usage = "Simple But Useful CA"
  app.Commands = []cli.Command{

    {
      Name: "server",
      Action: func(c *cli.Context) {
        fmt.Println("running")
        server.Run()
      },
    },

  }

  app.Run(os.Args)

}
