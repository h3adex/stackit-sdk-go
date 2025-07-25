/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

import (
	"testing"
)

// isEnum

func TestUpdateScrapeConfigPayloadScheme_UnmarshalJSON(t *testing.T) {
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
				src: []byte(`"http"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 2`,
			args: args{
				src: []byte(`"https"`),
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
			v := UpdateScrapeConfigPayloadScheme("")
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
