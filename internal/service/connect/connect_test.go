// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connect_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

// Serialized acceptance tests due to Connect account limits (max 2 parallel tests)
func TestAccConnect_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"BotAssociation": {
			acctest.CtBasic:    testAccBotAssociation_basic,
			"disappears":       testAccBotAssociation_disappears,
			"dataSource_basic": testAccBotAssociationDataSource_basic,
		},
		"ContactFlow": {
			acctest.CtBasic:   testAccContactFlow_basic,
			"disappears":      testAccContactFlow_disappears,
			"filename":        testAccContactFlow_filename,
			"dataSource_id":   testAccContactFlowDataSource_contactFlowID,
			"dataSource_name": testAccContactFlowDataSource_name,
		},
		"ContactFlowModule": {
			acctest.CtBasic:   testAccContactFlowModule_basic,
			"disappears":      testAccContactFlowModule_disappears,
			"filename":        testAccContactFlowModule_filename,
			"dataSource_id":   testAccContactFlowModuleDataSource_contactFlowModuleID,
			"dataSource_name": testAccContactFlowModuleDataSource_name,
		},
		"HoursOfOperation": {
			acctest.CtBasic:   testAccHoursOfOperation_basic,
			"disappears":      testAccHoursOfOperation_disappears,
			"tags":            testAccHoursOfOperation_updateTags,
			"config":          testAccHoursOfOperation_updateConfig,
			"dataSource_id":   testAccHoursOfOperationDataSource_hoursOfOperationID,
			"dataSource_name": testAccHoursOfOperationDataSource_name,
		},
		"Instance": {
			acctest.CtBasic:    testAccInstance_basic,
			"directory":        testAccInstance_directory,
			"saml":             testAccInstance_saml,
			"dataSource_basic": testAccInstanceDataSource_basic,
		},
		"InstanceStorageConfig": {
			acctest.CtBasic:                             testAccInstanceStorageConfig_basic,
			"disappears":                                testAccInstanceStorageConfig_disappears,
			"KinesisFirehoseConfig_FirehoseARN":         testAccInstanceStorageConfig_KinesisFirehoseConfig_FirehoseARN,
			"KinesisStreamConfig_StreamARN":             testAccInstanceStorageConfig_KinesisStreamConfig_StreamARN,
			"KinesisVideoStreamConfig_EncryptionConfig": testAccInstanceStorageConfig_KinesisVideoStreamConfig_EncryptionConfig,
			"KinesisVideoStreamConfig_Prefix":           testAccInstanceStorageConfig_KinesisVideoStreamConfig_Prefix,
			"KinesisVideoStreamConfig_Retention":        testAccInstanceStorageConfig_KinesisVideoStreamConfig_Retention,
			"S3Config_BucketName":                       testAccInstanceStorageConfig_S3Config_BucketName,
			"S3Config_BucketPrefix":                     testAccInstanceStorageConfig_S3Config_BucketPrefix,
			"S3Config_EncryptionConfig":                 testAccInstanceStorageConfig_S3Config_EncryptionConfig,
			"dataSource_KinesisFirehoseConfig":          testAccInstanceStorageConfigDataSource_KinesisFirehoseConfig,
			"dataSource_KinesisStreamConfig":            testAccInstanceStorageConfigDataSource_KinesisStreamConfig,
			"dataSource_KinesisVideoStreamConfig":       testAccInstanceStorageConfigDataSource_KinesisVideoStreamConfig,
			"dataSource_S3Config":                       testAccInstanceStorageConfigDataSource_S3Config,
		},
		"LambdaFunctionAssociation": {
			acctest.CtBasic:    testAccLambdaFunctionAssociation_basic,
			"disappears":       testAccLambdaFunctionAssociation_disappears,
			"dataSource_basic": testAccLambdaFunctionAssociationDataSource_basic,
		},
		"PhoneNumber": {
			acctest.CtBasic: testAccPhoneNumber_basic,
			"disappears":    testAccPhoneNumber_disappears,
			"tags":          testAccPhoneNumber_tags,
			"description":   testAccPhoneNumber_description,
			"prefix":        testAccPhoneNumber_prefix,
			"targetARN":     testAccPhoneNumber_targetARN,
		},
		"Prompt": {
			"dataSource_name": testAccPromptDataSource_name,
		},
		"Queue": {
			acctest.CtBasic:        testAccQueue_basic,
			"disappears":           testAccQueue_disappears,
			"tags":                 testAccQueue_updateTags,
			"hoursOfOperationId":   testAccQueue_updateHoursOfOperationId,
			"maxContacts":          testAccQueue_updateMaxContacts,
			"outboundCallerConfig": testAccQueue_updateOutboundCallerConfig,
			"status":               testAccQueue_updateStatus,
			"quickConnectIds":      testAccQueue_updateQuickConnectIds,
			"dataSource_id":        testAccQueueDataSource_queueID,
			"dataSource_name":      testAccQueueDataSource_name,
		},
		"QuickConnect": {
			acctest.CtBasic:   testAccQuickConnect_phoneNumber,
			"disappears":      testAccQuickConnect_disappears,
			"tags":            testAccQuickConnect_updateTags,
			"dataSource_id":   testAccQuickConnectDataSource_id,
			"dataSource_name": testAccQuickConnectDataSource_name,
		},
		"RoutingProfile": {
			acctest.CtBasic:                testAccRoutingProfile_basic,
			"disappears":                   testAccRoutingProfile_disappears,
			"tags":                         testAccRoutingProfile_updateTags,
			"concurrency":                  testAccRoutingProfile_updateConcurrency,
			"defaultOutboundQueue":         testAccRoutingProfile_updateDefaultOutboundQueue,
			"queues":                       testAccRoutingProfile_updateQueues,
			"createQueueBatchAssociations": testAccRoutingProfile_createQueueConfigsBatchedAssociateDisassociate,
			"updateQueueBatchAssociations": testAccRoutingProfile_updateQueueConfigsBatchedAssociateDisassociate,
			"dataSource_id":                testAccRoutingProfileDataSource_routingProfileID,
			"dataSource_name":              testAccRoutingProfileDataSource_name,
		},
		"SecurityProfile": {
			acctest.CtBasic:   testAccSecurityProfile_basic,
			"disappears":      testAccSecurityProfile_disappears,
			"tags":            testAccSecurityProfile_updateTags,
			"permissions":     testAccSecurityProfile_updatePermissions,
			"dataSource_id":   testAccSecurityProfileDataSource_securityProfileID,
			"dataSource_name": testAccSecurityProfileDataSource_name,
		},
		"User": {
			acctest.CtBasic:      testAccUser_basic,
			"disappears":         testAccUser_disappears,
			"tags":               testAccUser_updateTags,
			"hierarchyGroupId":   testAccUser_updateHierarchyGroupId,
			"identityInfo":       testAccUser_updateIdentityInfo,
			"phoneConfig":        testAccUser_updatePhoneConfig,
			"routingProfileId":   testAccUser_updateRoutingProfileId,
			"securityProfileIds": testAccUser_updateSecurityProfileIds,
			"dataSource_id":      testAccUserDataSource_userID,
			"dataSource_name":    testAccUserDataSource_name,
		},
		"UserHierarchyGroup": {
			acctest.CtBasic:   testAccUserHierarchyGroup_basic,
			"disappears":      testAccUserHierarchyGroup_disappears,
			"updateTags":      testAccUserHierarchyGroup_updateTags,
			"parentGroupId":   testAccUserHierarchyGroup_parentGroupId,
			"dataSource_id":   testAccUserHierarchyGroupDataSource_hierarchyGroupID,
			"dataSource_name": testAccUserHierarchyGroupDataSource_name,
		},
		"UserHierarchyStructure": {
			acctest.CtBasic: testAccUserHierarchyStructure_basic,
			"disappears":    testAccUserHierarchyStructure_disappears,
			"dataSource_id": testAccUserHierarchyStructureDataSource_instanceID,
		},
		"Vocabulary": {
			acctest.CtBasic:   testAccVocabulary_basic,
			"disappears":      testAccVocabulary_disappears,
			"tags":            testAccVocabulary_updateTags,
			"dataSource_id":   testAccVocabularyDataSource_vocabularyID,
			"dataSource_name": testAccVocabularyDataSource_name,
		},
	}

	acctest.RunSerialTests2Levels(t, testCases, 0)
}
