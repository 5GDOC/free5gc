/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

// This data type is defined in the same way as the MediaSubComponent data type, but with the OpenAPI nullable property set to true. Removable attributes marBwDland marBwUl are defined with the corresponding removable data type.
type MediaSubComponentRm struct {
	EthfDescs []EthFlowDescription `json:"ethfDescs,omitempty" bson:"ethfDescs"`
	FNum      int32                `json:"fNum" bson:"fNum"`
	FDescs    []string             `json:"fDescs,omitempty" bson:"fDescs"`
	FStatus   FlowStatus           `json:"fStatus,omitempty" bson:"fStatus"`
	MarBwDl   string               `json:"marBwDl,omitempty" bson:"marBwDl"`
	MarBwUl   string               `json:"marBwUl,omitempty" bson:"marBwUl"`
	// this data type is defined in the same way as the TosTrafficClass data type, but with the OpenAPI nullable property set to true
	TosTrCl   string    `json:"tosTrCl,omitempty" bson:"tosTrCl"`
	FlowUsage FlowUsage `json:"flowUsage,omitempty" bson:"flowUsage"`
}
