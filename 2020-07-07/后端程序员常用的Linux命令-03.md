作为一个java后端开发者，在日常工作中肯定会有和linux打交道的时候。下面我整理了一些开发中可能需要用到的一些命令（写这篇随笔前，我也没有什么操作经验，主要防止自己忘记这些操作命令，以后有用到新的命令再修改）。

以下的操作是通过SecureCRT对服务器进行远程连接，使用WinSCP来进行文件上传。

一、用户相关的一些命令

1.新建用户 adduser username （只有root有权限新增用户）

2.修改用户密码 passwd username 

3.切换用户 su username

4.删除用户 userdel username

5.以管理员的身份进行操作 sudo

 

二、文件的操作命令

1.列出文件的基本信息 ls -l 

　　-rw-rw-r-- 1 ubuntu ubuntu  162 Mar 26 14:49 t.tar.gz

　　第一位字符d表示目录，- 表示文件

　　后面9位分别是当前用户、当前用户所在组和其他用户对文件的权限每个占三位。文件的权限中 - 表示没有权限； r表示读，数字为4；w表示写，数字为2；x表示执行，数字为1。

　　后面表示文件所属的用户和组，以及文件的大小，创建日期和文件名

2.修改文件操作权限  

　　chmod u+x filename 表示给所属用户增加执行的权限

　　chmod g-w filename 表示给组减少写的权限

　　chmod o+w filename 表示给其他用户增加写的权限

　　chmod 777 filename 表示给所属用户，所属组，其他用户开放所有权限；7是通过4+2+1得到

3.文件的操作

　　pwd 显示当前所在的路径

　　mkdir filename 创建文件夹

　　mkdir -p /a/b 创建嵌套的目录

　　touch filename 创建文件

　　cp filename new_filename 复制文件

　　mv filename path 移动文件（如果路径中带文件名，则会给文件重命名）

　　rm -r filename 删除文件或者文件夹，（-r是递归的意思）

　　rm -f 强制删除目录或文件

　　rmdir filename 删除文件夹（只能是空文件夹）

　　find / -name "install.log“ 查找 可以用*作为通配符

4.文件的内容操作

　　find /path -name "*findname*" -type d 查询某个文件夹下面名字包含filename的文件夹

　　find /path -type d |wc -l 统计某路径下文件夹的个数

　　find /path -type f | wc -l 统计某路径下文件的个数

　　cat 查看文件内容

　　wc -lwc filename 统计文件的行数，字符数，字节数 l表示行数，w表示字符数，c表示字节数

　　> 覆盖式输出重定向符 >>追加式输出重定向符

　　vi filename 对文件进行编辑。

　　vi有三种模式，分别是命令模式、插入模式和末行模式。

　　默认进入命令模式，其他模式按esc也可进入命令模式

　　　　　　*光标← ↑ → ↓　　上下左右*

　　　　　　Page Down 或Ctrl+F　　下翻一页

　　　　　　Page Up 或Ctrl+B　　上翻一页

　　　　　　Home或^ 或0　　跳至行首

　　　　　　End 或$ 　　跳至行尾

　　　　　　#→　　右跳#个字符

　　　　　　#← 　　左跳#个字符

　　　　　　1G 或 gg　　跳至首行

　　　　　　G　　跳至尾行

　　　　　　#G　　跳至第#行

　　　　　　:set nu 　　示行数

　　　　　　:set nonu 　显示行数

　　　　　　x或del　　删除光标处的单个字符

　　　　　　dd　　删除当前行

　　　　　　#dd　　删除当前行开始的#行

　　　　　　d^　　删除当前行首至光标的字符

　　　　　　d$　　删除当前行尾至光标的字符

　　　　　　yy　　复制当前行

　　　　　　#yy　　复制当前行开始的#行

　　　　　　p　　粘贴至光标的右侧

　　　　　　P　　粘贴至光标的左侧

　　　　　　/word　　向下查找字符串word

　　　　　　?word　　向上查找字符串word

　　　　　　n　　光标定位至下一个匹配字符

　　　　　　N　　光标定位至上一个匹配字符

　　　　　　u　　单次撤销 多次恢复

　　　　　　U　　撤销所有编辑

 

　　插入模式：命令模式下，i 、a、o、 Insert即可进入。该模式可用于编辑文本

　　末行模式：shit+: 即可进入

　　　　　　　　 :w  保存文件 

　　　　　　　　 : /filename  另存文件为fielname

　　　　　　　　:q 未修改时退出

　　　　　　　　:q! 不保存且强制退出

　　　　　　　　:wq 保存修改并退出

 　　　　　　　:s /old/new 前行的第一个old替换为new

　　　　　　　　:s /old/new/g 前行的所有old替换为new

　　　　　　　　:#,# s /old/new/g   #行间的所有old替换为new

　　　　　　　　:% s /old/new/g 当前文件所有old替换为new

　　　　　　　　:s /old/new/c  old替换为new时提示确认替换

　　　　　　　　:g/str1/s//str2/g str2 替换所有的str1

 

二、程序员日常维护项目所用的命令

　　1.找到服务器对应的进程 ps -ef |grep "tomcat" 这样就能知道服务器所对应的进程，以及服务器安装的目录。

　　2.关闭服务器可以用kill -9 pid 杀死进程。也可以到tomcat的bin目录下./shutdown.sh。开启服务器./startup.sh

　　3.用winscp，上传修改后的文件。

　　4.日志查看，tail -200f catalina.out 显示日志的最后200行。ctrl+c 退出查看。

 

三、jdk和tomcat的安装和卸载（我个人习惯安装在/opt目录下）

　　打包 tar -cvf 1.tar 1.txt

　　打包并压缩 tar -cvzf 1.tar.gz 1.txt

　　压缩 gzip 1.txt

　　解压 gzip -d 1.txt.zip

　　解压缩并拆包 tar -zxvf 1.tar.gz

　　拆包 tar -xvf XXXXX.tar

　　解压.tar.xz文件，要先执行 xz -d XXXXXX.tar.xz 将该文件变为.tar文件然后再解包

　　有两种方式安装，自动和手动。

 

①从官网下载所需版本jdk安装包，上传到服务器，并解压。

设置环境变量：vim /etc/profile  

在文件的末尾加上

　　export JAVA_HOME="/opt/jdk1.8.0_131"  
　　export PATH="${JAVA_HOME}/bin:$PATH"

修改完成后用source /etc/profile执行profile文件，然后用java -version 进行测试。

 

②从官网下载对应的版本tomcat安装包，上传到服务器，并解压。

1、查看/etc/profile文件中有没有配置环境变量。

2、修改tomcat的conf目录下server.xml文件，修改http1.1监听的端口改为80。

3、启动服务器tomcat的bin目录下./startup.sh

4、关闭服务器tomcat的bin目录下./shutdown.sh

5、查看端口号被哪个线程占用 lsof -i:80

6、根据pid查询进程相关信息；cd /proc/5941;然后ls -ail 即可。

 

 

 

四、其他命令记录

　　1.卸载软件 apt-get purge XXXX 