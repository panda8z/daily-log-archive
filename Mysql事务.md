# MySQL 事务 - 学习笔记

## 前言

### MySQL数据库的架构图

牢记此图，我们接下来的所有内容都将被这张图完整概括。
> 图片来自网络，侵删

![](https://tva1.sinaimg.cn/large/007S8ZIlgy1gjuw8z8z04j31fw0u0ah1.jpg)



### 官方解释

MySQL手册：[MySQL :: MySQL 8.0 Reference Manual](https://dev.mysql.com/doc/refman/8.0/en/)

在MySQL官方的文档里 对 事务的定义是没有解释的。

**数据库事务**是一个数据库技术领域的概念， 维基百科的解释如下：

[Database transaction - Wikipedia](https://en.wikipedia.org/wiki/Database_transaction)

**wiki原文：**

A **database transaction** symbolizes a unit of work performed within a [database management system](https://en.wikipedia.org/wiki/Database_management_system) (or similar system) against a database, and treated in a coherent and reliable way independent of other transactions.

**机翻：**

数据库事务象征着在数据库管理系统（或类似系统）中针对数据库执行的工作单元，并且以与其他事务无关的一致且可靠的方式进行处理。

### 通过转账操作理解事务

关于银行账户转账操作，账户转账是**一个完整的业务**，**最小的单元**，**不可再分**，也就是说银行账户转账是一个**事务**。

**以下是银行账户表t_act(账号、余额)，进行转账操作：**

actno  |     balance
--|--
1        |   500
2        |  100


#### 转账操作

```sql
update t_act set balance=400 where actno=1;
update t_act set balance=200 where actno=2;
```

以上两台DML语句必须同时成功或者同时失败。

每一句都作为事务的一个最小单元不可再分。

当第一条DML语句执行成功后，并不能将底层数据库中的第一个账户的数据修改，只是将操作记录了一下；

这个记录是在内存中完成的；

当第二条DML语句执行成功后，和底层数据库文件中的数据完成同步。

若第二条DML语句执行失败，则清空所有的历史操作记录。

要完成以上的功能必须借助**事务**

## 事务(Transaction)的基本概念

### 事务的定义

- **定义：** **数据库事务**是构成单一逻辑工作单元的操作集合。

- **表现：** 一个事务由批量的 **DML**(`insert`、`update`、`delete`)语句共同联合完成。

只有 **DML** 语句才有事务这个概念。

根据业务逻辑不同，具体某个事务中的 **DML** 语句的个数不同

**注意点：**

1. **数据库事务**是一组（可以包含一个或多个操作）数据库操作， 当这些操作构成一个逻辑上的整体。
2. 构成逻辑整体的这一组操作，要么全部执行成功，要么全部不执行。
3. 不管事务是否成功执行，数据库总能保持一致性状态。

> 思考： 引入事务时如何解决问题呢？ 

#### 关于事务的一些术语

- 开启事务：Start Transaction
- 事务结束：End Transaction
- 提交事务：Commit Transaction
- 回滚事务：Rollback Transaction

###  事务的特点

#### 原子性(Atomicity)
  - 将事务中的所有操作作为一个整体，其像原子一样不可分割。要么全部成功，要么全部失败。
#### 一致性(Consistency)

  - 事务的执行结果必须使数据库从一个一致性状态到另一个一致性状态。
  - 什么是 **一致性状态**？
    1. 系统的状态满足数据的完整性约束（主码，参照完整性，check 约束等）。
    2. 系统的状态反应数据库本应描述的现实世界的真实状态，比如转账前后两个账户的金额总和应该保持不变。
#### 隔离性(Isolation)
  - 并发执行的事务不会互相影响。
  - 他们对数据库的影响和他们串行执行时一样。
  - 比如：多个用户同时往一个账户转账，最后账户的结果应该和他们按先后次序转账的结果一样。

#### 持久性(Durability)

  - 事务一旦提交，其对数据库的更新就是持久的。
  - 任何事务或系统故障都不会导致数据丢失。

## InnerDB 隔离级别

![image-20200922195759124](https://tva1.sinaimg.cn/large/007S8ZIlgy1gjuw926apjj30yo0hk76q.jpg)

![image-20200922195808629](https://tva1.sinaimg.cn/large/007S8ZIlgy1gjuw90rrtgj30yk0f80v0.jpg)

### 隔离级别分类

#### read uncommitted （读未提交）

- 事物A和事物B，事物A未提交的数据，事物B可以读取到
- 这里读取到的数据叫做“脏数据”
- 这种隔离级别最低，这种级别一般是在理论上存在，数据库隔离级别一般都高于该级别

#### read committed （读已提交）

- 事物A和事物B，事物A提交的数据，事物B才能读取到
- 这种隔离级别高于读未提交
- 换句话说，对方事物提交之后的数据，我当前事物才能读取到
- 这种级别可以避免“脏数据”
- 这种隔离级别会导致“不可重复读取”
- Oracle默认隔离级别

#### repeatable read （可重复读）

- 事务A和事务B，事务A提交之后的数据，事务B读取不到
- 事务B是可重复读取数据
- 这种隔离级别高于读已提交
- 换句话说，对方提交之后的数据，我还是读取不到
- 这种隔离级别可以避免“不可重复读取”，达到可重复读取
- 比如1点和2点读到数据是同一个
- MySQL默认级别
- 虽然可以达到可重复读取，但是会导致“幻像读”

#### serializable （串行化）

- 事务A和事务B，事务A在操作数据库时，事务B只能排队等待
- 这种隔离级别很少使用，吞吐量太低，用户体验差
- 这种级别可以避免“幻像读”，每一次读取的都是数据库中真实存在数据，事务A与事务B串行，而不并发

### 隔离级别与一致性关系

## ![隔离级别和一致性的关系](https://img-blog.csdn.net/2018032313015577?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L3dfbGludXg=/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

| 隔离级别 | 脏读 | 不可重复读 | 幻读      |
| -------- | ---- | ---------- | --------- |
| 读未提交 | ✅    | ✅          | ✅         |
| 读已提交 | ❌    | ✅          | ✅         |
| 可重复读 | ❌    | ❌          | 对InnoDB❌ |
| 串行化   | ❌    | ❌          | ❌         |

> 图例：
>
> - ❌：不可能
>
> - ✅：可能



## MySQL的日志种类

### server 层日志

- binlog 二进制日志
- error log 错误日志
- relay log 主从相关日志
- slow log 慢日志

### InnoDB层日志

- undo log 回滚日志
- redo log 重做日志



## 持久性(Durability)的实现原理：redo log

![image-20200922192420580](https://tva1.sinaimg.cn/large/007S8ZIlgy1gjuw904bh7j311m0miaf9.jpg)

![image-20200922195707511](https://tva1.sinaimg.cn/large/007S8ZIlgy1gjuw91iestj30zg0j0441.jpg)

![image-20200922195735861](https://tva1.sinaimg.cn/large/007S8ZIlgy1gjuw8zp1wvj314o0m2tg4.jpg)

### redo log 是什么

redo log (**重做日志**) ，是用来实现事务的持久性.

由两部分组成：

- 重做日志缓冲（redo log buffer）,在内存中.

- 重做日志文件（redo log）,在磁盘中。

当**事务提交之后**会把所有修改信息都会存到该日志中。

和 Undo Log 相反， Redo Log 记录的是新数据的备份。

在事务提交前，只要将 Redo Log 持久化即可， 不需要将数据持久化。 当系统崩溃时，虽然数据没有持久化，但是 Redo Log 已经持久化，系统可以根据 Redo Log 的内容，将所有数据恢复到最新的状态。

### redo log 有什么作用

为了提升性能 MySQL 不会把每次的修改都实时同步到磁盘，而是会先存到**缓冲池（Buffer Pool）**里头，把这个当作缓存来用。然后使用后台线程去做**缓冲池和磁盘之间的同步**。

那么，如果还没来的同步的时候宕机或断电了怎么办？还没来得及执行将缓冲池数据同步到磁盘的操作。这样导致了丢失部分已提交事务的修改信息。

通过 redo log来记录已成功提交事务的修改信息，并且会把redo log持久化到磁盘，系统重启之后在读取redo log恢复最新数据。

总结：
redo log是用来恢复数据的 用于保障，已提交事务的持久化特性

## 原子性(Atomicity)的实现原理：undo log

Undo Log 是为了实现事务的原子性。

在 MySQL 数据库 InnoDB 存储引擎中，Undo Log 还用来实现 **多版本并发控制（MVCC：multi version concurrency control）**。

在操作任何数据之前， 首先将数据备份到一个地方（这个存储数据备份的地方称为 Undo Log），然后再进行数据的修改。

如果出现了错误或者用户执行了 `ROLLBACK` 语句, 系统可以利用 Undo Log 中的数据备份将数据恢复到事务开始之前的状态。

注意： Undo Log 是逻辑日志， 可以理解为：

- 当  `delete` 一条记录时，Undo Log 中会记录一条对应的 `insert` 记录
- 当  `insert` 一条记录时，Undo Log 中会记录一条对应的 `delete` 记录
- 当  `update` 一条记录时，Undo Log 中会记录一条相反的 `update` 记录

#### 和事务相关的两条重要的SQL语句(TCL)

- commit：提交
- rollback：回滚

#### 事务开启的标志

- 任何一条DML语句(insert、update、delete)执行，标志事务的开启

#### 事务结束标志(提交或者回滚)

-  提交：成功的结束，将所有的DML语句操作历史记录和底层硬盘数据来一次同步
-  回滚：失败的结束，将所有的DML语句操作历史记录全部清空

#### 事务与数据库底层数据

在事物进行过程中，未结束之前，DML语句是不会更改底层数据，只是将历史操作记录一下，在内存中完成记录。

只有在事物结束的时候，而且是成功的结束的时候，才会修改底层硬盘文件中的数据

## MySQL中的事务提交与回滚

在MySQL中，默认情况下，事务是自动提交的，也就是说，只要执行一条DML语句就开启了事物，并且提交了事务。但是，这种自动提交机制是可以关闭的

### 对t_user进行提交和回滚操作

#### 提交操作(事务成功)

- start transaction

- DML语句

- commit

  ```
  mysql> start transaction;#手动开启事务
  mysql> insert into t_user(name) values('pp');
  mysql> commit;#commit之后即可改变底层数据库数据
  mysql> select * from t_user;
  +----+------+
  | id | name |
  +----+------+
  |  1 | jay  |
  |  2 | man  |
  |  3 | pp   |
  +----+------+
  3 rows in set (0.00 sec)
  ```

#### 回滚操作(事务失败)

- start transaction

- DML语句

- rollback

  ```
  mysql> start transaction;
  mysql> insert into t_user(name) values('yy');
  mysql> rollback;
  mysql> select * from t_user;
  +----+------+
  | id | name |
  +----+------+
  |  1 | jay  |
  |  2 | man  |
  |  3 | pp   |
  +----+------+
  3 rows in set (0.00 sec)
  ```



### 设置事务隔离级别

#### 方式一: 在my.ini文件中使用transaction-isolation选项来设置服务器的缺省事务隔离级别。

- 该选项值可以是：

  ```
  – READ-UNCOMMITTED
  – READ-COMMITTED
  – REPEATABLE-READ
  – SERIALIZABLE
  
  •   例如：
  [mysqld]
  transaction-isolation = READ-COMMITTED
  ```

#### 方式二： 通过命令动态设置隔离级别

- 隔离级别也可以在运行的服务器中动态设置，应使用 `SET TRANSACTION ISOLATION LEVEL` 语句。
- 其语法模式为：

```sql
  SET [GLOBAL | SESSION] TRANSACTION ISOLATION LEVEL <isolation-level>
```

  其中的`<isolation-level>`可以是：
  –   **READ UNCOMMITTED**
  –   **READ COMMITTED**
  –   **REPEATABLE READ**
  –   **SERIALIZABLE**

> 例如： SET TRANSACTION ISOLATION LEVEL REPEATABLE READ;

### 隔离级别的作用范围

#### 事务隔离级别的作用范围分为两种： 

- 全局级：对所有的会话有效 

- 会话级：只对当前的会话有效 


#### 设置会话级隔离级别为READ COMMITTED

```sql 
SET TRANSACTION ISOLATION LEVEL READ COMMITTED；
```

或：

```sql
SET SESSION TRANSACTION ISOLATION LEVEL READ COMMITTED；
```

#### 设置全局级隔离级别为READ COMMITTED

```sql
  mysql> SET GLOBAL TRANSACTION ISOLATION LEVEL READ COMMITTED；
```

#### 设置全局级隔离级别为READ COMMITTED

```sql
  mysql> SET GLOBAL TRANSACTION ISOLATION LEVEL READ COMMITTED;
```

