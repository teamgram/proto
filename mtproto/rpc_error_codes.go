// Copyright 2024 Teamgram Authors
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

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// https://core.telegram.org/api/errors
/*
Error handling
There will be errors when working with the API, and they must be correctly handled on the client.

An error is characterized by several parameters:

Error Code
Numerical value similar to HTTP status. Contains information on the type of error that occurred: for example, a data input error, privacy error, or server error. This is a required parameter.

Error Type
A string literal in the form of /[A-Z_0-9]+/, which summarizes the problem. For example, AUTH_KEY_UNREGISTERED. This is an optional parameter.

Error Constructors
There should be a way to handle errors that are returned in rpc_error constructors.

Below is a list of error codes and their meanings:
*/

const (
	/*
		303 SEE_OTHER
		The request must be repeated, but directed to a different data center.

		Examples of Errors:
			FILE_MIGRATE_X: the file to be accessed is currently stored in a different data center.
			PHONE_MIGRATE_X: the phone number a user is trying to use for authorization is associated with a different data center.
			NETWORK_MIGRATE_X: the source IP address is associated with a different data center (for registration)
			USER_MIGRATE_X: the user whose identity is being used to execute queries is associated with a different data center (for registration)

		In all these cases, the error description’s string literal contains the number of the data center (instead of the X) to which the repeated query must be sent.
		More information about redirects between data centers »
	*/
	ErrSeeOther codes.Code = 303

	/*
		400 BAD_REQUEST
		The query contains errors. In the event that a request was created using a form and contains user generated data, the user should be notified that the data must be corrected before the query is repeated.

		Examples of Errors:
			FIRSTNAME_INVALID: The first name is invalid
			LASTNAME_INVALID: The last name is invalid
			PHONE_NUMBER_INVALID: The phone number is invalid
			PHONE_CODE_HASH_EMPTY: phone_code_hash is missing
			PHONE_CODE_EMPTY: phone_code is missing
			PHONE_CODE_EXPIRED: The confirmation code has expired
			API_ID_INVALID: The api_id/api_hash combination is invalid
			PHONE_NUMBER_OCCUPIED: The phone number is already in use
			PHONE_NUMBER_UNOCCUPIED: The phone number is not yet being used
			USERS_TOO_FEW: Not enough users (to create a chat, for example)
			USERS_TOO_MUCH: The maximum number of users has been exceeded (to create a chat, for example)
			TYPE_CONSTRUCTOR_INVALID: The type constructor is invalid
			FILE_PART_INVALID: The file part number is invalid
			FILE_PARTS_INVALID: The number of file parts is invalid
			FILE_PART_Х_MISSING: Part X (where X is a number) of the file is missing from storage
			MD5_CHECKSUM_INVALID: The MD5 checksums do not match
			PHOTO_INVALID_DIMENSIONS: The photo dimensions are invalid
			FIELD_NAME_INVALID: The field with the name FIELD_NAME is invalid
			FIELD_NAME_EMPTY: The field with the name FIELD_NAME is missing
			MSG_WAIT_FAILED: A request that must be completed before processing the current request returned an error
			MSG_WAIT_TIMEOUT: A request that must be completed before processing the current request didn't finish processing yet
	*/
	ErrBadRequest codes.Code = 400

	/*
		401 UNAUTHORIZED
		There was an unauthorized attempt to use functionality available only to authorized users.

		Examples of Errors:
			AUTH_KEY_UNREGISTERED: The key is not registered in the system
			AUTH_KEY_INVALID: The key is invalid
			USER_DEACTIVATED: The user has been deleted/deactivated
			SESSION_REVOKED: The authorization has been invalidated, because of the user terminating all sessions
			SESSION_EXPIRED: The authorization has expired
			AUTH_KEY_PERM_EMPTY: The method is unavailble for temporary authorization key, not bound to permanent
	*/
	ErrUnauthorized codes.Code = 401

	/*
		403 FORBIDDEN
		Privacy violation. For example, an attempt to write a message to someone who has blacklisted the current user.
	*/
	ErrForbidden codes.Code = 403

	/*
		404 NOT_FOUND
		An attempt to invoke a non-existent object, such as a method.
	*/
	ErrNotFound codes.Code = 404

	/*
		406 NOT_ACCEPTABLE

		Similar to 400 BAD_REQUEST, but the app must display the error to the user a bit differently.
		Do not display any visible error to the user when receiving the rpc_error constructor: instead, wait for an updateServiceNotification update, and handle it as usual.
		Basically, an updateServiceNotification popup update will be emitted independently (ie NOT as an Updates constructor inside rpc_result but as a normal update) immediately after emission of a 406 rpc_error: the update will contain the actual localized error message to show to the user with a UI popup.

		An exception to this is the AUTH_KEY_DUPLICATED error, which is only emitted if any of the non-media DC detects that an authorized session is sending requests in parallel from two separate TCP connections, from the same or different IP addresses.
		Note that parallel connections are still allowed and actually recommended for media DCs.
		Also note that by session we mean a logged-in session identified by an authorization constructor, fetchable using account.getAuthorizations, not an MTProto session.

		If the client receives an AUTH_KEY_DUPLICATED error, the session was already invalidated by the server and the user must generate a new auth key and login again.
	*/
	ErrNotAcceptable codes.Code = 406

	/*
		420 FLOOD
		The maximum allowed number of attempts to invoke the given method with the given input parameters has been exceeded. For example, in an attempt to request a large number of text messages (SMS) for the same phone number.

		Error Example:
			FLOOD_WAIT_X: A wait of X seconds is required (where X is a number)
	*/
	ErrFlood codes.Code = 420

	/*
		500 INTERNAL
		An internal server error occurred while a request was being processed; for example, there was a disruption while accessing a database or file storage.

		If a client receives a 500 error, or you believe this error should not have occurred, please collect as much information as possible about the query and error and send it to the developers.
	*/
	ErrInternal codes.Code = 500

	/*
		Other Error Codes
		If a server returns an error with a code other than the ones listed above, it may be considered the same as a 500 error and treated as an internal server error.
	*/

	// ErrTimeOut503
	// TODO(@benqi): ???, if ecode < 0, panic
	// | -503 | Timeout | Timeout while fetching data |
	ErrTimeOut503 codes.Code = 5030000

	// ErrNotReturnClient ErrNotReturnClient
	ErrNotReturnClient codes.Code = 700

	// ErrDatabase db error
	ErrDatabase codes.Code = 600
)

// NewErrFileMigrateX
// 303 SEE_OTHER
//
// | 303 | NETWORK_MIGRATE_X | Repeat the query to data-center X |
// | 303 | PHONE_MIGRATE_X | Repeat the query to data-center X |
//

// NewErrFileMigrateX
// FILE_MIGRATE_X: the file to be accessed is currently stored in a different data center.
func NewErrFileMigrateX(dc int32) error {
	return status.Errorf(ErrSeeOther, "FILE_MIGRATE_%d", dc)
}

// NewErrPhoneMigrateX
// | 303 | PHONE_MIGRATE_X | Repeat the query to data-center X. |
func NewErrPhoneMigrateX(dc int32) error {
	return status.Errorf(ErrSeeOther, "PHONE_MIGRATE_%d", dc)
}

// NewErrNetworkMigrateX
// | 303 | NETWORK_MIGRATE_X | Repeat the query to data-center X. |
// NETWORK_MIGRATE_X: the source IP address is associated with a different data center (for registration)
func NewErrNetworkMigrateX(dc int32) error {
	return status.Errorf(ErrSeeOther, "NETWORK_MIGRATE_%d", dc)
}

// NewErrUserMigrateX
// USER_MIGRATE_X: the user whose identity is being used to execute queries is associated with a different data center (for registration)
func NewErrUserMigrateX(dc int32) error {
	return status.Errorf(ErrSeeOther, "USER_MIGRATE_%d", dc)
}

// NewErrFloodWaitX 420 FLOOD
//
// FLOOD_WAIT_X: A wait of X seconds is required (where X is a number)
func NewErrFloodWaitX(second int32) error {
	return status.Errorf(ErrFlood, "FLOOD_WAIT_%d", second)
}

// NewErrSlowModeWaitX
// | 420 | SLOWMODE_WAIT_X | Slowmode is enabled in this chat: you must wait for the specified number of seconds before sending another message to the chat. |
func NewErrSlowModeWaitX(second int32) error {
	return status.Errorf(ErrFlood, "SLOWMODE_WAIT_%d", second)
}

// NewErrTakeoutInitDelayX
// | 420 | TAKEOUT_INIT_DELAY_X | Wait X seconds before initing takeout. |
func NewErrTakeoutInitDelayX(second int32) error {
	return status.Errorf(ErrFlood, "TAKEOUT_INIT_DELAY_%d", second)
}

// NewErr2faConfirmWaitX
// | 420 | 2FA_CONFIRM_WAIT_X | Since this account is active and protected by a 2FA password, we will delete it in 1 week for security purposes. You can cancel this process at any time, you'll be able to reset your account in X seconds. |
func NewErr2faConfirmWaitX(second int32) error {
	return status.Errorf(ErrFlood, "2FA_CONFIRM_WAIT_%d", second)
}

// 420 ErrFlood codes.Code = 420
var (
	// ErrP0nyFloodwait
	// | 420 | P0NY_FLOODWAIT |   |
	ErrP0nyFloodwait = status.Error(ErrFlood, "P0NY_FLOODWAIT")

	//// ErrSlowmodeWait_%d
	//// | 420 | SLOWMODE_WAIT_%d | Slowmode is enabled in this chat: wait %d seconds before sending another message to this chat. |
	//ErrSlowmodeWait_%d = status.Error(420, "SLOWMODE_WAIT_%d")
	//
	//// ErrTakeoutInitDelay_%d
	//// | 420 | TAKEOUT_INIT_DELAY_%d | Wait %d seconds before initializing takeout. |
	//ErrTakeoutInitDelay_%d = status.Error(420, "TAKEOUT_INIT_DELAY_%d")
	//
	//// Err2faConfirmWait_%d
	//// | 420 | 2FA_CONFIRM_WAIT_%d | Since this account is active and protected by a 2FA password, we will delete it in 1 week for security purposes. You can cancel this process at any time, you'll be able to reset your account in %d seconds. |
	//Err2faConfirmWait_%d = status.Error(420, "2FA_CONFIRM_WAIT_%d")
)

// 400 BAD_REQUEST

// NewErrFileReferenceX
// | 400 | FILE_REFERENCE_* | The file reference expired, it must be refreshed. |
func NewErrFileReferenceX(second int32) error {
	return status.Errorf(ErrBadRequest, "FILE_REFERENCE_%d", second)
}

// NewErrPasswordTooFreshX
// PASSWORD_TOO_FRESH_%d
// | 400 | PASSWORD_TOO_FRESH_%d | The password was modified less than 24 hours ago, try again in %d seconds. |
func NewErrPasswordTooFreshX(second int32) error {
	return status.Errorf(ErrBadRequest, "PASSWORD_TOO_FRESH_%d", second)
}

// NewErrEmailUnconfirmedX
// ErrEmailUnconfirmed_%d
// | 400 | EMAIL_UNCONFIRMED_%d | The provided email isn't confirmed, %d is the length of the verification code that was just sent to the email: use [account.verifyEmail](https://core.telegram.org/method/account.verifyEmail) to enter the received verification code and enable the recovery email. |
func NewErrEmailUnconfirmedX(len int32) error {
	return status.Errorf(ErrBadRequest, "EMAIL_UNCONFIRMED_%d", len)
}

// NewErrSessionTooFreshX
// ErrSessionTooFresh_%d
// | 400 | SESSION_TOO_FRESH_%d | This session was created less than 24 hours ago, try again in %d seconds. |
func NewErrSessionTooFreshX(second int32) error {
	return status.Errorf(ErrBadRequest, "SESSION_TOO_FRESH_%d", second)
}

// NewFilePartXMissing
// FILE_PART_Х_MISSING: Part X (where X is a number) of the file is missing from storage. Try repeating the method call to resave the part.
func NewFilePartXMissing(x int32) error {
	return status.Errorf(ErrBadRequest, "FILE_PART_%d_MISSING", x)
}

// NewEmailUnconfirmedX
// 400	EMAIL_UNCONFIRMED_X	The provided email isn't confirmed, X is the length of the verification code that was just sent to the email: use account.verifyEmail to enter the received verification code and enable the recovery email.
func NewEmailUnconfirmedX(x int) error {
	return status.Errorf(ErrBadRequest, "EMAIL_UNCONFIRMED_%d", x)
}

