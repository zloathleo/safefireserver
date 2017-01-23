namespace go powermax.rpc
namespace cpp powermax.rpc

// double实时值
struct DoubleValue{
 1: string tag, 
 2: i64 tm,
 3: i32 status,
 4: double value,
}

// 结果
struct SimpleReturn{
 1: string code, 
 2: string msg, 
}

service RpcService {
    SimpleReturn ping(); 
    SimpleReturn sendDoubleData(1:list<DoubleValue> doubleValue); 
}