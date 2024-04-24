package middlewares

import "testing"

func Test_verifySignature(t *testing.T) {
	type args struct {
		payloadBody     string
		secretToken     string
		signatureHeader string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				payloadBody:     "Hello, World!",
				secretToken:     "It's a Secret to Everybody",
				signatureHeader: "sha256=757107ea0eb2509fc211221cce984b8a37570b6d7586c22c46f4379c8b043e17",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				payloadBody:     "Hello, World!",
				secretToken:     "It's a Secret to Everybody1",
				signatureHeader: "sha256=757107ea0eb2509fc211221cce984b8a37570b6d7586c22c46f4379c8b043e17",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := verifySignature(tt.args.payloadBody, tt.args.secretToken, tt.args.signatureHeader); (err != nil) != tt.wantErr {
				t.Errorf("verifySignature() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
