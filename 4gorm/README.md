# GORM MySQL 增删改查示例

这是一个使用GORM连接MySQL数据库进行增删改查操作的完整示例。

## 项目结构

```
4gorm/
├── go.mod              # Go模块文件
├── main.go             # 主程序文件（基础CRUD演示）
├── global/
│   └── db.go           # 全局数据库配置
├── models/
│   ├── user_model.go              # 用户模型定义（包含一对一关系）
│   └── department_employee_model.go # 部门员工模型定义（包含一对多关系）
├── 部门员工关系演示/
│   └── main.go         # 部门员工一对多关系演示程序
└── README.md           # 说明文档
```

## 环境要求

- Go 1.23.2 或更高版本
- MySQL 数据库
- 数据库名：`gorm_study`
- 用户名：`root`
- 密码：`mima123`

## 数据库配置

确保MySQL服务正在运行，并且已经创建了名为 `gorm_study` 的数据库：

```sql
CREATE DATABASE gorm_study;
```

## 运行步骤

### 基础CRUD演示
1. 进入项目目录：
```bash
cd 4gorm
```

2. 下载依赖：
```bash
go mod tidy
```

3. 运行基础CRUD演示：
```bash
go run main.go
```

### 一对一关系演示
1. 进入一对一关系演示目录：
```bash
cd 4gorm/一对一关系演示
```

2. 运行一对一关系演示：
```bash
go run main.go
```

### 部门员工一对多关系演示
1. 进入部门员工关系演示目录：
```bash
cd 4gorm/部门员工关系演示
```

2. 运行部门员工关系演示：
```bash
go run main.go
```

## 功能说明

### 基础CRUD演示 (main.go)
程序演示了以下GORM操作：

#### 1. 数据库连接
- 连接到本地MySQL数据库
- 自动迁移数据库表结构

#### 2. 创建操作 (Create)
- 创建新用户记录
- 自动生成ID和时间戳

#### 3. 查询操作 (Read)
- 查询所有用户
- 根据ID查询单个用户
- 根据条件查询用户（如按姓名查询）

#### 4. 更新操作 (Update)
- 更新用户信息
- 自动更新时间戳

#### 5. 删除操作 (Delete)
- 软删除（逻辑删除）
- 永久删除（物理删除）
- 查询软删除的记录

### 一对一关系演示 (一对一关系演示/main.go)
专门演示GORM的一对一关联操作：

#### 1. 关联模型创建
- 创建用户和用户详情记录
- 建立一对一关联关系

#### 2. 预加载查询
- 使用Preload查询用户及其详情
- 从用户详情反向查询用户信息

#### 3. 关联数据操作
- 更新关联的详情信息
- 删除关联的详情记录
- 查询所有用户及其关联的详情

### 部门员工一对多关系演示 (部门员工关系演示/main.go)
专门演示GORM的一对多关联操作：

#### 1. 部门模型 (DepartmentModel)
- 部门ID、名称、编码、经理、位置
- 一对多关联到员工列表

#### 2. 员工模型 (EmployeeModel)
- 员工ID、姓名、工号、职位、薪资、邮箱、电话
- 多对一关联到部门

#### 3. 一对多关联操作
- 创建部门和员工记录
- 使用Preload查询部门及其员工
- 从员工查询所属部门信息
- 更新员工信息和薪资
- 统计各部门员工数量

## 模型结构

### 用户模型 (一对一关系)
```go
type UserModel struct {
    ID              uint             `gorm:"primaryKey"`
    Name            string           `gorm:"size:100;not null"`
    Age             int              `gorm:"not null;default:18"`
    UserDetailModel *UserDetailModel `gorm:"foreignKey:UserID"`
    CreatedAt       time.Time
    UpdatedAt       time.Time
    DeletedAt       gorm.DeletedAt `gorm:"index"`
}

type UserDetailModel struct {
    ID        int
    UserID    int       `gorm:"unique"`
    Email     string    `gorm:"size:32"`
    UserModel UserModel `gorm:"foreignKey:UserID"`
}
```

### 部门员工模型 (一对多关系)
```go
type DepartmentModel struct {
    ID           uint             `gorm:"primaryKey"`
    Name         string           `gorm:"size:50;not null"`
    Code         string           `gorm:"size:20;unique;not null"`
    Manager      string           `gorm:"size:30"`
    Location     string           `gorm:"size:100"`
    EmployeeList []EmployeeModel  `gorm:"foreignKey:DepartmentID"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
    DeletedAt    gorm.DeletedAt  `gorm:"index"`
}

type EmployeeModel struct {
    ID           uint             `gorm:"primaryKey"`
    Name         string           `gorm:"size:30;not null"`
    EmployeeNo   string           `gorm:"size:20;unique;not null"`
    Position     string           `gorm:"size:30"`
    Salary       float64          `gorm:"type:decimal(10,2)"`
    Email        string           `gorm:"size:50"`
    Phone        string           `gorm:"size:20"`
    DepartmentID uint             `gorm:"not null"`
    Department   DepartmentModel  `gorm:"foreignKey:DepartmentID"`
    CreatedAt    time.Time
    UpdatedAt    time.Time
    DeletedAt    gorm.DeletedAt  `gorm:"index"`
}
```

## 主要功能函数

### 数据库配置 (global/db.go)
- `InitDB()` - 初始化数据库连接
- `MigrateDB()` - 自动迁移数据库表
- `GetDB()` - 获取数据库实例

### 业务功能 (main.go)
- `createUser()` - 创建用户
- `getAllUsers()` - 查询所有用户
- `getUserByID()` - 根据ID查询用户
- `getUserByName()` - 根据姓名查询用户
- `updateUser()` - 更新用户信息
- `deleteUser()` - 软删除用户
- `permanentlyDeleteUser()` - 永久删除用户
- `createUserDetail()` - 创建用户详情
- `getUserWithDetail()` - 查询用户及其详情

## 注意事项

1. 确保MySQL服务正在运行
2. 确保数据库 `gorm_study` 已创建
3. 确保用户名和密码正确
4. 程序会自动创建 `user_models` 表
5. 支持软删除功能，删除的记录不会真正从数据库中移除
