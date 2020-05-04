package tests

import "github.com/revel/revel"

// DocTest is
type DocTest struct {
	revel.TestSuite
}

// Before is
func (t *DocTest) Before() {
	println("Set up")
}

// TestThatDocIndexPageWorks is
func (t DocTest) TestThatDocIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

// After is
func (t *DocTest) After() {
	println("Tear down")
}
