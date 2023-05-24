package vm

import "testing"

func Test_GetInstance(t *testing.T) {
	vm, err := GetInstance("develop-339203", "asia-southeast1-b", "lab-zkp-00-a9c34c1")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", vm)
	if vm == nil {
		t.Fatal("not found instance")
	}
	if err = StopInstance(vm, true); err != nil {
		t.Fatal(err)
	}

}
