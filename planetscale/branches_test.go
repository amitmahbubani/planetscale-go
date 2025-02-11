package planetscale

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	qt "github.com/frankban/quicktest"
)

const testBranch = "planetscale-go-test-db-branch"

func TestDatabaseBranches_Create(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"id":"planetscale-go-test-db-branch","type":"database_branch","name":"planetscale-go-test-db-branch","notes":"This is a test DB created from the planetscale-go API library","created_at":"2021-01-14T10:19:23.000Z","updated_at":"2021-01-14T10:19:23.000Z", "region": {"slug": "us-west", "display_name": "US West"}}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()
	org := "my-org"
	name := "planetscale-go-test-db"
	notes := "This is a test DB created from the planetscale-go API library"

	db, err := client.DatabaseBranches.Create(ctx, &CreateDatabaseBranchRequest{
		Organization: org,
		Database:     name,
		Region:       "us-west",
		Name:         testBranch,
		Notes:        notes,
	})

	want := &DatabaseBranch{
		Name:  testBranch,
		Notes: notes,
		Region: Region{
			Slug: "us-west",
			Name: "US West",
		},
		CreatedAt: time.Date(2021, time.January, 14, 10, 19, 23, 000, time.UTC),
		UpdatedAt: time.Date(2021, time.January, 14, 10, 19, 23, 000, time.UTC),
	}

	c.Assert(err, qt.IsNil)
	c.Assert(db, qt.DeepEquals, want)
}

func TestDatabaseBranches_List(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"data":[{"id":"planetscale-go-test-db-branch","type":"database_branch","name":"planetscale-go-test-db-branch","notes":"This is a test DB created from the planetscale-go API library","created_at":"2021-01-14T10:19:23.000Z","updated_at":"2021-01-14T10:19:23.000Z"}]}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()
	org := "my-org"
	name := "planetscale-go-test-db"
	notes := "This is a test DB created from the planetscale-go API library"

	db, err := client.DatabaseBranches.List(ctx, &ListDatabaseBranchesRequest{
		Organization: org,
		Database:     name,
	})

	want := []*DatabaseBranch{{
		Name:      testBranch,
		Notes:     notes,
		CreatedAt: time.Date(2021, time.January, 14, 10, 19, 23, 000, time.UTC),
		UpdatedAt: time.Date(2021, time.January, 14, 10, 19, 23, 000, time.UTC),
	}}

	c.Assert(err, qt.IsNil)
	c.Assert(db, qt.DeepEquals, want)
}

func TestDatabaseBranches_ListEmpty(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"data":[]}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()
	org := "my-org"
	name := "planetscale-go-test-db"

	db, err := client.DatabaseBranches.List(ctx, &ListDatabaseBranchesRequest{
		Organization: org,
		Database:     name,
	})

	c.Assert(err, qt.IsNil)
	c.Assert(db, qt.HasLen, 0)
}

func TestDatabaseBranches_Get(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"id":"planetscale-go-test-db-branch","type":"database_branch","name":"planetscale-go-test-db-branch","notes":"This is a test DB created from the planetscale-go API library","created_at":"2021-01-14T10:19:23.000Z","updated_at":"2021-01-14T10:19:23.000Z"}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()
	org := "my-org"
	name := "planetscale-go-test-db"
	notes := "This is a test DB created from the planetscale-go API library"

	db, err := client.DatabaseBranches.Get(ctx, &GetDatabaseBranchRequest{
		Organization: org,
		Database:     name})

	want := &DatabaseBranch{
		Name:      testBranch,
		Notes:     notes,
		CreatedAt: time.Date(2021, time.January, 14, 10, 19, 23, 000, time.UTC),
		UpdatedAt: time.Date(2021, time.January, 14, 10, 19, 23, 000, time.UTC),
	}

	c.Assert(err, qt.IsNil)
	c.Assert(db, qt.DeepEquals, want)
}

func TestDatabaseBranches_Status(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{
    "id": "development",
	"type": "BranchStatus",
	"ready": true,
	"created_at": "2021-01-14T10:19:23.000Z",
	"updated_at": "2021-01-14T10:19:23.000Z",
	"credentials": {
		"mysql_gateway_host": "test-host",
		"mysql_gateway_port": 3306,
		"mysql_gateway_user": "root",
		"mysql_gateway_pass": "password"
	}
}`

		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)

	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()
	org := "my-org"
	name := "planetscale-go-test-db"

	db, err := client.DatabaseBranches.GetStatus(ctx, &GetDatabaseBranchStatusRequest{
		Organization: org,
		Database:     name,
		Branch:       testBranch,
	})

	want := &DatabaseBranchStatus{
		Ready: true,
		Credentials: DatabaseBranchCredentials{
			GatewayHost: "test-host",
			GatewayPort: 3306,
			User:        "root",
			Password:    "password",
		},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(db, qt.DeepEquals, want)
}

func TestBranches_Diff(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"data":[{"name": "foo"}, {"name": "bar"}]}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	diffs, err := client.DatabaseBranches.Diff(ctx, &DiffBranchRequest{
		Organization: "foo",
		Database:     "bar",
		Branch:       "baz",
	})

	want := []*Diff{
		{Name: "foo"},
		{Name: "bar"},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(diffs, qt.DeepEquals, want)
}

func TestBranches_Schema(t *testing.T) {
	c := qt.New(t)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		out := `{"data":[{"name": "foo"}, {"name": "bar"}]}`
		_, err := w.Write([]byte(out))
		c.Assert(err, qt.IsNil)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	schemas, err := client.DatabaseBranches.Schema(ctx, &BranchSchemaRequest{
		Organization: "foo",
		Database:     "bar",
		Branch:       "baz",
	})

	want := []*Diff{
		{Name: "foo"},
		{Name: "bar"},
	}

	c.Assert(err, qt.IsNil)
	c.Assert(schemas, qt.DeepEquals, want)
}

func TestBranches_RefreshSchema(t *testing.T) {
	c := qt.New(t)

	wantURL := "/v1/organizations/my-org/databases/planetscale-go-test-db/branches/planetscale-go-test-db-branch/refresh-schema"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(201)
		c.Assert(r.URL.String(), qt.DeepEquals, wantURL)
	}))

	client, err := NewClient(WithBaseURL(ts.URL))
	c.Assert(err, qt.IsNil)

	ctx := context.Background()

	err = client.DatabaseBranches.RefreshSchema(ctx, &RefreshSchemaRequest{
		Organization: testOrg,
		Database:     testDatabase,
		Branch:       testBranch,
	})
	c.Assert(err, qt.IsNil)
}
