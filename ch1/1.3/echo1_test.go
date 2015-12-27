// Exercise 1.3
// Experiment to measure the difference in running time between our potentially
// inefficient versions and the one that uses strings.Join (Section 1.6
// illustrates part of the time package and Section 11.4 shows how to write
// benchmark tests for systematic performance evaluation.)
package ch1

import "testing"

func BenchmarkEcho1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo1()
	}
}
