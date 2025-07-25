/*
STACKIT LogMe API

The STACKIT LogMe API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package logme

import (
	"testing"
)

// isEnum

func TestInstanceStatus_UnmarshalJSON(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: `success - possible enum value no. 1`,
			args: args{
				src: []byte(`"active"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 2`,
			args: args{
				src: []byte(`"failed"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 3`,
			args: args{
				src: []byte(`"stopped"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 4`,
			args: args{
				src: []byte(`"creating"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 5`,
			args: args{
				src: []byte(`"deleting"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 6`,
			args: args{
				src: []byte(`"updating"`),
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				src: []byte("\"FOOBAR\""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := InstanceStatus("")
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
