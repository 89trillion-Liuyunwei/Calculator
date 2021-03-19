### 技术选型
框架: gin

日志处理: gin


### 接口风格

 RESTFUL

###  项目结构
* api - handler函数
* route - 路由注册
* model - 数据计算
* log - 日志处理
* vendor - 项目依赖其他开源项目目录
* main.go - 程序执行入口



### 计算接口



| 	参数名称| 参数类型  |   备注
|-------| --------| --------|
| url |	localhost：8080/calculation|访问地址
| data |String	|  计算数据
| method|POST	|  访问类型

###请求样例
http://localhost:8080/calculation
body{

data: 2 +4/2

}

###Response
json数据

| 	参数名称| 参数类型  |  备注 
|-------| --------|--------|
|code |	int|状态吗
| data |int	| 计算结果  
| message|String	|  提示信息

###返回样例

{
code:0,
data:12,
message:"success"}

{
code:1,
data:0,
message:"请不要提交空数据"}

{
code:1,
data:0,
message:"表达式其中除数为0"}

{
code:1,
data:0,
message:"预计算的字符串格式错误"}