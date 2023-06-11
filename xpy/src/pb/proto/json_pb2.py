# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/json.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import descriptor_pb2 as google_dot_protobuf_dot_descriptor__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x10proto/json.proto\x12\x04json\x1a google/protobuf/descriptor.proto\"A\n\x0eMessageOptions\x12\x0e\n\x06ignore\x18\x01 \x01(\x08\x12\x1f\n\x17\x64isallow_unknown_fields\x18\x02 \x01(\x08\"[\n\x0cOneofOptions\x12\x11\n\x04json\x18\x01 \x01(\tH\x00\x88\x01\x01\x12\x0e\n\x06ignore\x18\x02 \x01(\x08\x12\x11\n\tomitempty\x18\x03 \x01(\x08\x12\x0c\n\x04hide\x18\x04 \x01(\x08\x42\x07\n\x05_json\"\xfd\x01\n\x0c\x46ieldOptions\x12\x11\n\x04json\x18\x01 \x01(\tH\x01\x88\x01\x01\x12\x0e\n\x06ignore\x18\x02 \x01(\x08\x12\x11\n\tomitempty\x18\x03 \x01(\x08\x12\x17\n\x0fuse_enum_string\x18\x04 \x01(\x08\x12\x1e\n\x04\x65num\x18\x0b \x01(\x0b\x32\x0e.json.TypeEnumH\x00\x12\x1c\n\x03\x61ny\x18\x0c \x01(\x0b\x32\r.json.TypeAnyH\x00\x12&\n\x08\x64uration\x18\r \x01(\x0b\x32\x12.json.TypeDurationH\x00\x12(\n\ttimestamp\x18\x0e \x01(\x0b\x32\x13.json.TypeTimestampH\x00\x42\x05\n\x03WKTB\x07\n\x05_json\"@\n\x08TypeEnum\x12\x14\n\nuse_number\x18\x01 \x01(\x08H\x00\x12\x14\n\nuse_string\x18\x02 \x01(\x08H\x00\x42\x08\n\x06\x46ormat\"7\n\x07TypeAny\x12\x10\n\x06native\x18\x01 \x01(\x08H\x00\x12\x10\n\x06\x65xpand\x18\x02 \x01(\x08H\x00\x42\x08\n\x06\x46ormat\"\xba\x01\n\x0cTypeDuration\x12\x10\n\x06native\x18\x01 \x01(\x08H\x00\x12\x10\n\x06string\x18\x02 \x01(\x08H\x00\x12\x15\n\x0bnanoseconds\x18\x03 \x01(\x08H\x00\x12\x16\n\x0cmicroseconds\x18\x04 \x01(\x08H\x00\x12\x16\n\x0cmilliseconds\x18\x05 \x01(\x08H\x00\x12\x11\n\x07seconds\x18\x06 \x01(\x08H\x00\x12\x11\n\x07minutes\x18\x07 \x01(\x08H\x00\x12\x0f\n\x05hours\x18\x08 \x01(\x08H\x00\x42\x08\n\x06\x46ormat\"\x81\x02\n\rTypeTimestamp\x12\x10\n\x06native\x18\x01 \x01(\x08H\x00\x12\x35\n\x0btime_layout\x18\x02 \x01(\x0b\x32\x1e.json.TypeTimestamp.TimeLayoutH\x00\x12\x13\n\tunix_nano\x18\x03 \x01(\x08H\x00\x12\x14\n\nunix_micro\x18\x04 \x01(\x08H\x00\x12\x14\n\nunix_milli\x18\x05 \x01(\x08H\x00\x12\x12\n\x08unix_sec\x18\x06 \x01(\x08H\x00\x1aH\n\nTimeLayout\x12\x0e\n\x06golang\x18\x01 \x01(\t\x12\x0c\n\x04java\x18\x02 \x01(\t\x12\x0c\n\x04rust\x18\x03 \x01(\t\x12\x0e\n\x06python\x18\x04 \x01(\tB\x08\n\x06\x46ormat:H\n\x07message\x12\x1f.google.protobuf.MessageOptions\x18\xa0\xf4\x03 \x01(\x0b\x32\x14.json.MessageOptions:B\n\x05\x66ield\x12\x1d.google.protobuf.FieldOptions\x18\xa1\xf4\x03 \x01(\x0b\x32\x12.json.FieldOptions:B\n\x05oneof\x12\x1d.google.protobuf.OneofOptions\x18\xa2\xf4\x03 \x01(\x0b\x32\x12.json.OneofOptionsB]\n\x1fio.github.yu31.protoc.pb.pbjsonB\x06PBJSONP\x01Z0github.com/yu31/protoc-plugin-json/xgo/pb/pbjsonb\x06proto3')


MESSAGE_FIELD_NUMBER = 64032
message = DESCRIPTOR.extensions_by_name['message']
FIELD_FIELD_NUMBER = 64033
field = DESCRIPTOR.extensions_by_name['field']
ONEOF_FIELD_NUMBER = 64034
oneof = DESCRIPTOR.extensions_by_name['oneof']

