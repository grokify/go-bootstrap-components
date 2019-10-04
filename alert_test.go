package bootstrap

import (
	"testing"
)

var alertBorderedTests = []struct {
	vAlert      AlertBordered
	vAlertIndex int
	want        string
}{
	{PrimaryB, 0, "Primary"},
	{SecondaryB, 1, "Secondary"},
	{SuccessB, 2, "Success"},
	{DangerB, 3, "Danger"},
	{WarningB, 4, "Warning"},
	{InfoB, 5, "Info"},
	{DarkB, 6, "Dark"},
	{PrimaryB, 7, "Primary"},
}

func TestAlertBordered(t *testing.T) {
	for _, tt := range alertBorderedTests {
		got := tt.vAlert.String()
		if got != tt.want {
			t.Errorf("AlertBodered: %v.String(): want[%v], got [%v]", tt.vAlert.String(), tt.want, got)
		}
		got2index := tt.vAlertIndex
		got2alert := NewAlertBordered(got2index)
		got2string := got2alert.String()
		if got2string != tt.want {
			t.Errorf("AlertBodered: %v.String(): want[%v], got [%v]", got2alert.String(), tt.want, got2string)
		}
		got3index := tt.vAlertIndex
		got3alert := AlertBordered(got3index)
		got3string := got3alert.String()
		if got3string != tt.want {
			t.Errorf("AlertBodered: %v.String(): want[%v], got [%v]", got3alert.String(), tt.want, got3string)
		}
	}
}

var alertBorderedDivTests = []struct {
	vAlert AlertBordered
	vText  string
	want   string
}{
	{PrimaryB, "foo bar", `<div class="alert alert-primary" role="alert">foo bar</div>`},
}

func TestAlertBorderedDiv(t *testing.T) {
	for _, tt := range alertBorderedDivTests {
		got := tt.vAlert.DivHtml(tt.vText)
		if got != tt.want {
			t.Errorf("AlertBodered: %v.DivHtml(): want[%v], got [%v]", tt.vAlert.String(), tt.want, got)
		}
	}
}
