package main

import (
	"fmt"
	"strings"
)

type CommandStatus struct {
	Command string
	Status  int
}

func NewCommandStatus(command string, status int) *CommandStatus {
	return &CommandStatus{command, status}
}

func (cs *CommandStatus) Format(msgFmt string) string {
	return format(msgFmt, replacements{
		"@cmd@":    cs.Command,
		"@status@": fmt.Sprintf("%d", cs.Status),
	})
}

type replacements map[string]string

func format(s string, m replacements) string {
	for k, v := range m {
		s = strings.Replace(s, k, v, 1)
	}
	return s
}
