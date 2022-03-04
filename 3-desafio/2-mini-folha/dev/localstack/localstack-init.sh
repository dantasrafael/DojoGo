#!/bin/bash

awslocal sns create-topic --name COLABORADORES
awslocal sqs create-queue --queue-name COLABORADORES_FOLHA
awslocal sns subscribe --topic-arn arn:aws:sns:us-east-1:000000000000:COLABORADORES \
         --protocol sqs \
         --notification-endpoint arn:aws:sqs:us-east-1:queue:COLABORADORES_FOLHA
