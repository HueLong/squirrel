<?php
/**
 * Created by PhpStorm.
 * User: bob
 * Date: 2020/5/27
 * Time: 16:36
 */

namespace App\Constants;

/*
 * code 码基本规则
 * 10 Debug
 * 20 Info
 * 30 Warning
 * 40-99 Error
 */
class ExceptionCodeConst
{
    /*参数验证失败*/
    const PARAMS_VALIDATE_ERROR = 10422;

    /*Ajax 错误码*/
    const AJAX_C_ERROR = 30601;

    /*从数据库查询数据不存在错误*/
    const DATA_NOT_EXIT_ERROR = 10404;
    const API_COMMON_ERROR = 10405;

    /*redis限制锁*/
    const API_REDIS_ERROR = 10408;

//    const SUCCESS = 0 ;

    /** avrokafak异常码 */
    const AVROKAFKA_ERROR = 30001;

    const INVOKE_OTHER_API_ERROR = 40001;


}
