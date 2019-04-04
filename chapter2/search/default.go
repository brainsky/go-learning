package search

// defaultMatcher implement default matcher
type defaultMatcher struct{}

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

// Search implement default matcher
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
