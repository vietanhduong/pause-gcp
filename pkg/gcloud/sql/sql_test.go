package sql

import "testing"

func TestGetInstance(t *testing.T) {
	instance, err := GetInstance("develop-339203", "ota")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", instance)
	if instance == nil {
		t.Fatal("instance not found")
	}

	if err = StartInstance(instance); err != nil {
		t.Fatal(err)
	}
}