var (
	// ErrMethodNotImpl
	// @benqi Add By NebulaChat, not impl the method
	// METHOD_NOT_IMPL: The method not impl
	ErrMethodNotImpl = status.Error(ErrBadRequest, "METHOD_NOT_IMPL")

	// ErrGroupCallParticipantInvalid GROUPCALL_PARTICIPANT_INVALID
	ErrGroupCallParticipantInvalid = status.Error(ErrBadRequest, "GROUPCALL_PARTICIPANT_INVALID")

	// ErrCheckSumInvalid
	// MD5_CHECKSUM_INVALID: The file’s checksum did not match the md5_checksum parameter
	ErrCheckSumInvalid = status.Error(ErrBadRequest, "MD5_CHECKSUM_INVALID")

	// ErrGroupCallInvalid GROUPCALL_INVALID
	ErrGroupCallInvalid = status.Error(ErrBadRequest, "GROUPCALL_INVALID")

	// ErrThemeSlugOccupied
	// THEME_SLUG_OCCUPIED
	ErrThemeSlugOccupied = status.Error(ErrBadRequest, "THEME_SLUG_OCCUPIED")

	// ErrShortnameOccupyFailed
	// | 400 | SHORTNAME_OCCUPY_FAILED | An internal error occurred. |
	ErrShortnameOccupyFailed = status.Error(ErrBadRequest, "SHORTNAME_OCCUPY_FAILED")

	// ErrThemeSlugInvalid
	// THEME_SLUG_INVALID 400 The input theme slug was not valid
	ErrThemeSlugInvalid = status.Error(ErrBadRequest, "THEME_SLUG_INVALID")

	// ErrApiServerNeeded
	// | 400 | API_SERVER_NEEDED | This method be used by api server |
	ErrApiServerNeeded = status.Error(ErrBadRequest, "API_SERVER_NEEDED")
	// ErrInputConstructorInvalid
	// 400	INPUT_CONSTRUCTOR_INVALID	The provided constructor is invalid
	ErrInputConstructorInvalid = status.Error(ErrBadRequest, "INPUT_CONSTRUCTOR_INVALID")

	// ErrEncryptedChatIdInvalid
	// | 400 | ENCRYPTED_CHAT_ID_INVALID | The encrypted chat id is invalid |
	ErrEncryptedChatIdInvalid = status.Error(ErrBadRequest, "ENCRYPTED_CHAT_ID_INVALID")

	// ErrAuthTokenAccepted
	// 400 - AUTH_TOKEN_ALREADY_ACCEPTED, the authorization token was already used
	ErrAuthTokenAccepted = status.Error(ErrBadRequest, "AUTH_TOKEN_ALREADY_ACCEPTED")

	// ErrBotMethodInvalid
	// | 400 | BOT_METHOD_INVALID | This method can't be used by a bot |
	ErrBotMethodInvalid = status.Error(ErrBadRequest, "BOT_METHOD_INVALID")

	// ErrInputRequestInvalid
	// INPUT_REQUEST_INVALID: The method not impl
	ErrInputRequestInvalid = status.Error(ErrBadRequest, "INPUT_REQUEST_INVALID")

	// ErrEnterpriseIsBlocked ErrEnterpriseIsBlocked
	ErrEnterpriseIsBlocked = status.Error(ErrBadRequest, "ERR_ENTERPRISE_IS_BLOCKED")

	// ErrErrBadRequest ErrErrBadRequest
	ErrErrBadRequest = status.Error(ErrBadRequest, "ERR_BAD_REQUEST")

	// ErrTimeTooBig
	// 400	TIME_TOO_BIG
	ErrTimeTooBig = status.Error(ErrBadRequest, "TIME_TOO_BIG")

	// ErrFilterIdInvalid
	// | 400 | FILTER_ID_INVALID | The specified filter ID is invalid. |
	ErrFilterIdInvalid = status.Error(ErrBadRequest, "FILTER_ID_INVALID")

	// ErrLangCodeInvalid
	// | 400 | LANG_CODE_INVALID | The specified language code is invalid. |
	ErrLangCodeInvalid = status.Error(ErrBadRequest, "LANG_CODE_INVALID")

	// ErrReactionsTooMany
	// | 400 | REACTIONS_TOO_MANY | The message already has exactly `reactions_uniq_max` reaction emojis, you can't react with a new emoji, see [the docs for more info &raquo;](/api/config#client-configuration). |
	ErrReactionsTooMany = status.Error(ErrBadRequest, "REACTIONS_TOO_MANY")

	// ErrScheduleTooMuch
	// | 400 | SCHEDULE_TOO_MUCH | There are too many scheduled messages. |
	ErrScheduleTooMuch = status.Error(ErrBadRequest, "SCHEDULE_TOO_MUCH")

	// ErrUserBlocked
	// | 400 | USER_BLOCKED | User blocked. |
	ErrUserBlocked = status.Error(ErrBadRequest, "USER_BLOCKED")

	// ErrFileTokenInvalid
	// | 400 | FILE_TOKEN_INVALID | The specified file token is invalid. |
	ErrFileTokenInvalid = status.Error(ErrBadRequest, "FILE_TOKEN_INVALID")

	// ErrMessageNotModified
	// | 400 | MESSAGE_NOT_MODIFIED | The provided message data is identical to the previous message data, the message wasn't modified. |
	ErrMessageNotModified = status.Error(ErrBadRequest, "MESSAGE_NOT_MODIFIED")

	// ErrPollOptionDuplicate
	// | 400 | POLL_OPTION_DUPLICATE | Duplicate poll options provided. |
	ErrPollOptionDuplicate = status.Error(ErrBadRequest, "POLL_OPTION_DUPLICATE")

	// ErrWebdocumentUrlInvalid
	// | 400 | WEBDOCUMENT_URL_INVALID | The specified webdocument URL is invalid. |
	ErrWebdocumentUrlInvalid = status.Error(ErrBadRequest, "WEBDOCUMENT_URL_INVALID")

	// ErrGraphInvalidReload
	// | 400 | GRAPH_INVALID_RELOAD | Invalid graph token provided, please reload the stats and provide the updated token. |
	ErrGraphInvalidReload = status.Error(ErrBadRequest, "GRAPH_INVALID_RELOAD")

	// ErrInvoicePayloadInvalid
	// | 400 | INVOICE_PAYLOAD_INVALID | The specified invoice payload is invalid. |
	ErrInvoicePayloadInvalid = status.Error(ErrBadRequest, "INVOICE_PAYLOAD_INVALID")

	// ErrChatInvalid
	// | 400 | CHAT_INVALID | Invalid chat. |
	ErrChatInvalid = status.Error(ErrBadRequest, "CHAT_INVALID")

	// ErrBotGroupsBlocked
	// | 400 | BOT_GROUPS_BLOCKED | This bot can't be added to groups. |
	ErrBotGroupsBlocked = status.Error(ErrBadRequest, "BOT_GROUPS_BLOCKED")

	// ErrGraphExpiredReload
	// | 400 | GRAPH_EXPIRED_RELOAD | This graph has expired, please obtain a new graph token. |
	ErrGraphExpiredReload = status.Error(ErrBadRequest, "GRAPH_EXPIRED_RELOAD")

	// ErrSecondsInvalid
	// | 400 | SECONDS_INVALID | Invalid duration provided. |
	ErrSecondsInvalid = status.Error(ErrBadRequest, "SECONDS_INVALID")

	// ErrArticleTitleEmpty
	// | 400 | ARTICLE_TITLE_EMPTY | The title of the article is empty. |
	ErrArticleTitleEmpty = status.Error(ErrBadRequest, "ARTICLE_TITLE_EMPTY")

	// ErrPhotoCropSizeSmall
	// | 400 | PHOTO_CROP_SIZE_SMALL | Photo is too small. |
	ErrPhotoCropSizeSmall = status.Error(ErrBadRequest, "PHOTO_CROP_SIZE_SMALL")

	// ErrPrivacyKeyInvalid
	// | 400 | PRIVACY_KEY_INVALID | The privacy key is invalid. |
	ErrPrivacyKeyInvalid = status.Error(ErrBadRequest, "PRIVACY_KEY_INVALID")

	// ErrQuizMultipleInvalid
	// | 400 | QUIZ_MULTIPLE_INVALID | Quizzes can't have the multiple_choice flag set! |
	ErrQuizMultipleInvalid = status.Error(ErrBadRequest, "QUIZ_MULTIPLE_INVALID")

	// ErrUserNotMutualContact
	// | 400 | USER_NOT_MUTUAL_CONTACT | The provided user is not a mutual contact. |
	ErrUserNotMutualContact = status.Error(ErrBadRequest, "USER_NOT_MUTUAL_CONTACT")

	// ErrChatForwardsRestricted
	// | 400 | CHAT_FORWARDS_RESTRICTED | You can't forward messages from a protected chat. |
	ErrChatForwardsRestricted = status.Error(ErrBadRequest, "CHAT_FORWARDS_RESTRICTED")

	// ErrEntitiesTooLong
	// | 400 | ENTITIES_TOO_LONG | You provided too many styled message entities. |
	ErrEntitiesTooLong = status.Error(ErrBadRequest, "ENTITIES_TOO_LONG")

	// ErrPasswordRecoveryExpired
	// | 400 | PASSWORD_RECOVERY_EXPIRED | The recovery code has expired. |
	ErrPasswordRecoveryExpired = status.Error(ErrBadRequest, "PASSWORD_RECOVERY_EXPIRED")

	// ErrPhonePasswordProtected
	// | 400 | PHONE_PASSWORD_PROTECTED | This phone is password protected. |
	ErrPhonePasswordProtected = status.Error(ErrBadRequest, "PHONE_PASSWORD_PROTECTED")

	// ErrAuthTokenException
	// | 400 | AUTH_TOKEN_EXCEPTION | An error occurred while importing the auth token. |
	ErrAuthTokenException = status.Error(ErrBadRequest, "AUTH_TOKEN_EXCEPTION")

	// ErrTranscriptionFailed
	// | 400 | TRANSCRIPTION_FAILED | Audio transcription failed. |
	ErrTranscriptionFailed = status.Error(ErrBadRequest, "TRANSCRIPTION_FAILED")

	// ErrTtlMediaInvalid
	// | 400 | TTL_MEDIA_INVALID | Invalid media Time To Live was provided. |
	ErrTtlMediaInvalid = status.Error(ErrBadRequest, "TTL_MEDIA_INVALID")

	// ErrEncryptionAlreadyAccepted
	// | 400 | ENCRYPTION_ALREADY_ACCEPTED | Secret chat already accepted. |
	ErrEncryptionAlreadyAccepted = status.Error(ErrBadRequest, "ENCRYPTION_ALREADY_ACCEPTED")

	// ErrPollAnswersInvalid
	// | 400 | POLL_ANSWERS_INVALID | Invalid poll answers were provided. |
	ErrPollAnswersInvalid = status.Error(ErrBadRequest, "POLL_ANSWERS_INVALID")

	// ErrQuizCorrectAnswersEmpty
	// | 400 | QUIZ_CORRECT_ANSWERS_EMPTY | No correct quiz answer was specified. |
	ErrQuizCorrectAnswersEmpty = status.Error(ErrBadRequest, "QUIZ_CORRECT_ANSWERS_EMPTY")

	// ErrWebpushAuthInvalid
	// | 400 | WEBPUSH_AUTH_INVALID | The specified web push authentication secret is invalid. |
	ErrWebpushAuthInvalid = status.Error(ErrBadRequest, "WEBPUSH_AUTH_INVALID")

	// ErrFileTitleEmpty
	// | 400 | FILE_TITLE_EMPTY | An empty file title was specified. |
	ErrFileTitleEmpty = status.Error(ErrBadRequest, "FILE_TITLE_EMPTY")

	// ErrContactReqMissing
	// | 400 | CONTACT_REQ_MISSING | Missing contact request. |
	ErrContactReqMissing = status.Error(ErrBadRequest, "CONTACT_REQ_MISSING")

	// ErrEmojiNotModified
	// | 400 | EMOJI_NOT_MODIFIED | The theme wasn't changed. |
	ErrEmojiNotModified = status.Error(ErrBadRequest, "EMOJI_NOT_MODIFIED")

	// ErrChatSendInlineForbidden
	// | 400 | CHAT_SEND_INLINE_FORBIDDEN | You can't send inline messages in this group. |
	ErrChatSendInlineForbidden = status.Error(ErrBadRequest, "CHAT_SEND_INLINE_FORBIDDEN")

	// ErrTtlPeriodInvalid
	// | 400 | TTL_PERIOD_INVALID | The specified TTL period is invalid. |
	ErrTtlPeriodInvalid = status.Error(ErrBadRequest, "TTL_PERIOD_INVALID")

	// ErrSwitchPmTextEmpty
	// | 400 | SWITCH_PM_TEXT_EMPTY | The switch_pm.text field was empty. |
	ErrSwitchPmTextEmpty = status.Error(ErrBadRequest, "SWITCH_PM_TEXT_EMPTY")

	// ErrMultiMediaTooLong
	// | 400 | MULTI_MEDIA_TOO_LONG | Too many media files for album. |
	ErrMultiMediaTooLong = status.Error(ErrBadRequest, "MULTI_MEDIA_TOO_LONG")

	// ErrUsernameNotOccupied
	// | 400 | USERNAME_NOT_OCCUPIED | The provided username is not occupied. |
	ErrUsernameNotOccupied = status.Error(ErrBadRequest, "USERNAME_NOT_OCCUPIED")

	// ErrTakeoutRequired
	// | 400 | TAKEOUT_REQUIRED | A takeout session has to be initialized, first. |
	ErrTakeoutRequired = status.Error(ErrBadRequest, "TAKEOUT_REQUIRED")

	// ErrQueryIdEmpty
	// | 400 | QUERY_ID_EMPTY | The query ID is empty. |
	ErrQueryIdEmpty = status.Error(ErrBadRequest, "QUERY_ID_EMPTY")

	// ErrPhoneHashExpired
	// | 400 | PHONE_HASH_EXPIRED | An invalid or expired `phone_code_hash` was provided. |
	ErrPhoneHashExpired = status.Error(ErrBadRequest, "PHONE_HASH_EXPIRED")

	// ErrWebdocumentInvalid
	// | 400 | WEBDOCUMENT_INVALID | Invalid webdocument URL provided. |
	ErrWebdocumentInvalid = status.Error(ErrBadRequest, "WEBDOCUMENT_INVALID")

	// ErrLinkNotModified
	// | 400 | LINK_NOT_MODIFIED | Discussion link not modified. |
	ErrLinkNotModified = status.Error(ErrBadRequest, "LINK_NOT_MODIFIED")

	// ErrPhotoIdInvalid
	// | 400 | PHOTO_ID_INVALID | Photo ID invalid. |
	ErrPhotoIdInvalid = status.Error(ErrBadRequest, "PHOTO_ID_INVALID")

	// ErrBotChannelsNa
	// | 400 | BOT_CHANNELS_NA | Bots can't edit admin privileges. |
	ErrBotChannelsNa = status.Error(ErrBadRequest, "BOT_CHANNELS_NA")

	// ErrQueryTooShort
	// | 400 | QUERY_TOO_SHORT | The query string is too short. |
	ErrQueryTooShort = status.Error(ErrBadRequest, "QUERY_TOO_SHORT")

	// ErrScheduleBotNotAllowed
	// | 400 | SCHEDULE_BOT_NOT_ALLOWED | Bots cannot schedule messages. |
	ErrScheduleBotNotAllowed = status.Error(ErrBadRequest, "SCHEDULE_BOT_NOT_ALLOWED")

	// ErrJoinAsPeerInvalid
	// | 400 | JOIN_AS_PEER_INVALID | The specified peer cannot be used to join a group call. |
	ErrJoinAsPeerInvalid = status.Error(ErrBadRequest, "JOIN_AS_PEER_INVALID")

	// ErrScoreInvalid
	// | 400 | SCORE_INVALID | The specified game score is invalid. |
	ErrScoreInvalid = status.Error(ErrBadRequest, "SCORE_INVALID")

	// ErrLangCodeNotSupported
	// | 400 | LANG_CODE_NOT_SUPPORTED | The specified language code is not supported. |
	ErrLangCodeNotSupported = status.Error(ErrBadRequest, "LANG_CODE_NOT_SUPPORTED")

	// ErrReplyMarkupBuyEmpty
	// | 400 | REPLY_MARKUP_BUY_EMPTY | Reply markup for buy button empty. |
	ErrReplyMarkupBuyEmpty = status.Error(ErrBadRequest, "REPLY_MARKUP_BUY_EMPTY")

	// ErrFilePartsInvalid
	// | 400 | FILE_PARTS_INVALID | The number of file parts is invalid. |
	ErrFilePartsInvalid = status.Error(ErrBadRequest, "FILE_PARTS_INVALID")

	// ErrReplyToInvalid
	// | 400 | REPLY_TO_INVALID | The specified `reply_to` field is invalid. |
	ErrReplyToInvalid = status.Error(ErrBadRequest, "REPLY_TO_INVALID")

	// ErrTokenTypeInvalid
	// | 400 | TOKEN_TYPE_INVALID | The specified token type is invalid. |
	ErrTokenTypeInvalid = status.Error(ErrBadRequest, "TOKEN_TYPE_INVALID")

	// ErrPhotoCropFileMissing
	// | 400 | PHOTO_CROP_FILE_MISSING | Photo crop file missing. |
	ErrPhotoCropFileMissing = status.Error(ErrBadRequest, "PHOTO_CROP_FILE_MISSING")

	// ErrRsaDecryptFailed
	// | 400 | RSA_DECRYPT_FAILED | Internal RSA decryption failed. |
	ErrRsaDecryptFailed = status.Error(ErrBadRequest, "RSA_DECRYPT_FAILED")

	// ErrCallAlreadyDeclined
	// | 400 | CALL_ALREADY_DECLINED | The call was already declined. |
	ErrCallAlreadyDeclined = status.Error(ErrBadRequest, "CALL_ALREADY_DECLINED")

	// ErrAccessTokenInvalid
	// | 400 | ACCESS_TOKEN_INVALID | Access token invalid. |
	ErrAccessTokenInvalid = status.Error(ErrBadRequest, "ACCESS_TOKEN_INVALID")

	// ErrChannelsTooMuch
	// | 400 | CHANNELS_TOO_MUCH | You have joined too many channels/supergroups. |
	ErrChannelsTooMuch = status.Error(ErrBadRequest, "CHANNELS_TOO_MUCH")

	// ErrFileReferenceInvalid
	// | 400 | FILE_REFERENCE_INVALID | The specified [file reference](https://core.telegram.org/api/file_reference) is invalid. |
	ErrFileReferenceInvalid = status.Error(ErrBadRequest, "FILE_REFERENCE_INVALID")

	// ErrImportFileInvalid
	// | 400 | IMPORT_FILE_INVALID | The specified chat export file is invalid. |
	ErrImportFileInvalid = status.Error(ErrBadRequest, "IMPORT_FILE_INVALID")

	// ErrWebpageCurlFailed
	// | 400 | WEBPAGE_CURL_FAILED | Failure while fetching the webpage with cURL. |
	ErrWebpageCurlFailed = status.Error(ErrBadRequest, "WEBPAGE_CURL_FAILED")

	// ErrBroadcastIdInvalid
	// | 400 | BROADCAST_ID_INVALID | Broadcast ID invalid. |
	ErrBroadcastIdInvalid = status.Error(ErrBadRequest, "BROADCAST_ID_INVALID")

	// ErrEmailHashExpired
	// | 400 | EMAIL_HASH_EXPIRED | Email hash expired. |
	ErrEmailHashExpired = status.Error(ErrBadRequest, "EMAIL_HASH_EXPIRED")

	// ErrPasswordRecoveryNa
	// | 400 | PASSWORD_RECOVERY_NA | No email was set, can't recover password via email. |
	ErrPasswordRecoveryNa = status.Error(ErrBadRequest, "PASSWORD_RECOVERY_NA")

	// ErrUsernamePurchaseAvailable
	// | 400 | USERNAME_PURCHASE_AVAILABLE | The specified username can be purchased on https://fragment.com. |
	ErrUsernamePurchaseAvailable = status.Error(ErrBadRequest, "USERNAME_PURCHASE_AVAILABLE")

	// ErrChatIdEmpty
	// | 400 | CHAT_ID_EMPTY | The provided chat ID is empty. |
	ErrChatIdEmpty = status.Error(ErrBadRequest, "CHAT_ID_EMPTY")

	// ErrPasswordHashInvalid
	// | 400 | PASSWORD_HASH_INVALID | The provided password hash is invalid. |
	ErrPasswordHashInvalid = status.Error(ErrBadRequest, "PASSWORD_HASH_INVALID")

	// ErrResultIdDuplicate
	// | 400 | RESULT_ID_DUPLICATE | You provided a duplicate result ID. |
	ErrResultIdDuplicate = status.Error(ErrBadRequest, "RESULT_ID_DUPLICATE")

	// ErrTopicNotModified
	// | 400 | TOPIC_NOT_MODIFIED | The updated topic info is equal to the current topic info, nothing was changed. |
	ErrTopicNotModified = status.Error(ErrBadRequest, "TOPIC_NOT_MODIFIED")

	// ErrUntilDateInvalid
	// | 400 | UNTIL_DATE_INVALID | Invalid until date provided. |
	ErrUntilDateInvalid = status.Error(ErrBadRequest, "UNTIL_DATE_INVALID")

	// ErrGroupcallAlreadyDiscarded
	// | 400 | GROUPCALL_ALREADY_DISCARDED | The group call was already discarded. |
	ErrGroupcallAlreadyDiscarded = status.Error(ErrBadRequest, "GROUPCALL_ALREADY_DISCARDED")

	// ErrMessagePollClosed
	// | 400 | MESSAGE_POLL_CLOSED | Poll closed. |
	ErrMessagePollClosed = status.Error(ErrBadRequest, "MESSAGE_POLL_CLOSED")

	// ErrPhotoInvalid
	// | 400 | PHOTO_INVALID | Photo invalid. |
	ErrPhotoInvalid = status.Error(ErrBadRequest, "PHOTO_INVALID")

	// ErrAuthTokenInvalidx
	// | 400 | AUTH_TOKEN_INVALIDX | The specified auth token is invalid. |
	ErrAuthTokenInvalidx = status.Error(ErrBadRequest, "AUTH_TOKEN_INVALIDX")

	// ErrOffsetPeerIdInvalid
	// | 400 | OFFSET_PEER_ID_INVALID | The provided offset peer is invalid. |
	ErrOffsetPeerIdInvalid = status.Error(ErrBadRequest, "OFFSET_PEER_ID_INVALID")

	// ErrPollAnswerInvalid
	// | 400 | POLL_ANSWER_INVALID | One of the poll answers is not acceptable. |
	ErrPollAnswerInvalid = status.Error(ErrBadRequest, "POLL_ANSWER_INVALID")

	// ErrAuthTokenAlreadyAccepted
	// | 400 | AUTH_TOKEN_ALREADY_ACCEPTED | The specified auth token was already accepted. |
	ErrAuthTokenAlreadyAccepted = status.Error(ErrBadRequest, "AUTH_TOKEN_ALREADY_ACCEPTED")

	// ErrStartParamInvalid
	// | 400 | START_PARAM_INVALID | Start parameter invalid. |
	ErrStartParamInvalid = status.Error(ErrBadRequest, "START_PARAM_INVALID")

	// ErrMediaGroupedInvalid
	// | 400 | MEDIA_GROUPED_INVALID | You tried to send media of different types in an album. |
	ErrMediaGroupedInvalid = status.Error(ErrBadRequest, "MEDIA_GROUPED_INVALID")

	// ErrChannelPrivate
	// | 400 | CHANNEL_PRIVATE | You haven't joined this channel/supergroup. |
	ErrChannelPrivate = status.Error(ErrBadRequest, "CHANNEL_PRIVATE")

	// ErrInviteRequestSent
	// | 400 | INVITE_REQUEST_SENT | You have successfully requested to join this chat or channel. |
	ErrInviteRequestSent = status.Error(ErrBadRequest, "INVITE_REQUEST_SENT")

	// ErrLastnameInvalid
	// | 400 | LASTNAME_INVALID | The last name is invalid. |
	ErrLastnameInvalid = status.Error(ErrBadRequest, "LASTNAME_INVALID")

	// ErrUserBotRequired
	// | 400 | USER_BOT_REQUIRED | This method can only be called by a bot. |
	ErrUserBotRequired = status.Error(ErrBadRequest, "USER_BOT_REQUIRED")

	// ErrChannelParicipantMissing
	// | 400 | CHANNEL_PARICIPANT_MISSING | The current user is not in the channel. |
	ErrChannelParicipantMissing = status.Error(ErrBadRequest, "CHANNEL_PARICIPANT_MISSING")

	// ErrGeneralModifyIconForbidden
	// | 400 | GENERAL_MODIFY_ICON_FORBIDDEN | You can't modify the icon of the "General" topic. |
	ErrGeneralModifyIconForbidden = status.Error(ErrBadRequest, "GENERAL_MODIFY_ICON_FORBIDDEN")

	// ErrInputTextEmpty
	// | 400 | INPUT_TEXT_EMPTY | The specified text is empty. |
	ErrInputTextEmpty = status.Error(ErrBadRequest, "INPUT_TEXT_EMPTY")

	// ErrMessageTooLong
	// | 400 | MESSAGE_TOO_LONG | The provided message is too long. |
	ErrMessageTooLong = status.Error(ErrBadRequest, "MESSAGE_TOO_LONG")

	// ErrSearchQueryEmpty
	// | 400 | SEARCH_QUERY_EMPTY | The search query is empty. |
	ErrSearchQueryEmpty = status.Error(ErrBadRequest, "SEARCH_QUERY_EMPTY")

	// ErrTypesEmpty
	// | 400 | TYPES_EMPTY | No top peer type was provided. |
	ErrTypesEmpty = status.Error(ErrBadRequest, "TYPES_EMPTY")

	// ErrChannelTooBig
	// | 400 | CHANNEL_TOO_BIG | This channel has too many participants (>1000) to be deleted. |
	ErrChannelTooBig = status.Error(ErrBadRequest, "CHANNEL_TOO_BIG")

	// ErrTopicsEmpty
	// | 400 | TOPICS_EMPTY | You specified no topic IDs. |
	ErrTopicsEmpty = status.Error(ErrBadRequest, "TOPICS_EMPTY")

	// ErrApiIdPublishedFlood
	// | 400 | API_ID_PUBLISHED_FLOOD | This API id was published somewhere, you can't use it now. |
	ErrApiIdPublishedFlood = status.Error(ErrBadRequest, "API_ID_PUBLISHED_FLOOD")

	// ErrFilePartInvalid
	// | 400 | FILE_PART_INVALID | The file part number is invalid. |
	ErrFilePartInvalid = status.Error(ErrBadRequest, "FILE_PART_INVALID")

	// ErrUserBotInvalid
	// | 400 | USER_BOT_INVALID | User accounts must provide the `bot` method parameter when calling this method. If there is no such method parameter, this method can only be invoked by bot accounts. |
	ErrUserBotInvalid = status.Error(ErrBadRequest, "USER_BOT_INVALID")

	// ErrUserInvalid
	// | 400 | USER_INVALID | Invalid user provided. |
	ErrUserInvalid = status.Error(ErrBadRequest, "USER_INVALID")

	// ErrVideoContentTypeInvalid
	// | 400 | VIDEO_CONTENT_TYPE_INVALID | The video's content type is invalid. |
	ErrVideoContentTypeInvalid = status.Error(ErrBadRequest, "VIDEO_CONTENT_TYPE_INVALID")

	// ErrAlbumPhotosTooMany
	// | 400 | ALBUM_PHOTOS_TOO_MANY | You have uploaded too many profile photos, delete some before retrying. |
	ErrAlbumPhotosTooMany = status.Error(ErrBadRequest, "ALBUM_PHOTOS_TOO_MANY")

	// ErrMessageEmpty
	// | 400 | MESSAGE_EMPTY | The provided message is empty. |
	ErrMessageEmpty = status.Error(ErrBadRequest, "MESSAGE_EMPTY")

	// ErrPublicKeyRequired
	// | 400 | PUBLIC_KEY_REQUIRED | A public key is required. |
	ErrPublicKeyRequired = status.Error(ErrBadRequest, "PUBLIC_KEY_REQUIRED")

	// ErrScheduleStatusPrivate
	// | 400 | SCHEDULE_STATUS_PRIVATE | Can't schedule until user is online, if the user's last seen timestamp is hidden by their privacy settings. |
	ErrScheduleStatusPrivate = status.Error(ErrBadRequest, "SCHEDULE_STATUS_PRIVATE")

	// ErrWallpaperInvalid
	// | 400 | WALLPAPER_INVALID | The specified wallpaper is invalid. |
	ErrWallpaperInvalid = status.Error(ErrBadRequest, "WALLPAPER_INVALID")

	// ErrEmoticonEmpty
	// | 400 | EMOTICON_EMPTY | The emoji is empty. |
	ErrEmoticonEmpty = status.Error(ErrBadRequest, "EMOTICON_EMPTY")

	// ErrChatRestricted
	// | 400 | CHAT_RESTRICTED | You can't send messages in this chat, you were restricted. |
	ErrChatRestricted = status.Error(ErrBadRequest, "CHAT_RESTRICTED")

	// ErrReactionInvalid
	// | 400 | REACTION_INVALID | The specified reaction is invalid. |
	ErrReactionInvalid = status.Error(ErrBadRequest, "REACTION_INVALID")

	// ErrAutoarchiveNotAvailable
	// | 400 | AUTOARCHIVE_NOT_AVAILABLE | The autoarchive setting is not available at this time: please check the value of the [autoarchive_setting_available field in client config &raquo;](https://core.telegram.org/api/config#client-configuration) before calling this method. |
	ErrAutoarchiveNotAvailable = status.Error(ErrBadRequest, "AUTOARCHIVE_NOT_AVAILABLE")

	// ErrSendAsPeerInvalid
	// | 400 | SEND_AS_PEER_INVALID | You can't send messages as the specified peer. |
	ErrSendAsPeerInvalid = status.Error(ErrBadRequest, "SEND_AS_PEER_INVALID")

	// ErrTempAuthKeyEmpty
	// | 400 | TEMP_AUTH_KEY_EMPTY | No temporary auth key provided. |
	ErrTempAuthKeyEmpty = status.Error(ErrBadRequest, "TEMP_AUTH_KEY_EMPTY")

	// ErrUserpicUploadRequired
	// | 400 | USERPIC_UPLOAD_REQUIRED | You must have a profile picture to publish your geolocation. |
	ErrUserpicUploadRequired = status.Error(ErrBadRequest, "USERPIC_UPLOAD_REQUIRED")

	// ErrResetRequestMissing
	// | 400 | RESET_REQUEST_MISSING | No password reset is in progress. |
	ErrResetRequestMissing = status.Error(ErrBadRequest, "RESET_REQUEST_MISSING")

	// ErrMaxQtsInvalid
	// | 400 | MAX_QTS_INVALID | The specified max_qts is invalid. |
	ErrMaxQtsInvalid = status.Error(ErrBadRequest, "MAX_QTS_INVALID")

	// ErrEmailInvalid
	// | 400 | EMAIL_INVALID | The specified email is invalid. |
	ErrEmailInvalid = status.Error(ErrBadRequest, "EMAIL_INVALID")

	// ErrEmailUnconfirmed_%d
	// | 400 | EMAIL_UNCONFIRMED_%d | The provided email isn't confirmed, %d is the length of the verification code that was just sent to the email: use [account.verifyEmail](https://core.telegram.org/method/account.verifyEmail) to enter the received verification code and enable the recovery email. |
	// ErrEmailUnconfirmed_%d = status.Error(ErrBadRequest, "EMAIL_UNCONFIRMED_%d")

	// ErrMediaFileInvalid
	// | 400 | MEDIA_FILE_INVALID | The specified media file is invalid. |
	ErrMediaFileInvalid = status.Error(ErrBadRequest, "MEDIA_FILE_INVALID")

	// ErrCodeEmpty
	// | 400 | CODE_EMPTY | The provided code is empty. |
	ErrCodeEmpty = status.Error(ErrBadRequest, "CODE_EMPTY")

	// ErrPrivacyTooLong
	// | 400 | PRIVACY_TOO_LONG | Too many privacy rules were specified, the current limit is 1000. |
	ErrPrivacyTooLong = status.Error(ErrBadRequest, "PRIVACY_TOO_LONG")

	// ErrGroupcallSsrcDuplicateMuch
	// | 400 | GROUPCALL_SSRC_DUPLICATE_MUCH | The app needs to retry joining the group call with a new SSRC value. |
	ErrGroupcallSsrcDuplicateMuch = status.Error(ErrBadRequest, "GROUPCALL_SSRC_DUPLICATE_MUCH")

	// ErrReplyMarkupTooLong
	// | 400 | REPLY_MARKUP_TOO_LONG | The specified reply_markup is too long. |
	ErrReplyMarkupTooLong = status.Error(ErrBadRequest, "REPLY_MARKUP_TOO_LONG")

	// ErrThemeFileInvalid
	// | 400 | THEME_FILE_INVALID | Invalid theme file provided. |
	ErrThemeFileInvalid = status.Error(ErrBadRequest, "THEME_FILE_INVALID")

	// ErrUsernameInvalid
	// | 400 | USERNAME_INVALID | The provided username is not valid. |
	ErrUsernameInvalid = status.Error(ErrBadRequest, "USERNAME_INVALID")

	// ErrFreshChangeAdminsForbidden
	// | 400 | FRESH_CHANGE_ADMINS_FORBIDDEN | You were just elected admin, you can't add or modify other admins yet. |
	ErrFreshChangeAdminsForbidden = status.Error(ErrBadRequest, "FRESH_CHANGE_ADMINS_FORBIDDEN")

	// ErrBotsTooMuch
	// | 400 | BOTS_TOO_MUCH | There are too many bots in this chat/channel. |
	ErrBotsTooMuch = status.Error(ErrBadRequest, "BOTS_TOO_MUCH")

	// ErrBannedRightsInvalid
	// | 400 | BANNED_RIGHTS_INVALID | You provided some invalid flags in the banned rights. |
	ErrBannedRightsInvalid = status.Error(ErrBadRequest, "BANNED_RIGHTS_INVALID")

	// ErrSendMessageTypeInvalid
	// | 400 | SEND_MESSAGE_TYPE_INVALID | The message type is invalid. |
	ErrSendMessageTypeInvalid = status.Error(ErrBadRequest, "SEND_MESSAGE_TYPE_INVALID")

	// ErrHashInvalid
	// | 400 | HASH_INVALID | The provided hash is invalid. |
	ErrHashInvalid = status.Error(ErrBadRequest, "HASH_INVALID")

	// ErrHideRequesterMissing
	// | 400 | HIDE_REQUESTER_MISSING | The join request was missing or was already handled. |
	ErrHideRequesterMissing = status.Error(ErrBadRequest, "HIDE_REQUESTER_MISSING")

	// ErrMsgIdInvalid
	// | 400 | MSG_ID_INVALID | Invalid message ID provided. |
	ErrMsgIdInvalid = status.Error(ErrBadRequest, "MSG_ID_INVALID")

	// ErrPrivacyValueInvalid
	// | 400 | PRIVACY_VALUE_INVALID | The specified privacy rule combination is invalid. |
	ErrPrivacyValueInvalid = status.Error(ErrBadRequest, "PRIVACY_VALUE_INVALID")

	// ErrThemeFormatInvalid
	// | 400 | THEME_FORMAT_INVALID | Invalid theme format provided. |
	ErrThemeFormatInvalid = status.Error(ErrBadRequest, "THEME_FORMAT_INVALID")

	// ErrTopicCloseSeparately
	// | 400 | TOPIC_CLOSE_SEPARATELY |  |
	ErrTopicCloseSeparately = status.Error(ErrBadRequest, "TOPIC_CLOSE_SEPARATELY")

	// ErrUserAdminInvalid
	// | 400 | USER_ADMIN_INVALID | You're not an admin. |
	ErrUserAdminInvalid = status.Error(ErrBadRequest, "USER_ADMIN_INVALID")

	// ErrAuthTokenExpired
	// | 400 | AUTH_TOKEN_EXPIRED | The authorization token has expired. |
	ErrAuthTokenExpired = status.Error(ErrBadRequest, "AUTH_TOKEN_EXPIRED")

	// ErrSha256HashInvalid
	// | 400 | SHA256_HASH_INVALID | The provided SHA256 hash is invalid. |
	ErrSha256HashInvalid = status.Error(ErrBadRequest, "SHA256_HASH_INVALID")

	// ErrStickerVideoNowebm
	// | 400 | STICKER_VIDEO_NOWEBM | The specified video sticker is not in webm format. |
	ErrStickerVideoNowebm = status.Error(ErrBadRequest, "STICKER_VIDEO_NOWEBM")

	// ErrSearchWithLinkNotSupported
	// | 400 | SEARCH_WITH_LINK_NOT_SUPPORTED | You cannot provide a search query and an invite link at the same time. |
	ErrSearchWithLinkNotSupported = status.Error(ErrBadRequest, "SEARCH_WITH_LINK_NOT_SUPPORTED")

	// ErrEncryptionDeclined
	// | 400 | ENCRYPTION_DECLINED | The secret chat was declined. |
	ErrEncryptionDeclined = status.Error(ErrBadRequest, "ENCRYPTION_DECLINED")

	// ErrOffsetInvalid
	// | 400 | OFFSET_INVALID | The provided offset is invalid. |
	ErrOffsetInvalid = status.Error(ErrBadRequest, "OFFSET_INVALID")

	// ErrUsernameNotModified
	// | 400 | USERNAME_NOT_MODIFIED | The username was not modified. |
	ErrUsernameNotModified = status.Error(ErrBadRequest, "USERNAME_NOT_MODIFIED")

	// ErrEmailVerifyExpired
	// | 400 | EMAIL_VERIFY_EXPIRED | The verification email has expired. |
	ErrEmailVerifyExpired = status.Error(ErrBadRequest, "EMAIL_VERIFY_EXPIRED")

	// ErrEntityMentionUserInvalid
	// | 400 | ENTITY_MENTION_USER_INVALID | You mentioned an invalid user. |
	ErrEntityMentionUserInvalid = status.Error(ErrBadRequest, "ENTITY_MENTION_USER_INVALID")

	// ErrFileReferenceEmpty
	// | 400 | FILE_REFERENCE_EMPTY | An empty [file reference](https://core.telegram.org/api/file_reference) was specified. |
	ErrFileReferenceEmpty = status.Error(ErrBadRequest, "FILE_REFERENCE_EMPTY")

	// ErrImportIdInvalid
	// | 400 | IMPORT_ID_INVALID | The specified import ID is invalid. |
	ErrImportIdInvalid = status.Error(ErrBadRequest, "IMPORT_ID_INVALID")

	// ErrInviteForbiddenWithJoinas
	// | 400 | INVITE_FORBIDDEN_WITH_JOINAS | If the user has anonymously joined a group call as a channel, they can't invite other users to the group call because that would cause deanonymization, because the invite would be sent using the original user ID, not the anonymized channel ID. |
	ErrInviteForbiddenWithJoinas = status.Error(ErrBadRequest, "INVITE_FORBIDDEN_WITH_JOINAS")

	// ErrMessageIdInvalid
	// | 400 | MESSAGE_ID_INVALID | The provided message id is invalid. |
	ErrMessageIdInvalid = status.Error(ErrBadRequest, "MESSAGE_ID_INVALID")

	// ErrResultIdInvalid
	// | 400 | RESULT_ID_INVALID | One of the specified result IDs is invalid. |
	ErrResultIdInvalid = status.Error(ErrBadRequest, "RESULT_ID_INVALID")

	// ErrStickerFileInvalid
	// | 400 | STICKER_FILE_INVALID | Sticker file invalid. |
	ErrStickerFileInvalid = status.Error(ErrBadRequest, "STICKER_FILE_INVALID")

	// ErrBroadcastPublicVotersForbidden
	// | 400 | BROADCAST_PUBLIC_VOTERS_FORBIDDEN | You can't forward polls with public voters. |
	ErrBroadcastPublicVotersForbidden = status.Error(ErrBadRequest, "BROADCAST_PUBLIC_VOTERS_FORBIDDEN")

	// ErrFirstnameInvalid
	// | 400 | FIRSTNAME_INVALID | The first name is invalid. |
	ErrFirstnameInvalid = status.Error(ErrBadRequest, "FIRSTNAME_INVALID")

	// ErrResultsTooMuch
	// | 400 | RESULTS_TOO_MUCH | Too many results were provided. |
	ErrResultsTooMuch = status.Error(ErrBadRequest, "RESULTS_TOO_MUCH")

	// ErrUsernameOccupied
	// | 400 | USERNAME_OCCUPIED | The provided username is already occupied. |
	ErrUsernameOccupied = status.Error(ErrBadRequest, "USERNAME_OCCUPIED")

	// ErrCreateCallFailed
	// | 400 | CREATE_CALL_FAILED | An error occurred while creating the call. |
	ErrCreateCallFailed = status.Error(ErrBadRequest, "CREATE_CALL_FAILED")

	// ErrOptionInvalid
	// | 400 | OPTION_INVALID | Invalid option selected. |
	ErrOptionInvalid = status.Error(ErrBadRequest, "OPTION_INVALID")

	// ErrPackShortNameInvalid
	// | 400 | PACK_SHORT_NAME_INVALID | Short pack name invalid. |
	ErrPackShortNameInvalid = status.Error(ErrBadRequest, "PACK_SHORT_NAME_INVALID")

	// ErrQuizCorrectAnswerInvalid
	// | 400 | QUIZ_CORRECT_ANSWER_INVALID | An invalid value was provided to the correct_answers field. |
	ErrQuizCorrectAnswerInvalid = status.Error(ErrBadRequest, "QUIZ_CORRECT_ANSWER_INVALID")

	// ErrSmsCodeCreateFailed
	// | 400 | SMS_CODE_CREATE_FAILED | An error occurred while creating the SMS code. |
	ErrSmsCodeCreateFailed = status.Error(ErrBadRequest, "SMS_CODE_CREATE_FAILED")

	// ErrBotDomainInvalid
	// | 400 | BOT_DOMAIN_INVALID | Bot domain invalid. |
	ErrBotDomainInvalid = status.Error(ErrBadRequest, "BOT_DOMAIN_INVALID")

	// ErrInlineResultExpired
	// | 400 | INLINE_RESULT_EXPIRED | The inline query expired. |
	ErrInlineResultExpired = status.Error(ErrBadRequest, "INLINE_RESULT_EXPIRED")

	// ErrPhoneNumberFlood
	// | 400 | PHONE_NUMBER_FLOOD | You asked for the code too many times. |
	ErrPhoneNumberFlood = status.Error(ErrBadRequest, "PHONE_NUMBER_FLOOD")

	// ErrUsernamesActiveTooMuch
	// | 400 | USERNAMES_ACTIVE_TOO_MUCH | The maximum number of active usernames was reached. |
	ErrUsernamesActiveTooMuch = status.Error(ErrBadRequest, "USERNAMES_ACTIVE_TOO_MUCH")

	// ErrChatlistExcludeInvalid
	// | 400 | CHATLIST_EXCLUDE_INVALID |  |
	ErrChatlistExcludeInvalid = status.Error(ErrBadRequest, "CHATLIST_EXCLUDE_INVALID")

	// ErrDcIdInvalid
	// | 400 | DC_ID_INVALID | The provided DC ID is invalid. |
	ErrDcIdInvalid = status.Error(ErrBadRequest, "DC_ID_INVALID")

	// ErrInviteHashInvalid
	// | 400 | INVITE_HASH_INVALID | The invite hash is invalid. |
	ErrInviteHashInvalid = status.Error(ErrBadRequest, "INVITE_HASH_INVALID")

	// ErrWebpushTokenInvalid
	// | 400 | WEBPUSH_TOKEN_INVALID | The specified web push token is invalid. |
	ErrWebpushTokenInvalid = status.Error(ErrBadRequest, "WEBPUSH_TOKEN_INVALID")

	// ErrChannelsAdminLocatedTooMuch
	// | 400 | CHANNELS_ADMIN_LOCATED_TOO_MUCH | The user has reached the limit of public geogroups. |
	ErrChannelsAdminLocatedTooMuch = status.Error(ErrBadRequest, "CHANNELS_ADMIN_LOCATED_TOO_MUCH")

	// ErrTopicIdInvalid
	// | 400 | TOPIC_ID_INVALID | The specified topic ID is invalid. |
	ErrTopicIdInvalid = status.Error(ErrBadRequest, "TOPIC_ID_INVALID")

	// ErrUserIsBot
	// | 400 | USER_IS_BOT | Bots can't send messages to other bots. |
	ErrUserIsBot = status.Error(ErrBadRequest, "USER_IS_BOT")

	// ErrRandomIdInvalid
	// | 400 | RANDOM_ID_INVALID | A provided random ID is invalid. |
	ErrRandomIdInvalid = status.Error(ErrBadRequest, "RANDOM_ID_INVALID")

	// ErrTaskAlreadyExists
	// | 400 | TASK_ALREADY_EXISTS | An email reset was already requested. |
	ErrTaskAlreadyExists = status.Error(ErrBadRequest, "TASK_ALREADY_EXISTS")

	// ErrBotAppInvalid
	// | 400 | BOT_APP_INVALID | The specified bot app is invalid. |
	ErrBotAppInvalid = status.Error(ErrBadRequest, "BOT_APP_INVALID")

	// ErrSettingsInvalid
	// | 400 | SETTINGS_INVALID | Invalid settings were provided. |
	ErrSettingsInvalid = status.Error(ErrBadRequest, "SETTINGS_INVALID")

	// ErrPhotoInvalidDimensions
	// | 400 | PHOTO_INVALID_DIMENSIONS | The photo dimensions are invalid. |
	ErrPhotoInvalidDimensions = status.Error(ErrBadRequest, "PHOTO_INVALID_DIMENSIONS")

	// ErrGifIdInvalid
	// | 400 | GIF_ID_INVALID | The provided GIF ID is invalid. |
	ErrGifIdInvalid = status.Error(ErrBadRequest, "GIF_ID_INVALID")

	// ErrPremiumAccountRequired
	// | 400 | PREMIUM_ACCOUNT_REQUIRED | A premium account is required to execute this action. |
	ErrPremiumAccountRequired = status.Error(ErrBadRequest, "PREMIUM_ACCOUNT_REQUIRED")

	// ErrDhGAInvalid
	// | 400 | DH_G_A_INVALID | g_a invalid. |
	ErrDhGAInvalid = status.Error(ErrBadRequest, "DH_G_A_INVALID")

	// ErrImportTokenInvalid
	// | 400 | IMPORT_TOKEN_INVALID | The specified token is invalid. |
	ErrImportTokenInvalid = status.Error(ErrBadRequest, "IMPORT_TOKEN_INVALID")

	// ErrLimitInvalid
	// | 400 | LIMIT_INVALID | The provided limit is invalid. |
	ErrLimitInvalid = status.Error(ErrBadRequest, "LIMIT_INVALID")

	// ErrMediaTypeInvalid
	// | 400 | MEDIA_TYPE_INVALID | The specified media type cannot be used in stories. |
	ErrMediaTypeInvalid = status.Error(ErrBadRequest, "MEDIA_TYPE_INVALID")

	// ErrRandomLengthInvalid
	// | 400 | RANDOM_LENGTH_INVALID | Random length invalid. |
	ErrRandomLengthInvalid = status.Error(ErrBadRequest, "RANDOM_LENGTH_INVALID")

	// ErrEntityBoundsInvalid
	// | 400 | ENTITY_BOUNDS_INVALID | A specified [entity offset or length](/api/entities#entity-length) is invalid, see [here &raquo;](/api/entities#entity-length) for info on how to properly compute the entity offset/length. |
	ErrEntityBoundsInvalid = status.Error(ErrBadRequest, "ENTITY_BOUNDS_INVALID")

	// ErrParticipantIdInvalid
	// | 400 | PARTICIPANT_ID_INVALID | The specified participant ID is invalid. |
	ErrParticipantIdInvalid = status.Error(ErrBadRequest, "PARTICIPANT_ID_INVALID")

	// ErrStickerDocumentInvalid
	// | 400 | STICKER_DOCUMENT_INVALID | The specified sticker document is invalid. |
	ErrStickerDocumentInvalid = status.Error(ErrBadRequest, "STICKER_DOCUMENT_INVALID")

	// ErrMd5ChecksumInvalid
	// | 400 | MD5_CHECKSUM_INVALID | The MD5 checksums do not match. |
	ErrMd5ChecksumInvalid = status.Error(ErrBadRequest, "MD5_CHECKSUM_INVALID")

	// ErrPackTitleInvalid
	// | 400 | PACK_TITLE_INVALID | The stickerpack title is invalid. |
	ErrPackTitleInvalid = status.Error(ErrBadRequest, "PACK_TITLE_INVALID")

	// ErrPhoneCodeHashEmpty
	// | 400 | PHONE_CODE_HASH_EMPTY | phone_code_hash is missing. |
	ErrPhoneCodeHashEmpty = status.Error(ErrBadRequest, "PHONE_CODE_HASH_EMPTY")

	// ErrWebpushKeyInvalid
	// | 400 | WEBPUSH_KEY_INVALID | The specified web push elliptic curve Diffie-Hellman public key is invalid. |
	ErrWebpushKeyInvalid = status.Error(ErrBadRequest, "WEBPUSH_KEY_INVALID")

	// ErrAudioTitleEmpty
	// | 400 | AUDIO_TITLE_EMPTY | An empty audio title was provided. |
	ErrAudioTitleEmpty = status.Error(ErrBadRequest, "AUDIO_TITLE_EMPTY")

	// ErrPasswordRequired
	// | 400 | PASSWORD_REQUIRED | A [2FA password](https://core.telegram.org/api/srp) must be configured to use Telegram Passport. |
	ErrPasswordRequired = status.Error(ErrBadRequest, "PASSWORD_REQUIRED")

	// ErrSessionTooFresh_%d
	// | 400 | SESSION_TOO_FRESH_%d | This session was created less than 24 hours ago, try again in %d seconds. |
	// ErrSessionTooFresh_%d = status.Error(ErrBadRequest, "SESSION_TOO_FRESH_%d")

	// ErrAccessTokenExpired
	// | 400 | ACCESS_TOKEN_EXPIRED | Access token expired. |
	ErrAccessTokenExpired = status.Error(ErrBadRequest, "ACCESS_TOKEN_EXPIRED")

	// ErrMsgWaitFailed
	// | 400 | MSG_WAIT_FAILED | A waiting call returned an error. |
	ErrMsgWaitFailed = status.Error(ErrBadRequest, "MSG_WAIT_FAILED")

	// ErrWcConvertUrlInvalid
	// | 400 | WC_CONVERT_URL_INVALID | WC convert URL invalid. |
	ErrWcConvertUrlInvalid = status.Error(ErrBadRequest, "WC_CONVERT_URL_INVALID")

	// ErrYouBlockedUser
	// | 400 | YOU_BLOCKED_USER | You blocked this user. |
	ErrYouBlockedUser = status.Error(ErrBadRequest, "YOU_BLOCKED_USER")

	// ErrExternalUrlInvalid
	// | 400 | EXTERNAL_URL_INVALID | External URL invalid. |
	ErrExternalUrlInvalid = status.Error(ErrBadRequest, "EXTERNAL_URL_INVALID")

	// ErrScheduleDateTooLate
	// | 400 | SCHEDULE_DATE_TOO_LATE | You can't schedule a message this far in the future. |
	ErrScheduleDateTooLate = status.Error(ErrBadRequest, "SCHEDULE_DATE_TOO_LATE")

	// ErrPinnedDialogsTooMuch
	// | 400 | PINNED_DIALOGS_TOO_MUCH | Too many pinned dialogs. |
	ErrPinnedDialogsTooMuch = status.Error(ErrBadRequest, "PINNED_DIALOGS_TOO_MUCH")

	// ErrChatLinkExists
	// | 400 | CHAT_LINK_EXISTS | The chat is public, you can't hide the history to new users. |
	ErrChatLinkExists = status.Error(ErrBadRequest, "CHAT_LINK_EXISTS")

	// ErrCurrencyTotalAmountInvalid
	// | 400 | CURRENCY_TOTAL_AMOUNT_INVALID | The total amount of all prices is invalid. |
	ErrCurrencyTotalAmountInvalid = status.Error(ErrBadRequest, "CURRENCY_TOTAL_AMOUNT_INVALID")

	// ErrStickerVideoNodoc
	// | 400 | STICKER_VIDEO_NODOC | You must send the video sticker as a document. |
	ErrStickerVideoNodoc = status.Error(ErrBadRequest, "STICKER_VIDEO_NODOC")

	// ErrWebdocumentSizeTooBig
	// | 400 | WEBDOCUMENT_SIZE_TOO_BIG | Webdocument is too big! |
	ErrWebdocumentSizeTooBig = status.Error(ErrBadRequest, "WEBDOCUMENT_SIZE_TOO_BIG")

	// ErrBotInlineDisabled
	// | 400 | BOT_INLINE_DISABLED | This bot can't be used in inline mode. |
	ErrBotInlineDisabled = status.Error(ErrBadRequest, "BOT_INLINE_DISABLED")

	// ErrStickersetInvalid
	// | 400 | STICKERSET_INVALID | The provided sticker set is invalid. |
	ErrStickersetInvalid = status.Error(ErrBadRequest, "STICKERSET_INVALID")

	// ErrChatIdInvalid
	// | 400 | CHAT_ID_INVALID | The provided chat id is invalid. |
	ErrChatIdInvalid = status.Error(ErrBadRequest, "CHAT_ID_INVALID")

	// ErrTtlDaysInvalid
	// | 400 | TTL_DAYS_INVALID | The provided TTL is invalid. |
	ErrTtlDaysInvalid = status.Error(ErrBadRequest, "TTL_DAYS_INVALID")

	// ErrChannelInvalid
	// | 400 | CHANNEL_INVALID | The provided channel is invalid. |
	ErrChannelInvalid = status.Error(ErrBadRequest, "CHANNEL_INVALID")

	// ErrChannelTooLarge
	// | 400 | CHANNEL_TOO_LARGE | Channel is too large to be deleted; this error is issued when trying to delete channels with more than 1000 members (subject to change). |
	ErrChannelTooLarge = status.Error(ErrBadRequest, "CHANNEL_TOO_LARGE")

	// ErrMediaTtlInvalid
	// | 400 | MEDIA_TTL_INVALID | The specified media TTL is invalid. |
	ErrMediaTtlInvalid = status.Error(ErrBadRequest, "MEDIA_TTL_INVALID")

	// ErrPhoneNumberUnoccupied
	// | 400 | PHONE_NUMBER_UNOCCUPIED | The phone number is not yet being used. |
	ErrPhoneNumberUnoccupied = status.Error(ErrBadRequest, "PHONE_NUMBER_UNOCCUPIED")

	// ErrPhotoFileMissing
	// | 400 | PHOTO_FILE_MISSING | Profile photo file missing. |
	ErrPhotoFileMissing = status.Error(ErrBadRequest, "PHOTO_FILE_MISSING")

	// ErrButtonUrlInvalid
	// | 400 | BUTTON_URL_INVALID | Button URL invalid. |
	ErrButtonUrlInvalid = status.Error(ErrBadRequest, "BUTTON_URL_INVALID")

	// ErrFilePartSizeInvalid
	// | 400 | FILE_PART_SIZE_INVALID | The provided file part size is invalid. |
	ErrFilePartSizeInvalid = status.Error(ErrBadRequest, "FILE_PART_SIZE_INVALID")

	// ErrMaxIdInvalid
	// | 400 | MAX_ID_INVALID | The provided max ID is invalid. |
	ErrMaxIdInvalid = status.Error(ErrBadRequest, "MAX_ID_INVALID")

	// ErrStoryPeriodInvalid
	// | 400 | STORY_PERIOD_INVALID |  |
	ErrStoryPeriodInvalid = status.Error(ErrBadRequest, "STORY_PERIOD_INVALID")

	// ErrContactNameEmpty
	// | 400 | CONTACT_NAME_EMPTY | Contact name empty. |
	ErrContactNameEmpty = status.Error(ErrBadRequest, "CONTACT_NAME_EMPTY")

	// ErrPasswordTooFresh_%d
	// | 400 | PASSWORD_TOO_FRESH_%d | The password was modified less than 24 hours ago, try again in %d seconds. |
	// ErrPasswordTooFresh_%d = status.Error(ErrBadRequest, "PASSWORD_TOO_FRESH_%d")

	// ErrRightsNotModified
	// | 400 | RIGHTS_NOT_MODIFIED | The new admin rights are equal to the old rights, no change was made. |
	ErrRightsNotModified = status.Error(ErrBadRequest, "RIGHTS_NOT_MODIFIED")

	// ErrStickersEmpty
	// | 400 | STICKERS_EMPTY | No sticker provided. |
	ErrStickersEmpty = status.Error(ErrBadRequest, "STICKERS_EMPTY")

	// ErrFolderIdInvalid
	// | 400 | FOLDER_ID_INVALID | Invalid folder ID. |
	ErrFolderIdInvalid = status.Error(ErrBadRequest, "FOLDER_ID_INVALID")

	// ErrGameBotInvalid
	// | 400 | GAME_BOT_INVALID | Bots can't send another bot's game. |
	ErrGameBotInvalid = status.Error(ErrBadRequest, "GAME_BOT_INVALID")

	// ErrUsersTooFew
	// | 400 | USERS_TOO_FEW | Not enough users (to create a chat, for example). |
	ErrUsersTooFew = status.Error(ErrBadRequest, "USERS_TOO_FEW")

	// ErrContactIdInvalid
	// | 400 | CONTACT_ID_INVALID | The provided contact ID is invalid. |
	ErrContactIdInvalid = status.Error(ErrBadRequest, "CONTACT_ID_INVALID")

	// ErrPhoneNumberOccupied
	// | 400 | PHONE_NUMBER_OCCUPIED | The phone number is already in use. |
	ErrPhoneNumberOccupied = status.Error(ErrBadRequest, "PHONE_NUMBER_OCCUPIED")

	// ErrUserIdInvalid
	// | 400 | USER_ID_INVALID | The provided user ID is invalid. |
	ErrUserIdInvalid = status.Error(ErrBadRequest, "USER_ID_INVALID")

	// ErrAdminRankInvalid
	// | 400 | ADMIN_RANK_INVALID | The specified admin rank is invalid. |
	ErrAdminRankInvalid = status.Error(ErrBadRequest, "ADMIN_RANK_INVALID")

	// ErrChatDiscussionUnallowed
	// | 400 | CHAT_DISCUSSION_UNALLOWED | You can't enable forum topics in a discussion group linked to a channel. |
	ErrChatDiscussionUnallowed = status.Error(ErrBadRequest, "CHAT_DISCUSSION_UNALLOWED")

	// ErrChatInvitePermanent
	// | 400 | CHAT_INVITE_PERMANENT | You can't set an expiration date on permanent invite links. |
	ErrChatInvitePermanent = status.Error(ErrBadRequest, "CHAT_INVITE_PERMANENT")

	// ErrMediaPrevInvalid
	// | 400 | MEDIA_PREV_INVALID | Previous media invalid. |
	ErrMediaPrevInvalid = status.Error(ErrBadRequest, "MEDIA_PREV_INVALID")

	// ErrQuizAnswerMissing
	// | 400 | QUIZ_ANSWER_MISSING | You can forward a quiz while hiding the original author only after choosing an option in the quiz. |
	ErrQuizAnswerMissing = status.Error(ErrBadRequest, "QUIZ_ANSWER_MISSING")

	// ErrResultTypeInvalid
	// | 400 | RESULT_TYPE_INVALID | Result type invalid. |
	ErrResultTypeInvalid = status.Error(ErrBadRequest, "RESULT_TYPE_INVALID")

	// ErrChannelsAdminPublicTooMuch
	// | 400 | CHANNELS_ADMIN_PUBLIC_TOO_MUCH | You're admin of too many public channels, make some channels private to change the username of this channel. |
	ErrChannelsAdminPublicTooMuch = status.Error(ErrBadRequest, "CHANNELS_ADMIN_PUBLIC_TOO_MUCH")

	// ErrUserAlreadyInvited
	// | 400 | USER_ALREADY_INVITED | You have already invited this user. |
	ErrUserAlreadyInvited = status.Error(ErrBadRequest, "USER_ALREADY_INVITED")

	// ErrBotInvalid
	// | 400 | BOT_INVALID | This is not a valid bot. |
	ErrBotInvalid = status.Error(ErrBadRequest, "BOT_INVALID")

	// ErrPhoneCodeExpired
	// | 400 | PHONE_CODE_EXPIRED | The phone code you provided has expired. |
	ErrPhoneCodeExpired = status.Error(ErrBadRequest, "PHONE_CODE_EXPIRED")

	// ErrReplyMarkupInvalid
	// | 400 | REPLY_MARKUP_INVALID | The provided reply markup is invalid. |
	ErrReplyMarkupInvalid = status.Error(ErrBadRequest, "REPLY_MARKUP_INVALID")

	// ErrStickerVideoBig
	// | 400 | STICKER_VIDEO_BIG | The specified video sticker is too big. |
	ErrStickerVideoBig = status.Error(ErrBadRequest, "STICKER_VIDEO_BIG")

	// ErrThemeMimeInvalid
	// | 400 | THEME_MIME_INVALID | The theme's MIME type is invalid. |
	ErrThemeMimeInvalid = status.Error(ErrBadRequest, "THEME_MIME_INVALID")

	// ErrUsageLimitInvalid
	// | 400 | USAGE_LIMIT_INVALID | The specified usage limit is invalid. |
	ErrUsageLimitInvalid = status.Error(ErrBadRequest, "USAGE_LIMIT_INVALID")

	// ErrGeoPointInvalid
	// | 400 | GEO_POINT_INVALID | Invalid geoposition provided. |
	ErrGeoPointInvalid = status.Error(ErrBadRequest, "GEO_POINT_INVALID")

	// ErrStartParamTooLong
	// | 400 | START_PARAM_TOO_LONG | Start parameter is too long. |
	ErrStartParamTooLong = status.Error(ErrBadRequest, "START_PARAM_TOO_LONG")

	// ErrPhotoExtInvalid
	// | 400 | PHOTO_EXT_INVALID | The extension of the photo is invalid. |
	ErrPhotoExtInvalid = status.Error(ErrBadRequest, "PHOTO_EXT_INVALID")

	// ErrMethodInvalid
	// | 400 | METHOD_INVALID | The specified method is invalid. |
	ErrMethodInvalid = status.Error(ErrBadRequest, "METHOD_INVALID")

	// ErrContactAddMissing
	// | 400 | CONTACT_ADD_MISSING | Contact to add is missing. |
	ErrContactAddMissing = status.Error(ErrBadRequest, "CONTACT_ADD_MISSING")

	// ErrFilePartTooBig
	// | 400 | FILE_PART_TOO_BIG | The uploaded file part is too big. |
	ErrFilePartTooBig = status.Error(ErrBadRequest, "FILE_PART_TOO_BIG")

	// ErrResultIdEmpty
	// | 400 | RESULT_ID_EMPTY | Result ID empty. |
	ErrResultIdEmpty = status.Error(ErrBadRequest, "RESULT_ID_EMPTY")

	// ErrSlowmodeMultiMsgsDisabled
	// | 400 | SLOWMODE_MULTI_MSGS_DISABLED | Slowmode is enabled, you cannot forward multiple messages to this group. |
	ErrSlowmodeMultiMsgsDisabled = status.Error(ErrBadRequest, "SLOWMODE_MULTI_MSGS_DISABLED")

	// ErrConnectionAppVersionEmpty
	// | 400 | CONNECTION_APP_VERSION_EMPTY | App version is empty. |
	ErrConnectionAppVersionEmpty = status.Error(ErrBadRequest, "CONNECTION_APP_VERSION_EMPTY")

	// ErrPhoneCodeEmpty
	// | 400 | PHONE_CODE_EMPTY | phone_code is missing. |
	ErrPhoneCodeEmpty = status.Error(ErrBadRequest, "PHONE_CODE_EMPTY")

	// ErrStoryIdEmpty
	// | 400 | STORY_ID_EMPTY | You specified no story IDs. |
	ErrStoryIdEmpty = status.Error(ErrBadRequest, "STORY_ID_EMPTY")

	// ErrTopicTitleEmpty
	// | 400 | TOPIC_TITLE_EMPTY | The specified topic title is empty. |
	ErrTopicTitleEmpty = status.Error(ErrBadRequest, "TOPIC_TITLE_EMPTY")

	// ErrCodeHashInvalid
	// | 400 | CODE_HASH_INVALID | Code hash invalid. |
	ErrCodeHashInvalid = status.Error(ErrBadRequest, "CODE_HASH_INVALID")

	// ErrDocumentInvalid
	// | 400 | DOCUMENT_INVALID | The specified document is invalid. |
	ErrDocumentInvalid = status.Error(ErrBadRequest, "DOCUMENT_INVALID")

	// ErrEmojiInvalid
	// | 400 | EMOJI_INVALID | The specified theme emoji is valid. |
	ErrEmojiInvalid = status.Error(ErrBadRequest, "EMOJI_INVALID")

	// ErrFileContentTypeInvalid
	// | 400 | FILE_CONTENT_TYPE_INVALID | File content-type is invalid. |
	ErrFileContentTypeInvalid = status.Error(ErrBadRequest, "FILE_CONTENT_TYPE_INVALID")

	// ErrStickerTgsNodoc
	// | 400 | STICKER_TGS_NODOC | You must send the animated sticker as a document. |
	ErrStickerTgsNodoc = status.Error(ErrBadRequest, "STICKER_TGS_NODOC")

	// ErrButtonTextInvalid
	// | 400 | BUTTON_TEXT_INVALID | The specified button text is invalid. |
	ErrButtonTextInvalid = status.Error(ErrBadRequest, "BUTTON_TEXT_INVALID")

	// ErrPersistentTimestampEmpty
	// | 400 | PERSISTENT_TIMESTAMP_EMPTY | Persistent timestamp empty. |
	ErrPersistentTimestampEmpty = status.Error(ErrBadRequest, "PERSISTENT_TIMESTAMP_EMPTY")

	// ErrMessageIdsEmpty
	// | 400 | MESSAGE_IDS_EMPTY | No message ids were provided. |
	ErrMessageIdsEmpty = status.Error(ErrBadRequest, "MESSAGE_IDS_EMPTY")

	// ErrUserIsBlocked
	// | 400 | USER_IS_BLOCKED | You were blocked by this user. |
	ErrUserIsBlocked = status.Error(ErrBadRequest, "USER_IS_BLOCKED")

	// ErrVoiceMessagesForbidden
	// | 400 | VOICE_MESSAGES_FORBIDDEN | This user's privacy settings forbid you from sending voice messages. |
	ErrVoiceMessagesForbidden = status.Error(ErrBadRequest, "VOICE_MESSAGES_FORBIDDEN")

	// ErrBotOnesideNotAvail
	// | 400 | BOT_ONESIDE_NOT_AVAIL | Bots can't pin messages in PM just for themselves. |
	ErrBotOnesideNotAvail = status.Error(ErrBadRequest, "BOT_ONESIDE_NOT_AVAIL")

	// ErrWallpaperMimeInvalid
	// | 400 | WALLPAPER_MIME_INVALID | The specified wallpaper MIME type is invalid. |
	ErrWallpaperMimeInvalid = status.Error(ErrBadRequest, "WALLPAPER_MIME_INVALID")

	// ErrGroupedMediaInvalid
	// | 400 | GROUPED_MEDIA_INVALID | Invalid grouped media. |
	ErrGroupedMediaInvalid = status.Error(ErrBadRequest, "GROUPED_MEDIA_INVALID")

	// ErrLocationInvalid
	// | 400 | LOCATION_INVALID | The provided location is invalid. |
	ErrLocationInvalid = status.Error(ErrBadRequest, "LOCATION_INVALID")

	// ErrParticipantVersionOutdated
	// | 400 | PARTICIPANT_VERSION_OUTDATED | The other participant does not use an up to date telegram client with support for calls. |
	ErrParticipantVersionOutdated = status.Error(ErrBadRequest, "PARTICIPANT_VERSION_OUTDATED")

	// ErrRangesInvalid
	// | 400 | RANGES_INVALID | Invalid range provided. |
	ErrRangesInvalid = status.Error(ErrBadRequest, "RANGES_INVALID")

	// ErrReactionEmpty
	// | 400 | REACTION_EMPTY | Empty reaction provided. |
	ErrReactionEmpty = status.Error(ErrBadRequest, "REACTION_EMPTY")

	// ErrAuthBytesInvalid
	// | 400 | AUTH_BYTES_INVALID | The provided authorization is invalid. |
	ErrAuthBytesInvalid = status.Error(ErrBadRequest, "AUTH_BYTES_INVALID")

	// ErrPhoneNumberBanned
	// | 400 | PHONE_NUMBER_BANNED | The provided phone number is banned from telegram. |
	ErrPhoneNumberBanned = status.Error(ErrBadRequest, "PHONE_NUMBER_BANNED")

	// ErrStickerGifDimensions
	// | 400 | STICKER_GIF_DIMENSIONS | The specified video sticker has invalid dimensions. |
	ErrStickerGifDimensions = status.Error(ErrBadRequest, "STICKER_GIF_DIMENSIONS")

	// ErrStickerPngDimensions
	// | 400 | STICKER_PNG_DIMENSIONS | Sticker png dimensions invalid. |
	ErrStickerPngDimensions = status.Error(ErrBadRequest, "STICKER_PNG_DIMENSIONS")

	// ErrToLangInvalid
	// | 400 | TO_LANG_INVALID | The specified destination language is invalid. |
	ErrToLangInvalid = status.Error(ErrBadRequest, "TO_LANG_INVALID")

	// ErrInputChatlistInvalid
	// | 400 | INPUT_CHATLIST_INVALID | The specified folder is invalid. |
	ErrInputChatlistInvalid = status.Error(ErrBadRequest, "INPUT_CHATLIST_INVALID")

	// ErrFileReferenceExpired
	// | 400 | FILE_REFERENCE_EXPIRED | File reference expired, it must be refetched as described in [the documentation](https://core.telegram.org/api/file_reference). |
	ErrFileReferenceExpired = status.Error(ErrBadRequest, "FILE_REFERENCE_EXPIRED")

	// ErrCodeInvalid
	// | 400 | CODE_INVALID | Code invalid. |
	ErrCodeInvalid = status.Error(ErrBadRequest, "CODE_INVALID")

	// ErrStickersTooMuch
	// | 400 | STICKERS_TOO_MUCH | There are too many stickers in this stickerpack, you can't add any more. |
	ErrStickersTooMuch = status.Error(ErrBadRequest, "STICKERS_TOO_MUCH")

	// ErrEncryptionIdInvalid
	// | 400 | ENCRYPTION_ID_INVALID | The provided secret chat ID is invalid. |
	ErrEncryptionIdInvalid = status.Error(ErrBadRequest, "ENCRYPTION_ID_INVALID")

	// ErrFromMessageBotDisabled
	// | 400 | FROM_MESSAGE_BOT_DISABLED | Bots can't use fromMessage min constructors. |
	ErrFromMessageBotDisabled = status.Error(ErrBadRequest, "FROM_MESSAGE_BOT_DISABLED")

	// ErrButtonDataInvalid
	// | 400 | BUTTON_DATA_INVALID | The data of one or more of the buttons you provided is invalid. |
	ErrButtonDataInvalid = status.Error(ErrBadRequest, "BUTTON_DATA_INVALID")

	// ErrMessageEditTimeExpired
	// | 400 | MESSAGE_EDIT_TIME_EXPIRED | You can't edit this message anymore, too much time has passed since its creation. |
	ErrMessageEditTimeExpired = status.Error(ErrBadRequest, "MESSAGE_EDIT_TIME_EXPIRED")

	// ErrPhoneNumberAppSignupForbidden
	// | 400 | PHONE_NUMBER_APP_SIGNUP_FORBIDDEN | You can't sign up using this app. |
	ErrPhoneNumberAppSignupForbidden = status.Error(ErrBadRequest, "PHONE_NUMBER_APP_SIGNUP_FORBIDDEN")

	// ErrBotCommandDescriptionInvalid
	// | 400 | BOT_COMMAND_DESCRIPTION_INVALID | The specified command description is invalid. |
	ErrBotCommandDescriptionInvalid = status.Error(ErrBadRequest, "BOT_COMMAND_DESCRIPTION_INVALID")

	// ErrPackShortNameOccupied
	// | 400 | PACK_SHORT_NAME_OCCUPIED | A stickerpack with this name already exists. |
	ErrPackShortNameOccupied = status.Error(ErrBadRequest, "PACK_SHORT_NAME_OCCUPIED")

	// ErrUserKicked
	// | 400 | USER_KICKED | This user was kicked from this supergroup/channel. |
	ErrUserKicked = status.Error(ErrBadRequest, "USER_KICKED")

	// ErrVideoTitleEmpty
	// | 400 | VIDEO_TITLE_EMPTY | The specified video title is empty. |
	ErrVideoTitleEmpty = status.Error(ErrBadRequest, "VIDEO_TITLE_EMPTY")

	// ErrAboutTooLong
	// | 400 | ABOUT_TOO_LONG | About string too long. |
	ErrAboutTooLong = status.Error(ErrBadRequest, "ABOUT_TOO_LONG")

	// ErrErrorTextEmpty
	// | 400 | ERROR_TEXT_EMPTY | The provided error message is empty. |
	ErrErrorTextEmpty = status.Error(ErrBadRequest, "ERROR_TEXT_EMPTY")

	// ErrMediaNewInvalid
	// | 400 | MEDIA_NEW_INVALID | The new media is invalid. |
	ErrMediaNewInvalid = status.Error(ErrBadRequest, "MEDIA_NEW_INVALID")

	// ErrPeersListEmpty
	// | 400 | PEERS_LIST_EMPTY | The specified list of peers is empty. |
	ErrPeersListEmpty = status.Error(ErrBadRequest, "PEERS_LIST_EMPTY")

	// ErrTempAuthKeyAlreadyBound
	// | 400 | TEMP_AUTH_KEY_ALREADY_BOUND | The passed temporary key is already bound to another **perm_auth_key_id**. |
	ErrTempAuthKeyAlreadyBound = status.Error(ErrBadRequest, "TEMP_AUTH_KEY_ALREADY_BOUND")

	// ErrUserNotParticipant
	// | 400 | USER_NOT_PARTICIPANT | You're not a member of this supergroup/channel. |
	ErrUserNotParticipant = status.Error(ErrBadRequest, "USER_NOT_PARTICIPANT")

	// ErrUserVolumeInvalid
	// | 400 | USER_VOLUME_INVALID | The specified user volume is invalid. |
	ErrUserVolumeInvalid = status.Error(ErrBadRequest, "USER_VOLUME_INVALID")

	// ErrChannelForumMissing
	// | 400 | CHANNEL_FORUM_MISSING | This supergroup is not a forum. |
	ErrChannelForumMissing = status.Error(ErrBadRequest, "CHANNEL_FORUM_MISSING")

	// ErrRandomIdEmpty
	// | 400 | RANDOM_ID_EMPTY | Random ID empty. |
	ErrRandomIdEmpty = status.Error(ErrBadRequest, "RANDOM_ID_EMPTY")

	// ErrMediaEmpty
	// | 400 | MEDIA_EMPTY | The provided media object is invalid. |
	ErrMediaEmpty = status.Error(ErrBadRequest, "MEDIA_EMPTY")

	// ErrThemeTitleInvalid
	// | 400 | THEME_TITLE_INVALID | The specified theme title is invalid. |
	ErrThemeTitleInvalid = status.Error(ErrBadRequest, "THEME_TITLE_INVALID")

	// ErrUserChannelsTooMuch
	// | 400 | USER_CHANNELS_TOO_MUCH | One of the users you tried to add is already in too many channels/supergroups. |
	ErrUserChannelsTooMuch = status.Error(ErrBadRequest, "USER_CHANNELS_TOO_MUCH")

	// ErrEmailNotSetup
	// | 400 | EMAIL_NOT_SETUP | In order to change the login email with emailVerifyPurposeLoginChange, an existing login email must already be set using emailVerifyPurposeLoginSetup. |
	ErrEmailNotSetup = status.Error(ErrBadRequest, "EMAIL_NOT_SETUP")

	// ErrPhotoContentTypeInvalid
	// | 400 | PHOTO_CONTENT_TYPE_INVALID | Photo mime-type invalid. |
	ErrPhotoContentTypeInvalid = status.Error(ErrBadRequest, "PHOTO_CONTENT_TYPE_INVALID")

	// ErrStickerIdInvalid
	// | 400 | STICKER_ID_INVALID | The provided sticker ID is invalid. |
	ErrStickerIdInvalid = status.Error(ErrBadRequest, "STICKER_ID_INVALID")

	// ErrChatAboutNotModified
	// | 400 | CHAT_ABOUT_NOT_MODIFIED | About text has not changed. |
	ErrChatAboutNotModified = status.Error(ErrBadRequest, "CHAT_ABOUT_NOT_MODIFIED")

	// ErrDataJsonInvalid
	// | 400 | DATA_JSON_INVALID | The provided JSON data is invalid. |
	ErrDataJsonInvalid = status.Error(ErrBadRequest, "DATA_JSON_INVALID")

	// ErrEmailUnconfirmed
	// | 400 | EMAIL_UNCONFIRMED | Email unconfirmed. |
	ErrEmailUnconfirmed = status.Error(ErrBadRequest, "EMAIL_UNCONFIRMED")

	// ErrNewSettingsInvalid
	// | 400 | NEW_SETTINGS_INVALID | The new password settings are invalid. |
	ErrNewSettingsInvalid = status.Error(ErrBadRequest, "NEW_SETTINGS_INVALID")

	// ErrDataInvalid
	// | 400 | DATA_INVALID | Encrypted data invalid. |
	ErrDataInvalid = status.Error(ErrBadRequest, "DATA_INVALID")

	// ErrGroupcallInvalid
	// | 400 | GROUPCALL_INVALID | The specified group call is invalid. |
	ErrGroupcallInvalid = status.Error(ErrBadRequest, "GROUPCALL_INVALID")

	// ErrInviteHashEmpty
	// | 400 | INVITE_HASH_EMPTY | The invite hash is empty. |
	ErrInviteHashEmpty = status.Error(ErrBadRequest, "INVITE_HASH_EMPTY")

	// ErrMaxDateInvalid
	// | 400 | MAX_DATE_INVALID | The specified maximum date is invalid. |
	ErrMaxDateInvalid = status.Error(ErrBadRequest, "MAX_DATE_INVALID")

	// ErrPhoneNotOccupied
	// | 400 | PHONE_NOT_OCCUPIED | No user is associated to the specified phone number. |
	ErrPhoneNotOccupied = status.Error(ErrBadRequest, "PHONE_NOT_OCCUPIED")

	// ErrSendMessageMediaInvalid
	// | 400 | SEND_MESSAGE_MEDIA_INVALID | Invalid media provided. |
	ErrSendMessageMediaInvalid = status.Error(ErrBadRequest, "SEND_MESSAGE_MEDIA_INVALID")

	// ErrAdminRankEmojiNotAllowed
	// | 400 | ADMIN_RANK_EMOJI_NOT_ALLOWED | An admin rank cannot contain emojis. |
	ErrAdminRankEmojiNotAllowed = status.Error(ErrBadRequest, "ADMIN_RANK_EMOJI_NOT_ALLOWED")

	// ErrGroupcallJoinMissing
	// | 400 | GROUPCALL_JOIN_MISSING | You haven't joined this group call. |
	ErrGroupcallJoinMissing = status.Error(ErrBadRequest, "GROUPCALL_JOIN_MISSING")

	// ErrGroupcallNotModified
	// | 400 | GROUPCALL_NOT_MODIFIED | Group call settings weren't modified. |
	ErrGroupcallNotModified = status.Error(ErrBadRequest, "GROUPCALL_NOT_MODIFIED")

	// ErrPollOptionInvalid
	// | 400 | POLL_OPTION_INVALID | Invalid poll option provided. |
	ErrPollOptionInvalid = status.Error(ErrBadRequest, "POLL_OPTION_INVALID")

	// ErrTokenInvalid
	// | 400 | TOKEN_INVALID | The provided token is invalid. |
	ErrTokenInvalid = status.Error(ErrBadRequest, "TOKEN_INVALID")

	// ErrButtonTypeInvalid
	// | 400 | BUTTON_TYPE_INVALID | The type of one or more of the buttons you provided is invalid. |
	ErrButtonTypeInvalid = status.Error(ErrBadRequest, "BUTTON_TYPE_INVALID")

	// ErrMediaVideoStoryMissing
	// | 400 | MEDIA_VIDEO_STORY_MISSING |  |
	ErrMediaVideoStoryMissing = status.Error(ErrBadRequest, "MEDIA_VIDEO_STORY_MISSING")

	// ErrPhoneNumberInvalid
	// | 400 | PHONE_NUMBER_INVALID | The phone number is invalid. |
	ErrPhoneNumberInvalid = status.Error(ErrBadRequest, "PHONE_NUMBER_INVALID")

	// ErrPinRestricted
	// | 400 | PIN_RESTRICTED | You can't pin messages. |
	ErrPinRestricted = status.Error(ErrBadRequest, "PIN_RESTRICTED")

	// ErrTitleInvalid
	// | 400 | TITLE_INVALID | The specified stickerpack title is invalid. |
	ErrTitleInvalid = status.Error(ErrBadRequest, "TITLE_INVALID")

	// ErrGraphOutdatedReload
	// | 400 | GRAPH_OUTDATED_RELOAD | The graph is outdated, please get a new async token using stats.getBroadcastStats. |
	ErrGraphOutdatedReload = status.Error(ErrBadRequest, "GRAPH_OUTDATED_RELOAD")

	// ErrParticipantsTooFew
	// | 400 | PARTICIPANTS_TOO_FEW | Not enough participants. |
	ErrParticipantsTooFew = status.Error(ErrBadRequest, "PARTICIPANTS_TOO_FEW")

	// ErrStickerThumbPngNopng
	// | 400 | STICKER_THUMB_PNG_NOPNG | Incorrect stickerset thumb file provided, PNG / WEBP expected. |
	ErrStickerThumbPngNopng = status.Error(ErrBadRequest, "STICKER_THUMB_PNG_NOPNG")

	// ErrStickerThumbTgsNotgs
	// | 400 | STICKER_THUMB_TGS_NOTGS | Incorrect stickerset TGS thumb file provided. |
	ErrStickerThumbTgsNotgs = status.Error(ErrBadRequest, "STICKER_THUMB_TGS_NOTGS")

	// ErrAdminIdInvalid
	// | 400 | ADMIN_ID_INVALID | The specified admin ID is invalid. |
	ErrAdminIdInvalid = status.Error(ErrBadRequest, "ADMIN_ID_INVALID")

	// ErrNextOffsetInvalid
	// | 400 | NEXT_OFFSET_INVALID | The specified offset is longer than 64 bytes. |
	ErrNextOffsetInvalid = status.Error(ErrBadRequest, "NEXT_OFFSET_INVALID")

	// ErrParticipantJoinMissing
	// | 400 | PARTICIPANT_JOIN_MISSING | Trying to enable a presentation, when the user hasn't joined the Video Chat with [phone.joinGroupCall](https://core.telegram.org/method/phone.joinGroupCall). |
	ErrParticipantJoinMissing = status.Error(ErrBadRequest, "PARTICIPANT_JOIN_MISSING")

	// ErrPhotoThumbUrlEmpty
	// | 400 | PHOTO_THUMB_URL_EMPTY | Photo thumbnail URL is empty. |
	ErrPhotoThumbUrlEmpty = status.Error(ErrBadRequest, "PHOTO_THUMB_URL_EMPTY")

	// ErrVideoFileInvalid
	// | 400 | VIDEO_FILE_INVALID | The specified video file is invalid. |
	ErrVideoFileInvalid = status.Error(ErrBadRequest, "VIDEO_FILE_INVALID")

	// ErrEmoticonInvalid
	// | 400 | EMOTICON_INVALID | The specified emoji is invalid. |
	ErrEmoticonInvalid = status.Error(ErrBadRequest, "EMOTICON_INVALID")

	// ErrExpireDateInvalid
	// | 400 | EXPIRE_DATE_INVALID | The specified expiration date is invalid. |
	ErrExpireDateInvalid = status.Error(ErrBadRequest, "EXPIRE_DATE_INVALID")

	// ErrInputUserDeactivated
	// | 400 | INPUT_USER_DEACTIVATED | The specified user was deleted. |
	ErrInputUserDeactivated = status.Error(ErrBadRequest, "INPUT_USER_DEACTIVATED")

	// ErrBotScoreNotModified
	// | 400 | BOT_SCORE_NOT_MODIFIED | The score wasn't modified. |
	ErrBotScoreNotModified = status.Error(ErrBadRequest, "BOT_SCORE_NOT_MODIFIED")

	// ErrFilterIncludeEmpty
	// | 400 | FILTER_INCLUDE_EMPTY | The include_peers vector of the filter is empty. |
	ErrFilterIncludeEmpty = status.Error(ErrBadRequest, "FILTER_INCLUDE_EMPTY")

	// ErrUrlInvalid
	// | 400 | URL_INVALID | Invalid URL provided. |
	ErrUrlInvalid = status.Error(ErrBadRequest, "URL_INVALID")

	// ErrChatPublicRequired
	// | 400 | CHAT_PUBLIC_REQUIRED | You can only enable join requests in public groups. |
	ErrChatPublicRequired = status.Error(ErrBadRequest, "CHAT_PUBLIC_REQUIRED")

	// ErrBroadcastRequired
	// | 400 | BROADCAST_REQUIRED | This method can only be called on a channel, please use stats.getMegagroupStats for supergroups. |
	ErrBroadcastRequired = status.Error(ErrBadRequest, "BROADCAST_REQUIRED")

	// ErrFileIdInvalid
	// | 400 | FILE_ID_INVALID | The provided file id is invalid. |
	ErrFileIdInvalid = status.Error(ErrBadRequest, "FILE_ID_INVALID")

	// ErrInputFilterInvalid
	// | 400 | INPUT_FILTER_INVALID | The specified filter is invalid. |
	ErrInputFilterInvalid = status.Error(ErrBadRequest, "INPUT_FILTER_INVALID")

	// ErrMegagroupPrehistoryHidden
	// | 400 | MEGAGROUP_PREHISTORY_HIDDEN | Group with hidden history for new members can't be set as discussion groups. |
	ErrMegagroupPrehistoryHidden = status.Error(ErrBadRequest, "MEGAGROUP_PREHISTORY_HIDDEN")

	// ErrPhotoSaveFileInvalid
	// | 400 | PHOTO_SAVE_FILE_INVALID | Internal issues, try again later. |
	ErrPhotoSaveFileInvalid = status.Error(ErrBadRequest, "PHOTO_SAVE_FILE_INVALID")

	// ErrAuthTokenInvalid
	// | 400 | AUTH_TOKEN_INVALID | The specified auth token is invalid. |
	ErrAuthTokenInvalid = status.Error(ErrBadRequest, "AUTH_TOKEN_INVALID")

	// ErrStickerPngNopng
	// | 400 | STICKER_PNG_NOPNG | One of the specified stickers is not a valid PNG file. |
	ErrStickerPngNopng = status.Error(ErrBadRequest, "STICKER_PNG_NOPNG")

	// ErrUserBot
	// | 400 | USER_BOT | Bots can only be admins in channels. |
	ErrUserBot = status.Error(ErrBadRequest, "USER_BOT")

	// ErrInviteHashExpired
	// | 400 | INVITE_HASH_EXPIRED | The invite link has expired. |
	ErrInviteHashExpired = status.Error(ErrBadRequest, "INVITE_HASH_EXPIRED")

	// ErrChatNotModified
	// | 400 | CHAT_NOT_MODIFIED | No changes were made to chat information because the new information you passed is identical to the current information. |
	ErrChatNotModified = status.Error(ErrBadRequest, "CHAT_NOT_MODIFIED")

	// ErrDataTooLong
	// | 400 | DATA_TOO_LONG | Data too long. |
	ErrDataTooLong = status.Error(ErrBadRequest, "DATA_TOO_LONG")

	// ErrImportFormatUnrecognized
	// | 400 | IMPORT_FORMAT_UNRECOGNIZED | The specified chat export file was exported from an unsupported chat app. |
	ErrImportFormatUnrecognized = status.Error(ErrBadRequest, "IMPORT_FORMAT_UNRECOGNIZED")

	// ErrApiIdInvalid
	// | 400 | API_ID_INVALID | API ID invalid. |
	ErrApiIdInvalid = status.Error(ErrBadRequest, "API_ID_INVALID")

	// ErrQuizCorrectAnswersTooMuch
	// | 400 | QUIZ_CORRECT_ANSWERS_TOO_MUCH | You specified too many correct answers in a quiz, quizzes can only have one right answer! |
	ErrQuizCorrectAnswersTooMuch = status.Error(ErrBadRequest, "QUIZ_CORRECT_ANSWERS_TOO_MUCH")

	// ErrGifContentTypeInvalid
	// | 400 | GIF_CONTENT_TYPE_INVALID | GIF content-type invalid. |
	ErrGifContentTypeInvalid = status.Error(ErrBadRequest, "GIF_CONTENT_TYPE_INVALID")

	// ErrNewSaltInvalid
	// | 400 | NEW_SALT_INVALID | The new salt is invalid. |
	ErrNewSaltInvalid = status.Error(ErrBadRequest, "NEW_SALT_INVALID")

	// ErrTopicDeleted
	// | 400 | TOPIC_DELETED | The specified topic was deleted. |
	ErrTopicDeleted = status.Error(ErrBadRequest, "TOPIC_DELETED")

	// ErrUserCreator
	// | 400 | USER_CREATOR | You can't leave this channel, because you're its creator. |
	ErrUserCreator = status.Error(ErrBadRequest, "USER_CREATOR")

	// ErrFileReference_*
	// | 400 | FILE_REFERENCE_* | The file reference expired, it [must be refreshed](https://core.telegram.org/api/file_reference). |
	// ErrFileReference_* = status.Error(ErrBadRequest, "FILE_REFERENCE_*")

	// ErrStartParamEmpty
	// | 400 | START_PARAM_EMPTY | The start parameter is empty. |
	ErrStartParamEmpty = status.Error(ErrBadRequest, "START_PARAM_EMPTY")

	// ErrChatRevokeDateUnsupported
	// | 400 | CHAT_REVOKE_DATE_UNSUPPORTED | `min_date` and `max_date` are not available for using with non-user peers. |
	ErrChatRevokeDateUnsupported = status.Error(ErrBadRequest, "CHAT_REVOKE_DATE_UNSUPPORTED")

	// ErrPaymentProviderInvalid
	// | 400 | PAYMENT_PROVIDER_INVALID | The specified payment provider is invalid. |
	ErrPaymentProviderInvalid = status.Error(ErrBadRequest, "PAYMENT_PROVIDER_INVALID")

	// ErrPeerIdInvalid
	// | 400 | PEER_ID_INVALID | The provided peer id is invalid. |
	ErrPeerIdInvalid = status.Error(ErrBadRequest, "PEER_ID_INVALID")

	// ErrSrpIdInvalid
	// | 400 | SRP_ID_INVALID | Invalid SRP ID provided. |
	ErrSrpIdInvalid = status.Error(ErrBadRequest, "SRP_ID_INVALID")

	// ErrTopicClosed
	// | 400 | TOPIC_CLOSED | This topic was closed, you can't send messages to it anymore. |
	ErrTopicClosed = status.Error(ErrBadRequest, "TOPIC_CLOSED")

	// ErrInviteSlugEmpty
	// | 400 | INVITE_SLUG_EMPTY | The specified invite slug is empty. |
	ErrInviteSlugEmpty = status.Error(ErrBadRequest, "INVITE_SLUG_EMPTY")

	// ErrFolderIdEmpty
	// | 400 | FOLDER_ID_EMPTY | An empty folder ID was specified. |
	ErrFolderIdEmpty = status.Error(ErrBadRequest, "FOLDER_ID_EMPTY")

	// ErrImageProcessFailed
	// | 400 | IMAGE_PROCESS_FAILED | Failure while processing image. |
	ErrImageProcessFailed = status.Error(ErrBadRequest, "IMAGE_PROCESS_FAILED")

	// ErrInviteSlugExpired
	// | 400 | INVITE_SLUG_EXPIRED | The specified chat folder link has expired. |
	ErrInviteSlugExpired = status.Error(ErrBadRequest, "INVITE_SLUG_EXPIRED")

	// ErrMsgTooOld
	// | 400 | MSG_TOO_OLD | [`chat_read_mark_expire_period` seconds](https://core.telegram.org/api/config#chat-read-mark-expire-period) have passed since the message was sent, read receipts were deleted. |
	ErrMsgTooOld = status.Error(ErrBadRequest, "MSG_TOO_OLD")

	// ErrStickerpackStickersTooMuch
	// | 400 | STICKERPACK_STICKERS_TOO_MUCH | There are too many stickers in this stickerpack, you can't add any more. |
	ErrStickerpackStickersTooMuch = status.Error(ErrBadRequest, "STICKERPACK_STICKERS_TOO_MUCH")

	// ErrChatAdminRequired
	// | 400 | CHAT_ADMIN_REQUIRED | You must be an admin in this chat to do this. |
	ErrChatAdminRequired = status.Error(ErrBadRequest, "CHAT_ADMIN_REQUIRED")

	// ErrFileEmtpy
	// | 400 | FILE_EMTPY | An empty file was provided. |
	ErrFileEmtpy = status.Error(ErrBadRequest, "FILE_EMTPY")

	// ErrFilterTitleEmpty
	// | 400 | FILTER_TITLE_EMPTY | The title field of the filter is empty. |
	ErrFilterTitleEmpty = status.Error(ErrBadRequest, "FILTER_TITLE_EMPTY")

	// ErrUsersTooMuch
	// | 400 | USERS_TOO_MUCH | The maximum number of users has been exceeded (to create a chat, for example). |
	ErrUsersTooMuch = status.Error(ErrBadRequest, "USERS_TOO_MUCH")

	// ErrEncryptedMessageInvalid
	// | 400 | ENCRYPTED_MESSAGE_INVALID | Encrypted message invalid. |
	ErrEncryptedMessageInvalid = status.Error(ErrBadRequest, "ENCRYPTED_MESSAGE_INVALID")

	// ErrUserAlreadyParticipant
	// | 400 | USER_ALREADY_PARTICIPANT | The user is already in the group. |
	ErrUserAlreadyParticipant = status.Error(ErrBadRequest, "USER_ALREADY_PARTICIPANT")

	// ErrBotResponseTimeout
	// | 400 | BOT_RESPONSE_TIMEOUT | A timeout occurred while fetching data from the bot. |
	ErrBotResponseTimeout = status.Error(ErrBadRequest, "BOT_RESPONSE_TIMEOUT")

	// ErrMegagroupRequired
	// | 400 | MEGAGROUP_REQUIRED | You can only use this method on a supergroup. |
	ErrMegagroupRequired = status.Error(ErrBadRequest, "MEGAGROUP_REQUIRED")

	// ErrNewSettingsEmpty
	// | 400 | NEW_SETTINGS_EMPTY | No password is set on the current account, and no new password was specified in `new_settings`. |
	ErrNewSettingsEmpty = status.Error(ErrBadRequest, "NEW_SETTINGS_EMPTY")

	// ErrStickerInvalid
	// | 400 | STICKER_INVALID | The provided sticker is invalid. |
	ErrStickerInvalid = status.Error(ErrBadRequest, "STICKER_INVALID")

	// ErrTokenEmpty
	// | 400 | TOKEN_EMPTY | The specified token is empty. |
	ErrTokenEmpty = status.Error(ErrBadRequest, "TOKEN_EMPTY")

	// ErrMediaCaptionTooLong
	// | 400 | MEDIA_CAPTION_TOO_LONG | The caption is too long. |
	ErrMediaCaptionTooLong = status.Error(ErrBadRequest, "MEDIA_CAPTION_TOO_LONG")

	// ErrLangPackInvalid
	// | 400 | LANG_PACK_INVALID | The provided language pack is invalid. |
	ErrLangPackInvalid = status.Error(ErrBadRequest, "LANG_PACK_INVALID")

	// ErrMediaInvalid
	// | 400 | MEDIA_INVALID | Media invalid. |
	ErrMediaInvalid = status.Error(ErrBadRequest, "MEDIA_INVALID")

	// ErrBankCardNumberInvalid
	// | 400 | BANK_CARD_NUMBER_INVALID | The specified card number is invalid. |
	ErrBankCardNumberInvalid = status.Error(ErrBadRequest, "BANK_CARD_NUMBER_INVALID")

	// ErrCallPeerInvalid
	// | 400 | CALL_PEER_INVALID | The provided call peer object is invalid. |
	ErrCallPeerInvalid = status.Error(ErrBadRequest, "CALL_PEER_INVALID")

	// ErrCdnMethodInvalid
	// | 400 | CDN_METHOD_INVALID | You can't call this method in a CDN DC. |
	ErrCdnMethodInvalid = status.Error(ErrBadRequest, "CDN_METHOD_INVALID")

	// ErrConnectionApiIdInvalid
	// | 400 | CONNECTION_API_ID_INVALID | The provided API id is invalid. |
	ErrConnectionApiIdInvalid = status.Error(ErrBadRequest, "CONNECTION_API_ID_INVALID")

	// ErrEncryptionAlreadyDeclined
	// | 400 | ENCRYPTION_ALREADY_DECLINED | The secret chat was already declined. |
	ErrEncryptionAlreadyDeclined = status.Error(ErrBadRequest, "ENCRYPTION_ALREADY_DECLINED")

	// ErrPeerIdNotSupported
	// | 400 | PEER_ID_NOT_SUPPORTED | The provided peer ID is not supported. |
	ErrPeerIdNotSupported = status.Error(ErrBadRequest, "PEER_ID_NOT_SUPPORTED")

	// ErrStickerEmojiInvalid
	// | 400 | STICKER_EMOJI_INVALID | Sticker emoji invalid. |
	ErrStickerEmojiInvalid = status.Error(ErrBadRequest, "STICKER_EMOJI_INVALID")

	// ErrAdminsTooMuch
	// | 400 | ADMINS_TOO_MUCH | There are too many admins. |
	ErrAdminsTooMuch = status.Error(ErrBadRequest, "ADMINS_TOO_MUCH")

	// ErrConnectionLayerInvalid
	// | 400 | CONNECTION_LAYER_INVALID | Layer invalid. |
	ErrConnectionLayerInvalid = status.Error(ErrBadRequest, "CONNECTION_LAYER_INVALID")

	// ErrPollQuestionInvalid
	// | 400 | POLL_QUESTION_INVALID | One of the poll questions is not acceptable. |
	ErrPollQuestionInvalid = status.Error(ErrBadRequest, "POLL_QUESTION_INVALID")

	// ErrChannelIdInvalid
	// | 400 | CHANNEL_ID_INVALID | The specified supergroup ID is invalid. |
	ErrChannelIdInvalid = status.Error(ErrBadRequest, "CHANNEL_ID_INVALID")

	// ErrMegagroupIdInvalid
	// | 400 | MEGAGROUP_ID_INVALID | Invalid supergroup ID. |
	ErrMegagroupIdInvalid = status.Error(ErrBadRequest, "MEGAGROUP_ID_INVALID")

	// ErrPasswordMissing
	// | 400 | PASSWORD_MISSING | You must enable 2FA in order to transfer ownership of a channel. |
	ErrPasswordMissing = status.Error(ErrBadRequest, "PASSWORD_MISSING")

	// ErrPeerHistoryEmpty
	// | 400 | PEER_HISTORY_EMPTY | You can't pin an empty chat with a user. |
	ErrPeerHistoryEmpty = status.Error(ErrBadRequest, "PEER_HISTORY_EMPTY")

	// ErrCallOccupyFailed
	// | 400 | CALL_OCCUPY_FAILED | The call failed because the user is already making another call. |
	ErrCallOccupyFailed = status.Error(ErrBadRequest, "CALL_OCCUPY_FAILED")

	// ErrFilePartEmpty
	// | 400 | FILE_PART_EMPTY | The provided file part is empty. |
	ErrFilePartEmpty = status.Error(ErrBadRequest, "FILE_PART_EMPTY")

	// ErrPasswordEmpty
	// | 400 | PASSWORD_EMPTY | The provided password is empty. |
	ErrPasswordEmpty = status.Error(ErrBadRequest, "PASSWORD_EMPTY")

	// ErrPersistentTimestampInvalid
	// | 400 | PERSISTENT_TIMESTAMP_INVALID | Persistent timestamp invalid. |
	ErrPersistentTimestampInvalid = status.Error(ErrBadRequest, "PERSISTENT_TIMESTAMP_INVALID")

	// ErrShortNameOccupied
	// | 400 | SHORT_NAME_OCCUPIED | The specified short name is already in use. |
	ErrShortNameOccupied = status.Error(ErrBadRequest, "SHORT_NAME_OCCUPIED")

	// ErrWebdocumentMimeInvalid
	// | 400 | WEBDOCUMENT_MIME_INVALID | Invalid webdocument mime type provided. |
	ErrWebdocumentMimeInvalid = status.Error(ErrBadRequest, "WEBDOCUMENT_MIME_INVALID")

	// ErrCallProtocolFlagsInvalid
	// | 400 | CALL_PROTOCOL_FLAGS_INVALID | Call protocol flags invalid. |
	ErrCallProtocolFlagsInvalid = status.Error(ErrBadRequest, "CALL_PROTOCOL_FLAGS_INVALID")

	// ErrFromPeerInvalid
	// | 400 | FROM_PEER_INVALID | The specified from_id is invalid. |
	ErrFromPeerInvalid = status.Error(ErrBadRequest, "FROM_PEER_INVALID")

	// ErrOptionsTooMuch
	// | 400 | OPTIONS_TOO_MUCH | Too many options provided. |
	ErrOptionsTooMuch = status.Error(ErrBadRequest, "OPTIONS_TOO_MUCH")

	// ErrWebpageMediaEmpty
	// | 400 | WEBPAGE_MEDIA_EMPTY | Webpage media empty. |
	ErrWebpageMediaEmpty = status.Error(ErrBadRequest, "WEBPAGE_MEDIA_EMPTY")

	// ErrChatAboutTooLong
	// | 400 | CHAT_ABOUT_TOO_LONG | Chat about too long. |
	ErrChatAboutTooLong = status.Error(ErrBadRequest, "CHAT_ABOUT_TOO_LONG")

	// ErrInviteRevokedMissing
	// | 400 | INVITE_REVOKED_MISSING | The specified invite link was already revoked or is invalid. |
	ErrInviteRevokedMissing = status.Error(ErrBadRequest, "INVITE_REVOKED_MISSING")

	// ErrInvitesTooMuch
	// | 400 | INVITES_TOO_MUCH | The maximum number of per-folder invites specified by the `chatlist_invites_limit_default`/`chatlist_invites_limit_premium` [client configuration parameters &raquo;](/api/config#chatlist-invites-limit-default) was reached. |
	ErrInvitesTooMuch = status.Error(ErrBadRequest, "INVITES_TOO_MUCH")

	// ErrScheduleDateInvalid
	// | 400 | SCHEDULE_DATE_INVALID | Invalid schedule date provided. |
	ErrScheduleDateInvalid = status.Error(ErrBadRequest, "SCHEDULE_DATE_INVALID")

	// ErrWallpaperFileInvalid
	// | 400 | WALLPAPER_FILE_INVALID | The specified wallpaper file is invalid. |
	ErrWallpaperFileInvalid = status.Error(ErrBadRequest, "WALLPAPER_FILE_INVALID")

	// ErrFilePartLengthInvalid
	// | 400 | FILE_PART_LENGTH_INVALID | The length of a file part is invalid. |
	ErrFilePartLengthInvalid = status.Error(ErrBadRequest, "FILE_PART_LENGTH_INVALID")

	// ErrFilePartSizeChanged
	// | 400 | FILE_PART_SIZE_CHANGED | Provided file part size has changed. |
	ErrFilePartSizeChanged = status.Error(ErrBadRequest, "FILE_PART_SIZE_CHANGED")

	// ErrThemeInvalid
	// | 400 | THEME_INVALID | Invalid theme provided. |
	ErrThemeInvalid = status.Error(ErrBadRequest, "THEME_INVALID")

	// ErrChatTitleEmpty
	// | 400 | CHAT_TITLE_EMPTY | No chat title provided. |
	ErrChatTitleEmpty = status.Error(ErrBadRequest, "CHAT_TITLE_EMPTY")

	// ErrButtonUserPrivacyRestricted
	// | 400 | BUTTON_USER_PRIVACY_RESTRICTED | The privacy setting of the user specified in a [inputKeyboardButtonUserProfile](/constructor/inputKeyboardButtonUserProfile) button do not allow creating such a button. |
	ErrButtonUserPrivacyRestricted = status.Error(ErrBadRequest, "BUTTON_USER_PRIVACY_RESTRICTED")

	// ErrDateEmpty
	// | 400 | DATE_EMPTY | Date empty. |
	ErrDateEmpty = status.Error(ErrBadRequest, "DATE_EMPTY")

	// ErrPhotoContentUrlEmpty
	// | 400 | PHOTO_CONTENT_URL_EMPTY | Photo URL invalid. |
	ErrPhotoContentUrlEmpty = status.Error(ErrBadRequest, "PHOTO_CONTENT_URL_EMPTY")

	// ErrBotCommandInvalid
	// | 400 | BOT_COMMAND_INVALID | The specified command is invalid. |
	ErrBotCommandInvalid = status.Error(ErrBadRequest, "BOT_COMMAND_INVALID")

	// ErrGroupcallForbidden
	// | 400 | GROUPCALL_FORBIDDEN | The group call has already ended. |
	ErrGroupcallForbidden = status.Error(ErrBadRequest, "GROUPCALL_FORBIDDEN")

	// ErrPhoneCodeInvalid
	// | 400 | PHONE_CODE_INVALID | The provided phone code is invalid. |
	ErrPhoneCodeInvalid = status.Error(ErrBadRequest, "PHONE_CODE_INVALID")

	// ErrShortNameInvalid
	// | 400 | SHORT_NAME_INVALID | The specified short name is invalid. |
	ErrShortNameInvalid = status.Error(ErrBadRequest, "SHORT_NAME_INVALID")

	// ErrUserBannedInChannel
	// | 400 | USER_BANNED_IN_CHANNEL | You're banned from sending messages in supergroups/channels. |
	ErrUserBannedInChannel = status.Error(ErrBadRequest, "USER_BANNED_IN_CHANNEL")

	// ErrAudioContentUrlEmpty
	// | 400 | AUDIO_CONTENT_URL_EMPTY | The remote URL specified in the content field is empty. |
	ErrAudioContentUrlEmpty = status.Error(ErrBadRequest, "AUDIO_CONTENT_URL_EMPTY")

	// ErrEmojiMarkupInvalid
	// | 400 | EMOJI_MARKUP_INVALID | The specified `video_emoji_markup` was invalid. |
	ErrEmojiMarkupInvalid = status.Error(ErrBadRequest, "EMOJI_MARKUP_INVALID")

	// ErrStickerMimeInvalid
	// | 400 | STICKER_MIME_INVALID | The specified sticker MIME type is invalid. |
	ErrStickerMimeInvalid = status.Error(ErrBadRequest, "STICKER_MIME_INVALID")

	// ErrTmpPasswordDisabled
	// | 400 | TMP_PASSWORD_DISABLED | The temporary password is disabled. |
	ErrTmpPasswordDisabled = status.Error(ErrBadRequest, "TMP_PASSWORD_DISABLED")

	// ErrBotMissing
	// | 400 | BOT_MISSING | Only bots can call this method, please use [@stickers](https://t.me/stickers) if you're a user. |
	ErrBotMissing = status.Error(ErrBadRequest, "BOT_MISSING")

	// ErrOrderInvalid
	// | 400 | ORDER_INVALID | The specified username order is invalid. |
	ErrOrderInvalid = status.Error(ErrBadRequest, "ORDER_INVALID")

	// ErrQueryIdInvalid
	// | 400 | QUERY_ID_INVALID | The query ID is invalid. |
	ErrQueryIdInvalid = status.Error(ErrBadRequest, "QUERY_ID_INVALID")

	// ErrForumEnabled
	// | 400 | FORUM_ENABLED |  |
	ErrForumEnabled = status.Error(ErrBadRequest, "FORUM_ENABLED")

	// ErrFilterNotSupported
	// | 400 | FILTER_NOT_SUPPORTED | The specified filter cannot be used in this context. |
	ErrFilterNotSupported = status.Error(ErrBadRequest, "FILTER_NOT_SUPPORTED")

	// ErrSrpPasswordChanged
	// | 400 | SRP_PASSWORD_CHANGED | Password has changed. |
	ErrSrpPasswordChanged = status.Error(ErrBadRequest, "SRP_PASSWORD_CHANGED")

	// ErrStickerTgsNotgs
	// | 400 | STICKER_TGS_NOTGS | Invalid TGS sticker provided. |
	ErrStickerTgsNotgs = status.Error(ErrBadRequest, "STICKER_TGS_NOTGS")

	// ErrStoriesNeverCreated
	// | 400 | STORIES_NEVER_CREATED |  |
	ErrStoriesNeverCreated = status.Error(ErrBadRequest, "STORIES_NEVER_CREATED")

	// ErrEmoticonStickerpackMissing
	// | 400 | EMOTICON_STICKERPACK_MISSING | inputStickerSetDice.emoji cannot be empty. |
	ErrEmoticonStickerpackMissing = status.Error(ErrBadRequest, "EMOTICON_STICKERPACK_MISSING")

	// ErrCallAlreadyAccepted
	// | 400 | CALL_ALREADY_ACCEPTED | The call was already accepted. |
	ErrCallAlreadyAccepted = status.Error(ErrBadRequest, "CALL_ALREADY_ACCEPTED")

	// ErrChatTooBig
	// | 400 | CHAT_TOO_BIG | This method is not available for groups with more than `chat_read_mark_size_threshold` members, [see client configuration &raquo;](https://core.telegram.org/api/config#client-configuration). |
	ErrChatTooBig = status.Error(ErrBadRequest, "CHAT_TOO_BIG")

	// ErrExportCardInvalid
	// | 400 | EXPORT_CARD_INVALID | Provided card is invalid. |
	ErrExportCardInvalid = status.Error(ErrBadRequest, "EXPORT_CARD_INVALID")

	// ErrMinDateInvalid
	// | 400 | MIN_DATE_INVALID | The specified minimum date is invalid. |
	ErrMinDateInvalid = status.Error(ErrBadRequest, "MIN_DATE_INVALID")

	// ErrRevoteNotAllowed
	// | 400 | REVOTE_NOT_ALLOWED | You cannot change your vote. |
	ErrRevoteNotAllowed = status.Error(ErrBadRequest, "REVOTE_NOT_ALLOWED")

	// ErrBotPaymentsDisabled
	// | 400 | BOT_PAYMENTS_DISABLED | Please enable bot payments in botfather before calling this method. |
	ErrBotPaymentsDisabled = status.Error(ErrBadRequest, "BOT_PAYMENTS_DISABLED")

	// ErrMessageNotReadYet
	// 400	MESSAGE_NOT_READ_YET	The specified message wasn't read yet.
	ErrMessageNotReadYet = status.Error(ErrBadRequest, "MESSAGE_NOT_READ_YET")
)

