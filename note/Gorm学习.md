<center>Gorm学习</center>

**戈姆**

- ORM和Gorm：
  - ORM（Object-Relational Mapping）概念：理解ORM的作用和原理。
  - 为什么选择Gorm：了解使用Gorm的优势。
- 安装和配置Gorm：
  - 安装Gorm：了解如何在Go项目中添加Gorm。
  - 数据库连接：学习如何配置和连接到MySQL的数据库。
- Gorm基础操作：
  - 模型定义：使用结构体定义数据库表。
  - CRUD操作：创建（Create）、读取（Read）、更新（Update）、删除（Delete）数据。
  - 查询：简单查询、条件查询、排序、分页等。

## 一、ORM和Grom

* ORM（Object-Relational Mapping）概念

​	ORM（对象关系映射）是一种编程技术，它将面向对象的编程语言与关系型数据库之间进行映射，从而实现将对象模型与数据库模型进行交互的过程。ORM 的主要目标是简化开发人员在应用程序和数据库之间的数据交互，提供一种面向对象的方式来处理数据库操作，而无需直接编写和执行 SQL 查询。

* ORM 的作用和原理：

1. 数据模型映射：ORM 将应用程序中的对象模型与数据库中的表和列进行映射。它通过定义对象和表之间的映射关系，使开发人员可以使用面向对象的方式来操作数据库，而无需关注底层的 SQL 查询和数据库操作。
2. 数据操作：ORM 提供了一组方法和 API，用于执行常见的数据库操作，如插入、更新、删除和查询数据。开发人员可以使用对象或类的方法来执行这些操作，而不必直接编写 SQL 查询。
3. 关系处理：ORM 可以处理对象之间的关系，如一对一、一对多和多对多关系。它提供了一种简化的方式来管理对象之间的关联，并自动处理关联的数据操作。
4. 数据库迁移：ORM 通常提供数据库迁移工具，用于管理数据库模式的变化。它可以自动创建或更新数据库表结构，以便与对象模型保持同步

*  Gorm 作为 ORM 框架的优势

1. 简单易用：Gorm 提供了简洁、直观的 API，使得使用它来执行数据库操作变得非常简单。它的语法和用法都很容易理解，降低了学习和使用的门槛。
2. 功能丰富：Gorm 提供了丰富的功能集，包括数据关联处理、预加载、事务支持、数据库迁移等。它支持多种数据库后端，如 MySQL、PostgreSQL、SQLite 等，以及常见的数据库操作，如查询、插入、更新和删除。
3. 性能优化：Gorm 在性能方面进行了优化，具有高效的查询生成和数据加载机制。它提供了一些性能调优选项，如缓存、批量操作等，以提高应用程序的性能。
4. 社区支持和活跃度：Gorm 是一个受欢迎的开源项目，拥有活跃的社区支持。它的文档齐全，有很多示例和教程可供参考，开发者可以轻松地获取帮助和支持。

## 二、安装和配置Gorm

要在 Go 项目中使用 Gorm，需要先安装 Gorm 包及其依赖

1. 打开终端或命令提示符。
2. 运行以下命令来安装 Gorm 包以及数据库驱动：

```shell

go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

这将下载并安装 Gorm 包及其相关依赖。

```go
import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)
```

下列例子为使用Grom连接数据库

```go
package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 在这里进行数据库操作
}
```

在上述代码中，需要将 `user` 和 `password` 替换为实际的 MySQL 用户名和密码，`database` 替换为要连接的数据库名称。

## 三、Gorm基础操作

`db` 变量是一个 Gorm 的数据库连接对象，可以根据实际情况进行初始化和配置。

```go
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
```

### 1.模型定义

在 Gorm 中，可以使用结构体定义数据库表。每个字段代表表的一个列，字段的标签（tag）提供了与数据库列的映射关系。

```go
type User struct {
    ID   uint
    Name string
    Age  int
}
```

在上述示例中，定义了一个名为 `User` 的结构体，它将映射到数据库中的 `users` 表。结构体的字段 `ID`、`Name` 和 `Age` 分别对应表中的列。

### 2.CRUD操作

- 创建（Create）数据：

```go
user := User{Name: "John", Age: 25}
db.Create(&user)
```

上述代码将创建一个名为 `user` 的新记录，并将其插入到数据库中。

- 读取（Read）数据：

```go
var user User
db.First(&user, 1) // 根据主键查询第一条记录
db.Find(&users)    // 查询所有记录
```

上述代码演示了如何根据主键或查询条件从数据库中读取数据。`First` 方法将返回满足查询条件的第一条记录，而 `Find` 方法将返回所有符合条件的记录。

- 更新（Update）数据：

```go
db.Model(&user).Update("Age", 30)
```

上述代码将更新 `user` 对象在数据库中的 `Age` 字段为 30。

- 删除（Delete）数据：

```go
db.Delete(&user)
```

上述代码将从数据库中删除 `user` 对象对应的记录。

### 3.查询

- 简单查询：

```go
var users []User
db.Find(&users) // 查询所有记录
```

上述代码将从数据库中获取所有 `User` 对象的记录，并将其存储在 `users` 切片中。

- 条件查询：

```go
var user User
db.Where("name = ?", "John").First(&user)
```

上述代码将根据条件查询数据库中符合条件的第一条记录。

- 排序：

```go
var users []User
db.Order("age desc").Find(&users)
```

上述代码将根据 `age` 字段的降序对记录进行排序。

- 分页：

```go
var users []User
db.Limit(10).Offset(20).Find(&users)
```

上述代码将从数据库中获取第 20 条记录开始的 10 条记录。

上述示例中的 

## 四、模型定义

### 1.命名策略

gorm采用的命名策略是，表名是蛇形复数，字段是蛇形单数

```go
 type Student struct{   
 	Name string    
 	Age  int
 	MyStudent string
 }
