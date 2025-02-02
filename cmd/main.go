package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

const (
	star            = "star"
	info            = "info"
	approve         = "approve"
	revoke          = "revoke"
	comment         = "comment"
	overviewComment = "overviewComment"
	deleteComment   = "deleteComment"
	editComment     = "editComment"
	reply           = "reply"
	listDiscussions = "listDiscussions"
)

func main() {

	branchName, err := getCurrentBranch()
	errCheck(err)
	if branchName == "main" || branchName == "master" {
		return
	}

	var c Client
	errCheck(c.Init(branchName))

	switch c.command {
	case star:
		errCheck(c.Star())
	case approve:
		errCheck(c.Approve())
	case revoke:
		errCheck(c.Revoke())
	case comment:
		errCheck(c.Comment())
	case deleteComment:
		errCheck(c.DeleteComment())
	case editComment:
		errCheck(c.EditComment())
	case overviewComment:
		errCheck(c.OverviewComment())
	case info:
		errCheck(c.Info())
	case reply:
		errCheck(c.Reply())
	case listDiscussions:
		errCheck(c.ListDiscussions())
	default:
		c.Usage("command")
	}
}

func errCheck(err error) {
	if err != nil {
		log.Fatalf("Failure: %s", err)
		os.Exit(1)
	}
}

/* Gets the current branch */
func getCurrentBranch() (res string, e error) {
	gitCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")

	output, err := gitCmd.Output()
	if err != nil {
		return "", fmt.Errorf("Error running git rev-parse: %w", err)
	}

	return strings.TrimSpace(string(output)), nil

}
