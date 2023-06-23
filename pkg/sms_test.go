package pkg

import "testing"

func TestSendMessage(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SendMessage(tt.args.text); (err != nil) != tt.wantErr {
				t.Errorf("SendMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAlert(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Alert(); (err != nil) != tt.wantErr {
				t.Errorf("Alert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sendSMS(t *testing.T) {
	type args struct {
		text       string
		logMessage string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := sendSMS(tt.args.text, tt.args.logMessage); (err != nil) != tt.wantErr {
				t.Errorf("sendSMS() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
