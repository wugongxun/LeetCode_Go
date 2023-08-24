package main

func main() {

}

type NestedIterator struct {
	nums  []int
	index int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	nestedIterator := NestedIterator{nums: make([]int, 0), index: 0}
	nestedIterator.addToNums(nestedList)
	return &nestedIterator
}

func (this *NestedIterator) addToNums(nestedList []*NestedInteger) {
	for i := 0; i < len(nestedList); i++ {
		if nestedList[i].IsInteger() {
			this.nums = append(this.nums, nestedList[i].GetInteger())
		} else {
			this.addToNums(nestedList[i].GetList())
		}
	}
}

func (this *NestedIterator) Next() int {
	val := this.nums[this.index]
	this.index++
	return val
}

func (this *NestedIterator) HasNext() bool {
	return this.index < len(this.nums)
}

type NestedInteger struct {
}

func (this NestedInteger) IsInteger() bool           {}
func (this NestedInteger) GetInteger() int           {}
func (n *NestedInteger) SetInteger(value int)        {}
func (this *NestedInteger) Add(elem NestedInteger)   {}
func (this NestedInteger) GetList() []*NestedInteger {}
