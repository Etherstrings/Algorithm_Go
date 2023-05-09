package L2383

func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	var Esum = 0
	for _, v := range energy {
		Esum += v
	}
	var sum = 0
	if initialEnergy <= Esum {
		sum += Esum - initialEnergy + 1
	}
	for _, v := range experience {
		if initialExperience <= v {
			var target = v + 1
			sum += target - initialExperience
			initialExperience = target + v
		} else {
			initialExperience = initialExperience + v
		}
	}
	return sum
}
