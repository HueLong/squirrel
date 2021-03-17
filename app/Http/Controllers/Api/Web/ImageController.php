<?php


namespace App\Http\Controllers\Api\Web;


use App\Http\Controllers\Controller;
use App\Http\Traits\ApiValidatesRequests;
use App\Repository\Api\ImageRepository;
use Illuminate\Http\Request;

class ImageController extends Controller
{
    use ApiValidatesRequests;

    public function listImage(Request $request)
    {
        $data = app(ImageRepository::class)->listImage($request);
        return $this->apiResponse(200, '成功返回', $data);
    }

    public function saveImage(Request $request)
    {
        $rule = [
            'img_url' => 'required|string',
        ];
        $message = [
            'img_url.required' => '图片地址不能为空',
        ];
        $this->apiValidate($request, $rule, $message);
        $res = app(ImageRepository::class)->saveImage($request);
        return $this->apiResponse($res['code'], $res['msg']);
    }
}
