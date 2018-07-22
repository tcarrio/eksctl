package builder

import (
	gfn "github.com/awslabs/goformation/cloudformation"
)

const (
	cfnOutputNodeInstanceRoleARN = "NodeInstanceRole"

	iamPolicyAmazonEKSServicePolicyARN = "arn:aws:iam::aws:policy/AmazonEKSServicePolicy"
	iamPolicyAmazonEKSClusterPolicyARN = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"

	iamPolicyAmazonEKSWorkerNodePolicyARN           = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"
	iamPolicyAmazonEKSCNIPolicyARN                  = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"
	iamPolicyAmazonEC2ContainerRegistryPowerUserARN = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryPowerUser"
	iamPolicyAmazonEC2ContainerRegistryReadOnlyARN  = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
)

var (
	defaultPolicyARNs = []*gfn.StringIntrinsic{
		gfn.NewString(iamPolicyAmazonEKSWorkerNodePolicyARN),
		gfn.NewString(iamPolicyAmazonEKSCNIPolicyARN),
	}
)

func makePolicyDocument(statement map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []interface{}{
			statement,
		},
	}
}

func makeAssumeRolePolicyDocument(service string) map[string]interface{} {
	return makePolicyDocument(map[string]interface{}{
		"Effect": "Allow",
		"Principal": map[string][]string{
			"Service": []string{service},
		},
		"Action": []string{"sts:AssumeRole"},
	})
}

func (c *clusterResourceSet) addResourcesForIAM() {
	refSR := c.newResource("ServiceRole", &gfn.AWSIAMRole{
		AssumeRolePolicyDocument: makeAssumeRolePolicyDocument("eks.amazonaws.com"),
		ManagedPolicyArns: []*gfn.StringIntrinsic{
			gfn.NewString(iamPolicyAmazonEKSServicePolicyARN),
			gfn.NewString(iamPolicyAmazonEKSClusterPolicyARN),
		},
	})
	c.newResource("PolicyNLB", &gfn.AWSIAMPolicy{
		PolicyName: makeName("PolicyNLB"),
		Roles:      makeSlice(refSR),
		PolicyDocument: makePolicyDocument(map[string]interface{}{
			"Effect":   "Allow",
			"Resource": "*",
			"Action": []string{
				"elasticloadbalancing:*",
				"ec2:CreateSecurityGroup",
				"ec2:Describe*",
			},
		}),
	})
}

func (n *nodeGroupResourceSet) addResourcesForIAM() {
	n.newResource("NodeInstanceProfile", &gfn.AWSIAMInstanceProfile{
		Path: gfn.NewString("/"),
		Roles: []*gfn.StringIntrinsic{
			n.newResource("NodeInstanceRole", &gfn.AWSIAMRole{
				Path: gfn.NewString("/"),
				AssumeRolePolicyDocument: makeAssumeRolePolicyDocument("ec2.amazonaws.com"),
				ManagedPolicyArns:        defaultPolicyARNs, // TODO parametrise
			}),
		},
	})
	n.newOutputFromAtt(cfnOutputNodeInstanceRoleARN, "NodeInstanceRole.Arn", true)
}
