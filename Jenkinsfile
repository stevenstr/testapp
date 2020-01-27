podTemplate(containers: [
    containerTemplate(name: 'golang', image: 'golang:1.13.0', ttyEnabled: true, command: 'cat')
  ]) {

    node(POD_LABEL) {
        stage('Get a Golang project') {
            container('golang') {
                stage('Build a Go project') {
                    sh 'go version'
                }
            }
        }

    }
}
