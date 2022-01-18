package precise_query

import (
	"build_VToken_index"
	"build_dictionary"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"
)

func MatchSearch(searchStr string, root *build_dictionary.TrieTreeNode, indexRoot *build_VToken_index.IndexTreeNode, qmin int, qmax int) []int {
	start2 := time.Now()
	var vgMap map[int][]string
	vgMap = make(map[int][]string)
	build_VToken_index.VGCons(root, qmin, qmax, searchStr, vgMap)
	fmt.Println(vgMap)
	var resArr []int
	preSeaPosition := 0
	var preInverPositionDis []int
	var nowInverPositionDis []int
	fields := strings.Fields(searchStr)
	for i := 0; i < len(fields); i++ { // 0 1 3   len(searchStr)
		tokenArr := vgMap[i]
		if tokenArr != nil {
			nowSeaPosition := i
			invertIndex = nil
			invertIndex2 = nil
			searchIndexTreeFromLeaves(tokenArr, indexRoot, 0)
			searchListsTreeFromLeaves(indexNode)
			invertIndex = append(invertIndex, invertIndex2...)
			invertIndex = RemoveSliceInvertIndex(invertIndex)
			sort.SliceStable(invertIndex, func(i, j int) bool {
				if invertIndex[i].Sid < invertIndex[j].Sid {
					return true
				}
				return false
			})
			if invertIndex == nil {
				return nil
			}
			if i == 0 {
				for j := 0; j < len(invertIndex); j++ {
					sid := invertIndex[j].Sid
					preInverPositionDis = append(preInverPositionDis, 0)
					nowInverPositionDis = append(nowInverPositionDis, invertIndex[j].Position)
					resArr = append(resArr, sid)
				}
			} else {
				for j := 0; j < len(resArr); j++ { //遍历之前合并好的resArr
					sidResArr := resArr[j]
					var k int
					for k = 0; k < len(invertIndex); k++ {
						sid := invertIndex[k].Sid
						if sidResArr == sid {
							nowInverPositionDis[j] = invertIndex[k].Position
							if nowInverPositionDis[j]-preInverPositionDis[j] == nowSeaPosition-preSeaPosition {
								break
							}
						}
					}
					if k == len(invertIndex) { //新的倒排表id不在之前合并好的结果集resArr 把此id从resArr删除
						resArr = append(resArr[:j], resArr[j+1:]...)
						preInverPositionDis = append(preInverPositionDis[:j], preInverPositionDis[j+1:]...)
						nowInverPositionDis = append(nowInverPositionDis[:j], nowInverPositionDis[j+1:]...)
						j--
					}
				}
			}
			preSeaPosition = nowSeaPosition
			//fmt.Println(preInverPositionDis)
			//fmt.Println(nowInverPositionDis)
			copy(preInverPositionDis, nowInverPositionDis)
		}
	}
	elapsed2 := time.Since(start2)
	fmt.Println("精确查询花费时间（ms）：", elapsed2)
	return resArr
}

var invertIndex []build_VToken_index.Inverted_index
var indexNode *build_VToken_index.IndexTreeNode

//查询当前串对应的倒排表（叶子节点）
func searchIndexTreeFromLeaves(tokenArr []string, indexRoot *build_VToken_index.IndexTreeNode, i int) {
	if indexRoot == nil {
		return
	}
	for j := 0; j < len(indexRoot.Children); j++ {
		if i < len(tokenArr)-1 && tokenArr[i] == indexRoot.Children[j].Data {
			searchIndexTreeFromLeaves(tokenArr, indexRoot.Children[j], i+1)
		}
		if i == len(tokenArr)-1 && tokenArr[i] == indexRoot.Children[j].Data { //找到那一层的倒排表
			for k := 0; k < len(indexRoot.Children[j].InvertedIndexList); k++ {
				invertIndex = append(invertIndex, *indexRoot.Children[j].InvertedIndexList[k])
			}
			indexNode = indexRoot.Children[j]
		}
	}
}

var invertIndex2 []build_VToken_index.Inverted_index

func searchListsTreeFromLeaves(indexNode *build_VToken_index.IndexTreeNode) {
	if indexNode != nil {
		for l := 0; l < len(indexNode.Children); l++ {
			if indexNode.Children[l].InvertedIndexList != nil {
				for k := 0; k < len(indexNode.Children[l].InvertedIndexList); k++ {
					invertIndex2 = append(invertIndex2, *indexNode.Children[l].InvertedIndexList[k])
				}
			}
			searchListsTreeFromLeaves(indexNode.Children[l])
		}
	}
}

func RemoveSliceInvertIndex(invertIndex []build_VToken_index.Inverted_index) (ret []build_VToken_index.Inverted_index) {
	n := len(invertIndex)
	for i := 0; i < n; i++ {
		state := false
		for j := i + 1; j < n; j++ {
			if j > 0 && reflect.DeepEqual(invertIndex[i], invertIndex[j]) {
				state = true
				break
			}
		}
		if !state {
			ret = append(ret, invertIndex[i])
		}
	}
	return
}
