package bootstrap

import (
	"testing"
)

var alertColoredTests = []struct {
	vAlert      AlertColored
	vAlertIndex int
	want        string
}{
	{PrimaryB, 0, "Primary"},
	{SecondaryB, 1, "Secondary"},
	{SuccessB, 2, "Success"},
	{DangerB, 3, "Danger"},
	{WarningB, 4, "Warning"},
	{InfoB, 5, "Info"},
	{PrimaryB, 6, "Primary"},
}

func TestAlertColored(t *testing.T) {
	for _, tt := range alertColoredTests {
		got := tt.vAlert.String()
		if got != tt.want {
			t.Errorf("AlertColored: %v.String(): want[%v], got [%v]", tt.vAlert.String(), tt.want, got)
		}
		got2index := tt.vAlertIndex
		got2alert := NewAlertColored(got2index)
		got2string := got2alert.String()
		if got2string != tt.want {
			t.Errorf("AlertColored: %v.String(): want[%v], got [%v]", got2alert.String(), tt.want, got2string)
		}
		got3index := tt.vAlertIndex
		got3alert := NewAlertColored(got3index)
		got3string := got3alert.String()
		if got3string != tt.want {
			t.Errorf("AlertColored: %v.String(): want[%v], got [%v]", got3alert.String(), tt.want, got3string)
		}
	}
}

var alertColoredDivTests = []struct {
	vAlert AlertColored
	vText  string
	want   string
}{
	{PrimaryB, "foo bar", `<div class="alert alert-primary" role="alert">foo bar</div>`},
}

func TestAlertBorderedDiv(t *testing.T) {
	for _, tt := range alertColoredDivTests {
		got := tt.vAlert.DivHtml(tt.vText)
		if got != tt.want {
			t.Errorf("AlertColored: %v.DivHtml(): want[%v], got [%v]", tt.vAlert.String(), tt.want, got)
		}
	}
}
