/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2024-present,  Teamgram Authors.
 *  All rights reserved.
 *
 * Author: Benqi (wubenqi@gmail.com)
 */

package mt

import (
	"github.com/teamgram/proto/v2/iface"
)

const (
	ClazzName_resPQ                      = "resPQ"
	ClazzName_p_q_inner_data             = "p_q_inner_data"
	ClazzName_p_q_inner_data_dc          = "p_q_inner_data_dc"
	ClazzName_p_q_inner_data_temp        = "p_q_inner_data_temp"
	ClazzName_p_q_inner_data_temp_dc     = "p_q_inner_data_temp_dc"
	ClazzName_bind_auth_key_inner        = "bind_auth_key_inner"
	ClazzName_server_DH_params_fail      = "server_DH_params_fail"
	ClazzName_server_DH_params_ok        = "server_DH_params_ok"
	ClazzName_server_DH_inner_data       = "server_DH_inner_data"
	ClazzName_client_DH_inner_data       = "client_DH_inner_data"
	ClazzName_dh_gen_ok                  = "dh_gen_ok"
	ClazzName_dh_gen_retry               = "dh_gen_retry"
	ClazzName_dh_gen_fail                = "dh_gen_fail"
	ClazzName_destroy_auth_key_ok        = "destroy_auth_key_ok"
	ClazzName_destroy_auth_key_none      = "destroy_auth_key_none"
	ClazzName_destroy_auth_key_fail      = "destroy_auth_key_fail"
	ClazzName_req_pq                     = "req_pq"
	ClazzName_req_pq_multi               = "req_pq_multi"
	ClazzName_req_DH_params              = "req_DH_params"
	ClazzName_set_client_DH_params       = "set_client_DH_params"
	ClazzName_destroy_auth_key           = "destroy_auth_key"
	ClazzName_msgs_ack                   = "msgs_ack"
	ClazzName_bad_msg_notification       = "bad_msg_notification"
	ClazzName_bad_server_salt            = "bad_server_salt"
	ClazzName_msgs_state_req             = "msgs_state_req"
	ClazzName_msgs_state_info            = "msgs_state_info"
	ClazzName_msgs_all_info              = "msgs_all_info"
	ClazzName_msg_detailed_info          = "msg_detailed_info"
	ClazzName_msg_new_detailed_info      = "msg_new_detailed_info"
	ClazzName_msg_resend_req             = "msg_resend_req"
	ClazzName_rpc_error                  = "rpc_error"
	ClazzName_rpc_answer_unknown         = "rpc_answer_unknown"
	ClazzName_rpc_answer_dropped_running = "rpc_answer_dropped_running"
	ClazzName_rpc_answer_dropped         = "rpc_answer_dropped"
	ClazzName_future_salt                = "future_salt"
	ClazzName_future_salts               = "future_salts"
	ClazzName_pong                       = "pong"
	ClazzName_destroy_session_ok         = "destroy_session_ok"
	ClazzName_destroy_session_none       = "destroy_session_none"
	ClazzName_new_session_created        = "new_session_created"
	ClazzName_http_wait                  = "http_wait"
	ClazzName_ipPort                     = "ipPort"
	ClazzName_ipPortSecret               = "ipPortSecret"
	ClazzName_accessPointRule            = "accessPointRule"
	ClazzName_help_configSimple          = "help_configSimple"
	ClazzName_tlsClientHello             = "tlsClientHello"
	ClazzName_tlsBlockString             = "tlsBlockString"
	ClazzName_tlsBlockRandom             = "tlsBlockRandom"
	ClazzName_tlsBlockZero               = "tlsBlockZero"
	ClazzName_tlsBlockDomain             = "tlsBlockDomain"
	ClazzName_tlsBlockGrease             = "tlsBlockGrease"
	ClazzName_tlsBlockPublicKey          = "tlsBlockPublicKey"
	ClazzName_tlsBlockScope              = "tlsBlockScope"
	ClazzName_rpc_drop_answer            = "rpc_drop_answer"
	ClazzName_get_future_salts           = "get_future_salts"
	ClazzName_ping                       = "ping"
	ClazzName_ping_delay_disconnect      = "ping_delay_disconnect"
	ClazzName_destroy_session            = "destroy_session"
	ClazzName_test_useError              = "test_useError"
	ClazzName_test_useConfigSimple       = "test_useConfigSimple"
	ClazzName_test_parseInputAppEvent    = "test_parseInputAppEvent"
)

