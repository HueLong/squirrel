# model层

## 介绍

- model是用来定义数据库的关联实体，如抽奖对应lottery_setting的数据表
- 本文件夹内只允许定义数据库关联的结构体、重写gorm的相关函数如TableName，数据查询方法，scopes等，绝不允许书写处理数据逻辑和缓存相关

## 文件内书写规范

由于golang的特性，一个文件夹内，属于同一个package，所以的方法都是开放、直接调用的，如model 下的user.go,lottery.go
里面的方法都是通过model.Func调用的，如果函数过多，就会混乱，不方便调用，所以做出如下规范

- 定义一个结构体，结构体名称为实体名+Model,如UserModel
- 每个实体下的查询方法必须实现该实体，如 func (UserModel) GetUserInfo()

同一个文件内可以定义结构体，书写scope，调用查询方法，现强行规范如下

- 首先定义结构体
- 如需要，则重写TableName方法
- 写Scopes
- 写查询方法

如何过滤不必要的数据