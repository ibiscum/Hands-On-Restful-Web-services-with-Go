package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	// define flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "save",
			Value: "no",
			Usage: "should save to database (yes/no)",
		},
	}

	app.Version = "1.0"

	// define action
	app.Action = func(c *cli.Context) error {
		var args []string
		if c.NArg() > 0 {
			// Fetch arguments in a array
			args = c.Args()
			personName := args[0]
			marks := args[1:]
			log.Println("Person:", personName)
			log.Println("Marks:", marks)
		}

		// check the flag value
		if c.String("save") == "no" {
			log.Println("Skipping saving to the database")
		} else {
			// Add database logic here
			log.Println("Saving to the database", args)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
