package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli/v2"
  "github.com/zshift/luxafor"
)

func main() {
  app := &cli.App{
    Name: "Luxafor Flag CLI",
    Usage: "Talk to your flag",
    Flags: []cli.Flag {
      &cli.StringFlag{
        Name:    "leds",
        Aliases: []string{"l"},
        Value:   "all",
        Usage:   "specify which led to change, or all, front, back",
        // Destination: &leds,
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
      fmt.Println(device)

      // --leds flag
      // fmt.Println(leds)
      // switch leds {
      //   case "all":
      //     // Set all
      //   case "front":
      //     // set front
      //   case "back":
      //     // set back
      //   case "1":
      //     // set first
      //   case "2":
      //     // set second
      //   case "3":
      //     // set third
      //   case "4":
      //     // set fourth
      //   case "5":
      //     // set fifth
      //   case "6":
      //     // set sixth
      //   default:
      //     // Set all
      // }

      return nil
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
