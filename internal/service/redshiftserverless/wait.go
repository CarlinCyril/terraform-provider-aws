package redshiftserverless

import (
	"time"

	"github.com/aws/aws-sdk-go/service/redshiftserverless"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func waitNamespaceDeleted(conn *redshiftserverless.RedshiftServerless, name string) (*redshiftserverless.Namespace, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			redshiftserverless.NamespaceStatusDeleting,
		},
		Target:  []string{},
		Refresh: statusNamespace(conn, name),
		Timeout: 10 * time.Minute,
	}

	outputRaw, err := stateConf.WaitForState()

	if output, ok := outputRaw.(*redshiftserverless.Namespace); ok {
		return output, err
	}

	return nil, err
}

func waitNamespaceUpdated(conn *redshiftserverless.RedshiftServerless, name string) (*redshiftserverless.Namespace, error) {
	stateConf := &resource.StateChangeConf{
		Pending: []string{
			redshiftserverless.NamespaceStatusModifying,
		},
		Target: []string{
			redshiftserverless.NamespaceStatusAvailable,
		},
		Refresh: statusNamespace(conn, name),
		Timeout: 10 * time.Minute,
	}

	outputRaw, err := stateConf.WaitForState()

	if output, ok := outputRaw.(*redshiftserverless.Namespace); ok {
		return output, err
	}

	return nil, err
}
