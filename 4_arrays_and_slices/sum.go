package arraysandslices

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAllTrails(numbersToSum ...[]int) (sums []int) {
	/*
	   In this implementation, we are worrying less about capacity. We start with an empty
	   slice sums and append to it the result of Sum as we work through the varargs.
	*/
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			/*
			   Slices can be sliced! The syntax is slice[low:high]. If you omit the value on one of the sides of the : it captures everything to that side of it.
			*/
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return
}
