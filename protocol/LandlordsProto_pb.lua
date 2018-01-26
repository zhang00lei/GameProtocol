-- Generated By protoc-gen-lua Do not Edit
local protobuf = require "protobuf"
module('LandlordsProto_pb')


local CS_PLAYERLOGIN = protobuf.Descriptor();
local CS_PLAYERLOGIN_ACCOUNT_FIELD = protobuf.FieldDescriptor();
local CS_PLAYERLOGIN_PWD_FIELD = protobuf.FieldDescriptor();
local SC_PLAYERLOGIN = protobuf.Descriptor();
local SC_PLAYERLOGIN_RESULT_FIELD = protobuf.FieldDescriptor();
local SC_PLAYERLOGIN_PLAYERID_FIELD = protobuf.FieldDescriptor();
local SC_PLAYERLOGIN_TEMP_FIELD = protobuf.FieldDescriptor();
local CS_PLAYERREGISTER = protobuf.Descriptor();
local CS_PLAYERREGISTER_ACCOUNT_FIELD = protobuf.FieldDescriptor();
local CS_PLAYERREGISTER_PWD_FIELD = protobuf.FieldDescriptor();
local SC_PLAYERREGISTER = protobuf.Descriptor();
local SC_PLAYERREGISTER_RESULT_FIELD = protobuf.FieldDescriptor();
local PLAYERINFO = protobuf.Descriptor();
local PLAYERINFO_PLAYERID_FIELD = protobuf.FieldDescriptor();
local PLAYERINFO_PLAYERACCOUNT_FIELD = protobuf.FieldDescriptor();
local PLAYERINFO_PLAYERNAME_FIELD = protobuf.FieldDescriptor();
local PLAYERINFO_PLAYERCOIN_FIELD = protobuf.FieldDescriptor();
local CS_PLAYERINFO = protobuf.Descriptor();
local CS_PLAYERINFO_PLAYERID_FIELD = protobuf.FieldDescriptor();
local SC_PLAYERINFO = protobuf.Descriptor();
local SC_PLAYERINFO_PLAYERINFO_FIELD = protobuf.FieldDescriptor();
local SC_PLAYERINFO_RESULT_FIELD = protobuf.FieldDescriptor();
local CS_HEARTBEAT = protobuf.Descriptor();
local SC_HEARTBEAT = protobuf.Descriptor();
local SC_HEARTBEAT_SERVERTIME_FIELD = protobuf.FieldDescriptor();
local CS_PLAYERREADY = protobuf.Descriptor();
local CS_PLAYERREADY_PLAYERID_FIELD = protobuf.FieldDescriptor();
local SC_PLAYERREADY = protobuf.Descriptor();
local SC_PLAYERREADY_RESULT_FIELD = protobuf.FieldDescriptor();
local CS_PLAYERCANCELREADY = protobuf.Descriptor();
local CS_PLAYERCANCELREADY_PLAYERID_FIELD = protobuf.FieldDescriptor();
local SC_PLAYERCANCELREADY = protobuf.Descriptor();
local SC_PLAYERCANCELREADY_RESULT_FIELD = protobuf.FieldDescriptor();
local SC_PLAYERCARD = protobuf.Descriptor();
local SC_PLAYERCARD_CARDID_FIELD = protobuf.FieldDescriptor();

CS_PLAYERLOGIN_ACCOUNT_FIELD.name = "account"
CS_PLAYERLOGIN_ACCOUNT_FIELD.full_name = ".msg.CS_PlayerLogin.account"
CS_PLAYERLOGIN_ACCOUNT_FIELD.number = 1
CS_PLAYERLOGIN_ACCOUNT_FIELD.index = 0
CS_PLAYERLOGIN_ACCOUNT_FIELD.label = 1
CS_PLAYERLOGIN_ACCOUNT_FIELD.has_default_value = false
CS_PLAYERLOGIN_ACCOUNT_FIELD.default_value = ""
CS_PLAYERLOGIN_ACCOUNT_FIELD.type = 9
CS_PLAYERLOGIN_ACCOUNT_FIELD.cpp_type = 9

