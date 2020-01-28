// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UpdateGlossaryDetails Properties used in glossary update operations.
type UpdateGlossaryDetails struct {

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the glossary.
	Description *string `mandatory:"false" json:"description"`

	// OCID of the user who is the owner of the glossary.
	Owner *string `mandatory:"false" json:"owner"`

	// Status of the approval process workflow for this business glossary.
	WorkflowStatus TermWorkflowStatusEnum `mandatory:"false" json:"workflowStatus,omitempty"`
}

func (m UpdateGlossaryDetails) String() string {
	return common.PointerString(m)
}
