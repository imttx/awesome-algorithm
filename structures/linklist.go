package structures

type ListNode struct {
	Val  int
	Next *ListNode
}

type ListNodeWithRandom struct {
	Val    int
	Next   *ListNodeWithRandom
	Random *ListNodeWithRandom
}
