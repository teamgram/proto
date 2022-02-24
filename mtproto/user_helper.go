// Copyright 2022 Teamgram Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: teamgramio (teamgram.io@gmail.com)
//

package mtproto

func ToUserIdByInput(userSelfId int64, inputUser *InputUser) int64 {
	switch inputUser.PredicateName {
	case Predicate_inputUserEmpty:
		return 0
	case Predicate_inputUserSelf:
		return userSelfId
	case Predicate_inputUser:
		return inputUser.UserId
	default:
		// log.Errorf("invalid inputUser classID - %v", inputUser)
		return 0
	}
}

func ToUserIdListByInput(userSelfId int64, inputUsers []*InputUser) []int64 {
	idList := make([]int64, 0, len(inputUsers))
	for _, user := range inputUsers {
		id := ToUserIdByInput(userSelfId, user)
		if id > 0 {
			idList = append(idList, id)
		} else {
			// ignore in
		}
	}
	return idList
}

func isUserDeleted(user *User) bool {
	return user == nil || user.PredicateName == Predicate_userEmpty || user.Deleted
}

func isUserContact(user *User) bool {
	return user != nil && (user.Contact || user.MutualContact)
}

func isUserSelf(user *User) bool {
	return user != nil && (user.Contact || user.MutualContact)
}

func GetUserName(user *User) string {
	if user == nil || isUserDeleted(user) {
		return "Deleted Account"
	}

	firstName := user.GetFirstName().GetValue()
	lastName := user.GetLastName().GetValue()

	if firstName == "" && lastName == "" {
		return ""
	} else if firstName == "" {
		return lastName
	} else if lastName == "" {
		return firstName
	} else {
		return firstName + " " + lastName
	}
}
