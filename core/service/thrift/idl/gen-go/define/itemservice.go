// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package define

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type ItemService interface {
	// Parameters:
	//  - ItemId
	//  - SkuId
	GetSku(itemId int64, skuId int64) (r *Sku, err error)
	// Parameters:
	//  - ItemId
	GetItemSkuJson(itemId int64) (r string, err error)
	// Parameters:
	//  - ItemId
	//  - IType
	GetItemDetailData(itemId int64, iType int32) (r string, err error)
}

type ItemServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewItemServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ItemServiceClient {
	return &ItemServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}

func NewItemServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ItemServiceClient {
	return &ItemServiceClient{Transport: t,
		ProtocolFactory: nil,
		InputProtocol:   iprot,
		OutputProtocol:  oprot,
		SeqId:           0,
	}
}

// Parameters:
//  - ItemId
//  - SkuId
func (p *ItemServiceClient) GetSku(itemId int64, skuId int64) (r *Sku, err error) {
	if err = p.sendGetSku(itemId, skuId); err != nil {
		return
	}
	return p.recvGetSku()
}

func (p *ItemServiceClient) sendGetSku(itemId int64, skuId int64) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("GetSku", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := ItemServiceGetSkuArgs{
		ItemId: itemId,
		SkuId:  skuId,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ItemServiceClient) recvGetSku() (value *Sku, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "GetSku" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "GetSku failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "GetSku failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error280 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error281 error
		error281, err = error280.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error281
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "GetSku failed: invalid message type")
		return
	}
	result := ItemServiceGetSkuResult{}
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
//  - ItemId
func (p *ItemServiceClient) GetItemSkuJson(itemId int64) (r string, err error) {
	if err = p.sendGetItemSkuJson(itemId); err != nil {
		return
	}
	return p.recvGetItemSkuJson()
}

func (p *ItemServiceClient) sendGetItemSkuJson(itemId int64) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("GetItemSkuJson", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := ItemServiceGetItemSkuJsonArgs{
		ItemId: itemId,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ItemServiceClient) recvGetItemSkuJson() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "GetItemSkuJson" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "GetItemSkuJson failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "GetItemSkuJson failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error282 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error283 error
		error283, err = error282.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error283
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "GetItemSkuJson failed: invalid message type")
		return
	}
	result := ItemServiceGetItemSkuJsonResult{}
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
//  - ItemId
//  - IType
func (p *ItemServiceClient) GetItemDetailData(itemId int64, iType int32) (r string, err error) {
	if err = p.sendGetItemDetailData(itemId, iType); err != nil {
		return
	}
	return p.recvGetItemDetailData()
}

func (p *ItemServiceClient) sendGetItemDetailData(itemId int64, iType int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("GetItemDetailData", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := ItemServiceGetItemDetailDataArgs{
		ItemId: itemId,
		IType:  iType,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ItemServiceClient) recvGetItemDetailData() (value string, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	method, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "GetItemDetailData" {
		err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "GetItemDetailData failed: wrong method name")
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "GetItemDetailData failed: out of sequence response")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error284 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error285 error
		error285, err = error284.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error285
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "GetItemDetailData failed: invalid message type")
		return
	}
	result := ItemServiceGetItemDetailDataResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	value = result.GetSuccess()
	return
}

type ItemServiceProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      ItemService
}

