# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: claim.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='claim.proto',
  package='pb',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x0b\x63laim.proto\x12\x02pb\"\x80\x01\n\x05\x43laim\x12\x1c\n\x04type\x18\x01 \x01(\x0e\x32\x0e.pb.Claim.Type\x12\x1a\n\x06stream\x18\x02 \x01(\x0b\x32\n.pb.Stream\x12\x1c\n\x07\x63hannel\x18\x03 \x01(\x0b\x32\x0b.pb.Channel\"\x1f\n\x04Type\x12\n\n\x06STREAM\x10\x00\x12\x0b\n\x07\x43HANNEL\x10\x01\"\x98\x01\n\x07\x43hannel\x12\x12\n\npublic_key\x18\x01 \x01(\x0c\x12\r\n\x05title\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x15\n\rcontact_email\x18\x04 \x01(\t\x12\x14\n\x0chomepage_url\x18\x05 \x01(\t\x12\x15\n\rthumbnail_url\x18\x10 \x01(\t\x12\x11\n\tcover_url\x18\x11 \x01(\t\"\xab\x02\n\x06Stream\x12\x0c\n\x04hash\x18\x01 \x01(\x0c\x12\x10\n\x08language\x18\x02 \x01(\t\x12\r\n\x05title\x18\x03 \x01(\t\x12\x0e\n\x06\x61uthor\x18\x04 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x05 \x01(\t\x12\x12\n\nmedia_type\x18\x06 \x01(\t\x12\x0f\n\x07license\x18\x07 \x01(\t\x12\x16\n\x04\x66ile\x18\x08 \x01(\x0b\x32\x08.pb.File\x12\x14\n\x03\x66\x65\x65\x18\x10 \x01(\x0b\x32\x07.pb.Fee\x12\x13\n\x0blicense_url\x18\x11 \x01(\t\x12\x15\n\rthumbnail_url\x18\x12 \x01(\t\x12\x10\n\x08\x64uration\x18\x13 \x01(\r\x12\x0c\n\x04tags\x18\x14 \x03(\t\x12\x14\n\x0crelease_time\x18\x15 \x01(\x03\x12\x18\n\x05video\x18\x16 \x01(\x0b\x32\t.pb.Video\"h\n\x03\x46\x65\x65\x12\"\n\x08\x63urrency\x18\x01 \x01(\x0e\x32\x10.pb.Fee.Currency\x12\x0f\n\x07\x61\x64\x64ress\x18\x02 \x01(\x0c\x12\x0e\n\x06\x61mount\x18\x03 \x01(\x04\"\x1c\n\x08\x43urrency\x12\x07\n\x03LBC\x10\x00\x12\x07\n\x03USD\x10\x01\"\"\n\x04\x46ile\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x0c\n\x04size\x18\x02 \x01(\x04\"&\n\x05Video\x12\r\n\x05width\x18\x01 \x01(\r\x12\x0e\n\x06height\x18\x02 \x01(\rb\x06proto3')
)



_CLAIM_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='pb.Claim.Type',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='STREAM', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CHANNEL', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=117,
  serialized_end=148,
)
_sym_db.RegisterEnumDescriptor(_CLAIM_TYPE)

_FEE_CURRENCY = _descriptor.EnumDescriptor(
  name='Currency',
  full_name='pb.Fee.Currency',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='LBC', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='USD', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=683,
  serialized_end=711,
)
_sym_db.RegisterEnumDescriptor(_FEE_CURRENCY)


_CLAIM = _descriptor.Descriptor(
  name='Claim',
  full_name='pb.Claim',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='pb.Claim.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='stream', full_name='pb.Claim.stream', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='channel', full_name='pb.Claim.channel', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _CLAIM_TYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=20,
  serialized_end=148,
)


_CHANNEL = _descriptor.Descriptor(
  name='Channel',
  full_name='pb.Channel',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='public_key', full_name='pb.Channel.public_key', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='title', full_name='pb.Channel.title', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='pb.Channel.description', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='contact_email', full_name='pb.Channel.contact_email', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='homepage_url', full_name='pb.Channel.homepage_url', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='thumbnail_url', full_name='pb.Channel.thumbnail_url', index=5,
      number=16, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='cover_url', full_name='pb.Channel.cover_url', index=6,
      number=17, type=9, cpp_type=9, label=1,
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
  serialized_start=151,
  serialized_end=303,
)


