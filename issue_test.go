// (c) 2023 Alex Suslov
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of
// this software and associated documentation files (the "Software"), to deal in
// the Software without restriction, including without limitation the rights to
// use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is furnished to do
// so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package v2

import (
	"testing"

	gojira "github.com/andygrunwald/go-jira"
)

func TestIssueService_Issue(t *testing.T) {
	sd := SD{}
	is := sd.IssueService()

	type args struct {
		issueIdOrKey string
		result       interface{}
	}
	tests := []struct {
		name string
		//fields  fields
		args    args
		wantErr bool
	}{
		{
			"get issue",
			args{"BSI-2184", &gojira.Issue{}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			I := is
			if err := I.Issue(tt.args.issueIdOrKey, tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("Issue() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
