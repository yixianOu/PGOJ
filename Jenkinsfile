pipeline {
    agent any // 指定任何可用的代理节点上运行此流水线

    environment { // 定义环境变量
        PROJECT_NAME = 'yixianOu/PGOJ'
        BRANCH_NAME = 'services/users'
    }

    stages { // 定义流水线的不同阶段
        stage('Checkout') { // 第一个阶段：从代码仓库检出代码
            steps {
                git branch: "${BRANCH_NAME}", url: 'https://github.com/user/repo.git' // 使用git插件检出代码
            }
        }

        stage('Build') { // 构建阶段
            steps {
                sh 'docker-compose -f ./docker-users/docker-compose-users.yml' up -d // 使用shell命令调用Maven进行构建
            }
        }

        stage('Test') { // 测试阶段
            steps {
                echo 'skip test' // 运行测试
            }
        }

        stage('Deploy') { // 部署阶段
            steps {
                sh 'ls'
            }
        }
    }

    post { // 流水线执行后的操作
        always {
            cleanWs() // 清理工作空间
        }
        success {
            echo 'Pipeline completed successfully!'
        }
        failure {
            echo 'Pipeline failed, please check the logs for more information.'
        }
    }
}