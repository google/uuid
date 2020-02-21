package uuid

import (
	"bytes"
	"testing"
)

func TestUnmarshalGQL(t *testing.T) {
	tests := []struct {
		name    string
		uuid    UUID
		args    interface{}
		wantErr bool
	}{
		{
			name:    "Without quite",
			uuid:    testUUID,
			args:    `f47ac10b-58cc-0372-8567-0e02b2c3d479`,
			wantErr: false,
		},
		{
			name:    "With quite",
			uuid:    testUUID,
			args:    `"f47ac10b-58cc-0372-8567-0e02b2c3d479"`,
			wantErr: false,
		},
		{
			name:    "Bad size uuid with quite",
			uuid:    testUUID,
			args:    `"47ac10b-58cc-0372-8567-0e02b2c3d479"`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.uuid.UnmarshalGQL(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalGQL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMarshalGQL(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	testUUID.MarshalGQL(buf)
	got := buf.String()
	want := `"f47ac10b-58cc-0372-8567-0e02b2c3d479"`
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
