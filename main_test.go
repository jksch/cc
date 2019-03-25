package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		name string
		args []string
		err  string
	}{
		{
			name: "Fail app setup",
			args: []string{
				"app",
				"-db_path", "",
			},
			err: "app database setup, ",
		},
		{
			name: "Fail app setup",
			args: []string{
				"app",
				"-db_path", "app.db",
				"-addr", "127", // invalid server addr
			},
			err: "app run, ",
		},
	}

	for ti, tt := range tests {
		ti, tt := ti, tt
		t.Run(fmt.Sprintf("%d. %s", ti, tt.name), func(t *testing.T) {
			os.Args = tt.args

			err := errStr(run())
			if !strings.HasPrefix(err, tt.err) {
				t.Errorf("%d. %s, exp err prefix: '%s' got: '%s'", ti, tt.name, tt.err, err)
			}
			if err != "" && tt.err == "" {
				t.Errorf("%d. %s, exp no err got: '%s'", ti, tt.name, err)
			}
		})
	}
}

func TestMainLogsError(t *testing.T) {
	os.Args = []string{
		"app",
		"-db_path", "",
	}

	backup := os.Stdout
	f, err := ioutil.TempFile("", "stdout")
	logErr(t, err)
	defer func() {
		logErr(t, os.Remove(f.Name()))
	}()

	os.Stdout = f

	main()

	f.Close()
	os.Stdout = backup

	b, err := ioutil.ReadFile(f.Name())
	logErr(t, err)

	exp := "Error on "
	got := string(b)
	if !strings.Contains(got, exp) {
		t.Errorf("Exp err prefix: '%s' got: '%s'", exp, got)
	}
}

func errStr(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}

func logErr(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("Unexpected error, %v", err)
	}
}
