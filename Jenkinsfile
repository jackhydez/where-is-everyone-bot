pipeline {
    environment {
        dockerRepo = "where-is-everyone-bot"

        dockerImageVersioned = ""
        dockerImageLatest = ""
    }

    agent any

    stages {
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
                    sh "make run-prod"
                }
            }
        }
        stage('Cleaning up containers') {
            steps {
                sh "make clean-containers"
            }
        }
        stage('Cleaning up images') {
            steps {
                sh "make clean-images"
            }
        }
    }
}