/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

// ConstructorList
// RequestList

package mt

import (
	"github.com/teamgram/proto/v2/iface"
)

func init() {
	// Constructor
	iface.RegisterClazzID(1973679973, func() iface.TLObject { return &TLBindAuthKeyInner{ClazzID: 1973679973} })        // 0x75a3f765
	iface.RegisterClazzID(1715713620, func() iface.TLObject { return &TLClientDHInnerData{ClazzID: 1715713620} })       // 0x6643b654
	iface.RegisterClazzID(-161422892, func() iface.TLObject { return &TLDestroyAuthKeyOk{ClazzID: -161422892} })        // 0xf660e1d4
	iface.RegisterClazzID(178201177, func() iface.TLObject { return &TLDestroyAuthKeyNone{ClazzID: 178201177} })        // 0xa9f2259
	iface.RegisterClazzID(-368010477, func() iface.TLObject { return &TLDestroyAuthKeyFail{ClazzID: -368010477} })      // 0xea109b13
	iface.RegisterClazzID(-2083955988, func() iface.TLObject { return &TLPQInnerData{ClazzID: -2083955988} })           // 0x83c95aec
	iface.RegisterClazzID(-1443537003, func() iface.TLObject { return &TLPQInnerDataDc{ClazzID: -1443537003} })         // 0xa9f55f95
	iface.RegisterClazzID(1013613780, func() iface.TLObject { return &TLPQInnerDataTemp{ClazzID: 1013613780} })         // 0x3c6a84d4
	iface.RegisterClazzID(1459478408, func() iface.TLObject { return &TLPQInnerDataTempDc{ClazzID: 1459478408} })       // 0x56fddf88
	iface.RegisterClazzID(85337187, func() iface.TLObject { return &TLResPQ{ClazzID: 85337187} })                       // 0x5162463
	iface.RegisterClazzID(-1249309254, func() iface.TLObject { return &TLServerDHInnerData{ClazzID: -1249309254} })     // 0xb5890dba
	iface.RegisterClazzID(2043348061, func() iface.TLObject { return &TLServerDHParamsFail{ClazzID: 2043348061} })      // 0x79cb045d
	iface.RegisterClazzID(-790100132, func() iface.TLObject { return &TLServerDHParamsOk{ClazzID: -790100132} })        // 0xd0e8075c
	iface.RegisterClazzID(1003222836, func() iface.TLObject { return &TLDhGenOk{ClazzID: 1003222836} })                 // 0x3bcbf734
	iface.RegisterClazzID(1188831161, func() iface.TLObject { return &TLDhGenRetry{ClazzID: 1188831161} })              // 0x46dc1fb9
	iface.RegisterClazzID(-1499615742, func() iface.TLObject { return &TLDhGenFail{ClazzID: -1499615742} })             // 0xa69dae02
	iface.RegisterClazzID(1182381663, func() iface.TLObject { return &TLAccessPointRule{ClazzID: 1182381663} })         // 0x4679b65f
	iface.RegisterClazzID(-1477445615, func() iface.TLObject { return &TLBadMsgNotification{ClazzID: -1477445615} })    // 0xa7eff811
	iface.RegisterClazzID(-307542917, func() iface.TLObject { return &TLBadServerSalt{ClazzID: -307542917} })           // 0xedab447b
	iface.RegisterClazzID(-501201412, func() iface.TLObject { return &TLDestroySessionOk{ClazzID: -501201412} })        // 0xe22045fc
	iface.RegisterClazzID(1658015945, func() iface.TLObject { return &TLDestroySessionNone{ClazzID: 1658015945} })      // 0x62d350c9
	iface.RegisterClazzID(155834844, func() iface.TLObject { return &TLFutureSalt{ClazzID: 155834844} })                // 0x949d9dc
	iface.RegisterClazzID(-1370486635, func() iface.TLObject { return &TLFutureSalts{ClazzID: -1370486635} })           // 0xae500895
	iface.RegisterClazzID(1515793004, func() iface.TLObject { return &TLHelpConfigSimple{ClazzID: 1515793004} })        // 0x5a592a6c
	iface.RegisterClazzID(-1835453025, func() iface.TLObject { return &TLHttpWait{ClazzID: -1835453025} })              // 0x9299359f
	iface.RegisterClazzID(-734810765, func() iface.TLObject { return &TLIpPort{ClazzID: -734810765} })                  // 0xd433ad73
	iface.RegisterClazzID(932718150, func() iface.TLObject { return &TLIpPortSecret{ClazzID: 932718150} })              // 0x37982646
	iface.RegisterClazzID(661470918, func() iface.TLObject { return &TLMsgDetailedInfo{ClazzID: 661470918} })           // 0x276d3ec6
	iface.RegisterClazzID(-2137147681, func() iface.TLObject { return &TLMsgNewDetailedInfo{ClazzID: -2137147681} })    // 0x809db6df
	iface.RegisterClazzID(2105940488, func() iface.TLObject { return &TLMsgResendReq{ClazzID: 2105940488} })            // 0x7d861a08
	iface.RegisterClazzID(1658238041, func() iface.TLObject { return &TLMsgsAck{ClazzID: 1658238041} })                 // 0x62d6b459
	iface.RegisterClazzID(-1933520591, func() iface.TLObject { return &TLMsgsAllInfo{ClazzID: -1933520591} })           // 0x8cc0d131
	iface.RegisterClazzID(81704317, func() iface.TLObject { return &TLMsgsStateInfo{ClazzID: 81704317} })               // 0x4deb57d
	iface.RegisterClazzID(-630588590, func() iface.TLObject { return &TLMsgsStateReq{ClazzID: -630588590} })            // 0xda69fb52
	iface.RegisterClazzID(-1631450872, func() iface.TLObject { return &TLNewSessionCreated{ClazzID: -1631450872} })     // 0x9ec20908
	iface.RegisterClazzID(880243653, func() iface.TLObject { return &TLPong{ClazzID: 880243653} })                      // 0x347773c5
	iface.RegisterClazzID(1579864942, func() iface.TLObject { return &TLRpcAnswerUnknown{ClazzID: 1579864942} })        // 0x5e2ad36e
	iface.RegisterClazzID(-847714938, func() iface.TLObject { return &TLRpcAnswerDroppedRunning{ClazzID: -847714938} }) // 0xcd78e586
	iface.RegisterClazzID(-1539647305, func() iface.TLObject { return &TLRpcAnswerDropped{ClazzID: -1539647305} })      // 0xa43ad8b7
	iface.RegisterClazzID(558156313, func() iface.TLObject { return &TLRpcError{ClazzID: 558156313} })                  // 0x2144ca19
	iface.RegisterClazzID(1108910436, func() iface.TLObject { return &TLTlsBlockString{ClazzID: 1108910436} })          // 0x4218a164
	iface.RegisterClazzID(1296942110, func() iface.TLObject { return &TLTlsBlockRandom{ClazzID: 1296942110} })          // 0x4d4dc41e
	iface.RegisterClazzID(154352379, func() iface.TLObject { return &TLTlsBlockZero{ClazzID: 154352379} })              // 0x9333afb
	iface.RegisterClazzID(283665263, func() iface.TLObject { return &TLTlsBlockDomain{ClazzID: 283665263} })            // 0x10e8636f
	iface.RegisterClazzID(-428498495, func() iface.TLObject { return &TLTlsBlockGrease{ClazzID: -428498495} })          // 0xe675a1c1
	iface.RegisterClazzID(-1632019620, func() iface.TLObject { return &TLTlsBlockPublicKey{ClazzID: -1632019620} })     // 0x9eb95b5c
	iface.RegisterClazzID(-416951217, func() iface.TLObject { return &TLTlsBlockScope{ClazzID: -416951217} })           // 0xe725d44f
	iface.RegisterClazzID(1817363588, func() iface.TLObject { return &TLTlsClientHello{ClazzID: 1817363588} })          // 0x6c52c484

	// Method
	iface.RegisterClazzID(1615239032, func() iface.TLObject { return &TLReqPq{ClazzID: 1615239032} })                    // 0x60469778
	iface.RegisterClazzID(-1099002127, func() iface.TLObject { return &TLReqPqMulti{ClazzID: -1099002127} })             // 0xbe7e8ef1
	iface.RegisterClazzID(-686627650, func() iface.TLObject { return &TLReqDHParams{ClazzID: -686627650} })              // 0xd712e4be
	iface.RegisterClazzID(-184262881, func() iface.TLObject { return &TLSetClientDHParams{ClazzID: -184262881} })        // 0xf5045f1f
	iface.RegisterClazzID(-784117408, func() iface.TLObject { return &TLDestroyAuthKey{ClazzID: -784117408} })           // 0xd1435160
	iface.RegisterClazzID(1491380032, func() iface.TLObject { return &TLRpcDropAnswer{ClazzID: 1491380032} })            // 0x58e4a740
	iface.RegisterClazzID(-1188971260, func() iface.TLObject { return &TLGetFutureSalts{ClazzID: -1188971260} })         // 0xb921bd04
	iface.RegisterClazzID(2059302892, func() iface.TLObject { return &TLPing{ClazzID: 2059302892} })                     // 0x7abe77ec
	iface.RegisterClazzID(-213746804, func() iface.TLObject { return &TLPingDelayDisconnect{ClazzID: -213746804} })      // 0xf3427b8c
	iface.RegisterClazzID(-414113498, func() iface.TLObject { return &TLDestroySession{ClazzID: -414113498} })           // 0xe7512126
	iface.RegisterClazzID(-294277375, func() iface.TLObject { return &TLTestUseError{ClazzID: -294277375} })             // 0xee75af01
	iface.RegisterClazzID(-105401795, func() iface.TLObject { return &TLTestUseConfigSimple{ClazzID: -105401795} })      // 0xf9b7b23d
	iface.RegisterClazzID(-1156741135, func() iface.TLObject { return &TLTestParseInputAppEvent{ClazzID: -1156741135} }) // 0xbb0d87f1
}
