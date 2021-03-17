<?php


namespace App\Repository\Api;


use App\Models\Image;

class ImageRepository
{
    public function listImage($request)
    {
        return Image::query()
            ->select('id', 'title', 'desc', 'img_url')
            ->get();
    }

    public function saveImage($request)
    {
        $imgUrl = $request->input('img_url', '');
        $title = $request->input('title', '');
        $title = explode('.', $title);
        $Image['title'] = $title[0];
        $Image['img_url'] = $imgUrl;
        $res = Image::query()->insert($Image);
        if ($res) {
            return ['code' => 200, 'msg' => '添加成功'];
        } else {
            return ['code' => 0, 'msg' => '添加失败'];
        }
    }
}
