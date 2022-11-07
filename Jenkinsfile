pipeline {
    environment {
        dockerRepo = "where-is-everyone-bot"

        dockerImageVersioned = ""
        dockerImageLatest = ""
    }

    agent any

    stages {
        stage('Cleaning up containers') {
            steps {
                sh "make clean-containers"
            }
        }
        stage("Building docker image"){
            steps{
                script{
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
        stage('Cleaning up images') {
            steps {
                sh "make clean-images"
            }
        }
    }
}