```

### 2.模型定义

模型使用普通结构体定义。这些结构可以包含具有基本 Go 类型的字段、这些类型的指针或别名，甚至是自定义类型，只要它们实现包中的 [Scanner](https://pkg.go.dev/database/sql/?tab=doc#Scanner) 和 [Valuer](https://pkg.go.dev/database/sql/driver#Valuer) 接口即可`database/sql`

```go
type User struct {
  ID           uint           // Standard field for the primary key
  Name         string         // A regular string field
  Email        *string        // A pointer to a string, allowing for null values
  Age          uint8          // An unsigned 8-bit integer
  Birthday     *time.Time     // A pointer to time.Time, can be null
  MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
  ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
  CreatedAt    time.Time      // Automatically managed by GORM for creation time
  UpdatedAt    time.Time      // Automatically managed by GORM for update time
}
```



### 3.字段级权限控制

GORM 允许用标签控制字段级别的权限。 这样就可以让一个字段的权限是只读、只写、只创建、只更新或者被忽略

**注意：** 使用 GORM Migrator 创建表时，不会创建被忽略的字段

```go
type User struct {
  Name string `gorm:"<-:create"` // 允许读和创建
  Name string `gorm:"<-:update"` // 允许读和更新
  Name string `gorm:"<-"`        // 允许读和写（创建和更新）
  Name string `gorm:"<-:false"`  // 允许读，禁止写
  Name string `gorm:"->"`        // 只读（除非有自定义配置，否则禁止写）
  Name string `gorm:"->;<-:create"` // 允许读和写
  Name string `gorm:"->:false;<-:create"` // 仅创建（禁止从 db 读）
  Name string `gorm:"-"`  // 通过 struct 读写会忽略该字段
  Name string `gorm:"-:all"`        // 通过 struct 读写、迁移会忽略该字段
  Name string `gorm:"-:migration"`  // 通过 struct 迁移会忽略该字段
}
```

### 4.创建/更新时间追踪（纳秒、毫秒、秒、Time）

GORM 约定使用 、 追踪创建/更新时间。 如果您定义了这种字段，GORM 在创建、更新时会自动填充 [当前时间](https://gorm.io/zh_CN/docs/gorm_config.html#now_func)`CreatedAt``UpdatedAt`

```go
type User struct {
  CreatedAt time.Time // 在创建时，如果该字段值为零值，则使用当前时间填充
  UpdatedAt int       // 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
  Updated   int64 `gorm:"autoUpdateTime:nano"` // 使用时间戳纳秒数填充更新时间
  Updated   int64 `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
  Created   int64 `gorm:"autoCreateTime"`      // 使用时间戳秒数填充创建时间
}
```

### 5.嵌入结构体

对于匿名字段，GORM 会将其字段包含在父结构体中，例如：

```go
type User struct {
  gorm.Model
  Name string
}
// 等效于
type User struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
  Name string
}
```

对于正常的结构体字段，可以通过标签 将其嵌入，例如：`embedded`

```go
type Author struct {
    Name  string
    Email string
}

type Blog struct {
  ID      int
  Author  Author `gorm:"embedded"`
  Upvotes int32
}
// 等效于
type Blog struct {
  ID    int64
  Name  string
  Email string
  Upvotes  int32
}
```

可以使用标签 来为 db 中的字段名添加前缀，例如：`embeddedPrefix`

```go
type Blog struct {
  ID      int
  Author  Author `gorm:"embedded;embeddedPrefix:author_"`
  Upvotes int32
}
// 等效于
type Blog struct {
  ID          int64
  AuthorName string
  AuthorEmail string
  Upvotes     int32
}
```

### 6.字段标签

声明 model 时，tag 是可选的，GORM 支持以下 tag： tag 名大小写不敏感，但建议使用 风格`camelCase`

| 标签名                          | 说明                                                         |
| :------------------------------ | :----------------------------------------------------------- |
| 列                              | 指定 db 列名                                                 |
| 类型                            | 列数据类型，推荐使用兼容性好的通用类型，例如：所有数据库都支持 bool、int、uint、float、string、time、bytes 并且可以和其他标签一起使用，例如：、, … 像 这样指定数据库数据类型也是支持的。 在使用指定数据库数据类型时，它需要是完整的数据库数据类型，如：`not null``size``autoIncrement``varbinary(8)``MEDIUMINT UNSIGNED not NULL AUTO_INCREMENT` |
| 序列化程序                      | 指定将数据序列化或反序列化到数据库中的序列化器, 例如: `serializer:json/gob/unixtime` |
| 大小                            | 定义列数据类型的大小或长度，例如 `size: 256`                 |
| primaryKey                      | 将列定义为主键                                               |
| 独特                            | 将列定义为唯一键                                             |
| 违约                            | 定义列的默认值                                               |
| 精度                            | 指定列的精度                                                 |
| 规模                            | 指定列大小                                                   |
| 非空                            | 指定列为 NOT NULL                                            |
| 自动增量                        | 指定列为自动增长                                             |
| autoIncrement增量               | 自动步长，控制连续记录之间的间隔                             |
| 嵌入式                          | 嵌套字段                                                     |
| embedded前缀                    | 嵌入字段的列名前缀                                           |
| autoCreateTime（自动创建时间）  | 创建时追踪当前时间，对于 字段，它会追踪时间戳秒数，您可以使用 / 来追踪纳秒、毫秒时间戳，例如：`int``nano``milli``autoCreateTime:nano` |
| autoUpdateTime （自动更新时间） | 创建/更新时追踪当前时间，对于 字段，它会追踪时间戳秒数，您可以使用 / 来追踪纳秒、毫秒时间戳，例如：`int``nano``milli``autoUpdateTime:milli` |
| 指数                            | 根据参数创建索引，多个字段使用相同的名称则创建复合索引，查看 [索引](https://gorm.io/zh_CN/docs/indexes.html) 获取详情 |
| 唯一索引                        | 与 相同，但创建的是唯一索引`index`                           |
| 检查                            | 创建检查约束，例如 ，查看 [约束](https://gorm.io/zh_CN/docs/constraints.html) 获取详情`check:age > 13` |
| <-                              | 设置字段写入的权限， 只创建、 只更新、 无写入权限、 创建和更新权限`<-:create``<-:update``<-:false``<-` |
| ->                              | 设置字段读的权限， 无读权限`->:false`                        |
| -                               | 忽略该字段， 表示无读写， 表示无迁移权限， 表示无读写迁移权限`-``-:migration``-:all` |
| 评论                            | 迁移时为字段添加注释                                         |

## 五、CRUD接口

### 1.创建

#### 1.创建记录

```go
user := User{Name: "Yisheng", Age: 18, Birthday: time.Now()}

result := db.Create(&user) // 通过数据的指针来创建

user.ID             // 返回插入数据的主键
result.Error        // 返回 error
result.RowsAffected // 返回插入记录的条数
```

还可以使用 `Create()` 创建多项记录：

```go
users := []*User{
    User{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
    User{Name: "Jackson", Age: 19, Birthday: time.Now()},
}

result := db.Create(users) // 传递切片以插入多行数据

result.Error        // 返回 error
result.RowsAffected // 返回插入记录的条数
```

注意：无法向‘create’传递结构体，所以应该传入数据的指针

#### 2.用指定的字段创建记录

创建记录并为指定字段赋值。

```go
db.Select("Name", "Age", "CreatedAt").Create(&user)
// INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")
```

创建记录并忽略传递给 ‘Omit’ 的字段值

```go
db.Omit("Name", "Age", "CreatedAt").Create(&user)
// INSERT INTO `users` (`birthday`,`updated_at`) VALUES ("2020-01-01 00:00:00.000", "2020-07-04 11:05:21.775")
```

#### 3.批量插入

要高效地插入大量记录，请将切片传递给`Create`方法。 GORM 将生成一条 SQL 来插入所有数据，以返回所有主键值，并触发 `Hook` 方法。 当这些记录可以被分割成多个批次时，GORM会开启一个事务</0>来处理它们。

```go
var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
db.Create(&users)

for _, user := range users {
  user.ID // 1,2,3
}
```

你可以通过`db.CreateInBatches`方法来指定批量插入的批次大小

```go
var users = []User{{Name: "jinzhu_1"}, ...., {Name: "jinzhu_10000"}}

// batch size 100
db.CreateInBatches(users, 100)
```

#### 4.创建钩子（Hook方法）

GROM允许用户通过实现这些接口 `BeforeSave`, `BeforeCreate`, `AfterSave`, `AfterCreate`来自定义钩子。 这些钩子方法会在创建一条记录时被调用，关于钩子的生命周期请参阅[Hooks](https://gorm.io/zh_CN/docs/hooks.html)。

```go
func (u *User) BeforeCreate(tx *gorm.DB)(err error) {// 接收一个指向 gorm.DB 对象的指针作为参数，并返回一个 error 类型的值。 
  u.UUID = uuid.New()//这行代码使用 uuid.New() 函数生成一个新的 UUID（通用唯一标识符），并将其赋值给 User 结构体中的 UUID 字段。这可以用来为新创建的记录生成一个唯一标识符。

    if u.Role == "admin" {
        return errors.New("invalid role")
    }
    return
}//可以在创建之前检查是否已经存在
```

如果想跳过`Hooks`方法，可以使用`SkipHooks`会话模式，例子如下

```go
DB.Session(&gorm.Session{SkipHooks: true}).Create(&user)

DB.Session(&gorm.Session{SkipHooks: true}).Create(&users)

DB.Session(&gorm.Session{SkipHooks: true}).CreateInBatches(users, 100)
```

#### 5.根据 Map 创建

GORM支持通过 `map[string]interface{}` 与 `[]map[string]interface{}{}`来创建记录。

```go
db.Model(&User{}).Create(map[string]interface{}{
  "Name": "jinzhu", "Age": 18,
})

// batch insert from `[]map[string]interface{}{}`
db.Model(&User{}).Create([]map[string]interface{}{//这是 Create 方法的调用，用于插入多个记录。它接收一个 []map[string]interface{} 类型(地图切片)的参数，其中每个 map 对象表示一个记录，键值对表示要插入的列和对应的值。
  {"Name": "jinzhu_1", "Age": 18},
  {"Name": "jinzhu_2", "Age": 20},
})
```

注意：当使用map来创建时，钩子方法不会执行，关联不会被保存且不会回写主键

#### 6.默认值

可以通过结构体Tag `default`来定义字段的默认值，示例如下：

```go
type User struct {
  ID   int64
  Name string `gorm:"default:galeone"`
  Age  int64  `gorm:"default:18"`
}
```

### 2.查询

由此进一步理解数据库中数据的动态过程，查询不是字面意义的查询，而是将库中的数据提取出来放入新的变量，然后通过新变量输出在控制台中显示查询的数据（此过程中内存的占用等一类问题仍需理解）

#### 1.检索

`First` and `Last` 方法会按主键排序找到第一条记录和最后一条记录 (分别)。 只有在目标 struct 是指针或者通过 `db.Model()` 指定 model 时，该方法才有效。 此外，如果相关 model 没有定义主键，那么将按 model 的第一个字段进行排序，如下的例子中皆将数据提取到心得变量里，可用fmt打印

```go
var user User
var users []User

// 获取第一条记录（主键升序）
db.First(&user)
// SELECT * FROM users ORDER BY id LIMIT 1;

// 获取一条记录，没有指定排序字段
db.Take(&user)
// SELECT * FROM users LIMIT 1;

// 获取最后一条记录（主键降序）
db.Last(&user)
// SELECT * FROM users ORDER BY id DESC LIMIT 1;

//根据主键检索
db.First(&user, 10)
// SELECT * FROM users WHERE id = 10;

db.Find(&users, []int{1,2,3})
// SELECT * FROM users WHERE id IN (1,2,3);

//检索全部对象
db.Find(&users)
// SELECT * FROM users;
```

#### 2.条件

##### 1.String 条件

```go
// Get first matched record
db.Where("name = ?", "jinzhu").First(&user)
// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;

// Get all matched records
db.Where("name <> ?", "jinzhu").Find(&users)
// SELECT * FROM users WHERE name <> 'jinzhu';

// IN
db.Where("name IN ?", []string{"jinzhu", "jinzhu 2"}).Find(&users)
// SELECT * FROM users WHERE name IN ('jinzhu','jinzhu 2');

// LIKE
db.Where("name LIKE ?", "%jin%").Find(&users)
// SELECT * FROM users WHERE name LIKE '%jin%';

// AND
db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' AND age >= 22;

// Time
db.Where("updated_at > ?", lastWeek).Find(&users)
// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

// BETWEEN
db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-08 00:00:00';
```

如果对象设置了主键，条件查询将不会覆盖主键的值，而是用 And 连接条件。 例如：

```go
var user = User{ID: 10}
db.Where("id = ?", 20).First(&user)
// SELECT * FROM users WHERE id = 10 and id = 20 ORDER BY id ASC LIMIT 1
```

##### 2.Struct & Map 条件

```go
// Struct
db.Where(&User{Name: "jinzhu", Age: 20}).First(&user)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 ORDER BY id LIMIT 1;

// Map
db.Where(map[string]interface{}{"name": "jinzhu", "age": 20}).Find(&users)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 20;

// Slice of primary keys
db.Where([]int64{20, 21, 22}).Find(&users)
// SELECT * FROM users WHERE id IN (20, 21, 22);

//注意使用 struct 进行查询时，GORM 将仅使用非零字段进行查询，这意味着如果字段的值为 、 或其他零值，则不会用于构建查询条件，例如：0''false
```

##### 3.指定结构体查询字段

使用 struct 进行搜索时，可以通过将相关字段名称或 dbname 传递给 来指定要在查询条件中使用的结构中的哪些特定值，例如：`Where()`

```go
db.Where(&User{Name: "jinzhu"}, "name", "Age").Find(&users)
// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;

db.Where(&User{Name: "jinzhu"}, "Age").Find(&users)
// SELECT * FROM users WHERE age = 0;
```

##### 4.内联条件

查询条件可以内联到方法中，类似于 和 。`First``Find``Where`

```go
// Get by primary key if it were a non-integer type
db.First(&users, "id = ?", "string_primary_key")
// SELECT * FROM users WHERE id = 'string_primary_key';

// Plain SQL
db.Find(&users, "name = ?", "jinzhu")
// SELECT * FROM users WHERE name = "jinzhu";

db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

// Struct
db.Find(&users, User{Age: 20})
// SELECT * FROM users WHERE age = 20;

// Map
db.Find(&users, map[string]interface{}{"age": 20})
// SELECT * FROM users WHERE age = 20;
```

##### 5.Not 条件

构建 NOT 条件，工作原理类似于`Where`

```go
db.Not("name = ?", "jinzhu").First(&user)
// SELECT * FROM users WHERE NOT name = "jinzhu" ORDER BY id LIMIT 1;

// Not In
db.Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&users)
// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

// Struct
db.Not(User{Name: "jinzhu", Age: 18}).First(&user)
// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;

// Not In slice of primary keys
db.Not([]int64{1,2,3}).First(&user)
// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;
```

##### 6.Or 条件

```go
db.Where("role = ?", "admin").Or("role = ?", "super_admin").Find(&users)
// SELECT * FROM users WHERE role = 'admin' OR role = 'super_admin';

// Struct
db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);

// Map
db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)
// SELECT * FROM users WHERE name = 'jinzhu' OR (name = 'jinzhu 2' AND age = 18);
```

#### 3.其他（仍需补充）

##### 1.选择特定字段

`Select`允许您指定要从数据库中检索的字段。否则，GORM 将默认选择所有字段。

```go
db.Select("name", "age").Find(&users)
// SELECT name, age FROM users;

db.Select([]string{"name", "age"}).Find(&users)
// SELECT name, age FROM users;

db.Table("users").Select("COALESCE(age,?)", 42).Rows()
// SELECT COALESCE(age,'42') FROM users;
```

##### 2.排序

指定从数据库检索记录时的顺序


在 GORM 中，可以使用 `Order` 方法对查询结果进行排序。该方法接受一个字符串作为参数，用于指定排序的规则。

例如，假设我们有一个 `User` 结构体，其中包含 `Name` 字段，我们想按照用户姓名的字母顺序对查询结果进行排序，可以这样使用：

```go
var users []User
db.Order("name").Find(&users)
```

上述代码将会按照 `name` 字段的升序排列用户记录。

如果要进行降序排列，可以在排序规则字符串中添加 `-` 符号：

```go
var users []User
db.Order("name DESC").Find(&users)
```

这样，查询结果将按照 `name` 字段的降序排列。

除了单个字段之外，您还可以通过多个字段进行排序，只需在排序规则字符串中添加多个字段名称即可：

```go
var users []User
db.Order("age DESC, name").Find(&users)
```

上述代码将首先按照 `age` 字段的降序排列，然后在相同 `age` 的记录中按照 `name` 字段的升序排列。

总之，`Order` 方法允许以指定的顺序对查询结果进行排序，使能够根据需要轻松地对数据进行排序。

##### 3.分页

根据进一步查询理解，分页查询不是将每一部分塞入不同结构体中，而是将整个表中的数据根据自定义的记录数分成一个个部分，再根据偏移量使其跳过已经显示的记录使其衔接至上一部分，然后将这一个个部分存储到切片中

###### 1.普通分页查询

可以使用 `Limit` 和 `Offset` 方法来实现分页查询。

```go
var users []User
pageSize := 10   // 每页记录数
pageNum := 1     // 当前页码

db.Model(&User{}).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users)
```

在这个示例中，假设每页显示10条记录，并且当前页码为1。`Limit` 方法用于指定每页的记录数，`Offset` 方法用于指定查询的偏移量，即要跳过的记录数。根据当前页码和每页记录数的关系，我们通过 `(pageNum - 1) * pageSize` 来计算偏移量。

此代码将从 `User` 表中查询指定页码的数据，并将结果存储在 `users` 切片中。

###### 2.基于游标的分页查询

可以使用 `Where` 方法结合游标来实现基于游标的分页查询。下面是一个示例代码：

```go
var users []User
pageSize := 10   // 每页记录数
cursor := 0      // 游标，初始值为0

db.Where("id > ?", cursor).Limit(pageSize).Find(&users)

// 更新游标
if len(users) > 0 {
    lastUser := users[len(users)-1]
    cursor = lastUser.ID
}
```

在这个示例中，我们假设每页显示10条记录，并且初始游标值为0。使用 `Where` 方法来指定查询条件，`id > ?` 表示只查询 `id` 大于当前游标值的记录。使用 `Limit` 方法来限制每页的记录数为10。

通过执行 `Find` 方法，我们可以获取到当前页的记录，并将结果存储在 `users` 切片中。

在更新游标时，我们检查是否有返回的记录，如果有，我们可以从最后一条记录中获取新的游标值，这样就可以定位到下一页的起始位置。

通过循环执行以上步骤，你可以实现基于游标的分页查询，避免了使用 `Offset` 导致的性能问题。

### 3.高级查询

#### 1.智能选择字段

在 GORM 中，您可以使用 [`Select`](https://gorm.io/zh_CN/docs/query.html) 方法有效地选择特定字段。 这在Model字段较多但只需要其中部分的时候尤其有用，比如编写API响应。

```go
type User struct {
  ID     uint
  Name   string
  Age    int
  Gender string
  // 很多很多字段
}

type APIUser struct {
  ID   uint
  Name string
}

// 在查询时，GORM 会自动选择 `id `, `name` 字段
db.Model(&User{}).Limit(10).Find(&APIUser{})
// SQL: SELECT `id`, `name` FROM `users` LIMIT 10
```

#### 2.锁

GORM 支持多种类型的锁，例如：

```go
// 基本的 FOR UPDATE 锁
db.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&users)
// SQL: SELECT * FROM `users` FOR UPDATE
```

上述语句将会在事务（transaction）中锁定选中行（selected rows）。 可以被用于以下场景：当你准备在事务（transaction）中更新（update）一些行（rows）时，并且想要在本事务完成前，阻止（prevent）其他的事务（other transactions）修改你准备更新的选中行。

`Strength` 也可以被设置为 `SHARE` ，这种锁只允许其他事务读取（read）被锁定的内容，而无法修改（update）或者删除（delete）。

```go
db.Clauses(clause.Locking{
  Strength: "SHARE",
  Table: clause.Table{Name: clause.CurrentTable},
}).Find(&users)
// SQL: SELECT * FROM `users` FOR SHARE OF `users`
```

`Table`选项用于指定将要被锁定的表。 这在你想要 join 多个表，并且锁定其一时非常有用。

你也可以提供如 `NOWAIT` 的Options，这将尝试获取一个锁，如果锁不可用，导致了获取失败，函数将会立即返回一个error。 当一个事务等待其他事务释放它们的锁时，此Options（Nowait）可以阻止这种行为

```go
db.Clauses(clause.Locking{
  Strength: "UPDATE",
  Options: "NOWAIT",
}).Find(&users)
// SQL: SELECT * FROM `users` FOR UPDATE NOWAIT
```

Options也可以是`SKIP LOCKED`，设置后将跳过所有已经被其他事务锁定的行（any rows that are already locked by other transactions.）。 这次高并发情况下非常有用：那时你可能会想要对未经其他事务锁定的行进行操作（process ）。

#### 3.子查询

子查询（Subquery）是SQL中非常强大的功能，它允许嵌套查询。 当你使用 *gorm.DB 对象作为参数时，GORM 可以自动生成子查询。

```go
// 简单的子查询
db.Where("amount > (?)", db.Table("orders").Select("AVG(amount)")).Find(&orders)
// SQL: SELECT * FROM "orders" WHERE amount > (SELECT AVG(amount) FROM "orders");

// 这行代码使用了一个简单的子查询，子查询的结果作为主查询的条件。在这个例子中，子查询部分是 db.Table("orders").Select("AVG(amount)")，它会返回 orders 表中 amount 列的平均值。主查询部分是 db.Where("amount > (?)", subQuery)，它会选择 orders 表中满足条件（amount 大于平均值）的记录。生成的 SQL 查询语句中会包含子查询。

//内嵌子查询
subQuery := db.Select("AVG(age)").Where("name LIKE ?", "name%").Table("users")
db.Select("AVG(age) as avgage").Group("name").Having("AVG(age) > (?)", subQuery).Find(&results)
// SQL: SELECT AVG(age) as avgage FROM `users` GROUP BY `name` HAVING AVG(age) > (SELECT AVG(age) FROM `users` WHERE name LIKE "name%")
```

### 6.更新

#### 1.普通选项

#####   1.保存所有字段

`Save` 会保存所有的字段，即使字段是零值

```go
db.First(&user)

user.Name = "jinzhu 2"
user.Age = 100
db.Save(&user)
// UPDATE users SET name='jinzhu 2', age=100, birthday='2016-01-01', updated_at = '2013-11-17 21:34:10' WHERE id=111;
```

`保存` 是一个组合函数。 如果保存值不包含主键，它将执行 `Create`，否则它将执行 `Update` (包含所有字段)。

```go
db.Save(&User{Name: "jinzhu", Age: 100})
// INSERT INTO `users` (`name`,`age`,`birthday`,`update_at`) VALUES ("jinzhu",100,"0000-00-00 00:00:00","0000-00-00 00:00:00")

db.Save(&User{ID: 1, Name: "jinzhu", Age: 100})
// UPDATE `users` SET `name`="jinzhu",`age`=100,`birthday`="0000-00-00 00:00:00",`update_at`="0000-00-00 00:00:00" WHERE `id` = 1
```

##### 2.更新单个列

当使用 `Update` 更新单列时，需要有一些条件，否则将会引起`ErrMissingWhereClause` 错误，查看 [阻止全局更新](https://gorm.io/zh_CN/docs/update.html#block_global_updates) 了解详情。 当使用 `Model` 方法，并且它有主键值时，主键将会被用于构建条件，例如：

```go
// 根据条件更新
db.Model(&User{}).Where("active = ?", true).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE active=true;

// User 的 ID 是 `111`
db.Model(&user).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111;

// 根据条件和 model 的值进行更新
db.Model(&user).Where("active = ?", true).Update("name", "hello")
// UPDATE users SET name='hello', updated_at='2013-11-17 21:34:10' WHERE id=111 AND active=true;
```

##### 3.更新多列

`Updates` 方法支持 `struct` 和 `map[string]interface{}` 参数。当使用 `struct` 更新时，默认情况下GORM 只会更新非零值的字段

```go
// 根据 `struct` 更新属性，只会更新非零值的字段
db.Model(&user).Updates(User{Name: "hello", Age: 18, Active: false})
// UPDATE users SET name='hello', age=18, updated_at = '2013-11-17 21:34:10' WHERE id = 111;

// 根据 `map` 更新属性
db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET name='hello', age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;
```

```go
data := map[string]interface{}{
    "name":  "Alice",
    "age":   25,
    "email": "alice@example.com",
}

if err := db.Table("users").Where("id = ?", id).Updates(data).Error; err != nil {
    // 处理更新失败的情况
}

//在这个例子中，我们直接使用 Table 方法指定要更新的数据库表名（假设为 "users"），而不是使用 Model 方法指定模型对象。

```

**注意：** 使用 struct 更新时, GORM 将只更新非零值字段。 可能想用 `map` 来更新属性，或者使用 `Select` 声明字段来更新,其他的更新操作逻辑与之前的示例相同。

##### 4.更新选定字段

如果您想要在更新时选择、忽略某些字段，您可以使用 `Select`、`Omit`

```go
// 选择 Map 的字段
// User 的 ID 是 `111`:
db.Model(&user).Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET name='hello' WHERE id=111;

db.Model(&user).Omit("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false})
// UPDATE users SET age=18, active=false, updated_at='2013-11-17 21:34:10' WHERE id=111;

// 选择 Struct 的字段（会选中零值的字段）
db.Model(&user).Select("Name", "Age").Updates(User{Name: "new_name", Age: 0})
// UPDATE users SET name='new_name', age=0 WHERE id=111;

// 选择所有字段（选择包括零值字段的所有字段）
db.Model(&user).Select("*").Updates(User{Name: "jinzhu", Role: "admin", Age: 0})

// 选择除 Role 外的所有字段（包括零值字段的所有字段）
db.Model(&user).Select("*").Omit("Role").Updates(User{Name: "jinzhu", Role: "admin", Age: 0})
```

##### 5.更新 Hook

GORM 支持的 hook 包括：`BeforeSave`, `BeforeUpdate`, `AfterSave`, `AfterUpdate`. 更新记录时将调用这些方法，查看 [Hooks](https://gorm.io/zh_CN/docs/hooks.html) 获取详细信息

```go
func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
    if u.Role == "admin" {
        return errors.New("admin user not allowed to update")
    }
    return
}
```

##### 6.批量更新

如果我们没有指定具有主键值的记录，则 GORM 将执行批量更新`Model`

```go
// Update with struct
db.Model(User{}).Where("role = ?", "admin").Updates(User{Name: "hello", Age: 18})
// UPDATE users SET name='hello', age=18 WHERE role = 'admin';

// Update with map
db.Table("users").Where("id IN ?", []int{10, 11}).Updates(map[string]interface{}{"name": "hello", "age": 18})
// UPDATE users SET name='hello', age=18 WHERE id IN (10, 11);
```

#### 2.高级选项

##### 1.使用 SQL 表达式更新

GORM 允许使用 SQL 表达式更新列，例如：

```go
// product's ID is `3`
db.Model(&product).Update("price", gorm.Expr("price * ? + ?", 2, 100))
// UPDATE "products" SET "price" = price * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;

db.Model(&product).Updates(map[string]interface{}{"price": gorm.Expr("price * ? + ?", 2, 100)})
// UPDATE "products" SET "price" = price * 2 + 100, "updated_at" = '2013-11-17 21:34:10' WHERE "id" = 3;

db.Model(&product).UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = 3;

db.Model(&product).Where("quantity > 1").UpdateColumn("quantity", gorm.Expr("quantity - ?", 1))
// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = 3 AND quantity > 1;
```

### 7.删除

#### 1.删除一条记录

删除一条记录时，删除对象需要指定主键，否则会触发 [批量删除](https://gorm.io/zh_CN/docs/delete.html#batch_delete)，例如：

```go
// Email 的 ID 是 `10`
db.Delete(&email)
// DELETE from emails where id = 10;

// 带额外条件的删除
db.Where("name = ?", "jinzhu").Delete(&email)
// DELETE from emails where id = 10 AND name = "jinzhu";
```

#### 2.根据主键删除

GORM 允许通过主键(可以是复合主键)和内联条件来删除对象，它可以使用数字（如以下例子。也可以使用字符串——译者注）。查看 [查询-内联条件（Query Inline Conditions）](https://gorm.io/zh_CN/docs/query.html#inline_conditions) 了解详情。

```go
db.Delete(&User{}, 10)
// DELETE FROM users WHERE id = 10;

db.Delete(&User{}, "10")
// DELETE FROM users WHERE id = 10;

db.Delete(&users, []int{1,2,3})
// DELETE FROM users WHERE id IN (1,2,3);
```

#### 3.钩子函数

对于删除操作，GORM 支持 `BeforeDelete`、`AfterDelete` Hook，在删除记录时会调用这些方法，查看 [Hook](https://gorm.io/zh_CN/docs/hooks.html) 获取详情

```go
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
    if u.Role == "admin" {
        return errors.New("admin user not allowed to delete")
    }
    return
}
```

#### 4.批量删除

如果指定的值不包括主属性，那么 GORM 会执行批量删除，它将删除所有匹配的记录

```go
db.Where("email LIKE ?", "%jinzhu%").Delete(&Email{})
// DELETE from emails where email LIKE "%jinzhu%";

db.Delete(&Email{}, "email LIKE ?", "%jinzhu%")
// DELETE from emails where email LIKE "%jinzhu%";
```

可以将一个主键切片传递给`Delete` 方法，以便更高效的删除数据量大的记录

```go
var users = []User{{ID: 1}, {ID: 2}, {ID: 3}}
db.Delete(&users)
// DELETE FROM users WHERE id IN (1,2,3);

db.Delete(&users, "name LIKE ?", "%jinzhu%")
// DELETE FROM users WHERE name LIKE "%jinzhu%" AND id IN (1,2,3); 
```

#### 5.返回删除行的数据

返回被删除的数据，仅当数据库支持回写功能时才能正常运行，如下例：

```go
// 回写所有的列
var users []User
DB.Clauses(clause.Returning{}).Where("role = ?", "admin").Delete(&users)
// DELETE FROM `users` WHERE role = "admin" RETURNING *
// users => []User{{ID: 1, Name: "jinzhu", Role: "admin", Salary: 100}, {ID: 2, Name: "jinzhu.2", Role: "admin", Salary: 1000}}

// 回写指定的列
DB.Clauses(clause.Returning{Columns: []clause.Column{{Name: "name"}, {Name: "salary"}}}).Where("role = ?", "admin").Delete(&users)
// DELETE FROM `users` WHERE role = "admin" RETURNING `name`, `salary`
// users => []User{{ID: 0, Name: "jinzhu", Role: "", Salary: 100}, {ID: 0, Name: "jinzhu.2", Role: "", Salary: 1000}}
```

#### 6.软删除

如果你的模型包含了 `gorm.DeletedAt`字段（该字段也被包含在`gorm.Model`中），那么该模型将会自动获得软删除的能力！

当调用`Delete`时，GORM并不会从数据库中删除该记录，而是将该记录的`DeleteAt`设置为当前时间，而后的一般查询方法将无法查找到此条记录。

```go
// user's ID is `111`
db.Delete(&user)
// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE id = 111;

// Batch Delete
db.Where("age = ?", 20).Delete(&User{})
// UPDATE users SET deleted_at="2013-10-29 10:23" WHERE age = 20;

// Soft deleted records will be ignored when querying
db.Where("age = 20").Find(&user)
// SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;
```

如果你并不想嵌套`gorm.Model`，你也可以像下方例子那样开启软删除特性：

```go
type User struct {
  ID      int
  Deleted gorm.DeletedAt
  Name    string
}
```

##### 查找被软删除的记录

你可以使用`Unscoped`来查询到被软删除的记录

```
db.Unscoped().Where("age = 20").Find(&users)
// SELECT * FROM users WHERE age = 20;
```

#### 7.永久删除（物理删除）

你可以使用 `Unscoped`来永久删除匹配的记录

```go
db.Unscoped().Delete(&order)
// DELETE FROM orders WHERE id=10;
```
