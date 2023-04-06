# Welcome to Summary Service

Hello! My name is Javier and in this repository we are going to find a service that processes, sends by mail and saves in a database a financial summary of a user. 

# Design

This repository is structured according to the Hexagonal Architecture based on what we can find in [this article](https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3).

## The structure

The structure of the service is as follow

![summary-service](https://user-images.githubusercontent.com/69270095/230511544-a01b31ac-dc1e-4180-9c17-fe91e323f082.png)

We have a lambda that is triggered for S3 bucket, so when we upload a file to the bucket this send the info to the lambda. The lambda process the info, send it for mail using SES and save it in Dynamodb table. There are many ways to improve the service, but I didn't have much time to do so.

## How to test service

Well this services is deployed in my personal account of AWS, so if you want to test it you have to contact me, because if I make this publicly available, it could generate costs for me.