func (p *ItemServiceProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *ItemServiceProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *ItemServiceProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewItemServiceProcessor(handler ItemService) *ItemServiceProcessor {

	self286 := &ItemServiceProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self286.processorMap["GetSku"] = &itemServiceProcessorGetSku{handler: handler}
	self286.processorMap["GetItemSkuJson"] = &itemServiceProcessorGetItemSkuJson{handler: handler}
	self286.processorMap["GetItemDetailData"] = &itemServiceProcessorGetItemDetailData{handler: handler}
	return self286
}

func (p *ItemServiceProcessor) Process(iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x287 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x287.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return false, x287

}

type itemServiceProcessorGetSku struct {
	handler ItemService
}

func (p *itemServiceProcessorGetSku) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ItemServiceGetSkuArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GetSku", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ItemServiceGetSkuResult{}
	var retval *Sku
	var err2 error
	if retval, err2 = p.handler.GetSku(args.ItemId, args.SkuId); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetSku: "+err2.Error())
		oprot.WriteMessageBegin("GetSku", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("GetSku", thrift.REPLY, seqId); err2 != nil {
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

type itemServiceProcessorGetItemSkuJson struct {
	handler ItemService
}

func (p *itemServiceProcessorGetItemSkuJson) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ItemServiceGetItemSkuJsonArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GetItemSkuJson", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ItemServiceGetItemSkuJsonResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.GetItemSkuJson(args.ItemId); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetItemSkuJson: "+err2.Error())
		oprot.WriteMessageBegin("GetItemSkuJson", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("GetItemSkuJson", thrift.REPLY, seqId); err2 != nil {
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

type itemServiceProcessorGetItemDetailData struct {
	handler ItemService
}

func (p *itemServiceProcessorGetItemDetailData) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := ItemServiceGetItemDetailDataArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("GetItemDetailData", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := ItemServiceGetItemDetailDataResult{}
	var retval string
	var err2 error
	if retval, err2 = p.handler.GetItemDetailData(args.ItemId, args.IType); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing GetItemDetailData: "+err2.Error())
		oprot.WriteMessageBegin("GetItemDetailData", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("GetItemDetailData", thrift.REPLY, seqId); err2 != nil {
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

// Attributes:
//  - ItemId
//  - SkuId
type ItemServiceGetSkuArgs struct {
	ItemId int64 `thrift:"itemId,1" json:"itemId"`
	SkuId  int64 `thrift:"skuId,2" json:"skuId"`
}

func NewItemServiceGetSkuArgs() *ItemServiceGetSkuArgs {
	return &ItemServiceGetSkuArgs{}
}

func (p *ItemServiceGetSkuArgs) GetItemId() int64 {
	return p.ItemId
}

func (p *ItemServiceGetSkuArgs) GetSkuId() int64 {
	return p.SkuId
}
func (p *ItemServiceGetSkuArgs) Read(iprot thrift.TProtocol) error {
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
		case 2:
			if err := p.readField2(iprot); err != nil {
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

func (p *ItemServiceGetSkuArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ItemId = v
	}
	return nil
}

func (p *ItemServiceGetSkuArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.SkuId = v
	}
	return nil
}

func (p *ItemServiceGetSkuArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetSku_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
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

func (p *ItemServiceGetSkuArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("itemId", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:itemId: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.ItemId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.itemId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:itemId: ", p), err)
	}
	return err
}

func (p *ItemServiceGetSkuArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("skuId", thrift.I64, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:skuId: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.SkuId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.skuId (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:skuId: ", p), err)
	}
	return err
}

func (p *ItemServiceGetSkuArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ItemServiceGetSkuArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ItemServiceGetSkuResult struct {
	Success *Sku `thrift:"success,0" json:"success,omitempty"`
}

func NewItemServiceGetSkuResult() *ItemServiceGetSkuResult {
	return &ItemServiceGetSkuResult{}
}

var ItemServiceGetSkuResult_Success_DEFAULT *Sku

func (p *ItemServiceGetSkuResult) GetSuccess() *Sku {
	if !p.IsSetSuccess() {
		return ItemServiceGetSkuResult_Success_DEFAULT
	}
	return p.Success
}
func (p *ItemServiceGetSkuResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ItemServiceGetSkuResult) Read(iprot thrift.TProtocol) error {
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

func (p *ItemServiceGetSkuResult) readField0(iprot thrift.TProtocol) error {
	p.Success = &Sku{}
	if err := p.Success.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Success), err)
	}
	return nil
}

func (p *ItemServiceGetSkuResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetSku_result"); err != nil {
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

func (p *ItemServiceGetSkuResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *ItemServiceGetSkuResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ItemServiceGetSkuResult(%+v)", *p)
}

// Attributes:
//  - ItemId
type ItemServiceGetItemSkuJsonArgs struct {
	ItemId int64 `thrift:"itemId,1" json:"itemId"`
}

func NewItemServiceGetItemSkuJsonArgs() *ItemServiceGetItemSkuJsonArgs {
	return &ItemServiceGetItemSkuJsonArgs{}
}

func (p *ItemServiceGetItemSkuJsonArgs) GetItemId() int64 {
	return p.ItemId
}
func (p *ItemServiceGetItemSkuJsonArgs) Read(iprot thrift.TProtocol) error {
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

func (p *ItemServiceGetItemSkuJsonArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ItemId = v
	}
	return nil
}

func (p *ItemServiceGetItemSkuJsonArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetItemSkuJson_args"); err != nil {
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

func (p *ItemServiceGetItemSkuJsonArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("itemId", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:itemId: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.ItemId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.itemId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:itemId: ", p), err)
	}
	return err
}

func (p *ItemServiceGetItemSkuJsonArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ItemServiceGetItemSkuJsonArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ItemServiceGetItemSkuJsonResult struct {
	Success *string `thrift:"success,0" json:"success,omitempty"`
}

func NewItemServiceGetItemSkuJsonResult() *ItemServiceGetItemSkuJsonResult {
	return &ItemServiceGetItemSkuJsonResult{}
}

var ItemServiceGetItemSkuJsonResult_Success_DEFAULT string

func (p *ItemServiceGetItemSkuJsonResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return ItemServiceGetItemSkuJsonResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *ItemServiceGetItemSkuJsonResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ItemServiceGetItemSkuJsonResult) Read(iprot thrift.TProtocol) error {
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

func (p *ItemServiceGetItemSkuJsonResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *ItemServiceGetItemSkuJsonResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetItemSkuJson_result"); err != nil {
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

func (p *ItemServiceGetItemSkuJsonResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *ItemServiceGetItemSkuJsonResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ItemServiceGetItemSkuJsonResult(%+v)", *p)
}

// Attributes:
//  - ItemId
//  - IType
type ItemServiceGetItemDetailDataArgs struct {
	ItemId int64 `thrift:"itemId,1" json:"itemId"`
	IType  int32 `thrift:"iType,2" json:"iType"`
}

func NewItemServiceGetItemDetailDataArgs() *ItemServiceGetItemDetailDataArgs {
	return &ItemServiceGetItemDetailDataArgs{}
}

func (p *ItemServiceGetItemDetailDataArgs) GetItemId() int64 {
	return p.ItemId
}

func (p *ItemServiceGetItemDetailDataArgs) GetIType() int32 {
	return p.IType
}
func (p *ItemServiceGetItemDetailDataArgs) Read(iprot thrift.TProtocol) error {
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
		case 2:
			if err := p.readField2(iprot); err != nil {
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

func (p *ItemServiceGetItemDetailDataArgs) readField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.ItemId = v
	}
	return nil
}

func (p *ItemServiceGetItemDetailDataArgs) readField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.IType = v
	}
	return nil
}

func (p *ItemServiceGetItemDetailDataArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetItemDetailData_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
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

func (p *ItemServiceGetItemDetailDataArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("itemId", thrift.I64, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:itemId: ", p), err)
	}
	if err := oprot.WriteI64(int64(p.ItemId)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.itemId (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:itemId: ", p), err)
	}
	return err
}

func (p *ItemServiceGetItemDetailDataArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("iType", thrift.I32, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:iType: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.IType)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.iType (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:iType: ", p), err)
	}
	return err
}

func (p *ItemServiceGetItemDetailDataArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ItemServiceGetItemDetailDataArgs(%+v)", *p)
}

// Attributes:
//  - Success
type ItemServiceGetItemDetailDataResult struct {
	Success *string `thrift:"success,0" json:"success,omitempty"`
}

func NewItemServiceGetItemDetailDataResult() *ItemServiceGetItemDetailDataResult {
	return &ItemServiceGetItemDetailDataResult{}
}

var ItemServiceGetItemDetailDataResult_Success_DEFAULT string

func (p *ItemServiceGetItemDetailDataResult) GetSuccess() string {
	if !p.IsSetSuccess() {
		return ItemServiceGetItemDetailDataResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *ItemServiceGetItemDetailDataResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ItemServiceGetItemDetailDataResult) Read(iprot thrift.TProtocol) error {
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

func (p *ItemServiceGetItemDetailDataResult) readField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 0: ", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *ItemServiceGetItemDetailDataResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("GetItemDetailData_result"); err != nil {
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

func (p *ItemServiceGetItemDetailDataResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRING, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Success)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.success (0) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *ItemServiceGetItemDetailDataResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("ItemServiceGetItemDetailDataResult(%+v)", *p)
}
