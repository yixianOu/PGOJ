## 使用go-zero实现的oj后端

#### 目录结构：

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
│  │  │  │  │  └─judgeStatus【⑥】  
│  │  │  │  ├─svc  
│  │  │  │  └─types  
│  │  │  └─logs  
│  │  └─rpc  
│  │      ├─etc  
│  │      ├─internal  
│  │      │  ├─code  
│  │      │  ├─config  
│  │      │  ├─logic【⑤】  
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
├─k8s_files（使用k8s部署中间件和服务）【⑧】  
│  ├─ingress  
│  ├─judge  
│  ├─minio  
│  ├─mysql【⑨】  
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



##### 【①】接口：error，xcode，结构体：code。为了实现client端从调用rpc函数返回的error中提取 gRPC status，定义了types.status结构体，通过*组合*（嵌入命名结构体）的方式 得到的xcode.status结构体实现了xcode接口。

1. xcode.go：定义实现了Error接口的Xcode接口和实现了Error接口Code结构体。
2. status.pb.go：types.status结构体用于grpc状态码的通信。
3. status.go：将code或error先转化为xcode.status，以便grpc传输。
4. common.go：定义通用的项目错误码。
5. response.go：定义了错误处理函数，被api服务用于溯源error，得到xcode，返回业务错误码和错误信息。

##### 【②】业务error message的转换与传输，手动兼容error、xcode，不产生新error。

1. ServerErrorInterceptor.go：将实现了error、xcode的结构体（xcode.status，code）转化为grpcStatus结构体（存入**details**字段）并且响应给客户端。
2. ClientErrorInterceptor.go：将从Server收到的error转化为grpcStatus结构体，然后再（根据grpcStatus的**details**字段）转化为**xcode接口**，转化成error并返回。

##### 【③】以users模块为例，介绍此项目的层次结构以及微服务开发流程。

1. 每个微服务都采用model——>rpc——>api三层架构：
   1. 每个model对应一张表，负责向rpc提供**操作单个table**的方法。
   2. rpc层负责根据业务，联合**操作同一模块的多张表**，比如多对多关系：problem，tag，problem_tag，会向api层隐藏problem_tag的存在。除此之外，rpc还负责中间件（如redis，nats，minio）的调用。
   3. api纯粹负责业务逻辑，可以**操作不同模块的数据表**，比如根据用户名找到其创建的题目，先向users模块查找用户名对应的user_id，再根据user_id向problems模块查找记录。除此之外，api还负责**jwt用户鉴权**。
2. model层：
   1. 通过命令：goctl model mysql datasource -c -d . --url "root:root@tcp(127.0.0.1:3306)/oj_micro" --table user_login
      生成自带缓存的**基础crud操作**。
   2. 确认**userloginmodel_gen.go**内容：定义了与table对应的*userLoginModel*接口，以及实现了此接口的*defaultUserLoginModel*结构体。
   3. 编辑userloginmodel.go：
      1. 对于UserLoginModel接口，它**嵌入**了userloginmodel_gen.go定义的***u**serLoginModel*接口，因此继承了后者的方法签名。也可以为***U**serLoginModel*接口添加自定义方法签名。
      2. 对于customUserLoginModel结构体，它**继承**（嵌入匿名结构体）了userloginmodel_gen.go定义的*defaultUserLoginModel*结构体，因此也实现了**u**serLoginModel的方法签名。*customUserLoginModel*要**实现** *UserLoginModel*接口。
      3. 本项目采用**squirrel**第三方库，实现自定义的数据库查找操作。
