package commands

import (
	"fmt"
	"mu/internal/config"
	"mu/internal/tmux"
	"mu/internal/utils"
	"os"
	"path/filepath"
	"strings"
)

func New(name string) error {
	baseDir, err := config.GetBaseDir();
	if err != nil {
		return err;
	}

	layoutsDir := filepath.Join(baseDir, "layouts/");

	fileName := name+".kdl";
	exists, err := utils.FileExistsInDir(fileName, layoutsDir);
	if err != nil {
		return err;
	}

	layoutPath := filepath.Join(layoutsDir, name+".kdl");
	if !exists {

		cont, err := config.GetBaseFile();
		if err != nil {
			return err;
		}

		baseNewLayout := fmt.Sprintf(cont, name);

		f, err := os.OpenFile(layoutPath, os.O_CREATE|os.O_EXCL|os.O_APPEND|os.O_RDWR, 0666);
		if err != nil {
			return err;
		}
		defer f.Close();
		f.WriteString(baseNewLayout);
	}
	err = utils.OpenEditor(layoutPath);
	if err != nil {
		return fmt.Errorf("could not open $EDITOR: %v: but base was wrote to `%v`\n", err, layoutPath);
	}

    return nil;
}

func Edit() error {
	layoutsDir, err := config.GetLayoutsDir();
	if err != nil {
		return err;
	}
	err = utils.OpenEditor(layoutsDir);
	if err != nil {
		return err;
	}

	return nil;
}

func Start(name string) error {
	layoutsDir, err := config.GetLayoutsDir();
	if err != nil {
		return err;
	}

	fileName := name+".kdl";
	exists, err := utils.FileExistsInDir(fileName, layoutsDir);
	if err != nil {
		return err;
	}
	if !exists {
		return fmt.Errorf("layout: `%v` doesn't exists", name);
	}

	filePath := filepath.Join(layoutsDir, fileName);
	dat, err := os.ReadFile(filePath);
	if err != nil {
		return err;
	}

	l, err := config.ParseLayout(dat)
	if err != nil {
		return err;
	}

	running, err := tmux.SessionIsRunning(l.Name);
	if err != nil {
		return err;
	}

	if !running {
		err = tmux.LayoutToSession(l);
		if err != nil {
			return err;
		}
	}

	if l.Attach {
		err = tmux.GoToSession(l.Name);
		if err != nil {
			return err;
		}
	}
	return nil;
}

func Attach(name string) error {
	return tmux.GoToSession(name);
}

func List() error {
	dir, err := config.GetLayoutsDir();
	if err != nil {
		return err;
	}

	dirs, err := os.ReadDir(dir);
	if err != nil {
		return err;
	}

	for _,e := range dirs {
		name := strings.TrimSuffix(e.Name(), ".kdl");
		fmt.Printf("%v\n", name);
	}
	return nil;
}

