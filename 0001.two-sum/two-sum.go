package Problem0001

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, a := range nums {
		//a+b=target  m[a] = i
		//if ok == true 说明i之前存在nums[j] == b
		if j, ok := m[target-a]; ok {
			return []int{j, i}
		}
		m[a] = i
	}
	return nil
}
