// properties([pipelineTriggers([githubPush()])])

pipeline {
    environment {
        dockerRepo = "where-is-everyone-bot"

        dockerImageVersioned = ""
        dockerImageLatest = ""
    }

    agent any

    stages {
        stage('Cleaning up docker images') {
            steps {
                sh "docker rmi $(docker images -a -q)"
            }
        }
        stage('Cleaning up') {
            steps {
                sh "make clean"
            }
        }
        stage("Building docker image"){
            steps{
                script{
                    // dockerImageVersioned = docker.build dockerRepo + ":$BUILD_NUMBER"
                    // dockerImageLatest = docker.build dockerRepo + ":latest"
                    // sh "docker stop $(docker ps -a -q)"
	                // sh "docker rm $(docker ps -a -q)"
	                // sh "docker rmi $(docker images -a -q)"
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
    }

    /* Cleanup workspace */
//     post {
//        always {
//            deleteDir()
//        }
//    }
}