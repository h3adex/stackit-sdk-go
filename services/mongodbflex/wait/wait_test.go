package wait

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stackitcloud/stackit-sdk-go/core/oapierror"
	"github.com/stackitcloud/stackit-sdk-go/core/utils"
	"github.com/stackitcloud/stackit-sdk-go/services/mongodbflex"
)

// Used for testing instance operations
type apiClientInstanceMocked struct {
	instanceId        string
	instanceState     mongodbflex.InstanceStatus
	instanceIsDeleted bool
	instanceGetFails  bool

	backupId             string
	restoreState         string
	listRestoreJobsFails bool
}

const testRegion = "eu01"

func (a *apiClientInstanceMocked) GetInstanceExecute(_ context.Context, _, _, _ string) (*mongodbflex.InstanceResponse, error) {
	if a.instanceGetFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	if a.instanceIsDeleted {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 404,
		}
	}

	return &mongodbflex.InstanceResponse{
		Item: &mongodbflex.Instance{
			Id:     &a.instanceId,
			Status: &a.instanceState,
		},
	}, nil
}

func (a *apiClientInstanceMocked) ListRestoreJobsExecute(_ context.Context, _, _, _ string) (*mongodbflex.ListRestoreJobsResponse, error) {
	if a.listRestoreJobsFails {
		return nil, &oapierror.GenericOpenAPIError{
			StatusCode: 500,
		}
	}

	return &mongodbflex.ListRestoreJobsResponse{
		Items: &[]mongodbflex.RestoreInstanceStatus{
			{
				Status:   &a.restoreState,
				BackupID: &a.backupId,
			},
		},
	}, nil
}

func TestCreateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc                string
		instanceGetFails    bool
		instanceState       mongodbflex.InstanceStatus
		usersGetErrorStatus int
		wantErr             bool
		wantResp            bool
	}{
		{
			desc:             "create_succeeded",
			instanceGetFails: false,
			instanceState:    mongodbflex.INSTANCESTATUS_READY,
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc:             "create_failed",
			instanceGetFails: false,
			instanceState:    mongodbflex.INSTANCESTATUS_FAILED,
			wantErr:          true,
			wantResp:         true,
		},
		{
			desc:             "create_failed_2",
			instanceGetFails: false,
			instanceState:    InstanceStateEmpty,
			wantErr:          true,
			wantResp:         false,
		},
		{
			desc:             "instance_get_fails",
			instanceGetFails: true,
			wantErr:          true,
			wantResp:         false,
		},
		{
			desc:             "timeout",
			instanceGetFails: false,
			instanceState:    mongodbflex.INSTANCESTATUS_PROCESSING,
			wantErr:          true,
			wantResp:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			instanceId := "foo-bar"

			apiClient := &apiClientInstanceMocked{
				instanceId:       instanceId,
				instanceState:    tt.instanceState,
				instanceGetFails: tt.instanceGetFails,
			}

			var wantRes *mongodbflex.InstanceResponse
			if tt.wantResp {
				wantRes = &mongodbflex.InstanceResponse{
					Item: &mongodbflex.Instance{
						Id:     &instanceId,
						Status: utils.Ptr(tt.instanceState),
					},
				}
			}

			handler := CreateInstanceWaitHandler(context.Background(), apiClient, "", instanceId, testRegion)

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).SetSleepBeforeWait(1 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestUpdateInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		instanceGetFails bool
		instanceState    mongodbflex.InstanceStatus
		wantErr          bool
		wantResp         bool
	}{
		{
			desc:             "update_succeeded",
			instanceGetFails: false,
			instanceState:    mongodbflex.INSTANCESTATUS_READY,
			wantErr:          false,
			wantResp:         true,
		},
		{
			desc:             "update_failed",
			instanceGetFails: false,
			instanceState:    mongodbflex.INSTANCESTATUS_FAILED,
			wantErr:          true,
			wantResp:         true,
		},
		{
			desc:             "update_failed_2",
			instanceGetFails: false,
			instanceState:    InstanceStateEmpty,
			wantErr:          true,
			wantResp:         false,
		},
		{
			desc:             "get_fails",
			instanceGetFails: true,
			wantErr:          true,
			wantResp:         false,
		},
		{
			desc:             "timeout",
			instanceGetFails: false,
			instanceState:    mongodbflex.INSTANCESTATUS_PROCESSING,
			wantErr:          true,
			wantResp:         false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			instanceId := "foo-bar"

			apiClient := &apiClientInstanceMocked{
				instanceId:       instanceId,
				instanceState:    tt.instanceState,
				instanceGetFails: tt.instanceGetFails,
			}

			var wantRes *mongodbflex.InstanceResponse
			if tt.wantResp {
				wantRes = &mongodbflex.InstanceResponse{
					Item: &mongodbflex.Instance{
						Id:     &instanceId,
						Status: utils.Ptr(tt.instanceState),
					},
				}
			}

			handler := UpdateInstanceWaitHandler(context.Background(), apiClient, "", instanceId, testRegion)

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}

func TestDeleteInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc             string
		instanceGetFails bool
		instanceState    mongodbflex.InstanceStatus
		wantErr          bool
	}{
		{
			desc:             "delete_succeeded",
			instanceGetFails: false,
			instanceState:    mongodbflex.INSTANCESTATUS_READY,
			wantErr:          false,
		},
		{
			desc:             "delete_failed",
			instanceGetFails: false,
			instanceState:    mongodbflex.INSTANCESTATUS_FAILED,
			wantErr:          true,
		},
		{
			desc:             "get_fails",
			instanceGetFails: true,
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			instanceId := "foo-bar"

			apiClient := &apiClientInstanceMocked{
				instanceGetFails:  tt.instanceGetFails,
				instanceIsDeleted: tt.instanceState == mongodbflex.INSTANCESTATUS_READY,
				instanceId:        instanceId,
				instanceState:     tt.instanceState,
			}

			handler := DeleteInstanceWaitHandler(context.Background(), apiClient, "", instanceId, testRegion)

			_, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRestoreInstanceWaitHandler(t *testing.T) {
	tests := []struct {
		desc                 string
		listRestoreJobsFails bool
		restoreState         string
		wantErr              bool
		wantResp             bool
	}{
		{
			desc:                 "restore_succeeded",
			listRestoreJobsFails: false,
			restoreState:         RestoreJobFinished,
			wantErr:              false,
			wantResp:             true,
		},
		{
			desc:                 "restore_broken",
			listRestoreJobsFails: false,
			restoreState:         RestoreJobBroken,
			wantErr:              true,
			wantResp:             true,
		},
		{
			desc:                 "restore_killed",
			listRestoreJobsFails: false,
			restoreState:         RestoreJobKilled,
			wantErr:              true,
			wantResp:             true,
		},
		{
			desc:                 "list_fails",
			listRestoreJobsFails: true,
			wantErr:              true,
			wantResp:             false,
		},
		{
			desc:                 "timeout",
			listRestoreJobsFails: false,
			restoreState:         RestoreJobProcessing,
			wantErr:              true,
			wantResp:             false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			backupId := "foo-bar"

			apiClient := &apiClientInstanceMocked{
				backupId:             backupId,
				restoreState:         tt.restoreState,
				listRestoreJobsFails: tt.listRestoreJobsFails,
			}

			var wantRes *mongodbflex.ListRestoreJobsResponse
			if tt.wantResp {
				wantRes = &mongodbflex.ListRestoreJobsResponse{
					Items: &[]mongodbflex.RestoreInstanceStatus{
						{
							Status:   utils.Ptr(tt.restoreState),
							BackupID: &backupId,
						},
					},
				}
			}

			handler := RestoreInstanceWaitHandler(context.Background(), apiClient, "", "", backupId, testRegion)

			gotRes, err := handler.SetTimeout(10 * time.Millisecond).WaitWithContext(context.Background())

			if (err != nil) != tt.wantErr {
				t.Fatalf("handler error = %v, wantErr %v", err, tt.wantErr)
			}
			if !cmp.Equal(gotRes, wantRes) {
				t.Fatalf("handler gotRes = %v, want %v", gotRes, wantRes)
			}
		})
	}
}
