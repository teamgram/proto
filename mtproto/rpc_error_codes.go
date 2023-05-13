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

	// ErrChatAboutNotModified
	// | 400 | CHAT_ABOUT_NOT_MODIFIED | About text has not changed. |
	ErrChatAboutNotModified = status.Error(400, "CHAT_ABOUT_NOT_MODIFIED")

	//// ErrFileReference_*
	//// | 400 | FILE_REFERENCE_* | The file reference expired, it [must be refreshed](https://core.telegram.org/api/file_reference). |
	//ErrFileReference_* = status.Error(400, "FILE_REFERENCE_*")

	// ErrHideRequesterMissing
	// | 400 | HIDE_REQUESTER_MISSING | The join request was missing or was already handled. |
	ErrHideRequesterMissing = status.Error(400, "HIDE_REQUESTER_MISSING")

	// ErrPeerIdNotSupported
	// | 400 | PEER_ID_NOT_SUPPORTED | The provided peer ID is not supported. |
	ErrPeerIdNotSupported = status.Error(400, "PEER_ID_NOT_SUPPORTED")

	// ErrSwitchPmTextEmpty
	// | 400 | SWITCH_PM_TEXT_EMPTY | The switch_pm.text field was empty. |
	ErrSwitchPmTextEmpty = status.Error(400, "SWITCH_PM_TEXT_EMPTY")

	// ErrFromMessageBotDisabled
	// | 400 | FROM_MESSAGE_BOT_DISABLED | Bots can't use fromMessage min constructors. |
	ErrFromMessageBotDisabled = status.Error(400, "FROM_MESSAGE_BOT_DISABLED")

	// ErrMediaCaptionTooLong
	// | 400 | MEDIA_CAPTION_TOO_LONG | The caption is too long. |
	ErrMediaCaptionTooLong = status.Error(400, "MEDIA_CAPTION_TOO_LONG")

	// ErrPinnedDialogsTooMuch
	// | 400 | PINNED_DIALOGS_TOO_MUCH | Too many pinned dialogs. |
	ErrPinnedDialogsTooMuch = status.Error(400, "PINNED_DIALOGS_TOO_MUCH")

	// ErrAdminIdInvalid
	// | 400 | ADMIN_ID_INVALID | The specified admin ID is invalid. |
	ErrAdminIdInvalid = status.Error(400, "ADMIN_ID_INVALID")

	// ErrFolderIdInvalid
	// | 400 | FOLDER_ID_INVALID | Invalid folder ID. |
	ErrFolderIdInvalid = status.Error(400, "FOLDER_ID_INVALID")

	// ErrPhoneNumberFlood
	// | 400 | PHONE_NUMBER_FLOOD | You asked for the code too many times. |
	ErrPhoneNumberFlood = status.Error(400, "PHONE_NUMBER_FLOOD")

	// ErrStickerThumbTgsNotgs
	// | 400 | STICKER_THUMB_TGS_NOTGS | Incorrect stickerset TGS thumb file provided. |
	ErrStickerThumbTgsNotgs = status.Error(400, "STICKER_THUMB_TGS_NOTGS")

	// ErrStickerVideoBig
	// | 400 | STICKER_VIDEO_BIG | The specified video sticker is too big. |
	ErrStickerVideoBig = status.Error(400, "STICKER_VIDEO_BIG")

	// ErrChannelsTooMuch
	// | 400 | CHANNELS_TOO_MUCH | You have joined too many channels/supergroups. |
	ErrChannelsTooMuch = status.Error(400, "CHANNELS_TOO_MUCH")

	// ErrCodeInvalid
	// | 400 | CODE_INVALID | Code invalid. |
	ErrCodeInvalid = status.Error(400, "CODE_INVALID")

	// ErrWebpageMediaEmpty
	// | 400 | WEBPAGE_MEDIA_EMPTY | Webpage media empty. |
	ErrWebpageMediaEmpty = status.Error(400, "WEBPAGE_MEDIA_EMPTY")

	// ErrSrpPasswordChanged
	// | 400 | SRP_PASSWORD_CHANGED | Password has changed. |
	ErrSrpPasswordChanged = status.Error(400, "SRP_PASSWORD_CHANGED")

	// ErrUsageLimitInvalid
	// | 400 | USAGE_LIMIT_INVALID | The specified usage limit is invalid. |
	ErrUsageLimitInvalid = status.Error(400, "USAGE_LIMIT_INVALID")

	// ErrUserKicked
	// | 400 | USER_KICKED | This user was kicked from this supergroup/channel. |
	ErrUserKicked = status.Error(400, "USER_KICKED")

	// ErrTempAuthKeyEmpty
	// | 400 | TEMP_AUTH_KEY_EMPTY | No temporary auth key provided. |
	ErrTempAuthKeyEmpty = status.Error(400, "TEMP_AUTH_KEY_EMPTY")

	// ErrThemeInvalid
	// | 400 | THEME_INVALID | Invalid theme provided. |
	ErrThemeInvalid = status.Error(400, "THEME_INVALID")

	// ErrVideoTitleEmpty
	// | 400 | VIDEO_TITLE_EMPTY | The specified video title is empty. |
	ErrVideoTitleEmpty = status.Error(400, "VIDEO_TITLE_EMPTY")

	// ErrWallpaperFileInvalid
	// | 400 | WALLPAPER_FILE_INVALID | The specified wallpaper file is invalid. |
	ErrWallpaperFileInvalid = status.Error(400, "WALLPAPER_FILE_INVALID")

	// ErrStartParamEmpty
	// | 400 | START_PARAM_EMPTY | The start parameter is empty. |
	ErrStartParamEmpty = status.Error(400, "START_PARAM_EMPTY")

	// ErrMethodInvalid
	// | 400 | METHOD_INVALID | The specified method is invalid. |
	ErrMethodInvalid = status.Error(400, "METHOD_INVALID")

	// ErrPhoneHashExpired
	// | 400 | PHONE_HASH_EXPIRED | An invalid or expired `phone_code_hash` was provided. |
	ErrPhoneHashExpired = status.Error(400, "PHONE_HASH_EXPIRED")

	// ErrAutoarchiveNotAvailable
	// | 400 | AUTOARCHIVE_NOT_AVAILABLE | The autoarchive setting is not available at this time: please check the value of the [autoarchive_setting_available field in client config &raquo;](https://core.telegram.org/api/config#client-configuration) before calling this method. |
	ErrAutoarchiveNotAvailable = status.Error(400, "AUTOARCHIVE_NOT_AVAILABLE")

	// ErrRightsNotModified
	// | 400 | RIGHTS_NOT_MODIFIED | The new admin rights are equal to the old rights, no change was made. |
	ErrRightsNotModified = status.Error(400, "RIGHTS_NOT_MODIFIED")

	// ErrStickerPngNopng
	// | 400 | STICKER_PNG_NOPNG | One of the specified stickers is not a valid PNG file. |
	ErrStickerPngNopng = status.Error(400, "STICKER_PNG_NOPNG")

	//// ErrStickersetInvalid
	//// | 400 | STICKERSET_INVALID | The provided sticker set is invalid. |
	//ErrStickersetInvalid = status.Error(400, "STICKERSET_INVALID")

	// ErrAdminsTooMuch
	// | 400 | ADMINS_TOO_MUCH | There are too many admins. |
	ErrAdminsTooMuch = status.Error(400, "ADMINS_TOO_MUCH")

	// ErrPeerIdInvalid
	// | 400 | PEER_ID_INVALID | The provided peer id is invalid. |
	ErrPeerIdInvalid = status.Error(400, "PEER_ID_INVALID")

	// ErrMessageIdsEmpty
	// | 400 | MESSAGE_IDS_EMPTY | No message ids were provided. |
	ErrMessageIdsEmpty = status.Error(400, "MESSAGE_IDS_EMPTY")

	// ErrMsgTooOld
	// | 400 | MSG_TOO_OLD | [`chat_read_mark_expire_period` seconds](https://core.telegram.org/api/config#chat-read-mark-expire-period) have passed since the message was sent, read receipts were deleted. |
	ErrMsgTooOld = status.Error(400, "MSG_TOO_OLD")

	// ErrMultiMediaTooLong
	// | 400 | MULTI_MEDIA_TOO_LONG | Too many media files for album. |
	ErrMultiMediaTooLong = status.Error(400, "MULTI_MEDIA_TOO_LONG")

	// ErrImportIdInvalid
	// | 400 | IMPORT_ID_INVALID | The specified import ID is invalid. |
	ErrImportIdInvalid = status.Error(400, "IMPORT_ID_INVALID")

	// ErrShortNameOccupied
	// | 400 | SHORT_NAME_OCCUPIED | The specified short name is already in use. |
	ErrShortNameOccupied = status.Error(400, "SHORT_NAME_OCCUPIED")

	// ErrUserCreator
	// | 400 | USER_CREATOR | You can't leave this channel, because you're its creator. |
	ErrUserCreator = status.Error(400, "USER_CREATOR")

	// ErrChannelsAdminLocatedTooMuch
	// | 400 | CHANNELS_ADMIN_LOCATED_TOO_MUCH | The user has reached the limit of public geogroups. |
	ErrChannelsAdminLocatedTooMuch = status.Error(400, "CHANNELS_ADMIN_LOCATED_TOO_MUCH")

	// ErrQueryIdEmpty
	// | 400 | QUERY_ID_EMPTY | The query ID is empty. |
	ErrQueryIdEmpty = status.Error(400, "QUERY_ID_EMPTY")

	// ErrInviteHashEmpty
	// | 400 | INVITE_HASH_EMPTY | The invite hash is empty. |
	ErrInviteHashEmpty = status.Error(400, "INVITE_HASH_EMPTY")

	//// ErrChannelTooLarge
	//// | 400 | CHANNEL_TOO_LARGE | Channel is too large to be deleted; this error is issued when trying to delete channels with more than 1000 members (subject to change). |
	//ErrChannelTooLarge = status.Error(400, "CHANNEL_TOO_LARGE")

	// ErrChatIdEmpty
	// | 400 | CHAT_ID_EMPTY | The provided chat ID is empty. |
	ErrChatIdEmpty = status.Error(400, "CHAT_ID_EMPTY")

	// ErrConnectionLayerInvalid
	// | 400 | CONNECTION_LAYER_INVALID | Layer invalid. |
	ErrConnectionLayerInvalid = status.Error(400, "CONNECTION_LAYER_INVALID")

	// ErrAuthTokenException
	// | 400 | AUTH_TOKEN_EXCEPTION | An error occurred while importing the auth token. |
	ErrAuthTokenException = status.Error(400, "AUTH_TOKEN_EXCEPTION")

	// ErrMessageNotModified
	// | 400 | MESSAGE_NOT_MODIFIED | The provided message data is identical to the previous message data, the message wasn't modified. |
	ErrMessageNotModified = status.Error(400, "MESSAGE_NOT_MODIFIED")

	// ErrQuizCorrectAnswersTooMuch
	// | 400 | QUIZ_CORRECT_ANSWERS_TOO_MUCH | You specified too many correct answers in a quiz, quizzes can only have one right answer! |
	ErrQuizCorrectAnswersTooMuch = status.Error(400, "QUIZ_CORRECT_ANSWERS_TOO_MUCH")

	// ErrWebpushAuthInvalid
	// | 400 | WEBPUSH_AUTH_INVALID | The specified web push authentication secret is invalid. |
	ErrWebpushAuthInvalid = status.Error(400, "WEBPUSH_AUTH_INVALID")

	// ErrPhoneNotOccupied
	// | 400 | PHONE_NOT_OCCUPIED | No user is associated to the specified phone number. |
	ErrPhoneNotOccupied = status.Error(400, "PHONE_NOT_OCCUPIED")

	// ErrScoreInvalid
	// | 400 | SCORE_INVALID | The specified game score is invalid. |
	ErrScoreInvalid = status.Error(400, "SCORE_INVALID")

	// ErrTypesEmpty
	// | 400 | TYPES_EMPTY | No top peer type was provided. |
	ErrTypesEmpty = status.Error(400, "TYPES_EMPTY")

	// ErrFileContentTypeInvalid
	// | 400 | FILE_CONTENT_TYPE_INVALID | File content-type is invalid. |
	ErrFileContentTypeInvalid = status.Error(400, "FILE_CONTENT_TYPE_INVALID")

	// ErrGameBotInvalid
	// | 400 | GAME_BOT_INVALID | Bots can't send another bot's game. |
	ErrGameBotInvalid = status.Error(400, "GAME_BOT_INVALID")

	// ErrInvoicePayloadInvalid
	// | 400 | INVOICE_PAYLOAD_INVALID | The specified invoice payload is invalid. |
	ErrInvoicePayloadInvalid = status.Error(400, "INVOICE_PAYLOAD_INVALID")

	// ErrCdnMethodInvalid
	// | 400 | CDN_METHOD_INVALID | You can't call this method in a CDN DC. |
	ErrCdnMethodInvalid = status.Error(400, "CDN_METHOD_INVALID")

	// ErrContactIdInvalid
	// | 400 | CONTACT_ID_INVALID | The provided contact ID is invalid. |
	ErrContactIdInvalid = status.Error(400, "CONTACT_ID_INVALID")

	// ErrExpireDateInvalid
	// | 400 | EXPIRE_DATE_INVALID | The specified expiration date is invalid. |
	ErrExpireDateInvalid = status.Error(400, "EXPIRE_DATE_INVALID")

	// ErrRsaDecryptFailed
	// | 400 | RSA_DECRYPT_FAILED | Internal RSA decryption failed. |
	ErrRsaDecryptFailed = status.Error(400, "RSA_DECRYPT_FAILED")

	// ErrEmailNotSetup
	// | 400 | EMAIL_NOT_SETUP | In order to change the login email with emailVerifyPurposeLoginChange, an existing login email must already be set using emailVerifyPurposeLoginSetup. |
	ErrEmailNotSetup = status.Error(400, "EMAIL_NOT_SETUP")

	// ErrInviteRevokedMissing
	// | 400 | INVITE_REVOKED_MISSING | The specified invite link was already revoked or is invalid. |
	ErrInviteRevokedMissing = status.Error(400, "INVITE_REVOKED_MISSING")

	// ErrTranscriptionFailed
	// | 400 | TRANSCRIPTION_FAILED | Audio transcription failed. |
	ErrTranscriptionFailed = status.Error(400, "TRANSCRIPTION_FAILED")

	// ErrBroadcastPublicVotersForbidden
	// | 400 | BROADCAST_PUBLIC_VOTERS_FORBIDDEN | You can't forward polls with public voters. |
	ErrBroadcastPublicVotersForbidden = status.Error(400, "BROADCAST_PUBLIC_VOTERS_FORBIDDEN")

	// ErrHashInvalid
	// | 400 | HASH_INVALID | The provided hash is invalid. |
	ErrHashInvalid = status.Error(400, "HASH_INVALID")

	// ErrImportFormatUnrecognized
	// | 400 | IMPORT_FORMAT_UNRECOGNIZED | The specified chat export file was exported from an unsupported chat app. |
	ErrImportFormatUnrecognized = status.Error(400, "IMPORT_FORMAT_UNRECOGNIZED")

	// ErrQuizMultipleInvalid
	// | 400 | QUIZ_MULTIPLE_INVALID | Quizzes can't have the multiple_choice flag set! |
	ErrQuizMultipleInvalid = status.Error(400, "QUIZ_MULTIPLE_INVALID")

	// ErrAboutTooLong
	// | 400 | ABOUT_TOO_LONG | About string too long. |
	ErrAboutTooLong = status.Error(400, "ABOUT_TOO_LONG")

	// ErrEncryptionAlreadyAccepted
	// | 400 | ENCRYPTION_ALREADY_ACCEPTED | Secret chat already accepted. |
	ErrEncryptionAlreadyAccepted = status.Error(400, "ENCRYPTION_ALREADY_ACCEPTED")

	// ErrArticleTitleEmpty
	// | 400 | ARTICLE_TITLE_EMPTY | The title of the article is empty. |
	ErrArticleTitleEmpty = status.Error(400, "ARTICLE_TITLE_EMPTY")

	// ErrImportTokenInvalid
	// | 400 | IMPORT_TOKEN_INVALID | The specified token is invalid. |
	ErrImportTokenInvalid = status.Error(400, "IMPORT_TOKEN_INVALID")

	// ErrDataJsonInvalid
	// | 400 | DATA_JSON_INVALID | The provided JSON data is invalid. |
	ErrDataJsonInvalid = status.Error(400, "DATA_JSON_INVALID")

	// ErrFileTitleEmpty
	// | 400 | FILE_TITLE_EMPTY | An empty file title was specified. |
	ErrFileTitleEmpty = status.Error(400, "FILE_TITLE_EMPTY")

	// ErrGroupcallAlreadyDiscarded
	// | 400 | GROUPCALL_ALREADY_DISCARDED | The group call was already discarded. |
	ErrGroupcallAlreadyDiscarded = status.Error(400, "GROUPCALL_ALREADY_DISCARDED")

	// ErrNewSaltInvalid
	// | 400 | NEW_SALT_INVALID | The new salt is invalid. |
	ErrNewSaltInvalid = status.Error(400, "NEW_SALT_INVALID")

	// ErrAuthBytesInvalid
	// | 400 | AUTH_BYTES_INVALID | The provided authorization is invalid. |
	ErrAuthBytesInvalid = status.Error(400, "AUTH_BYTES_INVALID")

	// ErrDataInvalid
	// | 400 | DATA_INVALID | Encrypted data invalid. |
	ErrDataInvalid = status.Error(400, "DATA_INVALID")

	// ErrStickerPngDimensions
	// | 400 | STICKER_PNG_DIMENSIONS | Sticker png dimensions invalid. |
	ErrStickerPngDimensions = status.Error(400, "STICKER_PNG_DIMENSIONS")

	// ErrUsernameInvalid
	// | 400 | USERNAME_INVALID | The provided username is not valid. |
	ErrUsernameInvalid = status.Error(400, "USERNAME_INVALID")

	// ErrUsernameNotModified
	// | 400 | USERNAME_NOT_MODIFIED | The username was not modified. |
	ErrUsernameNotModified = status.Error(400, "USERNAME_NOT_MODIFIED")

	// ErrSendAsPeerInvalid
	// | 400 | SEND_AS_PEER_INVALID | You can't send messages as the specified peer. |
	ErrSendAsPeerInvalid = status.Error(400, "SEND_AS_PEER_INVALID")

	// ErrInviteSlugEmpty
	// | 400 | INVITE_SLUG_EMPTY | The specified invite slug is empty. |
	ErrInviteSlugEmpty = status.Error(400, "INVITE_SLUG_EMPTY")

	// ErrScheduleTooMuch
	// | 400 | SCHEDULE_TOO_MUCH | There are too many scheduled messages. |
	ErrScheduleTooMuch = status.Error(400, "SCHEDULE_TOO_MUCH")

	// ErrStickerpackStickersTooMuch
	// | 400 | STICKERPACK_STICKERS_TOO_MUCH | There are too many stickers in this stickerpack, you can't add any more. |
	ErrStickerpackStickersTooMuch = status.Error(400, "STICKERPACK_STICKERS_TOO_MUCH")

	// ErrMessageIdInvalid
	// | 400 | MESSAGE_ID_INVALID | The provided message id is invalid. |
	ErrMessageIdInvalid = status.Error(400, "MESSAGE_ID_INVALID")

	// ErrFileTokenInvalid
	// | 400 | FILE_TOKEN_INVALID | The specified file token is invalid. |
	ErrFileTokenInvalid = status.Error(400, "FILE_TOKEN_INVALID")

	// ErrEmojiNotModified
	// | 400 | EMOJI_NOT_MODIFIED | The theme wasn't changed. |
	ErrEmojiNotModified = status.Error(400, "EMOJI_NOT_MODIFIED")

	// ErrInputTextEmpty
	// | 400 | INPUT_TEXT_EMPTY | The specified text is empty. |
	ErrInputTextEmpty = status.Error(400, "INPUT_TEXT_EMPTY")

	//// ErrUserChannelsTooMuch
	//// | 400 | USER_CHANNELS_TOO_MUCH | One of the users you tried to add is already in too many channels/supergroups. |
	//ErrUserChannelsTooMuch = status.Error(400, "USER_CHANNELS_TOO_MUCH")

	// ErrWebdocumentInvalid
	// | 400 | WEBDOCUMENT_INVALID | Invalid webdocument URL provided. |
	ErrWebdocumentInvalid = status.Error(400, "WEBDOCUMENT_INVALID")

	// ErrAuthTokenAlreadyAccepted
	// | 400 | AUTH_TOKEN_ALREADY_ACCEPTED | The specified auth token was already accepted. |
	ErrAuthTokenAlreadyAccepted = status.Error(400, "AUTH_TOKEN_ALREADY_ACCEPTED")

	// ErrPhoneNumberUnoccupied
	// | 400 | PHONE_NUMBER_UNOCCUPIED | The phone number is not yet being used. |
	ErrPhoneNumberUnoccupied = status.Error(400, "PHONE_NUMBER_UNOCCUPIED")

	// ErrPhotoInvalid
	// | 400 | PHOTO_INVALID | Photo invalid. |
	ErrPhotoInvalid = status.Error(400, "PHOTO_INVALID")

	// ErrPrivacyKeyInvalid
	// | 400 | PRIVACY_KEY_INVALID | The privacy key is invalid. |
	ErrPrivacyKeyInvalid = status.Error(400, "PRIVACY_KEY_INVALID")

	// ErrResultTypeInvalid
	// | 400 | RESULT_TYPE_INVALID | Result type invalid. |
	ErrResultTypeInvalid = status.Error(400, "RESULT_TYPE_INVALID")

	// ErrLangPackInvalid
	// | 400 | LANG_PACK_INVALID | The provided language pack is invalid. |
	ErrLangPackInvalid = status.Error(400, "LANG_PACK_INVALID")

	// ErrUntilDateInvalid
	// | 400 | UNTIL_DATE_INVALID | Invalid until date provided. |
	ErrUntilDateInvalid = status.Error(400, "UNTIL_DATE_INVALID")

	//// ErrChannelPrivate
	//// | 400 | CHANNEL_PRIVATE | You haven't joined this channel/supergroup. |
	//ErrChannelPrivate = status.Error(400, "CHANNEL_PRIVATE")

	// ErrPollOptionDuplicate
	// | 400 | POLL_OPTION_DUPLICATE | Duplicate poll options provided. |
	ErrPollOptionDuplicate = status.Error(400, "POLL_OPTION_DUPLICATE")

	// ErrScheduleBotNotAllowed
	// | 400 | SCHEDULE_BOT_NOT_ALLOWED | Bots cannot schedule messages. |
	ErrScheduleBotNotAllowed = status.Error(400, "SCHEDULE_BOT_NOT_ALLOWED")

	// ErrUserIsBlocked
	// | 400 | USER_IS_BLOCKED | You were blocked by this user. |
	ErrUserIsBlocked = status.Error(400, "USER_IS_BLOCKED")

	//// ErrUserpicUploadRequired
	//// | 400 | USERPIC_UPLOAD_REQUIRED | You must have a profile picture to publish your geolocation. |
	//ErrUserpicUploadRequired = status.Error(400, "USERPIC_UPLOAD_REQUIRED")

	// ErrButtonUrlInvalid
	// | 400 | BUTTON_URL_INVALID | Button URL invalid. |
	ErrButtonUrlInvalid = status.Error(400, "BUTTON_URL_INVALID")

	// ErrMediaPrevInvalid
	// | 400 | MEDIA_PREV_INVALID | Previous media invalid. |
	ErrMediaPrevInvalid = status.Error(400, "MEDIA_PREV_INVALID")

	// ErrResultIdEmpty
	// | 400 | RESULT_ID_EMPTY | Result ID empty. |
	ErrResultIdEmpty = status.Error(400, "RESULT_ID_EMPTY")

	// ErrPhoneCodeInvalid
	// | 400 | PHONE_CODE_INVALID | The provided phone code is invalid. |
	ErrPhoneCodeInvalid = status.Error(400, "PHONE_CODE_INVALID")

	// ErrPhoneNumberOccupied
	// | 400 | PHONE_NUMBER_OCCUPIED | The phone number is already in use. |
	ErrPhoneNumberOccupied = status.Error(400, "PHONE_NUMBER_OCCUPIED")

	// ErrUserBotRequired
	// | 400 | USER_BOT_REQUIRED | This method can only be called by a bot. |
	ErrUserBotRequired = status.Error(400, "USER_BOT_REQUIRED")

	// ErrChatInvitePermanent
	// | 400 | CHAT_INVITE_PERMANENT | You can't set an expiration date on permanent invite links. |
	ErrChatInvitePermanent = status.Error(400, "CHAT_INVITE_PERMANENT")

	// ErrOffsetInvalid
	// | 400 | OFFSET_INVALID | The provided offset is invalid. |
	ErrOffsetInvalid = status.Error(400, "OFFSET_INVALID")

	// ErrGifContentTypeInvalid
	// | 400 | GIF_CONTENT_TYPE_INVALID | GIF content-type invalid. |
	ErrGifContentTypeInvalid = status.Error(400, "GIF_CONTENT_TYPE_INVALID")

	// ErrThemeMimeInvalid
	// | 400 | THEME_MIME_INVALID | The theme's MIME type is invalid. |
	ErrThemeMimeInvalid = status.Error(400, "THEME_MIME_INVALID")

	// ErrOffsetPeerIdInvalid
	// | 400 | OFFSET_PEER_ID_INVALID | The provided offset peer is invalid. |
	ErrOffsetPeerIdInvalid = status.Error(400, "OFFSET_PEER_ID_INVALID")

	// ErrDcIdInvalid
	// | 400 | DC_ID_INVALID | The provided DC ID is invalid. |
	ErrDcIdInvalid = status.Error(400, "DC_ID_INVALID")

	// ErrSettingsInvalid
	// | 400 | SETTINGS_INVALID | Invalid settings were provided. |
	ErrSettingsInvalid = status.Error(400, "SETTINGS_INVALID")

	// ErrBannedRightsInvalid
	// | 400 | BANNED_RIGHTS_INVALID | You provided some invalid flags in the banned rights. |
	ErrBannedRightsInvalid = status.Error(400, "BANNED_RIGHTS_INVALID")

	// ErrWebdocumentSizeTooBig
	// | 400 | WEBDOCUMENT_SIZE_TOO_BIG | Webdocument is too big! |
	ErrWebdocumentSizeTooBig = status.Error(400, "WEBDOCUMENT_SIZE_TOO_BIG")

	// ErrSearchWithLinkNotSupported
	// | 400 | SEARCH_WITH_LINK_NOT_SUPPORTED | You cannot provide a search query and an invite link at the same time. |
	ErrSearchWithLinkNotSupported = status.Error(400, "SEARCH_WITH_LINK_NOT_SUPPORTED")

	// ErrCallProtocolFlagsInvalid
	// | 400 | CALL_PROTOCOL_FLAGS_INVALID | Call protocol flags invalid. |
	ErrCallProtocolFlagsInvalid = status.Error(400, "CALL_PROTOCOL_FLAGS_INVALID")

	// ErrExternalUrlInvalid
	// | 400 | EXTERNAL_URL_INVALID | External URL invalid. |
	ErrExternalUrlInvalid = status.Error(400, "EXTERNAL_URL_INVALID")

	// ErrUserAlreadyInvited
	// | 400 | USER_ALREADY_INVITED | You have already invited this user. |
	ErrUserAlreadyInvited = status.Error(400, "USER_ALREADY_INVITED")

	// ErrAdminRankInvalid
	// | 400 | ADMIN_RANK_INVALID | The specified admin rank is invalid. |
	ErrAdminRankInvalid = status.Error(400, "ADMIN_RANK_INVALID")

	// ErrFilterTitleEmpty
	// | 400 | FILTER_TITLE_EMPTY | The title field of the filter is empty. |
	ErrFilterTitleEmpty = status.Error(400, "FILTER_TITLE_EMPTY")

	// ErrOrderInvalid
	// | 400 | ORDER_INVALID | The specified username order is invalid. |
	ErrOrderInvalid = status.Error(400, "ORDER_INVALID")

	// ErrFilterIncludeEmpty
	// | 400 | FILTER_INCLUDE_EMPTY | The include_peers vector of the filter is empty. |
	ErrFilterIncludeEmpty = status.Error(400, "FILTER_INCLUDE_EMPTY")

	// ErrWebpushTokenInvalid
	// | 400 | WEBPUSH_TOKEN_INVALID | The specified web push token is invalid. |
	ErrWebpushTokenInvalid = status.Error(400, "WEBPUSH_TOKEN_INVALID")

	// ErrApiIdPublishedFlood
	// | 400 | API_ID_PUBLISHED_FLOOD | This API id was published somewhere, you can't use it now. |
	ErrApiIdPublishedFlood = status.Error(400, "API_ID_PUBLISHED_FLOOD")

	// ErrLocationInvalid
	// | 400 | LOCATION_INVALID | The provided location is invalid. |
	ErrLocationInvalid = status.Error(400, "LOCATION_INVALID")

	// ErrPhoneNumberAppSignupForbidden
	// | 400 | PHONE_NUMBER_APP_SIGNUP_FORBIDDEN | You can't sign up using this app. |
	ErrPhoneNumberAppSignupForbidden = status.Error(400, "PHONE_NUMBER_APP_SIGNUP_FORBIDDEN")

	// ErrRandomIdInvalid
	// | 400 | RANDOM_ID_INVALID | A provided random ID is invalid. |
	ErrRandomIdInvalid = status.Error(400, "RANDOM_ID_INVALID")

	// ErrStartParamTooLong
	// | 400 | START_PARAM_TOO_LONG | Start parameter is too long. |
	ErrStartParamTooLong = status.Error(400, "START_PARAM_TOO_LONG")

	// ErrButtonTextInvalid
	// | 400 | BUTTON_TEXT_INVALID | The specified button text is invalid. |
	ErrButtonTextInvalid = status.Error(400, "BUTTON_TEXT_INVALID")

	// ErrPackShortNameOccupied
	// | 400 | PACK_SHORT_NAME_OCCUPIED | A stickerpack with this name already exists. |
	ErrPackShortNameOccupied = status.Error(400, "PACK_SHORT_NAME_OCCUPIED")

	// ErrPhotoContentTypeInvalid
	// | 400 | PHOTO_CONTENT_TYPE_INVALID | Photo mime-type invalid. |
	ErrPhotoContentTypeInvalid = status.Error(400, "PHOTO_CONTENT_TYPE_INVALID")

	// ErrAudioTitleEmpty
	// | 400 | AUDIO_TITLE_EMPTY | An empty audio title was provided. |
	ErrAudioTitleEmpty = status.Error(400, "AUDIO_TITLE_EMPTY")

	// ErrFirstnameInvalid
	// | 400 | FIRSTNAME_INVALID | The first name is invalid. |
	ErrFirstnameInvalid = status.Error(400, "FIRSTNAME_INVALID")

	// ErrResetRequestMissing
	// | 400 | RESET_REQUEST_MISSING | No password reset is in progress. |
	ErrResetRequestMissing = status.Error(400, "RESET_REQUEST_MISSING")

	// ErrJoinAsPeerInvalid
	// | 400 | JOIN_AS_PEER_INVALID | The specified peer cannot be used to join a group call. |
	ErrJoinAsPeerInvalid = status.Error(400, "JOIN_AS_PEER_INVALID")

	// ErrParticipantJoinMissing
	// | 400 | PARTICIPANT_JOIN_MISSING | Trying to enable a presentation, when the user hasn't joined the Video Chat with [phone.joinGroupCall](https://core.telegram.org/method/phone.joinGroupCall). |
	ErrParticipantJoinMissing = status.Error(400, "PARTICIPANT_JOIN_MISSING")

	// ErrSecondsInvalid
	// | 400 | SECONDS_INVALID | Invalid duration provided. |
	ErrSecondsInvalid = status.Error(400, "SECONDS_INVALID")

	// ErrMessagePollClosed
	// | 400 | MESSAGE_POLL_CLOSED | Poll closed. |
	ErrMessagePollClosed = status.Error(400, "MESSAGE_POLL_CLOSED")

	// ErrTitleInvalid
	// | 400 | TITLE_INVALID | The specified stickerpack title is invalid. |
	ErrTitleInvalid = status.Error(400, "TITLE_INVALID")

	// ErrPhotoCropSizeSmall
	// | 400 | PHOTO_CROP_SIZE_SMALL | Photo is too small. |
	ErrPhotoCropSizeSmall = status.Error(400, "PHOTO_CROP_SIZE_SMALL")

	// ErrStickerMimeInvalid
	// | 400 | STICKER_MIME_INVALID | The specified sticker MIME type is invalid. |
	ErrStickerMimeInvalid = status.Error(400, "STICKER_MIME_INVALID")

	// ErrPollQuestionInvalid
	// | 400 | POLL_QUESTION_INVALID | One of the poll questions is not acceptable. |
	ErrPollQuestionInvalid = status.Error(400, "POLL_QUESTION_INVALID")

	// ErrScheduleDateInvalid
	// | 400 | SCHEDULE_DATE_INVALID | Invalid schedule date provided. |
	ErrScheduleDateInvalid = status.Error(400, "SCHEDULE_DATE_INVALID")

	// ErrUsersTooMuch
	// | 400 | USERS_TOO_MUCH | The maximum number of users has been exceeded (to create a chat, for example). |
	ErrUsersTooMuch = status.Error(400, "USERS_TOO_MUCH")

	// ErrGeoPointInvalid
	// | 400 | GEO_POINT_INVALID | Invalid geoposition provided. |
	ErrGeoPointInvalid = status.Error(400, "GEO_POINT_INVALID")

	// ErrStickerFileInvalid
	// | 400 | STICKER_FILE_INVALID | Sticker file invalid. |
	ErrStickerFileInvalid = status.Error(400, "STICKER_FILE_INVALID")

	// ErrEmojiInvalid
	// | 400 | EMOJI_INVALID | The specified theme emoji is valid. |
	ErrEmojiInvalid = status.Error(400, "EMOJI_INVALID")

	// ErrEncryptionAlreadyDeclined
	// | 400 | ENCRYPTION_ALREADY_DECLINED | The secret chat was already declined. |
	ErrEncryptionAlreadyDeclined = status.Error(400, "ENCRYPTION_ALREADY_DECLINED")

	// ErrVoiceMessagesForbidden
	// | 400 | VOICE_MESSAGES_FORBIDDEN | This user's privacy settings forbid you from sending voice messages. |
	ErrVoiceMessagesForbidden = status.Error(400, "VOICE_MESSAGES_FORBIDDEN")

	// ErrLimitInvalid
	// | 400 | LIMIT_INVALID | The provided limit is invalid. |
	ErrLimitInvalid = status.Error(400, "LIMIT_INVALID")

	// ErrUsernamesActiveTooMuch
	// | 400 | USERNAMES_ACTIVE_TOO_MUCH | The maximum number of active usernames was reached. |
	ErrUsernamesActiveTooMuch = status.Error(400, "USERNAMES_ACTIVE_TOO_MUCH")

	// ErrBotResponseTimeout
	// | 400 | BOT_RESPONSE_TIMEOUT | A timeout occurred while fetching data from the bot. |
	ErrBotResponseTimeout = status.Error(400, "BOT_RESPONSE_TIMEOUT")

	// ErrThemeFormatInvalid
	// | 400 | THEME_FORMAT_INVALID | Invalid theme format provided. |
	ErrThemeFormatInvalid = status.Error(400, "THEME_FORMAT_INVALID")

	// ErrWebpushKeyInvalid
	// | 400 | WEBPUSH_KEY_INVALID | The specified web push elliptic curve Diffie-Hellman public key is invalid. |
	ErrWebpushKeyInvalid = status.Error(400, "WEBPUSH_KEY_INVALID")

	// ErrYouBlockedUser
	// | 400 | YOU_BLOCKED_USER | You blocked this user. |
	ErrYouBlockedUser = status.Error(400, "YOU_BLOCKED_USER")

	// ErrCallAlreadyAccepted
	// | 400 | CALL_ALREADY_ACCEPTED | The call was already accepted. |
	ErrCallAlreadyAccepted = status.Error(400, "CALL_ALREADY_ACCEPTED")

	// ErrCurrencyTotalAmountInvalid
	// | 400 | CURRENCY_TOTAL_AMOUNT_INVALID | The total amount of all prices is invalid. |
	ErrCurrencyTotalAmountInvalid = status.Error(400, "CURRENCY_TOTAL_AMOUNT_INVALID")

	// ErrParticipantIdInvalid
	// | 400 | PARTICIPANT_ID_INVALID | The specified participant ID is invalid. |
	ErrParticipantIdInvalid = status.Error(400, "PARTICIPANT_ID_INVALID")

	// ErrConnectionApiIdInvalid
	// | 400 | CONNECTION_API_ID_INVALID | The provided API id is invalid. |
	ErrConnectionApiIdInvalid = status.Error(400, "CONNECTION_API_ID_INVALID")

	// ErrUserAlreadyParticipant
	// | 400 | USER_ALREADY_PARTICIPANT | The user is already in the group. |
	ErrUserAlreadyParticipant = status.Error(400, "USER_ALREADY_PARTICIPANT")

	// ErrUserBot
	// | 400 | USER_BOT | Bots can only be admins in channels. |
	ErrUserBot = status.Error(400, "USER_BOT")

	// ErrUsernameNotOccupied
	// | 400 | USERNAME_NOT_OCCUPIED | The provided username is not occupied. |
	ErrUsernameNotOccupied = status.Error(400, "USERNAME_NOT_OCCUPIED")

	// ErrQuizAnswerMissing
	// | 400 | QUIZ_ANSWER_MISSING | You can forward a quiz while hiding the original author only after choosing an option in the quiz. |
	ErrQuizAnswerMissing = status.Error(400, "QUIZ_ANSWER_MISSING")

	// ErrPasswordHashInvalid
	// | 400 | PASSWORD_HASH_INVALID | The provided password hash is invalid. |
	ErrPasswordHashInvalid = status.Error(400, "PASSWORD_HASH_INVALID")

	// ErrUserAdminInvalid
	// | 400 | USER_ADMIN_INVALID | You're not an admin. |
	ErrUserAdminInvalid = status.Error(400, "USER_ADMIN_INVALID")

	// ErrChatAboutTooLong
	// | 400 | CHAT_ABOUT_TOO_LONG | Chat about too long. |
	ErrChatAboutTooLong = status.Error(400, "CHAT_ABOUT_TOO_LONG")

	// ErrFilterNotSupported
	// | 400 | FILTER_NOT_SUPPORTED | The specified filter cannot be used in this context. |
	ErrFilterNotSupported = status.Error(400, "FILTER_NOT_SUPPORTED")

	// ErrWebdocumentMimeInvalid
	// | 400 | WEBDOCUMENT_MIME_INVALID | Invalid webdocument mime type provided. |
	ErrWebdocumentMimeInvalid = status.Error(400, "WEBDOCUMENT_MIME_INVALID")

	// ErrPhonePasswordProtected
	// | 400 | PHONE_PASSWORD_PROTECTED | This phone is password protected. |
	ErrPhonePasswordProtected = status.Error(400, "PHONE_PASSWORD_PROTECTED")

	// ErrSearchQueryEmpty
	// | 400 | SEARCH_QUERY_EMPTY | The search query is empty. |
	ErrSearchQueryEmpty = status.Error(400, "SEARCH_QUERY_EMPTY")

	// ErrChatRevokeDateUnsupported
	// | 400 | CHAT_REVOKE_DATE_UNSUPPORTED | `min_date` and `max_date` are not available for using with non-user peers. |
	ErrChatRevokeDateUnsupported = status.Error(400, "CHAT_REVOKE_DATE_UNSUPPORTED")

	// ErrPhotoInvalidDimensions
	// | 400 | PHOTO_INVALID_DIMENSIONS | The photo dimensions are invalid. |
	ErrPhotoInvalidDimensions = status.Error(400, "PHOTO_INVALID_DIMENSIONS")

	// ErrAlbumPhotosTooMany
	// | 400 | ALBUM_PHOTOS_TOO_MANY | You have uploaded too many profile photos, delete some before retrying. |
	ErrAlbumPhotosTooMany = status.Error(400, "ALBUM_PHOTOS_TOO_MANY")

	// ErrCallPeerInvalid
	// | 400 | CALL_PEER_INVALID | The provided call peer object is invalid. |
	ErrCallPeerInvalid = status.Error(400, "CALL_PEER_INVALID")

	// ErrPrivacyTooLong
	// | 400 | PRIVACY_TOO_LONG | Too many privacy rules were specified, the current limit is 1000. |
	ErrPrivacyTooLong = status.Error(400, "PRIVACY_TOO_LONG")

	// ErrUserBannedInChannel
	// | 400 | USER_BANNED_IN_CHANNEL | You're banned from sending messages in supergroups/channels. |
	ErrUserBannedInChannel = status.Error(400, "USER_BANNED_IN_CHANNEL")

	// ErrWcConvertUrlInvalid
	// | 400 | WC_CONVERT_URL_INVALID | WC convert URL invalid. |
	ErrWcConvertUrlInvalid = status.Error(400, "WC_CONVERT_URL_INVALID")

	// ErrChatTitleEmpty
	// | 400 | CHAT_TITLE_EMPTY | No chat title provided. |
	ErrChatTitleEmpty = status.Error(400, "CHAT_TITLE_EMPTY")

	// ErrCodeHashInvalid
	// | 400 | CODE_HASH_INVALID | Code hash invalid. |
	ErrCodeHashInvalid = status.Error(400, "CODE_HASH_INVALID")

	// ErrFilterIdInvalid
	// | 400 | FILTER_ID_INVALID | The specified filter ID is invalid. |
	ErrFilterIdInvalid = status.Error(400, "FILTER_ID_INVALID")

	// ErrPasswordRequired
	// | 400 | PASSWORD_REQUIRED | A [2FA password](https://core.telegram.org/api/srp) must be configured to use Telegram Passport. |
	ErrPasswordRequired = status.Error(400, "PASSWORD_REQUIRED")

	// ErrReactionsTooMany
	// | 400 | REACTIONS_TOO_MANY | The message already has exactly `reactions_uniq_max` reaction emojis, you can't react with a new emoji, see [the docs for more info &raquo;](/api/config#client-configuration). |
	ErrReactionsTooMany = status.Error(400, "REACTIONS_TOO_MANY")

	// ErrTokenEmpty
	// | 400 | TOKEN_EMPTY | The specified token is empty. |
	ErrTokenEmpty = status.Error(400, "TOKEN_EMPTY")

	// ErrChatForwardsRestricted
	// | 400 | CHAT_FORWARDS_RESTRICTED | You can't forward messages from a protected chat. |
	ErrChatForwardsRestricted = status.Error(400, "CHAT_FORWARDS_RESTRICTED")

	// ErrEmoticonStickerpackMissing
	// | 400 | EMOTICON_STICKERPACK_MISSING | inputStickerSetDice.emoji cannot be empty. |
	ErrEmoticonStickerpackMissing = status.Error(400, "EMOTICON_STICKERPACK_MISSING")

	// ErrGraphInvalidReload
	// | 400 | GRAPH_INVALID_RELOAD | Invalid graph token provided, please reload the stats and provide the updated token. |
	ErrGraphInvalidReload = status.Error(400, "GRAPH_INVALID_RELOAD")

	// ErrLinkNotModified
	// | 400 | LINK_NOT_MODIFIED | Discussion link not modified. |
	ErrLinkNotModified = status.Error(400, "LINK_NOT_MODIFIED")

	// ErrPinRestricted
	// | 400 | PIN_RESTRICTED | You can't pin messages. |
	ErrPinRestricted = status.Error(400, "PIN_RESTRICTED")

	// ErrSha256HashInvalid
	// | 400 | SHA256_HASH_INVALID | The provided SHA256 hash is invalid. |
	ErrSha256HashInvalid = status.Error(400, "SHA256_HASH_INVALID")

	// ErrUserVolumeInvalid
	// | 400 | USER_VOLUME_INVALID | The specified user volume is invalid. |
	ErrUserVolumeInvalid = status.Error(400, "USER_VOLUME_INVALID")

	// ErrUserNotParticipant
	// | 400 | USER_NOT_PARTICIPANT | You're not a member of this supergroup/channel. |
	ErrUserNotParticipant = status.Error(400, "USER_NOT_PARTICIPANT")

	// ErrBotMissing
	// | 400 | BOT_MISSING | Only bots can call this method, please use [@stickers](https://t.me/stickers) if you're a user. |
	ErrBotMissing = status.Error(400, "BOT_MISSING")

	// ErrExportCardInvalid
	// | 400 | EXPORT_CARD_INVALID | Provided card is invalid. |
	ErrExportCardInvalid = status.Error(400, "EXPORT_CARD_INVALID")

	// ErrGraphExpiredReload
	// | 400 | GRAPH_EXPIRED_RELOAD | This graph has expired, please obtain a new graph token. |
	ErrGraphExpiredReload = status.Error(400, "GRAPH_EXPIRED_RELOAD")

	//// ErrChatSendInlineForbidden
	//// | 400 | CHAT_SEND_INLINE_FORBIDDEN | You can't send inline messages in this group. |
	//ErrChatSendInlineForbidden = status.Error(400, "CHAT_SEND_INLINE_FORBIDDEN")

	// ErrDhGAInvalid
	// | 400 | DH_G_A_INVALID | g_a invalid. |
	ErrDhGAInvalid = status.Error(400, "DH_G_A_INVALID")

	// ErrMaxIdInvalid
	// | 400 | MAX_ID_INVALID | The provided max ID is invalid. |
	ErrMaxIdInvalid = status.Error(400, "MAX_ID_INVALID")

	// ErrReactionEmpty
	// | 400 | REACTION_EMPTY | Empty reaction provided. |
	ErrReactionEmpty = status.Error(400, "REACTION_EMPTY")

	// ErrMediaInvalid
	// | 400 | MEDIA_INVALID | Media invalid. |
	ErrMediaInvalid = status.Error(400, "MEDIA_INVALID")

	// ErrParticipantsTooFew
	// | 400 | PARTICIPANTS_TOO_FEW | Not enough participants. |
	ErrParticipantsTooFew = status.Error(400, "PARTICIPANTS_TOO_FEW")

	// ErrFilePartsInvalid
	// | 400 | FILE_PARTS_INVALID | The number of file parts is invalid. |
	ErrFilePartsInvalid = status.Error(400, "FILE_PARTS_INVALID")

	// ErrChatAdminRequired
	// | 400 | CHAT_ADMIN_REQUIRED | You must be an admin in this chat to do this. |
	ErrChatAdminRequired = status.Error(400, "CHAT_ADMIN_REQUIRED")

	// ErrContactNameEmpty
	// | 400 | CONTACT_NAME_EMPTY | Contact name empty. |
	ErrContactNameEmpty = status.Error(400, "CONTACT_NAME_EMPTY")

	// ErrTokenTypeInvalid
	// | 400 | TOKEN_TYPE_INVALID | The specified token type is invalid. |
	ErrTokenTypeInvalid = status.Error(400, "TOKEN_TYPE_INVALID")

	// ErrPersistentTimestampEmpty
	// | 400 | PERSISTENT_TIMESTAMP_EMPTY | Persistent timestamp empty. |
	ErrPersistentTimestampEmpty = status.Error(400, "PERSISTENT_TIMESTAMP_EMPTY")

	// ErrSlowmodeMultiMsgsDisabled
	// | 400 | SLOWMODE_MULTI_MSGS_DISABLED | Slowmode is enabled, you cannot forward multiple messages to this group. |
	ErrSlowmodeMultiMsgsDisabled = status.Error(400, "SLOWMODE_MULTI_MSGS_DISABLED")

	// ErrMediaGroupedInvalid
	// | 400 | MEDIA_GROUPED_INVALID | You tried to send media of different types in an album. |
	ErrMediaGroupedInvalid = status.Error(400, "MEDIA_GROUPED_INVALID")

	// ErrPackShortNameInvalid
	// | 400 | PACK_SHORT_NAME_INVALID | Short pack name invalid. |
	ErrPackShortNameInvalid = status.Error(400, "PACK_SHORT_NAME_INVALID")

	// ErrPhotoExtInvalid
	// | 400 | PHOTO_EXT_INVALID | The extension of the photo is invalid. |
	ErrPhotoExtInvalid = status.Error(400, "PHOTO_EXT_INVALID")

	// ErrStickerThumbPngNopng
	// | 400 | STICKER_THUMB_PNG_NOPNG | Incorrect stickerset thumb file provided, PNG / WEBP expected. |
	ErrStickerThumbPngNopng = status.Error(400, "STICKER_THUMB_PNG_NOPNG")

	// ErrThemeFileInvalid
	// | 400 | THEME_FILE_INVALID | Invalid theme file provided. |
	ErrThemeFileInvalid = status.Error(400, "THEME_FILE_INVALID")

	//// ErrCallOccupyFailed
	//// | 400 | CALL_OCCUPY_FAILED | The call failed because the user is already making another call. |
	//ErrCallOccupyFailed = status.Error(400, "CALL_OCCUPY_FAILED")

	// ErrChatDiscussionUnallowed
	// | 400 | CHAT_DISCUSSION_UNALLOWED | You can't enable forum topics in a discussion group linked to a channel. |
	ErrChatDiscussionUnallowed = status.Error(400, "CHAT_DISCUSSION_UNALLOWED")

	// ErrPasswordRecoveryNa
	// | 400 | PASSWORD_RECOVERY_NA | No email was set, can't recover password via email. |
	ErrPasswordRecoveryNa = status.Error(400, "PASSWORD_RECOVERY_NA")

	//// ErrTakeoutRequired
	//// | 400 | TAKEOUT_REQUIRED | A takeout session has to be initialized, first. |
	//ErrTakeoutRequired = status.Error(400, "TAKEOUT_REQUIRED")

	// ErrThemeTitleInvalid
	// | 400 | THEME_TITLE_INVALID | The specified theme title is invalid. |
	ErrThemeTitleInvalid = status.Error(400, "THEME_TITLE_INVALID")

	// ErrEmojiMarkupInvalid
	// | 400 | EMOJI_MARKUP_INVALID | The specified `video_emoji_markup` was invalid. |
	ErrEmojiMarkupInvalid = status.Error(400, "EMOJI_MARKUP_INVALID")

	// ErrVideoContentTypeInvalid
	// | 400 | VIDEO_CONTENT_TYPE_INVALID | The video's content type is invalid. |
	ErrVideoContentTypeInvalid = status.Error(400, "VIDEO_CONTENT_TYPE_INVALID")

	// ErrFilePartInvalid
	// | 400 | FILE_PART_INVALID | The file part number is invalid. |
	ErrFilePartInvalid = status.Error(400, "FILE_PART_INVALID")

	// ErrAccessTokenExpired
	// | 400 | ACCESS_TOKEN_EXPIRED | Access token expired. |
	ErrAccessTokenExpired = status.Error(400, "ACCESS_TOKEN_EXPIRED")

	// ErrAdminRankEmojiNotAllowed
	// | 400 | ADMIN_RANK_EMOJI_NOT_ALLOWED | An admin rank cannot contain emojis. |
	ErrAdminRankEmojiNotAllowed = status.Error(400, "ADMIN_RANK_EMOJI_NOT_ALLOWED")

	// ErrButtonTypeInvalid
	// | 400 | BUTTON_TYPE_INVALID | The type of one or more of the buttons you provided is invalid. |
	ErrButtonTypeInvalid = status.Error(400, "BUTTON_TYPE_INVALID")

	// ErrStickerTgsNotgs
	// | 400 | STICKER_TGS_NOTGS | Invalid TGS sticker provided. |
	ErrStickerTgsNotgs = status.Error(400, "STICKER_TGS_NOTGS")

	//// ErrUserNotMutualContact
	//// | 400 | USER_NOT_MUTUAL_CONTACT | The provided user is not a mutual contact. |
	//ErrUserNotMutualContact = status.Error(400, "USER_NOT_MUTUAL_CONTACT")

	// ErrInputFilterInvalid
	// | 400 | INPUT_FILTER_INVALID | The specified filter is invalid. |
	ErrInputFilterInvalid = status.Error(400, "INPUT_FILTER_INVALID")

	// ErrMediaNewInvalid
	// | 400 | MEDIA_NEW_INVALID | The new media is invalid. |
	ErrMediaNewInvalid = status.Error(400, "MEDIA_NEW_INVALID")

	// ErrAudioContentUrlEmpty
	// | 400 | AUDIO_CONTENT_URL_EMPTY | The remote URL specified in the content field is empty. |
	ErrAudioContentUrlEmpty = status.Error(400, "AUDIO_CONTENT_URL_EMPTY")

	// ErrPollOptionInvalid
	// | 400 | POLL_OPTION_INVALID | Invalid poll option provided. |
	ErrPollOptionInvalid = status.Error(400, "POLL_OPTION_INVALID")

	// ErrResultIdInvalid
	// | 400 | RESULT_ID_INVALID | One of the specified result IDs is invalid. |
	ErrResultIdInvalid = status.Error(400, "RESULT_ID_INVALID")

	// ErrGroupedMediaInvalid
	// | 400 | GROUPED_MEDIA_INVALID | Invalid grouped media. |
	ErrGroupedMediaInvalid = status.Error(400, "GROUPED_MEDIA_INVALID")

	// ErrTtlDaysInvalid
	// | 400 | TTL_DAYS_INVALID | The provided TTL is invalid. |
	ErrTtlDaysInvalid = status.Error(400, "TTL_DAYS_INVALID")

	// ErrGraphOutdatedReload
	// | 400 | GRAPH_OUTDATED_RELOAD | The graph is outdated, please get a new async token using stats.getBroadcastStats. |
	ErrGraphOutdatedReload = status.Error(400, "GRAPH_OUTDATED_RELOAD")

	// ErrCallAlreadyDeclined
	// | 400 | CALL_ALREADY_DECLINED | The call was already declined. |
	ErrCallAlreadyDeclined = status.Error(400, "CALL_ALREADY_DECLINED")

	// ErrScheduleDateTooLate
	// | 400 | SCHEDULE_DATE_TOO_LATE | You can't schedule a message this far in the future. |
	ErrScheduleDateTooLate = status.Error(400, "SCHEDULE_DATE_TOO_LATE")

	// ErrSmsCodeCreateFailed
	// | 400 | SMS_CODE_CREATE_FAILED | An error occurred while creating the SMS code. |
	ErrSmsCodeCreateFailed = status.Error(400, "SMS_CODE_CREATE_FAILED")

	// ErrPollAnswerInvalid
	// | 400 | POLL_ANSWER_INVALID | One of the poll answers is not acceptable. |
	ErrPollAnswerInvalid = status.Error(400, "POLL_ANSWER_INVALID")

	// ErrStickerGifDimensions
	// | 400 | STICKER_GIF_DIMENSIONS | The specified video sticker has invalid dimensions. |
	ErrStickerGifDimensions = status.Error(400, "STICKER_GIF_DIMENSIONS")

	//// ErrUserInvalid
	//// | 400 | USER_INVALID | Invalid user provided. |
	//ErrUserInvalid = status.Error(400, "USER_INVALID")

	// ErrWebpageCurlFailed
	// | 400 | WEBPAGE_CURL_FAILED | Failure while fetching the webpage with cURL. |
	ErrWebpageCurlFailed = status.Error(400, "WEBPAGE_CURL_FAILED")

	// ErrBotPaymentsDisabled
	// | 400 | BOT_PAYMENTS_DISABLED | Please enable bot payments in botfather before calling this method. |
	ErrBotPaymentsDisabled = status.Error(400, "BOT_PAYMENTS_DISABLED")

	// ErrChatNotModified
	// | 400 | CHAT_NOT_MODIFIED | The pinned message wasn't modified. |
	ErrChatNotModified = status.Error(400, "CHAT_NOT_MODIFIED")

	// ErrTmpPasswordDisabled
	// | 400 | TMP_PASSWORD_DISABLED | The temporary password is disabled. |
	ErrTmpPasswordDisabled = status.Error(400, "TMP_PASSWORD_DISABLED")

	// ErrStickerVideoNowebm
	// | 400 | STICKER_VIDEO_NOWEBM | The specified video sticker is not in webm format. |
	ErrStickerVideoNowebm = status.Error(400, "STICKER_VIDEO_NOWEBM")

	// ErrAuthTokenExpired
	// | 400 | AUTH_TOKEN_EXPIRED | The authorization token has expired. |
	ErrAuthTokenExpired = status.Error(400, "AUTH_TOKEN_EXPIRED")

	// ErrEncryptionDeclined
	// | 400 | ENCRYPTION_DECLINED | The secret chat was declined. |
	ErrEncryptionDeclined = status.Error(400, "ENCRYPTION_DECLINED")

	// ErrFilePartEmpty
	// | 400 | FILE_PART_EMPTY | The provided file part is empty. |
	ErrFilePartEmpty = status.Error(400, "FILE_PART_EMPTY")

	// ErrFilePartTooBig
	// | 400 | FILE_PART_TOO_BIG | The uploaded file part is too big. |
	ErrFilePartTooBig = status.Error(400, "FILE_PART_TOO_BIG")

	// ErrUsernameOccupied
	// | 400 | USERNAME_OCCUPIED | The provided username is already occupied. |
	ErrUsernameOccupied = status.Error(400, "USERNAME_OCCUPIED")

	// ErrBotInlineDisabled
	// | 400 | BOT_INLINE_DISABLED | This bot can't be used in inline mode. |
	ErrBotInlineDisabled = status.Error(400, "BOT_INLINE_DISABLED")

	// ErrEmailHashExpired
	// | 400 | EMAIL_HASH_EXPIRED | Email hash expired. |
	ErrEmailHashExpired = status.Error(400, "EMAIL_HASH_EXPIRED")

	// ErrStickerVideoNodoc
	// | 400 | STICKER_VIDEO_NODOC | You must send the video sticker as a document. |
	ErrStickerVideoNodoc = status.Error(400, "STICKER_VIDEO_NODOC")

	//// ErrTopicDeleted
	//// | 400 | TOPIC_DELETED | The specified topic was deleted. |
	//ErrTopicDeleted = status.Error(400, "TOPIC_DELETED")

	// ErrFileReferenceExpired
	// | 400 | FILE_REFERENCE_EXPIRED | File reference expired, it must be refetched as described in [the documentation](https://core.telegram.org/api/file_reference). |
	ErrFileReferenceExpired = status.Error(400, "FILE_REFERENCE_EXPIRED")

	// ErrMessageEmpty
	// | 400 | MESSAGE_EMPTY | The provided message is empty. |
	ErrMessageEmpty = status.Error(400, "MESSAGE_EMPTY")

	// ErrNewSettingsEmpty
	// | 400 | NEW_SETTINGS_EMPTY | No password is set on the current account, and no new password was specified in `new_settings`. |
	ErrNewSettingsEmpty = status.Error(400, "NEW_SETTINGS_EMPTY")

	// ErrBroadcastRequired
	// | 400 | BROADCAST_REQUIRED | This method can only be called on a channel, please use stats.getMegagroupStats for supergroups. |
	ErrBroadcastRequired = status.Error(400, "BROADCAST_REQUIRED")

	// ErrLangCodeNotSupported
	// | 400 | LANG_CODE_NOT_SUPPORTED | The specified language code is not supported. |
	ErrLangCodeNotSupported = status.Error(400, "LANG_CODE_NOT_SUPPORTED")

	// ErrVideoFileInvalid
	// | 400 | VIDEO_FILE_INVALID | The specified video file is invalid. |
	ErrVideoFileInvalid = status.Error(400, "VIDEO_FILE_INVALID")

	// ErrReplyMarkupBuyEmpty
	// | 400 | REPLY_MARKUP_BUY_EMPTY | Reply markup for buy button empty. |
	ErrReplyMarkupBuyEmpty = status.Error(400, "REPLY_MARKUP_BUY_EMPTY")

	// ErrWallpaperMimeInvalid
	// | 400 | WALLPAPER_MIME_INVALID | The specified wallpaper MIME type is invalid. |
	ErrWallpaperMimeInvalid = status.Error(400, "WALLPAPER_MIME_INVALID")

	// ErrOptionsTooMuch
	// | 400 | OPTIONS_TOO_MUCH | Too many options provided. |
	ErrOptionsTooMuch = status.Error(400, "OPTIONS_TOO_MUCH")

	// ErrPasswordEmpty
	// | 400 | PASSWORD_EMPTY | The provided password is empty. |
	ErrPasswordEmpty = status.Error(400, "PASSWORD_EMPTY")

	// ErrPhoneCodeExpired
	// | 400 | PHONE_CODE_EXPIRED | The phone code you provided has expired. |
	ErrPhoneCodeExpired = status.Error(400, "PHONE_CODE_EXPIRED")

	// ErrTaskAlreadyExists
	// | 400 | TASK_ALREADY_EXISTS | An email reset was already requested. |
	ErrTaskAlreadyExists = status.Error(400, "TASK_ALREADY_EXISTS")

	// ErrShortNameInvalid
	// | 400 | SHORT_NAME_INVALID | The specified short name is invalid. |
	ErrShortNameInvalid = status.Error(400, "SHORT_NAME_INVALID")

	// ErrChannelIdInvalid
	// | 400 | CHANNEL_ID_INVALID | The specified supergroup ID is invalid. |
	ErrChannelIdInvalid = status.Error(400, "CHANNEL_ID_INVALID")

	// ErrContactReqMissing
	// | 400 | CONTACT_REQ_MISSING | Missing contact request. |
	ErrContactReqMissing = status.Error(400, "CONTACT_REQ_MISSING")

	// ErrResultIdDuplicate
	// | 400 | RESULT_ID_DUPLICATE | You provided a duplicate result ID. |
	ErrResultIdDuplicate = status.Error(400, "RESULT_ID_DUPLICATE")

	//// ErrUserBotInvalid
	//// | 400 | USER_BOT_INVALID | User accounts must provide the `bot` method parameter when calling this method. If there is no such method parameter, this method can only be invoked by bot accounts. |
	//ErrUserBotInvalid = status.Error(400, "USER_BOT_INVALID")

	// ErrPhoneCodeHashEmpty
	// | 400 | PHONE_CODE_HASH_EMPTY | phone_code_hash is missing. |
	ErrPhoneCodeHashEmpty = status.Error(400, "PHONE_CODE_HASH_EMPTY")

	// ErrAuthTokenInvalid
	// | 400 | AUTH_TOKEN_INVALID | The specified auth token is invalid. |
	ErrAuthTokenInvalid = status.Error(400, "AUTH_TOKEN_INVALID")

	// ErrDocumentInvalid
	// | 400 | DOCUMENT_INVALID | The specified document is invalid. |
	ErrDocumentInvalid = status.Error(400, "DOCUMENT_INVALID")

	// ErrFilePartSizeChanged
	// | 400 | FILE_PART_SIZE_CHANGED | Provided file part size has changed. |
	ErrFilePartSizeChanged = status.Error(400, "FILE_PART_SIZE_CHANGED")

	// ErrLangCodeInvalid
	// | 400 | LANG_CODE_INVALID | The specified language code is invalid. |
	ErrLangCodeInvalid = status.Error(400, "LANG_CODE_INVALID")

	// ErrStickerTgsNodoc
	// | 400 | STICKER_TGS_NODOC | You must send the animated sticker as a document. |
	ErrStickerTgsNodoc = status.Error(400, "STICKER_TGS_NODOC")

	// ErrFolderIdEmpty
	// | 400 | FOLDER_ID_EMPTY | An empty folder ID was specified. |
	ErrFolderIdEmpty = status.Error(400, "FOLDER_ID_EMPTY")

	// ErrFromPeerInvalid
	// | 400 | FROM_PEER_INVALID | The specified from_id is invalid. |
	ErrFromPeerInvalid = status.Error(400, "FROM_PEER_INVALID")

	// ErrErrorTextEmpty
	// | 400 | ERROR_TEXT_EMPTY | The provided error message is empty. |
	ErrErrorTextEmpty = status.Error(400, "ERROR_TEXT_EMPTY")

	// ErrFileReferenceEmpty
	// | 400 | FILE_REFERENCE_EMPTY | An empty [file reference](https://core.telegram.org/api/file_reference) was specified. |
	ErrFileReferenceEmpty = status.Error(400, "FILE_REFERENCE_EMPTY")

	// ErrMediaEmpty
	// | 400 | MEDIA_EMPTY | The provided media object is invalid. |
	ErrMediaEmpty = status.Error(400, "MEDIA_EMPTY")

	// ErrRevoteNotAllowed
	// | 400 | REVOTE_NOT_ALLOWED | You cannot change your vote. |
	ErrRevoteNotAllowed = status.Error(400, "REVOTE_NOT_ALLOWED")

	// ErrConnectionAppVersionEmpty
	// | 400 | CONNECTION_APP_VERSION_EMPTY | App version is empty. |
	ErrConnectionAppVersionEmpty = status.Error(400, "CONNECTION_APP_VERSION_EMPTY")

	// ErrStickerDocumentInvalid
	// | 400 | STICKER_DOCUMENT_INVALID | The specified sticker document is invalid. |
	ErrStickerDocumentInvalid = status.Error(400, "STICKER_DOCUMENT_INVALID")

	// ErrUsernamePurchaseAvailable
	// | 400 | USERNAME_PURCHASE_AVAILABLE | The specified username can be purchased on https://fragment.com. |
	ErrUsernamePurchaseAvailable = status.Error(400, "USERNAME_PURCHASE_AVAILABLE")

	// ErrPrivacyValueInvalid
	// | 400 | PRIVACY_VALUE_INVALID | The specified privacy rule combination is invalid. |
	ErrPrivacyValueInvalid = status.Error(400, "PRIVACY_VALUE_INVALID")

	// ErrRandomIdEmpty
	// | 400 | RANDOM_ID_EMPTY | Random ID empty. |
	ErrRandomIdEmpty = status.Error(400, "RANDOM_ID_EMPTY")

	// ErrBotInvalid
	// | 400 | BOT_INVALID | This is not a valid bot. |
	ErrBotInvalid = status.Error(400, "BOT_INVALID")

	// ErrTtlMediaInvalid
	// | 400 | TTL_MEDIA_INVALID | Invalid media Time To Live was provided. |
	ErrTtlMediaInvalid = status.Error(400, "TTL_MEDIA_INVALID")

	// ErrBotChannelsNa
	// | 400 | BOT_CHANNELS_NA | Bots can't edit admin privileges. |
	ErrBotChannelsNa = status.Error(400, "BOT_CHANNELS_NA")

	// ErrChannelInvalid
	// | 400 | CHANNEL_INVALID | The provided channel is invalid. |
	ErrChannelInvalid = status.Error(400, "CHANNEL_INVALID")

	// ErrMaxDateInvalid
	// | 400 | MAX_DATE_INVALID | The specified maximum date is invalid. |
	ErrMaxDateInvalid = status.Error(400, "MAX_DATE_INVALID")

	// ErrQueryTooShort
	// | 400 | QUERY_TOO_SHORT | The query string is too short. |
	ErrQueryTooShort = status.Error(400, "QUERY_TOO_SHORT")

	// ErrQuizCorrectAnswerInvalid
	// | 400 | QUIZ_CORRECT_ANSWER_INVALID | An invalid value was provided to the correct_answers field. |
	ErrQuizCorrectAnswerInvalid = status.Error(400, "QUIZ_CORRECT_ANSWER_INVALID")

	// ErrImageProcessFailed
	// | 400 | IMAGE_PROCESS_FAILED | Failure while processing image. |
	ErrImageProcessFailed = status.Error(400, "IMAGE_PROCESS_FAILED")

	// ErrPersistentTimestampInvalid
	// | 400 | PERSISTENT_TIMESTAMP_INVALID | Persistent timestamp invalid. |
	ErrPersistentTimestampInvalid = status.Error(400, "PERSISTENT_TIMESTAMP_INVALID")

	// ErrPhotoIdInvalid
	// | 400 | PHOTO_ID_INVALID | Photo ID invalid. |
	ErrPhotoIdInvalid = status.Error(400, "PHOTO_ID_INVALID")

	// ErrPhotoThumbUrlEmpty
	// | 400 | PHOTO_THUMB_URL_EMPTY | Photo thumbnail URL is empty. |
	ErrPhotoThumbUrlEmpty = status.Error(400, "PHOTO_THUMB_URL_EMPTY")

	// ErrAccessTokenInvalid
	// | 400 | ACCESS_TOKEN_INVALID | Access token invalid. |
	ErrAccessTokenInvalid = status.Error(400, "ACCESS_TOKEN_INVALID")

	// ErrMaxQtsInvalid
	// | 400 | MAX_QTS_INVALID | The specified max_qts is invalid. |
	ErrMaxQtsInvalid = status.Error(400, "MAX_QTS_INVALID")

	// ErrPaymentProviderInvalid
	// | 400 | PAYMENT_PROVIDER_INVALID | The specified payment provider is invalid. |
	ErrPaymentProviderInvalid = status.Error(400, "PAYMENT_PROVIDER_INVALID")

	// ErrInviteRequestSent
	// | 400 | INVITE_REQUEST_SENT | You have successfully requested to join this chat or channel. |
	ErrInviteRequestSent = status.Error(400, "INVITE_REQUEST_SENT")

	// ErrPhoneNumberBanned
	// | 400 | PHONE_NUMBER_BANNED | The provided phone number is banned from telegram. |
	ErrPhoneNumberBanned = status.Error(400, "PHONE_NUMBER_BANNED")

	// ErrButtonUserPrivacyRestricted
	// | 400 | BUTTON_USER_PRIVACY_RESTRICTED | The privacy setting of the user specified in a [inputKeyboardButtonUserProfile](/constructor/inputKeyboardButtonUserProfile) button do not allow creating such a button. |
	ErrButtonUserPrivacyRestricted = status.Error(400, "BUTTON_USER_PRIVACY_RESTRICTED")

	// ErrInputUserDeactivated
	// | 400 | INPUT_USER_DEACTIVATED | The specified user was deleted. |
	ErrInputUserDeactivated = status.Error(400, "INPUT_USER_DEACTIVATED")

	// ErrParticipantVersionOutdated
	// | 400 | PARTICIPANT_VERSION_OUTDATED | The other participant does not use an up to date telegram client with support for calls. |
	ErrParticipantVersionOutdated = status.Error(400, "PARTICIPANT_VERSION_OUTDATED")

	// ErrChannelForumMissing
	// | 400 | CHANNEL_FORUM_MISSING | This supergroup is not a forum. |
	ErrChannelForumMissing = status.Error(400, "CHANNEL_FORUM_MISSING")

	// ErrContactAddMissing
	// | 400 | CONTACT_ADD_MISSING | Contact to add is missing. |
	ErrContactAddMissing = status.Error(400, "CONTACT_ADD_MISSING")

	// ErrLastnameInvalid
	// | 400 | LASTNAME_INVALID | The last name is invalid. |
	ErrLastnameInvalid = status.Error(400, "LASTNAME_INVALID")

	//// ErrInviteHashExpired
	//// | 400 | INVITE_HASH_EXPIRED | The invite link has expired. |
	//ErrInviteHashExpired = status.Error(400, "INVITE_HASH_EXPIRED")

	// ErrTokenInvalid
	// | 400 | TOKEN_INVALID | The provided token is invalid. |
	ErrTokenInvalid = status.Error(400, "TOKEN_INVALID")

	// ErrChatRestricted
	// | 400 | CHAT_RESTRICTED | You can't send messages in this chat, you were restricted. |
	ErrChatRestricted = status.Error(400, "CHAT_RESTRICTED")

	// ErrReplyMarkupInvalid
	// | 400 | REPLY_MARKUP_INVALID | The provided reply markup is invalid. |
	ErrReplyMarkupInvalid = status.Error(400, "REPLY_MARKUP_INVALID")

	// ErrWebdocumentUrlInvalid
	// | 400 | WEBDOCUMENT_URL_INVALID | The specified webdocument URL is invalid. |
	ErrWebdocumentUrlInvalid = status.Error(400, "WEBDOCUMENT_URL_INVALID")

	// ErrBotCommandDescriptionInvalid
	// | 400 | BOT_COMMAND_DESCRIPTION_INVALID | The specified command description is invalid. |
	ErrBotCommandDescriptionInvalid = status.Error(400, "BOT_COMMAND_DESCRIPTION_INVALID")

	// ErrChannelParicipantMissing
	// | 400 | CHANNEL_PARICIPANT_MISSING | The current user is not in the channel. |
	ErrChannelParicipantMissing = status.Error(400, "CHANNEL_PARICIPANT_MISSING")

	// ErrEntityMentionUserInvalid
	// | 400 | ENTITY_MENTION_USER_INVALID | You mentioned an invalid user. |
	ErrEntityMentionUserInvalid = status.Error(400, "ENTITY_MENTION_USER_INVALID")

	// ErrPasswordRecoveryExpired
	// | 400 | PASSWORD_RECOVERY_EXPIRED | The recovery code has expired. |
	ErrPasswordRecoveryExpired = status.Error(400, "PASSWORD_RECOVERY_EXPIRED")

	// ErrTtlPeriodInvalid
	// | 400 | TTL_PERIOD_INVALID | The specified TTL period is invalid. |
	ErrTtlPeriodInvalid = status.Error(400, "TTL_PERIOD_INVALID")

	// ErrUsersTooFew
	// | 400 | USERS_TOO_FEW | Not enough users (to create a chat, for example). |
	ErrUsersTooFew = status.Error(400, "USERS_TOO_FEW")

	// ErrMsgIdInvalid
	// | 400 | MSG_ID_INVALID | Invalid message ID provided. |
	ErrMsgIdInvalid = status.Error(400, "MSG_ID_INVALID")

	// ErrPackTitleInvalid
	// | 400 | PACK_TITLE_INVALID | The stickerpack title is invalid. |
	ErrPackTitleInvalid = status.Error(400, "PACK_TITLE_INVALID")

	// ErrEntityBoundsInvalid
	// | 400 | ENTITY_BOUNDS_INVALID | A specified [entity offset or length](/api/entities#entity-length) is invalid, see [here &raquo;](/api/entities#entity-length) for info on how to properly compute the entity offset/length. |
	ErrEntityBoundsInvalid = status.Error(400, "ENTITY_BOUNDS_INVALID")

	// ErrNextOffsetInvalid
	// | 400 | NEXT_OFFSET_INVALID | The specified offset is longer than 64 bytes. |
	ErrNextOffsetInvalid = status.Error(400, "NEXT_OFFSET_INVALID")

	// ErrUserBlocked
	// | 400 | USER_BLOCKED | User blocked. |
	ErrUserBlocked = status.Error(400, "USER_BLOCKED")

	// ErrBotAppInvalid
	// | 400 | BOT_APP_INVALID | The specified bot app is invalid. |
	ErrBotAppInvalid = status.Error(400, "BOT_APP_INVALID")

	// ErrBotsTooMuch
	// | 400 | BOTS_TOO_MUCH | There are too many bots in this chat/channel. |
	ErrBotsTooMuch = status.Error(400, "BOTS_TOO_MUCH")

	// ErrFilePartSizeInvalid
	// | 400 | FILE_PART_SIZE_INVALID | The provided file part size is invalid. |
	ErrFilePartSizeInvalid = status.Error(400, "FILE_PART_SIZE_INVALID")

	// ErrInviteForbiddenWithJoinas
	// | 400 | INVITE_FORBIDDEN_WITH_JOINAS | If the user has anonymously joined a group call as a channel, they can't invite other users to the group call because that would cause deanonymization, because the invite would be sent using the original user ID, not the anonymized channel ID. |
	ErrInviteForbiddenWithJoinas = status.Error(400, "INVITE_FORBIDDEN_WITH_JOINAS")

	// ErrMessageTooLong
	// | 400 | MESSAGE_TOO_LONG | The provided message is too long. |
	ErrMessageTooLong = status.Error(400, "MESSAGE_TOO_LONG")

	// ErrStartParamInvalid
	// | 400 | START_PARAM_INVALID | Start parameter invalid. |
	ErrStartParamInvalid = status.Error(400, "START_PARAM_INVALID")

	// ErrChatTooBig
	// | 400 | CHAT_TOO_BIG | This method is not available for groups with more than `chat_read_mark_size_threshold` members, [see client configuration &raquo;](https://core.telegram.org/api/config#client-configuration). |
	ErrChatTooBig = status.Error(400, "CHAT_TOO_BIG")

	// ErrStickerEmojiInvalid
	// | 400 | STICKER_EMOJI_INVALID | Sticker emoji invalid. |
	ErrStickerEmojiInvalid = status.Error(400, "STICKER_EMOJI_INVALID")

	// ErrWallpaperInvalid
	// | 400 | WALLPAPER_INVALID | The specified wallpaper is invalid. |
	ErrWallpaperInvalid = status.Error(400, "WALLPAPER_INVALID")

	// ErrTempAuthKeyAlreadyBound
	// | 400 | TEMP_AUTH_KEY_ALREADY_BOUND | The passed temporary key is already bound to another **perm_auth_key_id**. |
	ErrTempAuthKeyAlreadyBound = status.Error(400, "TEMP_AUTH_KEY_ALREADY_BOUND")

	// ErrChatPublicRequired
	// | 400 | CHAT_PUBLIC_REQUIRED | You can only enable join requests in public groups. |
	ErrChatPublicRequired = status.Error(400, "CHAT_PUBLIC_REQUIRED")

	// ErrPhotoSaveFileInvalid
	// | 400 | PHOTO_SAVE_FILE_INVALID | Internal issues, try again later. |
	ErrPhotoSaveFileInvalid = status.Error(400, "PHOTO_SAVE_FILE_INVALID")

	// ErrPollAnswersInvalid
	// | 400 | POLL_ANSWERS_INVALID | Invalid poll answers were provided. |
	ErrPollAnswersInvalid = status.Error(400, "POLL_ANSWERS_INVALID")

	// ErrPublicKeyRequired
	// | 400 | PUBLIC_KEY_REQUIRED | A public key is required. |
	ErrPublicKeyRequired = status.Error(400, "PUBLIC_KEY_REQUIRED")

	// ErrReplyMarkupTooLong
	// | 400 | REPLY_MARKUP_TOO_LONG | The specified reply_markup is too long. |
	ErrReplyMarkupTooLong = status.Error(400, "REPLY_MARKUP_TOO_LONG")

	// ErrEmoticonEmpty
	// | 400 | EMOTICON_EMPTY | The emoji is empty. |
	ErrEmoticonEmpty = status.Error(400, "EMOTICON_EMPTY")

	// ErrEmoticonInvalid
	// | 400 | EMOTICON_INVALID | The specified emoji is invalid. |
	ErrEmoticonInvalid = status.Error(400, "EMOTICON_INVALID")

	// ErrInvitesTooMuch
	// | 400 | INVITES_TOO_MUCH | The maximum number of per-folder invites specified by the `chatlist_invites_limit_default`/`chatlist_invites_limit_premium` [client configuration parameters &raquo;](/api/config#chatlist-invites-limit-default) was reached. |
	ErrInvitesTooMuch = status.Error(400, "INVITES_TOO_MUCH")

	// ErrUserIdInvalid
	// | 400 | USER_ID_INVALID | The provided user ID is invalid. |
	ErrUserIdInvalid = status.Error(400, "USER_ID_INVALID")

	// ErrCodeEmpty
	// | 400 | CODE_EMPTY | The provided code is empty. |
	ErrCodeEmpty = status.Error(400, "CODE_EMPTY")

	// ErrMd5ChecksumInvalid
	// | 400 | MD5_CHECKSUM_INVALID | The MD5 checksums do not match. |
	ErrMd5ChecksumInvalid = status.Error(400, "MD5_CHECKSUM_INVALID")

	// ErrSendMessageMediaInvalid
	// | 400 | SEND_MESSAGE_MEDIA_INVALID | Invalid media provided. |
	ErrSendMessageMediaInvalid = status.Error(400, "SEND_MESSAGE_MEDIA_INVALID")

	// ErrApiIdInvalid
	// | 400 | API_ID_INVALID | API ID invalid. |
	ErrApiIdInvalid = status.Error(400, "API_ID_INVALID")

	// ErrChatLinkExists
	// | 400 | CHAT_LINK_EXISTS | The chat is public, you can't hide the history to new users. |
	ErrChatLinkExists = status.Error(400, "CHAT_LINK_EXISTS")

	// ErrGroupcallInvalid
	// | 400 | GROUPCALL_INVALID | The specified group call is invalid. |
	ErrGroupcallInvalid = status.Error(400, "GROUPCALL_INVALID")

	// ErrNewSettingsInvalid
	// | 400 | NEW_SETTINGS_INVALID | The new password settings are invalid. |
	ErrNewSettingsInvalid = status.Error(400, "NEW_SETTINGS_INVALID")

	// ErrBotGroupsBlocked
	// | 400 | BOT_GROUPS_BLOCKED | This bot can't be added to groups. |
	ErrBotGroupsBlocked = status.Error(400, "BOT_GROUPS_BLOCKED")

	// ErrBotScoreNotModified
	// | 400 | BOT_SCORE_NOT_MODIFIED | The score wasn't modified. |
	ErrBotScoreNotModified = status.Error(400, "BOT_SCORE_NOT_MODIFIED")

	// ErrMessageEditTimeExpired
	// | 400 | MESSAGE_EDIT_TIME_EXPIRED | You can't edit this message anymore, too much time has passed since its creation. |
	ErrMessageEditTimeExpired = status.Error(400, "MESSAGE_EDIT_TIME_EXPIRED")

	// ErrGifIdInvalid
	// | 400 | GIF_ID_INVALID | The provided GIF ID is invalid. |
	ErrGifIdInvalid = status.Error(400, "GIF_ID_INVALID")

	// ErrRangesInvalid
	// | 400 | RANGES_INVALID | Invalid range provided. |
	ErrRangesInvalid = status.Error(400, "RANGES_INVALID")

	// ErrChatIdInvalid
	// | 400 | CHAT_ID_INVALID | The provided chat id is invalid. |
	ErrChatIdInvalid = status.Error(400, "CHAT_ID_INVALID")

	// ErrFileEmtpy
	// | 400 | FILE_EMTPY | An empty file was provided. |
	ErrFileEmtpy = status.Error(400, "FILE_EMTPY")

	// ErrFileIdInvalid
	// | 400 | FILE_ID_INVALID | The provided file id is invalid. |
	ErrFileIdInvalid = status.Error(400, "FILE_ID_INVALID")

	// ErrResultsTooMuch
	// | 400 | RESULTS_TOO_MUCH | Too many results were provided. |
	ErrResultsTooMuch = status.Error(400, "RESULTS_TOO_MUCH")

	// ErrPeersListEmpty
	// | 400 | PEERS_LIST_EMPTY | The specified list of peers is empty. |
	ErrPeersListEmpty = status.Error(400, "PEERS_LIST_EMPTY")

	// ErrPhotoCropFileMissing
	// | 400 | PHOTO_CROP_FILE_MISSING | Photo crop file missing. |
	ErrPhotoCropFileMissing = status.Error(400, "PHOTO_CROP_FILE_MISSING")

	// ErrPhotoFileMissing
	// | 400 | PHOTO_FILE_MISSING | Profile photo file missing. |
	ErrPhotoFileMissing = status.Error(400, "PHOTO_FILE_MISSING")

	// ErrQuizCorrectAnswersEmpty
	// | 400 | QUIZ_CORRECT_ANSWERS_EMPTY | No correct quiz answer was specified. |
	ErrQuizCorrectAnswersEmpty = status.Error(400, "QUIZ_CORRECT_ANSWERS_EMPTY")

	// ErrStickerIdInvalid
	// | 400 | STICKER_ID_INVALID | The provided sticker ID is invalid. |
	ErrStickerIdInvalid = status.Error(400, "STICKER_ID_INVALID")

	// ErrTopicNotModified
	// | 400 | TOPIC_NOT_MODIFIED | The updated topic info is equal to the current topic info, nothing was changed. |
	ErrTopicNotModified = status.Error(400, "TOPIC_NOT_MODIFIED")

	// ErrMegagroupPrehistoryHidden
	// | 400 | MEGAGROUP_PREHISTORY_HIDDEN | Group with hidden history for new members can't be set as discussion groups. |
	ErrMegagroupPrehistoryHidden = status.Error(400, "MEGAGROUP_PREHISTORY_HIDDEN")

	// ErrUrlInvalid
	// | 400 | URL_INVALID | Invalid URL provided. |
	ErrUrlInvalid = status.Error(400, "URL_INVALID")

	// ErrBotOnesideNotAvail
	// | 400 | BOT_ONESIDE_NOT_AVAIL | Bots can't pin messages in PM just for themselves. |
	ErrBotOnesideNotAvail = status.Error(400, "BOT_ONESIDE_NOT_AVAIL")

	// ErrGroupcallSsrcDuplicateMuch
	// | 400 | GROUPCALL_SSRC_DUPLICATE_MUCH | The app needs to retry joining the group call with a new SSRC value. |
	ErrGroupcallSsrcDuplicateMuch = status.Error(400, "GROUPCALL_SSRC_DUPLICATE_MUCH")

	// ErrGroupcallJoinMissing
	// | 400 | GROUPCALL_JOIN_MISSING | You haven't joined this group call. |
	ErrGroupcallJoinMissing = status.Error(400, "GROUPCALL_JOIN_MISSING")

	// ErrDataTooLong
	// | 400 | DATA_TOO_LONG | Data too long. |
	ErrDataTooLong = status.Error(400, "DATA_TOO_LONG")

	// ErrEmailInvalid
	// | 400 | EMAIL_INVALID | The specified email is invalid. |
	ErrEmailInvalid = status.Error(400, "EMAIL_INVALID")

	// ErrSendMessageTypeInvalid
	// | 400 | SEND_MESSAGE_TYPE_INVALID | The message type is invalid. |
	ErrSendMessageTypeInvalid = status.Error(400, "SEND_MESSAGE_TYPE_INVALID")

	// ErrChannelsAdminPublicTooMuch
	// | 400 | CHANNELS_ADMIN_PUBLIC_TOO_MUCH | You're admin of too many public channels, make some channels private to change the username of this channel. |
	ErrChannelsAdminPublicTooMuch = status.Error(400, "CHANNELS_ADMIN_PUBLIC_TOO_MUCH")

	// ErrChatInvalid
	// | 400 | CHAT_INVALID | Invalid chat. |
	ErrChatInvalid = status.Error(400, "CHAT_INVALID")

	// ErrMinDateInvalid
	// | 400 | MIN_DATE_INVALID | The specified minimum date is invalid. |
	ErrMinDateInvalid = status.Error(400, "MIN_DATE_INVALID")

	// ErrQueryIdInvalid
	// | 400 | QUERY_ID_INVALID | The query ID is invalid. |
	ErrQueryIdInvalid = status.Error(400, "QUERY_ID_INVALID")

	// ErrStickerInvalid
	// | 400 | STICKER_INVALID | The provided sticker is invalid. |
	ErrStickerInvalid = status.Error(400, "STICKER_INVALID")

	// ErrAuthTokenInvalidx
	// | 400 | AUTH_TOKEN_INVALIDX | The specified auth token is invalid. |
	ErrAuthTokenInvalidx = status.Error(400, "AUTH_TOKEN_INVALIDX")

	// ErrSrpIdInvalid
	// | 400 | SRP_ID_INVALID | Invalid SRP ID provided. |
	ErrSrpIdInvalid = status.Error(400, "SRP_ID_INVALID")

	// ErrEncryptionIdInvalid
	// | 400 | ENCRYPTION_ID_INVALID | The provided secret chat ID is invalid. |
	ErrEncryptionIdInvalid = status.Error(400, "ENCRYPTION_ID_INVALID")

	// ErrOptionInvalid
	// | 400 | OPTION_INVALID | Invalid option selected. |
	ErrOptionInvalid = status.Error(400, "OPTION_INVALID")

	// ErrFilePartLengthInvalid
	// | 400 | FILE_PART_LENGTH_INVALID | The length of a file part is invalid. |
	ErrFilePartLengthInvalid = status.Error(400, "FILE_PART_LENGTH_INVALID")

	// ErrFileReferenceInvalid
	// | 400 | FILE_REFERENCE_INVALID | The specified [file reference](https://core.telegram.org/api/file_reference) is invalid. |
	ErrFileReferenceInvalid = status.Error(400, "FILE_REFERENCE_INVALID")

	// ErrMegagroupRequired
	// | 400 | MEGAGROUP_REQUIRED | You can only use this method on a supergroup. |
	ErrMegagroupRequired = status.Error(400, "MEGAGROUP_REQUIRED")

	// ErrPeerHistoryEmpty
	// | 400 | PEER_HISTORY_EMPTY | You can't pin an empty chat with a user. |
	ErrPeerHistoryEmpty = status.Error(400, "PEER_HISTORY_EMPTY")

	// ErrBankCardNumberInvalid
	// | 400 | BANK_CARD_NUMBER_INVALID | The specified card number is invalid. |
	ErrBankCardNumberInvalid = status.Error(400, "BANK_CARD_NUMBER_INVALID")

	// ErrCreateCallFailed
	// | 400 | CREATE_CALL_FAILED | An error occurred while creating the call. |
	ErrCreateCallFailed = status.Error(400, "CREATE_CALL_FAILED")

	// ErrGroupcallNotModified
	// | 400 | GROUPCALL_NOT_MODIFIED | Group call settings weren't modified. |
	ErrGroupcallNotModified = status.Error(400, "GROUPCALL_NOT_MODIFIED")

	// ErrBotDomainInvalid
	// | 400 | BOT_DOMAIN_INVALID | Bot domain invalid. |
	ErrBotDomainInvalid = status.Error(400, "BOT_DOMAIN_INVALID")

	// ErrChannelTooBig
	// | 400 | CHANNEL_TOO_BIG | This channel has too many participants (>1000) to be deleted. |
	ErrChannelTooBig = status.Error(400, "CHANNEL_TOO_BIG")

	// ErrMsgWaitFailed
	// | 400 | MSG_WAIT_FAILED | A waiting call returned an error. |
	ErrMsgWaitFailed = status.Error(400, "MSG_WAIT_FAILED")

	// ErrStickersEmpty
	// | 400 | STICKERS_EMPTY | No sticker provided. |
	ErrStickersEmpty = status.Error(400, "STICKERS_EMPTY")

	// ErrMegagroupIdInvalid
	// | 400 | MEGAGROUP_ID_INVALID | Invalid supergroup ID. |
	ErrMegagroupIdInvalid = status.Error(400, "MEGAGROUP_ID_INVALID")

	// ErrScheduleStatusPrivate
	// | 400 | SCHEDULE_STATUS_PRIVATE | Can't schedule until user is online, if the user's last seen timestamp is hidden by their privacy settings. |
	ErrScheduleStatusPrivate = status.Error(400, "SCHEDULE_STATUS_PRIVATE")

	// ErrButtonDataInvalid
	// | 400 | BUTTON_DATA_INVALID | The data of one or more of the buttons you provided is invalid. |
	ErrButtonDataInvalid = status.Error(400, "BUTTON_DATA_INVALID")

	// ErrEncryptedMessageInvalid
	// | 400 | ENCRYPTED_MESSAGE_INVALID | Encrypted message invalid. |
	ErrEncryptedMessageInvalid = status.Error(400, "ENCRYPTED_MESSAGE_INVALID")

	// ErrInlineResultExpired
	// | 400 | INLINE_RESULT_EXPIRED | The inline query expired. |
	ErrInlineResultExpired = status.Error(400, "INLINE_RESULT_EXPIRED")

	// ErrStickersTooMuch
	// | 400 | STICKERS_TOO_MUCH | There are too many stickers in this stickerpack, you can't add any more. |
	ErrStickersTooMuch = status.Error(400, "STICKERS_TOO_MUCH")

	// ErrTopicIdInvalid
	// | 400 | TOPIC_ID_INVALID | The specified topic ID is invalid. |
	ErrTopicIdInvalid = status.Error(400, "TOPIC_ID_INVALID")

	// ErrPhotoContentUrlEmpty
	// | 400 | PHOTO_CONTENT_URL_EMPTY | Photo URL invalid. |
	ErrPhotoContentUrlEmpty = status.Error(400, "PHOTO_CONTENT_URL_EMPTY")

	// ErrEmailUnconfirmed
	// | 400 | EMAIL_UNCONFIRMED | Email unconfirmed. |
	ErrEmailUnconfirmed = status.Error(400, "EMAIL_UNCONFIRMED")

	// ErrEmailVerifyExpired
	// | 400 | EMAIL_VERIFY_EXPIRED | The verification email has expired. |
	ErrEmailVerifyExpired = status.Error(400, "EMAIL_VERIFY_EXPIRED")

	// ErrImportFileInvalid
	// | 400 | IMPORT_FILE_INVALID | The specified chat export file is invalid. |
	ErrImportFileInvalid = status.Error(400, "IMPORT_FILE_INVALID")

	// ErrBroadcastIdInvalid
	// | 400 | BROADCAST_ID_INVALID | Broadcast ID invalid. |
	ErrBroadcastIdInvalid = status.Error(400, "BROADCAST_ID_INVALID")

	// ErrToLangInvalid
	// | 400 | TO_LANG_INVALID | The specified destination language is invalid. |
	ErrToLangInvalid = status.Error(400, "TO_LANG_INVALID")

	// ErrUserIsBot
	// | 400 | USER_IS_BOT | Bots can't send messages to other bots. |
	ErrUserIsBot = status.Error(400, "USER_IS_BOT")

	//// ErrFreshChangeAdminsForbidden
	//// | 400 | FRESH_CHANGE_ADMINS_FORBIDDEN | You were just elected admin, you can't add or modify other admins yet. |
	//ErrFreshChangeAdminsForbidden = status.Error(400, "FRESH_CHANGE_ADMINS_FORBIDDEN")

	//// ErrGroupcallForbidden
	//// | 400 | GROUPCALL_FORBIDDEN | The group call has already ended. |
	//ErrGroupcallForbidden = status.Error(400, "GROUPCALL_FORBIDDEN")

	// ErrPasswordMissing
	// | 400 | PASSWORD_MISSING | You must enable 2FA in order to transfer ownership of a channel. |
	ErrPasswordMissing = status.Error(400, "PASSWORD_MISSING")

	//// ErrPhoneNumberInvalid
	//// | 400 | PHONE_NUMBER_INVALID | The phone number is invalid. |
	//ErrPhoneNumberInvalid = status.Error(400, "PHONE_NUMBER_INVALID")

	// ErrEntitiesTooLong
	// | 400 | ENTITIES_TOO_LONG | You provided too many styled message entities. |
	ErrEntitiesTooLong = status.Error(400, "ENTITIES_TOO_LONG")

	//// ErrTopicClosed
	//// | 400 | TOPIC_CLOSED | This topic was closed, you can't send messages to it anymore. |
	//ErrTopicClosed = status.Error(400, "TOPIC_CLOSED")

	//// ErrPasswordTooFresh_%d
	//// | 400 | PASSWORD_TOO_FRESH_%d | The password was modified less than 24 hours ago, try again in %d seconds. |
	//ErrPasswordTooFresh_%d = status.Error(400, "PASSWORD_TOO_FRESH_%d")

	// ErrReactionInvalid
	// | 400 | REACTION_INVALID | The specified reaction is invalid. |
	ErrReactionInvalid = status.Error(400, "REACTION_INVALID")

	// ErrInviteHashInvalid
	// | 400 | INVITE_HASH_INVALID | The invite hash is invalid. |
	ErrInviteHashInvalid = status.Error(400, "INVITE_HASH_INVALID")

	// ErrPhoneCodeEmpty
	// | 400 | PHONE_CODE_EMPTY | phone_code is missing. |
	ErrPhoneCodeEmpty = status.Error(400, "PHONE_CODE_EMPTY")

	// ErrRandomLengthInvalid
	// | 400 | RANDOM_LENGTH_INVALID | Random length invalid. |
	ErrRandomLengthInvalid = status.Error(400, "RANDOM_LENGTH_INVALID")

	// ErrBotCommandInvalid
	// | 400 | BOT_COMMAND_INVALID | The specified command is invalid. |
	ErrBotCommandInvalid = status.Error(400, "BOT_COMMAND_INVALID")

	// ErrDateEmpty
	// | 400 | DATE_EMPTY | Date empty. |
	ErrDateEmpty = status.Error(400, "DATE_EMPTY")
)

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
	// ErrChatSendPlainForbidden
	// | 403 | CHAT_SEND_PLAIN_FORBIDDEN | You can't send non-media (text) messages in this chat. |
	ErrChatSendPlainForbidden = status.Error(403, "CHAT_SEND_PLAIN_FORBIDDEN")

	// ErrChatSendVoicesForbidden
	// | 403 | CHAT_SEND_VOICES_FORBIDDEN | You can't send voice recordings in this chat. |
	ErrChatSendVoicesForbidden = status.Error(403, "CHAT_SEND_VOICES_FORBIDDEN")

	// ErrChannelPublicGroupNa
	// | 403 | CHANNEL_PUBLIC_GROUP_NA | channel/supergroup not available. |
	ErrChannelPublicGroupNa = status.Error(403, "CHANNEL_PUBLIC_GROUP_NA")

	// ErrUserPrivacyRestricted
	// | 403 | USER_PRIVACY_RESTRICTED | The user's privacy settings do not allow you to do this. |
	ErrUserPrivacyRestricted = status.Error(403, "USER_PRIVACY_RESTRICTED")

	// ErrChatSendInlineForbidden
	// | 403 | CHAT_SEND_INLINE_FORBIDDEN | You can't send inline messages in this group. |
	ErrChatSendInlineForbidden = status.Error(403, "CHAT_SEND_INLINE_FORBIDDEN")

	// ErrChatSendMediaForbidden
	// | 403 | CHAT_SEND_MEDIA_FORBIDDEN | You can't send media in this chat. |
	ErrChatSendMediaForbidden = status.Error(403, "CHAT_SEND_MEDIA_FORBIDDEN")

	// ErrChatSendPollForbidden
	// | 403 | CHAT_SEND_POLL_FORBIDDEN | You can't send polls in this chat. |
	ErrChatSendPollForbidden = status.Error(403, "CHAT_SEND_POLL_FORBIDDEN")

	// ErrGroupcallAlreadyStarted
	// | 403 | GROUPCALL_ALREADY_STARTED | The groupcall has already started, you can join directly using [phone.joinGroupCall](https://core.telegram.org/method/phone.joinGroupCall). |
	ErrGroupcallAlreadyStarted = status.Error(403, "GROUPCALL_ALREADY_STARTED")

	// ErrTakeoutRequired
	// | 403 | TAKEOUT_REQUIRED | A takeout session has to be initialized, first. |
	ErrTakeoutRequired = status.Error(403, "TAKEOUT_REQUIRED")

	// ErrUserBotInvalid
	// | 403 | USER_BOT_INVALID | User accounts must provide the `bot` method parameter when calling this method. If there is no such method parameter, this method can only be invoked by bot accounts. |
	ErrUserBotInvalid = status.Error(403, "USER_BOT_INVALID")

	// ErrUserDeleted
	// | 403 | USER_DELETED | You can't send this secret message because the other participant deleted their account. |
	ErrUserDeleted = status.Error(403, "USER_DELETED")

	//// ErrUserIsBlocked
	//// | 403 | USER_IS_BLOCKED | You were blocked by this user. |
	//ErrUserIsBlocked = status.Error(403, "USER_IS_BLOCKED")

	// ErrChatSendDocsForbidden
	// | 403 | CHAT_SEND_DOCS_FORBIDDEN | You can't send documents in this chat. |
	ErrChatSendDocsForbidden = status.Error(403, "CHAT_SEND_DOCS_FORBIDDEN")

	// ErrInlineBotRequired
	// | 403 | INLINE_BOT_REQUIRED | Only the inline bot can edit message. |
	ErrInlineBotRequired = status.Error(403, "INLINE_BOT_REQUIRED")

	// ErrMessageAuthorRequired
	// | 403 | MESSAGE_AUTHOR_REQUIRED | Message author required. |
	ErrMessageAuthorRequired = status.Error(403, "MESSAGE_AUTHOR_REQUIRED")

	// ErrMessageDeleteForbidden
	// | 403 | MESSAGE_DELETE_FORBIDDEN | You can't delete one of the messages you tried to delete, most likely because it is a service message. |
	ErrMessageDeleteForbidden = status.Error(403, "MESSAGE_DELETE_FORBIDDEN")

	//// ErrUserRestricted
	//// | 403 | USER_RESTRICTED | You're spamreported, you can't create channels or chats. |
	//ErrUserRestricted = status.Error(403, "USER_RESTRICTED")

	// ErrChatSendGifsForbidden
	// | 403 | CHAT_SEND_GIFS_FORBIDDEN | You can't send gifs in this chat. |
	ErrChatSendGifsForbidden = status.Error(403, "CHAT_SEND_GIFS_FORBIDDEN")

	// ErrUserChannelsTooMuch
	// | 403 | USER_CHANNELS_TOO_MUCH | One of the users you tried to add is already in too many channels/supergroups. |
	ErrUserChannelsTooMuch = status.Error(403, "USER_CHANNELS_TOO_MUCH")

	// ErrChatSendGameForbidden
	// | 403 | CHAT_SEND_GAME_FORBIDDEN | You can't send a game to this chat. |
	ErrChatSendGameForbidden = status.Error(403, "CHAT_SEND_GAME_FORBIDDEN")

	// ErrChatSendPhotosForbidden
	// | 403 | CHAT_SEND_PHOTOS_FORBIDDEN | You can't send photos in this chat. |
	ErrChatSendPhotosForbidden = status.Error(403, "CHAT_SEND_PHOTOS_FORBIDDEN")

	// ErrChatWriteForbidden
	// | 403 | CHAT_WRITE_FORBIDDEN | You can't write in this chat. |
	ErrChatWriteForbidden = status.Error(403, "CHAT_WRITE_FORBIDDEN")

	// ErrGroupcallForbidden
	// | 403 | GROUPCALL_FORBIDDEN | The group call has already ended. |
	ErrGroupcallForbidden = status.Error(403, "GROUPCALL_FORBIDDEN")

	// ErrPremiumAccountRequired
	// | 403 | PREMIUM_ACCOUNT_REQUIRED | A premium account is required to execute this action. |
	ErrPremiumAccountRequired = status.Error(403, "PREMIUM_ACCOUNT_REQUIRED")

	// ErrUserInvalid
	// | 403 | USER_INVALID | Invalid user provided. |
	ErrUserInvalid = status.Error(403, "USER_INVALID")

	// ErrUserNotMutualContact
	// | 403 | USER_NOT_MUTUAL_CONTACT | The provided user is not a mutual contact. |
	ErrUserNotMutualContact = status.Error(403, "USER_NOT_MUTUAL_CONTACT")

	// ErrChatAdminInviteRequired
	// | 403 | CHAT_ADMIN_INVITE_REQUIRED | You do not have the rights to do this. |
	ErrChatAdminInviteRequired = status.Error(403, "CHAT_ADMIN_INVITE_REQUIRED")

	// ErrChatSendAudiosForbidden
	// | 403 | CHAT_SEND_AUDIOS_FORBIDDEN | You can't send audio messages in this chat. |
	ErrChatSendAudiosForbidden = status.Error(403, "CHAT_SEND_AUDIOS_FORBIDDEN")

	// ErrChatSendStickersForbidden
	// | 403 | CHAT_SEND_STICKERS_FORBIDDEN | You can't send stickers in this chat. |
	ErrChatSendStickersForbidden = status.Error(403, "CHAT_SEND_STICKERS_FORBIDDEN")

	// ErrChatSendVideosForbidden
	// | 403 | CHAT_SEND_VIDEOS_FORBIDDEN | You can't send videos in this chat. |
	ErrChatSendVideosForbidden = status.Error(403, "CHAT_SEND_VIDEOS_FORBIDDEN")

	//// ErrParticipantJoinMissing
	//// | 403 | PARTICIPANT_JOIN_MISSING | Trying to enable a presentation, when the user hasn't joined the Video Chat with [phone.joinGroupCall](https://core.telegram.org/method/phone.joinGroupCall). |
	//ErrParticipantJoinMissing = status.Error(403, "PARTICIPANT_JOIN_MISSING")

	// ErrPollVoteRequired
	// | 403 | POLL_VOTE_REQUIRED | Cast a vote in the poll before calling this method. |
	ErrPollVoteRequired = status.Error(403, "POLL_VOTE_REQUIRED")

	// ErrSensitiveChangeForbidden
	// | 403 | SENSITIVE_CHANGE_FORBIDDEN | You can't change your sensitive content settings. |
	ErrSensitiveChangeForbidden = status.Error(403, "SENSITIVE_CHANGE_FORBIDDEN")

	//// ErrChatAdminRequired
	//// | 403 | CHAT_ADMIN_REQUIRED | You must be an admin in this chat to do this. |
	//ErrChatAdminRequired = status.Error(403, "CHAT_ADMIN_REQUIRED")

	// ErrChatGuestSendForbidden
	// | 403 | CHAT_GUEST_SEND_FORBIDDEN | You join the discussion group before commenting, see [here &raquo;](/api/discussion#requiring-users-to-join-the-group) for more info. |
	ErrChatGuestSendForbidden = status.Error(403, "CHAT_GUEST_SEND_FORBIDDEN")

	// ErrEditBotInviteForbidden
	// | 403 | EDIT_BOT_INVITE_FORBIDDEN | Normal users can't edit invites that were created by bots. |
	ErrEditBotInviteForbidden = status.Error(403, "EDIT_BOT_INVITE_FORBIDDEN")

	// ErrPublicChannelMissing
	// | 403 | PUBLIC_CHANNEL_MISSING | You can only export group call invite links for public chats or channels. |
	ErrPublicChannelMissing = status.Error(403, "PUBLIC_CHANNEL_MISSING")

	// ErrRightForbidden
	// | 403 | RIGHT_FORBIDDEN | Your admin rights do not allow you to do this. |
	ErrRightForbidden = status.Error(403, "RIGHT_FORBIDDEN")

	// ErrBroadcastForbidden
	// | 403 | BROADCAST_FORBIDDEN | Participants of polls in channels should stay anonymous. |
	ErrBroadcastForbidden = status.Error(403, "BROADCAST_FORBIDDEN")
)

