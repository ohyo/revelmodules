package tests

import "github.com/revel/revel/testing"

// WikiTest is
type WikiTest struct {
	testing.TestSuite
}

// Before is
func (t *WikiTest) Before() {
	println("Wiki Set up")
}

// TestThatWikiIndexPageWorks is
func (t *WikiTest) TestThatWikiIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

// After is
func (t *WikiTest) After() {
	println("Wiki Tear down")
}
