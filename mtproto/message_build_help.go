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

type MessageBuildEntry struct {
	Text           string
	Param          string
	EntityType     string
	EntityUrl      string
	EntityUserId   int64
	EntityLanguage string
}

type MessageBuildHelper []MessageBuildEntry

/*
TODO(@benqi):

	messageEntityUnknown#bb92ba95 offset:int length:int = MessageEntity;
	messageEntityMention#fa04579d offset:int length:int = MessageEntity;
	messageEntityHashtag#6f635b0d offset:int length:int = MessageEntity;
	messageEntityBotCommand#6cef8ac7 offset:int length:int = MessageEntity;
	messageEntityUrl#6ed02538 offset:int length:int = MessageEntity;
	messageEntityEmail#64e475c2 offset:int length:int = MessageEntity;
	messageEntityBold#bd610bc9 offset:int length:int = MessageEntity;
	messageEntityItalic#826f8b60 offset:int length:int = MessageEntity;
	messageEntityCode#28a20571 offset:int length:int = MessageEntity;
	messageEntityPre#73924be0 offset:int length:int language:string = MessageEntity;
	messageEntityTextUrl#76a6d327 offset:int length:int url:string = MessageEntity;
	messageEntityMentionName#352dca58 offset:int length:int user_id:int = MessageEntity;
	inputMessageEntityMentionName#208e68c9 offset:int length:int user_id:InputUser = MessageEntity;
	messageEntityPhone#9b69e34b offset:int length:int = MessageEntity;
	messageEntityCashtag#4c4e743f offset:int length:int = MessageEntity;
	messageEntityUnderline#9c4e7e8b offset:int length:int = MessageEntity;
	messageEntityStrike#bf0693d4 offset:int length:int = MessageEntity;
	messageEntityBlockquote#20df5d0 offset:int length:int = MessageEntity;
*/
func MakeTextAndMessageEntities(m MessageBuildHelper) (text string, entities []*MessageEntity) {
	if len(m) == 0 {
		return
	}

	var (
		offset int
		length int
	)
	for i := 0; i < len(m); i++ {
		text += m[i].Text
		offset = len(text)
		length = len(m[i].Param)
		if length > 0 {
			switch m[i].EntityType {
			case Predicate_messageEntityUnknown:
				entities = append(entities, MakeTLMessageEntityUnknown(&MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case Predicate_messageEntityMention:
				entities = append(entities, MakeTLMessageEntityMention(&MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case Predicate_messageEntityHashtag:
				entities = append(entities, MakeTLMessageEntityHashtag(&MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case Predicate_messageEntityBotCommand:
				entities = append(entities, MakeTLMessageEntityBotCommand(&MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case Predicate_messageEntityUrl:
				entities = append(entities, MakeTLMessageEntityUrl(&MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case Predicate_messageEntityEmail:
				entities = append(entities, MakeTLMessageEntityEmail(&MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case Predicate_messageEntityBold:
				entities = append(entities, MakeTLMessageEntityBold(&MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case Predicate_messageEntityItalic:
				entities = append(entities, MakeTLMessageEntityItalic(&MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case Predicate_messageEntityCode:
				entities = append(entities, MakeTLMessageEntityCode(&MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case Predicate_messageEntityPre:
				entities = append(entities, MakeTLMessageEntityPre(&MessageEntity{
					Offset:   int32(offset),
					Length:   int32(length),
					Language: m[i].EntityLanguage,
				}).To_MessageEntity())
			case Predicate_messageEntityTextUrl:
				entities = append(entities, MakeTLMessageEntityTextUrl(&MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
					Url:    m[i].EntityUrl,
				}).To_MessageEntity())
			case Predicate_messageEntityMentionName:
				entities = append(entities, MakeTLMessageEntityMentionName(&MessageEntity{
					Offset:       int32(offset),
					Length:       int32(length),
					UserId_INT64: m[i].EntityUserId,
				}).To_MessageEntity())
			case Predicate_messageEntityPhone:
				entities = append(entities, MakeTLMessageEntityPhone(&MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case Predicate_messageEntityCashtag:
				entities = append(entities, MakeTLMessageEntityCashtag(&MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case Predicate_messageEntityUnderline:
				entities = append(entities, MakeTLMessageEntityUnderline(&MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case Predicate_messageEntityStrike:
				entities = append(entities, MakeTLMessageEntityStrike(&MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			case Predicate_messageEntityBlockquote:
				entities = append(entities, MakeTLMessageEntityBlockquote(&MessageEntity{
					Offset: int32(offset),
					Length: int32(length),
				}).To_MessageEntity())
			}
		}
		text = text + m[i].Param
	}

	return
}
