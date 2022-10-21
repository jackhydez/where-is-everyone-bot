properties([pipelineTriggers([githubPush()])])

pipeline {
    environment {
        dockerRepo = "jackhydez/where-is-everyone-bot"

        dockerImageVersioned = ""
        dockerImageLatest = ""
    }

    agent any

    stages {
        stage("Building docker image"){
            steps{
                script{
                    dockerImageVersioned = docker.build dockerRepo + ":$BUILD_NUMBER"
                    dockerImageLatest = docker.build dockerRepo + ":latest"
                    sh "make build"
                }
            }
        }
        stage("Run docker container"){
            steps{
                script{
                    sh "make run"
                }
            }
        }
        stage('Cleaning up') {
            steps {
                sh "docker rmi $dockerRepo:$BUILD_NUMBER"
            }
        }
    }

    /* Cleanup workspace */
//     post {
//        always {
//            deleteDir()
//        }
//    }
}