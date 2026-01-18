package commands

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	isVerbose       bool
	setVersion      string
	setDate         string
	setCommit       string
	expectedVersion string
	expectedDate    string
	expectedCommit  string
}

func TestVersionCommand(t *testing.T) {
	testMap := map[string]testStruct{
		"default build": {
			isVerbose:       false,
			expectedVersion: "dev",
		},
		"set version": {
			isVerbose:       false,
			setVersion:      "test-a",
			expectedVersion: "test-a",
		},
		"verbose: set version": {
			isVerbose:       true,
			setVersion:      "test-a",
			expectedVersion: "test-a",
		},
		"verbose: set date, commit, version": {
			isVerbose:       true,
			setVersion:      "test-b",
			expectedVersion: "test-b",
			setCommit:       "test-c",
			expectedCommit:  "test-c",
			setDate:         "2020-04-06T00:00:00Z",
			expectedDate:    "2020-04-06T00:00:00Z",
		},
		"verbose: set date, commit": {
			isVerbose:      true,
			setCommit:      "test-c",
			expectedCommit: "test-c",
			setDate:        "2020-04-06T00:00:00Z",
			expectedDate:   "2020-04-06T00:00:00Z",
		},
	}

	for name, td := range testMap {
		t.Run(name, func(t *testing.T) {
			origVersion, origCommit, origDate := Version, Commit, Date
			defer func() {
				Version, Commit, Date = origVersion, origCommit, origDate
			}()

			if td.setVersion != "" {
				Version = td.setVersion
			}
			if td.setDate != "" {
				Date = td.setDate
			}
			if td.setCommit != "" {
				Commit = td.setCommit
			}

			cmd := versionCommand()
			buf := new(bytes.Buffer)
			cmd.SetOut(buf)
			cmd.SetErr(buf)

			if td.isVerbose {
				cmd.SetArgs([]string{"--verbose"})
			}

			err := cmd.Execute()

			assert.NoError(t, err)
			fmt.Println(buf.String())
			assert.Contains(t, buf.String(), td.expectedVersion)

			if td.isVerbose {
				assert.Contains(t, buf.String(), td.expectedCommit)
				assert.Contains(t, buf.String(), td.expectedDate)
			} else {
				assert.NotContains(t, buf.String(), "Commit:")
				assert.NotContains(t, buf.String(), "Date:")
			}
		})
	}
}
