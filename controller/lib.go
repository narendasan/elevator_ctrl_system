package controller

/*FindMaxMin (function)
 * return the maximum and the minimum of an unsorted list
 * @param {[]int} elements slice of ints
 * @return {int} smallest min
 * @return {int} biggest max
 */
func FindMaxMin(x []int) (smallest, biggest int) {
    smallest, biggest = x[0], x[0]
    for _, v := range x {
        if v > biggest {
            biggest = v
        }
        if v < smallest {
            smallest = v
        }
    }
    return
}

/*RemoveDuplicates (function)
 * Remove duplicates entries from a slice
 * @param {[]int} elements slice of ints
 * @return {[]int} deduplicated slice
 */
func RemoveDuplicates(elements []int) []int {
    // Use map to record duplicates as we find them.
    encountered := map[int]bool{}
    result := []int{}

    for v := range elements {
        if encountered[elements[v]] == true {
	         // Do not add duplicate.
	    } else {
	        // Record this element as an encountered element.
	        encountered[elements[v]] = true
	        // Append to result slice.
	        result = append(result, elements[v])
	   }
    }
    // Return the new slice.
    return result
}

/*Abs (function)
 * return absolute value of int
 * @param {int} x
 * @return {int} absolute value of x
 */
func Abs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}
