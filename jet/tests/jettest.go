package tests

import (
	"github.com/revel/revel/testing"
)

// JetTest is
type JetTest struct {
	testing.TestSuite
}

// Before is
func (t *JetTest) Before() {
	println("Jet Set up")
}

// TestThatJetIndexPageWorks is
func (t *JetTest) TestThatJetIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

// After is
func (t *JetTest) After() {
	println("Jet Tear down")
}
