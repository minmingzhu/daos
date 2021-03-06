/* Generated by the protocol buffer compiler.  DO NOT EDIT! */
/* Generated from: auth.proto */

/* Do not generate deprecated warnings for self */
#ifndef PROTOBUF_C__NO_DEPRECATED
#define PROTOBUF_C__NO_DEPRECATED
#endif

#include "auth.pb-c.h"
void   auth__token__init
                     (Auth__Token         *message)
{
  static const Auth__Token init_value = AUTH__TOKEN__INIT;
  *message = init_value;
}
size_t auth__token__get_packed_size
                     (const Auth__Token *message)
{
  assert(message->base.descriptor == &auth__token__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t auth__token__pack
                     (const Auth__Token *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &auth__token__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t auth__token__pack_to_buffer
                     (const Auth__Token *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &auth__token__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Auth__Token *
       auth__token__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Auth__Token *)
     protobuf_c_message_unpack (&auth__token__descriptor,
                                allocator, len, data);
}
void   auth__token__free_unpacked
                     (Auth__Token *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &auth__token__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   auth__sys__init
                     (Auth__Sys         *message)
{
  static const Auth__Sys init_value = AUTH__SYS__INIT;
  *message = init_value;
}
size_t auth__sys__get_packed_size
                     (const Auth__Sys *message)
{
  assert(message->base.descriptor == &auth__sys__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t auth__sys__pack
                     (const Auth__Sys *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &auth__sys__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t auth__sys__pack_to_buffer
                     (const Auth__Sys *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &auth__sys__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Auth__Sys *
       auth__sys__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Auth__Sys *)
     protobuf_c_message_unpack (&auth__sys__descriptor,
                                allocator, len, data);
}
void   auth__sys__free_unpacked
                     (Auth__Sys *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &auth__sys__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   auth__credential__init
                     (Auth__Credential         *message)
{
  static const Auth__Credential init_value = AUTH__CREDENTIAL__INIT;
  *message = init_value;
}
size_t auth__credential__get_packed_size
                     (const Auth__Credential *message)
{
  assert(message->base.descriptor == &auth__credential__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t auth__credential__pack
                     (const Auth__Credential *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &auth__credential__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t auth__credential__pack_to_buffer
                     (const Auth__Credential *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &auth__credential__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Auth__Credential *
       auth__credential__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Auth__Credential *)
     protobuf_c_message_unpack (&auth__credential__descriptor,
                                allocator, len, data);
}
void   auth__credential__free_unpacked
                     (Auth__Credential *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &auth__credential__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   auth__get_cred_resp__init
                     (Auth__GetCredResp         *message)
{
  static const Auth__GetCredResp init_value = AUTH__GET_CRED_RESP__INIT;
  *message = init_value;
}
size_t auth__get_cred_resp__get_packed_size
                     (const Auth__GetCredResp *message)
{
  assert(message->base.descriptor == &auth__get_cred_resp__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t auth__get_cred_resp__pack
                     (const Auth__GetCredResp *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &auth__get_cred_resp__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t auth__get_cred_resp__pack_to_buffer
                     (const Auth__GetCredResp *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &auth__get_cred_resp__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Auth__GetCredResp *
       auth__get_cred_resp__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Auth__GetCredResp *)
     protobuf_c_message_unpack (&auth__get_cred_resp__descriptor,
                                allocator, len, data);
}
void   auth__get_cred_resp__free_unpacked
                     (Auth__GetCredResp *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &auth__get_cred_resp__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   auth__validate_cred_req__init
                     (Auth__ValidateCredReq         *message)
{
  static const Auth__ValidateCredReq init_value = AUTH__VALIDATE_CRED_REQ__INIT;
  *message = init_value;
}
size_t auth__validate_cred_req__get_packed_size
                     (const Auth__ValidateCredReq *message)
{
  assert(message->base.descriptor == &auth__validate_cred_req__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t auth__validate_cred_req__pack
                     (const Auth__ValidateCredReq *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &auth__validate_cred_req__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t auth__validate_cred_req__pack_to_buffer
                     (const Auth__ValidateCredReq *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &auth__validate_cred_req__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Auth__ValidateCredReq *
       auth__validate_cred_req__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Auth__ValidateCredReq *)
     protobuf_c_message_unpack (&auth__validate_cred_req__descriptor,
                                allocator, len, data);
}
void   auth__validate_cred_req__free_unpacked
                     (Auth__ValidateCredReq *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &auth__validate_cred_req__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
void   auth__validate_cred_resp__init
                     (Auth__ValidateCredResp         *message)
{
  static const Auth__ValidateCredResp init_value = AUTH__VALIDATE_CRED_RESP__INIT;
  *message = init_value;
}
size_t auth__validate_cred_resp__get_packed_size
                     (const Auth__ValidateCredResp *message)
{
  assert(message->base.descriptor == &auth__validate_cred_resp__descriptor);
  return protobuf_c_message_get_packed_size ((const ProtobufCMessage*)(message));
}
size_t auth__validate_cred_resp__pack
                     (const Auth__ValidateCredResp *message,
                      uint8_t       *out)
{
  assert(message->base.descriptor == &auth__validate_cred_resp__descriptor);
  return protobuf_c_message_pack ((const ProtobufCMessage*)message, out);
}
size_t auth__validate_cred_resp__pack_to_buffer
                     (const Auth__ValidateCredResp *message,
                      ProtobufCBuffer *buffer)
{
  assert(message->base.descriptor == &auth__validate_cred_resp__descriptor);
  return protobuf_c_message_pack_to_buffer ((const ProtobufCMessage*)message, buffer);
}
Auth__ValidateCredResp *
       auth__validate_cred_resp__unpack
                     (ProtobufCAllocator  *allocator,
                      size_t               len,
                      const uint8_t       *data)
{
  return (Auth__ValidateCredResp *)
     protobuf_c_message_unpack (&auth__validate_cred_resp__descriptor,
                                allocator, len, data);
}
void   auth__validate_cred_resp__free_unpacked
                     (Auth__ValidateCredResp *message,
                      ProtobufCAllocator *allocator)
{
  if(!message)
    return;
  assert(message->base.descriptor == &auth__validate_cred_resp__descriptor);
  protobuf_c_message_free_unpacked ((ProtobufCMessage*)message, allocator);
}
static const ProtobufCFieldDescriptor auth__token__field_descriptors[2] =
{
  {
    "flavor",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_ENUM,
    0,   /* quantifier_offset */
    offsetof(Auth__Token, flavor),
    &auth__flavor__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "data",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_BYTES,
    0,   /* quantifier_offset */
    offsetof(Auth__Token, data),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned auth__token__field_indices_by_name[] = {
  1,   /* field[1] = data */
  0,   /* field[0] = flavor */
};
static const ProtobufCIntRange auth__token__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 2 }
};
const ProtobufCMessageDescriptor auth__token__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "auth.Token",
  "Token",
  "Auth__Token",
  "auth",
  sizeof(Auth__Token),
  2,
  auth__token__field_descriptors,
  auth__token__field_indices_by_name,
  1,  auth__token__number_ranges,
  (ProtobufCMessageInit) auth__token__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor auth__sys__field_descriptors[6] =
{
  {
    "stamp",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_UINT64,
    0,   /* quantifier_offset */
    offsetof(Auth__Sys, stamp),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "machinename",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Auth__Sys, machinename),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "user",
    3,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Auth__Sys, user),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "group",
    4,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Auth__Sys, group),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "groups",
    5,
    PROTOBUF_C_LABEL_REPEATED,
    PROTOBUF_C_TYPE_STRING,
    offsetof(Auth__Sys, n_groups),
    offsetof(Auth__Sys, groups),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "secctx",
    6,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Auth__Sys, secctx),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned auth__sys__field_indices_by_name[] = {
  3,   /* field[3] = group */
  4,   /* field[4] = groups */
  1,   /* field[1] = machinename */
  5,   /* field[5] = secctx */
  0,   /* field[0] = stamp */
  2,   /* field[2] = user */
};
static const ProtobufCIntRange auth__sys__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 6 }
};
const ProtobufCMessageDescriptor auth__sys__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "auth.Sys",
  "Sys",
  "Auth__Sys",
  "auth",
  sizeof(Auth__Sys),
  6,
  auth__sys__field_descriptors,
  auth__sys__field_indices_by_name,
  1,  auth__sys__number_ranges,
  (ProtobufCMessageInit) auth__sys__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor auth__credential__field_descriptors[3] =
{
  {
    "token",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_MESSAGE,
    0,   /* quantifier_offset */
    offsetof(Auth__Credential, token),
    &auth__token__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "verifier",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_MESSAGE,
    0,   /* quantifier_offset */
    offsetof(Auth__Credential, verifier),
    &auth__token__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "origin",
    3,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_STRING,
    0,   /* quantifier_offset */
    offsetof(Auth__Credential, origin),
    NULL,
    &protobuf_c_empty_string,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned auth__credential__field_indices_by_name[] = {
  2,   /* field[2] = origin */
  0,   /* field[0] = token */
  1,   /* field[1] = verifier */
};
static const ProtobufCIntRange auth__credential__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 3 }
};
const ProtobufCMessageDescriptor auth__credential__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "auth.Credential",
  "Credential",
  "Auth__Credential",
  "auth",
  sizeof(Auth__Credential),
  3,
  auth__credential__field_descriptors,
  auth__credential__field_indices_by_name,
  1,  auth__credential__number_ranges,
  (ProtobufCMessageInit) auth__credential__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor auth__get_cred_resp__field_descriptors[2] =
{
  {
    "status",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_INT32,
    0,   /* quantifier_offset */
    offsetof(Auth__GetCredResp, status),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "cred",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_MESSAGE,
    0,   /* quantifier_offset */
    offsetof(Auth__GetCredResp, cred),
    &auth__credential__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned auth__get_cred_resp__field_indices_by_name[] = {
  1,   /* field[1] = cred */
  0,   /* field[0] = status */
};
static const ProtobufCIntRange auth__get_cred_resp__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 2 }
};
const ProtobufCMessageDescriptor auth__get_cred_resp__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "auth.GetCredResp",
  "GetCredResp",
  "Auth__GetCredResp",
  "auth",
  sizeof(Auth__GetCredResp),
  2,
  auth__get_cred_resp__field_descriptors,
  auth__get_cred_resp__field_indices_by_name,
  1,  auth__get_cred_resp__number_ranges,
  (ProtobufCMessageInit) auth__get_cred_resp__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor auth__validate_cred_req__field_descriptors[1] =
{
  {
    "cred",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_MESSAGE,
    0,   /* quantifier_offset */
    offsetof(Auth__ValidateCredReq, cred),
    &auth__credential__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned auth__validate_cred_req__field_indices_by_name[] = {
  0,   /* field[0] = cred */
};
static const ProtobufCIntRange auth__validate_cred_req__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 1 }
};
const ProtobufCMessageDescriptor auth__validate_cred_req__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "auth.ValidateCredReq",
  "ValidateCredReq",
  "Auth__ValidateCredReq",
  "auth",
  sizeof(Auth__ValidateCredReq),
  1,
  auth__validate_cred_req__field_descriptors,
  auth__validate_cred_req__field_indices_by_name,
  1,  auth__validate_cred_req__number_ranges,
  (ProtobufCMessageInit) auth__validate_cred_req__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCFieldDescriptor auth__validate_cred_resp__field_descriptors[2] =
{
  {
    "status",
    1,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_INT32,
    0,   /* quantifier_offset */
    offsetof(Auth__ValidateCredResp, status),
    NULL,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
  {
    "token",
    2,
    PROTOBUF_C_LABEL_NONE,
    PROTOBUF_C_TYPE_MESSAGE,
    0,   /* quantifier_offset */
    offsetof(Auth__ValidateCredResp, token),
    &auth__token__descriptor,
    NULL,
    0,             /* flags */
    0,NULL,NULL    /* reserved1,reserved2, etc */
  },
};
static const unsigned auth__validate_cred_resp__field_indices_by_name[] = {
  0,   /* field[0] = status */
  1,   /* field[1] = token */
};
static const ProtobufCIntRange auth__validate_cred_resp__number_ranges[1 + 1] =
{
  { 1, 0 },
  { 0, 2 }
};
const ProtobufCMessageDescriptor auth__validate_cred_resp__descriptor =
{
  PROTOBUF_C__MESSAGE_DESCRIPTOR_MAGIC,
  "auth.ValidateCredResp",
  "ValidateCredResp",
  "Auth__ValidateCredResp",
  "auth",
  sizeof(Auth__ValidateCredResp),
  2,
  auth__validate_cred_resp__field_descriptors,
  auth__validate_cred_resp__field_indices_by_name,
  1,  auth__validate_cred_resp__number_ranges,
  (ProtobufCMessageInit) auth__validate_cred_resp__init,
  NULL,NULL,NULL    /* reserved[123] */
};
static const ProtobufCEnumValue auth__flavor__enum_values_by_number[2] =
{
  { "AUTH_NONE", "AUTH__FLAVOR__AUTH_NONE", 0 },
  { "AUTH_SYS", "AUTH__FLAVOR__AUTH_SYS", 1 },
};
static const ProtobufCIntRange auth__flavor__value_ranges[] = {
{0, 0},{0, 2}
};
static const ProtobufCEnumValueIndex auth__flavor__enum_values_by_name[2] =
{
  { "AUTH_NONE", 0 },
  { "AUTH_SYS", 1 },
};
const ProtobufCEnumDescriptor auth__flavor__descriptor =
{
  PROTOBUF_C__ENUM_DESCRIPTOR_MAGIC,
  "auth.Flavor",
  "Flavor",
  "Auth__Flavor",
  "auth",
  2,
  auth__flavor__enum_values_by_number,
  2,
  auth__flavor__enum_values_by_name,
  1,
  auth__flavor__value_ranges,
  NULL,NULL,NULL,NULL   /* reserved[1234] */
};
