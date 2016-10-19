node {
   stage('Preparation') {
      git 'https://github.com/abtris/bee.git'

   }
   stage('Build') {
      if (isUnix()) {
         env.GOPATH="/Users/abtris/go"
         env.GOROOT="/usr/local/opt/go/libexec"
         sh 'glide --no-color install'
         sh 'make release'
      }
   }
   stage('Results') {
      archive 'release/*'
   }
}
