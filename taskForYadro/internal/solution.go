package internal

import "sort"

func Checksortability(container [][]int32) bool {
	n := len(container)

	containerCapacities := make([]int32, n)
	typeDemands := make([]int32, n)
	for i := range container {
		for j, v := range container[i] {
			containerCapacities[i] += v
			typeDemands[j] += v
		}
	}

	sort.Slice(containerCapacities, func(i, j int) bool {
		return containerCapacities[i] < containerCapacities[j]
	})
	sort.Slice(typeDemands, func(i, j int) bool {
		return typeDemands[i] < typeDemands[j]
	})

	for i := range containerCapacities {
		if containerCapacities[i] != typeDemands[i] {
			return false
		}
	}

	return true
}