CS_PLAYERLOGIN_PWD_FIELD.name = "pwd"
CS_PLAYERLOGIN_PWD_FIELD.full_name = ".msg.CS_PlayerLogin.pwd"
CS_PLAYERLOGIN_PWD_FIELD.number = 2
CS_PLAYERLOGIN_PWD_FIELD.index = 1
CS_PLAYERLOGIN_PWD_FIELD.label = 1
CS_PLAYERLOGIN_PWD_FIELD.has_default_value = false
CS_PLAYERLOGIN_PWD_FIELD.default_value = ""
CS_PLAYERLOGIN_PWD_FIELD.type = 9
CS_PLAYERLOGIN_PWD_FIELD.cpp_type = 9

CS_PLAYERLOGIN.name = "CS_PlayerLogin"
CS_PLAYERLOGIN.full_name = ".msg.CS_PlayerLogin"
CS_PLAYERLOGIN.nested_types = {}
CS_PLAYERLOGIN.enum_types = {}
CS_PLAYERLOGIN.fields = {CS_PLAYERLOGIN_ACCOUNT_FIELD, CS_PLAYERLOGIN_PWD_FIELD}
CS_PLAYERLOGIN.is_extendable = false
CS_PLAYERLOGIN.extensions = {}
SC_PLAYERLOGIN_RESULT_FIELD.name = "result"
SC_PLAYERLOGIN_RESULT_FIELD.full_name = ".msg.SC_PlayerLogin.result"
SC_PLAYERLOGIN_RESULT_FIELD.number = 1
SC_PLAYERLOGIN_RESULT_FIELD.index = 0
SC_PLAYERLOGIN_RESULT_FIELD.label = 1
SC_PLAYERLOGIN_RESULT_FIELD.has_default_value = false
SC_PLAYERLOGIN_RESULT_FIELD.default_value = 0
SC_PLAYERLOGIN_RESULT_FIELD.type = 17
SC_PLAYERLOGIN_RESULT_FIELD.cpp_type = 1

SC_PLAYERLOGIN_PLAYERID_FIELD.name = "playerId"
SC_PLAYERLOGIN_PLAYERID_FIELD.full_name = ".msg.SC_PlayerLogin.playerId"
SC_PLAYERLOGIN_PLAYERID_FIELD.number = 2
SC_PLAYERLOGIN_PLAYERID_FIELD.index = 1
SC_PLAYERLOGIN_PLAYERID_FIELD.label = 1
SC_PLAYERLOGIN_PLAYERID_FIELD.has_default_value = false
SC_PLAYERLOGIN_PLAYERID_FIELD.default_value = 0
SC_PLAYERLOGIN_PLAYERID_FIELD.type = 4
SC_PLAYERLOGIN_PLAYERID_FIELD.cpp_type = 4

SC_PLAYERLOGIN_TEMP_FIELD.name = "temp"
SC_PLAYERLOGIN_TEMP_FIELD.full_name = ".msg.SC_PlayerLogin.temp"
SC_PLAYERLOGIN_TEMP_FIELD.number = 3
SC_PLAYERLOGIN_TEMP_FIELD.index = 2
SC_PLAYERLOGIN_TEMP_FIELD.label = 3
SC_PLAYERLOGIN_TEMP_FIELD.has_default_value = false
SC_PLAYERLOGIN_TEMP_FIELD.default_value = {}
SC_PLAYERLOGIN_TEMP_FIELD.type = 17
SC_PLAYERLOGIN_TEMP_FIELD.cpp_type = 1

