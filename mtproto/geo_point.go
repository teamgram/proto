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

// MakeGeoPointByInput
// inputGeoPointEmpty#e4c123d6 = InputGeoPoint;
// inputGeoPoint#f3b7acc9 lat:double long:double = InputGeoPoint;
//
// geoPointEmpty#1117dd5f = GeoPoint;
// geoPoint#296f104 long:double lat:double access_hash:long = GeoPoint;
func MakeGeoPointByInput(geoPoint *InputGeoPoint) (geo *GeoPoint) {
	switch geoPoint.PredicateName {
	case Predicate_inputGeoPointEmpty:
		geo = MakeTLGeoPointEmpty(nil).To_GeoPoint()
	case Predicate_inputGeoPoint:
		// TODO(@benqi): check access_hash
		geo = MakeTLGeoPoint(&GeoPoint{
			Long: geoPoint.Long,
			Lat:  geoPoint.Lat,
		}).To_GeoPoint()
	default:
		// log.Errorf("invalid predicateName - %v", geoPoint)
	}
	return
}
