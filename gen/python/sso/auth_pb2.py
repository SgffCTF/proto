# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: sso/auth.proto
# Protobuf Python Version: 5.28.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    28,
    1,
    '',
    'sso/auth.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0esso/auth.proto\x12\x04\x61uth\"C\n\x0fRegisterRequest\x12\x10\n\x08username\x18\x01 \x01(\t\x12\x0c\n\x04role\x18\x02 \x01(\t\x12\x10\n\x08password\x18\x03 \x01(\t\"\x1e\n\x10RegisterResponse\x12\n\n\x02id\x18\x01 \x01(\x03\"2\n\x0cLoginRequest\x12\x10\n\x08username\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\"\x1e\n\rLoginResponse\x12\r\n\x05token\x18\x01 \x01(\t2s\n\x04\x41uth\x12\x39\n\x08Register\x12\x15.auth.RegisterRequest\x1a\x16.auth.RegisterResponse\x12\x30\n\x05Login\x12\x12.auth.LoginRequest\x1a\x13.auth.LoginResponseB\x1aZ\x18\x63tf-service.sso.v1;ssov1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'sso.auth_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\030ctf-service.sso.v1;ssov1'
  _globals['_REGISTERREQUEST']._serialized_start=24
  _globals['_REGISTERREQUEST']._serialized_end=91
  _globals['_REGISTERRESPONSE']._serialized_start=93
  _globals['_REGISTERRESPONSE']._serialized_end=123
  _globals['_LOGINREQUEST']._serialized_start=125
  _globals['_LOGINREQUEST']._serialized_end=175
  _globals['_LOGINRESPONSE']._serialized_start=177
  _globals['_LOGINRESPONSE']._serialized_end=207
  _globals['_AUTH']._serialized_start=209
  _globals['_AUTH']._serialized_end=324
# @@protoc_insertion_point(module_scope)