SC_PLAYERLOGIN.name = "SC_PlayerLogin"
SC_PLAYERLOGIN.full_name = ".msg.SC_PlayerLogin"
SC_PLAYERLOGIN.nested_types = {}
SC_PLAYERLOGIN.enum_types = {}
SC_PLAYERLOGIN.fields = {SC_PLAYERLOGIN_RESULT_FIELD, SC_PLAYERLOGIN_PLAYERID_FIELD, SC_PLAYERLOGIN_TEMP_FIELD}
SC_PLAYERLOGIN.is_extendable = false
SC_PLAYERLOGIN.extensions = {}
CS_PLAYERREGISTER_ACCOUNT_FIELD.name = "account"
CS_PLAYERREGISTER_ACCOUNT_FIELD.full_name = ".msg.CS_PlayerRegister.account"
CS_PLAYERREGISTER_ACCOUNT_FIELD.number = 1
CS_PLAYERREGISTER_ACCOUNT_FIELD.index = 0
CS_PLAYERREGISTER_ACCOUNT_FIELD.label = 1
CS_PLAYERREGISTER_ACCOUNT_FIELD.has_default_value = false
CS_PLAYERREGISTER_ACCOUNT_FIELD.default_value = ""
CS_PLAYERREGISTER_ACCOUNT_FIELD.type = 9
CS_PLAYERREGISTER_ACCOUNT_FIELD.cpp_type = 9

CS_PLAYERREGISTER_PWD_FIELD.name = "pwd"
CS_PLAYERREGISTER_PWD_FIELD.full_name = ".msg.CS_PlayerRegister.pwd"
CS_PLAYERREGISTER_PWD_FIELD.number = 2
CS_PLAYERREGISTER_PWD_FIELD.index = 1
CS_PLAYERREGISTER_PWD_FIELD.label = 1
CS_PLAYERREGISTER_PWD_FIELD.has_default_value = false
CS_PLAYERREGISTER_PWD_FIELD.default_value = ""
CS_PLAYERREGISTER_PWD_FIELD.type = 9
CS_PLAYERREGISTER_PWD_FIELD.cpp_type = 9

CS_PLAYERREGISTER.name = "CS_PlayerRegister"
CS_PLAYERREGISTER.full_name = ".msg.CS_PlayerRegister"
CS_PLAYERREGISTER.nested_types = {}
CS_PLAYERREGISTER.enum_types = {}
CS_PLAYERREGISTER.fields = {CS_PLAYERREGISTER_ACCOUNT_FIELD, CS_PLAYERREGISTER_PWD_FIELD}
CS_PLAYERREGISTER.is_extendable = false
CS_PLAYERREGISTER.extensions = {}
SC_PLAYERREGISTER_RESULT_FIELD.name = "result"
SC_PLAYERREGISTER_RESULT_FIELD.full_name = ".msg.SC_PlayerRegister.result"
SC_PLAYERREGISTER_RESULT_FIELD.number = 1
SC_PLAYERREGISTER_RESULT_FIELD.index = 0
SC_PLAYERREGISTER_RESULT_FIELD.label = 1
SC_PLAYERREGISTER_RESULT_FIELD.has_default_value = false
SC_PLAYERREGISTER_RESULT_FIELD.default_value = 0
SC_PLAYERREGISTER_RESULT_FIELD.type = 17
SC_PLAYERREGISTER_RESULT_FIELD.cpp_type = 1

SC_PLAYERREGISTER.name = "SC_PlayerRegister"
SC_PLAYERREGISTER.full_name = ".msg.SC_PlayerRegister"
SC_PLAYERREGISTER.nested_types = {}
SC_PLAYERREGISTER.enum_types = {}
SC_PLAYERREGISTER.fields = {SC_PLAYERREGISTER_RESULT_FIELD}
SC_PLAYERREGISTER.is_extendable = false
SC_PLAYERREGISTER.extensions = {}
PLAYERINFO_PLAYERID_FIELD.name = "playerId"
PLAYERINFO_PLAYERID_FIELD.full_name = ".msg.PlayerInfo.playerId"
PLAYERINFO_PLAYERID_FIELD.number = 1
PLAYERINFO_PLAYERID_FIELD.index = 0
PLAYERINFO_PLAYERID_FIELD.label = 1
PLAYERINFO_PLAYERID_FIELD.has_default_value = false
PLAYERINFO_PLAYERID_FIELD.default_value = 0
PLAYERINFO_PLAYERID_FIELD.type = 4
PLAYERINFO_PLAYERID_FIELD.cpp_type = 4

