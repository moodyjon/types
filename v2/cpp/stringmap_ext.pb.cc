// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: stringmap_ext.proto

#include "stringmap_ext.pb.h"

#include <algorithm>

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/wire_format_lite.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>

PROTOBUF_PRAGMA_INIT_SEG

namespace _pb = ::PROTOBUF_NAMESPACE_ID;
namespace _pbi = _pb::internal;

namespace pb {
PROTOBUF_CONSTEXPR StringMap_SEntry_DoNotUse::StringMap_SEntry_DoNotUse(
    ::_pbi::ConstantInitialized){}
struct StringMap_SEntry_DoNotUseDefaultTypeInternal {
  PROTOBUF_CONSTEXPR StringMap_SEntry_DoNotUseDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~StringMap_SEntry_DoNotUseDefaultTypeInternal() {}
  union {
    StringMap_SEntry_DoNotUse _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 StringMap_SEntry_DoNotUseDefaultTypeInternal _StringMap_SEntry_DoNotUse_default_instance_;
PROTOBUF_CONSTEXPR StringMap::StringMap(
    ::_pbi::ConstantInitialized)
  : s_(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized{})
  , schema_(&::_pbi::fixed_address_empty_string, ::_pbi::ConstantInitialized{}){}
struct StringMapDefaultTypeInternal {
  PROTOBUF_CONSTEXPR StringMapDefaultTypeInternal()
      : _instance(::_pbi::ConstantInitialized{}) {}
  ~StringMapDefaultTypeInternal() {}
  union {
    StringMap _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT PROTOBUF_ATTRIBUTE_INIT_PRIORITY1 StringMapDefaultTypeInternal _StringMap_default_instance_;
}  // namespace pb
static ::_pb::Metadata file_level_metadata_stringmap_5fext_2eproto[2];
static constexpr ::_pb::EnumDescriptor const** file_level_enum_descriptors_stringmap_5fext_2eproto = nullptr;
static constexpr ::_pb::ServiceDescriptor const** file_level_service_descriptors_stringmap_5fext_2eproto = nullptr;

const uint32_t TableStruct_stringmap_5fext_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  PROTOBUF_FIELD_OFFSET(::pb::StringMap_SEntry_DoNotUse, _has_bits_),
  PROTOBUF_FIELD_OFFSET(::pb::StringMap_SEntry_DoNotUse, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::pb::StringMap_SEntry_DoNotUse, key_),
  PROTOBUF_FIELD_OFFSET(::pb::StringMap_SEntry_DoNotUse, value_),
  0,
  1,
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::pb::StringMap, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  ~0u,  // no _inlined_string_donated_
  PROTOBUF_FIELD_OFFSET(::pb::StringMap, schema_),
  PROTOBUF_FIELD_OFFSET(::pb::StringMap, s_),
};
static const ::_pbi::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, 8, -1, sizeof(::pb::StringMap_SEntry_DoNotUse)},
  { 10, -1, -1, sizeof(::pb::StringMap)},
};

static const ::_pb::Message* const file_default_instances[] = {
  &::pb::_StringMap_SEntry_DoNotUse_default_instance_._instance,
  &::pb::_StringMap_default_instance_._instance,
};

const char descriptor_table_protodef_stringmap_5fext_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n\023stringmap_ext.proto\022\002pb\"f\n\tStringMap\022\016"
  "\n\006schema\030\001 \001(\t\022\037\n\001s\030\002 \003(\0132\024.pb.StringMap"
  ".SEntry\032(\n\006SEntry\022\013\n\003key\030\001 \001(\t\022\r\n\005value\030"
  "\002 \001(\t:\0028\001b\006proto3"
  ;
static ::_pbi::once_flag descriptor_table_stringmap_5fext_2eproto_once;
const ::_pbi::DescriptorTable descriptor_table_stringmap_5fext_2eproto = {
    false, false, 137, descriptor_table_protodef_stringmap_5fext_2eproto,
    "stringmap_ext.proto",
    &descriptor_table_stringmap_5fext_2eproto_once, nullptr, 0, 2,
    schemas, file_default_instances, TableStruct_stringmap_5fext_2eproto::offsets,
    file_level_metadata_stringmap_5fext_2eproto, file_level_enum_descriptors_stringmap_5fext_2eproto,
    file_level_service_descriptors_stringmap_5fext_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK const ::_pbi::DescriptorTable* descriptor_table_stringmap_5fext_2eproto_getter() {
  return &descriptor_table_stringmap_5fext_2eproto;
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY2 static ::_pbi::AddDescriptorsRunner dynamic_init_dummy_stringmap_5fext_2eproto(&descriptor_table_stringmap_5fext_2eproto);
namespace pb {

// ===================================================================

StringMap_SEntry_DoNotUse::StringMap_SEntry_DoNotUse() {}
StringMap_SEntry_DoNotUse::StringMap_SEntry_DoNotUse(::PROTOBUF_NAMESPACE_ID::Arena* arena)
    : SuperType(arena) {}
void StringMap_SEntry_DoNotUse::MergeFrom(const StringMap_SEntry_DoNotUse& other) {
  MergeFromInternal(other);
}
::PROTOBUF_NAMESPACE_ID::Metadata StringMap_SEntry_DoNotUse::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_stringmap_5fext_2eproto_getter, &descriptor_table_stringmap_5fext_2eproto_once,
      file_level_metadata_stringmap_5fext_2eproto[0]);
}

// ===================================================================

class StringMap::_Internal {
 public:
};

StringMap::StringMap(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                         bool is_message_owned)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena, is_message_owned),
  s_(arena) {
  SharedCtor();
  if (arena != nullptr && !is_message_owned) {
    arena->OwnCustomDestructor(this, &StringMap::ArenaDtor);
  }
  // @@protoc_insertion_point(arena_constructor:pb.StringMap)
}
StringMap::StringMap(const StringMap& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  s_.MergeFrom(from.s_);
  schema_.InitDefault();
  #ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
    schema_.Set("", GetArenaForAllocation());
  #endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
  if (!from._internal_schema().empty()) {
    schema_.Set(from._internal_schema(), 
      GetArenaForAllocation());
  }
  // @@protoc_insertion_point(copy_constructor:pb.StringMap)
}

inline void StringMap::SharedCtor() {
schema_.InitDefault();
#ifdef PROTOBUF_FORCE_COPY_DEFAULT_STRING
  schema_.Set("", GetArenaForAllocation());
#endif // PROTOBUF_FORCE_COPY_DEFAULT_STRING
}

StringMap::~StringMap() {
  // @@protoc_insertion_point(destructor:pb.StringMap)
  if (auto *arena = _internal_metadata_.DeleteReturnArena<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>()) {
  (void)arena;
    ArenaDtor(this);
    return;
  }
  SharedDtor();
}

inline void StringMap::SharedDtor() {
  GOOGLE_DCHECK(GetArenaForAllocation() == nullptr);
  s_.Destruct();
  schema_.Destroy();
}

void StringMap::ArenaDtor(void* object) {
  StringMap* _this = reinterpret_cast< StringMap* >(object);
  _this->s_.Destruct();
}
void StringMap::SetCachedSize(int size) const {
  _cached_size_.Set(size);
}

void StringMap::Clear() {
// @@protoc_insertion_point(message_clear_start:pb.StringMap)
  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  s_.Clear();
  schema_.ClearToEmpty();
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* StringMap::_InternalParse(const char* ptr, ::_pbi::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    uint32_t tag;
    ptr = ::_pbi::ReadTag(ptr, &tag);
    switch (tag >> 3) {
      // string schema = 1;
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 10)) {
          auto str = _internal_mutable_schema();
          ptr = ::_pbi::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(ptr);
          CHK_(::_pbi::VerifyUTF8(str, "pb.StringMap.schema"));
        } else
          goto handle_unusual;
        continue;
      // map<string, string> s = 2;
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<uint8_t>(tag) == 18)) {
          ptr -= 1;
          do {
            ptr += 1;
            ptr = ctx->ParseMessage(&s_, ptr);
            CHK_(ptr);
            if (!ctx->DataAvailable(ptr)) break;
          } while (::PROTOBUF_NAMESPACE_ID::internal::ExpectTag<18>(ptr));
        } else
          goto handle_unusual;
        continue;
      default:
        goto handle_unusual;
    }  // switch
  handle_unusual:
    if ((tag == 0) || ((tag & 7) == 4)) {
      CHK_(ptr);
      ctx->SetLastTag(tag);
      goto message_done;
    }
    ptr = UnknownFieldParse(
        tag,
        _internal_metadata_.mutable_unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(),
        ptr, ctx);
    CHK_(ptr != nullptr);
  }  // while