// 401 UNAUTHORIZED
var (
	// ErrSessionPasswordNeeded
	// 401	SESSION_PASSWORD_NEEDED	The user has enabled 2FA, more steps are needed
	ErrSessionPasswordNeeded = status.Error(ErrUnauthorized, "SESSION_PASSWORD_NEEDED")

	// ErrAuthKeyUnregistered
	// AUTH_KEY_UNREGISTERED: The key is not registered in the system
	ErrAuthKeyUnregistered = status.Error(ErrUnauthorized, "AUTH_KEY_UNREGISTERED")

	// ErrAuthKeyInvalid
	// | 401 | AUTH_KEY_INVALID | Auth key invalid. |
	ErrAuthKeyInvalid = status.Error(ErrUnauthorized, "AUTH_KEY_INVALID")

	// ErrUserDeactivated
	// USER_DEACTIVATED: The user has been deleted/deactivated
	ErrUserDeactivated = status.Error(ErrUnauthorized, "USER_DEACTIVATED")

	// ErrSessionRevoked
	// SESSION_REVOKED: The authorization has been invalidated, because of the user terminating all sessions
	ErrSessionRevoked = status.Error(ErrUnauthorized, "SESSION_REVOKED")

	// ErrSessionExpired
	// SESSION_EXPIRED: The authorization has expired
	ErrSessionExpired = status.Error(ErrUnauthorized, "SESSION_EXPIRED")

	// ErrAuthKeyPermEmpty
	// | 401 | AUTH_KEY_PERM_EMPTY | The temporary auth key must be binded to the permanent auth key to use these methods. |
	ErrAuthKeyPermEmpty = status.Error(ErrUnauthorized, "AUTH_KEY_PERM_EMPTY")

	// ErrActiveUserRequired
	// ACTIVE_USER_REQUIRED	401	The method is only available to already activated users
	ErrActiveUserRequired = status.Error(ErrUnauthorized, "ACTIVE_USER_REQUIRED")
)

