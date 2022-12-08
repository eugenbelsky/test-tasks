
# How-to

```bash
docker build . -t eugnebelskyi/nginx-joynnginx-joyn:latest
docker push eugnebelskyi/nginx-joynnginx-joyn:latest
```
Note: this image is already pushed to dockerhub and is ready to use

run in [kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

```bash
kind cluster create
kubectl apply -f k8s/deployment.yaml
kubect get pods 
kubect get svc
```
