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
	"regexp"
	"strings"
)

func IsStickerDocument(document *Document) bool {
	for _, attribute := range document.GetAttributes() {
		if attribute.PredicateName == Predicate_documentAttributeSticker {
			return true
		}
	}
	return false
}

func IsMaskDocument(document *Document) bool {
	for _, attribute := range document.GetAttributes() {
		if attribute.PredicateName == Predicate_documentAttributeSticker && attribute.Mask {
			return true
		}
	}
	return false
}

func IsVoiceDocument(document *Document) bool {
	for _, attribute := range document.GetAttributes() {
		if attribute.PredicateName == Predicate_documentAttributeAudio {
			return attribute.Voice
		}
	}
	return false
}

func IsMusicDocument(document *Document) bool {
	for _, attribute := range document.GetAttributes() {
		if attribute.PredicateName == Predicate_documentAttributeAudio {
			return !attribute.Voice
		}

		if document.MimeType != "" {
			mime := strings.ToLower(document.MimeType)
			switch mime {
			case "audio/flac":
				return true
			case "audio/ogg":
				return true
			case "audio/opus":
				return true
			case "audio/x-opus+ogg":
				return true
			case "application/octet-stream":
				dFileName := getDocumentFileName(document)
				if strings.HasSuffix(dFileName, ".opus") {
					return true
				}
			}
		}
	}
	return false
}

func fixFileName(fileName string) (n string) {
	if fileName != "" {
		re, _ := regexp.Compile("[\u0001-\u001f<>:\"/\\\\|?*\u007f]+")
		n = strings.TrimSpace(re.ReplaceAllString(fileName, ""))
	}
	return
}

func getDocumentFileName(document *Document) (fileName string) {
	for _, attribute := range document.GetAttributes() {
		if attribute.PredicateName == Predicate_documentAttributeFilename {
			fileName = attribute.FileName
		}
	}
	return fixFileName(fileName)
}

func IsVideoDocument(document *Document) bool {
	if document == nil {
		return false
	}

	var (
		isAnimated = false
		isVideo    = false
		width      int32
		height     int32
	)
	for _, attribute := range document.GetAttributes() {
		if attribute.PredicateName == Predicate_documentAttributeVideo {
			if attribute.RoundMessage {
				return false
			}
			isVideo = true
			width = attribute.W
			height = attribute.H
		} else if attribute.PredicateName == Predicate_documentAttributeAnimated {
			isAnimated = true
		}
	}

	if isAnimated && (width > 1280 || height > 1280) {
		isAnimated = false
	}
	if !isVideo && "video/x-matroska" == document.MimeType {
		isVideo = true
	}
	return isVideo && !isAnimated
}

func IsDocumentHasThumb(document *Document) bool {
	if document == nil || len(document.Thumbs) > 0 {
		return false
	}

	for _, photoSize := range document.Thumbs {
		if photoSize != nil &&
			photoSize.PredicateName != Predicate_photoSizeEmpty {
			return true
		}

	}
	return false
}

func IsGifDocument(document *Document) bool {
	return document != nil &&
		len(document.Thumbs) > 0 &&
		(document.MimeType == "image/gif" || IsNewGifDocument(document))
}

func IsRoundVideoDocument(document *Document) bool {
	if document == nil || document.MimeType != "video/mp4" {
		return false
	}

	var (
		width  int32
		height int32
		round  bool
	)

	for _, attribute := range document.GetAttributes() {
		if attribute.PredicateName == Predicate_documentAttributeVideo {
			width = attribute.W
			height = attribute.H
			round = attribute.RoundMessage
		}
	}
	if round && width <= 1280 && height <= 1280 {
		return true
	}
	return false
}

func IsNewGifDocument(document *Document) bool {
	if document == nil || document.MimeType != "video/mp4" {
		return false
	}

	var (
		width    int32
		height   int32
		animated bool
	)

	for _, attribute := range document.GetAttributes() {
		if attribute.PredicateName == Predicate_documentAttributeAnimated {
			animated = true
		} else if attribute.PredicateName == Predicate_documentAttributeVideo {
			width = attribute.W
			height = attribute.H
		}
	}
	if animated && width <= 1280 && height <= 1280 {
		return true
	}
	return false
}

// GetDocumentAttribute
/*
documentAttributeImageSize#6c37c15c w:int h:int = DocumentAttribute;
documentAttributeAnimated#11b58939 = DocumentAttribute;
documentAttributeSticker#6319d612 flags:# mask:flags.1?true alt:string stickerset:InputStickerSet mask_coords:flags.0?MaskCoords = DocumentAttribute;
documentAttributeVideo#ef02ce6 flags:# round_message:flags.0?true supports_streaming:flags.1?true duration:int w:int h:int = DocumentAttribute;
documentAttributeAudio#9852f9c6 flags:# voice:flags.10?true duration:int title:flags.0?string performer:flags.1?string waveform:flags.2?bytes = DocumentAttribute;
documentAttributeFilename#15590068 file_name:string = DocumentAttribute;
documentAttributeHasStickers#9801d2f7 = DocumentAttribute;
*/
func GetDocumentAttribute(attributes []*DocumentAttribute, name string) (attribute *DocumentAttribute) {
	for _, a := range attributes {
		if a.PredicateName == name {
			switch name {
			case Predicate_documentAttributeImageSize:
				attribute = a
				break
			case Predicate_documentAttributeAnimated:
				attribute = a
				break
			case Predicate_documentAttributeSticker:
				attribute = a
				break
			case Predicate_documentAttributeVideo:
				attribute = a
				break
			case Predicate_documentAttributeAudio:
				attribute = a
				break
			case Predicate_documentAttributeFilename:
				attribute = a
				break
			case Predicate_documentAttributeHasStickers:
				attribute = a
				break
			}
		}
	}
	return
}

func GetVideoDuration(attributes []*DocumentAttribute) int32 {
	for _, a := range attributes {
		if a.PredicateName == Predicate_documentAttributeVideo {
			return a.Duration
		}
	}
	return 0
}

func GetAudioDuration(attributes []*DocumentAttribute) int32 {
	for _, a := range attributes {
		if a.PredicateName == Predicate_documentAttributeAudio {
			return a.Duration
		}
	}
	return 0
}

func FindDocumentById(documents []*Document, id int64) *Document {
	for _, d := range documents {
		if d.GetId() == id {
			return d
		}
	}
	return nil
}
