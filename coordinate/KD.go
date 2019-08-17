package main
import "sort"

type KDNode struct {
	Left  []Geo
	Right []Geo
	T     Geo
}

type KDTree struct {
	Root KDNode
}

func (k KDTree)CreateKDTree(items []Geo, depth int) (kd KDNode) {
	if len(items) {
		return
	}
	sort.Sort(GeoXSort(items))
	currentIndex := len(items) / 2
	kd=KDNode{Left:,Right:,T:}
}