// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: stream.proto

#ifndef GOOGLE_PROTOBUF_INCLUDED_stream_2eproto
#define GOOGLE_PROTOBUF_INCLUDED_stream_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3020000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3020001 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/metadata_lite.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/generated_enum_reflection.h>
#include <google/protobuf/unknown_field_set.h>
#include "metadata.pb.h"
#include "source.pb.h"
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_stream_2eproto
PROTOBUF_NAMESPACE_OPEN
namespace internal {
class AnyMetadata;
}  // namespace internal
PROTOBUF_NAMESPACE_CLOSE

// Internal implementation detail -- do not use these members.
struct TableStruct_stream_2eproto {
  static const uint32_t offsets[];
};
extern const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_stream_2eproto;
namespace legacy_pb {
class Stream;
struct StreamDefaultTypeInternal;
extern StreamDefaultTypeInternal _Stream_default_instance_;
}  // namespace legacy_pb
PROTOBUF_NAMESPACE_OPEN
template<> ::legacy_pb::Stream* Arena::CreateMaybeMessage<::legacy_pb::Stream>(Arena*);
PROTOBUF_NAMESPACE_CLOSE
namespace legacy_pb {

enum Stream_Version : int {
  Stream_Version_UNKNOWN_VERSION = 0,
  Stream_Version__0_0_1 = 1
};
bool Stream_Version_IsValid(int value);
constexpr Stream_Version Stream_Version_Version_MIN = Stream_Version_UNKNOWN_VERSION;
constexpr Stream_Version Stream_Version_Version_MAX = Stream_Version__0_0_1;
constexpr int Stream_Version_Version_ARRAYSIZE = Stream_Version_Version_MAX + 1;

const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor* Stream_Version_descriptor();
template<typename T>
inline const std::string& Stream_Version_Name(T enum_t_value) {
  static_assert(::std::is_same<T, Stream_Version>::value ||
    ::std::is_integral<T>::value,
    "Incorrect type passed to function Stream_Version_Name.");
  return ::PROTOBUF_NAMESPACE_ID::internal::NameOfEnum(
    Stream_Version_descriptor(), enum_t_value);
}
inline bool Stream_Version_Parse(
    ::PROTOBUF_NAMESPACE_ID::ConstStringParam name, Stream_Version* value) {
  return ::PROTOBUF_NAMESPACE_ID::internal::ParseNamedEnum<Stream_Version>(
    Stream_Version_descriptor(), name, value);
}
// ===================================================================

class Stream final :
    public ::PROTOBUF_NAMESPACE_ID::Message /* @@protoc_insertion_point(class_definition:legacy_pb.Stream) */ {
 public:
  inline Stream() : Stream(nullptr) {}
  ~Stream() override;
  explicit PROTOBUF_CONSTEXPR Stream(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized);

  Stream(const Stream& from);
  Stream(Stream&& from) noexcept
    : Stream() {
    *this = ::std::move(from);
  }

  inline Stream& operator=(const Stream& from) {
    CopyFrom(from);
    return *this;
  }
  inline Stream& operator=(Stream&& from) noexcept {
    if (this == &from) return *this;
    if (GetOwningArena() == from.GetOwningArena()
  #ifdef PROTOBUF_FORCE_COPY_IN_MOVE
        && GetOwningArena() != nullptr
  #endif  // !PROTOBUF_FORCE_COPY_IN_MOVE
    ) {
      InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }

  inline const ::PROTOBUF_NAMESPACE_ID::UnknownFieldSet& unknown_fields() const {
    return _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance);
  }
  inline ::PROTOBUF_NAMESPACE_ID::UnknownFieldSet* mutable_unknown_fields() {
    return _internal_metadata_.mutable_unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
  }

  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* descriptor() {
    return GetDescriptor();
  }
  static const ::PROTOBUF_NAMESPACE_ID::Descriptor* GetDescriptor() {
    return default_instance().GetMetadata().descriptor;
  }
  static const ::PROTOBUF_NAMESPACE_ID::Reflection* GetReflection() {
    return default_instance().GetMetadata().reflection;
  }
  static const Stream& default_instance() {
    return *internal_default_instance();
  }
  static inline const Stream* internal_default_instance() {
    return reinterpret_cast<const Stream*>(
               &_Stream_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  friend void swap(Stream& a, Stream& b) {
    a.Swap(&b);
  }
  inline void Swap(Stream* other) {
    if (other == this) return;
  #ifdef PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() != nullptr &&
        GetOwningArena() == other->GetOwningArena()) {
   #else  // PROTOBUF_FORCE_COPY_IN_SWAP
    if (GetOwningArena() == other->GetOwningArena()) {
  #endif  // !PROTOBUF_FORCE_COPY_IN_SWAP
      InternalSwap(other);
    } else {
      ::PROTOBUF_NAMESPACE_ID::internal::GenericSwap(this, other);
    }
  }
  void UnsafeArenaSwap(Stream* other) {
    if (other == this) return;
    GOOGLE_DCHECK(GetOwningArena() == other->GetOwningArena());
    InternalSwap(other);
  }

  // implements Message ----------------------------------------------

  Stream* New(::PROTOBUF_NAMESPACE_ID::Arena* arena = nullptr) const final {
    return CreateMaybeMessage<Stream>(arena);
  }
  using ::PROTOBUF_NAMESPACE_ID::Message::CopyFrom;
  void CopyFrom(const Stream& from);
  using ::PROTOBUF_NAMESPACE_ID::Message::MergeFrom;
  void MergeFrom(const Stream& from);
  private:
  static void MergeImpl(::PROTOBUF_NAMESPACE_ID::Message* to, const ::PROTOBUF_NAMESPACE_ID::Message& from);
  public:
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  const char* _InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) final;
  uint8_t* _InternalSerialize(
      uint8_t* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(Stream* other);

  private:
  friend class ::PROTOBUF_NAMESPACE_ID::internal::AnyMetadata;
  static ::PROTOBUF_NAMESPACE_ID::StringPiece FullMessageName() {
    return "legacy_pb.Stream";
  }
  protected:
  explicit Stream(::PROTOBUF_NAMESPACE_ID::Arena* arena,
                       bool is_message_owned = false);
  public:

  static const ClassData _class_data_;
  const ::PROTOBUF_NAMESPACE_ID::Message::ClassData*GetClassData() const final;

  ::PROTOBUF_NAMESPACE_ID::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  typedef Stream_Version Version;
  static constexpr Version UNKNOWN_VERSION =
    Stream_Version_UNKNOWN_VERSION;
  static constexpr Version _0_0_1 =
    Stream_Version__0_0_1;
  static inline bool Version_IsValid(int value) {
    return Stream_Version_IsValid(value);
  }
  static constexpr Version Version_MIN =
    Stream_Version_Version_MIN;
  static constexpr Version Version_MAX =
    Stream_Version_Version_MAX;
  static constexpr int Version_ARRAYSIZE =
    Stream_Version_Version_ARRAYSIZE;
  static inline const ::PROTOBUF_NAMESPACE_ID::EnumDescriptor*
  Version_descriptor() {
    return Stream_Version_descriptor();
  }
  template<typename T>
  static inline const std::string& Version_Name(T enum_t_value) {
    static_assert(::std::is_same<T, Version>::value ||
      ::std::is_integral<T>::value,
      "Incorrect type passed to function Version_Name.");
    return Stream_Version_Name(enum_t_value);
  }
  static inline bool Version_Parse(::PROTOBUF_NAMESPACE_ID::ConstStringParam name,
      Version* value) {
    return Stream_Version_Parse(name, value);
  }

  // accessors -------------------------------------------------------

  enum : int {
    kMetadataFieldNumber = 2,
    kSourceFieldNumber = 3,
    kVersionFieldNumber = 1,
  };
  // required .legacy_pb.Metadata metadata = 2;
  bool has_metadata() const;
  private:
  bool _internal_has_metadata() const;
  public:
  void clear_metadata();
  const ::legacy_pb::Metadata& metadata() const;
  PROTOBUF_NODISCARD ::legacy_pb::Metadata* release_metadata();
  ::legacy_pb::Metadata* mutable_metadata();
  void set_allocated_metadata(::legacy_pb::Metadata* metadata);
  private:
  const ::legacy_pb::Metadata& _internal_metadata() const;
  ::legacy_pb::Metadata* _internal_mutable_metadata();
  public:
  void unsafe_arena_set_allocated_metadata(
      ::legacy_pb::Metadata* metadata);
  ::legacy_pb::Metadata* unsafe_arena_release_metadata();

  // required .legacy_pb.Source source = 3;
  bool has_source() const;
  private:
  bool _internal_has_source() const;
  public:
  void clear_source();
  const ::legacy_pb::Source& source() const;
  PROTOBUF_NODISCARD ::legacy_pb::Source* release_source();
  ::legacy_pb::Source* mutable_source();
  void set_allocated_source(::legacy_pb::Source* source);
  private:
  const ::legacy_pb::Source& _internal_source() const;
  ::legacy_pb::Source* _internal_mutable_source();
  public:
  void unsafe_arena_set_allocated_source(
      ::legacy_pb::Source* source);
  ::legacy_pb::Source* unsafe_arena_release_source();

  // required .legacy_pb.Stream.Version version = 1;
  bool has_version() const;
  private:
  bool _internal_has_version() const;
  public:
  void clear_version();
  ::legacy_pb::Stream_Version version() const;
  void set_version(::legacy_pb::Stream_Version value);
  private:
  ::legacy_pb::Stream_Version _internal_version() const;
  void _internal_set_version(::legacy_pb::Stream_Version value);
  public:

  // @@protoc_insertion_point(class_scope:legacy_pb.Stream)
 private:
  class _Internal;

  // helper for ByteSizeLong()
  size_t RequiredFieldsByteSizeFallback() const;

  template <typename T> friend class ::PROTOBUF_NAMESPACE_ID::Arena::InternalHelper;
  typedef void InternalArenaConstructable_;
  typedef void DestructorSkippable_;
  ::PROTOBUF_NAMESPACE_ID::internal::HasBits<1> _has_bits_;
  mutable ::PROTOBUF_NAMESPACE_ID::internal::CachedSize _cached_size_;
  ::legacy_pb::Metadata* metadata_;
  ::legacy_pb::Source* source_;
  int version_;
  friend struct ::TableStruct_stream_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// Stream

// required .legacy_pb.Stream.Version version = 1;
inline bool Stream::_internal_has_version() const {
  bool value = (_has_bits_[0] & 0x00000004u) != 0;
  return value;
}
inline bool Stream::has_version() const {
  return _internal_has_version();
}
inline void Stream::clear_version() {
  version_ = 0;
  _has_bits_[0] &= ~0x00000004u;
}
inline ::legacy_pb::Stream_Version Stream::_internal_version() const {
  return static_cast< ::legacy_pb::Stream_Version >(version_);
}
inline ::legacy_pb::Stream_Version Stream::version() const {
  // @@protoc_insertion_point(field_get:legacy_pb.Stream.version)
  return _internal_version();
}
inline void Stream::_internal_set_version(::legacy_pb::Stream_Version value) {
  assert(::legacy_pb::Stream_Version_IsValid(value));
  _has_bits_[0] |= 0x00000004u;
  version_ = value;
}
inline void Stream::set_version(::legacy_pb::Stream_Version value) {
  _internal_set_version(value);
  // @@protoc_insertion_point(field_set:legacy_pb.Stream.version)
}

// required .legacy_pb.Metadata metadata = 2;
inline bool Stream::_internal_has_metadata() const {
  bool value = (_has_bits_[0] & 0x00000001u) != 0;
  PROTOBUF_ASSUME(!value || metadata_ != nullptr);
  return value;
}
inline bool Stream::has_metadata() const {
  return _internal_has_metadata();
}
inline const ::legacy_pb::Metadata& Stream::_internal_metadata() const {
  const ::legacy_pb::Metadata* p = metadata_;
  return p != nullptr ? *p : reinterpret_cast<const ::legacy_pb::Metadata&>(
      ::legacy_pb::_Metadata_default_instance_);
}
inline const ::legacy_pb::Metadata& Stream::metadata() const {
  // @@protoc_insertion_point(field_get:legacy_pb.Stream.metadata)
  return _internal_metadata();
}
inline void Stream::unsafe_arena_set_allocated_metadata(
    ::legacy_pb::Metadata* metadata) {
  if (GetArenaForAllocation() == nullptr) {
    delete reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(metadata_);
  }
  metadata_ = metadata;
  if (metadata) {
    _has_bits_[0] |= 0x00000001u;
  } else {
    _has_bits_[0] &= ~0x00000001u;
  }
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:legacy_pb.Stream.metadata)
}
inline ::legacy_pb::Metadata* Stream::release_metadata() {
  _has_bits_[0] &= ~0x00000001u;
  ::legacy_pb::Metadata* temp = metadata_;
  metadata_ = nullptr;
#ifdef PROTOBUF_FORCE_COPY_IN_RELEASE
  auto* old =  reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(temp);
  temp = ::PROTOBUF_NAMESPACE_ID::internal::DuplicateIfNonNull(temp);
  if (GetArenaForAllocation() == nullptr) { delete old; }
#else  // PROTOBUF_FORCE_COPY_IN_RELEASE
  if (GetArenaForAllocation() != nullptr) {
    temp = ::PROTOBUF_NAMESPACE_ID::internal::DuplicateIfNonNull(temp);
  }
#endif  // !PROTOBUF_FORCE_COPY_IN_RELEASE
  return temp;
}
inline ::legacy_pb::Metadata* Stream::unsafe_arena_release_metadata() {
  // @@protoc_insertion_point(field_release:legacy_pb.Stream.metadata)
  _has_bits_[0] &= ~0x00000001u;
  ::legacy_pb::Metadata* temp = metadata_;
  metadata_ = nullptr;
  return temp;
}
inline ::legacy_pb::Metadata* Stream::_internal_mutable_metadata() {
  _has_bits_[0] |= 0x00000001u;
  if (metadata_ == nullptr) {
    auto* p = CreateMaybeMessage<::legacy_pb::Metadata>(GetArenaForAllocation());
    metadata_ = p;
  }
  return metadata_;
}
inline ::legacy_pb::Metadata* Stream::mutable_metadata() {
  ::legacy_pb::Metadata* _msg = _internal_mutable_metadata();
  // @@protoc_insertion_point(field_mutable:legacy_pb.Stream.metadata)
  return _msg;
}
inline void Stream::set_allocated_metadata(::legacy_pb::Metadata* metadata) {
  ::PROTOBUF_NAMESPACE_ID::Arena* message_arena = GetArenaForAllocation();
  if (message_arena == nullptr) {
    delete reinterpret_cast< ::PROTOBUF_NAMESPACE_ID::MessageLite*>(metadata_);
  }
  if (metadata) {
    ::PROTOBUF_NAMESPACE_ID::Arena* submessage_arena =
        ::PROTOBUF_NAMESPACE_ID::Arena::InternalGetOwningArena(
                reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(metadata));
    if (message_arena != submessage_arena) {
      metadata = ::PROTOBUF_NAMESPACE_ID::internal::GetOwnedMessage(
          message_arena, metadata, submessage_arena);
    }
    _has_bits_[0] |= 0x00000001u;
  } else {
    _has_bits_[0] &= ~0x00000001u;
  }
  metadata_ = metadata;
  // @@protoc_insertion_point(field_set_allocated:legacy_pb.Stream.metadata)
}

// required .legacy_pb.Source source = 3;
inline bool Stream::_internal_has_source() const {
  bool value = (_has_bits_[0] & 0x00000002u) != 0;
  PROTOBUF_ASSUME(!value || source_ != nullptr);
  return value;
}
inline bool Stream::has_source() const {
  return _internal_has_source();
}
inline const ::legacy_pb::Source& Stream::_internal_source() const {
  const ::legacy_pb::Source* p = source_;
  return p != nullptr ? *p : reinterpret_cast<const ::legacy_pb::Source&>(
      ::legacy_pb::_Source_default_instance_);
}
inline const ::legacy_pb::Source& Stream::source() const {
  // @@protoc_insertion_point(field_get:legacy_pb.Stream.source)
  return _internal_source();
}
inline void Stream::unsafe_arena_set_allocated_source(
    ::legacy_pb::Source* source) {
  if (GetArenaForAllocation() == nullptr) {
    delete reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(source_);
  }
  source_ = source;
  if (source) {
    _has_bits_[0] |= 0x00000002u;
  } else {
    _has_bits_[0] &= ~0x00000002u;
  }
  // @@protoc_insertion_point(field_unsafe_arena_set_allocated:legacy_pb.Stream.source)
}
inline ::legacy_pb::Source* Stream::release_source() {
  _has_bits_[0] &= ~0x00000002u;
  ::legacy_pb::Source* temp = source_;
  source_ = nullptr;
#ifdef PROTOBUF_FORCE_COPY_IN_RELEASE
  auto* old =  reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(temp);
  temp = ::PROTOBUF_NAMESPACE_ID::internal::DuplicateIfNonNull(temp);
  if (GetArenaForAllocation() == nullptr) { delete old; }
#else  // PROTOBUF_FORCE_COPY_IN_RELEASE
  if (GetArenaForAllocation() != nullptr) {
    temp = ::PROTOBUF_NAMESPACE_ID::internal::DuplicateIfNonNull(temp);
  }
#endif  // !PROTOBUF_FORCE_COPY_IN_RELEASE
  return temp;
}
inline ::legacy_pb::Source* Stream::unsafe_arena_release_source() {
  // @@protoc_insertion_point(field_release:legacy_pb.Stream.source)
  _has_bits_[0] &= ~0x00000002u;
  ::legacy_pb::Source* temp = source_;
  source_ = nullptr;
  return temp;
}
inline ::legacy_pb::Source* Stream::_internal_mutable_source() {
  _has_bits_[0] |= 0x00000002u;
  if (source_ == nullptr) {
    auto* p = CreateMaybeMessage<::legacy_pb::Source>(GetArenaForAllocation());
    source_ = p;
  }
  return source_;
}
inline ::legacy_pb::Source* Stream::mutable_source() {
  ::legacy_pb::Source* _msg = _internal_mutable_source();
  // @@protoc_insertion_point(field_mutable:legacy_pb.Stream.source)
  return _msg;
}
inline void Stream::set_allocated_source(::legacy_pb::Source* source) {
  ::PROTOBUF_NAMESPACE_ID::Arena* message_arena = GetArenaForAllocation();
  if (message_arena == nullptr) {
    delete reinterpret_cast< ::PROTOBUF_NAMESPACE_ID::MessageLite*>(source_);
  }
  if (source) {
    ::PROTOBUF_NAMESPACE_ID::Arena* submessage_arena =
        ::PROTOBUF_NAMESPACE_ID::Arena::InternalGetOwningArena(
                reinterpret_cast<::PROTOBUF_NAMESPACE_ID::MessageLite*>(source));
    if (message_arena != submessage_arena) {
      source = ::PROTOBUF_NAMESPACE_ID::internal::GetOwnedMessage(
          message_arena, source, submessage_arena);
    }
    _has_bits_[0] |= 0x00000002u;
  } else {
    _has_bits_[0] &= ~0x00000002u;
  }
  source_ = source;
  // @@protoc_insertion_point(field_set_allocated:legacy_pb.Stream.source)
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)

}  // namespace legacy_pb

PROTOBUF_NAMESPACE_OPEN

template <> struct is_proto_enum< ::legacy_pb::Stream_Version> : ::std::true_type {};
template <>
inline const EnumDescriptor* GetEnumDescriptor< ::legacy_pb::Stream_Version>() {
  return ::legacy_pb::Stream_Version_descriptor();
}

PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // GOOGLE_PROTOBUF_INCLUDED_GOOGLE_PROTOBUF_INCLUDED_stream_2eproto
