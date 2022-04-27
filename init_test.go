package zconfig

import (
	"context"
	"testing"
)

type initTest struct {
	Initialized bool
}

func (i *initTest) Init(ctx context.Context) error {
	i.Initialized = true
	return nil
}

func TestInitialize(t *testing.T) {
	initMe := new(initTest)
	err := NewProcessor(Initialize).Process(context.Background(), initMe)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !initMe.Initialized {
		t.Fatal("struct was not initialized as expected")
	}
}
