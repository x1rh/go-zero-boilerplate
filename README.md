## 安装依赖
推荐使用go的tools

## 一般流程
- 在app文件夹下创建对应业务的文件夹，例如user-service，这是一个rpc服务。
- 编写proto文件，一个rpc服务推荐仅一个proto文件，使用插件进行注释；使用makefile生成代码。
- 使用`app/makefile`生成apigateway依赖的所有protosets 
- 启动所有的rpc后，再启动apigateway
- 目前把所有的数据库model放在全局进行依赖，在日常开发中有少数model会被多个服务依赖

## todo
- 合并swagger
- 检查合并的proto文件是否有问题
- 检查认证逻辑
- 检查header和metadata传递