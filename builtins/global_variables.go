package builtins

import "github.com/nikolalohinski/gonja/v2/exec"

var GlobalVariables = exec.NewContext(map[string]interface{}{
	"gonja": map[string]interface{}{
		"version": "0.0.0+trunk",
	},
})
