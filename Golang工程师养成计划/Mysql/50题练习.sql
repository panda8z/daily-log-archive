-- 面试遇到的题：
-- 101: 查询课程1的成绩比课程2的成绩高的所有学生的学号
-- 102：查询平均成绩大于60分的同学的学号和平均成绩
-- 103: 查询所有同学的学号、姓名、选课数、总成绩
-- 104: 查询姓“张”的老师的个数


-- 1、查询"01"课程比"02"课程成绩高的学生的信息及课程分数	
-- solution-1-1
SELECT a.*,
    b.sscore AS 01_Score,
    c.sscore AS 02_score
FROM Student a
    JOIN Score b ON a.sno = b.sno
    AND b.cno = '01'
    LEFT JOIN Score c ON a.sno = c.sno
    AND c.cno = '02'
WHERE b.sscore > c.sscore;
-- :output
-- 02	钱电	1990-12-21	男	70	60
-- 04	李云	1990-08-06	男	50	30
-- solution-1-2
SELECT a.*,
    b.sscore AS 01_score,
    c.sscore AS 02_score
FROM Student a,
    Score b,
    Score c
WHERE a.sno = b.sno
    AND a.sno = c.sno
    AND b.cno = '01'
    AND c.cno = '02'
    AND b.sscore > c.sscore;
-- 2、查询"01"课程比"02"课程成绩低的学生的信息及课程分数
-- solution-2-1
SELECT a.*,
    b.sscore AS 01_score,
    c.sscore AS 02_score
FROM Student a
    LEFT JOIN Score b ON a.sno = b.sno
    AND b.cno = '01'
    JOIN Score c ON a.sno = c.sno
    AND c.cno = '02'
WHERE b.sscore < c.sscore;
-- 	solution-2-2
SELECT a.*,
    b.sscore AS 01_score,
    c.sscore AS 02_score
FROM Student a,
    Score b,
    Score c
WHERE a.sno = b.sno
    AND a.sno = c.sno
    AND b.cno = '01'
    AND c.cno = '02'
    AND b.sscore < c.sscore;
-- 3、查询平均成绩大于等于60分的同学的学生编号和学生姓名和平均成绩
-- solution-3-1
SELECT a.sno,
    a.sname,
    ROUND(AVG(b.sscore), 2) AS avg_score
FROM Student a
    JOIN Score b ON b.sno = a.sno
GROUP BY a.sno,
    a.sname
HAVING avg_score >= 60;
-- 测试 仅仅分数查询
SELECT a.sno,
    a.sname,
    b.sscore
FROM Student a
    JOIN Score b ON b.sno = a.sno -- 测试 仅仅平均分查询
SELECT a.sno,
    a.sname,
    ROUND(AVG(b.sscore), 2) AS avg_score
FROM Student a
    JOIN Score b ON b.sno = a.sno
GROUP BY a.sno,
    a.sname;
-- 4、查询平均成绩小于60分的同学的学生编号和学生姓名和平均成绩
-- (包括有成绩的和无成绩的)
-- 测试4-1 仅查出小于60分的，不包含没有成绩的人	
SELECT a.sno,
    a.sname,
    ROUND(AVG(b.sscore), 2) AS avg_score
FROM Student a
    LEFT JOIN Score b ON a.sno = b.sno
GROUP BY a.sname,
    a.sno
HAVING avg_score < 60;
-- solution-4-1
SELECT a.sno,
    a.sname,
    ROUND(AVG(b.sscore), 2) AS avg_score
FROM Student a
    LEFT JOIN Score b ON a.sno = b.sno
GROUP BY a.sname,
    a.sno
HAVING avg_score < 60
UNION
SELECT a.sno,
    a.sname,
    0 AS avg_score
FROM Student a
WHERE a.sno NOT IN (
        SELECT DISTINCT sno
        FROM Score
    );