func init() {
	// RegisterClazzNameList
	iface.RegisterClazzName(ClazzName_resPQ, 0, 85337187)                        // 5162463
	iface.RegisterClazzName(ClazzName_p_q_inner_data, 0, -2083955988)            // 83c95aec
	iface.RegisterClazzName(ClazzName_p_q_inner_data_dc, 0, -1443537003)         // a9f55f95
	iface.RegisterClazzName(ClazzName_p_q_inner_data_temp, 0, 1013613780)        // 3c6a84d4
	iface.RegisterClazzName(ClazzName_p_q_inner_data_temp_dc, 0, 1459478408)     // 56fddf88
	iface.RegisterClazzName(ClazzName_bind_auth_key_inner, 0, 1973679973)        // 75a3f765
	iface.RegisterClazzName(ClazzName_server_DH_params_fail, 0, 2043348061)      // 79cb045d
	iface.RegisterClazzName(ClazzName_server_DH_params_ok, 0, -790100132)        // d0e8075c
	iface.RegisterClazzName(ClazzName_server_DH_inner_data, 0, -1249309254)      // b5890dba
	iface.RegisterClazzName(ClazzName_client_DH_inner_data, 0, 1715713620)       // 6643b654
	iface.RegisterClazzName(ClazzName_dh_gen_ok, 0, 1003222836)                  // 3bcbf734
	iface.RegisterClazzName(ClazzName_dh_gen_retry, 0, 1188831161)               // 46dc1fb9
	iface.RegisterClazzName(ClazzName_dh_gen_fail, 0, -1499615742)               // a69dae02
	iface.RegisterClazzName(ClazzName_destroy_auth_key_ok, 0, -161422892)        // f660e1d4
	iface.RegisterClazzName(ClazzName_destroy_auth_key_none, 0, 178201177)       // a9f2259
	iface.RegisterClazzName(ClazzName_destroy_auth_key_fail, 0, -368010477)      // ea109b13
	iface.RegisterClazzName(ClazzName_req_pq, 0, 1615239032)                     // 60469778
	iface.RegisterClazzName(ClazzName_req_pq_multi, 0, -1099002127)              // be7e8ef1
	iface.RegisterClazzName(ClazzName_req_DH_params, 0, -686627650)              // d712e4be
	iface.RegisterClazzName(ClazzName_set_client_DH_params, 0, -184262881)       // f5045f1f
	iface.RegisterClazzName(ClazzName_destroy_auth_key, 0, -784117408)           // d1435160
	iface.RegisterClazzName(ClazzName_msgs_ack, 0, 1658238041)                   // 62d6b459
	iface.RegisterClazzName(ClazzName_bad_msg_notification, 0, -1477445615)      // a7eff811
	iface.RegisterClazzName(ClazzName_bad_server_salt, 0, -307542917)            // edab447b
	iface.RegisterClazzName(ClazzName_msgs_state_req, 0, -630588590)             // da69fb52
	iface.RegisterClazzName(ClazzName_msgs_state_info, 0, 81704317)              // 4deb57d
	iface.RegisterClazzName(ClazzName_msgs_all_info, 0, -1933520591)             // 8cc0d131
	iface.RegisterClazzName(ClazzName_msg_detailed_info, 0, 661470918)           // 276d3ec6
	iface.RegisterClazzName(ClazzName_msg_new_detailed_info, 0, -2137147681)     // 809db6df
	iface.RegisterClazzName(ClazzName_msg_resend_req, 0, 2105940488)             // 7d861a08
	iface.RegisterClazzName(ClazzName_rpc_error, 0, 558156313)                   // 2144ca19
	iface.RegisterClazzName(ClazzName_rpc_answer_unknown, 0, 1579864942)         // 5e2ad36e
	iface.RegisterClazzName(ClazzName_rpc_answer_dropped_running, 0, -847714938) // cd78e586
	iface.RegisterClazzName(ClazzName_rpc_answer_dropped, 0, -1539647305)        // a43ad8b7
	iface.RegisterClazzName(ClazzName_future_salt, 0, 155834844)                 // 949d9dc
	iface.RegisterClazzName(ClazzName_future_salts, 0, -1370486635)              // ae500895
	iface.RegisterClazzName(ClazzName_pong, 0, 880243653)                        // 347773c5
	iface.RegisterClazzName(ClazzName_destroy_session_ok, 0, -501201412)         // e22045fc
	iface.RegisterClazzName(ClazzName_destroy_session_none, 0, 1658015945)       // 62d350c9
	iface.RegisterClazzName(ClazzName_new_session_created, 0, -1631450872)       // 9ec20908
	iface.RegisterClazzName(ClazzName_http_wait, 0, -1835453025)                 // 9299359f
	iface.RegisterClazzName(ClazzName_ipPort, 0, -734810765)                     // d433ad73
	iface.RegisterClazzName(ClazzName_ipPortSecret, 0, 932718150)                // 37982646
	iface.RegisterClazzName(ClazzName_accessPointRule, 0, 1182381663)            // 4679b65f
	iface.RegisterClazzName(ClazzName_help_configSimple, 0, 1515793004)          // 5a592a6c
	iface.RegisterClazzName(ClazzName_tlsClientHello, 0, 1817363588)             // 0x6c52c484
	iface.RegisterClazzName(ClazzName_tlsBlockString, 0, 1108910436)             // 0x4218a164
	iface.RegisterClazzName(ClazzName_tlsBlockRandom, 0, 1296942110)             // 0x4d4dc41e
	iface.RegisterClazzName(ClazzName_tlsBlockZero, 0, 154352379)                // 0x9333afb
	iface.RegisterClazzName(ClazzName_tlsBlockDomain, 0, 283665263)              // 0x10e8636f
	iface.RegisterClazzName(ClazzName_tlsBlockGrease, 0, -428498495)             // 0xe675a1c1
	iface.RegisterClazzName(ClazzName_tlsBlockPublicKey, 0, -1632019620)         // 0x9eb95b5c
	iface.RegisterClazzName(ClazzName_tlsBlockScope, 0, -416951217)              // 0xe725d44f
	iface.RegisterClazzName(ClazzName_rpc_drop_answer, 0, 1491380032)            // 58e4a740
	iface.RegisterClazzName(ClazzName_get_future_salts, 0, -1188971260)          // b921bd04
	iface.RegisterClazzName(ClazzName_ping, 0, 2059302892)                       // 7abe77ec
	iface.RegisterClazzName(ClazzName_ping_delay_disconnect, 0, -213746804)      // f3427b8c
	iface.RegisterClazzName(ClazzName_destroy_session, 0, -414113498)            // e7512126
	iface.RegisterClazzName(ClazzName_test_useError, 0, -294277375)              // 0xee75af01
	iface.RegisterClazzName(ClazzName_test_useConfigSimple, 0, -105401795)       // 0xf9b7b23d
	iface.RegisterClazzName(ClazzName_test_parseInputAppEvent, 0, -1156741135)   // 0xbb0d87f1

	//RegisterClazzIDNameList
	iface.RegisterClazzIDName(ClazzName_resPQ, 85337187)                        // 5162463
	iface.RegisterClazzIDName(ClazzName_p_q_inner_data, -2083955988)            // 83c95aec
	iface.RegisterClazzIDName(ClazzName_p_q_inner_data_dc, -1443537003)         // a9f55f95
	iface.RegisterClazzIDName(ClazzName_p_q_inner_data_temp, 1013613780)        // 3c6a84d4
	iface.RegisterClazzIDName(ClazzName_p_q_inner_data_temp_dc, 1459478408)     // 56fddf88
	iface.RegisterClazzIDName(ClazzName_bind_auth_key_inner, 1973679973)        // 75a3f765
	iface.RegisterClazzIDName(ClazzName_server_DH_params_fail, 2043348061)      // 79cb045d
	iface.RegisterClazzIDName(ClazzName_server_DH_params_ok, -790100132)        // d0e8075c
	iface.RegisterClazzIDName(ClazzName_server_DH_inner_data, -1249309254)      // b5890dba
	iface.RegisterClazzIDName(ClazzName_client_DH_inner_data, 1715713620)       // 6643b654
	iface.RegisterClazzIDName(ClazzName_dh_gen_ok, 1003222836)                  // 3bcbf734
	iface.RegisterClazzIDName(ClazzName_dh_gen_retry, 1188831161)               // 46dc1fb9
	iface.RegisterClazzIDName(ClazzName_dh_gen_fail, -1499615742)               // a69dae02
	iface.RegisterClazzIDName(ClazzName_destroy_auth_key_ok, -161422892)        // f660e1d4
	iface.RegisterClazzIDName(ClazzName_destroy_auth_key_none, 178201177)       // a9f2259
	iface.RegisterClazzIDName(ClazzName_destroy_auth_key_fail, -368010477)      // ea109b13
	iface.RegisterClazzIDName(ClazzName_req_pq, 1615239032)                     // 60469778
	iface.RegisterClazzIDName(ClazzName_req_pq_multi, -1099002127)              // be7e8ef1
	iface.RegisterClazzIDName(ClazzName_req_DH_params, -686627650)              // d712e4be
	iface.RegisterClazzIDName(ClazzName_set_client_DH_params, -184262881)       // f5045f1f
	iface.RegisterClazzIDName(ClazzName_destroy_auth_key, -784117408)           // d1435160
	iface.RegisterClazzIDName(ClazzName_msgs_ack, 1658238041)                   // 62d6b459
	iface.RegisterClazzIDName(ClazzName_bad_msg_notification, -1477445615)      // a7eff811
	iface.RegisterClazzIDName(ClazzName_bad_server_salt, -307542917)            // edab447b
	iface.RegisterClazzIDName(ClazzName_msgs_state_req, -630588590)             // da69fb52
	iface.RegisterClazzIDName(ClazzName_msgs_state_info, 81704317)              // 4deb57d
	iface.RegisterClazzIDName(ClazzName_msgs_all_info, -1933520591)             // 8cc0d131
	iface.RegisterClazzIDName(ClazzName_msg_detailed_info, 661470918)           // 276d3ec6
	iface.RegisterClazzIDName(ClazzName_msg_new_detailed_info, -2137147681)     // 809db6df
	iface.RegisterClazzIDName(ClazzName_msg_resend_req, 2105940488)             // 7d861a08
	iface.RegisterClazzIDName(ClazzName_rpc_error, 558156313)                   // 2144ca19
	iface.RegisterClazzIDName(ClazzName_rpc_answer_unknown, 1579864942)         // 5e2ad36e
	iface.RegisterClazzIDName(ClazzName_rpc_answer_dropped_running, -847714938) // cd78e586
	iface.RegisterClazzIDName(ClazzName_rpc_answer_dropped, -1539647305)        // a43ad8b7
	iface.RegisterClazzIDName(ClazzName_future_salt, 155834844)                 // 949d9dc
	iface.RegisterClazzIDName(ClazzName_future_salts, -1370486635)              // ae500895
	iface.RegisterClazzIDName(ClazzName_pong, 880243653)                        // 347773c5
	iface.RegisterClazzIDName(ClazzName_destroy_session_ok, -501201412)         // e22045fc
	iface.RegisterClazzIDName(ClazzName_destroy_session_none, 1658015945)       // 62d350c9
	iface.RegisterClazzIDName(ClazzName_new_session_created, -1631450872)       // 9ec20908
	iface.RegisterClazzIDName(ClazzName_http_wait, -1835453025)                 // 9299359f
	iface.RegisterClazzIDName(ClazzName_ipPort, -734810765)                     // d433ad73
	iface.RegisterClazzIDName(ClazzName_ipPortSecret, 932718150)                // 37982646
	iface.RegisterClazzIDName(ClazzName_accessPointRule, 1182381663)            // 4679b65f
	iface.RegisterClazzIDName(ClazzName_help_configSimple, 1515793004)          // 5a592a6c
	iface.RegisterClazzIDName(ClazzName_tlsClientHello, 1817363588)             // 0x6c52c484
	iface.RegisterClazzIDName(ClazzName_tlsBlockString, 1108910436)             // 0x4218a164
	iface.RegisterClazzIDName(ClazzName_tlsBlockRandom, 1296942110)             // 0x4d4dc41e
	iface.RegisterClazzIDName(ClazzName_tlsBlockZero, 154352379)                // 0x9333afb
	iface.RegisterClazzIDName(ClazzName_tlsBlockDomain, 283665263)              // 0x10e8636f
	iface.RegisterClazzIDName(ClazzName_tlsBlockGrease, -428498495)             // 0xe675a1c1
	iface.RegisterClazzIDName(ClazzName_tlsBlockPublicKey, -1632019620)         // 0x9eb95b5c
	iface.RegisterClazzIDName(ClazzName_tlsBlockScope, -416951217)              // 0xe725d44f
	iface.RegisterClazzIDName(ClazzName_rpc_drop_answer, 1491380032)            // 58e4a740
	iface.RegisterClazzIDName(ClazzName_get_future_salts, -1188971260)          // b921bd04
	iface.RegisterClazzIDName(ClazzName_ping, 2059302892)                       // 7abe77ec
	iface.RegisterClazzIDName(ClazzName_ping_delay_disconnect, -213746804)      // f3427b8c
	iface.RegisterClazzIDName(ClazzName_destroy_session, -414113498)            // e7512126
	iface.RegisterClazzIDName(ClazzName_test_useError, -294277375)              // 0xee75af01
	iface.RegisterClazzIDName(ClazzName_test_useConfigSimple, -105401795)       // 0xf9b7b23d
	iface.RegisterClazzIDName(ClazzName_test_parseInputAppEvent, -1156741135)   // 0xbb0d87f1
}
