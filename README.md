## 使用go-zero实现的oj后端

PGOJ:backend
│  .gitignore
│  go.mod
│  go.sum
│  README.md
│
├─common
│  ├─dataType
│  │      type.go
│  │
│  ├─interceptors
│  │      clienterrorinterceptor.go
│  │      servererrorinterceptor.go
│  │
│  ├─jwt
│  │      token.go
│  │
│  └─xcode
│      │  common.xcode.go
│      │  response.go
│      │  status.go
│      │  xcode.go
│      │
│      └─types
│              status.pb.go
│              status.proto
│
├─doc
│
├─docker_files
│  │  docker-compose-etcd-local.yml
│  │  docker-compose-etcd.yml
│  │  docker-compose-minio.yml
│  │  docker-compose-redis-local.yml
│  │  docker-compose-redis.yml
│  │
│  ├─docker-mysql
│  │  │  docker-compose-mysql.yml
│  │  │
│  │  ├─master-conf
│  │  │      my.cnf
│  │  │
│  │  ├─mount_data
│  │  └─slave-conf
│  │          my.cnf
│  │
│  └─docker-nginx
│      │  docker-compose-nginx.yml
│      │  docker-compose.yml
│      │  local_nginx.conf
│      │  nginx.conf
│      │
│      └─templates
├─judgeStatus
│  │  .gitignore
│  │  Jenkinsfile
│  │
│  ├─cmd
│  │  │  Dockerfile-judge_api
│  │  │  Dockerfile-judge_rpc
│  │  │
│  │  ├─api
│  │  │  │  judge-api
│  │  │  │  judgestatus.go
│  │  │  │
│  │  │  ├─desc
│  │  │  │      judgeStatus.api
│  │  │  │      judgeStatus_info.api
│  │  │  │
│  │  │  ├─etc
│  │  │  │      judgestatus-api.yaml
│  │  │  │
│  │  │  ├─internal
│  │  │  │  ├─config
│  │  │  │  │      config.go
│  │  │  │  │
│  │  │  │  ├─handler
│  │  │  │  │  │  routes.go
│  │  │  │  │  │
│  │  │  │  │  └─judgeStatus
│  │  │  │  │          deletejudgestatushandler.go
│  │  │  │  │          getjudgestatushandler.go
│  │  │  │  │          listjudgestatushandler.go
│  │  │  │  │          sendjudgestatushandler.go
│  │  │  │  │
│  │  │  │  ├─logic
│  │  │  │  │  └─judgeStatus
│  │  │  │  │          deletejudgestatuslogic.go
│  │  │  │  │          getjudgestatuslogic.go
│  │  │  │  │          listjudgestatuslogic.go
│  │  │  │  │          sendjudgestatuslogic.go
│  │  │  │  │
│  │  │  │  ├─svc
│  │  │  │  │      servicecontext.go
│  │  │  │  │
│  │  │  │  └─types
│  │  │  │          types.go
│  │  │  │
│  │  │  └─logs
│  │  │
│  │  └─rpc
│  │      │  judge-rpc
│  │      │  judgestatus.go
│  │      │  judgeStatus.proto
│  │      │
│  │      ├─etc
│  │      │      judgestatus.yaml
│  │      │
│  │      ├─internal
│  │      │  ├─code
│  │      │  │      code.go
│  │      │  │
│  │      │  ├─config
│  │      │  │      config.go
│  │      │  │
│  │      │  ├─logic
│  │      │  │      addjudgestatuslogic.go
│  │      │  │      deljudgestatuslogic.go
│  │      │  │      getjudgestatusbyidlogic.go
│  │      │  │      searchjudgestatuslogic.go
│  │      │  │      updatejudgestatuslogic.go
│  │      │  │
│  │      │  ├─server
│  │      │  │      judgeserviceserver.go
│  │      │  │
│  │      │  └─svc
│  │      │          servicecontext.go
│  │      │
│  │      ├─judgeservice
│  │      │      judgeservice.go
│  │      │
│  │      ├─logs
│  │      │
│  │      └─pb
│  │              judgestatus.pb.go
│  │              judgestatus_grpc.pb.go
│  │
│  ├─docker-judge
│  │  │  docker-compose-judge.yml
│  │  │  judgestatus-api.yaml
│  │  │  judgestatus.yaml
│  │  │
│  │  └─logs
│  │
│  └─model
│          judgestatusmodel.go
│          judgestatusmodel_gen.go
│          vars.go
│
├─k8s_files
│  │  deployment-etcd.yaml
│  │  deployment-redis.yaml
│  │  storage-class.yaml
│  │
│  ├─ingress
│  │      configmap-nginx.yaml
│  │      deployment-nginx.yaml
│  │      ingress-traefik.yaml
│  │      service-nginx.yaml
│  │
│  ├─judge
│  │      configmap-judge-api.yaml
│  │      configmap-judge-rpc.yaml
│  │      deployment-judge-api.yaml
│  │      deployment-judge-rpc.yaml
│  │
│  ├─minio
│  │      pv-minio.yaml
│  │      service-minio-backend.yaml
│  │      service-minio-judge.yaml
│  │      statefulSet-minio.yaml
│  │
│  ├─mysql
│  │      service-mysql-cluster.yaml
│  │      service-mysql-judge.yaml
│  │
│  ├─nats
│  │      configmap-nats.yaml
│  │      deployment-nats.yaml
│  │      pvc-nats.yaml
│  │      service-nats.yaml
│  │
│  ├─other
│  │      configmap-other-api.yaml
│  │      deployment-other-api.yaml
│  │
│  ├─problems
│  │      configmap-problems-api.yaml
│  │      configmap-problems-rpc.yaml
│  │      deployment-problems-api.yaml
│  │      deployment-problems-rpc.yaml
│  │
│  └─users
│          configmap-users-api.yaml
│          configmap-users-rpc.yaml
│          deployment-users-api.yaml
│          deployment-users-rpc.yaml
│          pvc-users.yaml
│
├─other
│  │  .gitignore
│  │  Dockerfile
│  │  Jenkinsfile
│  │  other-api
│  │  other.api
│  │  upload.go
│  │
│  ├─docker-other
│  │  │  docker-compose-other.yml
│  │  │  upload-api.yaml
│  │  │
│  │  └─logs
│  │
│  ├─etc
│  │      upload-api.yaml
│  │
│  ├─internal
│  │  ├─code
│  │  │      code.go
│  │  │
│  │  ├─config
│  │  │      config.go
│  │  │
│  │  ├─handler
│  │  │  │  routes.go
│  │  │  │
│  │  │  └─upload_file
│  │  │          addtestcasehandler.go
│  │  │          deletetestcasehandler.go
│  │  │          updateusercoverhandler.go
│  │  │
│  │  ├─logic
│  │  │  └─upload_file
│  │  │          addtestcaselogic.go
│  │  │          deletetestcaselogic.go
│  │  │          updateusercoverlogic.go
│  │  │
│  │  ├─svc
│  │  │  │  servicecontext.go
│  │  │  │
│  │  │  └─pkg
│  │  │          file.go
│  │  │
│  │  └─types
│  │          types.go
│  │
│  └─logs
│
├─problems
│  │  .gitignore
│  │  Jenkinsfile
│  │
│  ├─cmd
│  │  │  Dockerfile-problems_api
│  │  │  Dockerfile-problems_rpc
│  │  │
│  │  ├─api
│  │  │  │  problems-api
│  │  │  │  problems.go
│  │  │  │
│  │  │  ├─desc
│  │  │  │  │  problem.api
│  │  │  │  │
│  │  │  │  ├─problem
│  │  │  │  │      problem_data.api
│  │  │  │  │      problem_info.api
│  │  │  │  │
│  │  │  │  ├─tag
│  │  │  │  │      tag.api
│  │  │  │  │
│  │  │  │  └─test_case
│  │  │  │          test_case.api
│  │  │  │
│  │  │  ├─etc
│  │  │  │      problems-api.yaml
│  │  │  │
│  │  │  ├─internal
│  │  │  │  ├─code
│  │  │  │  │      code.go
│  │  │  │  │
│  │  │  │  ├─config
│  │  │  │  │      config.go
│  │  │  │  │
│  │  │  │  ├─handler
│  │  │  │  │  │  routes.go
│  │  │  │  │  │
│  │  │  │  │  ├─problems
│  │  │  │  │  │      addproblemhandler.go
│  │  │  │  │  │      deleteproblemhandler.go
│  │  │  │  │  │      getproblembyidhandler.go
│  │  │  │  │  │      listproblemsbytagidhandler.go
│  │  │  │  │  │      searchproblemshandler.go
│  │  │  │  │  │      updateproblemhandler.go
│  │  │  │  │  │
│  │  │  │  │  ├─problem_data
│  │  │  │  │  │      getproblemdatahandler.go
│  │  │  │  │  │      searchproblemdatashandler.go
│  │  │  │  │  │      searchproblemsdatahandler.go
│  │  │  │  │  │      updateproblemdatahandler.go
│  │  │  │  │  │
│  │  │  │  │  ├─tags
│  │  │  │  │  │      addtaghandler.go
│  │  │  │  │  │      deletetaghandler.go
│  │  │  │  │  │      gettagbyidhandler.go
│  │  │  │  │  │      listtagsbyproblemidhandler.go
│  │  │  │  │  │      searchtagshandler.go
│  │  │  │  │  │      updatetaghandler.go
│  │  │  │  │  │
│  │  │  │  │  └─test_case
│  │  │  │  │          gettestcasebyidhandler.go
│  │  │  │  │          searchtestcaseshandler.go
│  │  │  │  │
│  │  │  │  ├─logic
│  │  │  │  │  ├─problems
│  │  │  │  │  │      addproblemlogic.go
│  │  │  │  │  │      deleteproblemlogic.go
│  │  │  │  │  │      getproblembyidlogic.go
│  │  │  │  │  │      listproblemsbytagidlogic.go
│  │  │  │  │  │      searchproblemslogic.go
│  │  │  │  │  │      updateproblemlogic.go
│  │  │  │  │  │
│  │  │  │  │  ├─problem_data
│  │  │  │  │  │      getproblemdatalogic.go
│  │  │  │  │  │      searchproblemsdatalogic.go
│  │  │  │  │  │      updateproblemdatalogic.go
│  │  │  │  │  │
│  │  │  │  │  ├─tags
│  │  │  │  │  │      addtaglogic.go
│  │  │  │  │  │      deletetaglogic.go
│  │  │  │  │  │      gettagbyidlogic.go
│  │  │  │  │  │      listtagsbyproblemidlogic.go
│  │  │  │  │  │      searchtagslogic.go
│  │  │  │  │  │      updatetaglogic.go
│  │  │  │  │  │
│  │  │  │  │  └─test_case
│  │  │  │  │          gettestcasebyidlogic.go
│  │  │  │  │          searchtestcaseslogic.go
│  │  │  │  │
│  │  │  │  ├─svc
│  │  │  │  │      servicecontext.go
│  │  │  │  │
│  │  │  │  └─types
│  │  │  │          types.go
│  │  │  │
│  │  │  └─logs
│  │  │
│  │  └─rpc
│  │      │  problem.go
│  │      │  problem.proto
│  │      │  problems-rpc
│  │      │
│  │      ├─etc
│  │      │      problem.yaml
│  │      │
│  │      ├─internal
│  │      │  ├─code
│  │      │  │      code.go
│  │      │  │
│  │      │  ├─config
│  │      │  │      config.go
│  │      │  │
│  │      │  ├─logic
│  │      │  │      addproblemlogic.go
│  │      │  │      addtaglogic.go
│  │      │  │      addtestcaseslogic.go
│  │      │  │      delproblemlogic.go
│  │      │  │      deltaglogic.go
│  │      │  │      deltestcaseslogic.go
│  │      │  │      getproblembyidlogic.go
│  │      │  │      getproblemdatabyidlogic.go
│  │      │  │      getproblemdatabyproblemidlogic.go
│  │      │  │      gettagbyidlogic.go
│  │      │  │      gettestcasesbyidlogic.go
│  │      │  │      gettestcasesbyproblemidandtestgrouplogic.go
│  │      │  │      listproblemsbytagidlogic.go
│  │      │  │      listtagsbyproblemidlogic.go
│  │      │  │      searchproblemdatalogic.go
│  │      │  │      searchproblemlogic.go
│  │      │  │      searchtaglogic.go
│  │      │  │      searchtestcaseslogic.go
│  │      │  │      updateproblemdatalogic.go
│  │      │  │      updateproblemlogic.go
│  │      │  │      updatetaglogic.go
│  │      │  │      updatetestcaseslogic.go
│  │      │  │
│  │      │  ├─server
│  │      │  │      problemserviceserver.go
│  │      │  │
│  │      │  └─svc
│  │      │          servicecontext.go
│  │      │
│  │      ├─logs
│  │      │
│  │      ├─pb
│  │      │      problem.pb.go
│  │      │      problem_grpc.pb.go
│  │      │
│  │      └─problemservice
│  │              problemservice.go
│  │
│  ├─docker-problems
│  │  │  docker-compose-problems.yml
│  │  │  problem.yaml
│  │  │  problems-api.yaml
│  │  │
│  │  └─logs
│  │
│  └─model
│          problemdatamodel.go
│          problemdatamodel_gen.go
│          problemmodel.go
│          problemmodel_gen.go
│          problemtagmodel.go
│          problemtagmodel_gen.go
│          tagmodel.go
│          tagmodel_gen.go
│          testcasesmodel.go
│          testcasesmodel_gen.go
│          vars.go
│
└─users
    │  .gitignore
    │  Jenkinsfile
    │
    ├─cmd
    │  │  Dockerfile-users_api
    │  │  Dockerfile-users_rpc
    │  │
    │  ├─api
    │  │  │  users-api
    │  │  │  users.go
    │  │  │
    │  │  ├─desc
    │  │  │  │  user.api
    │  │  │  │
    │  │  │  ├─user_info
    │  │  │  │      user_info.api
    │  │  │  │
    │  │  │  └─user_profile
    │  │  │          user_profile.api
    │  │  │
    │  │  ├─etc
    │  │  │      users-api.yaml
    │  │  │
    │  │  ├─internal
    │  │  │  ├─code
    │  │  │  │      code.go
    │  │  │  │
    │  │  │  ├─config
    │  │  │  │      config.go
    │  │  │  │
    │  │  │  ├─handler
    │  │  │  │  │  routes.go
    │  │  │  │  │
    │  │  │  │  ├─other
    │  │  │  │  │      sendemailtouserhandler.go
    │  │  │  │  │
    │  │  │  │  ├─users_info
    │  │  │  │  │      authenticateuserhandler.go
    │  │  │  │  │      createuserhandler.go
    │  │  │  │  │      deleteuserhandler.go
    │  │  │  │  │      getuserinfohandler.go
    │  │  │  │  │      listuserhandler.go
    │  │  │  │  │      loginuserhandler.go
    │  │  │  │  │      loginwithcodehandler.go
    │  │  │  │  │      updateuserinfohandler.go
    │  │  │  │  │
    │  │  │  │  └─users_profile
    │  │  │  │          getprofilehandler.go
    │  │  │  │          listprofilehandler.go
    │  │  │  │          refreshuserratinghandler.go
    │  │  │  │          updateprofilehandler.go
    │  │  │  │
    │  │  │  ├─logic
    │  │  │  │  ├─other
    │  │  │  │  │      sendemailtouserlogic.go
    │  │  │  │  │
    │  │  │  │  ├─users_info
    │  │  │  │  │      authenticateuserlogic.go
    │  │  │  │  │      createuserlogic.go
    │  │  │  │  │      deleteuserlogic.go
    │  │  │  │  │      getuserinfologic.go
    │  │  │  │  │      listuserlogic.go
    │  │  │  │  │      loginuserlogic.go
    │  │  │  │  │      loginwithcodelogic.go
    │  │  │  │  │      updateuserinfologic.go
    │  │  │  │  │
    │  │  │  │  └─users_profile
    │  │  │  │          getprofilelogic.go
    │  │  │  │          listprofilelogic.go
    │  │  │  │          refreshuserratinglogic.go
    │  │  │  │          updateprofilelogic.go
    │  │  │  │
    │  │  │  ├─svc
    │  │  │  │  │  servicecontext.go
    │  │  │  │  │
    │  │  │  │  └─email
    │  │  │  │          email.go
    │  │  │  │          redis.go
    │  │  │  │
    │  │  │  └─types
    │  │  │          types.go
    │  │  │
    │  │  └─logs
    │  │
    │  └─rpc
    │      │  user.go
    │      │  user.proto
    │      │  users-rpc
    │      │
    │      ├─etc
    │      │      user.yaml
    │      │
    │      ├─internal
    │      │  ├─code
    │      │  │      code.go
    │      │  │
    │      │  ├─config
    │      │  │      config.go
    │      │  │
    │      │  ├─logic
    │      │  │      adduserloginlogic.go
    │      │  │      deluserloginlogic.go
    │      │  │      getrankbyuseridlogic.go
    │      │  │      getuserloginbyidlogic.go
    │      │  │      getuserprofilebyidlogic.go
    │      │  │      loginuserlogic.go
    │      │  │      loginwithcodelogic.go
    │      │  │      partialupdateuserlogic.go
    │      │  │      searchuserloginlogic.go
    │      │  │      searchuserprofilelogic.go
    │      │  │      updateuserloginlogic.go
    │      │  │      updateuserprofilelogic.go
    │      │  │
    │      │  ├─server
    │      │  │      userserviceserver.go
    │      │  │
    │      │  └─svc
    │      │          servicecontext.go
    │      │
    │      ├─logs
    │      │
    │      ├─pb
    │      │      user.pb.go
    │      │      user_grpc.pb.go
    │      │
    │      └─userservice
    │              userservice.go
    │
    ├─docker-users
    │  │  docker-compose-users.yml
    │  │  user.yaml
    │  │  users-api.yaml
    │  │
    │  └─logs
    │
    └─model
            userloginmodel.go
            userloginmodel_gen.go
            userprofilemodel.go
            userprofilemodel_gen.go
            vars.go