package services

// import (
// 	"testing"
// )

// func TestProccessMessage(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		message string
// 		wantErr bool
// 	}{
// 		{
// 			name:    "valid message",
// 			message: "100 food",
// 			wantErr: false,
// 		},
// 		{
// 			name:    "empty message",
// 			message: "",
// 			wantErr: true,
// 		},
// 		{
// 			name:    "message with leading/trailing spaces",
// 			message: "  50 restaurant  ",
// 			wantErr: false,
// 		},
// 		{
// 			name:    "only one word",
// 			message: "food",
// 			wantErr: true,
// 		},
// 		{
// 			name:    "first word is not a number",
// 			message: "abc food",
// 			wantErr: true,
// 		},
// 		{
// 			name:    "float number",
// 			message: "25.50 coffee",
// 			wantErr: false,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := ProccessMessage(tt.message)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("ProccessMessage() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
