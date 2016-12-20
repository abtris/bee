pipeline {
  agent any
  environment {
    PACKAGE="github.com/abtris/bee"
    GOPATH="/Users/abtris/go"
    GOROOT="/usr/local/opt/go/libexec"
  }
  stages {
     stage('Preparation') {
        steps {
          git 'https://github.com/abtris/bee.git'
        }
     }
     stage('Deps') {
        steps {
           sh 'glide --no-color install'
           sh 'mkdir -p release'
       }
     }
     stage('Test') {
        steps {
         sh 'make xunit'
        }
     }
     stage('Build') {
        steps {
           parallel (
              linux64: { sh "GOOS=linux GOARCH=amd64 go build -o release/bee-linux-amd64 ${PACKAGE}" },
              linux32: { sh "GOOS=linux GOARCH=386 go build -o release/bee-linux-386 ${PACKAGE}" },
              mac64: { sh "GOOS=darwin GOARCH=amd64 go build -o release/bee-darwin-amd64 ${PACKAGE}" },
              win64: { sh "GOOS=windows GOARCH=amd64 go build -o release/bee-windows-amd64 ${PACKAGE}" }
            )
        }
     }
     stage('Results') {
        steps {
          archive 'release/*'
          junit 'tests.xml'
        }
     }
  }
}
