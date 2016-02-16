package commands

import (
	"flag"
	"fmt"
	"gopkg.in/fsnotify.v1"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

var openFlagSet = flag.NewFlagSet("open", flag.ExitOnError)

var (
	openEditor   = openFlagSet.String("e", "", "Editor to open files with")
	openScpFlags = openFlagSet.String("scp-flags", "", "Additional scp flags")
)

const defaultEditor = "vim"
const localStorage = "/tmp/ssedit"

// openFile opens a file or directory from a remote server and stores it in a
// temp directory locally while edits are made.
func openFile(args []string) {
	if len(args) < 2 {
		fmt.Printf("Error: expecting two arguments\n")
		os.Exit(0)
	}
	openFlagSet.Parse(args)
	args = openFlagSet.Args()

	editor := defaultEditor
	if *openEditor != "" {
		editor = *openEditor
	}
	scpFlags := []string{}
	if *openScpFlags != "" {
		scpFlags = strings.Split(*openScpFlags, " ")
	}

	fmt.Printf("Starting ssedit session...\n")
	fmt.Printf("Ctrl+C to exit\n")

	dirs := strings.Split(args[1], "/")
	dirs = dirs[:len(dirs)-1]

	os.MkdirAll(localStorage+strings.Join(dirs, "/"), 0755)

	hostPath := fmt.Sprintf("%s:%s", args[0], args[1])
	localPath := fmt.Sprintf("/tmp/ssedit%s", args[1])

	cmdFlags := scpFlags
	cmdFlags = append(cmdFlags, "-r", "-C", hostPath, localPath)

	cmd := exec.Command("scp", cmdFlags...)

	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go startWatcher(localPath, hostPath)

	startEditor(editor, localPath)

	wg.Wait()
}

// startEditor starts the user's default text editor for use with ssedit.
func startEditor(editor string, localPath string) {
	cmd := exec.Command(editor, localPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()

	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

// startWatcher starts the file / directory watcher for new events to then
// transfer the updated file back to the remote.
func startWatcher(localPath string, hostPath string) {
	file, err := os.Open(localPath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	var remotefilePath string
	fi, err := file.Stat()
	switch {
	case err != nil:
		log.Fatal(err)
		os.Exit(1)
	case fi.IsDir():
		hostParts := strings.Split(hostPath, "/")
		remotefilePath = strings.Join(hostParts[:len(hostParts)-1], "/")
	default:
		remotefilePath = hostPath
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op == 2 || event.Op == 18 {
					cmd := exec.Command("scp", "-r", localPath, remotefilePath)
					err := cmd.Start()
					if err != nil {
						log.Fatal(err)
					}
					err = cmd.Wait()
					fmt.Printf("Saved to: %s\n", remotefilePath)
				}
			case err := <-watcher.Errors:
				fmt.Println("Error:", err)
			}
		}
	}()

	err = watcher.Add(localPath)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

// openCmd defines the "open" subcommand
var openCmd = &Command{
	Usage: func(cmd string) {
		fmt.Printf("Usage: %s [-t] [<remote> <path>]\n\nOptions:\n", cmd)
		openFlagSet.PrintDefaults()
	},
	RunMethod: func(args []string) error {
		openFile(args)
		return nil
	},
}
