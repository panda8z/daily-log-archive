## 什么是哈希表？

资料：[哈希表_百度百科](https://baike.baidu.com/item/%E5%93%88%E5%B8%8C%E8%A1%A8/5981869?fr=aladdin)

哈希表（Hash table，也叫散列表），是根据关键码值(Key value)而直接进行访问的数据结构。

也就是说，它通过把关键码值映射到表中一个位置来访问记录，以加快查找的速度。

这个映射函数叫做散列函数，存放记录的数组叫做散列表。

`记录的存储位置=f(关键字)`

这里的对应关系f称为散列函数，又称为哈希（Hash函数），采用散列技术将记录存储在一块连续的存储空间中，这块连续存储空间称为散列表或哈希表（Hash table）。

哈希表hashtable(key，value) 就是把Key通过一个固定的算法函数既所谓的哈希函数转换成一个整型数字，然后就将该数字对数组长度进行取余，取余结果就当作数组的下标，将value存储在以该数字为下标的数组空间里。

## 实现思路

**数组的特点是：寻址容易，插入和删除困难；**

**而链表的特点是：寻址困难，插入和删除容易。**

那么我们能不能综合两者的特性，做出一种寻址容易，插入删除也容易的数据结构？答案是肯定的，这就是我们要提起的哈希表，哈希表有多种不同的实现方法，我接下来解释的是最常用的一种方法——拉链法，我们可以理解为“链表的数组”，如图：

![img](/Users/zcj/panda/git4me/daily-log-archive/2020-05-22/Go语言的散列表哈希表HashTable实现/6534448-9dc8e65a49c2d619.png)

1. 左边是个数组，数组的每个成员包括一个指针，指向一个链表的头，当然这个链表可能为空，也可能元素很多。
2. 我们根据元素的一些特征把元素分配到不同的链表中去，也是根据这些特征，找到正确的链表，再从链表中找出这个元素

下面是使用go模拟哈希表代码的简单实现：

```go
package HashMp

import (
 "MyHashMap/LinkedNodes"
 "fmt"
)

//定义数组的全局变量
var arr [16] *LinkedNodes.Nodes

//创建16个顶层节点，放到数组中
func CreateArry (){
 var ar = [16]*LinkedNodes.Nodes{}

 for i := 0; i < 16 ; i++ {
 ar[i] = LinkedNodes.CreateArryNode("顶层节点","顶层节点")
 }
 //赋值给全局变量
 arr = ar
 //fmt.Println(ar)
}

//向数组中添加键值对
func AddKVToArr(k,v string){
 //先计算出要添加的数据存储到哪个下角标中，这里调用从网上找的算法
 var corner = HashCode(k)
 var head *LinkedNodes.Nodes = arr[corner]
 //调用添加方法
 LinkedNodes.AddChilddNode(k,v,head)
}


//获取数据
func GetValue(k string) string{
 //先判断是哪个下标存储
 var corner = HashCode(k)
 //获取头节点
 var head *LinkedNodes.Nodes = arr[corner]
 //通过头节点遍历
 for {
 if head.Data.K == k {
 fmt.Println(head.Data.V)
 break
 }else {
 head = head.NextNode
 }
 }
 return ""
}


//将key转换成数组下标的散列算法，范围16之间
func HashCode(key string) int{
 var index int = 0
 index = int(key[0])
 for  k:= 0; k< len(key) ; k++  {
 index *= (1103515245 + int(key[k]))
 }
 index >>= 27
 index &= 16 - 1

 return index
}
```



```go
package LinkedNodes

import "fmt"

//申明全局变量，保存头节点
var heads *Nodes  //头节点 为了遍历使用
var currs *Nodes  //当前节点

//定义结构存储，存储数组每个下标所包含的单独的map数据
type MP struct {
 K string
 V string
}
//创建结构体，用以申明每个节点
type Nodes struct {
 Data MP  // 数据信息
 NextNode *Nodes  //下一个节点的地址
}

//创建头节点
func CreateArryNode(k,v string) *Nodes{
 //创建Node头节点
 var node *Nodes = new(Nodes)
 //封装数据
 node.Data.V = v
 node.Data.K = k
 //指定下一个节点地址，因为还没添加所以是nil
 node.NextNode = nil

 //第一次创建头节点
 heads = node
 currs = node

 return node
}

//向指定的节点中添加数据 第二个参数：指定哪一个节点
func AddChilddNode (k,v string,currs *Nodes) *Nodes{
 var newNode *Nodes = new(Nodes)
 //添加信息
 newNode.Data.K = k
 newNode.Data.V = v

 newNode.NextNode = nil
 //挂接节点
 currs.NextNode = newNode
 currs = newNode
 //fmt.Println(curr)
 return newNode
}

//遍历指定的节点链表
func ShowNode(n *Nodes){
 var node = n
 for  {
 if node.NextNode != nil{
 fmt.Println(node.Data)
 node = node.NextNode
 }else {
 fmt.Println(node.Data)
 break
 }
 }
}

//计算节点个数
func NodesCount ()int{
 var n = heads
 var flag int  //临时存储节点个数变量
 for  {
 if n.NextNode != nil{
 flag+=1
 //fmt.Println(n.data)
 n = n.NextNode
 }else {
 flag+=1
 //fmt.Println(n.data)
 break
 }
 }
 //fmt.Println("节点个数是：",flag)
 return flag
}
```

测试：



```go
package main

import "MyHashMap/HashMp"

//程序入口，主执行
func main () {

 //创建数组，添加顶层节点
 HashMp.CreateArry()
 //随机向数组的每个下标添加子节点
 HashMp.AddKVToArr("abc","世界")
 HashMp.AddKVToArr("def","和平")

 HashMp.GetValue("abc")
}
```



1人点赞



[技术分享]()





作者：李_MAX
链接：https://www.jianshu.com/p/1b62ee8b62ce
来源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。