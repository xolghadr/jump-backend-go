package commands

import (
	"errors"
	"strconv"
	"strings"
	"vc/workdir"
)

type VC struct {
	wd            *workdir.WorkDir
	commits       []Commit
	stagedFiles   map[string]string
	lastCommitted map[string]string // Track last committed content
}

type Commit struct {
	name string
	wd   *workdir.WorkDir
}

type Status struct {
	ModifiedFiles []string
	StagedFiles   []string
}

func Init(wd *workdir.WorkDir) *VC {
	vc := &VC{
		wd:            wd,
		commits:       []Commit{},
		stagedFiles:   make(map[string]string),
		lastCommitted: make(map[string]string), // Initialize the map
	}

	for file, fileContent := range wd.Files {
		vc.lastCommitted[file] = fileContent
	}
	return vc
}

func (vc *VC) Status() Status {
	modifiedFiles := []string{}
	for _, file := range vc.wd.ListFilesRoot() {
		if vc.wd.Files[file] == vc.stagedFiles[file] {
			continue
		}
		// Check if the file has been modified
		if lastContent, exists := vc.lastCommitted[file]; exists {
			if vc.wd.Files[file] != lastContent {
				modifiedFiles = append(modifiedFiles, file)
			}
		} else {
			// If the file was never committed, consider it modified if it exists
			if vc.wd.Files[file] != "" {
				modifiedFiles = append(modifiedFiles, file)
			}
		}
	}
	stagedFiles := []string{}
	for file := range vc.stagedFiles {
		stagedFiles = append(stagedFiles, file)
	}
	return Status{ModifiedFiles: modifiedFiles, StagedFiles: stagedFiles}
}

func (vc *VC) AddFile(wd *workdir.WorkDir, path string) error {
	if file, exists := wd.Files[path]; exists {
		vc.stagedFiles[path] = file
		return nil
	}
	return errors.New("file not found")
}

func (vc *VC) AddAll() {
	for _, file := range vc.wd.ListFilesRoot() {
		vc.AddFile(vc.wd, file)
	}
}

func (vc *VC) Add(filePath ...string) {
	for _, path := range filePath {
		vc.AddFile(vc.wd, path)
	}
}

func (vc *VC) Commit(message string) {
	newCommit := Commit{name: message, wd: vc.wd.Clone()}
	vc.commits = append(vc.commits, newCommit)
	for file := range vc.stagedFiles {
		vc.lastCommitted[file] = vc.wd.Files[file] // Store the committed content
	}
	vc.stagedFiles = make(map[string]string) // Clear staged files after commit
}

func (vc *VC) GetWorkDir() *workdir.WorkDir {
	return vc.wd
}

func (vc *VC) Log() []string {
	log := make([]string, len(vc.commits))
	for i, commit := range vc.commits {
		log[len(vc.commits)-i-1] = commit.name
	}
	return log
}

func (vc *VC) Checkout(commitName string) (*workdir.WorkDir, error) {
	if strings.HasPrefix(commitName, "^") {
		return vc.commits[len(vc.commits)-1-len(commitName)].wd.Clone(), nil
	}
	if strings.HasPrefix(commitName, "~") {
		number, err := strconv.Atoi(commitName[1:])
		if err == nil {
			return vc.commits[len(vc.commits)-1-number].wd.Clone(), nil
		}
	}
	return nil, errors.New("invalid commit name")
}