-- 测试4-2 从Score表里查出不同的sno
SELECT DISTINCT sno
from Score;
-- 5、查询所有同学的学生编号、学生姓名、选课总数、所有课程的总成绩
SELECT a.sno,
    a.sname,
    COUNT(b.cno) AS sum_course,
    SUM(b.sscore) AS sum_score
FROM Student a
    LEFT JOIN Score b ON a.sno = b.sno
GROUP BY a.sno,
    a.sname;
-- 6、查询"李"姓老师的数量 
SELECT COUNT(tno)
FROM Teacher
WHERE tname LIKE '李%';
-- 7、查询学过"张三"老师授课的同学的信息 
-- 查出老师的编号
SELECT tno
FROM Teacher
WHERE tname = '张三'
);
-- 查出老师讲授的课程信息
SELECT c.cno
FROM Course c
WHERE tno = (
        SELECT tno
        FROM Teacher
        WHERE tname = '张三'
    );
-- 最终整合在一起：查出成绩里有这个课程的学生的信息
SELECT a.*
FROM Student a
    LEFT JOIN Score b ON b.sno = a.sno
WHERE b.cno IN (
        SELECT c.cno
        FROM Course c
        WHERE c.tno = (
                SELECT tno
                FROM Teacher
                WHERE tname = '张三'
            )
    );
-- 8、查询没学过"张三"老师授课的同学的信息 
-- 在上一题的基础上，在所有学生里编号不在上一题的结果里的学生就是答案
SELECT *
FROM Student
WHERE sno NOT IN (
        SELECT a.sno
        FROM Student a
            LEFT JOIN Score b ON b.sno = a.sno
        WHERE b.cno = (
                SELECT c.cno
                FROM Course c
                WHERE c.tno = (
                        SELECT tno
                        FROM Teacher
                        WHERE tname = '张三'
                    )
            )
    )
GROUP BY sno,
    sname;
-- 9、查询学过编号为"01"并且也学过编号为"02"的课程的同学的信息
-- solution-9-1
SELECT a.*
FROM Student a,
    Score b,
    Score c
WHERE b.sno = a.sno
    AND b.sno = c.sno
    AND b.cno = '01'
    AND c.cno = '02';
-- 10、查询学过编号为"01"但是没有学过编号为"02"的课程的同学的信息
-- 学过01的学生
SELECT sno
FROM Score
WHERE cno = '01';
-- 学过02的学生
SELECT sno
FROM Score
WHERE cno = '02';
-- solution-10-1
SELECT *
FROM Student
WHERE sno IN (
        SELECT sno
        FROM Score
        WHERE cno = '01'
    )
    AND sno NOT IN (
        SELECT sno
        FROM Score
        WHERE cno = '02'
    );
-- 11、查询没有学全所有课程的同学的信息 
-- 查出所有有成绩的同学
SELECT a.*
FROM Student a,
    Score b
WHERE a.sno = b.sno
GROUP BY a.sno,
    a.sname;
-- 没有学习任何课程的人 这个和题无关
SELECT *
FROM Student
WHERE sno NOT IN (
        SELECT a.sno
        FROM Student a,
            Score b
        WHERE a.sno = b.sno
        GROUP BY a.sno
    );
-- solution-11-1
-- 课程数 大于 学过的数量的人
-- 课程数量
SELECT COUNT(*)
FROM Course;
SELECT a.*
FROM Student a
    LEFT JOIN Score b ON a.sno = b.sno
GROUP BY a.sno
HAVING COUNT(b.cno) < (
        SELECT COUNT(*)
        FROM Course
    );
-- solution-11-2
-- 先查出学过全部课程的人，再看看谁不在这其中
SELECT *
FROM Student
WHERE sno NOT IN (
        SELECT sno
        FROM Score
        GROUP BY sno
        HAVING count(*) =(
                SELECT count(DISTINCT cno)
                FROM Course
            )
    );
-- 12、查询至少有一门课与学号为"01"的同学所学相同的同学的信息
-- 学号为01的同学学的课程编号
SELECT
	cno 