// 403 FORBIDDEN
var (

	// ErrChatSendGameForbidden
	// | 403 | CHAT_SEND_GAME_FORBIDDEN | You can't send a game to this chat. |
	ErrChatSendGameForbidden = status.Error(ErrForbidden, "CHAT_SEND_GAME_FORBIDDEN")

	// ErrForbiddenGroupcallForbidden
	// | 403 | GROUPCALL_FORBIDDEN | The group call has already ended. |
	ErrForbiddenGroupcallForbidden = status.Error(ErrForbidden, "GROUPCALL_FORBIDDEN")

	// ErrMessageAuthorRequired
	// | 403 | MESSAGE_AUTHOR_REQUIRED | Message author required. |
	ErrMessageAuthorRequired = status.Error(ErrForbidden, "MESSAGE_AUTHOR_REQUIRED")

	// ErrPublicChannelMissing
	// | 403 | PUBLIC_CHANNEL_MISSING | You can only export group call invite links for public chats or channels. |
	ErrPublicChannelMissing = status.Error(ErrForbidden, "PUBLIC_CHANNEL_MISSING")

	// ErrForbiddenUserBotInvalid
	// | 403 | USER_BOT_INVALID | User accounts must provide the `bot` method parameter when calling this method. If there is no such method parameter, this method can only be invoked by bot accounts. |
	ErrForbiddenUserBotInvalid = status.Error(ErrForbidden, "USER_BOT_INVALID")

	// ErrForbiddenUserIsBlocked
	// | 403 | USER_IS_BLOCKED | You were blocked by this user. |
	ErrForbiddenUserIsBlocked = status.Error(ErrForbidden, "USER_IS_BLOCKED")

	// ErrChatGuestSendForbidden
	// | 403 | CHAT_GUEST_SEND_FORBIDDEN | You join the discussion group before commenting, see [here &raquo;](/api/discussion#requiring-users-to-join-the-group) for more info. |
	ErrChatGuestSendForbidden = status.Error(ErrForbidden, "CHAT_GUEST_SEND_FORBIDDEN")

	// ErrChatSendAudiosForbidden
	// | 403 | CHAT_SEND_AUDIOS_FORBIDDEN | You can't send audio messages in this chat. |
	ErrChatSendAudiosForbidden = status.Error(ErrForbidden, "CHAT_SEND_AUDIOS_FORBIDDEN")

	// ErrChatSendPlainForbidden
	// | 403 | CHAT_SEND_PLAIN_FORBIDDEN | You can't send non-media (text) messages in this chat. |
	ErrChatSendPlainForbidden = status.Error(ErrForbidden, "CHAT_SEND_PLAIN_FORBIDDEN")

	// ErrEditBotInviteForbidden
	// | 403 | EDIT_BOT_INVITE_FORBIDDEN | Normal users can't edit invites that were created by bots. |
	ErrEditBotInviteForbidden = status.Error(ErrForbidden, "EDIT_BOT_INVITE_FORBIDDEN")

	// ErrBroadcastForbidden
	// | 403 | BROADCAST_FORBIDDEN | Channel poll voters and reactions cannot be fetched to prevent deanonymization. |
	ErrBroadcastForbidden = status.Error(ErrForbidden, "BROADCAST_FORBIDDEN")

	// ErrForbiddenChatAdminRequired
	// | 403 | CHAT_ADMIN_REQUIRED | You must be an admin in this chat to do this. |
	ErrForbiddenChatAdminRequired = status.Error(ErrForbidden, "CHAT_ADMIN_REQUIRED")

	// ErrChatSendPollForbidden
	// | 403 | CHAT_SEND_POLL_FORBIDDEN | You can't send polls in this chat. |
	ErrChatSendPollForbidden = status.Error(ErrForbidden, "CHAT_SEND_POLL_FORBIDDEN")

	// ErrChatSendVideosForbidden
	// | 403 | CHAT_SEND_VIDEOS_FORBIDDEN | You can't send videos in this chat. |
	ErrChatSendVideosForbidden = status.Error(ErrForbidden, "CHAT_SEND_VIDEOS_FORBIDDEN")

	// ErrPollVoteRequired
	// | 403 | POLL_VOTE_REQUIRED | Cast a vote in the poll before calling this method. |
	ErrPollVoteRequired = status.Error(ErrForbidden, "POLL_VOTE_REQUIRED")

	// ErrForbiddenUserInvalid
	// | 403 | USER_INVALID | Invalid user provided. |
	ErrForbiddenUserInvalid = status.Error(ErrForbidden, "USER_INVALID")

	// ErrAnonymousReactionsDisabled
	// | 403 | ANONYMOUS_REACTIONS_DISABLED |  |
	ErrAnonymousReactionsDisabled = status.Error(ErrForbidden, "ANONYMOUS_REACTIONS_DISABLED")

	// ErrChannelPublicGroupNa
	// | 403 | CHANNEL_PUBLIC_GROUP_NA | channel/supergroup not available. |
	ErrChannelPublicGroupNa = status.Error(ErrForbidden, "CHANNEL_PUBLIC_GROUP_NA")

	// ErrChatSendPhotosForbidden
	// | 403 | CHAT_SEND_PHOTOS_FORBIDDEN | You can't send photos in this chat. |
	ErrChatSendPhotosForbidden = status.Error(ErrForbidden, "CHAT_SEND_PHOTOS_FORBIDDEN")

	// ErrMessageDeleteForbidden
	// | 403 | MESSAGE_DELETE_FORBIDDEN | You can't delete one of the messages you tried to delete, most likely because it is a service message. |
	ErrMessageDeleteForbidden = status.Error(ErrForbidden, "MESSAGE_DELETE_FORBIDDEN")

	// ErrForbiddenParticipantJoinMissing
	// | 403 | PARTICIPANT_JOIN_MISSING | Trying to enable a presentation, when the user hasn't joined the Video Chat with [phone.joinGroupCall](https://core.telegram.org/method/phone.joinGroupCall). |
	ErrForbiddenParticipantJoinMissing = status.Error(ErrForbidden, "PARTICIPANT_JOIN_MISSING")

	// ErrSensitiveChangeForbidden
	// | 403 | SENSITIVE_CHANGE_FORBIDDEN | You can't change your sensitive content settings. |
	ErrSensitiveChangeForbidden = status.Error(ErrForbidden, "SENSITIVE_CHANGE_FORBIDDEN")

	// ErrChatSendStickersForbidden
	// | 403 | CHAT_SEND_STICKERS_FORBIDDEN | You can't send stickers in this chat. |
	ErrChatSendStickersForbidden = status.Error(ErrForbidden, "CHAT_SEND_STICKERS_FORBIDDEN")

	// ErrForbiddenUserChannelsTooMuch
	// | 403 | USER_CHANNELS_TOO_MUCH | One of the users you tried to add is already in too many channels/supergroups. |
	ErrForbiddenUserChannelsTooMuch = status.Error(ErrForbidden, "USER_CHANNELS_TOO_MUCH")

	// ErrChatSendDocsForbidden
	// | 403 | CHAT_SEND_DOCS_FORBIDDEN | You can't send documents in this chat. |
	ErrChatSendDocsForbidden = status.Error(ErrForbidden, "CHAT_SEND_DOCS_FORBIDDEN")

	// ErrForbiddenChatSendInlineForbidden
	// | 403 | CHAT_SEND_INLINE_FORBIDDEN | You can't send inline messages in this group. |
	ErrForbiddenChatSendInlineForbidden = status.Error(ErrForbidden, "CHAT_SEND_INLINE_FORBIDDEN")

	// ErrRightForbidden
	// | 403 | RIGHT_FORBIDDEN | Your admin rights do not allow you to do this. |
	ErrRightForbidden = status.Error(ErrForbidden, "RIGHT_FORBIDDEN")

	// ErrForbiddenTakeoutRequired
	// | 403 | TAKEOUT_REQUIRED | A takeout session has to be initialized, first. |
	ErrForbiddenTakeoutRequired = status.Error(ErrForbidden, "TAKEOUT_REQUIRED")

	// ErrUserDeleted
	// | 403 | USER_DELETED | You can't send this secret message because the other participant deleted their account. |
	ErrUserDeleted = status.Error(ErrForbidden, "USER_DELETED")

	// ErrChatAdminInviteRequired
	// | 403 | CHAT_ADMIN_INVITE_REQUIRED | You do not have the rights to do this. |
	ErrChatAdminInviteRequired = status.Error(ErrForbidden, "CHAT_ADMIN_INVITE_REQUIRED")

	// ErrChatSendGifsForbidden
	// | 403 | CHAT_SEND_GIFS_FORBIDDEN | You can't send gifs in this chat. |
	ErrChatSendGifsForbidden = status.Error(ErrForbidden, "CHAT_SEND_GIFS_FORBIDDEN")

	// ErrChatSendVoicesForbidden
	// | 403 | CHAT_SEND_VOICES_FORBIDDEN | You can't send voice recordings in this chat. |
	ErrChatSendVoicesForbidden = status.Error(ErrForbidden, "CHAT_SEND_VOICES_FORBIDDEN")

	// ErrChatWriteForbidden
	// | 403 | CHAT_WRITE_FORBIDDEN | You can't write in this chat. |
	ErrChatWriteForbidden = status.Error(ErrForbidden, "CHAT_WRITE_FORBIDDEN")

	// ErrInlineBotRequired
	// | 403 | INLINE_BOT_REQUIRED | Only the inline bot can edit message. |
	ErrInlineBotRequired = status.Error(ErrForbidden, "INLINE_BOT_REQUIRED")

	// ErrUserRestricted
	// | 403 | USER_RESTRICTED | You're spamreported, you can't create channels or chats. |
	ErrUserRestricted = status.Error(ErrForbidden, "USER_RESTRICTED")

	// ErrChatSendMediaForbidden
	// | 403 | CHAT_SEND_MEDIA_FORBIDDEN | You can't send media in this chat. |
	ErrChatSendMediaForbidden = status.Error(ErrForbidden, "CHAT_SEND_MEDIA_FORBIDDEN")

	// ErrGroupcallAlreadyStarted
	// | 403 | GROUPCALL_ALREADY_STARTED | The groupcall has already started, you can join directly using [phone.joinGroupCall](https://core.telegram.org/method/phone.joinGroupCall). |
	ErrGroupcallAlreadyStarted = status.Error(ErrForbidden, "GROUPCALL_ALREADY_STARTED")

	// ErrForbiddenPremiumAccountRequired
	// | 403 | PREMIUM_ACCOUNT_REQUIRED | A premium account is required to execute this action. |
	ErrForbiddenPremiumAccountRequired = status.Error(ErrForbidden, "PREMIUM_ACCOUNT_REQUIRED")

	// ErrForbiddenUserNotMutualContact
	// | 403 | USER_NOT_MUTUAL_CONTACT | The provided user is not a mutual contact. |
	ErrForbiddenUserNotMutualContact = status.Error(ErrForbidden, "USER_NOT_MUTUAL_CONTACT")

	// ErrUserPrivacyRestricted
	// | 403 | USER_PRIVACY_RESTRICTED | The user's privacy settings do not allow you to do this. |
	ErrUserPrivacyRestricted = status.Error(ErrForbidden, "USER_PRIVACY_RESTRICTED")

	// ErrForbiddenVoiceMessagesForbidden
	// | 403 | VOICE_MESSAGES_FORBIDDEN | This user's privacy settings forbid you from sending voice messages. |
	ErrForbiddenVoiceMessagesForbidden = status.Error(ErrForbidden, "VOICE_MESSAGES_FORBIDDEN")
)