message_done:
  return ptr;
failure:
  ptr = nullptr;
  goto message_done;
#undef CHK_
}

uint8_t* StringMap::_InternalSerialize(
    uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:pb.StringMap)
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  // string schema = 1;
  if (!this->_internal_schema().empty()) {
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
      this->_internal_schema().data(), static_cast<int>(this->_internal_schema().length()),
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE,
      "pb.StringMap.schema");
    target = stream->WriteStringMaybeAliased(
        1, this->_internal_schema(), target);
  }

  // map<string, string> s = 2;
  if (!this->_internal_s().empty()) {
    using MapType = ::_pb::Map<std::string, std::string>;
    using WireHelper = StringMap_SEntry_DoNotUse::Funcs;
    const auto& map_field = this->_internal_s();
    auto check_utf8 = [](const MapType::value_type& entry) {
      (void)entry;
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
        entry.first.data(), static_cast<int>(entry.first.length()),
        ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE,
        "pb.StringMap.SEntry.key");
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
        entry.second.data(), static_cast<int>(entry.second.length()),
        ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE,
        "pb.StringMap.SEntry.value");
    };

    if (stream->IsSerializationDeterministic() && map_field.size() > 1) {
      for (const auto& entry : ::_pbi::MapSorterPtr<MapType>(map_field)) {
        target = WireHelper::InternalSerialize(2, entry.first, entry.second, target, stream);
        check_utf8(entry);
      }
    } else {
      for (const auto& entry : map_field) {
        target = WireHelper::InternalSerialize(2, entry.first, entry.second, target, stream);
        check_utf8(entry);
      }
    }
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::_pbi::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:pb.StringMap)
  return target;
}

