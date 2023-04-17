package curl

import ()

// assemble curl command from args
func assembleCurlCommand(args []string) string {
	var curlCommand string
	for _, arg := range args {
		curlCommand = curlCommand + " " + arg
	}
	return curlCommand
}