FROM
	Score a 
WHERE
	sno = '01';
-- 查出所有至少有一门课程和01号学生一样的学生编号
SELECT DISTINCT a.sno
FROM Score a
WHERE a.cno IN (
        SELECT a.cno
        FROM Score a
        WHERE a.sno = '01'
    );
-- solution-12-1
SELECT *
FROM Student
WHERE sno IN (
        SELECT DISTINCT a.sno
        FROM Score a
        WHERE a.cno IN (
                SELECT a.cno
                FROM Score a
                WHERE a.sno = '01'
            )
    );
-- solution-12-2
SELECT DISTINCT
	a.* 
FROM
	Student a,
	Score b 
WHERE
	b.sno = a.sno 
	AND b.cno IN ( SELECT cno FROM Score WHERE sno = '01' );
	
-- 13、查询和"01"号的同学学习的课程完全相同的其他同学的信息 

-- 学号为01的同学学的课程编号
SELECT
	cno 
FROM
	Score a 
WHERE
	sno = '01';
	
-- 1. 找到01同学学过的数量相同的同学A。
-- 2. 找到没学过01同学的课程的同学B。
-- 3. 在A不在B的同学就是我们要找的。
-- 4. 最后去掉01同学自己就行了。

-- solution-13-1

SELECT
d.*
FROM
Student d
WHERE
	d.sno IN (
		SELECT 
			a.sno
		FROM 
			Score a 
		GROUP BY a.sno
		HAVING COUNT(a.cno) = (SELECT COUNT(cno) FROM Score WHERE sno = '01')
	) 
	AND d.sno NOT IN (
		SELECT 
			b.sno
		FROM
			Score b 
		WHERE b.cno NOT IN (SELECT cno FROM Score WHERE sno = '01')
		GROUP BY b.sno
	)
	AND d.sno NOT IN ('01');

-- 除了01同学的其他同学学过的课 打组 
SELECT
		sno,
		group_concat( cno ORDER BY cno ) group1 
	FROM
		Score 
	WHERE
		sno != '01' 
	GROUP BY
		sno ;
		
-- 01同学学过的课 打组

	SELECT
		group_concat( cno ORDER BY cno ) group2 
	FROM
		Score 
	WHERE
		sno = '01' 
	GROUP BY
		sno;

-- solution-13-2
-- 找到两组人一样的部分 并结合学生表筛选出学生信息
SELECT
	a.* 
FROM
	(
SELECT sno,group_concat(cno ORDER BY cno) group1 FROM Score WHERE sno !='01' GROUP BY sno
	) t1
	INNER JOIN (
SELECT sno,group_concat(cno ORDER BY cno) group2 FROM Score WHERE sno='01' GROUP BY sno
	) t2 ON t1.group1 = t2.group2
	INNER JOIN Student a  ON a.sno = t1.sno;
	
-- 练习
SELECT
a.* 
FROM
(SELECT sno,GROUP_CONCAT(cno ORDER BY cno) group1 FROM Score WHERE sno !='01' GROUP BY sno) t1
INNER JOIN 
(SELECT sno,GROUP_CONCAT(cno ORDER BY cno) group2 FROM Score WHERE sno='01' GROUP BY sno) t2 
ON t1.group1=t2.group2
INNER JOIN Student a 
ON t1.sno = a.sno;

	
	
	

-- 14、查询没学过"张三"老师讲授的任一门课程的学生姓名 

-- 找到张三老师的编号
SELECT
	tno 
FROM
	Teacher 
WHERE
	tname = '张三';
	
-- 找到张三教的课程编号

SELECT
	cno 
FROM
	Course 
WHERE
	tno = ( SELECT tno FROM Teacher WHERE tname = '张三' );
	
-- 找到学过张三将的任意一门课的学生编号
SELECT
	sno 
FROM
	Score 
