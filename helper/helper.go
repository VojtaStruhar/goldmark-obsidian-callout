package helper

import (
	"errors"
)

func IndexOf(byteArray []byte, target byte) (int, error) {
	for i, b := range byteArray {
		if b == target {
			return i, nil
		}
	}
	return -1, errors.New("index not found")
}

type CalloutType int

const (
	Note CalloutType = iota
	Info
	Abstract
	Todo
	Tip
	Success
	Question
	Warning
	Failure
	Danger
	Bug
	Example
	Quote
)

// These are the string values that dictate the type of the callout. Some have aliases.
// https://help.obsidian.md/Editing+and+formatting/Callouts
var CalloutTypeMapping = map[string]CalloutType{
	"!note":      Note, // This is also the default one, if the type is unrecognized
	"!info":      Info,
	"!abstract":  Abstract,
	"!summary":   Abstract,
	"!tldr":      Abstract,
	"!todo":      Todo,
	"!tip":       Tip,
	"!hint":      Tip,
	"!important": Tip,
	"!success":   Success,
	"!check":     Success,
	"!done":      Success,
	"!question":  Question,
	"!help":      Question,
	"!faq":       Question,
	"!warning":   Warning,
	"!caution":   Warning,
	"!attention": Warning,
	"!failure":   Failure,
	"!fail":      Failure,
	"!missing":   Failure,
	"!danger":    Danger,
	"!error":     Danger,
	"!bug":       Bug,
	"!example":   Example,
	"!quote":     Quote,
	"!cite":      Quote,
}

var CalloutTypeStringMapping = map[CalloutType]string{
	Note:     "Note",
	Info:     "info",
	Tip:      "tip",
	Abstract: "abstract",
	Todo:     "todo",
	Success:  "success",
	Question: "question",
	Warning:  "warning",
	Failure:  "failure",
	Danger:   "danger",
	Bug:      "bug",
	Example:  "example",
	Quote:    "quote",
}
