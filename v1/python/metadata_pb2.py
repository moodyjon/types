# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: metadata.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import fee_pb2 as fee__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0emetadata.proto\x12\tlegacy_pb\x1a\tfee.proto\"\xfc\x0e\n\x08Metadata\x12,\n\x07version\x18\x01 \x02(\x0e\x32\x1b.legacy_pb.Metadata.Version\x12.\n\x08language\x18\x02 \x02(\x0e\x32\x1c.legacy_pb.Metadata.Language\x12\r\n\x05title\x18\x03 \x02(\t\x12\x13\n\x0b\x64\x65scription\x18\x04 \x02(\t\x12\x0e\n\x06\x61uthor\x18\x05 \x02(\t\x12\x0f\n\x07license\x18\x06 \x02(\t\x12\x0c\n\x04nsfw\x18\x07 \x02(\x08\x12\x1b\n\x03\x66\x65\x65\x18\x08 \x01(\x0b\x32\x0e.legacy_pb.Fee\x12\x11\n\tthumbnail\x18\t \x01(\t\x12\x0f\n\x07preview\x18\n \x01(\t\x12\x12\n\nlicenseUrl\x18\x0b \x01(\t\"N\n\x07Version\x12\x13\n\x0fUNKNOWN_VERSION\x10\x00\x12\n\n\x06_0_0_1\x10\x01\x12\n\n\x06_0_0_2\x10\x02\x12\n\n\x06_0_0_3\x10\x03\x12\n\n\x06_0_1_0\x10\x04\"\x99\x0c\n\x08Language\x12\x14\n\x10UNKNOWN_LANGUAGE\x10\x00\x12\x06\n\x02\x65n\x10\x01\x12\x06\n\x02\x61\x61\x10\x02\x12\x06\n\x02\x61\x62\x10\x03\x12\x06\n\x02\x61\x65\x10\x04\x12\x06\n\x02\x61\x66\x10\x05\x12\x06\n\x02\x61k\x10\x06\x12\x06\n\x02\x61m\x10\x07\x12\x06\n\x02\x61n\x10\x08\x12\x06\n\x02\x61r\x10\t\x12\x06\n\x02\x61s\x10\n\x12\x06\n\x02\x61v\x10\x0b\x12\x06\n\x02\x61y\x10\x0c\x12\x06\n\x02\x61z\x10\r\x12\x06\n\x02\x62\x61\x10\x0e\x12\x06\n\x02\x62\x65\x10\x0f\x12\x06\n\x02\x62g\x10\x10\x12\x06\n\x02\x62h\x10\x11\x12\x06\n\x02\x62i\x10\x12\x12\x06\n\x02\x62m\x10\x13\x12\x06\n\x02\x62n\x10\x14\x12\x06\n\x02\x62o\x10\x15\x12\x06\n\x02\x62r\x10\x16\x12\x06\n\x02\x62s\x10\x17\x12\x06\n\x02\x63\x61\x10\x18\x12\x06\n\x02\x63\x65\x10\x19\x12\x06\n\x02\x63h\x10\x1a\x12\x06\n\x02\x63o\x10\x1b\x12\x06\n\x02\x63r\x10\x1c\x12\x06\n\x02\x63s\x10\x1d\x12\x06\n\x02\x63u\x10\x1e\x12\x06\n\x02\x63v\x10\x1f\x12\x06\n\x02\x63y\x10 \x12\x06\n\x02\x64\x61\x10!\x12\x06\n\x02\x64\x65\x10\"\x12\x06\n\x02\x64v\x10#\x12\x06\n\x02\x64z\x10$\x12\x06\n\x02\x65\x65\x10%\x12\x06\n\x02\x65l\x10&\x12\x06\n\x02\x65o\x10\'\x12\x06\n\x02\x65s\x10(\x12\x06\n\x02\x65t\x10)\x12\x06\n\x02\x65u\x10*\x12\x06\n\x02\x66\x61\x10+\x12\x06\n\x02\x66\x66\x10,\x12\x06\n\x02\x66i\x10-\x12\x06\n\x02\x66j\x10.\x12\x06\n\x02\x66o\x10/\x12\x06\n\x02\x66r\x10\x30\x12\x06\n\x02\x66y\x10\x31\x12\x06\n\x02ga\x10\x32\x12\x06\n\x02gd\x10\x33\x12\x06\n\x02gl\x10\x34\x12\x06\n\x02gn\x10\x35\x12\x06\n\x02gu\x10\x36\x12\x06\n\x02gv\x10\x37\x12\x06\n\x02ha\x10\x38\x12\x06\n\x02he\x10\x39\x12\x06\n\x02hi\x10:\x12\x06\n\x02ho\x10;\x12\x06\n\x02hr\x10<\x12\x06\n\x02ht\x10=\x12\x06\n\x02hu\x10>\x12\x06\n\x02hy\x10?\x12\x06\n\x02hz\x10@\x12\x06\n\x02ia\x10\x41\x12\x06\n\x02id\x10\x42\x12\x06\n\x02ie\x10\x43\x12\x06\n\x02ig\x10\x44\x12\x06\n\x02ii\x10\x45\x12\x06\n\x02ik\x10\x46\x12\x06\n\x02io\x10G\x12\x06\n\x02is\x10H\x12\x06\n\x02it\x10I\x12\x06\n\x02iu\x10J\x12\x06\n\x02ja\x10K\x12\x06\n\x02jv\x10L\x12\x06\n\x02ka\x10M\x12\x06\n\x02kg\x10N\x12\x06\n\x02ki\x10O\x12\x06\n\x02kj\x10P\x12\x06\n\x02kk\x10Q\x12\x06\n\x02kl\x10R\x12\x06\n\x02km\x10S\x12\x06\n\x02kn\x10T\x12\x06\n\x02ko\x10U\x12\x06\n\x02kr\x10V\x12\x06\n\x02ks\x10W\x12\x06\n\x02ku\x10X\x12\x06\n\x02kv\x10Y\x12\x06\n\x02kw\x10Z\x12\x06\n\x02ky\x10[\x12\x06\n\x02la\x10\\\x12\x06\n\x02lb\x10]\x12\x06\n\x02lg\x10^\x12\x06\n\x02li\x10_\x12\x06\n\x02ln\x10`\x12\x06\n\x02lo\x10\x61\x12\x06\n\x02lt\x10\x62\x12\x06\n\x02lu\x10\x63\x12\x06\n\x02lv\x10\x64\x12\x06\n\x02mg\x10\x65\x12\x06\n\x02mh\x10\x66\x12\x06\n\x02mi\x10g\x12\x06\n\x02mk\x10h\x12\x06\n\x02ml\x10i\x12\x06\n\x02mn\x10j\x12\x06\n\x02mr\x10k\x12\x06\n\x02ms\x10l\x12\x06\n\x02mt\x10m\x12\x06\n\x02my\x10n\x12\x06\n\x02na\x10o\x12\x06\n\x02nb\x10p\x12\x06\n\x02nd\x10q\x12\x06\n\x02ne\x10r\x12\x06\n\x02ng\x10s\x12\x06\n\x02nl\x10t\x12\x06\n\x02nn\x10u\x12\x06\n\x02no\x10v\x12\x06\n\x02nr\x10w\x12\x06\n\x02nv\x10x\x12\x06\n\x02ny\x10y\x12\x06\n\x02oc\x10z\x12\x06\n\x02oj\x10{\x12\x06\n\x02om\x10|\x12\x06\n\x02or\x10}\x12\x06\n\x02os\x10~\x12\x06\n\x02pa\x10\x7f\x12\x07\n\x02pi\x10\x80\x01\x12\x07\n\x02pl\x10\x81\x01\x12\x07\n\x02ps\x10\x82\x01\x12\x07\n\x02pt\x10\x83\x01\x12\x07\n\x02qu\x10\x84\x01\x12\x07\n\x02rm\x10\x85\x01\x12\x07\n\x02rn\x10\x86\x01\x12\x07\n\x02ro\x10\x87\x01\x12\x07\n\x02ru\x10\x88\x01\x12\x07\n\x02rw\x10\x89\x01\x12\x07\n\x02sa\x10\x8a\x01\x12\x07\n\x02sc\x10\x8b\x01\x12\x07\n\x02sd\x10\x8c\x01\x12\x07\n\x02se\x10\x8d\x01\x12\x07\n\x02sg\x10\x8e\x01\x12\x07\n\x02si\x10\x8f\x01\x12\x07\n\x02sk\x10\x90\x01\x12\x07\n\x02sl\x10\x91\x01\x12\x07\n\x02sm\x10\x92\x01\x12\x07\n\x02sn\x10\x93\x01\x12\x07\n\x02so\x10\x94\x01\x12\x07\n\x02sq\x10\x95\x01\x12\x07\n\x02sr\x10\x96\x01\x12\x07\n\x02ss\x10\x97\x01\x12\x07\n\x02st\x10\x98\x01\x12\x07\n\x02su\x10\x99\x01\x12\x07\n\x02sv\x10\x9a\x01\x12\x07\n\x02sw\x10\x9b\x01\x12\x07\n\x02ta\x10\x9c\x01\x12\x07\n\x02te\x10\x9d\x01\x12\x07\n\x02tg\x10\x9e\x01\x12\x07\n\x02th\x10\x9f\x01\x12\x07\n\x02ti\x10\xa0\x01\x12\x07\n\x02tk\x10\xa1\x01\x12\x07\n\x02tl\x10\xa2\x01\x12\x07\n\x02tn\x10\xa3\x01\x12\x07\n\x02to\x10\xa4\x01\x12\x07\n\x02tr\x10\xa5\x01\x12\x07\n\x02ts\x10\xa6\x01\x12\x07\n\x02tt\x10\xa7\x01\x12\x07\n\x02tw\x10\xa8\x01\x12\x07\n\x02ty\x10\xa9\x01\x12\x07\n\x02ug\x10\xaa\x01\x12\x07\n\x02uk\x10\xab\x01\x12\x07\n\x02ur\x10\xac\x01\x12\x07\n\x02uz\x10\xad\x01\x12\x07\n\x02ve\x10\xae\x01\x12\x07\n\x02vi\x10\xaf\x01\x12\x07\n\x02vo\x10\xb0\x01\x12\x07\n\x02wa\x10\xb1\x01\x12\x07\n\x02wo\x10\xb2\x01\x12\x07\n\x02xh\x10\xb3\x01\x12\x07\n\x02yi\x10\xb4\x01\x12\x07\n\x02yo\x10\xb5\x01\x12\x07\n\x02za\x10\xb6\x01\x12\x07\n\x02zh\x10\xb7\x01\x12\x07\n\x02zu\x10\xb8\x01')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'metadata_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  _METADATA._serialized_start=41
  _METADATA._serialized_end=1957
  _METADATA_VERSION._serialized_start=315
  _METADATA_VERSION._serialized_end=393
  _METADATA_LANGUAGE._serialized_start=396
  _METADATA_LANGUAGE._serialized_end=1957
# @@protoc_insertion_point(module_scope)
