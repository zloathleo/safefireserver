namespace go safefire.com.powermax.rpc
namespace cpp safefire.com.powermax.rpc


service TestService {
    void ping();
    list<string> funCall(1:i64 callTime, 2:string funCode, 3:map<string, string> paramMap);
}