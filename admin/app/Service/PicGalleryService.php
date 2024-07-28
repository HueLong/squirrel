<?php

namespace App\Service;

use App\Admin\Repositories\PicGallery;

class PicGalleryService extends BaseService
{

    private $picGalleryRep;

    public function __construct()
    {
        $this->picGalleryRep = new PicGallery();
    }

    public function batchInsert($name, $desc, $picList)
    {
        $insertData = [];
        $nowTime = time();
        foreach ($picList as $v) {
            $data['name'] = $name;
            $data['desc'] = $desc;
            $data['pic_url'] = $v;
            $data['created_at'] = $nowTime;
            $insertData[] = $data;
        }
        return $this->picGalleryRep->batchInsert($insertData);
    }

}
