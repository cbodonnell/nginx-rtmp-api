pipeline {
    agent any
    environment {
        GOROOT = "${tool type: 'go', name: 'go1.15.6'}/go"
    }
    stages {
        stage('build') {
            steps {
                echo 'building...'
                sh 'echo $GOROOT'
                sh '$GOROOT/bin/go build'
            }
        }
        stage('test') {
            steps {
                echo 'testing...'
            }
        }
        stage('deploy') {
            steps {
                echo 'deploying...'
                sh 'sudo systemctl stop nginx-rtmp-api'
                sh 'sudo cp nginx-rtmp-api /etc/nginx-rtmp-api/nginx-rtmp-api'
                sh 'sudo systemctl start nginx-rtmp-api'
            }
        }
    }
    post {
        cleanup {
            deleteDir()
        }
    }
}