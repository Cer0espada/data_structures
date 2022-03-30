package Sorts

//Helper method that is the workhorse of the MergeSort Algorithm
//It creates and new array by comparing the individual values of each subarray and returning a sorted new array
func merge(left, right []int)([]int){
	result := []int{}
	leftIndex := 0
	rightIndex := 0

	for leftIndex < len(left) && rightIndex < len(right){
		if (left[leftIndex] < right[rightIndex]){
			result = append(result, left[leftIndex])
			leftIndex++
		}else{
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}
	return append(result, left[leftIndex], right[rightIndex])
}
///MergeSort 
//
//Divide and conquer Algorithm
//Works by recursively breaking down the input array into smaller arrays in which they are recursively sorted and then merged to get a final array
//
//Time Complexity is O(nlog(n))
func MergeSort(xi []int)([]int){
	if len(xi) == 1{
		return []int{}
	}

	mid := len(xi) /2

	left := xi[:mid]
	right := xi[mid:]

	MergeSort(left)
	MergeSort(right)
	return merge(MergeSort(left), MergeSort(right))
}


//QuickSort
//
//QuickSort is a divide and conquer algorithm, it uses a pivot to a random point in an array , and then use that point as a reference to select different points in a slice to compare and swap
//QuickSort also more memory efficient than mergesort as it uses the array in place
//Time Complexity is O(nlog(n))
func QuickSort(xi []int)([]int){

	if len(xi) ==0{
		return []int{}
	}

	left := []int{}
	right := []int{}
	pivot := xi[0]

	for i:=1 ; i <len(xi);i++{
		if xi[i] < pivot{
			left = append(left, xi[i])
		}else{
			right = append(right,xi[i])
		}
	}

	combination := append(QuickSort(left), QuickSort(right)...)
	
	return  append(combination, pivot)
}