PLAYERINFO_PLAYERACCOUNT_FIELD.name = "playerAccount"
PLAYERINFO_PLAYERACCOUNT_FIELD.full_name = ".msg.PlayerInfo.playerAccount"
PLAYERINFO_PLAYERACCOUNT_FIELD.number = 2
PLAYERINFO_PLAYERACCOUNT_FIELD.index = 1
PLAYERINFO_PLAYERACCOUNT_FIELD.label = 1
PLAYERINFO_PLAYERACCOUNT_FIELD.has_default_value = false
PLAYERINFO_PLAYERACCOUNT_FIELD.default_value = ""
PLAYERINFO_PLAYERACCOUNT_FIELD.type = 9
PLAYERINFO_PLAYERACCOUNT_FIELD.cpp_type = 9

PLAYERINFO_PLAYERNAME_FIELD.name = "playerName"
PLAYERINFO_PLAYERNAME_FIELD.full_name = ".msg.PlayerInfo.playerName"
PLAYERINFO_PLAYERNAME_FIELD.number = 3
PLAYERINFO_PLAYERNAME_FIELD.index = 2
PLAYERINFO_PLAYERNAME_FIELD.label = 1
PLAYERINFO_PLAYERNAME_FIELD.has_default_value = false
PLAYERINFO_PLAYERNAME_FIELD.default_value = ""
PLAYERINFO_PLAYERNAME_FIELD.type = 9
PLAYERINFO_PLAYERNAME_FIELD.cpp_type = 9

PLAYERINFO_PLAYERCOIN_FIELD.name = "playerCoin"
PLAYERINFO_PLAYERCOIN_FIELD.full_name = ".msg.PlayerInfo.playerCoin"
PLAYERINFO_PLAYERCOIN_FIELD.number = 4
PLAYERINFO_PLAYERCOIN_FIELD.index = 3
PLAYERINFO_PLAYERCOIN_FIELD.label = 1
PLAYERINFO_PLAYERCOIN_FIELD.has_default_value = false
PLAYERINFO_PLAYERCOIN_FIELD.default_value = 0
PLAYERINFO_PLAYERCOIN_FIELD.type = 4
PLAYERINFO_PLAYERCOIN_FIELD.cpp_type = 4

PLAYERINFO.name = "PlayerInfo"
PLAYERINFO.full_name = ".msg.PlayerInfo"
PLAYERINFO.nested_types = {}
PLAYERINFO.enum_types = {}
PLAYERINFO.fields = {PLAYERINFO_PLAYERID_FIELD, PLAYERINFO_PLAYERACCOUNT_FIELD, PLAYERINFO_PLAYERNAME_FIELD, PLAYERINFO_PLAYERCOIN_FIELD}
PLAYERINFO.is_extendable = false
PLAYERINFO.extensions = {}
CS_PLAYERINFO_PLAYERID_FIELD.name = "playerId"
CS_PLAYERINFO_PLAYERID_FIELD.full_name = ".msg.CS_PlayerInfo.playerId"
CS_PLAYERINFO_PLAYERID_FIELD.number = 1
CS_PLAYERINFO_PLAYERID_FIELD.index = 0
CS_PLAYERINFO_PLAYERID_FIELD.label = 1
CS_PLAYERINFO_PLAYERID_FIELD.has_default_value = false
CS_PLAYERINFO_PLAYERID_FIELD.default_value = 0
CS_PLAYERINFO_PLAYERID_FIELD.type = 17
CS_PLAYERINFO_PLAYERID_FIELD.cpp_type = 1

