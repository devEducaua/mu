package tmux

import (
	"fmt"
	"mu/internal/utils"
	"os"
	"strings"
)

func GoToSession(session string) error {
	env := os.Getenv("TMUX");

	cmd := "-t"
	if session == "" { cmd = ""; }
	if env != "" {
		return utils.RunTmuxCommand("switch-client", cmd, session);
	}
	 return utils.RunTmuxCommand("attach-session", cmd, session);
}

func SessionIsRunning(session string) (bool, error) {

	output, err := utils.RunTmuxCommandWithOutput("list-sessions");

	if strings.HasPrefix(string(output), "no server running on ") {
		return false, nil;
	}

	if err != nil {
		return false, err;
	}

	for l := range strings.SplitSeq(string(output), "\n") {
		line := strings.TrimSpace(l);
		if strings.HasPrefix(line, session+":") {
			return true, nil;
		}
	}

	return false, nil;
}

func splitWindow(session, root, direction string, size int64, window any) error {
	if direction != "-h" && direction != "-v" {
		return fmt.Errorf("invalid direction to split window: `%v`", direction);
	}

	command := []string{"split-window", "-c", root, "-t", fmt.Sprintf("%v:%v", session, window), direction, "-p", fmt.Sprintf("%v", size)};
    err := utils.RunTmuxCommand(command...);
    if err != nil {
        return err;
    }
    return nil;
}

func newWindow(session, name, root string) error {
    command := []string{"new-window", "-d", "-c", root, "-n", name, "-t", session};
    err := utils.RunTmuxCommand(command...);
    if err != nil {
        return err;
    }
    return nil;
}

func renameWindow(session, name string, idx any) error {
	command := []string{"rename-window", "-t", fmt.Sprintf("%v:%v", session, idx), name};
    err := utils.RunTmuxCommand(command...);
    if err != nil {
        return err;
    }
    return nil;
}

func sendKeys(session string, window any, paneIndex int, keys string) error {
    command := []string{"send-keys", "-t", fmt.Sprintf("%v:%v.%v", session, window, paneIndex), keys, "C-m"};
    err := utils.RunTmuxCommand(command...);
    if err != nil {
        return err;
    }
	return nil;
}
