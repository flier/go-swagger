package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFindPetsByTagsParams creates a new FindPetsByTagsParams object
// with the default values initialized.
func NewFindPetsByTagsParams() FindPetsByTagsParams {
	var ()
	return FindPetsByTagsParams{}
}

// FindPetsByTagsParams contains all the bound params for the find pets by tags operation
// typically these are obtained from a http.Request
//
// swagger:parameters findPetsByTags
type FindPetsByTagsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*Tags to filter by
	  In: query
	  Collection Format: multi
	*/
	Tags []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *FindPetsByTagsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qTags, qhkTags, _ := qs.GetOK("tags")
	if err := o.bindTags(qTags, qhkTags, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindPetsByTagsParams) bindTags(rawData []string, hasKey bool, formats strfmt.Registry) error {

	raw := rawData
	size := len(raw)

	if size == 0 {
		return nil
	}

	ic := raw
	isz := size
	var ir []string
	iValidateElement := func(i int, tagsI string) *errors.Validation {

		return nil
	}

	for i := 0; i < isz; i++ {

		if err := iValidateElement(i, ic[i]); err != nil {
			return err
		}
		ir = append(ir, ic[i])
	}

	o.Tags = ir

	return nil
}
