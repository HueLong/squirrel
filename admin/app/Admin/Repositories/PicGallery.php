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
}
