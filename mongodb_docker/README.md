### This docker is for installing the mongodb and running in the particular port

Steps to Run the simple mongodb server

1. Clean up the existing the docker images (Be caution it will remove every images in the server)

 ```
 docker rm $(docker ps -a -q)
 docker rmi $(docker images -q)
 ```

2. Build the docker container with the below command
```
docker build -t mongodb .'
```

3. Run the docker container with the below command
```
docker run -it -p 5000:5000 -d mongodb'
```