WHERE
	cno IN (
	SELECT
		cno 
	FROM
		Course 
	WHERE
	tno = ( SELECT tno FROM Teacher WHERE tname = '张三' ));

	
-- solution-14-1

SELECT a.* 
FROM Student a 
WHERE a.sno NOT IN (
	SELECT sno FROM Score WHERE cno IN (
			SELECT con FROM Course WHERE tno=(
						SELECT tno FROM Teacher WHERE tname='张三'))
	);

-- 练习
SELECT sname
FROM Student
WHERE
sno NOT IN (
	SELECT b.sno 
	FROM Score b 
	WHERE b.cno IN(
		SELECT cno 
		FROM Course 
		WHERE tno=(SELECT tno FROM Teacher WHERE tname = '张三')
	)
)

-- 分析
-- 1. 找出所有同学姓名 sname
-- 2. 找到老师的编号 tno
-- 3. 找到老师教授的课程编号 cno
-- 4. 找到学过老师教授的课程的学生 sno
-- 5. 找到所有学生里面没学过老师课程的学生。END


-- 15、查询两门及其以上不及格课程的同学的学号，姓名及其平均成绩 







-- solution-15-1
SELECT
	a.sno,
	a.sname,
	ROUND( AVG( b.sscore ), 2 ) 
FROM
	Student a,
	Score b 
WHERE
	a.sno = b.sno 
	AND a.sno IN (
	SELECT
		sno,
		COUNT(1)
	FROM
		Score 
	WHERE
		sscore < 60 GROUP BY sno  HAVING COUNT( 1 ) >= 2 
	) 
GROUP BY
	a.sno,
	a.sname;
		
		
		
-- 分析
-- 1. 找到成绩表里成绩不及格的学生编号列表。 sno 
-- 2. 找到学生表里对应编号的学生信息 sno， sname， AVG(sscore)

SELECT 
	a.sno,
	a.sname,
	ROUND(AVG(b.sscore),2) AS avg_score
FROM
	Student a,
	Score b
WHERE
 a.sno IN (
	 SELECT
	 sno
	 FROM
	 Score
	 WHERE
	 sscore < 60  GROUP BY sno HAVING COUNT(1) >=2
 )
 GROUP BY
	a.sno,
	a.sname;
	
	
-- 16、检索"01"课程分数小于60，按分数降序排列的学生信息


-- 分析
-- 1. 找出01课程对应的分数小于60的同学编号 sno  sscore
-- 2. 降序
-- solution-16-1
SELECT
 a.*,
 b.sscore  AS 01_score
FROM
	Student a,
	Score b
WHERE
 b.sno = a.sno
 AND b.cno = '01'
 AND b.sscore < 60
 ORDER BY b.sscore DESC;

	
-- 17、按平均成绩从高到低显示所有学生的所有课程的成绩以及平均成绩


SELECT
	b.*,
(SELECT sscore FROM Score WHERE sno=a.sno AND cno='01') AS 语文,
(SELECT sscore FROM Score WHERE sno=a.sno AND cno='02') AS 数学,
(SELECT sscore FROM Score WHERE sno=a.sno AND cno='03') AS 英语,
ROUND( avg( sscore ), 2 ) AS 平均分 
FROM
	Score a 
	RIGHT JOIN
	Student b
ON a.sno = b.sno
GROUP BY
	a.sno 
ORDER BY
	平均分 DESC;
	


-- @喝完这杯还有一箱的写法


SELECT 
b.*,
MAX(CASE a.cno WHEN '01' THEN a.sscore END) AS 语文,
MAX(CASE a.cno WHEN '02' THEN a.sscore END) AS 数学,
MAX(CASE a.cno WHEN '03' THEN a.sscore END) AS 英语,
ROUND(AVG(a.sscore),2) AS 平均分
FROM 
Score a RIGHT JOIN
Student b
ON a.sno=b.sno 
GROUP BY b.sno 
ORDER BY 平均分 DESC;