// 406 NOT_ACCEPTABLE

// NewErrPreviousChatImportActiveWaitX
// ErrPreviousChatImportActiveWait_%dmin
// | 406 | PREVIOUS_CHAT_IMPORT_ACTIVE_WAIT_%dMIN | Import for this chat is already in progress, wait %d minutes before starting a new one. |
func NewErrPreviousChatImportActiveWaitX(minute int32) error {
	return status.Errorf(ErrBadRequest, "PREVIOUS_CHAT_IMPORT_ACTIVE_WAIT_%ddMIN", minute)
}

var (

	// ErrSendCodeUnavailable
	// | 406 | SEND_CODE_UNAVAILABLE | Returned when all available options for this type of number were already used (e.g. flash-call, then SMS, then this error might be returned to trigger a second resend). |
	ErrSendCodeUnavailable = status.Error(ErrNotAcceptable, "SEND_CODE_UNAVAILABLE")

	// ErrNotAcceptableTopicDeleted
	// | 406 | TOPIC_DELETED | The specified topic was deleted. |
	ErrNotAcceptableTopicDeleted = status.Error(ErrNotAcceptable, "TOPIC_DELETED")

	// ErrNotAcceptableUserpicUploadRequired
	// | 406 | USERPIC_UPLOAD_REQUIRED | You must have a profile picture to publish your geolocation. |
	ErrNotAcceptableUserpicUploadRequired = status.Error(ErrNotAcceptable, "USERPIC_UPLOAD_REQUIRED")

	// ErrNotAcceptableChannelPrivate
	// | 406 | CHANNEL_PRIVATE | You haven't joined this channel/supergroup. |
	ErrNotAcceptableChannelPrivate = status.Error(ErrNotAcceptable, "CHANNEL_PRIVATE")

	// ErrNotAcceptableFreshChangeAdminsForbidden
	// | 406 | FRESH_CHANGE_ADMINS_FORBIDDEN | You were just elected admin, you can't add or modify other admins yet. |
	ErrNotAcceptableFreshChangeAdminsForbidden = status.Error(ErrNotAcceptable, "FRESH_CHANGE_ADMINS_FORBIDDEN")

	// ErrFreshChangePhoneForbidden
	// | 406 | FRESH_CHANGE_PHONE_FORBIDDEN | You can't change phone number right after logging in, please wait at least 24 hours. |
	ErrFreshChangePhoneForbidden = status.Error(ErrNotAcceptable, "FRESH_CHANGE_PHONE_FORBIDDEN")

	// ErrNotAcceptablePhoneNumberInvalid
	// | 406 | PHONE_NUMBER_INVALID | The phone number is invalid. |
	ErrNotAcceptablePhoneNumberInvalid = status.Error(ErrNotAcceptable, "PHONE_NUMBER_INVALID")

	// ErrNotAcceptableStickersetInvalid
	// | 406 | STICKERSET_INVALID | The provided sticker set is invalid. |
	ErrNotAcceptableStickersetInvalid = status.Error(ErrNotAcceptable, "STICKERSET_INVALID")

	// ErrUserpicPrivacyRequired
	// | 406 | USERPIC_PRIVACY_REQUIRED | You need to disable privacy settings for your profile picture in order to make your geolocation public. |
	ErrUserpicPrivacyRequired = status.Error(ErrNotAcceptable, "USERPIC_PRIVACY_REQUIRED")

	// ErrNotAcceptableUserRestricted
	// | 406 | USER_RESTRICTED | You're spamreported, you can't create channels or chats. |
	ErrNotAcceptableUserRestricted = status.Error(ErrNotAcceptable, "USER_RESTRICTED")

	// ErrNotAcceptableChatForwardsRestricted
	// | 406 | CHAT_FORWARDS_RESTRICTED | You can't forward messages from a protected chat. |
	ErrNotAcceptableChatForwardsRestricted = status.Error(ErrNotAcceptable, "CHAT_FORWARDS_RESTRICTED")

	// ErrFilerefUpgradeNeeded
	// | 406 | FILEREF_UPGRADE_NEEDED | The client has to be updated in order to support [file references](https://core.telegram.org/api/file_reference). |
	ErrFilerefUpgradeNeeded = status.Error(ErrNotAcceptable, "FILEREF_UPGRADE_NEEDED")

	// ErrFreshResetAuthorisationForbidden
	// | 406 | FRESH_RESET_AUTHORISATION_FORBIDDEN | You can't logout other sessions if less than 24 hours have passed since you logged on the current session. |
	ErrFreshResetAuthorisationForbidden = status.Error(ErrNotAcceptable, "FRESH_RESET_AUTHORISATION_FORBIDDEN")

	// ErrNotAcceptableInviteHashExpired
	// | 406 | INVITE_HASH_EXPIRED | The invite link has expired. |
	ErrNotAcceptableInviteHashExpired = status.Error(ErrNotAcceptable, "INVITE_HASH_EXPIRED")

	// ErrPaymentUnsupported
	// | 406 | PAYMENT_UNSUPPORTED | A detailed description of the error will be received separately as described [here &raquo;](https://core.telegram.org/api/errors#406-not-acceptable). |
	ErrPaymentUnsupported = status.Error(ErrNotAcceptable, "PAYMENT_UNSUPPORTED")

	// ErrPhonePasswordFlood
	// | 406 | PHONE_PASSWORD_FLOOD | You have tried logging in too many times. |
	ErrPhonePasswordFlood = status.Error(ErrNotAcceptable, "PHONE_PASSWORD_FLOOD")

	// ErrStickersetOwnerAnonymous
	// | 406 | STICKERSET_OWNER_ANONYMOUS | Provided stickerset can't be installed as group stickerset to prevent admin deanonymization. |
	ErrStickersetOwnerAnonymous = status.Error(ErrNotAcceptable, "STICKERSET_OWNER_ANONYMOUS")

	// ErrCallProtocolCompatLayerInvalid
	// | 406 | CALL_PROTOCOL_COMPAT_LAYER_INVALID | The other side of the call does not support any of the VoIP protocols supported by the local client, as specified by the `protocol.layer` and `protocol.library_versions` fields. |
	ErrCallProtocolCompatLayerInvalid = status.Error(ErrNotAcceptable, "CALL_PROTOCOL_COMPAT_LAYER_INVALID")

	// ErrNotAcceptableChannelTooLarge
	// | 406 | CHANNEL_TOO_LARGE | Channel is too large to be deleted; this error is issued when trying to delete channels with more than 1000 members (subject to change). |
	ErrNotAcceptableChannelTooLarge = status.Error(ErrNotAcceptable, "CHANNEL_TOO_LARGE")

	// ErrPreviousChatImportActiveWait_%dmin
	// | 406 | PREVIOUS_CHAT_IMPORT_ACTIVE_WAIT_%dMIN | Import for this chat is already in progress, wait %d minutes before starting a new one. |
	// ErrPreviousChatImportActiveWait_%dmin = status.Error(ErrNotAcceptable, "PREVIOUS_CHAT_IMPORT_ACTIVE_WAIT_%dMIN")

	// ErrNotAcceptableTopicClosed
	// | 406 | TOPIC_CLOSED | This topic was closed, you can't send messages to it anymore. |
	ErrNotAcceptableTopicClosed = status.Error(ErrNotAcceptable, "TOPIC_CLOSED")
)

