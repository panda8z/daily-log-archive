(window.webpackJsonp=window.webpackJsonp||[]).push([[14],{364:function(v,_,i){"use strict";i.r(_);var l=i(42),t=Object(l.a)({},(function(){var v=this,_=v.$createElement,i=v._self._c||_;return i("ContentSlotsDistributor",{attrs:{"slot-key":v.$parent.slotKey}},[i("h1",{attrs:{id:"阅读《rust权威指南》"}},[i("a",{staticClass:"header-anchor",attrs:{href:"#阅读《rust权威指南》"}},[v._v("#")]),v._v(" 阅读《Rust权威指南》")]),v._v(" "),i("ul",[i("li",[v._v("一、入门指南\n"),i("ul",[i("li",[v._v("安装 rustup 工具.")]),v._v(" "),i("li",[v._v("rustup 命令行工具.")]),v._v(" "),i("li",[v._v("rustc 命令行工具.")]),v._v(" "),i("li",[v._v("helloworld程序.")]),v._v(" "),i("li",[v._v("cargo 命令行工具.")]),v._v(" "),i("li",[v._v("rustup doc 本地文档入口.")]),v._v(" "),i("li",[v._v("rustup docs --book rust英文书.")])])]),v._v(" "),i("li",[v._v("二、编写一个猜数游戏\n"),i("ul",[i("li",[v._v("fn关键字.")]),v._v(" "),i("li",[v._v("main函数定义.")]),v._v(" "),i("li",[v._v("println!()方法使用.")]),v._v(" "),i("li",[v._v("语句后要有分号.")]),v._v(" "),i("li",[v._v("{} 占位符.")]),v._v(" "),i("li",[v._v("let 关键字定义变量.")]),v._v(" "),i("li",[v._v("mut关键字定义可变变量.")]),v._v(" "),i("li",[v._v("use std::io 导入包.")]),v._v(" "),i("li",[v._v(":: 运算符 调用包方法 io::stdin().")]),v._v(" "),i("li",[v._v("String 类型.")]),v._v(" "),i("li",[v._v("Result类型和execpt()方法处理错误.")]),v._v(" "),i("li",[v._v("使用 第三方库 rand.")]),v._v(" "),i("li",[v._v("cargo build 下载依赖包并编译.")]),v._v(" "),i("li",[v._v("cargo update 更新新版本依赖包.")]),v._v(" "),i("li",[v._v("cargo.lock文件锁定依赖包版本.")]),v._v(" "),i("li",[v._v("cargo run 编译并运行程序.")]),v._v(" "),i("li",[v._v("match表达式的基本使用.")]),v._v(" "),i("li",[v._v("接触了 => 运算符,用于模式匹配.")]),v._v(" "),i("li",[v._v("loop {} 表达式做循环.")]),v._v(" "),i("li",[v._v("break语句结束循环.")])])]),v._v(" "),i("li",[v._v("三、通用编程概念\n"),i("ul",[i("li",[v._v("变量, 用let关键字声明.")]),v._v(" "),i("li",[v._v("可变性, 用 mut 关键字控制.")]),v._v(" "),i("li",[v._v("标量类型 scalar.")]),v._v(" "),i("li",[v._v("复合类型 compound.")]),v._v(" "),i("li",[v._v("有符号整型 i8,i16,i32,i64,isize.")]),v._v(" "),i("li",[v._v("无符号整型 u8,u16,u32,u64,usize.")]),v._v(" "),i("li",[v._v("浮点型 f32, f64,遵循IEEE-754s")]),v._v(" "),i("li",[v._v("整数字面量 76_000, 0xFDA9, 0o77, 0b1101_0011, b'A'.")]),v._v(" "),i("li",[v._v("数值运算符 +, -, *, /, %.")]),v._v(" "),i("li",[v._v("布尔类型 bool true false.")]),v._v(" "),i("li",[v._v("字符类型 描述单个字符:‘a’ ‘5’ ‘*’ ‘¥’.")]),v._v(" "),i("li",[v._v("元组 tuple 可放不同类型元素 定长 索引取值 (u8,i8,f32) (1,-2,3.0).")]),v._v(" "),i("li",[v._v("数组 array 只能装同类型元素 定长 索引和迭代器取值 [1,2,3,4,5].")]),v._v(" "),i("li",[v._v("Rust使用蛇形命名法(snake case)规范函数和变量标识符.")]),v._v(" "),i("li",[v._v("函数 用 fn 关键字定义 ,fn+函数名+参数列表+返回值列表+函数体.")]),v._v(" "),i("li",[v._v("函数参数和返回值必须明确指定类型.")]),v._v(" "),i("li",[v._v("函数参数是 parameters 中文叫形式参数.")]),v._v(" "),i("li",[v._v("给函数传递具体值的时候的参数叫 arguments 中文叫实际参数.")]),v._v(" "),i("li",[v._v("语句 statement 仅有操作没有返回值.")]),v._v(" "),i("li",[v._v("表达式 expression 有具体计算后的值返回.")]),v._v(" "),i("li",[v._v("{} 可以定义一个代码块,也是个表达式,有返回值.")]),v._v(" "),i("li",[v._v("Rust默认最后一个表达式的值是函数返回值.")]),v._v(" "),i("li",[v._v("-> 运算符定义了 函数的返回值列表.")]),v._v(" "),i("li",[v._v("注释,单行注释就是 // 文档注释在第十章才讲到.")]),v._v(" "),i("li",[v._v("if else表达式控制流: if+判断条件表达式+{} + else if + 判断条件 + {} + else + {}.")]),v._v(" "),i("li",[v._v("Rust提供三种循环: while loop 和 for")]),v._v(" "),i("li",[v._v("loop {} 无脑循环.")]),v._v(" "),i("li",[v._v("while + 判断条件 + {} 正常条件循环.")]),v._v(" "),i("li",[v._v("for item in list.iter() 集合遍历循环.")]),v._v(" "),i("li",[v._v("for - in 表达式集合遍历循环.")]),v._v(" "),i("li",[v._v(".. 表达式生成Range序列,(1..4)会生成一个序列目前不清楚是元组还是数组.")])])]),v._v(" "),i("li",[v._v("四、认识所有权\n"),i("ul",[i("li",[v._v("所有权:rust核心概念之一, 影响深远,认真学习!")]),v._v(" "),i("li",[v._v("Rust 使用包含特定规则的所有权系统来管理内存分配和使用.")]),v._v(" "),i("li",[v._v("编译器会在编译期发现并提出规避风险的措施.")]),v._v(" "),i("li",[v._v("栈,先进后出的数据结构,摞盘子模型.")]),v._v(" "),i("li",[v._v("堆,大块连续内存区域,管理比较松散,使用此区域前要分配不同大小的区域.")]),v._v(" "),i("li",[v._v("计算机对堆的操作比栈的操作跟费时费力.")]),v._v(" "),i("li",[v._v("所有权规则1:Rust中的每一个值都有一个对应的变量作为它的所有者.")]),v._v(" "),i("li",[v._v("所有权规则2:在同一时间内,值有且仅有一个所有者.")]),v._v(" "),i("li",[v._v("所有权规则3:当所有者离开自己的作用域时,它持有的值就会被释放掉.")]),v._v(" "),i("li",[v._v("作用域是一个对象在程序中有效的范围.")]),v._v(" "),i("li",[v._v("就目前知识而言: 代码块,函数结束就失去了作用域.")]),v._v(" "),i("li",[v._v("String类型是标准库提供的字符串数据类型.")]),v._v(" "),i("li",[v._v("let s = String::from(“Hello”) 创建新的不可变字符串变量.")]),v._v(" "),i("li",[v._v("let mut s = String::from(“Hello”) 创建新的可变字符串变量.")]),v._v(" "),i("li",[v._v("就字符串字面量而言,编译期知道其具体大小所以在栈上分配.")]),v._v(" "),i("li",[v._v("可变字符串变量,编译时未知其具体大小,所以在堆上分配.")]),v._v(" "),i("li",[v._v("与有GC语言不同,Rust内存会自动地在拥有它的变量离开作用域后进行释放.")]),v._v(" "),i("li",[v._v("Rust会在作用域结束的地方自动调用 drop()函数.")]),v._v(" "),i("li",[v._v("资源获取即初始化 Resource Acquisition Is Initialization RAII.")]),v._v(" "),i("li",[v._v("变量和数据的交互方式: 移动 move.")]),v._v(" "),i("li",[v._v("Rust永远不会自动地创建数据地深度拷贝.")]),v._v(" "),i("li",[v._v("变量和数据地交互方式: 克隆 clone.")]),v._v(" "),i("li",[v._v("堆上地数据类型变量可以调用 clone() trait 产生新的堆分配变量.")]),v._v(" "),i("li",[v._v("栈上变量和数据通过复制来交互 copy trait.")]),v._v(" "),i("li",[v._v("将变量传递给函数将会触发移动或复制,就像是赋值语句一样.")]),v._v(" "),i("li",[v._v("函数在返回值过程中也会发生所有权地转移.")]),v._v(" "),i("li",[v._v("变量所有权转移的模式1: 将一个值赋值给另一个变量时就会转移所有权.")]),v._v(" "),i("li",[v._v("变量所有权转移的模式2: 当一个持有堆数据的变量离开作用域时它的数据将会被 drop() 清理,除非这些数据的所有权发生了转移.")]),v._v(" "),i("li",[v._v("Rust提供了引用功能来简化所有权频繁更换的场景.")]),v._v(" "),i("li",[v._v("&s1 是String类型数据变量s1的引用")]),v._v(" "),i("li",[v._v("& 代表了引用予以,它允许在不获取所有权的前提下使用值.")]),v._v(" "),i("li",[v._v("解引用 dereferencing,使用 * 运算符,15章详细介绍.")]),v._v(" "),i("li",[v._v("对于特定作用域中的特定数据来说,一次只能使用一个可变引用.")]),v._v(" "),i("li",[v._v("数据竞争 data race.")]),v._v(" "),i("li",[v._v("data race1: 两个或两个以上的指针同时访问同一空间.")]),v._v(" "),i("li",[v._v("data race2: 其中至少有一个指针会向空间中写入数据.")]),v._v(" "),i("li",[v._v("data race3: 没有同步数据访问的机制.")]),v._v(" "),i("li",[v._v("任何数据竞争的情况出现在Rust里都编译不过.")]),v._v(" "),i("li",[v._v("不能在拥有不可变引用的同时创建可变应用.")]),v._v(" "),i("li",[v._v("Rust的编译器会确保引用永远不会进入悬垂状态.")]),v._v(" "),i("li",[v._v("新的概念 生命周期 会在 第10章详解.")]),v._v(" "),i("li",[v._v("引用规则1: 在任何一段给定的时间里,要么只能拥有一个可变应用,要么只能拥有任意数量的不可变引用.")]),v._v(" "),i("li",[v._v("引用规则2: 引用总是有效的.")]),v._v(" "),i("li",[v._v("切片 "),i("code",[v._v("slice")]),v._v(" Rust里面一个不持有所有权的数据类型.")]),v._v(" "),i("li",[v._v("字符串切片的类型写作:&str")]),v._v(" "),i("li",[v._v("还有很多其他类型的切片.&[i32] 数组切片.")]),v._v(" "),i("li",[v._v("字符串字面量就是切片.")])])]),v._v(" "),i("li",[v._v("五、使用结构体来组织相关联的数据\n"),i("ul",[i("li",[v._v("结构体,是一种自定义数据类型,它允许我们命名多个相关的值并将它们组成一个有机结合体.")]),v._v(" "),i("li",[v._v("和元组一样,结构体中的的数据可以拥有不同类型.")]),v._v(" "),i("li",[v._v("和元组不一样, 结构体需要给每个数据赋予名字以便清楚的表明它们的意义.")]),v._v(" "),i("li",[v._v("关键字 "),i("code",[v._v("struct")]),v._v("定义结构体.")]),v._v(" "),i("li",[v._v("通过点号来访问结构体实例中的特定字段.")]),v._v(" "),i("li",[v._v("一旦实例可变,实例中的所有字段都将是可变的.")]),v._v(" "),i("li",[v._v("结构体初始化支持 字段初始化简写 语法 field init shorthand.")]),v._v(" "),i("li",[v._v("用类似元组的方式来定义元组结构体,不需要对字段命名.")]),v._v(" "),i("li",[v._v("Rust允许空结构体.后续章节会解释用法.")]),v._v(" "),i("li",[v._v("Rust独有的生命周期概念保证了结构体实例中引用数据的有效期不短于实例本身.")]),v._v(" "),i("li",[v._v("分别用硬编码,元组,结构体实现长方形面积的求解程序.")]),v._v(" "),i("li",[v._v("结构体可以定义方法.")]),v._v(" "),i("li",[v._v("方法总被定义在某个结构体的上下文上.")]),v._v(" "),i("li",[i("code",[v._v("#[derive(Debug)]")]),v._v(" 注解 让打印更人性化.")]),v._v(" "),i("li",[v._v("Rust有自动引用和解引用功能,直接用引用点元素就能访问数据.")]),v._v(" "),i("li",[v._v("方法中的第一个参数 可能是 "),i("code",[v._v("self")]),v._v(" 或 "),i("code",[v._v("&self")]),v._v(" 表示结构体本身或起引用.")]),v._v(" "),i("li",[v._v("不接受self参数的函数成文关联函数,association function,一般用于初始化结构体实例.")]),v._v(" "),i("li",[v._v("所有方法被放在 关键字 "),i("code",[v._v("impl + 结构体名 + {}")]),v._v(" 定义的代码块中.")]),v._v(" "),i("li",[v._v("多个 "),i("code",[v._v("impl")]),v._v(" 代码块可以同时存在.")])])]),v._v(" "),i("li",[v._v("六、枚举与模式匹配\n"),i("ul",[i("li",[v._v("枚举类型,简称为枚举,它允许我们列举所有可能的值来定义一个类型.")]),v._v(" "),i("li",[v._v("枚举中的元素被称为 枚举变体 variant.")]),v._v(" "),i("li",[v._v("用 :`: 运算符访问枚举变体.")]),v._v(" "),i("li",[v._v("枚举允许我们直接将其关联的数据嵌入枚举变体内.")]),v._v(" "),i("li",[v._v("IP相关的编码和操作标准库内置了一套开箱即用的定义.")]),v._v(" "),i("li",[v._v("枚举同样可以使用 "),i("code",[v._v("impl")]),v._v(" 关键字定义的代码块来添加方法.")]),v._v(" "),i("li",[v._v("标准库中定义了一个枚举 Option 非常重要的核心概念.")]),v._v(" "),i("li",[v._v("Option类型描述了一种值可能不存在的情形,所以它广泛引用在各个地方.")]),v._v(" "),i("li",[v._v("使用"),i("code",[v._v("option")]),v._v("概念的类型系统,可以让编译器自动检查我们是否妥善处理了所有应该被处理的情况.")]),v._v(" "),i("li",[v._v("引发空值缺陷问的的关键不是空值概念本身,而是那些具体的实现措施.")]),v._v(" "),i("li",[v._v("<T> 是泛型参数, 第10章讨论其细节.")]),v._v(" "),i("li",[v._v("Rust中,无论在什么地方,只要一个值的类型不是 Option<T> 的,就可以安全假设其为非空.")]),v._v(" "),i("li",[v._v("match 表达式 是可以用来处理枚举的控制流结构.")]),v._v(" "),i("li",[v._v("match 表达式, 允许我们基于枚举拥有的变体来决定运行分支,并可通过模式匹配值来获取变体内数据.")])])])])])}),[],!1,null,null,null);_.default=t.exports}}]);