## 使用go-zero实现的oj后端

PGOJ:backend
├─common（各个微服务共享的基础组件）
│  ├─dataType（为const数据起别名，避免硬编码）
│  ├─interceptors（grpc状态码与错误码的转换）【②】
│  ├─jwt（定义并生成 JWT ）
│  └─xcode（自定义业务错误码）【①】
│      └─types
├─doc（api文档）
├─docker_files（使用docker部署中间件）
│  ├─docker-mysql（mysql主从部署）
│  │  ├─master-conf
│  │  ├─mount_data
│  │  └─slave-conf
│  ├─docker-nats（消息中间件nats）
│  └─docker-nginx（反向代理nginx）
│      └─templates（网页静态文件）
├─judgeStatus（判题模块）
│  ├─cmd
│  │  ├─api
│  │  │  ├─desc
│  │  │  ├─etc
│  │  │  ├─internal
│  │  │  │  ├─config
│  │  │  │  ├─handler
│  │  │  │  │  └─judgeStatus
│  │  │  │  ├─logic
│  │  │  │  │  └─judgeStatus
│  │  │  │  ├─svc
│  │  │  │  └─types
│  │  │  └─logs
│  │  └─rpc
│  │      ├─etc
│  │      ├─internal
│  │      │  ├─code
│  │      │  ├─config
│  │      │  ├─logic
│  │      │  ├─server
│  │      │  └─svc
│  │      ├─judgeservice
│  │      ├─logs
│  │      └─pb
│  ├─docker-judge（判题模块docker部署）
│  │  └─logs
│  │      ├─api
│  │      └─rpc
│  └─model
├─k8s_files（使用k8s部署中间件和服务）
│  ├─ingress
│  ├─judge
│  ├─minio
│  ├─mysql
│  ├─nats
│  ├─other
│  ├─problems
│  └─users
├─other（文件模块）
│  ├─docker-other（文件模块docker部署）
│  │  └─logs
│  ├─etc
│  ├─internal
│  │  ├─code
│  │  ├─config
│  │  ├─handler
│  │  │  └─upload_file
│  │  ├─logic
│  │  │  └─upload_file【④】
│  │  ├─svc
│  │  │  └─pkg
│  │  └─types
│  └─logs
├─problems（题目模块）
│  ├─cmd
│  │  ├─api
│  │  │  ├─desc
│  │  │  │  ├─problem
│  │  │  │  ├─tag
│  │  │  │  └─test_case
│  │  │  ├─etc
│  │  │  ├─internal
│  │  │  │  ├─code
│  │  │  │  ├─config
│  │  │  │  ├─handler
│  │  │  │  │  ├─problems
│  │  │  │  │  ├─problem_data
│  │  │  │  │  ├─tags
│  │  │  │  │  └─test_case
│  │  │  │  ├─logic
│  │  │  │  │  ├─problems
│  │  │  │  │  ├─problem_data
│  │  │  │  │  ├─tags
│  │  │  │  │  └─test_case
│  │  │  │  ├─svc
│  │  │  │  └─types
│  │  │  └─logs
│  │  └─rpc
│  │      ├─etc
│  │      ├─internal
│  │      │  ├─code
│  │      │  ├─config
│  │      │  ├─logic
│  │      │  ├─server
│  │      │  └─svc
│  │      ├─logs
│  │      ├─pb
│  │      └─problemservice
│  ├─docker-problems（题目模块docker部署）
│  │  └─logs
│  │      ├─api
│  │      └─rpc
│  └─model
└─users（用户模块）【③】
    ├─cmd
    │  ├─api
    │  │  ├─desc
    │  │  │  ├─user_info
    │  │  │  └─user_profile
    │  │  ├─etc
    │  │  ├─internal
    │  │  │  ├─code
    │  │  │  ├─config
    │  │  │  ├─handler
    │  │  │  │  ├─other
    │  │  │  │  ├─users_info
    │  │  │  │  └─users_profile
    │  │  │  ├─logic
    │  │  │  │  ├─other
    │  │  │  │  ├─users_info
    │  │  │  │  └─users_profile
    │  │  │  ├─svc
    │  │  │  │  └─email
    │  │  │  └─types
    │  │  └─logs
    │  └─rpc
    │      ├─etc
    │      ├─internal
    │      │  ├─code
    │      │  ├─config
    │      │  ├─logic
    │      │  ├─server
    │      │  └─svc
    │      ├─logs
    │      ├─pb
    │      └─userservice
    ├─docker-users（用户模块部署）
    │  └─logs
    │      ├─api
    │      └─rpc
    └─model



##### 【①】接口：error，xcode，结构体：code。为了实现client从error中提取 gRPC status，定义了types.status，通过*组合* 得到的xcode.status实现了xcode接口。

1. xcode.go：定义实现了Error的Xcode接口和Code结构体。
2. status.pb.go：types.status用于grpc的状态码通信。
3. status.go：将code或error先转化为xcode.status，以便grpc传输。
4. common.go：定义通用的项目错误码。
5. response.go：api服务用于溯源error，得到xcode，返回业务错误码和错误信息

##### 【②】业务error message的转换与传输，手动兼容error、xcode，不产生新error。

1. ServerErrorInterceptor.go：将error、xcode（xcode.status，code）转化为grpcStatus（存入**details**字段）并且返回。
2. ClientErrorInterceptor.go：将从Server收到的error转化为grpcStatus，然后再（根据grpcStatus的**details**字段）转化为**xcode接口**，转化成error并返回。

