package main

import (
  "fmt"
  "log"
  "os"
  "strconv"

  "github.com/urfave/cli/v2"
  "github.com/zshift/luxafor"
)

func main() {
  var leds string

  app := &cli.App{
    Name: "Luxafor Flag CLI",
    Usage: "Talk to your flag",
    Flags: []cli.Flag {
      &cli.StringFlag{
        Name:        "leds",
        Aliases:     []string{"l"},
        Value:       "all",
        Usage:       "specify which led to change. Can also be 'all', 'front', or 'back'",
        Destination: &leds,
      },
    },
    Action: func(c *cli.Context) error {
      // Get device
      luxs := luxafor.Enumerate()
      if len(luxs) == 0 {
        fmt.Println("No attached devices. Exiting.")
        os.Exit(0)
      }

      device := luxs[1]

      // Red
      r, err := strconv.ParseUint(c.Args().Get(0), 10, 8)
      if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
      }

      // Green
      g, err := strconv.ParseUint(c.Args().Get(1), 10, 8)
      if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
      }

      // Blue
      b, err := strconv.ParseUint(c.Args().Get(2), 10, 8)
      if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
      }

      // --leds flag
      switch leds {
        case "all":
          err = device.Set(luxafor.All, uint8(r), uint8(g), uint8(b))
          if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
          }

        case "front":
          err = device.Set(luxafor.FrontAll, uint8(r), uint8(g), uint8(b))
          if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
          }
        case "back":
          err = device.Set(luxafor.BackAll, uint8(r), uint8(g), uint8(b))
          if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
          }
        case "1":
          err = device.Set(luxafor.FrontTop, uint8(r), uint8(g), uint8(b))
          if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
          }
        case "2":
          err = device.Set(luxafor.FrontMiddle, uint8(r), uint8(g), uint8(b))
          if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
          }
        case "3":
          err = device.Set(luxafor.FrontBottom, uint8(r), uint8(g), uint8(b))
          if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
          }
        case "4":
          err = device.Set(luxafor.BackTop, uint8(r), uint8(g), uint8(b))
          if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
          }
        case "5":
          err = device.Set(luxafor.BackMiddle, uint8(r), uint8(g), uint8(b))
          if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
          }
        case "6":
          err = device.Set(luxafor.BackBottom, uint8(r), uint8(g), uint8(b))
          if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
          }
        default:
          fmt.Println("Specify an led with --leds flag")
          os.Exit(1)
      }

      return nil
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