-- 18.查询各科成绩最高分、最低分和平均分：以如下形式显示：课程ID，课程name，最高分，最低分，平均分，及格率，中等率，优良率，优秀率 
-- 注： 及格为 >= 60 ， 中等为 ： 70 -80 ， 优良为 ： 80 -90 ， 优秀为 ： >= 90
SELECT 
	a.cno,
	b.cname,
	MAX(sscore) AS 最高分,
	MIN(sscore) AS  最低分,
	ROUND(AVG(sscore), 2) AS 平均分,
	ROUND(100*(SUM(CASE WHEN a.sscore>=60 THEN 1 ELSE 0 END)/SUM(CASE WHEN a.sscore THEN 1 ELSE 0 END)),2) AS 及格率,
	ROUND(100*(SUM(CASE WHEN a.sscore>=70 AND a.sscore<=80 THEN 1 ELSE 0 END)/SUM(CASE WHEN a.sscore THEN 1 ELSE 0 END)),2) AS 中等率,
	ROUND(100*(SUM(CASE WHEN a.sscore>=80 AND a.sscore<=90 THEN 1 ELSE 0 END)/SUM(CASE WHEN a.sscore THEN 1 ELSE 0 END)),2) AS 优良率,
	ROUND(100*(SUM(CASE WHEN a.sscore>=90 THEN 1 ELSE 0 END)/SUM(CASE WHEN a.sscore THEN 1 ELSE 0 END)),2) AS 优秀率
FROM Score a LEFT JOIN Course b ON a.cno=b.cno 
GROUP BY a.cno,b.cname
		
		
-- 19、按各科成绩进行排序，并显示排名


SELECT 
a.*,
COUNT(b.sscore)+1 AS 排名
FROM Score a  JOIN Score b 
WHERE a.sscore < b.sscore AND a.cno = b.cno
GROUP BY 
a.cno, 
a.sno
ORDER BY 
排名 ASC;

-- mysql没有rank函数
SELECT 
a.sno,
a.cno,
@i := @i+1 AS i保留排名,
@k := (CASE WHEN @Score  = a.sscore THEN @k ELSE @i END) AS rank不保留排名, 
@Score := a.sscore AS Score 
FROM 
(SELECT sno,cno,sscore FROM Score GROUP BY sno,cno,sscore ORDER BY sscore DESC) a,
(SELECT @k :=0,@i :=0,@Score :=0) s ;
		
-- @k1051785839的写法
(
SELECT*FROM (
	SELECT t1.cno,t1.sscore,(
		SELECT count(DISTINCT t2.sscore) FROM Score t2 WHERE t2.sscore>=t1.sscore AND t2.cno='01') rank FROM Score t1 WHERE t1.cno='01' ORDER BY t1.sscore DESC) t1) 
UNION (
SELECT*FROM (
	SELECT t1.cno,t1.sscore,(
		SELECT count(DISTINCT t2.sscore) FROM Score t2 WHERE t2.sscore>=t1.sscore AND t2.cno='02') rank FROM Score t1 WHERE t1.cno='02' ORDER BY t1.sscore DESC) t) 
UNION (
SELECT*FROM (
	SELECT t1.cno,t1.sscore,(
		SELECT count(DISTINCT t2.sscore) FROM Score t2 WHERE t2.sscore>=t1.sscore AND t2.cno='03') rank FROM Score t1 WHERE t1.cno='03' ORDER BY t1.sscore DESC) t3);
		
		
-- 20、查询学生的总成绩并进行排名

-- solution-20-1
SELECT 
a.sno,
@i := @i+1 AS i,
@k := (CASE WHEN @Score=a.sum_score THEN @k ELSE @i END) AS 排名,
@Score := a.sum_score AS Score
FROM 
(SELECT sno,SUM(sscore) AS sum_score FROM Score GROUP BY sno ORDER BY sum_score DESC) a,
(SELECT @k :=0,@i :=0,@Score :=0) s;