CS_PLAYERINFO.name = "CS_PlayerInfo"
CS_PLAYERINFO.full_name = ".msg.CS_PlayerInfo"
CS_PLAYERINFO.nested_types = {}
CS_PLAYERINFO.enum_types = {}
CS_PLAYERINFO.fields = {CS_PLAYERINFO_PLAYERID_FIELD}
CS_PLAYERINFO.is_extendable = false
CS_PLAYERINFO.extensions = {}
SC_PLAYERINFO_PLAYERINFO_FIELD.name = "playerInfo"
SC_PLAYERINFO_PLAYERINFO_FIELD.full_name = ".msg.SC_PlayerInfo.playerInfo"
SC_PLAYERINFO_PLAYERINFO_FIELD.number = 1
SC_PLAYERINFO_PLAYERINFO_FIELD.index = 0
SC_PLAYERINFO_PLAYERINFO_FIELD.label = 1
SC_PLAYERINFO_PLAYERINFO_FIELD.has_default_value = false
SC_PLAYERINFO_PLAYERINFO_FIELD.default_value = nil
SC_PLAYERINFO_PLAYERINFO_FIELD.message_type = PLAYERINFO
SC_PLAYERINFO_PLAYERINFO_FIELD.type = 11
SC_PLAYERINFO_PLAYERINFO_FIELD.cpp_type = 10

SC_PLAYERINFO_RESULT_FIELD.name = "result"
SC_PLAYERINFO_RESULT_FIELD.full_name = ".msg.SC_PlayerInfo.result"
SC_PLAYERINFO_RESULT_FIELD.number = 2
SC_PLAYERINFO_RESULT_FIELD.index = 1
SC_PLAYERINFO_RESULT_FIELD.label = 1
SC_PLAYERINFO_RESULT_FIELD.has_default_value = false
SC_PLAYERINFO_RESULT_FIELD.default_value = 0
SC_PLAYERINFO_RESULT_FIELD.type = 17
SC_PLAYERINFO_RESULT_FIELD.cpp_type = 1

SC_PLAYERINFO.name = "SC_PlayerInfo"
SC_PLAYERINFO.full_name = ".msg.SC_PlayerInfo"
SC_PLAYERINFO.nested_types = {}
SC_PLAYERINFO.enum_types = {}
SC_PLAYERINFO.fields = {SC_PLAYERINFO_PLAYERINFO_FIELD, SC_PLAYERINFO_RESULT_FIELD}
SC_PLAYERINFO.is_extendable = false
SC_PLAYERINFO.extensions = {}
CS_HEARTBEAT.name = "CS_Heartbeat"
CS_HEARTBEAT.full_name = ".msg.CS_Heartbeat"
CS_HEARTBEAT.nested_types = {}
CS_HEARTBEAT.enum_types = {}
CS_HEARTBEAT.fields = {}
CS_HEARTBEAT.is_extendable = false
CS_HEARTBEAT.extensions = {}
SC_HEARTBEAT_SERVERTIME_FIELD.name = "serverTime"
SC_HEARTBEAT_SERVERTIME_FIELD.full_name = ".msg.SC_Heartbeat.serverTime"
SC_HEARTBEAT_SERVERTIME_FIELD.number = 1
SC_HEARTBEAT_SERVERTIME_FIELD.index = 0
SC_HEARTBEAT_SERVERTIME_FIELD.label = 1
SC_HEARTBEAT_SERVERTIME_FIELD.has_default_value = false
SC_HEARTBEAT_SERVERTIME_FIELD.default_value = 0
SC_HEARTBEAT_SERVERTIME_FIELD.type = 18
SC_HEARTBEAT_SERVERTIME_FIELD.cpp_type = 2

