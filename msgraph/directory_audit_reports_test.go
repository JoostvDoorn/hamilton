package msgraph_test

import (
	"testing"

	"github.com/manicminer/hamilton/auth"
	"github.com/manicminer/hamilton/internal/test"
	"github.com/manicminer/hamilton/msgraph"
)

type DirectoryAuditReportsClientTest struct {
	connection   *test.Connection
	client       *msgraph.DirectoryAuditReportsClient
	randomString string
}

func TestDirectoryAuditReportsTest(t *testing.T) {
	c := DirectoryAuditReportsClientTest{
		connection:   test.NewConnection(auth.MsGraph, auth.TokenVersion2),
		randomString: test.RandomString(),
	}
	c.client = msgraph.NewDirectoryAuditReportsClient(c.connection.AuthConfig.TenantID)
	c.client.BaseClient.Authorizer = c.connection.Authorizer

	auditLogs := testDirectoryAuditReports_List(t, c)
	testDirectoryAuditReports_Get(t, c, *(*auditLogs)[0].Id)
}

func testDirectoryAuditReports_List(t *testing.T, c DirectoryAuditReportsClientTest) (dirLogs *[]msgraph.AuditLog) {
	dirLogs, status, err := c.client.List(c.connection.Context, "")

	if status < 200 || status >= 300 {
		t.Fatalf("DirectoryAuditReportsClient.List(): invalid status: %d", status)
	}

	if err != nil {
		t.Fatalf("DirectoryAuditReportsClient.List(): %v", err)
	}

	if dirLogs == nil {
		t.Fatal("DirectoryAuditReportsClient.List():logs was nil")
	}
	return
}

func testDirectoryAuditReports_Get(t *testing.T, c DirectoryAuditReportsClientTest, id string) (dirLog *msgraph.AuditLog) {
	dirLog, status, err := c.client.Get(c.connection.Context, id)
	if err != nil {
		t.Fatalf("DirectoryAuditReportsClient.Get(): %v", err)
	}
	if status < 200 || status >= 300 {
		t.Fatalf("DirectoryAuditReportsClient.Get(): invalid status: %d", status)
	}
	if dirLog == nil {
		t.Fatal("DirectoryAuditReportsClient.Get(): domain was nil")
	}
	return
}
