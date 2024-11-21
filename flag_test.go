package looseflag

import (
	"fmt"
	"testing"
)

func printFlagSet(f *FlagSet) {
	fmt.Printf("beforeArgs: %+v\n", f.beforeArgs)
	fmt.Printf("afterArgs: %+v\n", f.afterArgs)

	for n, v := range f.options {
		fmt.Printf("options[%s] = %+v\n", n, v)
	}
}

func TestFlagSet_Parse(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		isBool  []string
		args    args
		wantErr bool
	}{
		{
			name: "base test",
			isBool: []string{
				"pack",
			},
			args: args{
				args: []string{
					"-o",
					"b026/_pkg_.a",
					"-trimpath",
					"b026=>",
					"-p",
					"internal/race",
					"-std",
					"-complete",
					"-buildid",
					"qnoO7tCfilLekDY7LXog/qnoO7tCfilLekDY7LXog",
					"-goversion",
					"go1.20.1",
					"-shared",
					"-c=4",
					"-nolocalimports",
					"-importcfg",
					"b026/importcfg",
					"-pack",
					"$GOPATH/src/internal/race/doc.go",
					"$GOPATH/src/internal/race/norace.go",
				},
			},
			wantErr: false,
		},
		{
			name: "base test",
			isBool: []string{
				"pack",
			},
			args: args{
				args: []string{
					"$GOPATH/src/internal/race/doc.go",
					"$GOPATH/src/internal/race/norace.go",
					"-o",
					"b026/_pkg_.a",
					"-trimpath",
					"b026=>",
					"-p",
					"internal/race",
					"-std",
					"-complete",
					"-buildid",
					"qnoO7tCfilLekDY7LXog/qnoO7tCfilLekDY7LXog",
					"-goversion",
					"go1.20.1",
					"-shared",
					"-c=4",
					"-nolocalimports",
					"-importcfg",
					"b026/importcfg",
					"-pack",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFlagSet(tt.name, 0)
			f.SetBoolVals(tt.isBool...)
			if err := f.Parse(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}

			if s, ok := f.GetString("o"); !ok || s != "b026/_pkg_.a" {
				t.Errorf("o = %v, want %v", s, "b026/_pkg_.a")
			}

			if s, ok := f.GetString("trimpath"); !ok || s != "b026=>" {
				t.Errorf("trimpath = %v, want %v", s, "b026=>")
			}

			if s, ok := f.GetString("p"); !ok || s != "internal/race" {
				t.Errorf("p = %v, want %v", s, "internal/race")
			}

			if s, ok := f.GetBool("std"); !ok || s != true {
				t.Errorf("std = %v, want %v", s, true)
			}

			if s, ok := f.GetBool("complete"); !ok || s != true {
				t.Errorf("complete = %v, want %v", s, true)
			}

			if s, ok := f.GetString("buildid"); !ok || s != "qnoO7tCfilLekDY7LXog/qnoO7tCfilLekDY7LXog" {
				t.Errorf("buildid = %v, want %v", s, "qnoO7tCfilLekDY7LXog/qnoO7tCfilLekDY7LXog")
			}

			if s, ok := f.GetString("goversion"); !ok || s != "go1.20.1" {
				t.Errorf("goversion = %v, want %v", s, "go1.20.1")
			}

			if s, ok := f.GetBool("shared"); !ok || s != true {
				t.Errorf("shared = %v, want %v", s, true)
			}

			if s, ok := f.GetInt("c"); !ok || s != 4 {
				t.Errorf("c = %v, want %v", s, 4)
			}

			if s, ok := f.GetString("importcfg"); !ok || s != "b026/importcfg" {
				t.Errorf("importcfg = %v, want %v", s, "b026/importcfg")
			}

			if s, ok := f.GetBool("nolocalimports"); !ok || s != true {
				t.Errorf("nolocalimports = %v, want %v", s, true)
			}

			if s, ok := f.GetBool("pack"); !ok || s != true {
				t.Errorf("pack = %v, want %v", s, true)
			}

			otherArgs := f.Args()
			if otherArgs[0] != "$GOPATH/src/internal/race/doc.go" || otherArgs[1] != "$GOPATH/src/internal/race/norace.go" {
				t.Errorf("Args error")
			}
		})
	}
}
