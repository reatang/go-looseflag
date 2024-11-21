// Package looseflag 实现了宽松的命令行解析方案
// 目的主要是解决某些工具的参数动态变化，如果需要实现接收参数就会很麻烦的问题（说的就是 go build）
// 当前只实现了 string/bool/int 的解析
//
// 想要看到更多的描述内容, 请访问 https://github.com/reatang/go-looseflag
package looseflag // Package github.com/reatang/go-looseflag

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ErrLoose ErrorHandling = iota
)

var CommandLine = NewFlagSet("looseflag", ErrLoose)

type ErrorHandling int

type FlagSet struct {
	name          string
	errorHandling ErrorHandling

	beforeArgs []string
	afterArgs  []string

	isBoolArgs []string

	options map[string]any
}

func NewFlagSet(name string, errorHandling ErrorHandling) *FlagSet {
	return &FlagSet{
		name:          name,
		errorHandling: errorHandling,

		beforeArgs: make([]string, 0),
		afterArgs:  make([]string, 0),

		isBoolArgs: make([]string, 0),
		options:    make(map[string]any),
	}
}

func (f *FlagSet) SetBoolArgs(b ...string) {
	f.isBoolArgs = append(f.isBoolArgs, b...)
}

func (f *FlagSet) Parse(args []string) error {
	for {
		a := args[0]
		_, s := parseOptions(a)
		if s != parseOptionNameIsValue {
			break
		}

		f.beforeArgs = append(f.beforeArgs, args[0])
		args = args[1:]
	}

	next := func() {
		if len(args) > 0 {
			args = args[1:]
		}
	}

	for len(args) > 0 {
		name, s := parseOptions(args[0])
		if s == parseOptionIsName {
			if eqIndex := strings.Index(name, "="); eqIndex >= 0 {
				v := name[eqIndex+1:]
				if b, ok := parseBool(v); ok {
					f.options[name[:eqIndex]] = b
				} else {
					f.options[name[:eqIndex]] = v
				}
			} else if sliceContains(f.isBoolArgs, name) {
				f.options[name] = true
			} else {
				if len(args) == 0 {
					f.options[name] = true
					return nil
				}

				next()

				value, s2 := parseOptions(args[0])
				if s2 != parseOptionNameIsValue {
					f.options[name] = true
					continue
				}

				if b, ok := parseBool(value); ok {
					f.options[name] = b
				} else {
					f.options[name] = value
				}
			}
		} else if s == parseOptionNameIsValue {
			f.afterArgs = append(f.afterArgs, name)
		} else if s == parseOptionStopParse {
			f.afterArgs = append(f.afterArgs, args[1:]...)
			return nil
		} else {
			return fmt.Errorf("%s: unknown option", f.name)
		}

		next()
	}

	return nil
}

const (
	parseOptionErr = iota - 1
	parseOptionIsName
	parseOptionNameIsValue
	parseOptionStopParse
)

func parseOptions(optName string) (string, int) {
	if optName == "" {
		return "", parseOptionErr
	}

	if optName[0] == '-' {
		if optName[1] == '-' {
			if len(optName) == 2 {
				return "", parseOptionStopParse
			} else {
				return optName[2:], parseOptionIsName
			}
		}

		return optName[1:], parseOptionIsName
	}

	return optName, parseOptionNameIsValue
}

func parseBool(value string) (bool, bool) {
	if value == "true" || value == "1" || value == "yes" || value == "on" {
		return true, true
	} else if value == "false" || value == "0" || value == "no" || value == "off" {
		return false, true
	}

	return false, false
}

func (f *FlagSet) GetString(name string) (string, bool) {
	if v, ok := f.options[name]; ok {
		switch _v := v.(type) {
		case string:
			return _v, true
		}
	}

	return "", false
}

func (f *FlagSet) GetBool(name string) (bool, bool) {
	if v, ok := f.options[name]; ok {
		switch _v := v.(type) {
		case bool:
			return _v, true
		}
	}

	return false, false
}

func (f *FlagSet) GetInt(name string) (int, bool) {
	if v, ok := f.options[name]; ok {
		switch _v := v.(type) {
		case int:
			return _v, true
		case string:
			atoi, err := strconv.Atoi(_v)
			if err != nil {
				return 0, false
			}
			return atoi, true
		}
	}

	return 0, false
}

func (f *FlagSet) Args() []string {
	return append(f.beforeArgs, f.afterArgs...)
}

func sliceContains(s []string, v string) bool {
	for i := range s {
		if v == s[i] {
			return true
		}
	}

	return false
}

func Parse() error {
	return CommandLine.Parse(os.Args[1:])
}

func GetString(name string) (string, bool) {
	return CommandLine.GetString(name)
}

func GetBool(name string) (bool, bool) {
	return CommandLine.GetBool(name)
}

func GetInt(name string) (int, bool) {
	return CommandLine.GetInt(name)
}

func Args() []string {
	return CommandLine.Args()
}
