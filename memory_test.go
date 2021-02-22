package gmonitor

import "testing"

func TestMemory_Stat(t *testing.T) {
	v := &Memory{}
	err := v.Stat()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("total:", v.TotalText)
	t.Log("available:", v.AvailableText)
	t.Log("used:", v.UsedText)
	t.Logf("used percent: %.1f%%", v.UsedPercent)
}
