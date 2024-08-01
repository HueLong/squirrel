<?php

namespace App\Models;

use Dcat\Admin\Traits\HasDateTimeFormatter;
use Illuminate\Database\Eloquent\SoftDeletes;
use Illuminate\Database\Eloquent\Model;

class PicGroup extends Model
{
	use HasDateTimeFormatter;
//    use SoftDeletes;

    protected $table = 'pic_group';

    public function setUpdatedAtAttribute($value)
    {
        $this->attributes['updated_at'] = $value ? strtotime($value) : 0;
    }

    public function setCreatedAtAttribute($value)
    {
        $this->attributes['created_at'] = $value ? strtotime($value) : 0;
    }

}
