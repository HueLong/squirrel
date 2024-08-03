<?php

namespace App\Admin\Repositories;

use App\Models\PicGallery as Model;
use Dcat\Admin\Repositories\EloquentRepository;

class PicGallery extends EloquentRepository
{
    /**
     * Model.
     *
     * @var string
     */
    protected $eloquentClass = Model::class;


    public function batchInsert(array $list): bool
    {
        return $this->model->newQuery()->insert($list);
    }

    public function getPicListByGroupId($groupId): array
    {
        $list = $this->model->newQuery()
            ->join("pic_gallery_group as gg", "gg.pic_id", "=", "pic_gallery.id")
            ->where("gg.group_id", $groupId)
            ->select("pic_gallery.id", "pic_gallery.name", "pic_gallery.pic_url")
            ->get();
        return $list ? $list->toArray() : [];
    }
}