SC_HEARTBEAT.name = "SC_Heartbeat"
SC_HEARTBEAT.full_name = ".msg.SC_Heartbeat"
SC_HEARTBEAT.nested_types = {}
SC_HEARTBEAT.enum_types = {}
SC_HEARTBEAT.fields = {SC_HEARTBEAT_SERVERTIME_FIELD}
SC_HEARTBEAT.is_extendable = false
SC_HEARTBEAT.extensions = {}
CS_PLAYERREADY_PLAYERID_FIELD.name = "playerId"
CS_PLAYERREADY_PLAYERID_FIELD.full_name = ".msg.CS_PlayerReady.playerId"
CS_PLAYERREADY_PLAYERID_FIELD.number = 1
CS_PLAYERREADY_PLAYERID_FIELD.index = 0
CS_PLAYERREADY_PLAYERID_FIELD.label = 1
CS_PLAYERREADY_PLAYERID_FIELD.has_default_value = false
CS_PLAYERREADY_PLAYERID_FIELD.default_value = 0
CS_PLAYERREADY_PLAYERID_FIELD.type = 4
CS_PLAYERREADY_PLAYERID_FIELD.cpp_type = 4

CS_PLAYERREADY.name = "CS_PlayerReady"
CS_PLAYERREADY.full_name = ".msg.CS_PlayerReady"
CS_PLAYERREADY.nested_types = {}
CS_PLAYERREADY.enum_types = {}
CS_PLAYERREADY.fields = {CS_PLAYERREADY_PLAYERID_FIELD}
CS_PLAYERREADY.is_extendable = false
CS_PLAYERREADY.extensions = {}
SC_PLAYERREADY_RESULT_FIELD.name = "result"
SC_PLAYERREADY_RESULT_FIELD.full_name = ".msg.SC_PlayerReady.result"
SC_PLAYERREADY_RESULT_FIELD.number = 1
SC_PLAYERREADY_RESULT_FIELD.index = 0
SC_PLAYERREADY_RESULT_FIELD.label = 1
SC_PLAYERREADY_RESULT_FIELD.has_default_value = false
SC_PLAYERREADY_RESULT_FIELD.default_value = 0
SC_PLAYERREADY_RESULT_FIELD.type = 17
SC_PLAYERREADY_RESULT_FIELD.cpp_type = 1

SC_PLAYERREADY.name = "SC_PlayerReady"
SC_PLAYERREADY.full_name = ".msg.SC_PlayerReady"
SC_PLAYERREADY.nested_types = {}
SC_PLAYERREADY.enum_types = {}
SC_PLAYERREADY.fields = {SC_PLAYERREADY_RESULT_FIELD}
SC_PLAYERREADY.is_extendable = false
SC_PLAYERREADY.extensions = {}
CS_PLAYERCANCELREADY_PLAYERID_FIELD.name = "playerId"
CS_PLAYERCANCELREADY_PLAYERID_FIELD.full_name = ".msg.CS_PlayerCancelReady.playerId"
CS_PLAYERCANCELREADY_PLAYERID_FIELD.number = 1
CS_PLAYERCANCELREADY_PLAYERID_FIELD.index = 0
CS_PLAYERCANCELREADY_PLAYERID_FIELD.label = 1
CS_PLAYERCANCELREADY_PLAYERID_FIELD.has_default_value = false
CS_PLAYERCANCELREADY_PLAYERID_FIELD.default_value = 0
CS_PLAYERCANCELREADY_PLAYERID_FIELD.type = 4
CS_PLAYERCANCELREADY_PLAYERID_FIELD.cpp_type = 4

