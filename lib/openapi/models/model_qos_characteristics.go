/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type QosCharacteristics struct {
	Var5qi int32 `json:"5qi" bson:"5qi"`

	ResourceType QosResourceType `json:"resourceType" bson:"resourceType"`

	PriorityLevel int32 `json:"priorityLevel" bson:"priorityLevel"`

	PacketDelayBudget int32 `json:"packetDelayBudget" bson:"packetDelayBudget"`

	PacketErrorRate string `json:"packetErrorRate" bson:"packetErrorRate"`

	AveragingWindow int32 `json:"averagingWindow,omitempty" bson:"averagingWindow"`

	MaxDataBurstVol int32 `json:"maxDataBurstVol,omitempty" bson:"maxDataBurstVol"`
}