// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package rpc

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type RpcService interface {
	Ping() (r *SimpleReturn, err error)
	// Parameters:
	//  - DoubleValue
	SendDoubleData(doubleValue []*DoubleValue) (r *SimpleReturn, err error)
}

type RpcServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewRpcServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *RpcServiceClient {
	return &RpcServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewRpcServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *RpcServiceClient {
	return &RpcServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

func (p *RpcServiceClient) Ping() (r *SimpleReturn, err error) {
	if err = p.sendPing(); err != nil {
		return
	}
	return p.recvPing()
}

func (p *RpcServiceClient) sendPing() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("ping", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := RpcServicePingArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *RpcServiceClient) recvPing() (value *SimpleReturn, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "ping" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "ping failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "ping failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1 error
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "ping failed: invalid message type")
		return
	}
	result := RpcServicePingResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - DoubleValue
func (p *RpcServiceClient) SendDoubleData(doubleValue []*DoubleValue) (r *SimpleReturn, err error) {
	if err = p.sendSendDoubleData(doubleValue); err != nil {
		return
	}
	return p.recvSendDoubleData()
}

func (p *RpcServiceClient) sendSendDoubleData(doubleValue []*DoubleValue) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("sendDoubleData", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := RpcServiceSendDoubleDataArgs{
		DoubleValue: doubleValue,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *RpcServiceClient) recvSendDoubleData() (value *SimpleReturn, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "sendDoubleData" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "sendDoubleData failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "sendDoubleData failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "sendDoubleData failed: invalid message type")
		return
	}
	result := RpcServiceSendDoubleDataResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type RpcServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      RpcService
}

func (p *RpcServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *RpcServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *RpcServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewRpcServiceProcessor(handler RpcService) *RpcServiceProcessor {

	self4 := &RpcServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self4.processorMap["ping"] = &rpcServiceProcessorPing{handler: handler}
	self4.processorMap["sendDoubleData"] = &rpcServiceProcessorSendDoubleData{handler: handler}
	return self4
}

func (p *RpcServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x5 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x5.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x5

}

type rpcServiceProcessorPing struct {
	handler RpcService
}

func (p *rpcServiceProcessorPing) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := RpcServicePingArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("ping", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := RpcServicePingResult{}
	var retval *SimpleReturn
	var err2 error
	if retval, err2 = p.handler.Ping(); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing ping: "+err2.Error())
		oprot.WriteMessageBegin("ping", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("ping", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type rpcServiceProcessorSendDoubleData struct {
	handler RpcService
}

func (p *rpcServiceProcessorSendDoubleData) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := RpcServiceSendDoubleDataArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("sendDoubleData", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := RpcServiceSendDoubleDataResult{}
	var retval *SimpleReturn
	var err2 error
	if retval, err2 = p.handler.SendDoubleData(args.DoubleValue); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing sendDoubleData: "+err2.Error())
		oprot.WriteMessageBegin("sendDoubleData", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("sendDoubleData", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type RpcServicePingArgs struct {
}

func NewRpcServicePingArgs() *RpcServicePingArgs {
	return &RpcServicePingArgs{}
}

func (p *RpcServicePingArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RpcServicePingArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ping_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RpcServicePingArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RpcServicePingArgs(%+v)", *p)
}

// Attributes:
//  - Success
type RpcServicePingResult struct {
	Success *SimpleReturn `thrift:"success,0" json:"success,omitempty"`
}

func NewRpcServicePingResult() *RpcServicePingResult {
	return &RpcServicePingResult{}
}

var RpcServicePingResult_Success_DEFAULT *SimpleReturn

func (p *RpcServicePingResult) GetSuccess() *SimpleReturn {
	if !p.IsSetSuccess() {
		return RpcServicePingResult_Success_DEFAULT
	}
	return p.Success
}
func (p *RpcServicePingResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RpcServicePingResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RpcServicePingResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &SimpleReturn{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *RpcServicePingResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("ping_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RpcServicePingResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *RpcServicePingResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RpcServicePingResult(%+v)", *p)
}

// Attributes:
//  - DoubleValue
type RpcServiceSendDoubleDataArgs struct {
	DoubleValue []*DoubleValue `thrift:"doubleValue,1" json:"doubleValue"`
}

func NewRpcServiceSendDoubleDataArgs() *RpcServiceSendDoubleDataArgs {
	return &RpcServiceSendDoubleDataArgs{}
}

func (p *RpcServiceSendDoubleDataArgs) GetDoubleValue() []*DoubleValue {
	return p.DoubleValue
}
func (p *RpcServiceSendDoubleDataArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.readField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RpcServiceSendDoubleDataArgs) readField1(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return thrift.PrependError("error reading list begin: ", err)
	}
	tSlice := make([]*DoubleValue, 0, size)
	p.DoubleValue = tSlice
	for i := 0; i < size; i++ {
		_elem6 := &DoubleValue{}
		if err := _elem6.Read(iprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", _elem6), err)
		}
		p.DoubleValue = append(p.DoubleValue, _elem6)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return thrift.PrependError("error reading list end: ", err)
	}
	return nil
}

func (p *RpcServiceSendDoubleDataArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("sendDoubleData_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RpcServiceSendDoubleDataArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("doubleValue", thrift.LIST, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:doubleValue: ", p), err)
	}
	if err := oprot.WriteListBegin(thrift.STRUCT, len(p.DoubleValue)); err != nil {
		return thrift.PrependError("error writing list begin: ", err)
	}
	for _, v := range p.DoubleValue {
		if err := v.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", v), err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return thrift.PrependError("error writing list end: ", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:doubleValue: ", p), err)
	}
	return err
}

func (p *RpcServiceSendDoubleDataArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RpcServiceSendDoubleDataArgs(%+v)", *p)
}

// Attributes:
//  - Success
type RpcServiceSendDoubleDataResult struct {
	Success *SimpleReturn `thrift:"success,0" json:"success,omitempty"`
}

func NewRpcServiceSendDoubleDataResult() *RpcServiceSendDoubleDataResult {
	return &RpcServiceSendDoubleDataResult{}
}

var RpcServiceSendDoubleDataResult_Success_DEFAULT *SimpleReturn

func (p *RpcServiceSendDoubleDataResult) GetSuccess() *SimpleReturn {
	if !p.IsSetSuccess() {
		return RpcServiceSendDoubleDataResult_Success_DEFAULT
	}
	return p.Success
}
func (p *RpcServiceSendDoubleDataResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *RpcServiceSendDoubleDataResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.readField0(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *RpcServiceSendDoubleDataResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &SimpleReturn{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *RpcServiceSendDoubleDataResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("sendDoubleData_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *RpcServiceSendDoubleDataResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *RpcServiceSendDoubleDataResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("RpcServiceSendDoubleDataResult(%+v)", *p)
}
