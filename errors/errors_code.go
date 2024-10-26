package errors

// 成功返回
const OK int = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码
const SERVER_COMMON_ERROR int = 100001
const REUQEST_PARAM_ERROR int = 100002
const TOKEN_EXPIRE_ERROR int = 100003
const TOKEN_GENERATE_ERROR int = 100004
const DB_ERROR int = 100005
const DB_UPDATE_AFFECTED_ZERO_ERROR int = 100006

// 用户模块
const ParamErrorCode int = 100502
