package status

import (
	"bufio"
	"io"
	"strings"
)

type Output struct {
	Branch string
	FilesStatus []string
}

//Parse parses a git status output command
//It is compatible with the short version of the git status command
func ParseShort(r io.Reader) Output {

	s := bufio.NewScanner(r)

	var branch string
	//Extract branch name
	for s.Scan() {
		if len(s.Text()) < 1 {
			continue
		}

		branch = parseBranch(s.Text())
		break
	}


	var fs []string
	for s.Scan() {
		if len(s.Text()) < 1 {
			continue
		}
		fs = append(fs, s.Text())
	}

	return Output{
		Branch:      branch,
		FilesStatus: fs,
	}
}

func parseBranch(input string) string {
	s := bufio.NewScanner(strings.NewReader(input))
	s.Split(bufio.ScanWords)

	//check if input is a status branch line output
	s.Scan()
	if s.Text() != "##" {
		return ""
	}

	//read next word and return the branch name
	s.Scan()
	b := strings.Split(s.Text(), "...")
	return b[0]
}