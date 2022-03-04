#!/bin/bash

awslocal sns create-topic --name EMPLOYEE
awslocal sqs create-queue --queue-name EMPLOYEE_PAYROLL
awslocal sns subscribe --topic-arn arn:aws:sns:us-east-1:000000000000:EMPLOYEE \
         --protocol sqs \
         --notification-endpoint arn:aws:sqs:us-east-1:queue:EMPLOYEE_PAYROLL
