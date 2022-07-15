package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
	"strings"
)

type aa struct {
	i int64
}

type bb struct {
	Aa *aa //8
	//Name string  //16
	//Age int //8
	//by byte //8
	//bys []byte //24
}

func main() {
}

type People interface {
	Show()
}

type Student struct {
	N string
}

func (receiver *Student) Show() {

}

func live() People {
	var s *Student
	return s
}

func defer1() {
	defer func() {
		fmt.Println("defer 1")
	}()
	defer func() {
		fmt.Println("defer 1.2")
	}()

}
func defer2() {
	defer func() {
		fmt.Println("defer 2")
	}()
}

func strToInt(str string) int {
	str = strings.TrimSpace(str)
	if len(str) == 0 {
		return 0
	}

	if len(str) > 10 {
		if str[0] == '-' {
			return math.MinInt32
		}
		return math.MaxInt32
	}
	n := 0
	fuhao := 1
	if str[0] == '+' || str[0] == '-' {
		if str[0] == '-' {
			fuhao = -1
		}
		str = str[1:]
	}
	for _, bb := range []byte(str) {
		bb -= '0'
		if bb > 9 {
			break
		}
		n = n*10 + int(bb)
	}
	return n * fuhao
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (hp) Push(interface{})     {}
func (hp) Pop() (_ interface{}) { return }
func smallestK(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	h := &hp{arr[:k]}
	heap.Init(h)
	for _, v := range arr[k:] {
		if h.IntSlice[0] > v {
			h.IntSlice[0] = v
			heap.Fix(h, 0)
		}
	}
	return h.IntSlice
}

func main1() {
	return
	//var node = &TreeNode{
	//	Val: 4,
	//	Left: &TreeNode{
	//		Val: 2,
	//		Left: &TreeNode{
	//			Val:   1,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &TreeNode{
	//			Val:   3,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &TreeNode{
	//		Val:  5,
	//		Left: nil,
	//		Right: &TreeNode{
	//			Val:   6,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}
	//isValidBST(node)

	//pz([]int{3,10,81,0})
	//return
	//b := make(chan struct{})
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//go func(close <-chan struct{}) {
	//	fmt.Println("dsfsdfds")
	//	wg.Done()
	//}(b)
	//wg.Wait()

}

//滑动窗口
func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	lenx := math.MaxInt32
	ansL, ansR := -1, -1
	//"ADOBECODEBANCAB", "ABC"
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < lenx {
				lenx = r - l + 1
				ansL, ansR = l, l+lenx
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]]--
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}

	return s[ansL:ansR]
}

func simplePath(x string) string {
	var p []string
	for _, nx := range strings.Split(x, "/") {
		if nx == ".." {
			if len(p) > 0 {
				p = p[:len(p)-1]
			}
		} else if nx != "." && nx != "" {
			p = append(p, nx)
		}
	}
	return "/" + strings.Join(p, "/")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var stack []*TreeNode
	inorder := math.MinInt64
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.Val <= inorder {
			return false
		}
		inorder = root.Val
		root = root.Right
	}
	return true
}

func quickSort(arr []int, left, right int) {
	if left < right {
		pivot := arr[right]
		j := left - 1
		for i := left; i < right; i++ {
			if arr[i] < pivot {
				j++
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
		arr[right], arr[j+1] = arr[j+1], arr[right]
		j += 1
		quickSort(arr, left, j-1)
		quickSort(arr, j+1, right)
	}
}

func tt(v int) int {

	if v == 2 {
		return 1
	}

	if v < 2 {
		return -1
	}

	t := v / 3
	t2 := v % 3

	return t + tt(t+t2)
}

func sortArray(nums []int) []int {
	// 计数排序，基于哈希思想的稳定外排序算法，空间换时间，时间O(n)，空间O(n)
	// 数据量大时，空间占用大
	// 空间换时间，通过开辟额外数据空间存储索引号记录数组的值和数组额个数
	// 思路：
	// 1.找出待排序的数组的最大值和最小值
	// 2.创建数组存放各元素的出现次数，先于[min, max]之间
	// 3.统计数组值的个数
	// 4.反向填充数组，填充时注意,num[i]=j+min，
	// j-前面需要略过的数的个数，两个维度，依次递增的数j++，一个是重复的数的计数j-不变
	if len(nums) == 0 {
		return nums
	}
	// 获取最大最小值
	minAndMax := func(nums []int) (min, max int) {
		minNum := math.MaxInt32
		maxNum := math.MinInt32
		for i := 0; i < len(nums); i++ {
			if nums[i] < minNum {
				minNum = nums[i]
			}
			if nums[i] > maxNum {
				maxNum = nums[i]
			}
		}
		return minNum, maxNum
	}
	min_, max_ := minAndMax(nums)
	// 中转数组存放遍历元素
	// 空间只需要min-max
	tmpNums := make([]int, max_-min_+1)
	// 遍历原数组
	for i := 0; i < len(nums); i++ {
		tmpNums[nums[i]-min_]++
	}
	// 遍历中转数组填入原数组
	j := 0
	for i := 0; i < len(nums); i++ {
		// 如果对应数字cnt=0，说明可以计入下一位数字
		for tmpNums[j] == 0 {
			j++
		}
		// 填入数字
		nums[i] = j + min_
		// 填一个数字，对应数字cnt--
		tmpNums[j]--
	}
	return nums
}
