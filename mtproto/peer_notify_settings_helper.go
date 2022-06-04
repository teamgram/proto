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

func MakePeerNotifySettings(settings *InputPeerNotifySettings) (*PeerNotifySettings, error) {
	notifySettings := MakeTLPeerNotifySettings(&PeerNotifySettings{
		ShowPreviews: settings.GetShowPreviews(),
		Silent:       settings.GetSilent(),
		MuteUntil:    settings.GetMuteUntil(),
		Sound:        settings.GetSound_FLAGSTRING(),
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
		ShowPreviews: ToBool(true),
		Silent:       ToBool(true),
		MuteUntil:    &types.Int32Value{Value: 0},
		Sound:        &types.StringValue{Value: "default"},
	}).To_PeerNotifySettings()

	return settings
}
