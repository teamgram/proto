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
	"github.com/gogo/protobuf/types"
)

var (
	DefaultPeerNotifySettings = MakeDefaultPeerNotifySettings()
)

// MakePeerNotifySettings
// peerNotifySettings#99622c0c flags:# show_previews:flags.0?Bool silent:flags.1?Bool mute_until:flags.2?int ios_sound:flags.3?NotificationSound android_sound:flags.4?NotificationSound other_sound:flags.5?NotificationSound stories_muted:flags.6?Bool stories_hide_sender:flags.7?Bool stories_ios_sound:flags.8?NotificationSound stories_android_sound:flags.9?NotificationSound stories_other_sound:flags.10?NotificationSound = PeerNotifySettings;
func MakePeerNotifySettings(settings *InputPeerNotifySettings) (*PeerNotifySettings, error) {
	notifySettings := MakeTLPeerNotifySettings(&PeerNotifySettings{
		ShowPreviews:        settings.GetShowPreviews(),
		Silent:              settings.GetSilent(),
		MuteUntil:           settings.GetMuteUntil(),
		IosSound:            nil,
		AndroidSound:        nil,
		OtherSound:          nil,
		StoriesMuted:        nil,
		StoriesHideSender:   nil,
		StoriesIosSound:     nil,
		StoriesAndroidSound: nil,
		StoriesOtherSound:   nil,
	}).To_PeerNotifySettings()

	return notifySettings, nil
}

// MakeDefaultPeerNotifySettings
/*
   { updateNotifySettings
     peer: { notifyChats },
     notify_settings: { peerNotifySettings
       flags: 15 [INT],
       show_previews: { boolTrue },
       silent: { boolFalse },
       mute_until: 0 [INT],
       sound: "100" [STRING],
     },
   },
*/
func MakeDefaultPeerNotifySettings() *PeerNotifySettings {
	settings := MakeTLPeerNotifySettings(&PeerNotifySettings{
		ShowPreviews:        ToBool(true),
		Silent:              ToBool(true),
		MuteUntil:           &types.Int32Value{Value: 0},
		IosSound:            nil,
		AndroidSound:        nil,
		OtherSound:          nil,
		StoriesMuted:        nil,
		StoriesHideSender:   nil,
		StoriesIosSound:     nil,
		StoriesAndroidSound: nil,
		StoriesOtherSound:   nil,
	}).To_PeerNotifySettings()

	return settings
}
