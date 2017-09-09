package twofer

// One for you, One for me.
func ShareWith(you string) string {
	if you == "" {
		you = "you"
	}
	return "One for " + you + ", one for me."
}
