package manager

func (c *StackCollection) DeprecatedDeleteStackVPC() error {
	return c.DeleteStack("EKS-" + c.Spec.ClusterName + "-VPC")
}

func (c *StackCollection) DeprecatedDeleteStackServiceRole() error {
	return c.DeleteStack("EKS-" + c.Spec.ClusterName + "-ServiceRole")
}

// func (c *StackCollection) stackParamsDefaultNodeGroup() map[string]string {
// 	regionalAMIs := map[string]string{
// 		// TODO: https://github.com/weaveworks/eksctl/issues/49
// 		// currently source of truth for these is here:
// 		// https://docs.aws.amazon.com/eks/latest/userguide/launch-workers.html
// 		"us-west-2": "ami-73a6e20b",
// 		"us-east-1": "ami-dea4d5a1",
// 	}

// 	if c.Spec.NodeAMI == "" {
// 		c.Spec.NodeAMI = regionalAMIs[c.Spec.Region]
// 	}

// 	if c.Spec.MinNodes == 0 && c.Spec.MaxNodes == 0 {
// 		c.Spec.MinNodes = c.Spec.Nodes
// 		c.Spec.MaxNodes = c.Spec.Nodes
// 	}

// 	if len(c.Spec.PolicyARNs) == 0 {
// 		c.Spec.PolicyARNs = defaultPolicyARNs
// 	}
// 	if c.Spec.Addons.WithIAM.PolicyAmazonEC2ContainerRegistryPowerUser {
// 		c.Spec.PolicyARNs = append(c.Spec.PolicyARNs, iamAmazonEC2ContainerRegistryPowerUserARN)
// 	} else {
// 		c.Spec.PolicyARNs = append(c.Spec.PolicyARNs, iamAmazonEC2ContainerRegistryReadOnlyARN)
// 	}

// 	// params := map[string]string{
// 	// 	"ClusterName":                      c.Spec.ClusterName,
// 	// 	"NodeGroupName":                    "default",
// 	// 	"KeyName":                          c.Spec.keyName,
// 	// 	"NodeImageId":                      c.Spec.NodeAMI,
// 	// 	"NodeInstanceType":                 c.Spec.NodeType,
// 	// 	"NodeAutoScalingGroupMinSize":      fmt.Sprintf("%d", c.Spec.MinNodes),
// 	// 	"NodeAutoScalingGroupMaxSize":      fmt.Sprintf("%d", c.Spec.MaxNodes),
// 	// 	"ClusterControlPlaneSecurityGroup": c.Spec.securityGroup,
// 	// 	"Subnets":                          c.Spec.subnetsList,
// 	// 	"VpcId":                            c.Spec.clusterVPC,
// 	// 	"ManagedPolicyArns":                strings.Join(c.Spec.PolicyARNs, ","),
// 	// }

// 	return params
// }

// func (c *StackCollection) createStackDefaultNodeGroup(errs chan error) error {
// 	name := c.stackNameDefaultNodeGroup()
// 	logger.Info("creating DefaultNodeGroup stack %q", name)
// 	templateBody, err := amazonEksNodegroupYamlBytes()
// 	if err != nil {
// 		return errors.Wrap(err, "decompressing bundled template for DefaultNodeGroup stack")
// 	}

// 	stackChan := make(chan Stack)
// 	taskErrs := make(chan error)

// 	if err := c.CreateStack(name, templateBody, c.stackParamsDefaultNodeGroup(), true, stackChan, taskErrs); err != nil {
// 		return err
// 	}

// 	go func() {
// 		defer close(errs)
// 		defer close(stackChan)

// 		if err := <-taskErrs; err != nil {
// 			errs <- err
// 			return
// 		}

// 		s := <-stackChan

// 		logger.Debug("created DefaultNodeGroup stack %q – processing outputs", name)

// 		nodeInstanceRoleARN := GetOutput(&s, "NodeInstanceRole")
// 		if nodeInstanceRoleARN == nil {
// 			errs <- fmt.Errorf("NodeInstanceRole is nil")
// 			return
// 		}
// 		c.Spec.nodeInstanceRoleARN = *nodeInstanceRoleARN

// 		logger.Debug("clusterConfig = %#v", c.Spec)
// 		logger.Success("created DefaultNodeGroup stack %q", name)

// 		errs <- nil
// 	}()
// 	return nil
// }

func (c *StackCollection) DeprecatedDeleteStackDefaultNodeGroup() error {
	return c.DeleteStack("EKS-" + c.Spec.ClusterName + "-DefaultNodeGroup")
}
