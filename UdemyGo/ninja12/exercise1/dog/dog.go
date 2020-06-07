// Package dog helps you calculate dogs years.
package dog

import "fmt"

// Years take human years and turns them into dogs years.
func Years(year int) string {
	return fmt.Sprintf("Dog Year is: %v", year*7)
}