// 406 NOT_ACCEPTABLE

// NewErrPreviousChatImportActiveWaitX
// ErrPreviousChatImportActiveWait_%dmin
// | 406 | PREVIOUS_CHAT_IMPORT_ACTIVE_WAIT_%dMIN | Import for this chat is already in progress, wait %d minutes before starting a new one. |
func NewErrPreviousChatImportActiveWaitX(minute int32) error {
	return status.Errorf(ErrBadRequest, "PREVIOUS_CHAT_IMPORT_ACTIVE_WAIT_%ddMIN", minute)
}

var (
	// ErrChannelPrivate
	// | 406 | CHANNEL_PRIVATE | You haven't joined this channel/supergroup. |
	ErrChannelPrivate = status.Error(406, "CHANNEL_PRIVATE")

	// ErrChannelTooLarge
	// | 406 | CHANNEL_TOO_LARGE | Channel is too large to be deleted; this error is issued when trying to delete channels with more than 1000 members (subject to change). |
	ErrChannelTooLarge = status.Error(406, "CHANNEL_TOO_LARGE")

	// ErrFreshChangePhoneForbidden
	// | 406 | FRESH_CHANGE_PHONE_FORBIDDEN | You can't change phone number right after logging in, please wait at least 24 hours. |
	ErrFreshChangePhoneForbidden = status.Error(406, "FRESH_CHANGE_PHONE_FORBIDDEN")

	// ErrPaymentUnsupported
	// | 406 | PAYMENT_UNSUPPORTED | A detailed description of the error will be received separately as described [here &raquo;](https://core.telegram.org/api/errors#406-not-acceptable). |
	ErrPaymentUnsupported = status.Error(406, "PAYMENT_UNSUPPORTED")

	// ErrPhonePasswordFlood
	// | 406 | PHONE_PASSWORD_FLOOD | You have tried logging in too many times. |
	ErrPhonePasswordFlood = status.Error(406, "PHONE_PASSWORD_FLOOD")

	// ErrFreshChangeAdminsForbidden
	// | 406 | FRESH_CHANGE_ADMINS_FORBIDDEN | You were just elected admin, you can't add or modify other admins yet. |
	ErrFreshChangeAdminsForbidden = status.Error(406, "FRESH_CHANGE_ADMINS_FORBIDDEN")

	// ErrFreshResetAuthorisationForbidden
	// | 406 | FRESH_RESET_AUTHORISATION_FORBIDDEN | You can't logout other sessions if less than 24 hours have passed since you logged on the current session. |
	ErrFreshResetAuthorisationForbidden = status.Error(406, "FRESH_RESET_AUTHORISATION_FORBIDDEN")

	// ErrInviteHashExpired
	// | 406 | INVITE_HASH_EXPIRED | The invite link has expired. |
	ErrInviteHashExpired = status.Error(406, "INVITE_HASH_EXPIRED")

	// ErrStickersetOwnerAnonymous
	// | 406 | STICKERSET_OWNER_ANONYMOUS | Provided stickerset can't be installed as group stickerset to prevent admin deanonymization. |
	ErrStickersetOwnerAnonymous = status.Error(406, "STICKERSET_OWNER_ANONYMOUS")

	// ErrTopicDeleted
	// | 406 | TOPIC_DELETED | The specified topic was deleted. |
	ErrTopicDeleted = status.Error(406, "TOPIC_DELETED")

	// ErrUserpicUploadRequired
	// | 406 | USERPIC_UPLOAD_REQUIRED | You must have a profile picture to publish your geolocation. |
	ErrUserpicUploadRequired = status.Error(406, "USERPIC_UPLOAD_REQUIRED")

	//// ErrChatForwardsRestricted
	//// | 406 | CHAT_FORWARDS_RESTRICTED | You can't forward messages from a protected chat. |
	//ErrChatForwardsRestricted = status.Error(406, "CHAT_FORWARDS_RESTRICTED")

	// ErrFilerefUpgradeNeeded
	// | 406 | FILEREF_UPGRADE_NEEDED | The client has to be updated in order to support [file references](https://core.telegram.org/api/file_reference). |
	ErrFilerefUpgradeNeeded = status.Error(406, "FILEREF_UPGRADE_NEEDED")

	// ErrTopicClosed
	// | 406 | TOPIC_CLOSED | This topic was closed, you can't send messages to it anymore. |
	ErrTopicClosed = status.Error(406, "TOPIC_CLOSED")

	// ErrUserpicPrivacyRequired
	// | 406 | USERPIC_PRIVACY_REQUIRED | You need to disable privacy settings for your profile picture in order to make your geolocation public. |
	ErrUserpicPrivacyRequired = status.Error(406, "USERPIC_PRIVACY_REQUIRED")

	// ErrPhoneNumberInvalid
	// | 406 | PHONE_NUMBER_INVALID | The phone number is invalid. |
	ErrPhoneNumberInvalid = status.Error(406, "PHONE_NUMBER_INVALID")

	// ErrSendCodeUnavailable
	// | 406 | SEND_CODE_UNAVAILABLE | Returned when all available options for this type of number were already used (e.g. flash-call, then SMS, then this error might be returned to trigger a second resend). |
	ErrSendCodeUnavailable = status.Error(406, "SEND_CODE_UNAVAILABLE")

	// ErrStickersetInvalid
	// | 406 | STICKERSET_INVALID | The provided sticker set is invalid. |
	ErrStickersetInvalid = status.Error(406, "STICKERSET_INVALID")

	// ErrUserRestricted
	// | 406 | USER_RESTRICTED | You're spamreported, you can't create channels or chats. |
	ErrUserRestricted = status.Error(406, "USER_RESTRICTED")
)