CS_PLAYERCANCELREADY.name = "CS_PlayerCancelReady"
CS_PLAYERCANCELREADY.full_name = ".msg.CS_PlayerCancelReady"
CS_PLAYERCANCELREADY.nested_types = {}
CS_PLAYERCANCELREADY.enum_types = {}
CS_PLAYERCANCELREADY.fields = {CS_PLAYERCANCELREADY_PLAYERID_FIELD}
CS_PLAYERCANCELREADY.is_extendable = false
CS_PLAYERCANCELREADY.extensions = {}
SC_PLAYERCANCELREADY_RESULT_FIELD.name = "result"
SC_PLAYERCANCELREADY_RESULT_FIELD.full_name = ".msg.SC_PlayerCancelReady.result"
SC_PLAYERCANCELREADY_RESULT_FIELD.number = 1
SC_PLAYERCANCELREADY_RESULT_FIELD.index = 0
SC_PLAYERCANCELREADY_RESULT_FIELD.label = 1
SC_PLAYERCANCELREADY_RESULT_FIELD.has_default_value = false
SC_PLAYERCANCELREADY_RESULT_FIELD.default_value = 0
SC_PLAYERCANCELREADY_RESULT_FIELD.type = 17
SC_PLAYERCANCELREADY_RESULT_FIELD.cpp_type = 1

SC_PLAYERCANCELREADY.name = "SC_PlayerCancelReady"
SC_PLAYERCANCELREADY.full_name = ".msg.SC_PlayerCancelReady"
SC_PLAYERCANCELREADY.nested_types = {}
SC_PLAYERCANCELREADY.enum_types = {}
SC_PLAYERCANCELREADY.fields = {SC_PLAYERCANCELREADY_RESULT_FIELD}
SC_PLAYERCANCELREADY.is_extendable = false
SC_PLAYERCANCELREADY.extensions = {}
SC_PLAYERCARD_CARDID_FIELD.name = "cardId"
SC_PLAYERCARD_CARDID_FIELD.full_name = ".msg.SC_PlayerCard.cardId"
SC_PLAYERCARD_CARDID_FIELD.number = 1
SC_PLAYERCARD_CARDID_FIELD.index = 0
SC_PLAYERCARD_CARDID_FIELD.label = 3
SC_PLAYERCARD_CARDID_FIELD.has_default_value = false
SC_PLAYERCARD_CARDID_FIELD.default_value = {}
SC_PLAYERCARD_CARDID_FIELD.type = 17
SC_PLAYERCARD_CARDID_FIELD.cpp_type = 1

SC_PLAYERCARD.name = "SC_PlayerCard"
SC_PLAYERCARD.full_name = ".msg.SC_PlayerCard"
SC_PLAYERCARD.nested_types = {}
SC_PLAYERCARD.enum_types = {}
SC_PLAYERCARD.fields = {SC_PLAYERCARD_CARDID_FIELD}
SC_PLAYERCARD.is_extendable = false
SC_PLAYERCARD.extensions = {}

CS_Heartbeat = protobuf.Message(CS_HEARTBEAT)
CS_PlayerCancelReady = protobuf.Message(CS_PLAYERCANCELREADY)
CS_PlayerInfo = protobuf.Message(CS_PLAYERINFO)
CS_PlayerLogin = protobuf.Message(CS_PLAYERLOGIN)
CS_PlayerReady = protobuf.Message(CS_PLAYERREADY)
CS_PlayerRegister = protobuf.Message(CS_PLAYERREGISTER)
PlayerInfo = protobuf.Message(PLAYERINFO)
SC_Heartbeat = protobuf.Message(SC_HEARTBEAT)
SC_PlayerCancelReady = protobuf.Message(SC_PLAYERCANCELREADY)
SC_PlayerCard = protobuf.Message(SC_PLAYERCARD)
SC_PlayerInfo = protobuf.Message(SC_PLAYERINFO)
SC_PlayerLogin = protobuf.Message(SC_PLAYERLOGIN)
SC_PlayerReady = protobuf.Message(SC_PLAYERREADY)
SC_PlayerRegister = protobuf.Message(SC_PLAYERREGISTER)

