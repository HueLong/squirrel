<?php

namespace App\Enum;

class PicGroupEnum
{

    //状态
    public const STATUS_ON = 1;
    public const STATUS_OFF = 0;

    public static array $status = [
        self::STATUS_ON => '开启',
        self::STATUS_OFF => '关闭'
    ];

}
