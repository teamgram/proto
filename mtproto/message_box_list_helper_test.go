// Copyright 2024 Teamgram Authors
//  All rights reserved.
//
// Author: Benqi (wubenqi@gmail.com)
//

package mtproto

import (
	"testing"
)

func TestMakeBoxListByBoxList(t *testing.T) {
	// Create a slice of MessageBox instances
	boxList := []*MessageBox{
		&MessageBox{ /* initialize fields here */ },
		&MessageBox{ /* initialize fields here */ },
		// Add more MessageBox instances as needed
	}

	// Call the function under test
	result := MakeBoxListByBoxList(boxList)

	// Check the result
	if result == nil {
		t.Errorf("MakeBoxListByBoxList returned nil")
	}

	if len(result.BoxList) != len(boxList) {
		t.Errorf("MakeBoxListByBoxList returned wrong length, got: %d, want: %d", len(result.BoxList), len(boxList))
	}

	// Add more checks as needed, for example to verify the contents of the returned MessageBoxListHelper
}
