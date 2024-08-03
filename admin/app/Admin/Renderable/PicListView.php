<?php


namespace App\Admin\Renderable;


use App\Service\PicGalleryService;
use Dcat\Admin\Support\LazyRenderable;

//商品列表
class PicListView extends LazyRenderable
{

    public function render()
    {
        $groupId = $this->payload['group_id'] ?? 0;

        //查询商品名称
        $picService = new PicGalleryService();
        $picList = $picService->getPicListByGroupId($groupId);

        //渲染数据
        return view('admin.pic.pic_list', ['list' => $picList]);
    }
}