_STREAM = _descriptor.Descriptor(
  name='Stream',
  full_name='pb.Stream',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='hash', full_name='pb.Stream.hash', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='language', full_name='pb.Stream.language', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='title', full_name='pb.Stream.title', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='author', full_name='pb.Stream.author', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='pb.Stream.description', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='media_type', full_name='pb.Stream.media_type', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='license', full_name='pb.Stream.license', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='file', full_name='pb.Stream.file', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='fee', full_name='pb.Stream.fee', index=8,
      number=16, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='license_url', full_name='pb.Stream.license_url', index=9,
      number=17, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='thumbnail_url', full_name='pb.Stream.thumbnail_url', index=10,
      number=18, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='duration', full_name='pb.Stream.duration', index=11,
      number=19, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='tags', full_name='pb.Stream.tags', index=12,
      number=20, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='release_time', full_name='pb.Stream.release_time', index=13,
      number=21, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='video', full_name='pb.Stream.video', index=14,
      number=22, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=306,
  serialized_end=605,
)


_FEE = _descriptor.Descriptor(
  name='Fee',
  full_name='pb.Fee',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='currency', full_name='pb.Fee.currency', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='address', full_name='pb.Fee.address', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='amount', full_name='pb.Fee.amount', index=2,
      number=3, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _FEE_CURRENCY,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=607,
  serialized_end=711,
)


_FILE = _descriptor.Descriptor(
  name='File',
  full_name='pb.File',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='pb.File.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='size', full_name='pb.File.size', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=713,
  serialized_end=747,
)


_VIDEO = _descriptor.Descriptor(
  name='Video',
  full_name='pb.Video',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='width', full_name='pb.Video.width', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='height', full_name='pb.Video.height', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=749,
  serialized_end=787,
)

_CLAIM.fields_by_name['type'].enum_type = _CLAIM_TYPE
_CLAIM.fields_by_name['stream'].message_type = _STREAM
_CLAIM.fields_by_name['channel'].message_type = _CHANNEL
_CLAIM_TYPE.containing_type = _CLAIM
_STREAM.fields_by_name['file'].message_type = _FILE
_STREAM.fields_by_name['fee'].message_type = _FEE
_STREAM.fields_by_name['video'].message_type = _VIDEO
_FEE.fields_by_name['currency'].enum_type = _FEE_CURRENCY
_FEE_CURRENCY.containing_type = _FEE
DESCRIPTOR.message_types_by_name['Claim'] = _CLAIM
DESCRIPTOR.message_types_by_name['Channel'] = _CHANNEL
DESCRIPTOR.message_types_by_name['Stream'] = _STREAM
DESCRIPTOR.message_types_by_name['Fee'] = _FEE
DESCRIPTOR.message_types_by_name['File'] = _FILE
DESCRIPTOR.message_types_by_name['Video'] = _VIDEO
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Claim = _reflection.GeneratedProtocolMessageType('Claim', (_message.Message,), dict(
  DESCRIPTOR = _CLAIM,
  __module__ = 'claim_pb2'
  # @@protoc_insertion_point(class_scope:pb.Claim)
  ))
_sym_db.RegisterMessage(Claim)

Channel = _reflection.GeneratedProtocolMessageType('Channel', (_message.Message,), dict(
  DESCRIPTOR = _CHANNEL,
  __module__ = 'claim_pb2'
  # @@protoc_insertion_point(class_scope:pb.Channel)
  ))
_sym_db.RegisterMessage(Channel)

Stream = _reflection.GeneratedProtocolMessageType('Stream', (_message.Message,), dict(
  DESCRIPTOR = _STREAM,
  __module__ = 'claim_pb2'
  # @@protoc_insertion_point(class_scope:pb.Stream)
  ))
_sym_db.RegisterMessage(Stream)

Fee = _reflection.GeneratedProtocolMessageType('Fee', (_message.Message,), dict(
  DESCRIPTOR = _FEE,
  __module__ = 'claim_pb2'
  # @@protoc_insertion_point(class_scope:pb.Fee)
  ))
_sym_db.RegisterMessage(Fee)

File = _reflection.GeneratedProtocolMessageType('File', (_message.Message,), dict(
  DESCRIPTOR = _FILE,
  __module__ = 'claim_pb2'
  # @@protoc_insertion_point(class_scope:pb.File)
  ))
_sym_db.RegisterMessage(File)

Video = _reflection.GeneratedProtocolMessageType('Video', (_message.Message,), dict(
  DESCRIPTOR = _VIDEO,
  __module__ = 'claim_pb2'
  # @@protoc_insertion_point(class_scope:pb.Video)
  ))
_sym_db.RegisterMessage(Video)


# @@protoc_insertion_point(module_scope)