3. rpc层：
   1. 编写**protobuf**文件，通过命令：goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=. ，生成rpc服务提供的自定义方法调用。
   2. 在etc配置文件、config结构体里面，添加database、cache、etcd等资源设置。  
      在**servicecontext.go**中，为*ServiceContext*结构体添加*UserLoginModel*等字段，然后修改NewServiceContext函数使用config进行（数据库、缓存等）资源的实例化（连接）。
   3. 在internal包内新建code.go，声明了此微服务相关的xcode（**业务错误码**）全局变量。  
      在启动入口user.go的main函数中，添加【②】中自定义的**grpc拦截器** *ServerErrorInterceptor*。
   4. logic方法：主要使用三个结构体：logic结构体，req结构体，resp结构体。
      1. logic结构体由ctx，svcCtx，logger组成。**svcCtx**是在main.go启动的时候就创建的（数据库、缓存等）资源实例连接。
      2. req结构体和logic结构体的*ctx*都是**api层调用**的时候传入的参数。logic结构体的**logger**是根据这个*ctx*创建的日志追踪器。
      3. logic方法通过resp结构体（或err）**响应**处理结果（或错误码）。
   5. 解释**userserviceserver.go**和**userservice.go**和**user.go**的作用：
      1. rpc层实现的logic方法，被**UserServiceServer**结构体的方法调用。**UserServiceServer**在进程启动时以单例模式被注册到pb服务端。
      2. api层调用的rpc方法，在**userservice.go**中被声明。**UserService**结构体的方法会调用pb客户端，向pb服务端发起请求，pb服务端根据请求调用UserServiceServer结构体的方法并响应。
      3. *user.go*的**main**函数读取配置文件，然后以单例模式**初始化**中间件资源（如mysql，redis，etcd等）的连接，然后**启动**grpc服务。
4. api层：
   1. 编写**api文件**，通过命令：goctl api go --api desc/user.api --dir .
      生成api服务提供api访问。
   2. 在etc配置文件、config结构体里，添加所需的rpc服务、鉴权信息等设置。  
      在**servicecontext.go**，为*ServiceContext*添加*userservice.UserService*等字段（用于调用服务），然后修改NewServiceContext函数使用config进行实例化（连接）。  
      在初始化rpc客户端时，会添加【②】中自定义的**grpc拦截器** *ClientErrorInterceptor*。
   3. 在启动入口users.go的**main函数**中，向http服务注册了【①】中定义的**xcode.ErrHanler**错误处理函数。  
      为了实现需要鉴权的api访问，还向http服务注册了**WithUnauthorizedCallback**函数，它会检查（对需要鉴权的api的）请求的header是否携带authentication字段，**携带则将jwt的json数据存入ctx键值对中**，不携带则响应鉴权失败错误码。  
      如果handler对Requst的参数校验失败，（为了触发错误处理函数**xcode.ErrHanler**），则向ResponseWriter写入**错误码**xcode.InvalidParams.WithDetails(err.Error())而不是err。
   4. logic方法：主要使用三个结构体：logic结构体，req结构体，resp结构体。
      1. logic结构体由ctx，svcCtx，logger组成。**svcCtx**是在main.go启动的时候就创建的资源连接（如rpc服务等）
      2. req结构体和logic结构体的*ctx*都是handler函数解析**http.Request**得到的变量。logic结构体的**logger**是根据这个*ctx*创建的日志追踪器。
      3. logic方法通过resp结构体（或err）**响应**处理结果（或错误码）。
      4. 需要**鉴权**的api操作，会**从ctx中读取json键值对**，如jwt过期时间，用户等级，用户ID。
   5. 解释types.go和routes.go和users.go的作用：
      1. **types.go**定义了所有logic方法的**请求体和响应体**。
      2. **routes.go**的RegisterHandlers函数将所有**Handler方法**绑定到对应的*URI*（Method+Path）。
         *handler*函数会使用Request的**ctx**和router传入的**svcCtx**，创建api对应的结构体实例并调用其**唯一的方法**：*logic*。
      3. **users.go**的main函数会先根据RestConf**初始化**api server，然后调用RegisterHandlers函数，将所有*URI* **注册**到server并**启动**server提供*URL*访问。

##### 【④】other模块的api使用minio实现文件上传和图片访问。

