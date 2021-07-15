// Code generated by go-swagger; DO NOT EDIT.

package etcdbackupconfig

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListEtcdBackupConfigParams creates a new ListEtcdBackupConfigParams object
// with the default values initialized.
func NewListEtcdBackupConfigParams() *ListEtcdBackupConfigParams {
	var ()
	return &ListEtcdBackupConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListEtcdBackupConfigParamsWithTimeout creates a new ListEtcdBackupConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListEtcdBackupConfigParamsWithTimeout(timeout time.Duration) *ListEtcdBackupConfigParams {
	var ()
	return &ListEtcdBackupConfigParams{

		timeout: timeout,
	}
}

// NewListEtcdBackupConfigParamsWithContext creates a new ListEtcdBackupConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewListEtcdBackupConfigParamsWithContext(ctx context.Context) *ListEtcdBackupConfigParams {
	var ()
	return &ListEtcdBackupConfigParams{

		Context: ctx,
	}
}

// NewListEtcdBackupConfigParamsWithHTTPClient creates a new ListEtcdBackupConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListEtcdBackupConfigParamsWithHTTPClient(client *http.Client) *ListEtcdBackupConfigParams {
	var ()
	return &ListEtcdBackupConfigParams{
		HTTPClient: client,
	}
}

/*ListEtcdBackupConfigParams contains all the parameters to send to the API endpoint
for the list etcd backup config operation typically these are written to a http.Request
*/
type ListEtcdBackupConfigParams struct {

	/*ClusterID*/
	ClusterID string
	/*ProjectID*/
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list etcd backup config params
func (o *ListEtcdBackupConfigParams) WithTimeout(timeout time.Duration) *ListEtcdBackupConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list etcd backup config params
func (o *ListEtcdBackupConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list etcd backup config params
func (o *ListEtcdBackupConfigParams) WithContext(ctx context.Context) *ListEtcdBackupConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list etcd backup config params
func (o *ListEtcdBackupConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list etcd backup config params
func (o *ListEtcdBackupConfigParams) WithHTTPClient(client *http.Client) *ListEtcdBackupConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list etcd backup config params
func (o *ListEtcdBackupConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the list etcd backup config params
func (o *ListEtcdBackupConfigParams) WithClusterID(clusterID string) *ListEtcdBackupConfigParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the list etcd backup config params
func (o *ListEtcdBackupConfigParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the list etcd backup config params
func (o *ListEtcdBackupConfigParams) WithProjectID(projectID string) *ListEtcdBackupConfigParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list etcd backup config params
func (o *ListEtcdBackupConfigParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListEtcdBackupConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
