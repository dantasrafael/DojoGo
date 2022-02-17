#!/bin/bash

awslocal sns create-topic --name SCHOOL_ENROLLMENT
awslocal sqs create-queue --queue-name SCHOOL_ENROLLMENT_FINANCIAL
awslocal sns subscribe --topic-arn arn:aws:sns:us-east-1:000000000000:SCHOOL_ENROLLMENT \
         --protocol sqs \
         --notification-endpoint arn:aws:sqs:us-east-1:queue:SCHOOL_ENROLLMENT_FINANCIAL

awslocal sns create-topic --name FINANCIAL_INSTALLMENT
awslocal sqs create-queue --queue-name FINANCIAL_INSTALLMENT_SCHOOL
awslocal sns subscribe --topic-arn arn:aws:sns:us-east-1:000000000000:FINANCIAL_INSTALLMENT \
         --protocol sqs \
         --notification-endpoint arn:aws:sqs:us-east-1:queue:FINANCIAL_INSTALLMENT_SCHOOL