1. handler函数使用**r.FormFile**方法从Request中取出file和fileHeader，使用**context.WithValue**方法将file和fileHeader存入Request的**context**中，然后调用对应的logic方法。
2. 添加测试样例文件：
   1. 检查解析jwt得到的字段（是否过期，用户权限）："access_expire"，"role_level"
   2. **检查**数据库：是否已存在testcase记录。如果存在则：要删除旧记录及其文件之后，才能上传新文件。（文件改了但是testgroup没改的情况）
   3. 从ctx中取出file和fileHeader，与minio.Client，FileType等信息一同传入PutFileToMinio函数。
   4. 在PutFileToMinio函数中，检查文件的大小和后缀名，然后根据业务生成minio**文件路径**："\$problemID"+"/"+"inputfile/outputfile"+"\$filename"（路径不包含testgroup导致需要两次检查）。  
      使用**StatObject**方法**检查文件路径**，防止相同文件名导致的**覆盖**（文件没改但是testGroup改了的情况），文件不存在则使用**PutObject**方法上传此文件。
   5. 向数据库插入此testcase记录。
   6. 修改后：minio文件名不依赖\$filename了，而是problemID+testGroup，因为路径加上了testGroup，所以minio的文件路径是全局唯一的，不需要检查覆盖和冲突的情况了。
3. 更新用户头像：
   1. 检查解析jwt得到的字段（是否过期，用户ID）："access_expire"，"user_id"
   2. 删除用户旧头像：先用**ListObjects**方法根据**Prefix**获得ObjectInfo，然后用**RemoveObjects**方法根据ObjectInfo删除文件。（**ObjectInfo可以为空**，兼容了用户初次上传头像的情况）
   3. 从ctx中取出file和fileHeader，与minio.Client，FileType等信息一同传入PutFileToMinio函数。
   4. 在PutFileToMinio函数中，检查文件的大小和后缀名，然后根据业务生成minio**文件路径**："user_cover"+"/"+"\$user_id"+"\$GetFileExt(fileName)"   
      使用**PutObject**方法上传此图片，并且通过**minio.PutObjectOptions**设置参数ContentDisposition: "inline"和ContentType："image/$GetFileExt(fileName)"，以便访问图片时预览而非下载。
   5. 拼接（DomainName+BucketName+filePath）图片的**URL**并更新此用户的数据库记录。
4. 删除测试样例文件：
   1. 检查解析jwt得到的字段（是否过期，用户权限）："access_expire"，"role_level"
   2. 根据test_id查询数据库得到**文件路径**：InputFileName和OutputFileName。
   3. 使用**RemoveObject**方法删除FileName对应的minio文件。
   4. 从数据库删除此testcase记录。

##### 【5】通过消息队列发送判题

1. rpc层的addjudgestatuslogic.go负责使用**nats**向判题机**发送判题任务**并且**监听**判题任务的**多个处理结果**。

2. 首先约定好判题任务和判题结果的json结构体：judgeTest，testCaseResult。

3. rpc方法首先根据这个判题请求向数据库**insert**判题记录并且得到**"$judgeId"**。然后根据判题请求创建判题任务judgeTest，并通过nats的**JetStream.PublishAsync**方法将这个judgeTest发到nats服务器的特定Subject："ToJudger"，返回**PubAck**。  
   提前通过**js.CreateOrUpdateStream**方法，在nats服务器创建的stream（"judge_status"），会将任务放入一个消息队列里面，这个消息队列可以保证消息的**exactly once**消费。

4. 判题机通过创建属于这个（名为"judge_status"的）stream的**consumer**，能通过**pull**模式向nats服务器的消息队列（存着发到"ToJudger"主题的消息），拉取judgeTest任务。拉取动作是**阻塞**的，可自定义**timeout**不过期以避免循环pull

5. 介绍**NatsClient.Subscribe**方法，此方法接收两个参数：Subject和MsgHandler，返回一个subscription。  
   为了区分每次判题，我们将判题结果发到不同的主题，项目使用了 **"$judgeId"**作为Subject。  
   subscription（pub/sub消息传递机制）可以**控制**对当前主题消息的读取操作。在此项目用以接收从nats服务器**push**到订阅者的消息，并使用MsgHandler处理最新的消息。判题结束后使用**Unsubscribe**方法取消对当前主题的订阅。

