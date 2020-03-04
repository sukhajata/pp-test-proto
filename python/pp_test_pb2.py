# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: pp-test.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='pp-test.proto',
  package='test',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\rpp-test.proto\x12\x04test\"\xd0\x01\n\x0fReceivedMessage\x12,\n\x06source\x18\x01 \x01(\x0e\x32\x1c.test.ReceivedMessage.Source\x12\x0c\n\x04type\x18\x02 \x01(\t\x12\x11\n\tdeviceEUI\x18\x03 \x01(\t\x12\x11\n\ttimestamp\x18\x04 \x01(\t\x12\r\n\x05phase\x18\x05 \x01(\x05\x12\x0f\n\x07\x63ontent\x18\x06 \x01(\t\";\n\x06Source\x12\r\n\tNO_SOURCE\x10\x00\x12\x08\n\x04MQTT\x10\x01\x12\t\n\x05KAFKA\x10\x02\x12\r\n\tTIMESCALE\x10\x03\"\x19\n\x08Response\x12\r\n\x05reply\x18\x01 \x01(\t2K\n\x0bTestService\x12<\n\x11OnMessageReceived\x12\x15.test.ReceivedMessage\x1a\x0e.test.Response\"\x00\x62\x06proto3')
)



_RECEIVEDMESSAGE_SOURCE = _descriptor.EnumDescriptor(
  name='Source',
  full_name='test.ReceivedMessage.Source',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NO_SOURCE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MQTT', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='KAFKA', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TIMESCALE', index=3, number=3,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=173,
  serialized_end=232,
)
_sym_db.RegisterEnumDescriptor(_RECEIVEDMESSAGE_SOURCE)


_RECEIVEDMESSAGE = _descriptor.Descriptor(
  name='ReceivedMessage',
  full_name='test.ReceivedMessage',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='source', full_name='test.ReceivedMessage.source', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='test.ReceivedMessage.type', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='deviceEUI', full_name='test.ReceivedMessage.deviceEUI', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='test.ReceivedMessage.timestamp', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='phase', full_name='test.ReceivedMessage.phase', index=4,
      number=5, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='content', full_name='test.ReceivedMessage.content', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _RECEIVEDMESSAGE_SOURCE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=24,
  serialized_end=232,
)


_RESPONSE = _descriptor.Descriptor(
  name='Response',
  full_name='test.Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='reply', full_name='test.Response.reply', index=0,
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
  serialized_start=234,
  serialized_end=259,
)

_RECEIVEDMESSAGE.fields_by_name['source'].enum_type = _RECEIVEDMESSAGE_SOURCE
_RECEIVEDMESSAGE_SOURCE.containing_type = _RECEIVEDMESSAGE
DESCRIPTOR.message_types_by_name['ReceivedMessage'] = _RECEIVEDMESSAGE
DESCRIPTOR.message_types_by_name['Response'] = _RESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ReceivedMessage = _reflection.GeneratedProtocolMessageType('ReceivedMessage', (_message.Message,), {
  'DESCRIPTOR' : _RECEIVEDMESSAGE,
  '__module__' : 'pp_test_pb2'
  # @@protoc_insertion_point(class_scope:test.ReceivedMessage)
  })
_sym_db.RegisterMessage(ReceivedMessage)

Response = _reflection.GeneratedProtocolMessageType('Response', (_message.Message,), {
  'DESCRIPTOR' : _RESPONSE,
  '__module__' : 'pp_test_pb2'
  # @@protoc_insertion_point(class_scope:test.Response)
  })
_sym_db.RegisterMessage(Response)



_TESTSERVICE = _descriptor.ServiceDescriptor(
  name='TestService',
  full_name='test.TestService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=261,
  serialized_end=336,
  methods=[
  _descriptor.MethodDescriptor(
    name='OnMessageReceived',
    full_name='test.TestService.OnMessageReceived',
    index=0,
    containing_service=None,
    input_type=_RECEIVEDMESSAGE,
    output_type=_RESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_TESTSERVICE)

DESCRIPTOR.services_by_name['TestService'] = _TESTSERVICE

# @@protoc_insertion_point(module_scope)
