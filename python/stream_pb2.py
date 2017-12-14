# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: stream.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import metadata_pb2 as metadata__pb2
import source_pb2 as source__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='stream.proto',
  package='pb',
  syntax='proto2',
  serialized_pb=_b('\n\x0cstream.proto\x12\x02pb\x1a\x0emetadata.proto\x1a\x0csource.proto\"\x95\x01\n\x06Stream\x12#\n\x07version\x18\x01 \x02(\x0e\x32\x12.pb.Stream.Version\x12\x1e\n\x08metadata\x18\x02 \x02(\x0b\x32\x0c.pb.Metadata\x12\x1a\n\x06source\x18\x03 \x02(\x0b\x32\n.pb.Source\"*\n\x07Version\x12\x13\n\x0fUNKNOWN_VERSION\x10\x00\x12\n\n\x06_0_0_1\x10\x01')
  ,
  dependencies=[metadata__pb2.DESCRIPTOR,source__pb2.DESCRIPTOR,])



_STREAM_VERSION = _descriptor.EnumDescriptor(
  name='Version',
  full_name='pb.Stream.Version',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN_VERSION', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='_0_0_1', index=1, number=1,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=158,
  serialized_end=200,
)
_sym_db.RegisterEnumDescriptor(_STREAM_VERSION)


_STREAM = _descriptor.Descriptor(
  name='Stream',
  full_name='pb.Stream',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='version', full_name='pb.Stream.version', index=0,
      number=1, type=14, cpp_type=8, label=2,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='metadata', full_name='pb.Stream.metadata', index=1,
      number=2, type=11, cpp_type=10, label=2,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='source', full_name='pb.Stream.source', index=2,
      number=3, type=11, cpp_type=10, label=2,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _STREAM_VERSION,
  ],
  options=None,
  is_extendable=False,
  syntax='proto2',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=51,
  serialized_end=200,
)

_STREAM.fields_by_name['version'].enum_type = _STREAM_VERSION
_STREAM.fields_by_name['metadata'].message_type = metadata__pb2._METADATA
_STREAM.fields_by_name['source'].message_type = source__pb2._SOURCE
_STREAM_VERSION.containing_type = _STREAM
DESCRIPTOR.message_types_by_name['Stream'] = _STREAM
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Stream = _reflection.GeneratedProtocolMessageType('Stream', (_message.Message,), dict(
  DESCRIPTOR = _STREAM,
  __module__ = 'stream_pb2'
  # @@protoc_insertion_point(class_scope:pb.Stream)
  ))
_sym_db.RegisterMessage(Stream)


# @@protoc_insertion_point(module_scope)
