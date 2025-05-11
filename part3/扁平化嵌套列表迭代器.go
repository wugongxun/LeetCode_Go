package main

func main() {

}

type NestedIterator struct {
	nums []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	var nums []int
	var dfs func([]*NestedInteger)
	dfs = func(nestedList []*NestedInteger) {
		for _, nest := range nestedList {
			if nest.IsInteger() {
				nums = append(nums, nest.GetInteger())
			} else {
				dfs(nest.GetList())
			}
		}
	}
	dfs(nestedList)
	return &NestedIterator{nums}
}

func (nestedIterator *NestedIterator) Next() int {
	v := nestedIterator.nums[0]
	nestedIterator.nums = nestedIterator.nums[1:]
	return v
}

func (nestedIterator *NestedIterator) HasNext() bool {
	return len(nestedIterator.nums) > 0
}

type NestedInteger struct {
	isInteger bool
	integer   int
	list      []*NestedInteger
}

func (nestedInteger *NestedInteger) IsInteger() bool {
	return nestedInteger.isInteger
}
func (nestedInteger *NestedInteger) GetInteger() int {
	return nestedInteger.integer
}
func (nestedInteger *NestedInteger) SetInteger(value int) {
	nestedInteger.integer = value
}
func (nestedInteger *NestedInteger) Add(elem NestedInteger) {
	if elem.IsInteger() {
		nestedInteger.list = append(nestedInteger.list, &NestedInteger{true, elem.GetInteger(), nil})
	} else {
		nestedInteger.list = append(nestedInteger.list, elem.list...)
	}
}
func (nestedInteger *NestedInteger) GetList() []*NestedInteger {
	return nestedInteger.list
}
