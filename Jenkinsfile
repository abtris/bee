node {
   stage('Preparation') {
      git 'https://github.com/abtris/bee.git'
      poll: true

   }
   stage('Deps') {
       env.PACKAGE="github.com/abtris/bee"
       env.GOPATH="/Users/abtris/go"
       env.GOROOT="/usr/local/opt/go/libexec"
       sh 'glide --no-color install'
       sh 'mkdir -p release'
   }
   stage('Test') {
       sh 'make xunit'
   }
   stage('Build') {
         parallel (
            linux64: { sh "GOOS=linux GOARCH=amd64 go build -o release/bee-linux-amd64 ${PACKAGE}" },
            linux32: { sh "GOOS=linux GOARCH=386 go build -o release/bee-linux-386 ${PACKAGE}" },
            mac64: { sh "GOOS=darwin GOARCH=amd64 go build -o release/bee-darwin-amd64 ${PACKAGE}" },
            win64: { sh "GOOS=windows GOARCH=amd64 go build -o release/bee-windows-amd64 ${PACKAGE}" }
          )
   }
   stage('Results') {
      archive 'release/*'
      junit 'tests.xml'
   }
}
