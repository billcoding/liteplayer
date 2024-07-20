package envs

import "os"

func bindEnv(name, defVal string) (bindVal string) {
	bindVal = defVal
	if val := os.Getenv(name); val != "" {
		bindVal = val
	}
	return
}
