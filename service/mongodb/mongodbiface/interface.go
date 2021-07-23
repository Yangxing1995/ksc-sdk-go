// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mongodbiface provides an interface to enable mocking the mongodb service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mongodbiface

import (
	"github.com/KscSDK/ksc-sdk-go/service/mongodb"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// MongodbAPI provides an interface to enable mocking the
// mongodb.Mongodb service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // mongodb.
//    func myFunc(svc mongodbiface.MongodbAPI) bool {
//        // Make svc.AddClusterNode request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := mongodb.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMongodbClient struct {
//        mongodbiface.MongodbAPI
//    }
//    func (m *mockMongodbClient) AddClusterNode(input *map[string]interface{}) (*map[string]interface{}, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMongodbClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type MongodbAPI interface {
	AddClusterNode(*map[string]interface{}) (*map[string]interface{}, error)
	AddClusterNodeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddClusterNodeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddSecondaryInstance(*map[string]interface{}) (*map[string]interface{}, error)
	AddSecondaryInstanceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddSecondaryInstanceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	AddSecurityGroupRule(*map[string]interface{}) (*map[string]interface{}, error)
	AddSecurityGroupRuleWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	AddSecurityGroupRuleRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateMongoDBInstance(*map[string]interface{}) (*map[string]interface{}, error)
	CreateMongoDBInstanceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateMongoDBInstanceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateMongoDBShardInstance(*map[string]interface{}) (*map[string]interface{}, error)
	CreateMongoDBShardInstanceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateMongoDBShardInstanceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	CreateMongoDBSnapshot(*map[string]interface{}) (*map[string]interface{}, error)
	CreateMongoDBSnapshotWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateMongoDBSnapshotRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteClusterNode(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteClusterNodeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteClusterNodeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteMongoDBInstance(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteMongoDBInstanceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteMongoDBInstanceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteMongoDBSnapshot(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteMongoDBSnapshotWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteMongoDBSnapshotRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DeleteSecurityGroupRules(*map[string]interface{}) (*map[string]interface{}, error)
	DeleteSecurityGroupRulesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DeleteSecurityGroupRulesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeMongoDBInstance(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeMongoDBInstanceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeMongoDBInstanceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeMongoDBInstanceNode(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeMongoDBInstanceNodeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeMongoDBInstanceNodeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeMongoDBInstances(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeMongoDBInstancesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeMongoDBInstancesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeMongoDBShardNode(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeMongoDBShardNodeWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeMongoDBShardNodeRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeMongoDBSnapshot(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeMongoDBSnapshotWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeMongoDBSnapshotRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeRegions(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeRegionsWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeRegionsRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ListSecurityGroupRules(*map[string]interface{}) (*map[string]interface{}, error)
	ListSecurityGroupRulesWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ListSecurityGroupRulesRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RenameMongoDBInstance(*map[string]interface{}) (*map[string]interface{}, error)
	RenameMongoDBInstanceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RenameMongoDBInstanceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RenameMongoDBSnapshot(*map[string]interface{}) (*map[string]interface{}, error)
	RenameMongoDBSnapshotWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RenameMongoDBSnapshotRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	ResetPasswordMongoDBInstance(*map[string]interface{}) (*map[string]interface{}, error)
	ResetPasswordMongoDBInstanceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	ResetPasswordMongoDBInstanceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	RestartMongoDBInstance(*map[string]interface{}) (*map[string]interface{}, error)
	RestartMongoDBInstanceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	RestartMongoDBInstanceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	SetMongoDBTimingSnapshot(*map[string]interface{}) (*map[string]interface{}, error)
	SetMongoDBTimingSnapshotWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	SetMongoDBTimingSnapshotRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateMongoDBInstance(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateMongoDBInstanceWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateMongoDBInstanceRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	UpdateMongoDBInstanceCluster(*map[string]interface{}) (*map[string]interface{}, error)
	UpdateMongoDBInstanceClusterWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	UpdateMongoDBInstanceClusterRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ MongodbAPI = (*mongodb.Mongodb)(nil)
