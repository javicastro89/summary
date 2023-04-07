# Welcome to Summary Service

Hello! My name is Javier and in this repository we are going to find a service that processes, sends by mail and saves in a database a financial summary of a user. 

# Design

 This repository is structured according to the Hexagonal Architecture based on what we can find in [this article](https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3).

## The architecture

>The architecture of the service is as follow

![summary-service](https://user-images.githubusercontent.com/69270095/230511544-a01b31ac-dc1e-4180-9c17-fe91e323f082.png)

We have a lambda that is triggered by a S3 bucket, so when we upload a file to the bucket this trigger the lambda. The lambda process the info, send it for mail using SES and save some info about the summary in a Dynamodb table. 

## How to test service

To test the service you have to: 

> - [Create a S3 bucket](https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-bucket.html).
> 
> - Then we have to [deploy](https://docs.aws.amazon.com/lambda/latest/dg/gettingstarted-package.html) this lambda and set a trigger from the S3 bucket we have created (setting the corresponding permissions).
> 
>  - To deploy the lambda we can use the *.zip* file that is inside the **bin** folder, if you use this *.zip* file you have to set the runtime as **Custom runtime on Amazon Linux 2** and the architecture as **arm64** because this lambda is build to run using [graviton 2](https://aws.amazon.com/blogs/aws/aws-lambda-functions-powered-by-aws-graviton2-processor-run-your-functions-on-arm-and-get-up-to-34-better-price-performance/) architecture. [Here](https://catalog.workshops.aws/serverless-optimization/en-US/graviton/3-deploy-graviton-function) is a step-by-step on how to migrate a lambda function to **graviton 2**. 
>  - Then we have to create a dynamodb table and set as **Partition key** the field **Email**, (we have to add to the lambda the corresponding permissions to access this table). In order for the lambda to access the table, it is necessary to pass the table name through the **TABLE_NAME** environment variable.
>  
>  - Finally, we have to use an email service, in this case I have used SES. With the host, port, username and password provided by SES that we can find in its configuration, we pass this data to the lambda through the environment variables **HOST**, **PORT**, **USERNAME** and **PASSWORD** respectively. [Here](https://blog.devgenius.io/sending-emails-with-golang-and-amazon-ses-31f25a0f2acb) is an example of how we can configure the SES service. Also, you have to put the e-mail address from where the mail is going to be sent, this has to be passed to the lambda through the environment variable **ORIGIN_EMAIL**.
>  
>  - Well we are ready to test the service, for that we can use the file test_summary.csv that is in the bin folder, we have to change the id field by any other or other emails. We can put different emails and the service separates them as transactions of different people and processes them separately sending the corresponding email with their movements to each user. Then we upload this file to the S3 bucket and we will receive the information in the email or emails that we have configured.
