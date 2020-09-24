package main

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strconv"
	"sync"
)

type NodeType []uint32

func (n NodeType) Len() int {
	return len(n)
}

func (n NodeType) Less(i, j int) bool {
	return n[i] < n[j]
}

func (n NodeType) Swap(i, j int) {
	n[j], n[i] = n[i], n[j]
}

// VirtualNode ..VirtualNode
// NewInstance
// getHashCode
// sortByHash
// getvirtualNode
// addVirtualNode
// removeVirtualNode
type VirtualNode struct {
	VirtualNodes map[uint32]*ServiceNode
	NodeKeys     NodeType
	sync.RWMutex
}

func (v *VirtualNode) getHashCode(nodeHash string) uint32 {
	//crc32方式hash code
	//return crc32.ChecksumIEEE([]byte(nodeHash))
	m := md5.New()
	m.Write([]byte(nodeHash))

	md5Str := hex.EncodeToString(m.Sum(nil))

	var h int = 0
	byteHash := []byte(md5Str)
	for i := 0; i < 32; i++ {
		h <<= 8
		h |= int(byteHash[i]) & 0xFF
	}
	return uint32(h)
}

func (v *VirtualNode) sortByHash() {
	v.NodeKeys = nil
	for k := range v.VirtualNodes {
		v.NodeKeys = append(v.NodeKeys, k)
	}
	sort.Sort(v.NodeKeys)
}

func (v *VirtualNode) getVirtualNode(routKey string) *ServiceNode {

	//并发读写map-加读锁,可并发读不可同时写
	v.RLock()
	defer v.RUnlock()

	index := 0
	hashCode := v.getHashCode(routKey)

	i := sort.Search(len(v.NodeKeys), func(i int) bool { return v.NodeKeys[i] > hashCode })

	//当i大于下标最大值时,证明没找到, 给到第0个虚拟节点, 当i小于node节点数时, index为当前节点
	if i < len(v.NodeKeys) {
		index = i
	} else {
		index = 0
	}

	return v.VirtualNodes[v.NodeKeys[index]]

}

func (v *VirtualNode) addVirtualNode(node *ServiceNode, virtualNum uint) {
	v.Lock()
	defer v.Unlock()

	for i := uint(0); i < virtualNum; i++ {
		hashStr := node.Ip + ":" + node.Port + ":" + strconv.Itoa(int(i))
		v.VirtualNodes[v.getHashCode(hashStr)] = node
	}

	v.sortByHash()
}

func (v *VirtualNode) removeVirtualNode(node *ServiceNode, virtualNum uint) {
	v.Lock()
	defer v.Unlock()

	for i := uint(0); i < virtualNum; i++ {
		hashStr := node.Ip + ":" + node.Port + ":" + strconv.Itoa(int(i))
		delete(v.VirtualNodes, v.getHashCode(hashStr))
	}

	v.sortByHash()
}

var virtualServiceNode = NewVirtualNode()

func NewVirtualNode() *VirtualNode {
	return &VirtualNode{
		VirtualNodes: map[uint32]*ServiceNode{},
	}
}
