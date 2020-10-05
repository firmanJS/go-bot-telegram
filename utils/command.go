package utils

func RequestCommand(command string) (action string) {
	if command == "/joke" {
		action = "imam ganteng"
	} else {
		action = "command not found"
	}
	return
}
