package coderank


func GradingStudents(grades []int32) []int32 {
	gradesToReturn := []int32{}
	for _, g := range grades {
		if isAFallingGrade(g) {
			gradesToReturn = append(gradesToReturn, g)
		} else {
			nextMult := getNextMultipleOf5(g)
			if (nextMult - g) < 3 {
				gradesToReturn = append(gradesToReturn, nextMult)
			} else {
				gradesToReturn = append(gradesToReturn, g)
			}
		}
	}
	return gradesToReturn
}

func isAFallingGrade(g int32) bool {
	return g < 38
}

func getNextMultipleOf5(g int32) int32 {
	i := g
	isMultipleOf5 := false
	for !isMultipleOf5 {
		if i % 5 == 0 {
			isMultipleOf5 = true
		} else {
			i++
		}
	}
	return i
}
