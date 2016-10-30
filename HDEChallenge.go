/*HDE Challenge 003
* Author: Matthew Jones
*
* Calculates the sum of squares given by standard input. Does not use FOR statement.
*/

package main

import "fmt"

func main() {
	//Variable declaration for the number of cases; 1 <= N <= 100
	//Unsigned int8 would be most suitable for this value
	var N uint8;
	fmt.Scan(&N);
	
	//Slice created to contain the calculated sums of squares; length equal to N, the first input on command line
	var sumsOfSquares = make([]int, N);
	
	//Populates the sumsOfSquares array with the sums of squares from the user input
	populateSums(0, sumsOfSquares);
	
	//Prints out every sum of squares calculated
	printArray(0, sumsOfSquares);
}

/*Prints out every value in a given array without FOR-looping
*
* Parameters: currentIndex - The current index to print out
*			  array - The array to print values from
*/
func printArray(currentIndex uint8, array []int){
	if currentIndex < uint8(len(array)){
		fmt.Println(array[currentIndex]);
		printArray(currentIndex + 1, array);
	}
}

/*Populates the empty sums array with the sum of the squares calculated
*
* Parameters: currentIndex - The current index of the array to populate
*			  sums - The array to populate; passed by reference automatically
*/
func populateSums(currentInt uint8, sums []int){
	if currentInt < uint8(len(sums)){
		var X int;
		fmt.Scan(&X);
		Y := make([]int8, X);
		fillArray(0, Y);
		sums[currentInt] = sumOfSquaresArray(0, 0, Y);
		populateSums(currentInt + 1, sums);
	}
}

/*Fills an empty array with the next X values entered by standard input,
* where X is the length of the array
*
* Paramaters: currentInt - The current index of the array to fill
			  array - The array to fill
*/
func fillArray(currentInt uint8, array []int8){
	if currentInt < uint8(len(array)){
		fmt.Scan(&array[currentInt]);
		fillArray(currentInt + 1, array);
	}
}

/*Iterates through an array and returns the sum of the squares of each 
* index (excluding negative values) and returns the final sum
*
* Parameters: runningTotal - The current total of the sum of squares
*			  currentIndex - Current index that is being added
			  array - The array being summed
*/
func sumOfSquaresArray(runningTotal int, currentIndex uint8, array []int8) int{
	if currentIndex < uint8(len(array)){
		if array[currentIndex] <= 0{
			return (runningTotal + sumOfSquaresArray(runningTotal, currentIndex + 1, array));
		} else{
			return (runningTotal + (int(array[currentIndex]) * int(array[currentIndex])) + sumOfSquaresArray(runningTotal, currentIndex + 1, array));
		}
	}
	
	return runningTotal;
}