6. **NatsClient.Subscribe**方法不是阻塞的，其工作机制是：发起一个**goroutine**，这个goroutine会执行**for循环**，循环的流程是：每当subscription**阻塞监听**到一个消息，就会调用一次（函数类型的入参）**MsgHandler**，并且将这个消息作为**nats.Msg**的实参**传入**MsgHandler进行处理。

7. MsgHandler的实现：使用了三个**外部变量**处理nats.Msg：waitgroup，无缓冲errChannel，pb.stream。  

   1. 由于业务场景是：判题机拉取判题任务之后，对**每个testcase**都会生成一个判题结果并通过nats发布消息。**所以**后端需要**监听**与testcase相等数量的消息，然后通过rpc的stream将消息**流式响应**到api层处理业务逻辑。处理期间如果有**Error**，则通过channel通知**主goroutine**，然后将waitgroup的**计数器置零**，返回业务错误码。
   2. 具体流程是：先在主goroutine初始化errChannel和waitgroup，指定监听个数**wg.Add(int(in.CaseNum))**。在MsgHandler中，先对接收到的msg的data进行json解码，然后将data通过pb的**stream.Send**方法发到api层，最后两个**计数器自减**：wg.Done()，atomic.AddInt64(&in.CaseNum, -1) 。期间出现的任何Error都会传入errChannel。

8. 发起订阅subject并处理nats消息的goroutine之后，还要发起一个goroutine用于**监听错误并进行善后处理**。这个goroutine接收两个入参：errChannel和codeChannel，内部**只有一个select语句**。这个select有3个case：  

   ```go
   case err = <-PubAck.Err():
   case <-l.ctx.Done():
   case err = <-errChannel:
   ```

   分别监听三种错误：**消息接收失败，超时，内部错误**。  
   每个case的错误处理逻辑相同：将两个**计数器置零**，然后将**业务错误码**传入codeChannel。

9. 为什么除了waitgroup内置计数器，还需要in.CaseNum计数器？为什么in.CaseNum通过原子操作自减？  
   waitgroup**不能读取**内置计数器的值，所以in.CaseNum与waitgroup计数器**同步自减**，以便出现Error时，执行**in.CaseNum次数**的wg.Done()，将waitgroup计算器置零，然后**唤醒**被wg.Wait()阻塞的主goroutine。  
   因为发起的两个goroutine会对in.CaseNum**并发读写**，所以使用**原子操作**避免**race condition**。

10. 被wg.Wait()阻塞的主goroutine在被唤醒后，通过**select**读取codeChannel并**返回业务错误码**，如果没有error（对codeChannel的读取会阻塞），则default**返回nil**。

11. 【⑥】判题请求的api层会在**for循环**中，通过**stream.Recv()**接收rpc层发来的多个判题结果（接收到**io.EOF时break**），然后根据业务要求，将多个判题结果**整合**成一个判题数据，**更新**rpc插入数据库的判题记录，并响应前端。

##### 【⑧】本项目使用k3s部署到云服务器，使用jenkins实现CICD。项目目录已包含jenkinsfile和kubectl配置文件

##### 【⑨】MySQL主从部署

1. 项目的业务要求，判题机**从minio拉取**testcase样例文件，保存到服务器**本地**。当pull到一个**判题任务**时，判题机应判断本地**是否存在**此题目对应的样例文件，以及这些样例文件**是否需要更新**。
2. 解决方案是：后端负责读写testcase表，判题机负责读testcase表和**同步本地old_case表**。针对**每一个判题任务**，判题机通过**比对**两张表的**时间戳**记录，判断是否需要从minio拉取样例文件，如果拉取了新的样例文件到本地，就**更新本地old_case表**，与testcase表保持一致。
3. 在分布式的情况下，需要通过**主从复制**实现不同服务器的MySQL的**testcase表同步**，而old_case表属于另一个**不同步的schema**，这样就实现了不同的服务器由本地的判题机来维护本地样例对应的old_case表。
