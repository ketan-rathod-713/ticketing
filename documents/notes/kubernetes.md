# Setting up ingress

- `minikube addons enable ingress`
- Above command will automatically download and configure ingress in our kubernetes cluster.
- Ingress is type of service which is used to expose something from the kubernetes cluster to the outside world.

- For hosts, we have to edit the etc/hosts file and update the ip of our local domain. We should write the ip of the ingress for it. Not that of minikube. 
- // TODO: why we are writting ip of ingress and not of kubernets cluster

## Testing services

- Initially if we haven't setup ingress controller for our service. Then we can make service of type nodePort and expose it. Then test it via `<minikubeIp>:<nodePort>`

## Final Product

- For final thing what we can do is to make cluster ip services of our applications and then using ingress controller define their path.
- By default cluster ip port can be accessible by the ingress controller.

## DockerFile For Golang Applications

- go build command is very important.