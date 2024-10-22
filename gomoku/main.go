package main

func main() {

	var UnsortedList = []int{5, 3, 1, 4, 2}

	var SortedList = Bubblesort(UnsortedList)
	println(SortedList)

}

func Bubblesort(UnsortedList []int) []int {

	var SortedList = UnsortedList
	for i := 0; i < len(SortedList); i++ {

		for j := 0; j < len(SortedList)-1; j++ {

			if SortedList[j] > SortedList[j+1] {

				SortedList[j], SortedList[j+1] = SortedList[j+1], SortedList[j]

			}

		}

	}

	return SortedList
}
