/**
 * Autogenerated by Thrift Compiler (0.9.3)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 *  @generated
 */
#include "RPCService_types.h"

#include <algorithm>
#include <ostream>

#include <thrift/TToString.h>

namespace powermax { namespace rpc {


DoubleValue::~DoubleValue() throw() {
}


void DoubleValue::__set_tag(const std::string& val) {
  this->tag = val;
}

void DoubleValue::__set_tm(const int64_t val) {
  this->tm = val;
}

void DoubleValue::__set_status(const int32_t val) {
  this->status = val;
}

void DoubleValue::__set_value(const double val) {
  this->value = val;
}

uint32_t DoubleValue::read(::apache::thrift::protocol::TProtocol* iprot) {

  apache::thrift::protocol::TInputRecursionTracker tracker(*iprot);
  uint32_t xfer = 0;
  std::string fname;
  ::apache::thrift::protocol::TType ftype;
  int16_t fid;

  xfer += iprot->readStructBegin(fname);

  using ::apache::thrift::protocol::TProtocolException;


  while (true)
  {
    xfer += iprot->readFieldBegin(fname, ftype, fid);
    if (ftype == ::apache::thrift::protocol::T_STOP) {
      break;
    }
    switch (fid)
    {
      case 1:
        if (ftype == ::apache::thrift::protocol::T_STRING) {
          xfer += iprot->readString(this->tag);
          this->__isset.tag = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 2:
        if (ftype == ::apache::thrift::protocol::T_I64) {
          xfer += iprot->readI64(this->tm);
          this->__isset.tm = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 3:
        if (ftype == ::apache::thrift::protocol::T_I32) {
          xfer += iprot->readI32(this->status);
          this->__isset.status = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 4:
        if (ftype == ::apache::thrift::protocol::T_DOUBLE) {
          xfer += iprot->readDouble(this->value);
          this->__isset.value = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      default:
        xfer += iprot->skip(ftype);
        break;
    }
    xfer += iprot->readFieldEnd();
  }

  xfer += iprot->readStructEnd();

  return xfer;
}

uint32_t DoubleValue::write(::apache::thrift::protocol::TProtocol* oprot) const {
  uint32_t xfer = 0;
  apache::thrift::protocol::TOutputRecursionTracker tracker(*oprot);
  xfer += oprot->writeStructBegin("DoubleValue");

  xfer += oprot->writeFieldBegin("tag", ::apache::thrift::protocol::T_STRING, 1);
  xfer += oprot->writeString(this->tag);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("tm", ::apache::thrift::protocol::T_I64, 2);
  xfer += oprot->writeI64(this->tm);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("status", ::apache::thrift::protocol::T_I32, 3);
  xfer += oprot->writeI32(this->status);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("value", ::apache::thrift::protocol::T_DOUBLE, 4);
  xfer += oprot->writeDouble(this->value);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldStop();
  xfer += oprot->writeStructEnd();
  return xfer;
}

void swap(DoubleValue &a, DoubleValue &b) {
  using ::std::swap;
  swap(a.tag, b.tag);
  swap(a.tm, b.tm);
  swap(a.status, b.status);
  swap(a.value, b.value);
  swap(a.__isset, b.__isset);
}

DoubleValue::DoubleValue(const DoubleValue& other0) {
  tag = other0.tag;
  tm = other0.tm;
  status = other0.status;
  value = other0.value;
  __isset = other0.__isset;
}
DoubleValue& DoubleValue::operator=(const DoubleValue& other1) {
  tag = other1.tag;
  tm = other1.tm;
  status = other1.status;
  value = other1.value;
  __isset = other1.__isset;
  return *this;
}
void DoubleValue::printTo(std::ostream& out) const {
  using ::apache::thrift::to_string;
  out << "DoubleValue(";
  out << "tag=" << to_string(tag);
  out << ", " << "tm=" << to_string(tm);
  out << ", " << "status=" << to_string(status);
  out << ", " << "value=" << to_string(value);
  out << ")";
}


SimpleReturn::~SimpleReturn() throw() {
}


void SimpleReturn::__set_code(const std::string& val) {
  this->code = val;
}

void SimpleReturn::__set_msg(const std::string& val) {
  this->msg = val;
}

uint32_t SimpleReturn::read(::apache::thrift::protocol::TProtocol* iprot) {

  apache::thrift::protocol::TInputRecursionTracker tracker(*iprot);
  uint32_t xfer = 0;
  std::string fname;
  ::apache::thrift::protocol::TType ftype;
  int16_t fid;

  xfer += iprot->readStructBegin(fname);

  using ::apache::thrift::protocol::TProtocolException;


  while (true)
  {
    xfer += iprot->readFieldBegin(fname, ftype, fid);
    if (ftype == ::apache::thrift::protocol::T_STOP) {
      break;
    }
    switch (fid)
    {
      case 1:
        if (ftype == ::apache::thrift::protocol::T_STRING) {
          xfer += iprot->readString(this->code);
          this->__isset.code = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      case 2:
        if (ftype == ::apache::thrift::protocol::T_STRING) {
          xfer += iprot->readString(this->msg);
          this->__isset.msg = true;
        } else {
          xfer += iprot->skip(ftype);
        }
        break;
      default:
        xfer += iprot->skip(ftype);
        break;
    }
    xfer += iprot->readFieldEnd();
  }

  xfer += iprot->readStructEnd();

  return xfer;
}

uint32_t SimpleReturn::write(::apache::thrift::protocol::TProtocol* oprot) const {
  uint32_t xfer = 0;
  apache::thrift::protocol::TOutputRecursionTracker tracker(*oprot);
  xfer += oprot->writeStructBegin("SimpleReturn");

  xfer += oprot->writeFieldBegin("code", ::apache::thrift::protocol::T_STRING, 1);
  xfer += oprot->writeString(this->code);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldBegin("msg", ::apache::thrift::protocol::T_STRING, 2);
  xfer += oprot->writeString(this->msg);
  xfer += oprot->writeFieldEnd();

  xfer += oprot->writeFieldStop();
  xfer += oprot->writeStructEnd();
  return xfer;
}

void swap(SimpleReturn &a, SimpleReturn &b) {
  using ::std::swap;
  swap(a.code, b.code);
  swap(a.msg, b.msg);
  swap(a.__isset, b.__isset);
}

SimpleReturn::SimpleReturn(const SimpleReturn& other2) {
  code = other2.code;
  msg = other2.msg;
  __isset = other2.__isset;
}
SimpleReturn& SimpleReturn::operator=(const SimpleReturn& other3) {
  code = other3.code;
  msg = other3.msg;
  __isset = other3.__isset;
  return *this;
}
void SimpleReturn::printTo(std::ostream& out) const {
  using ::apache::thrift::to_string;
  out << "SimpleReturn(";
  out << "code=" << to_string(code);
  out << ", " << "msg=" << to_string(msg);
  out << ")";
}

}} // namespace
