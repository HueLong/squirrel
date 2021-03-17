<?php
/**
 * Created by PhpStorm.
 * User: bob
 * Date: 2020/5/27
 * Time: 17:37
 */

namespace App\Constants;


class SmsConst
{
    const TYPE_GETCODE_REGISTER     = 'register';
    const TYPE_GETCODE_LOGIN        = 'login';
    const TYPE_GETCODE_NEWTEL       = 'newtel';
    const TYPE_GETCODE_AUTH         = 'auth';
    const TYPE_GETCODE_FORGET       = 'forget';

    //注册
    const TEMPLATE_CODE_REGISTER      = 'SMS_139242697';
    //登录
    const TEMPLATE_CODE_LOGIN         = 'SMS_137667017';
    //绑定
    const TEMPLATE_CODE_BINDING       = 'SMS_139237774';
    //解绑
    const TEMPLATE_CODE_UNBINDING     = 'SMS_139242699';
    //找回密码
    const TEMPLATE_CODE_FORGET        = 'SMS_157065473';
    //审核通过
    const TEMPLATE_CODE_CHECKED       = 'SMS_157065476';
    //审核未通过
    const TEMPLATE_CODE_NOTCHECK      = 'SMS_157065480';
    //充值成功
    const TEMPLATE_CODE_RECHARGE      = 'SMS_157065487';
    //实名认证
    const TEMPLATE_CODE_CERTIFICATION = 'SMS_161380284';
    //通用验证
    const TEMPLATE_CODE_COMMON        = 'SMS_191833531';
    //订单通知
    const TEMPLATE_CODE_ORDER         = 'SMS_195871376';
    //报名通知
    const TEMPLATE_CODE_JOIN          = 'SMS_198925572';
    //上课通知
    const TEMPLATE_CODE_CLASS_START   = 'SMS_203717401';
    //报名成功通知
    const TEMPLATE_CODE_JOIN_SUCCESS  = 'SMS_211980125';
    //报名失败通知
    const TEMPLATE_CODE_JOIN_ERROR    = 'SMS_212693612';


    //发送短信类型list
    public static $TYPE_GETCODE_LIST = [
        self::TYPE_GETCODE_REGISTER,
        self::TYPE_GETCODE_LOGIN,
        self::TYPE_GETCODE_NEWTEL,
        self::TYPE_GETCODE_AUTH,
        self::TYPE_GETCODE_FORGET,
    ];
}
