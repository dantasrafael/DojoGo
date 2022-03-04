#!/bin/bash

awslocal sns create-topic --name
awslocal sqs create-queue --queue-name SCHOOL_ENROLLMENT_FINANCIAL
awslocal sns subscribe --topic-arn arn:aws:sns:us-east-1:000000000000:SCHOOL_ENROLLMENT \
         --protocol sqs \
         --notification-endpoint arn:aws:sqs:us-east-1:queue:SCHOOL_ENROLLMENT_FINANCIAL
