package git

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"regexp"
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

func ShortStatus(path string) (io.Reader, error) {
	return execOutput(fmt.Sprintf("git -C %s status -s -b", path))
}

var RemoteName = func(path string) string {
	out, err := execOutput(fmt.Sprintf("git -C %s remote", path))
	if err != nil {
		return "origin"
	}

	s := bufio.NewScanner(out)

	//read only the first remote then breaks
	var remote string
	for s.Scan() {
		if len(s.Text()) < 1 {
			continue
		}

		remote = s.Text()
		break
	}

	return remote
}

func DefaultBranch(path string) string {
	remote := RemoteName(path)
	out, err := execOutput(fmt.Sprintf("git -C %s symbolic-ref --short refs/remotes/%s/HEAD ", path, remote))
	if err != nil {
		return "master"
	}

	d, err := ioutil.ReadAll(out)
	if err != nil {
		return "master"
	}

	re := regexp.MustCompile(fmt.Sprintf("%s/(.*)", remote))
	match := re.FindStringSubmatch(string(d))
	return match[1]

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

var execOutput = func(c string) (io.Reader, error) {
	out, err := exec.Command("/bin/sh", "-c", c).Output()

	return bytes.NewReader(out), err
}