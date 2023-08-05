pipeline {
    // install golang 1.14 on Jenkins node
    agent any
    tools {
        go 'go1.14'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage("unit-test") {
            steps {
                echo 'UNIT TEST EXECUTION STARTED'
                sh 'go test -v ./api/internal/app/db ./api/internal/controller/user ./api/internal/handler/rest/public/user ./api/internal/repository/user'
            }
        }
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'docker compose build'
            }
        }
    }
}