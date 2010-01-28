package main

var varsHolder map[string]string

func initializeVars() {
	var m = map[string]string{
		`PS1`: `\u@\h \w \$ `,
		`PS2`: `> `,
		`PS3`: `#? `,
		`PS4`: `+ `,
	}
	varsHolder = m
}

func GetLocalVar(name string) string { return varsHolder[name] }

func SetLocalVar(name, val string) { varsHolder[name] = val }
