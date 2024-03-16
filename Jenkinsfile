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
                sh "make stop-and-remove-container"
            }
        }
        stage('Cleaning up images') {
            steps {
                sh "make clean-images"
            }
        }
        stage("Building images"){
            steps{
                script{
                    sh "make build-prod"
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
    }
}