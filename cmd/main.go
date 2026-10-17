package main

import (
	"os"
	"fmt"
	"errors"
	"mu/internal/commands"
	"mu/internal/config"
)

func main() {
	argv := os.Args[1:];
	if err := handleCommandLineArgs(argv); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err);
		os.Exit(1);
	}
}

func handleCommandLineArgs(argv []string) error {
	conf, err := config.GetConfig();
	if err != nil {
		return err;
	}
	command := conf.DefaultCommand;
	if len(argv) > 0 {
		command = argv[0];
	}

	switch command {
	case "new", "nw", "n":
        if len(argv) < 2 {
            return errors.New("command: `new` expects argument for layout name");
        }
        name := argv[1];
        err = commands.New(name);
	case "edit", "ed", "e":
        err := commands.Edit();
        if err != nil {
            return err;
        }
	case "start", "st":
        if len(argv) < 2 {
            return errors.New("command: `start` expects argument for layout name");
        }
        name := argv[1];
        err = commands.Start(name);
	case "attach", "at", "a":
		var name string;
        if len(argv) >= 2 {
			name = argv[1];
        }
		err = commands.Attach(name);
	case "list", "ls":
        err = commands.List();
	default:
		return fmt.Errorf("unknown command: `%v`", command);
	}
	if err != nil {
		return err;
	}
	
	return nil;
}

