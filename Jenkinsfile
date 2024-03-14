pipeline {
    environment {
        dockerRepo = "where-is-everyone-bot"

        dockerImageVersioned = ""
        dockerImageLatest = ""
    }

    agent any

    stages {
        stage('Read .env file') {
            steps {
                script {
                    // Путь к вашему файлу .env
                    def envFilePath = '.env'
                    // Прочитать содержимое файла .env
                    def envFileContent = readFile(envFilePath).trim()
                    // Разделить содержимое файла на строки
                    def envLines = envFileContent.tokenize('\n')
                    // Пройти по каждой строке
                    envLines.each { line ->
                        // Разделить строку на имя переменной и значение
                        def parts = line.tokenize('=')
                        // Установить переменную окружения в Jenkins
                        env."${parts[0].trim()}" = parts[1].trim()
                    }
                }
            }
        }
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