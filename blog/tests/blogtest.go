package tests

import "github.com/revel/revel"

// BlogTest is
type BlogTest struct {
	revel.TestSuite
}

// Before is
func (t *BlogTest) Before() {
	println("Blog Set up")
}

// TestThatBlogIndexPageWorks is
func (t BlogTest) TestThatBlogIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

// After is
func (t *BlogTest) After() {
	println("Blog Tear down")
}
