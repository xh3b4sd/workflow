package env

import (
	"os"
	"path/filepath"
	"strings"
)

func Env() map[string]string {
	var err error

	var pat string
	{
		pat = ".workflow/"
	}

	var fil []os.DirEntry
	{
		fil, err = os.ReadDir(pat)
		if os.IsNotExist(err) {
			// fall through
		} else if err != nil {
			panic(err)
		}
	}

	env := map[string]string{}
	{
		for _, f := range fil {
			if f.IsDir() {
				continue
			}

			byt, err := os.ReadFile(filepath.Join(pat, f.Name()))
			if err != nil {
				panic(err)
			}

			env[f.Name()] = strings.TrimSpace(string(byt))
		}
	}

	return env
}
