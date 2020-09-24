### 回忆

还记得刚毕业入职到新公司的时候, 我的上级领导与前端同学解释后端技术栈庞杂. 大概记得举了一个例子是 “如何多台机器提供数据缓存存储服务?” , 扭头问了我一下, 当时直接说使用 hash取模 的方式分摊数据。 

接着我肯定被追问一台机器挂了怎么办, 怎么减少节点挂掉的影响, 结果是被鄙视了, 从那以后也就记住了 一致性hash 这个词. 

虽然工作时间也不短了, 但是现在再问我 一致性hash算法 究竟是啥, 我大概也只能回答出 一个圆环, 环里有很多虚拟节点, key hash后定位到对应的虚拟节点, 确从来没有自己动手写过一行代码.

老骥伏枥、志在千里. 我们开始吧~

### 一致性hash算法

一致性哈希算法在1997年由麻省理工学院的Karger等人在解决分布式Cache中提出的，设计目标是为了解决因特网中的热点(Hot spot)问题，初衷和CARP十分类似。一致性哈希修正了CARP使用的简单哈希算法带来的问题，使得DHT可以在P2P环境中真正得到应用.

一致性hash在数据存储领域中有广泛的应用, 目的主要是减少数据倾斜问题, 在节点失效、节点增加时, 只需影响少量数据.



![img](https://tva1.sinaimg.cn/large/007S8ZIlgy1gios56ih4nj30hz0dddhp.jpg)



我们看上图, 为一个环, 在环中我们根据hash计算放入4个node节点

我们又根据键值计算结果放入到对应离他最近的下一个节点.



![img](https://tva1.sinaimg.cn/large/007S8ZIlgy1gios5dke73j30jk0djgnu.jpg)



当我们新增node5节点时计算hash值放入环中, 仅将 node4 中部分数据(hash值小于node5) 重新定位到 node5 即可.

> 当移除节点node4时, 也仅将 node4 数据移入下一个节点 node3 即可.

思考: node4失效后, node4数据压力全部给到node3, node3的压力增大, 会不会发生链条反应, 导致所有节点崩溃?

这时我们需要增加虚拟节点来分担 node3 压力, 将实体节点通过 hash计算 分散更多的分布到环上, 相对来说数据 hash key  更能随机到不同的节点上, 理想状态下当其中一个节点失效时, 多个节点分摊数据压力

### 逻辑实现

MyServiceNode节点操作

```
package main

import (
    "fmt"
    "strconv"
)

//service node 结构体定义
type ServiceNode struct {
    Ip   string
    Port string
}
// 返回service node实例
func NewServiceNode(ip string, port string) *ServiceNode {
    return &ServiceNode{
        Ip:   ip,
        Port: port,
    }
}

func main() {

    //实例化三个实体节点
    node1 := NewServiceNode("127.0.0.1", "3305")
    node2 := NewServiceNode("127.0.0.1", "3306")
    node3 := NewServiceNode("127.0.0.1", "3307")

    //添加对应的虚拟化节点数1、1、3
    virtualNodeService.addVirtualNode(node1, 3)
    virtualNodeService.addVirtualNode(node2, 3)
    virtualNodeService.addVirtualNode(node3, 3)

    //打印节点列表
    PrintRouteList()

    //循环测试
    for i := 0; i < 20; i++ {
        //移除node2节点
        virtualNodeService.removeVirtualNode(node2, 3)
        //获取对应节点地址
        serviceNode := virtualNodeService.getVirtualNodel("get_cc" + strconv.Itoa(i))
        fmt.Println(serviceNode.Ip + ":" + serviceNode.Port)
    }
}

//打印节点列表
func PrintRouteList() {
    for key, val := range virtualNodeService.VirtualNodes {
        fmt.Println("key=" + strconv.Itoa(int(key)) + ", host=" + val.Ip + ":" + val.Port)
    }
}

复制代码
```

virtualNodeService 虚拟化节点

```
package main

import (
    "crypto/md5"
    "encoding/hex"
    "sort"
    "strconv"
    "sync"
)

var virtualNodeService = NewVirtualNode()

type NodeType []uint32

//Len()
func (s NodeType) Len() int {
    return len(s)
}

//Less():成绩将有低到高排序
func (s NodeType) Less(i, j int) bool {
    return s[i] < s[j]
}

//Swap()
func (s NodeType) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

//虚拟节点结构定义
type VirtualNode struct {
    VirtualNodes map[uint32]*ServiceNode
    NodeKeys     NodeType
    sync.RWMutex
}

//实例化虚拟节点对象
func NewVirtualNode() *VirtualNode {
    return &VirtualNode{
        VirtualNodes: map[uint32]*ServiceNode{},
    }
}

//添加虚拟节点
func (v *VirtualNode) addVirtualNode(serviceNode *ServiceNode, virtualNum uint) {

    //并发读写map-加锁
    v.Lock()
    defer v.Unlock()

    for i := uint(0); i < virtualNum; i++ {
        hashStr := serviceNode.Ip + ":" + serviceNode.Port + ":" + strconv.Itoa(int(i))
        v.VirtualNodes[v.getHashCode(hashStr)] = serviceNode
    }

    //虚拟节点hash值排序
    v.sortHash()
}

//移除虚拟节点
func (v *VirtualNode) removeVirtualNode(serviceNode *ServiceNode, virtualNum uint) {
    //并发读写map-加锁
    v.Lock()
    defer v.Unlock()

    for i := uint(0); i < virtualNum; i++ {
        hashStr := serviceNode.Ip + ":" + serviceNode.Port + ":" + strconv.Itoa(int(i))
        delete(v.VirtualNodes, v.getHashCode(hashStr))
    }

    v.sortHash()
}

//hash数值排序
func (v *VirtualNode) sortHash() {
    v.NodeKeys = nil
    for k := range v.VirtualNodes {
        v.NodeKeys = append(v.NodeKeys, k)
    }
    sort.Sort(v.NodeKeys)
}

//获取虚拟节点(二分查找)
func (v *VirtualNode) getVirtualNodel(routeKey string) *ServiceNode {
    //并发读写map-加读锁,可并发读不可同时写
    v.RLock()
    defer v.RUnlock()

    index := 0
    hashCode := v.getHashCode(routeKey)
    i := sort.Search(len(v.NodeKeys), func(i int) bool { return v.NodeKeys[i] > hashCode })
    //当i大于下标最大值时,证明没找到, 给到第0个虚拟节点, 当i小于node节点数时, index为当前节点
    if i < len(v.NodeKeys) {
        index = i
    } else {
        index = 0
    }

    //返回具体节点
    return v.VirtualNodes[v.NodeKeys[index]]
}

//获取hash code(采用md5字符串后计算)
func (v *VirtualNode) getHashCode(nodeHash string) uint32 {
    //crc32方式hash code
    //return crc32.ChecksumIEEE([]byte(nodeHash))
    md5 := md5.New()
    md5.Write([]byte(nodeHash))
    md5Str := hex.EncodeToString(md5.Sum(nil))

    h := 0
    byteHash := []byte(md5Str)
    for i := 0; i < 32; i++ {
        h <<= 8
        h |= int(byteHash[i]) & 0xFF
    }
    return uint32(h)
}

复制代码
```


作者：小宇渣渣渣
链接：https://juejin.im/post/6871260529801822216
来源：掘金
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。