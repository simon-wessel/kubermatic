//go:build ee

/*
                  Kubermatic Enterprise Read-Only License
                         Version 1.0 ("KERO-1.0”)
                     Copyright © 2021 Loodse GmbH

   1.	You may only view, read and display for studying purposes the source
      code of the software licensed under this license, and, to the extent
      explicitly provided under this license, the binary code.
   2.	Any use of the software which exceeds the foregoing right, including,
      without limitation, its execution, compilation, copying, modification
      and distribution, is expressly prohibited.
   3.	THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND,
      EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
      MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
      IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
      CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
      TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
      SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

   END OF TERMS AND CONDITIONS
*/

package metering

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"reflect"
	"sigs.k8s.io/controller-runtime/pkg/client"
	fakectrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"
	"testing"
)

func Test_createOrUpdateMeteringToolSecret(t *testing.T) {
	type args struct {
		ctx        context.Context
		seedClient client.Client
		secretData map[string][]byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "scenario 1: asdf",
			args: args{
				ctx: nil,
				seedClient: fakectrlruntimeclient.
					NewClientBuilder().
					WithScheme(scheme.Scheme).
					Build(),
				secretData: map[string][]byte{
					AccessKey: []byte("AKIAIOSFODNN7EXAMPLE"),
					SecretKey: []byte("wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"),
					Bucket:    []byte("metering-handler-example"),
					Endpoint:  []byte("https://s3.amazonaws.com"),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createOrUpdateMeteringToolSecret(tt.args.ctx, tt.args.seedClient, tt.args.secretData); (err != nil) != tt.wantErr {
				t.Errorf("createOrUpdateMeteringToolSecret() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		ctx := context.Background()

		// Check secret
		secretList := &corev1.SecretList{}
		if err := tt.args.seedClient.List(ctx, secretList); err != nil {
			t.Fatalf("Error listing secrets: %v", err)
		}
		if len(secretList.Items) != 1 {
			t.Fatalf("Expected exactly one secret, got %v", len(secretList.Items))
		}
		secret := secretList.Items[0]
		if secret.Name != SecretName {
			t.Fatalf("Expected secret name to be %v, got %v", SecretName, secret.Name)
		}
		if !reflect.DeepEqual(secret.Data, tt.args.secretData) {
			t.Fatal("Secret mismatch")
		}
	}
}