##### 【③】以users模块为例，介绍此项目的层次结构以及微服务开发流程。

1. 每个微服务都采用model——>rpc——>api三层架构：
   1. 每个model对应一张表，负责向rpc提供**操作单个table**的方法。
   2. rpc层负责根据业务，联合**操作同一模块的多张表**，比如多对多关系：problem，tag，problem_tag，会向api层隐藏problem_tag的操作。除此之外，rpc还负责中间件（如redis，nats，minio）的调用。
   3. api纯粹负责业务逻辑，可以**操作不同模块的数据表**，比如根据用户名找到其创建的题目，先向users模块查找user_id，再根据user_id向problems模块查找记录。除此之外，api还负责jwt用户鉴权。
2. model层：
   1. 通过命令：goctl model mysql datasource -c -d . --url "root:root@tcp(127.0.0.1:3306)/oj_micro" --table user_login
      生成自带缓存的**基础crud操作**。
   2. 确认**userloginmodel_gen.go**内容：定义了与table对应的*userLoginModel*接口，以及实现了此接口的*defaultUserLoginModel*结构体。
   3. 编辑userloginmodel.go：
      1. 对于UserLoginModel接口，它**嵌入**了userloginmodel_gen.go中的***u**serLoginModel*接口，可以为***U**serLoginModel*接口添加自定义方法签名。
      2. 对于customUserLoginModel结构体，它**继承**了userloginmodel_gen.go中的*defaultUserLoginModel*结构体，*customUserLoginModel*要**实现** *UserLoginModel*接口。
      3. 本项目采用**squirrel**第三方库，实现自定义的数据库查找操作。
3. rpc层：
   1. 编写**protobuf**文件，通过命令：goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=. ，生成rpc服务提供自定义方法调用。
   2. 在etc配置文件、config结构体里面，添加database、cache、etcd等设置。在**servicecontext.go**中，为*ServiceContext*结构体添加*UserLoginModel*等字段，然后修改NewServiceContext函数使用config进行资源实例化（连接）
   3. 在internal包内新建code.go，声明了此微服务相关的xcode（**业务错误码**）全局变量。在启动入口user.go的main函数中，添加【②】中自定义的**grpc拦截器** *ServerErrorInterceptor*。
   4. logic方法：主要使用三个结构体：logic结构体，req结构体，resp结构体。
      1. logic结构体由ctx，svcCtx，logger组成。**svcCtx**是在main.go启动的时候就创建的中间件资源连接。
      2. req结构体和logic结构体的**ctx**都是api层调用的时候传入的参数。logic结构体的**logger**是根据这个ctx创建的日志追踪器。
      3. logic方法通过resp结构体（或err）**响应**处理结果（或错误码）。
   5. 解释**userserviceserver.go**和**userservice.go**和**user.go**的作用：
      1. rpc层实现的logic方法，被**UserServiceServer**结构体的方法调用。**UserServiceServer**在进程启动时以单例模式被注册到pb服务端。
      2. api层调用的rpc方法，在**userservice.go**中被声明。**UserService**结构体的方法会调用pb客户端，向pb服务端发起请求。
      3. *user.go*的**main**函数读取配置文件，然后以单例模式初始化中间件（如mysql，redis，etcd）的连接，然后启动grpc服务。
4. api层：
   1. 编写**.api文件**，通过命令：goctl api go --api desc/user.api --dir .
      生成api服务提供api访问。
   2. 在etc配置文件、config结构体里面，添加需要的rpc服务，鉴权信息等设置。在**servicecontext.go**，为*ServiceContext*添加*userservice.UserService*等字段，然后修改NewServiceContext函数使用config进行实例化（连接）。在初始化rpc客户端时，会添加【②】中自定义的**grpc拦截器** *ClientErrorInterceptor*。
   3. 在启动入口users.go的main中，向http服务注册了【①】中定义的**xcode.ErrHanler**错误处理函数。为了实现需要鉴权的api访问，还向http服务注册了**WithUnauthorizedCallback**函数，会解析request的header是否携带authentication字段，**携带则将jwt的json数据存入ctx键值对中**，否则响应鉴权错误码。如果handler对Requst参数的解析失败，则返回错误码：**xcode.InvalidParams.WithDetails(err.Error())**
   4. logic方法：主要使用三个结构体：logic结构体，req结构体，resp结构体。
      1. logic结构体由ctx，svcCtx，logger组成。**svcCtx**是在main.go启动的时候就创建的rpc连接。
      2. req结构体和logic结构体的ctx都是handler函数解析**http.Request**得到的变量。logic结构体的**logger**是根据这个ctx创建的日志追踪器。
      3. logic方法通过resp结构体（或err）**响应**处理结果（或错误码）。
      4. 需要**鉴权**的api操作，会**从ctx中读取json键值对**，如jwt过期时间，用户等级，用户ID。
   5. 解释types.go和routes.go和users.go的作用：
      1. **types.go**定义了所有logic方法的请求体和响应体。
      2. **routes.go**的RegisterHandlers函数将所有Handler方法绑定到对应的*URI*（Method+Path）
      3. **users.go**的main函数会先根据RestConf初始化api server，然后调用RegisterHandlers函数，将所有*URI*注册到server并启动server提供*URL*访问。

##### 【④】使用minio实现文件上传和图片访问。

1. 

