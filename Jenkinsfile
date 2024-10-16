pipeline {
    agent {
        docker {
            image 'docker/compose:1.29.2'  // 使用包含 docker-compose 的 Docker 镜像
            args '-v /var/run/docker.sock:/var/run/docker.sock'  // 挂载 Docker 的 socket
        }
    }
    environment { // 定义环境变量
        PROJECT_URL = 'https://github.com/yixianOu/PGOJ.git'
        BRANCH_NAME = 'services/users'
    }

    stages { // 定义流水线的不同阶段
        stage('Checkout') { // 第一个阶段：从代码仓库检出代码
            steps {
                git branch: "${BRANCH_NAME}", url: "${PROJECT_URL}"
            }
        }

        stage('Build') { // 构建阶段
            steps {
                sh 'ls -l ./cmd/rpc'
            }
        }

        stage('Test') { // 测试阶段
            steps {
                echo 'skip test' // 运行测试
            }
        }

        stage('Deploy') { // 部署阶段
            steps {
                script {
                    // 构建 Docker 镜像
                    sh 'docker build -t oj-users-rpc -f ./docker-users/Dockerfile .'

                    // 运行容器并设置二进制文件的可执行权限
                    sh 'docker run --rm -v $(pwd):/app oj-users-rpc chmod +x /app/cmd/rpc/users-rpc'
                    sh 'docker run --rm -v $(pwd):/app oj-users-rpc chmod +x /app/cmd/api/users-api'

                    // 启动 docker-compose 服务
                    sh 'docker-compose -f ./docker-users/docker-compose-users.yml up -d'
                }
            }
        }
    }

    post { // 流水线执行后的操作
        success {
            echo 'Pipeline completed successfully!'
        }
        failure {
            echo 'Pipeline failed, please check the logs for more information.'
        }
    }
}