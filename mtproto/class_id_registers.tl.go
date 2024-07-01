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

package mtproto

var clazzIdRegisters2 = map[int32]func() TLObject{
	// parsed_manually_types
	1538843921: func() TLObject { // CRC32_message2
		return &TLMessage2{}
	},
	1945237724: func() TLObject { // CRC32_msg_container
		return &TLMsgContainer{}
	},
	530561358: func() TLObject { // CRC32_msg_copy
		return &TLMsgCopy{}
	},
	812830625: func() TLObject { // CRC32_gzip_packed
		return &TLGzipPacked{}
	},
	-212046591: func() TLObject { // CRC32_rpc_result
		return &TLRpcResult{}
	},

	// Constructor
	1973679973: func() TLObject { // 0x75a3f765
		o := MakeTLBindAuthKeyInner(nil)
		o.Data2.Constructor = 1973679973
		return o
	},
	1715713620: func() TLObject { // 0x6643b654
		o := MakeTLClient_DHInnerData(nil)
		o.Data2.Constructor = 1715713620
		return o
	},
	-161422892: func() TLObject { // 0xf660e1d4
		o := MakeTLDestroyAuthKeyOk(nil)
		o.Data2.Constructor = -161422892
		return o
	},
	178201177: func() TLObject { // 0xa9f2259
		o := MakeTLDestroyAuthKeyNone(nil)
		o.Data2.Constructor = 178201177
		return o
	},
	-368010477: func() TLObject { // 0xea109b13
		o := MakeTLDestroyAuthKeyFail(nil)
		o.Data2.Constructor = -368010477
		return o
	},
	-2083955988: func() TLObject { // 0x83c95aec
		o := MakeTLPQInnerData(nil)
		o.Data2.Constructor = -2083955988
		return o
	},
	-1443537003: func() TLObject { // 0xa9f55f95
		o := MakeTLPQInnerDataDc(nil)
		o.Data2.Constructor = -1443537003
		return o
	},
	1013613780: func() TLObject { // 0x3c6a84d4
		o := MakeTLPQInnerDataTemp(nil)
		o.Data2.Constructor = 1013613780
		return o
	},
	1459478408: func() TLObject { // 0x56fddf88
		o := MakeTLPQInnerDataTempDc(nil)
		o.Data2.Constructor = 1459478408
		return o
	},
	85337187: func() TLObject { // 0x5162463
		o := MakeTLResPQ(nil)
		o.Data2.Constructor = 85337187
		return o
	},
	-1249309254: func() TLObject { // 0xb5890dba
		o := MakeTLServer_DHInnerData(nil)
		o.Data2.Constructor = -1249309254
		return o
	},
	2043348061: func() TLObject { // 0x79cb045d
		o := MakeTLServer_DHParamsFail(nil)
		o.Data2.Constructor = 2043348061
		return o
	},
	-790100132: func() TLObject { // 0xd0e8075c
		o := MakeTLServer_DHParamsOk(nil)
		o.Data2.Constructor = -790100132
		return o
	},
	1003222836: func() TLObject { // 0x3bcbf734
		o := MakeTLDhGenOk(nil)
		o.Data2.Constructor = 1003222836
		return o
	},
	1188831161: func() TLObject { // 0x46dc1fb9
		o := MakeTLDhGenRetry(nil)
		o.Data2.Constructor = 1188831161
		return o
	},
	-1499615742: func() TLObject { // 0xa69dae02
		o := MakeTLDhGenFail(nil)
		o.Data2.Constructor = -1499615742
		return o
	},
	1182381663: func() TLObject { // 0x4679b65f
		o := MakeTLAccessPointRule(nil)
		o.Data2.Constructor = 1182381663
		return o
	},
	-1477445615: func() TLObject { // 0xa7eff811
		o := MakeTLBadMsgNotification(nil)
		o.Data2.Constructor = -1477445615
		return o
	},
	-307542917: func() TLObject { // 0xedab447b
		o := MakeTLBadServerSalt(nil)
		o.Data2.Constructor = -307542917
		return o
	},
	-501201412: func() TLObject { // 0xe22045fc
		o := MakeTLDestroySessionOk(nil)
		o.Data2.Constructor = -501201412
		return o
	},
	1658015945: func() TLObject { // 0x62d350c9
		o := MakeTLDestroySessionNone(nil)
		o.Data2.Constructor = 1658015945
		return o
	},
	155834844: func() TLObject { // 0x949d9dc
		o := MakeTLFutureSalt(nil)
		o.Data2.Constructor = 155834844
		return o
	},
	-1370486635: func() TLObject { // 0xae500895
		o := MakeTLFutureSalts(nil)
		o.Data2.Constructor = -1370486635
		return o
	},
	1515793004: func() TLObject { // 0x5a592a6c
		o := MakeTLHelpConfigSimple(nil)
		o.Data2.Constructor = 1515793004
		return o
	},
	-1835453025: func() TLObject { // 0x9299359f
		o := MakeTLHttpWait(nil)
		o.Data2.Constructor = -1835453025
		return o
	},
	-734810765: func() TLObject { // 0xd433ad73
		o := MakeTLIpPort(nil)
		o.Data2.Constructor = -734810765
		return o
	},
	932718150: func() TLObject { // 0x37982646
		o := MakeTLIpPortSecret(nil)
		o.Data2.Constructor = 932718150
		return o
	},
	661470918: func() TLObject { // 0x276d3ec6
		o := MakeTLMsgDetailedInfo(nil)
		o.Data2.Constructor = 661470918
		return o
	},
	-2137147681: func() TLObject { // 0x809db6df
		o := MakeTLMsgNewDetailedInfo(nil)
		o.Data2.Constructor = -2137147681
		return o
	},
	2105940488: func() TLObject { // 0x7d861a08
		o := MakeTLMsgResendReq(nil)
		o.Data2.Constructor = 2105940488
		return o
	},
	1658238041: func() TLObject { // 0x62d6b459
		o := MakeTLMsgsAck(nil)
		o.Data2.Constructor = 1658238041
		return o
	},
	-1933520591: func() TLObject { // 0x8cc0d131
		o := MakeTLMsgsAllInfo(nil)
		o.Data2.Constructor = -1933520591
		return o
	},
	81704317: func() TLObject { // 0x4deb57d
		o := MakeTLMsgsStateInfo(nil)
		o.Data2.Constructor = 81704317
		return o
	},
	-630588590: func() TLObject { // 0xda69fb52
		o := MakeTLMsgsStateReq(nil)
		o.Data2.Constructor = -630588590
		return o
	},
	-1631450872: func() TLObject { // 0x9ec20908
		o := MakeTLNewSessionCreated(nil)
		o.Data2.Constructor = -1631450872
		return o
	},
	880243653: func() TLObject { // 0x347773c5
		o := MakeTLPong(nil)
		o.Data2.Constructor = 880243653
		return o
	},
	1579864942: func() TLObject { // 0x5e2ad36e
		o := MakeTLRpcAnswerUnknown(nil)
		o.Data2.Constructor = 1579864942
		return o
	},
	-847714938: func() TLObject { // 0xcd78e586
		o := MakeTLRpcAnswerDroppedRunning(nil)
		o.Data2.Constructor = -847714938
		return o
	},
	-1539647305: func() TLObject { // 0xa43ad8b7
		o := MakeTLRpcAnswerDropped(nil)
		o.Data2.Constructor = -1539647305
		return o
	},
	558156313: func() TLObject { // 0x2144ca19
		o := MakeTLRpcError(nil)
		o.Data2.Constructor = 558156313
		return o
	},
	1108910436: func() TLObject { // 0x4218a164
		o := MakeTLTlsBlockString(nil)
		o.Data2.Constructor = 1108910436
		return o
	},
	1296942110: func() TLObject { // 0x4d4dc41e
		o := MakeTLTlsBlockRandom(nil)
		o.Data2.Constructor = 1296942110
		return o
	},
	154352379: func() TLObject { // 0x9333afb
		o := MakeTLTlsBlockZero(nil)
		o.Data2.Constructor = 154352379
		return o
	},
	283665263: func() TLObject { // 0x10e8636f
		o := MakeTLTlsBlockDomain(nil)
		o.Data2.Constructor = 283665263
		return o
	},
	-428498495: func() TLObject { // 0xe675a1c1
		o := MakeTLTlsBlockGrease(nil)
		o.Data2.Constructor = -428498495
		return o
	},
	-1632019620: func() TLObject { // 0x9eb95b5c
		o := MakeTLTlsBlockPublicKey(nil)
		o.Data2.Constructor = -1632019620
		return o
	},
	-416951217: func() TLObject { // 0xe725d44f
		o := MakeTLTlsBlockScope(nil)
		o.Data2.Constructor = -416951217
		return o
	},
	1817363588: func() TLObject { // 0x6c52c484
		o := MakeTLTlsClientHello(nil)
		o.Data2.Constructor = 1817363588
		return o
	},
	-1194283041: func() TLObject { // 0xb8d0afdf
		o := MakeTLAccountDaysTTL(nil)
		o.Data2.Constructor = -1194283041
		return o
	},
	-1389486888: func() TLObject { // 0xad2e1cd8
		o := MakeTLAccountAuthorizationForm(nil)
		o.Data2.Constructor = -1389486888
		return o
	},
	1275039392: func() TLObject { // 0x4bff8ea0
		o := MakeTLAccountAuthorizations(nil)
		o.Data2.Constructor = 1275039392
		return o
	},
	1674235686: func() TLObject { // 0x63cacf26
		o := MakeTLAccountAutoDownloadSettings(nil)
		o.Data2.Constructor = 1674235686
		return o
	},
	1279133341: func() TLObject { // 0x4c3e069d
		o := MakeTLAccountAutoSaveSettings(nil)
		o.Data2.Constructor = 1279133341
		return o
	},
	-331111727: func() TLObject { // 0xec43a2d1
		o := MakeTLAccountBusinessChatLinks(nil)
		o.Data2.Constructor = -331111727
		return o
	},
	400029819: func() TLObject { // 0x17d7f87b
		o := MakeTLAccountConnectedBots(nil)
		o.Data2.Constructor = 400029819
		return o
	},
	1474462241: func() TLObject { // 0x57e28221
		o := MakeTLAccountContentSettings(nil)
		o.Data2.Constructor = 1474462241
		return o
	},
	731303195: func() TLObject { // 0x2b96cd1b
		o := MakeTLAccountEmailVerified(nil)
		o.Data2.Constructor = 731303195
		return o
	},
	-507835039: func() TLObject { // 0xe1bb0d61
		o := MakeTLAccountEmailVerifiedLogin(nil)
		o.Data2.Constructor = -507835039
		return o
	},
	-796072379: func() TLObject { // 0xd08ce645
		o := MakeTLAccountEmojiStatusesNotModified(nil)
		o.Data2.Constructor = -796072379
		return o
	},
	-1866176559: func() TLObject { // 0x90c467d1
		o := MakeTLAccountEmojiStatuses(nil)
		o.Data2.Constructor = -1866176559
		return o
	},
	-1787080453: func() TLObject { // 0x957b50fb
		o := MakeTLAccountPassword(nil)
		o.Data2.Constructor = -1787080453
		return o
	},
	408623183: func() TLObject { // 0x185b184f
		o := MakeTLAccountPassword(nil)
		o.Data2.Constructor = 408623183
		return o
	},
	-1036572727: func() TLObject { // 0xc23727c9
		o := MakeTLAccountPasswordInputSettings(nil)
		o.Data2.Constructor = -1036572727
		return o
	},
	-1705233435: func() TLObject { // 0x9a5c33e5
		o := MakeTLAccountPasswordSettings(nil)
		o.Data2.Constructor = -1705233435
		return o
	},
	1352683077: func() TLObject { // 0x50a04e45
		o := MakeTLAccountPrivacyRules(nil)
		o.Data2.Constructor = 1352683077
		return o
	},
	-478701471: func() TLObject { // 0xe3779861
		o := MakeTLAccountResetPasswordFailedWait(nil)
		o.Data2.Constructor = -478701471
		return o
	},
	-370148227: func() TLObject { // 0xe9effc7d
		o := MakeTLAccountResetPasswordRequestedWait(nil)
		o.Data2.Constructor = -370148227
		return o
	},
	-383330754: func() TLObject { // 0xe926d63e
		o := MakeTLAccountResetPasswordOk(nil)
		o.Data2.Constructor = -383330754
		return o
	},
	-1708937439: func() TLObject { // 0x9a23af21
		o := MakeTLAccountResolvedBusinessChatLinks(nil)
		o.Data2.Constructor = -1708937439
		return o
	},
	-1222230163: func() TLObject { // 0xb7263f6d
		o := MakeTLAccountSavedRingtone(nil)
		o.Data2.Constructor = -1222230163
		return o
	},
	523271863: func() TLObject { // 0x1f307eb7
		o := MakeTLAccountSavedRingtoneConverted(nil)
		o.Data2.Constructor = 523271863
		return o
	},
	-67704655: func() TLObject { // 0xfbf6e8b1
		o := MakeTLAccountSavedRingtonesNotModified(nil)
		o.Data2.Constructor = -67704655
		return o
	},
	-1041683259: func() TLObject { // 0xc1e92cc5
		o := MakeTLAccountSavedRingtones(nil)
		o.Data2.Constructor = -1041683259
		return o
	},
	-2128640689: func() TLObject { // 0x811f854f
		o := MakeTLAccountSentEmailCode(nil)
		o.Data2.Constructor = -2128640689
		return o
	},
	1304052993: func() TLObject { // 0x4dba4501
		o := MakeTLAccountTakeout(nil)
		o.Data2.Constructor = 1304052993
		return o
	},
	-199313886: func() TLObject { // 0xf41eb622
		o := MakeTLAccountThemesNotModified(nil)
		o.Data2.Constructor = -199313886
		return o
	},
	-1707242387: func() TLObject { // 0x9a3d8c6d
		o := MakeTLAccountThemes(nil)
		o.Data2.Constructor = -1707242387
		return o
	},
	-614138572: func() TLObject { // 0xdb64fd34
		o := MakeTLAccountTmpPassword(nil)
		o.Data2.Constructor = -614138572
		return o
	},
	471437699: func() TLObject { // 0x1c199183
		o := MakeTLAccountWallPapersNotModified(nil)
		o.Data2.Constructor = 471437699
		return o
	},
	-842824308: func() TLObject { // 0xcdc3858c
		o := MakeTLAccountWallPapers(nil)
		o.Data2.Constructor = -842824308
		return o
	},
	-313079300: func() TLObject { // 0xed56c9fc
		o := MakeTLAccountWebAuthorizations(nil)
		o.Data2.Constructor = -313079300
		return o
	},
	1008422669: func() TLObject { // 0x3c1b4f0d
		o := MakeTLAppWebViewResultUrl(nil)
		o.Data2.Constructor = 1008422669
		return o
	},
	-653423106: func() TLObject { // 0xd90d8dfe
		o := MakeTLAttachMenuBot(nil)
		o.Data2.Constructor = -653423106
		return o
	},
	-928371502: func() TLObject { // 0xc8aa2cd2
		o := MakeTLAttachMenuBot(nil)
		o.Data2.Constructor = -928371502
		return o
	},
	-381896846: func() TLObject { // 0xe93cb772
		o := MakeTLAttachMenuBot(nil)
		o.Data2.Constructor = -381896846
		return o
	},
	-1297663893: func() TLObject { // 0xb2a7386b
		o := MakeTLAttachMenuBotIcon(nil)
		o.Data2.Constructor = -1297663893
		return o
	},
	1165423600: func() TLObject { // 0x4576f3f0
		o := MakeTLAttachMenuBotIconColor(nil)
		o.Data2.Constructor = 1165423600
		return o
	},
	-237467044: func() TLObject { // 0xf1d88a5c
		o := MakeTLAttachMenuBotsNotModified(nil)
		o.Data2.Constructor = -237467044
		return o
	},
	1011024320: func() TLObject { // 0x3c4301c0
		o := MakeTLAttachMenuBots(nil)
		o.Data2.Constructor = 1011024320
		return o
	},
	-1816172929: func() TLObject { // 0x93bf667f
		o := MakeTLAttachMenuBotsBot(nil)
		o.Data2.Constructor = -1816172929
		return o
	},
	2104224014: func() TLObject { // 0x7d6be90e
		o := MakeTLAttachMenuPeerTypeSameBotPM(nil)
		o.Data2.Constructor = 2104224014
		return o
	},
	-1020528102: func() TLObject { // 0xc32bfa1a
		o := MakeTLAttachMenuPeerTypeBotPM(nil)
		o.Data2.Constructor = -1020528102
		return o
	},
	-247016673: func() TLObject { // 0xf146d31f
		o := MakeTLAttachMenuPeerTypePM(nil)
		o.Data2.Constructor = -247016673
		return o
	},
	84480319: func() TLObject { // 0x509113f
		o := MakeTLAttachMenuPeerTypeChat(nil)
		o.Data2.Constructor = 84480319
		return o
	},
	2080104188: func() TLObject { // 0x7bfbdefc
		o := MakeTLAttachMenuPeerTypeBroadcast(nil)
		o.Data2.Constructor = 2080104188
		return o
	},
	-856756288: func() TLObject { // 0xcceeefc0
		o := MakeTLAuthKeyInfo(nil)
		o.Data2.Constructor = -856756288
		return o
	},
	782418132: func() TLObject { // 0x2ea2c0d4
		o := MakeTLAuthAuthorization(nil)
		o.Data2.Constructor = 782418132
		return o
	},
	872119224: func() TLObject { // 0x33fb7bb8
		o := MakeTLAuthAuthorization(nil)
		o.Data2.Constructor = 872119224
		return o
	},
	1148485274: func() TLObject { // 0x44747e9a
		o := MakeTLAuthAuthorizationSignUpRequired(nil)
		o.Data2.Constructor = 1148485274
		return o
	},
	1923290508: func() TLObject { // 0x72a3158c
		o := MakeTLAuthCodeTypeSms(nil)
		o.Data2.Constructor = 1923290508
		return o
	},
	1948046307: func() TLObject { // 0x741cd3e3
		o := MakeTLAuthCodeTypeCall(nil)
		o.Data2.Constructor = 1948046307
		return o
	},
	577556219: func() TLObject { // 0x226ccefb
		o := MakeTLAuthCodeTypeFlashCall(nil)
		o.Data2.Constructor = 577556219
		return o
	},
	-702884114: func() TLObject { // 0xd61ad6ee
		o := MakeTLAuthCodeTypeMissedCall(nil)
		o.Data2.Constructor = -702884114
		return o
	},
	116234636: func() TLObject { // 0x6ed998c
		o := MakeTLAuthCodeTypeFragmentSms(nil)
		o.Data2.Constructor = 116234636
		return o
	},
	-1271602504: func() TLObject { // 0xb434e2b8
		o := MakeTLAuthExportedAuthorization(nil)
		o.Data2.Constructor = -1271602504
		return o
	},
	-1012759713: func() TLObject { // 0xc3a2835f
		o := MakeTLAuthLoggedOut(nil)
		o.Data2.Constructor = -1012759713
		return o
	},
	1654593920: func() TLObject { // 0x629f1980
		o := MakeTLAuthLoginToken(nil)
		o.Data2.Constructor = 1654593920
		return o
	},
	110008598: func() TLObject { // 0x68e9916
		o := MakeTLAuthLoginTokenMigrateTo(nil)
		o.Data2.Constructor = 110008598
		return o
	},
	957176926: func() TLObject { // 0x390d5c5e
		o := MakeTLAuthLoginTokenSuccess(nil)
		o.Data2.Constructor = 957176926
		return o
	},
	326715557: func() TLObject { // 0x137948a5
		o := MakeTLAuthPasswordRecovery(nil)
		o.Data2.Constructor = 326715557
		return o
	},
	1577067778: func() TLObject { // 0x5e002502
		o := MakeTLAuthSentCode(nil)
		o.Data2.Constructor = 1577067778
		return o
	},
	596704836: func() TLObject { // 0x2390fe44
		o := MakeTLAuthSentCodeSuccess(nil)
		o.Data2.Constructor = 596704836
		return o
	},
	1035688326: func() TLObject { // 0x3dbb5986
		o := MakeTLAuthSentCodeTypeApp(nil)
		o.Data2.Constructor = 1035688326
		return o
	},
	-1073693790: func() TLObject { // 0xc000bba2
		o := MakeTLAuthSentCodeTypeSms(nil)
		o.Data2.Constructor = -1073693790
		return o
	},
	1398007207: func() TLObject { // 0x5353e5a7
		o := MakeTLAuthSentCodeTypeCall(nil)
		o.Data2.Constructor = 1398007207
		return o
	},
	-1425815847: func() TLObject { // 0xab03c6d9
		o := MakeTLAuthSentCodeTypeFlashCall(nil)
		o.Data2.Constructor = -1425815847
		return o
	},
	-2113903484: func() TLObject { // 0x82006484
		o := MakeTLAuthSentCodeTypeMissedCall(nil)
		o.Data2.Constructor = -2113903484
		return o
	},
	-196020837: func() TLObject { // 0xf450f59b
		o := MakeTLAuthSentCodeTypeEmailCode(nil)
		o.Data2.Constructor = -196020837
		return o
	},
	1511364673: func() TLObject { // 0x5a159841
		o := MakeTLAuthSentCodeTypeEmailCode(nil)
		o.Data2.Constructor = 1511364673
		return o
	},
	-1521934870: func() TLObject { // 0xa5491dea
		o := MakeTLAuthSentCodeTypeSetUpEmailRequired(nil)
		o.Data2.Constructor = -1521934870
		return o
	},
	-648651719: func() TLObject { // 0xd9565c39
		o := MakeTLAuthSentCodeTypeFragmentSms(nil)
		o.Data2.Constructor = -648651719
		return o
	},
	10475318: func() TLObject { // 0x9fd736
		o := MakeTLAuthSentCodeTypeFirebaseSms(nil)
		o.Data2.Constructor = 10475318
		return o
	},
	331943703: func() TLObject { // 0x13c90f17
		o := MakeTLAuthSentCodeTypeFirebaseSms(nil)
		o.Data2.Constructor = 331943703
		return o
	},
	-444918734: func() TLObject { // 0xe57b1432
		o := MakeTLAuthSentCodeTypeFirebaseSms(nil)
		o.Data2.Constructor = -444918734
		return o
	},
	-1542017919: func() TLObject { // 0xa416ac81
		o := MakeTLAuthSentCodeTypeSmsWord(nil)
		o.Data2.Constructor = -1542017919
		return o
	},
	-1284008785: func() TLObject { // 0xb37794af
		o := MakeTLAuthSentCodeTypeSmsPhrase(nil)
		o.Data2.Constructor = -1284008785
		return o
	},
	-1392388579: func() TLObject { // 0xad01d61d
		o := MakeTLAuthorization(nil)
		o.Data2.Constructor = -1392388579
		return o
	},
	-1163561432: func() TLObject { // 0xbaa57628
		o := MakeTLAutoDownloadSettings(nil)
		o.Data2.Constructor = -1163561432
		return o
	},
	-1896171181: func() TLObject { // 0x8efab953
		o := MakeTLAutoDownloadSettings(nil)
		o.Data2.Constructor = -1896171181
		return o
	},
	-532532493: func() TLObject { // 0xe04232f3
		o := MakeTLAutoDownloadSettings(nil)
		o.Data2.Constructor = -532532493
		return o
	},
	-2124403385: func() TLObject { // 0x81602d47
		o := MakeTLAutoSaveException(nil)
		o.Data2.Constructor = -2124403385
		return o
	},
	-934791986: func() TLObject { // 0xc84834ce
		o := MakeTLAutoSaveSettings(nil)
		o.Data2.Constructor = -934791986
		return o
	},
	-1815879042: func() TLObject { // 0x93c3e27e
		o := MakeTLAvailableEffect(nil)
		o.Data2.Constructor = -1815879042
		return o
	},
	-1065882623: func() TLObject { // 0xc077ec01
		o := MakeTLAvailableReaction(nil)
		o.Data2.Constructor = -1065882623
		return o
	},
	-177732982: func() TLObject { // 0xf568028a
		o := MakeTLBankCardOpenUrl(nil)
		o.Data2.Constructor = -177732982
		return o
	},
	-1012849566: func() TLObject { // 0xc3a12462
		o := MakeTLBaseThemeClassic(nil)
		o.Data2.Constructor = -1012849566
		return o
	},
	-69724536: func() TLObject { // 0xfbd81688
		o := MakeTLBaseThemeDay(nil)
		o.Data2.Constructor = -69724536
		return o
	},
	-1212997976: func() TLObject { // 0xb7b31ea8
		o := MakeTLBaseThemeNight(nil)
		o.Data2.Constructor = -1212997976
		return o
	},
	1834973166: func() TLObject { // 0x6d5f77ee
		o := MakeTLBaseThemeTinted(nil)
		o.Data2.Constructor = 1834973166
		return o
	},
	1527845466: func() TLObject { // 0x5b11125a
		o := MakeTLBaseThemeArctic(nil)
		o.Data2.Constructor = 1527845466
		return o
	},
	1821253126: func() TLObject { // 0x6c8e1e06
		o := MakeTLBirthday(nil)
		o.Data2.Constructor = 1821253126
		return o
	},
	1840491641: func() TLObject { // 0x6db3ac79
		o := MakeTLBizDataRaw(nil)
		o.Data2.Constructor = 1840491641
		return o
	},
	-1132882121: func() TLObject { // 0xbc799737
		o := MakeTLBoolFalse(nil)
		o.Data2.Constructor = -1132882121
		return o
	},
	-1720552011: func() TLObject { // 0x997275b5
		o := MakeTLBoolTrue(nil)
		o.Data2.Constructor = -1720552011
		return o
	},
	706514033: func() TLObject { // 0x2a1c8c71
		o := MakeTLBoost(nil)
		o.Data2.Constructor = 706514033
		return o
	},
	245261184: func() TLObject { // 0xe9e6380
		o := MakeTLBooster(nil)
		o.Data2.Constructor = 245261184
		return o
	},
	1571189943: func() TLObject { // 0x5da674b7
		o := MakeTLBotAppNotModified(nil)
		o.Data2.Constructor = 1571189943
		return o
	},
	-1778593322: func() TLObject { // 0x95fcd1d6
		o := MakeTLBotApp(nil)
		o.Data2.Constructor = -1778593322
		return o
	},
	-1989921868: func() TLObject { // 0x896433b4
		o := MakeTLBotBusinessConnection(nil)
		o.Data2.Constructor = -1989921868
		return o
	},
	-1032140601: func() TLObject { // 0xc27ac8c7
		o := MakeTLBotCommand(nil)
		o.Data2.Constructor = -1032140601
		return o
	},
	795652779: func() TLObject { // 0x2f6cb2ab
		o := MakeTLBotCommandScopeDefault(nil)
		o.Data2.Constructor = 795652779
		return o
	},
	1011811544: func() TLObject { // 0x3c4f04d8
		o := MakeTLBotCommandScopeUsers(nil)
		o.Data2.Constructor = 1011811544
		return o
	},
	1877059713: func() TLObject { // 0x6fe1a881
		o := MakeTLBotCommandScopeChats(nil)
		o.Data2.Constructor = 1877059713
		return o
	},
	-1180016534: func() TLObject { // 0xb9aa606a
		o := MakeTLBotCommandScopeChatAdmins(nil)
		o.Data2.Constructor = -1180016534
		return o
	},
	-610432643: func() TLObject { // 0xdb9d897d
		o := MakeTLBotCommandScopePeer(nil)
		o.Data2.Constructor = -610432643
		return o
	},
	1071145937: func() TLObject { // 0x3fd863d1
		o := MakeTLBotCommandScopePeerAdmins(nil)
		o.Data2.Constructor = 1071145937
		return o
	},
	169026035: func() TLObject { // 0xa1321f3
		o := MakeTLBotCommandScopePeerUser(nil)
		o.Data2.Constructor = 169026035
		return o
	},
	-319608864: func() TLObject { // 0xecf327e0
		o := MakeTLBotData(nil)
		o.Data2.Constructor = -319608864
		return o
	},
	-1892676777: func() TLObject { // 0x8f300b57
		o := MakeTLBotInfo(nil)
		o.Data2.Constructor = -1892676777
		return o
	},
	-468280483: func() TLObject { // 0xe4169b5d
		o := MakeTLBotInfo(nil)
		o.Data2.Constructor = -468280483
		return o
	},
	460632885: func() TLObject { // 0x1b74b335
		o := MakeTLBotInfo(nil)
		o.Data2.Constructor = 460632885
		return o
	},
	1984755728: func() TLObject { // 0x764cf810
		o := MakeTLBotInlineMessageMediaAuto(nil)
		o.Data2.Constructor = 1984755728
		return o
	},
	-1937807902: func() TLObject { // 0x8c7f65e2
		o := MakeTLBotInlineMessageText(nil)
		o.Data2.Constructor = -1937807902
		return o
	},
	85477117: func() TLObject { // 0x51846fd
		o := MakeTLBotInlineMessageMediaGeo(nil)
		o.Data2.Constructor = 85477117
		return o
	},
	-1970903652: func() TLObject { // 0x8a86659c
		o := MakeTLBotInlineMessageMediaVenue(nil)
		o.Data2.Constructor = -1970903652
		return o
	},
	416402882: func() TLObject { // 0x18d1cdc2
		o := MakeTLBotInlineMessageMediaContact(nil)
		o.Data2.Constructor = 416402882
		return o
	},
	894081801: func() TLObject { // 0x354a9b09
		o := MakeTLBotInlineMessageMediaInvoice(nil)
		o.Data2.Constructor = 894081801
		return o
	},
	-2137335386: func() TLObject { // 0x809ad9a6
		o := MakeTLBotInlineMessageMediaWebPage(nil)
		o.Data2.Constructor = -2137335386
		return o
	},
	295067450: func() TLObject { // 0x11965f3a
		o := MakeTLBotInlineResult(nil)
		o.Data2.Constructor = 295067450
		return o
	},
	400266251: func() TLObject { // 0x17db940b
		o := MakeTLBotInlineMediaResult(nil)
		o.Data2.Constructor = 400266251
		return o
	},
	1966318984: func() TLObject { // 0x7533a588
		o := MakeTLBotMenuButtonDefault(nil)
		o.Data2.Constructor = 1966318984
		return o
	},
	1113113093: func() TLObject { // 0x4258c205
		o := MakeTLBotMenuButtonCommands(nil)
		o.Data2.Constructor = 1113113093
		return o
	},
	-944407322: func() TLObject { // 0xc7b57ce6
		o := MakeTLBotMenuButton(nil)
		o.Data2.Constructor = -944407322
		return o
	},
	-391678544: func() TLObject { // 0xe8a775b0
		o := MakeTLBotsBotInfo(nil)
		o.Data2.Constructor = -391678544
		return o
	},
	-2076642874: func() TLObject { // 0x8438f1c6
		o := MakeTLBroadcastRevenueBalances(nil)
		o.Data2.Constructor = -2076642874
		return o
	},
	1434332356: func() TLObject { // 0x557e2cc4
		o := MakeTLBroadcastRevenueTransactionProceeds(nil)
		o.Data2.Constructor = 1434332356
		return o
	},
	1515784568: func() TLObject { // 0x5a590978
		o := MakeTLBroadcastRevenueTransactionWithdrawal(nil)
		o.Data2.Constructor = 1515784568
		return o
	},
	1121127726: func() TLObject { // 0x42d30d2e
		o := MakeTLBroadcastRevenueTransactionRefund(nil)
		o.Data2.Constructor = 1121127726
		return o
	},
	-283809188: func() TLObject { // 0xef156a5c
		o := MakeTLBusinessAwayMessage(nil)
		o.Data2.Constructor = -283809188
		return o
	},
	-910564679: func() TLObject { // 0xc9b9e2b9
		o := MakeTLBusinessAwayMessageScheduleAlways(nil)
		o.Data2.Constructor = -910564679
		return o
	},
	-1007487743: func() TLObject { // 0xc3f2f501
		o := MakeTLBusinessAwayMessageScheduleOutsideWorkHours(nil)
		o.Data2.Constructor = -1007487743
		return o
	},
	-867328308: func() TLObject { // 0xcc4d9ecc
		o := MakeTLBusinessAwayMessageScheduleCustom(nil)
		o.Data2.Constructor = -867328308
		return o
	},
	-1198722189: func() TLObject { // 0xb88cf373
		o := MakeTLBusinessBotRecipients(nil)
		o.Data2.Constructor = -1198722189
		return o
	},
	-1263638929: func() TLObject { // 0xb4ae666f
		o := MakeTLBusinessChatLink(nil)
		o.Data2.Constructor = -1263638929
		return o
	},
	-451302485: func() TLObject { // 0xe519abab
		o := MakeTLBusinessGreetingMessage(nil)
		o.Data2.Constructor = -451302485
		return o
	},
	1510606445: func() TLObject { // 0x5a0a066d
		o := MakeTLBusinessIntro(nil)
		o.Data2.Constructor = 1510606445
		return o
	},
	-1403249929: func() TLObject { // 0xac5c1af7
		o := MakeTLBusinessLocation(nil)
		o.Data2.Constructor = -1403249929
		return o
	},
	554733559: func() TLObject { // 0x21108ff7
		o := MakeTLBusinessRecipients(nil)
		o.Data2.Constructor = 554733559
		return o
	},
	302717625: func() TLObject { // 0x120b1ab9
		o := MakeTLBusinessWeeklyOpen(nil)
		o.Data2.Constructor = 302717625
		return o
	},
	-1936543592: func() TLObject { // 0x8c92b098
		o := MakeTLBusinessWorkHours(nil)
		o.Data2.Constructor = -1936543592
		return o
	},
	1462101002: func() TLObject { // 0x5725e40a
		o := MakeTLCdnConfig(nil)
		o.Data2.Constructor = 1462101002
		return o
	},
	-914167110: func() TLObject { // 0xc982eaba
		o := MakeTLCdnPublicKey(nil)
		o.Data2.Constructor = -914167110
		return o
	},
	531458253: func() TLObject { // 0x1fad68cd
		o := MakeTLChannelAdminLogEvent(nil)
		o.Data2.Constructor = 531458253
		return o
	},
	-421545947: func() TLObject { // 0xe6dfb825
		o := MakeTLChannelAdminLogEventActionChangeTitle(nil)
		o.Data2.Constructor = -421545947
		return o
	},
	1427671598: func() TLObject { // 0x55188a2e
		o := MakeTLChannelAdminLogEventActionChangeAbout(nil)
		o.Data2.Constructor = 1427671598
		return o
	},
	1783299128: func() TLObject { // 0x6a4afc38
		o := MakeTLChannelAdminLogEventActionChangeUsername(nil)
		o.Data2.Constructor = 1783299128
		return o
	},
	1129042607: func() TLObject { // 0x434bd2af
		o := MakeTLChannelAdminLogEventActionChangePhoto(nil)
		o.Data2.Constructor = 1129042607
		return o
	},
	460916654: func() TLObject { // 0x1b7907ae
		o := MakeTLChannelAdminLogEventActionToggleInvites(nil)
		o.Data2.Constructor = 460916654
		return o
	},
	648939889: func() TLObject { // 0x26ae0971
		o := MakeTLChannelAdminLogEventActionToggleSignatures(nil)
		o.Data2.Constructor = 648939889
		return o
	},
	-370660328: func() TLObject { // 0xe9e82c18
		o := MakeTLChannelAdminLogEventActionUpdatePinned(nil)
		o.Data2.Constructor = -370660328
		return o
	},
	1889215493: func() TLObject { // 0x709b2405
		o := MakeTLChannelAdminLogEventActionEditMessage(nil)
		o.Data2.Constructor = 1889215493
		return o
	},
	1121994683: func() TLObject { // 0x42e047bb
		o := MakeTLChannelAdminLogEventActionDeleteMessage(nil)
		o.Data2.Constructor = 1121994683
		return o
	},
	405815507: func() TLObject { // 0x183040d3
		o := MakeTLChannelAdminLogEventActionParticipantJoin(nil)
		o.Data2.Constructor = 405815507
		return o
	},
	-124291086: func() TLObject { // 0xf89777f2
		o := MakeTLChannelAdminLogEventActionParticipantLeave(nil)
		o.Data2.Constructor = -124291086
		return o
	},
	-484690728: func() TLObject { // 0xe31c34d8
		o := MakeTLChannelAdminLogEventActionParticipantInvite(nil)
		o.Data2.Constructor = -484690728
		return o
	},
	-422036098: func() TLObject { // 0xe6d83d7e
		o := MakeTLChannelAdminLogEventActionParticipantToggleBan(nil)
		o.Data2.Constructor = -422036098
		return o
	},
	-714643696: func() TLObject { // 0xd5676710
		o := MakeTLChannelAdminLogEventActionParticipantToggleAdmin(nil)
		o.Data2.Constructor = -714643696
		return o
	},
	-1312568665: func() TLObject { // 0xb1c3caa7
		o := MakeTLChannelAdminLogEventActionChangeStickerSet(nil)
		o.Data2.Constructor = -1312568665
		return o
	},
	1599903217: func() TLObject { // 0x5f5c95f1
		o := MakeTLChannelAdminLogEventActionTogglePreHistoryHidden(nil)
		o.Data2.Constructor = 1599903217
		return o
	},
	771095562: func() TLObject { // 0x2df5fc0a
		o := MakeTLChannelAdminLogEventActionDefaultBannedRights(nil)
		o.Data2.Constructor = 771095562
		return o
	},
	-1895328189: func() TLObject { // 0x8f079643
		o := MakeTLChannelAdminLogEventActionStopPoll(nil)
		o.Data2.Constructor = -1895328189
		return o
	},
	84703944: func() TLObject { // 0x50c7ac8
		o := MakeTLChannelAdminLogEventActionChangeLinkedChat(nil)
		o.Data2.Constructor = 84703944
		return o
	},
	241923758: func() TLObject { // 0xe6b76ae
		o := MakeTLChannelAdminLogEventActionChangeLocation(nil)
		o.Data2.Constructor = 241923758
		return o
	},
	1401984889: func() TLObject { // 0x53909779
		o := MakeTLChannelAdminLogEventActionToggleSlowMode(nil)
		o.Data2.Constructor = 1401984889
		return o
	},
	589338437: func() TLObject { // 0x23209745
		o := MakeTLChannelAdminLogEventActionStartGroupCall(nil)
		o.Data2.Constructor = 589338437
		return o
	},
	-610299584: func() TLObject { // 0xdb9f9140
		o := MakeTLChannelAdminLogEventActionDiscardGroupCall(nil)
		o.Data2.Constructor = -610299584
		return o
	},
	-115071790: func() TLObject { // 0xf92424d2
		o := MakeTLChannelAdminLogEventActionParticipantMute(nil)
		o.Data2.Constructor = -115071790
		return o
	},
	-431740480: func() TLObject { // 0xe64429c0
		o := MakeTLChannelAdminLogEventActionParticipantUnmute(nil)
		o.Data2.Constructor = -431740480
		return o
	},
	1456906823: func() TLObject { // 0x56d6a247
		o := MakeTLChannelAdminLogEventActionToggleGroupCallSetting(nil)
		o.Data2.Constructor = 1456906823
		return o
	},
	-23084712: func() TLObject { // 0xfe9fc158
		o := MakeTLChannelAdminLogEventActionParticipantJoinByInvite(nil)
		o.Data2.Constructor = -23084712
		return o
	},
	1557846647: func() TLObject { // 0x5cdada77
		o := MakeTLChannelAdminLogEventActionParticipantJoinByInvite(nil)
		o.Data2.Constructor = 1557846647
		return o
	},
	1515256996: func() TLObject { // 0x5a50fca4
		o := MakeTLChannelAdminLogEventActionExportedInviteDelete(nil)
		o.Data2.Constructor = 1515256996
		return o
	},
	1091179342: func() TLObject { // 0x410a134e
		o := MakeTLChannelAdminLogEventActionExportedInviteRevoke(nil)
		o.Data2.Constructor = 1091179342
		return o
	},
	-384910503: func() TLObject { // 0xe90ebb59
		o := MakeTLChannelAdminLogEventActionExportedInviteEdit(nil)
		o.Data2.Constructor = -384910503
		return o
	},
	1048537159: func() TLObject { // 0x3e7f6847
		o := MakeTLChannelAdminLogEventActionParticipantVolume(nil)
		o.Data2.Constructor = 1048537159
		return o
	},
	1855199800: func() TLObject { // 0x6e941a38
		o := MakeTLChannelAdminLogEventActionChangeHistoryTTL(nil)
		o.Data2.Constructor = 1855199800
		return o
	},
	-1347021750: func() TLObject { // 0xafb6144a
		o := MakeTLChannelAdminLogEventActionParticipantJoinByRequest(nil)
		o.Data2.Constructor = -1347021750
		return o
	},
	-886388890: func() TLObject { // 0xcb2ac766
		o := MakeTLChannelAdminLogEventActionToggleNoForwards(nil)
		o.Data2.Constructor = -886388890
		return o
	},
	663693416: func() TLObject { // 0x278f2868
		o := MakeTLChannelAdminLogEventActionSendMessage(nil)
		o.Data2.Constructor = 663693416
		return o
	},
	-1102180616: func() TLObject { // 0xbe4e0ef8
		o := MakeTLChannelAdminLogEventActionChangeAvailableReactions(nil)
		o.Data2.Constructor = -1102180616
		return o
	},
	-1661470870: func() TLObject { // 0x9cf7f76a
		o := MakeTLChannelAdminLogEventActionChangeAvailableReactions(nil)
		o.Data2.Constructor = -1661470870
		return o
	},
	-263212119: func() TLObject { // 0xf04fb3a9
		o := MakeTLChannelAdminLogEventActionChangeUsernames(nil)
		o.Data2.Constructor = -263212119
		return o
	},
	46949251: func() TLObject { // 0x2cc6383
		o := MakeTLChannelAdminLogEventActionToggleForum(nil)
		o.Data2.Constructor = 46949251
		return o
	},
	1483767080: func() TLObject { // 0x58707d28
		o := MakeTLChannelAdminLogEventActionCreateTopic(nil)
		o.Data2.Constructor = 1483767080
		return o
	},
	-261103096: func() TLObject { // 0xf06fe208
		o := MakeTLChannelAdminLogEventActionEditTopic(nil)
		o.Data2.Constructor = -261103096
		return o
	},
	-1374254839: func() TLObject { // 0xae168909
		o := MakeTLChannelAdminLogEventActionDeleteTopic(nil)
		o.Data2.Constructor = -1374254839
		return o
	},
	1569535291: func() TLObject { // 0x5d8d353b
		o := MakeTLChannelAdminLogEventActionPinTopic(nil)
		o.Data2.Constructor = 1569535291
		return o
	},
	1693675004: func() TLObject { // 0x64f36dfc
		o := MakeTLChannelAdminLogEventActionToggleAntiSpam(nil)
		o.Data2.Constructor = 1693675004
		return o
	},
	1469507456: func() TLObject { // 0x5796e780
		o := MakeTLChannelAdminLogEventActionChangePeerColor(nil)
		o.Data2.Constructor = 1469507456
		return o
	},
	1581742885: func() TLObject { // 0x5e477b25
		o := MakeTLChannelAdminLogEventActionChangeProfilePeerColor(nil)
		o.Data2.Constructor = 1581742885
		return o
	},
	834362706: func() TLObject { // 0x31bb5d52
		o := MakeTLChannelAdminLogEventActionChangeWallpaper(nil)
		o.Data2.Constructor = 834362706
		return o
	},
	1051328177: func() TLObject { // 0x3ea9feb1
		o := MakeTLChannelAdminLogEventActionChangeEmojiStatus(nil)
		o.Data2.Constructor = 1051328177
		return o
	},
	1188577451: func() TLObject { // 0x46d840ab
		o := MakeTLChannelAdminLogEventActionChangeEmojiStickerSet(nil)
		o.Data2.Constructor = 1188577451
		return o
	},
	1009460347: func() TLObject { // 0x3c2b247b
		o := MakeTLChannelAdminLogEventActionChangeColor(nil)
		o.Data2.Constructor = 1009460347
		return o
	},
	1147126836: func() TLObject { // 0x445fc434
		o := MakeTLChannelAdminLogEventActionChangeBackgroundEmoji(nil)
		o.Data2.Constructor = 1147126836
		return o
	},
	-368018716: func() TLObject { // 0xea107ae4
		o := MakeTLChannelAdminLogEventsFilter(nil)
		o.Data2.Constructor = -368018716
		return o
	},
	-1078612597: func() TLObject { // 0xbfb5ad8b
		o := MakeTLChannelLocationEmpty(nil)
		o.Data2.Constructor = -1078612597
		return o
	},
	547062491: func() TLObject { // 0x209b82db
		o := MakeTLChannelLocation(nil)
		o.Data2.Constructor = 547062491
		return o
	},
	-1798033689: func() TLObject { // 0x94d42ee7
		o := MakeTLChannelMessagesFilterEmpty(nil)
		o.Data2.Constructor = -1798033689
		return o
	},
	-847783593: func() TLObject { // 0xcd77d957
		o := MakeTLChannelMessagesFilter(nil)
		o.Data2.Constructor = -847783593
		return o
	},
	-1072953408: func() TLObject { // 0xc00c07c0
		o := MakeTLChannelParticipant(nil)
		o.Data2.Constructor = -1072953408
		return o
	},
	900251559: func() TLObject { // 0x35a8bfa7
		o := MakeTLChannelParticipantSelf(nil)
		o.Data2.Constructor = 900251559
		return o
	},
	803602899: func() TLObject { // 0x2fe601d3
		o := MakeTLChannelParticipantCreator(nil)
		o.Data2.Constructor = 803602899
		return o
	},
	885242707: func() TLObject { // 0x34c3bb53
		o := MakeTLChannelParticipantAdmin(nil)
		o.Data2.Constructor = 885242707
		return o
	},
	1844969806: func() TLObject { // 0x6df8014e
		o := MakeTLChannelParticipantBanned(nil)
		o.Data2.Constructor = 1844969806
		return o
	},
	453242886: func() TLObject { // 0x1b03f006
		o := MakeTLChannelParticipantLeft(nil)
		o.Data2.Constructor = 453242886
		return o
	},
	-566281095: func() TLObject { // 0xde3f3c79
		o := MakeTLChannelParticipantsRecent(nil)
		o.Data2.Constructor = -566281095
		return o
	},
	-1268741783: func() TLObject { // 0xb4608969
		o := MakeTLChannelParticipantsAdmins(nil)
		o.Data2.Constructor = -1268741783
		return o
	},
	-1548400251: func() TLObject { // 0xa3b54985
		o := MakeTLChannelParticipantsKicked(nil)
		o.Data2.Constructor = -1548400251
		return o
	},
	-1328445861: func() TLObject { // 0xb0d1865b
		o := MakeTLChannelParticipantsBots(nil)
		o.Data2.Constructor = -1328445861
		return o
	},
	338142689: func() TLObject { // 0x1427a5e1
		o := MakeTLChannelParticipantsBanned(nil)
		o.Data2.Constructor = 338142689
		return o
	},
	106343499: func() TLObject { // 0x656ac4b
		o := MakeTLChannelParticipantsSearch(nil)
		o.Data2.Constructor = 106343499
		return o
	},
	-1150621555: func() TLObject { // 0xbb6ae88d
		o := MakeTLChannelParticipantsContacts(nil)
		o.Data2.Constructor = -1150621555
		return o
	},
	-531931925: func() TLObject { // 0xe04b5ceb
		o := MakeTLChannelParticipantsMentions(nil)
		o.Data2.Constructor = -531931925
		return o
	},
	-309659827: func() TLObject { // 0xed8af74d
		o := MakeTLChannelsAdminLogResults(nil)
		o.Data2.Constructor = -309659827
		return o
	},
	-541588713: func() TLObject { // 0xdfb80317
		o := MakeTLChannelsChannelParticipant(nil)
		o.Data2.Constructor = -541588713
		return o
	},
	-1699676497: func() TLObject { // 0x9ab0feaf
		o := MakeTLChannelsChannelParticipants(nil)
		o.Data2.Constructor = -1699676497
		return o
	},
	-266911767: func() TLObject { // 0xf0173fe9
		o := MakeTLChannelsChannelParticipantsNotModified(nil)
		o.Data2.Constructor = -266911767
		return o
	},
	-191450938: func() TLObject { // 0xf496b0c6
		o := MakeTLChannelsSendAsPeers(nil)
		o.Data2.Constructor = -191450938
		return o
	},
	-2091463255: func() TLObject { // 0x8356cda9
		o := MakeTLChannelsSendAsPeers(nil)
		o.Data2.Constructor = -2091463255
		return o
	},
	-2073059774: func() TLObject { // 0x846f9e42
		o := MakeTLChannelsSponsoredMessageReportResultChooseOption(nil)
		o.Data2.Constructor = -2073059774
		return o
	},
	1044107055: func() TLObject { // 0x3e3bcf2f
		o := MakeTLChannelsSponsoredMessageReportResultAdsHidden(nil)
		o.Data2.Constructor = 1044107055
		return o
	},
	-1384544183: func() TLObject { // 0xad798849
		o := MakeTLChannelsSponsoredMessageReportResultReported(nil)
		o.Data2.Constructor = -1384544183
		return o
	},
	693512293: func() TLObject { // 0x29562865
		o := MakeTLChatEmpty(nil)
		o.Data2.Constructor = 693512293
		return o
	},
	1103884886: func() TLObject { // 0x41cbf256
		o := MakeTLChat(nil)
		o.Data2.Constructor = 1103884886
		return o
	},
	1704108455: func() TLObject { // 0x6592a1a7
		o := MakeTLChatForbidden(nil)
		o.Data2.Constructor = 1704108455
		return o
	},
	179174543: func() TLObject { // 0xaadfc8f
		o := MakeTLChannel(nil)
		o.Data2.Constructor = 179174543
		return o
	},
	-1903702824: func() TLObject { // 0x8e87ccd8
		o := MakeTLChannel(nil)
		o.Data2.Constructor = -1903702824
		return o
	},
	427944574: func() TLObject { // 0x1981ea7e
		o := MakeTLChannel(nil)
		o.Data2.Constructor = 427944574
		return o
	},
	-1795845413: func() TLObject { // 0x94f592db
		o := MakeTLChannel(nil)
		o.Data2.Constructor = -1795845413
		return o
	},
	-2094689180: func() TLObject { // 0x83259464
		o := MakeTLChannel(nil)
		o.Data2.Constructor = -2094689180
		return o
	},
	-2107528095: func() TLObject { // 0x8261ac61
		o := MakeTLChannel(nil)
		o.Data2.Constructor = -2107528095
		return o
	},
	399807445: func() TLObject { // 0x17d493d5
		o := MakeTLChannelForbidden(nil)
		o.Data2.Constructor = 399807445
		return o
	},
	1605510357: func() TLObject { // 0x5fb224d5
		o := MakeTLChatAdminRights(nil)
		o.Data2.Constructor = 1605510357
		return o
	},
	-219353309: func() TLObject { // 0xf2ecef23
		o := MakeTLChatAdminWithInvites(nil)
		o.Data2.Constructor = -219353309
		return o
	},
	-1626209256: func() TLObject { // 0x9f120418
		o := MakeTLChatBannedRights(nil)
		o.Data2.Constructor = -1626209256
		return o
	},
	640893467: func() TLObject { // 0x2633421b
		o := MakeTLChatFull(nil)
		o.Data2.Constructor = 640893467
		return o
	},
	-908914376: func() TLObject { // 0xc9d31138
		o := MakeTLChatFull(nil)
		o.Data2.Constructor = -908914376
		return o
	},
	-779165146: func() TLObject { // 0xd18ee226
		o := MakeTLChatFull(nil)
		o.Data2.Constructor = -779165146
		return o
	},
	-1146407795: func() TLObject { // 0xbbab348d
		o := MakeTLChannelFull(nil)
		o.Data2.Constructor = -1146407795
		return o
	},
	1153455271: func() TLObject { // 0x44c054a7
		o := MakeTLChannelFull(nil)
		o.Data2.Constructor = 1153455271
		return o
	},
	254528367: func() TLObject { // 0xf2bcb6f
		o := MakeTLChannelFull(nil)
		o.Data2.Constructor = 254528367
		return o
	},
	1915758525: func() TLObject { // 0x723027bd
		o := MakeTLChannelFull(nil)
		o.Data2.Constructor = 1915758525
		return o
	},
	-231385849: func() TLObject { // 0xf2355507
		o := MakeTLChannelFull(nil)
		o.Data2.Constructor = -231385849
		return o
	},
	-362240487: func() TLObject { // 0xea68a619
		o := MakeTLChannelFull(nil)
		o.Data2.Constructor = -362240487
		return o
	},
	-516145888: func() TLObject { // 0xe13c3d20
		o := MakeTLChannelFull(nil)
		o.Data2.Constructor = -516145888
		return o
	},
	1516793212: func() TLObject { // 0x5a686d7c
		o := MakeTLChatInviteAlready(nil)
		o.Data2.Constructor = 1516793212
		return o
	},
	-840897472: func() TLObject { // 0xcde0ec40
		o := MakeTLChatInvite(nil)
		o.Data2.Constructor = -840897472
		return o
	},
	806110401: func() TLObject { // 0x300c44c1
		o := MakeTLChatInvite(nil)
		o.Data2.Constructor = 806110401
		return o
	},
	1634294960: func() TLObject { // 0x61695cb0
		o := MakeTLChatInvitePeek(nil)
		o.Data2.Constructor = 1634294960
		return o
	},
	-1940201511: func() TLObject { // 0x8c5adfd9
		o := MakeTLChatInviteImporter(nil)
		o.Data2.Constructor = -1940201511
		return o
	},
	-264117680: func() TLObject { // 0xf041e250
		o := MakeTLChatOnlines(nil)
		o.Data2.Constructor = -264117680
		return o
	},
	-1070776313: func() TLObject { // 0xc02d4007
		o := MakeTLChatParticipant(nil)
		o.Data2.Constructor = -1070776313
		return o
	},
	-462696732: func() TLObject { // 0xe46bcee4
		o := MakeTLChatParticipantCreator(nil)
		o.Data2.Constructor = -462696732
		return o
	},
	-1600962725: func() TLObject { // 0xa0933f5b
		o := MakeTLChatParticipantAdmin(nil)
		o.Data2.Constructor = -1600962725
		return o
	},
	-2023500831: func() TLObject { // 0x8763d3e1
		o := MakeTLChatParticipantsForbidden(nil)
		o.Data2.Constructor = -2023500831
		return o
	},
	1018991608: func() TLObject { // 0x3cbc93f8
		o := MakeTLChatParticipants(nil)
		o.Data2.Constructor = 1018991608
		return o
	},
	935395612: func() TLObject { // 0x37c1011c
		o := MakeTLChatPhotoEmpty(nil)
		o.Data2.Constructor = 935395612
		return o
	},
	476978193: func() TLObject { // 0x1c6e1c11
		o := MakeTLChatPhoto(nil)
		o.Data2.Constructor = 476978193
		return o
	},
	-352570692: func() TLObject { // 0xeafc32bc
		o := MakeTLChatReactionsNone(nil)
		o.Data2.Constructor = -352570692
		return o
	},
	1385335754: func() TLObject { // 0x52928bca
		o := MakeTLChatReactionsAll(nil)
		o.Data2.Constructor = 1385335754
		return o
	},
	1713193015: func() TLObject { // 0x661d4037
		o := MakeTLChatReactionsSome(nil)
		o.Data2.Constructor = 1713193015
		return o
	},
	-91752871: func() TLObject { // 0xfa87f659
		o := MakeTLChatlistsChatlistInviteAlready(nil)
		o.Data2.Constructor = -91752871
		return o
	},
	500007837: func() TLObject { // 0x1dcd839d
		o := MakeTLChatlistsChatlistInvite(nil)
		o.Data2.Constructor = 500007837
		return o
	},
	-1816295539: func() TLObject { // 0x93bd878d
		o := MakeTLChatlistsChatlistUpdates(nil)
		o.Data2.Constructor = -1816295539
		return o
	},
	283567014: func() TLObject { // 0x10e6e3a6
		o := MakeTLChatlistsExportedChatlistInvite(nil)
		o.Data2.Constructor = 283567014
		return o
	},
	279670215: func() TLObject { // 0x10ab6dc7
		o := MakeTLChatlistsExportedInvites(nil)
		o.Data2.Constructor = 279670215
		return o
	},
	-1390068360: func() TLObject { // 0xad253d78
		o := MakeTLCodeSettings(nil)
		o.Data2.Constructor = -1390068360
		return o
	},
	-1973130814: func() TLObject { // 0x8a6469c2
		o := MakeTLCodeSettings(nil)
		o.Data2.Constructor = -1973130814
		return o
	},
	-870702050: func() TLObject { // 0xcc1a241e
		o := MakeTLConfig(nil)
		o.Data2.Constructor = -870702050
		return o
	},
	589653676: func() TLObject { // 0x232566ac
		o := MakeTLConfig(nil)
		o.Data2.Constructor = 589653676
		return o
	},
	856375399: func() TLObject { // 0x330b4067
		o := MakeTLConfig(nil)
		o.Data2.Constructor = 856375399
		return o
	},
	-1123645951: func() TLObject { // 0xbd068601
		o := MakeTLConnectedBot(nil)
		o.Data2.Constructor = -1123645951
		return o
	},
	-404121113: func() TLObject { // 0xe7e999e7
		o := MakeTLConnectedBot(nil)
		o.Data2.Constructor = -404121113
		return o
	},
	341499403: func() TLObject { // 0x145ade0b
		o := MakeTLContact(nil)
		o.Data2.Constructor = 341499403
		return o
	},
	496600883: func() TLObject { // 0x1d998733
		o := MakeTLContactBirthday(nil)
		o.Data2.Constructor = 496600883
		return o
	},
	-858039014: func() TLObject { // 0xccdb5d1a
		o := MakeTLContactData(nil)
		o.Data2.Constructor = -858039014
		return o
	},
	383348795: func() TLObject { // 0x16d9703b
		o := MakeTLContactStatus(nil)
		o.Data2.Constructor = 383348795
		return o
	},
	182326673: func() TLObject { // 0xade1591
		o := MakeTLContactsBlocked(nil)
		o.Data2.Constructor = 182326673
		return o
	},
	-513392236: func() TLObject { // 0xe1664194
		o := MakeTLContactsBlockedSlice(nil)
		o.Data2.Constructor = -513392236
		return o
	},
	290452237: func() TLObject { // 0x114ff30d
		o := MakeTLContactsContactBirthdays(nil)
		o.Data2.Constructor = 290452237
		return o
	},
	-1219778094: func() TLObject { // 0xb74ba9d2
		o := MakeTLContactsContactsNotModified(nil)
		o.Data2.Constructor = -1219778094
		return o
	},
	-353862078: func() TLObject { // 0xeae87e42
		o := MakeTLContactsContacts(nil)
		o.Data2.Constructor = -353862078
		return o
	},
	-1290580579: func() TLObject { // 0xb3134d9d
		o := MakeTLContactsFound(nil)
		o.Data2.Constructor = -1290580579
		return o
	},
	2010127419: func() TLObject { // 0x77d01c3b
		o := MakeTLContactsImportedContacts(nil)
		o.Data2.Constructor = 2010127419
		return o
	},
	2131196633: func() TLObject { // 0x7f077ad9
		o := MakeTLContactsResolvedPeer(nil)
		o.Data2.Constructor = 2131196633
		return o
	},
	-567906571: func() TLObject { // 0xde266ef5
		o := MakeTLContactsTopPeersNotModified(nil)
		o.Data2.Constructor = -567906571
		return o
	},
	1891070632: func() TLObject { // 0x70b772a8
		o := MakeTLContactsTopPeers(nil)
		o.Data2.Constructor = 1891070632
		return o
	},
	-1255369827: func() TLObject { // 0xb52c939d
		o := MakeTLContactsTopPeersDisabled(nil)
		o.Data2.Constructor = -1255369827
		return o
	},
	2104790276: func() TLObject { // 0x7d748d04
		o := MakeTLDataJSON(nil)
		o.Data2.Constructor = 2104790276
		return o
	},
	414687501: func() TLObject { // 0x18b7a10d
		o := MakeTLDcOption(nil)
		o.Data2.Constructor = 414687501
		return o
	},
	1135897376: func() TLObject { // 0x43b46b20
		o := MakeTLDefaultHistoryTTL(nil)
		o.Data2.Constructor = 1135897376
		return o
	},
	-712374074: func() TLObject { // 0xd58a08c6
		o := MakeTLDialog(nil)
		o.Data2.Constructor = -712374074
		return o
	},
	-1460809483: func() TLObject { // 0xa8edd0f5
		o := MakeTLDialog(nil)
		o.Data2.Constructor = -1460809483
		return o
	},
	1908216652: func() TLObject { // 0x71bd134c
		o := MakeTLDialogFolder(nil)
		o.Data2.Constructor = 1908216652
		return o
	},
	1605718587: func() TLObject { // 0x5fb5523b
		o := MakeTLDialogFilter(nil)
		o.Data2.Constructor = 1605718587
		return o
	},
	1949890536: func() TLObject { // 0x7438f7e8
		o := MakeTLDialogFilter(nil)
		o.Data2.Constructor = 1949890536
		return o
	},
	909284270: func() TLObject { // 0x363293ae
		o := MakeTLDialogFilterDefault(nil)
		o.Data2.Constructor = 909284270
		return o
	},
	-1612542300: func() TLObject { // 0x9fe28ea4
		o := MakeTLDialogFilterChatlist(nil)
		o.Data2.Constructor = -1612542300
		return o
	},
	-699792216: func() TLObject { // 0xd64a04a8
		o := MakeTLDialogFilterChatlist(nil)
		o.Data2.Constructor = -699792216
		return o
	},
	2004110666: func() TLObject { // 0x77744d4a
		o := MakeTLDialogFilterSuggested(nil)
		o.Data2.Constructor = 2004110666
		return o
	},
	-445792507: func() TLObject { // 0xe56dbf05
		o := MakeTLDialogPeer(nil)
		o.Data2.Constructor = -445792507
		return o
	},
	1363483106: func() TLObject { // 0x514519e2
		o := MakeTLDialogPeerFolder(nil)
		o.Data2.Constructor = 1363483106
		return o
	},
	922273905: func() TLObject { // 0x36f8c871
		o := MakeTLDocumentEmpty(nil)
		o.Data2.Constructor = 922273905
		return o
	},
	-1881881384: func() TLObject { // 0x8fd4c4d8
		o := MakeTLDocument(nil)
		o.Data2.Constructor = -1881881384
		return o
	},
	512177195: func() TLObject { // 0x1e87342b
		o := MakeTLDocument(nil)
		o.Data2.Constructor = 512177195
		return o
	},
	1815593308: func() TLObject { // 0x6c37c15c
		o := MakeTLDocumentAttributeImageSize(nil)
		o.Data2.Constructor = 1815593308
		return o
	},
	297109817: func() TLObject { // 0x11b58939
		o := MakeTLDocumentAttributeAnimated(nil)
		o.Data2.Constructor = 297109817
		return o
	},
	1662637586: func() TLObject { // 0x6319d612
		o := MakeTLDocumentAttributeSticker(nil)
		o.Data2.Constructor = 1662637586
		return o
	},
	-745541182: func() TLObject { // 0xd38ff1c2
		o := MakeTLDocumentAttributeVideo(nil)
		o.Data2.Constructor = -745541182
		return o
	},
	250621158: func() TLObject { // 0xef02ce6
		o := MakeTLDocumentAttributeVideo(nil)
		o.Data2.Constructor = 250621158
		return o
	},
	-1739392570: func() TLObject { // 0x9852f9c6
		o := MakeTLDocumentAttributeAudio(nil)
		o.Data2.Constructor = -1739392570
		return o
	},
	358154344: func() TLObject { // 0x15590068
		o := MakeTLDocumentAttributeFilename(nil)
		o.Data2.Constructor = 358154344
		return o
	},
	-1744710921: func() TLObject { // 0x9801d2f7
		o := MakeTLDocumentAttributeHasStickers(nil)
		o.Data2.Constructor = -1744710921
		return o
	},
	-48981863: func() TLObject { // 0xfd149899
		o := MakeTLDocumentAttributeCustomEmoji(nil)
		o.Data2.Constructor = -48981863
		return o
	},
	1431132616: func() TLObject { // 0x554d59c8
		o := MakeTLDouble(nil)
		o.Data2.Constructor = 1431132616
		return o
	},
	453805082: func() TLObject { // 0x1b0c841a
		o := MakeTLDraftMessageEmpty(nil)
		o.Data2.Constructor = 453805082
		return o
	},
	761606687: func() TLObject { // 0x2d65321f
		o := MakeTLDraftMessage(nil)
		o.Data2.Constructor = 761606687
		return o
	},
	1070397423: func() TLObject { // 0x3fccf7ef
		o := MakeTLDraftMessage(nil)
		o.Data2.Constructor = 1070397423
		return o
	},
	-40996577: func() TLObject { // 0xfd8e711f
		o := MakeTLDraftMessage(nil)
		o.Data2.Constructor = -40996577
		return o
	},
	-1842457175: func() TLObject { // 0x922e55a9
		o := MakeTLEmailVerificationCode(nil)
		o.Data2.Constructor = -1842457175
		return o
	},
	-611279166: func() TLObject { // 0xdb909ec2
		o := MakeTLEmailVerificationGoogle(nil)
		o.Data2.Constructor = -611279166
		return o
	},
	-1764723459: func() TLObject { // 0x96d074fd
		o := MakeTLEmailVerificationApple(nil)
		o.Data2.Constructor = -1764723459
		return o
	},
	1128644211: func() TLObject { // 0x4345be73
		o := MakeTLEmailVerifyPurposeLoginSetup(nil)
		o.Data2.Constructor = 1128644211
		return o
	},
	1383932651: func() TLObject { // 0x527d22eb
		o := MakeTLEmailVerifyPurposeLoginChange(nil)
		o.Data2.Constructor = 1383932651
		return o
	},
	-1141565819: func() TLObject { // 0xbbf51685
		o := MakeTLEmailVerifyPurposePassport(nil)
		o.Data2.Constructor = -1141565819
		return o
	},
	2056961449: func() TLObject { // 0x7a9abda9
		o := MakeTLEmojiGroup(nil)
		o.Data2.Constructor = 2056961449
		return o
	},
	-2133693241: func() TLObject { // 0x80d26cc7
		o := MakeTLEmojiGroupGreeting(nil)
		o.Data2.Constructor = -2133693241
		return o
	},
	154914612: func() TLObject { // 0x93bcf34
		o := MakeTLEmojiGroupPremium(nil)
		o.Data2.Constructor = 154914612
		return o
	},
	-709641735: func() TLObject { // 0xd5b3b9f9
		o := MakeTLEmojiKeyword(nil)
		o.Data2.Constructor = -709641735
		return o
	},
	594408994: func() TLObject { // 0x236df622
		o := MakeTLEmojiKeywordDeleted(nil)
		o.Data2.Constructor = 594408994
		return o
	},
	1556570557: func() TLObject { // 0x5cc761bd
		o := MakeTLEmojiKeywordsDifference(nil)
		o.Data2.Constructor = 1556570557
		return o
	},
	-1275374751: func() TLObject { // 0xb3fb5361
		o := MakeTLEmojiLanguage(nil)
		o.Data2.Constructor = -1275374751
		return o
	},
	1209970170: func() TLObject { // 0x481eadfa
		o := MakeTLEmojiListNotModified(nil)
		o.Data2.Constructor = 1209970170
		return o
	},
	2048790993: func() TLObject { // 0x7a1e11d1
		o := MakeTLEmojiList(nil)
		o.Data2.Constructor = 2048790993
		return o
	},
	769727150: func() TLObject { // 0x2de11aae
		o := MakeTLEmojiStatusEmpty(nil)
		o.Data2.Constructor = 769727150
		return o
	},
	-1835310691: func() TLObject { // 0x929b619d
		o := MakeTLEmojiStatus(nil)
		o.Data2.Constructor = -1835310691
		return o
	},
	-97474361: func() TLObject { // 0xfa30a8c7
		o := MakeTLEmojiStatusUntil(nil)
		o.Data2.Constructor = -97474361
		return o
	},
	-1519029347: func() TLObject { // 0xa575739d
		o := MakeTLEmojiURL(nil)
		o.Data2.Constructor = -1519029347
		return o
	},
	-1417756512: func() TLObject { // 0xab7ec0a0
		o := MakeTLEncryptedChatEmpty(nil)
		o.Data2.Constructor = -1417756512
		return o
	},
	1722964307: func() TLObject { // 0x66b25953
		o := MakeTLEncryptedChatWaiting(nil)
		o.Data2.Constructor = 1722964307
		return o
	},
	1223809356: func() TLObject { // 0x48f1d94c
		o := MakeTLEncryptedChatRequested(nil)
		o.Data2.Constructor = 1223809356
		return o
	},
	1643173063: func() TLObject { // 0x61f0d4c7
		o := MakeTLEncryptedChat(nil)
		o.Data2.Constructor = 1643173063
		return o
	},
	505183301: func() TLObject { // 0x1e1c7c45
		o := MakeTLEncryptedChatDiscarded(nil)
		o.Data2.Constructor = 505183301
		return o
	},
	-1038136962: func() TLObject { // 0xc21f497e
		o := MakeTLEncryptedFileEmpty(nil)
		o.Data2.Constructor = -1038136962
		return o
	},
	-1476358952: func() TLObject { // 0xa8008cd8
		o := MakeTLEncryptedFile(nil)
		o.Data2.Constructor = -1476358952
		return o
	},
	1248893260: func() TLObject { // 0x4a70994c
		o := MakeTLEncryptedFile(nil)
		o.Data2.Constructor = 1248893260
		return o
	},
	-317144808: func() TLObject { // 0xed18c118
		o := MakeTLEncryptedMessage(nil)
		o.Data2.Constructor = -317144808
		return o
	},
	594758406: func() TLObject { // 0x23734b06
		o := MakeTLEncryptedMessageService(nil)
		o.Data2.Constructor = 594758406
		return o
	},
	-994444869: func() TLObject { // 0xc4b9f9bb
		o := MakeTLError(nil)
		o.Data2.Constructor = -994444869
		return o
	},
	179611673: func() TLObject { // 0xab4a819
		o := MakeTLChatInviteExported(nil)
		o.Data2.Constructor = 179611673
		return o
	},
	-317687113: func() TLObject { // 0xed107ab7
		o := MakeTLChatInvitePublicJoinRequests(nil)
		o.Data2.Constructor = -317687113
		return o
	},
	206668204: func() TLObject { // 0xc5181ac
		o := MakeTLExportedChatlistInvite(nil)
		o.Data2.Constructor = 206668204
		return o
	},
	1103040667: func() TLObject { // 0x41bf109b
		o := MakeTLExportedContactToken(nil)
		o.Data2.Constructor = 1103040667
		return o
	},
	1571494644: func() TLObject { // 0x5dab1af4
		o := MakeTLExportedMessageLink(nil)
		o.Data2.Constructor = 1571494644
		return o
	},
	1070138683: func() TLObject { // 0x3fc9053b
		o := MakeTLExportedStoryLink(nil)
		o.Data2.Constructor = 1070138683
		return o
	},
	-1197736753: func() TLObject { // 0xb89bfccf
		o := MakeTLFactCheck(nil)
		o.Data2.Constructor = -1197736753
		return o
	},
	-207944868: func() TLObject { // 0xf39b035c
		o := MakeTLFileHash(nil)
		o.Data2.Constructor = -207944868
		return o
	},
	1648543603: func() TLObject { // 0x6242c773
		o := MakeTLFileHash(nil)
		o.Data2.Constructor = 1648543603
		return o
	},
	-11252123: func() TLObject { // 0xff544e65
		o := MakeTLFolder(nil)
		o.Data2.Constructor = -11252123
		return o
	},
	-373643672: func() TLObject { // 0xe9baa668
		o := MakeTLFolderPeer(nil)
		o.Data2.Constructor = -373643672
		return o
	},
	37687451: func() TLObject { // 0x23f109b
		o := MakeTLForumTopicDeleted(nil)
		o.Data2.Constructor = 37687451
		return o
	},
	1903173033: func() TLObject { // 0x71701da9
		o := MakeTLForumTopic(nil)
		o.Data2.Constructor = 1903173033
		return o
	},
	-394605632: func() TLObject { // 0xe87acbc0
		o := MakeTLFoundStory(nil)
		o.Data2.Constructor = -394605632
		return o
	},
	1857945489: func() TLObject { // 0x6ebdff91
		o := MakeTLFragmentCollectibleInfo(nil)
		o.Data2.Constructor = 1857945489
		return o
	},
	-1107729093: func() TLObject { // 0xbdf9653b
		o := MakeTLGame(nil)
		o.Data2.Constructor = -1107729093
		return o
	},
	286776671: func() TLObject { // 0x1117dd5f
		o := MakeTLGeoPointEmpty(nil)
		o.Data2.Constructor = 286776671
		return o
	},
	-1297942941: func() TLObject { // 0xb2a2f663
		o := MakeTLGeoPoint(nil)
		o.Data2.Constructor = -1297942941
		return o
	},
	-565420653: func() TLObject { // 0xde4c5d93
		o := MakeTLGeoPointAddress(nil)
		o.Data2.Constructor = -565420653
		return o
	},
	1934380235: func() TLObject { // 0x734c4ccb
		o := MakeTLGlobalPrivacySettings(nil)
		o.Data2.Constructor = 1934380235
		return o
	},
	-1096616924: func() TLObject { // 0xbea2f424
		o := MakeTLGlobalPrivacySettings(nil)
		o.Data2.Constructor = -1096616924
		return o
	},
	2004925620: func() TLObject { // 0x7780bcb4
		o := MakeTLGroupCallDiscarded(nil)
		o.Data2.Constructor = 2004925620
		return o
	},
	-711498484: func() TLObject { // 0xd597650c
		o := MakeTLGroupCall(nil)
		o.Data2.Constructor = -711498484
		return o
	},
	-341428482: func() TLObject { // 0xeba636fe
		o := MakeTLGroupCallParticipant(nil)
		o.Data2.Constructor = -341428482
		return o
	},
	1735736008: func() TLObject { // 0x67753ac8
		o := MakeTLGroupCallParticipantVideo(nil)
		o.Data2.Constructor = 1735736008
		return o
	},
	-592373577: func() TLObject { // 0xdcb118b7
		o := MakeTLGroupCallParticipantVideoSourceGroup(nil)
		o.Data2.Constructor = -592373577
		return o
	},
	-2132064081: func() TLObject { // 0x80eb48af
		o := MakeTLGroupCallStreamChannel(nil)
		o.Data2.Constructor = -2132064081
		return o
	},
	2094949405: func() TLObject { // 0x7cde641d
		o := MakeTLHelpAppConfigNotModified(nil)
		o.Data2.Constructor = 2094949405
		return o
	},
	-585598930: func() TLObject { // 0xdd18782e
		o := MakeTLHelpAppConfig(nil)
		o.Data2.Constructor = -585598930
		return o
	},
	-860107216: func() TLObject { // 0xccbbce30
		o := MakeTLHelpAppUpdate(nil)
		o.Data2.Constructor = -860107216
		return o
	},
	-1000708810: func() TLObject { // 0xc45a6536
		o := MakeTLHelpNoAppUpdate(nil)
		o.Data2.Constructor = -1000708810
		return o
	},
	-1815339214: func() TLObject { // 0x93cc1f32
		o := MakeTLHelpCountriesListNotModified(nil)
		o.Data2.Constructor = -1815339214
		return o
	},
	-2016381538: func() TLObject { // 0x87d0759e
		o := MakeTLHelpCountriesList(nil)
		o.Data2.Constructor = -2016381538
		return o
	},
	-1014526429: func() TLObject { // 0xc3878e23
		o := MakeTLHelpCountry(nil)
		o.Data2.Constructor = -1014526429
		return o
	},
	1107543535: func() TLObject { // 0x4203c5ef
		o := MakeTLHelpCountryCode(nil)
		o.Data2.Constructor = 1107543535
		return o
	},
	1722786150: func() TLObject { // 0x66afa166
		o := MakeTLHelpDeepLinkInfoEmpty(nil)
		o.Data2.Constructor = 1722786150
		return o
	},
	1783556146: func() TLObject { // 0x6a4ee832
		o := MakeTLHelpDeepLinkInfo(nil)
		o.Data2.Constructor = 1783556146
		return o
	},
	415997816: func() TLObject { // 0x18cb9f78
		o := MakeTLHelpInviteText(nil)
		o.Data2.Constructor = 415997816
		return o
	},
	-1078332329: func() TLObject { // 0xbfb9f457
		o := MakeTLHelpPassportConfigNotModified(nil)
		o.Data2.Constructor = -1078332329
		return o
	},
	-1600596305: func() TLObject { // 0xa098d6af
		o := MakeTLHelpPassportConfig(nil)
		o.Data2.Constructor = -1600596305
		return o
	},
	-1377014082: func() TLObject { // 0xadec6ebe
		o := MakeTLHelpPeerColorOption(nil)
		o.Data2.Constructor = -1377014082
		return o
	},
	-276549461: func() TLObject { // 0xef8430ab
		o := MakeTLHelpPeerColorOption(nil)
		o.Data2.Constructor = -276549461
		return o
	},
	324785199: func() TLObject { // 0x135bd42f
		o := MakeTLHelpPeerColorOption(nil)
		o.Data2.Constructor = 324785199
		return o
	},
	639736408: func() TLObject { // 0x26219a58
		o := MakeTLHelpPeerColorSet(nil)
		o.Data2.Constructor = 639736408
		return o
	},
	1987928555: func() TLObject { // 0x767d61eb
		o := MakeTLHelpPeerColorProfileSet(nil)
		o.Data2.Constructor = 1987928555
		return o
	},
	732034510: func() TLObject { // 0x2ba1f5ce
		o := MakeTLHelpPeerColorsNotModified(nil)
		o.Data2.Constructor = 732034510
		return o
	},
	16313608: func() TLObject { // 0xf8ed08
		o := MakeTLHelpPeerColors(nil)
		o.Data2.Constructor = 16313608
		return o
	},
	1395946908: func() TLObject { // 0x5334759c
		o := MakeTLHelpPremiumPromo(nil)
		o.Data2.Constructor = 1395946908
		return o
	},
	-1974518743: func() TLObject { // 0x8a4f3c29
		o := MakeTLHelpPremiumPromo(nil)
		o.Data2.Constructor = -1974518743
		return o
	},
	-1728664459: func() TLObject { // 0x98f6ac75
		o := MakeTLHelpPromoDataEmpty(nil)
		o.Data2.Constructor = -1728664459
		return o
	},
	-1942390465: func() TLObject { // 0x8c39793f
		o := MakeTLHelpPromoData(nil)
		o.Data2.Constructor = -1942390465
		return o
	},
	235081943: func() TLObject { // 0xe0310d7
		o := MakeTLHelpRecentMeUrls(nil)
		o.Data2.Constructor = 235081943
		return o
	},
	398898678: func() TLObject { // 0x17c6b5f6
		o := MakeTLHelpSupport(nil)
		o.Data2.Constructor = 398898678
		return o
	},
	-1945767479: func() TLObject { // 0x8c05f1c9
		o := MakeTLHelpSupportName(nil)
		o.Data2.Constructor = -1945767479
		return o
	},
	2013922064: func() TLObject { // 0x780a0310
		o := MakeTLHelpTermsOfService(nil)
		o.Data2.Constructor = 2013922064
		return o
	},
	-483352705: func() TLObject { // 0xe3309f7f
		o := MakeTLHelpTermsOfServiceUpdateEmpty(nil)
		o.Data2.Constructor = -483352705
		return o
	},
	686618977: func() TLObject { // 0x28ecf961
		o := MakeTLHelpTermsOfServiceUpdate(nil)
		o.Data2.Constructor = 686618977
		return o
	},
	-1761146676: func() TLObject { // 0x970708cc
		o := MakeTLHelpTimezonesListNotModified(nil)
		o.Data2.Constructor = -1761146676
		return o
	},
	2071260529: func() TLObject { // 0x7b74ed71
		o := MakeTLHelpTimezonesList(nil)
		o.Data2.Constructor = 2071260529
		return o
	},
	-206688531: func() TLObject { // 0xf3ae2eed
		o := MakeTLHelpUserInfoEmpty(nil)
		o.Data2.Constructor = -206688531
		return o
	},
	32192344: func() TLObject { // 0x1eb3758
		o := MakeTLHelpUserInfo(nil)
		o.Data2.Constructor = 32192344
		return o
	},
	1940093419: func() TLObject { // 0x73a379eb
		o := MakeTLHighScore(nil)
		o.Data2.Constructor = 1940093419
		return o
	},
	-1557334680: func() TLObject { // 0xa32cf568
		o := MakeTLImmutableChat(nil)
		o.Data2.Constructor = -1557334680
		return o
	},
	-100771298: func() TLObject { // 0xf9fe5a1e
		o := MakeTLImmutableChatParticipant(nil)
		o.Data2.Constructor = -100771298
		return o
	},
	972235212: func() TLObject { // 0x39f321cc
		o := MakeTLImmutableUser(nil)
		o.Data2.Constructor = 972235212
		return o
	},
	-1052885936: func() TLObject { // 0xc13e3c50
		o := MakeTLImportedContact(nil)
		o.Data2.Constructor = -1052885936
		return o
	},
	1008755359: func() TLObject { // 0x3c20629f
		o := MakeTLInlineBotSwitchPM(nil)
		o.Data2.Constructor = 1008755359
		return o
	},
	-1250781739: func() TLObject { // 0xb57295d5
		o := MakeTLInlineBotWebView(nil)
		o.Data2.Constructor = -1250781739
		return o
	},
	813821341: func() TLObject { // 0x3081ed9d
		o := MakeTLInlineQueryPeerTypeSameBotPM(nil)
		o.Data2.Constructor = 813821341
		return o
	},
	-2093215828: func() TLObject { // 0x833c0fac
		o := MakeTLInlineQueryPeerTypePM(nil)
		o.Data2.Constructor = -2093215828
		return o
	},
	-681130742: func() TLObject { // 0xd766c50a
		o := MakeTLInlineQueryPeerTypeChat(nil)
		o.Data2.Constructor = -681130742
		return o
	},
	1589952067: func() TLObject { // 0x5ec4be43
		o := MakeTLInlineQueryPeerTypeMegagroup(nil)
		o.Data2.Constructor = 1589952067
		return o
	},
	1664413338: func() TLObject { // 0x6334ee9a
		o := MakeTLInlineQueryPeerTypeBroadcast(nil)
		o.Data2.Constructor = 1664413338
		return o
	},
	238759180: func() TLObject { // 0xe3b2d0c
		o := MakeTLInlineQueryPeerTypeBotPM(nil)
		o.Data2.Constructor = 238759180
		return o
	},
	488313413: func() TLObject { // 0x1d1b1245
		o := MakeTLInputAppEvent(nil)
		o.Data2.Constructor = 488313413
		return o
	},
	-1457472134: func() TLObject { // 0xa920bd7a
		o := MakeTLInputBotAppID(nil)
		o.Data2.Constructor = -1457472134
		return o
	},
	-1869872121: func() TLObject { // 0x908c0407
		o := MakeTLInputBotAppShortName(nil)
		o.Data2.Constructor = -1869872121
		return o
	},
	864077702: func() TLObject { // 0x3380c786
		o := MakeTLInputBotInlineMessageMediaAuto(nil)
		o.Data2.Constructor = 864077702
		return o
	},
	1036876423: func() TLObject { // 0x3dcd7a87
		o := MakeTLInputBotInlineMessageText(nil)
		o.Data2.Constructor = 1036876423
		return o
	},
	-1768777083: func() TLObject { // 0x96929a85
		o := MakeTLInputBotInlineMessageMediaGeo(nil)
		o.Data2.Constructor = -1768777083
		return o
	},
	1098628881: func() TLObject { // 0x417bbf11
		o := MakeTLInputBotInlineMessageMediaVenue(nil)
		o.Data2.Constructor = 1098628881
		return o
	},
	-1494368259: func() TLObject { // 0xa6edbffd
		o := MakeTLInputBotInlineMessageMediaContact(nil)
		o.Data2.Constructor = -1494368259
		return o
	},
	1262639204: func() TLObject { // 0x4b425864
		o := MakeTLInputBotInlineMessageGame(nil)
		o.Data2.Constructor = 1262639204
		return o
	},
	-672693723: func() TLObject { // 0xd7e78225
		o := MakeTLInputBotInlineMessageMediaInvoice(nil)
		o.Data2.Constructor = -672693723
		return o
	},
	-1109605104: func() TLObject { // 0xbddcc510
		o := MakeTLInputBotInlineMessageMediaWebPage(nil)
		o.Data2.Constructor = -1109605104
		return o
	},
	-1995686519: func() TLObject { // 0x890c3d89
		o := MakeTLInputBotInlineMessageID(nil)
		o.Data2.Constructor = -1995686519
		return o
	},
	-1227287081: func() TLObject { // 0xb6d915d7
		o := MakeTLInputBotInlineMessageID64(nil)
		o.Data2.Constructor = -1227287081
		return o
	},
	-2000710887: func() TLObject { // 0x88bf9319
		o := MakeTLInputBotInlineResult(nil)
		o.Data2.Constructor = -2000710887
		return o
	},
	-1462213465: func() TLObject { // 0xa8d864a7
		o := MakeTLInputBotInlineResultPhoto(nil)
		o.Data2.Constructor = -1462213465
		return o
	},
	-459324: func() TLObject { // 0xfff8fdc4
		o := MakeTLInputBotInlineResultDocument(nil)
		o.Data2.Constructor = -459324
		return o
	},
	1336154098: func() TLObject { // 0x4fa417f2
		o := MakeTLInputBotInlineResultGame(nil)
		o.Data2.Constructor = 1336154098
		return o
	},
	-2094959136: func() TLObject { // 0x832175e0
		o := MakeTLInputBusinessAwayMessage(nil)
		o.Data2.Constructor = -2094959136
		return o
	},
	-991587810: func() TLObject { // 0xc4e5921e
		o := MakeTLInputBusinessBotRecipients(nil)
		o.Data2.Constructor = -991587810
		return o
	},
	292003751: func() TLObject { // 0x11679fa7
		o := MakeTLInputBusinessChatLink(nil)
		o.Data2.Constructor = 292003751
		return o
	},
	26528571: func() TLObject { // 0x194cb3b
		o := MakeTLInputBusinessGreetingMessage(nil)
		o.Data2.Constructor = 26528571
		return o
	},
	163867085: func() TLObject { // 0x9c469cd
		o := MakeTLInputBusinessIntro(nil)
		o.Data2.Constructor = 163867085
		return o
	},
	1871393450: func() TLObject { // 0x6f8b32aa
		o := MakeTLInputBusinessRecipients(nil)
		o.Data2.Constructor = 1871393450
		return o
	},
	-292807034: func() TLObject { // 0xee8c1e86
		o := MakeTLInputChannelEmpty(nil)
		o.Data2.Constructor = -292807034
		return o
	},
	-212145112: func() TLObject { // 0xf35aec28
		o := MakeTLInputChannel(nil)
		o.Data2.Constructor = -212145112
		return o
	},
	1536380829: func() TLObject { // 0x5b934f9d
		o := MakeTLInputChannelFromMessage(nil)
		o.Data2.Constructor = 1536380829
		return o
	},
	480546647: func() TLObject { // 0x1ca48f57
		o := MakeTLInputChatPhotoEmpty(nil)
		o.Data2.Constructor = 480546647
		return o
	},
	-1110593856: func() TLObject { // 0xbdcdaec0
		o := MakeTLInputChatUploadedPhoto(nil)
		o.Data2.Constructor = -1110593856
		return o
	},
	-968723890: func() TLObject { // 0xc642724e
		o := MakeTLInputChatUploadedPhoto(nil)
		o.Data2.Constructor = -968723890
		return o
	},
	-1991004873: func() TLObject { // 0x8953ad37
		o := MakeTLInputChatPhoto(nil)
		o.Data2.Constructor = -1991004873
		return o
	},
	-203367885: func() TLObject { // 0xf3e0da33
		o := MakeTLInputChatlistDialogFilter(nil)
		o.Data2.Constructor = -203367885
		return o
	},
	-1736378792: func() TLObject { // 0x9880f658
		o := MakeTLInputCheckPasswordEmpty(nil)
		o.Data2.Constructor = -1736378792
		return o
	},
	-763367294: func() TLObject { // 0xd27ff082
		o := MakeTLInputCheckPasswordSRP(nil)
		o.Data2.Constructor = -763367294
		return o
	},
	1968737087: func() TLObject { // 0x75588b3f
		o := MakeTLInputClientProxy(nil)
		o.Data2.Constructor = 1968737087
		return o
	},
	-476815191: func() TLObject { // 0xe39460a9
		o := MakeTLInputCollectibleUsername(nil)
		o.Data2.Constructor = -476815191
		return o
	},
	-1562241884: func() TLObject { // 0xa2e214a4
		o := MakeTLInputCollectiblePhone(nil)
		o.Data2.Constructor = -1562241884
		return o
	},
	-208488460: func() TLObject { // 0xf392b7f4
		o := MakeTLInputPhoneContact(nil)
		o.Data2.Constructor = -208488460
		return o
	},
	-55902537: func() TLObject { // 0xfcaafeb7
		o := MakeTLInputDialogPeer(nil)
		o.Data2.Constructor = -55902537
		return o
	},
	1684014375: func() TLObject { // 0x64600527
		o := MakeTLInputDialogPeerFolder(nil)
		o.Data2.Constructor = 1684014375
		return o
	},
	1928391342: func() TLObject { // 0x72f0eaae
		o := MakeTLInputDocumentEmpty(nil)
		o.Data2.Constructor = 1928391342
		return o
	},
	448771445: func() TLObject { // 0x1abfb575
		o := MakeTLInputDocument(nil)
		o.Data2.Constructor = 448771445
		return o
	},
	-247351839: func() TLObject { // 0xf141b5e1
		o := MakeTLInputEncryptedChat(nil)
		o.Data2.Constructor = -247351839
		return o
	},
	406307684: func() TLObject { // 0x1837c364
		o := MakeTLInputEncryptedFileEmpty(nil)
		o.Data2.Constructor = 406307684
		return o
	},
	1690108678: func() TLObject { // 0x64bd0306
		o := MakeTLInputEncryptedFileUploaded(nil)
		o.Data2.Constructor = 1690108678
		return o
	},
	1511503333: func() TLObject { // 0x5a17b5e5
		o := MakeTLInputEncryptedFile(nil)
		o.Data2.Constructor = 1511503333
		return o
	},
	767652808: func() TLObject { // 0x2dc173c8
		o := MakeTLInputEncryptedFileBigUploaded(nil)
		o.Data2.Constructor = 767652808
		return o
	},
	-181407105: func() TLObject { // 0xf52ff27f
		o := MakeTLInputFile(nil)
		o.Data2.Constructor = -181407105
		return o
	},
	-95482955: func() TLObject { // 0xfa4f0bb5
		o := MakeTLInputFileBig(nil)
		o.Data2.Constructor = -95482955
		return o
	},
	-539317279: func() TLObject { // 0xdfdaabe1
		o := MakeTLInputFileLocation(nil)
		o.Data2.Constructor = -539317279
		return o
	},
	-182231723: func() TLObject { // 0xf5235d55
		o := MakeTLInputEncryptedFileLocation(nil)
		o.Data2.Constructor = -182231723
		return o
	},
	-1160743548: func() TLObject { // 0xbad07584
		o := MakeTLInputDocumentFileLocation(nil)
		o.Data2.Constructor = -1160743548
		return o
	},
	-876089816: func() TLObject { // 0xcbc7ee28
		o := MakeTLInputSecureFileLocation(nil)
		o.Data2.Constructor = -876089816
		return o
	},
	700340377: func() TLObject { // 0x29be5899
		o := MakeTLInputTakeoutFileLocation(nil)
		o.Data2.Constructor = 700340377
		return o
	},
	1075322878: func() TLObject { // 0x40181ffe
		o := MakeTLInputPhotoFileLocation(nil)
		o.Data2.Constructor = 1075322878
		return o
	},
	-667654413: func() TLObject { // 0xd83466f3
		o := MakeTLInputPhotoLegacyFileLocation(nil)
		o.Data2.Constructor = -667654413
		return o
	},
	925204121: func() TLObject { // 0x37257e99
		o := MakeTLInputPeerPhotoFileLocation(nil)
		o.Data2.Constructor = 925204121
		return o
	},
	-1652231205: func() TLObject { // 0x9d84f3db
		o := MakeTLInputStickerSetThumb(nil)
		o.Data2.Constructor = -1652231205
		return o
	},
	93890858: func() TLObject { // 0x598a92a
		o := MakeTLInputGroupCallStream(nil)
		o.Data2.Constructor = 93890858
		return o
	},
	-70073706: func() TLObject { // 0xfbd2c296
		o := MakeTLInputFolderPeer(nil)
		o.Data2.Constructor = -70073706
		return o
	},
	53231223: func() TLObject { // 0x32c3e77
		o := MakeTLInputGameID(nil)
		o.Data2.Constructor = 53231223
		return o
	},
	-1020139510: func() TLObject { // 0xc331e80a
		o := MakeTLInputGameShortName(nil)
		o.Data2.Constructor = -1020139510
		return o
	},
	-457104426: func() TLObject { // 0xe4c123d6
		o := MakeTLInputGeoPointEmpty(nil)
		o.Data2.Constructor = -457104426
		return o
	},
	1210199983: func() TLObject { // 0x48222faf
		o := MakeTLInputGeoPoint(nil)
		o.Data2.Constructor = 1210199983
		return o
	},
	-659913713: func() TLObject { // 0xd8aa840f
		o := MakeTLInputGroupCall(nil)
		o.Data2.Constructor = -659913713
		return o
	},
	-977967015: func() TLObject { // 0xc5b56859
		o := MakeTLInputInvoiceMessage(nil)
		o.Data2.Constructor = -977967015
		return o
	},
	-1020867857: func() TLObject { // 0xc326caef
		o := MakeTLInputInvoiceSlug(nil)
		o.Data2.Constructor = -1020867857
		return o
	},
	-1734841331: func() TLObject { // 0x98986c0d
		o := MakeTLInputInvoicePremiumGiftCode(nil)
		o.Data2.Constructor = -1734841331
		return o
	},
	497236696: func() TLObject { // 0x1da33ad8
		o := MakeTLInputInvoiceStars(nil)
		o.Data2.Constructor = 497236696
		return o
	},
	-1771768449: func() TLObject { // 0x9664f57f
		o := MakeTLInputMediaEmpty(nil)
		o.Data2.Constructor = -1771768449
		return o
	},
	505969924: func() TLObject { // 0x1e287d04
		o := MakeTLInputMediaUploadedPhoto(nil)
		o.Data2.Constructor = 505969924
		return o
	},
	-1279654347: func() TLObject { // 0xb3ba0635
		o := MakeTLInputMediaPhoto(nil)
		o.Data2.Constructor = -1279654347
		return o
	},
	-104578748: func() TLObject { // 0xf9c44144
		o := MakeTLInputMediaGeoPoint(nil)
		o.Data2.Constructor = -104578748
		return o
	},
	-122978821: func() TLObject { // 0xf8ab7dfb
		o := MakeTLInputMediaContact(nil)
		o.Data2.Constructor = -122978821
		return o
	},
	1530447553: func() TLObject { // 0x5b38c6c1
		o := MakeTLInputMediaUploadedDocument(nil)
		o.Data2.Constructor = 1530447553
		return o
	},
	860303448: func() TLObject { // 0x33473058
		o := MakeTLInputMediaDocument(nil)
		o.Data2.Constructor = 860303448
		return o
	},
	-1052959727: func() TLObject { // 0xc13d1c11
		o := MakeTLInputMediaVenue(nil)
		o.Data2.Constructor = -1052959727
		return o
	},
	-440664550: func() TLObject { // 0xe5bbfe1a
		o := MakeTLInputMediaPhotoExternal(nil)
		o.Data2.Constructor = -440664550
		return o
	},
	-78455655: func() TLObject { // 0xfb52dc99
		o := MakeTLInputMediaDocumentExternal(nil)
		o.Data2.Constructor = -78455655
		return o
	},
	-750828557: func() TLObject { // 0xd33f43f3
		o := MakeTLInputMediaGame(nil)
		o.Data2.Constructor = -750828557
		return o
	},
	1080028941: func() TLObject { // 0x405fef0d
		o := MakeTLInputMediaInvoice(nil)
		o.Data2.Constructor = 1080028941
		return o
	},
	-1900697899: func() TLObject { // 0x8eb5a6d5
		o := MakeTLInputMediaInvoice(nil)
		o.Data2.Constructor = -1900697899
		return o
	},
	-646342540: func() TLObject { // 0xd9799874
		o := MakeTLInputMediaInvoice(nil)
		o.Data2.Constructor = -646342540
		return o
	},
	-1759532989: func() TLObject { // 0x971fa843
		o := MakeTLInputMediaGeoLive(nil)
		o.Data2.Constructor = -1759532989
		return o
	},
	261416433: func() TLObject { // 0xf94e5f1
		o := MakeTLInputMediaPoll(nil)
		o.Data2.Constructor = 261416433
		return o
	},
	-428884101: func() TLObject { // 0xe66fbf7b
		o := MakeTLInputMediaDice(nil)
		o.Data2.Constructor = -428884101
		return o
	},
	-1979852936: func() TLObject { // 0x89fdd778
		o := MakeTLInputMediaStory(nil)
		o.Data2.Constructor = -1979852936
		return o
	},
	-1702447729: func() TLObject { // 0x9a86b58f
		o := MakeTLInputMediaStory(nil)
		o.Data2.Constructor = -1702447729
		return o
	},
	-1038383031: func() TLObject { // 0xc21b8849
		o := MakeTLInputMediaWebPage(nil)
		o.Data2.Constructor = -1038383031
		return o
	},
	-1436147773: func() TLObject { // 0xaa661fc3
		o := MakeTLInputMediaPaidMedia(nil)
		o.Data2.Constructor = -1436147773
		return o
	},
	-1097470438: func() TLObject { // 0xbe95ee1a
		o := MakeTLInputMediaBizDataRaw(nil)
		o.Data2.Constructor = -1097470438
		return o
	},
	-1502174430: func() TLObject { // 0xa676a322
		o := MakeTLInputMessageID(nil)
		o.Data2.Constructor = -1502174430
		return o
	},
	-1160215659: func() TLObject { // 0xbad88395
		o := MakeTLInputMessageReplyTo(nil)
		o.Data2.Constructor = -1160215659
		return o
	},
	-2037963464: func() TLObject { // 0x86872538
		o := MakeTLInputMessagePinned(nil)
		o.Data2.Constructor = -2037963464
		return o
	},
	-1392895362: func() TLObject { // 0xacfa1a7e
		o := MakeTLInputMessageCallbackQuery(nil)
		o.Data2.Constructor = -1392895362
		return o
	},
	-1195615476: func() TLObject { // 0xb8bc5b0c
		o := MakeTLInputNotifyPeer(nil)
		o.Data2.Constructor = -1195615476
		return o
	},
	423314455: func() TLObject { // 0x193b4417
		o := MakeTLInputNotifyUsers(nil)
		o.Data2.Constructor = 423314455
		return o
	},
	1251338318: func() TLObject { // 0x4a95e84e
		o := MakeTLInputNotifyChats(nil)
		o.Data2.Constructor = 1251338318
		return o
	},
	-1311015810: func() TLObject { // 0xb1db7c7e
		o := MakeTLInputNotifyBroadcasts(nil)
		o.Data2.Constructor = -1311015810
		return o
	},
	1548122514: func() TLObject { // 0x5c467992
		o := MakeTLInputNotifyForumTopic(nil)
		o.Data2.Constructor = 1548122514
		return o
	},
	-1056001329: func() TLObject { // 0xc10eb2cf
		o := MakeTLInputPaymentCredentialsSaved(nil)
		o.Data2.Constructor = -1056001329
		return o
	},
	873977640: func() TLObject { // 0x3417d728
		o := MakeTLInputPaymentCredentials(nil)
		o.Data2.Constructor = 873977640
		return o
	},
	178373535: func() TLObject { // 0xaa1c39f
		o := MakeTLInputPaymentCredentialsApplePay(nil)
		o.Data2.Constructor = 178373535
		return o
	},
	-1966921727: func() TLObject { // 0x8ac32801
		o := MakeTLInputPaymentCredentialsGooglePay(nil)
		o.Data2.Constructor = -1966921727
		return o
	},
	2134579434: func() TLObject { // 0x7f3b18ea
		o := MakeTLInputPeerEmpty(nil)
		o.Data2.Constructor = 2134579434
		return o
	},
	2107670217: func() TLObject { // 0x7da07ec9
		o := MakeTLInputPeerSelf(nil)
		o.Data2.Constructor = 2107670217
		return o
	},
	900291769: func() TLObject { // 0x35a95cb9
		o := MakeTLInputPeerChat(nil)
		o.Data2.Constructor = 900291769
		return o
	},
	-571955892: func() TLObject { // 0xdde8a54c
		o := MakeTLInputPeerUser(nil)
		o.Data2.Constructor = -571955892
		return o
	},
	666680316: func() TLObject { // 0x27bcbbfc
		o := MakeTLInputPeerChannel(nil)
		o.Data2.Constructor = 666680316
		return o
	},
	-1468331492: func() TLObject { // 0xa87b0a1c
		o := MakeTLInputPeerUserFromMessage(nil)
		o.Data2.Constructor = -1468331492
		return o
	},
	-1121318848: func() TLObject { // 0xbd2a0840
		o := MakeTLInputPeerChannelFromMessage(nil)
		o.Data2.Constructor = -1121318848
		return o
	},
	-88014124: func() TLObject { // 0xfac102d4
		o := MakeTLInputPeerUsername(nil)
		o.Data2.Constructor = -88014124
		return o
	},
	-892638494: func() TLObject { // 0xcacb6ae2
		o := MakeTLInputPeerNotifySettings(nil)
		o.Data2.Constructor = -892638494
		return o
	},
	-551616469: func() TLObject { // 0xdf1f002b
		o := MakeTLInputPeerNotifySettings(nil)
		o.Data2.Constructor = -551616469
		return o
	},
	-1673717362: func() TLObject { // 0x9c3d198e
		o := MakeTLInputPeerNotifySettings(nil)
		o.Data2.Constructor = -1673717362
		return o
	},
	506920429: func() TLObject { // 0x1e36fded
		o := MakeTLInputPhoneCall(nil)
		o.Data2.Constructor = 506920429
		return o
	},
	483901197: func() TLObject { // 0x1cd7bf0d
		o := MakeTLInputPhotoEmpty(nil)
		o.Data2.Constructor = 483901197
		return o
	},
	1001634122: func() TLObject { // 0x3bb3b94a
		o := MakeTLInputPhoto(nil)
		o.Data2.Constructor = 1001634122
		return o
	},
	1335282456: func() TLObject { // 0x4f96cb18
		o := MakeTLInputPrivacyKeyStatusTimestamp(nil)
		o.Data2.Constructor = 1335282456
		return o
	},
	-1107622874: func() TLObject { // 0xbdfb0426
		o := MakeTLInputPrivacyKeyChatInvite(nil)
		o.Data2.Constructor = -1107622874
		return o
	},
	-88417185: func() TLObject { // 0xfabadc5f
		o := MakeTLInputPrivacyKeyPhoneCall(nil)
		o.Data2.Constructor = -88417185
		return o
	},
	-610373422: func() TLObject { // 0xdb9e70d2
		o := MakeTLInputPrivacyKeyPhoneP2P(nil)
		o.Data2.Constructor = -610373422
		return o
	},
	-1529000952: func() TLObject { // 0xa4dd4c08
		o := MakeTLInputPrivacyKeyForwards(nil)
		o.Data2.Constructor = -1529000952
		return o
	},
	1461304012: func() TLObject { // 0x5719bacc
		o := MakeTLInputPrivacyKeyProfilePhoto(nil)
		o.Data2.Constructor = 1461304012
		return o
	},
	55761658: func() TLObject { // 0x352dafa
		o := MakeTLInputPrivacyKeyPhoneNumber(nil)
		o.Data2.Constructor = 55761658
		return o
	},
	-786326563: func() TLObject { // 0xd1219bdd
		o := MakeTLInputPrivacyKeyAddedByPhone(nil)
		o.Data2.Constructor = -786326563
		return o
	},
	-1360618136: func() TLObject { // 0xaee69d68
		o := MakeTLInputPrivacyKeyVoiceMessages(nil)
		o.Data2.Constructor = -1360618136
		return o
	},
	941870144: func() TLObject { // 0x3823cc40
		o := MakeTLInputPrivacyKeyAbout(nil)
		o.Data2.Constructor = 941870144
		return o
	},
	-698740276: func() TLObject { // 0xd65a11cc
		o := MakeTLInputPrivacyKeyBirthday(nil)
		o.Data2.Constructor = -698740276
		return o
	},
	218751099: func() TLObject { // 0xd09e07b
		o := MakeTLInputPrivacyValueAllowContacts(nil)
		o.Data2.Constructor = 218751099
		return o
	},
	407582158: func() TLObject { // 0x184b35ce
		o := MakeTLInputPrivacyValueAllowAll(nil)
		o.Data2.Constructor = 407582158
		return o
	},
	320652927: func() TLObject { // 0x131cc67f
		o := MakeTLInputPrivacyValueAllowUsers(nil)
		o.Data2.Constructor = 320652927
		return o
	},
	195371015: func() TLObject { // 0xba52007
		o := MakeTLInputPrivacyValueDisallowContacts(nil)
		o.Data2.Constructor = 195371015
		return o
	},
	-697604407: func() TLObject { // 0xd66b66c9
		o := MakeTLInputPrivacyValueDisallowAll(nil)
		o.Data2.Constructor = -697604407
		return o
	},
	-1877932953: func() TLObject { // 0x90110467
		o := MakeTLInputPrivacyValueDisallowUsers(nil)
		o.Data2.Constructor = -1877932953
		return o
	},
	-2079962673: func() TLObject { // 0x840649cf
		o := MakeTLInputPrivacyValueAllowChatParticipants(nil)
		o.Data2.Constructor = -2079962673
		return o
	},
	-380694650: func() TLObject { // 0xe94f0f86
		o := MakeTLInputPrivacyValueDisallowChatParticipants(nil)
		o.Data2.Constructor = -380694650
		return o
	},
	793067081: func() TLObject { // 0x2f453e49
		o := MakeTLInputPrivacyValueAllowCloseFriends(nil)
		o.Data2.Constructor = 793067081
		return o
	},
	2009975281: func() TLObject { // 0x77cdc9f1
		o := MakeTLInputPrivacyValueAllowPremium(nil)
		o.Data2.Constructor = 2009975281
		return o
	},
	609840449: func() TLObject { // 0x24596d41
		o := MakeTLInputQuickReplyShortcut(nil)
		o.Data2.Constructor = 609840449
		return o
	},
	18418929: func() TLObject { // 0x1190cf1
		o := MakeTLInputQuickReplyShortcutId(nil)
		o.Data2.Constructor = 18418929
		return o
	},
	583071445: func() TLObject { // 0x22c0f6d5
		o := MakeTLInputReplyToMessage(nil)
		o.Data2.Constructor = 583071445
		return o
	},
	121554949: func() TLObject { // 0x73ec805
		o := MakeTLInputReplyToMessage(nil)
		o.Data2.Constructor = 121554949
		return o
	},
	-1672247580: func() TLObject { // 0x9c5386e4
		o := MakeTLInputReplyToMessage(nil)
		o.Data2.Constructor = -1672247580
		return o
	},
	1484862010: func() TLObject { // 0x5881323a
		o := MakeTLInputReplyToStory(nil)
		o.Data2.Constructor = 1484862010
		return o
	},
	363917955: func() TLObject { // 0x15b0f283
		o := MakeTLInputReplyToStory(nil)
		o.Data2.Constructor = 363917955
		return o
	},
	859091184: func() TLObject { // 0x3334b0f0
		o := MakeTLInputSecureFileUploaded(nil)
		o.Data2.Constructor = 859091184
		return o
	},
	1399317950: func() TLObject { // 0x5367e5be
		o := MakeTLInputSecureFile(nil)
		o.Data2.Constructor = 1399317950
		return o
	},
	-618540889: func() TLObject { // 0xdb21d0a7
		o := MakeTLInputSecureValue(nil)
		o.Data2.Constructor = -618540889
		return o
	},
	482797855: func() TLObject { // 0x1cc6e91f
		o := MakeTLInputSingleMedia(nil)
		o.Data2.Constructor = 482797855
		return o
	},
	543876817: func() TLObject { // 0x206ae6d1
		o := MakeTLInputStarsTransaction(nil)
		o.Data2.Constructor = 543876817
		return o
	},
	-4838507: func() TLObject { // 0xffb62b95
		o := MakeTLInputStickerSetEmpty(nil)
		o.Data2.Constructor = -4838507
		return o
	},
	-1645763991: func() TLObject { // 0x9de7a269
		o := MakeTLInputStickerSetID(nil)
		o.Data2.Constructor = -1645763991
		return o
	},
	-2044933984: func() TLObject { // 0x861cc8a0
		o := MakeTLInputStickerSetShortName(nil)
		o.Data2.Constructor = -2044933984
		return o
	},
	42402760: func() TLObject { // 0x28703c8
		o := MakeTLInputStickerSetAnimatedEmoji(nil)
		o.Data2.Constructor = 42402760
		return o
	},
	-427863538: func() TLObject { // 0xe67f520e
		o := MakeTLInputStickerSetDice(nil)
		o.Data2.Constructor = -427863538
		return o
	},
	215889721: func() TLObject { // 0xcde3739
		o := MakeTLInputStickerSetAnimatedEmojiAnimations(nil)
		o.Data2.Constructor = 215889721
		return o
	},
	-930399486: func() TLObject { // 0xc88b3b02
		o := MakeTLInputStickerSetPremiumGifts(nil)
		o.Data2.Constructor = -930399486
		return o
	},
	80008398: func() TLObject { // 0x4c4d4ce
		o := MakeTLInputStickerSetEmojiGenericAnimations(nil)
		o.Data2.Constructor = 80008398
		return o
	},
	701560302: func() TLObject { // 0x29d0f5ee
		o := MakeTLInputStickerSetEmojiDefaultStatuses(nil)
		o.Data2.Constructor = 701560302
		return o
	},
	1153562857: func() TLObject { // 0x44c1f8e9
		o := MakeTLInputStickerSetEmojiDefaultTopicIcons(nil)
		o.Data2.Constructor = 1153562857
		return o
	},
	1232373075: func() TLObject { // 0x49748553
		o := MakeTLInputStickerSetEmojiChannelDefaultStatuses(nil)
		o.Data2.Constructor = 1232373075
		return o
	},
	853188252: func() TLObject { // 0x32da9e9c
		o := MakeTLInputStickerSetItem(nil)
		o.Data2.Constructor = 853188252
		return o
	},
	-6249322: func() TLObject { // 0xffa0a496
		o := MakeTLInputStickerSetItem(nil)
		o.Data2.Constructor = -6249322
		return o
	},
	1251549527: func() TLObject { // 0x4a992157
		o := MakeTLInputStickeredMediaPhoto(nil)
		o.Data2.Constructor = 1251549527
		return o
	},
	70813275: func() TLObject { // 0x438865b
		o := MakeTLInputStickeredMediaDocument(nil)
		o.Data2.Constructor = 70813275
		return o
	},
	-1502273946: func() TLObject { // 0xa6751e66
		o := MakeTLInputStorePaymentPremiumSubscription(nil)
		o.Data2.Constructor = -1502273946
		return o
	},
	1634697192: func() TLObject { // 0x616f7fe8
		o := MakeTLInputStorePaymentGiftPremium(nil)
		o.Data2.Constructor = 1634697192
		return o
	},
	-1551868097: func() TLObject { // 0xa3805f3f
		o := MakeTLInputStorePaymentPremiumGiftCode(nil)
		o.Data2.Constructor = -1551868097
		return o
	},
	369444042: func() TLObject { // 0x160544ca
		o := MakeTLInputStorePaymentPremiumGiveaway(nil)
		o.Data2.Constructor = 369444042
		return o
	},
	2090038758: func() TLObject { // 0x7c9375e6
		o := MakeTLInputStorePaymentPremiumGiveaway(nil)
		o.Data2.Constructor = 2090038758
		return o
	},
	1326377183: func() TLObject { // 0x4f0ee8df
		o := MakeTLInputStorePaymentStars(nil)
		o.Data2.Constructor = 1326377183
		return o
	},
	1012306921: func() TLObject { // 0x3c5693e9
		o := MakeTLInputTheme(nil)
		o.Data2.Constructor = 1012306921
		return o
	},
	-175567375: func() TLObject { // 0xf5890df1
		o := MakeTLInputThemeSlug(nil)
		o.Data2.Constructor = -175567375
		return o
	},
	-1881255857: func() TLObject { // 0x8fde504f
		o := MakeTLInputThemeSettings(nil)
		o.Data2.Constructor = -1881255857
		return o
	},
	-1182234929: func() TLObject { // 0xb98886cf
		o := MakeTLInputUserEmpty(nil)
		o.Data2.Constructor = -1182234929
		return o
	},
	-138301121: func() TLObject { // 0xf7c1b13f
		o := MakeTLInputUserSelf(nil)
		o.Data2.Constructor = -138301121
		return o
	},
	-233744186: func() TLObject { // 0xf21158c6
		o := MakeTLInputUser(nil)
		o.Data2.Constructor = -233744186
		return o
	},
	497305826: func() TLObject { // 0x1da448e2
		o := MakeTLInputUserFromMessage(nil)
		o.Data2.Constructor = 497305826
		return o
	},
	-433014407: func() TLObject { // 0xe630b979
		o := MakeTLInputWallPaper(nil)
		o.Data2.Constructor = -433014407
		return o
	},
	1913199744: func() TLObject { // 0x72091c80
		o := MakeTLInputWallPaperSlug(nil)
		o.Data2.Constructor = 1913199744
		return o
	},
	-1770371538: func() TLObject { // 0x967a462e
		o := MakeTLInputWallPaperNoFile(nil)
		o.Data2.Constructor = -1770371538
		return o
	},
	-1678949555: func() TLObject { // 0x9bed434d
		o := MakeTLInputWebDocument(nil)
		o.Data2.Constructor = -1678949555
		return o
	},
	-1036396922: func() TLObject { // 0xc239d686
		o := MakeTLInputWebFileLocation(nil)
		o.Data2.Constructor = -1036396922
		return o
	},
	-1625153079: func() TLObject { // 0x9f2221c9
		o := MakeTLInputWebFileGeoPointLocation(nil)
		o.Data2.Constructor = -1625153079
		return o
	},
	-193992412: func() TLObject { // 0xf46fe924
		o := MakeTLInputWebFileAudioAlbumThumbLocation(nil)
		o.Data2.Constructor = -193992412
		return o
	},
	-1932527041: func() TLObject { // 0x8ccffa3f
		o := MakeTLInt32(nil)
		o.Data2.Constructor = -1932527041
		return o
	},
	1253220205: func() TLObject { // 0x4ab29f6d
		o := MakeTLLong(nil)
		o.Data2.Constructor = 1253220205
		return o
	},
	-1568590240: func() TLObject { // 0xa2813660
		o := MakeTLInt64(nil)
		o.Data2.Constructor = -1568590240
		return o
	},
	1572428309: func() TLObject { // 0x5db95a15
		o := MakeTLInvoice(nil)
		o.Data2.Constructor = 1572428309
		return o
	},
	1048946971: func() TLObject { // 0x3e85a91b
		o := MakeTLInvoice(nil)
		o.Data2.Constructor = 1048946971
		return o
	},
	215516896: func() TLObject { // 0xcd886e0
		o := MakeTLInvoice(nil)
		o.Data2.Constructor = 215516896
		return o
	},
	-1059185703: func() TLObject { // 0xc0de1bd9
		o := MakeTLJsonObjectValue(nil)
		o.Data2.Constructor = -1059185703
		return o
	},
	1064139624: func() TLObject { // 0x3f6d7b68
		o := MakeTLJsonNull(nil)
		o.Data2.Constructor = 1064139624
		return o
	},
	-952869270: func() TLObject { // 0xc7345e6a
		o := MakeTLJsonBool(nil)
		o.Data2.Constructor = -952869270
		return o
	},
	736157604: func() TLObject { // 0x2be0dfa4
		o := MakeTLJsonNumber(nil)
		o.Data2.Constructor = 736157604
		return o
	},
	-1222740358: func() TLObject { // 0xb71e767a
		o := MakeTLJsonString(nil)
		o.Data2.Constructor = -1222740358
		return o
	},
	-146520221: func() TLObject { // 0xf7444763
		o := MakeTLJsonArray(nil)
		o.Data2.Constructor = -146520221
		return o
	},
	-1715350371: func() TLObject { // 0x99c1d49d
		o := MakeTLJsonObject(nil)
		o.Data2.Constructor = -1715350371
		return o
	},
	-1560655744: func() TLObject { // 0xa2fa4880
		o := MakeTLKeyboardButton(nil)
		o.Data2.Constructor = -1560655744
		return o
	},
	629866245: func() TLObject { // 0x258aff05
		o := MakeTLKeyboardButtonUrl(nil)
		o.Data2.Constructor = 629866245
		return o
	},
	901503851: func() TLObject { // 0x35bbdb6b
		o := MakeTLKeyboardButtonCallback(nil)
		o.Data2.Constructor = 901503851
		return o
	},
	-1318425559: func() TLObject { // 0xb16a6c29
		o := MakeTLKeyboardButtonRequestPhone(nil)
		o.Data2.Constructor = -1318425559
		return o
	},
	-59151553: func() TLObject { // 0xfc796b3f
		o := MakeTLKeyboardButtonRequestGeoLocation(nil)
		o.Data2.Constructor = -59151553
		return o
	},
	-1816527947: func() TLObject { // 0x93b9fbb5
		o := MakeTLKeyboardButtonSwitchInline(nil)
		o.Data2.Constructor = -1816527947
		return o
	},
	90744648: func() TLObject { // 0x568a748
		o := MakeTLKeyboardButtonSwitchInline(nil)
		o.Data2.Constructor = 90744648
		return o
	},
	1358175439: func() TLObject { // 0x50f41ccf
		o := MakeTLKeyboardButtonGame(nil)
		o.Data2.Constructor = 1358175439
		return o
	},
	-1344716869: func() TLObject { // 0xafd93fbb
		o := MakeTLKeyboardButtonBuy(nil)
		o.Data2.Constructor = -1344716869
		return o
	},
	280464681: func() TLObject { // 0x10b78d29
		o := MakeTLKeyboardButtonUrlAuth(nil)
		o.Data2.Constructor = 280464681
		return o
	},
	-802258988: func() TLObject { // 0xd02e7fd4
		o := MakeTLInputKeyboardButtonUrlAuth(nil)
		o.Data2.Constructor = -802258988
		return o
	},
	-1144565411: func() TLObject { // 0xbbc7515d
		o := MakeTLKeyboardButtonRequestPoll(nil)
		o.Data2.Constructor = -1144565411
		return o
	},
	-376962181: func() TLObject { // 0xe988037b
		o := MakeTLInputKeyboardButtonUserProfile(nil)
		o.Data2.Constructor = -376962181
		return o
	},
	814112961: func() TLObject { // 0x308660c1
		o := MakeTLKeyboardButtonUserProfile(nil)
		o.Data2.Constructor = 814112961
		return o
	},
	326529584: func() TLObject { // 0x13767230
		o := MakeTLKeyboardButtonWebView(nil)
		o.Data2.Constructor = 326529584
		return o
	},
	-1598009252: func() TLObject { // 0xa0c0505c
		o := MakeTLKeyboardButtonSimpleWebView(nil)
		o.Data2.Constructor = -1598009252
		return o
	},
	1406648280: func() TLObject { // 0x53d7bfd8
		o := MakeTLKeyboardButtonRequestPeer(nil)
		o.Data2.Constructor = 1406648280
		return o
	},
	218842764: func() TLObject { // 0xd0b468c
		o := MakeTLKeyboardButtonRequestPeer(nil)
		o.Data2.Constructor = 218842764
		return o
	},
	-916050683: func() TLObject { // 0xc9662d05
		o := MakeTLInputKeyboardButtonRequestPeer(nil)
		o.Data2.Constructor = -916050683
		return o
	},
	2002815875: func() TLObject { // 0x77608b83
		o := MakeTLKeyboardButtonRow(nil)
		o.Data2.Constructor = 2002815875
		return o
	},
	-886477832: func() TLObject { // 0xcb296bf8
		o := MakeTLLabeledPrice(nil)
		o.Data2.Constructor = -886477832
		return o
	},
	-209337866: func() TLObject { // 0xf385c1f6
		o := MakeTLLangPackDifference(nil)
		o.Data2.Constructor = -209337866
		return o
	},
	-288727837: func() TLObject { // 0xeeca5ce3
		o := MakeTLLangPackLanguage(nil)
		o.Data2.Constructor = -288727837
		return o
	},
	-892239370: func() TLObject { // 0xcad181f6
		o := MakeTLLangPackString(nil)
		o.Data2.Constructor = -892239370
		return o
	},
	1816636575: func() TLObject { // 0x6c47ac9f
		o := MakeTLLangPackStringPluralized(nil)
		o.Data2.Constructor = 1816636575
		return o
	},
	695856818: func() TLObject { // 0x2979eeb2
		o := MakeTLLangPackStringDeleted(nil)
		o.Data2.Constructor = 695856818
		return o
	},
	-1361650766: func() TLObject { // 0xaed6dbb2
		o := MakeTLMaskCoords(nil)
		o.Data2.Constructor = -1361650766
		return o
	},
	-1098720356: func() TLObject { // 0xbe82db9c
		o := MakeTLMediaAreaVenue(nil)
		o.Data2.Constructor = -1098720356
		return o
	},
	-1300094593: func() TLObject { // 0xb282217f
		o := MakeTLInputMediaAreaVenue(nil)
		o.Data2.Constructor = -1300094593
		return o
	},
	-891992787: func() TLObject { // 0xcad5452d
		o := MakeTLMediaAreaGeoPoint(nil)
		o.Data2.Constructor = -891992787
		return o
	},
	-544523486: func() TLObject { // 0xdf8b3b22
		o := MakeTLMediaAreaGeoPoint(nil)
		o.Data2.Constructor = -544523486
		return o
	},
	340088945: func() TLObject { // 0x14455871
		o := MakeTLMediaAreaSuggestedReaction(nil)
		o.Data2.Constructor = 340088945
		return o
	},
	1996756655: func() TLObject { // 0x770416af
		o := MakeTLMediaAreaChannelPost(nil)
		o.Data2.Constructor = 1996756655
		return o
	},
	577893055: func() TLObject { // 0x2271f2bf
		o := MakeTLInputMediaAreaChannelPost(nil)
		o.Data2.Constructor = 577893055
		return o
	},
	926421125: func() TLObject { // 0x37381085
		o := MakeTLMediaAreaUrl(nil)
		o.Data2.Constructor = 926421125
		return o
	},
	-808853502: func() TLObject { // 0xcfc9e002
		o := MakeTLMediaAreaCoordinates(nil)
		o.Data2.Constructor = -808853502
		return o
	},
	64088654: func() TLObject { // 0x3d1ea4e
		o := MakeTLMediaAreaCoordinates(nil)
		o.Data2.Constructor = 64088654
		return o
	},
	-1868117372: func() TLObject { // 0x90a6ca84
		o := MakeTLMessageEmpty(nil)
		o.Data2.Constructor = -1868117372
		return o
	},
	-1808510398: func() TLObject { // 0x94345242
		o := MakeTLMessage(nil)
		o.Data2.Constructor = -1808510398
		return o
	},
	-1109353426: func() TLObject { // 0xbde09c2e
		o := MakeTLMessage(nil)
		o.Data2.Constructor = -1109353426
		return o
	},
	592953125: func() TLObject { // 0x2357bf25
		o := MakeTLMessage(nil)
		o.Data2.Constructor = 592953125
		return o
	},
	-1502839044: func() TLObject { // 0xa66c7efc
		o := MakeTLMessage(nil)
		o.Data2.Constructor = -1502839044
		return o
	},
	508332649: func() TLObject { // 0x1e4c8a69
		o := MakeTLMessage(nil)
		o.Data2.Constructor = 508332649
		return o
	},
	1992213009: func() TLObject { // 0x76bec211
		o := MakeTLMessage(nil)
		o.Data2.Constructor = 1992213009
		return o
	},
	940666592: func() TLObject { // 0x38116ee0
		o := MakeTLMessage(nil)
		o.Data2.Constructor = 940666592
		return o
	},
	721967202: func() TLObject { // 0x2b085862
		o := MakeTLMessageService(nil)
		o.Data2.Constructor = 721967202
		return o
	},
	-1230047312: func() TLObject { // 0xb6aef7b0
		o := MakeTLMessageActionEmpty(nil)
		o.Data2.Constructor = -1230047312
		return o
	},
	-1119368275: func() TLObject { // 0xbd47cbad
		o := MakeTLMessageActionChatCreate(nil)
		o.Data2.Constructor = -1119368275
		return o
	},
	-1247687078: func() TLObject { // 0xb5a1ce5a
		o := MakeTLMessageActionChatEditTitle(nil)
		o.Data2.Constructor = -1247687078
		return o
	},
	2144015272: func() TLObject { // 0x7fcb13a8
		o := MakeTLMessageActionChatEditPhoto(nil)
		o.Data2.Constructor = 2144015272
		return o
	},
	-1780220945: func() TLObject { // 0x95e3fbef
		o := MakeTLMessageActionChatDeletePhoto(nil)
		o.Data2.Constructor = -1780220945
		return o
	},
	365886720: func() TLObject { // 0x15cefd00
		o := MakeTLMessageActionChatAddUser(nil)
		o.Data2.Constructor = 365886720
		return o
	},
	-1539362612: func() TLObject { // 0xa43f30cc
		o := MakeTLMessageActionChatDeleteUser(nil)
		o.Data2.Constructor = -1539362612
		return o
	},
	51520707: func() TLObject { // 0x31224c3
		o := MakeTLMessageActionChatJoinedByLink(nil)
		o.Data2.Constructor = 51520707
		return o
	},
	-1781355374: func() TLObject { // 0x95d2ac92
		o := MakeTLMessageActionChannelCreate(nil)
		o.Data2.Constructor = -1781355374
		return o
	},
	-519864430: func() TLObject { // 0xe1037f92
		o := MakeTLMessageActionChatMigrateTo(nil)
		o.Data2.Constructor = -519864430
		return o
	},
	-365344535: func() TLObject { // 0xea3948e9
		o := MakeTLMessageActionChannelMigrateFrom(nil)
		o.Data2.Constructor = -365344535
		return o
	},
	-1799538451: func() TLObject { // 0x94bd38ed
		o := MakeTLMessageActionPinMessage(nil)
		o.Data2.Constructor = -1799538451
		return o
	},
	-1615153660: func() TLObject { // 0x9fbab604
		o := MakeTLMessageActionHistoryClear(nil)
		o.Data2.Constructor = -1615153660
		return o
	},
	-1834538890: func() TLObject { // 0x92a72876
		o := MakeTLMessageActionGameScore(nil)
		o.Data2.Constructor = -1834538890
		return o
	},
	-1892568281: func() TLObject { // 0x8f31b327
		o := MakeTLMessageActionPaymentSentMe(nil)
		o.Data2.Constructor = -1892568281
		return o
	},
	-1776926890: func() TLObject { // 0x96163f56
		o := MakeTLMessageActionPaymentSent(nil)
		o.Data2.Constructor = -1776926890
		return o
	},
	1080663248: func() TLObject { // 0x40699cd0
		o := MakeTLMessageActionPaymentSent(nil)
		o.Data2.Constructor = 1080663248
		return o
	},
	-2132731265: func() TLObject { // 0x80e11a7f
		o := MakeTLMessageActionPhoneCall(nil)
		o.Data2.Constructor = -2132731265
		return o
	},
	1200788123: func() TLObject { // 0x4792929b
		o := MakeTLMessageActionScreenshotTaken(nil)
		o.Data2.Constructor = 1200788123
		return o
	},
	-85549226: func() TLObject { // 0xfae69f56
		o := MakeTLMessageActionCustomAction(nil)
		o.Data2.Constructor = -85549226
		return o
	},
	-988359047: func() TLObject { // 0xc516d679
		o := MakeTLMessageActionBotAllowed(nil)
		o.Data2.Constructor = -988359047
		return o
	},
	-1410748418: func() TLObject { // 0xabe9affe
		o := MakeTLMessageActionBotAllowed(nil)
		o.Data2.Constructor = -1410748418
		return o
	},
	455635795: func() TLObject { // 0x1b287353
		o := MakeTLMessageActionSecureValuesSentMe(nil)
		o.Data2.Constructor = 455635795
		return o
	},
	-648257196: func() TLObject { // 0xd95c6154
		o := MakeTLMessageActionSecureValuesSent(nil)
		o.Data2.Constructor = -648257196
		return o
	},
	-202219658: func() TLObject { // 0xf3f25f76
		o := MakeTLMessageActionContactSignUp(nil)
		o.Data2.Constructor = -202219658
		return o
	},
	-1730095465: func() TLObject { // 0x98e0d697
		o := MakeTLMessageActionGeoProximityReached(nil)
		o.Data2.Constructor = -1730095465
		return o
	},
	2047704898: func() TLObject { // 0x7a0d7f42
		o := MakeTLMessageActionGroupCall(nil)
		o.Data2.Constructor = 2047704898
		return o
	},
	1345295095: func() TLObject { // 0x502f92f7
		o := MakeTLMessageActionInviteToGroupCall(nil)
		o.Data2.Constructor = 1345295095
		return o
	},
	1007897979: func() TLObject { // 0x3c134d7b
		o := MakeTLMessageActionSetMessagesTTL(nil)
		o.Data2.Constructor = 1007897979
		return o
	},
	-1441072131: func() TLObject { // 0xaa1afbfd
		o := MakeTLMessageActionSetMessagesTTL(nil)
		o.Data2.Constructor = -1441072131
		return o
	},
	-1281329567: func() TLObject { // 0xb3a07661
		o := MakeTLMessageActionGroupCallScheduled(nil)
		o.Data2.Constructor = -1281329567
		return o
	},
	-1434950843: func() TLObject { // 0xaa786345
		o := MakeTLMessageActionSetChatTheme(nil)
		o.Data2.Constructor = -1434950843
		return o
	},
	-339958837: func() TLObject { // 0xebbca3cb
		o := MakeTLMessageActionChatJoinedByRequest(nil)
		o.Data2.Constructor = -339958837
		return o
	},
	1205698681: func() TLObject { // 0x47dd8079
		o := MakeTLMessageActionWebViewDataSentMe(nil)
		o.Data2.Constructor = 1205698681
		return o
	},
	-1262252875: func() TLObject { // 0xb4c38cb5
		o := MakeTLMessageActionWebViewDataSent(nil)
		o.Data2.Constructor = -1262252875
		return o
	},
	-935499028: func() TLObject { // 0xc83d6aec
		o := MakeTLMessageActionGiftPremium(nil)
		o.Data2.Constructor = -935499028
		return o
	},
	-1415514682: func() TLObject { // 0xaba0f5c6
		o := MakeTLMessageActionGiftPremium(nil)
		o.Data2.Constructor = -1415514682
		return o
	},
	228168278: func() TLObject { // 0xd999256
		o := MakeTLMessageActionTopicCreate(nil)
		o.Data2.Constructor = 228168278
		return o
	},
	-1064024032: func() TLObject { // 0xc0944820
		o := MakeTLMessageActionTopicEdit(nil)
		o.Data2.Constructor = -1064024032
		return o
	},
	-1316338916: func() TLObject { // 0xb18a431c
		o := MakeTLMessageActionTopicEdit(nil)
		o.Data2.Constructor = -1316338916
		return o
	},
	1474192222: func() TLObject { // 0x57de635e
		o := MakeTLMessageActionSuggestProfilePhoto(nil)
		o.Data2.Constructor = 1474192222
		return o
	},
	827428507: func() TLObject { // 0x31518e9b
		o := MakeTLMessageActionRequestedPeer(nil)
		o.Data2.Constructor = 827428507
		return o
	},
	-25742243: func() TLObject { // 0xfe77345d
		o := MakeTLMessageActionRequestedPeer(nil)
		o.Data2.Constructor = -25742243
		return o
	},
	1348510708: func() TLObject { // 0x5060a3f4
		o := MakeTLMessageActionSetChatWallPaper(nil)
		o.Data2.Constructor = 1348510708
		return o
	},
	-1136350937: func() TLObject { // 0xbc44a927
		o := MakeTLMessageActionSetChatWallPaper(nil)
		o.Data2.Constructor = -1136350937
		return o
	},
	1737240073: func() TLObject { // 0x678c2e09
		o := MakeTLMessageActionGiftCode(nil)
		o.Data2.Constructor = 1737240073
		return o
	},
	-758129906: func() TLObject { // 0xd2cfdb0e
		o := MakeTLMessageActionGiftCode(nil)
		o.Data2.Constructor = -758129906
		return o
	},
	858499565: func() TLObject { // 0x332ba9ed
		o := MakeTLMessageActionGiveawayLaunch(nil)
		o.Data2.Constructor = 858499565
		return o
	},
	715107781: func() TLObject { // 0x2a9fadc5
		o := MakeTLMessageActionGiveawayResults(nil)
		o.Data2.Constructor = 715107781
		return o
	},
	-872240531: func() TLObject { // 0xcc02aa6d
		o := MakeTLMessageActionBoostApply(nil)
		o.Data2.Constructor = -872240531
		return o
	},
	-1816979384: func() TLObject { // 0x93b31848
		o := MakeTLMessageActionRequestedPeerSentMe(nil)
		o.Data2.Constructor = -1816979384
		return o
	},
	-1065845395: func() TLObject { // 0xc0787d6d
		o := MakeTLMessageActionSetSameChatWallPaper(nil)
		o.Data2.Constructor = -1065845395
		return o
	},
	-404267113: func() TLObject { // 0xe7e75f97
		o := MakeTLMessageActionAttachMenuBotAllowed(nil)
		o.Data2.Constructor = -404267113
		return o
	},
	805171639: func() TLObject { // 0x2ffdf1b7
		o := MakeTLMessageActionBizDataRaw(nil)
		o.Data2.Constructor = 805171639
		return o
	},
	964662120: func() TLObject { // 0x397f9368
		o := MakeTLMessageBox(nil)
		o.Data2.Constructor = 964662120
		return o
	},
	1393091297: func() TLObject { // 0x5308e2e1
		o := MakeTLMessageBoxList(nil)
		o.Data2.Constructor = 1393091297
		return o
	},
	-2136871889: func() TLObject { // 0x80a1ec2f
		o := MakeTLMessageBoxListSlice(nil)
		o.Data2.Constructor = -2136871889
		return o
	},
	-1148011883: func() TLObject { // 0xbb92ba95
		o := MakeTLMessageEntityUnknown(nil)
		o.Data2.Constructor = -1148011883
		return o
	},
	-100378723: func() TLObject { // 0xfa04579d
		o := MakeTLMessageEntityMention(nil)
		o.Data2.Constructor = -100378723
		return o
	},
	1868782349: func() TLObject { // 0x6f635b0d
		o := MakeTLMessageEntityHashtag(nil)
		o.Data2.Constructor = 1868782349
		return o
	},
	1827637959: func() TLObject { // 0x6cef8ac7
		o := MakeTLMessageEntityBotCommand(nil)
		o.Data2.Constructor = 1827637959
		return o
	},
	1859134776: func() TLObject { // 0x6ed02538
		o := MakeTLMessageEntityUrl(nil)
		o.Data2.Constructor = 1859134776
		return o
	},
	1692693954: func() TLObject { // 0x64e475c2
		o := MakeTLMessageEntityEmail(nil)
		o.Data2.Constructor = 1692693954
		return o
	},
	-1117713463: func() TLObject { // 0xbd610bc9
		o := MakeTLMessageEntityBold(nil)
		o.Data2.Constructor = -1117713463
		return o
	},
	-2106619040: func() TLObject { // 0x826f8b60
		o := MakeTLMessageEntityItalic(nil)
		o.Data2.Constructor = -2106619040
		return o
	},
	681706865: func() TLObject { // 0x28a20571
		o := MakeTLMessageEntityCode(nil)
		o.Data2.Constructor = 681706865
		return o
	},
	1938967520: func() TLObject { // 0x73924be0
		o := MakeTLMessageEntityPre(nil)
		o.Data2.Constructor = 1938967520
		return o
	},
	1990644519: func() TLObject { // 0x76a6d327
		o := MakeTLMessageEntityTextUrl(nil)
		o.Data2.Constructor = 1990644519
		return o
	},
	-595914432: func() TLObject { // 0xdc7b1140
		o := MakeTLMessageEntityMentionName(nil)
		o.Data2.Constructor = -595914432
		return o
	},
	546203849: func() TLObject { // 0x208e68c9
		o := MakeTLInputMessageEntityMentionName(nil)
		o.Data2.Constructor = 546203849
		return o
	},
	-1687559349: func() TLObject { // 0x9b69e34b
		o := MakeTLMessageEntityPhone(nil)
		o.Data2.Constructor = -1687559349
		return o
	},
	1280209983: func() TLObject { // 0x4c4e743f
		o := MakeTLMessageEntityCashtag(nil)
		o.Data2.Constructor = 1280209983
		return o
	},
	-1672577397: func() TLObject { // 0x9c4e7e8b
		o := MakeTLMessageEntityUnderline(nil)
		o.Data2.Constructor = -1672577397
		return o
	},
	-1090087980: func() TLObject { // 0xbf0693d4
		o := MakeTLMessageEntityStrike(nil)
		o.Data2.Constructor = -1090087980
		return o
	},
	1981704948: func() TLObject { // 0x761e6af4
		o := MakeTLMessageEntityBankCard(nil)
		o.Data2.Constructor = 1981704948
		return o
	},
	852137487: func() TLObject { // 0x32ca960f
		o := MakeTLMessageEntitySpoiler(nil)
		o.Data2.Constructor = 852137487
		return o
	},
	-925956616: func() TLObject { // 0xc8cf05f8
		o := MakeTLMessageEntityCustomEmoji(nil)
		o.Data2.Constructor = -925956616
		return o
	},
	-238245204: func() TLObject { // 0xf1ccaaac
		o := MakeTLMessageEntityBlockquote(nil)
		o.Data2.Constructor = -238245204
		return o
	},
	34469328: func() TLObject { // 0x20df5d0
		o := MakeTLMessageEntityBlockquote(nil)
		o.Data2.Constructor = 34469328
		return o
	},
	-1386050360: func() TLObject { // 0xad628cc8
		o := MakeTLMessageExtendedMediaPreview(nil)
		o.Data2.Constructor = -1386050360
		return o
	},
	-297296796: func() TLObject { // 0xee479c64
		o := MakeTLMessageExtendedMedia(nil)
		o.Data2.Constructor = -297296796
		return o
	},
	1313731771: func() TLObject { // 0x4e4df4bb
		o := MakeTLMessageFwdHeader(nil)
		o.Data2.Constructor = 1313731771
		return o
	},
	1601666510: func() TLObject { // 0x5f777dce
		o := MakeTLMessageFwdHeader(nil)
		o.Data2.Constructor = 1601666510
		return o
	},
	-1387279939: func() TLObject { // 0xad4fc9bd
		o := MakeTLMessageInteractionCounters(nil)
		o.Data2.Constructor = -1387279939
		return o
	},
	1038967584: func() TLObject { // 0x3ded6320
		o := MakeTLMessageMediaEmpty(nil)
		o.Data2.Constructor = 1038967584
		return o
	},
	1766936791: func() TLObject { // 0x695150d7
		o := MakeTLMessageMediaPhoto(nil)
		o.Data2.Constructor = 1766936791
		return o
	},
	1457575028: func() TLObject { // 0x56e0d474
		o := MakeTLMessageMediaGeo(nil)
		o.Data2.Constructor = 1457575028
		return o
	},
	1882335561: func() TLObject { // 0x70322949
		o := MakeTLMessageMediaContact(nil)
		o.Data2.Constructor = 1882335561
		return o
	},
	-1618676578: func() TLObject { // 0x9f84f49e
		o := MakeTLMessageMediaUnsupported(nil)
		o.Data2.Constructor = -1618676578
		return o
	},
	1291114285: func() TLObject { // 0x4cf4d72d
		o := MakeTLMessageMediaDocument(nil)
		o.Data2.Constructor = 1291114285
		return o
	},
	-1666158377: func() TLObject { // 0x9cb070d7
		o := MakeTLMessageMediaDocument(nil)
		o.Data2.Constructor = -1666158377
		return o
	},
	-571405253: func() TLObject { // 0xddf10c3b
		o := MakeTLMessageMediaWebPage(nil)
		o.Data2.Constructor = -571405253
		return o
	},
	-1557277184: func() TLObject { // 0xa32dd600
		o := MakeTLMessageMediaWebPage(nil)
		o.Data2.Constructor = -1557277184
		return o
	},
	784356159: func() TLObject { // 0x2ec0533f
		o := MakeTLMessageMediaVenue(nil)
		o.Data2.Constructor = 784356159
		return o
	},
	-38694904: func() TLObject { // 0xfdb19008
		o := MakeTLMessageMediaGame(nil)
		o.Data2.Constructor = -38694904
		return o
	},
	-156940077: func() TLObject { // 0xf6a548d3
		o := MakeTLMessageMediaInvoice(nil)
		o.Data2.Constructor = -156940077
		return o
	},
	-2074799289: func() TLObject { // 0x84551347
		o := MakeTLMessageMediaInvoice(nil)
		o.Data2.Constructor = -2074799289
		return o
	},
	-1186937242: func() TLObject { // 0xb940c666
		o := MakeTLMessageMediaGeoLive(nil)
		o.Data2.Constructor = -1186937242
		return o
	},
	1272375192: func() TLObject { // 0x4bd6e798
		o := MakeTLMessageMediaPoll(nil)
		o.Data2.Constructor = 1272375192
		return o
	},
	1065280907: func() TLObject { // 0x3f7ee58b
		o := MakeTLMessageMediaDice(nil)
		o.Data2.Constructor = 1065280907
		return o
	},
	1758159491: func() TLObject { // 0x68cb6283
		o := MakeTLMessageMediaStory(nil)
		o.Data2.Constructor = 1758159491
		return o
	},
	-877523576: func() TLObject { // 0xcbb20d88
		o := MakeTLMessageMediaStory(nil)
		o.Data2.Constructor = -877523576
		return o
	},
	-626162256: func() TLObject { // 0xdaad85b0
		o := MakeTLMessageMediaGiveaway(nil)
		o.Data2.Constructor = -626162256
		return o
	},
	1478887012: func() TLObject { // 0x58260664
		o := MakeTLMessageMediaGiveaway(nil)
		o.Data2.Constructor = 1478887012
		return o
	},
	-963047320: func() TLObject { // 0xc6991068
		o := MakeTLMessageMediaGiveawayResults(nil)
		o.Data2.Constructor = -963047320
		return o
	},
	-1467669359: func() TLObject { // 0xa8852491
		o := MakeTLMessageMediaPaidMedia(nil)
		o.Data2.Constructor = -1467669359
		return o
	},
	2124445994: func() TLObject { // 0x7ea0792a
		o := MakeTLMessageMediaBizDataRaw(nil)
		o.Data2.Constructor = 2124445994
		return o
	},
	-1938180548: func() TLObject { // 0x8c79b63c
		o := MakeTLMessagePeerReaction(nil)
		o.Data2.Constructor = -1938180548
		return o
	},
	-1319698788: func() TLObject { // 0xb156fe9c
		o := MakeTLMessagePeerReaction(nil)
		o.Data2.Constructor = -1319698788
		return o
	},
	1370914559: func() TLObject { // 0x51b67eff
		o := MakeTLMessagePeerReaction(nil)
		o.Data2.Constructor = 1370914559
		return o
	},
	-1228133028: func() TLObject { // 0xb6cc2d5c
		o := MakeTLMessagePeerVote(nil)
		o.Data2.Constructor = -1228133028
		return o
	},
	1959634180: func() TLObject { // 0x74cda504
		o := MakeTLMessagePeerVoteInputOption(nil)
		o.Data2.Constructor = 1959634180
		return o
	},
	1177089766: func() TLObject { // 0x4628f6e6
		o := MakeTLMessagePeerVoteMultiple(nil)
		o.Data2.Constructor = 1177089766
		return o
	},
	182649427: func() TLObject { // 0xae30253
		o := MakeTLMessageRange(nil)
		o.Data2.Constructor = 182649427
		return o
	},
	1328256121: func() TLObject { // 0x4f2b9479
		o := MakeTLMessageReactions(nil)
		o.Data2.Constructor = 1328256121
		return o
	},
	-2083123262: func() TLObject { // 0x83d60fc2
		o := MakeTLMessageReplies(nil)
		o.Data2.Constructor = -2083123262
		return o
	},
	-1346631205: func() TLObject { // 0xafbc09db
		o := MakeTLMessageReplyHeader(nil)
		o.Data2.Constructor = -1346631205
		return o
	},
	1860946621: func() TLObject { // 0x6eebcabd
		o := MakeTLMessageReplyHeader(nil)
		o.Data2.Constructor = 1860946621
		return o
	},
	-1495959709: func() TLObject { // 0xa6d57763
		o := MakeTLMessageReplyHeader(nil)
		o.Data2.Constructor = -1495959709
		return o
	},
	240843065: func() TLObject { // 0xe5af939
		o := MakeTLMessageReplyStoryHeader(nil)
		o.Data2.Constructor = 240843065
		return o
	},
	-1667711039: func() TLObject { // 0x9c98bfc1
		o := MakeTLMessageReplyStoryHeader(nil)
		o.Data2.Constructor = -1667711039
		return o
	},
	886196148: func() TLObject { // 0x34d247b4
		o := MakeTLMessageUserVote(nil)
		o.Data2.Constructor = 886196148
		return o
	},
	1017491692: func() TLObject { // 0x3ca5b0ec
		o := MakeTLMessageUserVoteInputOption(nil)
		o.Data2.Constructor = 1017491692
		return o
	},
	-1973033641: func() TLObject { // 0x8a65e557
		o := MakeTLMessageUserVoteMultiple(nil)
		o.Data2.Constructor = -1973033641
		return o
	},
	1163625789: func() TLObject { // 0x455b853d
		o := MakeTLMessageViews(nil)
		o.Data2.Constructor = 1163625789
		return o
	},
	1474492012: func() TLObject { // 0x57e2f66c
		o := MakeTLInputMessagesFilterEmpty(nil)
		o.Data2.Constructor = 1474492012
		return o
	},
	-1777752804: func() TLObject { // 0x9609a51c
		o := MakeTLInputMessagesFilterPhotos(nil)
		o.Data2.Constructor = -1777752804
		return o
	},
	-1614803355: func() TLObject { // 0x9fc00e65
		o := MakeTLInputMessagesFilterVideo(nil)
		o.Data2.Constructor = -1614803355
		return o
	},
	1458172132: func() TLObject { // 0x56e9f0e4
		o := MakeTLInputMessagesFilterPhotoVideo(nil)
		o.Data2.Constructor = 1458172132
		return o
	},
	-1629621880: func() TLObject { // 0x9eddf188
		o := MakeTLInputMessagesFilterDocument(nil)
		o.Data2.Constructor = -1629621880
		return o
	},
	2129714567: func() TLObject { // 0x7ef0dd87
		o := MakeTLInputMessagesFilterUrl(nil)
		o.Data2.Constructor = 2129714567
		return o
	},
	-3644025: func() TLObject { // 0xffc86587
		o := MakeTLInputMessagesFilterGif(nil)
		o.Data2.Constructor = -3644025
		return o
	},
	1358283666: func() TLObject { // 0x50f5c392
		o := MakeTLInputMessagesFilterVoice(nil)
		o.Data2.Constructor = 1358283666
		return o
	},
	928101534: func() TLObject { // 0x3751b49e
		o := MakeTLInputMessagesFilterMusic(nil)
		o.Data2.Constructor = 928101534
		return o
	},
	975236280: func() TLObject { // 0x3a20ecb8
		o := MakeTLInputMessagesFilterChatPhotos(nil)
		o.Data2.Constructor = 975236280
		return o
	},
	-2134272152: func() TLObject { // 0x80c99768
		o := MakeTLInputMessagesFilterPhoneCalls(nil)
		o.Data2.Constructor = -2134272152
		return o
	},
	2054952868: func() TLObject { // 0x7a7c17a4
		o := MakeTLInputMessagesFilterRoundVoice(nil)
		o.Data2.Constructor = 2054952868
		return o
	},
	-1253451181: func() TLObject { // 0xb549da53
		o := MakeTLInputMessagesFilterRoundVideo(nil)
		o.Data2.Constructor = -1253451181
		return o
	},
	-1040652646: func() TLObject { // 0xc1f8e69a
		o := MakeTLInputMessagesFilterMyMentions(nil)
		o.Data2.Constructor = -1040652646
		return o
	},
	-419271411: func() TLObject { // 0xe7026d0d
		o := MakeTLInputMessagesFilterGeo(nil)
		o.Data2.Constructor = -419271411
		return o
	},
	-530392189: func() TLObject { // 0xe062db83
		o := MakeTLInputMessagesFilterContacts(nil)
		o.Data2.Constructor = -530392189
		return o
	},
	464520273: func() TLObject { // 0x1bb00451
		o := MakeTLInputMessagesFilterPinned(nil)
		o.Data2.Constructor = 464520273
		return o
	},
	-275956116: func() TLObject { // 0xef8d3e6c
		o := MakeTLMessagesAffectedFoundMessages(nil)
		o.Data2.Constructor = -275956116
		return o
	},
	-1269012015: func() TLObject { // 0xb45c69d1
		o := MakeTLMessagesAffectedHistory(nil)
		o.Data2.Constructor = -1269012015
		return o
	},
	-2066640507: func() TLObject { // 0x84d19185
		o := MakeTLMessagesAffectedMessages(nil)
		o.Data2.Constructor = -2066640507
		return o
	},
	-395967805: func() TLObject { // 0xe86602c3
		o := MakeTLMessagesAllStickersNotModified(nil)
		o.Data2.Constructor = -395967805
		return o
	},
	-843329861: func() TLObject { // 0xcdbbcebb
		o := MakeTLMessagesAllStickers(nil)
		o.Data2.Constructor = -843329861
		return o
	},
	1338747336: func() TLObject { // 0x4fcba9c8
		o := MakeTLMessagesArchivedStickers(nil)
		o.Data2.Constructor = 1338747336
		return o
	},
	-772957605: func() TLObject { // 0xd1ed9a5b
		o := MakeTLMessagesAvailableEffectsNotModified(nil)
		o.Data2.Constructor = -772957605
		return o
	},
	-1109696146: func() TLObject { // 0xbddb616e
		o := MakeTLMessagesAvailableEffects(nil)
		o.Data2.Constructor = -1109696146
		return o
	},
	-1626924713: func() TLObject { // 0x9f071957
		o := MakeTLMessagesAvailableReactionsNotModified(nil)
		o.Data2.Constructor = -1626924713
		return o
	},
	1989032621: func() TLObject { // 0x768e3aad
		o := MakeTLMessagesAvailableReactions(nil)
		o.Data2.Constructor = 1989032621
		return o
	},
	-347034123: func() TLObject { // 0xeb50adf5
		o := MakeTLMessagesBotApp(nil)
		o.Data2.Constructor = -347034123
		return o
	},
	911761060: func() TLObject { // 0x36585ea4
		o := MakeTLMessagesBotCallbackAnswer(nil)
		o.Data2.Constructor = 911761060
		return o
	},
	-534646026: func() TLObject { // 0xe021f2f6
		o := MakeTLMessagesBotResults(nil)
		o.Data2.Constructor = -534646026
		return o
	},
	-1803769784: func() TLObject { // 0x947ca848
		o := MakeTLMessagesBotResults(nil)
		o.Data2.Constructor = -1803769784
		return o
	},
	-1231326505: func() TLObject { // 0xb69b72d7
		o := MakeTLMessagesChatAdminsWithInvites(nil)
		o.Data2.Constructor = -1231326505
		return o
	},
	-438840932: func() TLObject { // 0xe5d7d19c
		o := MakeTLMessagesChatFull(nil)
		o.Data2.Constructor = -438840932
		return o
	},
	-2118733814: func() TLObject { // 0x81b6b00a
		o := MakeTLMessagesChatInviteImporters(nil)
		o.Data2.Constructor = -2118733814
		return o
	},
	1694474197: func() TLObject { // 0x64ff9fd5
		o := MakeTLMessagesChats(nil)
		o.Data2.Constructor = 1694474197
		return o
	},
	-1663561404: func() TLObject { // 0x9cd81144
		o := MakeTLMessagesChatsSlice(nil)
		o.Data2.Constructor = -1663561404
		return o
	},
	-1571952873: func() TLObject { // 0xa24de717
		o := MakeTLMessagesCheckedHistoryImportPeer(nil)
		o.Data2.Constructor = -1571952873
		return o
	},
	-1058912715: func() TLObject { // 0xc0e24635
		o := MakeTLMessagesDhConfigNotModified(nil)
		o.Data2.Constructor = -1058912715
		return o
	},
	740433629: func() TLObject { // 0x2c221edd
		o := MakeTLMessagesDhConfig(nil)
		o.Data2.Constructor = 740433629
		return o
	},
	718878489: func() TLObject { // 0x2ad93719
		o := MakeTLMessagesDialogFilters(nil)
		o.Data2.Constructor = 718878489
		return o
	},
	364538944: func() TLObject { // 0x15ba6c40
		o := MakeTLMessagesDialogs(nil)
		o.Data2.Constructor = 364538944
		return o
	},
	1910543603: func() TLObject { // 0x71e094f3
		o := MakeTLMessagesDialogsSlice(nil)
		o.Data2.Constructor = 1910543603
		return o
	},
	-253500010: func() TLObject { // 0xf0e3e596
		o := MakeTLMessagesDialogsNotModified(nil)
		o.Data2.Constructor = -253500010
		return o
	},
	-1506535550: func() TLObject { // 0xa6341782
		o := MakeTLMessagesDiscussionMessage(nil)
		o.Data2.Constructor = -1506535550
		return o
	},
	1874111879: func() TLObject { // 0x6fb4ad87
		o := MakeTLMessagesEmojiGroupsNotModified(nil)
		o.Data2.Constructor = 1874111879
		return o
	},
	-2011186869: func() TLObject { // 0x881fb94b
		o := MakeTLMessagesEmojiGroups(nil)
		o.Data2.Constructor = -2011186869
		return o
	},
	410107472: func() TLObject { // 0x1871be50
		o := MakeTLMessagesExportedChatInvite(nil)
		o.Data2.Constructor = 410107472
		return o
	},
	572915951: func() TLObject { // 0x222600ef
		o := MakeTLMessagesExportedChatInviteReplaced(nil)
		o.Data2.Constructor = 572915951
		return o
	},
	-1111085620: func() TLObject { // 0xbdc62dcc
		o := MakeTLMessagesExportedChatInvites(nil)
		o.Data2.Constructor = -1111085620
		return o
	},
	-1634752813: func() TLObject { // 0x9e8fa6d3
		o := MakeTLMessagesFavedStickersNotModified(nil)
		o.Data2.Constructor = -1634752813
		return o
	},
	750063767: func() TLObject { // 0x2cb51097
		o := MakeTLMessagesFavedStickers(nil)
		o.Data2.Constructor = 750063767
		return o
	},
	-958657434: func() TLObject { // 0xc6dc0c66
		o := MakeTLMessagesFeaturedStickersNotModified(nil)
		o.Data2.Constructor = -958657434
		return o
	},
	-1103615738: func() TLObject { // 0xbe382906
		o := MakeTLMessagesFeaturedStickers(nil)
		o.Data2.Constructor = -1103615738
		return o
	},
	-2067782896: func() TLObject { // 0x84c02310
		o := MakeTLMessagesFeaturedStickers(nil)
		o.Data2.Constructor = -2067782896
		return o
	},
	913709011: func() TLObject { // 0x367617d3
		o := MakeTLMessagesForumTopics(nil)
		o.Data2.Constructor = 913709011
		return o
	},
	223655517: func() TLObject { // 0xd54b65d
		o := MakeTLMessagesFoundStickerSetsNotModified(nil)
		o.Data2.Constructor = 223655517
		return o
	},
	-1963942446: func() TLObject { // 0x8af09dd2
		o := MakeTLMessagesFoundStickerSets(nil)
		o.Data2.Constructor = -1963942446
		return o
	},
	-1707344487: func() TLObject { // 0x9a3bfd99
		o := MakeTLMessagesHighScores(nil)
		o.Data2.Constructor = -1707344487
		return o
	},
	375566091: func() TLObject { // 0x1662af0b
		o := MakeTLMessagesHistoryImport(nil)
		o.Data2.Constructor = 375566091
		return o
	},
	1578088377: func() TLObject { // 0x5e0fb7b9
		o := MakeTLMessagesHistoryImportParsed(nil)
		o.Data2.Constructor = 1578088377
		return o
	},
	-1456996667: func() TLObject { // 0xa927fec5
		o := MakeTLMessagesInactiveChats(nil)
		o.Data2.Constructor = -1456996667
		return o
	},
	2136862630: func() TLObject { // 0x7f5defa6
		o := MakeTLMessagesInvitedUsers(nil)
		o.Data2.Constructor = 2136862630
		return o
	},
	649453030: func() TLObject { // 0x26b5dde6
		o := MakeTLMessagesMessageEditData(nil)
		o.Data2.Constructor = 649453030
		return o
	},
	834488621: func() TLObject { // 0x31bd492d
		o := MakeTLMessagesMessageReactionsList(nil)
		o.Data2.Constructor = 834488621
		return o
	},
	-1228606141: func() TLObject { // 0xb6c4f543
		o := MakeTLMessagesMessageViews(nil)
		o.Data2.Constructor = -1228606141
		return o
	},
	-1938715001: func() TLObject { // 0x8c718e87
		o := MakeTLMessagesMessages(nil)
		o.Data2.Constructor = -1938715001
		return o
	},
	978610270: func() TLObject { // 0x3a54685e
		o := MakeTLMessagesMessagesSlice(nil)
		o.Data2.Constructor = 978610270
		return o
	},
	-948520370: func() TLObject { // 0xc776ba4e
		o := MakeTLMessagesChannelMessages(nil)
		o.Data2.Constructor = -948520370
		return o
	},
	1682413576: func() TLObject { // 0x64479808
		o := MakeTLMessagesChannelMessages(nil)
		o.Data2.Constructor = 1682413576
		return o
	},
	1951620897: func() TLObject { // 0x74535f21
		o := MakeTLMessagesMessagesNotModified(nil)
		o.Data2.Constructor = 1951620897
		return o
	},
	-83926371: func() TLObject { // 0xfaff629d
		o := MakeTLMessagesMyStickers(nil)
		o.Data2.Constructor = -83926371
		return o
	},
	863093588: func() TLObject { // 0x3371c354
		o := MakeTLMessagesPeerDialogs(nil)
		o.Data2.Constructor = 863093588
		return o
	},
	1753266509: func() TLObject { // 0x6880b94d
		o := MakeTLMessagesPeerSettings(nil)
		o.Data2.Constructor = 1753266509
		return o
	},
	-963811691: func() TLObject { // 0xc68d6695
		o := MakeTLMessagesQuickReplies(nil)
		o.Data2.Constructor = -963811691
		return o
	},
	1603398491: func() TLObject { // 0x5f91eb5b
		o := MakeTLMessagesQuickRepliesNotModified(nil)
		o.Data2.Constructor = 1603398491
		return o
	},
	-1334846497: func() TLObject { // 0xb06fdbdf
		o := MakeTLMessagesReactionsNotModified(nil)
		o.Data2.Constructor = -1334846497
		return o
	},
	-352454890: func() TLObject { // 0xeafdf716
		o := MakeTLMessagesReactions(nil)
		o.Data2.Constructor = -352454890
		return o
	},
	186120336: func() TLObject { // 0xb17f890
		o := MakeTLMessagesRecentStickersNotModified(nil)
		o.Data2.Constructor = 186120336
		return o
	},
	-1999405994: func() TLObject { // 0x88d37c56
		o := MakeTLMessagesRecentStickers(nil)
		o.Data2.Constructor = -1999405994
		return o
	},
	-130358751: func() TLObject { // 0xf83ae221
		o := MakeTLMessagesSavedDialogs(nil)
		o.Data2.Constructor = -130358751
		return o
	},
	1153080793: func() TLObject { // 0x44ba9dd9
		o := MakeTLMessagesSavedDialogsSlice(nil)
		o.Data2.Constructor = 1153080793
		return o
	},
	-1071681560: func() TLObject { // 0xc01f6fe8
		o := MakeTLMessagesSavedDialogsNotModified(nil)
		o.Data2.Constructor = -1071681560
		return o
	},
	-402498398: func() TLObject { // 0xe8025ca2
		o := MakeTLMessagesSavedGifsNotModified(nil)
		o.Data2.Constructor = -402498398
		return o
	},
	-2069878259: func() TLObject { // 0x84a02a0d
		o := MakeTLMessagesSavedGifs(nil)
		o.Data2.Constructor = -2069878259
		return o
	},
	-2003084817: func() TLObject { // 0x889b59ef
		o := MakeTLMessagesSavedReactionTagsNotModified(nil)
		o.Data2.Constructor = -2003084817
		return o
	},
	844731658: func() TLObject { // 0x3259950a
		o := MakeTLMessagesSavedReactionTags(nil)
		o.Data2.Constructor = 844731658
		return o
	},
	-398136321: func() TLObject { // 0xe844ebff
		o := MakeTLMessagesSearchCounter(nil)
		o.Data2.Constructor = -398136321
		return o
	},
	343859772: func() TLObject { // 0x147ee23c
		o := MakeTLMessagesSearchResultsCalendar(nil)
		o.Data2.Constructor = 343859772
		return o
	},
	1404185519: func() TLObject { // 0x53b22baf
		o := MakeTLMessagesSearchResultsPositions(nil)
		o.Data2.Constructor = 1404185519
		return o
	},
	1443858741: func() TLObject { // 0x560f8935
		o := MakeTLMessagesSentEncryptedMessage(nil)
		o.Data2.Constructor = 1443858741
		return o
	},
	-1802240206: func() TLObject { // 0x9493ff32
		o := MakeTLMessagesSentEncryptedFile(nil)
		o.Data2.Constructor = -1802240206
		return o
	},
	-907141753: func() TLObject { // 0xc9ee1d87
		o := MakeTLMessagesSponsoredMessages(nil)
		o.Data2.Constructor = -907141753
		return o
	},
	1705297877: func() TLObject { // 0x65a4c7d5
		o := MakeTLMessagesSponsoredMessages(nil)
		o.Data2.Constructor = 1705297877
		return o
	},
	406407439: func() TLObject { // 0x1839490f
		o := MakeTLMessagesSponsoredMessagesEmpty(nil)
		o.Data2.Constructor = 406407439
		return o
	},
	1846886166: func() TLObject { // 0x6e153f16
		o := MakeTLMessagesStickerSet(nil)
		o.Data2.Constructor = 1846886166
		return o
	},
	-1240849242: func() TLObject { // 0xb60a24a6
		o := MakeTLMessagesStickerSet(nil)
		o.Data2.Constructor = -1240849242
		return o
	},
	-738646805: func() TLObject { // 0xd3f924eb
		o := MakeTLMessagesStickerSetNotModified(nil)
		o.Data2.Constructor = -738646805
		return o
	},
	946083368: func() TLObject { // 0x38641628
		o := MakeTLMessagesStickerSetInstallResultSuccess(nil)
		o.Data2.Constructor = 946083368
		return o
	},
	904138920: func() TLObject { // 0x35e410a8
		o := MakeTLMessagesStickerSetInstallResultArchive(nil)
		o.Data2.Constructor = 904138920
		return o
	},
	-244016606: func() TLObject { // 0xf1749a22
		o := MakeTLMessagesStickersNotModified(nil)
		o.Data2.Constructor = -244016606
		return o
	},
	816245886: func() TLObject { // 0x30a6ec7e
		o := MakeTLMessagesStickers(nil)
		o.Data2.Constructor = 816245886
		return o
	},
	-809903785: func() TLObject { // 0xcfb9d957
		o := MakeTLMessagesTranscribedAudio(nil)
		o.Data2.Constructor = -809903785
		return o
	},
	-1821037486: func() TLObject { // 0x93752c52
		o := MakeTLMessagesTranscribedAudio(nil)
		o.Data2.Constructor = -1821037486
		return o
	},
	870003448: func() TLObject { // 0x33db32f8
		o := MakeTLMessagesTranslateResult(nil)
		o.Data2.Constructor = 870003448
		return o
	},
	1741309751: func() TLObject { // 0x67ca4737
		o := MakeTLMessagesTranslateNoResult(nil)
		o.Data2.Constructor = 1741309751
		return o
	},
	-1575684144: func() TLObject { // 0xa214f7d0
		o := MakeTLMessagesTranslateResultText(nil)
		o.Data2.Constructor = -1575684144
		return o
	},
	1218005070: func() TLObject { // 0x4899484e
		o := MakeTLMessagesVotesList(nil)
		o.Data2.Constructor = 1218005070
		return o
	},
	136574537: func() TLObject { // 0x823f649
		o := MakeTLMessagesVotesList(nil)
		o.Data2.Constructor = 136574537
		return o
	},
	-44166467: func() TLObject { // 0xfd5e12bd
		o := MakeTLMessagesWebPage(nil)
		o.Data2.Constructor = -44166467
		return o
	},
	1653379620: func() TLObject { // 0x628c9224
		o := MakeTLMissingInvitee(nil)
		o.Data2.Constructor = 1653379620
		return o
	},
	-34609042: func() TLObject { // 0xfdefe86e
		o := MakeTLMutableChat(nil)
		o.Data2.Constructor = -34609042
		return o
	},
	917538818: func() TLObject { // 0x36b08802
		o := MakeTLMutableUsers(nil)
		o.Data2.Constructor = 917538818
		return o
	},
	-1001897636: func() TLObject { // 0xc448415c
		o := MakeTLMyBoost(nil)
		o.Data2.Constructor = -1001897636
		return o
	},
	-1910892683: func() TLObject { // 0x8e1a1775
		o := MakeTLNearestDc(nil)
		o.Data2.Constructor = -1910892683
		return o
	},
	-1746354498: func() TLObject { // 0x97e8bebe
		o := MakeTLNotificationSoundDefault(nil)
		o.Data2.Constructor = -1746354498
		return o
	},
	1863070943: func() TLObject { // 0x6f0c34df
		o := MakeTLNotificationSoundNone(nil)
		o.Data2.Constructor = 1863070943
		return o
	},
	-2096391452: func() TLObject { // 0x830b9ae4
		o := MakeTLNotificationSoundLocal(nil)
		o.Data2.Constructor = -2096391452
		return o
	},
	-9666487: func() TLObject { // 0xff6c8049
		o := MakeTLNotificationSoundRingtone(nil)
		o.Data2.Constructor = -9666487
		return o
	},
	-1613493288: func() TLObject { // 0x9fd40bd8
		o := MakeTLNotifyPeer(nil)
		o.Data2.Constructor = -1613493288
		return o
	},
	-1261946036: func() TLObject { // 0xb4c83b4c
		o := MakeTLNotifyUsers(nil)
		o.Data2.Constructor = -1261946036
		return o
	},
	-1073230141: func() TLObject { // 0xc007cec3
		o := MakeTLNotifyChats(nil)
		o.Data2.Constructor = -1073230141
		return o
	},
	-703403793: func() TLObject { // 0xd612e8ef
		o := MakeTLNotifyBroadcasts(nil)
		o.Data2.Constructor = -703403793
		return o
	},
	577659656: func() TLObject { // 0x226e6308
		o := MakeTLNotifyForumTopic(nil)
		o.Data2.Constructor = 577659656
		return o
	},
	1450380236: func() TLObject { // 0x56730bcc
		o := MakeTLNull(nil)
		o.Data2.Constructor = 1450380236
		return o
	},
	1001931436: func() TLObject { // 0x3bb842ac
		o := MakeTLOutboxReadDate(nil)
		o.Data2.Constructor = 1001931436
		return o
	},
	-1738178803: func() TLObject { // 0x98657f0d
		o := MakeTLPage(nil)
		o.Data2.Constructor = -1738178803
		return o
	},
	324435594: func() TLObject { // 0x13567e8a
		o := MakeTLPageBlockUnsupported(nil)
		o.Data2.Constructor = 324435594
		return o
	},
	1890305021: func() TLObject { // 0x70abc3fd
		o := MakeTLPageBlockTitle(nil)
		o.Data2.Constructor = 1890305021
		return o
	},
	-1879401953: func() TLObject { // 0x8ffa9a1f
		o := MakeTLPageBlockSubtitle(nil)
		o.Data2.Constructor = -1879401953
		return o
	},
	-1162877472: func() TLObject { // 0xbaafe5e0
		o := MakeTLPageBlockAuthorDate(nil)
		o.Data2.Constructor = -1162877472
		return o
	},
	-1076861716: func() TLObject { // 0xbfd064ec
		o := MakeTLPageBlockHeader(nil)
		o.Data2.Constructor = -1076861716
		return o
	},
	-248793375: func() TLObject { // 0xf12bb6e1
		o := MakeTLPageBlockSubheader(nil)
		o.Data2.Constructor = -248793375
		return o
	},
	1182402406: func() TLObject { // 0x467a0766
		o := MakeTLPageBlockParagraph(nil)
		o.Data2.Constructor = 1182402406
		return o
	},
	-1066346178: func() TLObject { // 0xc070d93e
		o := MakeTLPageBlockPreformatted(nil)
		o.Data2.Constructor = -1066346178
		return o
	},
	1216809369: func() TLObject { // 0x48870999
		o := MakeTLPageBlockFooter(nil)
		o.Data2.Constructor = 1216809369
		return o
	},
	-618614392: func() TLObject { // 0xdb20b188
		o := MakeTLPageBlockDivider(nil)
		o.Data2.Constructor = -618614392
		return o
	},
	-837994576: func() TLObject { // 0xce0d37b0
		o := MakeTLPageBlockAnchor(nil)
		o.Data2.Constructor = -837994576
		return o
	},
	-454524911: func() TLObject { // 0xe4e88011
		o := MakeTLPageBlockList(nil)
		o.Data2.Constructor = -454524911
		return o
	},
	641563686: func() TLObject { // 0x263d7c26
		o := MakeTLPageBlockBlockquote(nil)
		o.Data2.Constructor = 641563686
		return o
	},
	1329878739: func() TLObject { // 0x4f4456d3
		o := MakeTLPageBlockPullquote(nil)
		o.Data2.Constructor = 1329878739
		return o
	},
	391759200: func() TLObject { // 0x1759c560
		o := MakeTLPageBlockPhoto(nil)
		o.Data2.Constructor = 391759200
		return o
	},
	2089805750: func() TLObject { // 0x7c8fe7b6
		o := MakeTLPageBlockVideo(nil)
		o.Data2.Constructor = 2089805750
		return o
	},
	972174080: func() TLObject { // 0x39f23300
		o := MakeTLPageBlockCover(nil)
		o.Data2.Constructor = 972174080
		return o
	},
	-1468953147: func() TLObject { // 0xa8718dc5
		o := MakeTLPageBlockEmbed(nil)
		o.Data2.Constructor = -1468953147
		return o
	},
	-229005301: func() TLObject { // 0xf259a80b
		o := MakeTLPageBlockEmbedPost(nil)
		o.Data2.Constructor = -229005301
		return o
	},
	1705048653: func() TLObject { // 0x65a0fa4d
		o := MakeTLPageBlockCollage(nil)
		o.Data2.Constructor = 1705048653
		return o
	},
	52401552: func() TLObject { // 0x31f9590
		o := MakeTLPageBlockSlideshow(nil)
		o.Data2.Constructor = 52401552
		return o
	},
	-283684427: func() TLObject { // 0xef1751b5
		o := MakeTLPageBlockChannel(nil)
		o.Data2.Constructor = -283684427
		return o
	},
	-2143067670: func() TLObject { // 0x804361ea
		o := MakeTLPageBlockAudio(nil)
		o.Data2.Constructor = -2143067670
		return o
	},
	504660880: func() TLObject { // 0x1e148390
		o := MakeTLPageBlockKicker(nil)
		o.Data2.Constructor = 504660880
		return o
	},
	-1085412734: func() TLObject { // 0xbf4dea82
		o := MakeTLPageBlockTable(nil)
		o.Data2.Constructor = -1085412734
		return o
	},
	-1702174239: func() TLObject { // 0x9a8ae1e1
		o := MakeTLPageBlockOrderedList(nil)
		o.Data2.Constructor = -1702174239
		return o
	},
	1987480557: func() TLObject { // 0x76768bed
		o := MakeTLPageBlockDetails(nil)
		o.Data2.Constructor = 1987480557
		return o
	},
	370236054: func() TLObject { // 0x16115a96
		o := MakeTLPageBlockRelatedArticles(nil)
		o.Data2.Constructor = 370236054
		return o
	},
	-1538310410: func() TLObject { // 0xa44f3ef6
		o := MakeTLPageBlockMap(nil)
		o.Data2.Constructor = -1538310410
		return o
	},
	1869903447: func() TLObject { // 0x6f747657
		o := MakeTLPageCaption(nil)
		o.Data2.Constructor = 1869903447
		return o
	},
	-1188055347: func() TLObject { // 0xb92fb6cd
		o := MakeTLPageListItemText(nil)
		o.Data2.Constructor = -1188055347
		return o
	},
	635466748: func() TLObject { // 0x25e073fc
		o := MakeTLPageListItemBlocks(nil)
		o.Data2.Constructor = 635466748
		return o
	},
	1577484359: func() TLObject { // 0x5e068047
		o := MakeTLPageListOrderedItemText(nil)
		o.Data2.Constructor = 1577484359
		return o
	},
	-1730311882: func() TLObject { // 0x98dd8936
		o := MakeTLPageListOrderedItemBlocks(nil)
		o.Data2.Constructor = -1730311882
		return o
	},
	-1282352120: func() TLObject { // 0xb390dc08
		o := MakeTLPageRelatedArticle(nil)
		o.Data2.Constructor = -1282352120
		return o
	},
	878078826: func() TLObject { // 0x34566b6a
		o := MakeTLPageTableCell(nil)
		o.Data2.Constructor = 878078826
		return o
	},
	-524237339: func() TLObject { // 0xe0c0c5e5
		o := MakeTLPageTableRow(nil)
		o.Data2.Constructor = -524237339
		return o
	},
	-732254058: func() TLObject { // 0xd45ab096
		o := MakeTLPasswordKdfAlgoUnknown(nil)
		o.Data2.Constructor = -732254058
		return o
	},
	982592842: func() TLObject { // 0x3a912d4a
		o := MakeTLPasswordKdfAlgoModPow(nil)
		o.Data2.Constructor = 982592842
		return o
	},
	-368917890: func() TLObject { // 0xea02c27e
		o := MakeTLPaymentCharge(nil)
		o.Data2.Constructor = -368917890
		return o
	},
	-1996951013: func() TLObject { // 0x88f8f21b
		o := MakeTLPaymentFormMethod(nil)
		o.Data2.Constructor = -1996951013
		return o
	},
	-1868808300: func() TLObject { // 0x909c3f94
		o := MakeTLPaymentRequestedInfo(nil)
		o.Data2.Constructor = -1868808300
		return o
	},
	-842892769: func() TLObject { // 0xcdc27a1f
		o := MakeTLPaymentSavedCredentialsCard(nil)
		o.Data2.Constructor = -842892769
		return o
	},
	1042605427: func() TLObject { // 0x3e24e573
		o := MakeTLPaymentsBankCardData(nil)
		o.Data2.Constructor = 1042605427
		return o
	},
	675942550: func() TLObject { // 0x284a1096
		o := MakeTLPaymentsCheckedGiftCode(nil)
		o.Data2.Constructor = 675942550
		return o
	},
	-1222446760: func() TLObject { // 0xb722f158
		o := MakeTLPaymentsCheckedGiftCode(nil)
		o.Data2.Constructor = -1222446760
		return o
	},
	-1362048039: func() TLObject { // 0xaed0cbd9
		o := MakeTLPaymentsExportedInvoice(nil)
		o.Data2.Constructor = -1362048039
		return o
	},
	1130879648: func() TLObject { // 0x4367daa0
		o := MakeTLPaymentsGiveawayInfo(nil)
		o.Data2.Constructor = 1130879648
		return o
	},
	13456752: func() TLObject { // 0xcd5570
		o := MakeTLPaymentsGiveawayInfoResults(nil)
		o.Data2.Constructor = 13456752
		return o
	},
	-1610250415: func() TLObject { // 0xa0058751
		o := MakeTLPaymentsPaymentForm(nil)
		o.Data2.Constructor = -1610250415
		return o
	},
	-1340916937: func() TLObject { // 0xb0133b37
		o := MakeTLPaymentsPaymentForm(nil)
		o.Data2.Constructor = -1340916937
		return o
	},
	378828315: func() TLObject { // 0x1694761b
		o := MakeTLPaymentsPaymentForm(nil)
		o.Data2.Constructor = 378828315
		return o
	},
	2079764828: func() TLObject { // 0x7bf6b15c
		o := MakeTLPaymentsPaymentFormStars(nil)
		o.Data2.Constructor = 2079764828
		return o
	},
	1891958275: func() TLObject { // 0x70c4fe03
		o := MakeTLPaymentsPaymentReceipt(nil)
		o.Data2.Constructor = 1891958275
		return o
	},
	-625215430: func() TLObject { // 0xdabbf83a
		o := MakeTLPaymentsPaymentReceiptStars(nil)
		o.Data2.Constructor = -625215430
		return o
	},
	1314881805: func() TLObject { // 0x4e5f810d
		o := MakeTLPaymentsPaymentResult(nil)
		o.Data2.Constructor = 1314881805
		return o
	},
	-666824391: func() TLObject { // 0xd8411139
		o := MakeTLPaymentsPaymentVerificationNeeded(nil)
		o.Data2.Constructor = -666824391
		return o
	},
	-74456004: func() TLObject { // 0xfb8fe43c
		o := MakeTLPaymentsSavedInfo(nil)
		o.Data2.Constructor = -74456004
		return o
	},
	961445665: func() TLObject { // 0x394e7f21
		o := MakeTLPaymentsStarsRevenueAdsAccountUrl(nil)
		o.Data2.Constructor = 961445665
		return o
	},
	-919881925: func() TLObject { // 0xc92bb73b
		o := MakeTLPaymentsStarsRevenueStats(nil)
		o.Data2.Constructor = -919881925
		return o
	},
	497778871: func() TLObject { // 0x1dab80b7
		o := MakeTLPaymentsStarsRevenueWithdrawalUrl(nil)
		o.Data2.Constructor = 497778871
		return o
	},
	-1930105248: func() TLObject { // 0x8cf4ee60
		o := MakeTLPaymentsStarsStatus(nil)
		o.Data2.Constructor = -1930105248
		return o
	},
	-784000893: func() TLObject { // 0xd1451883
		o := MakeTLPaymentsValidatedRequestedInfo(nil)
		o.Data2.Constructor = -784000893
		return o
	},
	1498486562: func() TLObject { // 0x59511722
		o := MakeTLPeerUser(nil)
		o.Data2.Constructor = 1498486562
		return o
	},
	918946202: func() TLObject { // 0x36c6019a
		o := MakeTLPeerChat(nil)
		o.Data2.Constructor = 918946202
		return o
	},
	-1566230754: func() TLObject { // 0xa2a5371e
		o := MakeTLPeerChannel(nil)
		o.Data2.Constructor = -1566230754
		return o
	},
	-386039788: func() TLObject { // 0xe8fd8014
		o := MakeTLPeerBlocked(nil)
		o.Data2.Constructor = -386039788
		return o
	},
	-1253352753: func() TLObject { // 0xb54b5acf
		o := MakeTLPeerColor(nil)
		o.Data2.Constructor = -1253352753
		return o
	},
	-901375139: func() TLObject { // 0xca461b5d
		o := MakeTLPeerLocated(nil)
		o.Data2.Constructor = -901375139
		return o
	},
	-118740917: func() TLObject { // 0xf8ec284b
		o := MakeTLPeerSelfLocated(nil)
		o.Data2.Constructor = -118740917
		return o
	},
	-1721619444: func() TLObject { // 0x99622c0c
		o := MakeTLPeerNotifySettings(nil)
		o.Data2.Constructor = -1721619444
		return o
	},
	-1472527322: func() TLObject { // 0xa83b0426
		o := MakeTLPeerNotifySettings(nil)
		o.Data2.Constructor = -1472527322
		return o
	},
	-1353671392: func() TLObject { // 0xaf509d20
		o := MakeTLPeerNotifySettings(nil)
		o.Data2.Constructor = -1353671392
		return o
	},
	-1395233698: func() TLObject { // 0xacd66c5e
		o := MakeTLPeerSettings(nil)
		o.Data2.Constructor = -1395233698
		return o
	},
	-1525149427: func() TLObject { // 0xa518110d
		o := MakeTLPeerSettings(nil)
		o.Data2.Constructor = -1525149427
		return o
	},
	-1707742823: func() TLObject { // 0x9a35e999
		o := MakeTLPeerStories(nil)
		o.Data2.Constructor = -1707742823
		return o
	},
	602876837: func() TLObject { // 0x23ef2ba5
		o := MakeTLPeerUtil(nil)
		o.Data2.Constructor = 602876837
		return o
	},
	1399245077: func() TLObject { // 0x5366c915
		o := MakeTLPhoneCallEmpty(nil)
		o.Data2.Constructor = 1399245077
		return o
	},
	-987599081: func() TLObject { // 0xc5226f17
		o := MakeTLPhoneCallWaiting(nil)
		o.Data2.Constructor = -987599081
		return o
	},
	347139340: func() TLObject { // 0x14b0ed0c
		o := MakeTLPhoneCallRequested(nil)
		o.Data2.Constructor = 347139340
		return o
	},
	912311057: func() TLObject { // 0x3660c311
		o := MakeTLPhoneCallAccepted(nil)
		o.Data2.Constructor = 912311057
		return o
	},
	810769141: func() TLObject { // 0x30535af5
		o := MakeTLPhoneCall(nil)
		o.Data2.Constructor = 810769141
		return o
	},
	-1770029977: func() TLObject { // 0x967f7c67
		o := MakeTLPhoneCall(nil)
		o.Data2.Constructor = -1770029977
		return o
	},
	1355435489: func() TLObject { // 0x50ca4de1
		o := MakeTLPhoneCallDiscarded(nil)
		o.Data2.Constructor = 1355435489
		return o
	},
	-2048646399: func() TLObject { // 0x85e42301
		o := MakeTLPhoneCallDiscardReasonMissed(nil)
		o.Data2.Constructor = -2048646399
		return o
	},
	-527056480: func() TLObject { // 0xe095c1a0
		o := MakeTLPhoneCallDiscardReasonDisconnect(nil)
		o.Data2.Constructor = -527056480
		return o
	},
	1471006352: func() TLObject { // 0x57adc690
		o := MakeTLPhoneCallDiscardReasonHangup(nil)
		o.Data2.Constructor = 1471006352
		return o
	},
	-84416311: func() TLObject { // 0xfaf7e8c9
		o := MakeTLPhoneCallDiscardReasonBusy(nil)
		o.Data2.Constructor = -84416311
		return o
	},
	-58224696: func() TLObject { // 0xfc878fc8
		o := MakeTLPhoneCallProtocol(nil)
		o.Data2.Constructor = -58224696
		return o
	},
	-1665063993: func() TLObject { // 0x9cc123c7
		o := MakeTLPhoneConnection(nil)
		o.Data2.Constructor = -1665063993
		return o
	},
	-1655957568: func() TLObject { // 0x9d4c17c0
		o := MakeTLPhoneConnection(nil)
		o.Data2.Constructor = -1655957568
		return o
	},
	1667228533: func() TLObject { // 0x635fe375
		o := MakeTLPhoneConnectionWebrtc(nil)
		o.Data2.Constructor = 1667228533
		return o
	},
	541839704: func() TLObject { // 0x204bd158
		o := MakeTLPhoneExportedGroupCallInvite(nil)
		o.Data2.Constructor = 541839704
		return o
	},
	-1636664659: func() TLObject { // 0x9e727aad
		o := MakeTLPhoneGroupCall(nil)
		o.Data2.Constructor = -1636664659
		return o
	},
	-790330702: func() TLObject { // 0xd0e482b2
		o := MakeTLPhoneGroupCallStreamChannels(nil)
		o.Data2.Constructor = -790330702
		return o
	},
	767505458: func() TLObject { // 0x2dbf3432
		o := MakeTLPhoneGroupCallStreamRtmpUrl(nil)
		o.Data2.Constructor = 767505458
		return o
	},
	-193506890: func() TLObject { // 0xf47751b6
		o := MakeTLPhoneGroupParticipants(nil)
		o.Data2.Constructor = -193506890
		return o
	},
	-1343921601: func() TLObject { // 0xafe5623f
		o := MakeTLPhoneJoinAsPeers(nil)
		o.Data2.Constructor = -1343921601
		return o
	},
	-326966976: func() TLObject { // 0xec82e140
		o := MakeTLPhonePhoneCall(nil)
		o.Data2.Constructor = -326966976
		return o
	},
	590459437: func() TLObject { // 0x2331b22d
		o := MakeTLPhotoEmpty(nil)
		o.Data2.Constructor = 590459437
		return o
	},
	-82216347: func() TLObject { // 0xfb197a65
		o := MakeTLPhoto(nil)
		o.Data2.Constructor = -82216347
		return o
	},
	236446268: func() TLObject { // 0xe17e23c
		o := MakeTLPhotoSizeEmpty(nil)
		o.Data2.Constructor = 236446268
		return o
	},
	1976012384: func() TLObject { // 0x75c78e60
		o := MakeTLPhotoSize(nil)
		o.Data2.Constructor = 1976012384
		return o
	},
	35527382: func() TLObject { // 0x21e1ad6
		o := MakeTLPhotoCachedSize(nil)
		o.Data2.Constructor = 35527382
		return o
	},
	-525288402: func() TLObject { // 0xe0b0bc2e
		o := MakeTLPhotoStrippedSize(nil)
		o.Data2.Constructor = -525288402
		return o
	},
	-96535659: func() TLObject { // 0xfa3efb95
		o := MakeTLPhotoSizeProgressive(nil)
		o.Data2.Constructor = -96535659
		return o
	},
	-668906175: func() TLObject { // 0xd8214d41
		o := MakeTLPhotoPathSize(nil)
		o.Data2.Constructor = -668906175
		return o
	},
	539045032: func() TLObject { // 0x20212ca8
		o := MakeTLPhotosPhoto(nil)
		o.Data2.Constructor = 539045032
		return o
	},
	-1916114267: func() TLObject { // 0x8dca6aa5
		o := MakeTLPhotosPhotos(nil)
		o.Data2.Constructor = -1916114267
		return o
	},
	352657236: func() TLObject { // 0x15051f54
		o := MakeTLPhotosPhotosSlice(nil)
		o.Data2.Constructor = 352657236
		return o
	},
	1484026161: func() TLObject { // 0x58747131
		o := MakeTLPoll(nil)
		o.Data2.Constructor = 1484026161
		return o
	},
	-2032041631: func() TLObject { // 0x86e18161
		o := MakeTLPoll(nil)
		o.Data2.Constructor = -2032041631
		return o
	},
	-15277366: func() TLObject { // 0xff16e2ca
		o := MakeTLPollAnswer(nil)
		o.Data2.Constructor = -15277366
		return o
	},
	1823064809: func() TLObject { // 0x6ca9c2e9
		o := MakeTLPollAnswer(nil)
		o.Data2.Constructor = 1823064809
		return o
	},
	997055186: func() TLObject { // 0x3b6ddad2
		o := MakeTLPollAnswerVoters(nil)
		o.Data2.Constructor = 997055186
		return o
	},
	2061444128: func() TLObject { // 0x7adf2420
		o := MakeTLPollResults(nil)
		o.Data2.Constructor = 2061444128
		return o
	},
	-591909213: func() TLObject { // 0xdcb82ea3
		o := MakeTLPollResults(nil)
		o.Data2.Constructor = -591909213
		return o
	},
	1558266229: func() TLObject { // 0x5ce14175
		o := MakeTLPopularContact(nil)
		o.Data2.Constructor = 1558266229
		return o
	},
	512535275: func() TLObject { // 0x1e8caaeb
		o := MakeTLPostAddress(nil)
		o.Data2.Constructor = 512535275
		return o
	},
	-419066241: func() TLObject { // 0xe7058e7f
		o := MakeTLPostInteractionCountersMessage(nil)
		o.Data2.Constructor = -419066241
		return o
	},
	-1974989273: func() TLObject { // 0x8a480e27
		o := MakeTLPostInteractionCountersStory(nil)
		o.Data2.Constructor = -1974989273
		return o
	},
	383118531: func() TLObject { // 0x16d5ecc3
		o := MakeTLPredefinedUser(nil)
		o.Data2.Constructor = 383118531
		return o
	},
	629052971: func() TLObject { // 0x257e962b
		o := MakeTLPremiumGiftCodeOption(nil)
		o.Data2.Constructor = 629052971
		return o
	},
	1958953753: func() TLObject { // 0x74c34319
		o := MakeTLPremiumGiftOption(nil)
		o.Data2.Constructor = 1958953753
		return o
	},
	1596792306: func() TLObject { // 0x5f2d1df2
		o := MakeTLPremiumSubscriptionOption(nil)
		o.Data2.Constructor = 1596792306
		return o
	},
	-1225711938: func() TLObject { // 0xb6f11ebe
		o := MakeTLPremiumSubscriptionOption(nil)
		o.Data2.Constructor = -1225711938
		return o
	},
	-2030542532: func() TLObject { // 0x86f8613c
		o := MakeTLPremiumBoostsList(nil)
		o.Data2.Constructor = -2030542532
		return o
	},
	1230586490: func() TLObject { // 0x4959427a
		o := MakeTLPremiumBoostsStatus(nil)
		o.Data2.Constructor = 1230586490
		return o
	},
	-1696454430: func() TLObject { // 0x9ae228e2
		o := MakeTLPremiumMyBoosts(nil)
		o.Data2.Constructor = -1696454430
		return o
	},
	-1303143084: func() TLObject { // 0xb2539d54
		o := MakeTLPrepaidGiveaway(nil)
		o.Data2.Constructor = -1303143084
		return o
	},
	-1137792208: func() TLObject { // 0xbc2eab30
		o := MakeTLPrivacyKeyStatusTimestamp(nil)
		o.Data2.Constructor = -1137792208
		return o
	},
	1343122938: func() TLObject { // 0x500e6dfa
		o := MakeTLPrivacyKeyChatInvite(nil)
		o.Data2.Constructor = 1343122938
		return o
	},
	1030105979: func() TLObject { // 0x3d662b7b
		o := MakeTLPrivacyKeyPhoneCall(nil)
		o.Data2.Constructor = 1030105979
		return o
	},
	961092808: func() TLObject { // 0x39491cc8
		o := MakeTLPrivacyKeyPhoneP2P(nil)
		o.Data2.Constructor = 961092808
		return o
	},
	1777096355: func() TLObject { // 0x69ec56a3
		o := MakeTLPrivacyKeyForwards(nil)
		o.Data2.Constructor = 1777096355
		return o
	},
	-1777000467: func() TLObject { // 0x96151fed
		o := MakeTLPrivacyKeyProfilePhoto(nil)
		o.Data2.Constructor = -1777000467
		return o
	},
	-778378131: func() TLObject { // 0xd19ae46d
		o := MakeTLPrivacyKeyPhoneNumber(nil)
		o.Data2.Constructor = -778378131
		return o
	},
	1124062251: func() TLObject { // 0x42ffd42b
		o := MakeTLPrivacyKeyAddedByPhone(nil)
		o.Data2.Constructor = 1124062251
		return o
	},
	110621716: func() TLObject { // 0x697f414
		o := MakeTLPrivacyKeyVoiceMessages(nil)
		o.Data2.Constructor = 110621716
		return o
	},
	-1534675103: func() TLObject { // 0xa486b761
		o := MakeTLPrivacyKeyAbout(nil)
		o.Data2.Constructor = -1534675103
		return o
	},
	536913176: func() TLObject { // 0x2000a518
		o := MakeTLPrivacyKeyBirthday(nil)
		o.Data2.Constructor = 536913176
		return o
	},
	-1810715178: func() TLObject { // 0x9412add6
		o := MakeTLPrivacyKeyRules(nil)
		o.Data2.Constructor = -1810715178
		return o
	},
	-123988: func() TLObject { // 0xfffe1bac
		o := MakeTLPrivacyValueAllowContacts(nil)
		o.Data2.Constructor = -123988
		return o
	},
	1698855810: func() TLObject { // 0x65427b82
		o := MakeTLPrivacyValueAllowAll(nil)
		o.Data2.Constructor = 1698855810
		return o
	},
	-1198497870: func() TLObject { // 0xb8905fb2
		o := MakeTLPrivacyValueAllowUsers(nil)
		o.Data2.Constructor = -1198497870
		return o
	},
	-125240806: func() TLObject { // 0xf888fa1a
		o := MakeTLPrivacyValueDisallowContacts(nil)
		o.Data2.Constructor = -125240806
		return o
	},
	-1955338397: func() TLObject { // 0x8b73e763
		o := MakeTLPrivacyValueDisallowAll(nil)
		o.Data2.Constructor = -1955338397
		return o
	},
	-463335103: func() TLObject { // 0xe4621141
		o := MakeTLPrivacyValueDisallowUsers(nil)
		o.Data2.Constructor = -463335103
		return o
	},
	1796427406: func() TLObject { // 0x6b134e8e
		o := MakeTLPrivacyValueAllowChatParticipants(nil)
		o.Data2.Constructor = 1796427406
		return o
	},
	1103656293: func() TLObject { // 0x41c87565
		o := MakeTLPrivacyValueDisallowChatParticipants(nil)
		o.Data2.Constructor = 1103656293
		return o
	},
	-135735141: func() TLObject { // 0xf7e8d89b
		o := MakeTLPrivacyValueAllowCloseFriends(nil)
		o.Data2.Constructor = -135735141
		return o
	},
	-320241333: func() TLObject { // 0xece9814b
		o := MakeTLPrivacyValueAllowPremium(nil)
		o.Data2.Constructor = -320241333
		return o
	},
	32685898: func() TLObject { // 0x1f2bf4a
		o := MakeTLPublicForwardMessage(nil)
		o.Data2.Constructor = 32685898
		return o
	},
	-302797360: func() TLObject { // 0xedf3add0
		o := MakeTLPublicForwardStory(nil)
		o.Data2.Constructor = -302797360
		return o
	},
	110563371: func() TLObject { // 0x697102b
		o := MakeTLQuickReply(nil)
		o.Data2.Constructor = 110563371
		return o
	},
	2046153753: func() TLObject { // 0x79f5d419
		o := MakeTLReactionEmpty(nil)
		o.Data2.Constructor = 2046153753
		return o
	},
	455247544: func() TLObject { // 0x1b2286b8
		o := MakeTLReactionEmoji(nil)
		o.Data2.Constructor = 455247544
		return o
	},
	-1992950669: func() TLObject { // 0x8935fc73
		o := MakeTLReactionCustomEmoji(nil)
		o.Data2.Constructor = -1992950669
		return o
	},
	-1546531968: func() TLObject { // 0xa3d1cb80
		o := MakeTLReactionCount(nil)
		o.Data2.Constructor = -1546531968
		return o
	},
	1873957073: func() TLObject { // 0x6fb250d1
		o := MakeTLReactionCount(nil)
		o.Data2.Constructor = 1873957073
		return o
	},
	-1161583078: func() TLObject { // 0xbac3a61a
		o := MakeTLReactionNotificationsFromContacts(nil)
		o.Data2.Constructor = -1161583078
		return o
	},
	1268654752: func() TLObject { // 0x4b9e22a0
		o := MakeTLReactionNotificationsFromAll(nil)
		o.Data2.Constructor = 1268654752
		return o
	},
	1457736048: func() TLObject { // 0x56e34970
		o := MakeTLReactionsNotifySettings(nil)
		o.Data2.Constructor = 1457736048
		return o
	},
	1246753138: func() TLObject { // 0x4a4ff172
		o := MakeTLReadParticipantDate(nil)
		o.Data2.Constructor = 1246753138
		return o
	},
	-1551583367: func() TLObject { // 0xa384b779
		o := MakeTLReceivedNotifyMessage(nil)
		o.Data2.Constructor = -1551583367
		return o
	},
	1189204285: func() TLObject { // 0x46e1d13d
		o := MakeTLRecentMeUrlUnknown(nil)
		o.Data2.Constructor = 1189204285
		return o
	},
	-1188296222: func() TLObject { // 0xb92c09e2
		o := MakeTLRecentMeUrlUser(nil)
		o.Data2.Constructor = -1188296222
		return o
	},
	-1294306862: func() TLObject { // 0xb2da71d2
		o := MakeTLRecentMeUrlChat(nil)
		o.Data2.Constructor = -1294306862
		return o
	},
	-347535331: func() TLObject { // 0xeb49081d
		o := MakeTLRecentMeUrlChatInvite(nil)
		o.Data2.Constructor = -347535331
		return o
	},
	-1140172836: func() TLObject { // 0xbc0a57dc
		o := MakeTLRecentMeUrlStickerSet(nil)
		o.Data2.Constructor = -1140172836
		return o
	},
	-1606526075: func() TLObject { // 0xa03e5b85
		o := MakeTLReplyKeyboardHide(nil)
		o.Data2.Constructor = -1606526075
		return o
	},
	-2035021048: func() TLObject { // 0x86b40b08
		o := MakeTLReplyKeyboardForceReply(nil)
		o.Data2.Constructor = -2035021048
		return o
	},
	-2049074735: func() TLObject { // 0x85dd99d1
		o := MakeTLReplyKeyboardMarkup(nil)
		o.Data2.Constructor = -2049074735
		return o
	},
	1218642516: func() TLObject { // 0x48a30254
		o := MakeTLReplyInlineMarkup(nil)
		o.Data2.Constructor = 1218642516
		return o
	},
	1490799288: func() TLObject { // 0x58dbcab8
		o := MakeTLInputReportReasonSpam(nil)
		o.Data2.Constructor = 1490799288
		return o
	},
	505595789: func() TLObject { // 0x1e22c78d
		o := MakeTLInputReportReasonViolence(nil)
		o.Data2.Constructor = 505595789
		return o
	},
	777640226: func() TLObject { // 0x2e59d922
		o := MakeTLInputReportReasonPornography(nil)
		o.Data2.Constructor = 777640226
		return o
	},
	-1376497949: func() TLObject { // 0xadf44ee3
		o := MakeTLInputReportReasonChildAbuse(nil)
		o.Data2.Constructor = -1376497949
		return o
	},
	-1041980751: func() TLObject { // 0xc1e4a2b1
		o := MakeTLInputReportReasonOther(nil)
		o.Data2.Constructor = -1041980751
		return o
	},
	-1685456582: func() TLObject { // 0x9b89f93a
		o := MakeTLInputReportReasonCopyright(nil)
		o.Data2.Constructor = -1685456582
		return o
	},
	-606798099: func() TLObject { // 0xdbd4feed
		o := MakeTLInputReportReasonGeoIrrelevant(nil)
		o.Data2.Constructor = -606798099
		return o
	},
	-170010905: func() TLObject { // 0xf5ddd6e7
		o := MakeTLInputReportReasonFake(nil)
		o.Data2.Constructor = -170010905
		return o
	},
	177124030: func() TLObject { // 0xa8eb2be
		o := MakeTLInputReportReasonIllegalDrugs(nil)
		o.Data2.Constructor = 177124030
		return o
	},
	-1631091139: func() TLObject { // 0x9ec7863d
		o := MakeTLInputReportReasonPersonalDetails(nil)
		o.Data2.Constructor = -1631091139
		return o
	},
	1597737472: func() TLObject { // 0x5f3b8a00
		o := MakeTLRequestPeerTypeUser(nil)
		o.Data2.Constructor = 1597737472
		return o
	},
	-906990053: func() TLObject { // 0xc9f06e1b
		o := MakeTLRequestPeerTypeChat(nil)
		o.Data2.Constructor = -906990053
		return o
	},
	865857388: func() TLObject { // 0x339bef6c
		o := MakeTLRequestPeerTypeBroadcast(nil)
		o.Data2.Constructor = 865857388
		return o
	},
	-701500310: func() TLObject { // 0xd62ff46a
		o := MakeTLRequestedPeerUser(nil)
		o.Data2.Constructor = -701500310
		return o
	},
	1929860175: func() TLObject { // 0x7307544f
		o := MakeTLRequestedPeerChat(nil)
		o.Data2.Constructor = 1929860175
		return o
	},
	-1952185372: func() TLObject { // 0x8ba403e4
		o := MakeTLRequestedPeerChannel(nil)
		o.Data2.Constructor = -1952185372
		return o
	},
	-797791052: func() TLObject { // 0xd072acb4
		o := MakeTLRestrictionReason(nil)
		o.Data2.Constructor = -797791052
		return o
	},
	-599948721: func() TLObject { // 0xdc3d824f
		o := MakeTLTextEmpty(nil)
		o.Data2.Constructor = -599948721
		return o
	},
	1950782688: func() TLObject { // 0x744694e0
		o := MakeTLTextPlain(nil)
		o.Data2.Constructor = 1950782688
		return o
	},
	1730456516: func() TLObject { // 0x6724abc4
		o := MakeTLTextBold(nil)
		o.Data2.Constructor = 1730456516
		return o
	},
	-653089380: func() TLObject { // 0xd912a59c
		o := MakeTLTextItalic(nil)
		o.Data2.Constructor = -653089380
		return o
	},
	-1054465340: func() TLObject { // 0xc12622c4
		o := MakeTLTextUnderline(nil)
		o.Data2.Constructor = -1054465340
		return o
	},
	-1678197867: func() TLObject { // 0x9bf8bb95
		o := MakeTLTextStrike(nil)
		o.Data2.Constructor = -1678197867
		return o
	},
	1816074681: func() TLObject { // 0x6c3f19b9
		o := MakeTLTextFixed(nil)
		o.Data2.Constructor = 1816074681
		return o
	},
	1009288385: func() TLObject { // 0x3c2884c1
		o := MakeTLTextUrl(nil)
		o.Data2.Constructor = 1009288385
		return o
	},
	-564523562: func() TLObject { // 0xde5a0dd6
		o := MakeTLTextEmail(nil)
		o.Data2.Constructor = -564523562
		return o
	},
	2120376535: func() TLObject { // 0x7e6260d7
		o := MakeTLTextConcat(nil)
		o.Data2.Constructor = 2120376535
		return o
	},
	-311786236: func() TLObject { // 0xed6a8504
		o := MakeTLTextSubscript(nil)
		o.Data2.Constructor = -311786236
		return o
	},
	-939827711: func() TLObject { // 0xc7fb5e01
		o := MakeTLTextSuperscript(nil)
		o.Data2.Constructor = -939827711
		return o
	},
	55281185: func() TLObject { // 0x34b8621
		o := MakeTLTextMarked(nil)
		o.Data2.Constructor = 55281185
		return o
	},
	483104362: func() TLObject { // 0x1ccb966a
		o := MakeTLTextPhone(nil)
		o.Data2.Constructor = 483104362
		return o
	},
	136105807: func() TLObject { // 0x81ccf4f
		o := MakeTLTextImage(nil)
		o.Data2.Constructor = 136105807
		return o
	},
	894777186: func() TLObject { // 0x35553762
		o := MakeTLTextAnchor(nil)
		o.Data2.Constructor = 894777186
		return o
	},
	289586518: func() TLObject { // 0x1142bd56
		o := MakeTLSavedPhoneContact(nil)
		o.Data2.Constructor = 289586518
		return o
	},
	-1115174036: func() TLObject { // 0xbd87cb6c
		o := MakeTLSavedDialog(nil)
		o.Data2.Constructor = -1115174036
		return o
	},
	-881854424: func() TLObject { // 0xcb6ff828
		o := MakeTLSavedReactionTag(nil)
		o.Data2.Constructor = -881854424
		return o
	},
	-911191137: func() TLObject { // 0xc9b0539f
		o := MakeTLSearchResultsCalendarPeriod(nil)
		o.Data2.Constructor = -911191137
		return o
	},
	2137295719: func() TLObject { // 0x7f648b67
		o := MakeTLSearchResultPosition(nil)
		o.Data2.Constructor = 2137295719
		return o
	},
	871426631: func() TLObject { // 0x33f0ea47
		o := MakeTLSecureCredentialsEncrypted(nil)
		o.Data2.Constructor = 871426631
		return o
	},
	-1964327229: func() TLObject { // 0x8aeabec3
		o := MakeTLSecureData(nil)
		o.Data2.Constructor = -1964327229
		return o
	},
	1679398724: func() TLObject { // 0x64199744
		o := MakeTLSecureFileEmpty(nil)
		o.Data2.Constructor = 1679398724
		return o
	},
	2097791614: func() TLObject { // 0x7d09c27e
		o := MakeTLSecureFile(nil)
		o.Data2.Constructor = 2097791614
		return o
	},
	-534283678: func() TLObject { // 0xe0277a62
		o := MakeTLSecureFile(nil)
		o.Data2.Constructor = -534283678
		return o
	},
	4883767: func() TLObject { // 0x4a8537
		o := MakeTLSecurePasswordKdfAlgoUnknown(nil)
		o.Data2.Constructor = 4883767
		return o
	},
	-1141711456: func() TLObject { // 0xbbf2dda0
		o := MakeTLSecurePasswordKdfAlgoPBKDF2(nil)
		o.Data2.Constructor = -1141711456
		return o
	},
	-2042159726: func() TLObject { // 0x86471d92
		o := MakeTLSecurePasswordKdfAlgoSHA512(nil)
		o.Data2.Constructor = -2042159726
		return o
	},
	2103482845: func() TLObject { // 0x7d6099dd
		o := MakeTLSecurePlainPhone(nil)
		o.Data2.Constructor = 2103482845
		return o
	},
	569137759: func() TLObject { // 0x21ec5a5f
		o := MakeTLSecurePlainEmail(nil)
		o.Data2.Constructor = 569137759
		return o
	},
	-2103600678: func() TLObject { // 0x829d99da
		o := MakeTLSecureRequiredType(nil)
		o.Data2.Constructor = -2103600678
		return o
	},
	41187252: func() TLObject { // 0x27477b4
		o := MakeTLSecureRequiredTypeOneOf(nil)
		o.Data2.Constructor = 41187252
		return o
	},
	354925740: func() TLObject { // 0x1527bcac
		o := MakeTLSecureSecretSettings(nil)
		o.Data2.Constructor = 354925740
		return o
	},
	411017418: func() TLObject { // 0x187fa0ca
		o := MakeTLSecureValue(nil)
		o.Data2.Constructor = 411017418
		return o
	},
	-391902247: func() TLObject { // 0xe8a40bd9
		o := MakeTLSecureValueErrorData(nil)
		o.Data2.Constructor = -391902247
		return o
	},
	12467706: func() TLObject { // 0xbe3dfa
		o := MakeTLSecureValueErrorFrontSide(nil)
		o.Data2.Constructor = 12467706
		return o
	},
	-2037765467: func() TLObject { // 0x868a2aa5
		o := MakeTLSecureValueErrorReverseSide(nil)
		o.Data2.Constructor = -2037765467
		return o
	},
	-449327402: func() TLObject { // 0xe537ced6
		o := MakeTLSecureValueErrorSelfie(nil)
		o.Data2.Constructor = -449327402
		return o
	},
	2054162547: func() TLObject { // 0x7a700873
		o := MakeTLSecureValueErrorFile(nil)
		o.Data2.Constructor = 2054162547
		return o
	},
	1717706985: func() TLObject { // 0x666220e9
		o := MakeTLSecureValueErrorFiles(nil)
		o.Data2.Constructor = 1717706985
		return o
	},
	-2036501105: func() TLObject { // 0x869d758f
		o := MakeTLSecureValueError(nil)
		o.Data2.Constructor = -2036501105
		return o
	},
	-1592506512: func() TLObject { // 0xa1144770
		o := MakeTLSecureValueErrorTranslationFile(nil)
		o.Data2.Constructor = -1592506512
		return o
	},
	878931416: func() TLObject { // 0x34636dd8
		o := MakeTLSecureValueErrorTranslationFiles(nil)
		o.Data2.Constructor = 878931416
		return o
	},
	-316748368: func() TLObject { // 0xed1ecdb0
		o := MakeTLSecureValueHash(nil)
		o.Data2.Constructor = -316748368
		return o
	},
	-1658158621: func() TLObject { // 0x9d2a81e3
		o := MakeTLSecureValueTypePersonalDetails(nil)
		o.Data2.Constructor = -1658158621
		return o
	},
	1034709504: func() TLObject { // 0x3dac6a00
		o := MakeTLSecureValueTypePassport(nil)
		o.Data2.Constructor = 1034709504
		return o
	},
	115615172: func() TLObject { // 0x6e425c4
		o := MakeTLSecureValueTypeDriverLicense(nil)
		o.Data2.Constructor = 115615172
		return o
	},
	-1596951477: func() TLObject { // 0xa0d0744b
		o := MakeTLSecureValueTypeIdentityCard(nil)
		o.Data2.Constructor = -1596951477
		return o
	},
	-1717268701: func() TLObject { // 0x99a48f23
		o := MakeTLSecureValueTypeInternalPassport(nil)
		o.Data2.Constructor = -1717268701
		return o
	},
	-874308058: func() TLObject { // 0xcbe31e26
		o := MakeTLSecureValueTypeAddress(nil)
		o.Data2.Constructor = -874308058
		return o
	},
	-63531698: func() TLObject { // 0xfc36954e
		o := MakeTLSecureValueTypeUtilityBill(nil)
		o.Data2.Constructor = -63531698
		return o
	},
	-1995211763: func() TLObject { // 0x89137c0d
		o := MakeTLSecureValueTypeBankStatement(nil)
		o.Data2.Constructor = -1995211763
		return o
	},
	-1954007928: func() TLObject { // 0x8b883488
		o := MakeTLSecureValueTypeRentalAgreement(nil)
		o.Data2.Constructor = -1954007928
		return o
	},
	-1713143702: func() TLObject { // 0x99e3806a
		o := MakeTLSecureValueTypePassportRegistration(nil)
		o.Data2.Constructor = -1713143702
		return o
	},
	-368907213: func() TLObject { // 0xea02ec33
		o := MakeTLSecureValueTypeTemporaryRegistration(nil)
		o.Data2.Constructor = -368907213
		return o
	},
	-1289704741: func() TLObject { // 0xb320aadb
		o := MakeTLSecureValueTypePhone(nil)
		o.Data2.Constructor = -1289704741
		return o
	},
	-1908627474: func() TLObject { // 0x8e3ca7ee
		o := MakeTLSecureValueTypeEmail(nil)
		o.Data2.Constructor = -1908627474
		return o
	},
	-1206095820: func() TLObject { // 0xb81c7034
		o := MakeTLSendAsPeer(nil)
		o.Data2.Constructor = -1206095820
		return o
	},
	381645902: func() TLObject { // 0x16bf744e
		o := MakeTLSendMessageTypingAction(nil)
		o.Data2.Constructor = 381645902
		return o
	},
	-44119819: func() TLObject { // 0xfd5ec8f5
		o := MakeTLSendMessageCancelAction(nil)
		o.Data2.Constructor = -44119819
		return o
	},
	-1584933265: func() TLObject { // 0xa187d66f
		o := MakeTLSendMessageRecordVideoAction(nil)
		o.Data2.Constructor = -1584933265
		return o
	},
	-378127636: func() TLObject { // 0xe9763aec
		o := MakeTLSendMessageUploadVideoAction(nil)
		o.Data2.Constructor = -378127636
		return o
	},
	-718310409: func() TLObject { // 0xd52f73f7
		o := MakeTLSendMessageRecordAudioAction(nil)
		o.Data2.Constructor = -718310409
		return o
	},
	-212740181: func() TLObject { // 0xf351d7ab
		o := MakeTLSendMessageUploadAudioAction(nil)
		o.Data2.Constructor = -212740181
		return o
	},
	-774682074: func() TLObject { // 0xd1d34a26
		o := MakeTLSendMessageUploadPhotoAction(nil)
		o.Data2.Constructor = -774682074
		return o
	},
	-1441998364: func() TLObject { // 0xaa0cd9e4
		o := MakeTLSendMessageUploadDocumentAction(nil)
		o.Data2.Constructor = -1441998364
		return o
	},
	393186209: func() TLObject { // 0x176f8ba1
		o := MakeTLSendMessageGeoLocationAction(nil)
		o.Data2.Constructor = 393186209
		return o
	},
	1653390447: func() TLObject { // 0x628cbc6f
		o := MakeTLSendMessageChooseContactAction(nil)
		o.Data2.Constructor = 1653390447
		return o
	},
	-580219064: func() TLObject { // 0xdd6a8f48
		o := MakeTLSendMessageGamePlayAction(nil)
		o.Data2.Constructor = -580219064
		return o
	},
	-1997373508: func() TLObject { // 0x88f27fbc
		o := MakeTLSendMessageRecordRoundAction(nil)
		o.Data2.Constructor = -1997373508
		return o
	},
	608050278: func() TLObject { // 0x243e1c66
		o := MakeTLSendMessageUploadRoundAction(nil)
		o.Data2.Constructor = 608050278
		return o
	},
	-651419003: func() TLObject { // 0xd92c2285
		o := MakeTLSpeakingInGroupCallAction(nil)
		o.Data2.Constructor = -651419003
		return o
	},
	-606432698: func() TLObject { // 0xdbda9246
		o := MakeTLSendMessageHistoryImportAction(nil)
		o.Data2.Constructor = -606432698
		return o
	},
	-1336228175: func() TLObject { // 0xb05ac6b1
		o := MakeTLSendMessageChooseStickerAction(nil)
		o.Data2.Constructor = -1336228175
		return o
	},
	630664139: func() TLObject { // 0x25972bcb
		o := MakeTLSendMessageEmojiInteraction(nil)
		o.Data2.Constructor = 630664139
		return o
	},
	-1234857938: func() TLObject { // 0xb665902e
		o := MakeTLSendMessageEmojiInteractionSeen(nil)
		o.Data2.Constructor = -1234857938
		return o
	},
	-1239335713: func() TLObject { // 0xb6213cdf
		o := MakeTLShippingOption(nil)
		o.Data2.Constructor = -1239335713
		return o
	},
	-2010155333: func() TLObject { // 0x882f76bb
		o := MakeTLSimpleWebViewResultUrl(nil)
		o.Data2.Constructor = -2010155333
		return o
	},
	-425595208: func() TLObject { // 0xe6a1eeb8
		o := MakeTLSmsJob(nil)
		o.Data2.Constructor = -425595208
		return o
	},
	-594852657: func() TLObject { // 0xdc8b44cf
		o := MakeTLSmsjobsEligibleToJoin(nil)
		o.Data2.Constructor = -594852657
		return o
	},
	720277905: func() TLObject { // 0x2aee9191
		o := MakeTLSmsjobsStatus(nil)
		o.Data2.Constructor = 720277905
		return o
	},
	-1108478618: func() TLObject { // 0xbdedf566
		o := MakeTLSponsoredMessage(nil)
		o.Data2.Constructor = -1108478618
		return o
	},
	-313293833: func() TLObject { // 0xed5383f7
		o := MakeTLSponsoredMessage(nil)
		o.Data2.Constructor = -313293833
		return o
	},
	-626000021: func() TLObject { // 0xdaafff6b
		o := MakeTLSponsoredMessage(nil)
		o.Data2.Constructor = -626000021
		return o
	},
	-64636888: func() TLObject { // 0xfc25b828
		o := MakeTLSponsoredMessage(nil)
		o.Data2.Constructor = -64636888
		return o
	},
	981691896: func() TLObject { // 0x3a836df8
		o := MakeTLSponsoredMessage(nil)
		o.Data2.Constructor = 981691896
		return o
	},
	1124938064: func() TLObject { // 0x430d3150
		o := MakeTLSponsoredMessageReportOption(nil)
		o.Data2.Constructor = 1124938064
		return o
	},
	1035529315: func() TLObject { // 0x3db8ec63
		o := MakeTLSponsoredWebPage(nil)
		o.Data2.Constructor = 1035529315
		return o
	},
	2033461574: func() TLObject { // 0x79342946
		o := MakeTLStarsRevenueStatus(nil)
		o.Data2.Constructor = 2033461574
		return o
	},
	198776256: func() TLObject { // 0xbd915c0
		o := MakeTLStarsTopupOption(nil)
		o.Data2.Constructor = 198776256
		return o
	},
	766853519: func() TLObject { // 0x2db5418f
		o := MakeTLStarsTransaction(nil)
		o.Data2.Constructor = 766853519
		return o
	},
	-1442789224: func() TLObject { // 0xaa00c898
		o := MakeTLStarsTransaction(nil)
		o.Data2.Constructor = -1442789224
		return o
	},
	-865044046: func() TLObject { // 0xcc7079b2
		o := MakeTLStarsTransaction(nil)
		o.Data2.Constructor = -865044046
		return o
	},
	-1779253276: func() TLObject { // 0x95f2bfe4
		o := MakeTLStarsTransactionPeerUnsupported(nil)
		o.Data2.Constructor = -1779253276
		return o
	},
	-1269320843: func() TLObject { // 0xb457b375
		o := MakeTLStarsTransactionPeerAppStore(nil)
		o.Data2.Constructor = -1269320843
		return o
	},
	2069236235: func() TLObject { // 0x7b560a0b
		o := MakeTLStarsTransactionPeerPlayMarket(nil)
		o.Data2.Constructor = 2069236235
		return o
	},
	621656824: func() TLObject { // 0x250dbaf8
		o := MakeTLStarsTransactionPeerPremiumBot(nil)
		o.Data2.Constructor = 621656824
		return o
	},
	-382740222: func() TLObject { // 0xe92fd902
		o := MakeTLStarsTransactionPeerFragment(nil)
		o.Data2.Constructor = -382740222
		return o
	},
	-670195363: func() TLObject { // 0xd80da15d
		o := MakeTLStarsTransactionPeer(nil)
		o.Data2.Constructor = -670195363
		return o
	},
	1617438738: func() TLObject { // 0x60682812
		o := MakeTLStarsTransactionPeerAds(nil)
		o.Data2.Constructor = 1617438738
		return o
	},
	-884757282: func() TLObject { // 0xcb43acde
		o := MakeTLStatsAbsValueAndPrev(nil)
		o.Data2.Constructor = -884757282
		return o
	},
	-1237848657: func() TLObject { // 0xb637edaf
		o := MakeTLStatsDateRangeDays(nil)
		o.Data2.Constructor = -1237848657
		return o
	},
	1244130093: func() TLObject { // 0x4a27eb2d
		o := MakeTLStatsGraphAsync(nil)
		o.Data2.Constructor = 1244130093
		return o
	},
	-1092839390: func() TLObject { // 0xbedc9822
		o := MakeTLStatsGraphError(nil)
		o.Data2.Constructor = -1092839390
		return o
	},
	-1901828938: func() TLObject { // 0x8ea464b6
		o := MakeTLStatsGraph(nil)
		o.Data2.Constructor = -1901828938
		return o
	},
	-682079097: func() TLObject { // 0xd7584c87
		o := MakeTLStatsGroupTopAdmin(nil)
		o.Data2.Constructor = -682079097
		return o
	},
	1398765469: func() TLObject { // 0x535f779d
		o := MakeTLStatsGroupTopInviter(nil)
		o.Data2.Constructor = 1398765469
		return o
	},
	-1660637285: func() TLObject { // 0x9d04af9b
		o := MakeTLStatsGroupTopPoster(nil)
		o.Data2.Constructor = -1660637285
		return o
	},
	-875679776: func() TLObject { // 0xcbce2fe0
		o := MakeTLStatsPercentValue(nil)
		o.Data2.Constructor = -875679776
		return o
	},
	1202287072: func() TLObject { // 0x47a971e0
		o := MakeTLStatsURL(nil)
		o.Data2.Constructor = 1202287072
		return o
	},
	1409802903: func() TLObject { // 0x5407e297
		o := MakeTLStatsBroadcastRevenueStats(nil)
		o.Data2.Constructor = 1409802903
		return o
	},
	-797226067: func() TLObject { // 0xd07b4bad
		o := MakeTLStatsBroadcastRevenueStats(nil)
		o.Data2.Constructor = -797226067
		return o
	},
	-2028632986: func() TLObject { // 0x87158466
		o := MakeTLStatsBroadcastRevenueTransactions(nil)
		o.Data2.Constructor = -2028632986
		return o
	},
	-328886473: func() TLObject { // 0xec659737
		o := MakeTLStatsBroadcastRevenueWithdrawalUrl(nil)
		o.Data2.Constructor = -328886473
		return o
	},
	963421692: func() TLObject { // 0x396ca5fc
		o := MakeTLStatsBroadcastStats(nil)
		o.Data2.Constructor = 963421692
		return o
	},
	-1107852396: func() TLObject { // 0xbdf78394
		o := MakeTLStatsBroadcastStats(nil)
		o.Data2.Constructor = -1107852396
		return o
	},
	-276825834: func() TLObject { // 0xef7ff916
		o := MakeTLStatsMegagroupStats(nil)
		o.Data2.Constructor = -276825834
		return o
	},
	2145983508: func() TLObject { // 0x7fe91c14
		o := MakeTLStatsMessageStats(nil)
		o.Data2.Constructor = 2145983508
		return o
	},
	-1986399595: func() TLObject { // 0x8999f295
		o := MakeTLStatsMessageStats(nil)
		o.Data2.Constructor = -1986399595
		return o
	},
	-1828487648: func() TLObject { // 0x93037e20
		o := MakeTLStatsPublicForwards(nil)
		o.Data2.Constructor = -1828487648
		return o
	},
	1355613820: func() TLObject { // 0x50cd067c
		o := MakeTLStatsStoryStats(nil)
		o.Data2.Constructor = 1355613820
		return o
	},
	-50416996: func() TLObject { // 0xfcfeb29c
		o := MakeTLStickerKeyword(nil)
		o.Data2.Constructor = -50416996
		return o
	},
	313694676: func() TLObject { // 0x12b299d4
		o := MakeTLStickerPack(nil)
		o.Data2.Constructor = 313694676
		return o
	},
	768691932: func() TLObject { // 0x2dd14edc
		o := MakeTLStickerSet(nil)
		o.Data2.Constructor = 768691932
		return o
	},
	-673242758: func() TLObject { // 0xd7df217a
		o := MakeTLStickerSet(nil)
		o.Data2.Constructor = -673242758
		return o
	},
	1678812626: func() TLObject { // 0x6410a5d2
		o := MakeTLStickerSetCovered(nil)
		o.Data2.Constructor = 1678812626
		return o
	},
	872932635: func() TLObject { // 0x3407e51b
		o := MakeTLStickerSetMultiCovered(nil)
		o.Data2.Constructor = 872932635
		return o
	},
	1087454222: func() TLObject { // 0x40d13c0e
		o := MakeTLStickerSetFullCovered(nil)
		o.Data2.Constructor = 1087454222
		return o
	},
	451763941: func() TLObject { // 0x1aed5ee5
		o := MakeTLStickerSetFullCovered(nil)
		o.Data2.Constructor = 451763941
		return o
	},
	2008112412: func() TLObject { // 0x77b15d1c
		o := MakeTLStickerSetNoCovered(nil)
		o.Data2.Constructor = 2008112412
		return o
	},
	-2046910401: func() TLObject { // 0x85fea03f
		o := MakeTLStickersSuggestedShortName(nil)
		o.Data2.Constructor = -2046910401
		return o
	},
	-1432995067: func() TLObject { // 0xaa963b05
		o := MakeTLStorageFileUnknown(nil)
		o.Data2.Constructor = -1432995067
		return o
	},
	1086091090: func() TLObject { // 0x40bc6f52
		o := MakeTLStorageFilePartial(nil)
		o.Data2.Constructor = 1086091090
		return o
	},
	8322574: func() TLObject { // 0x7efe0e
		o := MakeTLStorageFileJpeg(nil)
		o.Data2.Constructor = 8322574
		return o
	},
	-891180321: func() TLObject { // 0xcae1aadf
		o := MakeTLStorageFileGif(nil)
		o.Data2.Constructor = -891180321
		return o
	},
	172975040: func() TLObject { // 0xa4f63c0
		o := MakeTLStorageFilePng(nil)
		o.Data2.Constructor = 172975040
		return o
	},
	-1373745011: func() TLObject { // 0xae1e508d
		o := MakeTLStorageFilePdf(nil)
		o.Data2.Constructor = -1373745011
		return o
	},
	1384777335: func() TLObject { // 0x528a0677
		o := MakeTLStorageFileMp3(nil)
		o.Data2.Constructor = 1384777335
		return o
	},
	1258941372: func() TLObject { // 0x4b09ebbc
		o := MakeTLStorageFileMov(nil)
		o.Data2.Constructor = 1258941372
		return o
	},
	-1278304028: func() TLObject { // 0xb3cea0e4
		o := MakeTLStorageFileMp4(nil)
		o.Data2.Constructor = -1278304028
		return o
	},
	276907596: func() TLObject { // 0x1081464c
		o := MakeTLStorageFileWebp(nil)
		o.Data2.Constructor = 276907596
		return o
	},
	1898850301: func() TLObject { // 0x712e27fd
		o := MakeTLStoriesStealthMode(nil)
		o.Data2.Constructor = 1898850301
		return o
	},
	291044926: func() TLObject { // 0x1158fe3e
		o := MakeTLStoriesAllStoriesNotModified(nil)
		o.Data2.Constructor = 291044926
		return o
	},
	1205903486: func() TLObject { // 0x47e0a07e
		o := MakeTLStoriesAllStoriesNotModified(nil)
		o.Data2.Constructor = 1205903486
		return o
	},
	1862033025: func() TLObject { // 0x6efc5e81
		o := MakeTLStoriesAllStories(nil)
		o.Data2.Constructor = 1862033025
		return o
	},
	1369278878: func() TLObject { // 0x519d899e
		o := MakeTLStoriesAllStories(nil)
		o.Data2.Constructor = 1369278878
		return o
	},
	-2086796248: func() TLObject { // 0x839e0428
		o := MakeTLStoriesAllStories(nil)
		o.Data2.Constructor = -2086796248
		return o
	},
	-203604707: func() TLObject { // 0xf3dd3d1d
		o := MakeTLStoriesBoostersList(nil)
		o.Data2.Constructor = -203604707
		return o
	},
	-440292772: func() TLObject { // 0xe5c1aa5c
		o := MakeTLStoriesBoostsStatus(nil)
		o.Data2.Constructor = -440292772
		return o
	},
	1726619631: func() TLObject { // 0x66ea1fef
		o := MakeTLStoriesBoostsStatus(nil)
		o.Data2.Constructor = 1726619631
		return o
	},
	-1021889145: func() TLObject { // 0xc3173587
		o := MakeTLStoriesCanApplyBoostOk(nil)
		o.Data2.Constructor = -1021889145
		return o
	},
	1898726997: func() TLObject { // 0x712c4655
		o := MakeTLStoriesCanApplyBoostReplace(nil)
		o.Data2.Constructor = 1898726997
		return o
	},
	-488736969: func() TLObject { // 0xe2de7737
		o := MakeTLStoriesFoundStories(nil)
		o.Data2.Constructor = -488736969
		return o
	},
	-890861720: func() TLObject { // 0xcae68768
		o := MakeTLStoriesPeerStories(nil)
		o.Data2.Constructor = -890861720
		return o
	},
	1673780490: func() TLObject { // 0x63c3dd0a
		o := MakeTLStoriesStories(nil)
		o.Data2.Constructor = 1673780490
		return o
	},
	1574486984: func() TLObject { // 0x5dd8c3c8
		o := MakeTLStoriesStories(nil)
		o.Data2.Constructor = 1574486984
		return o
	},
	1340440049: func() TLObject { // 0x4fe57df1
		o := MakeTLStoriesStories(nil)
		o.Data2.Constructor = 1340440049
		return o
	},
	-1436583780: func() TLObject { // 0xaa5f789c
		o := MakeTLStoriesStoryReactionsList(nil)
		o.Data2.Constructor = -1436583780
		return o
	},
	-560009955: func() TLObject { // 0xde9eed1d
		o := MakeTLStoriesStoryViews(nil)
		o.Data2.Constructor = -560009955
		return o
	},
	1507299269: func() TLObject { // 0x59d78fc5
		o := MakeTLStoriesStoryViewsList(nil)
		o.Data2.Constructor = 1507299269
		return o
	},
	1189722604: func() TLObject { // 0x46e9b9ec
		o := MakeTLStoriesStoryViewsList(nil)
		o.Data2.Constructor = 1189722604
		return o
	},
	-79726676: func() TLObject { // 0xfb3f77ac
		o := MakeTLStoriesStoryViewsList(nil)
		o.Data2.Constructor = -79726676
		return o
	},
	933691231: func() TLObject { // 0x37a6ff5f
		o := MakeTLStoriesUserStories(nil)
		o.Data2.Constructor = 933691231
		return o
	},
	-1205411504: func() TLObject { // 0xb826e150
		o := MakeTLStoryFwdHeader(nil)
		o.Data2.Constructor = -1205411504
		return o
	},
	1374088783: func() TLObject { // 0x51e6ee4f
		o := MakeTLStoryItemDeleted(nil)
		o.Data2.Constructor = 1374088783
		return o
	},
	-5388013: func() TLObject { // 0xffadc913
		o := MakeTLStoryItemSkipped(nil)
		o.Data2.Constructor = -5388013
		return o
	},
	2041735716: func() TLObject { // 0x79b26a24
		o := MakeTLStoryItem(nil)
		o.Data2.Constructor = 2041735716
		return o
	},
	-1352440415: func() TLObject { // 0xaf6365a1
		o := MakeTLStoryItem(nil)
		o.Data2.Constructor = -1352440415
		return o
	},
	1153718222: func() TLObject { // 0x44c457ce
		o := MakeTLStoryItem(nil)
		o.Data2.Constructor = 1153718222
		return o
	},
	1445635639: func() TLObject { // 0x562aa637
		o := MakeTLStoryItem(nil)
		o.Data2.Constructor = 1445635639
		return o
	},
	1620104917: func() TLObject { // 0x6090d6d5
		o := MakeTLStoryReaction(nil)
		o.Data2.Constructor = 1620104917
		return o
	},
	-1146411453: func() TLObject { // 0xbbab2643
		o := MakeTLStoryReactionPublicForward(nil)
		o.Data2.Constructor = -1146411453
		return o
	},
	-808644845: func() TLObject { // 0xcfcd0f13
		o := MakeTLStoryReactionPublicRepost(nil)
		o.Data2.Constructor = -808644845
		return o
	},
	-1329730875: func() TLObject { // 0xb0bdeac5
		o := MakeTLStoryView(nil)
		o.Data2.Constructor = -1329730875
		return o
	},
	-1491424062: func() TLObject { // 0xa71aacc2
		o := MakeTLStoryView(nil)
		o.Data2.Constructor = -1491424062
		return o
	},
	-1870436597: func() TLObject { // 0x9083670b
		o := MakeTLStoryViewPublicForward(nil)
		o.Data2.Constructor = -1870436597
		return o
	},
	-1116418231: func() TLObject { // 0xbd74cf49
		o := MakeTLStoryViewPublicRepost(nil)
		o.Data2.Constructor = -1116418231
		return o
	},
	-1923523370: func() TLObject { // 0x8d595cd6
		o := MakeTLStoryViews(nil)
		o.Data2.Constructor = -1923523370
		return o
	},
	-968094825: func() TLObject { // 0xc64c0b97
		o := MakeTLStoryViews(nil)
		o.Data2.Constructor = -968094825
		return o
	},
	-748199729: func() TLObject { // 0xd36760cf
		o := MakeTLStoryViews(nil)
		o.Data2.Constructor = -748199729
		return o
	},
	194458693: func() TLObject { // 0xb973445
		o := MakeTLString(nil)
		o.Data2.Constructor = 194458693
		return o
	},
	1964978502: func() TLObject { // 0x751f3146
		o := MakeTLTextWithEntities(nil)
		o.Data2.Constructor = 1964978502
		return o
	},
	-1609668650: func() TLObject { // 0xa00e67d6
		o := MakeTLTheme(nil)
		o.Data2.Constructor = -1609668650
		return o
	},
	-94849324: func() TLObject { // 0xfa58b6d4
		o := MakeTLThemeSettings(nil)
		o.Data2.Constructor = -94849324
		return o
	},
	-7173643: func() TLObject { // 0xff9289f5
		o := MakeTLTimezone(nil)
		o.Data2.Constructor = -7173643
		return o
	},
	-305282981: func() TLObject { // 0xedcdc05b
		o := MakeTLTopPeer(nil)
		o.Data2.Constructor = -305282981
		return o
	},
	-1419371685: func() TLObject { // 0xab661b5b
		o := MakeTLTopPeerCategoryBotsPM(nil)
		o.Data2.Constructor = -1419371685
		return o
	},
	344356834: func() TLObject { // 0x148677e2
		o := MakeTLTopPeerCategoryBotsInline(nil)
		o.Data2.Constructor = 344356834
		return o
	},
	104314861: func() TLObject { // 0x637b7ed
		o := MakeTLTopPeerCategoryCorrespondents(nil)
		o.Data2.Constructor = 104314861
		return o
	},
	-1122524854: func() TLObject { // 0xbd17a14a
		o := MakeTLTopPeerCategoryGroups(nil)
		o.Data2.Constructor = -1122524854
		return o
	},
	371037736: func() TLObject { // 0x161d9628
		o := MakeTLTopPeerCategoryChannels(nil)
		o.Data2.Constructor = 371037736
		return o
	},
	511092620: func() TLObject { // 0x1e76a78c
		o := MakeTLTopPeerCategoryPhoneCalls(nil)
		o.Data2.Constructor = 511092620
		return o
	},
	-1472172887: func() TLObject { // 0xa8406ca9
		o := MakeTLTopPeerCategoryForwardUsers(nil)
		o.Data2.Constructor = -1472172887
		return o
	},
	-68239120: func() TLObject { // 0xfbeec0f0
		o := MakeTLTopPeerCategoryForwardChats(nil)
		o.Data2.Constructor = -68239120
		return o
	},
	-75283823: func() TLObject { // 0xfb834291
		o := MakeTLTopPeerCategoryPeers(nil)
		o.Data2.Constructor = -75283823
		return o
	},
	1072550713: func() TLObject { // 0x3fedd339
		o := MakeTLTrue(nil)
		o.Data2.Constructor = 1072550713
		return o
	},
	522914557: func() TLObject { // 0x1f2b0afd
		o := MakeTLUpdateNewMessage(nil)
		o.Data2.Constructor = 522914557
		return o
	},
	1318109142: func() TLObject { // 0x4e90bfd6
		o := MakeTLUpdateMessageID(nil)
		o.Data2.Constructor = 1318109142
		return o
	},
	-1576161051: func() TLObject { // 0xa20db0e5
		o := MakeTLUpdateDeleteMessages(nil)
		o.Data2.Constructor = -1576161051
		return o
	},
	-1071741569: func() TLObject { // 0xc01e857f
		o := MakeTLUpdateUserTyping(nil)
		o.Data2.Constructor = -1071741569
		return o
	},
	-2092401936: func() TLObject { // 0x83487af0
		o := MakeTLUpdateChatUserTyping(nil)
		o.Data2.Constructor = -2092401936
		return o
	},
	125178264: func() TLObject { // 0x7761198
		o := MakeTLUpdateChatParticipants(nil)
		o.Data2.Constructor = 125178264
		return o
	},
	-440534818: func() TLObject { // 0xe5bdf8de
		o := MakeTLUpdateUserStatus(nil)
		o.Data2.Constructor = -440534818
		return o
	},
	-1484486364: func() TLObject { // 0xa7848924
		o := MakeTLUpdateUserName(nil)
		o.Data2.Constructor = -1484486364
		return o
	},
	-1007549728: func() TLObject { // 0xc3f202e0
		o := MakeTLUpdateUserName(nil)
		o.Data2.Constructor = -1007549728
		return o
	},
	-1991136273: func() TLObject { // 0x8951abef
		o := MakeTLUpdateNewAuthorization(nil)
		o.Data2.Constructor = -1991136273
		return o
	},
	314359194: func() TLObject { // 0x12bcbd9a
		o := MakeTLUpdateNewEncryptedMessage(nil)
		o.Data2.Constructor = 314359194
		return o
	},
	386986326: func() TLObject { // 0x1710f156
		o := MakeTLUpdateEncryptedChatTyping(nil)
		o.Data2.Constructor = 386986326
		return o
	},
	-1264392051: func() TLObject { // 0xb4a2e88d
		o := MakeTLUpdateEncryption(nil)
		o.Data2.Constructor = -1264392051
		return o
	},
	956179895: func() TLObject { // 0x38fe25b7
		o := MakeTLUpdateEncryptedMessagesRead(nil)
		o.Data2.Constructor = 956179895
		return o
	},
	1037718609: func() TLObject { // 0x3dda5451
		o := MakeTLUpdateChatParticipantAdd(nil)
		o.Data2.Constructor = 1037718609
		return o
	},
	-483443337: func() TLObject { // 0xe32f3d77
		o := MakeTLUpdateChatParticipantDelete(nil)
		o.Data2.Constructor = -483443337
		return o
	},
	-1906403213: func() TLObject { // 0x8e5e9873
		o := MakeTLUpdateDcOptions(nil)
		o.Data2.Constructor = -1906403213
		return o
	},
	-1094555409: func() TLObject { // 0xbec268ef
		o := MakeTLUpdateNotifySettings(nil)
		o.Data2.Constructor = -1094555409
		return o
	},
	-337352679: func() TLObject { // 0xebe46819
		o := MakeTLUpdateServiceNotification(nil)
		o.Data2.Constructor = -337352679
		return o
	},
	-298113238: func() TLObject { // 0xee3b272a
		o := MakeTLUpdatePrivacy(nil)
		o.Data2.Constructor = -298113238
		return o
	},
	88680979: func() TLObject { // 0x5492a13
		o := MakeTLUpdateUserPhone(nil)
		o.Data2.Constructor = 88680979
		return o
	},
	-1667805217: func() TLObject { // 0x9c974fdf
		o := MakeTLUpdateReadHistoryInbox(nil)
		o.Data2.Constructor = -1667805217
		return o
	},
	791617983: func() TLObject { // 0x2f2f21bf
		o := MakeTLUpdateReadHistoryOutbox(nil)
		o.Data2.Constructor = 791617983
		return o
	},
	2139689491: func() TLObject { // 0x7f891213
		o := MakeTLUpdateWebPage(nil)
		o.Data2.Constructor = 2139689491
		return o
	},
	-131960447: func() TLObject { // 0xf8227181
		o := MakeTLUpdateReadMessagesContents(nil)
		o.Data2.Constructor = -131960447
		return o
	},
	1757493555: func() TLObject { // 0x68c13933
		o := MakeTLUpdateReadMessagesContents(nil)
		o.Data2.Constructor = 1757493555
		return o
	},
	277713951: func() TLObject { // 0x108d941f
		o := MakeTLUpdateChannelTooLong(nil)
		o.Data2.Constructor = 277713951
		return o
	},
	1666927625: func() TLObject { // 0x635b4c09
		o := MakeTLUpdateChannel(nil)
		o.Data2.Constructor = 1666927625
		return o
	},
	1656358105: func() TLObject { // 0x62ba04d9
		o := MakeTLUpdateNewChannelMessage(nil)
		o.Data2.Constructor = 1656358105
		return o
	},
	-1842450928: func() TLObject { // 0x922e6e10
		o := MakeTLUpdateReadChannelInbox(nil)
		o.Data2.Constructor = -1842450928
		return o
	},
	-1020437742: func() TLObject { // 0xc32d5b12
		o := MakeTLUpdateDeleteChannelMessages(nil)
		o.Data2.Constructor = -1020437742
		return o
	},
	-232346616: func() TLObject { // 0xf226ac08
		o := MakeTLUpdateChannelMessageViews(nil)
		o.Data2.Constructor = -232346616
		return o
	},
	-674602590: func() TLObject { // 0xd7ca61a2
		o := MakeTLUpdateChatParticipantAdmin(nil)
		o.Data2.Constructor = -674602590
		return o
	},
	1753886890: func() TLObject { // 0x688a30aa
		o := MakeTLUpdateNewStickerSet(nil)
		o.Data2.Constructor = 1753886890
		return o
	},
	196268545: func() TLObject { // 0xbb2d201
		o := MakeTLUpdateStickerSetsOrder(nil)
		o.Data2.Constructor = 196268545
		return o
	},
	834816008: func() TLObject { // 0x31c24808
		o := MakeTLUpdateStickerSets(nil)
		o.Data2.Constructor = 834816008
		return o
	},
	1135492588: func() TLObject { // 0x43ae3dec
		o := MakeTLUpdateStickerSets(nil)
		o.Data2.Constructor = 1135492588
		return o
	},
	-1821035490: func() TLObject { // 0x9375341e
		o := MakeTLUpdateSavedGifs(nil)
		o.Data2.Constructor = -1821035490
		return o
	},
	1232025500: func() TLObject { // 0x496f379c
		o := MakeTLUpdateBotInlineQuery(nil)
		o.Data2.Constructor = 1232025500
		return o
	},
	317794823: func() TLObject { // 0x12f12a07
		o := MakeTLUpdateBotInlineSend(nil)
		o.Data2.Constructor = 317794823
		return o
	},
	457133559: func() TLObject { // 0x1b3f4df7
		o := MakeTLUpdateEditChannelMessage(nil)
		o.Data2.Constructor = 457133559
		return o
	},
	-1177566067: func() TLObject { // 0xb9cfc48d
		o := MakeTLUpdateBotCallbackQuery(nil)
		o.Data2.Constructor = -1177566067
		return o
	},
	-469536605: func() TLObject { // 0xe40370a3
		o := MakeTLUpdateEditMessage(nil)
		o.Data2.Constructor = -469536605
		return o
	},
	1763610706: func() TLObject { // 0x691e9052
		o := MakeTLUpdateInlineBotCallbackQuery(nil)
		o.Data2.Constructor = 1763610706
		return o
	},
	-1218471511: func() TLObject { // 0xb75f99a9
		o := MakeTLUpdateReadChannelOutbox(nil)
		o.Data2.Constructor = -1218471511
		return o
	},
	457829485: func() TLObject { // 0x1b49ec6d
		o := MakeTLUpdateDraftMessage(nil)
		o.Data2.Constructor = 457829485
		return o
	},
	-299124375: func() TLObject { // 0xee2bb969
		o := MakeTLUpdateDraftMessage(nil)
		o.Data2.Constructor = -299124375
		return o
	},
	1461528386: func() TLObject { // 0x571d2742
		o := MakeTLUpdateReadFeaturedStickers(nil)
		o.Data2.Constructor = 1461528386
		return o
	},
	-1706939360: func() TLObject { // 0x9a422c20
		o := MakeTLUpdateRecentStickers(nil)
		o.Data2.Constructor = -1706939360
		return o
	},
	-1574314746: func() TLObject { // 0xa229dd06
		o := MakeTLUpdateConfig(nil)
		o.Data2.Constructor = -1574314746
		return o
	},
	861169551: func() TLObject { // 0x3354678f
		o := MakeTLUpdatePtsChanged(nil)
		o.Data2.Constructor = 861169551
		return o
	},
	791390623: func() TLObject { // 0x2f2ba99f
		o := MakeTLUpdateChannelWebPage(nil)
		o.Data2.Constructor = 791390623
		return o
	},
	1852826908: func() TLObject { // 0x6e6fe51c
		o := MakeTLUpdateDialogPinned(nil)
		o.Data2.Constructor = 1852826908
		return o
	},
	-99664734: func() TLObject { // 0xfa0f3ca2
		o := MakeTLUpdatePinnedDialogs(nil)
		o.Data2.Constructor = -99664734
		return o
	},
	-2095595325: func() TLObject { // 0x8317c0c3
		o := MakeTLUpdateBotWebhookJSON(nil)
		o.Data2.Constructor = -2095595325
		return o
	},
	-1684914010: func() TLObject { // 0x9b9240a6
		o := MakeTLUpdateBotWebhookJSONQuery(nil)
		o.Data2.Constructor = -1684914010
		return o
	},
	-1246823043: func() TLObject { // 0xb5aefd7d
		o := MakeTLUpdateBotShippingQuery(nil)
		o.Data2.Constructor = -1246823043
		return o
	},
	-1934976362: func() TLObject { // 0x8caa9a96
		o := MakeTLUpdateBotPrecheckoutQuery(nil)
		o.Data2.Constructor = -1934976362
		return o
	},
	-1425052898: func() TLObject { // 0xab0f6b1e
		o := MakeTLUpdatePhoneCall(nil)
		o.Data2.Constructor = -1425052898
		return o
	},
	1180041828: func() TLObject { // 0x46560264
		o := MakeTLUpdateLangPackTooLong(nil)
		o.Data2.Constructor = 1180041828
		return o
	},
	1442983757: func() TLObject { // 0x56022f4d
		o := MakeTLUpdateLangPack(nil)
		o.Data2.Constructor = 1442983757
		return o
	},
	-451831443: func() TLObject { // 0xe511996d
		o := MakeTLUpdateFavedStickers(nil)
		o.Data2.Constructor = -451831443
		return o
	},
	-366410403: func() TLObject { // 0xea29055d
		o := MakeTLUpdateChannelReadMessagesContents(nil)
		o.Data2.Constructor = -366410403
		return o
	},
	1153291573: func() TLObject { // 0x44bdd535
		o := MakeTLUpdateChannelReadMessagesContents(nil)
		o.Data2.Constructor = 1153291573
		return o
	},
	1887741886: func() TLObject { // 0x7084a7be
		o := MakeTLUpdateContactsReset(nil)
		o.Data2.Constructor = 1887741886
		return o
	},
	-1304443240: func() TLObject { // 0xb23fc698
		o := MakeTLUpdateChannelAvailableMessages(nil)
		o.Data2.Constructor = -1304443240
		return o
	},
	-513517117: func() TLObject { // 0xe16459c3
		o := MakeTLUpdateDialogUnreadMark(nil)
		o.Data2.Constructor = -513517117
		return o
	},
	-1398708869: func() TLObject { // 0xaca1657b
		o := MakeTLUpdateMessagePoll(nil)
		o.Data2.Constructor = -1398708869
		return o
	},
	1421875280: func() TLObject { // 0x54c01850
		o := MakeTLUpdateChatDefaultBannedRights(nil)
		o.Data2.Constructor = 1421875280
		return o
	},
	422972864: func() TLObject { // 0x19360dc0
		o := MakeTLUpdateFolderPeers(nil)
		o.Data2.Constructor = 422972864
		return o
	},
	1786671974: func() TLObject { // 0x6a7e7366
		o := MakeTLUpdatePeerSettings(nil)
		o.Data2.Constructor = 1786671974
		return o
	},
	-1263546448: func() TLObject { // 0xb4afcfb0
		o := MakeTLUpdatePeerLocated(nil)
		o.Data2.Constructor = -1263546448
		return o
	},
	967122427: func() TLObject { // 0x39a51dfb
		o := MakeTLUpdateNewScheduledMessage(nil)
		o.Data2.Constructor = 967122427
		return o
	},
	-1870238482: func() TLObject { // 0x90866cee
		o := MakeTLUpdateDeleteScheduledMessages(nil)
		o.Data2.Constructor = -1870238482
		return o
	},
	-2112423005: func() TLObject { // 0x8216fba3
		o := MakeTLUpdateTheme(nil)
		o.Data2.Constructor = -2112423005
		return o
	},
	-2027964103: func() TLObject { // 0x871fb939
		o := MakeTLUpdateGeoLiveViewed(nil)
		o.Data2.Constructor = -2027964103
		return o
	},
	1448076945: func() TLObject { // 0x564fe691
		o := MakeTLUpdateLoginToken(nil)
		o.Data2.Constructor = 1448076945
		return o
	},
	619974263: func() TLObject { // 0x24f40e77
		o := MakeTLUpdateMessagePollVote(nil)
		o.Data2.Constructor = 619974263
		return o
	},
	274961865: func() TLObject { // 0x106395c9
		o := MakeTLUpdateMessagePollVote(nil)
		o.Data2.Constructor = 274961865
		return o
	},
	654302845: func() TLObject { // 0x26ffde7d
		o := MakeTLUpdateDialogFilter(nil)
		o.Data2.Constructor = 654302845
		return o
	},
	-1512627963: func() TLObject { // 0xa5d72105
		o := MakeTLUpdateDialogFilterOrder(nil)
		o.Data2.Constructor = -1512627963
		return o
	},
	889491791: func() TLObject { // 0x3504914f
		o := MakeTLUpdateDialogFilters(nil)
		o.Data2.Constructor = 889491791
		return o
	},
	643940105: func() TLObject { // 0x2661bf09
		o := MakeTLUpdatePhoneCallSignalingData(nil)
		o.Data2.Constructor = 643940105
		return o
	},
	-761649164: func() TLObject { // 0xd29a27f4
		o := MakeTLUpdateChannelMessageForwards(nil)
		o.Data2.Constructor = -761649164
		return o
	},
	-693004986: func() TLObject { // 0xd6b19546
		o := MakeTLUpdateReadChannelDiscussionInbox(nil)
		o.Data2.Constructor = -693004986
		return o
	},
	1767677564: func() TLObject { // 0x695c9e7c
		o := MakeTLUpdateReadChannelDiscussionOutbox(nil)
		o.Data2.Constructor = 1767677564
		return o
	},
	-337610926: func() TLObject { // 0xebe07752
		o := MakeTLUpdatePeerBlocked(nil)
		o.Data2.Constructor = -337610926
		return o
	},
	610945826: func() TLObject { // 0x246a4b22
		o := MakeTLUpdatePeerBlocked(nil)
		o.Data2.Constructor = 610945826
		return o
	},
	-1937192669: func() TLObject { // 0x8c88c923
		o := MakeTLUpdateChannelUserTyping(nil)
		o.Data2.Constructor = -1937192669
		return o
	},
	-309990731: func() TLObject { // 0xed85eab5
		o := MakeTLUpdatePinnedMessages(nil)
		o.Data2.Constructor = -309990731
		return o
	},
	1538885128: func() TLObject { // 0x5bb98608
		o := MakeTLUpdatePinnedChannelMessages(nil)
		o.Data2.Constructor = 1538885128
		return o
	},
	-124097970: func() TLObject { // 0xf89a6a4e
		o := MakeTLUpdateChat(nil)
		o.Data2.Constructor = -124097970
		return o
	},
	-219423922: func() TLObject { // 0xf2ebdb4e
		o := MakeTLUpdateGroupCallParticipants(nil)
		o.Data2.Constructor = -219423922
		return o
	},
	347227392: func() TLObject { // 0x14b24500
		o := MakeTLUpdateGroupCall(nil)
		o.Data2.Constructor = 347227392
		return o
	},
	-1147422299: func() TLObject { // 0xbb9bb9a5
		o := MakeTLUpdatePeerHistoryTTL(nil)
		o.Data2.Constructor = -1147422299
		return o
	},
	-796432838: func() TLObject { // 0xd087663a
		o := MakeTLUpdateChatParticipant(nil)
		o.Data2.Constructor = -796432838
		return o
	},
	-1738720581: func() TLObject { // 0x985d3abb
		o := MakeTLUpdateChannelParticipant(nil)
		o.Data2.Constructor = -1738720581
		return o
	},
	-997782967: func() TLObject { // 0xc4870a49
		o := MakeTLUpdateBotStopped(nil)
		o.Data2.Constructor = -997782967
		return o
	},
	192428418: func() TLObject { // 0xb783982
		o := MakeTLUpdateGroupCallConnection(nil)
		o.Data2.Constructor = 192428418
		return o
	},
	1299263278: func() TLObject { // 0x4d712f2e
		o := MakeTLUpdateBotCommands(nil)
		o.Data2.Constructor = 1299263278
		return o
	},
	1885586395: func() TLObject { // 0x7063c3db
		o := MakeTLUpdatePendingJoinRequests(nil)
		o.Data2.Constructor = 1885586395
		return o
	},
	299870598: func() TLObject { // 0x11dfa986
		o := MakeTLUpdateBotChatInviteRequester(nil)
		o.Data2.Constructor = 299870598
		return o
	},
	1578843320: func() TLObject { // 0x5e1b3cb8
		o := MakeTLUpdateMessageReactions(nil)
		o.Data2.Constructor = 1578843320
		return o
	},
	357013699: func() TLObject { // 0x154798c3
		o := MakeTLUpdateMessageReactions(nil)
		o.Data2.Constructor = 357013699
		return o
	},
	397910539: func() TLObject { // 0x17b7a20b
		o := MakeTLUpdateAttachMenuBots(nil)
		o.Data2.Constructor = 397910539
		return o
	},
	361936797: func() TLObject { // 0x1592b79d
		o := MakeTLUpdateWebViewResultSent(nil)
		o.Data2.Constructor = 361936797
		return o
	},
	347625491: func() TLObject { // 0x14b85813
		o := MakeTLUpdateBotMenuButton(nil)
		o.Data2.Constructor = 347625491
		return o
	},
	1960361625: func() TLObject { // 0x74d8be99
		o := MakeTLUpdateSavedRingtones(nil)
		o.Data2.Constructor = 1960361625
		return o
	},
	8703322: func() TLObject { // 0x84cd5a
		o := MakeTLUpdateTranscribedAudio(nil)
		o.Data2.Constructor = 8703322
		return o
	},
	-78886548: func() TLObject { // 0xfb4c496c
		o := MakeTLUpdateReadFeaturedEmojiStickers(nil)
		o.Data2.Constructor = -78886548
		return o
	},
	674706841: func() TLObject { // 0x28373599
		o := MakeTLUpdateUserEmojiStatus(nil)
		o.Data2.Constructor = 674706841
		return o
	},
	821314523: func() TLObject { // 0x30f443db
		o := MakeTLUpdateRecentEmojiStatuses(nil)
		o.Data2.Constructor = 821314523
		return o
	},
	1870160884: func() TLObject { // 0x6f7863f4
		o := MakeTLUpdateRecentReactions(nil)
		o.Data2.Constructor = 1870160884
		return o
	},
	-2030252155: func() TLObject { // 0x86fccf85
		o := MakeTLUpdateMoveStickerSetToTop(nil)
		o.Data2.Constructor = -2030252155
		return o
	},
	-710666460: func() TLObject { // 0xd5a41724
		o := MakeTLUpdateMessageExtendedMedia(nil)
		o.Data2.Constructor = -710666460
		return o
	},
	1517529484: func() TLObject { // 0x5a73a98c
		o := MakeTLUpdateMessageExtendedMedia(nil)
		o.Data2.Constructor = 1517529484
		return o
	},
	422509539: func() TLObject { // 0x192efbe3
		o := MakeTLUpdateChannelPinnedTopic(nil)
		o.Data2.Constructor = 422509539
		return o
	},
	-158027602: func() TLObject { // 0xf694b0ae
		o := MakeTLUpdateChannelPinnedTopic(nil)
		o.Data2.Constructor = -158027602
		return o
	},
	-31881726: func() TLObject { // 0xfe198602
		o := MakeTLUpdateChannelPinnedTopics(nil)
		o.Data2.Constructor = -31881726
		return o
	},
	542282808: func() TLObject { // 0x20529438
		o := MakeTLUpdateUser(nil)
		o.Data2.Constructor = 542282808
		return o
	},
	-335171433: func() TLObject { // 0xec05b097
		o := MakeTLUpdateAutoSaveSettings(nil)
		o.Data2.Constructor = -335171433
		return o
	},
	1974712216: func() TLObject { // 0x75b3b798
		o := MakeTLUpdateStory(nil)
		o.Data2.Constructor = 1974712216
		return o
	},
	542785843: func() TLObject { // 0x205a4133
		o := MakeTLUpdateStory(nil)
		o.Data2.Constructor = 542785843
		return o
	},
	-145845461: func() TLObject { // 0xf74e932b
		o := MakeTLUpdateReadStories(nil)
		o.Data2.Constructor = -145845461
		return o
	},
	-21679014: func() TLObject { // 0xfeb5345a
		o := MakeTLUpdateReadStories(nil)
		o.Data2.Constructor = -21679014
		return o
	},
	468923833: func() TLObject { // 0x1bf335b9
		o := MakeTLUpdateStoryID(nil)
		o.Data2.Constructor = 468923833
		return o
	},
	738741697: func() TLObject { // 0x2c084dc1
		o := MakeTLUpdateStoriesStealthMode(nil)
		o.Data2.Constructor = 738741697
		return o
	},
	2103604867: func() TLObject { // 0x7d627683
		o := MakeTLUpdateSentStoryReaction(nil)
		o.Data2.Constructor = 2103604867
		return o
	},
	-475579104: func() TLObject { // 0xe3a73d20
		o := MakeTLUpdateSentStoryReaction(nil)
		o.Data2.Constructor = -475579104
		return o
	},
	-1873947492: func() TLObject { // 0x904dd49c
		o := MakeTLUpdateBotChatBoost(nil)
		o.Data2.Constructor = -1873947492
		return o
	},
	129403168: func() TLObject { // 0x7b68920
		o := MakeTLUpdateChannelViewForumAsMessages(nil)
		o.Data2.Constructor = 129403168
		return o
	},
	-1371598819: func() TLObject { // 0xae3f101d
		o := MakeTLUpdatePeerWallpaper(nil)
		o.Data2.Constructor = -1371598819
		return o
	},
	-1407069234: func() TLObject { // 0xac21d3ce
		o := MakeTLUpdateBotMessageReaction(nil)
		o.Data2.Constructor = -1407069234
		return o
	},
	164329305: func() TLObject { // 0x9cb7759
		o := MakeTLUpdateBotMessageReactions(nil)
		o.Data2.Constructor = 164329305
		return o
	},
	-1364222348: func() TLObject { // 0xaeaf9e74
		o := MakeTLUpdateSavedDialogPinned(nil)
		o.Data2.Constructor = -1364222348
		return o
	},
	1751942566: func() TLObject { // 0x686c85a6
		o := MakeTLUpdatePinnedSavedDialogs(nil)
		o.Data2.Constructor = 1751942566
		return o
	},
	969307186: func() TLObject { // 0x39c67432
		o := MakeTLUpdateSavedReactionTags(nil)
		o.Data2.Constructor = 969307186
		return o
	},
	-245208620: func() TLObject { // 0xf16269d4
		o := MakeTLUpdateSmsJob(nil)
		o.Data2.Constructor = -245208620
		return o
	},
	-112784718: func() TLObject { // 0xf9470ab2
		o := MakeTLUpdateQuickReplies(nil)
		o.Data2.Constructor = -112784718
		return o
	},
	-180508905: func() TLObject { // 0xf53da717
		o := MakeTLUpdateNewQuickReply(nil)
		o.Data2.Constructor = -180508905
		return o
	},
	1407644140: func() TLObject { // 0x53e6f1ec
		o := MakeTLUpdateDeleteQuickReply(nil)
		o.Data2.Constructor = 1407644140
		return o
	},
	1040518415: func() TLObject { // 0x3e050d0f
		o := MakeTLUpdateQuickReplyMessage(nil)
		o.Data2.Constructor = 1040518415
		return o
	},
	1450174413: func() TLObject { // 0x566fe7cd
		o := MakeTLUpdateDeleteQuickReplyMessages(nil)
		o.Data2.Constructor = 1450174413
		return o
	},
	-1964652166: func() TLObject { // 0x8ae5c97a
		o := MakeTLUpdateBotBusinessConnect(nil)
		o.Data2.Constructor = -1964652166
		return o
	},
	-1646578564: func() TLObject { // 0x9ddb347c
		o := MakeTLUpdateBotNewBusinessMessage(nil)
		o.Data2.Constructor = -1646578564
		return o
	},
	132077692: func() TLObject { // 0x7df587c
		o := MakeTLUpdateBotEditBusinessMessage(nil)
		o.Data2.Constructor = 132077692
		return o
	},
	-1607821266: func() TLObject { // 0xa02a982e
		o := MakeTLUpdateBotDeleteBusinessMessage(nil)
		o.Data2.Constructor = -1607821266
		return o
	},
	405070859: func() TLObject { // 0x1824e40b
		o := MakeTLUpdateNewStoryReaction(nil)
		o.Data2.Constructor = 405070859
		return o
	},
	-539401739: func() TLObject { // 0xdfd961f5
		o := MakeTLUpdateBroadcastRevenueTransactions(nil)
		o.Data2.Constructor = -539401739
		return o
	},
	263737752: func() TLObject { // 0xfb85198
		o := MakeTLUpdateStarsBalance(nil)
		o.Data2.Constructor = 263737752
		return o
	},
	513998247: func() TLObject { // 0x1ea2fda7
		o := MakeTLUpdateBusinessBotCallbackQuery(nil)
		o.Data2.Constructor = 513998247
		return o
	},
	-1518030823: func() TLObject { // 0xa584b019
		o := MakeTLUpdateStarsRevenueStatus(nil)
		o.Data2.Constructor = -1518030823
		return o
	},
	-856651050: func() TLObject { // 0xccf08ad6
		o := MakeTLUpdateGroupInvitePrivacyForbidden(nil)
		o.Data2.Constructor = -856651050
		return o
	},
	-232290676: func() TLObject { // 0xf227868c
		o := MakeTLUpdateUserPhoto(nil)
		o.Data2.Constructor = -232290676
		return o
	},
	-2083620338: func() TLObject { // 0x83ce7a0e
		o := MakeTLUpdateBizDataRaw(nil)
		o.Data2.Constructor = -2083620338
		return o
	},
	-1877696350: func() TLObject { // 0x9014a0a2
		o := MakeTLUpdateList(nil)
		o.Data2.Constructor = -1877696350
		return o
	},
	-484987010: func() TLObject { // 0xe317af7e
		o := MakeTLUpdatesTooLong(nil)
		o.Data2.Constructor = -484987010
		return o
	},
	826001400: func() TLObject { // 0x313bc7f8
		o := MakeTLUpdateShortMessage(nil)
		o.Data2.Constructor = 826001400
		return o
	},
	1299050149: func() TLObject { // 0x4d6deea5
		o := MakeTLUpdateShortChatMessage(nil)
		o.Data2.Constructor = 1299050149
		return o
	},
	2027216577: func() TLObject { // 0x78d4dec1
		o := MakeTLUpdateShort(nil)
		o.Data2.Constructor = 2027216577
		return o
	},
	1918567619: func() TLObject { // 0x725b04c3
		o := MakeTLUpdatesCombined(nil)
		o.Data2.Constructor = 1918567619
		return o
	},
	1957577280: func() TLObject { // 0x74ae4240
		o := MakeTLUpdates(nil)
		o.Data2.Constructor = 1957577280
		return o
	},
	-1877614335: func() TLObject { // 0x9015e101
		o := MakeTLUpdateShortSentMessage(nil)
		o.Data2.Constructor = -1877614335
		return o
	},
	294964541: func() TLObject { // 0x1194cd3d
		o := MakeTLUpdateAccountResetAuthorization(nil)
		o.Data2.Constructor = 294964541
		return o
	},
	1041346555: func() TLObject { // 0x3e11affb
		o := MakeTLUpdatesChannelDifferenceEmpty(nil)
		o.Data2.Constructor = 1041346555
		return o
	},
	-1531132162: func() TLObject { // 0xa4bcc6fe
		o := MakeTLUpdatesChannelDifferenceTooLong(nil)
		o.Data2.Constructor = -1531132162
		return o
	},
	543450958: func() TLObject { // 0x2064674e
		o := MakeTLUpdatesChannelDifference(nil)
		o.Data2.Constructor = 543450958
		return o
	},
	1567990072: func() TLObject { // 0x5d75a138
		o := MakeTLUpdatesDifferenceEmpty(nil)
		o.Data2.Constructor = 1567990072
		return o
	},
	16030880: func() TLObject { // 0xf49ca0
		o := MakeTLUpdatesDifference(nil)
		o.Data2.Constructor = 16030880
		return o
	},
	-1459938943: func() TLObject { // 0xa8fb1981
		o := MakeTLUpdatesDifferenceSlice(nil)
		o.Data2.Constructor = -1459938943
		return o
	},
	1258196845: func() TLObject { // 0x4afe8f6d
		o := MakeTLUpdatesDifferenceTooLong(nil)
		o.Data2.Constructor = 1258196845
		return o
	},
	-1519637954: func() TLObject { // 0xa56c2a3e
		o := MakeTLUpdatesState(nil)
		o.Data2.Constructor = -1519637954
		return o
	},
	-290921362: func() TLObject { // 0xeea8e46e
		o := MakeTLUploadCdnFileReuploadNeeded(nil)
		o.Data2.Constructor = -290921362
		return o
	},
	-1449145777: func() TLObject { // 0xa99fca4f
		o := MakeTLUploadCdnFile(nil)
		o.Data2.Constructor = -1449145777
		return o
	},
	157948117: func() TLObject { // 0x96a18d5
		o := MakeTLUploadFile(nil)
		o.Data2.Constructor = 157948117
		return o
	},
	-242427324: func() TLObject { // 0xf18cda44
		o := MakeTLUploadFileCdnRedirect(nil)
		o.Data2.Constructor = -242427324
		return o
	},
	568808380: func() TLObject { // 0x21e753bc
		o := MakeTLUploadWebFile(nil)
		o.Data2.Constructor = 568808380
		return o
	},
	-1831650802: func() TLObject { // 0x92d33a0e
		o := MakeTLUrlAuthResultRequest(nil)
		o.Data2.Constructor = -1831650802
		return o
	},
	-1886646706: func() TLObject { // 0x8f8c0e4e
		o := MakeTLUrlAuthResultAccepted(nil)
		o.Data2.Constructor = -1886646706
		return o
	},
	-1445536993: func() TLObject { // 0xa9d6db1f
		o := MakeTLUrlAuthResultDefault(nil)
		o.Data2.Constructor = -1445536993
		return o
	},
	-742634630: func() TLObject { // 0xd3bc4b7a
		o := MakeTLUserEmpty(nil)
		o.Data2.Constructor = -742634630
		return o
	},
	559694904: func() TLObject { // 0x215c4438
		o := MakeTLUser(nil)
		o.Data2.Constructor = 559694904
		return o
	},
	-346018011: func() TLObject { // 0xeb602f25
		o := MakeTLUser(nil)
		o.Data2.Constructor = -346018011
		return o
	},
	-1414139616: func() TLObject { // 0xabb5f120
		o := MakeTLUser(nil)
		o.Data2.Constructor = -1414139616
		return o
	},
	-1885878744: func() TLObject { // 0x8f97c628
		o := MakeTLUser(nil)
		o.Data2.Constructor = -1885878744
		return o
	},
	1570352622: func() TLObject { // 0x5d99adee
		o := MakeTLUser(nil)
		o.Data2.Constructor = 1570352622
		return o
	},
	1073147056: func() TLObject { // 0x3ff6ecb0
		o := MakeTLUser(nil)
		o.Data2.Constructor = 1073147056
		return o
	},
	509150835: func() TLObject { // 0x1e590673
		o := MakeTLUserData(nil)
		o.Data2.Constructor = 509150835
		return o
	},
	-862357728: func() TLObject { // 0xcc997720
		o := MakeTLUserFull(nil)
		o.Data2.Constructor = -862357728
		return o
	},
	587153029: func() TLObject { // 0x22ff3e85
		o := MakeTLUserFull(nil)
		o.Data2.Constructor = 587153029
		return o
	},
	-1179571092: func() TLObject { // 0xb9b12c6c
		o := MakeTLUserFull(nil)
		o.Data2.Constructor = -1179571092
		return o
	},
	1340198022: func() TLObject { // 0x4fe1cc86
		o := MakeTLUserFull(nil)
		o.Data2.Constructor = 1340198022
		return o
	},
	-1813324973: func() TLObject { // 0x93eadb53
		o := MakeTLUserFull(nil)
		o.Data2.Constructor = -1813324973
		return o
	},
	-120378643: func() TLObject { // 0xf8d32aed
		o := MakeTLUserFull(nil)
		o.Data2.Constructor = -120378643
		return o
	},
	-994968513: func() TLObject { // 0xc4b1fc3f
		o := MakeTLUserFull(nil)
		o.Data2.Constructor = -994968513
		return o
	},
	-1938625919: func() TLObject { // 0x8c72ea81
		o := MakeTLUserFull(nil)
		o.Data2.Constructor = -1938625919
		return o
	},
	-818518751: func() TLObject { // 0xcf366521
		o := MakeTLUserFull(nil)
		o.Data2.Constructor = -818518751
		return o
	},
	1326562017: func() TLObject { // 0x4f11bae1
		o := MakeTLUserProfilePhotoEmpty(nil)
		o.Data2.Constructor = 1326562017
		return o
	},
	-2100168954: func() TLObject { // 0x82d1f706
		o := MakeTLUserProfilePhoto(nil)
		o.Data2.Constructor = -2100168954
		return o
	},
	164646985: func() TLObject { // 0x9d05049
		o := MakeTLUserStatusEmpty(nil)
		o.Data2.Constructor = 164646985
		return o
	},
	-306628279: func() TLObject { // 0xedb93949
		o := MakeTLUserStatusOnline(nil)
		o.Data2.Constructor = -306628279
		return o
	},
	9203775: func() TLObject { // 0x8c703f
		o := MakeTLUserStatusOffline(nil)
		o.Data2.Constructor = 9203775
		return o
	},
	2065268168: func() TLObject { // 0x7b197dc8
		o := MakeTLUserStatusRecently(nil)
		o.Data2.Constructor = 2065268168
		return o
	},
	-496024847: func() TLObject { // 0xe26f42f1
		o := MakeTLUserStatusRecently(nil)
		o.Data2.Constructor = -496024847
		return o
	},
	1410997530: func() TLObject { // 0x541a1d1a
		o := MakeTLUserStatusLastWeek(nil)
		o.Data2.Constructor = 1410997530
		return o
	},
	129960444: func() TLObject { // 0x7bf09fc
		o := MakeTLUserStatusLastWeek(nil)
		o.Data2.Constructor = 129960444
		return o
	},
	1703516023: func() TLObject { // 0x65899777
		o := MakeTLUserStatusLastMonth(nil)
		o.Data2.Constructor = 1703516023
		return o
	},
	2011940674: func() TLObject { // 0x77ebc742
		o := MakeTLUserStatusLastMonth(nil)
		o.Data2.Constructor = 2011940674
		return o
	},
	-2045664768: func() TLObject { // 0x8611a200
		o := MakeTLUserStories(nil)
		o.Data2.Constructor = -2045664768
		return o
	},
	-1274595769: func() TLObject { // 0xb4073647
		o := MakeTLUsername(nil)
		o.Data2.Constructor = -1274595769
		return o
	},
	997004590: func() TLObject { // 0x3b6d152e
		o := MakeTLUsersUserFull(nil)
		o.Data2.Constructor = 997004590
		return o
	},
	-567037804: func() TLObject { // 0xde33b094
		o := MakeTLVideoSize(nil)
		o.Data2.Constructor = -567037804
		return o
	},
	-128171716: func() TLObject { // 0xf85c413c
		o := MakeTLVideoSizeEmojiMarkup(nil)
		o.Data2.Constructor = -128171716
		return o
	},
	228623102: func() TLObject { // 0xda082fe
		o := MakeTLVideoSizeStickerMarkup(nil)
		o.Data2.Constructor = 228623102
		return o
	},
	470303800: func() TLObject { // 0x1c084438
		o := MakeTLVoid(nil)
		o.Data2.Constructor = 470303800
		return o
	},
	-1539849235: func() TLObject { // 0xa437c3ed
		o := MakeTLWallPaper(nil)
		o.Data2.Constructor = -1539849235
		return o
	},
	-528465642: func() TLObject { // 0xe0804116
		o := MakeTLWallPaperNoFile(nil)
		o.Data2.Constructor = -528465642
		return o
	},
	925826256: func() TLObject { // 0x372efcd0
		o := MakeTLWallPaperSettings(nil)
		o.Data2.Constructor = 925826256
		return o
	},
	499236004: func() TLObject { // 0x1dc1bca4
		o := MakeTLWallPaperSettings(nil)
		o.Data2.Constructor = 499236004
		return o
	},
	-1493633966: func() TLObject { // 0xa6f8f452
		o := MakeTLWebAuthorization(nil)
		o.Data2.Constructor = -1493633966
		return o
	},
	475467473: func() TLObject { // 0x1c570ed1
		o := MakeTLWebDocument(nil)
		o.Data2.Constructor = 475467473
		return o
	},
	-104284986: func() TLObject { // 0xf9c8bcc6
		o := MakeTLWebDocumentNoProxy(nil)
		o.Data2.Constructor = -104284986
		return o
	},
	555358088: func() TLObject { // 0x211a1788
		o := MakeTLWebPageEmpty(nil)
		o.Data2.Constructor = 555358088
		return o
	},
	-350980120: func() TLObject { // 0xeb1477e8
		o := MakeTLWebPageEmpty(nil)
		o.Data2.Constructor = -350980120
		return o
	},
	-1328464313: func() TLObject { // 0xb0d13e47
		o := MakeTLWebPagePending(nil)
		o.Data2.Constructor = -1328464313
		return o
	},
	-981018084: func() TLObject { // 0xc586da1c
		o := MakeTLWebPagePending(nil)
		o.Data2.Constructor = -981018084
		return o
	},
	-392411726: func() TLObject { // 0xe89c45b2
		o := MakeTLWebPage(nil)
		o.Data2.Constructor = -392411726
		return o
	},
	1930545681: func() TLObject { // 0x7311ca11
		o := MakeTLWebPageNotModified(nil)
		o.Data2.Constructor = 1930545681
		return o
	},
	1421174295: func() TLObject { // 0x54b56617
		o := MakeTLWebPageAttributeTheme(nil)
		o.Data2.Constructor = 1421174295
		return o
	},
	781501415: func() TLObject { // 0x2e94c3e7
		o := MakeTLWebPageAttributeStory(nil)
		o.Data2.Constructor = 781501415
		return o
	},
	-1818605967: func() TLObject { // 0x939a4671
		o := MakeTLWebPageAttributeStory(nil)
		o.Data2.Constructor = -1818605967
		return o
	},
	1355547603: func() TLObject { // 0x50cc03d3
		o := MakeTLWebPageAttributeStickerSet(nil)
		o.Data2.Constructor = 1355547603
		return o
	},
	211046684: func() TLObject { // 0xc94511c
		o := MakeTLWebViewMessageSent(nil)
		o.Data2.Constructor = 211046684
		return o
	},
	1294139288: func() TLObject { // 0x4d22ff98
		o := MakeTLWebViewResultUrl(nil)
		o.Data2.Constructor = 1294139288
		return o
	},
	202659196: func() TLObject { // 0xc14557c
		o := MakeTLWebViewResultUrl(nil)
		o.Data2.Constructor = 202659196
		return o
	},

	// Method
	1615239032: func() TLObject { // 0x60469778
		return &TLReqPq{
			Constructor: 1615239032,
		}
	},
	-1099002127: func() TLObject { // 0xbe7e8ef1
		return &TLReqPqMulti{
			Constructor: -1099002127,
		}
	},
	-686627650: func() TLObject { // 0xd712e4be
		return &TLReq_DHParams{
			Constructor: -686627650,
		}
	},
	-184262881: func() TLObject { // 0xf5045f1f
		return &TLSetClient_DHParams{
			Constructor: -184262881,
		}
	},
	-784117408: func() TLObject { // 0xd1435160
		return &TLDestroyAuthKey{
			Constructor: -784117408,
		}
	},
	1491380032: func() TLObject { // 0x58e4a740
		return &TLRpcDropAnswer{
			Constructor: 1491380032,
		}
	},
	-1188971260: func() TLObject { // 0xb921bd04
		return &TLGetFutureSalts{
			Constructor: -1188971260,
		}
	},
	2059302892: func() TLObject { // 0x7abe77ec
		return &TLPing{
			Constructor: 2059302892,
		}
	},
	-213746804: func() TLObject { // 0xf3427b8c
		return &TLPingDelayDisconnect{
			Constructor: -213746804,
		}
	},
	-414113498: func() TLObject { // 0xe7512126
		return &TLDestroySession{
			Constructor: -414113498,
		}
	},
	-294277375: func() TLObject { // 0xee75af01
		return &TLTestUseError{
			Constructor: -294277375,
		}
	},
	-105401795: func() TLObject { // 0xf9b7b23d
		return &TLTestUseConfigSimple{
			Constructor: -105401795,
		}
	},
	-1156741135: func() TLObject { // 0xbb0d87f1
		return &TLTestParseInputAppEvent{
			Constructor: -1156741135,
		}
	},
	-878758099: func() TLObject { // 0xcb9f372d
		return &TLInvokeAfterMsg{
			Constructor: -878758099,
		}
	},
	1036301552: func() TLObject { // 0x3dc4b4f0
		return &TLInvokeAfterMsgs{
			Constructor: 1036301552,
		}
	},
	-1043505495: func() TLObject { // 0xc1cd5ea9
		return &TLInitConnection{
			Constructor: -1043505495,
		}
	},
	2018609336: func() TLObject { // 0x785188b8
		return &TLInitConnection{
			Constructor: 2018609336,
		}
	},
	-627372787: func() TLObject { // 0xda9b0d0d
		return &TLInvokeWithLayer{
			Constructor: -627372787,
		}
	},
	-1080796745: func() TLObject { // 0xbf9459b7
		return &TLInvokeWithoutUpdates{
			Constructor: -1080796745,
		}
	},
	911373810: func() TLObject { // 0x365275f2
		return &TLInvokeWithMessagesRange{
			Constructor: 911373810,
		}
	},
	-1398145746: func() TLObject { // 0xaca9fd2e
		return &TLInvokeWithTakeout{
			Constructor: -1398145746,
		}
	},
	-584540274: func() TLObject { // 0xdd289f8e
		return &TLInvokeWithBusinessConnection{
			Constructor: -584540274,
		}
	},
	502868356: func() TLObject { // 0x1df92984
		return &TLInvokeWithGooglePlayIntegrity{
			Constructor: 502868356,
		}
	},
	229528824: func() TLObject { // 0xdae54f8
		return &TLInvokeWithApnsSecret{
			Constructor: 229528824,
		}
	},
	-1502141361: func() TLObject { // 0xa677244f
		return &TLAuthSendCode{
			Constructor: -1502141361,
		}
	},
	-1429752041: func() TLObject { // 0xaac7b717
		return &TLAuthSignUp{
			Constructor: -1429752041,
		}
	},
	-2131827673: func() TLObject { // 0x80eee427
		return &TLAuthSignUp{
			Constructor: -2131827673,
		}
	},
	-1923962543: func() TLObject { // 0x8d52a951
		return &TLAuthSignIn{
			Constructor: -1923962543,
		}
	},
	-1126886015: func() TLObject { // 0xbcd51581
		return &TLAuthSignIn{
			Constructor: -1126886015,
		}
	},
	1047706137: func() TLObject { // 0x3e72ba19
		return &TLAuthLogOut{
			Constructor: 1047706137,
		}
	},
	-1616179942: func() TLObject { // 0x9fab0d1a
		return &TLAuthResetAuthorizations{
			Constructor: -1616179942,
		}
	},
	-440401971: func() TLObject { // 0xe5bfffcd
		return &TLAuthExportAuthorization{
			Constructor: -440401971,
		}
	},
	-1518699091: func() TLObject { // 0xa57a7dad
		return &TLAuthImportAuthorization{
			Constructor: -1518699091,
		}
	},
	-841733627: func() TLObject { // 0xcdd42a05
		return &TLAuthBindTempAuthKey{
			Constructor: -841733627,
		}
	},
	1738800940: func() TLObject { // 0x67a3ff2c
		return &TLAuthImportBotAuthorization{
			Constructor: 1738800940,
		}
	},
	-779399914: func() TLObject { // 0xd18b4d16
		return &TLAuthCheckPassword{
			Constructor: -779399914,
		}
	},
	-661144474: func() TLObject { // 0xd897bc66
		return &TLAuthRequestPasswordRecovery{
			Constructor: -661144474,
		}
	},
	923364464: func() TLObject { // 0x37096c70
		return &TLAuthRecoverPassword{
			Constructor: 923364464,
		}
	},
	-890997469: func() TLObject { // 0xcae47523
		return &TLAuthResendCode{
			Constructor: -890997469,
		}
	},
	1056025023: func() TLObject { // 0x3ef1a9bf
		return &TLAuthResendCode{
			Constructor: 1056025023,
		}
	},
	520357240: func() TLObject { // 0x1f040578
		return &TLAuthCancelCode{
			Constructor: 520357240,
		}
	},
	-1907842680: func() TLObject { // 0x8e48a188
		return &TLAuthDropTempAuthKeys{
			Constructor: -1907842680,
		}
	},
	-1210022402: func() TLObject { // 0xb7e085fe
		return &TLAuthExportLoginToken{
			Constructor: -1210022402,
		}
	},
	-1783866140: func() TLObject { // 0x95ac5ce4
		return &TLAuthImportLoginToken{
			Constructor: -1783866140,
		}
	},
	-392909491: func() TLObject { // 0xe894ad4d
		return &TLAuthAcceptLoginToken{
			Constructor: -392909491,
		}
	},
	221691769: func() TLObject { // 0xd36bf79
		return &TLAuthCheckRecoveryPassword{
			Constructor: 221691769,
		}
	},
	767062953: func() TLObject { // 0x2db873a9
		return &TLAuthImportWebTokenAuthorization{
			Constructor: 767062953,
		}
	},
	-1908857314: func() TLObject { // 0x8e39261e
		return &TLAuthRequestFirebaseSms{
			Constructor: -1908857314,
		}
	},
	-1991881904: func() TLObject { // 0x89464b50
		return &TLAuthRequestFirebaseSms{
			Constructor: -1991881904,
		}
	},
	2123760019: func() TLObject { // 0x7e960193
		return &TLAuthResetLoginEmail{
			Constructor: 2123760019,
		}
	},
	-878841866: func() TLObject { // 0xcb9deff6
		return &TLAuthReportMissingCode{
			Constructor: -878841866,
		}
	},
	-326762118: func() TLObject { // 0xec86017a
		return &TLAccountRegisterDevice{
			Constructor: -326762118,
		}
	},
	1669245048: func() TLObject { // 0x637ea878
		return &TLAccountRegisterDevice{
			Constructor: 1669245048,
		}
	},
	1779249670: func() TLObject { // 0x6a0d3206
		return &TLAccountUnregisterDevice{
			Constructor: 1779249670,
		}
	},
	1707432768: func() TLObject { // 0x65c55b40
		return &TLAccountUnregisterDevice{
			Constructor: 1707432768,
		}
	},
	-2067899501: func() TLObject { // 0x84be5b93
		return &TLAccountUpdateNotifySettings{
			Constructor: -2067899501,
		}
	},
	313765169: func() TLObject { // 0x12b3ad31
		return &TLAccountGetNotifySettings{
			Constructor: 313765169,
		}
	},
	-612493497: func() TLObject { // 0xdb7e1747
		return &TLAccountResetNotifySettings{
			Constructor: -612493497,
		}
	},
	2018596725: func() TLObject { // 0x78515775
		return &TLAccountUpdateProfile{
			Constructor: 2018596725,
		}
	},
	1713919532: func() TLObject { // 0x6628562c
		return &TLAccountUpdateStatus{
			Constructor: 1713919532,
		}
	},
	127302966: func() TLObject { // 0x7967d36
		return &TLAccountGetWallPapers{
			Constructor: 127302966,
		}
	},
	-977650298: func() TLObject { // 0xc5ba3d86
		return &TLAccountReportPeer{
			Constructor: -977650298,
		}
	},
	655677548: func() TLObject { // 0x2714d86c
		return &TLAccountCheckUsername{
			Constructor: 655677548,
		}
	},
	1040964988: func() TLObject { // 0x3e0bdd7c
		return &TLAccountUpdateUsername{
			Constructor: 1040964988,
		}
	},
	-623130288: func() TLObject { // 0xdadbc950
		return &TLAccountGetPrivacy{
			Constructor: -623130288,
		}
	},
	-906486552: func() TLObject { // 0xc9f81ce8
		return &TLAccountSetPrivacy{
			Constructor: -906486552,
		}
	},
	-1564422284: func() TLObject { // 0xa2c0cf74
		return &TLAccountDeleteAccount{
			Constructor: -1564422284,
		}
	},
	1099779595: func() TLObject { // 0x418d4e0b
		return &TLAccountDeleteAccount{
			Constructor: 1099779595,
		}
	},
	150761757: func() TLObject { // 0x8fc711d
		return &TLAccountGetAccountTTL{
			Constructor: 150761757,
		}
	},
	608323678: func() TLObject { // 0x2442485e
		return &TLAccountSetAccountTTL{
			Constructor: 608323678,
		}
	},
	-2108208411: func() TLObject { // 0x82574ae5
		return &TLAccountSendChangePhoneCode{
			Constructor: -2108208411,
		}
	},
	1891839707: func() TLObject { // 0x70c32edb
		return &TLAccountChangePhone{
			Constructor: 1891839707,
		}
	},
	954152242: func() TLObject { // 0x38df3532
		return &TLAccountUpdateDeviceLocked{
			Constructor: 954152242,
		}
	},
	-484392616: func() TLObject { // 0xe320c158
		return &TLAccountGetAuthorizations{
			Constructor: -484392616,
		}
	},
	-545786948: func() TLObject { // 0xdf77f3bc
		return &TLAccountResetAuthorization{
			Constructor: -545786948,
		}
	},
	1418342645: func() TLObject { // 0x548a30f5
		return &TLAccountGetPassword{
			Constructor: 1418342645,
		}
	},
	-1663767815: func() TLObject { // 0x9cd4eaf9
		return &TLAccountGetPasswordSettings{
			Constructor: -1663767815,
		}
	},
	-1516564433: func() TLObject { // 0xa59b102f
		return &TLAccountUpdatePasswordSettings{
			Constructor: -1516564433,
		}
	},
	457157256: func() TLObject { // 0x1b3faa88
		return &TLAccountSendConfirmPhoneCode{
			Constructor: 457157256,
		}
	},
	1596029123: func() TLObject { // 0x5f2178c3
		return &TLAccountConfirmPhone{
			Constructor: 1596029123,
		}
	},
	1151208273: func() TLObject { // 0x449e0b51
		return &TLAccountGetTmpPassword{
			Constructor: 1151208273,
		}
	},
	405695855: func() TLObject { // 0x182e6d6f
		return &TLAccountGetWebAuthorizations{
			Constructor: 405695855,
		}
	},
	755087855: func() TLObject { // 0x2d01b9ef
		return &TLAccountResetWebAuthorization{
			Constructor: 755087855,
		}
	},
	1747789204: func() TLObject { // 0x682d2594
		return &TLAccountResetWebAuthorizations{
			Constructor: 1747789204,
		}
	},
	-1299661699: func() TLObject { // 0xb288bc7d
		return &TLAccountGetAllSecureValues{
			Constructor: -1299661699,
		}
	},
	1936088002: func() TLObject { // 0x73665bc2
		return &TLAccountGetSecureValue{
			Constructor: 1936088002,
		}
	},
	-1986010339: func() TLObject { // 0x899fe31d
		return &TLAccountSaveSecureValue{
			Constructor: -1986010339,
		}
	},
	-1199522741: func() TLObject { // 0xb880bc4b
		return &TLAccountDeleteSecureValue{
			Constructor: -1199522741,
		}
	},
	-1456907910: func() TLObject { // 0xa929597a
		return &TLAccountGetAuthorizationForm{
			Constructor: -1456907910,
		}
	},
	-202552205: func() TLObject { // 0xf3ed4c73
		return &TLAccountAcceptAuthorization{
			Constructor: -202552205,
		}
	},
	-1516022023: func() TLObject { // 0xa5a356f9
		return &TLAccountSendVerifyPhoneCode{
			Constructor: -1516022023,
		}
	},
	1305716726: func() TLObject { // 0x4dd3a7f6
		return &TLAccountVerifyPhone{
			Constructor: 1305716726,
		}
	},
	-1730136133: func() TLObject { // 0x98e037bb
		return &TLAccountSendVerifyEmailCode{
			Constructor: -1730136133,
		}
	},
	1880182943: func() TLObject { // 0x7011509f
		return &TLAccountSendVerifyEmailCode{
			Constructor: 1880182943,
		}
	},
	53322959: func() TLObject { // 0x32da4cf
		return &TLAccountVerifyEmail32DA4CF{
			Constructor: 53322959,
		}
	},
	-1896617296: func() TLObject { // 0x8ef3eab0
		return &TLAccountInitTakeoutSession{
			Constructor: -1896617296,
		}
	},
	-262453244: func() TLObject { // 0xf05b4804
		return &TLAccountInitTakeoutSession{
			Constructor: -262453244,
		}
	},
	489050862: func() TLObject { // 0x1d2652ee
		return &TLAccountFinishTakeoutSession{
			Constructor: 489050862,
		}
	},
	-1881204448: func() TLObject { // 0x8fdf1920
		return &TLAccountConfirmPasswordEmail{
			Constructor: -1881204448,
		}
	},
	2055154197: func() TLObject { // 0x7a7f2a15
		return &TLAccountResendPasswordEmail{
			Constructor: 2055154197,
		}
	},
	-1043606090: func() TLObject { // 0xc1cbd5b6
		return &TLAccountCancelPasswordEmail{
			Constructor: -1043606090,
		}
	},
	-1626880216: func() TLObject { // 0x9f07c728
		return &TLAccountGetContactSignUpNotification{
			Constructor: -1626880216,
		}
	},
	-806076575: func() TLObject { // 0xcff43f61
		return &TLAccountSetContactSignUpNotification{
			Constructor: -806076575,
		}
	},
	1398240377: func() TLObject { // 0x53577479
		return &TLAccountGetNotifyExceptions{
			Constructor: 1398240377,
		}
	},
	-57811990: func() TLObject { // 0xfc8ddbea
		return &TLAccountGetWallPaper{
			Constructor: -57811990,
		}
	},
	-476410109: func() TLObject { // 0xe39a8f03
		return &TLAccountUploadWallPaper{
			Constructor: -476410109,
		}
	},
	-578472351: func() TLObject { // 0xdd853661
		return &TLAccountUploadWallPaper{
			Constructor: -578472351,
		}
	},
	1817860919: func() TLObject { // 0x6c5a5b37
		return &TLAccountSaveWallPaper{
			Constructor: 1817860919,
		}
	},
	-18000023: func() TLObject { // 0xfeed5769
		return &TLAccountInstallWallPaper{
			Constructor: -18000023,
		}
	},
	-1153722364: func() TLObject { // 0xbb3b9804
		return &TLAccountResetWallPapers{
			Constructor: -1153722364,
		}
	},
	1457130303: func() TLObject { // 0x56da0b3f
		return &TLAccountGetAutoDownloadSettings{
			Constructor: 1457130303,
		}
	},
	1995661875: func() TLObject { // 0x76f36233
		return &TLAccountSaveAutoDownloadSettings{
			Constructor: 1995661875,
		}
	},
	473805619: func() TLObject { // 0x1c3db333
		return &TLAccountUploadTheme{
			Constructor: 473805619,
		}
	},
	1697530880: func() TLObject { // 0x652e4400
		return &TLAccountCreateTheme{
			Constructor: 1697530880,
		}
	},
	737414348: func() TLObject { // 0x2bf40ccc
		return &TLAccountUpdateTheme{
			Constructor: 737414348,
		}
	},
	-229175188: func() TLObject { // 0xf257106c
		return &TLAccountSaveTheme{
			Constructor: -229175188,
		}
	},
	-953697477: func() TLObject { // 0xc727bb3b
		return &TLAccountInstallTheme{
			Constructor: -953697477,
		}
	},
	978872812: func() TLObject { // 0x3a5869ec
		return &TLAccountGetTheme{
			Constructor: 978872812,
		}
	},
	-1919060949: func() TLObject { // 0x8d9d742b
		return &TLAccountGetTheme{
			Constructor: -1919060949,
		}
	},
	1913054296: func() TLObject { // 0x7206e458
		return &TLAccountGetThemes{
			Constructor: 1913054296,
		}
	},
	-1250643605: func() TLObject { // 0xb574b16b
		return &TLAccountSetContentSettings{
			Constructor: -1250643605,
		}
	},
	-1952756306: func() TLObject { // 0x8b9b4dae
		return &TLAccountGetContentSettings{
			Constructor: -1952756306,
		}
	},
	1705865692: func() TLObject { // 0x65ad71dc
		return &TLAccountGetMultiWallPapers{
			Constructor: 1705865692,
		}
	},
	-349483786: func() TLObject { // 0xeb2b4cf6
		return &TLAccountGetGlobalPrivacySettings{
			Constructor: -349483786,
		}
	},
	517647042: func() TLObject { // 0x1edaaac2
		return &TLAccountSetGlobalPrivacySettings{
			Constructor: 517647042,
		}
	},
	-91437323: func() TLObject { // 0xfa8cc6f5
		return &TLAccountReportProfilePhoto{
			Constructor: -91437323,
		}
	},
	-1828139493: func() TLObject { // 0x9308ce1b
		return &TLAccountResetPassword{
			Constructor: -1828139493,
		}
	},
	1284770294: func() TLObject { // 0x4c9409f6
		return &TLAccountDeclinePasswordReset{
			Constructor: 1284770294,
		}
	},
	-700916087: func() TLObject { // 0xd638de89
		return &TLAccountGetChatThemes{
			Constructor: -700916087,
		}
	},
	-1081501024: func() TLObject { // 0xbf899aa0
		return &TLAccountSetAuthorizationTTL{
			Constructor: -1081501024,
		}
	},
	1089766498: func() TLObject { // 0x40f48462
		return &TLAccountChangeAuthorizationSettings{
			Constructor: 1089766498,
		}
	},
	-510647672: func() TLObject { // 0xe1902288
		return &TLAccountGetSavedRingtones{
			Constructor: -510647672,
		}
	},
	1038768899: func() TLObject { // 0x3dea5b03
		return &TLAccountSaveRingtone{
			Constructor: 1038768899,
		}
	},
	-2095414366: func() TLObject { // 0x831a83a2
		return &TLAccountUploadRingtone{
			Constructor: -2095414366,
		}
	},
	-70001045: func() TLObject { // 0xfbd3de6b
		return &TLAccountUpdateEmojiStatus{
			Constructor: -70001045,
		}
	},
	-696962170: func() TLObject { // 0xd6753386
		return &TLAccountGetDefaultEmojiStatuses{
			Constructor: -696962170,
		}
	},
	257392901: func() TLObject { // 0xf578105
		return &TLAccountGetRecentEmojiStatuses{
			Constructor: 257392901,
		}
	},
	404757166: func() TLObject { // 0x18201aae
		return &TLAccountClearRecentEmojiStatuses{
			Constructor: 404757166,
		}
	},
	-279966037: func() TLObject { // 0xef500eab
		return &TLAccountReorderUsernames{
			Constructor: -279966037,
		}
	},
	1490465654: func() TLObject { // 0x58d6b376
		return &TLAccountToggleUsername{
			Constructor: 1490465654,
		}
	},
	-495647960: func() TLObject { // 0xe2750328
		return &TLAccountGetDefaultProfilePhotoEmojis{
			Constructor: -495647960,
		}
	},
	-1856479058: func() TLObject { // 0x915860ae
		return &TLAccountGetDefaultGroupPhotoEmojis{
			Constructor: -1856479058,
		}
	},
	-1379156774: func() TLObject { // 0xadcbbcda
		return &TLAccountGetAutoSaveSettings{
			Constructor: -1379156774,
		}
	},
	-694451359: func() TLObject { // 0xd69b8361
		return &TLAccountSaveAutoSaveSettings{
			Constructor: -694451359,
		}
	},
	1404829728: func() TLObject { // 0x53bc0020
		return &TLAccountDeleteAutoSaveExceptions{
			Constructor: 1404829728,
		}
	},
	-896866118: func() TLObject { // 0xca8ae8ba
		return &TLAccountInvalidateSignInCodes{
			Constructor: -896866118,
		}
	},
	2096079197: func() TLObject { // 0x7cefa15d
		return &TLAccountUpdateColor{
			Constructor: 2096079197,
		}
	},
	-1610494909: func() TLObject { // 0xa001cc43
		return &TLAccountUpdateColor{
			Constructor: -1610494909,
		}
	},
	-1509246514: func() TLObject { // 0xa60ab9ce
		return &TLAccountGetDefaultBackgroundEmojis{
			Constructor: -1509246514,
		}
	},
	1999087573: func() TLObject { // 0x7727a7d5
		return &TLAccountGetChannelDefaultEmojiStatuses{
			Constructor: 1999087573,
		}
	},
	900325589: func() TLObject { // 0x35a9e0d5
		return &TLAccountGetChannelRestrictedStatusEmojis{
			Constructor: 900325589,
		}
	},
	1258348646: func() TLObject { // 0x4b00e066
		return &TLAccountUpdateBusinessWorkHours{
			Constructor: 1258348646,
		}
	},
	-1637149926: func() TLObject { // 0x9e6b131a
		return &TLAccountUpdateBusinessLocation{
			Constructor: -1637149926,
		}
	},
	1724755908: func() TLObject { // 0x66cdafc4
		return &TLAccountUpdateBusinessGreetingMessage{
			Constructor: 1724755908,
		}
	},
	-1570078811: func() TLObject { // 0xa26a7fa5
		return &TLAccountUpdateBusinessAwayMessage{
			Constructor: -1570078811,
		}
	},
	1138250269: func() TLObject { // 0x43d8521d
		return &TLAccountUpdateConnectedBot{
			Constructor: 1138250269,
		}
	},
	-1674751363: func() TLObject { // 0x9c2d527d
		return &TLAccountUpdateConnectedBot{
			Constructor: -1674751363,
		}
	},
	1319421967: func() TLObject { // 0x4ea4c80f
		return &TLAccountGetConnectedBots{
			Constructor: 1319421967,
		}
	},
	1990746736: func() TLObject { // 0x76a86270
		return &TLAccountGetBotBusinessConnection{
			Constructor: 1990746736,
		}
	},
	-1508585420: func() TLObject { // 0xa614d034
		return &TLAccountUpdateBusinessIntro{
			Constructor: -1508585420,
		}
	},
	1684934807: func() TLObject { // 0x646e1097
		return &TLAccountToggleConnectedBotPaused{
			Constructor: 1684934807,
		}
	},
	1581481689: func() TLObject { // 0x5e437ed9
		return &TLAccountDisablePeerConnectedBot{
			Constructor: 1581481689,
		}
	},
	-865203183: func() TLObject { // 0xcc6e0c11
		return &TLAccountUpdateBirthday{
			Constructor: -865203183,
		}
	},
	-2007898482: func() TLObject { // 0x8851e68e
		return &TLAccountCreateBusinessChatLink{
			Constructor: -2007898482,
		}
	},
	-1942744913: func() TLObject { // 0x8c3410af
		return &TLAccountEditBusinessChatLink{
			Constructor: -1942744913,
		}
	},
	1611085428: func() TLObject { // 0x60073674
		return &TLAccountDeleteBusinessChatLink{
			Constructor: 1611085428,
		}
	},
	1869667809: func() TLObject { // 0x6f70dde1
		return &TLAccountGetBusinessChatLinks{
			Constructor: 1869667809,
		}
	},
	1418913262: func() TLObject { // 0x5492e5ee
		return &TLAccountResolveBusinessChatLink{
			Constructor: 1418913262,
		}
	},
	-649919008: func() TLObject { // 0xd94305e0
		return &TLAccountUpdatePersonalChannel{
			Constructor: -649919008,
		}
	},
	-1176919155: func() TLObject { // 0xb9d9a38d
		return &TLAccountToggleSponsoredMessages{
			Constructor: -1176919155,
		}
	},
	115172684: func() TLObject { // 0x6dd654c
		return &TLAccountGetReactionsNotifySettings{
			Constructor: 115172684,
		}
	},
	829220168: func() TLObject { // 0x316ce548
		return &TLAccountSetReactionsNotifySettings{
			Constructor: 829220168,
		}
	},
	227648840: func() TLObject { // 0xd91a548
		return &TLUsersGetUsers{
			Constructor: 227648840,
		}
	},
	-1240508136: func() TLObject { // 0xb60f5918
		return &TLUsersGetFullUser{
			Constructor: -1240508136,
		}
	},
	-1865902923: func() TLObject { // 0x90c894b5
		return &TLUsersSetSecureValueErrors{
			Constructor: -1865902923,
		}
	},
	-1507677680: func() TLObject { // 0xa622aa10
		return &TLUsersGetIsPremiumRequiredToContact{
			Constructor: -1507677680,
		}
	},
	2061264541: func() TLObject { // 0x7adc669d
		return &TLContactsGetContactIDs{
			Constructor: 2061264541,
		}
	},
	-995929106: func() TLObject { // 0xc4a353ee
		return &TLContactsGetStatuses{
			Constructor: -995929106,
		}
	},
	1574346258: func() TLObject { // 0x5dd69e12
		return &TLContactsGetContacts{
			Constructor: 1574346258,
		}
	},
	746589157: func() TLObject { // 0x2c800be5
		return &TLContactsImportContacts{
			Constructor: 746589157,
		}
	},
	157945344: func() TLObject { // 0x96a0e00
		return &TLContactsDeleteContacts{
			Constructor: 157945344,
		}
	},
	269745566: func() TLObject { // 0x1013fd9e
		return &TLContactsDeleteByPhones{
			Constructor: 269745566,
		}
	},
	774801204: func() TLObject { // 0x2e2e8734
		return &TLContactsBlock{
			Constructor: 774801204,
		}
	},
	1758204945: func() TLObject { // 0x68cc1411
		return &TLContactsBlock{
			Constructor: 1758204945,
		}
	},
	-1252994264: func() TLObject { // 0xb550d328
		return &TLContactsUnblock{
			Constructor: -1252994264,
		}
	},
	-1096393392: func() TLObject { // 0xbea65d50
		return &TLContactsUnblock{
			Constructor: -1096393392,
		}
	},
	-1702457472: func() TLObject { // 0x9a868f80
		return &TLContactsGetBlocked{
			Constructor: -1702457472,
		}
	},
	-176409329: func() TLObject { // 0xf57c350f
		return &TLContactsGetBlocked{
			Constructor: -176409329,
		}
	},
	301470424: func() TLObject { // 0x11f812d8
		return &TLContactsSearch{
			Constructor: 301470424,
		}
	},
	-113456221: func() TLObject { // 0xf93ccba3
		return &TLContactsResolveUsername{
			Constructor: -113456221,
		}
	},
	-1758168906: func() TLObject { // 0x973478b6
		return &TLContactsGetTopPeers{
			Constructor: -1758168906,
		}
	},
	451113900: func() TLObject { // 0x1ae373ac
		return &TLContactsResetTopPeerRating{
			Constructor: 451113900,
		}
	},
	-2020263951: func() TLObject { // 0x879537f1
		return &TLContactsResetSaved{
			Constructor: -2020263951,
		}
	},
	-2098076769: func() TLObject { // 0x82f1e39f
		return &TLContactsGetSaved{
			Constructor: -2098076769,
		}
	},
	-2062238246: func() TLObject { // 0x8514bdda
		return &TLContactsToggleTopPeers{
			Constructor: -2062238246,
		}
	},
	-386636848: func() TLObject { // 0xe8f463d0
		return &TLContactsAddContact{
			Constructor: -386636848,
		}
	},
	-130964977: func() TLObject { // 0xf831a20f
		return &TLContactsAcceptContact{
			Constructor: -130964977,
		}
	},
	-750207932: func() TLObject { // 0xd348bc44
		return &TLContactsGetLocated{
			Constructor: -750207932,
		}
	},
	698914348: func() TLObject { // 0x29a8962c
		return &TLContactsBlockFromReplies{
			Constructor: 698914348,
		}
	},
	-1963375804: func() TLObject { // 0x8af94344
		return &TLContactsResolvePhone{
			Constructor: -1963375804,
		}
	},
	-127582169: func() TLObject { // 0xf8654027
		return &TLContactsExportContactToken{
			Constructor: -127582169,
		}
	},
	318789512: func() TLObject { // 0x13005788
		return &TLContactsImportContactToken{
			Constructor: 318789512,
		}
	},
	-1167653392: func() TLObject { // 0xba6705f0
		return &TLContactsEditCloseFriends{
			Constructor: -1167653392,
		}
	},
	-1798939530: func() TLObject { // 0x94c65c76
		return &TLContactsSetBlocked{
			Constructor: -1798939530,
		}
	},
	-621959068: func() TLObject { // 0xdaeda864
		return &TLContactsGetBirthdays{
			Constructor: -621959068,
		}
	},
	1673946374: func() TLObject { // 0x63c66506
		return &TLMessagesGetMessages{
			Constructor: 1673946374,
		}
	},
	1109588596: func() TLObject { // 0x4222fa74
		return &TLMessagesGetMessages{
			Constructor: 1109588596,
		}
	},
	-1594569905: func() TLObject { // 0xa0f4cb4f
		return &TLMessagesGetDialogs{
			Constructor: -1594569905,
		}
	},
	1143203525: func() TLObject { // 0x4423e6c5
		return &TLMessagesGetHistory{
			Constructor: 1143203525,
		}
	},
	703497338: func() TLObject { // 0x29ee847a
		return &TLMessagesSearch{
			Constructor: 703497338,
		}
	},
	-1481316055: func() TLObject { // 0xa7b4e929
		return &TLMessagesSearch{
			Constructor: -1481316055,
		}
	},
	-1593989278: func() TLObject { // 0xa0fda762
		return &TLMessagesSearch{
			Constructor: -1593989278,
		}
	},
	238054714: func() TLObject { // 0xe306d3a
		return &TLMessagesReadHistory{
			Constructor: 238054714,
		}
	},
	-1332768214: func() TLObject { // 0xb08f922a
		return &TLMessagesDeleteHistory{
			Constructor: -1332768214,
		}
	},
	-443640366: func() TLObject { // 0xe58e95d2
		return &TLMessagesDeleteMessages{
			Constructor: -443640366,
		}
	},
	94983360: func() TLObject { // 0x5a954c0
		return &TLMessagesReceivedMessages{
			Constructor: 94983360,
		}
	},
	1486110434: func() TLObject { // 0x58943ee2
		return &TLMessagesSetTyping{
			Constructor: 1486110434,
		}
	},
	-1740662971: func() TLObject { // 0x983f9745
		return &TLMessagesSendMessage{
			Constructor: -1740662971,
		}
	},
	-537394132: func() TLObject { // 0xdff8042c
		return &TLMessagesSendMessage{
			Constructor: -537394132,
		}
	},
	671943023: func() TLObject { // 0x280d096f
		return &TLMessagesSendMessage{
			Constructor: 671943023,
		}
	},
	482476935: func() TLObject { // 0x1cc20387
		return &TLMessagesSendMessage{
			Constructor: 482476935,
		}
	},
	228423076: func() TLObject { // 0xd9d75a4
		return &TLMessagesSendMessage{
			Constructor: 228423076,
		}
	},
	2018673486: func() TLObject { // 0x7852834e
		return &TLMessagesSendMedia{
			Constructor: 2018673486,
		}
	},
	2077646913: func() TLObject { // 0x7bd66041
		return &TLMessagesSendMedia{
			Constructor: 2077646913,
		}
	},
	1926021693: func() TLObject { // 0x72ccc23d
		return &TLMessagesSendMedia{
			Constructor: 1926021693,
		}
	},
	1967638886: func() TLObject { // 0x7547c966
		return &TLMessagesSendMedia{
			Constructor: 1967638886,
		}
	},
	-497026848: func() TLObject { // 0xe25ff8e0
		return &TLMessagesSendMedia{
			Constructor: -497026848,
		}
	},
	-721186296: func() TLObject { // 0xd5039208
		return &TLMessagesForwardMessages{
			Constructor: -721186296,
		}
	},
	-966673468: func() TLObject { // 0xc661bbc4
		return &TLMessagesForwardMessages{
			Constructor: -966673468,
		}
	},
	-869258997: func() TLObject { // 0xcc30290b
		return &TLMessagesForwardMessages{
			Constructor: -869258997,
		}
	},
	-820669733: func() TLObject { // 0xcf1592db
		return &TLMessagesReportSpam{
			Constructor: -820669733,
		}
	},
	-270948702: func() TLObject { // 0xefd9a6a2
		return &TLMessagesGetPeerSettings{
			Constructor: -270948702,
		}
	},
	-1991005362: func() TLObject { // 0x8953ab4e
		return &TLMessagesReport{
			Constructor: -1991005362,
		}
	},
	1240027791: func() TLObject { // 0x49e9528f
		return &TLMessagesGetChats{
			Constructor: 1240027791,
		}
	},
	-1364194508: func() TLObject { // 0xaeb00b34
		return &TLMessagesGetFullChat{
			Constructor: -1364194508,
		}
	},
	1937260541: func() TLObject { // 0x73783ffd
		return &TLMessagesEditChatTitle{
			Constructor: 1937260541,
		}
	},
	903730804: func() TLObject { // 0x35ddd674
		return &TLMessagesEditChatPhoto{
			Constructor: 903730804,
		}
	},
	-876162809: func() TLObject { // 0xcbc6d107
		return &TLMessagesAddChatUserCBC6D107{
			Constructor: -876162809,
		}
	},
	-1575461717: func() TLObject { // 0xa2185cab
		return &TLMessagesDeleteChatUser{
			Constructor: -1575461717,
		}
	},
	-1831936556: func() TLObject { // 0x92ceddd4
		return &TLMessagesCreateChat92CEDDD4{
			Constructor: -1831936556,
		}
	},
	651135312: func() TLObject { // 0x26cf8950
		return &TLMessagesGetDhConfig{
			Constructor: 651135312,
		}
	},
	-162681021: func() TLObject { // 0xf64daf43
		return &TLMessagesRequestEncryption{
			Constructor: -162681021,
		}
	},
	1035731989: func() TLObject { // 0x3dbc0415
		return &TLMessagesAcceptEncryption{
			Constructor: 1035731989,
		}
	},
	-208425312: func() TLObject { // 0xf393aea0
		return &TLMessagesDiscardEncryption{
			Constructor: -208425312,
		}
	},
	2031374829: func() TLObject { // 0x791451ed
		return &TLMessagesSetEncryptedTyping{
			Constructor: 2031374829,
		}
	},
	2135648522: func() TLObject { // 0x7f4b690a
		return &TLMessagesReadEncryptedHistory{
			Constructor: 2135648522,
		}
	},
	1157265941: func() TLObject { // 0x44fa7a15
		return &TLMessagesSendEncrypted{
			Constructor: 1157265941,
		}
	},
	1431914525: func() TLObject { // 0x5559481d
		return &TLMessagesSendEncryptedFile{
			Constructor: 1431914525,
		}
	},
	852769188: func() TLObject { // 0x32d439a4
		return &TLMessagesSendEncryptedService{
			Constructor: 852769188,
		}
	},
	1436924774: func() TLObject { // 0x55a5bb66
		return &TLMessagesReceivedQueue{
			Constructor: 1436924774,
		}
	},
	1259113487: func() TLObject { // 0x4b0c8c0f
		return &TLMessagesReportEncryptedSpam{
			Constructor: 1259113487,
		}
	},
	916930423: func() TLObject { // 0x36a73f77
		return &TLMessagesReadMessageContents{
			Constructor: 916930423,
		}
	},
	-710552671: func() TLObject { // 0xd5a5d3a1
		return &TLMessagesGetStickers{
			Constructor: -710552671,
		}
	},
	-1197432408: func() TLObject { // 0xb8a0a1a8
		return &TLMessagesGetAllStickers{
			Constructor: -1197432408,
		}
	},
	-1956073268: func() TLObject { // 0x8b68b0cc
		return &TLMessagesGetWebPagePreview{
			Constructor: -1956073268,
		}
	},
	-1607670315: func() TLObject { // 0xa02ce5d5
		return &TLMessagesExportChatInvite{
			Constructor: -1607670315,
		}
	},
	1051570619: func() TLObject { // 0x3eadb1bb
		return &TLMessagesCheckChatInvite{
			Constructor: 1051570619,
		}
	},
	1817183516: func() TLObject { // 0x6c50051c
		return &TLMessagesImportChatInvite{
			Constructor: 1817183516,
		}
	},
	-928977804: func() TLObject { // 0xc8a0ec74
		return &TLMessagesGetStickerSet{
			Constructor: -928977804,
		}
	},
	639215886: func() TLObject { // 0x2619a90e
		return &TLMessagesGetStickerSet{
			Constructor: 639215886,
		}
	},
	-946871200: func() TLObject { // 0xc78fe460
		return &TLMessagesInstallStickerSet{
			Constructor: -946871200,
		}
	},
	-110209570: func() TLObject { // 0xf96e55de
		return &TLMessagesUninstallStickerSet{
			Constructor: -110209570,
		}
	},
	-421563528: func() TLObject { // 0xe6df7378
		return &TLMessagesStartBot{
			Constructor: -421563528,
		}
	},
	1468322785: func() TLObject { // 0x5784d3e1
		return &TLMessagesGetMessagesViews{
			Constructor: 1468322785,
		}
	},
	-1470377534: func() TLObject { // 0xa85bd1c2
		return &TLMessagesEditChatAdmin{
			Constructor: -1470377534,
		}
	},
	-1568189671: func() TLObject { // 0xa2875319
		return &TLMessagesMigrateChat{
			Constructor: -1568189671,
		}
	},
	1271290010: func() TLObject { // 0x4bc6589a
		return &TLMessagesSearchGlobal{
			Constructor: 1271290010,
		}
	},
	2016638777: func() TLObject { // 0x78337739
		return &TLMessagesReorderStickerSets{
			Constructor: 2016638777,
		}
	},
	-1309538785: func() TLObject { // 0xb1f2061f
		return &TLMessagesGetDocumentByHash{
			Constructor: -1309538785,
		}
	},
	864953444: func() TLObject { // 0x338e2464
		return &TLMessagesGetDocumentByHash{
			Constructor: 864953444,
		}
	},
	1559270965: func() TLObject { // 0x5cf09635
		return &TLMessagesGetSavedGifs{
			Constructor: 1559270965,
		}
	},
	846868683: func() TLObject { // 0x327a30cb
		return &TLMessagesSaveGif{
			Constructor: 846868683,
		}
	},
	1364105629: func() TLObject { // 0x514e999d
		return &TLMessagesGetInlineBotResults{
			Constructor: 1364105629,
		}
	},
	-1156406247: func() TLObject { // 0xbb12a419
		return &TLMessagesSetInlineBotResults{
			Constructor: -1156406247,
		}
	},
	-346119674: func() TLObject { // 0xeb5ea206
		return &TLMessagesSetInlineBotResults{
			Constructor: -346119674,
		}
	},
	1052698730: func() TLObject { // 0x3ebee86a
		return &TLMessagesSendInlineBotResult{
			Constructor: 1052698730,
		}
	},
	-138647366: func() TLObject { // 0xf7bc68ba
		return &TLMessagesSendInlineBotResult{
			Constructor: -138647366,
		}
	},
	-738468661: func() TLObject { // 0xd3fbdccb
		return &TLMessagesSendInlineBotResult{
			Constructor: -738468661,
		}
	},
	2057376407: func() TLObject { // 0x7aa11297
		return &TLMessagesSendInlineBotResult{
			Constructor: 2057376407,
		}
	},
	-39416522: func() TLObject { // 0xfda68d36
		return &TLMessagesGetMessageEditData{
			Constructor: -39416522,
		}
	},
	-539934715: func() TLObject { // 0xdfd14005
		return &TLMessagesEditMessage{
			Constructor: -539934715,
		}
	},
	1224152952: func() TLObject { // 0x48f71778
		return &TLMessagesEditMessage{
			Constructor: 1224152952,
		}
	},
	-2091549254: func() TLObject { // 0x83557dba
		return &TLMessagesEditInlineBotMessage{
			Constructor: -2091549254,
		}
	},
	-1824339449: func() TLObject { // 0x9342ca07
		return &TLMessagesGetBotCallbackAnswer{
			Constructor: -1824339449,
		}
	},
	-712043766: func() TLObject { // 0xd58f130a
		return &TLMessagesSetBotCallbackAnswer{
			Constructor: -712043766,
		}
	},
	-462373635: func() TLObject { // 0xe470bcfd
		return &TLMessagesGetPeerDialogs{
			Constructor: -462373635,
		}
	},
	-747452978: func() TLObject { // 0xd372c5ce
		return &TLMessagesSaveDraft{
			Constructor: -747452978,
		}
	},
	2146678790: func() TLObject { // 0x7ff3b806
		return &TLMessagesSaveDraft{
			Constructor: 2146678790,
		}
	},
	-1271718337: func() TLObject { // 0xb4331e3f
		return &TLMessagesSaveDraft{
			Constructor: -1271718337,
		}
	},
	-1137057461: func() TLObject { // 0xbc39e14b
		return &TLMessagesSaveDraft{
			Constructor: -1137057461,
		}
	},
	1782549861: func() TLObject { // 0x6a3f8d65
		return &TLMessagesGetAllDrafts{
			Constructor: 1782549861,
		}
	},
	1685588756: func() TLObject { // 0x64780b14
		return &TLMessagesGetFeaturedStickers{
			Constructor: 1685588756,
		}
	},
	1527873830: func() TLObject { // 0x5b118126
		return &TLMessagesReadFeaturedStickers{
			Constructor: 1527873830,
		}
	},
	-1649852357: func() TLObject { // 0x9da9403b
		return &TLMessagesGetRecentStickers{
			Constructor: -1649852357,
		}
	},
	958863608: func() TLObject { // 0x392718f8
		return &TLMessagesSaveRecentSticker{
			Constructor: 958863608,
		}
	},
	-1986437075: func() TLObject { // 0x8999602d
		return &TLMessagesClearRecentStickers{
			Constructor: -1986437075,
		}
	},
	1475442322: func() TLObject { // 0x57f17692
		return &TLMessagesGetArchivedStickers{
			Constructor: 1475442322,
		}
	},
	1678738104: func() TLObject { // 0x640f82b8
		return &TLMessagesGetMaskStickers{
			Constructor: 1678738104,
		}
	},
	-866424884: func() TLObject { // 0xcc5b67cc
		return &TLMessagesGetAttachedStickers{
			Constructor: -866424884,
		}
	},
	-1896289088: func() TLObject { // 0x8ef8ecc0
		return &TLMessagesSetGameScore{
			Constructor: -1896289088,
		}
	},
	363700068: func() TLObject { // 0x15ad9f64
		return &TLMessagesSetInlineGameScore{
			Constructor: 363700068,
		}
	},
	-400399203: func() TLObject { // 0xe822649d
		return &TLMessagesGetGameHighScores{
			Constructor: -400399203,
		}
	},
	258170395: func() TLObject { // 0xf635e1b
		return &TLMessagesGetInlineGameHighScores{
			Constructor: 258170395,
		}
	},
	-468934396: func() TLObject { // 0xe40ca104
		return &TLMessagesGetCommonChats{
			Constructor: -468934396,
		}
	},
	-1919511901: func() TLObject { // 0x8d9692a3
		return &TLMessagesGetWebPage8D9692A3{
			Constructor: -1919511901,
		}
	},
	-1489903017: func() TLObject { // 0xa731e257
		return &TLMessagesToggleDialogPin{
			Constructor: -1489903017,
		}
	},
	991616823: func() TLObject { // 0x3b1adf37
		return &TLMessagesReorderPinnedDialogs{
			Constructor: 991616823,
		}
	},
	-692498958: func() TLObject { // 0xd6b94df2
		return &TLMessagesGetPinnedDialogs{
			Constructor: -692498958,
		}
	},
	-436833542: func() TLObject { // 0xe5f672fa
		return &TLMessagesSetBotShippingResults{
			Constructor: -436833542,
		}
	},
	163765653: func() TLObject { // 0x9c2dd95
		return &TLMessagesSetBotPrecheckoutResults{
			Constructor: 163765653,
		}
	},
	345405816: func() TLObject { // 0x14967978
		return &TLMessagesUploadMedia{
			Constructor: 345405816,
		}
	},
	1369162417: func() TLObject { // 0x519bc2b1
		return &TLMessagesUploadMedia{
			Constructor: 1369162417,
		}
	},
	-1589618665: func() TLObject { // 0xa1405817
		return &TLMessagesSendScreenshotNotification{
			Constructor: -1589618665,
		}
	},
	-914493408: func() TLObject { // 0xc97df020
		return &TLMessagesSendScreenshotNotification{
			Constructor: -914493408,
		}
	},
	82946729: func() TLObject { // 0x4f1aaa9
		return &TLMessagesGetFavedStickers{
			Constructor: 82946729,
		}
	},
	-1174420133: func() TLObject { // 0xb9ffc55b
		return &TLMessagesFaveSticker{
			Constructor: -1174420133,
		}
	},
	-251140208: func() TLObject { // 0xf107e790
		return &TLMessagesGetUnreadMentions{
			Constructor: -251140208,
		}
	},
	1180140658: func() TLObject { // 0x46578472
		return &TLMessagesGetUnreadMentions{
			Constructor: 1180140658,
		}
	},
	921026381: func() TLObject { // 0x36e5bf4d
		return &TLMessagesReadMentions{
			Constructor: 921026381,
		}
	},
	251759059: func() TLObject { // 0xf0189d3
		return &TLMessagesReadMentions{
			Constructor: 251759059,
		}
	},
	1881817312: func() TLObject { // 0x702a40e0
		return &TLMessagesGetRecentLocations{
			Constructor: 1881817312,
		}
	},
	934757205: func() TLObject { // 0x37b74355
		return &TLMessagesSendMultiMedia{
			Constructor: 934757205,
		}
	},
	211175177: func() TLObject { // 0xc964709
		return &TLMessagesSendMultiMedia{
			Constructor: 211175177,
		}
	},
	1164872071: func() TLObject { // 0x456e8987
		return &TLMessagesSendMultiMedia{
			Constructor: 1164872071,
		}
	},
	-1225713124: func() TLObject { // 0xb6f11a1c
		return &TLMessagesSendMultiMedia{
			Constructor: -1225713124,
		}
	},
	-134016113: func() TLObject { // 0xf803138f
		return &TLMessagesSendMultiMedia{
			Constructor: -134016113,
		}
	},
	1347929239: func() TLObject { // 0x5057c497
		return &TLMessagesUploadEncryptedFile{
			Constructor: 1347929239,
		}
	},
	896555914: func() TLObject { // 0x35705b8a
		return &TLMessagesSearchStickerSets{
			Constructor: 896555914,
		}
	},
	486505992: func() TLObject { // 0x1cff7e08
		return &TLMessagesGetSplitRanges{
			Constructor: 486505992,
		}
	},
	-1031349873: func() TLObject { // 0xc286d98f
		return &TLMessagesMarkDialogUnread{
			Constructor: -1031349873,
		}
	},
	585256482: func() TLObject { // 0x22e24e22
		return &TLMessagesGetDialogUnreadMarks{
			Constructor: 585256482,
		}
	},
	2119757468: func() TLObject { // 0x7e58ee9c
		return &TLMessagesClearAllDrafts{
			Constructor: 2119757468,
		}
	},
	-760547348: func() TLObject { // 0xd2aaf7ec
		return &TLMessagesUpdatePinnedMessage{
			Constructor: -760547348,
		}
	},
	283795844: func() TLObject { // 0x10ea6184
		return &TLMessagesSendVote{
			Constructor: 283795844,
		}
	},
	1941660731: func() TLObject { // 0x73bb643b
		return &TLMessagesGetPollResults{
			Constructor: 1941660731,
		}
	},
	1848369232: func() TLObject { // 0x6e2be050
		return &TLMessagesGetOnlines{
			Constructor: 1848369232,
		}
	},
	-554301545: func() TLObject { // 0xdef60797
		return &TLMessagesEditChatAbout{
			Constructor: -554301545,
		}
	},
	-1517917375: func() TLObject { // 0xa5866b41
		return &TLMessagesEditChatDefaultBannedRights{
			Constructor: -1517917375,
		}
	},
	899735650: func() TLObject { // 0x35a0e062
		return &TLMessagesGetEmojiKeywords{
			Constructor: 899735650,
		}
	},
	352892591: func() TLObject { // 0x1508b6af
		return &TLMessagesGetEmojiKeywordsDifference{
			Constructor: 352892591,
		}
	},
	1318675378: func() TLObject { // 0x4e9963b2
		return &TLMessagesGetEmojiKeywordsLanguages{
			Constructor: 1318675378,
		}
	},
	-709817306: func() TLObject { // 0xd5b10c26
		return &TLMessagesGetEmojiURL{
			Constructor: -709817306,
		}
	},
	465367808: func() TLObject { // 0x1bbcf300
		return &TLMessagesGetSearchCounters{
			Constructor: 465367808,
		}
	},
	11435201: func() TLObject { // 0xae7cc1
		return &TLMessagesGetSearchCounters{
			Constructor: 11435201,
		}
	},
	1932455680: func() TLObject { // 0x732eef00
		return &TLMessagesGetSearchCounters{
			Constructor: 1932455680,
		}
	},
	428848198: func() TLObject { // 0x198fb446
		return &TLMessagesRequestUrlAuth{
			Constructor: 428848198,
		}
	},
	-1322487515: func() TLObject { // 0xb12c7125
		return &TLMessagesAcceptUrlAuth{
			Constructor: -1322487515,
		}
	},
	1336717624: func() TLObject { // 0x4facb138
		return &TLMessagesHidePeerSettingsBar{
			Constructor: 1336717624,
		}
	},
	-183077365: func() TLObject { // 0xf516760b
		return &TLMessagesGetScheduledHistory{
			Constructor: -183077365,
		}
	},
	-1111817116: func() TLObject { // 0xbdbb0464
		return &TLMessagesGetScheduledMessages{
			Constructor: -1111817116,
		}
	},
	-1120369398: func() TLObject { // 0xbd38850a
		return &TLMessagesSendScheduledMessages{
			Constructor: -1120369398,
		}
	},
	1504586518: func() TLObject { // 0x59ae2b16
		return &TLMessagesDeleteScheduledMessages{
			Constructor: 1504586518,
		}
	},
	-1200736242: func() TLObject { // 0xb86e380e
		return &TLMessagesGetPollVotes{
			Constructor: -1200736242,
		}
	},
	-1257951254: func() TLObject { // 0xb5052fea
		return &TLMessagesToggleStickerSets{
			Constructor: -1257951254,
		}
	},
	-271283063: func() TLObject { // 0xefd48c89
		return &TLMessagesGetDialogFiltersEFD48C89{
			Constructor: -271283063,
		}
	},
	-1566780372: func() TLObject { // 0xa29cd42c
		return &TLMessagesGetSuggestedDialogFilters{
			Constructor: -1566780372,
		}
	},
	450142282: func() TLObject { // 0x1ad4a04a
		return &TLMessagesUpdateDialogFilter{
			Constructor: 450142282,
		}
	},
	-983318044: func() TLObject { // 0xc563c1e4
		return &TLMessagesUpdateDialogFiltersOrder{
			Constructor: -983318044,
		}
	},
	2127598753: func() TLObject { // 0x7ed094a1
		return &TLMessagesGetOldFeaturedStickers{
			Constructor: 2127598753,
		}
	},
	584962828: func() TLObject { // 0x22ddd30c
		return &TLMessagesGetReplies{
			Constructor: 584962828,
		}
	},
	1147761405: func() TLObject { // 0x446972fd
		return &TLMessagesGetDiscussionMessage{
			Constructor: 1147761405,
		}
	},
	-147740172: func() TLObject { // 0xf731a9f4
		return &TLMessagesReadDiscussion{
			Constructor: -147740172,
		}
	},
	-299714136: func() TLObject { // 0xee22b9a8
		return &TLMessagesUnpinAllMessages{
			Constructor: -299714136,
		}
	},
	-265962357: func() TLObject { // 0xf025bc8b
		return &TLMessagesUnpinAllMessages{
			Constructor: -265962357,
		}
	},
	1540419152: func() TLObject { // 0x5bd0ee50
		return &TLMessagesDeleteChat{
			Constructor: 1540419152,
		}
	},
	-104078327: func() TLObject { // 0xf9cbe409
		return &TLMessagesDeletePhoneCallHistory{
			Constructor: -104078327,
		}
	},
	1140726259: func() TLObject { // 0x43fe19f3
		return &TLMessagesCheckHistoryImport{
			Constructor: 1140726259,
		}
	},
	873008187: func() TLObject { // 0x34090c3b
		return &TLMessagesInitHistoryImport{
			Constructor: 873008187,
		}
	},
	713433234: func() TLObject { // 0x2a862092
		return &TLMessagesUploadImportedMedia{
			Constructor: 713433234,
		}
	},
	-1271008444: func() TLObject { // 0xb43df344
		return &TLMessagesStartHistoryImport{
			Constructor: -1271008444,
		}
	},
	-1565154314: func() TLObject { // 0xa2b5a3f6
		return &TLMessagesGetExportedChatInvites{
			Constructor: -1565154314,
		}
	},
	1937010524: func() TLObject { // 0x73746f5c
		return &TLMessagesGetExportedChatInvite{
			Constructor: 1937010524,
		}
	},
	-1110823051: func() TLObject { // 0xbdca2f75
		return &TLMessagesEditExportedChatInvite{
			Constructor: -1110823051,
		}
	},
	1452833749: func() TLObject { // 0x56987bd5
		return &TLMessagesDeleteRevokedExportedChatInvites{
			Constructor: 1452833749,
		}
	},
	-731601877: func() TLObject { // 0xd464a42b
		return &TLMessagesDeleteExportedChatInvite{
			Constructor: -731601877,
		}
	},
	958457583: func() TLObject { // 0x3920e6ef
		return &TLMessagesGetAdminsWithInvites{
			Constructor: 958457583,
		}
	},
	-553329330: func() TLObject { // 0xdf04dd4e
		return &TLMessagesGetChatInviteImporters{
			Constructor: -553329330,
		}
	},
	-1207017500: func() TLObject { // 0xb80e5fe4
		return &TLMessagesSetHistoryTTL{
			Constructor: -1207017500,
		}
	},
	1573261059: func() TLObject { // 0x5dc60f03
		return &TLMessagesCheckHistoryImportPeer{
			Constructor: 1573261059,
		}
	},
	-432283329: func() TLObject { // 0xe63be13f
		return &TLMessagesSetChatTheme{
			Constructor: -432283329,
		}
	},
	834782287: func() TLObject { // 0x31c1c44f
		return &TLMessagesGetMessageReadParticipants31C1C44F{
			Constructor: 834782287,
		}
	},
	1789130429: func() TLObject { // 0x6aa3f6bd
		return &TLMessagesGetSearchResultsCalendar{
			Constructor: 1789130429,
		}
	},
	1240514025: func() TLObject { // 0x49f0bde9
		return &TLMessagesGetSearchResultsCalendar{
			Constructor: 1240514025,
		}
	},
	-1669386480: func() TLObject { // 0x9c7f2f10
		return &TLMessagesGetSearchResultsPositions{
			Constructor: -1669386480,
		}
	},
	1855292323: func() TLObject { // 0x6e9583a3
		return &TLMessagesGetSearchResultsPositions{
			Constructor: 1855292323,
		}
	},
	2145904661: func() TLObject { // 0x7fe7e815
		return &TLMessagesHideChatJoinRequest{
			Constructor: 2145904661,
		}
	},
	-528091926: func() TLObject { // 0xe085f4ea
		return &TLMessagesHideAllChatJoinRequests{
			Constructor: -528091926,
		}
	},
	-1323389022: func() TLObject { // 0xb11eafa2
		return &TLMessagesToggleNoForwards{
			Constructor: -1323389022,
		}
	},
	-855777386: func() TLObject { // 0xccfddf96
		return &TLMessagesSaveDefaultSendAs{
			Constructor: -855777386,
		}
	},
	-754091820: func() TLObject { // 0xd30d78d4
		return &TLMessagesSendReaction{
			Constructor: -754091820,
		}
	},
	627641572: func() TLObject { // 0x25690ce4
		return &TLMessagesSendReaction{
			Constructor: 627641572,
		}
	},
	-1950707482: func() TLObject { // 0x8bba90e6
		return &TLMessagesGetMessagesReactions{
			Constructor: -1950707482,
		}
	},
	1176190792: func() TLObject { // 0x461b3f48
		return &TLMessagesGetMessageReactionsList{
			Constructor: 1176190792,
		}
	},
	-521245833: func() TLObject { // 0xe0ee6b77
		return &TLMessagesGetMessageReactionsList{
			Constructor: -521245833,
		}
	},
	1511328724: func() TLObject { // 0x5a150bd4
		return &TLMessagesSetChatAvailableReactions{
			Constructor: 1511328724,
		}
	},
	-21928079: func() TLObject { // 0xfeb16771
		return &TLMessagesSetChatAvailableReactions{
			Constructor: -21928079,
		}
	},
	335875750: func() TLObject { // 0x14050ea6
		return &TLMessagesSetChatAvailableReactions{
			Constructor: 335875750,
		}
	},
	417243308: func() TLObject { // 0x18dea0ac
		return &TLMessagesGetAvailableReactions{
			Constructor: 417243308,
		}
	},
	1330094102: func() TLObject { // 0x4f47a016
		return &TLMessagesSetDefaultReaction{
			Constructor: 1330094102,
		}
	},
	-647969580: func() TLObject { // 0xd960c4d4
		return &TLMessagesSetDefaultReaction{
			Constructor: -647969580,
		}
	},
	1662529584: func() TLObject { // 0x63183030
		return &TLMessagesTranslateText{
			Constructor: 1662529584,
		}
	},
	617508334: func() TLObject { // 0x24ce6dee
		return &TLMessagesTranslateText{
			Constructor: 617508334,
		}
	},
	841173339: func() TLObject { // 0x3223495b
		return &TLMessagesGetUnreadReactions{
			Constructor: 841173339,
		}
	},
	-396644838: func() TLObject { // 0xe85bae1a
		return &TLMessagesGetUnreadReactions{
			Constructor: -396644838,
		}
	},
	1420459918: func() TLObject { // 0x54aa7f8e
		return &TLMessagesReadReactions{
			Constructor: 1420459918,
		}
	},
	-2099097129: func() TLObject { // 0x82e251d7
		return &TLMessagesReadReactions{
			Constructor: -2099097129,
		}
	},
	276705696: func() TLObject { // 0x107e31a0
		return &TLMessagesSearchSentMedia{
			Constructor: 276705696,
		}
	},
	385663691: func() TLObject { // 0x16fcc2cb
		return &TLMessagesGetAttachMenuBots{
			Constructor: 385663691,
		}
	},
	1998676370: func() TLObject { // 0x77216192
		return &TLMessagesGetAttachMenuBot{
			Constructor: 1998676370,
		}
	},
	1777704297: func() TLObject { // 0x69f59d69
		return &TLMessagesToggleBotInAttachMenu{
			Constructor: 1777704297,
		}
	},
	451818415: func() TLObject { // 0x1aee33af
		return &TLMessagesToggleBotInAttachMenu{
			Constructor: 451818415,
		}
	},
	647873217: func() TLObject { // 0x269dc2c1
		return &TLMessagesRequestWebView{
			Constructor: 647873217,
		}
	},
	395003915: func() TLObject { // 0x178b480b
		return &TLMessagesRequestWebView{
			Constructor: 395003915,
		}
	},
	-58219204: func() TLObject { // 0xfc87a53c
		return &TLMessagesRequestWebView{
			Constructor: -58219204,
		}
	},
	-1850648527: func() TLObject { // 0x91b15831
		return &TLMessagesRequestWebView{
			Constructor: -1850648527,
		}
	},
	262163967: func() TLObject { // 0xfa04dff
		return &TLMessagesRequestWebView{
			Constructor: 262163967,
		}
	},
	-1328014717: func() TLObject { // 0xb0d81a83
		return &TLMessagesProlongWebView{
			Constructor: -1328014717,
		}
	},
	2146648841: func() TLObject { // 0x7ff34309
		return &TLMessagesProlongWebView{
			Constructor: 2146648841,
		}
	},
	-362824498: func() TLObject { // 0xea5fbcce
		return &TLMessagesProlongWebView{
			Constructor: -362824498,
		}
	},
	-768945848: func() TLObject { // 0xd22ad148
		return &TLMessagesProlongWebView{
			Constructor: -768945848,
		}
	},
	1094336115: func() TLObject { // 0x413a3e73
		return &TLMessagesRequestSimpleWebView413A3E73{
			Constructor: 1094336115,
		}
	},
	172168437: func() TLObject { // 0xa4314f5
		return &TLMessagesSendWebViewResultMessage{
			Constructor: 172168437,
		}
	},
	-603831608: func() TLObject { // 0xdc0242c8
		return &TLMessagesSendWebViewData{
			Constructor: -603831608,
		}
	},
	647928393: func() TLObject { // 0x269e9a49
		return &TLMessagesTranscribeAudio{
			Constructor: 647928393,
		}
	},
	2132608815: func() TLObject { // 0x7f1d072f
		return &TLMessagesRateTranscribedAudio{
			Constructor: 2132608815,
		}
	},
	-643100844: func() TLObject { // 0xd9ab0f54
		return &TLMessagesGetCustomEmojiDocuments{
			Constructor: -643100844,
		}
	},
	-67329649: func() TLObject { // 0xfbfca18f
		return &TLMessagesGetEmojiStickers{
			Constructor: -67329649,
		}
	},
	248473398: func() TLObject { // 0xecf6736
		return &TLMessagesGetFeaturedEmojiStickers{
			Constructor: 248473398,
		}
	},
	1063567478: func() TLObject { // 0x3f64c076
		return &TLMessagesReportReaction{
			Constructor: 1063567478,
		}
	},
	-1149164102: func() TLObject { // 0xbb8125ba
		return &TLMessagesGetTopReactions{
			Constructor: -1149164102,
		}
	},
	960896434: func() TLObject { // 0x39461db2
		return &TLMessagesGetRecentReactions{
			Constructor: 960896434,
		}
	},
	-1644236876: func() TLObject { // 0x9dfeefb4
		return &TLMessagesClearRecentReactions{
			Constructor: -1644236876,
		}
	},
	-2064119788: func() TLObject { // 0x84f80814
		return &TLMessagesGetExtendedMedia{
			Constructor: -2064119788,
		}
	},
	-1632299963: func() TLObject { // 0x9eb51445
		return &TLMessagesSetDefaultHistoryTTL{
			Constructor: -1632299963,
		}
	},
	1703637384: func() TLObject { // 0x658b7188
		return &TLMessagesGetDefaultHistoryTTL{
			Constructor: 1703637384,
		}
	},
	-1850552224: func() TLObject { // 0x91b2d060
		return &TLMessagesSendBotRequestedPeer{
			Constructor: -1850552224,
		}
	},
	-29831141: func() TLObject { // 0xfe38d01b
		return &TLMessagesSendBotRequestedPeer{
			Constructor: -29831141,
		}
	},
	1955122779: func() TLObject { // 0x7488ce5b
		return &TLMessagesGetEmojiGroups{
			Constructor: 1955122779,
		}
	},
	785209037: func() TLObject { // 0x2ecd56cd
		return &TLMessagesGetEmojiStatusGroups{
			Constructor: 785209037,
		}
	},
	564480243: func() TLObject { // 0x21a548f3
		return &TLMessagesGetEmojiProfilePhotoGroups{
			Constructor: 564480243,
		}
	},
	739360983: func() TLObject { // 0x2c11c0d7
		return &TLMessagesSearchCustomEmoji{
			Constructor: 739360983,
		}
	},
	-461589127: func() TLObject { // 0xe47cb579
		return &TLMessagesTogglePeerTranslations{
			Constructor: -461589127,
		}
	},
	889046467: func() TLObject { // 0x34fdc5c3
		return &TLMessagesGetBotApp{
			Constructor: 889046467,
		}
	},
	1398901710: func() TLObject { // 0x53618bce
		return &TLMessagesRequestAppWebView53618BCE{
			Constructor: 1398901710,
		}
	},
	-1879389471: func() TLObject { // 0x8ffacae1
		return &TLMessagesSetChatWallPaper{
			Constructor: -1879389471,
		}
	},
	-1833678516: func() TLObject { // 0x92b4494c
		return &TLMessagesSearchEmojiStickerSets{
			Constructor: -1833678516,
		}
	},
	1401016858: func() TLObject { // 0x5381d21a
		return &TLMessagesGetSavedDialogs{
			Constructor: 1401016858,
		}
	},
	1033519437: func() TLObject { // 0x3d9a414d
		return &TLMessagesGetSavedHistory{
			Constructor: 1033519437,
		}
	},
	1855459371: func() TLObject { // 0x6e98102b
		return &TLMessagesDeleteSavedHistory{
			Constructor: 1855459371,
		}
	},
	-700607264: func() TLObject { // 0xd63d94e0
		return &TLMessagesGetPinnedSavedDialogs{
			Constructor: -700607264,
		}
	},
	-1400783906: func() TLObject { // 0xac81bbde
		return &TLMessagesToggleSavedDialogPin{
			Constructor: -1400783906,
		}
	},
	-1955502713: func() TLObject { // 0x8b716587
		return &TLMessagesReorderPinnedSavedDialogs{
			Constructor: -1955502713,
		}
	},
	909631579: func() TLObject { // 0x3637e05b
		return &TLMessagesGetSavedReactionTags{
			Constructor: 909631579,
		}
	},
	1981668047: func() TLObject { // 0x761ddacf
		return &TLMessagesGetSavedReactionTags{
			Constructor: 1981668047,
		}
	},
	1613331948: func() TLObject { // 0x60297dec
		return &TLMessagesUpdateSavedReactionTag{
			Constructor: 1613331948,
		}
	},
	-1107741656: func() TLObject { // 0xbdf93428
		return &TLMessagesGetDefaultTagReactions{
			Constructor: -1107741656,
		}
	},
	-1941176739: func() TLObject { // 0x8c4bfe5d
		return &TLMessagesGetOutboxReadDate{
			Constructor: -1941176739,
		}
	},
	-729550168: func() TLObject { // 0xd483f2a8
		return &TLMessagesGetQuickReplies{
			Constructor: -729550168,
		}
	},
	1613961479: func() TLObject { // 0x60331907
		return &TLMessagesReorderQuickReplies{
			Constructor: 1613961479,
		}
	},
	-237962285: func() TLObject { // 0xf1d0fbd3
		return &TLMessagesCheckQuickReplyShortcut{
			Constructor: -237962285,
		}
	},
	1543519471: func() TLObject { // 0x5c003cef
		return &TLMessagesEditQuickReplyShortcut{
			Constructor: 1543519471,
		}
	},
	1019234112: func() TLObject { // 0x3cc04740
		return &TLMessagesDeleteQuickReplyShortcut{
			Constructor: 1019234112,
		}
	},
	-1801153085: func() TLObject { // 0x94a495c3
		return &TLMessagesGetQuickReplyMessages{
			Constructor: -1801153085,
		}
	},
	1819610593: func() TLObject { // 0x6c750de1
		return &TLMessagesSendQuickReplyMessages{
			Constructor: 1819610593,
		}
	},
	857029332: func() TLObject { // 0x33153ad4
		return &TLMessagesSendQuickReplyMessages{
			Constructor: 857029332,
		}
	},
	-519706352: func() TLObject { // 0xe105e910
		return &TLMessagesDeleteQuickReplyMessages{
			Constructor: -519706352,
		}
	},
	-47326647: func() TLObject { // 0xfd2dda49
		return &TLMessagesToggleDialogFilterTags{
			Constructor: -47326647,
		}
	},
	-793386500: func() TLObject { // 0xd0b5e1fc
		return &TLMessagesGetMyStickers{
			Constructor: -793386500,
		}
	},
	500711669: func() TLObject { // 0x1dd840f5
		return &TLMessagesGetEmojiStickerGroups{
			Constructor: 500711669,
		}
	},
	-559805895: func() TLObject { // 0xdea20a39
		return &TLMessagesGetAvailableEffects{
			Constructor: -559805895,
		}
	},
	92925557: func() TLObject { // 0x589ee75
		return &TLMessagesEditFactCheck{
			Constructor: 92925557,
		}
	},
	-774204404: func() TLObject { // 0xd1da940c
		return &TLMessagesDeleteFactCheck{
			Constructor: -774204404,
		}
	},
	-1177696786: func() TLObject { // 0xb9cdc5ee
		return &TLMessagesGetFactCheck{
			Constructor: -1177696786,
		}
	},
	-304838614: func() TLObject { // 0xedd4882a
		return &TLUpdatesGetState{
			Constructor: -304838614,
		}
	},
	432207715: func() TLObject { // 0x19c2f763
		return &TLUpdatesGetDifference{
			Constructor: 432207715,
		}
	},
	630429265: func() TLObject { // 0x25939651
		return &TLUpdatesGetDifference{
			Constructor: 630429265,
		}
	},
	51854712: func() TLObject { // 0x3173d78
		return &TLUpdatesGetChannelDifference{
			Constructor: 51854712,
		}
	},
	166207545: func() TLObject { // 0x9e82039
		return &TLPhotosUpdateProfilePhoto{
			Constructor: 166207545,
		}
	},
	473782614: func() TLObject { // 0x1c3d5956
		return &TLPhotosUpdateProfilePhoto{
			Constructor: 473782614,
		}
	},
	1926525996: func() TLObject { // 0x72d4742c
		return &TLPhotosUpdateProfilePhoto{
			Constructor: 1926525996,
		}
	},
	59286453: func() TLObject { // 0x388a3b5
		return &TLPhotosUploadProfilePhoto{
			Constructor: 59286453,
		}
	},
	154966609: func() TLObject { // 0x93c9a51
		return &TLPhotosUploadProfilePhoto{
			Constructor: 154966609,
		}
	},
	-1980559511: func() TLObject { // 0x89f30f69
		return &TLPhotosUploadProfilePhoto{
			Constructor: -1980559511,
		}
	},
	-2016444625: func() TLObject { // 0x87cf7f2f
		return &TLPhotosDeletePhotos{
			Constructor: -2016444625,
		}
	},
	-1848823128: func() TLObject { // 0x91cd32a8
		return &TLPhotosGetUserPhotos{
			Constructor: -1848823128,
		}
	},
	-515093903: func() TLObject { // 0xe14c4a71
		return &TLPhotosUploadContactProfilePhoto{
			Constructor: -515093903,
		}
	},
	-1189444673: func() TLObject { // 0xb91a83bf
		return &TLPhotosUploadContactProfilePhoto{
			Constructor: -1189444673,
		}
	},
	-1291540959: func() TLObject { // 0xb304a621
		return &TLUploadSaveFilePart{
			Constructor: -1291540959,
		}
	},
	-1101843010: func() TLObject { // 0xbe5335be
		return &TLUploadGetFile{
			Constructor: -1101843010,
		}
	},
	-1319462148: func() TLObject { // 0xb15a9afc
		return &TLUploadGetFile{
			Constructor: -1319462148,
		}
	},
	-562337987: func() TLObject { // 0xde7b673d
		return &TLUploadSaveBigFilePart{
			Constructor: -562337987,
		}
	},
	619086221: func() TLObject { // 0x24e6818d
		return &TLUploadGetWebFile{
			Constructor: 619086221,
		}
	},
	962554330: func() TLObject { // 0x395f69da
		return &TLUploadGetCdnFile{
			Constructor: 962554330,
		}
	},
	536919235: func() TLObject { // 0x2000bcc3
		return &TLUploadGetCdnFile{
			Constructor: 536919235,
		}
	},
	-1691921240: func() TLObject { // 0x9b2754a8
		return &TLUploadReuploadCdnFile{
			Constructor: -1691921240,
		}
	},
	-1847836879: func() TLObject { // 0x91dc3f31
		return &TLUploadGetCdnFileHashes{
			Constructor: -1847836879,
		}
	},
	1302676017: func() TLObject { // 0x4da54231
		return &TLUploadGetCdnFileHashes{
			Constructor: 1302676017,
		}
	},
	-1856595926: func() TLObject { // 0x9156982a
		return &TLUploadGetFileHashes{
			Constructor: -1856595926,
		}
	},
	-956147407: func() TLObject { // 0xc7025931
		return &TLUploadGetFileHashes{
			Constructor: -956147407,
		}
	},
	-990308245: func() TLObject { // 0xc4f9186b
		return &TLHelpGetConfig{
			Constructor: -990308245,
		}
	},
	531836966: func() TLObject { // 0x1fb33026
		return &TLHelpGetNearestDc{
			Constructor: 531836966,
		}
	},
	1378703997: func() TLObject { // 0x522d5a7d
		return &TLHelpGetAppUpdate{
			Constructor: 1378703997,
		}
	},
	1295590211: func() TLObject { // 0x4d392343
		return &TLHelpGetInviteText{
			Constructor: 1295590211,
		}
	},
	-1663104819: func() TLObject { // 0x9cdf08cd
		return &TLHelpGetSupport{
			Constructor: -1663104819,
		}
	},
	-333262899: func() TLObject { // 0xec22cfcd
		return &TLHelpSetBotUpdatesStatus{
			Constructor: -333262899,
		}
	},
	1375900482: func() TLObject { // 0x52029342
		return &TLHelpGetCdnConfig{
			Constructor: 1375900482,
		}
	},
	1036054804: func() TLObject { // 0x3dc0f114
		return &TLHelpGetRecentMeUrls{
			Constructor: 1036054804,
		}
	},
	749019089: func() TLObject { // 0x2ca51fd1
		return &TLHelpGetTermsOfServiceUpdate{
			Constructor: 749019089,
		}
	},
	-294455398: func() TLObject { // 0xee72f79a
		return &TLHelpAcceptTermsOfService{
			Constructor: -294455398,
		}
	},
	1072547679: func() TLObject { // 0x3fedc75f
		return &TLHelpGetDeepLinkInfo{
			Constructor: 1072547679,
		}
	},
	1642330196: func() TLObject { // 0x61e3f854
		return &TLHelpGetAppConfig61E3F854{
			Constructor: 1642330196,
		}
	},
	1862465352: func() TLObject { // 0x6f02f748
		return &TLHelpSaveAppLog{
			Constructor: 1862465352,
		}
	},
	-966677240: func() TLObject { // 0xc661ad08
		return &TLHelpGetPassportConfig{
			Constructor: -966677240,
		}
	},
	-748624084: func() TLObject { // 0xd360e72c
		return &TLHelpGetSupportName{
			Constructor: -748624084,
		}
	},
	59377875: func() TLObject { // 0x38a08d3
		return &TLHelpGetUserInfo{
			Constructor: 59377875,
		}
	},
	1723407216: func() TLObject { // 0x66b91b70
		return &TLHelpEditUserInfo{
			Constructor: 1723407216,
		}
	},
	-1063816159: func() TLObject { // 0xc0977421
		return &TLHelpGetPromoData{
			Constructor: -1063816159,
		}
	},
	505748629: func() TLObject { // 0x1e251c95
		return &TLHelpHidePromoData{
			Constructor: 505748629,
		}
	},
	-183649631: func() TLObject { // 0xf50dbaa1
		return &TLHelpDismissSuggestion{
			Constructor: -183649631,
		}
	},
	1935116200: func() TLObject { // 0x735787a8
		return &TLHelpGetCountriesList{
			Constructor: 1935116200,
		}
	},
	-1206152236: func() TLObject { // 0xb81b93d4
		return &TLHelpGetPremiumPromo{
			Constructor: -1206152236,
		}
	},
	-629083089: func() TLObject { // 0xda80f42f
		return &TLHelpGetPeerColors{
			Constructor: -629083089,
		}
	},
	-1412453891: func() TLObject { // 0xabcfa9fd
		return &TLHelpGetPeerProfileColors{
			Constructor: -1412453891,
		}
	},
	1236468288: func() TLObject { // 0x49b30240
		return &TLHelpGetTimezonesList{
			Constructor: 1236468288,
		}
	},
	-871347913: func() TLObject { // 0xcc104937
		return &TLChannelsReadHistory{
			Constructor: -871347913,
		}
	},
	-2067661490: func() TLObject { // 0x84c1fd4e
		return &TLChannelsDeleteMessages{
			Constructor: -2067661490,
		}
	},
	-196443371: func() TLObject { // 0xf44a8315
		return &TLChannelsReportSpam{
			Constructor: -196443371,
		}
	},
	-1383294429: func() TLObject { // 0xad8c9a23
		return &TLChannelsGetMessages{
			Constructor: -1383294429,
		}
	},
	-1814580409: func() TLObject { // 0x93d7b347
		return &TLChannelsGetMessages{
			Constructor: -1814580409,
		}
	},
	2010044880: func() TLObject { // 0x77ced9d0
		return &TLChannelsGetParticipants{
			Constructor: 2010044880,
		}
	},
	-1599378234: func() TLObject { // 0xa0ab6cc6
		return &TLChannelsGetParticipant{
			Constructor: -1599378234,
		}
	},
	176122811: func() TLObject { // 0xa7f6bbb
		return &TLChannelsGetChannels{
			Constructor: 176122811,
		}
	},
	141781513: func() TLObject { // 0x8736a09
		return &TLChannelsGetFullChannel{
			Constructor: 141781513,
		}
	},
	-1862244601: func() TLObject { // 0x91006707
		return &TLChannelsCreateChannel{
			Constructor: -1862244601,
		}
	},
	1029681423: func() TLObject { // 0x3d5fb10f
		return &TLChannelsCreateChannel{
			Constructor: 1029681423,
		}
	},
	-751007486: func() TLObject { // 0xd33c8902
		return &TLChannelsEditAdmin{
			Constructor: -751007486,
		}
	},
	1450044624: func() TLObject { // 0x566decd0
		return &TLChannelsEditTitle{
			Constructor: 1450044624,
		}
	},
	-248621111: func() TLObject { // 0xf12e57c9
		return &TLChannelsEditPhoto{
			Constructor: -248621111,
		}
	},
	283557164: func() TLObject { // 0x10e6bd2c
		return &TLChannelsCheckUsername{
			Constructor: 283557164,
		}
	},
	890549214: func() TLObject { // 0x3514b3de
		return &TLChannelsUpdateUsername{
			Constructor: 890549214,
		}
	},
	615851205: func() TLObject { // 0x24b524c5
		return &TLChannelsJoinChannel{
			Constructor: 615851205,
		}
	},
	-130635115: func() TLObject { // 0xf836aa95
		return &TLChannelsLeaveChannel{
			Constructor: -130635115,
		}
	},
	-907854508: func() TLObject { // 0xc9e33d54
		return &TLChannelsInviteToChannelC9E33D54{
			Constructor: -907854508,
		}
	},
	-1072619549: func() TLObject { // 0xc0111fe3
		return &TLChannelsDeleteChannel{
			Constructor: -1072619549,
		}
	},
	-432034325: func() TLObject { // 0xe63fadeb
		return &TLChannelsExportMessageLink{
			Constructor: -432034325,
		}
	},
	527021574: func() TLObject { // 0x1f69b606
		return &TLChannelsToggleSignatures{
			Constructor: 527021574,
		}
	},
	-122669393: func() TLObject { // 0xf8b036af
		return &TLChannelsGetAdminedPublicChannels{
			Constructor: -122669393,
		}
	},
	-1763259007: func() TLObject { // 0x96e6cd81
		return &TLChannelsEditBanned{
			Constructor: -1763259007,
		}
	},
	870184064: func() TLObject { // 0x33ddf480
		return &TLChannelsGetAdminLog{
			Constructor: 870184064,
		}
	},
	-359881479: func() TLObject { // 0xea8ca4f9
		return &TLChannelsSetStickers{
			Constructor: -359881479,
		}
	},
	-357180360: func() TLObject { // 0xeab5dc38
		return &TLChannelsReadMessageContents{
			Constructor: -357180360,
		}
	},
	-1683319225: func() TLObject { // 0x9baa9647
		return &TLChannelsDeleteHistory9BAA9647{
			Constructor: -1683319225,
		}
	},
	-356796084: func() TLObject { // 0xeabbb94c
		return &TLChannelsTogglePreHistoryHidden{
			Constructor: -356796084,
		}
	},
	-2092831552: func() TLObject { // 0x8341ecc0
		return &TLChannelsGetLeftChannels{
			Constructor: -2092831552,
		}
	},
	-170208392: func() TLObject { // 0xf5dad378
		return &TLChannelsGetGroupsForDiscussion{
			Constructor: -170208392,
		}
	},
	1079520178: func() TLObject { // 0x40582bb2
		return &TLChannelsSetDiscussionGroup{
			Constructor: 1079520178,
		}
	},
	-1892102881: func() TLObject { // 0x8f38cd1f
		return &TLChannelsEditCreator{
			Constructor: -1892102881,
		}
	},
	1491484525: func() TLObject { // 0x58e63f6d
		return &TLChannelsEditLocation{
			Constructor: 1491484525,
		}
	},
	-304832784: func() TLObject { // 0xedd49ef0
		return &TLChannelsToggleSlowMode{
			Constructor: -304832784,
		}
	},
	300429806: func() TLObject { // 0x11e831ee
		return &TLChannelsGetInactiveChannels{
			Constructor: 300429806,
		}
	},
	187239529: func() TLObject { // 0xb290c69
		return &TLChannelsConvertToGigagroup{
			Constructor: 187239529,
		}
	},
	-1095836780: func() TLObject { // 0xbeaedb94
		return &TLChannelsViewSponsoredMessage{
			Constructor: -1095836780,
		}
	},
	-333377601: func() TLObject { // 0xec210fbf
		return &TLChannelsGetSponsoredMessages{
			Constructor: -333377601,
		}
	},
	231174382: func() TLObject { // 0xdc770ee
		return &TLChannelsGetSendAs{
			Constructor: 231174382,
		}
	},
	913655003: func() TLObject { // 0x367544db
		return &TLChannelsDeleteParticipantHistory{
			Constructor: 913655003,
		}
	},
	-456419968: func() TLObject { // 0xe4cb9580
		return &TLChannelsToggleJoinToSend{
			Constructor: -456419968,
		}
	},
	1277789622: func() TLObject { // 0x4c2985b6
		return &TLChannelsToggleJoinRequest{
			Constructor: 1277789622,
		}
	},
	-1268978403: func() TLObject { // 0xb45ced1d
		return &TLChannelsReorderUsernames{
			Constructor: -1268978403,
		}
	},
	1358053637: func() TLObject { // 0x50f24105
		return &TLChannelsToggleUsername{
			Constructor: 1358053637,
		}
	},
	170155475: func() TLObject { // 0xa245dd3
		return &TLChannelsDeactivateAllUsernames{
			Constructor: 170155475,
		}
	},
	-1540781271: func() TLObject { // 0xa4298b29
		return &TLChannelsToggleForum{
			Constructor: -1540781271,
		}
	},
	-200539612: func() TLObject { // 0xf40c0224
		return &TLChannelsCreateForumTopic{
			Constructor: -200539612,
		}
	},
	233136337: func() TLObject { // 0xde560d1
		return &TLChannelsGetForumTopics{
			Constructor: 233136337,
		}
	},
	-1333584199: func() TLObject { // 0xb0831eb9
		return &TLChannelsGetForumTopicsByID{
			Constructor: -1333584199,
		}
	},
	-186670715: func() TLObject { // 0xf4dfa185
		return &TLChannelsEditForumTopic{
			Constructor: -186670715,
		}
	},
	1820868141: func() TLObject { // 0x6c883e2d
		return &TLChannelsEditForumTopic{
			Constructor: 1820868141,
		}
	},
	1814925350: func() TLObject { // 0x6c2d9026
		return &TLChannelsUpdatePinnedForumTopic{
			Constructor: 1814925350,
		}
	},
	876830509: func() TLObject { // 0x34435f2d
		return &TLChannelsDeleteTopicHistory{
			Constructor: 876830509,
		}
	},
	693150095: func() TLObject { // 0x2950a18f
		return &TLChannelsReorderPinnedForumTopics{
			Constructor: 693150095,
		}
	},
	1760814315: func() TLObject { // 0x68f3e4eb
		return &TLChannelsToggleAntiSpam{
			Constructor: 1760814315,
		}
	},
	-1471109485: func() TLObject { // 0xa850a693
		return &TLChannelsReportAntiSpamFalsePositive{
			Constructor: -1471109485,
		}
	},
	1785624660: func() TLObject { // 0x6a6e7854
		return &TLChannelsToggleParticipantsHidden{
			Constructor: 1785624660,
		}
	},
	414170259: func() TLObject { // 0x18afbc93
		return &TLChannelsClickSponsoredMessage{
			Constructor: 414170259,
		}
	},
	-659933583: func() TLObject { // 0xd8aa3671
		return &TLChannelsUpdateColor{
			Constructor: -659933583,
		}
	},
	1645879327: func() TLObject { // 0x621a201f
		return &TLChannelsUpdateColor{
			Constructor: 1645879327,
		}
	},
	-1757889771: func() TLObject { // 0x9738bb15
		return &TLChannelsToggleViewForumAsMessages{
			Constructor: -1757889771,
		}
	},
	631707458: func() TLObject { // 0x25a71742
		return &TLChannelsGetChannelRecommendations{
			Constructor: 631707458,
		}
	},
	-2085155433: func() TLObject { // 0x83b70d97
		return &TLChannelsGetChannelRecommendations{
			Constructor: -2085155433,
		}
	},
	-254548312: func() TLObject { // 0xf0d3e6a8
		return &TLChannelsUpdateEmojiStatus{
			Constructor: -254548312,
		}
	},
	-1388733202: func() TLObject { // 0xad399cee
		return &TLChannelsSetBoostsToUnblockRestrictions{
			Constructor: -1388733202,
		}
	},
	1020866743: func() TLObject { // 0x3cd930b7
		return &TLChannelsSetEmojiStickers{
			Constructor: 1020866743,
		}
	},
	-1349519687: func() TLObject { // 0xaf8ff6b9
		return &TLChannelsReportSponsoredMessage{
			Constructor: -1349519687,
		}
	},
	-1696000743: func() TLObject { // 0x9ae91519
		return &TLChannelsRestrictSponsoredMessages{
			Constructor: -1696000743,
		}
	},
	-778069893: func() TLObject { // 0xd19f987b
		return &TLChannelsSearchPosts{
			Constructor: -778069893,
		}
	},
	-1440257555: func() TLObject { // 0xaa2769ed
		return &TLBotsSendCustomRequest{
			Constructor: -1440257555,
		}
	},
	-434028723: func() TLObject { // 0xe6213f4d
		return &TLBotsAnswerWebhookJSONQuery{
			Constructor: -434028723,
		}
	},
	85399130: func() TLObject { // 0x517165a
		return &TLBotsSetBotCommands{
			Constructor: 85399130,
		}
	},
	1032708345: func() TLObject { // 0x3d8de0f9
		return &TLBotsResetBotCommands{
			Constructor: 1032708345,
		}
	},
	-481554986: func() TLObject { // 0xe34c0dd6
		return &TLBotsGetBotCommands{
			Constructor: -481554986,
		}
	},
	1157944655: func() TLObject { // 0x4504d54f
		return &TLBotsSetBotMenuButton{
			Constructor: 1157944655,
		}
	},
	-1671369944: func() TLObject { // 0x9c60eb28
		return &TLBotsGetBotMenuButton{
			Constructor: -1671369944,
		}
	},
	2021942497: func() TLObject { // 0x788464e1
		return &TLBotsSetBotBroadcastDefaultAdminRights{
			Constructor: 2021942497,
		}
	},
	-1839281686: func() TLObject { // 0x925ec9ea
		return &TLBotsSetBotGroupDefaultAdminRights{
			Constructor: -1839281686,
		}
	},
	282013987: func() TLObject { // 0x10cf3123
		return &TLBotsSetBotInfo{
			Constructor: 282013987,
		}
	},
	-1553604742: func() TLObject { // 0xa365df7a
		return &TLBotsSetBotInfo{
			Constructor: -1553604742,
		}
	},
	-589753091: func() TLObject { // 0xdcd914fd
		return &TLBotsGetBotInfoDCD914FD{
			Constructor: -589753091,
		}
	},
	-1760972350: func() TLObject { // 0x9709b1c2
		return &TLBotsReorderUsernames{
			Constructor: -1760972350,
		}
	},
	87861619: func() TLObject { // 0x53ca973
		return &TLBotsToggleUsername{
			Constructor: 87861619,
		}
	},
	324662502: func() TLObject { // 0x1359f4e6
		return &TLBotsCanSendMessage{
			Constructor: 324662502,
		}
	},
	-248323089: func() TLObject { // 0xf132e3ef
		return &TLBotsAllowSendMessage{
			Constructor: -248323089,
		}
	},
	142591463: func() TLObject { // 0x87fc5e7
		return &TLBotsInvokeWebViewCustomMethod{
			Constructor: 142591463,
		}
	},
	924093883: func() TLObject { // 0x37148dbb
		return &TLPaymentsGetPaymentForm{
			Constructor: 924093883,
		}
	},
	-1976353651: func() TLObject { // 0x8a333c8d
		return &TLPaymentsGetPaymentForm{
			Constructor: -1976353651,
		}
	},
	611897804: func() TLObject { // 0x2478d1cc
		return &TLPaymentsGetPaymentReceipt{
			Constructor: 611897804,
		}
	},
	-1228345045: func() TLObject { // 0xb6c8f12b
		return &TLPaymentsValidateRequestedInfo{
			Constructor: -1228345045,
		}
	},
	-619695760: func() TLObject { // 0xdb103170
		return &TLPaymentsValidateRequestedInfo{
			Constructor: -619695760,
		}
	},
	755192367: func() TLObject { // 0x2d03522f
		return &TLPaymentsSendPaymentForm{
			Constructor: 755192367,
		}
	},
	818134173: func() TLObject { // 0x30c3bc9d
		return &TLPaymentsSendPaymentForm{
			Constructor: 818134173,
		}
	},
	578650699: func() TLObject { // 0x227d824b
		return &TLPaymentsGetSavedInfo{
			Constructor: 578650699,
		}
	},
	-667062079: func() TLObject { // 0xd83d70c1
		return &TLPaymentsClearSavedInfo{
			Constructor: -667062079,
		}
	},
	779736953: func() TLObject { // 0x2e79d779
		return &TLPaymentsGetBankCardData{
			Constructor: 779736953,
		}
	},
	261206117: func() TLObject { // 0xf91b065
		return &TLPaymentsExportInvoice{
			Constructor: 261206117,
		}
	},
	-92600067: func() TLObject { // 0xfa7b08fd
		return &TLPaymentsExportInvoice{
			Constructor: -92600067,
		}
	},
	-2131921795: func() TLObject { // 0x80ed747d
		return &TLPaymentsAssignAppStoreTransaction{
			Constructor: -2131921795,
		}
	},
	267129798: func() TLObject { // 0xfec13c6
		return &TLPaymentsAssignAppStoreTransaction{
			Constructor: 267129798,
		}
	},
	-537046829: func() TLObject { // 0xdffd50d3
		return &TLPaymentsAssignPlayMarketTransaction{
			Constructor: -537046829,
		}
	},
	1336560365: func() TLObject { // 0x4faa4aed
		return &TLPaymentsAssignPlayMarketTransaction{
			Constructor: 1336560365,
		}
	},
	-1614700874: func() TLObject { // 0x9fc19eb6
		return &TLPaymentsCanPurchasePremium{
			Constructor: -1614700874,
		}
	},
	-1435856696: func() TLObject { // 0xaa6a90c8
		return &TLPaymentsCanPurchasePremium{
			Constructor: -1435856696,
		}
	},
	660060756: func() TLObject { // 0x2757ba54
		return &TLPaymentsGetPremiumGiftCodeOptions{
			Constructor: 660060756,
		}
	},
	-1907247935: func() TLObject { // 0x8e51b4c1
		return &TLPaymentsCheckGiftCode{
			Constructor: -1907247935,
		}
	},
	-152934316: func() TLObject { // 0xf6e26854
		return &TLPaymentsApplyGiftCode{
			Constructor: -152934316,
		}
	},
	-198994907: func() TLObject { // 0xf4239425
		return &TLPaymentsGetGiveawayInfo{
			Constructor: -198994907,
		}
	},
	1609928480: func() TLObject { // 0x5ff58f20
		return &TLPaymentsLaunchPrepaidGiveaway{
			Constructor: 1609928480,
		}
	},
	-1072773165: func() TLObject { // 0xc00ec7d3
		return &TLPaymentsGetStarsTopupOptions{
			Constructor: -1072773165,
		}
	},
	273665959: func() TLObject { // 0x104fcfa7
		return &TLPaymentsGetStarsStatus{
			Constructor: 273665959,
		}
	},
	-1751937702: func() TLObject { // 0x97938d5a
		return &TLPaymentsGetStarsTransactions{
			Constructor: -1751937702,
		}
	},
	1731904249: func() TLObject { // 0x673ac2f9
		return &TLPaymentsGetStarsTransactions{
			Constructor: 1731904249,
		}
	},
	45839133: func() TLObject { // 0x2bb731d
		return &TLPaymentsSendStarsForm{
			Constructor: 45839133,
		}
	},
	632196938: func() TLObject { // 0x25ae8f4a
		return &TLPaymentsRefundStarsCharge{
			Constructor: 632196938,
		}
	},
	-652215594: func() TLObject { // 0xd91ffad6
		return &TLPaymentsGetStarsRevenueStats{
			Constructor: -652215594,
		}
	},
	331081907: func() TLObject { // 0x13bbe8b3
		return &TLPaymentsGetStarsRevenueWithdrawalUrl{
			Constructor: 331081907,
		}
	},
	-774377531: func() TLObject { // 0xd1d7efc5
		return &TLPaymentsGetStarsRevenueAdsAccountUrl{
			Constructor: -774377531,
		}
	},
	662973742: func() TLObject { // 0x27842d2e
		return &TLPaymentsGetStarsTransactionsByID{
			Constructor: 662973742,
		}
	},
	-1876841625: func() TLObject { // 0x9021ab67
		return &TLStickersCreateStickerSet{
			Constructor: -1876841625,
		}
	},
	-143257775: func() TLObject { // 0xf7760f51
		return &TLStickersRemoveStickerFromSet{
			Constructor: -143257775,
		}
	},
	-4795190: func() TLObject { // 0xffb6d4ca
		return &TLStickersChangeStickerPosition{
			Constructor: -4795190,
		}
	},
	-2041315650: func() TLObject { // 0x8653febe
		return &TLStickersAddStickerToSet{
			Constructor: -2041315650,
		}
	},
	-1486204014: func() TLObject { // 0xa76a5392
		return &TLStickersSetStickerSetThumb{
			Constructor: -1486204014,
		}
	},
	-1707717072: func() TLObject { // 0x9a364e30
		return &TLStickersSetStickerSetThumb{
			Constructor: -1707717072,
		}
	},
	676017721: func() TLObject { // 0x284b3639
		return &TLStickersCheckShortName{
			Constructor: 676017721,
		}
	},
	1303364867: func() TLObject { // 0x4dafc503
		return &TLStickersSuggestShortName{
			Constructor: 1303364867,
		}
	},
	-179077444: func() TLObject { // 0xf5537ebc
		return &TLStickersChangeSticker{
			Constructor: -179077444,
		}
	},
	306912256: func() TLObject { // 0x124b1c00
		return &TLStickersRenameStickerSet{
			Constructor: 306912256,
		}
	},
	-2022685804: func() TLObject { // 0x87704394
		return &TLStickersDeleteStickerSet{
			Constructor: -2022685804,
		}
	},
	1184253338: func() TLObject { // 0x4696459a
		return &TLStickersReplaceSticker{
			Constructor: 1184253338,
		}
	},
	1430593449: func() TLObject { // 0x55451fa9
		return &TLPhoneGetCallConfig{
			Constructor: 1430593449,
		}
	},
	1124046573: func() TLObject { // 0x42ff96ed
		return &TLPhoneRequestCall{
			Constructor: 1124046573,
		}
	},
	1003664544: func() TLObject { // 0x3bd2b4a0
		return &TLPhoneAcceptCall{
			Constructor: 1003664544,
		}
	},
	788404002: func() TLObject { // 0x2efe1722
		return &TLPhoneConfirmCall{
			Constructor: 788404002,
		}
	},
	399855457: func() TLObject { // 0x17d54f61
		return &TLPhoneReceivedCall{
			Constructor: 399855457,
		}
	},
	-1295269440: func() TLObject { // 0xb2cbc1c0
		return &TLPhoneDiscardCall{
			Constructor: -1295269440,
		}
	},
	1508562471: func() TLObject { // 0x59ead627
		return &TLPhoneSetCallRating{
			Constructor: 1508562471,
		}
	},
	662363518: func() TLObject { // 0x277add7e
		return &TLPhoneSaveCallDebug{
			Constructor: 662363518,
		}
	},
	-8744061: func() TLObject { // 0xff7a9383
		return &TLPhoneSendSignalingData{
			Constructor: -8744061,
		}
	},
	1221445336: func() TLObject { // 0x48cdc6d8
		return &TLPhoneCreateGroupCall{
			Constructor: 1221445336,
		}
	},
	-1322057861: func() TLObject { // 0xb132ff7b
		return &TLPhoneJoinGroupCall{
			Constructor: -1322057861,
		}
	},
	1342404601: func() TLObject { // 0x500377f9
		return &TLPhoneLeaveGroupCall{
			Constructor: 1342404601,
		}
	},
	2067345760: func() TLObject { // 0x7b393160
		return &TLPhoneInviteToGroupCall{
			Constructor: 2067345760,
		}
	},
	2054648117: func() TLObject { // 0x7a777135
		return &TLPhoneDiscardGroupCall{
			Constructor: 2054648117,
		}
	},
	1958458429: func() TLObject { // 0x74bbb43d
		return &TLPhoneToggleGroupCallSettings{
			Constructor: 1958458429,
		}
	},
	68699611: func() TLObject { // 0x41845db
		return &TLPhoneGetGroupCall{
			Constructor: 68699611,
		}
	},
	-984033109: func() TLObject { // 0xc558d8ab
		return &TLPhoneGetGroupParticipants{
			Constructor: -984033109,
		}
	},
	-1248003721: func() TLObject { // 0xb59cf977
		return &TLPhoneCheckGroupCall{
			Constructor: -1248003721,
		}
	},
	-248985848: func() TLObject { // 0xf128c708
		return &TLPhoneToggleGroupCallRecord{
			Constructor: -248985848,
		}
	},
	-1524155713: func() TLObject { // 0xa5273abf
		return &TLPhoneEditGroupCallParticipant{
			Constructor: -1524155713,
		}
	},
	480685066: func() TLObject { // 0x1ca6ac0a
		return &TLPhoneEditGroupCallTitle{
			Constructor: 480685066,
		}
	},
	-277077702: func() TLObject { // 0xef7c213a
		return &TLPhoneGetGroupCallJoinAs{
			Constructor: -277077702,
		}
	},
	-425040769: func() TLObject { // 0xe6aa647f
		return &TLPhoneExportGroupCallInvite{
			Constructor: -425040769,
		}
	},
	563885286: func() TLObject { // 0x219c34e6
		return &TLPhoneToggleGroupCallStartSubscription{
			Constructor: 563885286,
		}
	},
	1451287362: func() TLObject { // 0x5680e342
		return &TLPhoneStartScheduledGroupCall{
			Constructor: 1451287362,
		}
	},
	1465786252: func() TLObject { // 0x575e1f8c
		return &TLPhoneSaveDefaultGroupCallJoinAs{
			Constructor: 1465786252,
		}
	},
	-873829436: func() TLObject { // 0xcbea6bc4
		return &TLPhoneJoinGroupCallPresentation{
			Constructor: -873829436,
		}
	},
	475058500: func() TLObject { // 0x1c50d144
		return &TLPhoneLeaveGroupCallPresentation{
			Constructor: 475058500,
		}
	},
	447879488: func() TLObject { // 0x1ab21940
		return &TLPhoneGetGroupCallStreamChannels{
			Constructor: 447879488,
		}
	},
	-558650433: func() TLObject { // 0xdeb3abbf
		return &TLPhoneGetGroupCallStreamRtmpUrl{
			Constructor: -558650433,
		}
	},
	1092913030: func() TLObject { // 0x41248786
		return &TLPhoneSaveCallLog{
			Constructor: 1092913030,
		}
	},
	-219008246: func() TLObject { // 0xf2f2330a
		return &TLLangpackGetLangPack{
			Constructor: -219008246,
		}
	},
	-1699363442: func() TLObject { // 0x9ab5c58e
		return &TLLangpackGetLangPack{
			Constructor: -1699363442,
		}
	},
	-269862909: func() TLObject { // 0xefea3803
		return &TLLangpackGetStrings{
			Constructor: -269862909,
		}
	},
	773776152: func() TLObject { // 0x2e1ee318
		return &TLLangpackGetStrings{
			Constructor: 773776152,
		}
	},
	-845657435: func() TLObject { // 0xcd984aa5
		return &TLLangpackGetDifference{
			Constructor: -845657435,
		}
	},
	1120311183: func() TLObject { // 0x42c6978f
		return &TLLangpackGetLanguages{
			Constructor: 1120311183,
		}
	},
	-2146445955: func() TLObject { // 0x800fd57d
		return &TLLangpackGetLanguages{
			Constructor: -2146445955,
		}
	},
	1784243458: func() TLObject { // 0x6a596502
		return &TLLangpackGetLanguage{
			Constructor: 1784243458,
		}
	},
	1749536939: func() TLObject { // 0x6847d0ab
		return &TLFoldersEditPeerFolders{
			Constructor: 1749536939,
		}
	},
	-1421720550: func() TLObject { // 0xab42441a
		return &TLStatsGetBroadcastStats{
			Constructor: -1421720550,
		}
	},
	1646092192: func() TLObject { // 0x621d5fa0
		return &TLStatsLoadAsyncGraph{
			Constructor: 1646092192,
		}
	},
	-589330937: func() TLObject { // 0xdcdf8607
		return &TLStatsGetMegagroupStats{
			Constructor: -589330937,
		}
	},
	1595212100: func() TLObject { // 0x5f150144
		return &TLStatsGetMessagePublicForwards5F150144{
			Constructor: 1595212100,
		}
	},
	-1226791947: func() TLObject { // 0xb6e0a3f5
		return &TLStatsGetMessageStats{
			Constructor: -1226791947,
		}
	},
	927985472: func() TLObject { // 0x374fef40
		return &TLStatsGetStoryStats{
			Constructor: 927985472,
		}
	},
	-1505526026: func() TLObject { // 0xa6437ef6
		return &TLStatsGetStoryPublicForwards{
			Constructor: -1505526026,
		}
	},
	1977595505: func() TLObject { // 0x75dfb671
		return &TLStatsGetBroadcastRevenueStats{
			Constructor: 1977595505,
		}
	},
	711323507: func() TLObject { // 0x2a65ef73
		return &TLStatsGetBroadcastRevenueWithdrawalUrl{
			Constructor: 711323507,
		}
	},
	6891535: func() TLObject { // 0x69280f
		return &TLStatsGetBroadcastRevenueTransactions{
			Constructor: 6891535,
		}
	},
	-2072885362: func() TLObject { // 0x8472478e
		return &TLChatlistsExportChatlistInvite{
			Constructor: -2072885362,
		}
	},
	1906072670: func() TLObject { // 0x719c5c5e
		return &TLChatlistsDeleteExportedInvite{
			Constructor: 1906072670,
		}
	},
	1698543165: func() TLObject { // 0x653db63d
		return &TLChatlistsEditExportedInvite{
			Constructor: 1698543165,
		}
	},
	-838608253: func() TLObject { // 0xce03da83
		return &TLChatlistsGetExportedInvites{
			Constructor: -838608253,
		}
	},
	1103171583: func() TLObject { // 0x41c10fff
		return &TLChatlistsCheckChatlistInvite{
			Constructor: 1103171583,
		}
	},
	-1498291302: func() TLObject { // 0xa6b1e39a
		return &TLChatlistsJoinChatlistInvite{
			Constructor: -1498291302,
		}
	},
	-1992190687: func() TLObject { // 0x89419521
		return &TLChatlistsGetChatlistUpdates{
			Constructor: -1992190687,
		}
	},
	-527828747: func() TLObject { // 0xe089f8f5
		return &TLChatlistsJoinChatlistUpdates{
			Constructor: -527828747,
		}
	},
	1726252795: func() TLObject { // 0x66e486fb
		return &TLChatlistsHideChatlistUpdates{
			Constructor: 1726252795,
		}
	},
	-37955820: func() TLObject { // 0xfdbcd714
		return &TLChatlistsGetLeaveChatlistSuggestions{
			Constructor: -37955820,
		}
	},
	1962598714: func() TLObject { // 0x74fae13a
		return &TLChatlistsLeaveChatlist{
			Constructor: 1962598714,
		}
	},
	-941629475: func() TLObject { // 0xc7dfdfdd
		return &TLStoriesCanSendStory{
			Constructor: -941629475,
		}
	},
	-1325345699: func() TLObject { // 0xb100d45d
		return &TLStoriesCanSendStory{
			Constructor: -1325345699,
		}
	},
	-454661813: func() TLObject { // 0xe4e6694b
		return &TLStoriesSendStory{
			Constructor: -454661813,
		}
	},
	-1128843708: func() TLObject { // 0xbcb73644
		return &TLStoriesSendStory{
			Constructor: -1128843708,
		}
	},
	-732562196: func() TLObject { // 0xd455fcec
		return &TLStoriesSendStory{
			Constructor: -732562196,
		}
	},
	1112331386: func() TLObject { // 0x424cd47a
		return &TLStoriesSendStory{
			Constructor: 1112331386,
		}
	},
	-1249658298: func() TLObject { // 0xb583ba46
		return &TLStoriesEditStory{
			Constructor: -1249658298,
		}
	},
	-1447486748: func() TLObject { // 0xa9b91ae4
		return &TLStoriesEditStory{
			Constructor: -1447486748,
		}
	},
	716077633: func() TLObject { // 0x2aae7a41
		return &TLStoriesEditStory{
			Constructor: 716077633,
		}
	},
	-1369842849: func() TLObject { // 0xae59db5f
		return &TLStoriesDeleteStories{
			Constructor: -1369842849,
		}
	},
	-1244331561: func() TLObject { // 0xb5d501d7
		return &TLStoriesDeleteStories{
			Constructor: -1244331561,
		}
	},
	-1703566865: func() TLObject { // 0x9a75a1ef
		return &TLStoriesTogglePinned{
			Constructor: -1703566865,
		}
	},
	1365256516: func() TLObject { // 0x51602944
		return &TLStoriesTogglePinned{
			Constructor: 1365256516,
		}
	},
	-290400731: func() TLObject { // 0xeeb0d625
		return &TLStoriesGetAllStories{
			Constructor: -290400731,
		}
	},
	1478600156: func() TLObject { // 0x5821a5dc
		return &TLStoriesGetPinnedStories{
			Constructor: 1478600156,
		}
	},
	189206839: func() TLObject { // 0xb471137
		return &TLStoriesGetPinnedStories{
			Constructor: 189206839,
		}
	},
	-1271586794: func() TLObject { // 0xb4352016
		return &TLStoriesGetStoriesArchive{
			Constructor: -1271586794,
		}
	},
	526108114: func() TLObject { // 0x1f5bc5d2
		return &TLStoriesGetStoriesArchive{
			Constructor: 526108114,
		}
	},
	1467271796: func() TLObject { // 0x5774ca74
		return &TLStoriesGetStoriesByID{
			Constructor: 1467271796,
		}
	},
	1779814214: func() TLObject { // 0x6a15cf46
		return &TLStoriesGetStoriesByID{
			Constructor: 1779814214,
		}
	},
	2082822084: func() TLObject { // 0x7c2557c4
		return &TLStoriesToggleAllStoriesHidden{
			Constructor: 2082822084,
		}
	},
	-1521034552: func() TLObject { // 0xa556dac8
		return &TLStoriesReadStories{
			Constructor: -1521034552,
		}
	},
	-305852325: func() TLObject { // 0xedc5105b
		return &TLStoriesReadStories{
			Constructor: -305852325,
		}
	},
	-1308456197: func() TLObject { // 0xb2028afb
		return &TLStoriesIncrementStoryViews{
			Constructor: -1308456197,
		}
	},
	571629863: func() TLObject { // 0x22126127
		return &TLStoriesIncrementStoryViews{
			Constructor: 571629863,
		}
	},
	2127707223: func() TLObject { // 0x7ed23c57
		return &TLStoriesGetStoryViewsList{
			Constructor: 2127707223,
		}
	},
	-111189596: func() TLObject { // 0xf95f61a4
		return &TLStoriesGetStoryViewsList{
			Constructor: -111189596,
		}
	},
	1262182039: func() TLObject { // 0x4b3b5e97
		return &TLStoriesGetStoryViewsList{
			Constructor: 1262182039,
		}
	},
	685862088: func() TLObject { // 0x28e16cc8
		return &TLStoriesGetStoriesViews{
			Constructor: 685862088,
		}
	},
	-1703553370: func() TLObject { // 0x9a75d6a6
		return &TLStoriesGetStoriesViews{
			Constructor: -1703553370,
		}
	},
	2072899360: func() TLObject { // 0x7b8def20
		return &TLStoriesExportStoryLink{
			Constructor: 2072899360,
		}
	},
	384058318: func() TLObject { // 0x16e443ce
		return &TLStoriesExportStoryLink{
			Constructor: 384058318,
		}
	},
	421788300: func() TLObject { // 0x1923fa8c
		return &TLStoriesReport{
			Constructor: 421788300,
		}
	},
	-916725654: func() TLObject { // 0xc95be06a
		return &TLStoriesReport{
			Constructor: -916725654,
		}
	},
	1471926630: func() TLObject { // 0x57bbd166
		return &TLStoriesActivateStealthMode{
			Constructor: 1471926630,
		}
	},
	2144810674: func() TLObject { // 0x7fd736b2
		return &TLStoriesSendReaction{
			Constructor: 2144810674,
		}
	},
	1235921331: func() TLObject { // 0x49aaa9b3
		return &TLStoriesSendReaction{
			Constructor: 1235921331,
		}
	},
	743103056: func() TLObject { // 0x2c4ada50
		return &TLStoriesGetPeerStories{
			Constructor: 743103056,
		}
	},
	-1688541191: func() TLObject { // 0x9b5ae7f9
		return &TLStoriesGetAllReadPeerStories{
			Constructor: -1688541191,
		}
	},
	1398375363: func() TLObject { // 0x535983c3
		return &TLStoriesGetPeerMaxIDs{
			Constructor: 1398375363,
		}
	},
	-1519744160: func() TLObject { // 0xa56a8b60
		return &TLStoriesGetChatsToSend{
			Constructor: -1519744160,
		}
	},
	-1123805756: func() TLObject { // 0xbd0415c4
		return &TLStoriesTogglePeerStoriesHidden{
			Constructor: -1123805756,
		}
	},
	-1179482081: func() TLObject { // 0xb9b2881f
		return &TLStoriesGetStoryReactionsList{
			Constructor: -1179482081,
		}
	},
	187268763: func() TLObject { // 0xb297e9b
		return &TLStoriesTogglePinnedToTop{
			Constructor: 187268763,
		}
	},
	1827279210: func() TLObject { // 0x6cea116a
		return &TLStoriesSearchPosts{
			Constructor: 1827279210,
		}
	},
	1626764896: func() TLObject { // 0x60f67660
		return &TLPremiumGetBoostsList{
			Constructor: 1626764896,
		}
	},
	199719754: func() TLObject { // 0xbe77b4a
		return &TLPremiumGetMyBoosts{
			Constructor: 199719754,
		}
	},
	1803396934: func() TLObject { // 0x6b7da746
		return &TLPremiumApplyBoost{
			Constructor: 1803396934,
		}
	},
	70197089: func() TLObject { // 0x42f1f61
		return &TLPremiumGetBoostsStatus{
			Constructor: 70197089,
		}
	},
	965037343: func() TLObject { // 0x39854d1f
		return &TLPremiumGetUserBoosts{
			Constructor: 965037343,
		}
	},
	249313744: func() TLObject { // 0xedc39d0
		return &TLSmsjobsIsEligibleToJoin{
			Constructor: 249313744,
		}
	},
	-1488007635: func() TLObject { // 0xa74ece2d
		return &TLSmsjobsJoin{
			Constructor: -1488007635,
		}
	},
	-1734824589: func() TLObject { // 0x9898ad73
		return &TLSmsjobsLeave{
			Constructor: -1734824589,
		}
	},
	155164863: func() TLObject { // 0x93fa0bf
		return &TLSmsjobsUpdateSettings{
			Constructor: 155164863,
		}
	},
	279353576: func() TLObject { // 0x10a698e8
		return &TLSmsjobsGetStatus{
			Constructor: 279353576,
		}
	},
	2005766191: func() TLObject { // 0x778d902f
		return &TLSmsjobsGetSmsJob{
			Constructor: 2005766191,
		}
	},
	1327415076: func() TLObject { // 0x4f1ebf24
		return &TLSmsjobsFinishJob{
			Constructor: 1327415076,
		}
	},
	-1105295942: func() TLObject { // 0xbe1e85ba
		return &TLFragmentGetCollectibleInfo{
			Constructor: -1105295942,
		}
	},
	440815626: func() TLObject { // 0x1a46500a
		return &TLMessagesRequestSimpleWebView1A46500A{
			Constructor: 440815626,
		}
	},
	-1940243652: func() TLObject { // 0x8c5a3b3c
		return &TLMessagesRequestAppWebView8C5A3B3C{
			Constructor: -1940243652,
		}
	},
	-230206493: func() TLObject { // 0xf24753e3
		return &TLMessagesAddChatUserF24753E3{
			Constructor: -230206493,
		}
	},
	3450904: func() TLObject { // 0x34a818
		return &TLMessagesCreateChat34A818{
			Constructor: 3450904,
		}
	},
	429865580: func() TLObject { // 0x199f3a6c
		return &TLChannelsInviteToChannel199F3A6C{
			Constructor: 429865580,
		}
	},
	-241247891: func() TLObject { // 0xf19ed96d
		return &TLMessagesGetDialogFiltersF19ED96D{
			Constructor: -241247891,
		}
	},
	-1877938321: func() TLObject { // 0x9010ef6f
		return &TLHelpGetAppChangelog{
			Constructor: -1877938321,
		}
	},
	1445996571: func() TLObject { // 0x5630281b
		return &TLStatsGetMessagePublicForwards5630281B{
			Constructor: 1445996571,
		}
	},
	1279562866: func() TLObject { // 0x4c449472
		return &TLStoriesGetBoostsStatus{
			Constructor: 1279562866,
		}
	},
	863959424: func() TLObject { // 0x337ef980
		return &TLStoriesGetBoostersList{
			Constructor: 863959424,
		}
	},
	-620379715: func() TLObject { // 0xdb05c1bd
		return &TLStoriesCanApplyBoost{
			Constructor: -620379715,
		}
	},
	-224560085: func() TLObject { // 0xf29d7c2b
		return &TLStoriesApplyBoost{
			Constructor: -224560085,
		}
	},
	852135825: func() TLObject { // 0x32ca8f91
		return &TLMessagesGetWebPage32CA8F91{
			Constructor: 852135825,
		}
	},
	-904087125: func() TLObject { // 0xca1cb9ab
		return &TLUsersGetStoriesMaxIDs{
			Constructor: -904087125,
		}
	},
	1967110245: func() TLObject { // 0x753fb865
		return &TLContactsToggleStoriesHidden{
			Constructor: 1967110245,
		}
	},
	-1764415264: func() TLObject { // 0x96d528e0
		return &TLStoriesGetUserStories{
			Constructor: -1764415264,
		}
	},
	1922848300: func() TLObject { // 0x729c562c
		return &TLStoriesGetAllReadUserStories{
			Constructor: 1922848300,
		}
	},
	698084494: func() TLObject { // 0x299bec8e
		return &TLMessagesRequestSimpleWebView299BEC8E{
			Constructor: 698084494,
		}
	},
	-2023787330: func() TLObject { // 0x875f74be
		return &TLMessagesGetAllChats{
			Constructor: -2023787330,
		}
	},
	1978405606: func() TLObject { // 0x75ec12e6
		return &TLBotsGetBotInfo75EC12E6{
			Constructor: 1978405606,
		}
	},
	472471681: func() TLObject { // 0x1c295881
		return &TLFoldersDeleteFolder{
			Constructor: 472471681,
		}
	},
	745510839: func() TLObject { // 0x2c6f97b7
		return &TLMessagesGetMessageReadParticipants2C6F97B7{
			Constructor: 745510839,
		}
	},
	-1735311088: func() TLObject { // 0x98914110
		return &TLHelpGetAppConfig98914110{
			Constructor: -1735311088,
		}
	},
	164303470: func() TLObject { // 0x9cb126e
		return &TLMessagesCreateChat9CB126E{
			Constructor: 164303470,
		}
	},
	-323339813: func() TLObject { // 0xecba39db
		return &TLAccountVerifyEmailECBA39DB{
			Constructor: -323339813,
		}
	},
	1790652275: func() TLObject { // 0x6abb2f73
		return &TLMessagesRequestSimpleWebView6ABB2F73{
			Constructor: 1790652275,
		}
	},
	342791565: func() TLObject { // 0x146e958d
		return &TLPaymentsRequestRecurringPayment{
			Constructor: 342791565,
		}
	},
	-781917334: func() TLObject { // 0xd164e36a
		return &TLPaymentsRestorePlayMarketReceipt{
			Constructor: -781917334,
		}
	},
	-1355375294: func() TLObject { // 0xaf369d42
		return &TLChannelsDeleteHistoryAF369D42{
			Constructor: -1355375294,
		}
	},
	-1058929929: func() TLObject { // 0xc0e202f7
		return &TLHelpTest{
			Constructor: -1058929929,
		}
	},
	602071838: func() TLObject { // 0x23e2e31e
		return &TLPredefinedCreatePredefinedUser{
			Constructor: 602071838,
		}
	},
	316411194: func() TLObject { // 0x12dc0d3a
		return &TLPredefinedUpdatePredefinedUsername{
			Constructor: 316411194,
		}
	},
	752679237: func() TLObject { // 0x2cdcf945
		return &TLPredefinedUpdatePredefinedProfile{
			Constructor: 752679237,
		}
	},
	1060448425: func() TLObject { // 0x3f3528a9
		return &TLPredefinedUpdatePredefinedVerified{
			Constructor: 1060448425,
		}
	},
	-449440377: func() TLObject { // 0xe5361587
		return &TLPredefinedUpdatePredefinedCode{
			Constructor: -449440377,
		}
	},
	1375904789: func() TLObject { // 0x5202a415
		return &TLPredefinedGetPredefinedUser{
			Constructor: 1375904789,
		}
	},
	697834180: func() TLObject { // 0x29981ac4
		return &TLPredefinedGetPredefinedUsers{
			Constructor: 697834180,
		}
	},
	825513746: func() TLObject { // 0x31345712
		return &TLUsersGetMe{
			Constructor: 825513746,
		}
	},
	353634673: func() TLObject { // 0x15140971
		return &TLAccountUpdateVerified{
			Constructor: 353634673,
		}
	},
	-501253832: func() TLObject { // 0xe21f7938
		return &TLAuthToggleBan{
			Constructor: -501253832,
		}
	},
	1511592262: func() TLObject { // 0x5a191146
		return &TLBizInvokeBizDataRaw{
			Constructor: 1511592262,
		}
	},
}

func NewTLObjectByClassID(classId int32) TLObject {
	f, ok := clazzIdRegisters2[classId]
	if !ok {
		return nil
	}
	return f()
}

func CheckClassID(classId int32) (ok bool) {
	_, ok = clazzIdRegisters2[classId]
	return
}
