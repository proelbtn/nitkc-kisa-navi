# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: food/messages.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='food/messages.proto',
  package='food',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x13\x66ood/messages.proto\x12\x04\x66ood\"\x1b\n\x0b\x46oodRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"\x1c\n\x0c\x46oodResponse\x12\x0c\n\x04name\x18\x01 \x01(\tb\x06proto3')
)




_FOODREQUEST = _descriptor.Descriptor(
  name='FoodRequest',
  full_name='food.FoodRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='food.FoodRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=29,
  serialized_end=56,
)


_FOODRESPONSE = _descriptor.Descriptor(
  name='FoodResponse',
  full_name='food.FoodResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='food.FoodResponse.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=58,
  serialized_end=86,
)

DESCRIPTOR.message_types_by_name['FoodRequest'] = _FOODREQUEST
DESCRIPTOR.message_types_by_name['FoodResponse'] = _FOODRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

FoodRequest = _reflection.GeneratedProtocolMessageType('FoodRequest', (_message.Message,), dict(
  DESCRIPTOR = _FOODREQUEST,
  __module__ = 'food.messages_pb2'
  # @@protoc_insertion_point(class_scope:food.FoodRequest)
  ))
_sym_db.RegisterMessage(FoodRequest)

FoodResponse = _reflection.GeneratedProtocolMessageType('FoodResponse', (_message.Message,), dict(
  DESCRIPTOR = _FOODRESPONSE,
  __module__ = 'food.messages_pb2'
  # @@protoc_insertion_point(class_scope:food.FoodResponse)
  ))
_sym_db.RegisterMessage(FoodResponse)


# @@protoc_insertion_point(module_scope)
