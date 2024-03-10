<center>MySQL学习</center>

## 一、MySQL的前世今生

MySQL 是一款开源的关系型数据库管理系统（RDBMS），具有广泛的应用和用户群体。

1. 发展起源：MySQL 最初由瑞典的 MySQL AB 公司开发，由 Michael Widenius、David Axmark 和 Allan Larsson 创立。其初衷是为了提供一个轻量级、易于使用和高性能的数据库解决方案。
2. 成为事实上的开源数据库标准：由于其开源的特性、性能优势和广泛的社区支持，MySQL 逐渐成为事实上的开源数据库标准。它在 Web 应用、企业应用和云计算等领域得到广泛应用。
3. MySQL 的版本发展：MySQL 经历了多个版本的迭代和发展。其中，MySQL 5.1 提供了许多重要的功能和性能改进，MySQL 5.5 引入了 InnoDB 存储引擎作为默认引擎，MySQL 5.6 和 5.7 增加了更多的功能和优化，MySQL 8.0 引入了一些重要的特性，如窗口函数、CTE（公共表达式）等。
4. 衍生版本和分支：由于 MySQL 的开源性质，衍生出了许多基于 MySQL 的分支版本，例如 MariaDB、Percona Server 等。这些分支版本在功能和性能上做出了一些改进和扩展，丰富了 MySQL 生态系统。
5. MySQL 的应用领域：MySQL 在 Web 应用、企业应用、嵌入式系统和大数据等领域得到广泛应用。它被许多知名公司和组织使用，如 Facebook、Twitter、Uber等。

## 二、Go连接&操作MySQL

原生支持连接池，是并发安全的





## 三、MySQL少部分内容学习

ps:记得加（;）！！！

* 创建库

```sql
CREATE DATABASE YiSheng;
```

* 删除库

```sql
DROP DATABASE YiSheng;
```



### 1.数据表结构

数据表是 MySQL 中存储数据的基本单位。在创建数据表时，需要定义表的结构，包括表名、列名和列的数据类型等。eg：

```sql
CREATE TABLE students (
  id INT PRIMARY KEY,
  name VARCHAR(50),
  age INT,
  gender ENUM('Male', 'Female')
);
```

### 2.不同表之间的关联

在数据库设计中，不同的表之间可以建立关联关系，以便进行数据的关联查询和操作。常见的关联类型包括一对一关系、一对多关系和多对多关系。eg：

```sql
CREATE TABLE books (
  id INT PRIMARY KEY,
  title VARCHAR(100),
  author_id INT,
  FOREIGN KEY (author_id) REFERENCES authors(id)
);

CREATE TABLE authors (
  id INT PRIMARY KEY,
  name VARCHAR(50)
);
```

### 3.SQL 查询语句（CURD）

SQL 是用于在关系型数据库中执行各种操作的查询语言。CURD 是指对数据库的**增加(Create)**、**查询(Read)**、**更新(Update)**和**删除(Delete)**操作

#### 1.创建表：

```go
type User struct {
    gorm.Model
    Name  string
    Email string
}

// 迁移（创建）User表
err = db.AutoMigrate(&User{})
if err != nil {
    panic(err)
}
```

#### 2.插入数据：

```go
user := User{Name: "John Doe", Email: "john@example.com"}

// 创建记录
result := db.Create(&user)
if result.Error != nil {
    panic(result.Error)
}
```

#### 3.查询数据：

```go
// 查询单个记录
var user User
result := db.First(&user, 1) // 根据主键查询
if result.Error != nil {
    panic(result.Error)
}

// 查询多个记录
var users []User
result := db.Find(&users)
if result.Error != nil {
    panic(result.Error)
}
```

#### 4.更新数据：

```go
// 更新单个记录
result := db.Model(&user).Update("Name", "Jane Doe")
if result.Error != nil {
    panic(result.Error)
}

// 更新多个记录
result := db.Model(&User{}).Where("age < ?", 18).Update("Name", "Child")
if result.Error != nil {
    panic(result.Error)
}
```

#### 5.删除数据：

```go
// 删除单个记录
result := db.Delete(&user)
if result.Error != nil {
    panic(result.Error)
}

// 删除多个记录
result := db.Where("age > ?", 60).Delete(&User{})
if result.Error != nil {
    panic(result.Error)
}
```

## 四、后续继续了解

* ACID

1. 原子性：事务要么成功要么失败，没有中间状态
2. 一致性：数据库的完整性没有被破坏
3. 隔离性：事务之间是互相隔离的
4. 持久性：事务操作的结果是不会丢失的

* 索引

1. 索引的原理：B数和B+数
2. 索引的类型
3. 索引的命中
4. 分库分表
5. SQL注入
6. SQL慢查询优化
7. MySQL主从：

​		binlog

8. MySQL读写分离