// 500 InternalServerError
var (
	// StatusInternalServerError - StatusInternelServerError
	StatusInternalServerError = status.New(ErrInternal, "INTERNAL_SERVER_ERROR")

	// ErrInternalServerError
	// | 500 | INTERNAL_SERVER_ERROR |  |
	ErrInternalServerError = status.Error(ErrInternal, "INTERNAL_SERVER_ERROR")

	// ErrChatMembersChannel
	// | 500 | CHAT_MEMBERS_CHANNEL |  |
	ErrChatMembersChannel = status.Error(ErrInternal, "CHAT_MEMBERS_CHANNEL")

	// ErrChatIdGenerateFailed
	// | 500 | CHAT_ID_GENERATE_FAILED | Failure while generating the chat ID. |
	ErrChatIdGenerateFailed = status.Error(ErrInternal, "CHAT_ID_GENERATE_FAILED")

	// ErrPersistentTimestampOutdated
	// | 500 | PERSISTENT_TIMESTAMP_OUTDATED | Channel internal replication issues, try again later (treat this like an RPC_CALL_FAIL). |
	ErrPersistentTimestampOutdated = status.Error(ErrInternal, "PERSISTENT_TIMESTAMP_OUTDATED")

	// ErrMemberNotFound
	// | 500 | MEMBER_NOT_FOUND |  |
	ErrMemberNotFound = status.Error(ErrInternal, "MEMBER_NOT_FOUND")

	// ErrFileWriteEmpty
	// | 500 | FILE_WRITE_EMPTY |  |
	ErrFileWriteEmpty = status.Error(ErrInternal, "FILE_WRITE_EMPTY")

	// ErrAuthRestart
	// | 500 | AUTH_RESTART | Restart the authorization process. |
	ErrAuthRestart = status.Error(ErrInternal, "AUTH_RESTART")

	// ErrCdnUploadTimeout
	// | 500 | CDN_UPLOAD_TIMEOUT | A server-side timeout occurred while reuploading the file to the CDN DC. |
	ErrCdnUploadTimeout = status.Error(ErrInternal, "CDN_UPLOAD_TIMEOUT")

	// ErrChannelIdGenerateFailed
	// | 500 | CHANNEL_ID_GENERATE_FAILED |  |
	ErrChannelIdGenerateFailed = status.Error(ErrInternal, "CHANNEL_ID_GENERATE_FAILED")

	// ErrGroupcallAddParticipantsFailed
	// | 500 | GROUPCALL_ADD_PARTICIPANTS_FAILED |  |
	ErrGroupcallAddParticipantsFailed = status.Error(ErrInternal, "GROUPCALL_ADD_PARTICIPANTS_FAILED")

	// ErrMemberChatAddFailed
	// | 500 | MEMBER_CHAT_ADD_FAILED |  |
	ErrMemberChatAddFailed = status.Error(ErrInternal, "MEMBER_CHAT_ADD_FAILED")

	// ErrInternalMsgWaitFailed
	// | 500 | MSG_WAIT_FAILED | A waiting call returned an error. |
	ErrInternalMsgWaitFailed = status.Error(ErrInternal, "MSG_WAIT_FAILED")

	// ErrRandomIdDuplicate
	// | 500 | RANDOM_ID_DUPLICATE | You provided a random ID that was already used. |
	ErrRandomIdDuplicate = status.Error(ErrInternal, "RANDOM_ID_DUPLICATE")

	// ErrSignInFailed
	// | 500 | SIGN_IN_FAILED | Failure while signing in. |
	ErrSignInFailed = status.Error(ErrInternal, "SIGN_IN_FAILED")

	// ErrInternalCallOccupyFailed
	// | 500 | CALL_OCCUPY_FAILED | The call failed because the user is already making another call. |
	ErrInternalCallOccupyFailed = status.Error(ErrInternal, "CALL_OCCUPY_FAILED")
)

