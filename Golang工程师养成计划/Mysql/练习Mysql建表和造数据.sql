USE `school`;

DROP TABLE IF EXISTS `Student`;
CREATE TABLE `Student` (
	`sno` VARCHAR ( 20 ) COMMENT '学生编号',
	`sname` VARCHAR ( 20 ) NOT NULL DEFAULT 'xx' COMMENT '学生姓名',
	`sage` VARCHAR ( 20 ) NOT NULL DEFAULT '1970-01-01' COMMENT '生日年月日',
	`ssex` VARCHAR ( 20 ) NOT NULL DEFAULT '女' COMMENT '性别',
PRIMARY KEY ( `sno` )) ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT = '学生表';

DROP TABLE IF EXISTS `Teacher`;
CREATE TABLE `Teacher` (
	`tno` VARCHAR ( 20 ) COMMENT '教师编号',
	`tname` VARCHAR ( 20 ) NOT NULL DEFAULT '' COMMENT '教师姓名',
PRIMARY KEY ( `tno` )) ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT = '教师表';

DROP TABLE IF EXISTS `Course`;
CREATE TABLE `Course` (
	`cno` VARCHAR ( 20 ) COMMENT '课程编号',
	`cname` VARCHAR ( 20 ) NOT NULL DEFAULT '' COMMENT '课程名称',
	`tno` VARCHAR ( 20 ) NOT NULL DEFAULT '' COMMENT '教师编号',
PRIMARY KEY ( `cno` )) ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT = '课程表';

DROP TABLE IF EXISTS `Score`;
CREATE TABLE `Score` (
	`sno` VARCHAR ( 20 ) COMMENT '学生编号',
	`cno` VARCHAR ( 20 ) COMMENT '课程编号',
	`sscore` INT ( 3 ) COMMENT '所得分数',
PRIMARY KEY ( `sno`, `cno` )) ENGINE = INNODB DEFAULT CHARSET = utf8 COMMENT = '成绩表';

-- 学生测试数据 
INSERT INTO Student
VALUES
	( '01', '赵雷', '1990-01-01', '男' );
INSERT INTO Student
VALUES
	( '02', '钱电', '1990-12-21', '男' );
INSERT INTO Student
VALUES
	( '03', '孙风', '1990-05-20', '男' );
INSERT INTO Student
VALUES
	( '04', '李云', '1990-08-06', '男' );
INSERT INTO Student
VALUES
	( '05', '周梅', '1991-12-01', '女' );
INSERT INTO Student
VALUES
	( '06', '吴兰', '1992-03-01', '女' );
INSERT INTO Student
VALUES
	( '07', '郑竹', '1989-07-01', '女' );
INSERT INTO Student
VALUES
	( '08', '王菊', '1990-01-20', '女' );

-- 课程表测试数据 
INSERT INTO Course
VALUES
	( '01', '语文', '02' );
INSERT INTO Course
VALUES
	( '02', '数学', '01' );
INSERT INTO Course
VALUES
	( '03', '英语', '03' );

-- 教师表测试数据 
INSERT INTO Teacher
VALUES
	( '01', '张三' );
INSERT INTO Teacher
VALUES
	( '02', '李四' );
INSERT INTO Teacher
VALUES
	( '03', '王五' );

-- 成绩表测试数据 
INSERT INTO Score
VALUES
	( '01', '01', 80 );
INSERT INTO Score
VALUES
	( '01', '02', 90 );
INSERT INTO Score
VALUES
	( '01', '03', 99 );
INSERT INTO Score
VALUES
	( '02', '01', 70 );
INSERT INTO Score
VALUES
	( '02', '02', 60 );
INSERT INTO Score
VALUES
	( '02', '03', 80 );
INSERT INTO Score
VALUES
	( '03', '01', 80 );
INSERT INTO Score
VALUES
	( '03', '02', 80 );
INSERT INTO Score
VALUES
	( '03', '03', 80 );
INSERT INTO Score
VALUES
	( '04', '01', 50 );
INSERT INTO Score
VALUES
	( '04', '02', 30 );
INSERT INTO Score
VALUES
	( '04', '03', 20 );
INSERT INTO Score
VALUES
	( '05', '01', 76 );
INSERT INTO Score
VALUES
	( '05', '02', 87 );
INSERT INTO Score
VALUES
	( '06', '01', 31 );
INSERT INTO Score
VALUES
	( '06', '03', 34 );
INSERT INTO Score
VALUES
	( '07', '02', 89 );
INSERT INTO Score
VALUES
	( '07', '03', 98 );