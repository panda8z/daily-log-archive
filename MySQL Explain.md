# MySQL 性能分析之 Explain

##  性能瓶颈

  - MySQL Query Optimizer

  - MySQL常见瓶颈
  - CPU:CPU在饱和的时候一般发生在数据装入在内存或从磁盘上读取数据时候
  - IO:磁盘I/O瓶颈发生在装入数据远大于内存容量时
  - 服务器硬件的性能瓶颈：top,free,iostat和vmstat来查看系统的性能状态



### 是什么（查看执行计划）

  - 使用EXPLAIN关键字可以模拟优化器执行SQL语句，从而知道MySQL是
  - 如何处理你的SQL语句的。分析你的查询语句或是结构的性能瓶颈
  - 官网介绍

### 能干嘛

  - 表的读取顺序
  - 数据读取操作的操作类型
  - 哪些索引可以使用
  - 哪些索引被实际使用
  - 表之间的引用
  - 每张表有多少行被优化器查询

### 怎么玩

  - - Explain+SQL语句

  - 执行计划包含的信息

    ![image-20200923095646108](https://tva1.sinaimg.cn/large/007S8ZIlgy1gj0cnomakpj31do030whi.jpg)

  - ### 各个字段解释

#### id

  - - select查询的序列号，包含一组数字，表示查询中执行select子句或操作表的顺序
    
    - 三种情况
    
    - - id相同，执行顺序由上至下
    - id不同，如果是子查询，id的序号会递增，id值越大优先级越高，越先被执行
    - id相同不同，同时存在

#### select_type

    - 有哪些
      
    - 查询的类型，主要用于区别
      
    - 普通查询、联合查询、子查询等的复杂查询
      
    - - 1.SIMPLE
      
    - - 简单的select查询，查询中不包含子查询或者UNION
      
    - 2.PRIMARY
      
    - - 查询中若包含任何复杂的子部分，最外层查询则被标记为
      
    - 3.SUBQUERY
      
    - - 在SELECT或者WHERE列表中包含了子查询
      
    - 4.DERIVED
      
    - - 在FROM列表中包含的子查询被标记为DERIVED（衍生）
      - MySQL会递归执行这些子查询，把结果放在临时表里。
      
    - 5.UNION
      
    - - 若第二个SELECT出现在UNION之后，则被标记为UNION;
      - 若UNION包含在FROM子句的子查询中，外层SELECT将被标记为：DERIVED
      
    - 6.UNION RESULT
      
    - - 从UNION表获取结果的SELECT

#### table

  - 显示这一行的数据是关于哪张表的

#### type

- Subtopic
- 访问类型排列
- 显示查询使用了何种类型
- 从最好到最差依次是：
- system>const>eq_ref>ref>range>index>ALL
  - system
  - 表只有一行记录（等于系统表），这是const类型的特例，平时不会出现，这个也可以忽略不计
- const
  - 表示通过索引一次就找到了，const用于比较primary key或者unique索引。因为只匹配一行数据，所以很快。如将主键至于where列表中，MySQL就能将该查询转换为一个常量
- eq_ref
- - 唯一性索引，对于每个索引键，表中只有一条记录与之匹配，常见于主键或唯一索引扫描
- ref
- - 非唯一索引扫描，返回匹配某个单独值的所有行。
  - 本质上也是一种索引访问，它返回所有匹配某个单独值的行，然而，
  - 它可能会找到多个符合条件的行，所以他应该属于查找和扫描的混合体
- range
- - 只检索给定范围的行，使用一个索引来选择行。key列显示使用了哪个索引
  - 一般就是在你的where语句中出现了between、<、>、in等的查询
  - 这种范围扫描索引扫描比全表扫描要好，因为他只需要开始索引的某一点，而结束语另一点，不用扫描全部索引
- index
- - Full Index Scan,index与ALL区别为index类型只遍历索引树。这通常比ALL快，因为索引文件通常比数据文件小。
  - （也就是说虽然all和index都是读全表，但index是从索引中读取的，而all是从硬盘中读的）
- all
  - FullTable Scan,将遍历全表以找到匹配的行
- 备注：
- - 一般来说，得保证查询只是达到range级别，最好达到ref

#### possible_keys

  - - 显示可能应用在这张表中的索引,一个或多个。
    - 查询涉及的字段上若存在索引，则该索引将被列出，但不一定被查询实际使用

#### key

  - - 实际使用的索引。如果为null则没有使用索引
    - 查询中若使用了覆盖索引，则索引和查询的select字段重叠

#### key_len

  - - 表示索引中使用的字节数，可通过该列计算查询中使用的索引的长度。在不损失精确性的情况下，长度越短越好
    - key_len显示的值为索引最大可能长度，并非实际使用长度，即key_len是根据表定义计算而得，不是通过表内检索出的

#### ref

  - - 显示索引那一列被使用了，如果可能的话，是一个常数。那些列或常量被用于查找索引列上的值

#### rows

  - - 根据表统计信息及索引选用情况，大致估算出找到所需的记录所需要读取的行数

#### Extra

包含不适合在其他列中显示但十分重要的额外信息

1. Using filesort
   - 说明mysql会对数据使用一个外部的索引排序，而不是按照表内的索引顺序进行读取。
   - MySQL中无法利用索引完成排序操作成为“文件排序”
2. Using temporary
  - 使用了临时表保存中间结果，MySQL在对查询结果排序时使用临时表。常见于排序order by 和分组查询 group by
3. USING index
  - 表示相应的select操作中使用了覆盖索引（Coveing Index）,避免访问了表的数据行，效率不错！
   - 如果同时出现using where，表明索引被用来执行索引键值的查找；
   - 如果没有同时出现using where，表面索引用来读取数据而非执行查找动作。
   - 覆盖索引（Covering Index）
4. Using where 表面使用了where过滤
5. using join buffer 使用了连接缓存
6. impossible where where子句的值总是false，不能用来获取任何元组
7. select tables optimized away
  - 在没有GROUPBY子句的情况下，基于索引优化MIN/MAX操作或者
  - 对于MyISAM存储引擎优化COUNT(*)操作，不必等到执行阶段再进行计算，
  - 查询执行计划生成的阶段即完成优化。

8. distinct
  - 优化distinct，在找到第一匹配的元组后即停止找同样值的工作

### 热身Case