-- solution-20-2
SET @crank = 0;
SELECT
	q.sno,
	total,
	@crank := @crank + 1 AS 'rank' 
FROM
	(
	SELECT
		a.sno,
		SUM( a.sscore ) AS total 
	FROM
		Score a
	GROUP BY
		a.sno 
	ORDER BY
		total DESC 
	) q;


-- 21、查询不同老师所教不同课程平均分从高到低显示
SELECT a.tno,c.t_name,a.cno,ROUND(avg(sscore),2) AS avg_score FROM Course a LEFT JOIN Score b ON a.cno=b.cno LEFT JOIN Teacher c ON a.tno=c.tno GROUP BY a.cno,a.tno,c.t_name ORDER BY avg_score DESC;
-- 22、查询所有课程的成绩第2名到第3名的学生信息及该课程成绩
SELECT d.*,c.排名,c.sscore,c.cno FROM (SELECT a.sno,a.sscore,a.cno,@i :=@i+1 AS 排名 FROM Score a,(SELECT @i :=0) s WHERE a.cno='01' ORDER BY a.sscore DESC) c LEFT JOIN Student d ON c.sno=d.sno WHERE 排名 BETWEEN 2 AND 3 UNION SELECT d.*,c.排名,c.sscore,c.cno FROM (SELECT a.sno,a.sscore,a.cno,@j :=@j+1 AS 排名 FROM Score a,(SELECT @j :=0) s WHERE a.cno='02' ORDER BY a.sscore DESC) c LEFT JOIN Student d ON c.sno=d.sno WHERE 排名 BETWEEN 2 AND 3 UNION SELECT d.*,c.排名,c.sscore,c.cno FROM (SELECT a.sno,a.sscore,a.cno,@k :=@k+1 AS 排名 FROM Score a,(SELECT @k :=0) s WHERE a.cno='03' ORDER BY a.sscore DESC) c LEFT JOIN Student d ON c.sno=d.sno WHERE 排名 BETWEEN 2 AND 3;

-- 23、统计各科成绩各分数段人数：课程编号,课程名称,[100-85],[85-70],[70-60],[0-60]及所占百分比
SELECT DISTINCT f.cname,a.cno,b.`85-100`,b.百分比,c.`70-85`,c.百分比,d.`60-70`,d.百分比,e.`0-60`,e.百分比 FROM Score a LEFT JOIN (SELECT cno,SUM(CASE WHEN sscore> 85 AND sscore<=100 THEN 1 ELSE 0 END) AS `85-100`,ROUND(100*(SUM(CASE WHEN sscore> 85 AND sscore<=100 THEN 1 ELSE 0 END)/count(*)),2) AS 百分比 FROM Score GROUP BY cno) b ON a.cno=b.cno LEFT JOIN (SELECT cno,SUM(CASE WHEN sscore> 70 AND sscore<=85 THEN 1 ELSE 0 END) AS `70-85`,ROUND(100*(SUM(CASE WHEN sscore> 70 AND sscore<=85 THEN 1 ELSE 0 END)/count(*)),2) AS 百分比 FROM Score GROUP BY cno) c ON a.cno=c.cno LEFT JOIN (SELECT cno,SUM(CASE WHEN sscore> 60 AND sscore<=70 THEN 1 ELSE 0 END) AS `60-70`,ROUND(100*(SUM(CASE WHEN sscore> 60 AND sscore<=70 THEN 1 ELSE 0 END)/count(*)),2) AS 百分比 FROM Score GROUP BY cno) d ON a.cno=d.cno LEFT JOIN (SELECT cno,SUM(CASE WHEN sscore>=0 AND sscore<=60 THEN 1 ELSE 0 END) AS `0-60`,ROUND(100*(SUM(CASE WHEN sscore>=0 AND sscore<=60 THEN 1 ELSE 0 END)/count(*)),2) AS 百分比 FROM Score GROUP BY cno) e ON a.cno=e.cno LEFT JOIN Course f ON a.cno=f.cno;

