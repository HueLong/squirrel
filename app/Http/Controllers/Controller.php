<?php

namespace App\Http\Controllers;

use Illuminate\Foundation\Auth\Access\AuthorizesRequests;
use Illuminate\Foundation\Bus\DispatchesJobs;
use Illuminate\Foundation\Validation\ValidatesRequests;
use Illuminate\Routing\Controller as BaseController;

class Controller extends BaseController
{
    use AuthorizesRequests, DispatchesJobs, ValidatesRequests;

    public $res = [];

    /**
     * @param $code
     * @param $msg
     * @param array $data
     * @param int $status
     * @param array $headers
     * @param int $options
     * @return \Illuminate\Http\JsonResponse
     */
    protected function apiResponse($code, $msg, $data = [], $status = 200 , $headers = [], $options =0)
    {
        $returnJson = [
            'code'    => $code,
            'message' => $msg,
            'data'    => $data,
        ];

        return response()->json($returnJson, $status, $headers, $options);
    }

    // 发起外部网络请求
    protected function sendRequest($url, $option = array(), $header = array(), $type = 'POST', $assoc=false)
    {
        if (!empty($option) && $type == 'GET') {
            $url .= '?';
            foreach ($option as $key => $value) {
                $url .= $key.'='.$value.'&';
            }
        }
        $curl = curl_init(); // 启动一个CURL会话
        curl_setopt($curl, CURLOPT_URL, $url); // 要访问的地址
        curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, FALSE); // 对认证证书来源的检查
        curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, FALSE); // 从证书中检查SSL加密算法是否存在
        curl_setopt($curl, CURLOPT_USERAGENT, 'Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)'); // 模拟用户使用的浏览器
        if (!empty($option)) {
            $options = json_encode($option, JSON_UNESCAPED_UNICODE);
            curl_setopt($curl, CURLOPT_POSTFIELDS, $options); // Post提交的数据包
        }
        if (empty($header)) {
            $header = array(
                'Content-Type: application/json'
            );
        }
        curl_setopt($curl, CURLOPT_TIMEOUT, 30); // 设置超时限制防止死循环
        curl_setopt($curl, CURLOPT_HTTPHEADER, $header); // 设置HTTP头
        curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1); // 获取的信息以文件流的形式返回
        curl_setopt($curl, CURLOPT_CUSTOMREQUEST, $type);
        $result = curl_exec($curl); // 执行操作
        $result = json_decode($result, $assoc);
        curl_close($curl); // 关闭CURL会话
        return $result;
    }

    // 发起外部网络请求 --- 上传文件
    protected function sendFileRequest($url, $option)
    {
        $curl = curl_init(); // 启动一个CURL会话
        curl_setopt($curl, CURLOPT_URL, $url); // 要访问的地址
        curl_setopt($curl, CURLOPT_POST, 1);
        curl_setopt($curl, CURLOPT_POSTFIELDS, $option);
        curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1); // 获取的信息以文件流的形式返回
        // curl_setopt ( $curl, CURLOPT_CUSTOMREQUEST, 'PATCH' );
        $result = curl_exec($curl); // 执行操作
        $result = json_decode($result);
        curl_close($curl); // 关闭CURL会话
        return $result;
    }
    // 发起外部网络请求 --- 请求加载图片
    protected function getImageFromUrl($url, $option = array(), $header = array(), $type = 'POST')
    {
        $curl = curl_init(); // 启动一个CURL会话
        curl_setopt($curl, CURLOPT_URL, $url); // 要访问的地址
        curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, FALSE); // 对认证证书来源的检查
        curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, FALSE); // 从证书中检查SSL加密算法是否存在
        curl_setopt($curl, CURLOPT_USERAGENT, 'Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0)'); // 模拟用户使用的浏览器
        if (! empty($option)) {
            $options = json_encode($option, JSON_UNESCAPED_UNICODE);
            curl_setopt($curl, CURLOPT_POSTFIELDS, $options); // Post提交的数据包
        }
        if (empty($header)) {
            $header = array(
                'Content-Type: application/json'
            );
        }
        curl_setopt($curl, CURLOPT_TIMEOUT, 30); // 设置超时限制防止死循环
        curl_setopt($curl, CURLOPT_HTTPHEADER, $header); // 设置HTTP头
        curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1); // 获取的信息以文件流的形式返回
        curl_setopt($curl, CURLOPT_CUSTOMREQUEST, $type);
        $result = curl_exec($curl); // 执行操作
        // $result = json_decode($result);
        curl_close($curl); // 关闭CURL会话
        return $result;
    }

    /**
     * 判断是否微信内置浏览器访问
     * @return bool
     */
    protected function isWxClient()
    {
        return strpos($_SERVER['HTTP_USER_AGENT'], 'MicroMessenger') !== false;
    }

    /**
     * 是否是微信小程序
     */
    protected function isWxmin()
    {
        if (stripos($_SERVER['HTTP_USER_AGENT'], 'miniProgram') !== false) {
            return true;
        } else {
            return false;
        }
    }

    /**
     * 是否在支付宝打开
     */
    protected function isAliClient() {
        if(strpos($_SERVER['HTTP_USER_AGENT'], 'AlipayClient') !== false) {
            return true;
        }
        return false;
    }

    /**
     * 是否是移动端
     * @return bool
     */
    protected function isMobile()
    {
        // 如果有HTTP_X_WAP_PROFILE则一定是移动设备
        if (isset ($_SERVER['HTTP_X_WAP_PROFILE'])) {
            return true;
        }
        // 如果via信息含有wap则一定是移动设备,部分服务商会屏蔽该信息
        if (isset ($_SERVER['HTTP_VIA'])) {
            return stristr($_SERVER['HTTP_VIA'], "wap") ? true : false;// 找不到为flase,否则为TRUE
        }
        // 判断手机发送的客户端标志,兼容性有待提高
        if (isset ($_SERVER['HTTP_USER_AGENT'])) {
            $clientkeywords = array(
                'mobile',
                'nokia',
                'sony',
                'ericsson',
                'mot',
                'samsung',
                'htc',
                'sgh',
                'lg',
                'sharp',
                'sie-',
                'philips',
                'panasonic',
                'alcatel',
                'lenovo',
                'iphone',
                'ipod',
                'blackberry',
                'meizu',
                'android',
                'netfront',
                'symbian',
                'ucweb',
                'windowsce',
                'palm',
                'operamini',
                'operamobi',
                'openwave',
                'nexusone',
                'cldc',
                'midp',
                'wap'
            );
            // 从HTTP_USER_AGENT中查找手机浏览器的关键字
            if (preg_match("/(" . implode('|', $clientkeywords) . ")/i", strtolower($_SERVER['HTTP_USER_AGENT']))) {
                return true;
            }
        }
        if (isset ($_SERVER['HTTP_ACCEPT'])) {
            // 协议法，因为有可能不准确，放到最后判断
            // 如果只支持wml并且不支持html那一定是移动设备
            // 如果支持wml和html但是wml在html之前则是移动设备
            if ((strpos($_SERVER['HTTP_ACCEPT'], 'vnd.wap.wml') !== false) && (strpos($_SERVER['HTTP_ACCEPT'],
                        'text/html') === false || (strpos($_SERVER['HTTP_ACCEPT'],
                            'vnd.wap.wml') < strpos($_SERVER['HTTP_ACCEPT'], 'text/html')))) {
                return true;
            }
        }
        return false;
    }
}
