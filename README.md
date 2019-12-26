<p align="center">
  <h3 align="center">Launchpad</h3>
  <p align="center">Backend API for the launchpad application</p>
  <p align="center">
    <a href="https://goreportcard.com/report/github.com/getlaunchpad/server"><img src="https://goreportcard.com/badge/github.com/getlaunchpad/server"></a>
    <a href="https://circleci.com/gh/getlaunchpad/server"><img src="https://circleci.com/gh/getlaunchpad/server.svg?style=svg"></a>
    <a href="https://codecov.io/gh/getlaunchpad/server"><img src="https://codecov.io/gh/getlaunchpad/server/branch/master/graph/badge.svg" /></a>
  </p>
</p>

> Server still in very early phases of development

### Helpful Commands

##### Pushing code via GitHub Packages

If you need help getting your creds in order check out [this gist](https://gist.github.com/LucasStettner/66b2108d0fd9663f2c09db5556f69d39)

```shell
# Build docker container
$ docker build -t  docker.pkg.github.com/getlaunchpad/launchpad-server/launchpad:latest .

# Publish to Github Packages
$ docker push docker.pkg.github.com/getlaunchpad/launchpad-server/launchpad:latest
```

##### Deploying kubernetes locally

```shell
# Started k8 process
$ minikube start

# Apply k8 config
$ kubectl apply -f k8s-deployment.yml

# Open node port
$ minikube service launchpad-service
```
