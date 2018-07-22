package builder_test

import (
	"encoding/base64"

	cfn "github.com/aws/aws-sdk-go/service/cloudformation"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/weaveworks/eksctl/pkg/cfn/builder"
	"github.com/weaveworks/eksctl/pkg/eks/api"
)

func newStackWithOutputs(outputs map[string]string) cfn.Stack {
	s := cfn.Stack{}
	for k, v := range outputs {
		func(k, v string) {
			s.Outputs = append(s.Outputs,
				&cfn.Output{
					OutputKey:   &k,
					OutputValue: &v,
				})
		}(k, v)
	}
	return s
}

var _ = Describe("GetAllOutputs", func() {
	Describe("Cluster", func() {
		Describe("TODO", func() {
			caCert := "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUN5RENDQWJDZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQjRYRFRFNE1EWXdOekExTlRBMU5Wb1hEVEk0TURZd05EQTFOVEExTlZvd0ZURVRNQkVHQTFVRQpBeE1LYTNWaVpYSnVaWFJsY3pDQ0FTSXdEUVlKS29aSWh2Y05BUUVCQlFBRGdnRVBBRENDQVFvQ2dnRUJBTWJoCnpvZElYR0drckNSZE1jUmVEN0YvMnB1NFZweTdvd3FEVDgrdk9zeGs2bXFMNWxQd3ZicFhmYkE3R0xzMDVHa0wKaDdqL0ZjcU91cnMwUFZSK3N5REtuQXltdDFORWxGNllGQktSV1dUQ1hNd2lwN1pweW9XMXdoYTlJYUlPUGxCTQpPTEVlckRabFVrVDFVV0dWeVdsMmxPeFgxa2JhV2gvakptWWdkeW5jMXhZZ3kxa2JybmVMSkkwLzVUVTRCajJxClB1emtrYW5Xd3lKbGdXQzhBSXlpWW82WFh2UVZmRzYrM3RISE5XM1F1b3ZoRng2MTFOYnl6RUI3QTdtZGNiNmgKR0ZpWjdOeThHZnFzdjJJSmI2Nk9FVzBSdW9oY1k3UDZPdnZmYnlKREhaU2hqTStRWFkxQXN5b3g4Ri9UelhHSgpQUWpoWUZWWEVhZU1wQmJqNmNFQ0F3RUFBYU1qTUNFd0RnWURWUjBQQVFIL0JBUURBZ0trTUE4R0ExVWRFd0VCCi93UUZNQU1CQWY4d0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQkFCa2hKRVd4MHk1LzlMSklWdXJ1c1hZbjN6Z2EKRkZ6V0JsQU44WTlqUHB3S2t0Vy9JNFYyUGg3bWY2Z3ZwZ3Jhc2t1Slk1aHZPcDdBQmcxSTFhaHUxNUFpMUI0ZApuMllRaDlOaHdXM2pKMmhuRXk0VElpb0gza2JFdHRnUVB2bWhUQzNEYUJreEpkbmZJSEJCV1RFTTU1czRwRmxUClpzQVJ3aDc1Q3hYbjdScVU0akpKcWNPaTRjeU5qeFVpRDBqR1FaTmNiZWEyMkRCeTJXaEEzUWZnbGNScGtDVGUKRDVPS3NOWlF4MW9MZFAwci9TSmtPT1NPeUdnbVJURTIrODQxN21PRW02Z3RPMCszdWJkbXQ0aENsWEtFTTZYdwpuQWNlK0JxVUNYblVIN2ZNS3p2TDE5UExvMm5KbFU1TnlCbU1nL1pNVHVlUy80eFZmKy94WnpsQ0Q1WT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="
			caCertData, err := base64.StdEncoding.DecodeString(caCert)
			It("should not error", func() { Expect(err).NotTo(HaveOccurred()) })

			expected := &api.ClusterConfig{
				SecurityGroup:            "sg-0b44c48bcba5b7362",
				Subnets:                  []string{"subnet-0f98135715dfcf55f", "subnet-0ade11bad78dced9e", "subnet-0e2e63ff1712bf6ef"},
				VPC:                      "vpc-0e265ad953062b94b",
				Endpoint:                 "https://DE37D8AFB23F7275D2361AD6B2599143.yl4.us-west-2.eks.amazonaws.com",
				CertificateAuthorityData: caCertData,
				ARN:                 "arn:aws:eks:us-west-2:376248598259:cluster/ferocious-mushroom-1532594698",
				NodeInstanceRoleARN: "",
				AvailabilityZones:   []string{"us-west-2b", "us-west-2a", "us-west-2c"},
			}

			initial := &api.ClusterConfig{
				AvailabilityZones: []string{"us-west-2b", "us-west-2a", "us-west-2c"},
			}

			rs := NewClusterResourceSet()
			rs.AddAllResources(initial.AvailabilityZones)

			sampleStack := newStackWithOutputs(map[string]string{
				"SecurityGroup":            "sg-0b44c48bcba5b7362",
				"Subnets":                  "subnet-0f98135715dfcf55f,subnet-0ade11bad78dced9e,subnet-0e2e63ff1712bf6ef",
				"VPC":                      "vpc-0e265ad953062b94b",
				"Endpoint":                 "https://DE37D8AFB23F7275D2361AD6B2599143.yl4.us-west-2.eks.amazonaws.com",
				"CertificateAuthorityData": caCert,
				"ARN": "arn:aws:eks:us-west-2:376248598259:cluster/ferocious-mushroom-1532594698",
			})

			It("should not error", func() {
				err := rs.GetAllOutputs(sampleStack, initial)
				Expect(err).NotTo(HaveOccurred())
			})

			It("should be equal", func() {
				Expect(initial).To(Equal(expected))
			})
		})
	})
})
