package bootstrap

import (
	"fmt"
	"strings"

	"github.com/grokify/mogo/math/bigint"
)

type Alert int

const alertPrefix string = "alert-"

const (
	Primary Alert = iota
	Secondary
	Success
	Danger
	Warning
	Info
	Light
	Dark
)

var alerts = [...]string{
	"Primary",
	"Secondary",
	"Success",
	"Danger",
	"Warning",
	"Info",
	"Light",
	"Dark",
}

func NewAlert(idx int) Alert {
	return Alert(bigint.Int64Mod(int64(idx), int64(len(alerts))))
}

// Class returns the class name of the alert ("Primary", "Secondary", ...).
func (a Alert) String() string {
	//if Primary <= a && a <= Dark {
	if a >= Primary && a <= Dark {
		return alerts[a]
	}
	a2 := NewAlert(int(a))
	return alerts[a2]
}

// Class returns the class name of the alert ("Primary", "Secondary", ...).
func (a Alert) Class() string {
	return alertPrefix + strings.ToLower(a.String())
}

type AlertColored int

const (
	PrimaryC AlertColored = iota
	SecondaryC
	SuccessC
	DangerC
	WarningC
	InfoC
)

var alertscolored = [...]string{
	"Primary",
	"Secondary",
	"Success",
	"Danger",
	"Warning",
	"Info",
}

func NewAlertColored(idx int) AlertColored {
	return AlertColored(bigint.Int64Mod(int64(idx), int64(len(alertscolored))))
}

// Class returns the class name of the alert ("Primary", "Secondary", ...).
func (a AlertColored) String() string {
	if a >= PrimaryC && a <= InfoC {
		return alertscolored[a]
	}
	a2 := NewAlertColored(int(a))
	return alertscolored[a2]
}

// Class returns the class name of the alert ("Primary", "Secondary", ...).
func (a AlertColored) Class() string {
	return alertPrefix + strings.ToLower(a.String())
}

const alertDivFormat string = `<div class="alert alert-%s" role="alert">`

// Class returns the class name of the alert ("Primary", "Secondary", ...).
func (a AlertColored) DivHTML(innerHTML string) string {
	begin := fmt.Sprintf(alertDivFormat, strings.ToLower(a.String()))
	return begin + innerHTML + "</div>"
}
