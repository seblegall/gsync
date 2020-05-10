package git

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	cloneTimeOut = 30*time.Second
)

func Clone(repo, path string) error {
	return execWithTimeOut(fmt.Sprintf("git clone %s %s", repo, path), cloneTimeOut)
}


func Reset(path string) error {
	if err := execWithoutTimeOut(fmt.Sprintf("git -C %s reset --hard", path)); err != nil {
		return err
	}

	if err := execWithoutTimeOut(fmt.Sprintf("git -C %s clean -fd", path)); err != nil {
		return err
	}
	return nil
}

func Fetch(path string) error {
	return execWithoutTimeOut(fmt.Sprintf("git -C %s fetch --all -tp", path))
}

func Checkout(path, branch string) error {
	return execWithoutTimeOut(fmt.Sprintf("git -C %s checkout %s", path, branch))
}

func Rebase(path, branch string) error {
	return execWithoutTimeOut(fmt.Sprintf("git -C %s rebase %s", path, branch))
}

func ShortStatus(path string) (error, io.Reader) {
	return execOutput(fmt.Sprintf("git -C %s status -s -b", path))
}


func execWithoutTimeOut(c string) error {
	cmd := exec.Command("/bin/sh", "-c", c)

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		logrus.Warn("Error creating StdoutPipe for Cmd")
		return err
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			logrus.Debug(scanner.Text())
		}
	}()

	if err := cmd.Start(); err != nil {
		logrus.Warn("Error stating Cmd")
		return err
	}

	if err := cmd.Wait(); err != nil {
		logrus.Warn("The command did not succeed")
		return err
	}

	return nil
}


func execWithTimeOut(c string, t time.Duration) error {

	cmd := exec.Command("/bin/sh", "-c", c)

	cmdReader, err := cmd.StdoutPipe()
	if err != nil {
		logrus.Warn("Error creating StdoutPipe for Cmd")
		return err
	}

	scanner := bufio.NewScanner(cmdReader)
	go func() {
		for scanner.Scan() {
			logrus.Info(scanner.Text())
		}
	}()

	// Start process. Exit code 127 if process fail to start.
	if err := cmd.Start(); err != nil {
		logrus.Warn("Error stating Cmd")
		return err
	}

	var timer *time.Timer
	if t > 0 {
		timer = time.NewTimer(t)
		var err error
		go func(timer *time.Timer, cmd *exec.Cmd, err *error) {
			for range timer.C {
				e := cmd.Process.Kill()
				if e != nil {
					*err = errors.New("the command has timeout but the process could not be killed")
				} else {
					*err = errors.New("the command timed out")
				}
			}
		}(timer, cmd, &err)
	}

	err = cmd.Wait()

	if t > 0 {
		timer.Stop()
	}

	if err != nil {
		return errors.New("the command did not succeed")
	}

	return nil
}

func execOutput(c string) (error, io.Reader) {
	out, err := exec.Command("/bin/sh", "-c", c).Output()

	return err, bytes.NewReader(out)
}