// Package qselect implements QuickSelect
/*

The implementation here is based on the Rosetta Code implementation
    https://rosettacode.org/wiki/Quickselect_algorithm
plus some enhancements.

*/
package qselect

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

func Select(data Interface, k int) {

	left, right := 0, data.Len()-1

	for {
		// insertion sort for small ranges
		if right-left <= 20 {
			for i := left + 1; i <= right; i++ {
				for j := i; j > 0 && data.Less(j, j-1); j-- {
					data.Swap(j, j-1)
				}
			}

			return
		}

		// median-of-three to choose pivot
		pivotIndex := left + (right-left)/2
		if data.Less(right, left) {
			data.Swap(right, left)
		}
		if data.Less(pivotIndex, left) {
			data.Swap(pivotIndex, left)
		}
		if data.Less(right, pivotIndex) {
			data.Swap(right, pivotIndex)
		}

		// partition
		data.Swap(left, pivotIndex)
		ll := left + 1
		rr := right
		for ll <= rr {
			for ll <= right && data.Less(ll, left) {
				ll++
			}
			for rr >= left && data.Less(left, rr) {
				rr--
			}
			if ll <= rr {
				data.Swap(ll, rr)
				ll++
				rr--
			}
		}
		data.Swap(left, rr) // swap into right place
		pivotIndex = rr

		if k == pivotIndex {
			return
		}

		if k < pivotIndex {
			right = pivotIndex - 1
		} else {
			left = pivotIndex + 1
		}
	}
}
