## Latest Docker image with Newman for GitLab Continuous Integration Pipeline

### Description
This image runs Notifyer-1.0.0 on Golang 1.14.6

This image does not have an entrypoint so it can be used in GitLab CI Pipeline.

Docker hub url: 
 - [skeletondocker/notifyer-ci](https://hub.docker.com/r/skeletondocker/notifyer-ci/)

#### Gitlab CI
Example .gitlab-ci.yml:

```
...your pipeline config...
stages:
  - notify_release_ready

notify_release_ready:
  stage: notify_release_ready
  image: skeletondocker/notifyer-ci
  script:
    - notifyer release ready --text "Release ready! @Rishats"
```

#### Docker

Docker Pull Command

```
docker pull skeletondocker/notifyer-ci
```

Docker Run Command

```
docker run skeletondocker/notifyer-ci notifyer --help
```