-- 24、查询学生平均成绩及其名次
SELECT a.sno,@i :=@i+1 AS '不保留空缺排名',@k :=(CASE WHEN @avg_score=a.avg_s THEN @k ELSE @i END) AS '保留空缺排名',@avg_score :=avg_s AS '平均分' FROM (SELECT sno,ROUND(AVG(sscore),2) AS avg_s FROM Score GROUP BY sno ORDER BY avg_s DESC) a,(SELECT @avg_score :=0,@i :=0,@k :=0) b;

-- 25、查询各科成绩前三名的记录
-- 1.选出b表比a表成绩大的所有组
-- 2.选出比当前id成绩大的 小于三个的
SELECT a.sno,a.cno,a.sscore FROM Score a LEFT JOIN Score b ON a.cno=b.cno AND a.sscore< b.sscore GROUP BY a.sno,a.cno,a.sscore HAVING COUNT(b.sno)< 3 ORDER BY a.cno,a.sscore DESC;

-- 26、查询每门课程被选修的学生数
SELECT cno,count(sno) FROM Score a GROUP BY cno;

-- 27、查询出只有两门课程的全部学生的学号和姓名
SELECT sno,sname FROM Student WHERE sno IN (SELECT sno FROM Score GROUP BY sno HAVING COUNT(cno)=2);

-- 28、查询男生、女生人数
SELECT s_sex,COUNT(s_sex) AS 人数 FROM Student GROUP BY s_sex;

-- 29、查询名字中含有"风"字的学生信息
SELECT*FROM Student WHERE sname LIKE '%风%';

-- 30、查询同名同性学生名单，并统计同名人数
SELECT a.sname,a.s_sex,count(*) FROM Student a JOIN Student b ON a.sno !=b.sno AND a.sname=b.sname AND a.s_sex=b.s_sex GROUP BY a.sname,a.s_sex;

-- 31、查询1990年出生的学生名单
SELECT sname FROM Student WHERE s_birth LIKE '1990%';

-- 32、查询每门课程的平均成绩，结果按平均成绩降序排列，平均成绩相同时，按课程编号升序排列 
SELECT cno,ROUND(AVG(sscore),2) AS avg_score FROM Score GROUP BY cno ORDER BY avg_score DESC,cno ASC;
		
-- 33、查询平均成绩大于等于85的所有学生的学号、姓名和平均成绩
SELECT a.sno,b.sname,ROUND(avg(a.sscore),2) AS avg_score FROM Score a LEFT JOIN Student b ON a.sno=b.sno GROUP BY sno HAVING avg_score>=85;

-- 34、查询课程名称为"数学"，且分数低于60的学生姓名和分数
SELECT a.sname,b.sscore FROM Score b JOIN Student a ON a.sno=b.sno WHERE b.cno=(SELECT cno FROM Course WHERE cname='数学') AND b.sscore< 60;

-- 35、查询所有学生的课程及分数情况；
SELECT a.sno,a.sname,SUM(CASE c.cname WHEN '语文' THEN b.sscore ELSE 0 END) AS '语文',SUM(CASE c.cname WHEN '数学' THEN b.sscore ELSE 0 END) AS '数学',SUM(CASE c.cname WHEN '英语' THEN b.sscore ELSE 0 END) AS '英语',SUM(b.sscore) AS '总分' FROM Student a LEFT JOIN Score b ON a.sno=b.sno LEFT JOIN Course c ON b.cno=c.cno GROUP BY a.sno,a.sname;
		
-- 36、查询任何一门课程成绩在70分以上的姓名、课程名称和分数；
SELECT a.sname,b.cname,c.sscore FROM Course b LEFT JOIN Score c ON b.cno=c.cno LEFT JOIN Student a ON a.sno=c.sno WHERE c.sscore>=70

