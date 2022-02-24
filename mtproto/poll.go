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
	"encoding/json"
)

/*
   public static void updatePollResults(TLRPC.TL_messageMediaPoll media, TLRPC.TL_pollResults results) {
       if ((results.flags & 2) != 0) {
           byte[] chosen = null;
           if (results.min && media.results.results != null) {
               for (int b = 0, N2 = media.results.results.size(); b < N2; b++) {
                   TLRPC.TL_pollAnswerVoters answerVoters = media.results.results.get(b);
                   if (answerVoters.chosen) {
                       chosen = answerVoters.option;
                       break;
                   }
               }
           }
           media.results.results = results.results;
           if (chosen != null) {
               for (int b = 0, N2 = media.results.results.size(); b < N2; b++) {
                   TLRPC.TL_pollAnswerVoters answerVoters = media.results.results.get(b);
                   if (Arrays.equals(answerVoters.option, chosen)) {
                       answerVoters.chosen = true;
                       break;
                   }
               }
           }
           media.results.flags |= 2;
       }
       if ((results.flags & 4) != 0) {
           media.results.total_voters = results.total_voters;
           media.results.flags |= 4;
       }
   }

   public boolean isPollClosed() {
       if (type != TYPE_POLL) {
           return false;
       }
       return ((TLRPC.TL_messageMediaPoll) messageOwner.media).poll.closed;
   }

   public boolean isVoted() {
       if (type != TYPE_POLL) {
           return false;
       }
       TLRPC.TL_messageMediaPoll mediaPoll = (TLRPC.TL_messageMediaPoll) messageOwner.media;
       if (mediaPoll.results == null || mediaPoll.results.results.isEmpty()) {
           return false;
       }
       for (int a = 0, N = mediaPoll.results.results.size(); a < N; a++) {
           TLRPC.TL_pollAnswerVoters answer = mediaPoll.results.results.get(a);
           if (answer.chosen) {
               return true;
           }
       }
       return false;
   }

   public long getPollId() {
       if (type != TYPE_POLL) {
           return 0;
       }
       return ((TLRPC.TL_messageMediaPoll) messageOwner.media).poll.id;
   }
*/

/*
   media: { inputMediaPoll
     poll: { poll
       id: 9488326629078739604 [LONG],
       flags: 1 [INT],
       closed: YES [ BY BIT 0 IN FIELD flags ],
       question: "nnn" [STRING],
       answers: [ vector<0x0>
         { pollAnswer
           text: "bb" [STRING],
           option: "0" [STRING],
         },
         { pollAnswer
           text: "cc" [STRING],
           option: "1" [STRING],
         },
       ],
     },
   },
*/

/*
pollAnswer#6ca9c2e9 text:string option:bytes = PollAnswer;

poll#d5529d06 id:long flags:# closed:flags.0?true question:string answers:Vector<PollAnswer> = Poll;

pollAnswerVoters#3b6ddad2 flags:# chosen:flags.0?true option:bytes voters:int = PollAnswerVoters;

pollResults#5755785a flags:# min:flags.0?true results:flags.1?Vector<PollAnswerVoters> total_voters:flags.2?int = PollResults;
*/

// messageMediaPoll#4bd6e798 poll:Poll results:PollResults = MessageMedia;

type MediaPoll struct {
	UserId  int32        `json:"user_id"`
	Poll    *Poll        `json:"poll"`
	Results *PollResults `json:"results"`
}

func (m *MediaPoll) DebugString() string {
	jData, _ := json.Marshal(m)
	return string(jData)
}
func (m *MediaPoll) ToMessageMedia() *MessageMedia {
	return MakeTLMessageMediaPoll(&MessageMedia{
		Poll:    m.Poll,
		Results: m.Results,
	}).To_MessageMedia()
}

func (m *MediaPoll) ToUpdateMessagePoll() *Update {
	return MakeTLUpdateMessagePoll(&Update{
		PollId:  m.Poll.Id,
		Poll:    m.Poll,
		Results: m.Results,
	}).To_Update()
}

func GetPollIdByMessage(mediaPoll *MessageMedia) (int64, error) {
	if mediaPoll == nil {
		return 0, ErrMediaEmpty
	}
	if mediaPoll.PredicateName != Predicate_messageMediaPoll {
		return 0, ErrMediaInvalid
	}

	return mediaPoll.GetPoll().GetId(), nil
}

func GetPollByMessage(mediaPoll *MessageMedia) (*Poll, error) {
	if mediaPoll == nil {
		return nil, ErrMediaEmpty
	}
	if mediaPoll.PredicateName != Predicate_messageMediaPoll {
		return nil, ErrMediaInvalid
	}

	return mediaPoll.GetPoll(), nil
}
