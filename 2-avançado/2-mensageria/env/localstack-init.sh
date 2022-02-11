#!/bin/bash

awslocal sns create-topic --name GOJO_USER_CREATE
awslocal sqs create-queue --queue-name GOJO_USER_CREATE_CONSUMER
awslocal sns subscribe --topic-arn arn:aws:sns:us-east-1:000000000000:GOJO_USER_CREATE \
         --protocol sqs \
         --notification-endpoint arn:aws:sqs:us-east-1:queue:GOJO_USER_CREATE_CONSUMER