-- 37、查询不及格的课程
SELECT a.sno,a.cno,b.cname,a.sscore FROM Score a LEFT JOIN Course b ON a.cno=b.cno WHERE a.sscore< 60

 -- 38、查询课程编号为01且课程成绩在80分以上的学生的学号和姓名； 
SELECT a.sno,b.sname FROM Score a LEFT JOIN Student b ON a.sno=b.sno WHERE a.cno='01' AND a.sscore> 80

-- 39、求每门课程的学生人数
SELECT count(*) FROM Score GROUP BY cno;

-- 40、查询选修"张三"老师所授课程的学生中，成绩最高的学生信息及其成绩
-- 查询老师id
SELECT cno FROM Course c,Teacher d WHERE c.tno=d.tno AND d.t_name='张三'-- 查询最高分（可能有相同分数）
SELECT MAX(sscore) FROM Score WHERE cno='02'-- 查询信息
SELECT a.*,b.sscore,b.cno,c.cname FROM Student a LEFT JOIN Score b ON a.sno=b.sno LEFT JOIN Course c ON b.cno=c.cno WHERE b.cno=(SELECT cno FROM Course c,Teacher d WHERE c.tno=d.tno AND d.t_name='张三') AND b.sscore IN (SELECT MAX(sscore) FROM Score WHERE cno='02');

-- 41、查询不同课程成绩相同的学生的学生编号、课程编号、学生成绩 
SELECT DISTINCT b.sno,b.cno,b.sscore FROM Score a,Score b WHERE a.cno !=b.cno AND a.sscore=b.sscore;
		
-- 42、查询每门功成绩最好的前两名 
-- 牛逼的写法
SELECT a.sno,a.cno,a.sscore FROM Score a WHERE (SELECT COUNT(1) FROM Score b WHERE b.cno=a.cno AND b.sscore>=a.sscore)<=2 ORDER BY a.cno;

-- 43、统计每门课程的学生选修人数（超过5人的课程才统计）。要求输出课程号和选修人数，查询结果按人数降序排列，若人数相同，按课程号升序排列  
SELECT cno,count(*) AS total FROM Score GROUP BY cno HAVING total> 5 ORDER BY total,cno ASC;
		
-- 44、检索至少选修两门课程的学生学号 
SELECT sno,count(*) AS sel FROM Score GROUP BY sno HAVING sel>=2;

-- 45、查询选修了全部课程的学生信息 
SELECT*FROM Student WHERE sno IN (SELECT sno FROM Score GROUP BY sno HAVING count(*)=(SELECT count(*) FROM Course));

-- 46、查询各学生的年龄
-- 按照出生日期来算，当前月日 < 出生年月的月日则，年龄减一
SELECT s_birth,(DATE_FORMAT(NOW(),'%Y')-DATE_FORMAT(s_birth,'%Y')-(CASE WHEN DATE_FORMAT(NOW(),'%m%d')> DATE_FORMAT(s_birth,'%m%d') THEN 0 ELSE 1 END)) AS age FROM Student;

-- 47、查询本周过生日的学生
SELECT*FROM Student WHERE WEEK (DATE_FORMAT(NOW(),'%Y%m%d'))=WEEK (s_birth) SELECT*FROM Student WHERE YEARWEEK(s_birth)=YEARWEEK(DATE_FORMAT(NOW(),'%Y%m%d')) SELECT WEEK (DATE_FORMAT(NOW(),'%Y%m%d'));

-- 48、查询下周过生日的学生
SELECT*FROM Student WHERE WEEK (DATE_FORMAT(NOW(),'%Y%m%d'))+1=WEEK (s_birth);

-- 49、查询本月过生日的学生
SELECT*FROM Student WHERE MONTH (DATE_FORMAT(NOW(),'%Y%m%d'))=MONTH (s_birth);

-- 50、查询下月过生日的学生
SELECT*FROM Student WHERE MONTH (DATE_FORMAT(NOW(),'%Y%m%d'))+1=MONTH (s_birth);