// 500 InternalServerError
var (
	// StatusInternelServerError - StatusInternelServerError
	StatusInternelServerError = status.New(ErrInternal, "INTERNAL_SERVER_ERROR")

	// ErrInternelServerError
	// | 500 | INTERNAL_SERVER_ERROR |  |
	ErrInternelServerError = status.Error(ErrInternal, "INTERNAL_SERVER_ERROR")

	// ErrCallOccupyFailed
	// | 500 | CALL_OCCUPY_FAILED | The call failed because the user is already making another call. |
	ErrCallOccupyFailed = status.Error(500, "CALL_OCCUPY_FAILED")

	// ErrCdnUploadTimeout
	// | 500 | CDN_UPLOAD_TIMEOUT | A server-side timeout occurred while reuploading the file to the CDN DC. |
	ErrCdnUploadTimeout = status.Error(500, "CDN_UPLOAD_TIMEOUT")

	// ErrChatIdGenerateFailed
	// | 500 | CHAT_ID_GENERATE_FAILED | Failure while generating the chat ID. |
	ErrChatIdGenerateFailed = status.Error(500, "CHAT_ID_GENERATE_FAILED")

	// ErrPersistentTimestampOutdated
	// | 500 | PERSISTENT_TIMESTAMP_OUTDATED | Channel internal replication issues, try again later (treat this like an RPC_CALL_FAIL). |
	ErrPersistentTimestampOutdated = status.Error(500, "PERSISTENT_TIMESTAMP_OUTDATED")

	// ErrRandomIdDuplicate
	// | 500 | RANDOM_ID_DUPLICATE | You provided a random ID that was already used. |
	ErrRandomIdDuplicate = status.Error(500, "RANDOM_ID_DUPLICATE")

	// ErrSignInFailed
	// | 500 | SIGN_IN_FAILED | Failure while signing in. |
	ErrSignInFailed = status.Error(500, "SIGN_IN_FAILED")

	// ErrAuthRestart
	// | 500 | AUTH_RESTART | Restart the authorization process. |
	ErrAuthRestart = status.Error(500, "AUTH_RESTART")
)

// -503 Timeout
var (
	// StatusTimeout - StatusTimeout
	StatusTimeout = status.New(ErrTimeOut503, "Timeout")

	// ErrTimeout
	// | -503 | Timeout | Timeout while fetching data |
	ErrTimeout = status.Error(ErrTimeOut503, "Timeout")
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

//-500
var (
//// ErrInvalid MsgResendReq Query
//// | -500 | Invalid msg_resend_req query | Invalid msg_resend_req query. |
//ErrInvalid MsgResendReq Query = status.Error(-500, "Invalid msg_resend_req query")
//
//// ErrInvalid MsgsAck Query
//// | -500 | Invalid msgs_ack query | Invalid msgs_ack query. |
//ErrInvalid MsgsAck Query = status.Error(-500, "Invalid msgs_ack query")
//
//// ErrInvalid MsgsStateReq Query
//// | -500 | Invalid msgs_state_req query | Invalid msgs_state_req query. |
//ErrInvalid MsgsStateReq Query = status.Error(-500, "Invalid msgs_state_req query")

)
