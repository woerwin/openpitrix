// Code generated by go-swagger; DO NOT EDIT.

package cluster_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDescribeKeyPairsParams creates a new DescribeKeyPairsParams object
// with the default values initialized.
func NewDescribeKeyPairsParams() *DescribeKeyPairsParams {
	var ()
	return &DescribeKeyPairsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDescribeKeyPairsParamsWithTimeout creates a new DescribeKeyPairsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDescribeKeyPairsParamsWithTimeout(timeout time.Duration) *DescribeKeyPairsParams {
	var ()
	return &DescribeKeyPairsParams{

		timeout: timeout,
	}
}

// NewDescribeKeyPairsParamsWithContext creates a new DescribeKeyPairsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDescribeKeyPairsParamsWithContext(ctx context.Context) *DescribeKeyPairsParams {
	var ()
	return &DescribeKeyPairsParams{

		Context: ctx,
	}
}

// NewDescribeKeyPairsParamsWithHTTPClient creates a new DescribeKeyPairsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDescribeKeyPairsParamsWithHTTPClient(client *http.Client) *DescribeKeyPairsParams {
	var ()
	return &DescribeKeyPairsParams{
		HTTPClient: client,
	}
}

/*DescribeKeyPairsParams contains all the parameters to send to the API endpoint
for the describe key pairs operation typically these are written to a http.Request
*/
type DescribeKeyPairsParams struct {

	/*Description*/
	Description *string
	/*KeyPairID*/
	KeyPairID *string
	/*Limit*/
	Limit *int64
	/*Name*/
	Name *string
	/*Offset*/
	Offset *int64
	/*Owner*/
	Owner []string
	/*PubKey*/
	PubKey *string
	/*SearchWord*/
	SearchWord *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the describe key pairs params
func (o *DescribeKeyPairsParams) WithTimeout(timeout time.Duration) *DescribeKeyPairsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the describe key pairs params
func (o *DescribeKeyPairsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the describe key pairs params
func (o *DescribeKeyPairsParams) WithContext(ctx context.Context) *DescribeKeyPairsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the describe key pairs params
func (o *DescribeKeyPairsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the describe key pairs params
func (o *DescribeKeyPairsParams) WithHTTPClient(client *http.Client) *DescribeKeyPairsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the describe key pairs params
func (o *DescribeKeyPairsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the describe key pairs params
func (o *DescribeKeyPairsParams) WithDescription(description *string) *DescribeKeyPairsParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the describe key pairs params
func (o *DescribeKeyPairsParams) SetDescription(description *string) {
	o.Description = description
}

// WithKeyPairID adds the keyPairID to the describe key pairs params
func (o *DescribeKeyPairsParams) WithKeyPairID(keyPairID *string) *DescribeKeyPairsParams {
	o.SetKeyPairID(keyPairID)
	return o
}

// SetKeyPairID adds the keyPairId to the describe key pairs params
func (o *DescribeKeyPairsParams) SetKeyPairID(keyPairID *string) {
	o.KeyPairID = keyPairID
}

// WithLimit adds the limit to the describe key pairs params
func (o *DescribeKeyPairsParams) WithLimit(limit *int64) *DescribeKeyPairsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the describe key pairs params
func (o *DescribeKeyPairsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the describe key pairs params
func (o *DescribeKeyPairsParams) WithName(name *string) *DescribeKeyPairsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the describe key pairs params
func (o *DescribeKeyPairsParams) SetName(name *string) {
	o.Name = name
}

// WithOffset adds the offset to the describe key pairs params
func (o *DescribeKeyPairsParams) WithOffset(offset *int64) *DescribeKeyPairsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the describe key pairs params
func (o *DescribeKeyPairsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOwner adds the owner to the describe key pairs params
func (o *DescribeKeyPairsParams) WithOwner(owner []string) *DescribeKeyPairsParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the describe key pairs params
func (o *DescribeKeyPairsParams) SetOwner(owner []string) {
	o.Owner = owner
}

// WithPubKey adds the pubKey to the describe key pairs params
func (o *DescribeKeyPairsParams) WithPubKey(pubKey *string) *DescribeKeyPairsParams {
	o.SetPubKey(pubKey)
	return o
}

// SetPubKey adds the pubKey to the describe key pairs params
func (o *DescribeKeyPairsParams) SetPubKey(pubKey *string) {
	o.PubKey = pubKey
}

// WithSearchWord adds the searchWord to the describe key pairs params
func (o *DescribeKeyPairsParams) WithSearchWord(searchWord *string) *DescribeKeyPairsParams {
	o.SetSearchWord(searchWord)
	return o
}

// SetSearchWord adds the searchWord to the describe key pairs params
func (o *DescribeKeyPairsParams) SetSearchWord(searchWord *string) {
	o.SearchWord = searchWord
}

// WriteToRequest writes these params to a swagger request
func (o *DescribeKeyPairsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Description != nil {

		// query param description
		var qrDescription string
		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {
			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}

	}

	if o.KeyPairID != nil {

		// query param key_pair_id
		var qrKeyPairID string
		if o.KeyPairID != nil {
			qrKeyPairID = *o.KeyPairID
		}
		qKeyPairID := qrKeyPairID
		if qKeyPairID != "" {
			if err := r.SetQueryParam("key_pair_id", qKeyPairID); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Name != nil {

		// query param name
		var qrName string
		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {
			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesOwner := o.Owner

	joinedOwner := swag.JoinByFormat(valuesOwner, "multi")
	// query array param owner
	if err := r.SetQueryParam("owner", joinedOwner...); err != nil {
		return err
	}

	if o.PubKey != nil {

		// query param pub_key
		var qrPubKey string
		if o.PubKey != nil {
			qrPubKey = *o.PubKey
		}
		qPubKey := qrPubKey
		if qPubKey != "" {
			if err := r.SetQueryParam("pub_key", qPubKey); err != nil {
				return err
			}
		}

	}

	if o.SearchWord != nil {

		// query param search_word
		var qrSearchWord string
		if o.SearchWord != nil {
			qrSearchWord = *o.SearchWord
		}
		qSearchWord := qrSearchWord
		if qSearchWord != "" {
			if err := r.SetQueryParam("search_word", qSearchWord); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
