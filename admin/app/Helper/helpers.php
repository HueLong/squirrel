<?php


/**
 * 获取图片地址
 * @param string $src 原地址
 * @param string $size 尺寸
 * @return string
 */
function get_img_link(string $src, string $size): string
{
    if (empty($src)) {
        return '';
    }
    $parseData = parse_url($src);
    if (empty($parseData['path'])) {
        return '';
    }
    $sizeMap = [
        'xs' => '!th_50',
        's' => '!th_100',
        'xm' => '!th_300',
        'm' => '!th_400',
        'l' => '!th_750',
        'xl' => '!th_1000',
        'o' => '!pe',
        'act' => '!th_335_220',
        'cover' => '!th_90_70',
        'front' => '!th_335_0',
        'case_list' => '!th_750_0',
        'square_280' => '!th_280',
        'prize_cover' => '!th_470_365',
        'lottery_cover' => '!th_690_400',
        'xxm' => '!th_200',
        'free' => ''
    ];
    if (!in_array($size, array_keys($sizeMap))) {
        return '';
    }
    //清除后缀
    $path = preg_replace('/![0-9a-zA-Z_!]+/', '', $parseData['path']);
    if (strpos($path, '/') !== 0) {
        $path = '/' . $path;
    }
    $sizeStr = $sizeMap[$size];
    //尺寸等参数
    if (!empty($parseData['query'])) {
        $sizeStr .= '?' . $parseData['query'];
    }
    return env('CURRENT_IMG_CND') . $path . $sizeStr;
}
