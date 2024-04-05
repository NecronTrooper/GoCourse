package main

import (
	"context"
	"testing"
)

func TestQueryDatabase(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QueryDatabase(tt.args.ctx)
		})
	}
}
