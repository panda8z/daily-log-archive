# 网络

找出某程序(tomcat)的进程 ps -ef|grep tomcat   找出后如果要关闭 kill -9 pid
统计某程序(tomcat)连接数 ps -ef|grep tomcat|wc -l
查看tomcat日志文件包含某关键字的记录  grep '关键字' -C5 xxxxx.log   -C5表示包括上面5行和下面5行一起打印出来
查看文件 vi或vim /关键字 查找

 

列出所有 tcp 端口 netstat -at
只列出所有监听 tcp 端口 netstat -lt
找出特定端口运行的程序 netstat -p (p就是pid)
找出某程序运行的端口 netstat -ap | grep ssh
每隔一秒持续输出netstat信息 netstat -c
列出连接某服务端口最多的前20个ip（由多到少排序） netstat -nat | grep "192.168.1.15:22" |awk '{print $5}'|awk -F: '{print $1}'|sort|uniq -c|sort -nr|head -20
用tcpdump嗅探80端口的访问看看谁最高 tcpdump -i eth0 -tnn dst port 80 -c 1000 | awk -F"." '{print $1"."$2"."$3"."$4}' | sort | uniq -c | sort -nr |head -20
查看http的并发请求数与其TCP连接状态 netstat -na | awk '/^tcp/ {++S[$NF]} END {for(i in S) print i, S[i]}'

# 文件

按名称查找某个目录（可以是绝对路径也可以是相对路径）下的文件 find /usr/local/ -name *.conf
按文件权限查找文件 find /usr/local/nginx -perm 755
查找某目录下10天内修改过的文件 find /usr/local/nginx -mtime -10
查找某目录下10天前修改过的文件 find /usr/local/nginx -mtime +10