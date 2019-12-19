package main

import (
        "fmt"
        "os"
	"testing"
)

var cfg Config
var f *os.File

func TestReadFile(t *testing.T) {
    cases := []struct {Name string; A string; Expected *os.File} {
	{"read a yaml file", "testConfigs/readme.yml", f},
	{"read a txt file", "testConfigs/readme.txt", f},
	{"read a json", "testConfigs/readme.json", f},
        {"read a toml", "testConfigs/readme.toml", f},
        {"read a md file", "testConfigs/readme.md", f},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("read(%s) ", testCase.A), func(t *testing.T) {
            actual := readFile(testCase.A)
	    if (actual == testCase.Expected) {t.Errorf("TestCompare %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}

func TestParseYaml(t *testing.T) {
    cases := []struct {Name string; A *Config; B *os.File; Expected bool} {
	{"parse an empty yaml file", &cfg, readFile("testConfigs/empty.yml"), true},
	{"parse a standard odin yaml file", &cfg, readFile("testConfigs/prune_containers.yml"), true},
	{"parse a large yaml file", &cfg, readFile("testConfigs/large.yml"), true},
    }
    for i, testCase := range cases {
        t.Run(fmt.Sprintf("%v.parse(%s) ", testCase.B, testCase.A), func(t *testing.T) {
            actual := parseYaml(testCase.A, testCase.B)
	    if (actual != testCase.Expected) {t.Errorf("TestCompare %d failed - expected: '%v' got: '%v'", i+1, actual, testCase.Expected)}
        })
    }
}
