package lib

// slice内に指定の数値があるかをチェックして
// 要素番号を返す関数
func Index(slice []int, item int) int {
	for i := range slice {
		if slice[i] == item {
			return i
		}
	}
	return -1
}