size_t StringMap::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:pb.StringMap)
  size_t total_size = 0;

  uint32_t cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // map<string, string> s = 2;
  total_size += 1 *
      ::PROTOBUF_NAMESPACE_ID::internal::FromIntSize(this->_internal_s_size());
  for (::PROTOBUF_NAMESPACE_ID::Map< std::string, std::string >::const_iterator
      it = this->_internal_s().begin();
      it != this->_internal_s().end(); ++it) {
    total_size += StringMap_SEntry_DoNotUse::Funcs::ByteSizeLong(it->first, it->second);
  }

  // string schema = 1;
  if (!this->_internal_schema().empty()) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
        this->_internal_schema());
  }

  return MaybeComputeUnknownFieldsSize(total_size, &_cached_size_);
}

const ::PROTOBUF_NAMESPACE_ID::Message::ClassData StringMap::_class_data_ = {
    ::PROTOBUF_NAMESPACE_ID::Message::CopyWithSizeCheck,
    StringMap::MergeImpl
};
const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*StringMap::GetClassData() const { return &_class_data_; }

void StringMap::MergeImpl(::PROTOBUF_NAMESPACE_ID::Message* to,
                      const ::PROTOBUF_NAMESPACE_ID::Message& from) {
  static_cast<StringMap *>(to)->MergeFrom(
      static_cast<const StringMap &>(from));
}


void StringMap::MergeFrom(const StringMap& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:pb.StringMap)
  GOOGLE_DCHECK_NE(&from, this);
  uint32_t cached_has_bits = 0;
  (void) cached_has_bits;

  s_.MergeFrom(from.s_);
  if (!from._internal_schema().empty()) {
    _internal_set_schema(from._internal_schema());
  }
  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
}

void StringMap::CopyFrom(const StringMap& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:pb.StringMap)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool StringMap::IsInitialized() const {
  return true;
}

void StringMap::InternalSwap(StringMap* other) {
  using std::swap;
  auto* lhs_arena = GetArenaForAllocation();
  auto* rhs_arena = other->GetArenaForAllocation();
  _internal_metadata_.InternalSwap(&other->_internal_metadata_);
  s_.InternalSwap(&other->s_);
  ::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr::InternalSwap(
      &schema_, lhs_arena,
      &other->schema_, rhs_arena
  );
}

::PROTOBUF_NAMESPACE_ID::Metadata StringMap::GetMetadata() const {
  return ::_pbi::AssignDescriptors(
      &descriptor_table_stringmap_5fext_2eproto_getter, &descriptor_table_stringmap_5fext_2eproto_once,
      file_level_metadata_stringmap_5fext_2eproto[1]);
}

// @@protoc_insertion_point(namespace_scope)
}  // namespace pb
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::pb::StringMap_SEntry_DoNotUse*
Arena::CreateMaybeMessage< ::pb::StringMap_SEntry_DoNotUse >(Arena* arena) {
  return Arena::CreateMessageInternal< ::pb::StringMap_SEntry_DoNotUse >(arena);
}
template<> PROTOBUF_NOINLINE ::pb::StringMap*
Arena::CreateMaybeMessage< ::pb::StringMap >(Arena* arena) {
  return Arena::CreateMessageInternal< ::pb::StringMap >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
