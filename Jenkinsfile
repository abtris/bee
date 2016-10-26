node {
   stage('Preparation') {
      git 'https://github.com/abtris/bee.git'
   }
   stage('Deps') {
       env.PACKAGE="github.com/abtris/bee"
       env.GOPATH="~/go"
       env.GOROOT="/usr/local/opt/go/libexec"
       sh 'glide --no-color install'
       sh 'mkdir -p release'
   }
   stage('Test') {
       sh 'make xunit'
   }
   stage('Build') {
         parallel (
            linux64: { sh "goos=linux goarch=amd64 make build" },
            linux32: { sh "goos=linux goarch=386 make build" },
            mac64: { sh "goos=darwin goarch=amd64 make build" },
            win64: { sh "goos=windows goarch=amd64 make build" }
          )
   }
   stage('Results') {
      archive 'release/*'
      junit 'tests.xml'
   }
}
