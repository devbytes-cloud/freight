package commands

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type hookStructure struct {
	inputHooks     []string
	expectedHooks  map[string]struct{}
	expectedErr    bool
	expectedErrMgs error
}

func TestValidateAllowHooks(t *testing.T) {
	testData := map[string]hookStructure{
		"no input hooks": {
			expectedErr: false,
			inputHooks:  []string{},
			expectedHooks: map[string]struct{}{
				"pre-commit":         {},
				"prepare-commit-msg": {},
				"commit-msg":         {},
				"post-commit":        {},
				"post-checkout":      {},
			},
		},
		"only pre-commit hook": {
			expectedErr: false,
			inputHooks:  []string{"pre-commit"},
			expectedHooks: map[string]struct{}{
				"pre-commit": {},
			},
		},
		"invalid hook name hook": {
			expectedErr:    true,
			inputHooks:     []string{"invalid hook name"},
			expectedErrMgs: fmt.Errorf("invalid hook types: invalid hook name"),
			expectedHooks:  nil,
		},
		"multiple valid hooks": {
			expectedErr: false,
			inputHooks:  []string{"pre-commit", "commit-msg", "post-checkout"},
			expectedHooks: map[string]struct{}{
				"pre-commit":    {},
				"commit-msg":    {},
				"post-checkout": {},
			},
		},
		"multiple invalid hooks": {
			expectedErr:    true,
			inputHooks:     []string{"invalid1", "invalid2"},
			expectedErrMgs: fmt.Errorf("invalid hook types: invalid1, invalid2"),
			expectedHooks:  nil,
		},
		"mix of valid and invalid hooks": {
			expectedErr:    true,
			inputHooks:     []string{"pre-commit", "invalid-hook"},
			expectedErrMgs: fmt.Errorf("invalid hook types: invalid-hook"),
			expectedHooks:  nil,
		},
		"all valid hooks explicitly provided": {
			expectedErr: false,
			inputHooks:  []string{"pre-commit", "prepare-commit-msg", "commit-msg", "post-commit", "post-checkout"},
			expectedHooks: map[string]struct{}{
				"pre-commit":         {},
				"prepare-commit-msg": {},
				"commit-msg":         {},
				"post-commit":        {},
				"post-checkout":      {},
			},
		},
		"duplicate valid hooks": {
			expectedErr: false,
			inputHooks:  []string{"pre-commit", "pre-commit"},
			expectedHooks: map[string]struct{}{
				"pre-commit": {},
			},
		},
	}

	for name, test := range testData {
		t.Run(name, func(t *testing.T) {
			resp, err := validateAllowHooks(test.inputHooks)

			if test.expectedErr {
				assert.Error(t, err)
				assert.EqualError(t, err, test.expectedErrMgs.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedHooks, resp)
			}
		})
	}
}