// -503 Timeout
var (
	// StatusTimeout - StatusTimeout
	StatusTimeout = status.New(ErrTimeOut503, "Timeout")

	// ErrTimeout
	// | -503 | Timeout | Timeout while fetching data |
	ErrTimeout = status.Error(ErrTimeOut503, "Timeout")
)

// -500
var (
// // ErrInvalid MsgResendReq Query
// // | -500 | Invalid msg_resend_req query | Invalid msg_resend_req query. |
// ErrInvalid MsgResendReq Query = status.Error(-500, "Invalid msg_resend_req query")
//
// // ErrInvalid MsgsAck Query
// // | -500 | Invalid msgs_ack query | Invalid msgs_ack query. |
// ErrInvalid MsgsAck Query = status.Error(-500, "Invalid msgs_ack query")
//
// // ErrInvalid MsgsStateReq Query
// // | -500 | Invalid msgs_state_req query | Invalid msgs_state_req query. |
// ErrInvalid MsgsStateReq Query = status.Error(-500, "Invalid msgs_state_req query")
)

// 700
var (
	// ErrPushRpcClient
	// db error
	// TLRpcErrorCodes_NOTRETURN_CLIENT TLRpcErrorCodes = 700
	ErrPushRpcClient = status.Error(ErrNotReturnClient, "NOTRETURN_CLIENT")

	// ErrMigratedToChannel
	// MIGRATED_TO_CHANNEL
	ErrMigratedToChannel = status.Error(ErrNotReturnClient, "MIGRATED_TO_CHANNEL")
)

// NewErrRedirectToX
// REDIRECT_TO_SERVER
// ErrRedirectToServer = status.Error(ErrNotReturnClient, "REDIRECT_TO_SERVER")
func NewErrRedirectToX(v string) error {
	return status.Errorf(ErrNotReturnClient, "REDIRECT_TO_%s", v)
}

// StatusErrEqual is essentially a copy of testutils.StatusErrEqual(), to avoid a
// cyclic dependency.
func StatusErrEqual(err1, err2 error) bool {
	status1, ok := status.FromError(err1)
	if !ok {
		return false
	}
	status2, ok := status.FromError(err2)
	if !ok {
		return false
	}
	return status1.Code() == status2.Code() && status1.Message() == status2.Message()
}
