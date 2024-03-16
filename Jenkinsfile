pipeline {
    environment {
        dockerRepo = "where-is-everyone-bot"

        dockerImageVersioned = ""
        dockerImageLatest = ""
    }

    agent any

    stages {
        stage('Stop and remove containter') {
            steps {
                sh "make stop"
            }
        }
        stage('Cleaning up images') {
            steps {
                sh "make clean"
            }
        }
        stage("Building images"){
            steps{
                script{
                    sh "make build"
                }
            }
        }
        stage("Run containers"){
            steps{
                script{
                    sh "make run"
                }
            }
        }
    }
}