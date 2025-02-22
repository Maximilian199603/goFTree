package nerdtree

import ()

// An mapping from an filename to an empty string is not allowed
// as in the validation if the access to the dictionary succeeded
// there is an check if the return is an empty string
var fileNameToIcon = map[string]string{
	"dockerfile":    " ", // Docker
	".dockerignore": " ", // Docker

	".gitignore": "󰊢 ", // Git

	"license": " ", //License

	"go.mod": " ", // Go package management
	"go.sum": " ", // Go package management
}
