// Package app is a tools set that respond for a http deal layer
package app

// StartApp map and run http listener main program
func StartApp() {
	// map handlers
	maps()
	// start middleware
	middleware()
	// run listener application
	run()
}