_MESSAGEOPTIONS = DESCRIPTOR.message_types_by_name['MessageOptions']
_ONEOFOPTIONS = DESCRIPTOR.message_types_by_name['OneofOptions']
_FIELDOPTIONS = DESCRIPTOR.message_types_by_name['FieldOptions']
_TYPEENUM = DESCRIPTOR.message_types_by_name['TypeEnum']
_TYPEANY = DESCRIPTOR.message_types_by_name['TypeAny']
_TYPEDURATION = DESCRIPTOR.message_types_by_name['TypeDuration']
_TYPETIMESTAMP = DESCRIPTOR.message_types_by_name['TypeTimestamp']
_TYPETIMESTAMP_TIMELAYOUT = _TYPETIMESTAMP.nested_types_by_name['TimeLayout']
MessageOptions = _reflection.GeneratedProtocolMessageType('MessageOptions', (_message.Message,), {
  'DESCRIPTOR' : _MESSAGEOPTIONS,
  '__module__' : 'proto.json_pb2'
  # @@protoc_insertion_point(class_scope:json.MessageOptions)
  })
_sym_db.RegisterMessage(MessageOptions)

OneofOptions = _reflection.GeneratedProtocolMessageType('OneofOptions', (_message.Message,), {
  'DESCRIPTOR' : _ONEOFOPTIONS,
  '__module__' : 'proto.json_pb2'
  # @@protoc_insertion_point(class_scope:json.OneofOptions)
  })
_sym_db.RegisterMessage(OneofOptions)

FieldOptions = _reflection.GeneratedProtocolMessageType('FieldOptions', (_message.Message,), {
  'DESCRIPTOR' : _FIELDOPTIONS,
  '__module__' : 'proto.json_pb2'
  # @@protoc_insertion_point(class_scope:json.FieldOptions)
  })
_sym_db.RegisterMessage(FieldOptions)

TypeEnum = _reflection.GeneratedProtocolMessageType('TypeEnum', (_message.Message,), {
  'DESCRIPTOR' : _TYPEENUM,
  '__module__' : 'proto.json_pb2'
  # @@protoc_insertion_point(class_scope:json.TypeEnum)
  })
_sym_db.RegisterMessage(TypeEnum)

TypeAny = _reflection.GeneratedProtocolMessageType('TypeAny', (_message.Message,), {
  'DESCRIPTOR' : _TYPEANY,
  '__module__' : 'proto.json_pb2'
  # @@protoc_insertion_point(class_scope:json.TypeAny)
  })
_sym_db.RegisterMessage(TypeAny)

TypeDuration = _reflection.GeneratedProtocolMessageType('TypeDuration', (_message.Message,), {
  'DESCRIPTOR' : _TYPEDURATION,
  '__module__' : 'proto.json_pb2'
  # @@protoc_insertion_point(class_scope:json.TypeDuration)
  })
_sym_db.RegisterMessage(TypeDuration)

TypeTimestamp = _reflection.GeneratedProtocolMessageType('TypeTimestamp', (_message.Message,), {

  'TimeLayout' : _reflection.GeneratedProtocolMessageType('TimeLayout', (_message.Message,), {
    'DESCRIPTOR' : _TYPETIMESTAMP_TIMELAYOUT,
    '__module__' : 'proto.json_pb2'
    # @@protoc_insertion_point(class_scope:json.TypeTimestamp.TimeLayout)
    })
  ,
  'DESCRIPTOR' : _TYPETIMESTAMP,
  '__module__' : 'proto.json_pb2'
  # @@protoc_insertion_point(class_scope:json.TypeTimestamp)
  })
_sym_db.RegisterMessage(TypeTimestamp)
_sym_db.RegisterMessage(TypeTimestamp.TimeLayout)

if _descriptor._USE_C_DESCRIPTORS == False:
  google_dot_protobuf_dot_descriptor__pb2.MessageOptions.RegisterExtension(message)
  google_dot_protobuf_dot_descriptor__pb2.FieldOptions.RegisterExtension(field)
  google_dot_protobuf_dot_descriptor__pb2.OneofOptions.RegisterExtension(oneof)

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\037io.github.yu31.protoc.pb.pbjsonB\006PBJSONP\001Z0github.com/yu31/protoc-plugin-json/xgo/pb/pbjson'
  _MESSAGEOPTIONS._serialized_start=60
  _MESSAGEOPTIONS._serialized_end=125
  _ONEOFOPTIONS._serialized_start=127
  _ONEOFOPTIONS._serialized_end=218
  _FIELDOPTIONS._serialized_start=221
  _FIELDOPTIONS._serialized_end=474
  _TYPEENUM._serialized_start=476
  _TYPEENUM._serialized_end=540
  _TYPEANY._serialized_start=542
  _TYPEANY._serialized_end=597
  _TYPEDURATION._serialized_start=600
  _TYPEDURATION._serialized_end=786
  _TYPETIMESTAMP._serialized_start=789
  _TYPETIMESTAMP._serialized_end=1046
  _TYPETIMESTAMP_TIMELAYOUT._serialized_start=964
  _TYPETIMESTAMP_TIMELAYOUT._serialized_end=1036
# @@protoc_insertion_point(module_scope)
