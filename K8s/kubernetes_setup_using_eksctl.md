## Setup Kubernetes on Amazon EKS

You can follow same procedure in the official AWS document Getting started with Amazon EKS – eksctl
Pre-requisites:

    an EC2 Instance

## AWS EKS Setup

    Setup kubectl
    a. Download kubectl version 1.20
    b. Grant execution permissions to kubectl executable
    c. Move kubectl onto /usr/local/bin
    d. Test that your kubectl installation was successful

    curl -o kubectl https://amazon-eks.s3.us-west-2.amazonaws.com/1.19.6/2021-01-05/bin/linux/amd64/kubectl
    chmod +x ./kubectl
    mv ./kubectl /usr/local/bin 
    kubectl version --short --client

    Setup eksctl
    a. Download and extract the latest release
    b. Move the extracted binary to /usr/local/bin
    c. Test that your eksclt installation was successful

    curl --silent --location "https://github.com/weaveworks/eksctl/releases/latest/download/eksctl_$(uname -s)_amd64.tar.gz" | tar xz -C /tmp
    sudo mv /tmp/eksctl /usr/local/bin
    eksctl version

    Create an IAM Role and attache it to EC2 instance
    Note: create IAM user with programmatic access if your bootstrap system is outside of AWS
    IAM user should have access to
    IAM
    EC2
    VPC
    CloudFormation

    Create your cluster and nodes

    eksctl create cluster --name cluster-name  \
    --region region-name \
    --node-type instance-type \
    --nodes-min 2 \
    --nodes-max 2 \ 
    --zones <AZ-1>,<AZ-2>

    example:
    eksctl create cluster --name valaxy-cluster \
       --region ap-south-1 \
    --node-type t2.small \

    To delete the EKS clsuter

    eksctl delete cluster valaxy --region ap-south-1

    Validate your cluster using by creating by checking nodes and by creating a pod

    kubectl get nodes
    kubectl run pod tomcat --image=tomcat 