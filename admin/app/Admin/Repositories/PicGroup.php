<?php

namespace App\Admin\Repositories;

use App\Models\PicGroup as Model;
use Dcat\Admin\Repositories\EloquentRepository;

class PicGroup extends EloquentRepository
{
    /**
     * Model.
     *
     * @var string
     */
    protected $eloquentClass = Model::class;
}
