# ticketing

# Authentication service

- For now i am using localhost mongodb database instance, later on i will be switched to the docker containered database.

## Tasks to do
- token based authentication // store token in mongodb database for verification from client 
- role based access (user, Admin) // in future add casbin role authorization


## Dependencies

1. Docker
2. Kubernetes
3. Minikube with addons of ingress
4. Hosts file with required host configuration ( specify ip in host file that returned by `kubectl get ingress`)
5. 

## Running Project Using Skaffold

- It requires skaffold to be installed in machine
- Install skaffold using below command

```
# For Linux x86_64 (amd64)
curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-linux-amd64 && \
sudo install skaffold /usr/local/bin/

```

To start project in dev mode

```
start docker
minikube start
skaffold dev
```

Above command will start all services in dev mode. Changing anything in code will prompt it to regenerate the image and run it.