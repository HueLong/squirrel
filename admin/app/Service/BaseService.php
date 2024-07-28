<?php


namespace App\Service;

//service基类
class BaseService
{

    protected $uid;
    protected $error;

    /**
     * 获取错误信息
     * @return string
     */
    public function getError(): string
    {
        return $this->error ?? '';
    }

}
