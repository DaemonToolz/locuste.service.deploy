package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/keybase/go-ps"
)

// Run Exécute une commande
func Run(cmd ExecCommand) {
	isRunning := checkRunning(cmd.Application)
	version := AppVersion{
		Name:    cmd.Application,
		Version: cmd.Version,
	}
	log.Println(version, isRunning)
	switch cmd.Command {
	case Start:
		if !isRunning {
			rootVersion := GetDiskVersion(true)
			dir := fmt.Sprintf("/home/pi/Documents/locuste/%s/%s/%s", rootVersion.GlobalVersion, cmd.Application, cmd.Version)
			executablePath := fmt.Sprintf("%s/%s", dir, cmd.Application)
			startProcess(cmd.Application, dir, executablePath)
			version.IsRunning = true
		}
	case Stop:
		if isRunning {
			stopProcess(getProcess(cmd.Application).Pid())
			version.IsRunning = false
		}
	}

	broadcastProcessUpdate(version)

}

func startProcess(application, dir, path string) {

	var attr = os.ProcAttr{
		Dir: dir,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
		Env: os.Environ(),
	}
	process, err := os.StartProcess(application, []string{path}, &attr)
	if err == nil {
		err = process.Release()
		if err != nil {
			failOnError(err, "Impossible de détacher le processus")
		}

	} else {
		failOnError(err, "Impossible de démarrer le processus")
	}
}

func stopProcess(pid int) {
	syscall.Kill(pid, syscall.SIGINT)
}

func checkRunning(target string) bool {
	processes, err := ps.Processes()
	if err != nil {
		failOnError(err, "Error :")
		return false
	}
	for index := range processes {
		if strings.Contains(processes[index].Executable(), target) {
			return true
		}
	}

	return false
}

func getProcess(target string) ps.Process {
	processes, err := ps.Processes()
	if err != nil {
		failOnError(err, "Error :")
		return nil
	}

	for index := range processes {
		if strings.Contains(processes[index].Executable(), target) {
			return processes[index]
		}
	}

	return nil
}
