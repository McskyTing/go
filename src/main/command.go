package main

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func runInLinux(cmd string) string {
	fmt.Println("Running Linux Cmd:" + cmd)
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(result))
}

func runInWindows(cmd string) string {
	fmt.Println("Running Win Cmd:" + cmd)
	result, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(result))
}

func RunCommand(cmd string) string {
	if runtime.GOOS == "windows" {
		return runInWindows(cmd)
	} else {
		return runInLinux(cmd)
	}
}

func RunLinuxCommand(cmd string) string {
	if runtime.GOOS == "windows" {
		return ""
	} else {
		return runInLinux(cmd)
	}
}

func runInLinuxWithErr(cmd string) (string, error) {
	fmt.Println("Running Linux Cmd:" + cmd)
	result, err := exec.Command("/bin/sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(result)), err
}

func runInWindowsWithErr(cmd string) (string, error) {
	fmt.Println("Running Win Cmd:" + cmd)
	result, err := exec.Command("cmd", "/c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(result)), err
}

func RunCommandWithErr(cmd string) (string, error) {
	if runtime.GOOS == "windows" {
		return runInWindowsWithErr(cmd)
	} else {
		return runInLinuxWithErr(cmd)
	}
}

func RunLinuxCommandWithErr(cmd string) (string, error) {
	if runtime.GOOS == "windows" {
		return "", errors.New("could not run in windows OS")

	} else {
		return runInLinuxWithErr(cmd)
	}
}
