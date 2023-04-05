a todolist demo

## 1.业务需求

系统包含2种角色的用户：一般用户、管理员，不同角色的权限不一样。

一条待办事项至少包含如下属性（标题、正文、创建者、创建时间、修改时间）。

对于一般用户，只能操作自己创建的待办事项；对于管理者，能够操作所有人创建的待办事项。

### 一般用户
以下需求3～6要求用户必须登陆系统后才能操作。
1. 通过用户名+密码登陆系统
2. 登出系统
3. 新增一条（属于自己的）待办事项
4. 修改一条（属于自己的）待办事项
5. 删除一条（属于自己的）待办事项
6. 罗列所有（属于自己的）待办事项列表

### 管理员
以下所有需求都必须登陆系统。管理员除了具有上面一般用户的权限，还可以：
1. 新增一名一般用户
2. 删除一名一般用户
3. 修改一名一般用户的密码
4. 修改一条（属于任何人的）待办事项
5. 删除一条（属于任何人的）待办事项
6. 罗列所有（属于任何人的）待办事项列表

## 2. 技术需求

1. 数据不要求持久化存储到DB，可以保存在内存中，重启后丢失是可接受的
2. 管理员数据可以硬编码在代码中
3. 提供Swagger文档







